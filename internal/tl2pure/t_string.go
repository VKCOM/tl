package tl2pure

import (
	"fmt"

	"github.com/vkcom/tl/internal/tlast"
	"github.com/vkcom/tl/pkg/basictl"
)

type TypeInstanceString struct {
	TypeInstanceCommon
}

func (ins *TypeInstanceString) GoodForMapKey() bool {
	return true
}

func (ins *TypeInstanceString) FindCycle(c *cycleFinder) {
}

func (ins *TypeInstanceString) CreateValue() KernelValue {
	return &KernelValueString{}
}

func (ins *TypeInstanceString) SkipTL2(r []byte) ([]byte, error) {
	return basictl.SkipSizedValue(r)
}

func (k *Kernel) addString() {
	name := "string"
	comb := tlast.TL2Combinator{
		TypeDecl: tlast.TL2TypeDeclaration{
			Name: tlast.TL2TypeName{Name: name},
			Type: tlast.TL2TypeDefinition{IsConstructorFields: true}, // for the purpose of type check, this is object with no fields
		},
	}
	ins := TypeInstanceString{
		TypeInstanceCommon: TypeInstanceCommon{
			canonicalName: name,
			comb:          comb,
		},
	}
	ref := &TypeInstanceRef{
		ins: &ins,
	}
	kt := &KernelType{
		comb:      comb,
		instances: map[string]*TypeInstanceRef{name: ref},
	}
	if _, ok := k.instances[name]; ok {
		panic(fmt.Sprintf("error adding primitive type %s: exist in global list", name))
	}
	if err := k.addTip(kt); err != nil {
		panic(fmt.Sprintf("error adding primitive type %s: %v", name, err))
	}
	k.instances[name] = ref
	// k.instancesOrdered = append(k.instancesOrdered, ref) - we do not yet know if we need them here
}
