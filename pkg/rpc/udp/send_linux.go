// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

//go:build linux

package udp

import (
	"encoding/binary"
	"log"
	"syscall"
	"time"
	"unsafe"

	"github.com/VKCOM/tl/internal/vkgo/pkg/algo"
)

type MMsghdr struct {
	msgHdr syscall.Msghdr /* Message header */
	// TODO is it unsigned int ??
	msgLen uint32 /* Number of bytes transmitted */
}

// SysSendmmsg SENDMMSG syscall number https://filippo.io/linux-syscall-table/
const SysSendmmsg = 307

func (t *Transport) goSend() {
	defer t.wg.Done()

	rawConn, err := t.socket.SyscallConn()
	if err != nil {
		panic(err)
	}

	datagramSendQueueLocal := algo.CircularSlice[Datagram]{}
	iovs := make([]syscall.Iovec, 1)
	msgVec := make([]MMsghdr, 1)
	rawAddrs := make([]syscall.RawSockaddrInet4, 1)

	socketWriteEnd := time.Time{}
	for {
		t.writeMu.Lock()
		for t.datagramSendQueue.Len() == 0 && !t.closed {
			t.sendCV.Wait()
		}
		if t.closed {
			t.writeMu.Unlock()
			return
		}
		t.datagramSendQueue.Swap(&datagramSendQueueLocal)
		t.writeMu.Unlock()

		iovs = algo.ResizeSlice(iovs, datagramSendQueueLocal.Len())
		rawAddrs = algo.ResizeSlice(rawAddrs, datagramSendQueueLocal.Len())
		msgVec = algo.ResizeSlice(msgVec, datagramSendQueueLocal.Len())

		for i := 0; i < datagramSendQueueLocal.Len(); i++ {
			datagram := datagramSendQueueLocal.Index(i)

			// iov
			iovs[i].Base = &datagram.data[0]
			iovs[i].Len = uint64(len(datagram.data))

			// raw addr
			rawAddrs[i].Family = syscall.AF_INET
			rawAddrs[i].Addr = datagram.addr.Addr().As4()
			p := (*[2]byte)(unsafe.Pointer(&rawAddrs[i].Port))
			binary.BigEndian.PutUint16((*p)[:], datagram.addr.Port())

			// msg vec
			msgVec[i].msgHdr.Name = (*byte)(unsafe.Pointer(&rawAddrs[i]))
			msgVec[i].msgHdr.Namelen = syscall.SizeofSockaddrInet4
			msgVec[i].msgHdr.Iov = &iovs[i]
			msgVec[i].msgHdr.Iovlen = 1
		}

		msgsSent := 0
		// TODO write latency and write schedule latency metrics
		for msgsSent < datagramSendQueueLocal.Len() {
			socketWriteStart := time.Now()
			if t.debugUDPLatency && (socketWriteEnd != time.Time{}) {
				t.writeScheduleLatencyMetric(socketWriteStart.Sub(socketWriteEnd).Seconds() * 1000)
			}
			msgsSentWas := msgsSent
			err = rawConn.Write(func(fd uintptr) (done bool) {
				sent, _, errno := syscall.Syscall6(
					SysSendmmsg,
					fd,
					uintptr(unsafe.Pointer(&msgVec[msgsSent])),
					uintptr(len(msgVec)-msgsSent),
					uintptr(0), // flags
					0,
					0,
				)
				if errno != 0 {
					if errno == syscall.EAGAIN {
						return false
					}
					log.Panicf("sendmmsg returned errno 0x%x", int(errno))
					return true
				}
				msgsSent += int(sent)
				return true
			})
			if t.debugUDPLatency {
				socketWriteEnd = time.Now()
				t.socketWriteLatencyMetric(socketWriteEnd.Sub(socketWriteStart).Seconds() * 1000)
				t.writeMsgsCountMetric(msgsSent - msgsSentWas)
			}
			t.stats.DatagramWritten.Add(int64(msgsSent - msgsSentWas))
			for i := msgsSentWas; i < msgsSent; i++ {
				if msgVec[i].msgLen != uint32(len(datagramSendQueueLocal.Index(i).data)) {
					log.Printf("MMsghdr::msgLen(%d) != messageSize(%d) after sendmmsg() call", msgVec[i].msgLen, len(datagramSendQueueLocal.Index(i).data))
				}
				t.stats.DatagramSizeWritten.Add(int64(msgVec[i].msgLen))
				msgVec[i].msgLen = 0
			}

			if err != nil {
				if commonCloseError(err) {
					return
				}
				log.Printf("goWrite: failed to write datagram: %s", err)
				break
			}
		}

		// return message buffers back to goWrite() goroutine
		t.writeMu.Lock()
		for i := 0; i < datagramSendQueueLocal.Len(); i++ {
			datagramRef := datagramSendQueueLocal.IndexRef(i)
			datagramRef.data = datagramRef.data[:0]
			t.sendBuffers = append(t.sendBuffers, datagramRef.data)
			datagramRef.data = nil
		}
		t.writeMu.Unlock()
		datagramSendQueueLocal.Clear()
	}
}
