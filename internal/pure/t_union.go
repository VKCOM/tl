package pure

import (
	"fmt"

	"github.com/vkcom/tl/internal/tlast"
	"github.com/vkcom/tl/pkg/basictl"
)

type TypeInstanceUnion struct {
	TypeInstanceCommon
	def          tlast.TL2UnionType
	variantTypes []*TypeInstanceStruct
}

func (ins *TypeInstanceUnion) FindCycle(c *cycleFinder) {
	if !c.push(ins) {
		return
	}
	defer c.pop(ins)
	// any variant with a cycle is prohibited, because it could be set active
	for _, variant := range ins.variantTypes {
		variant.FindCycle(c)
	}
}

func (ins *TypeInstanceUnion) CreateValue() KernelValue {
	value := &KernelValueUnion{
		instance: ins,
		index:    0,
		variants: make([]KernelValueStruct, len(ins.variantTypes)),
	}
	for i, vt := range ins.variantTypes {
		value.variants[i] = vt.CreateValueObject()
	}
	return value
}

func (ins *TypeInstanceUnion) SkipTL2(r []byte) ([]byte, error) {
	return basictl.SkipSizedValue(r)
}

func (k *Kernel) createUnion(canonicalName string, def tlast.TL2UnionType,
	leftArgs []tlast.TL2TypeTemplate, actualArgs []tlast.TL2TypeArgument) (TypeInstance, error) {
	ins := &TypeInstanceUnion{
		TypeInstanceCommon: TypeInstanceCommon{
			canonicalName: canonicalName,
		},
		def:          def,
		variantTypes: make([]*TypeInstanceStruct, len(def.Variants)),
	}
	for i, variantDef := range def.Variants {
		element, err := k.createStruct(canonicalName+"__"+variantDef.Name,
			!variantDef.IsTypeAlias, variantDef.TypeAlias, variantDef.Fields, leftArgs, actualArgs, true, i, nil)
		if err != nil {
			return nil, fmt.Errorf("fail to resolve type of union %s element %d: %w", canonicalName, i, err)
		}
		ins.variantTypes[i] = element
	}
	return ins, nil
}
