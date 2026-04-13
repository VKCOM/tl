package speed

import (
	"capnproto.org/go/capnp/v3"
	"github.com/VKCOM/tl/internal/speed/speedcapnp"
)

func MakeCapnpPoint(arena capnp.Arena, p point) []byte {
	msg, seg, err := capnp.NewMessage(arena)
	if err != nil {
		panic(err)
	}
	res, err := speedcapnp.NewRootPoint(seg)
	if err != nil {
		panic(err)
	}
	res.SetX(p.x)
	res.SetY(p.y)
	res.SetZ(p.z)

	fb, err := msg.Marshal()
	if err != nil {
		panic(err)
	}
	return fb
}

func ReadCapnpPoint(buf []byte) (p point) {
	msg, err := capnp.Unmarshal(buf)
	if err != nil {
		panic(err)
	}

	// Again, don't worry about the meaning of "root" for now.
	// When in doubt, use the "root" version of functions.
	res, err := speedcapnp.ReadRootPoint(msg)
	if err != nil {
		panic(err)
	}
	return point{res.X(), res.Y(), res.Z()}
}
