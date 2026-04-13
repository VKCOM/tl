package speed

import (
	"testing"

	"google.golang.org/protobuf/encoding/protowire"

	"github.com/VKCOM/tl/pkg/basictl"
	"github.com/tinylib/msgp/msgp"
)

func BenchmarkStringReadTL1(b *testing.B) {
	values, buf := prepareStringsBuffer()
	for _, v := range values {
		buf = basictl.StringWrite(buf, v)
	}
	b.ReportAllocs()
	b.ResetTimer()
	var res []byte // reuse
	finish := b.N / len(values)
	for i := 0; i < finish; i++ {
		buf2 := buf
		for _, v := range values {
			var err error
			buf2, err = basictl.StringReadBytes(buf2, &res)
			if err != nil {
				b.Fatalf("bad")
			}
			if len(v) != len(res) {
				b.Fatalf("bad")
			}
		}
		if len(buf2) != 0 {
			b.Fatalf("bad")
		}
	}
	printSizes(b, int64(len(buf))*int64(finish))
}

func BenchmarkStringReadTLBad(b *testing.B) {
	values, buf := prepareStringsBuffer()
	for _, v := range values {
		buf = TL2StringWrite(buf, v)
	}
	b.ReportAllocs()
	b.ResetTimer()
	var res []byte // reuse
	finish := b.N / len(values)
	for i := 0; i < finish; i++ {
		buf2 := buf
		for _, v := range values {
			var err error
			buf2, err = TL2StringReadBytes(buf2, &res)
			if err != nil {
				b.Fatalf("bad")
			}
			if len(v) != len(res) {
				b.Fatalf("bad")
			}
		}
		if len(buf2) != 0 {
			b.Fatalf("bad")
		}
	}
	printSizes(b, int64(len(buf))*int64(finish))
}

func BenchmarkStringReadMsgp(b *testing.B) {
	values, buf := prepareStringsBuffer()
	for _, v := range values {
		buf = appendString(buf, v)
		//buf = msgp.AppendString(buf, v)
	}
	b.ReportAllocs()
	b.ResetTimer()
	var res []byte // reuse
	finish := b.N / len(values)
	for i := 0; i < finish; i++ {
		buf2 := buf
		for _, v := range values {
			var err error
			res, buf2, err = msgp.ReadStringAsBytes(buf2, res)
			if err != nil {
				b.Fatalf("bad")
			}
			if len(v) != len(res) {
				b.Fatalf("bad")
			}
		}
		if len(buf2) != 0 {
			b.Fatalf("bad")
		}
	}
	printSizes(b, int64(len(buf))*int64(finish))
}

func BenchmarkStringReadProtobuf(b *testing.B) {
	values, buf := prepareStringsBuffer()
	for _, v := range values {
		buf = protowire.AppendString(buf, v)
	}
	b.ReportAllocs()
	b.ResetTimer()
	var res []byte // reuse
	finish := b.N / len(values)
	for i := 0; i < finish; i++ {
		buf2 := buf
		for _, v := range values {
			str, n := protowire.ConsumeBytes(buf2)
			res = append(res[:0], str...)
			if n < 0 {
				b.Fatalf("bad")
			}
			if len(v) != len(res) {
				b.Fatalf("bad")
			}
			buf2 = buf2[n:]
		}
		if len(buf2) != 0 {
			b.Fatalf("bad")
		}
	}
	printSizes(b, int64(len(buf))*int64(finish))
}
