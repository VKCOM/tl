// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package basictl

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"pgregory.net/rapid"
)

var sideEffect int

func Benchmark_StringRead(b *testing.B) {
	b.ReportAllocs()
	var buf []byte
	buf = StringWrite(buf, "four")

	for i := 0; i < b.N; i++ {
		var str string
		if _, err := StringRead(buf, &str); err != nil {
			b.Fatal(err)
		}
		sideEffect += len(str)
	}
}

func Benchmark_StringRead2(b *testing.B) {
	b.ReportAllocs()
	var buf []byte
	buf = StringWrite(buf, "four")

	for i := 0; i < b.N; i++ {
		var str string
		if _, err := StringReadTL2(buf, &str); err != nil {
			b.Fatal(err)
		}
		sideEffect += len(str)
	}
}

func TestCases_String(t *testing.T) {
	src := []byte{4, 'f', 'o', 'u', 'r', 0, 0, 0}
	var str string
	r, err := StringRead(src, &str)
	if err != nil {
		t.Fatal(err)
	}
	if len(r) != 0 || str != "four" {
		t.Fatalf("wrong TL string format")
	}
	var strBytes []byte
	r, err = StringReadBytes(src, &strBytes)
	if err != nil {
		t.Fatal(err)
	}
	if len(r) != 0 || str != "four" {
		t.Fatalf("wrong TL string format")
	}
}

func TestCases_String2(t *testing.T) {
	src := []byte{4, 'f', 'o', 'u', 'r'}
	var str string
	r, err := StringReadTL2(src, &str)
	if err != nil {
		t.Fatal(err)
	}
	if len(r) != 0 || str != "four" {
		t.Fatalf("wrong TL2 string format")
	}
	var strBytes []byte
	r, err = StringReadTL2Bytes(src, &strBytes)
	if err != nil {
		t.Fatal(err)
	}
	if len(r) != 0 || str != "four" {
		t.Fatalf("wrong TL2 string format")
	}
}

func TestPanic_String(t *testing.T) {
	t.Parallel()

	rapid.Check(t, func(t *rapid.T) {
		in := rapid.String().Draw(t, "in")

		var out string
		_, _ = StringRead([]byte(in), &out)
	})
}

func TestPanic_String2(t *testing.T) {
	t.Parallel()

	rapid.Check(t, func(t *rapid.T) {
		in := rapid.String().Draw(t, "in")

		var out string
		_, _ = StringReadTL2([]byte(in), &out)
	})
}

func TestPanic_StringBytes(t *testing.T) {
	t.Parallel()

	rapid.Check(t, func(t *rapid.T) {
		in := rapid.String().Draw(t, "in")

		var out []byte
		_, _ = StringReadBytes([]byte(in), &out)
	})
}

func TestPanic_String2Bytes(t *testing.T) {
	t.Parallel()

	rapid.Check(t, func(t *rapid.T) {
		in := rapid.String().Draw(t, "in")

		var out []byte
		_, _ = StringReadTL2Bytes([]byte(in), &out)
	})
}

func TestRevStringRead(t *testing.T) {
	t.Parallel()

	rapid.Check(t, func(t *rapid.T) {
		var rw []byte
		var errW, errR error

		in := rapid.String().Draw(t, "in")
		rw, errW = StringWrite(rw, in), nil
		if errW != nil {
			t.Fatalf("failed to write %#v: %v", in, errW)
		}
		if len(rw)%4 != 0 {
			t.Fatalf("size not divisible by 4: %v", len(rw))
		}
		// if n := StringSize(len(in)); n != rw.Len() {
		//	t.Fatalf("invalid size: %v instead of %v", rw.Len(), n)
		// }

		var out string
		rw, errR = StringRead(rw, &out)
		if errR != nil {
			t.Fatalf("failed to read: %v", errR)
		}

		if in != out {
			t.Fatalf("got back %#v after writing %#v", out, in)
		}
		if len(rw) != 0 {
			t.Fatalf("%v unread bytes left", len(rw))
		}
	})
}

func TestRevStringReadBytes(t *testing.T) {
	t.Parallel()

	rapid.Check(t, func(t *rapid.T) {
		var rw []byte
		var errW, errR error

		in := rapid.SliceOf(rapid.Byte()).Draw(t, "in")
		rw, errW = StringWriteBytes(rw, in), nil
		if errW != nil {
			t.Fatalf("failed to write %#v: %v", in, errW)
		}
		if len(rw)%4 != 0 {
			t.Fatalf("size not divisible by 4: %v", len(rw))
		}

		out := rapid.SliceOf(rapid.Byte()).Draw(t, "out")
		rw, errR = StringReadBytes(rw, &out)
		if errR != nil {
			t.Fatalf("failed to read: %v", errR)
		}

		if !bytes.Equal(in, out) {
			t.Fatalf("got back %#v after writing %#v", out, in)
		}
		if len(rw) != 0 {
			t.Fatalf("%v unread bytes left", len(rw))
		}
	})
}

func TestRevStringRead2(t *testing.T) {
	t.Parallel()

	rapid.Check(t, func(t *rapid.T) {
		var rw []byte
		var errW, errR error

		in := rapid.String().Draw(t, "in")
		rw, errW = StringWriteTL2(rw, in), nil
		if errW != nil {
			t.Fatalf("failed to write %#v: %v", in, errW)
		}

		var out string
		rw, errR = StringReadTL2(rw, &out)
		if errR != nil {
			t.Fatalf("failed to read: %v", errR)
		}

		if in != out {
			t.Fatalf("got back %#v after writing %#v", out, in)
		}
		if len(rw) != 0 {
			t.Fatalf("%v unread bytes left", len(rw))
		}
	})
}

func TestRevStringRead2Bytes(t *testing.T) {
	t.Parallel()

	rapid.Check(t, func(t *rapid.T) {
		var rw []byte
		var errW, errR error

		in := rapid.SliceOf(rapid.Byte()).Draw(t, "in")
		rw, errW = StringWriteTL2Bytes(rw, in), nil
		if errW != nil {
			t.Fatalf("failed to write %#v: %v", in, errW)
		}

		out := rapid.SliceOf(rapid.Byte()).Draw(t, "out")
		rw, errR = StringReadTL2Bytes(rw, &out)
		if errR != nil {
			t.Fatalf("failed to read: %v", errR)
		}

		if !bytes.Equal(in, out) {
			t.Fatalf("got back %#v after writing %#v", out, in)
		}
		if len(rw) != 0 {
			t.Fatalf("%v unread bytes left", len(rw))
		}
	})
}

// We have no string version of this test because it works too slow.
// But we know StringRead is a wrapper around StringReadBytes, and StringWrite
func TestBuf_ByteSliceHuge(t *testing.T) {
	t.Parallel()

	prefixLen := bigStringLen - 10
	in := []byte(strings.Repeat("-", prefixLen))
	var out []byte
	var rw []byte

	rapid.Check(t, func(t *rapid.T) {
		var errW, errR error

		in = append(in[:prefixLen], rapid.SliceOf(rapid.Byte()).Draw(t, "in")...)
		rw, errW = StringWriteBytes(rw[:0], in), nil
		if errW != nil {
			t.Fatalf("failed to write %#v: %v", in, errW)
		}
		if len(rw)%4 != 0 {
			t.Fatalf("size not divisible by 4: %v", len(rw))
		}
		// if n := StringSize(len(in)); n != rw.Len() {
		//	t.Fatalf("invalid size: %v instead of %v", rw.Len(), n)
		// }

		out = append(out[:0], rapid.SliceOf(rapid.Byte()).Draw(t, "out")...)
		rw, errR = StringReadBytes(rw, &out)
		if errR != nil {
			t.Fatalf("failed to read: %v", errR)
		}

		if !bytes.Equal(in, out) {
			t.Fatalf("got back %#v after writing %#v", out, in)
		}
		if len(rw) != 0 {
			t.Fatalf("%v unread bytes left", len(rw))
		}
	})
}

func TestVectorBool(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		val := rapid.SliceOf(rapid.Bool()).Draw(t, "values")
		w := VectorBoolContentWriteTL2(nil, val)
		val2 := make([]bool, len(val))
		w2, err := VectorBoolContentReadTL2(w, val2)
		require.NoError(t, err)
		require.Equal(t, len(w2), 0)
		require.Equal(t, val, val2)
	})
}
