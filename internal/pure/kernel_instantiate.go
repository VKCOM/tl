// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package pure

import (
	"errors"
	"fmt"

	"github.com/vkcom/tl/internal/tlast"
)

type LocalArgHybrid struct {
	wrongTypeErr *tlast.ParseError // we must add all field names to local context, because they must correctly shadow names outside, but we check the type
	arg          tlast.TL2TypeArgument
	natArgs      []ActualNatArg
}

func (k *Kernel) resolveTypeHybrid(ctxTL2 bool, tr tlast.TL2TypeRef, leftArgs []tlast.TemplateArgument,
	actualArgs []LocalArgHybrid) (tlast.TL2TypeRef, []ActualNatArg, error) {
	ac, natArgs, err := k.resolveArgumentHybrid(ctxTL2, tlast.TL2TypeArgument{Type: tr}, leftArgs, actualArgs)
	if err != nil {
		return tr, nil, err
	}
	if ac.IsNumber {
		// TODO - beautiful test case,
		return tr, nil, fmt.Errorf("type argument %s resolved to a number %d", tr, ac.Number)
	}
	if ac.Type.String() == "*" {
		// TODO - beautiful test case,
		return tr, nil, tr.PR.BeautifulError(fmt.Errorf("type argument %s resolved to a nat argument %s", tr, ac.Type))
	}
	return ac.Type, natArgs, nil
}

func (k *Kernel) resolveArgumentHybrid(ctxTL2 bool, tr tlast.TL2TypeArgument, leftArgs []tlast.TemplateArgument,
	actualArgs []LocalArgHybrid) (tlast.TL2TypeArgument, []ActualNatArg, error) {
	before := tr
	was := before.String()
	tr, natArgs, err := k.resolveArgumentHybridImpl(ctxTL2, tr, leftArgs, actualArgs)
	after := before.String()
	if was != after {
		panic(fmt.Sprintf("tl2pure: internal error, resolveArgumentTL1 destroyed %s original value %s due to golang aliasing", after, was))
	}
	return tr, natArgs, err
}

func (k *Kernel) resolveArgumentHybridImpl(ctxTL2 bool, tr tlast.TL2TypeArgument, leftArgs []tlast.TemplateArgument,
	actualArgs []LocalArgHybrid) (tlast.TL2TypeArgument, []ActualNatArg, error) {
	if tr.IsNumber {
		return tr, nil, nil
	}
	if tr.Type.String() == "*" {
		return tr, nil, tr.Type.PR.BeautifulError(fmt.Errorf("internal error: resolving type more than once"))
	}
	if br := tr.Type.BracketType; br != nil {
		var natArgs []ActualNatArg
		bracketType := *tr.Type.BracketType
		if bracketType.HasIndex {
			ic, natArgs2, err := k.resolveArgumentHybrid(ctxTL2, bracketType.IndexType, leftArgs, actualArgs)
			if err != nil {
				return tr, nil, err
			}
			bracketType.IndexType = ic
			natArgs = append(natArgs, natArgs2...)

			// TODO - restore SourceField functions
			//if tName == "__tuple" {
			//	if sf := tr.T.Args[0].SourceField; sf != (tlast.CombinatorField{}) {
			//		field := &sf.Comb.Fields[sf.FieldIndex]
			//		if field.UsedAsMask {
			//			e3 := field.UsedAsMaskPR.BeautifulError(fmt.Errorf("used as mask here"))
			//			e3.PrintWarning(k.opts.ErrorWriter, nil)
			//			e1 := field.PRName.BeautifulError(fmt.Errorf("#-field %s is used as tuple size, while already being used as a field mask", field.FieldName))
			//			e2 := tr.T.Args[0].T.PR.BeautifulError(fmt.Errorf("used as size here"))
			//			return tr, nil, tlast.BeautifulError2(e1, e2)
			//		}
			//		field.UsedAsSize = true
			//		field.UsedAsSizePR = tr.T.Args[0].T.PR
			//	}
			//}
		}
		ac, natArgs2, err := k.resolveTypeHybrid(ctxTL2, bracketType.ArrayType, leftArgs, actualArgs)
		if err != nil {
			return tr, nil, err
		}
		bracketType.ArrayType = ac
		tr.Type.BracketType = &bracketType
		natArgs = append(natArgs, natArgs2...)
		return tr, natArgs, nil
	}
	someType := &tr.Type.SomeType
	// names found in local arguments have priority over global type names
	if someType.Name.Namespace == "" {
		for i, targ := range leftArgs {
			if targ.FieldName == someType.Name.Name {
				for _, arg := range someType.Arguments {
					e1 := arg.PR.BeautifulError(fmt.Errorf("reference to template argument %s cannot have arguments", targ.FieldName))
					e2 := targ.PR.BeautifulError(fmt.Errorf("declared here"))
					return tr, nil, tlast.BeautifulError2(e1, e2)
				}
				actualArg := actualArgs[i]
				if actualArg.wrongTypeErr != nil {
					e1 := tr.PR.BeautifulError(fmt.Errorf("reference %q should be to #-param or # field", targ.FieldName))
					return tr, nil, tlast.BeautifulError2(e1, actualArg.wrongTypeErr)
				}
				actualArg.arg.OriginalArgumentName = targ.FieldName // TODO - check if this is enough
				actualArg.arg.PR = someType.PR
				// TODO - check if all necessary PRs are set
				// actualArg.arg.T.PRArgs = tr.T.PRArgs
				if actualArg.arg.IsNumber || actualArg.arg.Type.String() == "*" {
					if someType.Bare {
						e1 := someType.PR.BeautifulError(fmt.Errorf("reference to #-param %q cannot be bare", targ.FieldName))
						e2 := targ.PR.BeautifulError(fmt.Errorf("declared here"))
						return tr, nil, tlast.BeautifulError2(e1, e2)
					}
					return actualArg.arg, actualArg.natArgs, nil
				}
				if someType.Bare && !actualArg.arg.Type.IsBracket() {
					actualArg.arg.Type.SomeType.Bare = true
				}
				return actualArg.arg, actualArg.natArgs, nil
			}
		}
		// probably ref to global type or a typo
	}
	tName := someType.Name.String()
	kt, ok := k.tips[tName]
	if !ok {
		return tr, nil, someType.PR.BeautifulError(fmt.Errorf("type or argument reference %s not found", tName))
	}
	if kt.isFunction {
		return tr, nil, kt.functionCanNotBeReferencedError(someType.PR)
	}
	if _, ok := kt.tl1Names[tName]; !ok && !ctxTL2 {
		for good := range kt.tl1Names {
			return tr, nil, someType.PR.BeautifulError(fmt.Errorf("type %s is TL2-specific, in TL1 please use %s instead", tName, good))
		}
		return tr, nil, someType.PR.BeautifulError(fmt.Errorf("type %s is TL2-specific and cannot be used from TL1", tName))
	}
	if someType.Bare && kt.builtinWrappedCanonicalName != "" {
		tName = kt.builtinWrappedCanonicalName
		kt, ok = k.tips[tName]
		if !ok {
			return tr, nil, someType.PR.BeautifulError(fmt.Errorf("internal error: built-in wrapped type %s not found", tName))
		}
		someType.Name = tlast.TL2TypeName{Name: tName}
		someType.Bare = false // not required and should not change canonical type
		//if tName == "tuple" {
		//	tr.T.Args[0], tr.T.Args[1] = tr.T.Args[1], tr.T.Args[0]
		//}
	}

	if kt.originTL2 && !ctxTL2 {
		e1 := someType.PR.BeautifulError(fmt.Errorf("TL1 combinator cannot reference TL2 combinator %s", someType.Name))
		e2 := kt.combTL2.TypeDecl.PRName.BeautifulError(fmt.Errorf("declared here"))
		return tr, nil, tlast.BeautifulError2(e1, e2)
	}
	//td := kt.combTL1[0]
	// checks below are redundant, but they catch type resolve errors early for beautiful errors
	// if modifying this code, modify also code in func (k *Kernel) resolveArgumentTL2Impl()
	if len(kt.templateArguments) > len(someType.Arguments) {
		arg := kt.templateArguments[len(someType.Arguments)]
		e1 := someType.PRArguments.CollapseToEnd().BeautifulError(fmt.Errorf("missing template argument %q here", arg.Name))
		e2 := arg.PR.BeautifulError(fmt.Errorf("declared here"))
		return tr, nil, tlast.BeautifulError2(e1, e2)
	}
	if len(kt.templateArguments) < len(someType.Arguments) {
		arg := someType.Arguments[len(kt.templateArguments)]
		e1 := arg.PR.BeautifulError(errors.New("excess template argument here"))
		// TODO - TemplateArgumentsPR for TL2 types
		//e2 := kt.TemplateArgumentsPR.BeautifulError(fmt.Errorf("arguments declared here"))
		return tr, nil, e1 // tlast.BeautifulError2(e1, e2)
	}

	someType.Arguments = append([]tlast.TL2TypeArgument{}, someType.Arguments...) // preserve original
	var natArgs []ActualNatArg
	for i, arg := range someType.Arguments {
		ta := kt.templateArguments[i]
		rt, natArgs2, err := k.resolveArgumentHybrid(ctxTL2, arg, leftArgs, actualArgs)
		if err != nil {
			return tr, nil, err
		}
		someType.Arguments[i] = rt
		natArgs = append(natArgs, natArgs2...)

		argNat := rt.IsNumber || rt.Type.String() == "*"
		if ta.Category.IsNat() && !argNat {
			e1 := arg.PR.BeautifulError(errors.New("template argument must be # here"))
			e2 := ta.PR.BeautifulError(fmt.Errorf("argument declared here"))
			return tr, nil, tlast.BeautifulError2(e1, e2)
		}
		if !ta.Category.IsNat() && argNat {
			e1 := arg.PR.BeautifulError(errors.New("template argument must be Type here"))
			e2 := ta.PR.BeautifulError(fmt.Errorf("argument declared here"))
			return tr, nil, tlast.BeautifulError2(e1, e2)
		}
		if ta.Category.IsNat() {
			if rt.IsNumber {
				if kt.targs[i].usedAsNatConst == nil {
					kt.targs[i].usedAsNatConst = map[uint32]struct{}{}
				}
				kt.targs[i].usedAsNatConst[rt.Number] = struct{}{}
			} else {
				kt.targs[i].usedAsNatVariable = true
			}
		}
	}
	return tr, natArgs, nil
}

func (k *Kernel) getArgNamespace2(rt tlast.TL2TypeRef) string {
	argNamespaces := map[string]struct{}{}
	k.collectArgsNamespaces2(tlast.TL2TypeArgument{Type: rt}, argNamespaces)
	if len(argNamespaces) == 1 {
		for ns := range argNamespaces {
			return ns
		}
	}
	if rt.BracketType != nil {
		return ""
	}
	return rt.SomeType.Name.Namespace
}

func (k *Kernel) collectArgsNamespaces2(rt tlast.TL2TypeArgument, argNamespaces map[string]struct{}) {
	// This is policy. You can adjust it, so more or less templates instantiations
	// are moved into namespace of arguments. Code should compile anyway.
	if rt.IsNumber || rt.Type.String() == "*" {
		return
	}
	if br := rt.Type.BracketType; br != nil {
		k.collectArgsNamespaces2(tlast.TL2TypeArgument{Type: br.ArrayType}, argNamespaces)
		if br.HasIndex {
			k.collectArgsNamespaces2(br.IndexType, argNamespaces)
		}
	} else {
		someType := rt.Type.SomeType
		if someType.Name.Namespace != "" {
			argNamespaces[someType.Name.Namespace] = struct{}{}
		}
		for _, arg := range someType.Arguments {
			k.collectArgsNamespaces2(arg, argNamespaces)
		}
	}
}
