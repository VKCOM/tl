package speed

import (
	"testing"

	"github.com/VKCOM/tl/internal/tlcodegen/test/gen/goldmaster/tlmemcache"
)

func BenchmarkAddReadTL(b *testing.B) {
	values, buf := prepareAddsBuffer()
	for _, v := range values {
		buf = v.WriteTL1(buf)
	}
	b.ReportAllocs()
	b.ResetTimer()
	finish := b.N / len(values)
	var res tlmemcache.AddBytes // reuse
	for i := 0; i < finish; i++ {
		buf2 := buf
		for _, v := range values {
			var err error
			buf2, err = res.ReadTL1(buf2)
			if err != nil {
				b.Fatalf("bad")
			}
			if res.Flags != v.Flags || res.Delay != v.Delay || len(res.Key) != len(v.Key) || len(res.Value) != len(v.Value) {
				b.Fatalf("bad")
			}
		}
		if len(buf2) != 0 {
			b.Fatalf("bad")
		}
	}
	printSizes(b, int64(len(buf))*int64(finish))
}

func BenchmarkAddReadTLBad(b *testing.B) {
	values, buf := prepareAddsBuffer()
	for _, v := range values {
		buf = writeTLBadAdd(buf, &v, writeExcessField)
	}
	b.ReportAllocs()
	b.ResetTimer()
	finish := b.N / len(values)
	var res tlmemcache.AddBytes // reuse
	for i := 0; i < finish; i++ {
		buf2 := buf
		for _, v := range values {
			var err error
			buf2, err = readTLBadAdd(buf2, &res)
			if err != nil {
				b.Fatalf("bad")
			}
			if res.Flags != v.Flags || res.Delay != v.Delay || len(res.Key) != len(v.Key) || len(res.Value) != len(v.Value) {
				b.Fatalf("bad")
			}
		}
		if len(buf2) != 0 {
			b.Fatalf("bad")
		}
	}
	printSizes(b, int64(len(buf))*int64(finish))
}

func BenchmarkAddReadMsgp(b *testing.B) {
	values, buf := prepareAddsBuffer()
	for _, v := range values {
		buf = writeMsgpAdd(&v, buf, writeExcessField)
	}
	b.ReportAllocs()
	b.ResetTimer()
	finish := b.N / len(values)
	var res tlmemcache.AddBytes // reuse
	for i := 0; i < finish; i++ {
		buf2 := buf
		for _, v := range values {
			var err error
			buf2, err = readMsgpAdds(&res, buf2)
			if err != nil {
				b.Fatalf("bad")
			}
			if res.Flags != v.Flags || res.Delay != v.Delay || len(res.Key) != len(v.Key) || len(res.Value) != len(v.Value) {
				b.Fatalf("bad")
			}
		}
		if len(buf2) != 0 {
			b.Fatalf("bad")
		}
	}
	printSizes(b, int64(len(buf))*int64(finish))
}

func BenchmarkAddReadMsgpFieldNames(b *testing.B) {
	values, buf := prepareAddsBuffer()
	for _, v := range values {
		buf = writeMsgpAddFieldNames(&v, buf, writeExcessField)
	}
	b.ReportAllocs()
	b.ResetTimer()
	finish := b.N / len(values)
	var res tlmemcache.AddBytes // reuse
	for i := 0; i < finish; i++ {
		buf2 := buf
		for _, v := range values {
			var err error
			buf2, err = readMsgpAddsFieldNames(&res, buf2)
			if err != nil {
				b.Fatalf("bad")
			}
			if res.Flags != v.Flags || res.Delay != v.Delay || len(res.Key) != len(v.Key) || len(res.Value) != len(v.Value) {
				b.Fatalf("bad")
			}
		}
		if len(buf2) != 0 {
			b.Fatalf("bad")
		}
	}
	printSizes(b, int64(len(buf))*int64(finish))
}

func BenchmarkAddReadMsgpArray(b *testing.B) {
	values, buf := prepareAddsBuffer()
	for _, v := range values {
		buf = writeMsgpAddArray(&v, buf, writeExcessField)
	}
	b.ReportAllocs()
	b.ResetTimer()
	finish := b.N / len(values)
	var res tlmemcache.AddBytes // reuse
	for i := 0; i < finish; i++ {
		buf2 := buf
		for _, v := range values {
			var err error
			buf2, err = readMsgpAddsArray(&res, buf2)
			if err != nil {
				b.Fatalf("bad")
			}
			if res.Flags != v.Flags || res.Delay != v.Delay || len(res.Key) != len(v.Key) || len(res.Value) != len(v.Value) {
				b.Fatalf("bad")
			}
		}
		if len(buf2) != 0 {
			b.Fatalf("bad")
		}
	}
	printSizes(b, int64(len(buf))*int64(finish))
}
