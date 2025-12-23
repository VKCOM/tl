package tl2pure

import (
	"math/rand"

	"github.com/vkcom/tl/pkg/basictl"
)

type TypeInstanceTupleVector struct {
	TypeInstanceCommon
	isTuple   bool
	count     int
	fieldType *TypeInstanceRef // TODO rename to elemType
}

type KernelValueTuple struct {
	instance *TypeInstanceTupleVector
	elements []KernelValue
}

func (ins *TypeInstanceTupleVector) FindCycle(c *cycleFinder) {
	if !c.push(ins) {
		return
	}
	defer c.pop(ins)
	if ins.isTuple {
		ins.fieldType.ins.FindCycle(c)
	}
}

func (ins *TypeInstanceTupleVector) CreateValue() KernelValue {
	value := &KernelValueTuple{
		instance: ins,
	}
	if ins.isTuple {
		value.resize(ins.count)
	}
	return value
}

func (ins *TypeInstanceTupleVector) SkipTL2(r []byte) ([]byte, error) {
	return basictl.SkipSizedValue(r)
}

func (v *KernelValueTuple) resize(count int) {
	v.elements = v.elements[:min(count, cap(v.elements))]
	for len(v.elements) < count {
		v.elements = append(v.elements, v.instance.fieldType.ins.CreateValue())
	}
	if len(v.elements) > count {
		v.elements = v.elements[:count]
	}
}

func (v *KernelValueTuple) Reset() {
	if !v.instance.isTuple {
		v.elements = v.elements[:0]
		return
	}
	for _, el := range v.elements {
		el.Reset()
	}
}

func (v *KernelValueTuple) Random(rg *rand.Rand) {
	if !v.instance.isTuple {
		count := 0
		if (rg.Uint32() & 3) != 0 { // many vectors empty
			count = 1 + rg.Intn(4)
		}
		v.resize(count)
	}
	for _, el := range v.elements {
		el.Random(rg)
	}
}

func (v *KernelValueTuple) WriteTL2(w []byte, optimizeEmpty bool, ctx *TL2Context) []byte {
	if len(v.elements) == 0 && optimizeEmpty {
		return w
	}

	oldLen := len(w)
	w = append(w, make([]byte, 16)...) // reserve space for

	firstUsedByte := len(w)

	w = basictl.TL2WriteSize(w, len(v.elements))

	for _, elem := range v.elements {
		w = elem.WriteTL2(w, false, ctx)
	}

	lastUsedByte := len(w)
	offset := basictl.TL2PutSize(w[oldLen:], lastUsedByte-firstUsedByte)
	offset += copy(w[oldLen+offset:], w[firstUsedByte:lastUsedByte])
	return w[:oldLen+offset]
}

func (v *KernelValueTuple) ReadTL2(r []byte, ctx *TL2Context) (_ []byte, err error) {
	currentSize := 0
	if r, currentSize, err = basictl.TL2ParseSize(r); err != nil {
		return r, err
	}
	if len(r) < currentSize {
		return r, basictl.TL2Error("not enough data: expected %d, got %d", currentSize, len(r))
	}

	currentR := r[:currentSize]
	r = r[currentSize:]

	elementCount := 0
	if currentSize != 0 {
		if currentR, elementCount, err = basictl.TL2ParseSize(currentR); err != nil {
			return r, err
		}
	}

	if !v.instance.isTuple {
		v.resize(elementCount)
	}
	for i := 0; i < min(len(v.elements), elementCount); i++ {
		if currentR, err = v.elements[i].ReadTL2(currentR, ctx); err != nil {
			return r, err
		}
	}
	for i := min(len(v.elements), elementCount); i < elementCount; i++ {
		if currentR, err = v.instance.fieldType.ins.SkipTL2(currentR); err != nil {
			return r, err
		}
	}
	for i := min(len(v.elements), elementCount); i < len(v.elements); i++ {
		v.elements[i].Reset()
	}
	return r, nil
}

func (v *KernelValueTuple) WriteJSON(w []byte, ctx *TL2Context) []byte {
	w = append(w, '[')
	first := true
	for _, el := range v.elements {
		if !first {
			w = append(w, ',')
		}
		first = false
		w = el.WriteJSON(w, ctx)
	}
	w = append(w, ']')
	return w
}

func (v *KernelValueTuple) Clone() KernelValue {
	clone := *v // TODO - copy slice
	for i, el := range clone.elements {
		clone.elements[i] = el.Clone()
	}
	return &clone
}

func (v *KernelValueTuple) CompareForMapKey(other KernelValue) int {
	return 0
}

func (k *Kernel) createTupleVector(canonicalName string, isTuple bool, count uint32, fieldType *TypeInstanceRef) TypeInstance {
	ins := &TypeInstanceTupleVector{
		TypeInstanceCommon: TypeInstanceCommon{
			canonicalName: canonicalName,
		},
		isTuple:   isTuple,
		count:     int(count),
		fieldType: fieldType,
	}
	return ins
}
