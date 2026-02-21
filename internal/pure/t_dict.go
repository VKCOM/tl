// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package pure

import (
	"fmt"

	"github.com/vkcom/tl/internal/tlast"
	"github.com/vkcom/tl/pkg/basictl"
)

type TypeInstanceDict struct {
	TypeInstanceCommon

	field     Field
	fieldType *TypeInstanceStruct // same as field.ins, but better typed
}

func (ins *TypeInstanceDict) Field() Field                   { return ins.field }
func (ins *TypeInstanceDict) FieldType() *TypeInstanceStruct { return ins.fieldType }

func (ins *TypeInstanceDict) FindCycle(c *cycleFinder) {
}

func (ins *TypeInstanceDict) GetChildren(children []TypeInstance, withReturnType bool) []TypeInstance {
	return append(children, ins.field.ins.ins)
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

func (k *Kernel) createDict(canonicalName string, keyType *TypeInstanceRef, fieldType *TypeInstanceRef) TypeInstance {
	ins := &TypeInstanceDict{
		TypeInstanceCommon: TypeInstanceCommon{
			canonicalName: canonicalName,
			tip:           nil, // TODO - dicts have no corresponding type
		},
		fieldType: &TypeInstanceStruct{
			TypeInstanceCommon: TypeInstanceCommon{
				canonicalName: canonicalName + "__elem",
				tip:           nil, //  TODO - TL2 dict elements have no corresponding type
			},
			isConstructorFields: true,
			fields: []Field{{
				name: "k",
				ins:  keyType,
			}, {
				name: "v",
				ins:  fieldType,
			}},
		},
	}
	return ins
}

func (k *Kernel) createDictTL1(canonicalName string, tip *KernelType,
	resolvedType tlast.TypeRef,
	leftArgs []tlast.TemplateArgument, actualArgs []tlast.ArithmeticOrType) (TypeInstance, error) {

	localArgs, natParams := k.getTL1Args(leftArgs, actualArgs)
	//log.Printf("natParams for dict %s: %s", canonicalName, strings.Join(natParams, ","))

	fieldT := tlast.TypeRef{Type: tlast.Name{Name: "t"}}

	rt, fieldNatArgs, err := k.resolveTypeTL1(fieldT, leftArgs, localArgs)
	if err != nil {
		return nil, fmt.Errorf("fail to resolve type of dict %s element: %w", canonicalName, err)
	}
	//log.Printf("resolveTypeTL2 of dict for %s element: %s -> %s", canonicalName, fieldT, rt.String())
	fieldIns, fieldBare, err := k.getInstanceTL1(rt, true)
	if err != nil {
		return nil, fmt.Errorf("fail to instantiate type of dict %s element: %w", canonicalName, err)
	}
	fieldInsStruct, ok := fieldIns.ins.(*TypeInstanceStruct)
	if !ok {
		return nil, fmt.Errorf("internal error: dict %s element is not a struct", canonicalName)
	}
	if !fieldInsStruct.fields[0].ins.ins.GoodForMapKey() {
		if len(rt.Args) < 1 { // should be impossible, but who knows
			return nil, rt.PR.BeautifulError(fmt.Errorf("dict %s key type must be bit, bool, string or integer", canonicalName))
		}
		return nil, rt.Args[0].T.PR.BeautifulError(fmt.Errorf("dict %s key type must be bit, bool, string or integer", canonicalName))
	}

	ins := &TypeInstanceDict{
		TypeInstanceCommon: TypeInstanceCommon{
			canonicalName: canonicalName,
			natParams:     natParams,
			tip:           tip,
			rt:            resolvedType,
			argNamespace:  k.getArgNamespace(resolvedType),
		},
		fieldType: fieldInsStruct,
	}
	ins.field = Field{
		owner:   ins,
		ins:     fieldIns,
		bare:    fieldBare,
		natArgs: fieldNatArgs,
	}
	return ins, nil
}
