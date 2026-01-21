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
	"github.com/vkcom/tl/internal/utils"
	"github.com/vkcom/tl/pkg/basictl"
)

type TypeInstanceUnion struct {
	TypeInstanceCommon
	variantNames         []tlast.TL2TypeName
	variantOriginalNames []string
	variantTypes         []*TypeInstanceStruct
	isEnum               bool
}

func (ins *TypeInstanceUnion) VariantNames() []tlast.TL2TypeName   { return ins.variantNames }
func (ins *TypeInstanceUnion) VariantOriginalNames() []string      { return ins.variantOriginalNames }
func (ins *TypeInstanceUnion) VariantTypes() []*TypeInstanceStruct { return ins.variantTypes }
func (ins *TypeInstanceUnion) ElementNatArgs() []ActualNatArg      { return nil } // TODO
func (ins *TypeInstanceUnion) IsEnum() bool                        { return ins.isEnum }

func (ins *TypeInstanceUnion) BoxedOnly() bool {
	return true
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
		isEnum:               true,
		variantNames:         make([]tlast.TL2TypeName, len(def.Variants)),
		variantOriginalNames: make([]string, len(def.Variants)),
		variantTypes:         make([]*TypeInstanceStruct, len(def.Variants)),
	}
	for i, variantDef := range def.Variants {
		element, err := k.createStruct(canonicalName+"__"+variantDef.Name, tip,
			!variantDef.IsTypeAlias, variantDef.TypeAlias, variantDef.Fields, leftArgs, actualArgs, true, i, nil)
		if err != nil {
			return nil, fmt.Errorf("fail to resolve type of union %s element %d: %w", canonicalName, i, err)
		}
		ins.variantTypes[i] = element
		ins.variantNames[i] = tlast.TL2TypeName{Namespace: "", Name: variantDef.Name}
		ins.variantOriginalNames[i] = variantDef.Name
		if len(element.fields) != 0 {
			ins.isEnum = false
		}
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
		isEnum:               true,
		variantTypes:         make([]*TypeInstanceStruct, len(definition)),
		variantOriginalNames: make([]string, len(definition)),
	}
	// Removing prefix/suffix common with union name.
	// We allow relaxed case match. To use strict match, we could remove all strings.ToLower() calls below
	typePrefix := strings.ToLower(utils.ToLowerFirst(definition[0].TypeDecl.Name.Name))
	typeSuffix := strings.ToLower(definition[0].TypeDecl.Name.Name)
	for _, typ := range definition {
		conName := strings.ToLower(typ.Construct.Name.Name)
		// if constructor is full prefix of type, we will shorten accessors
		// ab.saveStateOne = ab.SaveState; // item.AsOne()
		// ab.saveStateTwo = ab.SaveState; // item.AsTwo()
		if !strings.HasPrefix(conName, typePrefix) { // same check as in checkUnionElementsCompatibility
			typePrefix = ""
		}
		if !strings.HasSuffix(conName, typeSuffix) {
			typeSuffix = ""
		}
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
		ins.variantOriginalNames[i] = variantDef.Construct.Name.String()
		typeConstructName := variantDef.Construct.Name
		if typePrefix != "" && len(typePrefix) < len(typeConstructName.Name) {
			typeConstructName.Name = typeConstructName.Name[len(typePrefix):]
		} else if typeSuffix != "" && len(typeSuffix) < len(typeConstructName.Name) {
			typeConstructName.Name = typeConstructName.Name[:len(typeConstructName.Name)-len(typeSuffix)]
		}
		variantName := tlast.TL2TypeName{Namespace: "", Name: typeConstructName.Name}
		// check against already defined fields
		for _, usedName := range ins.variantNames {
			if usedName == variantName {
				variantName.Namespace = typeConstructName.Namespace // add namespace on collision
				break
			}
		}
		// check again
		for _, usedName := range ins.variantNames {
			if usedName == variantName {
				return nil, fmt.Errorf("cannot define TL1 union - prohibited variant name collision")
			}
		}
		ins.variantNames = append(ins.variantNames, variantName)
		if len(element.fields) != 0 {
			ins.isEnum = false
		}
	}
	return ins, nil
}
