// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package pure

import (
	"fmt"
	"log"

	"github.com/vkcom/tl/internal/tlast"
)

var (
	errSeeHere = fmt.Errorf("see here")
	// errFieldNameCollision     = fmt.Errorf("field name collision")
	// errNatParamNameCollision  = fmt.Errorf("nat-parametr name collision")
	// errTypeParamNameCollision = fmt.Errorf("type-parametr name collision ")
)

func (k *Kernel) CompileTL1() error {
	log.Printf("tl2pure: compiling %d TL1 combinators", len(k.filesTL1))
	// Collect unions, check that functions cannot form a union with each other or with normal singleConstructors
	allConstructors := map[string]*tlast.Combinator{}
	typeDescriptors := map[string][]*tlast.Combinator{}
	for _, typ := range k.filesTL1 {
		for _, f := range typ.Fields {
			if f.FieldName == "" && (len(typ.Fields) != 1 || f.Mask != nil) {
				return f.PR.BeautifulError(fmt.Errorf("anonymous fields are discouraged, except when used in '# a:[int]' pattern or when type has single anonymous field without fieldmask (typedef-like)"))
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
			// TODO - copy from tlgen.go
			//if len(typ.Modifiers) == 0 && doLint(typ.CommentRight) {
			//	e1 := typ.Construct.NamePR.CollapseToBegin().BeautifulError(fmt.Errorf("function constructor %q without modifier (identifier starting with '@') not recommended", typ.Construct.Name.String()))
			//	if gen.options.WarningsAreErrors {
			//		return e1
			//	}
			//	e1.PrintWarning(gen.options.ErrorWriter, nil)
			//}
		}
	}
	for _, typ := range typeDescriptors {
		tName := typ[0].TypeDecl.Name
		if len(typ) == 1 { // here there is no functions
			cName := typ[0].Construct.Name
			//typePrefix := strings.ToLower(utils.LowerFirst(tName.Name))

			if cName.Namespace != tName.Namespace {
				e1 := typ[0].Construct.NamePR.BeautifulError(fmt.Errorf("simple type constructor namespace should exactly match type namespace"))
				e2 := typ[0].TypeDecl.NamePR.BeautifulError(errSeeHere)
				return tlast.BeautifulError2(e1, e2)
			}
			// We temporarily allow relaxed case match. To use strict match, remove strings.ToLower() calls below
			// TODO - copy from tlgen.go
			//if EnableWarningsSimpleTypeName && strings.ToLower(cName.Name) != typePrefix &&
			//	!LegacyEnableWarningsSimpleTypeNameSkip(cName.String()) && doLint(typ[0].CommentRight) {
			//	e1 := typ[0].Construct.NamePR.BeautifulError(fmt.Errorf("simple type constructor name should differ from type name by case only"))
			//	e2 := typ[0].TypeDecl.NamePR.BeautifulError(errSeeHere)
			//	if gen.options.WarningsAreErrors {
			//		return tlast.BeautifulError2(e1, e2)
			//	}
			//	tlast.BeautifulError2(e1, e2).PrintWarning(gen.options.ErrorWriter, nil)
			//}
			kt := &KernelType{
				originTL2: false,
				combTL1:   typ,
				instances: map[string]*TypeInstanceRef{},
			}
			if err := k.addTip(kt, cName.String(), tName.String()); err != nil {
				return fmt.Errorf("error adding type %s: %w", typ[0].String(), err)
			}
			continue
		}
		if err := checkUnionElementsCompatibility(typ); err != nil {
			return err
		}
		kt := &KernelType{
			originTL2: false,
			combTL1:   typ,
			instances: map[string]*TypeInstanceRef{},
		}
		if err := k.addTip(kt, tName.String(), ""); err != nil {
			return fmt.Errorf("error adding type %s: %w", typ[0].String(), err)
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

func checkUnionElementsCompatibility(types []*tlast.Combinator) error {
	// We temporarily allow relaxed case match. To use strict match, remove strings.ToLower() calls below
	// TODO - copy from tlgen.go
	//typePrefix := strings.ToLower(utils.LowerFirst(types[0].TypeDecl.Name.Name))
	//typeSuffix := strings.ToLower(types[0].TypeDecl.Name.Name)
	//for _, typ := range types {
	//	conName := strings.ToLower(typ.Construct.Name.Name)
	//	if EnableWarningsUnionNamespace && typ.Construct.Name.Namespace != typ.TypeDecl.Name.Namespace &&
	//		!LegacyEnableWarningsUnionNamespaceSkip(typ.Construct.Name.Namespace, typ.TypeDecl.Name.Namespace) &&
	//		doLint(typ.CommentRight) {
	//		e1 := typ.Construct.NamePR.BeautifulError(fmt.Errorf("union constructor namespace %q should match type namespace %q", typ.Construct.Name.Namespace, typ.TypeDecl.Name.Namespace))
	//		e2 := typ.TypeDecl.NamePR.BeautifulError(errSeeHere)
	//		if options.WarningsAreErrors {
	//			return tlast.BeautifulError2(e1, e2)
	//		}
	//		tlast.BeautifulError2(e1, e2).PrintWarning(options.ErrorWriter, nil)
	//	}
	//	if EnableWarningsUnionNamePrefix &&
	//		!strings.HasPrefix(conName, typePrefix) &&
	//		!strings.HasSuffix(conName, typeSuffix) &&
	//		!LegacyEnableWarningsUnionNamePrefixSkip(typ.Construct.Name.Name, typePrefix, typeSuffix) &&
	//		doLint(typ.CommentRight) { // same check as in generateType
	//		e1 := typ.Construct.NamePR.BeautifulError(fmt.Errorf("union constructor should have type name prefix or suffix %q", typePrefix))
	//		e2 := typ.TypeDecl.NamePR.BeautifulError(errSeeHere)
	//		if options.WarningsAreErrors {
	//			return tlast.BeautifulError2(e1, e2)
	//		}
	//		tlast.BeautifulError2(e1, e2).PrintWarning(options.ErrorWriter, nil)
	//		continue
	//	}
	//	if EnableWarningsUnionNameExact && conName == typePrefix &&
	//		!LegacyEnableWarningsUnionNameExactSkip(typ.Construct.Name.String()) &&
	//		doLint(typ.CommentRight) {
	//		e1 := typ.Construct.NamePR.BeautifulError(fmt.Errorf("union constructor name should not exactly match type name %q", typePrefix))
	//		e2 := typ.TypeDecl.PR.BeautifulError(errSeeHere)
	//		if options.WarningsAreErrors {
	//			return tlast.BeautifulError2(e1, e2)
	//		}
	//		tlast.BeautifulError2(e1, e2).PrintWarning(options.ErrorWriter, nil)
	//	}
	//}
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
			// We cannot support this, because resolveType replaces parameter names into names of first union field
			if baseArg.IsNat != typArg.IsNat || baseArg.FieldName != typArg.FieldName {
				e1 := typArg.PR.BeautifulError(fmt.Errorf("union constructor %q has different argument name or type here %q", cur, typArg.FieldName))
				e2 := baseArg.PR.BeautifulError(errSeeHere)
				return tlast.BeautifulError2(e1, e2)
			}
		}
	}
	return nil
}
