package pure

import (
	"fmt"

	"github.com/vkcom/tl/internal/tlast"
)

type TypeInstanceAlias struct {
	TypeInstanceCommon
	fieldType *TypeInstanceRef
}

func (ins *TypeInstanceAlias) GoodForMapKey() bool {
	return ins.fieldType.ins.GoodForMapKey()
}

func (ins *TypeInstanceAlias) IsBit() bool {
	return ins.fieldType.ins.IsBit()
}

func (ins *TypeInstanceAlias) FindCycle(c *cycleFinder) {
	if !c.push(ins) {
		return
	}
	defer c.pop(ins)
	ins.fieldType.ins.FindCycle(c)
}

func (ins *TypeInstanceAlias) CreateValue() KernelValue {
	value := &KernelValueAlias{
		instance: ins,
		value:    ins.fieldType.ins.CreateValue(),
	}
	return value
}

func (ins *TypeInstanceAlias) SkipTL2(r []byte) ([]byte, error) {
	return ins.fieldType.ins.SkipTL2(r)
}

func (k *Kernel) createAlias(canonicalName string, alias tlast.TL2TypeRef,
	leftArgs []tlast.TL2TypeTemplate, actualArgs []tlast.TL2TypeArgument) (TypeInstance, error) {
	rt, err := k.resolveType(alias, leftArgs, actualArgs)
	if err != nil {
		return nil, fmt.Errorf("fail to resolve type of alias %s to %s: %w", canonicalName, alias, err)
	}
	fieldType, err := k.getInstance(rt)
	if err != nil {
		return nil, fmt.Errorf("fail to instantiate alias %s to %s: %w", canonicalName, alias, err)
	}
	ins := &TypeInstanceAlias{
		TypeInstanceCommon: TypeInstanceCommon{
			canonicalName: canonicalName,
		},
		fieldType: fieldType,
	}
	return ins, nil
}
