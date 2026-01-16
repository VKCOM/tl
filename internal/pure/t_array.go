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
	count    int
	elemType *TypeInstanceRef

	dynamicSize bool           // for TL1 only, false for TL2
	bare        bool           // for TL1 only, false for TL2
	natArgs     []ActualNatArg // for TL1 only, empty for TL2
}

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
		value.resize(ins.count)
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
		count:    int(count),
		elemType: fieldType,
	}
	return ins
}

func (k *Kernel) createArrayTL1(canonicalName string, isTuple bool,
	leftArgs []tlast.TemplateArgument, actualArgs []tlast.ArithmeticOrType) (TypeInstance, error) {

	localArgs, natParams := k.getTL1Args(actualArgs)
	log.Printf("natParams for vector %s: %s", canonicalName, strings.Join(natParams, ","))

	elementT := actualArgs[0].T

	rt, natArgs, err := k.resolveTypeTL1(elementT, leftArgs, localArgs)
	if err != nil {
		return nil, fmt.Errorf("fail to resolve type of vector %s element: %w", canonicalName, err)
	}
	log.Printf("resolveType of vector for %s element: %s -> %s", canonicalName, elementT, rt.String())
	fieldIns, err := k.getInstanceTL1(rt)
	if err != nil {
		return nil, fmt.Errorf("fail to instantiate type of vector %s element: %w", canonicalName, err)
	}

	ins := &TypeInstanceArray{
		TypeInstanceCommon: TypeInstanceCommon{
			canonicalName: canonicalName,
			NatParams:     natParams,
		},
		isTuple:  isTuple,
		elemType: fieldIns,
		bare:     rt.Bare,
		natArgs:  natArgs,
	}
	if isTuple {
		ins.dynamicSize = actualArgs[1].IsArith
		ins.count = int(actualArgs[1].Arith.Res)
	}
	return ins, nil
}
