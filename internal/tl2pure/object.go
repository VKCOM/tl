package tl2pure

import (
	"fmt"
	"math/rand"

	"github.com/vkcom/tl/internal/tlast"
)

type TypeInstanceObject struct {
	TypeInstanceCommon
	isConstructorFields bool
	constructorFields   []tlast.TL2Field
	fieldTypes          []*TypeInstanceRef
	isUnionElement      bool
	unionIndex          int
}

type KernelValueObject struct {
	instance *TypeInstanceObject
	fields   []KernelValue // nil if optional field not set, to break recursion
}

func (ins *TypeInstanceObject) GoodForMapKey() bool {
	return false
}

func (ins *TypeInstanceObject) FindCycle(c *cycleFinder) {
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

func (ins *TypeInstanceObject) CreateValue() KernelValue {
	v := ins.CreateValueObject()
	return &v
}

func (ins *TypeInstanceObject) CreateValueObject() KernelValueObject {
	value := KernelValueObject{
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

func (v *KernelValueObject) Random(rg *rand.Rand) {
	for i, fieldDef := range v.instance.constructorFields {
		ft := v.instance.fieldTypes[i]
		if fieldDef.IsOptional {
			set := rg.Uint32()&1 != 0
			if !set {
				v.fields[i] = nil
				continue
			}
			if v.fields[i] == nil {
				v.fields[i] = ft.ins.CreateValue()
			}
		}
		v.fields[i].Random(rg)
	}
}

func (v *KernelValueObject) WriteTL2(w []byte) []byte {
	for i, fieldDef := range v.instance.constructorFields {
		if fieldDef.IsOptional {
			if v.fields[i] != nil {
				w = append(w, 1)
				w = v.fields[i].WriteTL2(w)
			} else {
				w = append(w, 0)
			}
			continue
		}
		w = v.fields[i].WriteTL2(w)
	}
	return w
}

func (v *KernelValueObject) ReadTL2(w []byte) ([]byte, error) {
	return w, nil // TODO
}

func (v *KernelValueObject) WriteJSON(w []byte) []byte {
	if !v.instance.isConstructorFields {
		return v.fields[0].WriteJSON(w)
	}
	w = append(w, '{')
	first := true
	for i, fieldDef := range v.instance.constructorFields {
		if fieldDef.IsOptional {
			if v.fields[i] == nil {
				continue
			}
		}
		if !first {
			w = append(w, ',')
		}
		first = false
		w = append(w, "'"...)
		w = append(w, fieldDef.Name...)
		w = append(w, "':"...)
		w = v.fields[i].WriteJSON(w)
	}
	w = append(w, '}')
	return w
}

func (v *KernelValueObject) Clone() KernelValue {
	clone := v.CloneObject()
	return &clone
}

func (v *KernelValueObject) CloneObject() KernelValueObject {
	clone := *v
	for i, v := range clone.fields {
		if v != nil { // skip not set optional fields
			clone.fields[i] = v.Clone()
		}
	}
	return clone
}

func (v *KernelValueObject) CompareForMapKey(other KernelValue) int {
	return 0
}

func (k *Kernel) createObject(canonicalName string, declaration tlast.TL2TypeDeclaration,
	isConstructorFields bool, alias tlast.TL2TypeRef, constructorFields []tlast.TL2Field,
	templateArguments []tlast.TL2TypeTemplate, lrc []tlast.TL2TypeArgument,
	isUnionElement bool, unionIndex int) (*TypeInstanceObject, error) {

	ins := &TypeInstanceObject{
		TypeInstanceCommon: TypeInstanceCommon{
			canonicalName: canonicalName,
			declaration:   declaration,
		},
		isConstructorFields: isConstructorFields,
		constructorFields:   constructorFields,
		isUnionElement:      isUnionElement,
		unionIndex:          unionIndex,
	}
	if !isConstructorFields { // if we are here, this is union variant or function result, where alias is field 1
		ins.constructorFields = append(ins.constructorFields, tlast.TL2Field{Type: alias})
	}

	for _, fieldDef := range ins.constructorFields {
		rt, err := k.resolveType(fieldDef.Type, templateArguments, lrc)
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
