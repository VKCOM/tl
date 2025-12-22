package tl2pure

import (
	"math/rand"
	"slices"

	"github.com/vkcom/tl/internal/tlast"
)

type TypeInstanceMap struct {
	TypeInstanceCommon

	fieldType TypeInstanceObject
}

type KernelValueMap struct {
	instance *TypeInstanceMap
	elements []KernelValueObject
}

func (ins *TypeInstanceMap) GoodForMapKey() bool {
	return false
}

func (ins *TypeInstanceMap) FindCycle(c *cycleFinder) {
}

func (ins *TypeInstanceMap) CreateValue() KernelValue {
	value := &KernelValueMap{
		instance: ins,
	}
	return value
}

func (v *KernelValueMap) resize(count uint32) {
	for uint32(len(v.elements)) < count {
		v.elements = append(v.elements, v.instance.fieldType.CreateValueObject())
	}
	if uint32(len(v.elements)) > count {
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

func (v *KernelValueMap) Random(rg *rand.Rand) {
	var count uint32
	if (rg.Uint32() & 3) != 0 { // many vectors empty
		count = 1 + uint32(rg.Intn(4))
	}
	v.resize(count)
	for _, el := range v.elements {
		el.Random(rg)
	}
	v.sort()
}

func (v *KernelValueMap) WriteTL2(w []byte) []byte {
	v.sort()
	lenValue := KernelValueUint32{value: uint32(len(v.elements))}
	w = lenValue.WriteTL2(w)
	for _, el := range v.elements {
		w = el.WriteTL2(w)
	}
	return w
}

func (v *KernelValueMap) ReadTL2(w []byte) (_ []byte, err error) {
	lenValue := KernelValueUint32{}
	if w, err = lenValue.ReadTL2(w); err != nil {
		return w, nil
	}
	v.resize(lenValue.value)
	for _, el := range v.elements {
		if w, err = el.ReadTL2(w); err != nil {
			return w, nil
		}
	}
	v.sort()
	return w, nil
}

func (v *KernelValueMap) WriteJSON(w []byte) []byte {
	v.sort()
	w = append(w, '[')
	first := true
	for _, el := range v.elements {
		if !first {
			w = append(w, ',')
		}
		first = false
		w = el.WriteJSON(w)
	}
	w = append(w, ']')
	return w
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
