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

type LocalArg struct {
	wrongTypeErr *tlast.ParseError // we must add all field names to local context, because they must correctly shadow names outside, but we check the type
	arg          tlast.TL2TypeArgument
	natArgs      []ActualNatArg
}

func (k *Kernel) resolveType(ctxTL2 bool, tr tlast.TL2TypeRef, leftArgs []tlast.TL2TypeTemplate,
	actualArgs []LocalArg) (tlast.TL2TypeRef, []ActualNatArg, error) {
	ac, natArgs, err := k.resolveArgument(ctxTL2, tlast.TL2TypeArgument{Type: tr}, leftArgs, actualArgs)
	if err != nil {
		return tr, nil, err
	}
	if ac.IsNumber {
		// TODO - beautiful test case,
		return tr, nil, tr.PR.BeautifulError(fmt.Errorf("type argument %s resolved to a number %d", tr, ac.Number))
	}
	if ac.Type.String() == "*" {
		// TODO - beautiful test case,
		return tr, nil, tr.PR.BeautifulError(fmt.Errorf("type argument %s resolved to a nat argument %s", tr, ac.Type))
	}
	return ac.Type, natArgs, nil
}

func (k *Kernel) resolveArgument(ctxTL2 bool, tr tlast.TL2TypeArgument, leftArgs []tlast.TL2TypeTemplate,
	actualArgs []LocalArg) (tlast.TL2TypeArgument, []ActualNatArg, error) {
	before := tr
	was := before.String()
	tr, natArgs, err := k.resolveArgumentImpl(ctxTL2, tr, leftArgs, actualArgs)
	after := before.String()
	if was != after {
		panic(fmt.Sprintf("tl2pure: internal error, resolveArgumentTL1 destroyed %s original value %s due to golang aliasing", after, was))
	}
	return tr, natArgs, err
}

func (k *Kernel) resolveArgumentImpl(ctxTL2 bool, tr tlast.TL2TypeArgument, leftArgs []tlast.TL2TypeTemplate,
	actualArgs []LocalArg) (tlast.TL2TypeArgument, []ActualNatArg, error) {
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
			ic, natArgs2, err := k.resolveArgument(ctxTL2, bracketType.IndexType, leftArgs, actualArgs)
			if err != nil {
				return tr, nil, err
			}
			bracketType.IndexType = ic
			natArgs = append(natArgs, natArgs2...)

			//if sf := ic.SourceField; sf != (tlast.CombinatorField{}) {
			//	field := &sf.Comb.Fields[sf.FieldIndex]
			//	if field.UsedAsMask {
			//		e3 := field.UsedAsMaskPR.BeautifulError(fmt.Errorf("used as mask here"))
			//		e3.PrintWarning(k.opts.ErrorWriter, nil)
			//		e1 := field.PRName.BeautifulError(fmt.Errorf("#-field %s is used as tuple size, while already being used as a field mask", field.FieldName))
			//		e2 := ic.PR.BeautifulError(fmt.Errorf("used as size here"))
			//		return tr, nil, tlast.BeautifulError2(e1, e2)
			//	}
			//	field.UsedAsSize = true
			//	field.UsedAsSizePR = ic.PR
			//}
			//if sf, ok := ic.SourceFieldAny.(CombinatorField); ok {
			//	field := &sf.Ins.fields[sf.FieldIndex]
			//	if field.UsedAsMask {
			//		e3 := field.UsedAsMaskPR.BeautifulError(fmt.Errorf("used as mask here"))
			//		e3.PrintWarning(k.opts.ErrorWriter, nil)
			//		e1 := field.PRName.BeautifulError(fmt.Errorf("#-field %s is used as tuple size, while already being used as a field mask", field.name))
			//		e2 := ic.PR.BeautifulError(fmt.Errorf("used as size here"))
			//		return tr, nil, tlast.BeautifulError2(e1, e2)
			//	}
			//	field.UsedAsSize = true
			//	field.UsedAsSizePR = ic.PR
			//}
		}
		ac, natArgs2, err := k.resolveType(ctxTL2, bracketType.ArrayType, leftArgs, actualArgs)
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
			if targ.Name == someType.Name.Name {
				for _, arg := range someType.Arguments {
					e1 := arg.PR.BeautifulError(fmt.Errorf("reference to template argument %s cannot have arguments", targ.Name))
					e2 := targ.PR.BeautifulError(fmt.Errorf("declared here"))
					return tr, nil, tlast.BeautifulError2(e1, e2)
				}
				actualArg := actualArgs[i]
				if actualArg.wrongTypeErr != nil {
					e1 := tr.PR.BeautifulError(fmt.Errorf("reference %q should be to #-param or # field", targ.Name))
					return tr, nil, tlast.BeautifulError2(e1, actualArg.wrongTypeErr)
				}
				actualArg.arg.OriginalArgumentName = targ.Name // TODO - check if this is enough
				actualArg.arg.PR = someType.PR
				// TODO - check if all necessary PRs are set
				// actualArg.arg.T.PRArgs = tr.T.PRArgs
				if actualArg.arg.IsNumber || actualArg.arg.Type.String() == "*" {
					if someType.Bare {
						e1 := someType.PR.BeautifulError(fmt.Errorf("reference to #-param %q cannot be bare", targ.Name))
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
	}

	if kt.originTL2 && !ctxTL2 {
		// TODO - this is wrong, add hasTL1, hasTL2 properties instead
		e1 := someType.PR.BeautifulError(fmt.Errorf("TL1 combinator cannot reference TL2 combinator %s", someType.Name))
		e2 := kt.combTL2.TypeDecl.PRName.BeautifulError(fmt.Errorf("declared here"))
		return tr, nil, tlast.BeautifulError2(e1, e2)
	}
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
		rt, natArgs2, err := k.resolveArgument(ctxTL2, arg, leftArgs, actualArgs)
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

func (k *Kernel) resolveMaskTL1(mask tlast.FieldMask, leftArgs []tlast.TL2TypeTemplate,
	actualArgs []LocalArg, combinatorField tlast.CombinatorField, combinatorField2 CombinatorField) (ActualNatArg, error) {
	for i, targ := range leftArgs {
		if targ.Name == mask.MaskName {
			actualArg := actualArgs[i]
			if actualArg.wrongTypeErr != nil {
				e1 := mask.PRName.BeautifulError(fmt.Errorf("reference %q should be to #-param or # field", targ.Name))
				return ActualNatArg{}, tlast.BeautifulError2(e1, actualArg.wrongTypeErr)
			}
			if !targ.Category.IsNatValue {
				e1 := mask.PRName.BeautifulError(fmt.Errorf("fieldMask cannot reference Type-parameter %s", targ.Name))
				e2 := targ.PR.BeautifulError(fmt.Errorf("declared here"))
				return ActualNatArg{}, tlast.BeautifulError2(e1, e2)
			}
			if len(actualArg.natArgs) > 1 {
				return ActualNatArg{}, mask.PRName.BeautifulError(fmt.Errorf("internal error: fieldMask reference len(natArg) == %d for parameter %s", len(actualArg.natArgs), targ.Name))
			}
			if actualArg.arg.IsNumber {
				return ActualNatArg{
					isNumber: true,
					number:   actualArg.arg.Number,
				}, nil
			}
			//if sf := actualArg.arg.SourceField; sf != (tlast.CombinatorField{}) {
			//	field := &sf.Comb.Fields[sf.FieldIndex]
			//	if field.UsedAsSize {
			//		e3 := field.UsedAsSizePR.BeautifulError(fmt.Errorf("used as size here"))
			//		e3.PrintWarning(k.opts.ErrorWriter, nil)
			//		e1 := field.PRName.BeautifulError(fmt.Errorf("#-field %s is used as an field mask, while already being used as tuple size", field.FieldName))
			//		e2 := mask.PRName.BeautifulError(fmt.Errorf("used as mask here"))
			//		return ActualNatArg{}, tlast.BeautifulError2(e1, e2)
			//	}
			//	field.UsedAsMask = true
			//	field.UsedAsMaskPR = mask.PRName
			//	field.AffectedFields = append(field.AffectedFields, combinatorField)
			//}
			//if sf, ok := actualArg.arg.SourceFieldAny.(CombinatorField); ok {
			//	field := &sf.Ins.fields[sf.FieldIndex]
			//	if field.UsedAsSize {
			//		e3 := field.UsedAsSizePR.BeautifulError(fmt.Errorf("used as size here"))
			//		e3.PrintWarning(k.opts.ErrorWriter, nil)
			//		e1 := field.PRName.BeautifulError(fmt.Errorf("#-field %s is used as an field mask, while already being used as tuple size", field.name))
			//		e2 := mask.PRName.BeautifulError(fmt.Errorf("used as mask here"))
			//		return ActualNatArg{}, tlast.BeautifulError2(e1, e2)
			//	}
			//	field.UsedAsMask = true
			//	field.UsedAsMaskPR = mask.PRName
			//	// access safe, due to check before calling this function
			//	field.AffectedFields[mask.BitNumber] = append(field.AffectedFields[mask.BitNumber], combinatorField2)
			//}
			return actualArg.natArgs[0], nil
		}
	}
	return ActualNatArg{}, mask.PRName.BeautifulError(errors.New("fieldMask reference not found"))
}
