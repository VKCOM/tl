// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package pure

import (
	"fmt"
	"log"
	"strings"

	"github.com/vkcom/tl/internal/tlast"
	"github.com/vkcom/tl/pkg/basictl"
)

type TypeInstanceUnion struct {
	TypeInstanceCommon
	variantNames []string
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

func (k *Kernel) createUnion(canonicalName string, tip *KernelType, def tlast.TL2UnionType,
	leftArgs []tlast.TL2TypeTemplate, actualArgs []tlast.TL2TypeArgument) (TypeInstance, error) {
	ins := &TypeInstanceUnion{
		TypeInstanceCommon: TypeInstanceCommon{
			canonicalName: canonicalName,
			tip:           tip,
		},
		variantNames: make([]string, len(def.Variants)),
		variantTypes: make([]*TypeInstanceStruct, len(def.Variants)),
	}
	for i, variantDef := range def.Variants {
		element, err := k.createStruct(canonicalName+"__"+variantDef.Name, tip,
			!variantDef.IsTypeAlias, variantDef.TypeAlias, variantDef.Fields, leftArgs, actualArgs, true, i, nil)
		if err != nil {
			return nil, fmt.Errorf("fail to resolve type of union %s element %d: %w", canonicalName, i, err)
		}
		ins.variantTypes[i] = element
		ins.variantNames[i] = variantDef.Name
	}
	return ins, nil
}

func (k *Kernel) createUnionTL1FromTL1(canonicalName string, tip *KernelType,
	resolvedType tlast.TypeRef, definition []*tlast.Combinator,
	leftArgs []tlast.TemplateArgument, actualArgs []tlast.ArithmeticOrType) (TypeInstance, error) {

	_, natParams := k.getTL1Args(leftArgs, actualArgs)
	log.Printf("natParams for %s: %s", canonicalName, strings.Join(natParams, ","))

	ins := &TypeInstanceUnion{
		TypeInstanceCommon: TypeInstanceCommon{
			canonicalName: canonicalName,
			NatParams:     natParams,
			tip:           tip,
			rt:            resolvedType,
		},
		variantNames: make([]string, len(definition)),
		variantTypes: make([]*TypeInstanceStruct, len(definition)),
	}
	for i, variantDef := range definition {
		element, err := k.createStructTL1FromTL1(canonicalName+"__"+fmt.Sprintf("%d", i), tip,
			resolvedType,
			variantDef.Fields,
			leftArgs, actualArgs,
			true, i, nil)
		if err != nil {
			return nil, fmt.Errorf("fail to resolve type of union %s element %d: %w", canonicalName, i, err)
		}
		ins.variantTypes[i] = element
		ins.variantNames[i] = variantDef.Construct.Name.String()
	}
	return ins, nil
}
