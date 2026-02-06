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
	isTuple  bool
	count    uint32
	elemType *TypeInstanceRef

	elemBare    bool           // for TL1 only, false for TL2
	elemNatArgs []ActualNatArg // for TL1 only, empty for TL2
	dynamicSize bool           // for TL1 only, false for TL2
}

func (ins *TypeInstanceArray) IsTuple() bool               { return ins.isTuple }
func (ins *TypeInstanceArray) ElemType() TypeInstance      { return ins.elemType.ins }
func (ins *TypeInstanceArray) ElemBare() bool              { return ins.elemBare }
func (ins *TypeInstanceArray) ElemNatArgs() []ActualNatArg { return ins.elemNatArgs }
func (ins *TypeInstanceArray) DynamicSize() bool           { return ins.dynamicSize }
func (ins *TypeInstanceArray) Count() uint32               { return ins.count }

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

func (k *Kernel) createArray(canonicalName string, isTuple bool, count uint32, elemType *TypeInstanceRef) TypeInstance {
	if elemType.ins.IsBit() {
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
		elemType: elemType,
	}
	return ins
}

func (k *Kernel) createVectorTL1(canonicalName string,
	resolvedType tlast.TypeRef,
	leftArgs []tlast.TemplateArgument, actualArgs []tlast.ArithmeticOrType) (TypeInstance, error) {

	localArgs, natParams := k.getTL1Args(leftArgs, actualArgs)
	// log.Printf("natParams for vector %s: %s", canonicalName, strings.Join(natParams, ","))

	elementT := tlast.TypeRef{Type: tlast.Name{Name: "t"}}

	rt, elemNatArgs, err := k.resolveTypeTL1(elementT, leftArgs, localArgs)
	if err != nil {
		return nil, fmt.Errorf("fail to resolve type of vector %s element: %w", canonicalName, err)
	}
	// log.Printf("resolveType of vector for %s element: %s -> %s", canonicalName, elementT, rt.String())
	elemIns, elemBare, err := k.getInstanceTL1(rt, true, false)
	if err != nil {
		return nil, fmt.Errorf("fail to instantiate type of vector %s element: %w", canonicalName, err)
	}

	ins := &TypeInstanceArray{
		TypeInstanceCommon: TypeInstanceCommon{
			canonicalName: canonicalName,
			NatParams:     natParams,
			rt:            resolvedType,
			argNamespace:  k.getArgNamespace(resolvedType),
		},
		isTuple:     false,
		elemType:    elemIns,
		elemBare:    elemBare,
		elemNatArgs: elemNatArgs,
	}
	return ins, nil
}

func (k *Kernel) createTupleTL1(canonicalName string,
	resolvedType tlast.TypeRef,
	leftArgs []tlast.TemplateArgument, actualArgs []tlast.ArithmeticOrType) (TypeInstance, error) {

	localArgs, natParams := k.getTL1Args(leftArgs, actualArgs)
	// log.Printf("natParams for tuple %s: %s", canonicalName, strings.Join(natParams, ","))

	elementT := tlast.TypeRef{Type: tlast.Name{Name: "t"}}

	rt, natArgs, err := k.resolveTypeTL1(elementT, leftArgs, localArgs)
	if err != nil {
		return nil, fmt.Errorf("fail to resolve type of tuple %s element: %w", canonicalName, err)
	}
	// log.Printf("resolveType of tuple for %s element: %s -> %s", canonicalName, elementT, rt.String())
	fieldIns, fieldBare, err := k.getInstanceTL1(rt, true, false)
	if err != nil {
		return nil, fmt.Errorf("fail to instantiate type of tuple %s element: %w", canonicalName, err)
	}

	if sf := actualArgs[0].SourceField; sf != (tlast.CombinatorField{}) {
		field := &sf.Comb.Fields[sf.FieldIndex]
		if field.UsedAsMask {
			e3 := field.UsedAsMaskPR.BeautifulError(fmt.Errorf("used as mask here"))
			e3.PrintWarning(k.opts.ErrorWriter, nil)
			e1 := field.PRName.BeautifulError(fmt.Errorf("#-field %s is used as tuple size, while already being used as a field mask", field.FieldName))
			e2 := actualArgs[0].T.PR.BeautifulError(fmt.Errorf("used as size here"))
			return nil, tlast.BeautifulError2(e1, e2)
		}
		field.UsedAsSize = true
		field.UsedAsSizePR = actualArgs[0].T.PR
	}
	ins := &TypeInstanceArray{
		TypeInstanceCommon: TypeInstanceCommon{
			canonicalName: canonicalName,
			NatParams:     natParams,
			rt:            resolvedType,
			argNamespace:  k.getArgNamespace(resolvedType),
		},
		isTuple:     true,
		elemType:    fieldIns,
		elemBare:    fieldBare,
		elemNatArgs: natArgs,
		dynamicSize: !actualArgs[0].IsArith,
		count:       actualArgs[0].Arith.Res,
	}
	return ins, nil
}

func (k *Kernel) addTL1Brackets() {
	// for the purpose of type check, this is object with no fields, like __vector {t:Type} = ;
	combTL1 := &tlast.Combinator{
		Construct: tlast.Constructor{
			Name: tlast.Name{Name: "__vector"},
		},
		TemplateArguments: []tlast.TemplateArgument{{FieldName: "t", IsNat: false}},
	}
	kt := &KernelType{
		originTL2:     false,
		builtin:       true,
		combTL1:       []*tlast.Combinator{combTL1},
		instances:     map[string]*TypeInstanceRef{},
		tl1Names:      map[string]struct{}{"__vector": {}},
		tl2Names:      map[string]struct{}{},
		canonicalName: tlast.Name{Name: "__vector"},
		tl1name:       "__vector",
		canBeBare:     true,
		targs:         make([]KernelTypeTarg, 1),
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
		originTL2:     false,
		builtin:       true,
		combTL1:       []*tlast.Combinator{combTL1},
		instances:     map[string]*TypeInstanceRef{},
		tl1Names:      map[string]struct{}{"__tuple": {}},
		tl2Names:      map[string]struct{}{},
		canonicalName: tlast.Name{Name: "__tuple"},
		tl1name:       "__tuple",
		canBeBare:     true,
		targs:         make([]KernelTypeTarg, 2),
	}
	if err := k.addTip(kt, "__tuple", ""); err != nil {
		panic(fmt.Sprintf("error adding __tuple: %v", err))
	}
}
