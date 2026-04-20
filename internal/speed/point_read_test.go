package speed

import (
	"testing"

	"github.com/VKCOM/tl/internal/speed/speed_proto/pb"
	"github.com/VKCOM/tl/internal/speed/speed_tl/tl"
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/proto"

	"github.com/VKCOM/tl/pkg/basictl"
)

func BenchmarkPointReadTL(b *testing.B) {
	values, buf := preparePointsBuffer()
	for _, v := range values {
		buf = v.writeTL(buf)
	}
	b.ReportAllocs()
	b.ResetTimer()
	finish := b.N / len(values)
	for i := 0; i < finish; i++ {
		buf2 := buf
		for _, v := range values {
			var err error
			var p point
			buf2, err = p.readTL(buf2)
			if err != nil {
				b.Fatalf("bad")
			}
			if p != v {
				b.Fatalf("bad")
			}
		}
		if len(buf2) != 0 {
			b.Fatalf("bad")
		}
	}
	printSizes(b, int64(len(buf))*int64(finish))
}

func BenchmarkPointReadTLGen(b *testing.B) {
	values, buf, _ := prepareTLPointsBuffer()
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
			var p tl.Point
			buf2, err = p.ReadTL1(buf2)
			if err != nil {
				b.Fatalf("bad")
			}
			if p != v {
				b.Fatalf("bad")
			}
		}
		if len(buf2) != 0 {
			b.Fatalf("bad")
		}
	}
	printSizes(b, int64(len(buf))*int64(finish))
}

func BenchmarkPointReadTL2(b *testing.B) {
	values, buf := preparePointsBuffer()
	for _, v := range values {
		buf = v.writeTL2(buf, writeExcessField)
	}
	b.ReportAllocs()
	b.ResetTimer()
	finish := b.N / len(values)
	for i := 0; i < finish; i++ {
		buf2 := buf
		for _, v := range values {
			var err error
			var p point
			buf2, err = p.readTL2(buf2)
			if err != nil {
				b.Fatalf("bad")
			}
			if p != v {
				b.Fatalf("bad")
			}
		}
		if len(buf2) != 0 {
			b.Fatalf("bad")
		}
	}
	printSizes(b, int64(len(buf))*int64(finish))
}

func BenchmarkPointReadTL2Gen(b *testing.B) {
	values, buf, ctx := prepareTLPointsBuffer()
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
			var p tl.Point
			buf2, err = p.ReadTL2(buf2, nil)
			if err != nil {
				b.Fatalf("bad")
			}
			if p != v {
				b.Fatalf("bad")
			}
		}
		if len(buf2) != 0 {
			b.Fatalf("bad")
		}
	}
	printSizes(b, int64(len(buf))*int64(finish))
}

func BenchmarkPointReadTLBad(b *testing.B) {
	values, buf := preparePointsBuffer()
	for _, v := range values {
		buf = v.writeTLBad(buf, writeExcessField)
	}
	b.ReportAllocs()
	b.ResetTimer()
	finish := b.N / len(values)
	for i := 0; i < finish; i++ {
		buf2 := buf
		for _, v := range values {
			var err error
			var p point
			buf2, err = p.readTLBadX(buf2)
			if err != nil {
				b.Fatalf("bad")
			}
			if p != v {
				b.Fatalf("bad")
			}
		}
		if len(buf2) != 0 {
			b.Fatalf("bad")
		}
	}
	printSizes(b, int64(len(buf))*int64(finish))
}

func BenchmarkPointReadTL2Dumb(b *testing.B) {
	values, buf := preparePointsBuffer()
	for _, v := range values {
		buf = v.writeTL2Dumb(buf)
	}
	b.ReportAllocs()
	b.ResetTimer()
	finish := b.N / len(values)
	for i := 0; i < finish; i++ {
		buf2 := buf
		for _, v := range values {
			var err error
			var p point
			buf2, err = p.readTL2(buf2)
			if err != nil {
				b.Fatalf("bad")
			}
			if p != v {
				b.Fatalf("bad")
			}
		}
		if len(buf2) != 0 {
			b.Fatalf("bad")
		}
	}
	printSizes(b, int64(len(buf))*int64(finish))
}

func BenchmarkPointReadMsgp(b *testing.B) {
	values, buf := preparePointsBuffer()
	for _, v := range values {
		buf = v.writeMsgp(buf, writeExcessField)
	}
	b.ReportAllocs()
	b.ResetTimer()
	finish := b.N / len(values)
	for i := 0; i < finish; i++ {
		buf2 := buf
		for _, v := range values {
			var err error
			var res point // reuse
			buf2, err = res.readMsgp(buf2)
			if err != nil {
				b.Fatalf("bad")
			}
			if res != v {
				b.Fatalf("bad")
			}
		}
		if len(buf2) != 0 {
			b.Fatalf("bad")
		}
	}
	printSizes(b, int64(len(buf))*int64(finish))
}

func BenchmarkPointReadMsgpFieldNames(b *testing.B) {
	values, buf := preparePointsBuffer()
	for _, v := range values {
		buf = v.writeMsgpFieldNames(buf, writeExcessField)
	}
	b.ReportAllocs()
	b.ResetTimer()
	finish := b.N / len(values)
	for i := 0; i < finish; i++ {
		buf2 := buf
		for _, v := range values {
			var err error
			var res point // reuse
			buf2, err = res.readMsgpFieldNames(buf2)
			if err != nil {
				b.Fatalf("bad")
			}
			if res != v {
				b.Fatalf("bad")
			}
		}
		if len(buf2) != 0 {
			b.Fatalf("bad")
		}
	}
	printSizes(b, int64(len(buf))*int64(finish))
}

func BenchmarkPointReadMsgpArray(b *testing.B) {
	values, buf := preparePointsBuffer()
	for _, v := range values {
		buf = v.writeMsgpArray(buf, writeExcessField)
	}
	b.ReportAllocs()
	b.ResetTimer()
	finish := b.N / len(values)
	for i := 0; i < finish; i++ {
		buf2 := buf
		for _, v := range values {
			var err error
			var res point // reuse
			buf2, err = res.readMsgpArray(buf2)
			if err != nil {
				b.Fatalf("bad")
			}
			if res != v {
				b.Fatalf("bad")
			}
		}
		if len(buf2) != 0 {
			b.Fatalf("bad")
		}
	}
	printSizes(b, int64(len(buf))*int64(finish))
}

func BenchmarkPointReadProtobuf(b *testing.B) {
	values, buf := preparePointsBuffer()
	for _, v := range values {
		buf = protobufAppendPoint(buf, &v, writeExcessField)
	}
	b.ReportAllocs()
	b.ResetTimer()
	finish := b.N / len(values)
	for i := 0; i < finish; i++ {
		buf2 := buf
		for _, v := range values {
			var err error
			var res point // reuse
			str, n := protowire.ConsumeBytes(buf2)
			if n < 0 {
				b.Fatalf("bad")
			}
			buf2 = buf2[n:]
			err = protobufUnmarshalPoint(str, &res)
			if err != nil {
				b.Fatalf("bad")
			}
			if res != v {
				b.Fatalf("bad")
			}
		}
		if len(buf2) != 0 {
			b.Fatalf("bad")
		}
	}
	printSizes(b, int64(len(buf))*int64(finish))
}

func BenchmarkPointReadProtobufGen(b *testing.B) {
	values, buf := prepareProtoPointsBuffer()
	for v := range values {
		buf = protobufAppendPoint(buf, &point{x: values[v].X, y: values[v].Y, z: values[v].Z}, writeExcessField)
	}
	b.ReportAllocs()
	b.ResetTimer()
	finish := b.N / len(values)
	for i := 0; i < finish; i++ {
		buf2 := buf
		for v := range values {
			var err error
			var res pb.PointPB // reuse
			str, n := protowire.ConsumeBytes(buf2)
			if n < 0 {
				b.Fatalf("bad")
			}
			buf2 = buf2[n:]
			err = proto.UnmarshalOptions{}.Unmarshal(str, &res)
			if err != nil {
				b.Fatalf("bad")
			}
			if res.X != values[v].X || res.Y != values[v].Y || res.Z != values[v].Z {
				b.Fatalf("bad")
			}
		}
		if len(buf2) != 0 {
			b.Fatalf("bad")
		}
	}
	printSizes(b, int64(len(buf))*int64(finish))
}

func BenchmarkPointReadProtobufOptimistic(b *testing.B) {
	values, buf := preparePointsBuffer()
	for _, v := range values {
		buf = protobufAppendPoint(buf, &v, writeExcessField)
	}
	b.ReportAllocs()
	b.ResetTimer()
	finish := b.N / len(values)
	for i := 0; i < finish; i++ {
		buf2 := buf
		for _, v := range values {
			var err error
			var res point // reuse
			str, n := protowire.ConsumeBytes(buf2)
			if n < 0 {
				b.Fatalf("bad")
			}
			buf2 = buf2[n:]
			err = protobufUnmarshalPointOptimistic(str, &res)
			if err != nil {
				b.Fatalf("bad")
			}
			if res != v {
				b.Fatalf("bad")
			}
		}
		if len(buf2) != 0 {
			b.Fatalf("bad")
		}
	}
	printSizes(b, int64(len(buf))*int64(finish))
}

func BenchmarkPointReadJSON(b *testing.B) {
	values, buf := preparePointsBuffer()
	for _, v := range values {
		buf = JSONAppendPoint(buf, &v, writeExcessField)
	}
	b.ReportAllocs()
	b.ResetTimer()
	finish := b.N / len(values)
	for i := 0; i < finish; i++ {
		in := &basictl.JsonLexer{Data: buf}
		for _, v := range values {
			var res point // reuse
			if err := JSONReadPoint(&res, in); err != nil {
				b.Fatalf("bad")
			}
			if res != v {
				b.Fatalf("bad")
			}
		}
		if in.GetPos() != len(in.Data) {
			b.Fatalf("bad")
		}
	}
	printSizes(b, int64(len(buf))*int64(finish))
}
