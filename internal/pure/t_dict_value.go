package pure

import (
	"math/rand"
	"slices"
	"strings"

	"github.com/vkcom/tl/pkg/basictl"
)

type KernelValueDict struct {
	instance *TypeInstanceDict
	elements []KernelValueStruct // cap contains created elements
}

var _ KernelValue = &KernelValueDict{}

func (v *KernelValueDict) resize(count int) {
	v.elements = v.elements[:min(count, cap(v.elements))]
	for len(v.elements) < count {
		v.elements = append(v.elements, v.instance.fieldType.CreateValueObject())
	}
	if len(v.elements) > count {
		v.elements = v.elements[:count]
	}
}

func (v *KernelValueDict) sort() {
	slices.SortFunc(v.elements, func(a KernelValueStruct, b KernelValueStruct) int {
		return a.fields[0].CompareForMapKey(b.fields[0])
	})
	v.elements = slices.CompactFunc(v.elements, func(a KernelValueStruct, b KernelValueStruct) bool {
		return a.fields[0].CompareForMapKey(b.fields[0]) == 0
	})
}

func (v *KernelValueDict) Reset() {
	v.elements = v.elements[:0]
}

func (v *KernelValueDict) Random(rg *rand.Rand) {
	count := 0
	if (rg.Uint32() & 3) != 0 { // many vectors empty
		count = 1 + rg.Intn(4)
	}
	v.resize(count)
	for _, el := range v.elements {
		el.Random(rg)
	}
	v.sort()
}

func (v *KernelValueDict) WriteTL2(w *ByteBuilder, optimizeEmpty bool, onPath bool, level int, model *UIModel) {
	if len(v.elements) == 0 && optimizeEmpty {
		return
	}
	v.sort()

	firstUsedByte := w.ReserveSpaceForSize()
	w.WriteElementCount(len(v.elements))

	for _, elem := range v.elements {
		elem.WriteTL2(w, false, false, 0, model)
	}

	lastUsedByte := w.Len()
	w.FinishSize(firstUsedByte, lastUsedByte, optimizeEmpty)
}

func (v *KernelValueDict) ReadTL2(r []byte, ctx *TL2Context) (_ []byte, err error) {
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
		if elementCount > len(currentR) {
			return r, basictl.TL2ElementCountError(elementCount, currentR)
		}
	}

	v.resize(elementCount)
	for _, elem := range v.elements {
		if currentR, err = elem.ReadTL2(currentR, ctx); err != nil {
			return r, err
		}
	}
	v.sort()
	return r, nil
}

func (v *KernelValueDict) WriteJSON(w []byte, ctx *TL2Context) []byte {
	v.sort()
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

func (v *KernelValueDict) UIWrite(sb *strings.Builder, onPath bool, level int, model *UIModel) {
	sb.WriteString("<KernelValueDict>")
}

func (v *KernelValueDict) UIFixPath(side int, level int, model *UIModel) int {
	return 0
}

func (v *KernelValueDict) UIStartEdit(level int, model *UIModel, createMode int) {
}

func (v *KernelValueDict) UIKey(level int, model *UIModel, insert bool, delete bool, up bool, down bool) {
}

func (v *KernelValueDict) Clone() KernelValue {
	clone := *v
	for i, el := range clone.elements {
		clone.elements[i] = el.CloneObject()
	}
	return &clone
}

func (v *KernelValueDict) CompareForMapKey(other KernelValue) int {
	return 0
}
