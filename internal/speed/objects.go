package speed

import (
	"strings"
	"testing"

	"github.com/VKCOM/tl/internal/speed/speed_proto/pb"
	"github.com/VKCOM/tl/internal/tlcodegen/test/gen/goldmaster/tlmemcache"
)

type point struct {
	x uint32
	y uint32
	z uint32
}

func (p *point) Reset() {
	p.x = 0
	p.y = 0
	p.z = 0
}

const chunkSize = 1000
const bufferSize = 1024 * 1024 * 128
const writeExcessField = false

func printSizes(b *testing.B, bytes int64) {
	b.SetBytes(bytes / int64(b.N)) // it is int, so will not be super correct
	b.ReportMetric(float64(bytes)/float64(b.N), "bytes")
}

func prepareIntsBuffer() ([]uint32, []byte) {
	values := []uint32{1, 5, 10, 100, 1024, 10000000, 1000, 50, 200, 500}
	for len(values) < chunkSize {
		values = append(values, values...)
	}
	values = values[:chunkSize]
	buf := make([]byte, 0, bufferSize) // worst case
	return values, buf
}

func prepareStringsBuffer() ([]string, []byte) {
	values := []string{"", "", "22", "333", "4444", "55555", "666666", "7777777", strings.Repeat("aha", 128), strings.Repeat("qqqq", 6)}
	for len(values) < chunkSize {
		values = append(values, values...)
	}
	values = values[:chunkSize]
	buf := make([]byte, 0, bufferSize) // worst case
	return values, buf
}

var basePoints = []point{{1, 0, 1}, {5, 5, 0}, {0, 10, 0}, {100, 100, 100},
	{1024, 1024, 1024}, {10000000, 10000000, 10000000}, {1000, 1000, 1000},
	{50, 50, 50}, {200, 200, 200}, {0, 0, 0}}

func preparePointsBuffer() ([]point, []byte) {
	values := make([]point, len(basePoints))
	copy(values, basePoints)

	for len(values) < chunkSize {
		values = append(values, values...)
	}
	values = values[:chunkSize]
	buf := make([]byte, 0, bufferSize) // worst case
	return values, buf
}

func prepareGeneratedPointsBuffer() ([]pb.PointPB, []byte) {
	values := make([]pb.PointPB, 0)
	for _, p := range basePoints {
		values = append(values, pb.PointPB{X: p.x, Y: p.y, Z: p.z})
	}
	for len(values) < chunkSize {
		values = append(values, values...)
	}
	values = values[:chunkSize]
	buf := make([]byte, 0, bufferSize) // worst case
	return values, buf
}

func prepareAddsBuffer() ([]tlmemcache.Add, []byte) {
	str := "some not so loooooooooooooooooooooooooooooooooooooooooooooooooooooong value"

	var values []tlmemcache.Add
	for len(values) < 10 {
		str2 := str
		if len(str2) >= len(values)*5 {
			str2 = str2[:len(values)*5]
		}
		values = append(values, tlmemcache.Add{Key: "some", Flags: 1, Delay: 100, Value: str2})
	}
	for len(values) < chunkSize {
		values = append(values, values...)
	}
	values = values[:chunkSize]
	values[0].Key = strings.Repeat("more", 256) // test of objects >256 bytes
	buf := make([]byte, 0, bufferSize)          // worst case
	return values, buf
}
