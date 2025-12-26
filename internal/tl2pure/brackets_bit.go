package tl2pure

import (
	"math/rand"
	"strings"

	"github.com/vkcom/tl/pkg/basictl"
)

type TypeInstanceTupleVectorBit struct {
	TypeInstanceCommon
	isTuple bool
	count   int
}

type KernelValueTupleBit struct {
	instance *TypeInstanceTupleVectorBit
	elements []bool
}

func (ins *TypeInstanceTupleVectorBit) FindCycle(c *cycleFinder) {
}

func (ins *TypeInstanceTupleVectorBit) CreateValue() KernelValue {
	value := &KernelValueTupleBit{
		instance: ins,
	}
	if ins.isTuple {
		value.resize(ins.count)
	}
	return value
}

func (ins *TypeInstanceTupleVectorBit) SkipTL2(r []byte) ([]byte, error) {
	return basictl.SkipSizedValue(r)
}

func (v *KernelValueTupleBit) resize(count int) {
	v.elements = v.elements[:min(count, cap(v.elements))]
	if len(v.elements) < count {
		v.elements = append(v.elements, make([]bool, count-len(v.elements))...)
	}
	if len(v.elements) > count {
		v.elements = v.elements[:count]
	}
}

func (v *KernelValueTupleBit) Reset() {
	if !v.instance.isTuple {
		v.elements = v.elements[:0]
		return
	}
	clear(v.elements)
}

func (v *KernelValueTupleBit) Random(rg *rand.Rand) {
	if !v.instance.isTuple {
		count := 0
		if (rg.Uint32() & 3) != 0 { // many vectors empty
			count = 1 + rg.Intn(4)
		}
		v.resize(count)
	}
	for i := range v.elements {
		v.elements[i] = rg.Uint32()&1 != 0
	}
}

func (v *KernelValueTupleBit) WriteTL2(w []byte, optimizeEmpty bool, ctx *TL2Context) []byte {
	if len(v.elements) == 0 && optimizeEmpty {
		return w
	}

	oldLen := len(w)
	w = append(w, make([]byte, 16)...) // reserve space for

	firstUsedByte := len(w)

	w = basictl.TL2WriteSize(w, len(v.elements))

	w = basictl.VectorBoolContentWriteTL2(w, v.elements)

	lastUsedByte := len(w)
	offset := basictl.TL2PutSize(w[oldLen:], lastUsedByte-firstUsedByte)
	offset += copy(w[oldLen+offset:], w[firstUsedByte:lastUsedByte])
	return w[:oldLen+offset]
}

func (v *KernelValueTupleBit) ReadTL2(r []byte, ctx *TL2Context) (_ []byte, err error) {
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
		if !v.instance.isTuple && elementCount/8 > len(currentR) { // this is relaxed check, +7 could overflow
			return r, basictl.TL2ElementCountError(elementCount, currentR)
		}
	}
	if !v.instance.isTuple {
		v.resize(elementCount)
	}
	lastIndex := min(elementCount, elementCount)
	if _, err = basictl.VectorBoolContentReadTL2(currentR, v.elements[:lastIndex]); err != nil {
		return r, err
	}
	clear(v.elements[lastIndex:])
	// we skip excess element all at once. not one by one
	return r, nil
}

func (v *KernelValueTupleBit) WriteJSON(w []byte, ctx *TL2Context) []byte {
	w = append(w, '[')
	first := true
	for _, el := range v.elements {
		if !first {
			w = append(w, ',')
		}
		first = false
		value := KernelValueBool{el}
		w = value.WriteJSON(w, ctx)
	}
	w = append(w, ']')
	return w
}

func (v *KernelValueTupleBit) UIWrite(sb *strings.Builder, onPath bool, level int, path []int, model *UIModel) {
	sb.WriteString("<KernelValueTupleBit>")
}

func (v *KernelValueTupleBit) UIFixPath(level int, path []int) {
}

func (v *KernelValueTupleBit) Clone() KernelValue {
	return &KernelValueTupleBit{
		instance: v.instance,
		elements: append([]bool{}, v.elements...),
	}
}

func (v *KernelValueTupleBit) CompareForMapKey(other KernelValue) int {
	return 0
}
