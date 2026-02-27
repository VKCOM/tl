// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package pure

import (
	"fmt"

	"github.com/vkcom/tl/internal/tlast"
)

type TypeInstanceAlias struct {
	TypeInstanceCommon
	fieldType *TypeInstanceRef
	fieldBare bool // TODO - actually use it
}

func (ins *TypeInstanceAlias) GoodForMapKey() bool {
	return ins.fieldType.ins.GoodForMapKey()
}

func (ins *TypeInstanceAlias) IsBit() bool {
	return ins.fieldType.ins.IsBit()
}

func (ins *TypeInstanceAlias) FindCycle(c *cycleFinder) {
	if !c.push(ins) {
		return
	}
	defer c.pop(ins)
	ins.fieldType.ins.FindCycle(c)
}

func (ins *TypeInstanceAlias) GetChildren(children []TypeInstance, withReturnType bool) []TypeInstance {
	return append(children, ins.fieldType.ins)
}

func (ins *TypeInstanceAlias) CreateValue() KernelValue {
	value := &KernelValueAlias{
		instance: ins,
		value:    ins.fieldType.ins.CreateValue(),
	}
	return value
}

func (ins *TypeInstanceAlias) SkipTL2(r []byte) ([]byte, error) {
	return ins.fieldType.ins.SkipTL2(r)
}

func (k *Kernel) createAliasTL2(canonicalName string, tip *KernelType, resolvedType tlast.TL2TypeRef,
	alias tlast.TL2TypeRef,
	leftArgs []tlast.TL2TypeTemplate) (TypeInstance, error) {

	localArgs, natParams := k.getTL1ArgsHybrid(tip.templateArguments, resolvedType)
	if len(natParams) != 0 {
		panic("TODO - process TL1 types as usual")
	}
	rt, natArgs, err := k.resolveType(true, alias, leftArgs, localArgs)
	if len(natArgs) != 0 {
		panic("TODO - process TL1 types as usual")
	}
	if err != nil {
		return nil, fmt.Errorf("fail to resolve type of alias %s to %s: %w", canonicalName, alias, err)
	}
	fieldType, fieldBare, err := k.getInstance(rt, true)
	if err != nil {
		return nil, fmt.Errorf("fail to instantiate alias %s to %s: %w", canonicalName, alias, err)
	}
	ins := &TypeInstanceAlias{
		TypeInstanceCommon: TypeInstanceCommon{
			canonicalName: canonicalName,
			tip:           tip,
			natParams:     natParams,
			resolvedType:  resolvedType,
			isTopLevel:    tip.isTopLevel,
			hasTL2:        true,
			commentBefore: tip.combTL2.CommentBefore,
			commentRight:  "", // TODO - no comment right in TL2?
		},
		fieldType: fieldType,
		fieldBare: fieldBare,
	}
	return ins, nil
}
