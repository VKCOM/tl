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
			bytes = testObject.WriteTL2(nil, nil)
		}
		return bytes
	}
}

func BenchmarkTL1ReadRandomVector(b *testing.B) {
	b.ReportAllocs()

	valuesBytes := generateBenchmarkTestObjectBytes(b, 1)

	dst := tlbenchmarks.VrutoyTopLevelContainer{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := dst.Read(valuesBytes[i%NumberOfSamples])
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkTL2ReadRandomVector(b *testing.B) {
	b.ReportAllocs()

	valuesBytes := generateBenchmarkTestObjectBytes(b, 2)

	dst := tlbenchmarks.VrutoyTopLevelContainer{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := dst.ReadTL2(valuesBytes[i%NumberOfSamples], nil)
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
		_, err := dst.ReadTL2(value, nil)
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
		writeBuffer = dst.Write(writeBuffer[:0])
	}
}

func BenchmarkTL2WriteRandomVectorWithWriteBuffer(b *testing.B) {
	b.ReportAllocs()
	writeBuffer := make([]byte, 1000)
	context := basictl.TL2WriteContext{SizeBuffer: make([]int, 10000)}

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
		writeBuffer = dst.WriteTL2(writeBuffer[:0], &context)
	}
}

const NumberOfSamples = 50

func BenchmarkTL1WriteRandomArrayWithWriteBuffer(b *testing.B) {
	b.ReportAllocs()

	dsts := generateBenchmarkTestObjects(b)
	writeBuffer := make([]byte, 1000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		writeBuffer, _ = dsts[i%NumberOfSamples].Write(writeBuffer[:0])
	}
}

func BenchmarkTL2WriteRandomArrayWithWriteBuffer(b *testing.B) {
	b.ReportAllocs()

	dsts := generateBenchmarkTestObjects(b)
	writeBuffer := make([]byte, 1000)
	context := basictl.TL2WriteContext{SizeBuffer: make([]int, 10000)}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		writeBuffer = dsts[i%NumberOfSamples].WriteTL2(writeBuffer[:0], &context)
	}
}

func generateBenchmarkTestObjectBytes(b *testing.B, version int) [][]byte {
	valueGen := initTestValues[*tlbenchmarks.VrutoyTopLevelContainer](version, func() *tlbenchmarks.VrutoyTopLevelContainer {
		return &tlbenchmarks.VrutoyTopLevelContainer{}
	})

	dsts := make([][]byte, NumberOfSamples)

	for i := 0; i < NumberOfSamples; i++ {
		dsts[i] = valueGen(i)
	}
	return dsts
}

func generateBenchmarkTestObjects(b *testing.B) []tlbenchmarks.VrutoyTopLevelContainerWithDependency {
	valueGen := initTestValues[*tlbenchmarks.VrutoyTopLevelContainerWithDependency](1, func() *tlbenchmarks.VrutoyTopLevelContainerWithDependency {
		return &tlbenchmarks.VrutoyTopLevelContainerWithDependency{}
	})

	dsts := make([]tlbenchmarks.VrutoyTopLevelContainerWithDependency, NumberOfSamples)

	for i := 0; i < NumberOfSamples; i++ {
		value := valueGen(i)
		_, err2 := dsts[i].Read(value)
		if err2 != nil {
			b.Fail()
		}
	}
	return dsts
}
