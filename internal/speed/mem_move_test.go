package speed

import (
	"testing"
)

func BenchmarkMemMove(b *testing.B) {
	buf := make([]byte, 256)
	b.ReportAllocs()
	b.ResetTimer()
	var bytes int64
	for i := 0; i < b.N; i++ {
		bytes += int64(copy(buf[3:], buf))
		buf[0] = byte(bytes)
		buf[1] = byte(bytes >> 8)
		buf[2] = byte(bytes >> 16)
	}
	printSizes(b, bytes)
}
