package tl2pure

import (
	"math/rand"

	"github.com/vkcom/tl/internal/tlast"
)

type TypeInstanceUnion struct {
	TypeInstanceCommon
	def          tlast.TL2UnionType
	variantTypes []*TypeInstanceObject
}

type KernelValueUnion struct {
	instance *TypeInstanceUnion
	index    int
	variants []KernelValueObject // we remember state of all variants to improve editing experience
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
		variants: make([]KernelValueObject, len(ins.variantTypes)),
	}
	for i, vt := range ins.variantTypes {
		value.variants[i] = vt.CreateValueObject()
	}
	return value
}

func (v *KernelValueUnion) Random(rg *rand.Rand) {
	v.index = rg.Intn(len(v.variants))
	v.variants[v.index].Random(rg)
}

func (v *KernelValueUnion) WriteTL2(w []byte) []byte {
	return v.variants[v.index].WriteTL2(w)
}

func (v *KernelValueUnion) ReadTL2(w []byte) ([]byte, error) {
	return v.variants[v.index].ReadTL2(w)
}

func (v *KernelValueUnion) WriteJSON(w []byte) []byte {
	defVariant := v.instance.def.Variants[v.index]
	w = append(w, `{"type":"`...)
	w = append(w, defVariant.Name...)
	if len(v.instance.variantTypes[v.index].constructorFields) == 0 {
		return append(w, `"}`...)
	}
	w = append(w, `","value":`...)
	w = v.variants[v.index].WriteJSON(w)
	w = append(w, '}')
	return w
}

func (v *KernelValueUnion) Clone() KernelValue {
	clone := *v
	for i, v := range clone.variants {
		clone.variants[i] = v.CloneObject()
	}
	return &clone
}

func (v *KernelValueUnion) CompareForMapKey(other KernelValue) int {
	return 0
}

/*
func (k *Kernel) createUnion(canonicalName string, def tlast.TL2UnionType, lrc map[string]ResolvedArgument) (TypeInstance, error) {
	ins := &TypeInstanceUnion{
		TypeInstanceCommon: TypeInstanceCommon{
			canonicalName: canonicalName,
		},
		def:          def,
		variantTypes: make([]*TypeInstanceObject, len(def.Variants)),
	}
	for i, variantDef := range def.Variants {
		element, err := k.createObject(canonicalName+"__"+variantDef.Name, tlast.TL2TypeDeclaration{},
			!variantDef.IsTypeAlias, variantDef.TypeAlias, variantDef.Fields, lrc, true, i)
		if err != nil {
			return nil, fmt.Errorf("fail to resolve type of union %s element %d: %w", canonicalName, i, err)
		}
		ins.variantTypes[i] = element
	}
	return ins, nil
}
*/
