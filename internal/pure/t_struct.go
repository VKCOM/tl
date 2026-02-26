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
	owner TypeInstance
	name  string
	ins   *TypeInstanceRef

	commentBefore string
	commentRight  string

	// though all TL2 types are bare, we still set Boxed for unions, because we want
	// vector<Union> and []Union to reference the same generated type
	bare bool

	fieldMask *ActualNatArg
	bitNumber uint32 // only used when fieldMask != nil

	maskTL2Bit *int

	natArgs []ActualNatArg // for TL1 only, empty for TL2
	//rt      tlast.TypeRef  // for TL1 only, empty for TL2
}

func (f Field) OwnerTypeInstance() TypeInstance { return f.owner }

func (f Field) Bare() bool                 { return f.bare }
func (f Field) Name() string               { return f.name }
func (f Field) CommentBefore() string      { return f.commentBefore }
func (f Field) CommentRight() string       { return f.commentRight }
func (f Field) TypeInstance() TypeInstance { return f.ins.ins }
func (f Field) FieldMask() *ActualNatArg   { return f.fieldMask }
func (f Field) BitNumber() uint32          { return f.bitNumber }
func (f Field) NatArgs() []ActualNatArg    { return f.natArgs }

// we do not know if this object is used by some other TL2 object when we generate this,
// so we return nil if owner does not marked as one needing TL2
func (f Field) MaskTL2Bit() *int {
	if !f.owner.Common().HasTL2() {
		return nil
	}
	return f.maskTL2Bit
}

func (f Field) IsBit() bool {
	if f.ins.ins.IsBit() {
		return true
	}
	return f.fieldMask != nil && f.ins.ins.CanonicalName() == "True"
}

type TypeInstanceStruct struct {
	TypeInstanceCommon
	isConstructorFields bool
	fields              []Field
	isUnionElement      bool
	unionIndex          int
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

func (ins *TypeInstanceStruct) IsUnwrap() bool {
	return ins.isUnwrap
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

	localArgs, natParams := k.getTL1ArgsHybrid(tip.templateArguments, resolvedType)

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
		},
		isConstructorFields: isConstructorFields,
		isUnionElement:      isUnionElement,
		unionIndex:          unionIndex,
		resultType:          resultType,
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
			// commentRight:  fieldDef., CommentRight - TODO
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
func (k *Kernel) fillNatParamHybrid(rt tlast.TL2TypeArgument, natParams *[]string, prefix string) {
	if rt.IsNumber {
		return
	}
	if rt.Type.String() == "*" {
		*natParams = append(*natParams, prefix)
		return
	}
	if br := rt.Type.BracketType; br != nil {
		if !br.HasIndex {
			k.fillNatParamHybrid(tlast.TL2TypeArgument{Type: br.ArrayType}, natParams, prefix+"t")
		}
		if br.IndexType.IsNumber { // not sure, may be, we better use "n" for index always
			k.fillNatParamHybrid(br.IndexType, natParams, prefix+"n")
		} else {
			k.fillNatParamHybrid(br.IndexType, natParams, prefix+"f")
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
		k.fillNatParamHybrid(arg, natParams, prefix+leftArg.FieldName)
	}
}

func (k *Kernel) getTL1ArgHybrid(arg tlast.TL2TypeArgument, targName string) (localArgs []LocalArgHybrid, natParams []string) {
	var localNatParams []string
	k.fillNatParamHybrid(arg, &localNatParams, targName)
	if len(localNatParams) == 1 {
		localNatParams[0] = targName
	}
	natParams = append(natParams, localNatParams...)
	localArg := LocalArgHybrid{
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

func (k *Kernel) getTL1ArgsHybrid(leftArgs []tlast.TL2TypeTemplate, resolvedType2 tlast.TL2TypeRef) (localArgs []LocalArgHybrid, natParams []string) {
	actualArgs := resolvedType2.SomeType.Arguments // empty if brackets
	if br := resolvedType2.BracketType; br != nil {
		if br.HasIndex {
			actualArgs = append(actualArgs, br.IndexType)
		}
		actualArgs = append(actualArgs, tlast.TL2TypeArgument{Type: br.ArrayType})
	}
	for i, arg := range actualArgs {
		leftArg := leftArgs[i]
		args, params := k.getTL1ArgHybrid(arg, leftArg.Name)
		natParams = append(natParams, params...)
		localArgs = append(localArgs, args...)
	}
	return
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
		if f.FieldName == "" && (len(fieldsAfterReplace) != 1 || f.Mask != nil) {
			return nil, nil, nil, f.PR.BeautifulError(fmt.Errorf("anonymous fields are only allowed when used in '# a:[int]' pattern or when type has single anonymous field without fieldmask (typedef-like)"))
		}
	}
	return fieldsAfterReplace, typesAfterReplace, originalFieldIndices, nil
}

func (k *Kernel) createStructTL1FromTL1(canonicalName string, tip *KernelType,
	resolvedType2 tlast.TL2TypeRef, def *tlast.Combinator,
	isUnionElement bool, unionIndex int) (*TypeInstanceStruct, error) {
	leftArgs := append([]tlast.TL2TypeTemplate{}, tip.templateArguments...) // prevent golang aliasing

	localArgs, natParams := k.getTL1ArgsHybrid(tip.templateArguments, resolvedType2)
	// fmt.Printf("natParams for %s: %s\n", canonicalName, strings.Join(natParams, ","))

	ins := &TypeInstanceStruct{
		TypeInstanceCommon: TypeInstanceCommon{
			canonicalName: canonicalName,
			tlName:        tlast.TL2TypeName(def.Construct.Name),
			tlTag:         def.Construct.ID,
			natParams:     natParams,
			tip:           tip,
			isTopLevel:    tip.isTopLevel, // both single types and union elements
			rt2:           resolvedType2,
			argNamespace:  k.getArgNamespace(resolvedType2),
			hasTL2:        false, // could be marked later
		},
		isConstructorFields: true,
		isUnionElement:      isUnionElement,
		unionIndex:          unionIndex,
		isUnwrap:            tip.builtinWrappedCanonicalName != "",
	}
	nextTL2MaskBit := 0
	fieldsAfterReplace, typesAfterReplace, originalFieldIndices, err := k.replaceTL1Brackets(def)
	if err != nil {
		return nil, err
	}
	if tip.canonicalName.String() == "vector" || tip.canonicalName.String() == "tuple" {
		ins.isUnwrap = true
	}
	if isDict, dictFieldT := k.IsDict(tip); isDict {
		fieldT := tlast.TypeRef{
			Type: tlast.Name(dictFieldT.canonicalName),
			Bare: true,
		}
		// TODO - I'm not sure if passing PR of actualArgs is correct
		for i, targ := range tip.combTL1[0].TemplateArguments {
			fieldT.Args = append(fieldT.Args, tlast.ArithmeticOrType{
				T: tlast.TypeRef{
					Type: tlast.Name{Name: targ.FieldName},
					PR:   resolvedType2.SomeType.Arguments[i].PR,
				},
			})
		}
		// TODO - PR below
		ins.isUnwrap = true
		fieldsAfterReplace = []tlast.Field{{
			FieldName: "",
			FieldType: tlast.TypeRef{
				Type: tlast.Name{Name: "__dict"},
				Bare: true, // TODO - remove
				Args: []tlast.ArithmeticOrType{{
					T: fieldT,
				}},
			},
		}}
		typesAfterReplace = []tlast.TL2TypeRef{{
			SomeType: tlast.TL2TypeApplication{
				Name: tlast.TL2TypeName{Name: "__dict"},
				Bare: true, // TODO - remove
				Arguments: []tlast.TL2TypeArgument{{
					Type: k.convertTypeRef(fieldT),
				}},
			},
			PR: tlast.PositionRange{},
		}}
		originalFieldIndices = []int{0}
	}

	for i, fieldDef := range fieldsAfterReplace {
		fieldType := typesAfterReplace[i]
		originalFieldIndex := originalFieldIndices[i]
		rt, natArgs, err := k.resolveType(false, fieldType, leftArgs, localArgs)
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
			name:          fieldDef.FieldName,
			commentBefore: fieldDef.CommentBefore,
			commentRight:  fieldDef.CommentRight,
			ins:           fieldIns,
			natArgs:       natArgs,
			bare:          fieldBare,
		}
		if !fieldBare && fieldIns.ins != nil && fieldIns.ins.CanonicalName() == "True" &&
			!purelegacy.AllowTrueBoxed(def.Construct.Name.String(), fieldDef.FieldName) &&
			utils.DoLint(fieldDef.CommentRight) {
			// We compare type by name, because there is examples of other true types which are to be extended
			// to unions or have added fields in the future
			e1 := fieldDef.FieldType.PR.BeautifulError(fmt.Errorf("true type fields should be bare, use 'true' or '%%True' instead"))
			if k.opts.WarningsAreErrors {
				return nil, e1
			}
			e1.PrintWarning(k.opts.ErrorWriter, nil)
		}
		if fieldDef.Mask != nil {
			if fieldDef.Mask.BitNumber >= 32 {
				return nil, fieldDef.Mask.PRBits.BeautifulError(fmt.Errorf("bitmask (%d) must be in range [0..31]", fieldDef.Mask.BitNumber))
			}
			fieldMask, err := k.resolveMaskTL1(*fieldDef.Mask, leftArgs, localArgs,
				tlast.CombinatorField{Comb: def, FieldIndex: originalFieldIndex})
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
			e1 := fieldDef.FieldType.PR.BeautifulError(fmt.Errorf("using Bool type under fields mask produces 3rd state, use 'true' instead of 'Bool' or add // tlgen:nolint to the right"))
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
			if fieldDef.FieldType.String() != "#" {
				localArgs = append(localArgs, LocalArgHybrid{
					wrongTypeErr: fieldDef.PRName.BeautifulError(fmt.Errorf("defined here")),
				})
			} else {
				localArgs = append(localArgs, LocalArgHybrid{
					wrongTypeErr: nil,
					arg: tlast.TL2TypeArgument{
						Type:        tlast.TL2TypeRef{SomeType: tlast.TL2TypeApplication{Name: tlast.TL2TypeName{Name: "*"}}},
						PR:          fieldDef.PR,
						SourceField: tlast.CombinatorField{Comb: def, FieldIndex: originalFieldIndex},
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
	return ins, nil
}
