package speed

import (
	"testing"
)

func BenchmarkAddWriteTL(b *testing.B) {
	values, buf := prepareAddsBuffer()
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

func BenchmarkAddWriteTLBad(b *testing.B) {
	values, buf := prepareAddsBuffer()
	b.ReportAllocs()
	b.ResetTimer()
	var total int64
	finish := b.N / len(values)
	for i := 0; i < finish; i++ {
		for _, v := range values {
			buf = writeTLBadAdd(buf, &v, writeExcessField)
		}
		if len(buf) > bufferSize/2 {
			total += int64(len(buf))
			buf = buf[:0]
		}
	}
	printSizes(b, total+int64(len(buf)))
}

func BenchmarkAddWriteTLBadDumb(b *testing.B) {
	values, buf := prepareAddsBuffer()
	b.ReportAllocs()
	b.ResetTimer()
	var total int64
	finish := b.N / len(values)
	for i := 0; i < finish; i++ {
		for _, v := range values {
			buf = writeTLBadAddDumb(buf, &v)
		}
		if len(buf) > bufferSize/2 {
			total += int64(len(buf))
			buf = buf[:0]
		}
	}
	printSizes(b, total+int64(len(buf)))
}

func BenchmarkAddWriteMsgp(b *testing.B) {
	values, buf := prepareAddsBuffer()
	b.ReportAllocs()
	b.ResetTimer()
	var total int64
	finish := b.N / len(values)
	for i := 0; i < finish; i++ {
		for _, v := range values {
			buf = writeMsgpAdd(&v, buf, writeExcessField)
		}
		if len(buf) > bufferSize/2 {
			total += int64(len(buf))
			buf = buf[:0]
		}
	}
	printSizes(b, total+int64(len(buf)))
}

func BenchmarkAddWriteMsgpFieldNames(b *testing.B) {
	values, buf := prepareAddsBuffer()
	b.ReportAllocs()
	b.ResetTimer()
	var total int64
	finish := b.N / len(values)
	for i := 0; i < finish; i++ {
		for _, v := range values {
			buf = writeMsgpAddFieldNames(&v, buf, writeExcessField)
		}
		if len(buf) > bufferSize/2 {
			total += int64(len(buf))
			buf = buf[:0]
		}
	}
	printSizes(b, total+int64(len(buf)))
}

func BenchmarkAddWriteMsgpArray(b *testing.B) {
	values, buf := prepareAddsBuffer()
	b.ReportAllocs()
	b.ResetTimer()
	var total int64
	finish := b.N / len(values)
	for i := 0; i < finish; i++ {
		for _, v := range values {
			buf = writeMsgpAddArray(&v, buf, writeExcessField)
		}
		if len(buf) > bufferSize/2 {
			total += int64(len(buf))
			buf = buf[:0]
		}
	}
	printSizes(b, total+int64(len(buf)))
}
