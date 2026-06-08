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

func (t *Transport) goReceive() {
	defer t.wg.Done()

	socketReadEnd := time.Time{}
	for {
		var buffer []byte
		t.writeMu.Lock()
		if n := len(t.receiveBuffers) - 1; n >= 0 {
			buffer = t.receiveBuffers[n]
			t.receiveBuffers[n] = nil
			t.receiveBuffers = t.receiveBuffers[:n]
			buffer = algo.ResizeSlice(buffer, 2000)
		} else {
			buffer = make([]byte, 2000)
		}
		t.writeMu.Unlock()

		socketReadStart := time.Now()
		if t.debugUDPLatency && (socketReadEnd != time.Time{}) {
			t.readScheduleLatencyMetric(socketReadStart.Sub(socketReadEnd).Seconds() * 1000)
		}
		// TODO process n bytes in case when err != nil and n > 0 !!!
		n, addr, err := t.socket.ReadFromUDPAddrPort(buffer)
		if t.debugUDPLatency {
			socketReadEnd = time.Now()
			t.socketReadLatencyMetric(socketReadEnd.Sub(socketReadStart).Seconds() * 1000)
			t.readMsgsCountMetric(1)
		}
		t.stats.DatagramRead.Add(1)
		t.stats.DatagramSizeRead.Add(int64(n))
		if t.debugUdpRPC >= 2 {
			log.Printf("goReceive() <- udp datagram from %s (%d bytes)", addr.String(), n)
		}

		t.writeMu.Lock()
		t.datagramReceiveQueue.PushBack(Datagram{
			data: buffer[:n],
			addr: addr,
		})
		// TODO remove second mutex lock
		t.writeMu.Unlock()
		t.receiveCV.Signal()

		if err != nil {
			if commonCloseError(err) {
				return
			}
			if t.debugUdpRPC >= 2 {
				log.Printf("[ <- goReceive ] network error: %v", err) // TODO - rare log. Happens when attaching or detaching adapters.
			}
			continue
		}
	}
}
