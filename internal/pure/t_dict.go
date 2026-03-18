// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package pure

import (
	"fmt"

	"github.com/VKCOM/tl/internal/tlast"
	"github.com/VKCOM/tl/pkg/basictl"
)

type TypeInstanceDict struct {
	TypeInstanceCommon

	field     Field
	fieldType *TypeInstanceStruct // same as field.ins, but better typed
}

func (ins *TypeInstanceDict) Field() Field                   { return ins.field }
func (ins *TypeInstanceDict) FieldType() *TypeInstanceStruct { return ins.fieldType }

func (ins *TypeInstanceDict) FindCycle(c *cycleFinder, prName tlast.PositionRange) {
}

func (ins *TypeInstanceDict) GetChildren(children []TypeInstance, withReturnType bool) []TypeInstance {
	return append(children, ins.field.ins.ins)
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

	_, natParams := k.fillLocalArg(tlast.TL2TypeArgument{Type: resolvedType}, "t")
	//_, natParams := k.fillLocalArg(tlast.TL2TypeArgument{Type: resolvedType.BracketType.ArrayType}, "t")

	fieldNatArgs := k.natParamsToActualNatArgs(natParams)

	fieldIns, fieldBare, err := k.getInstance(fieldRt, true)
	if err != nil {
		return nil, fmt.Errorf("fail to instantiate type of dict %s element: %w", canonicalName, err)
	}
	if fieldIns.ins == nil {
		return nil, fmt.Errorf("internal error: recursive dict %s element not supported", canonicalName)
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
			hasFetcher:    k.resolvedTypeNeedsFetcher(resolvedType),
			argNamespace:  k.getArgNamespace(resolvedType),
		},
		fieldType: fieldInsStruct,
	}
	ins.field = Field{
		owner:   ins,
		ins:     fieldIns,
		bare:    fieldBare,
		natArgs: fieldNatArgs,
		pr:      resolvedType.BracketType.ArrayType.PR, // to print recursive cycles, we do not need key type
	}
	return ins, nil
}
