package speed

import (
	"testing"

	"google.golang.org/protobuf/encoding/protowire"

	"github.com/VKCOM/tl/pkg/basictl"
)

func BenchmarkIntWriteTL(b *testing.B) {
	values, buf := prepareIntsBuffer()
	b.ReportAllocs()
	b.ResetTimer()
	var total int64
	finish := b.N / len(values)
	for i := 0; i < finish; i++ {
		for _, v := range values {
			buf = basictl.NatWrite(buf, v)
		}
		if len(buf) > bufferSize/2 {
			total += int64(len(buf))
			buf = buf[:0]
		}
	}
	printSizes(b, total+int64(len(buf)))
}

func BenchmarkIntWriteTL2(b *testing.B) {
	BenchmarkIntWriteTL(b)
}

func BenchmarkIntWriteMsgP(b *testing.B) {
	values, buf := prepareIntsBuffer()

	b.ReportAllocs()
	b.ResetTimer()
	var total int64
	finish := b.N / len(values)
	for i := 0; i < finish; i++ {
		for _, v := range values {
			buf = appendUint32(buf, v)
		}
		if len(buf) > bufferSize/2 {
			total += int64(len(buf))
			buf = buf[:0]
		}
	}
	printSizes(b, total+int64(len(buf)))
}

func BenchmarkIntWriteProtobuf(b *testing.B) {
	values, buf := prepareIntsBuffer()
	b.ReportAllocs()
	b.ResetTimer()
	var total int64
	finish := b.N / len(values)
	for i := 0; i < finish; i++ {
		for _, v := range values {
			// buf = binary.AppendUvarint(buf, uint64(v))
			buf = protowire.AppendVarint(buf, uint64(v))
		}
		if len(buf) > bufferSize/2 {
			total += int64(len(buf))
			buf = buf[:0]
		}
	}
	printSizes(b, total+int64(len(buf)))
}
