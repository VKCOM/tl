package speed

import (
	"testing"

	"capnproto.org/go/capnp/v3"
	"google.golang.org/protobuf/proto"

	flatbuffers "github.com/google/flatbuffers/go"
)

func BenchmarkPointWriteTL(b *testing.B) {
	values, buf := preparePointsBuffer()
	b.ReportAllocs()
	b.ResetTimer()
	var total int64
	finish := b.N / len(values)
	for i := 0; i < finish; i++ {
		for _, v := range values {
			buf = v.writeTL(buf)
		}
		if len(buf) > bufferSize/2 {
			total += int64(len(buf))
			buf = buf[:0]
		}
	}
	printSizes(b, total+int64(len(buf)))
}

func BenchmarkPointWriteTLGen(b *testing.B) {
	values, buf, _ := prepareTLPointsBuffer()
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

func BenchmarkPointWriteTLBad(b *testing.B) {
	values, buf := preparePointsBuffer()
	b.ReportAllocs()
	b.ResetTimer()
	var total int64
	finish := b.N / len(values)
	for i := 0; i < finish; i++ {
		for _, v := range values {
			buf = v.writeTLBad(buf, writeExcessField)
		}
		if len(buf) > bufferSize/2 {
			total += int64(len(buf))
			buf = buf[:0]
		}
	}
	printSizes(b, total+int64(len(buf)))
}

func BenchmarkPointWriteTL2(b *testing.B) {
	values, buf := preparePointsBuffer()
	b.ReportAllocs()
	b.ResetTimer()
	var total int64
	finish := b.N / len(values)
	for i := 0; i < finish; i++ {
		for _, v := range values {
			buf = v.writeTL2(buf, writeExcessField)
		}
		if len(buf) > bufferSize/2 {
			total += int64(len(buf))
			buf = buf[:0]
		}
	}
	printSizes(b, total+int64(len(buf)))
}

func BenchmarkPointWriteTL2Dumb(b *testing.B) {
	values, buf := preparePointsBuffer()
	b.ReportAllocs()
	b.ResetTimer()
	var total int64
	finish := b.N / len(values)
	for i := 0; i < finish; i++ {
		for _, v := range values {
			buf = v.writeTL2Dumb(buf)
		}
		if len(buf) > bufferSize/2 {
			total += int64(len(buf))
			buf = buf[:0]
		}
	}
	printSizes(b, total+int64(len(buf)))
}

func BenchmarkPointWriteTL2Gen(b *testing.B) {
	values, buf, ctx := prepareTLPointsBuffer()

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

func BenchmarkPointWriteMsgp(b *testing.B) {
	values, buf := preparePointsBuffer()
	b.ReportAllocs()
	b.ResetTimer()
	var total int64
	finish := b.N / len(values)
	for i := 0; i < finish; i++ {
		for _, v := range values {
			buf = v.writeMsgp(buf, writeExcessField)
		}
		if len(buf) > bufferSize/2 {
			total += int64(len(buf))
			buf = buf[:0]
		}
	}
	printSizes(b, total+int64(len(buf)))
}

func BenchmarkPointWriteMsgpFieldNames(b *testing.B) {
	values, buf := preparePointsBuffer()
	b.ReportAllocs()
	b.ResetTimer()
	var total int64
	finish := b.N / len(values)
	for i := 0; i < finish; i++ {
		for _, v := range values {
			buf = v.writeMsgpFieldNames(buf, writeExcessField)
		}
		if len(buf) > bufferSize/2 {
			total += int64(len(buf))
			buf = buf[:0]
		}
	}
	printSizes(b, total+int64(len(buf)))
}

func BenchmarkPointWriteMsgpArray(b *testing.B) {
	values, buf := preparePointsBuffer()
	b.ReportAllocs()
	b.ResetTimer()
	var total int64
	finish := b.N / len(values)
	for i := 0; i < finish; i++ {
		for _, v := range values {
			buf = v.writeMsgpArray(buf, writeExcessField)
		}
		if len(buf) > bufferSize/2 {
			total += int64(len(buf))
			buf = buf[:0]
		}
	}
	printSizes(b, total+int64(len(buf)))
}

func BenchmarkPointWriteProtobuf(b *testing.B) {
	values, buf := preparePointsBuffer()
	b.ReportAllocs()
	b.ResetTimer()
	var total int64
	finish := b.N / len(values)
	for i := 0; i < finish; i++ {
		for _, v := range values {
			buf = protobufAppendPoint(buf, &v, writeExcessField)
		}
		if len(buf) > bufferSize/2 {
			total += int64(len(buf))
			buf = buf[:0]
		}
	}
	printSizes(b, total+int64(len(buf)))
}

func BenchmarkPointWriteProtobufGen(b *testing.B) {
	values, buf := prepareProtoPointsBuffer()
	opts := proto.MarshalOptions{}
	b.ReportAllocs()
	b.ResetTimer()
	var total int64
	finish := b.N / len(values)
	for i := 0; i < finish; i++ {
		for vIndex := range values {
			buf, _ = opts.MarshalAppend(buf, &values[vIndex])
		}
		if len(buf) > bufferSize/2 {
			total += int64(len(buf))
			buf = buf[:0]
		}
	}
	printSizes(b, total+int64(len(buf)))
}

func TestFlatbuffers(t *testing.T) {
	fb := flatbuffers.NewBuilder(0)
	p1 := point{x: 1, y: 0, z: 3}
	buf := MakeFlatbuffersPoint(fb, p1)
	p2 := ReadFlatbuffersPoint(buf)
	if p1 != p2 {
		t.Fail()
	}
}

func TestCapnproto(t *testing.T) {
	arena := capnp.SingleSegment(nil)
	p1 := point{x: 1, y: 0, z: 3}
	buf := MakeCapnpPoint(arena, p1)
	p2 := ReadCapnpPoint(buf)
	if p1 != p2 {
		t.Fail()
	}
}

func BenchmarkPointWriteFlatbuffers(b *testing.B) {
	fb := flatbuffers.NewBuilder(0)

	values, buf := preparePointsBuffer()
	b.ReportAllocs()
	b.ResetTimer()
	var total int64
	finish := b.N / len(values)
	for i := 0; i < finish; i++ {
		for _, v := range values {
			str := MakeFlatbuffersPoint(fb, v)
			total += int64(len(str))
		}
	}
	printSizes(b, total+int64(len(buf)))
}

func BenchmarkPointWriteJSON(b *testing.B) {
	values, buf := preparePointsBuffer()
	b.ReportAllocs()
	b.ResetTimer()
	var total int64
	finish := b.N / len(values)
	for i := 0; i < finish; i++ {
		for _, v := range values {
			buf = JSONAppendPoint(buf, &v, writeExcessField)
		}
		if len(buf) > bufferSize/2 {
			total += int64(len(buf))
			buf = buf[:0]
		}
	}
	printSizes(b, total+int64(len(buf)))
}

func BenchmarkPointWriteCapnp(b *testing.B) {
	arena := capnp.SingleSegment(nil)
	values, buf := preparePointsBuffer()
	b.ReportAllocs()
	b.ResetTimer()
	var total int64
	finish := b.N / len(values)
	for i := 0; i < finish; i++ {
		for _, v := range values {
			arena.Release()
			str := MakeCapnpPoint(arena, v)
			total += int64(len(str))
		}
	}
	printSizes(b, total+int64(len(buf)))
}

// n*[point]
// [Size] [Count] [Element] [Element] [Element] [Element] [Element]

// 32*[byte]
// [32] [Count] bytes
