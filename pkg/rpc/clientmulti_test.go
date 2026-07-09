// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package rpc

import (
	"context"
	"errors"
	"net"
	"sync"
	"testing"

	"pgregory.net/rapid"
)

func TestDoMultiErrorError(t *testing.T) {
	t.Parallel()

	addr := NetAddr{
		Network: "tcp4",
		Address: "127.0.0.1:10000",
	}
	wrappedErr := errors.New("boom")

	tests := []struct {
		name     string
		input    DoMultiError
		expected string
	}{
		{
			name: "with wrapped error",
			input: DoMultiError{
				Addr: addr,
				Err:  wrappedErr,
				msg:  "failed to prepare request for",
			},
			expected: "failed to prepare request for tcp4://127.0.0.1:10000: boom",
		},
		{
			name: "without wrapped error",
			input: DoMultiError{
				Addr: addr,
				msg:  "failed to prepare request for",
			},
			expected: "failed to prepare request for tcp4://127.0.0.1:10000",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.input.Error(); got != tt.expected {
				t.Fatalf("unexpected Error() string: got %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestDoMultiErrorUnwrap(t *testing.T) {
	t.Parallel()

	innerErr := errors.New("inner")
	err := DoMultiError{
		Addr: NetAddr{
			Network: "tcp4",
			Address: "127.0.0.1:10000",
		},
		Err: innerErr,
		msg: "failed to handle response from",
	}

	if !errors.Is(err, innerErr) {
		t.Fatalf("expected wrapped error %q, got %q", innerErr, err)
	}
}

func TestRPCMultiRoundtrip(t *testing.T) {
	t.Parallel()

	// this is not really a property-based test, since it is not deterministic
	// however, biased integer generators from rapid are very convenient
	rapid.Check(t, testRPCMultiRoundtrip)
}

func TestDoMultiReturnsDoMultiErrorWithActorID(t *testing.T) {
	t.Parallel()

	var c ClientImpl
	prepareErr := errors.New("prepare failed")
	addr := NetAddr{
		Network: "tcp4",
		Address: "127.0.0.1:10000",
	}
	const actorID int64 = 42

	err := c.DoMulti(
		context.Background(),
		[]NetAddr{addr},
		func(_ NetAddr, req *Request) error {
			req.ActorID = actorID
			return prepareErr
		},
		func(_ NetAddr, _ *Response, _ error) error { return nil },
	)
	if err == nil {
		t.Fatal("expected error")
	}

	var doMultiErr DoMultiError
	if !errors.As(err, &doMultiErr) {
		t.Fatalf("expected DoMultiError, got %T", err)
	}
	if doMultiErr.Addr != addr {
		t.Fatalf("unexpected addr in error: got %+v, want %+v", doMultiErr.Addr, addr)
	}
	if doMultiErr.ActorID != actorID {
		t.Fatalf("unexpected actorID in error: got %d, want %d", doMultiErr.ActorID, actorID)
	}
	if !errors.Is(err, prepareErr) {
		t.Fatalf("expected wrapped prepare error %q, got %q", prepareErr, err)
	}
}

func TestDoMultiReturnsDoMultiErrorWithActorIDFromProcessResponse(t *testing.T) {
	t.Parallel()

	ln, err := net.Listen("tcp4", "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()

	s := NewServer(
		ServerWithHandler(handler),
		ServerWithCryptoKeys(testCryptoKeys),
	)
	serverErr := make(chan error, 1)
	go func() {
		serverErr <- s.Serve(ln)
	}()
	defer func() {
		if closeErr := s.Close(); closeErr != nil {
			t.Errorf("failed to close server: %v", closeErr)
		}
		if serveErr := <-serverErr; serveErr != nil {
			t.Errorf("server serve error: %v", serveErr)
		}
	}()

	c := NewClient(
		ClientWithProtocolVersion(LatestProtocolVersion),
		ClientWithCryptoKey(testCryptoKeys[0]),
	)
	defer func() {
		if closeErr := c.Close(); closeErr != nil {
			t.Errorf("failed to close client: %v", closeErr)
		}
	}()

	addr := NetAddr{
		Network: "tcp4",
		Address: ln.Addr().String(),
	}
	const actorID int64 = 777
	processErr := errors.New("process failed")

	err = c.DoMulti(
		context.Background(),
		[]NetAddr{addr},
		func(_ NetAddr, req *Request) error {
			_ = prepareTestRequest(req)
			req.ActorID = actorID
			return nil
		},
		func(_ NetAddr, _ *Response, _ error) error {
			return processErr
		},
	)
	if err == nil {
		t.Fatal("expected error")
	}

	var doMultiErr DoMultiError
	if !errors.As(err, &doMultiErr) {
		t.Fatalf("expected DoMultiError, got %T", err)
	}
	if doMultiErr.Addr != addr {
		t.Fatalf("unexpected addr in error: got %+v, want %+v", doMultiErr.Addr, addr)
	}
	if doMultiErr.ActorID != actorID {
		t.Fatalf("unexpected actorID in error: got %d, want %d", doMultiErr.ActorID, actorID)
	}
	if !errors.Is(err, processErr) {
		t.Fatalf("expected wrapped process error %q, got %q", processErr, err)
	}
}

func testRPCMultiRoundtrip(t *rapid.T) {
	ln, err := net.Listen("tcp4", "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	}

	clients := rapid.SliceOf(rapid.Custom(genClient)).Draw(t, "clients")
	numRequests := rapid.IntRange(1, 10).Draw(t, "numRequests")

	s := NewServer(
		ServerWithHandler(handler),
		ServerWithCryptoKeys(testCryptoKeys),
		ServerWithMaxConns(rapid.IntRange(0, 3).Draw(t, "maxConns")),
		ServerWithMaxWorkers(rapid.IntRange(-1, 3).Draw(t, "maxWorkers")),
		ServerWithConnReadBufSize(rapid.IntRange(0, 64).Draw(t, "connReadBufSize")),
		ServerWithConnWriteBufSize(rapid.IntRange(0, 64).Draw(t, "connWriteBufSize")),
		ServerWithRequestBufSize(rapid.IntRange(512, 1024).Draw(t, "requestBufSize")),
		ServerWithResponseBufSize(rapid.IntRange(512, 1024).Draw(t, "responseBufSize")),
	)
	serverErr := make(chan error)
	go func() {
		serverErr <- s.Serve(ln)
	}()

	var wg sync.WaitGroup
	for _, c := range clients {
		wg.Add(1)
		go func(c Client) {
			defer wg.Done()

			m := c.Multi(numRequests)
			defer m.Close()

			queryIDs := map[int64]struct{}{}
			queryIDToBodyCopy := map[int64]string{}

			for j := 0; j < numRequests; j++ {
				req := c.GetRequest()
				queryID := req.QueryID()
				bodyCopy := prepareTestRequest(req)

				err := m.Start(context.Background(), "tcp4", ln.Addr().String(), req)
				if err != nil {
					t.Errorf("failed to start request %v: %v", j, err)
				}

				queryIDs[queryID] = struct{}{}
				queryIDToBodyCopy[queryID] = bodyCopy
			}

			for k := 0; k < numRequests; k++ {
				var queryID int64
				var resp *Response
				var err error
				if k%2 == 0 {
					for qID := range queryIDs {
						queryID = qID // get the first request ID from the map
						break
					}
					resp, err = m.Wait(context.Background(), queryID)
				} else {
					queryID, resp, err = m.WaitAny(context.Background())
				}

				bodyCopy := queryIDToBodyCopy[queryID]
				delete(queryIDToBodyCopy, queryID)
				delete(queryIDs, queryID)
				checkTestResponse(t, resp, err, bodyCopy)
				c.PutResponse(resp)
			}

			err := c.Close()
			if err != nil {
				t.Errorf("failed to close client: %v", err)
			}
		}(c)
	}

	wg.Wait()

	err = s.Close()
	if err != nil {
		t.Fatal(err)
	}
	err = <-serverErr
	if err != nil {
		t.Fatal(err)
	}
}
