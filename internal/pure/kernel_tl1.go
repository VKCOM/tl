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
)

var (
	errSeeHere = fmt.Errorf("see here")
	// errFieldNameCollision     = fmt.Errorf("field name collision")
	// errNatParamNameCollision  = fmt.Errorf("nat-parameter name collision")
	// errTypeParamNameCollision = fmt.Errorf("type-parameter name collision ")
)

//func (k *Kernel) shouldSkipDefinition(typ *tlast.Combinator) bool {
//	return typ.Construct.Name.String() == "vector" || typ.Construct.Name.String() == "tuple"
//}

func (k *Kernel) CompileBoolTL1(tlType []*tlast.Combinator) error {
	// if type is
	// 1. enum with 2 elements, 0 template arguments
	// 2. has name "Bool"
	// 3. fields have names "boolFalse" and "boolTrue"
	if len(tlType) != 2 ||
		len(tlType[0].Fields) != 0 || len(tlType[1].Fields) != 0 ||
		len(tlType[0].TemplateArguments) != 0 || len(tlType[1].TemplateArguments) != 0 {
		return fmt.Errorf("kernel supports only exact TL1 Bool definition: 'boolFalse#<magic1> = Bool; boolTrue#<magic2> = Bool;'")
	}
	if tlType[0].Construct.Name.String() != "boolFalse" ||
		tlType[1].Construct.Name.String() != "boolTrue" {
		return fmt.Errorf("kernel supports only exact TL1 Bool definition: 'boolFalse#<magic1> = Bool; boolTrue#<magic2> = Bool;'")
	}
	tip, ok := k.tips["bool"]
	if !ok {
		return tlType[0].Construct.NamePR.BeautifulError(errors.New("internal error builtin type not found"))
	}
	if _, ok2 := k.tips["Bool"]; ok2 || len(tip.combTL1) != 1 {
		// TODO - see here
		return tlType[0].TypeDecl.NamePR.BeautifulError(errors.New("builtin type Bool already defined as not builtin"))
	}
	tip.combTL1 = tlType
	tip.originTL2 = false // allow references from TL1
	k.tips["Bool"] = tip
	tip.tl1Names["Bool"] = struct{}{}
	tip.tl1BoxedName = tlast.Name{Name: "Bool"}
	// we do not allow references to boxed wrappers from TL2
	return nil
}

func (k *Kernel) CompileBuiltinTL1(typ *tlast.Combinator) error {
	for _, arg := range typ.TemplateArguments {
		return arg.PR.BeautifulError(errors.New("built-in wrapper cannot have template arguments"))
	}
	addBuiltin := func(tl2name string, bigName string) error {
		if typ.TypeDecl.Name.String() != bigName {
			return typ.TypeDecl.NamePR.BeautifulError(fmt.Errorf("built-in wrapper must have type name %s", bigName))
		}
		//tip, ok := k.tips[tl2name]
		//if !ok {
		//	return typ.Construct.NamePR.BeautifulError(errors.New("internal error builtin type not found"))
		//}
		if _, ok2 := k.tips[bigName]; ok2 {
			// TODO - see here
			return typ.TypeDecl.NamePR.BeautifulError(errors.New("builtin type already defined"))
		}
		combTL1 := *typ
		combTL1.Builtin = false
		combTL1.Fields = []tlast.Field{{
			FieldType: tlast.TypeRef{Type: typ.Construct.Name},
		}}
		kt := &KernelType{
			originTL2:      false,
			combTL1:        []*tlast.Combinator{&combTL1},
			instances:      map[string]*TypeInstanceRef{},
			tl1Names:       map[string]struct{}{bigName: {}},
			tl2Names:       map[string]struct{}{},
			canonicalName:  typ.TypeDecl.Name,
			historicalName: typ.TypeDecl.Name,
			tl1BoxedName:   typ.TypeDecl.Name,
			isTopLevel:     true,

			builtinWrappedCanonicalName: typ.Construct.Name.String(),
		}
		if err := k.addTip(kt, bigName, ""); err != nil {
			return err
		}
		//k.tips[bigName] = kt
		//k.tipsTopLevel = append(k.tipsTopLevel, tip)
		//tip.tl1Names[bigName] = struct{}{}
		//tip.tl1BoxedName = tlast.Name{Name: bigName}
		// we do not allow references to boxed wrappers from TL2
		return nil
	}
	switch typ.Construct.Name.String() {
	case "int":
		return addBuiltin("int32", "Int")
	case "long":
		return addBuiltin("int64", "Long")
	case "string":
		return addBuiltin("string", "String")
	case "float":
		return addBuiltin("float32", "Float")
	case "double":
		return addBuiltin("float64", "Double")
	}
	return typ.Construct.NamePR.BeautifulError(fmt.Errorf("built-in wrapper must have constructor name %s equal to some built-in type", typ.Construct.Name.String()))
}

func (k *Kernel) CompileTL1() error {
	log.Printf("tl2pure: compiling %d TL1 combinators", len(k.filesTL1))
	// Collect unions, check that functions cannot form a union with each other or with normal singleConstructors
	allConstructors := map[string]*tlast.Combinator{}
	typeDescriptors := map[string][]*tlast.Combinator{}
	for i := range k.filesTL1 {
		k.filesTL1[i].OriginalOrderIndex = i
	}
	var boolCombinators []*tlast.Combinator
	for _, typ := range k.filesTL1 {
		if typ.Builtin {
			if err := k.CompileBuiltinTL1(typ); err != nil {
				return err
			}
			continue
		}
		if typ.TypeDecl.Name.String() == "Bool" {
			boolCombinators = append(boolCombinators, typ)
			continue
		}
		for _, field := range typ.Fields {
			if field.Excl && !purelegacy.EnableExclamation(typ.Construct.Name.String()) {
				return field.PR.BeautifulError(fmt.Errorf("new !X function wrappers are forbidden"))
			}
		}
		conName := typ.Construct.Name.String()
		if col, ok := allConstructors[conName]; ok {
			// typeA = TypeA;
			// typeA = TypeB;
			e1 := typ.Construct.NamePR.BeautifulError(fmt.Errorf("constructor name %q is used again here", conName))
			e2 := col.Construct.NamePR.BeautifulError(errSeeHere)
			return tlast.BeautifulError2(e1, e2)
		}
		allConstructors[conName] = typ
		if !typ.IsFunction {
			typeName := typ.TypeDecl.Name.String()
			if len(typ.TemplateArguments) > len(typ.TypeDecl.Arguments) {
				// rightLeftArgs {X:Type} {Y:#} = RightLeftArgs X; <- bad
				arg := typ.TemplateArguments[len(typ.TypeDecl.Arguments)]
				return typ.TypeDecl.PR.CollapseToEnd().BeautifulError(fmt.Errorf("type declaration %q is missing template argument %q here", typeName, arg.FieldName))
			}
			if len(typ.TemplateArguments) < len(typ.TypeDecl.Arguments) {
				// rightLeftArgs {X:Type} {Y:#} = RightLeftArgs X Y Y; <- bad
				arg := typ.TypeDecl.Arguments[len(typ.TemplateArguments)]
				pr := typ.TypeDecl.ArgumentsPR[len(typ.TemplateArguments)]
				return pr.BeautifulError(fmt.Errorf("type declaration %q has excess template argument %q here", typeName, arg))
			}
			for j, t := range typ.TemplateArguments {
				if t.FieldName != typ.TypeDecl.Arguments[j] {
					// rightLeftArgs {X:Type} {Y:#} = RightLeftArgs Y X;   <- bad
					pr := typ.TypeDecl.ArgumentsPR[j]
					typArg := typ.TypeDecl.Arguments[j]
					e1 := pr.BeautifulError(fmt.Errorf("type declaration %q has wrong template argument name %q here", typeName, typArg))
					e2 := t.PR.BeautifulError(errSeeHere)
					return tlast.BeautifulError2(e1, e2)
				}
			}
			if typeName == "_" { // prohibit boxed type
				return fmt.Errorf("internal error - underscore as a typename prohibited, must not pass AST parser")
			}
			typeDescriptors[typeName] = append(typeDescriptors[typeName], typ)
		} else {
			for _, t := range typ.TemplateArguments {
				if t.IsNat {
					// @read funWithArg {fields_mask: #} => True;
					return t.PR.BeautifulError(fmt.Errorf("function declaration %q cannot have template arguments", conName))
				}
				// TODO - sort out things with rpc wrapping later which has a form
				// @readwrite tree_stats.preferMaster {X:Type} query:!X = X;
			}
			if len(typ.Modifiers) == 0 && utils.DoLint(typ.CommentRight) {
				e1 := typ.Construct.NamePR.CollapseToBegin().BeautifulError(fmt.Errorf("function constructor %q without modifier (identifier starting with '@') not recommended", typ.Construct.Name.String()))
				if k.opts.WarningsAreErrors {
					return e1
				}
				e1.PrintWarning(k.opts.ErrorWriter, nil)
			}
		}
		var nc NameCollision
		for _, targ := range typ.TemplateArguments {
			if err := nc.AddUniqueName(targ.FieldName, targ.PR, "template argument"); err != nil {
				return err
			}
		}
		nc.ResetNormalized()
		for _, field := range typ.Fields {
			if field.FieldName == "" {
				continue
			}
			if err := nc.AddUniqueName(field.FieldName, field.PR, "field"); err != nil {
				return err
			}
		}
	}
	if len(boolCombinators) != 0 {
		if err := k.CompileBoolTL1(boolCombinators); err != nil {
			return err
		}
	}
	// in order for deterministic migration
	for _, comb := range k.filesTL1 {
		if comb.IsFunction {
			cName := comb.Construct.Name
			kt := &KernelType{
				originTL2:  false,
				combTL1:    []*tlast.Combinator{comb},
				instances:  map[string]*TypeInstanceRef{},
				isFunction: true,
				isTopLevel: true,
				// functions have no canonical name, because there is no references to functions
				// also they have no TL1 names or TL2 names set.
				canonicalName:  cName,
				historicalName: cName,
				canBeBare:      true,
			}
			for _, m := range comb.Modifiers {
				kt.annotations = append(kt.annotations, m.Name)
			}
			if err := k.addTip(kt, cName.String(), ""); err != nil {
				return fmt.Errorf("error adding function %s: %w", cName.String(), err)
			}
			continue
		}
		typ, ok := typeDescriptors[comb.TypeDecl.Name.String()]
		if !ok {
			continue
		}
		delete(typeDescriptors, comb.TypeDecl.Name.String())
		tName := typ[0].TypeDecl.Name
		cName := typ[0].Construct.Name
		if len(typ) == 1 {
			if typ[0].IsFunction {
				return fmt.Errorf("internal error - function %q must not be in type descriptors", cName)
			}
			typePrefix := strings.ToLower(utils.ToLowerFirst(tName.Name))

			if cName.Namespace != tName.Namespace {
				e1 := typ[0].Construct.NamePR.BeautifulError(fmt.Errorf("simple type constructor namespace should exactly match type namespace"))
				e2 := typ[0].TypeDecl.NamePR.BeautifulError(errSeeHere)
				return tlast.BeautifulError2(e1, e2)
			}
			// We temporarily allow relaxed case match. To use strict match, remove strings.ToLower() calls below
			if strings.ToLower(cName.Name) != typePrefix &&
				!purelegacy.EnableWarningsSimpleTypeNameSkip(cName.String()) && utils.DoLint(typ[0].CommentRight) {
				e1 := typ[0].Construct.NamePR.BeautifulError(fmt.Errorf("simple type constructor name should differ from type name by case only"))
				e2 := typ[0].TypeDecl.NamePR.BeautifulError(errSeeHere)
				if k.opts.WarningsAreErrors {
					return tlast.BeautifulError2(e1, e2)
				}
				tlast.BeautifulError2(e1, e2).PrintWarning(k.opts.ErrorWriter, nil)
			}
			kt := &KernelType{
				originTL2:      false,
				combTL1:        typ,
				instances:      map[string]*TypeInstanceRef{},
				isTopLevel:     len(typ[0].TemplateArguments) == 0,
				tl1Names:       map[string]struct{}{cName.String(): {}, tName.String(): {}},
				tl2Names:       map[string]struct{}{cName.String(): {}, tName.String(): {}},
				canonicalName:  cName,
				historicalName: cName,
				tl1BoxedName:   tName,
				canBeBare:      true,
				targs:          make([]KernelTypeTarg, len(typ[0].TemplateArguments)),
			}
			for _, m := range comb.Modifiers {
				kt.annotations = append(kt.annotations, m.Name)
			}
			if err := k.addTip(kt, cName.String(), tName.String()); err != nil {
				return typ[0].Construct.NamePR.BeautifulError(fmt.Errorf("error adding type %s: %w", cName, err))
			}
			continue
		}
		if err := k.checkUnionElementsCompatibility(typ); err != nil {
			return err
		}
		kt := &KernelType{
			originTL2:      false,
			combTL1:        typ,
			instances:      map[string]*TypeInstanceRef{},
			isTopLevel:     len(typ[0].TemplateArguments) == 0,
			tl1Names:       map[string]struct{}{tName.String(): {}},
			tl2Names:       map[string]struct{}{tName.String(): {}},
			canonicalName:  tName,
			historicalName: tName,
			tl1BoxedName:   tName,
			targs:          make([]KernelTypeTarg, len(typ[0].TemplateArguments)),
		}
		for _, m := range comb.Modifiers {
			return m.PR.BeautifulError(fmt.Errorf("annotations in TL1 are not supported for union %s", tName))
		}
		if err := k.addTip(kt, tName.String(), ""); err != nil {
			return err
		}
	}
	//for _, comb := range k.filesTL1 {
	//	log.Printf("tl2pure: compiling %s", comb)
	//	kt := &KernelType{
	//		originTL2: false,
	//		combTL1:   comb,
	//		instances: map[string]*TypeInstanceRef{},
	//	}
	//	if !comb.IsFunction {
	//		namesNormalized := map[string]int{}
	//		names := map[string]int{}
	//		for i, targ := range comb.TemplateArguments {
	//			if _, ok := names[targ.FieldName]; ok {
	//				return fmt.Errorf("error adding type %s: template argument %s name collision", comb.Construct.Name, targ.FieldName)
	//			}
	//			nn := k.normalizeName(targ.FieldName)
	//			if _, ok := names[nn]; ok {
	//				return fmt.Errorf("error adding type %s: template argument %s normalized name collision", comb.Construct.Name, nn)
	//			}
	//			namesNormalized[nn] = i
	//		}
	//	}
	//	if err := k.addTip(kt); err != nil {
	//		return fmt.Errorf("error adding type %s: %w", comb.String(), err)
	//	}
	//}
	return nil
}

func (k *Kernel) checkUnionElementsCompatibility(types []*tlast.Combinator) error {
	// We temporarily allow relaxed case match. To use strict match, remove strings.ToLower() calls below
	typePrefix := strings.ToLower(utils.ToLowerFirst(types[0].TypeDecl.Name.Name))
	typeSuffix := strings.ToLower(types[0].TypeDecl.Name.Name)
	for _, typ := range types {
		conName := strings.ToLower(typ.Construct.Name.Name)
		if typ.Construct.Name.Namespace != typ.TypeDecl.Name.Namespace &&
			!purelegacy.EnableWarningsUnionNamespaceSkip(typ.Construct.Name.Namespace, typ.TypeDecl.Name.Namespace) &&
			utils.DoLint(typ.CommentRight) {
			e1 := typ.Construct.NamePR.BeautifulError(fmt.Errorf("union constructor namespace %q should match type namespace %q", typ.Construct.Name.Namespace, typ.TypeDecl.Name.Namespace))
			e2 := typ.TypeDecl.NamePR.BeautifulError(errSeeHere)
			if k.opts.WarningsAreErrors {
				return tlast.BeautifulError2(e1, e2)
			}
			tlast.BeautifulError2(e1, e2).PrintWarning(k.opts.ErrorWriter, nil)
		}
		if !strings.HasPrefix(conName, typePrefix) &&
			!strings.HasSuffix(conName, typeSuffix) &&
			!purelegacy.EnableWarningsUnionNamePrefixSkip(typ.Construct.Name.Name, typePrefix, typeSuffix) &&
			utils.DoLint(typ.CommentRight) { // same check as in generateType
			e1 := typ.Construct.NamePR.BeautifulError(fmt.Errorf("union constructor should have type name prefix or suffix %q", typePrefix))
			e2 := typ.TypeDecl.NamePR.BeautifulError(errSeeHere)
			if k.opts.WarningsAreErrors {
				return tlast.BeautifulError2(e1, e2)
			}
			tlast.BeautifulError2(e1, e2).PrintWarning(k.opts.ErrorWriter, nil)
			continue
		}
		if conName == typePrefix &&
			!purelegacy.EnableWarningsUnionNameExactSkip(typ.Construct.Name.String()) &&
			utils.DoLint(typ.CommentRight) {
			e1 := typ.Construct.NamePR.BeautifulError(fmt.Errorf("union constructor name should not exactly match type name %q", typePrefix))
			e2 := typ.TypeDecl.PR.BeautifulError(errSeeHere)
			if k.opts.WarningsAreErrors {
				return tlast.BeautifulError2(e1, e2)
			}
			tlast.BeautifulError2(e1, e2).PrintWarning(k.opts.ErrorWriter, nil)
		}
	}
	base := types[0]
	for _, typ := range types[1:] {
		cur := typ.Construct.Name.String()
		if len(typ.TemplateArguments) < len(base.TemplateArguments) {
			baseArg := base.TemplateArguments[len(typ.TemplateArguments)]
			// unionArgs2 {A:Type} {B:#} a:A = UnionArgs A B;
			// unionArgs1 {X:Type} a:X = UnionArgs X;
			e1 := typ.TemplateArgumentsPR.CollapseToEnd().BeautifulError(fmt.Errorf("union constructor %q has missing argument %q here", cur, baseArg.FieldName))
			e2 := baseArg.PR.BeautifulError(errSeeHere)
			return tlast.BeautifulError2(e1, e2)
		}
		if len(typ.TemplateArguments) > len(base.TemplateArguments) {
			typArg := typ.TemplateArguments[len(base.TemplateArguments)]
			// unionArgs1 {X:Type} a:X = UnionArgs X;
			// unionArgs2 {A:Type} {B:#} a:A = UnionArgs A B;
			e1 := typArg.PR.BeautifulError(fmt.Errorf("union constructor %q has excess argument %q here", cur, typArg.FieldName))
			e2 := base.TemplateArgumentsPR.CollapseToEnd().BeautifulError(errSeeHere)
			return tlast.BeautifulError2(e1, e2)
		}
		for i, typArg := range typ.TemplateArguments {
			baseArg := base.TemplateArguments[i]
			// unionArgs1 {X:Type} {Y:#} a:X = UnionArgs X Y;
			// unionArgs2 {A:Type} {B:Type} a:A = UnionArgs A B;
			// We cannot support this, because resolveTypeTL2 replaces parameter names into names of first union field
			if baseArg.IsNat != typArg.IsNat || baseArg.FieldName != typArg.FieldName {
				e1 := typArg.PR.BeautifulError(fmt.Errorf("union constructor %q has different argument name or type here %q", cur, typArg.FieldName))
				e2 := baseArg.PR.BeautifulError(errSeeHere)
				return tlast.BeautifulError2(e1, e2)
			}
		}
	}
	return nil
}
