package speed

import (
	"testing"

	"github.com/VKCOM/tl/internal/speed/speed_tl/basictl"
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
