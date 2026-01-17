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

type ActualNatArg struct {
	isNumber   bool
	number     uint32
	isField    bool // otherwise it is # param with name
	fieldIndex int
	name       string // param name
}

func (arg *ActualNatArg) IsNumber() bool {
	return arg.isNumber
}

func (arg *ActualNatArg) Number() uint32 {
	return arg.number
}

func (arg *ActualNatArg) IsField() bool {
	return arg.isField
}

func (arg *ActualNatArg) FieldIndex() int {
	return arg.fieldIndex
}

func (arg *ActualNatArg) Name() string {
	return arg.name
}

type Field struct {
	name string
	ins  *TypeInstanceRef

	bare bool // for TL1 only, false for TL2
	//recursive bool

	fieldMask *ActualNatArg
	bitNumber uint32 // only used when fieldMask != nil

	natArgs []ActualNatArg // for TL1 only, empty for TL2
	//rt      tlast.TypeRef  // for TL1 only, empty for TL2
}

func (f *Field) Bare() bool                 { return f.bare }
func (f *Field) Name() string               { return f.name }
func (f *Field) TypeInstance() TypeInstance { return f.ins.ins }
func (f *Field) FieldMask() *ActualNatArg   { return f.fieldMask }
func (f *Field) BitNumber() uint32          { return f.bitNumber }
func (f *Field) NatArgs() []ActualNatArg    { return f.natArgs }

type TypeInstanceStruct struct {
	TypeInstanceCommon
	isConstructorFields bool
	fields              []Field
	isUnionElement      bool
	unionIndex          int

	// if function
	resultType TypeInstance
}

func (ins *TypeInstanceStruct) Fields() []Field {
	return ins.fields
}

func (ins *TypeInstanceStruct) FindCycle(c *cycleFinder) {
	if !c.push(ins) {
		return
	}
	defer c.pop(ins)
	for _, ft := range ins.fields {
		if ft.fieldMask == nil {
			ft.ins.ins.FindCycle(c)
		}
	}
}

func (ins *TypeInstanceStruct) CreateValue() KernelValue {
	v := ins.CreateValueObject()
	return &v
}

func (ins *TypeInstanceStruct) SkipTL2(r []byte) ([]byte, error) {
	return basictl.SkipSizedValue(r)
}

func (ins *TypeInstanceStruct) CreateValueObject() KernelValueStruct {
	value := KernelValueStruct{
		instance: ins,
		fields:   make([]KernelValue, len(ins.fields)),
	}
	for i, ft := range ins.fields {
		if ft.fieldMask == nil {
			value.fields[i] = ft.ins.ins.CreateValue()
		}
	}
	return value
}

func (k *Kernel) createStruct(canonicalName string, tip *KernelType,
	isConstructorFields bool, alias tlast.TL2TypeRef, constructorFields []tlast.TL2Field,
	leftArgs []tlast.TL2TypeTemplate, actualArgs []tlast.TL2TypeArgument,
	isUnionElement bool, unionIndex int, resultType TypeInstance) (*TypeInstanceStruct, error) {

	ins := &TypeInstanceStruct{
		TypeInstanceCommon: TypeInstanceCommon{
			canonicalName: canonicalName,
			tip:           tip,
		},
		isConstructorFields: isConstructorFields,
		isUnionElement:      isUnionElement,
		unionIndex:          unionIndex,
		resultType:          resultType,
	}
	if !isConstructorFields { // if we are here, this is union variant or function result, where alias is field 1
		constructorFields = append(constructorFields, tlast.TL2Field{Type: alias})
	}

	for _, fieldDef := range constructorFields {
		rt, err := k.resolveType(fieldDef.Type, leftArgs, actualArgs)
		if err != nil {
			return nil, fmt.Errorf("fail to resolve type of object %s field %s: %w", canonicalName, fieldDef.Name, err)
		}
		fieldIns, err := k.GetInstance(rt)
		if err != nil {
			return nil, fmt.Errorf("fail to instantiate type of object %s field %s: %w", canonicalName, fieldDef.Name, err)
		}
		var fieldMask *ActualNatArg
		if fieldDef.IsOptional {
			fieldMask = &ActualNatArg{} // TODO - mark as TL2
		}
		ins.fields = append(ins.fields, Field{
			name:      fieldDef.Name,
			ins:       fieldIns,
			fieldMask: fieldMask,
		})
	}
	return ins, nil
}

func (k *Kernel) createStructTL1FromTL2(canonicalName string,
	constructorFields []tlast.Field,
	leftArgs []tlast.TemplateArgument, actualArgs []tlast.TL2TypeArgument,
	isUnionElement bool, unionIndex int, resultType TypeInstance) (*TypeInstanceStruct, error) {

	return nil, fmt.Errorf("TODO - not implemented yet")
	//ins := &TypeInstanceStruct{
	//	TypeInstanceCommon: TypeInstanceCommon{
	//		canonicalName: canonicalName,
	//	},
	//	isConstructorFields: false,
	//	isUnionElement:      isUnionElement,
	//	unionIndex:          unionIndex,
	//	resultType:          resultType,
	//}
	//
	//for _, fieldDef := range constructorFields {
	//	rt, err := k.resolveType(fieldDef., leftArgs, actualArgs)
	//	if err != nil {
	//		return nil, fmt.Errorf("fail to resolve type of object %s field %s: %w", canonicalName, fieldDef.Name, err)
	//	}
	//	fieldIns, err := k.GetInstance(rt)
	//	if err != nil {
	//		return nil, fmt.Errorf("fail to instantiate type of object %s field %s: %w", canonicalName, fieldDef.Name, err)
	//	}
	//	var fieldMask *ActualNatArg
	//	if fieldDef.IsOptional {
	//		fieldMask = &ActualNatArg{} // TODO - mark as TL2
	//	}
	//	ins.fields = append(ins.fields, Field{
	//		ins:       fieldIns,
	//		fieldMask: fieldMask,
	//	})
	//}
	//return ins, nil
}

// we want the same naming convention for nat params, as in old kernel,
// though it has no difference to semantic and can be simplified to p0, p1, p2, etc.
func (k *Kernel) fillNatParam(rt tlast.ArithmeticOrType, natParams *[]string, prefix string) {
	if rt.IsArith {
		return
	}
	if rt.T.String() == "*" {
		*natParams = append(*natParams, prefix)
		return
	}
	tip, ok := k.tips[rt.T.Type.String()]
	if !ok {
		panic("resolved type not found in global type map")
	}
	for i, arg := range rt.T.Args {
		leftArg := tip.combTL1[0].TemplateArguments[i]
		k.fillNatParam(arg, natParams, prefix+leftArg.FieldName)
	}
}

func (k *Kernel) getTL1Args(leftArgs []tlast.TemplateArgument, actualArgs []tlast.ArithmeticOrType) (localArgs []LocalArg, natParams []string) {
	for i, arg := range actualArgs {
		leftArg := leftArgs[i]
		var localNatParams []string
		k.fillNatParam(arg, &localNatParams, leftArg.FieldName)
		if len(localNatParams) == 1 {
			localNatParams[0] = leftArg.FieldName
		}
		natParams = append(natParams, localNatParams...)
		localArg := LocalArg{
			wrongTypeErr: nil,
			arg:          arg,
		}
		for _, param := range localNatParams {
			localArg.natArgs = append(localArg.natArgs, ActualNatArg{
				name: param,
			})
		}
		localArgs = append(localArgs, localArg)
	}
	return
}

func (k *Kernel) createStructTL1FromTL1(canonicalName string, tip *KernelType,
	resolvedType tlast.TypeRef, constructorFields []tlast.Field,
	leftArgs []tlast.TemplateArgument, actualArgs []tlast.ArithmeticOrType,
	isUnionElement bool, unionIndex int, resultType TypeInstance) (*TypeInstanceStruct, error) {

	localArgs, natParams := k.getTL1Args(leftArgs, actualArgs)
	log.Printf("natParams for %s: %s", canonicalName, strings.Join(natParams, ","))

	ins := &TypeInstanceStruct{
		TypeInstanceCommon: TypeInstanceCommon{
			canonicalName: canonicalName,
			NatParams:     natParams,
			tip:           tip,
			rt:            resolvedType,
		},
		isConstructorFields: true,
		isUnionElement:      isUnionElement,
		unionIndex:          unionIndex,
		resultType:          resultType,
	}
	for i, fieldDef := range constructorFields {
		rt, natArgs, err := k.resolveTypeTL1(fieldDef.FieldType, leftArgs, localArgs)
		if err != nil {
			return nil, fmt.Errorf("fail to resolve type of object %s field %s: %w", canonicalName, fieldDef.FieldName, err)
		}
		log.Printf("resolveType for %s field %s: %s -> %s", canonicalName, fieldDef.FieldName, fieldDef.FieldType.String(), rt.String())
		fieldIns, err := k.getInstanceTL1(rt, true)
		if err != nil {
			return nil, fmt.Errorf("fail to instantiate type of object %s field %s: %w", canonicalName, fieldDef.FieldName, err)
		}
		var fieldMask *ActualNatArg
		//if fieldDef.IsOptional {
		//	fieldMask = &ActualNatArg{} // TODO - mark as TL2
		//}
		ins.fields = append(ins.fields, Field{
			name:      fieldDef.FieldName,
			ins:       fieldIns,
			fieldMask: fieldMask,
			natArgs:   natArgs,
			bare:      rt.Bare,
		})
		if fieldDef.FieldName != "" {
			leftArgs = append(leftArgs, tlast.TemplateArgument{
				FieldName: fieldDef.FieldName,
				IsNat:     true,
				PR:        fieldDef.PR,
			})
			if fieldDef.FieldType.String() != "#" {
				localArgs = append(localArgs, LocalArg{
					wrongTypeErr: fmt.Errorf("only reference to field with type # is allowed"),
				})
			} else {
				localArgs = append(localArgs, LocalArg{
					wrongTypeErr: nil,
					arg:          tlast.ArithmeticOrType{T: tlast.TypeRef{Type: tlast.Name{Name: "*"}}},
					natArgs: []ActualNatArg{{
						isField:    true,
						fieldIndex: i,
					}},
				})
			}
		}
	}
	return ins, nil
}
