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

func (k *Kernel) createVectorTL1(canonicalName string,
	resolvedType2 tlast.TL2TypeRef) (TypeInstance, error) {

	_, natParams3 := k.getTL1ArgHybrid(tlast.TL2TypeArgument{Type: resolvedType2.BracketType.ArrayType}, "t")

	var fieldNatArgs3 []ActualNatArg
	for _, param := range natParams3 {
		fieldNatArgs3 = append(fieldNatArgs3, ActualNatArg{
			name: param,
		})
	}

	// log.Printf("resolveTypeTL2 of vector for %s field: %s -> %s", canonicalName, fieldT, rt.String())
	fieldIns, fieldBare, err := k.getInstanceTL1(resolvedType2.BracketType.ArrayType, true)
	if err != nil {
		return nil, fmt.Errorf("fail to instantiate type of vector %s field: %w", canonicalName, err)
	}

	ins := &TypeInstanceArray{
		TypeInstanceCommon: TypeInstanceCommon{
			canonicalName: canonicalName,
			natParams:     natParams3,
			tip:           nil, // TODO - try to live without brackets type at all
			rt2:           resolvedType2,
			argNamespace:  k.getArgNamespace2(resolvedType2),
		},
		isTuple: false,
	}

	ins.field = Field{
		owner:   ins,
		ins:     fieldIns,
		bare:    fieldBare,
		natArgs: fieldNatArgs3,
	}
	return ins, nil
}

func (k *Kernel) createTupleTL1(canonicalName string, resolvedType2 tlast.TL2TypeRef) (TypeInstance, error) {

	_, natParams3 := k.getTL1ArgHybrid(tlast.TL2TypeArgument{Type: resolvedType2.BracketType.ArrayType}, "t")
	// log.Printf("natParams for tuple %s: %s", canonicalName, strings.Join(natParams, ","))
	//if len(natParams) != 0 {
	//	fmt.Printf("tuple natparams %s\n", strings.Join(natParams, ","))
	//}
	var fieldNatArgs3 []ActualNatArg
	for _, param := range natParams3 {
		//if i == 0 && !resolvedType2.BracketType.IndexType.IsNumber {
		//	continue
		//}
		fieldNatArgs3 = append(fieldNatArgs3, ActualNatArg{
			name: param,
		})
	}
	if !resolvedType2.BracketType.IndexType.IsNumber {
		natParams3 = append([]string{"n"}, natParams3...)
	}

	// log.Printf("resolveTypeTL2 of tuple for %s field: %s -> %s", canonicalName, fieldT, rt.String())
	fieldIns, fieldBare, err := k.getInstanceTL1(resolvedType2.BracketType.ArrayType, true)
	if err != nil {
		return nil, fmt.Errorf("fail to instantiate type of tuple %s field: %w", canonicalName, err)
	}

	ins := &TypeInstanceArray{
		TypeInstanceCommon: TypeInstanceCommon{
			canonicalName: canonicalName,
			natParams:     natParams3,
			tip:           nil, // TODO - try to live without brackets type at all
			rt2:           resolvedType2,
			argNamespace:  k.getArgNamespace2(resolvedType2),
		},
		isTuple:     true,
		count:       resolvedType2.BracketType.IndexType.Number,
		dynamicSize: !resolvedType2.BracketType.IndexType.IsNumber,
	}
	ins.field = Field{
		owner:   ins,
		ins:     fieldIns,
		bare:    fieldBare,
		natArgs: fieldNatArgs3,
	}
	return ins, nil
}

func (k *Kernel) addTL1Brackets() {
	// for the purpose of type check, this is object with no fields, like __dict {t:Type} = ;
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
		canonicalName:     tlast.TL2TypeName{Name: "__dict"},
		historicalName:    tlast.TL2TypeName{Name: "BuiltinDict"},
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
		canonicalName:     tlast.TL2TypeName{Name: "__dict2"},
		historicalName:    tlast.TL2TypeName{Name: "BuiltinDict2"},
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
		canonicalName:     tlast.TL2TypeName{Name: "__dict_field"},
		historicalName:    tlast.TL2TypeName{Name: "DictField"},
		canBeBare:         true,
		templateArguments: k.convertTemplateArguments(combTL1.TemplateArguments),
		targs:             make([]KernelTypeTarg, 2),
	}
	if err := k.addTip(kt, "__dict_field", ""); err != nil {
		panic(fmt.Sprintf("error adding __dict_field: %v", err))
	}
}
