// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package pure

import (
	"fmt"
	"strings"

	"github.com/vkcom/tl/internal/tlast"
	"github.com/vkcom/tl/pkg/basictl"
)

type TypeInstanceUnion struct {
	TypeInstanceCommon
	variantNames             []string
	variantTL1ConstructNames []string // we keep them during migration to preserve legacy JSON format, if needed
	variantTypes             []*TypeInstanceStruct
	elementNatArgs           []ActualNatArg // empty for TL2
	isEnum                   bool
}

func (ins *TypeInstanceUnion) VariantNames() []string { return ins.variantNames }
func (ins *TypeInstanceUnion) VariantTL1ConstructNames() []string {
	return ins.variantTL1ConstructNames
}
func (ins *TypeInstanceUnion) VariantTypes() []*TypeInstanceStruct { return ins.variantTypes }
func (ins *TypeInstanceUnion) ElementNatArgs() []ActualNatArg      { return ins.elementNatArgs }
func (ins *TypeInstanceUnion) IsEnum() bool                        { return ins.isEnum }

func (ins *TypeInstanceUnion) BoxedOnly() bool {
	return true
}

// This is hint to generators
func (ins *TypeInstanceUnion) IsUnionMaybe() (isMaybe bool, elementField Field) {
	// this is not exhaustive, but good enough for now
	if len(ins.variantTypes) != 2 || ins.tip.canonicalName.String() != "Maybe" {
		return
	}
	if len(ins.variantTypes[0].fields) != 0 || len(ins.variantTypes[1].fields) != 1 {
		return
	}
	return true, ins.variantTypes[1].fields[0]
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
			isTopLevel:    tip.isTopLevel,
		},
		isEnum:                   true,
		variantNames:             make([]string, len(def.Variants)),
		variantTL1ConstructNames: make([]string, len(def.Variants)),
		variantTypes:             make([]*TypeInstanceStruct, len(def.Variants)),
	}
	for i, variantDef := range def.Variants {
		element, err := k.createStruct(canonicalName+"__"+variantDef.Name, tip,
			!variantDef.IsTypeAlias, variantDef.TypeAlias, variantDef.Fields, leftArgs, actualArgs, true, i, nil)
		if err != nil {
			return nil, fmt.Errorf("fail to resolve type of union %s element %d: %w", canonicalName, i, err)
		}
		ins.variantTypes[i] = element
		ins.variantNames[i] = variantDef.Name
		ins.variantTL1ConstructNames[i] = variantDef.Name
		if len(element.fields) != 0 {
			ins.isEnum = false
		}
	}
	return ins, nil
}

func (k *Kernel) createUnionTL1FromTL1(canonicalName string, tip *KernelType,
	resolvedType tlast.TypeRef, definition []*tlast.Combinator,
	leftArgs []tlast.TemplateArgument, actualArgs []tlast.ArithmeticOrType) (TypeInstance, error) {

	localArgs, natParams := k.getTL1Args(leftArgs, actualArgs)
	// log.Printf("natParams for %s: %s", canonicalName, strings.Join(natParams, ","))

	var natArgs []ActualNatArg
	for _, localArg := range localArgs { // pass all our parameters to our variant
		natArgs = append(natArgs, localArg.natArgs...)
	}
	// log.Printf("natArgs for %s union fields is: %v", canonicalName, natArgs)

	variantNames, err := k.VariantNames(definition)
	if err != nil {
		return nil, err
	}
	ins := &TypeInstanceUnion{
		TypeInstanceCommon: TypeInstanceCommon{
			canonicalName: canonicalName,
			tlName:        definition[0].TypeDecl.Name,
			tlTag:         0,
			natParams:     natParams,
			tip:           tip,
			isTopLevel:    false, // in TL1, union variants are top level, not union itself
			rt:            resolvedType,
			argNamespace:  k.getArgNamespace(resolvedType),
		},
		variantNames:             variantNames,
		variantTL1ConstructNames: make([]string, len(definition)),
		variantTypes:             make([]*TypeInstanceStruct, len(definition)),
		elementNatArgs:           natArgs,
		isEnum:                   true,
	}

	for i, variantDef := range definition {
		// do not change canonical names before removing long adapters,
		// otherwise long adapter discovery will break (see func (gen *genGo) findLongAdapter)
		argsStart := len(canonicalName)
		if st := strings.Index(canonicalName, "<"); st >= 0 {
			argsStart = st
		}
		variantCanonicalName := variantDef.Construct.Name.String() + canonicalName[argsStart:]
		element, err := k.createStructTL1FromTL1(variantCanonicalName, tip,
			resolvedType,
			variantDef,
			leftArgs, actualArgs,
			true, i)
		if err != nil {
			return nil, fmt.Errorf("fail to resolve type of union %s element %d: %w", canonicalName, i, err)
		}
		ins.variantTypes[i] = element
		ins.variantTL1ConstructNames[i] = variantDef.Construct.Name.String()
		if len(element.fields) != 0 {
			ins.isEnum = false
		}
	}
	return ins, nil
}
