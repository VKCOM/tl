package speed

import (
	"strings"
	"testing"

	"github.com/VKCOM/tl/internal/speed/speed_proto_fast/pb_fast"
	"github.com/VKCOM/tl/internal/speed/speed_tl/basictl"
	"github.com/VKCOM/tl/internal/speed/speed_tl/tl"
	tlmem "github.com/VKCOM/tl/internal/speed/speed_tl/tlmemcache"
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

func BenchmarkNewPoint(b *testing.B) {
	buffer_ := make([]byte, 0, bufferSize)

	benchBase[tl.Point, pb_fast.PointPB](b, benchDataGenerator[tl.Point, pb_fast.PointPB]{
		GenerateSamples: func() []tl.Point {
			ps, _, _ := prepareTLPointsBuffer()
			return ps
		},
		GenerateBuffer: func() []byte {
			return buffer_
		},
		MapSample: func(p tl.Point) pb_fast.PointPB {
			return pb_fast.PointPB{X: p.X, Y: p.Y, Z: p.Z}
		},

		TLProvider: TLData[tl.Point]{
			WriteTL1Boxed: func(x *tl.Point, buf []byte) []byte {
				return x.WriteTL1Boxed(buf)
			},
			ReadTL1Boxed: func(x *tl.Point, buf []byte) ([]byte, error) {
				return x.ReadTL1Boxed(buf)
			},
			WriteTL2: func(x *tl.Point, w []byte, tctx *basictl.TL2WriteContext) []byte {
				return x.WriteTL2(w, tctx)
			},
			ReadTL2: func(x *tl.Point, r []byte, tctx *basictl.TL2ReadContext) (_ []byte, err error) {
				return x.ReadTL2(r, tctx)
			},
		},

		ProtoProvider: ProtoData[pb_fast.PointPB]{
			Size: func(x *pb_fast.PointPB) int {
				return x.Size()
			},
			MarshalToSizedBuffer: func(x *pb_fast.PointPB, dAtA []byte) (int, error) {
				return x.MarshalToSizedBuffer(dAtA)
			},
			Marshal: func(x *pb_fast.PointPB) ([]byte, error) {
				return x.Marshal()
			},
			Unmarshal: func(x *pb_fast.PointPB, dAtA []byte) error {
				return x.Unmarshal(dAtA)
			},
		},
	})
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

func BenchmarkNewMemcacheValue(b *testing.B) {
	buffer_ := make([]byte, 0, bufferSize)

	benchBase[tlmem.Value, pb_fast.Value](b, benchDataGenerator[tlmem.Value, pb_fast.Value]{
		GenerateSamples: func() []tlmem.Value {
			ps, _ := prepareTLMemValuesBuffer()
			return ps
		},
		GenerateBuffer: func() []byte {
			return buffer_
		},
		MapSample: mapTLMemcacheValueToProto,

		TLProvider: TLData[tlmem.Value]{
			WriteTL1Boxed: func(x *tlmem.Value, buf []byte) []byte {
				return x.WriteTL1Boxed(buf)
			},
			ReadTL1Boxed: func(x *tlmem.Value, buf []byte) ([]byte, error) {
				return x.ReadTL1Boxed(buf)
			},
			WriteTL2: func(x *tlmem.Value, w []byte, tctx *basictl.TL2WriteContext) []byte {
				return x.WriteTL2(w, tctx)
			},
			ReadTL2: func(x *tlmem.Value, r []byte, tctx *basictl.TL2ReadContext) (_ []byte, err error) {
				return x.ReadTL2(r, tctx)
			},
		},

		ProtoProvider: ProtoData[pb_fast.Value]{
			Size: func(x *pb_fast.Value) int {
				return x.Size()
			},
			MarshalToSizedBuffer: func(x *pb_fast.Value, dAtA []byte) (int, error) {
				return x.MarshalToSizedBuffer(dAtA)
			},
			Marshal: func(x *pb_fast.Value) ([]byte, error) {
				return x.Marshal()
			},
			Unmarshal: func(x *pb_fast.Value, dAtA []byte) error {
				return x.Unmarshal(dAtA)
			},
		},
	})
}

func BenchmarkNewMemcacheValues(b *testing.B) {
	buffer_ := make([]byte, 0, bufferSize)

	benchBase[tlmem.Values, pb_fast.Values](b, benchDataGenerator[tlmem.Values, pb_fast.Values]{
		GenerateSamples: func() []tlmem.Values {
			ps, _ := prepareTLMemValuessBuffer()
			return ps
		},
		GenerateBuffer: func() []byte {
			return buffer_
		},
		MapSample: func(values tlmem.Values) pb_fast.Values {
			r := pb_fast.Values{Values: make([]*pb_fast.Value, len(values.Values))}

			for i, value := range values.Values {
				vv := mapTLMemcacheValueToProto(value)
				r.Values[i] = &vv
			}

			return r
		},

		TLProvider: TLData[tlmem.Values]{
			WriteTL1Boxed: func(x *tlmem.Values, buf []byte) []byte {
				return x.WriteTL1Boxed(buf)
			},
			ReadTL1Boxed: func(x *tlmem.Values, buf []byte) ([]byte, error) {
				return x.ReadTL1Boxed(buf)
			},
			WriteTL2: func(x *tlmem.Values, w []byte, tctx *basictl.TL2WriteContext) []byte {
				return x.WriteTL2(w, tctx)
			},
			ReadTL2: func(x *tlmem.Values, r []byte, tctx *basictl.TL2ReadContext) (_ []byte, err error) {
				return x.ReadTL2(r, tctx)
			},
		},

		ProtoProvider: ProtoData[pb_fast.Values]{
			Size: func(x *pb_fast.Values) int {
				return x.Size()
			},
			MarshalToSizedBuffer: func(x *pb_fast.Values, dAtA []byte) (int, error) {
				return x.MarshalToSizedBuffer(dAtA)
			},
			Marshal: func(x *pb_fast.Values) ([]byte, error) {
				return x.Marshal()
			},
			Unmarshal: func(x *pb_fast.Values, dAtA []byte) error {
				return x.Unmarshal(dAtA)
			},
		},
	})
}

func makePartialPoints(ps []tl.Point, defaultMask uint32, compact bool) []tl.PartialPoint {
	pps := make([]tl.PartialPoint, len(ps))
	for i, p := range ps {
		pp := tl.PartialPoint{}
		pp.X = p.X
		pp.Y = p.Y
		pp.Z = p.Z

		pp.Mask = uint32((i*107+31)%8) ^ defaultMask

		if compact {
			if pp.X == 0 {
				pp.Mask &^= 1 << 0
			}
			if pp.Y == 0 {
				pp.Mask &^= 1 << 1
			}
			if pp.Z == 0 {
				pp.Mask &^= 1 << 2
			}
		}

		pps[i] = pp
	}

	return pps
}

func makePartialPointPB(p tl.PartialPoint) pb_fast.PartialPointPB {
	return pb_fast.PartialPointPB{Mask: p.Mask, X: p.X, Y: p.Y, Z: p.Z}
}

func BenchmarkNewPartialPoint(b *testing.B) {
	buffer_ := make([]byte, 0, bufferSize)

	benchBase[tl.PartialPoint, pb_fast.PartialPointPB](b, benchDataGenerator[tl.PartialPoint, pb_fast.PartialPointPB]{
		GenerateSamples: func() []tl.PartialPoint {
			ps, _, _ := prepareTLPointsBuffer()
			return makePartialPoints(ps, 0, false)
		},
		GenerateBuffer: func() []byte {
			return buffer_
		},
		MapSample: makePartialPointPB,

		TLProvider: TLData[tl.PartialPoint]{
			WriteTL1Boxed: func(x *tl.PartialPoint, buf []byte) []byte {
				return x.WriteTL1Boxed(buf)
			},
			ReadTL1Boxed: func(x *tl.PartialPoint, buf []byte) ([]byte, error) {
				return x.ReadTL1Boxed(buf)
			},
			WriteTL2: func(x *tl.PartialPoint, w []byte, tctx *basictl.TL2WriteContext) []byte {
				return x.WriteTL2(w, tctx)
			},
			ReadTL2: func(x *tl.PartialPoint, r []byte, tctx *basictl.TL2ReadContext) (_ []byte, err error) {
				return x.ReadTL2(r, tctx)
			},
		},

		ProtoProvider: ProtoData[pb_fast.PartialPointPB]{
			Size: func(x *pb_fast.PartialPointPB) int {
				return x.Size()
			},
			MarshalToSizedBuffer: func(x *pb_fast.PartialPointPB, dAtA []byte) (int, error) {
				return x.MarshalToSizedBuffer(dAtA)
			},
			Marshal: func(x *pb_fast.PartialPointPB) ([]byte, error) {
				return x.Marshal()
			},
			Unmarshal: func(x *pb_fast.PartialPointPB, dAtA []byte) error {
				return x.Unmarshal(dAtA)
			},
		},
	})
}

func BenchmarkNewPartialPoints(b *testing.B) {
	buffer_ := make([]byte, 0, bufferSize)

	benchBase[tl.PartialPoints, pb_fast.PartialPointsPB](b, benchDataGenerator[tl.PartialPoints, pb_fast.PartialPointsPB]{
		GenerateSamples: func() []tl.PartialPoints {
			ps, _ := prepareTLPointssBuffer()
			pss := make([]tl.PartialPoints, len(ps))
			for i, po := range ps {
				pss[i] = tl.PartialPoints{Values: makePartialPoints(po.Values, uint32((i*73)+31)%8, false)}
			}
			return pss
		},
		GenerateBuffer: func() []byte {
			return buffer_
		},
		MapSample: func(p tl.PartialPoints) pb_fast.PartialPointsPB {
			pbb := pb_fast.PartialPointsPB{Points: make([]*pb_fast.PartialPointPB, 0)}
			for _, value := range p.Values {
				pb := makePartialPointPB(value)
				pbb.Points = append(pbb.Points, &pb)
			}
			return pbb
		},

		TLProvider: TLData[tl.PartialPoints]{
			WriteTL1Boxed: func(x *tl.PartialPoints, buf []byte) []byte {
				return x.WriteTL1Boxed(buf)
			},
			ReadTL1Boxed: func(x *tl.PartialPoints, buf []byte) ([]byte, error) {
				return x.ReadTL1Boxed(buf)
			},
			WriteTL2: func(x *tl.PartialPoints, w []byte, tctx *basictl.TL2WriteContext) []byte {
				return x.WriteTL2(w, tctx)
			},
			ReadTL2: func(x *tl.PartialPoints, r []byte, tctx *basictl.TL2ReadContext) (_ []byte, err error) {
				return x.ReadTL2(r, tctx)
			},
		},

		ProtoProvider: ProtoData[pb_fast.PartialPointsPB]{
			Size: func(x *pb_fast.PartialPointsPB) int {
				return x.Size()
			},
			MarshalToSizedBuffer: func(x *pb_fast.PartialPointsPB, dAtA []byte) (int, error) {
				return x.MarshalToSizedBuffer(dAtA)
			},
			Marshal: func(x *pb_fast.PartialPointsPB) ([]byte, error) {
				return x.Marshal()
			},
			Unmarshal: func(x *pb_fast.PartialPointsPB, dAtA []byte) error {
				return x.Unmarshal(dAtA)
			},
		},
	})
}

func BenchmarkNewPartialPointsCompact(b *testing.B) {
	buffer_ := make([]byte, 0, bufferSize)

	benchBase[tl.PartialPoints, pb_fast.PartialPointsPB](b, benchDataGenerator[tl.PartialPoints, pb_fast.PartialPointsPB]{
		GenerateSamples: func() []tl.PartialPoints {
			ps, _ := prepareTLPointssBuffer()
			pss := make([]tl.PartialPoints, len(ps))
			for i, po := range ps {
				pss[i] = tl.PartialPoints{Values: makePartialPoints(po.Values, uint32((i*73)+31)%8, true)}
			}
			return pss
		},
		GenerateBuffer: func() []byte {
			return buffer_
		},
		MapSample: func(p tl.PartialPoints) pb_fast.PartialPointsPB {
			pbb := pb_fast.PartialPointsPB{Points: make([]*pb_fast.PartialPointPB, 0)}
			for _, value := range p.Values {
				pb := makePartialPointPB(value)
				pbb.Points = append(pbb.Points, &pb)
			}
			return pbb
		},

		TLProvider: TLData[tl.PartialPoints]{
			WriteTL1Boxed: func(x *tl.PartialPoints, buf []byte) []byte {
				return x.WriteTL1Boxed(buf)
			},
			ReadTL1Boxed: func(x *tl.PartialPoints, buf []byte) ([]byte, error) {
				return x.ReadTL1Boxed(buf)
			},
			WriteTL2: func(x *tl.PartialPoints, w []byte, tctx *basictl.TL2WriteContext) []byte {
				return x.WriteTL2(w, tctx)
			},
			ReadTL2: func(x *tl.PartialPoints, r []byte, tctx *basictl.TL2ReadContext) (_ []byte, err error) {
				return x.ReadTL2(r, tctx)
			},
		},

		ProtoProvider: ProtoData[pb_fast.PartialPointsPB]{
			Size: func(x *pb_fast.PartialPointsPB) int {
				return x.Size()
			},
			MarshalToSizedBuffer: func(x *pb_fast.PartialPointsPB, dAtA []byte) (int, error) {
				return x.MarshalToSizedBuffer(dAtA)
			},
			Marshal: func(x *pb_fast.PartialPointsPB) ([]byte, error) {
				return x.Marshal()
			},
			Unmarshal: func(x *pb_fast.PartialPointsPB, dAtA []byte) error {
				return x.Unmarshal(dAtA)
			},
		},
	})
}

func makeUnionPoints(ps []tl.PartialPoint, shift int) []tl.UnionPoint {
	ups := make([]tl.UnionPoint, len(ps))
	for i, p := range ps {
		// random as i can
		rValue := (i*79 + 89) % 313
		if rValue < 313/5 {
			ups[i].SetError(tl.UnionPointError{Err: strings.Repeat("a", rValue)})
		} else {
			ups[i].SetValue(tl.UnionPointValue{
				Mask: p.Mask,
				X:    p.X,
				Y:    p.Y,
				Z:    p.Z,
			})
		}
	}

	return ups
}

func makeUnionPointPB(up tl.UnionPoint) pb_fast.UnionPointPB {
	switch {
	case up.IsValue():
		v, _ := up.AsValue()
		return pb_fast.UnionPointPB{Kind: &pb_fast.UnionPointPB_Value{Value: &pb_fast.UnionPointValuePB{
			Mask: v.Mask,
			X:    v.X,
			Y:    v.Y,
			Z:    v.Z,
		}}}
	default:
		v, _ := up.AsError()
		return pb_fast.UnionPointPB{Kind: &pb_fast.UnionPointPB_Error{Error: &pb_fast.UnionPointErrorPB{
			Err: v.Err,
		}}}
	}
}

func BenchmarkNewUnionPoint(b *testing.B) {
	buffer_ := make([]byte, 0, bufferSize)

	benchBase[tl.UnionPoint, pb_fast.UnionPointPB](b, benchDataGenerator[tl.UnionPoint, pb_fast.UnionPointPB]{
		GenerateSamples: func() []tl.UnionPoint {
			ps, _, _ := prepareTLPointsBuffer()
			return makeUnionPoints(makePartialPoints(ps, 0, false), 0)
		},
		GenerateBuffer: func() []byte {
			return buffer_
		},
		MapSample: makeUnionPointPB,

		TLProvider: TLData[tl.UnionPoint]{
			WriteTL1Boxed: func(x *tl.UnionPoint, buf []byte) []byte {
				return x.WriteTL1Boxed(buf)
			},
			ReadTL1Boxed: func(x *tl.UnionPoint, buf []byte) ([]byte, error) {
				return x.ReadTL1Boxed(buf)
			},
			WriteTL2: func(x *tl.UnionPoint, w []byte, tctx *basictl.TL2WriteContext) []byte {
				return x.WriteTL2(w, tctx)
			},
			ReadTL2: func(x *tl.UnionPoint, r []byte, tctx *basictl.TL2ReadContext) (_ []byte, err error) {
				return x.ReadTL2(r, tctx)
			},
		},

		ProtoProvider: ProtoData[pb_fast.UnionPointPB]{
			Size: func(x *pb_fast.UnionPointPB) int {
				return x.Size()
			},
			MarshalToSizedBuffer: func(x *pb_fast.UnionPointPB, dAtA []byte) (int, error) {
				return x.MarshalToSizedBuffer(dAtA)
			},
			Marshal: func(x *pb_fast.UnionPointPB) ([]byte, error) {
				return x.Marshal()
			},
			Unmarshal: func(x *pb_fast.UnionPointPB, dAtA []byte) error {
				return x.Unmarshal(dAtA)
			},
		},
	})
}

func BenchmarkNewUnionPoints(b *testing.B) {
	buffer_ := make([]byte, 0, bufferSize)

	benchBase[tl.UnionPoints, pb_fast.UnionPointsPB](b, benchDataGenerator[tl.UnionPoints, pb_fast.UnionPointsPB]{
		GenerateSamples: func() []tl.UnionPoints {
			ps, _ := prepareTLPointssBuffer()
			pss := make([]tl.UnionPoints, len(ps))
			for i, po := range ps {
				pss[i] = tl.UnionPoints{Values: makeUnionPoints(makePartialPoints(po.Values, uint32((i*73)+31)%8, true), i)}
			}
			return pss
		},
		GenerateBuffer: func() []byte {
			return buffer_
		},
		MapSample: func(p tl.UnionPoints) pb_fast.UnionPointsPB {
			pbb := pb_fast.UnionPointsPB{Points: make([]*pb_fast.UnionPointPB, 0)}
			for _, value := range p.Values {
				pb := makeUnionPointPB(value)
				pbb.Points = append(pbb.Points, &pb)
			}
			return pbb
		},

		TLProvider: TLData[tl.UnionPoints]{
			WriteTL1Boxed: func(x *tl.UnionPoints, buf []byte) []byte {
				return x.WriteTL1Boxed(buf)
			},
			ReadTL1Boxed: func(x *tl.UnionPoints, buf []byte) ([]byte, error) {
				return x.ReadTL1Boxed(buf)
			},
			WriteTL2: func(x *tl.UnionPoints, w []byte, tctx *basictl.TL2WriteContext) []byte {
				return x.WriteTL2(w, tctx)
			},
			ReadTL2: func(x *tl.UnionPoints, r []byte, tctx *basictl.TL2ReadContext) (_ []byte, err error) {
				return x.ReadTL2(r, tctx)
			},
		},

		ProtoProvider: ProtoData[pb_fast.UnionPointsPB]{
			Size: func(x *pb_fast.UnionPointsPB) int {
				return x.Size()
			},
			MarshalToSizedBuffer: func(x *pb_fast.UnionPointsPB, dAtA []byte) (int, error) {
				return x.MarshalToSizedBuffer(dAtA)
			},
			Marshal: func(x *pb_fast.UnionPointsPB) ([]byte, error) {
				return x.Marshal()
			},
			Unmarshal: func(x *pb_fast.UnionPointsPB, dAtA []byte) error {
				return x.Unmarshal(dAtA)
			},
		},
	})
}
