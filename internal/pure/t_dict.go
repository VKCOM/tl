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

func (k *Kernel) createDict(canonicalName string,
	resolvedType tlast.TL2TypeRef) (TypeInstance, error) {

	fieldRt := tlast.TL2TypeRef{
		SomeType: tlast.TL2TypeApplication{
			Name: tlast.TL2TypeName{Name: "__dict_field"},
			Arguments: []tlast.TL2TypeArgument{
				resolvedType.BracketType.IndexType, {
					Type: resolvedType.BracketType.ArrayType,
				},
			},
			PR: resolvedType.PR, // TODO - check all PRs
		},
		PR: resolvedType.PR, // TODO - check all PRs
	}

	_, natParams := k.getTL1ArgHybrid(tlast.TL2TypeArgument{Type: resolvedType}, "t")
	//_, natParams := k.getTL1ArgHybrid(tlast.TL2TypeArgument{Type: resolvedType.BracketType.ArrayType}, "t")

	var fieldNatArgs []ActualNatArg
	for _, param := range natParams {
		fieldNatArgs = append(fieldNatArgs, ActualNatArg{
			name: param,
		})
	}

	//ktField, ok := k.tips["__dict_field"]
	//if !ok {
	//	panic(fmt.Errorf("internal error - built in __dict_field type not found"))
	//}
	//
	//fieldIns := &TypeInstanceStruct{
	//	TypeInstanceCommon: TypeInstanceCommon{
	//		canonicalName: canonicalName,
	//		tlName:        ktField.canonicalName,
	//		tlTag:         0,
	//		natParams:     natParams,
	//		tip:           ktField,
	//		isTopLevel:    false,
	//		resolvedType:  resolvedType,
	//		argNamespace:  k.getArgNamespace(resolvedType),
	//		hasTL2:        false, // could be marked later
	//	},
	//	isConstructorFields: true,
	//	isUnionElement:      isUnionElement,
	//	unionIndex:          unionIndex,
	//	isUnwrap:            tip.builtinWrappedCanonicalName != "",
	//}

	//localArgs, natParams := k.getTL1ArgsHybrid(tip.templateArguments, resolvedType)
	//fmt.Printf("natParams for dict %s: %s\n", canonicalName, strings.Join(natParams, ","))

	//fieldT := tlast.TypeRef{
	//	Type: tlast.Name{Name: "__dict_field"},
	//	Args: []tlast.ArithmeticOrType{{
	//		T: tlast.TypeRef{Type: tlast.Name{Name: "k"}},
	//	}, {
	//		T: tlast.TypeRef{Type: tlast.Name{Name: "v"}},
	//	}},
	//}
	//rt, fieldNatArgs, err := k.resolveType(false, k.convertTypeRef(fieldT), ktField.templateArguments, localArgs)
	//if err != nil {
	//	return nil, err
	//}
	////fmt.Printf("resolveTypeTL2 of dict for %s element: %s -> %s\n", canonicalName, fieldT, rt.String())
	fieldIns, fieldBare, err := k.getInstance(fieldRt, true)
	if err != nil {
		return nil, fmt.Errorf("fail to instantiate type of dict %s element: %w", canonicalName, err)
	}
	fieldInsStruct, ok := fieldIns.ins.(*TypeInstanceStruct)
	if !ok {
		return nil, fmt.Errorf("internal error: dict %s element is not a struct", canonicalName)
	}
	if !fieldInsStruct.fields[0].ins.ins.GoodForMapKey() {
		return nil, resolvedType.BracketType.IndexType.PR.BeautifulError(fmt.Errorf("dict %s key type must be bit, bool, string or integer", canonicalName))
	}
	if !fieldBare {
		return nil, fmt.Errorf("internal error dict %s field is not bare", canonicalName)
	}
	ins := &TypeInstanceDict{
		TypeInstanceCommon: TypeInstanceCommon{
			canonicalName: canonicalName,
			natParams:     natParams,
			tip:           nil,
			resolvedType:  resolvedType,
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
	//ins := &TypeInstanceDict{
	//	TypeInstanceCommon: TypeInstanceCommon{
	//		canonicalName: canonicalName,
	//		tip:           nil, // TODO - dicts have no corresponding type
	//	},
	//	fieldType: &TypeInstanceStruct{
	//		TypeInstanceCommon: TypeInstanceCommon{
	//			canonicalName: canonicalName + "__elem",
	//			tip:           nil, //  TODO - TL2 dict elements have no corresponding type
	//		},
	//		isConstructorFields: true,
	//		fields: []Field{{
	//			name: "k",
	//			ins:  keyType,
	//		}, {
	//			name: "v",
	//			ins:  fieldType,
	//			bare: fieldBare,
	//		}},
	//	},
	//}
}
