package casetests

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/meta"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/tlbenchmarks"
	"github.com/vkcom/tl/pkg/basictl"
	"math/rand"
	"testing"
)

const seed int64 = 11234567

func initTestValues[T meta.Object](tlVersion int, constructor func() T) func(iteration int) []byte {
	r := basictl.NewRandGenerator(rand.New(rand.NewSource(seed)))

	return func(iteration int) []byte {
		testObject := constructor()
		testObject.FillRandom(r)

		var bytes []byte
		var err error

		if tlVersion == 1 {
			bytes, err = testObject.WriteGeneral(nil)
			if err != nil {
				panic("can't init such data")
			}
		} else {
			bytes, _ = testObject.WriteTL2(nil, nil)
		}
		return bytes
	}
}

func BenchmarkTL1ReadRandomVector(b *testing.B) {
	b.ReportAllocs()

	valueGen := initTestValues[*tlbenchmarks.VrutoyTopLevelContainer](1, func() *tlbenchmarks.VrutoyTopLevelContainer {
		return &tlbenchmarks.VrutoyTopLevelContainer{}
	})
	dst := tlbenchmarks.VrutoyTopLevelContainer{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		value := valueGen(i)
		b.StartTimer()
		_, err := dst.Read(value)
		if err != nil {
			b.Fail()
		}
	}
}

func BenchmarkTL2ReadRandomVector(b *testing.B) {
	b.ReportAllocs()

	valueGen := initTestValues[*tlbenchmarks.VrutoyTopLevelContainer](2, func() *tlbenchmarks.VrutoyTopLevelContainer {
		return &tlbenchmarks.VrutoyTopLevelContainer{}
	})
	dst := tlbenchmarks.VrutoyTopLevelContainer{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		value := valueGen(i)
		b.StartTimer()
		_, err := dst.ReadTL2(value)
		if err != nil {
			b.Fail()
		}
	}
}

func BenchmarkTL1ReadRandomArray(b *testing.B) {
	b.ReportAllocs()

	valueGen := initTestValues[*tlbenchmarks.VrutoyTopLevelContainerWithDependency](1, func() *tlbenchmarks.VrutoyTopLevelContainerWithDependency {
		return &tlbenchmarks.VrutoyTopLevelContainerWithDependency{}
	})
	dst := tlbenchmarks.VrutoyTopLevelContainerWithDependency{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		value := valueGen(i)
		b.StartTimer()
		_, err := dst.Read(value)
		if err != nil {
			b.Fail()
		}
	}
}

func BenchmarkTL2ReadRandomArray(b *testing.B) {
	b.ReportAllocs()

	valueGen := initTestValues[*tlbenchmarks.VrutoyTopLevelContainerWithDependency](2, func() *tlbenchmarks.VrutoyTopLevelContainerWithDependency {
		return &tlbenchmarks.VrutoyTopLevelContainerWithDependency{}
	})
	dst := tlbenchmarks.VrutoyTopLevelContainerWithDependency{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		value := valueGen(i)
		b.StartTimer()
		_, err := dst.ReadTL2(value)
		if err != nil {
			b.Fail()
		}
	}
}

func BenchmarkTL1WriteRandomVector(b *testing.B) {
	b.ReportAllocs()

	valueGen := initTestValues[*tlbenchmarks.VrutoyTopLevelContainer](1, func() *tlbenchmarks.VrutoyTopLevelContainer {
		return &tlbenchmarks.VrutoyTopLevelContainer{}
	})
	dst := tlbenchmarks.VrutoyTopLevelContainer{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		value := valueGen(i)
		_, err2 := dst.Read(value)
		if err2 != nil {
			b.Fail()
		}
		b.StartTimer()
		_ = dst.Write(nil)
	}
}

func BenchmarkTL1WriteRandomVectorWithWriteBuffer(b *testing.B) {
	b.ReportAllocs()
	writeBuffer := make([]byte, 1000)

	valueGen := initTestValues[*tlbenchmarks.VrutoyTopLevelContainer](1, func() *tlbenchmarks.VrutoyTopLevelContainer {
		return &tlbenchmarks.VrutoyTopLevelContainer{}
	})
	dst := tlbenchmarks.VrutoyTopLevelContainer{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		value := valueGen(i)
		_, err2 := dst.Read(value)
		if err2 != nil {
			b.Fail()
		}
		b.StartTimer()
		writeBuffer = dst.Write(writeBuffer[0:0])
	}
}

func BenchmarkTL2WriteRandomVector(b *testing.B) {
	b.ReportAllocs()
	writeBuffer := make([]byte, 1000)
	buffer := make([]int, 10000)

	valueGen := initTestValues[*tlbenchmarks.VrutoyTopLevelContainer](1, func() *tlbenchmarks.VrutoyTopLevelContainer {
		return &tlbenchmarks.VrutoyTopLevelContainer{}
	})
	dst := tlbenchmarks.VrutoyTopLevelContainer{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		value := valueGen(i)
		_, err2 := dst.Read(value)
		if err2 != nil {
			b.Fail()
		}
		b.StartTimer()
		writeBuffer, buffer = dst.WriteTL2(writeBuffer[0:0], buffer[0:0])
	}
}

func BenchmarkTL1WriteRandomArrayWithWriteBuffer(b *testing.B) {
	b.ReportAllocs()

	writeBuffer := make([]byte, 1000)

	valueGen := initTestValues[*tlbenchmarks.VrutoyTopLevelContainerWithDependency](1, func() *tlbenchmarks.VrutoyTopLevelContainerWithDependency {
		return &tlbenchmarks.VrutoyTopLevelContainerWithDependency{}
	})
	dst := tlbenchmarks.VrutoyTopLevelContainerWithDependency{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		value := valueGen(i)
		_, err2 := dst.Read(value)
		if err2 != nil {
			b.Fail()
		}
		b.StartTimer()
		writeBuffer, _ = dst.Write(writeBuffer[0:0])
	}
}

func BenchmarkTL2WriteRandomArrayWithWriteBuffer(b *testing.B) {
	b.ReportAllocs()
	writeBuffer := make([]byte, 1000)
	buffer := make([]int, 10000)

	valueGen := initTestValues[*tlbenchmarks.VrutoyTopLevelContainerWithDependency](1, func() *tlbenchmarks.VrutoyTopLevelContainerWithDependency {
		return &tlbenchmarks.VrutoyTopLevelContainerWithDependency{}
	})
	dst := tlbenchmarks.VrutoyTopLevelContainerWithDependency{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		value := valueGen(i)
		_, err2 := dst.Read(value)
		if err2 != nil {
			b.Fail()
		}
		b.StartTimer()
		writeBuffer, buffer = dst.WriteTL2(writeBuffer[0:0], buffer[0:0])
	}
}
