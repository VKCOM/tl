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

type TypeInstanceArray struct {
	TypeInstanceCommon
	isTuple     bool
	count       uint32
	dynamicSize bool // for TL1 only, false for TL2

	field Field
}

func (ins *TypeInstanceArray) IsTuple() bool     { return ins.isTuple }
func (ins *TypeInstanceArray) Count() uint32     { return ins.count }
func (ins *TypeInstanceArray) DynamicSize() bool { return ins.dynamicSize }
func (ins *TypeInstanceArray) Field() Field      { return ins.field }

func (ins *TypeInstanceArray) FindCycle(c *cycleFinder) {
	if !c.push(ins) {
		return
	}
	defer c.pop(ins)
	if ins.isTuple {
		ins.field.ins.ins.FindCycle(c)
	}
}

func (ins *TypeInstanceArray) GetChildren(children []TypeInstance, withReturnType bool) []TypeInstance {
	return append(children, ins.field.ins.ins)
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
		isTuple: isTuple,
		count:   count,
	}
	ins.field = Field{
		owner: ins,
		ins:   fieldType,
		bare:  true,
	}
	return ins
}

func (k *Kernel) createVectorTL1(canonicalName string, tip *KernelType,
	resolvedType tlast.TypeRef,
	leftArgs []tlast.TemplateArgument, actualArgs []tlast.ArithmeticOrType) (TypeInstance, error) {

	localArgs, natParams := k.getTL1Args(leftArgs, actualArgs)
	// log.Printf("natParams for vector %s: %s", canonicalName, strings.Join(natParams, ","))

	fieldT := tlast.TypeRef{Type: tlast.Name{Name: "t"}}

	rt, fieldNatArgs, err := k.resolveTypeTL1(fieldT, leftArgs, localArgs)
	if err != nil {
		return nil, fmt.Errorf("fail to resolve type of vector %s field: %w", canonicalName, err)
	}
	// log.Printf("resolveTypeTL2 of vector for %s field: %s -> %s", canonicalName, fieldT, rt.String())
	fieldIns, fieldBare, err := k.getInstanceTL1(rt, true)
	if err != nil {
		return nil, fmt.Errorf("fail to instantiate type of vector %s field: %w", canonicalName, err)
	}

	ins := &TypeInstanceArray{
		TypeInstanceCommon: TypeInstanceCommon{
			canonicalName: canonicalName,
			natParams:     natParams,
			tip:           tip,
			rt:            resolvedType,
			argNamespace:  k.getArgNamespace(resolvedType),
		},
		isTuple: false,
	}
	ins.field = Field{
		owner:   ins,
		ins:     fieldIns,
		bare:    fieldBare,
		natArgs: fieldNatArgs,
	}
	return ins, nil
}

func (k *Kernel) createTupleTL1(canonicalName string, tip *KernelType,
	resolvedType tlast.TypeRef,
	leftArgs []tlast.TemplateArgument, actualArgs []tlast.ArithmeticOrType) (TypeInstance, error) {

	localArgs, natParams := k.getTL1Args(leftArgs, actualArgs)
	// log.Printf("natParams for tuple %s: %s", canonicalName, strings.Join(natParams, ","))

	fieldT := tlast.TypeRef{Type: tlast.Name{Name: "t"}}

	rt, natArgs, err := k.resolveTypeTL1(fieldT, leftArgs, localArgs)
	if err != nil {
		return nil, fmt.Errorf("fail to resolve type of tuple %s field: %w", canonicalName, err)
	}
	// log.Printf("resolveTypeTL2 of tuple for %s field: %s -> %s", canonicalName, fieldT, rt.String())
	fieldIns, fieldBare, err := k.getInstanceTL1(rt, true)
	if err != nil {
		return nil, fmt.Errorf("fail to instantiate type of tuple %s field: %w", canonicalName, err)
	}

	ins := &TypeInstanceArray{
		TypeInstanceCommon: TypeInstanceCommon{
			canonicalName: canonicalName,
			natParams:     natParams,
			tip:           tip,
			rt:            resolvedType,
			argNamespace:  k.getArgNamespace(resolvedType),
		},
		isTuple:     true,
		count:       actualArgs[0].Arith.Res,
		dynamicSize: !actualArgs[0].IsArith,
	}
	ins.field = Field{
		owner:   ins,
		ins:     fieldIns,
		bare:    fieldBare,
		natArgs: natArgs,
	}
	return ins, nil
}

func (k *Kernel) addTL1Brackets() {
	// for the purpose of type check, this is object with no fields, like __vector {t:Type} = ;
	combTL1 := &tlast.Combinator{
		Construct: tlast.Constructor{
			Name: tlast.Name{Name: "__dict"},
		},
		TemplateArguments: []tlast.TemplateArgument{{FieldName: "t", IsNat: false}},
	}
	kt := &KernelType{
		originTL2:      false,
		builtin:        true,
		combTL1:        []*tlast.Combinator{combTL1},
		instances:      map[string]*TypeInstanceRef{},
		tl1Names:       map[string]struct{}{"__dict": {}},
		tl2Names:       map[string]struct{}{},
		canonicalName:  tlast.Name{Name: "__dict"},
		historicalName: tlast.Name{Name: "BuiltinDict"},
		canBeBare:      true,
		targs:          make([]KernelTypeTarg, 1),
	}
	if err := k.addTip(kt, "__dict", ""); err != nil {
		panic(fmt.Sprintf("error adding __dict: %v", err))
	}
	combTL1 = &tlast.Combinator{
		Construct: tlast.Constructor{
			Name: tlast.Name{Name: "__vector"},
		},
		TemplateArguments: []tlast.TemplateArgument{{FieldName: "t", IsNat: false}},
	}
	kt = &KernelType{
		originTL2:      false,
		builtin:        true,
		combTL1:        []*tlast.Combinator{combTL1},
		instances:      map[string]*TypeInstanceRef{},
		tl1Names:       map[string]struct{}{"__vector": {}},
		tl2Names:       map[string]struct{}{},
		canonicalName:  tlast.Name{Name: "__vector"},
		historicalName: tlast.Name{Name: "BuiltinVector"},
		canBeBare:      true,
		targs:          make([]KernelTypeTarg, 1),
	}
	if err := k.addTip(kt, "__vector", ""); err != nil {
		panic(fmt.Sprintf("error adding __vector: %v", err))
	}
	combTL1 = &tlast.Combinator{
		Construct: tlast.Constructor{
			Name: tlast.Name{Name: "__tuple"},
		},
		TemplateArguments: []tlast.TemplateArgument{{FieldName: "n", IsNat: true}, {FieldName: "t", IsNat: false}},
	}
	kt = &KernelType{
		originTL2:      false,
		builtin:        true,
		combTL1:        []*tlast.Combinator{combTL1},
		instances:      map[string]*TypeInstanceRef{},
		tl1Names:       map[string]struct{}{"__tuple": {}},
		tl2Names:       map[string]struct{}{},
		canonicalName:  tlast.Name{Name: "__tuple"},
		historicalName: tlast.Name{Name: "BuiltinTuple"},
		canBeBare:      true,
		targs:          make([]KernelTypeTarg, 2),
	}
	if err := k.addTip(kt, "__tuple", ""); err != nil {
		panic(fmt.Sprintf("error adding __tuple: %v", err))
	}
}
