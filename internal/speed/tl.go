package speed

import (
	"github.com/VKCOM/tl/pkg/basictl"
)

func (p *point) writeTL(buf []byte) []byte {
	buf = basictl.NatWrite(buf, p.x)
	buf = basictl.NatWrite(buf, p.y)
	return basictl.NatWrite(buf, p.z)
}

func (p *point) readTL(buf []byte) (_ []byte, err error) {
	if buf, err = basictl.NatRead(buf, &p.x); err != nil {
		return buf, err
	}
	if buf, err = basictl.NatRead(buf, &p.y); err != nil {
		return buf, err
	}
	if buf, err = basictl.NatRead(buf, &p.z); err != nil {
		return buf, err
	}
	return buf, nil
}
