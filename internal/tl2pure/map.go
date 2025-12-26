package tl2pure

import (
	"math/rand"
	"slices"
	"strings"

	"github.com/vkcom/tl/internal/tlast"
	"github.com/vkcom/tl/pkg/basictl"
)

type TypeInstanceMap struct {
	TypeInstanceCommon

	fieldType TypeInstanceObject
}

type KernelValueMap struct {
	instance *TypeInstanceMap
	elements []KernelValueObject // cap contains created elements
}

func (ins *TypeInstanceMap) FindCycle(c *cycleFinder) {
}

func (ins *TypeInstanceMap) CreateValue() KernelValue {
	value := &KernelValueMap{
		instance: ins,
	}
	return value
}

func (ins *TypeInstanceMap) SkipTL2(r []byte) ([]byte, error) {
	return basictl.SkipSizedValue(r)
}

func (v *KernelValueMap) resize(count int) {
	v.elements = v.elements[:min(count, cap(v.elements))]
	for len(v.elements) < count {
		v.elements = append(v.elements, v.instance.fieldType.CreateValueObject())
	}
	if len(v.elements) > count {
		v.elements = v.elements[:count]
	}
}

func (v *KernelValueMap) sort() {
	slices.SortFunc(v.elements, func(a KernelValueObject, b KernelValueObject) int {
		return a.fields[0].CompareForMapKey(b.fields[0])
	})
	v.elements = slices.CompactFunc(v.elements, func(a KernelValueObject, b KernelValueObject) bool {
		return a.fields[0].CompareForMapKey(b.fields[0]) == 0
	})
}

func (v *KernelValueMap) Reset() {
	v.elements = v.elements[:0]
}

func (v *KernelValueMap) Random(rg *rand.Rand) {
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

func (v *KernelValueMap) WriteTL2(w []byte, optimizeEmpty bool, ctx *TL2Context) []byte {
	v.sort()

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

func (v *KernelValueMap) ReadTL2(r []byte, ctx *TL2Context) (_ []byte, err error) {
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

func (v *KernelValueMap) WriteJSON(w []byte, ctx *TL2Context) []byte {
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

func (v *KernelValueMap) UIWrite(sb *strings.Builder, onPath bool, level int, path []int, model *UIModel) {
	sb.WriteString("<KernelValueMap>")
}

func (v *KernelValueMap) UIFixPath(level int, path []int) {
}

func (v *KernelValueMap) Clone() KernelValue {
	clone := *v
	for i, el := range clone.elements {
		clone.elements[i] = el.CloneObject()
	}
	return &clone
}

func (v *KernelValueMap) CompareForMapKey(other KernelValue) int {
	return 0
}

func (k *Kernel) createMap(canonicalName string, keyType *TypeInstanceRef, fieldType *TypeInstanceRef) TypeInstance {
	ins := &TypeInstanceMap{
		TypeInstanceCommon: TypeInstanceCommon{
			canonicalName: canonicalName,
		},
		fieldType: TypeInstanceObject{
			TypeInstanceCommon: TypeInstanceCommon{
				canonicalName: canonicalName + "__elem",
			},
			isConstructorFields: true,
			constructorFields: []tlast.TL2Field{{
				Name: "k",
			}, {
				Name: "v",
			}},
			fieldTypes: []*TypeInstanceRef{keyType, fieldType},
		},
	}
	return ins
}
