// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package udp

import (
	"testing"

	"github.com/VKCOM/tl/internal/vkgo/pkg/algo"
)

// Тест воспроизводит сброс соединения, когда один кусок сообщения уже успел
// освободить свой буфер в receive-пути, а другой кусок в том же окне ещё
// держит живую память.
// При reset нельзя освобождать память повторно и нельзя паниковать на
// двойном списании!
func TestConnectionResetGoReadUnlockedStateReleasesOnlyLiveMessageMemory(t *testing.T) {
	var deallocatorCalls int

	transport := &Transport{
		incomingMessagesMemoryLimit: 1 << 20,
		messageDeallocator: func(*[]byte) {
			deallocatorCalls++
		},
	}

	conn := &Connection{
		incoming: IncomingConnection{
			transport: transport,
			windowChunks: algo.NewTreeMap[uint32, IncomingChunk, seqNumCompT](
				&algo.SliceCacheAllocator[algo.TreeNode[algo.Entry[uint32, IncomingChunk]]]{},
			),
		},
	}
	conn.incoming.conn = conn

	message1Data := make([]byte, 1<<1)
	message1 := &IncomingMessage{ // not received message
		data:      &message1Data,
		seqNo:     111,
		parts:     1,
		remaining: 1,
	}
	message2Data := make([]byte, 1<<2)
	message2 := &IncomingMessage{ // received message
		data:      nil, // received
		seqNo:     112,
		parts:     1,
		remaining: 0,
	}
	message3Data := make([]byte, 1<<3)
	message3 := &IncomingMessage{ // one chunk received, second not yet
		data:      &message3Data,
		seqNo:     113,
		parts:     2,
		remaining: 1,
	}
	message4Data := make([]byte, 1<<4)
	message4 := &IncomingMessage{ // two chunks are received
		data:      nil, // received
		seqNo:     115,
		parts:     2,
		remaining: 0,
	}
	message5Data := make([]byte, 1<<5)
	message5 := &IncomingMessage{ // two chunks are not received
		data:      &message5Data,
		seqNo:     117,
		parts:     2,
		remaining: 2,
	}
	message6Data := make([]byte, 1<<6)
	message6 := &IncomingMessage{ // first chunk in window, second not in window
		data:      &message6Data,
		seqNo:     119,
		parts:     2,
		remaining: 2,
	}

	conn.incoming.windowChunks.Set(111, IncomingChunk{
		message:     message1,
		messageSize: 0, // not received chunk
	})
	conn.incoming.windowChunks.Set(112, IncomingChunk{
		message:     message2,
		messageSize: uint32(len(message2Data)), // received chunk
	})
	conn.incoming.windowChunks.Set(113, IncomingChunk{
		message:     message3,
		messageSize: uint32(len(message3Data)),
	})
	conn.incoming.windowChunks.Set(114, IncomingChunk{
		message:     message3,
		messageSize: 0,
	})
	conn.incoming.windowChunks.Set(115, IncomingChunk{
		message:     message4,
		messageSize: uint32(len(message4Data)),
	})
	conn.incoming.windowChunks.Set(116, IncomingChunk{
		message:     message4,
		messageSize: uint32(len(message4Data)),
	})
	conn.incoming.windowChunks.Set(117, IncomingChunk{
		message:     message5,
		messageSize: 0,
	})
	conn.incoming.windowChunks.Set(118, IncomingChunk{
		message:     message5,
		messageSize: 0,
	})
	conn.incoming.windowChunks.Set(119, IncomingChunk{
		message:     message6,
		messageSize: uint32(len(message6Data)),
	})

	conn.incoming.messagesBeginOffset = 2100
	conn.incoming.messagesTotalOffset = conn.incoming.messagesBeginOffset + int64(len(message1Data)) + int64(len(message2Data)) + int64(len(message3Data)) + int64(len(message4Data)) + int64(len(message5Data)) + int64(len(message6Data))

	transport.acquiredMemory = int64(len(message1Data)) + int64(len(message3Data)) + int64(len(message5Data)) + int64(len(message6Data))

	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("reset panicked: %v", r)
		}
	}()

	transport.checkInvariants()
	conn.checkInvariants(transport)

	conn.resetGoReadUnlockedState()

	if got := transport.acquiredMemory; got != 0 {
		t.Fatalf("acquiredMemory = %d, want 0", got)
	}
	if got := deallocatorCalls; got != 4 {
		t.Fatalf("deallocatorCalls = %d, want 4", got)
	}
	if got := len(transport.incomingMessagesPool); got != 4 {
		t.Fatalf("incomingMessagesPool len = %d, want 4", got)
	}
}
