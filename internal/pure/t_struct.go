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

	maskTL2Bit *int

	natArgs []ActualNatArg // for TL1 only, empty for TL2
	//rt      tlast.TypeRef  // for TL1 only, empty for TL2
}

func (f *Field) Bare() bool                 { return f.bare }
func (f *Field) Name() string               { return f.name }
func (f *Field) TypeInstance() TypeInstance { return f.ins.ins }
func (f *Field) FieldMask() *ActualNatArg   { return f.fieldMask }
func (f *Field) BitNumber() uint32          { return f.bitNumber }
func (f *Field) MaskTL2Bit() *int           { return f.maskTL2Bit }
func (f *Field) NatArgs() []ActualNatArg    { return f.natArgs }

type TypeInstanceStruct struct {
	TypeInstanceCommon
	isConstructorFields bool
	fields              []Field
	isUnionElement      bool
	unionIndex          int

	// if function
	resultType    TypeInstance
	resultNatArgs []ActualNatArg // for TL1 only, empty for TL2
}

func (ins *TypeInstanceStruct) Fields() []Field {
	return ins.fields
}

func (ins *TypeInstanceStruct) ResultType() TypeInstance {
	return ins.resultType
}

func (ins *TypeInstanceStruct) ResultNatArgs() []ActualNatArg {
	return ins.resultNatArgs
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
	tName := rt.T.Type.String()
	tip, ok := k.tips[tName]
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

func (k *Kernel) canonicalBrackets(fieldDef tlast.Field) error {
	if !fieldDef.IsRepeated {
		return nil // always canonical
	}
	if len(fieldDef.ScaleRepeat.Rep) != 1 {
		return fieldDef.ScaleRepeat.PR.BeautifulError(fmt.Errorf("brackets must contain single type"))
	}
	f := fieldDef.ScaleRepeat.Rep[0]
	if f.IsRepeated || f.FieldName != "" {
		return f.PR.BeautifulError(fmt.Errorf("brackets field should not be named or contain brackets"))
	}
	if f.Mask != nil && f.Excl {
		return f.PR.BeautifulError(fmt.Errorf("brackets field should not contain fieldsmask or exclamation"))
	}
	return nil
}

func (k *Kernel) createStructTL1FromTL1(canonicalName string, tip *KernelType,
	resolvedType tlast.TypeRef, def *tlast.Combinator,
	leftArgs []tlast.TemplateArgument, actualArgs []tlast.ArithmeticOrType,
	isUnionElement bool, unionIndex int) (*TypeInstanceStruct, error) {

	constructorFields := def.Fields

	localArgs, natParams := k.getTL1Args(leftArgs, actualArgs)
	log.Printf("natParams for %s: %s", canonicalName, strings.Join(natParams, ","))

	ins := &TypeInstanceStruct{
		TypeInstanceCommon: TypeInstanceCommon{
			canonicalName: canonicalName,
			NatParams:     natParams,
			tip:           tip,
			rt:            resolvedType,
			argNamespace:  k.getArgNamespace(resolvedType),
		},
		isConstructorFields: true,
		isUnionElement:      isUnionElement,
		unionIndex:          unionIndex,
	}
	nextTL2MaskBit := 0
	var fieldsAfterReplace []tlast.Field
	for i := 0; i < len(constructorFields); i++ {
		fieldDef := constructorFields[i]
		if fieldDef.FieldType.String() == "#" && fieldDef.FieldName == "" && i+1 < len(constructorFields) {
			nextFieldDef := constructorFields[i+1]
			if nextFieldDef.Mask != nil || !nextFieldDef.IsRepeated || nextFieldDef.ScaleRepeat.ExplicitScale {
				return nil, fieldDef.PR.BeautifulError(fmt.Errorf("anonymous # field must be followed by brackets with no fieldmask and no explicit scale repeat (# [...] or # a:[...])"))
			}
			if err := k.canonicalBrackets(fieldDef); err != nil {
				return nil, err
			}
			// we replace 2 fields with vector
			// hren # a:[int] = Hren;
			i++
			fieldDef = nextFieldDef
			fieldDef.FieldType.Args = []tlast.ArithmeticOrType{{T: fieldDef.ScaleRepeat.Rep[0].FieldType}}
			fieldDef.FieldType.Type = tlast.Name{Name: "__vector"}
			fieldDef.FieldType.Bare = true
		} else if fieldDef.IsRepeated && i == 0 && !fieldDef.ScaleRepeat.ExplicitScale &&
			len(tip.combTL1[0].TemplateArguments) != 0 {
			a := tip.combTL1[0].TemplateArguments[len(tip.combTL1[0].TemplateArguments)-1]
			if !a.IsNat {
				e1 := fieldDef.FieldType.PR.CollapseToBegin().BeautifulError(fmt.Errorf("anonymous scale repeat implicitly references last template parameter %q which should have type #", a.FieldName))
				e2 := a.PR.BeautifulError(fmt.Errorf("see here"))
				return nil, tlast.BeautifulError2(e1, e2)
			}
			if err := k.canonicalBrackets(fieldDef); err != nil {
				return nil, err
			}
			fieldDef.FieldType.Args = []tlast.ArithmeticOrType{{T: tlast.TypeRef{Type: tlast.Name{Name: a.FieldName}}}, {T: fieldDef.ScaleRepeat.Rep[0].FieldType}}
			fieldDef.FieldType.Type = tlast.Name{Name: "__tuple"}
			fieldDef.FieldType.Bare = true
		} else if fieldDef.IsRepeated {
			if err := k.canonicalBrackets(fieldDef); err != nil {
				return nil, err
			}
			if !fieldDef.ScaleRepeat.ExplicitScale {
				prevFieldDef := constructorFields[i-1] // never panics, due to checks above
				if prevFieldDef.FieldType.String() != "#" {
					e1 := fieldDef.FieldType.PR.CollapseToBegin().BeautifulError(fmt.Errorf("anonymous scale repeat implicitly references previous field %q, which should have type #", prevFieldDef.FieldName))
					e2 := prevFieldDef.PR.BeautifulError(fmt.Errorf("see here"))
					return nil, tlast.BeautifulError2(e1, e2)
				}
				fieldDef.ScaleRepeat.Scale = tlast.ScaleFactor{
					IsArith: false,
					Scale:   prevFieldDef.FieldName,
				}
			}
			fieldDef.FieldType.Args = []tlast.ArithmeticOrType{{}, {T: fieldDef.ScaleRepeat.Rep[0].FieldType}}
			fieldDef.FieldType.Type = tlast.Name{Name: "__tuple"}
			fieldDef.FieldType.Bare = true
			if fieldDef.ScaleRepeat.Scale.IsArith {
				fieldDef.FieldType.Args[0] = tlast.ArithmeticOrType{T: tlast.TypeRef{PR: fieldDef.ScaleRepeat.Scale.PR}, IsArith: true, Arith: fieldDef.ScaleRepeat.Scale.Arith}
			} else {
				fieldDef.FieldType.Args[0] = tlast.ArithmeticOrType{T: tlast.TypeRef{PR: fieldDef.ScaleRepeat.Scale.PR, Type: tlast.Name{Name: fieldDef.ScaleRepeat.Scale.Scale}}}
			}
		}
		rt, natArgs, err := k.resolveTypeTL1(fieldDef.FieldType, leftArgs, localArgs)
		if err != nil {
			return nil, err
		}
		log.Printf("resolveType for %s field %s: %s -> %s", canonicalName, fieldDef.FieldName, fieldDef.FieldType.String(), rt.String())
		fieldIns, fieldBare, err := k.getInstanceTL1(rt, true, false)
		if err != nil {
			return nil, err
		}
		newField := Field{
			name:    fieldDef.FieldName,
			ins:     fieldIns,
			natArgs: natArgs,
			bare:    fieldBare,
		}
		if fieldDef.Mask != nil {
			if fieldDef.Mask.BitNumber >= 32 {
				return nil, fieldDef.Mask.PRBits.BeautifulError(fmt.Errorf("bitmask (%d) must be in range [0..31]", fieldDef.Mask.BitNumber))
			}
			fieldMask, err := k.resolveMaskTL1(*fieldDef.Mask, leftArgs, localArgs)
			if err != nil {
				return nil, err
			}
			newField.bitNumber = fieldDef.Mask.BitNumber
			newField.fieldMask = &fieldMask
			maskBit := nextTL2MaskBit
			newField.maskTL2Bit = &maskBit
			nextTL2MaskBit++
		}

		ins.fields = append(ins.fields, newField)
		fieldsAfterReplace = append(fieldsAfterReplace, fieldDef)
		if fieldDef.FieldName != "" {
			leftArgs = append(leftArgs, tlast.TemplateArgument{
				FieldName: fieldDef.FieldName,
				IsNat:     true,
				PR:        fieldDef.PR,
			})
			if fieldDef.FieldType.String() != "#" {
				localArgs = append(localArgs, LocalArg{
					wrongTypeErr: fieldDef.PRName.BeautifulError(fmt.Errorf("defined here")),
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
	for _, f := range fieldsAfterReplace {
		if f.FieldName == "" && (len(ins.fields) != 1 || f.Mask != nil) {
			return nil, f.PR.BeautifulError(fmt.Errorf("anonymous fields are discouraged, except when used in '# a:[int]' pattern or when type has single anonymous field without fieldmask (typedef-like)"))
		}
	}
	if def.IsFunction {
		rt, natArgs, err := k.resolveTypeTL1(def.FuncDecl, leftArgs, localArgs)
		if err != nil {
			return nil, fmt.Errorf("fail to resolve function %s result type: %w", canonicalName, err)
		}
		log.Printf("resolveType for function %s result type: %s -> %s", canonicalName, def.FuncDecl.String(), rt.String())
		fieldIns, fieldBare, err := k.getInstanceTL1(rt, true, false)
		if err != nil {
			return nil, fmt.Errorf("fail to instantiate function %s result type: %w", canonicalName, err)
		}
		if fieldBare {
			// @read a.TypeA = int;
			// @read a.TypeB = %Int;
			return nil, def.FuncDecl.PR.BeautifulError(fmt.Errorf("function %q result cannot be bare", def.Construct.Name.String()))
		}
		ins.resultType = fieldIns.ins
		ins.resultNatArgs = natArgs
	}
	return ins, nil
}
