package tl2pure

import (
	"math/rand"
)

type TypeInstanceTupleVector struct {
	TypeInstanceCommon
	isTuple   bool
	count     uint32
	fieldType *TypeInstanceRef // TODO rename to elemType
}

type KernelValueTuple struct {
	instance *TypeInstanceTupleVector
	elements []KernelValue
}

func (ins *TypeInstanceTupleVector) GoodForMapKey() bool {
	return false
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

func (v *KernelValueTuple) resize(count uint32) {
	for uint32(len(v.elements)) < count {
		v.elements = append(v.elements, v.instance.fieldType.ins.CreateValue())
	}
	if uint32(len(v.elements)) > count {
		v.elements = v.elements[:count]
	}
}

func (v *KernelValueTuple) Random(rg *rand.Rand) {
	if !v.instance.isTuple {
		var count uint32
		if (rg.Uint32() & 3) != 0 { // many vectors empty
			count = 1 + uint32(rg.Intn(4))
		}
		v.resize(count)
	}
	for _, el := range v.elements {
		el.Random(rg)
	}
}

func (v *KernelValueTuple) WriteTL2(w []byte) []byte {
	lenValue := KernelValueUint32{value: uint32(len(v.elements))}
	w = lenValue.WriteTL2(w)
	for _, el := range v.elements {
		w = el.WriteTL2(w)
	}
	return w
}

func (v *KernelValueTuple) ReadTL2(w []byte) (_ []byte, err error) {
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
	return w, nil
}

func (v *KernelValueTuple) WriteJSON(w []byte) []byte {
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
		count:     count,
		fieldType: fieldType,
	}
	return ins
}
