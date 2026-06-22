// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package rpc

import "testing"

func TestUDPRequestBufPoolDropsLargeBuffers(t *testing.T) {
	t.Skip("flaky test, see https://teamcity.vkteam.ru/buildConfiguration/Backend_Vkgo_Tests_CheckBase/14028168?buildTab=log&linesState=624&logView=flowAware&focusLine=22473")

	const (
		requestBufSize  = 1024
		responseBufSize = 2048
	)

	s := NewServer(
		ServerWithRequestBufSize(requestBufSize),
		ServerWithResponseBufSize(responseBufSize),
	)

	pooled := s.allocateRequestBufUDP(responseBufSize)
	pooledCap := cap(*pooled)
	if pooledCap < responseBufSize {
		t.Fatalf("pooled buffer cap = %d, want at least %d", pooledCap, responseBufSize)
	}
	s.deAllocateRequestBufUDP(pooled)

	reused := s.allocateRequestBufUDP(requestBufSize)
	if cap(*reused) != pooledCap {
		t.Fatalf("reused buffer cap = %d, want %d", cap(*reused), pooledCap)
	}
	s.deAllocateRequestBufUDP(reused)

	large := s.allocateRequestBufUDP(responseBufSize + 1)
	if cap(*large) <= responseBufSize {
		t.Fatalf("large buffer cap = %d, want greater than %d", cap(*large), responseBufSize)
	}
	s.deAllocateRequestBufUDP(large)

	afterLarge := s.allocateRequestBufUDP(requestBufSize)
	if cap(*afterLarge) > responseBufSize {
		t.Fatalf("buffer pool retained large buffer cap = %d, limit %d", cap(*afterLarge), responseBufSize)
	}
}
