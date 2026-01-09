package pure

import (
	"github.com/vkcom/tl/internal/tlast"
	"github.com/vkcom/tl/pkg/basictl"
)

type TypeInstanceDict struct {
	TypeInstanceCommon

	fieldType TypeInstanceStruct
}

func (ins *TypeInstanceDict) FindCycle(c *cycleFinder) {
}

func (ins *TypeInstanceDict) CreateValue() KernelValue {
	value := &KernelValueDict{
		instance: ins,
	}
	return value
}

func (ins *TypeInstanceDict) SkipTL2(r []byte) ([]byte, error) {
	return basictl.SkipSizedValue(r)
}

func (k *Kernel) createMap(canonicalName string, keyType *TypeInstanceRef, fieldType *TypeInstanceRef) TypeInstance {
	ins := &TypeInstanceDict{
		TypeInstanceCommon: TypeInstanceCommon{
			canonicalName: canonicalName,
		},
		fieldType: TypeInstanceStruct{
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
