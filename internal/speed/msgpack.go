package speed

import (
	"math"

	"github.com/VKCOM/tl/internal/tlcodegen/test/gen/goldmaster/tlmemcache"
	"github.com/tinylib/msgp/msgp"
)

const last7 = 0x7f
const muint8 uint8 = 0xcc
const muint16 uint8 = 0xcd
const muint32 uint8 = 0xce // const muint64 uint8 = 0xcf
const mfixstr uint8 = 0xa0
const last5 = 0x1f
const mstr8 uint8 = 0xd9
const mstr16 uint8 = 0xda
const mstr32 uint8 = 0xdb

func wfixint(u uint8) byte {
	return u & last7
}

func wfixstr(u uint8) byte {
	return (u & last5) | mfixstr
}

func appendUint32(b []byte, u uint32) []byte {
	switch {
	case u <= (1<<7)-1:
		return append(b, wfixint(uint8(u)))
	case u <= math.MaxUint8:
		return append(b, muint8, uint8(u))
	case u <= math.MaxUint16:
		return append(b, muint16, byte(u>>8), byte(u))
	default:
		return append(b, muint32, byte(u>>24), byte(u>>16), byte(u>>8), byte(u))
	}
}

func appendString(b []byte, s string) []byte {
	u := len(s)
	switch {
	case u <= 31:
		b = append(b, wfixstr(uint8(u)))
	case u <= math.MaxUint8:
		b = append(b, mstr8, uint8(u))
	case u <= math.MaxUint16:
		b = append(b, mstr16, byte(u>>8), byte(u))
	default:
		b = append(b, mstr32, byte(u>>24), byte(u>>16), byte(u>>8), byte(u))
	}
	return append(b, s...)
}

func (p *point) writeMsgp(buf []byte, writeExcessField bool) []byte {
	count := uint32(0)
	if p.x != 0 {
		count++
	}
	if p.y != 0 {
		count++
	}
	if p.z != 0 {
		count++
	}
	if writeExcessField {
		count++
	}
	buf = msgp.AppendMapHeader(buf, count)
	if p.x != 0 {
		buf = append(buf, wfixint(0))
		buf = appendUint32(buf, p.x)
	}
	if p.y != 0 {
		buf = append(buf, wfixint(1))
		buf = appendUint32(buf, p.y)
	}
	if p.z != 0 {
		buf = append(buf, wfixint(2))
		buf = appendUint32(buf, p.z)
	}
	if writeExcessField {
		buf = append(buf, wfixint(3))
		buf = appendUint32(buf, p.x)
	}
	return buf
}

func (p *point) readMsgp(buf []byte) (_ []byte, err error) {
	*p = point{}
	var field uint32
	var numFields uint32
	numFields, buf, err = msgp.ReadMapHeaderBytes(buf)
	if err != nil {
		return nil, err
	}
	for numFields > 0 {
		numFields--
		field, buf, err = msgp.ReadUint32Bytes(buf)
		if err != nil {
			return nil, err
		}
		switch field {
		case 0:
			p.x, buf, err = msgp.ReadUint32Bytes(buf)
			if err != nil {
				return nil, err
			}
		case 1:
			p.y, buf, err = msgp.ReadUint32Bytes(buf)
			if err != nil {
				return nil, err
			}
		case 2:
			p.z, buf, err = msgp.ReadUint32Bytes(buf)
			if err != nil {
				return nil, err
			}
		default:
			buf, err = msgp.Skip(buf)
			if err != nil {
				return nil, msgp.WrapError(err)
			}
		}
	}
	return buf, nil
}

func writeMsgpAdd(p *tlmemcache.Add, buf []byte, writeExcessField bool) []byte {
	count := uint32(0)
	if p.Key != "" {
		count++
	}
	if p.Flags != 0 {
		count++
	}
	if p.Delay != 0 {
		count++
	}
	if p.Value != "" {
		count++
	}
	if writeExcessField {
		count++
	}
	buf = msgp.AppendMapHeader(buf, count)
	if p.Key != "" {
		buf = appendUint32(buf, 0)
		buf = appendString(buf, p.Key)
	}
	if p.Flags != 0 {
		buf = appendUint32(buf, 1)
		buf = appendUint32(buf, uint32(p.Flags))
	}
	if p.Delay != 0 {
		buf = appendUint32(buf, 2)
		buf = appendUint32(buf, uint32(p.Delay))
	}
	if p.Value != "" {
		buf = appendUint32(buf, 3)
		buf = appendString(buf, p.Value)
	}
	if writeExcessField {
		buf = appendUint32(buf, 4)
		buf = appendUint32(buf, uint32(p.Flags))
	}
	return buf
}

func readMsgpAdds(p *tlmemcache.AddBytes, buf []byte) (_ []byte, err error) {
	p.Reset()
	var field uint32
	var str []byte
	var numFields uint32
	numFields, buf, err = msgp.ReadMapHeaderBytes(buf)
	if err != nil {
		return nil, err
	}
	for numFields > 0 {
		numFields--
		field, buf, err = msgp.ReadUint32Bytes(buf)
		if err != nil {
			return nil, err
		}
		switch field {
		case 0:
			str, buf, err = msgp.ReadStringZC(buf)
			if err != nil {
				return nil, err
			}
			p.Key = append(p.Key[:0], str...)
		case 1:
			var value uint32
			value, buf, err = msgp.ReadUint32Bytes(buf)
			if err != nil {
				return nil, err
			}
			p.Flags = int32(value)
		case 2:
			var value uint32
			value, buf, err = msgp.ReadUint32Bytes(buf)
			if err != nil {
				return nil, err
			}
			p.Delay = int32(value)
		case 3:
			str, buf, err = msgp.ReadStringZC(buf)
			if err != nil {
				return nil, err
			}
			p.Value = append(p.Value[:0], str...)
		default:
			buf, err = msgp.Skip(buf)
			if err != nil {
				return nil, msgp.WrapError(err)
			}
		}
	}
	return buf, nil
}
