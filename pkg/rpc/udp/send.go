// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

//go:build !linux

package udp

import (
	"log"
	"time"

	"github.com/VKCOM/tl/internal/vkgo/pkg/algo"
)

func (t *Transport) goSend() {
	defer t.wg.Done()

	datagramSendQueueLocal := algo.CircularSlice[Datagram]{}
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

		for datagramSendQueueLocal.Len() > 0 {
			datagram := datagramSendQueueLocal.PopFront()

			socketWriteStart := time.Now()
			if t.debugUDPLatency && (socketWriteEnd != time.Time{}) {
				t.writeScheduleLatencyMetric(socketWriteStart.Sub(socketWriteEnd).Seconds() * 1000)
			}
			_, err := t.socket.WriteToUDPAddrPort(datagram.data, datagram.addr)
			if t.debugUDPLatency {
				socketWriteEnd = time.Now()
				t.socketWriteLatencyMetric(socketWriteEnd.Sub(socketWriteStart).Seconds() * 1000)
				t.writeMsgsCountMetric(1)
			}
			t.stats.DatagramWritten.Add(1)
			t.stats.DatagramSizeWritten.Add(int64(len(datagram.data)))
			if t.debugUdpRPC >= 2 {
				log.Printf("goSend() -> datagram to %s (%d bytes)", datagram.addr.String(), int64(len(datagram.data)))
			}
			datagram.data = datagram.data[:0]
			t.writeMu.Lock()
			t.sendBuffers = append(t.sendBuffers, datagram.data)
			t.writeMu.Unlock()
			datagram.data = nil

			if err != nil {
				if commonCloseError(err) {
					return
				}
				log.Printf("goSend: failed to write datagram: %s", err)
				continue
			}
		}
	}
}
