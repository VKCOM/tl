package speed

import (
	"testing"

	"google.golang.org/protobuf/encoding/protowire"

	"github.com/VKCOM/tl/pkg/basictl"
	"github.com/tinylib/msgp/msgp"
)

func BenchmarkIntReadTL(b *testing.B) {
	values, buf := prepareIntsBuffer()
	for _, v := range values {
		buf = basictl.NatWrite(buf, v)
	}
	b.ReportAllocs()
	b.ResetTimer()
	finish := b.N / len(values)
	for i := 0; i < finish; i++ {
		buf2 := buf
		for _, v := range values {
			var err error
			var res uint32
			buf2, err = basictl.NatRead(buf2, &res)
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

func BenchmarkIntReadTL2(b *testing.B) {
	BenchmarkIntReadTL(b)
}

func BenchmarkIntReadMsgp(b *testing.B) {
	values, buf := prepareIntsBuffer()
	for _, v := range values {
		buf = appendUint32(buf, v)
	}
	b.ReportAllocs()
	b.ResetTimer()
	finish := b.N / len(values)
	for i := 0; i < finish; i++ {
		buf2 := buf
		for _, v := range values {
			var err error
			var res uint32
			res, buf2, err = msgp.ReadUint32Bytes(buf2)
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

func BenchmarkIntReadProtobuf(b *testing.B) {
	values, buf := prepareIntsBuffer()
	for _, v := range values {
		buf = protowire.AppendVarint(buf, uint64(v))
		//buf = binary.AppendUvarint(buf, uint64(v))
	}
	b.ReportAllocs()
	b.ResetTimer()
	finish := b.N / len(values)
	for i := 0; i < finish; i++ {
		buf2 := buf
		for _, v := range values {
			res, n := protoReadUint32(buf2)
			// vv, n := binary.Varint(buf2)
			if n <= 0 {
				b.Fatalf("bad")
			}
			if res != v {
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
