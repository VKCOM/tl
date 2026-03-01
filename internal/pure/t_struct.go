// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package pure

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/vkcom/tl/internal/purelegacy"
	"github.com/vkcom/tl/internal/tlast"
	"github.com/vkcom/tl/internal/utils"
	"github.com/vkcom/tl/pkg/basictl"
)

type TypeInstanceStruct struct {
	TypeInstanceCommon
	isConstructorFields bool
	fields              []Field
	isUnionElement      bool
	unionIndex          int
	isTypedef           bool
	isAlias             bool
	isUnwrap            bool

	// if function
	resultType    TypeInstance
	resultNatArgs []ActualNatArg // for TL1 only, empty for TL2
	isResultAlias bool           // false for TL1 functions and TL2 functions with single unnamed field
	rpcPreferTL2  bool
}

func (ins *TypeInstanceStruct) Fields() []Field {
	return ins.fields
}

// this struct has empty name of the single field (without fieldsmask)
func (ins *TypeInstanceStruct) IsTypedef() bool {
	return ins.isTypedef
}

// Both TL1-style typedef (single anonymous field)
// And actual TL2-style alias.
// In both cases, there is no TL2 object wrapping around the aliasing type
// So TL2 serialization format is the same. TL1 format will differ if the first field is Boxed.
func (ins *TypeInstanceStruct) IsAlias() bool {
	return ins.isAlias
}

// Where this type is used during generation, we must instad use wrapped type.
// vector<int> is compiled into []int.
func (ins *TypeInstanceStruct) IsUnwrap() bool {
	return ins.isAlias && ins.isUnwrap // isAlias should always be set, if isUnwrap set
}

func (ins *TypeInstanceStruct) ResultType() TypeInstance {
	return ins.resultType
}

func (ins *TypeInstanceStruct) ResultNatArgs() []ActualNatArg {
	return ins.resultNatArgs
}

func (ins *TypeInstanceStruct) IsResultAlias() bool {
	return ins.isResultAlias
}

func (ins *TypeInstanceStruct) RPCPreferTL2() bool {
	return ins.rpcPreferTL2
}

// TODO - check how this works
func (ins *TypeInstanceStruct) GoodForMapKey() bool {
	return ins.isAlias && ins.fields[0].ins.ins.GoodForMapKey()
}

// TODO - check/decide how this should work
//func (ins *TypeInstanceAlias) IsBit() bool {
//	return ins.fieldType.ins.IsBit()
//}

// most generators will need to add !recursive
func (trw *TypeInstanceStruct) IsTypeDef() bool {
	return len(trw.fields) == 1 && trw.fields[0].name == "" && trw.fields[0].fieldMask == nil
}

// same code as in func (w *TypeInstanceCommon) TransformNatArgsToChild
func (ins *TypeInstanceStruct) ReplaceUnwrapArgs(natArgs []string) []string {
	// Caller called outer.Read(   , nat_x, nat_y)
	// outer has func Read(   ,nat_inner_x uint32, nat_inner_y uint32) {
	// which calls for example inner.Read(   , nat_inner_y, nat_inner_y)
	// in other words, outer passes some parameters to inner in some order, with potential repeats.
	// When unwrapping, we do the job of golang compiler, replacing references to outer nat parameters,
	// so that at the calling site outer.Read(   , nat_x, nat_y) is replaced to
	// inner.Read(   , nat_y, nat_y)
	var result []string
outer:
	for _, arg := range ins.fields[0].natArgs {
		if arg.IsNumber() || arg.IsField() {
			panic("cannot replace to child arith or field nat param")
		}
		for i, p := range ins.natParams {
			if p == arg.Name() {
				result = append(result, natArgs[i])
				continue outer
			}
		}
		log.Panicf("internal compiler error, nat parameter %s not found for unwrap type %s", arg.Name(), ins.canonicalName)
	}
	return result
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

func (ins *TypeInstanceStruct) GetChildren(children []TypeInstance, withReturnType bool) []TypeInstance {
	for _, ft := range ins.fields {
		children = append(children, ft.ins.ins)
	}
	if withReturnType && ins.resultType != nil {
		children = append(children, ins.resultType)
	}
	return children
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

func (k *Kernel) createStructTL2(canonicalName string, tip *KernelType, resolvedType tlast.TL2TypeRef,
	tlName tlast.TL2TypeName, tlTag uint32,
	isConstructorFields bool, alias tlast.TL2TypeRef, constructorFields []tlast.TL2Field,
	leftArgs []tlast.TL2TypeTemplate,
	isUnionElement bool, unionIndex int, resultType TypeInstance, resultAlias bool) (*TypeInstanceStruct, error) {

	localArgs, natParams := k.fillLocalArgs(tip.templateArguments, resolvedType)

	if len(natParams) != 0 {
		return nil, fmt.Errorf("internal error - TL2 struct %s has natparams %s", canonicalName, strings.Join(natParams, ","))
	}

	ins := &TypeInstanceStruct{
		TypeInstanceCommon: TypeInstanceCommon{
			canonicalName: canonicalName,
			tlName:        tlName,
			tlTag:         tlTag,
			natParams:     natParams,
			tip:           tip,
			isTopLevel:    tip.isTopLevel && !isUnionElement,
			argNamespace:  k.getArgNamespace(resolvedType),
			hasTL2:        true,
			commentBefore: tip.combTL2.CommentBefore,
			commentRight:  "", // TODO - no comment right in TL2?
		},
		isConstructorFields: isConstructorFields,
		isUnionElement:      isUnionElement,
		unionIndex:          unionIndex,
		resultType:          resultType,
		isAlias:             false, // TL2 has separate syntax for alias
		isResultAlias:       resultAlias,
		rpcPreferTL2:        resultType != nil && k.rpcPreferTL2WhiteList.HasName2(tlName),
	}
	if !isConstructorFields { // if we are here, this is union variant or function result, where alias is field 1
		constructorFields = append(constructorFields, tlast.TL2Field{Type: alias})
	}
	nextTL2MaskBit := 0
	for _, fieldDef := range constructorFields {
		rt, fieldNatArgs, err := k.resolveType(true, fieldDef.Type, leftArgs, localArgs)
		if err != nil {
			return nil, err
		}
		fieldIns, fieldBare, err := k.getInstance(rt, true)
		if err != nil {
			return nil, err
		}
		newField := Field{
			owner:   ins,
			name:    fieldDef.Name,
			ins:     fieldIns,
			bare:    fieldBare,
			natArgs: fieldNatArgs,
			// fieldMask:     fieldMask,
			commentBefore: fieldDef.CommentBefore,
			commentRight:  fieldDef.CommentRight,
			prName:        fieldDef.PRName,
		}
		if fieldDef.IsOptional && newField.ins.ins.IsBit() {
			// we allow optional bit through aliases or template arguments,
			// but we warn if used directly
			return nil, fieldDef.PR.BeautifulError(errors.New("bit field cannot be explicitly marked optional (despite being optional void internally)"))
		}
		if fieldDef.IsOptional || newField.IsBit() {
			maskBit := nextTL2MaskBit
			newField.maskTL2Bit = &maskBit
			nextTL2MaskBit++
		}
		ins.fields = append(ins.fields, newField)
	}
	return ins, nil
}

// we want the same naming convention for nat params, as in old kernel,
// though it has no difference to semantic and can be simplified to p0, p1, p2, etc.
func (k *Kernel) fillNatParamFromArg(rt tlast.TL2TypeArgument, natParams *[]string, prefix string) {
	if rt.IsNumber {
		return
	}
	if rt.Type.String() == "*" {
		*natParams = append(*natParams, prefix)
		return
	}
	if br := rt.Type.BracketType; br != nil {
		if !br.HasIndex {
			k.fillNatParamFromArg(tlast.TL2TypeArgument{Type: br.ArrayType}, natParams, prefix+"t")
			return
		}
		if br.IndexType.IsNumber {
			k.fillNatParamFromArg(br.IndexType, natParams, prefix+"n")
			k.fillNatParamFromArg(tlast.TL2TypeArgument{Type: br.ArrayType}, natParams, prefix+"t")
		} else {
			k.fillNatParamFromArg(br.IndexType, natParams, prefix+"k")
			k.fillNatParamFromArg(tlast.TL2TypeArgument{Type: br.ArrayType}, natParams, prefix+"v")
		}
		return
	}
	tName := rt.Type.SomeType.Name.String()
	tip, ok := k.tips[tName]
	if !ok {
		panic("resolved type not found in global type map")
	}
	for i, arg := range rt.Type.SomeType.Arguments {
		leftArg := tip.combTL1[0].TemplateArguments[i]
		k.fillNatParamFromArg(arg, natParams, prefix+leftArg.FieldName)
	}
}

// Collect nat params from type tree into linear array
func (k *Kernel) fillLocalArg(arg tlast.TL2TypeArgument, targName string) (localArgs []LocalArg, natParams []string) {
	var localNatParams []string
	k.fillNatParamFromArg(arg, &localNatParams, targName)
	if len(localNatParams) == 1 {
		localNatParams[0] = targName
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
	return
}

// Collect nat params from type tree into linear array
func (k *Kernel) fillLocalArgs(leftArgs []tlast.TL2TypeTemplate, resolvedType2 tlast.TL2TypeRef) (localArgs []LocalArg, natParams []string) {
	actualArgs := resolvedType2.SomeType.Arguments // empty if brackets
	if br := resolvedType2.BracketType; br != nil {
		if br.HasIndex {
			actualArgs = append(actualArgs, br.IndexType)
		}
		actualArgs = append(actualArgs, tlast.TL2TypeArgument{Type: br.ArrayType})
	}
	for i, arg := range actualArgs {
		leftArg := leftArgs[i]
		args, params := k.fillLocalArg(arg, leftArg.Name)
		natParams = append(natParams, params...)
		localArgs = append(localArgs, args...)
	}
	return
}

func (k *Kernel) natParamsToActualNatArgs(natParams []string) []ActualNatArg {
	var result []ActualNatArg
	for _, param := range natParams {
		//if i == 0 && !resolvedType.BracketType.IndexType.IsNumber {
		//	continue
		//}
		result = append(result, ActualNatArg{
			name: param,
		})
	}
	return result
}

func (k *Kernel) isGoodBrackets(fieldDef tlast.Field) error {
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

// bracket types in returned []tlast.Field are not touched at all, because we want TL2-style types
// which cannot be expressed in tlast.TypeRef, so for type resolution we use types in []tlast.TL2TypeRef
func (k *Kernel) replaceTL1Brackets(def *tlast.Combinator) ([]tlast.Field, []tlast.TL2TypeRef, []int, error) {
	var fieldsAfterReplace []tlast.Field
	var typesAfterReplace []tlast.TL2TypeRef
	var originalFieldIndices []int
	for i := 0; i < len(def.Fields); i++ {
		fieldDef := def.Fields[i]
		var type2 tlast.TL2TypeRef
		if fieldDef.FieldType.String() == "#" && fieldDef.FieldName == "" && i+1 < len(def.Fields) {
			nextFieldDef := def.Fields[i+1]
			if nextFieldDef.Mask != nil || !nextFieldDef.IsRepeated || nextFieldDef.ScaleRepeat.ExplicitScale {
				return nil, nil, nil, fieldDef.PR.BeautifulError(fmt.Errorf("anonymous # field must be followed by brackets with no fieldmask and no explicit scale repeat (# [...] or # a:[...])"))
			}
			if err := k.isGoodBrackets(fieldDef); err != nil {
				return nil, nil, nil, err
			}
			// we replace 2 fields with vector
			// hren # a:[int] = Hren;
			i++
			type2 = tlast.TL2TypeRef{
				BracketType: &tlast.TL2BracketType{
					ArrayType: k.convertTypeRef(nextFieldDef.ScaleRepeat.Rep[0].FieldType),
					PR:        nextFieldDef.PR,
				},
				PR: nextFieldDef.PR,
			}
			fieldDef = nextFieldDef
			//fieldDef.FieldType.Args = []tlast.ArithmeticOrType{
			//	{T: fieldDef.ScaleRepeat.Rep[0].FieldType},
			//}
			//fieldDef.FieldType.Type = tlast.Name{Name: "__vector"}
			//fieldDef.FieldType.Bare = true
		} else if fieldDef.IsRepeated && i == 0 && !fieldDef.ScaleRepeat.ExplicitScale &&
			len(def.TemplateArguments) != 0 {
			a := def.TemplateArguments[len(def.TemplateArguments)-1]
			if !a.IsNat {
				e1 := fieldDef.FieldType.PR.CollapseToBegin().BeautifulError(fmt.Errorf("anonymous scale repeat implicitly references last template parameter %q which should have type #", a.FieldName))
				e2 := a.PR.BeautifulError(fmt.Errorf("see here"))
				return nil, nil, nil, tlast.BeautifulError2(e1, e2)
			}
			if err := k.isGoodBrackets(fieldDef); err != nil {
				return nil, nil, nil, err
			}
			type2 = tlast.TL2TypeRef{
				BracketType: &tlast.TL2BracketType{
					HasIndex: true,
					IndexType: tlast.TL2TypeArgument{
						Type: tlast.TL2TypeRef{
							SomeType: tlast.TL2TypeApplication{
								Name:        tlast.TL2TypeName{Name: a.FieldName},
								PR:          a.PR,
								PRName:      a.PR,
								PRArguments: a.PR.CollapseToEnd(),
							},
							PR: a.PR,
						},
						PR: a.PR,
					},
					ArrayType: k.convertTypeRef(fieldDef.ScaleRepeat.Rep[0].FieldType),
					PR:        fieldDef.PR,
				},
				PR: fieldDef.PR,
			}
			//fieldDef.FieldType.Args = []tlast.ArithmeticOrType{
			//	{T: tlast.TypeRef{PR: a.PR, Type: tlast.Name{Name: a.FieldName}}},
			//	{T: fieldDef.ScaleRepeat.Rep[0].FieldType},
			//}
			//fieldDef.FieldType.Type = tlast.Name{Name: "__tuple"}
			//fieldDef.FieldType.Bare = true
		} else if fieldDef.IsRepeated {
			if err := k.isGoodBrackets(fieldDef); err != nil {
				return nil, nil, nil, err
			}
			scale := fieldDef.ScaleRepeat.Scale
			if !fieldDef.ScaleRepeat.ExplicitScale {
				prevFieldDef := def.Fields[i-1] // never panics, due to checks above
				if prevFieldDef.FieldType.String() != "#" {
					e1 := fieldDef.FieldType.PR.CollapseToBegin().BeautifulError(fmt.Errorf("anonymous scale repeat implicitly references previous field %q, which should have type #", prevFieldDef.FieldName))
					e2 := prevFieldDef.PR.BeautifulError(fmt.Errorf("see here"))
					return nil, nil, nil, tlast.BeautifulError2(e1, e2)
				}
				scale = tlast.ScaleFactor{
					IsArith: false,
					Scale:   prevFieldDef.FieldName,
					PR:      prevFieldDef.PRName,
				}
			}
			//fieldDef.FieldType.Args = []tlast.ArithmeticOrType{
			//	{},
			//	{T: fieldDef.ScaleRepeat.Rep[0].FieldType},
			//}
			//fieldDef.FieldType.Type = tlast.Name{Name: "__tuple"}
			//fieldDef.FieldType.Bare = true
			//if scale.IsArith {
			//	fieldDef.FieldType.Args[0] = tlast.ArithmeticOrType{T: tlast.TypeRef{PR: scale.PR}, IsArith: true, Arith: scale.Arith}
			//} else {
			//	fieldDef.FieldType.Args[0] = tlast.ArithmeticOrType{T: tlast.TypeRef{PR: scale.PR, Type: tlast.Name{Name: scale.Scale}}}
			//}
			type2 = tlast.TL2TypeRef{
				BracketType: &tlast.TL2BracketType{
					HasIndex: true,
					// IndexType: set below
					ArrayType: k.convertTypeRef(fieldDef.ScaleRepeat.Rep[0].FieldType),
					PR:        fieldDef.PR,
				},
				PR: fieldDef.PR,
			}
			if scale.IsArith {
				type2.BracketType.IndexType = tlast.TL2TypeArgument{
					IsNumber: true,
					Number:   scale.Arith.Res,
					PR:       scale.PR,
				}
			} else {
				type2.BracketType.IndexType = tlast.TL2TypeArgument{
					Type: tlast.TL2TypeRef{
						SomeType: tlast.TL2TypeApplication{
							Name:        tlast.TL2TypeName{Name: scale.Scale},
							PR:          scale.PR,
							PRName:      scale.PR,
							PRArguments: scale.PR.CollapseToEnd(),
						},
						PR: scale.PR,
					},
					PR: scale.PR,
				}
			}
		} else {
			type2 = k.convertTypeRef(fieldDef.FieldType)
		}
		fieldsAfterReplace = append(fieldsAfterReplace, fieldDef)
		typesAfterReplace = append(typesAfterReplace, type2)
		originalFieldIndices = append(originalFieldIndices, i)
	}
	for _, f := range fieldsAfterReplace {
		if f.FieldName == "" && def.IsFunction {
			return nil, nil, nil, f.PR.BeautifulError(fmt.Errorf("functions cannot have anonymous fields"))
		}
		if f.FieldName == "" && (len(fieldsAfterReplace) != 1 || f.Mask != nil) {
			return nil, nil, nil, f.PR.BeautifulError(fmt.Errorf("anonymous fields are only allowed when used in '# a:[int]' pattern or when type has single anonymous field without fieldmask (typedef-like)"))
		}
	}
	return fieldsAfterReplace, typesAfterReplace, originalFieldIndices, nil
}

func (k *Kernel) createStructTL1FromTL1(canonicalName string, tip *KernelType,
	resolvedType tlast.TL2TypeRef, def *tlast.Combinator,
	isUnionElement bool, unionIndex int) (*TypeInstanceStruct, error) {
	leftArgs := append([]tlast.TL2TypeTemplate{}, tip.templateArguments...) // prevent golang aliasing

	localArgs, natParams := k.fillLocalArgs(tip.templateArguments, resolvedType)
	// fmt.Printf("natParams for %s: %s\n", canonicalName, strings.Join(natParams, ","))

	ins := &TypeInstanceStruct{
		TypeInstanceCommon: TypeInstanceCommon{
			canonicalName: canonicalName,
			tlName:        tlast.TL2TypeName(def.Construct.Name),
			tlTag:         def.Construct.ID,
			natParams:     natParams,
			tip:           tip,
			isTopLevel:    tip.isTopLevel, // both single types and union elements
			resolvedType:  resolvedType,
			argNamespace:  k.getArgNamespace(resolvedType),
			hasTL2:        false, // could be marked later
			commentBefore: def.CommentBefore,
			commentRight:  def.CommentRight,
		},
		isConstructorFields: true,
		isUnionElement:      isUnionElement,
		unionIndex:          unionIndex,
		isUnwrap:            tip.builtinWrappedCanonicalName != "",
	}
	nextTL2MaskBit := 0
	fieldsAfterReplace, typesAfterReplace, _, err := k.replaceTL1Brackets(def)
	if err != nil {
		return nil, err
	}
	if tip.canonicalName.String() == "vector" || tip.canonicalName.String() == "tuple" {
		ins.isUnwrap = true
	}
	isDict, keyRT, elemRT, err := k.IsDictWrapper(tip, resolvedType)
	if err != nil {
		return nil, err
	}
	if isDict {
		// fmt.Printf("dict detected [%s]%s\n", keyRT.String(), elemRT.String())
		ins.isUnwrap = true
		if len(fieldsAfterReplace) != 1 || len(typesAfterReplace) != 1 {
			return nil, resolvedType.PR.BeautifulError(fmt.Errorf("internal error - during Dict detection"))
		}
	}

	for i, fieldDef := range fieldsAfterReplace {
		fieldType := typesAfterReplace[i]
		var natArgs []ActualNatArg
		var rt tlast.TL2TypeRef
		if isDict {
			// types are already resolved, we should never resolve twice, this is not correct
			rt = tlast.TL2TypeRef{
				BracketType: &tlast.TL2BracketType{
					IndexType: keyRT,
					ArrayType: elemRT,
					HasIndex:  true,
					PR:        resolvedType.PR,
				},
				PR: resolvedType.PR,
			}
			// pass all our nat params to the dict
			natArgs = k.natParamsToActualNatArgs(natParams)
		} else {
			rt, natArgs, err = k.resolveType(false, fieldType, leftArgs, localArgs)
			if err != nil {
				return nil, err
			}
		}
		// fmt.Printf("resolveTypeTL2 for %s field %s: %s -> %s\n", canonicalName, fieldDef.FieldName, fieldDef.FieldType.String(), rt.String())
		fieldIns, fieldBare, err := k.getInstance(rt, true)
		if err != nil {
			return nil, err
		}
		newField := Field{
			owner:         ins,
			name:          fieldDef.FieldName,
			commentBefore: fieldDef.CommentBefore,
			commentRight:  fieldDef.CommentRight,
			ins:           fieldIns,
			natArgs:       natArgs,
			bare:          fieldBare,
			prName:        fieldDef.PRName,
		}
		if !fieldBare && fieldIns.ins != nil && fieldIns.ins.CanonicalName() == "True" &&
			!purelegacy.AllowTrueBoxed(def.Construct.Name.String(), fieldDef.FieldName) &&
			utils.DoLint(fieldDef.CommentRight) {
			// We compare type by name, because there is examples of other true types which are to be extended
			// to unions or have added fields in the future
			e1 := typesAfterReplace[i].PR.BeautifulError(fmt.Errorf("true type fields should be bare, use 'true' or '%%True' instead"))
			if k.opts.WarningsAreErrors {
				return nil, e1
			}
			e1.PrintWarning(k.opts.ErrorWriter, nil)
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
		if fieldIns.ins != nil && fieldIns.ins.CanonicalName() == "bool" &&
			newField.fieldMask != nil && !newField.fieldMask.isNumber && newField.fieldMask.isField &&
			!purelegacy.AllowBoolFieldsmask(def.Construct.Name.String(), newField.name) &&
			utils.DoLint(fieldDef.CommentRight) {
			// We compare type by name to make warning more narrow at first.
			e1 := fieldType.PR.BeautifulError(fmt.Errorf("using Bool type under fields mask produces 3rd state, use 'true' instead of 'Bool' or add // tlgen:nolint to the right"))
			if k.opts.WarningsAreErrors {
				return nil, e1
			}
			e1.PrintWarning(k.opts.ErrorWriter, nil)
		}

		ins.fields = append(ins.fields, newField)
		if fieldDef.FieldName != "" {
			leftArgs = append(leftArgs, tlast.TL2TypeTemplate{
				Name:     fieldDef.FieldName,
				Category: tlast.TL2TypeCategory{IsNatValue: true},
				PR:       fieldDef.PR,
			})
			if fieldType.String() != "#" {
				localArgs = append(localArgs, LocalArg{
					wrongTypeErr: fieldDef.PRName.BeautifulError(fmt.Errorf("defined here")),
				})
			} else {
				localArgs = append(localArgs, LocalArg{
					wrongTypeErr: nil,
					arg: tlast.TL2TypeArgument{
						Type: tlast.TL2TypeRef{SomeType: tlast.TL2TypeApplication{Name: tlast.TL2TypeName{Name: "*"}}},
						PR:   fieldDef.PR,
					},
					natArgs: []ActualNatArg{{
						isField:    true,
						fieldIndex: i, // not originalFieldIndex
					}},
				})
			}
		}
	}
	if def.IsFunction {
		fieldType := k.convertTypeRef(def.FuncDecl)
		rt, natArgs, err := k.resolveType(false, fieldType, leftArgs, localArgs)
		if err != nil {
			return nil, err
		}
		// fmt.Printf("resolveTypeTL2 for function %s result type: %s -> %s\n", canonicalName, def.FuncDecl.String(), rt.String())
		fieldIns, fieldBare, err := k.getInstance(rt, true)
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
		ins.rpcPreferTL2 = k.rpcPreferTL2WhiteList.HasName(def.Construct.Name)
	}
	ins.isTypedef = !def.IsFunction && len(ins.fields) == 1 && ins.fields[0].name == "" && ins.fields[0].FieldMask() == nil
	ins.isAlias = !isUnionElement && ins.isTypedef

	return ins, nil
}

func (k *Kernel) createAliasTL2(canonicalName string, tip *KernelType, resolvedType tlast.TL2TypeRef,
	def tlast.TL2TypeDeclaration) (TypeInstance, error) {

	leftArgs := append([]tlast.TL2TypeTemplate{}, tip.templateArguments...) // prevent golang aliasing

	localArgs, natParams := k.fillLocalArgs(tip.templateArguments, resolvedType)
	// fmt.Printf("natParams for %s: %s\n", canonicalName, strings.Join(natParams, ","))

	ins := &TypeInstanceStruct{
		TypeInstanceCommon: TypeInstanceCommon{
			canonicalName: canonicalName,
			tlName:        def.Name,
			tlTag:         def.Magic,
			natParams:     natParams,
			tip:           tip,
			isTopLevel:    tip.isTopLevel, // both single types and union elements
			resolvedType:  resolvedType,
			argNamespace:  k.getArgNamespace(resolvedType),
			hasTL2:        true,
			commentBefore: tip.combTL2.CommentBefore,
			commentRight:  "", // there is no comment right in TL2 type
		},
		isConstructorFields: false,
		isUnwrap:            false,
	}

	var natArgs []ActualNatArg
	var rt tlast.TL2TypeRef

	rt, natArgs, err := k.resolveType(true, def.Type.TypeAlias, leftArgs, localArgs)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("resolveTypeTL2 for %s field %s: %s -> %s\n", canonicalName, fieldDef.FieldName, fieldDef.FieldType.String(), rt.String())
	fieldIns, fieldBare, err := k.getInstance(rt, true)
	if err != nil {
		return nil, err
	}
	newField := Field{
		owner:         ins,
		name:          "",
		commentBefore: "",
		commentRight:  "",
		ins:           fieldIns,
		natArgs:       natArgs,
		bare:          fieldBare,
		prName:        def.Type.TypeAlias.PR,
	}
	ins.fields = append(ins.fields, newField)

	ins.isTypedef = true
	ins.isAlias = true

	return ins, nil
}
