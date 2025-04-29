package basictl

import (
	"fmt"
	"io"
	"math"
)

func TL2ReadSize(r []byte) (int, []byte, error) {
	if len(r) == 0 {
		return 0, r, io.ErrUnexpectedEOF
	}
	b0 := r[0]

	var l int
	switch {
	case b0 <= tinyStringLen:
		l = int(b0)
		r = r[1:]
	case b0 == bigStringMarker:
		if len(r) < 4 {
			return 0, r, io.ErrUnexpectedEOF
		}
		l = (int(r[3]) << 16) + (int(r[2]) << 8) + (int(r[1]) << 0)
		r = r[4:]
		if l <= tinyStringLen {
			return 0, r, fmt.Errorf("non-canonical (big) string format for length: %d", l)
		}
	default: // hugeStringMarker
		if len(r) < 8 {
			return 0, r, io.ErrUnexpectedEOF
		}
		l64 := (int64(r[7]) << 48) + (int64(r[6]) << 40) + (int64(r[5]) << 32) + (int64(r[4]) << 24) + (int64(r[3]) << 16) + (int64(r[2]) << 8) + (int64(r[1]) << 0)
		if l64 > math.MaxInt {
			return 0, r, fmt.Errorf("string length cannot be represented on 32-bit platform: %d", l64)
		}
		l = int(l64)
		r = r[8:]
		if l <= bigStringLen {
			return 0, r, fmt.Errorf("non-canonical (huge) string format for length: %d", l)
		}
	}
	if l > 0 {
		if len(r) < l {
			return 0, r, io.ErrUnexpectedEOF
		}
	}
	return l, r, nil
}

func TL2WriteSize(w []byte, l int) []byte {
	switch {
	case l <= tinyStringLen:
		w = append(w, byte(l))
	case l <= bigStringLen:
		w = append(w, bigStringMarker, byte(l), byte(l>>8), byte(l>>16))
	default:
		if l > hugeStringLen { // for correctness only, we do not expect strings so huge
			l = hugeStringLen
		}
		w = append(w, hugeStringMarker, byte(l), byte(l>>8), byte(l>>16), byte(l>>24), byte(l>>32), byte(l>>40), byte(l>>48))
	}
	return w
}

func TL2CalculateSize(l int) int {
	switch {
	case l <= tinyStringLen:
		return 1
	case l <= bigStringLen:
		return 4
	default:
		return 8
	}
}
