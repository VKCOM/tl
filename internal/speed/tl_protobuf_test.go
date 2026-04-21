package speed

import (
	"testing"

	"github.com/VKCOM/tl/internal/speed/speed_proto_fast/pb_fast"
	"github.com/VKCOM/tl/internal/speed/speed_tl/basictl"
	"github.com/VKCOM/tl/internal/speed/speed_tl/tl"
)

func BenchmarkPointsWriteGen(b *testing.B) {
	values, buf := prepareTLPointssBuffer()

	b.ReportAllocs()
	b.ResetTimer()
	var total int64
	finish := b.N / len(values)
	for i := 0; i < finish; i++ {
		for _, v := range values {
			buf = v.WriteTL1(buf)
		}
		if len(buf) > bufferSize/2 {
			total += int64(len(buf))
			buf = buf[:0]
		}
	}

	printSizes(b, total+int64(len(buf)))
}

func BenchmarkPointsWriteTL2Gen(b *testing.B) {
	values, buf := prepareTLPointssBuffer()
	ctx := basictl.TL2WriteContext{SizeBuffer: make([]int, 1000)}

	b.ReportAllocs()
	b.ResetTimer()
	var total int64
	finish := b.N / len(values)
	for i := 0; i < finish; i++ {
		for _, v := range values {
			buf = v.WriteTL2(buf, &ctx)
		}
		if len(buf) > bufferSize/2 {
			total += int64(len(buf))
			buf = buf[:0]
		}
	}

	printSizes(b, total+int64(len(buf)))
}

func BenchmarkPointsWriteProtobufFastGen(b *testing.B) {
	values, buf := prepareProtoPointssBuffer()
	b.ReportAllocs()
	b.ResetTimer()
	var total int64
	finish := b.N / len(values)
	for i := 0; i < finish; i++ {
		for vIndex := range values {
			s := values[vIndex].Size()
			pointer := len(buf)
			buf = append(buf, make([]byte, s)...)

			_, err := values[vIndex].MarshalToSizedBuffer(buf[pointer:])
			if err != nil {
				b.Fatal(err)
			}
		}
		if len(buf) > bufferSize/2 {
			total += int64(len(buf))
			buf = buf[:0]
		}
	}
	printSizes(b, total+int64(len(buf)))
}

func BenchmarkMemcacheValueWriteGen(b *testing.B) {
	values, buf := prepareTLMemValuesBuffer()

	b.ReportAllocs()
	b.ResetTimer()
	var total int64
	finish := b.N / len(values)
	for i := 0; i < finish; i++ {
		for _, v := range values {
			buf = v.WriteTL1Boxed(buf)
		}
		if len(buf) > bufferSize/2 {
			total += int64(len(buf))
			buf = buf[:0]
		}
	}

	printSizes(b, total+int64(len(buf)))
}

func BenchmarkMemcacheValueWriteTL2Gen(b *testing.B) {
	values, buf := prepareTLMemValuesBuffer()
	ctx := basictl.TL2WriteContext{SizeBuffer: make([]int, 1000)}

	b.ReportAllocs()
	b.ResetTimer()
	var total int64
	finish := b.N / len(values)
	for i := 0; i < finish; i++ {
		for _, v := range values {
			buf = v.WriteTL2(buf, &ctx)
		}
		if len(buf) > bufferSize/2 {
			total += int64(len(buf))
			buf = buf[:0]
		}
	}

	printSizes(b, total+int64(len(buf)))
}

func BenchmarkMemcacheValueWriteProtobufFastGen(b *testing.B) {
	values, buf := prepareProtoMemValuesBuffer()
	b.ReportAllocs()
	b.ResetTimer()
	var total int64
	finish := b.N / len(values)
	for i := 0; i < finish; i++ {
		for vIndex := range values {
			s := values[vIndex].Size()
			pointer := len(buf)
			buf = append(buf, make([]byte, s)...)

			_, err := values[vIndex].MarshalToSizedBuffer(buf[pointer:])
			if err != nil {
				b.Fatal(err)
			}
		}
		if len(buf) > bufferSize/2 {
			total += int64(len(buf))
			buf = buf[:0]
		}
	}
	printSizes(b, total+int64(len(buf)))
}

func BenchmarkMemcacheValuesWriteGen(b *testing.B) {
	values, buf := prepareTLMemValuessBuffer()

	b.ReportAllocs()
	b.ResetTimer()
	var total int64
	finish := b.N / len(values)
	for i := 0; i < finish; i++ {
		for _, v := range values {
			buf = v.WriteTL1(buf)
		}
		if len(buf) > bufferSize/2 {
			total += int64(len(buf))
			buf = buf[:0]
		}
	}

	printSizes(b, total+int64(len(buf)))
}

func BenchmarkMemcacheValuesWriteTL2Gen(b *testing.B) {
	values, buf := prepareTLMemValuessBuffer()
	ctx := basictl.TL2WriteContext{SizeBuffer: make([]int, 1000)}

	b.ReportAllocs()
	b.ResetTimer()
	var total int64
	finish := b.N / len(values)
	for i := 0; i < finish; i++ {
		for _, v := range values {
			buf = v.WriteTL2(buf, &ctx)
		}
		if len(buf) > bufferSize/2 {
			total += int64(len(buf))
			buf = buf[:0]
		}
	}

	printSizes(b, total+int64(len(buf)))
}

func BenchmarkMemcacheValuesWriteProtobufFastGen(b *testing.B) {
	values, buf := prepareProtoMemValuessBuffer()
	b.ReportAllocs()
	b.ResetTimer()
	var total int64
	finish := b.N / len(values)
	for i := 0; i < finish; i++ {
		for vIndex := range values {
			s := values[vIndex].Size()
			pointer := len(buf)
			buf = append(buf, make([]byte, s)...)

			_, err := values[vIndex].MarshalToSizedBuffer(buf[pointer:])
			if err != nil {
				b.Fatal(err)
			}
		}
		if len(buf) > bufferSize/2 {
			total += int64(len(buf))
			buf = buf[:0]
		}
	}
	printSizes(b, total+int64(len(buf)))
}

func BenchmarkNewPoints(b *testing.B) {
	buffer_ := make([]byte, 0, bufferSize)

	benchBase[tl.Points, pb_fast.PointsPB](b, benchDataGenerator[tl.Points, pb_fast.PointsPB]{
		GenerateSamples: func() []tl.Points {
			ps, _ := prepareTLPointssBuffer()
			return ps
		},
		GenerateBuffer: func() []byte {
			return buffer_
		},
		MapSample: func(points tl.Points) pb_fast.PointsPB {
			pbb := pb_fast.PointsPB{Points: make([]*pb_fast.PointPB, len(points.Values))}
			for i, p := range points.Values {
				pbb.Points[i] = &pb_fast.PointPB{X: p.X, Y: p.Y, Z: p.Z}
			}
			return pbb
		},

		TLProvider: TLData[tl.Points]{
			WriteTL1Boxed: func(x *tl.Points, buf []byte) []byte {
				return x.WriteTL1Boxed(buf)
			},
			ReadTL1Boxed: func(x *tl.Points, buf []byte) ([]byte, error) {
				return x.ReadTL1Boxed(buf)
			},
			WriteTL2: func(x *tl.Points, w []byte, tctx *basictl.TL2WriteContext) []byte {
				return x.WriteTL2(w, tctx)
			},
			ReadTL2: func(x *tl.Points, r []byte, tctx *basictl.TL2ReadContext) (_ []byte, err error) {
				return x.ReadTL2(r, tctx)
			},
		},

		ProtoProvider: ProtoData[pb_fast.PointsPB]{
			Size: func(x *pb_fast.PointsPB) int {
				return x.Size()
			},
			MarshalToSizedBuffer: func(x *pb_fast.PointsPB, dAtA []byte) (int, error) {
				return x.MarshalToSizedBuffer(dAtA)
			},
			Marshal: func(x *pb_fast.PointsPB) ([]byte, error) {
				return x.Marshal()
			},
			Unmarshal: func(x *pb_fast.PointsPB, dAtA []byte) error {
				return x.Unmarshal(dAtA)
			},
		},
	})
}
