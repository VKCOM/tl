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

//func (k *Kernel) createArray(canonicalName string, tip *KernelType, resolvedType tlast.TypeRef,
//	isTuple bool, count uint32, fieldType *TypeInstanceRef, fieldBare bool) TypeInstance {
//	if fieldType.ins.IsBit() {
//		ins := &TypeInstanceArrayBit{
//			TypeInstanceCommon: TypeInstanceCommon{
//				canonicalName: canonicalName,
//				tip:           tip,
//				rt:            resolvedType,
//				argNamespace:  "", // k.getArgNamespace(resolvedType), // should be empty
//			},
//			isTuple: isTuple,
//			count:   int(count),
//		}
//		return ins
//	}
//	ins := &TypeInstanceArray{
//		TypeInstanceCommon: TypeInstanceCommon{
//			canonicalName: canonicalName,
//			tip:           tip,
//			rt:            resolvedType,
//			argNamespace:  k.getArgNamespace(resolvedType),
//		},
//		isTuple: isTuple,
//		count:   count,
//	}
//	ins.field = Field{
//		owner: ins,
//		ins:   fieldType,
//		bare:  fieldBare,
//	}
//	return ins
//}

func (k *Kernel) createVectorTL1(canonicalName string, tip *KernelType,
	resolvedType tlast.TypeRef, resolvedType2 tlast.TL2TypeRef,
	leftArgs []tlast.TemplateArgument) (TypeInstance, error) {

	// TODO - do not derive parameters from getTL1ArgsHybrid, call fillNatParamHybrid directly
	localArgs, natParams := k.getTL1Args(leftArgs, resolvedType.Args)
	localArgs2, natParams2 := k.getTL1ArgsHybrid(tip.templateArguments, resolvedType2)
	_, natParams3 := k.getTL1ArgHybrid(tlast.TL2TypeArgument{Type: resolvedType2.BracketType.ArrayType}, "t")
	if a, b := strings.Join(natParams, ","), strings.Join(natParams2, ","); a != b || len(localArgs) != len(localArgs2) {
		panic(fmt.Errorf("!equalNatParams %s %s", a, b))
	}
	// log.Printf("natParams for vector %s: %s", canonicalName, strings.Join(natParams, ","))

	fieldT := tlast.TypeRef{Type: tlast.Name{Name: "t"}}

	rt, fieldNatArgs, err := k.resolveTypeTL1(fieldT, leftArgs, localArgs)
	if err != nil {
		return nil, fmt.Errorf("fail to resolve type of vector %s field: %w", canonicalName, err)
	}
	rt2, fieldNatArgs2, err := k.resolveTypeHybrid(false, k.convertTypeRef(fieldT), leftArgs, localArgs2)
	if err != nil {
		return nil, err
	}
	k.equalTypes(rt, rt2)
	k.equalNatArgs(fieldNatArgs, fieldNatArgs2)

	var fieldNatArgs3 []ActualNatArg
	for _, param := range natParams3 {
		fieldNatArgs3 = append(fieldNatArgs3, ActualNatArg{
			name: param,
		})
	}
	k.equalTypes(rt, resolvedType2.BracketType.ArrayType)
	k.equalNatArgs(fieldNatArgs, fieldNatArgs3)

	// log.Printf("resolveTypeTL2 of vector for %s field: %s -> %s", canonicalName, fieldT, rt.String())
	fieldIns, fieldBare, err := k.getInstanceTL1(rt, rt2, true)
	if err != nil {
		return nil, fmt.Errorf("fail to instantiate type of vector %s field: %w", canonicalName, err)
	}

	ins := &TypeInstanceArray{
		TypeInstanceCommon: TypeInstanceCommon{
			canonicalName: canonicalName,
			natParams:     natParams,
			tip:           tip,
			rt:            resolvedType,
			rt2:           resolvedType2,
			argNamespace:  k.getArgNamespace(resolvedType),
			argNamespace2: k.getArgNamespace2(resolvedType2),
		},
		isTuple: false,
	}
	if ins.argNamespace != ins.argNamespace2 {
		panic("internal error getArgNamespace2")
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
	resolvedType tlast.TypeRef, resolvedType2 tlast.TL2TypeRef,
	leftArgs []tlast.TemplateArgument) (TypeInstance, error) {

	// TODO - do not derive parameters from getTL1ArgsHybrid, call fillNatParamHybrid directly
	localArgs, natParams := k.getTL1Args(leftArgs, resolvedType.Args)
	localArgs2, natParams2 := k.getTL1ArgsHybrid(tip.templateArguments, resolvedType2)
	_, natParams3 := k.getTL1ArgHybrid(tlast.TL2TypeArgument{Type: resolvedType2.BracketType.ArrayType}, "t")
	if a, b := strings.Join(natParams, ","), strings.Join(natParams2, ","); a != b || len(localArgs) != len(localArgs2) {
		panic(fmt.Errorf("!equalNatParams %s %s", a, b))
	}
	// log.Printf("natParams for tuple %s: %s", canonicalName, strings.Join(natParams, ","))
	//if len(natParams) != 0 {
	//	fmt.Printf("tuple natparams %s\n", strings.Join(natParams, ","))
	//}
	fieldT := tlast.TypeRef{Type: tlast.Name{Name: "t"}}
	rt, fieldNatArgs, err := k.resolveTypeTL1(fieldT, leftArgs, localArgs)
	if err != nil {
		return nil, fmt.Errorf("fail to resolve type of tuple %s field: %w", canonicalName, err)
	}
	rt2, fieldNatArgs2, err := k.resolveTypeHybrid(false, k.convertTypeRef(fieldT), leftArgs, localArgs2)
	if err != nil {
		return nil, err
	}
	k.equalTypes(rt, rt2)
	k.equalNatArgs(fieldNatArgs, fieldNatArgs2)

	var fieldNatArgs3 []ActualNatArg
	for _, param := range natParams3 {
		//if i == 0 && !resolvedType2.BracketType.IndexType.IsNumber {
		//	continue
		//}
		fieldNatArgs3 = append(fieldNatArgs3, ActualNatArg{
			name: param,
		})
	}
	k.equalTypes(rt, resolvedType2.BracketType.ArrayType)
	k.equalNatArgs(fieldNatArgs, fieldNatArgs3)

	// log.Printf("resolveTypeTL2 of tuple for %s field: %s -> %s", canonicalName, fieldT, rt.String())
	fieldIns, fieldBare, err := k.getInstanceTL1(rt, rt2, true)
	if err != nil {
		return nil, fmt.Errorf("fail to instantiate type of tuple %s field: %w", canonicalName, err)
	}

	if resolvedType2.BracketType.IndexType.IsNumber != resolvedType.Args[0].IsArith ||
		resolvedType2.BracketType.IndexType.Number != resolvedType.Args[0].Arith.Res {
		panic("tuple properties differ")
	}

	ins := &TypeInstanceArray{
		TypeInstanceCommon: TypeInstanceCommon{
			canonicalName: canonicalName,
			natParams:     natParams,
			tip:           tip,
			rt:            resolvedType,
			rt2:           resolvedType2,
			argNamespace:  k.getArgNamespace(resolvedType),
			argNamespace2: k.getArgNamespace2(resolvedType2),
		},
		isTuple:     true,
		count:       resolvedType2.BracketType.IndexType.Number,
		dynamicSize: !resolvedType2.BracketType.IndexType.IsNumber,
	}
	if ins.argNamespace != ins.argNamespace2 {
		panic("internal error getArgNamespace2")
	}
	ins.field = Field{
		owner:   ins,
		ins:     fieldIns,
		bare:    fieldBare,
		natArgs: fieldNatArgs,
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
		originTL2:         false,
		builtin:           true,
		combTL1:           []*tlast.Combinator{combTL1},
		instances:         map[string]*TypeInstanceRef{},
		tl1Names:          map[string]struct{}{"__dict": {}},
		tl2Names:          map[string]struct{}{},
		canonicalName:     tlast.Name{Name: "__dict"},
		historicalName:    tlast.Name{Name: "BuiltinDict"},
		canBeBare:         true,
		templateArguments: k.convertTemplateArguments(combTL1.TemplateArguments),
		targs:             make([]KernelTypeTarg, 1),
	}
	if err := k.addTip(kt, "__dict", ""); err != nil {
		panic(fmt.Sprintf("error adding __dict: %v", err))
	}
	combTL1 = &tlast.Combinator{
		Construct: tlast.Constructor{
			Name: tlast.Name{Name: "__dict2"},
		},
		TemplateArguments: []tlast.TemplateArgument{{FieldName: "k"}, {FieldName: "v"}},
	}
	kt = &KernelType{
		originTL2:         false,
		builtin:           true,
		combTL1:           []*tlast.Combinator{combTL1},
		instances:         map[string]*TypeInstanceRef{},
		tl1Names:          map[string]struct{}{"__dict2": {}},
		tl2Names:          map[string]struct{}{},
		canonicalName:     tlast.Name{Name: "__dict2"},
		historicalName:    tlast.Name{Name: "BuiltinDict2"},
		canBeBare:         true,
		templateArguments: k.convertTemplateArguments(combTL1.TemplateArguments),
		targs:             make([]KernelTypeTarg, 2),
	}
	if err := k.addTip(kt, "__dict2", ""); err != nil {
		panic(fmt.Sprintf("error adding __dict2: %v", err))
	}
	combTL1 = &tlast.Combinator{
		Construct: tlast.Constructor{
			Name: tlast.Name{Name: "__dict_field"},
		},
		TemplateArguments: []tlast.TemplateArgument{{FieldName: "k", IsNat: false}, {FieldName: "v", IsNat: false}},
		Fields: []tlast.Field{{
			FieldName: "key",
			FieldType: tlast.TypeRef{Type: tlast.Name{Name: "k"}},
		}, {
			FieldName: "value",
			FieldType: tlast.TypeRef{Type: tlast.Name{Name: "v"}},
		}},
	}
	kt = &KernelType{
		originTL2:         false,
		builtin:           true,
		combTL1:           []*tlast.Combinator{combTL1},
		instances:         map[string]*TypeInstanceRef{},
		tl1Names:          map[string]struct{}{"__dict_field": {}},
		tl2Names:          map[string]struct{}{},
		canonicalName:     tlast.Name{Name: "__dict_field"},
		historicalName:    tlast.Name{Name: "DictField"},
		canBeBare:         true,
		templateArguments: k.convertTemplateArguments(combTL1.TemplateArguments),
		targs:             make([]KernelTypeTarg, 2),
	}
	if err := k.addTip(kt, "__dict_field", ""); err != nil {
		panic(fmt.Sprintf("error adding __dict_field: %v", err))
	}
	combTL1 = &tlast.Combinator{
		Construct: tlast.Constructor{
			Name: tlast.Name{Name: "__vector"},
		},
		TemplateArguments: []tlast.TemplateArgument{{FieldName: "t", IsNat: false}},
	}
	kt = &KernelType{
		originTL2:         false,
		builtin:           true,
		combTL1:           []*tlast.Combinator{combTL1},
		instances:         map[string]*TypeInstanceRef{},
		tl1Names:          map[string]struct{}{"__vector": {}},
		tl2Names:          map[string]struct{}{},
		canonicalName:     tlast.Name{Name: "__vector"},
		historicalName:    tlast.Name{Name: "BuiltinVector"},
		canBeBare:         true,
		templateArguments: k.convertTemplateArguments(combTL1.TemplateArguments),
		targs:             make([]KernelTypeTarg, 1),
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
		originTL2:         false,
		builtin:           true,
		combTL1:           []*tlast.Combinator{combTL1},
		instances:         map[string]*TypeInstanceRef{},
		tl1Names:          map[string]struct{}{"__tuple": {}},
		tl2Names:          map[string]struct{}{},
		canonicalName:     tlast.Name{Name: "__tuple"},
		historicalName:    tlast.Name{Name: "BuiltinTuple"},
		canBeBare:         true,
		templateArguments: k.convertTemplateArguments(combTL1.TemplateArguments),
		targs:             make([]KernelTypeTarg, 2),
	}
	if err := k.addTip(kt, "__tuple", ""); err != nil {
		panic(fmt.Sprintf("error adding __tuple: %v", err))
	}
}
