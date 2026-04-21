package speed

import (
	"slices"
	"testing"

	"github.com/VKCOM/tl/internal/speed/speed_proto_fast/pb_fast"
	"github.com/VKCOM/tl/internal/speed/speed_tl/basictl"
	"github.com/VKCOM/tl/internal/speed/speed_tl/tl"
	"google.golang.org/protobuf/encoding/protowire"
)

func BenchmarkPointsReadTLGen(b *testing.B) {
	values, buf := prepareTLPointssBuffer()
	for _, v := range values {
		buf = v.WriteTL1(buf)
	}
	b.ReportAllocs()
	b.ResetTimer()
	finish := b.N / len(values)
	for i := 0; i < finish; i++ {
		buf2 := buf
		for _, v := range values {
			var err error
			var p tl.Points
			buf2, err = p.ReadTL1(buf2)
			if err != nil {
				b.Fatalf("bad")
			}
			if !slices.Equal(p.Values, v.Values) {
				b.Fatalf("bad")
			}
		}
		if len(buf2) != 0 {
			b.Fatalf("bad")
		}
	}
	printSizes(b, int64(len(buf))*int64(finish))
}

func BenchmarkPointsReadTL2Gen(b *testing.B) {
	values, buf := prepareTLPointssBuffer()
	ctx := basictl.TL2WriteContext{SizeBuffer: make([]int, 1000)}
	for _, v := range values {
		buf = v.WriteTL2(buf, &ctx)
	}
	b.ReportAllocs()
	b.ResetTimer()
	finish := b.N / len(values)
	for i := 0; i < finish; i++ {
		buf2 := buf
		for _, v := range values {
			var err error
			var p tl.Points
			buf2, err = p.ReadTL2(buf2, nil)
			if err != nil {
				b.Fatalf("bad")
			}
			if !slices.Equal(p.Values, v.Values) {
				b.Fatalf("bad")
			}
		}
		if len(buf2) != 0 {
			b.Fatalf("bad")
		}
	}
	printSizes(b, int64(len(buf))*int64(finish))
}

func BenchmarkPointsReadProtoFastGen(b *testing.B) {
	values, buf := prepareProtoPointssBuffer()
	for _, v := range values {
		tmpBuf, err := v.Marshal()
		if err != nil {
			b.Fatalf("bad")
		}
		buf = protowire.AppendBytes(buf, tmpBuf)
	}
	b.ReportAllocs()
	b.ResetTimer()
	finish := b.N / len(values)
	for i := 0; i < finish; i++ {
		buf2 := buf
		for j, v := range values {
			var err error
			var p pb_fast.PointsPB
			str, n := protowire.ConsumeBytes(buf2)
			if n < 0 {
				b.Fatalf("bad")
			}
			buf2 = buf2[n:]
			err = p.Unmarshal(str)
			if err != nil {
				b.Fatalf("bad")
			}

			_ = j
			_ = v
			//if !compareProtoPoints(&p, &v) {
			//	b.Fatalf("bad %d", j)
			//}
		}
		if len(buf2) != 0 {
			b.Fatalf("bad")
		}
	}
	printSizes(b, int64(len(buf))*int64(finish))
}

func compareProtoPoints(a, b *pb_fast.PointsPB) bool {
	ap := a.Points
	bp := b.Points

	if len(ap) != len(bp) {
		return false
	}

	for i := 0; i < len(ap); i++ {
		pa := ap[i]
		pb := bp[i]

		// repeated message в Go обычно []*PointPB, поэтому учитываем nil
		if pa == nil || pb == nil {
			if pa != pb {
				return false
			}
			continue
		}

		if pa.X != pb.X || pa.Y != pb.Y || pa.Z != pb.Z {
			return false
		}
	}

	return true
}
