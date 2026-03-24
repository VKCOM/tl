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

func (ins *TypeInstanceArray) FindCycle(c *cycleFinder, prName tlast.PositionRange) {
	if !c.push(ins, prName) {
		return
	}
	defer c.pop(ins)
	if ins.isTuple {
		ins.field.ins.ins.FindCycle(c, ins.field.pr)
	}
}

func (ins *TypeInstanceArray) GetChildren(children []TypeInstance, withReturnType bool) []TypeInstance {
	return append(children, ins.field.ins.ins)
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
	resolvedType tlast.TL2TypeRef) (TypeInstance, error) {

	_, natParams := k.fillLocalArg(tlast.TL2TypeArgument{Type: resolvedType.BracketType.ArrayType}, "t", nil)

	fieldNatArgs := k.natParamsToActualNatArgs(natParams)

	// fmt.Printf("resolveTypeTL2 of vector for %s field: %s -> %s\n", canonicalName, fieldT, rt.String())
	fieldIns, fieldBare, err := k.getInstance(resolvedType.BracketType.ArrayType, true)
	if err != nil {
		return nil, fmt.Errorf("fail to instantiate type of vector %s field: %w", canonicalName, err)
	}

	ins := &TypeInstanceArray{
		TypeInstanceCommon: TypeInstanceCommon{
			canonicalName: canonicalName,
			natParams:     natParams,
			tip:           nil, // TODO - try to live without brackets type at all
			resolvedType:  resolvedType,
			hasFetcher:    k.resolvedTypeNeedsFetcher(resolvedType),
			argNamespace:  k.getArgNamespace(resolvedType),
		},
		isTuple: false,
	}

	ins.field = Field{
		owner:   ins,
		ins:     fieldIns,
		bare:    fieldBare,
		natArgs: fieldNatArgs,
		pr:      resolvedType.BracketType.ArrayType.PR,
	}
	return ins, nil
}

func (k *Kernel) createTupleTL1(canonicalName string, resolvedType tlast.TL2TypeRef) (TypeInstance, error) {

	_, natParams := k.fillLocalArg(tlast.TL2TypeArgument{Type: resolvedType.BracketType.ArrayType}, "t", nil)
	// fmt.Printf("natParams for tuple %s: %s\n", canonicalName, strings.Join(natParams, ","))
	fieldNatArgs := k.natParamsToActualNatArgs(natParams)

	if !resolvedType.BracketType.IndexType.IsNumber {
		natParams = append([]string{"n"}, natParams...)
		for i := range fieldNatArgs {
			fieldNatArgs[i].index++
		}
	}

	// fmr.Printf("resolveTypeTL2 of tuple for %s field: %s -> %s\n", canonicalName, fieldT, rt.String())
	fieldIns, fieldBare, err := k.getInstance(resolvedType.BracketType.ArrayType, true)
	if err != nil {
		return nil, fmt.Errorf("fail to instantiate type of tuple %s field: %w", canonicalName, err)
	}

	ins := &TypeInstanceArray{
		TypeInstanceCommon: TypeInstanceCommon{
			canonicalName: canonicalName,
			natParams:     natParams,
			tip:           nil, // TODO - try to live without brackets type at all
			resolvedType:  resolvedType,
			hasFetcher:    k.resolvedTypeNeedsFetcher(resolvedType),
			argNamespace:  k.getArgNamespace(resolvedType),
		},
		isTuple:     true,
		count:       resolvedType.BracketType.IndexType.Number,
		dynamicSize: !resolvedType.BracketType.IndexType.IsNumber,
	}
	ins.field = Field{
		owner:   ins,
		ins:     fieldIns,
		bare:    fieldBare,
		natArgs: fieldNatArgs,
		pr:      resolvedType.BracketType.ArrayType.PR,
	}
	return ins, nil
}

func (k *Kernel) addTL1Brackets() {
	// TODO - parse primitive definitions from local TL so Beautiful Errors work correctly for int<int>, etc.
	str := "__dict_field {k:Type} {v:Type} key:k value:v = DictField;"
	combs, err := tlast.ParseTLFile(str, "builtin.tl", tlast.LexerOptions{AllowBuiltin: true, AllowDirty: true})
	if err != nil || len(combs) != 1 {
		panic("error adding built in types")
	}
	combTL1 := combs[0]
	combTL1.Construct.ID = 0
	combTL1.TypeDecl.Name = tlast.Name{}
	kt := &KernelType{
		originTL2:         false,
		builtin:           true,
		combTL1:           []*tlast.Combinator{combTL1},
		instances:         map[string]*TypeInstanceRef{},
		tl1Names:          map[string]struct{}{"__dict_field": {}},
		tl2Names:          map[string]struct{}{},
		canonicalName:     tlast.TL2TypeName{Name: "__dict_field"},
		canBeBare:         true,
		templateArguments: k.convertTemplateArguments(combTL1.TemplateArguments),
		targs:             make([]KernelTypeTarg, 2),
	}
	if err := k.addTip(kt, "__dict_field", ""); err != nil {
		panic(fmt.Sprintf("error adding __dict_field: %v", err))
	}
}
