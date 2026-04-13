package speed

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"

	"github.com/VKCOM/tl/internal/tlcodegen/test/gen/goldmaster/tlmemcache"
	"github.com/VKCOM/tl/pkg/basictl"
)

const (
	tinyStringLen    = 253
	bigStringMarker  = 0xfe
	hugeStringMarker = 0xff
	bigStringLen     = (1 << 24) - 1
	hugeStringLen    = (1 << 56) - 1
)

func TL2StringLen(u int) int {
	switch {
	case u <= tinyStringLen:
		return 1 + u
	case u <= bigStringLen:
		return 4 + u
	default:
		return 8 + u
	}
}

func TL2AppendLen(b []byte, u int64) []byte {
	switch {
	case u <= tinyStringLen:
		return append(b, uint8(u))
	case u <= bigStringLen:
		return binary.LittleEndian.AppendUint32(b, uint32((u<<8)+bigStringMarker))
	default:
		return binary.LittleEndian.AppendUint64(b, uint64((u<<8)+hugeStringMarker))
	}
}

func TL2StringWrite(b []byte, s string) []byte {
	return append(TL2AppendLen(b, int64(len(s))), s...)
}

func TL2StringReadBytes(r []byte, dst *[]byte) ([]byte, error) {
	if len(r) == 0 {
		return r, io.ErrUnexpectedEOF
	}
	b0 := r[0]

	var l int
	switch {
	case b0 <= tinyStringLen:
		l = int(b0)
		r = r[1:]
		if l == 0 {
			*dst = (*dst)[:0]
			return r, nil
		}
	case b0 == bigStringMarker:
		if len(r) < 4 {
			return r, io.ErrUnexpectedEOF
		}
		l = int(binary.LittleEndian.Uint32(r) >> 8)
		r = r[4:]
		if l <= tinyStringLen {
			return r, fmt.Errorf("non-canonical (big) string format for length: %d", l)
		}
	default: // hugeStringMarker
		if len(r) < 8 {
			return r, io.ErrUnexpectedEOF
		}
		l64 := binary.LittleEndian.Uint64(r) >> 8
		if l64 > math.MaxInt {
			return r, fmt.Errorf("string length cannot be represented on 32-bit platform: %d", l64)
		}
		l = int(l64)
		r = r[8:]
		if l <= bigStringLen {
			return r, fmt.Errorf("non-canonical (huge) string format for length: %d", l)
		}
	}

	if len(r) < l {
		return r, io.ErrUnexpectedEOF
	}
	// *dst = append((*dst)[:0], r[:l]...)
	// Allocate only after we know there is enough bytes in reader
	if cap(*dst) < l {
		*dst = make([]byte, l)
	} else {
		*dst = (*dst)[:l]
	}
	copy(*dst, r)
	return r[l:], nil
}

func TL2StringReadBytesRange(r []byte) (_ []byte, dst []byte, err error) {
	if len(r) == 0 {
		return r, nil, io.ErrUnexpectedEOF
	}
	b0 := r[0]
	if b0 <= tinyStringLen {
		l := int(b0)
		return r[1+l:], r[1 : 1+l], nil
	}
	return tl2StringReadBytesRangeSlow(r, b0)
}

func tl2StringReadBytesRangeSlow(r []byte, b0 byte) (_ []byte, dst []byte, err error) {
	var l int
	switch {
	case b0 == bigStringMarker:
		if len(r) < 4 {
			return r, nil, io.ErrUnexpectedEOF
		}
		l = int(binary.LittleEndian.Uint32(r) >> 8)
		r = r[4:]
		if l <= tinyStringLen {
			return r, nil, fmt.Errorf("non-canonical (big) string format for length: %d", l)
		}
	default: // hugeStringMarker
		if len(r) < 8 {
			return r, nil, io.ErrUnexpectedEOF
		}
		l64 := binary.LittleEndian.Uint64(r) >> 8
		if l64 > math.MaxInt {
			return r, nil, fmt.Errorf("string length cannot be represented on 32-bit platform: %d", l64)
		}
		l = int(l64)
		r = r[8:]
		if l <= bigStringLen {
			return r, nil, fmt.Errorf("non-canonical (huge) string format for length: %d", l)
		}
	}
	return r[l:], r[:l], nil
}

/*
Function without explicit checks
func TL2StringReadBytesRangePanic(r []byte) (n int, l int) {
	b0 := r[0]

	switch {
	case b0 <= tinyStringLen:
		l = int(b0)
		n = 1
	case b0 == bigStringMarker:
		//l = int(*(*uint32)(unsafe.Pointer(&r[0]))) >> 8
		l = int(binary.LittleEndian.Uint32(r) >> 8)
		//l = (int(r[3]) << 16) + (int(r[2]) << 8) + (int(r[1]) << 0)
		n = 4
		if l <= tinyStringLen {
			panic("non-canonical (big) string format for length")
		}
	default: // hugeStringMarker
		l64 := binary.LittleEndian.Uint64(r) >> 8
		if l64 > math.MaxInt {
			panic("string length cannot be represented on 32-bit platform")
		}
		l = int(l64)
		n = 8
		if l <= bigStringLen {
			panic("non-canonical (huge) string format for length")
		}
	}
	if n+l > len(r) {
		panic("string overflow")
	}
	return n, l
}
*/

func (p *point) writeTLBad(buf []byte, writeExcessField bool) []byte {
	initial := len(buf)
	buf = append(buf, 0, 0)
	fm := 0
	if p.x != 0 {
		buf = basictl.NatWrite(buf, p.x)
		fm |= 2
	}
	if p.y != 0 {
		buf = basictl.NatWrite(buf, p.y)
		fm |= 4
	}
	if p.z != 0 {
		buf = basictl.NatWrite(buf, p.z)
		fm |= 8
	}
	if writeExcessField {
		buf = basictl.NatWrite(buf, 0)
		fm |= 16
	}
	if fm == 0 { // save byte on transmitting it
		return buf[:initial+1]
	}
	buf[initial] = byte(len(buf) - (initial + 1))
	buf[initial+1] = byte(fm)
	return buf
}

func (p *point) writeTL2(buf []byte, writeExcessField bool) []byte {
	fmPos := len(buf)
	buf = append(buf, 0)
	fm := 0
	if p.x != 0 {
		buf = basictl.NatWrite(buf, p.x)
		fm |= 1 << 1
	}
	if p.y != 0 {
		buf = basictl.NatWrite(buf, p.y)
		fm |= 1 << 4
	}
	fm2Pos := len(buf)
	buf = append(buf, 0, 0)
	fm2 := 0
	if p.z != 0 {
		buf = basictl.NatWrite(buf, p.z)
		fm2 |= 1 << 0
	}
	if writeExcessField {
		buf = basictl.NatWrite(buf, 0)
		fm2 |= 1 << 3
	}
	if fm2 == 0 {
		buf = buf[:fm2Pos]
	} else {
		fm |= 1 << 7
		buf[fm2Pos] = byte(fm2)
		buf[fm2Pos+1] = byte(fm2 >> 16)
	}
	buf[fmPos] = byte(fm)
	return buf
}

func (p *point) writeTL2Dumb(buf []byte) []byte {
	buf = append(buf, byte(1<<1)+byte(1<<4)+byte(1<<7))
	buf = basictl.NatWrite(buf, p.x)
	buf = basictl.NatWrite(buf, p.y)
	buf = append(buf, byte(1<<0), 0)
	buf = basictl.NatWrite(buf, p.z)
	return buf
}

func (p *point) readTLBad(buf []byte) (_ []byte, err error) {
	var str []byte
	if buf, str, err = TL2StringReadBytesRange(buf); err != nil {
		return buf, err
	}
	fm := 0
	if len(str) != 0 {
		fm = int(str[0])
		str = str[1:]
	}
	if fm&2 != 0 {
		if str, err = basictl.NatRead(str, &p.x); err != nil {
			return buf, err
		}
	} else {
		p.x = 0
	}
	if fm&4 != 0 {
		if str, err = basictl.NatRead(str, &p.y); err != nil {
			return buf, err
		}
	} else {
		p.y = 0
	}
	if fm&8 != 0 {
		if _, err = basictl.NatRead(str, &p.z); err != nil {
			return buf, err
		}
	} else {
		p.z = 0
	}
	return buf, nil
}

func (p *point) readTLBadX(buf []byte) (_ []byte, err error) {
	if len(buf) == 0 {
		return buf, io.ErrUnexpectedEOF
	}
	b0 := buf[0]
	if b0 > tinyStringLen { // slowpath
		return p.readTLBad(buf)
	}
	str := buf[1 : 1+int(b0)]
	buf = buf[1+int(b0):]

	fm := 0
	if len(str) != 0 {
		fm = int(str[0])
		str = str[1:]
	}
	if fm&2 != 0 {
		if str, err = basictl.NatRead(str, &p.x); err != nil {
			return buf, err
		}
	} else {
		p.x = 0
	}
	if fm&4 != 0 {
		if str, err = basictl.NatRead(str, &p.y); err != nil {
			return buf, err
		}
	} else {
		p.y = 0
	}
	if fm&8 != 0 {
		if _, err = basictl.NatRead(str, &p.z); err != nil {
			return buf, err
		}
	} else {
		p.z = 0
	}
	return buf, nil
}

func (p *point) readTL2(buf []byte) (_ []byte, err error) {
	p.Reset()
	if len(buf) == 0 {
		return buf, io.EOF
	}
	fm := buf[0]
	if fm&1 != 0 {
		return buf, io.EOF // union constructor
	}
	buf = buf[1:]
	fm >>= 1
	if t := fm & 3; t != 0 {
		if t != 1 {
			return buf, err // wrong type
		}
		if buf, err = basictl.NatRead(buf, &p.x); err != nil {
			return buf, err
		}
	}
	fm >>= 3
	if t := fm & 3; t != 0 {
		if t != 1 {
			return buf, err // wrong type
		}
		if buf, err = basictl.NatRead(buf, &p.y); err != nil {
			return buf, err
		}
	}
	fm >>= 3
	if fm == 0 {
		return buf, nil
	}
	if len(buf) < 2 {
		return buf, io.EOF
	}
	fm2 := (int(buf[1]) << 8) + int(buf[0])
	buf = buf[2:]
	if t := fm2 & 3; t != 0 {
		if t != 1 {
			return buf, err // wrong type
		}
		if buf, err = basictl.NatRead(buf, &p.z); err != nil {
			return buf, err
		}
	}
	fm2 >>= 3
	if t := fm2 & 3; t != 0 {
		if t != 1 {
			return buf, err // wrong type
		}
		var tmp uint32
		if buf, err = basictl.NatRead(buf, &tmp); err != nil {
			return buf, err
		}
	}
	fm2 >>= 3
	if fm2 != 0 {
		panic("bad")
	}
	// TODO - skip here
	return buf, nil
}

/*
Function wihout explicit checks
func (p *point) readTLBadPanic(buf []byte) (_ []byte, err error) {
	pos, length := TL2StringReadBytesRangePanic(buf)
	length += pos // length is actually a limit
	fm := 0
	if pos != length {
		fm = int(buf[pos])
		pos++
	}
	if pos != length && fm&2 != 0 {
		p.x = binary.LittleEndian.Uint32(buf[pos:])
		pos += 4
	} else {
		p.x = 0
	}
	if pos != length && fm&4 != 0 {
		p.y = binary.LittleEndian.Uint32(buf[pos:])
		pos += 4
	} else {
		p.y = 0
	}
	if pos != length && fm&8 != 0 {
		p.z = binary.LittleEndian.Uint32(buf[pos:])
		pos += 4
	} else {
		p.z = 0
	}
	return buf[length:], nil
}
*/

func writeTLBadAdd(buf []byte, p *tlmemcache.Add, writeExcessField bool) []byte {
	initial := len(buf)
	buf = append(buf, 0, 0)
	fm := 0

	if len(p.Key) != 0 {
		buf = TL2StringWrite(buf, p.Key)
		fm |= 2
	}
	if p.Flags != 0 {
		buf = basictl.IntWrite(buf, p.Flags)
		fm |= 4
	}
	if p.Delay != 0 {
		buf = basictl.IntWrite(buf, p.Delay)
		fm |= 8
	}
	if len(p.Value) != 0 {
		buf = TL2StringWrite(buf, p.Value)
		fm |= 16
	}
	if writeExcessField {
		buf = basictl.NatWrite(buf, 0)
		fm |= 32
	}
	if fm == 0 { // save byte on transmitting it
		return buf[:initial+1]
	}
	buf[initial+1] = byte(fm)
	objSize := len(buf) - (initial + 1)
	if objSize <= tinyStringLen {
		buf[initial] = byte(objSize)
		return buf
	}
	buf = append(buf, 0, 0, 0)
	copy(buf[initial+3:], buf[initial:])
	binary.LittleEndian.PutUint32(buf[initial:], uint32((objSize<<8)+bigStringMarker))
	return buf
}

func writeTLBadAddDumb(buf []byte, p *tlmemcache.Add) []byte {
	s := 1 + TL2StringLen(len(p.Key)) + 2*4 + TL2StringLen(len(p.Value))
	buf = TL2AppendLen(buf, int64(s))
	buf = append(buf, byte(2+4+8+16))
	buf = TL2StringWrite(buf, p.Key)
	buf = basictl.IntWrite(buf, p.Flags)
	buf = basictl.IntWrite(buf, p.Delay)
	buf = TL2StringWrite(buf, p.Value)
	return buf
}

func readTLBadAdd(buf []byte, p *tlmemcache.AddBytes) (_ []byte, err error) {
	var str []byte
	if buf, str, err = TL2StringReadBytesRange(buf); err != nil {
		return buf, err
	}
	fm := 0
	if len(str) != 0 {
		fm = int(str[0])
		str = str[1:]
	}
	if fm&2 != 0 {
		if str, err = TL2StringReadBytes(str, &p.Key); err != nil {
			return buf, err
		}
	} else {
		p.Key = p.Key[:0]
	}
	if fm&4 != 0 {
		if str, err = basictl.IntRead(str, &p.Flags); err != nil {
			return buf, err
		}
	} else {
		p.Flags = 0
	}
	if fm&8 != 0 {
		if str, err = basictl.IntRead(str, &p.Delay); err != nil {
			return buf, err
		}
	} else {
		p.Delay = 0
	}
	if fm&16 != 0 {
		if _, err = TL2StringReadBytes(str, &p.Value); err != nil {
			return buf, err
		}
	} else {
		p.Value = p.Value[:0]
	}
	return buf, nil
}
