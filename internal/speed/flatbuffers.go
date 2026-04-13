package speed

import (
	"github.com/VKCOM/tl/internal/speed/speed"
	flatbuffers "github.com/google/flatbuffers/go"
)

func MakeFlatbuffersPoint(fb *flatbuffers.Builder, p point) []byte {
	// re-use the already-allocated Builder:
	fb.Reset()

	pp := speed.PointFlat{}
	pp.Table()

	// write the User object:
	speed.PointFlatStart(fb)
	speed.PointFlatAddX(fb, p.x)
	speed.PointFlatAddY(fb, p.y)
	speed.PointFlatAddZ(fb, p.z)

	user_position := speed.PointFlatEnd(fb)

	// finish the write operations by our User the root object:
	fb.Finish(user_position)

	// return the byte slice containing encoded data:
	return fb.Bytes[fb.Head():]
}

func ReadFlatbuffersPoint(buf []byte) (p point) {
	user := speed.GetRootAsPointFlat(buf, 0)

	return point{
		x: user.X(),
		y: user.Y(),
		z: user.Z(),
	}
}
