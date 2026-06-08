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
	"net/netip"
	"syscall"
	"time"
	"unsafe"

	"github.com/VKCOM/tl/internal/vkgo/pkg/algo"
)

const sysRecvmmsg = 299
const iovsSize = 100
const bufferSize = 2000

func (t *Transport) goReceive() {
	defer t.wg.Done()

	rawConn, err := t.socket.SyscallConn()
	if err != nil {
		panic(err)
	}

	iovs := make([]syscall.Iovec, iovsSize)
	rawAddrs := make([]syscall.RawSockaddrInet4, iovsSize)
	msgVec := make([]MMsghdr, iovsSize)
	localReceiveBuffers := make([][]byte, iovsSize)
	for i := 0; i < iovsSize; i++ {
		localReceiveBuffers[i] = make([]byte, bufferSize)
	}

	socketReadEnd := time.Time{}
	msgsReceived := 0
	for {
		if msgsReceived > 0 {
			t.writeMu.Lock()
			buffersTaken := min(msgsReceived, len(t.receiveBuffers))
			for i := 0; i < buffersTaken; i++ {
				n := len(t.receiveBuffers) - 1
				buffer := t.receiveBuffers[n]
				t.receiveBuffers[n] = nil
				t.receiveBuffers = t.receiveBuffers[:n]
				buffer = algo.ResizeSlice(buffer, bufferSize)
				localReceiveBuffers[i] = buffer
			}
			t.writeMu.Unlock()

			if buffersTaken < msgsReceived {
				// still not enough
				needBuffers := msgsReceived - buffersTaken
				for i := 0; i < needBuffers; i++ {
					// TODO manage uncontrollable growth of receive buffers !!!!!!!
					buffer := make([]byte, bufferSize)
					localReceiveBuffers[buffersTaken+i] = buffer
				}
			}
		}

		// TODO latency and schedule latency metrics
		for i := 0; i < len(localReceiveBuffers); i++ {
			iovs[i].Base = &localReceiveBuffers[i][0]
			iovs[i].Len = uint64(len(localReceiveBuffers[i]))

			msgVec[i].msgHdr.Name = (*byte)(unsafe.Pointer(&rawAddrs[i]))
			msgVec[i].msgHdr.Namelen = syscall.SizeofSockaddrInet4
			msgVec[i].msgHdr.Iov = &iovs[i]
			msgVec[i].msgHdr.Iovlen = 1
		}

		socketReadStart := time.Now()
		if t.debugUDPLatency && (socketReadEnd != time.Time{}) {
			t.readScheduleLatencyMetric(socketReadStart.Sub(socketReadEnd).Seconds() * 1000)
		}
		err = rawConn.Read(func(fd uintptr) (done bool) {
			var errno syscall.Errno
			received, _, errno := syscall.Syscall6(
				sysRecvmmsg,
				fd,
				uintptr(unsafe.Pointer(&msgVec[0])),
				uintptr(len(msgVec)),
				uintptr(0), // flags
				0,
				0,
			)
			if errno != 0 {
				if errno == syscall.EAGAIN {
					return false
				}
				log.Panicf("errno 0x%x", int(errno))
			}
			msgsReceived = int(received)
			return true
		})
		if t.debugUDPLatency {
			socketReadEnd = time.Now()
			t.socketReadLatencyMetric(socketReadEnd.Sub(socketReadStart).Seconds() * 1000)
			t.readMsgsCountMetric(msgsReceived)
		}
		t.stats.DatagramRead.Add(int64(msgsReceived))
		for i := 0; i < msgsReceived; i++ {
			t.stats.DatagramSizeRead.Add(int64(msgVec[i].msgLen))
		}

		if err != nil {
			if commonCloseError(err) {
				return
			}
			if t.debugUdpRPC >= 2 {
				log.Printf("[ <- goReceive ] network error: %v", err) // TODO - rare log. Happens when attaching or detaching adapters.
			}
			continue
		}

		t.writeMu.Lock()
		for i := 0; i < msgsReceived; i++ {
			if rawAddrs[i].Addr == [4]byte{0, 0, 0, 0} {
				println("zero ip (0.0.0.0) in datagram from recvmmsg")
				t.receiveBuffers = append(t.receiveBuffers, localReceiveBuffers[i])
				localReceiveBuffers[i] = nil
				continue
			}
			if rawAddrs[i].Port == 0 {
				println("zero port in datagram from recvmmsg")
				t.receiveBuffers = append(t.receiveBuffers, localReceiveBuffers[i])
				localReceiveBuffers[i] = nil
				continue
			}
			p := (*[2]byte)(unsafe.Pointer(&rawAddrs[i].Port))
			port := binary.BigEndian.Uint16((*p)[:])
			addr := netip.AddrPortFrom(netip.AddrFrom4(rawAddrs[i].Addr), port)
			t.datagramReceiveQueue.PushBack(Datagram{
				data: localReceiveBuffers[i][:msgVec[i].msgLen],
				addr: addr,
			})
			localReceiveBuffers[i] = nil
		}
		// TODO remove second mutex lock
		t.writeMu.Unlock()
		t.receiveCV.Signal()
	}
}
