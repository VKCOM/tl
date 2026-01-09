package tl2pure

import (
	"github.com/vkcom/tl/pkg/basictl"
)

type TypeInstanceArray struct {
	TypeInstanceCommon
	isTuple   bool
	count     int
	fieldType *TypeInstanceRef // TODO rename to elemType
}

func (ins *TypeInstanceArray) FindCycle(c *cycleFinder) {
	if !c.push(ins) {
		return
	}
	defer c.pop(ins)
	if ins.isTuple {
		ins.fieldType.ins.FindCycle(c)
	}
}

func (ins *TypeInstanceArray) CreateValue() KernelValue {
	value := &KernelValueArray{
		instance: ins,
	}
	if ins.isTuple {
		value.resize(ins.count)
	}
	return value
}

func (ins *TypeInstanceArray) SkipTL2(r []byte) ([]byte, error) {
	return basictl.SkipSizedValue(r)
}

func (k *Kernel) createTupleVector(canonicalName string, isTuple bool, count uint32, fieldType *TypeInstanceRef) TypeInstance {
	if fieldType.ins.IsBit() {
		ins := &TypeInstanceArrayBit{
			TypeInstanceCommon: TypeInstanceCommon{
				canonicalName: canonicalName,
			},
			isTuple: isTuple,
			count:   int(count),
		}
		return ins
	}
	ins := &TypeInstanceArray{
		TypeInstanceCommon: TypeInstanceCommon{
			canonicalName: canonicalName,
		},
		isTuple:   isTuple,
		count:     int(count),
		fieldType: fieldType,
	}
	return ins
}
