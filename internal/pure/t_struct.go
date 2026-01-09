package pure

import (
	"fmt"

	"github.com/vkcom/tl/internal/tlast"
	"github.com/vkcom/tl/pkg/basictl"
)

type TypeInstanceStruct struct {
	TypeInstanceCommon
	isConstructorFields bool
	constructorFields   []tlast.TL2Field
	fieldTypes          []*TypeInstanceRef
	isUnionElement      bool
	unionIndex          int

	// if function
	resultType TypeInstance
}

func (ins *TypeInstanceStruct) FindCycle(c *cycleFinder) {
	if !c.push(ins) {
		return
	}
	defer c.pop(ins)
	for i, fieldDef := range ins.constructorFields {
		ft := ins.fieldTypes[i]
		if fieldDef.IsOptional {
			ft.ins.FindCycle(c)
		}
	}
}

func (ins *TypeInstanceStruct) CreateValue() KernelValue {
	v := ins.CreateValueObject()
	return &v
}

func (ins *TypeInstanceStruct) SkipTL2(r []byte) ([]byte, error) {
	return basictl.SkipSizedValue(r)
}

func (ins *TypeInstanceStruct) CreateValueObject() KernelValueStruct {
	value := KernelValueStruct{
		instance: ins,
		fields:   make([]KernelValue, len(ins.fieldTypes)),
	}
	for i, fieldDef := range ins.constructorFields {
		ft := ins.fieldTypes[i]
		if !fieldDef.IsOptional {
			value.fields[i] = ft.ins.CreateValue()
		}
	}
	return value
}

func (k *Kernel) createObject(canonicalName string,
	isConstructorFields bool, alias tlast.TL2TypeRef, constructorFields []tlast.TL2Field,
	leftArgs []tlast.TL2TypeTemplate, actualArgs []tlast.TL2TypeArgument,
	isUnionElement bool, unionIndex int, resultType TypeInstance) (*TypeInstanceStruct, error) {

	ins := &TypeInstanceStruct{
		TypeInstanceCommon: TypeInstanceCommon{
			canonicalName: canonicalName,
		},
		isConstructorFields: isConstructorFields,
		constructorFields:   constructorFields,
		isUnionElement:      isUnionElement,
		unionIndex:          unionIndex,
		resultType:          resultType,
	}
	if !isConstructorFields { // if we are here, this is union variant or function result, where alias is field 1
		ins.constructorFields = append(ins.constructorFields, tlast.TL2Field{Type: alias})
	}

	for _, fieldDef := range ins.constructorFields {
		rt, err := k.resolveType(fieldDef.Type, leftArgs, actualArgs)
		if err != nil {
			return nil, fmt.Errorf("fail to resolve type of object %s field %s: %w", canonicalName, fieldDef.Name, err)
		}
		fieldIns, err := k.getInstance(rt)
		if err != nil {
			return nil, fmt.Errorf("fail to insantiate type of object %s field %s: %w", canonicalName, fieldDef.Name, err)
		}
		ins.fieldTypes = append(ins.fieldTypes, fieldIns)
	}
	return ins, nil
}
