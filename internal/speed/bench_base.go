package speed

import (
	"testing"

	"github.com/VKCOM/tl/internal/speed/speed_tl/basictl"
	"google.golang.org/protobuf/encoding/protowire"
)

type TLData[T any] struct {
	WriteTL1Boxed func(x *T, buf []byte) []byte
	ReadTL1Boxed  func(x *T, buf []byte) ([]byte, error)

	WriteTL2 func(x *T, w []byte, tctx *basictl.TL2WriteContext) []byte
	ReadTL2  func(x *T, r []byte, tctx *basictl.TL2ReadContext) (_ []byte, err error)
}

type ProtoData[P any] struct {
	Size func(x *P) int

	MarshalToSizedBuffer func(x *P, dAtA []byte) (int, error)
	Marshal              func(x *P) (dAtA []byte, err error)

	Unmarshal func(x *P, dAtA []byte) error
}

type benchDataGenerator[T, P any] struct {
	TLProvider    TLData[T]
	ProtoProvider ProtoData[P]

	GenerateSamples func() []T
	GenerateBuffer  func() []byte
	MapSample       func(T) P

	CompareTLResult    func(T, T) bool
	CompareProtoResult func(P, P) bool
}

func benchBase[T, P any](b *testing.B, gen benchDataGenerator[T, P]) {
	values := gen.GenerateSamples()
	protoValues := make([]P, len(values))
	for i := range protoValues {
		protoValues[i] = gen.MapSample(values[i])
	}

	b.Run("Read", func(b *testing.B) {
		b.Run("TL1", func(b *testing.B) {
			buf := gen.GenerateBuffer()
			for _, v := range values {
				buf = gen.TLProvider.WriteTL1Boxed(&v, buf)
			}
			b.ReportAllocs()
			b.ResetTimer()
			finish := b.N / len(values)

			var p T

			for i := 0; i < finish; i++ {
				buf2 := buf
				for _, v := range values {
					var err error

					buf2, err = gen.TLProvider.ReadTL1Boxed(&p, buf2)
					if err != nil {
						b.Fatalf("bad")
					}
					if gen.CompareTLResult != nil && !gen.CompareTLResult(p, v) {
						b.Fatalf("bad")
					}
				}
				if len(buf2) != 0 {
					b.Fatalf("bad")
				}
			}
			printSizes(b, int64(len(buf))*int64(finish))
		})

		b.Run("TL2", func(b *testing.B) {
			buf := gen.GenerateBuffer()
			ctx := basictl.TL2WriteContext{SizeBuffer: make([]int, 10000)}

			for _, v := range values {
				buf = gen.TLProvider.WriteTL2(&v, buf, &ctx)
			}

			b.ReportAllocs()
			b.ResetTimer()
			finish := b.N / len(values)

			var p T

			for i := 0; i < finish; i++ {
				buf2 := buf
				for _, v := range values {
					var err error

					buf2, err = gen.TLProvider.ReadTL2(&p, buf2, nil)
					if err != nil {
						b.Fatalf("bad")
					}
					if gen.CompareTLResult != nil && !gen.CompareTLResult(p, v) {
						b.Fatalf("bad")
					}
				}
				if len(buf2) != 0 {
					b.Fatalf("bad")
				}
			}
			printSizes(b, int64(len(buf))*int64(finish))
		})

		b.Run("Proto", func(b *testing.B) {
			buf := gen.GenerateBuffer()
			for _, v := range protoValues {
				tmpBuf, err := gen.ProtoProvider.Marshal(&v)
				if err != nil {
					b.Fatalf("bad")
				}
				buf = protowire.AppendBytes(buf, tmpBuf)
			}
			b.ReportAllocs()
			b.ResetTimer()
			finish := b.N / len(values)

			var p P

			for i := 0; i < finish; i++ {
				buf2 := buf
				for _, v := range protoValues {
					var err error

					str, n := protowire.ConsumeBytes(buf2)
					if n < 0 {
						b.Fatalf("bad")
					}
					buf2 = buf2[n:]
					err = gen.ProtoProvider.Unmarshal(&p, str)
					if err != nil {
						b.Fatalf("bad")
					}

					if gen.CompareProtoResult != nil && !gen.CompareProtoResult(p, v) {
						b.Fatalf("bad")
					}
				}
				if len(buf2) != 0 {
					b.Fatalf("bad")
				}
			}
			printSizes(b, int64(len(buf))*int64(finish))
		})
	})

	b.Run("Write", func(b *testing.B) {
		b.Run("TL1", func(b *testing.B) {
			buf := gen.GenerateBuffer()

			b.ReportAllocs()
			b.ResetTimer()
			var total int64
			finish := b.N / len(values)
			for i := 0; i < finish; i++ {
				for j := range values {
					buf = gen.TLProvider.WriteTL1Boxed(&values[j], buf)
				}
				if len(buf) > bufferSize/2 {
					total += int64(len(buf))
					buf = buf[:0]
				}
			}

			printSizes(b, total+int64(len(buf)))
		})

		b.Run("TL2", func(b *testing.B) {
			buf := gen.GenerateBuffer()
			ctx := basictl.TL2WriteContext{SizeBuffer: make([]int, 1000)}

			b.ReportAllocs()
			b.ResetTimer()
			var total int64
			finish := b.N / len(values)
			for i := 0; i < finish; i++ {
				for j := range values {
					buf = gen.TLProvider.WriteTL2(&values[j], buf, &ctx)
				}
				if len(buf) > bufferSize/2 {
					total += int64(len(buf))
					buf = buf[:0]
				}
			}

			printSizes(b, total+int64(len(buf)))
		})

		b.Run("Proto", func(b *testing.B) {
			buf := gen.GenerateBuffer()

			b.ReportAllocs()
			b.ResetTimer()
			var total int64
			finish := b.N / len(protoValues)
			for i := 0; i < finish; i++ {
				for vIndex := range protoValues {
					s := gen.ProtoProvider.Size(&protoValues[vIndex])
					pointer := len(buf)
					buf = append(buf, make([]byte, s)...)

					_, err := gen.ProtoProvider.MarshalToSizedBuffer(&protoValues[vIndex], buf[pointer:])
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
		})
	})
}
