package speed

import (
	"github.com/VKCOM/tl/internal/tlcodegen/test/gen/goldmaster/tlmemcache"
	"github.com/tinylib/msgp/msgp"
)

func (p *point) writeMsgpFieldNames(buf []byte, writeExcessField bool) []byte {
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
		buf = appendString(buf, "one")
		buf = appendUint32(buf, p.x)
	}
	if p.y != 0 {
		buf = appendString(buf, "two")
		buf = appendUint32(buf, p.y)
	}
	if p.z != 0 {
		buf = appendString(buf, "three")
		buf = appendUint32(buf, p.z)
	}
	if writeExcessField {
		buf = appendString(buf, "four")
		buf = appendUint32(buf, p.x)
	}
	return buf
}

func (p *point) readMsgpFieldNames(buf []byte) (_ []byte, err error) {
	*p = point{}
	var field []byte
	var numFields uint32
	numFields, buf, err = msgp.ReadMapHeaderBytes(buf)
	if err != nil {
		return nil, err
	}
	for numFields > 0 {
		numFields--
		field, buf, err = msgp.ReadMapKeyZC(buf)
		if err != nil {
			return nil, err
		}
		switch msgp.UnsafeString(field) {
		case "one":
			p.x, buf, err = msgp.ReadUint32Bytes(buf)
			if err != nil {
				return nil, err
			}
		case "two":
			p.y, buf, err = msgp.ReadUint32Bytes(buf)
			if err != nil {
				return nil, err
			}
		case "three":
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

func writeMsgpAddFieldNames(p *tlmemcache.Add, buf []byte, writeExcessField bool) []byte {
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
		buf = appendString(buf, "key")
		buf = appendString(buf, p.Key)
	}
	if p.Flags != 0 {
		buf = appendString(buf, "flags")
		buf = appendUint32(buf, uint32(p.Flags))
	}
	if p.Delay != 0 {
		buf = appendString(buf, "delay")
		buf = appendUint32(buf, uint32(p.Delay))
	}
	if p.Value != "" {
		buf = appendString(buf, "value")
		buf = appendString(buf, p.Value)
	}
	if writeExcessField {
		buf = appendString(buf, "four")
		buf = appendUint32(buf, uint32(p.Flags))
	}
	return buf
}

func readMsgpAddsFieldNames(p *tlmemcache.AddBytes, buf []byte) (_ []byte, err error) {
	p.Reset()
	var field []byte
	var str []byte
	var numFields uint32
	numFields, buf, err = msgp.ReadMapHeaderBytes(buf)
	if err != nil {
		return nil, err
	}
	for numFields > 0 {
		numFields--
		field, buf, err = msgp.ReadMapKeyZC(buf)
		if err != nil {
			return nil, err
		}
		switch msgp.UnsafeString(field) {
		case "key":
			str, buf, err = msgp.ReadStringZC(buf)
			if err != nil {
				return nil, err
			}
			p.Key = append(p.Key[:0], str...)
		case "flags":
			var value uint32
			value, buf, err = msgp.ReadUint32Bytes(buf)
			if err != nil {
				return nil, err
			}
			p.Flags = int32(value)
		case "delay":
			var value uint32
			value, buf, err = msgp.ReadUint32Bytes(buf)
			if err != nil {
				return nil, err
			}
			p.Delay = int32(value)
		case "value":
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
