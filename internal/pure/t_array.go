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

type TypeInstanceArray struct {
	TypeInstanceCommon
	isTuple  bool
	count    uint32
	elemType *TypeInstanceRef

	dynamicSize bool           // for TL1 only, false for TL2
	elemBare    bool           // for TL1 only, false for TL2
	natArgs     []ActualNatArg // for TL1 only, empty for TL2
}

func (ins *TypeInstanceArray) IsTuple() bool               { return ins.isTuple }
func (ins *TypeInstanceArray) ElemType() TypeInstance      { return ins.elemType.ins }
func (ins *TypeInstanceArray) ElemBare() bool              { return ins.elemBare }
func (ins *TypeInstanceArray) DynamicSize() bool           { return ins.dynamicSize }
func (ins *TypeInstanceArray) Count() uint32               { return ins.count }
func (ins *TypeInstanceArray) ElemNatArgs() []ActualNatArg { return ins.natArgs }

func (ins *TypeInstanceArray) FindCycle(c *cycleFinder) {
	if !c.push(ins) {
		return
	}
	defer c.pop(ins)
	if ins.isTuple {
		ins.elemType.ins.FindCycle(c)
	}
}

func (ins *TypeInstanceArray) CreateValue() KernelValue {
	value := &KernelValueArray{
		instance: ins,
	}
	if ins.isTuple {
		value.resize(int(ins.count))
	}
	return value
}

func (ins *TypeInstanceArray) SkipTL2(r []byte) ([]byte, error) {
	return basictl.SkipSizedValue(r)
}

func (k *Kernel) createArray(canonicalName string, isTuple bool, count uint32, fieldType *TypeInstanceRef) TypeInstance {
	if fieldType.ins.IsBit() {
		ins := &TypeInstanceArrayBit{
			TypeInstanceCommon: TypeInstanceCommon{
				canonicalName: canonicalName,
				tip:           nil, // TODO - arrays have no corresponding type
			},
			isTuple: isTuple,
			count:   int(count),
		}
		return ins
	}
	ins := &TypeInstanceArray{
		TypeInstanceCommon: TypeInstanceCommon{
			canonicalName: canonicalName,
			tip:           nil, // TODO - arrays have no corresponding type
		},
		isTuple:  isTuple,
		count:    count,
		elemType: fieldType,
	}
	return ins
}

func (k *Kernel) createVectorTL1(canonicalName string,
	resolvedType tlast.TypeRef,
	leftArgs []tlast.TemplateArgument, actualArgs []tlast.ArithmeticOrType) (TypeInstance, error) {

	localArgs, natParams := k.getTL1Args(leftArgs, actualArgs)
	log.Printf("natParams for vector %s: %s", canonicalName, strings.Join(natParams, ","))

	//elementT := actualArgs[0].T
	elementT := tlast.TypeRef{Type: tlast.Name{Name: "t"}}

	rt, natArgs, err := k.resolveTypeTL1(elementT, leftArgs, localArgs)
	if err != nil {
		return nil, fmt.Errorf("fail to resolve type of vector %s element: %w", canonicalName, err)
	}
	log.Printf("resolveType of vector for %s element: %s -> %s", canonicalName, elementT, rt.String())
	fieldIns, fieldBare, err := k.getInstanceTL1(rt, true)
	if err != nil {
		return nil, fmt.Errorf("fail to instantiate type of vector %s element: %w", canonicalName, err)
	}

	ins := &TypeInstanceArray{
		TypeInstanceCommon: TypeInstanceCommon{
			canonicalName: canonicalName,
			NatParams:     natParams,
			rt:            resolvedType,
		},
		isTuple:  false,
		elemType: fieldIns,
		elemBare: fieldBare,
		natArgs:  natArgs,
	}
	return ins, nil
}

func (k *Kernel) createTupleTL1(canonicalName string,
	resolvedType tlast.TypeRef,
	leftArgs []tlast.TemplateArgument, actualArgs []tlast.ArithmeticOrType) (TypeInstance, error) {

	localArgs, natParams := k.getTL1Args(leftArgs, actualArgs)
	log.Printf("natParams for tuple %s: %s", canonicalName, strings.Join(natParams, ","))

	// elementT := actualArgs[1].T
	elementT := tlast.TypeRef{Type: tlast.Name{Name: "t"}}

	rt, natArgs, err := k.resolveTypeTL1(elementT, leftArgs, localArgs)
	if err != nil {
		return nil, fmt.Errorf("fail to resolve type of tuple %s element: %w", canonicalName, err)
	}
	log.Printf("resolveType of tuple for %s element: %s -> %s", canonicalName, elementT, rt.String())
	fieldIns, fieldBare, err := k.getInstanceTL1(rt, true)
	if err != nil {
		return nil, fmt.Errorf("fail to instantiate type of tuple %s element: %w", canonicalName, err)
	}

	ins := &TypeInstanceArray{
		TypeInstanceCommon: TypeInstanceCommon{
			canonicalName: canonicalName,
			NatParams:     natParams,
			rt:            resolvedType,
		},
		isTuple:     true,
		elemType:    fieldIns,
		elemBare:    fieldBare,
		natArgs:     natArgs,
		dynamicSize: !actualArgs[0].IsArith,
		count:       actualArgs[0].Arith.Res,
	}
	return ins, nil
}
