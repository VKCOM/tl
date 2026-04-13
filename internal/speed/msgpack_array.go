package speed

import (
	"github.com/VKCOM/tl/internal/tlcodegen/test/gen/goldmaster/tlmemcache"
	"github.com/tinylib/msgp/msgp"
)

func (p *point) writeMsgpArray(buf []byte, writeExcessField bool) []byte {
	count := uint32(3)
	if writeExcessField {
		count++
	}
	buf = msgp.AppendArrayHeader(buf, count)
	buf = appendUint32(buf, p.x)
	buf = appendUint32(buf, p.y)
	buf = appendUint32(buf, p.z)
	if writeExcessField {
		buf = appendUint32(buf, p.x)
	}
	return buf
}

func (p *point) readMsgpArray(buf []byte) (_ []byte, err error) {
	*p = point{}
	var numFields uint32
	numFields, buf, err = msgp.ReadArrayHeaderBytes(buf)
	if err != nil {
		return nil, err
	}
	if numFields == 0 {
		return buf, nil
	} else {
		numFields--
		p.x, buf, err = msgp.ReadUint32Bytes(buf)
		if err != nil {
			return nil, err
		}
	}
	if numFields == 0 {
		return buf, nil
	} else {
		numFields--
		p.y, buf, err = msgp.ReadUint32Bytes(buf)
		if err != nil {
			return nil, err
		}
	}
	if numFields == 0 {
		return buf, nil
	} else {
		numFields--
		p.z, buf, err = msgp.ReadUint32Bytes(buf)
		if err != nil {
			return nil, err
		}
	}
	for numFields > 0 {
		numFields--
		buf, err = msgp.Skip(buf)
		if err != nil {
			return nil, msgp.WrapError(err)
		}
	}
	return buf, nil
}

func writeMsgpAddArray(p *tlmemcache.Add, buf []byte, writeExcessField bool) []byte {
	count := uint32(4)
	if writeExcessField {
		count++
	}
	buf = msgp.AppendArrayHeader(buf, count)
	buf = appendString(buf, p.Key)
	buf = appendUint32(buf, uint32(p.Flags))
	buf = appendUint32(buf, uint32(p.Delay))
	buf = appendString(buf, p.Value)
	buf = appendUint32(buf, uint32(p.Flags))
	return buf
}

func readMsgpAddsArray(p *tlmemcache.AddBytes, buf []byte) (_ []byte, err error) {
	p.Reset()
	var str []byte
	var numFields uint32
	numFields, buf, err = msgp.ReadArrayHeaderBytes(buf)
	if err != nil {
		return nil, err
	}
	if numFields == 0 {
		return buf, nil
	} else {
		numFields--
		str, buf, err = msgp.ReadStringZC(buf)
		if err != nil {
			return nil, err
		}
		p.Key = append(p.Key[:0], str...)
	}
	if numFields == 0 {
		return buf, nil
	} else {
		numFields--
		var value uint32
		value, buf, err = msgp.ReadUint32Bytes(buf)
		if err != nil {
			return nil, err
		}
		p.Flags = int32(value)
	}
	if numFields == 0 {
		return buf, nil
	} else {
		var value uint32
		value, buf, err = msgp.ReadUint32Bytes(buf)
		if err != nil {
			return nil, err
		}
		p.Delay = int32(value)
	}
	if numFields == 0 {
		return buf, nil
	} else {
		numFields--
		str, buf, err = msgp.ReadStringZC(buf)
		if err != nil {
			return nil, err
		}
		p.Value = append(p.Value[:0], str...)
	}
	for numFields > 0 {
		numFields--
		buf, err = msgp.Skip(buf)
		if err != nil {
			return nil, msgp.WrapError(err)
		}
	}
	return buf, nil
}
