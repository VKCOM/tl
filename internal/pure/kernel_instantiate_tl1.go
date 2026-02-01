// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package pure

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/vkcom/tl/internal/tlast"
)

type LocalArg struct {
	wrongTypeErr *tlast.ParseError // we must add all field names to local context, because they must correctly shadow names outside, but we check the type
	arg          tlast.ArithmeticOrType
	natArgs      []ActualNatArg
}

// top level types do not have bare/boxed in their names, instead bare is returned from function
func (k *Kernel) canonicalStringTL1(tr tlast.TypeRef, top bool, allowFunctions bool) (_ string, bare bool, _ error) {
	var s strings.Builder

	tName := tr.Type.String()
	kt, ok := k.tips[tName]
	if !ok {
		return "", false, tr.PR.BeautifulError(fmt.Errorf("type reference %s not found", tName))
	}
	if kt.isFunction && !allowFunctions {
		e1 := tr.PR.BeautifulError(fmt.Errorf("function %s cannot be referenced", tName))
		if kt.originTL2 {
			// TODO - beautiful
			return "", false, e1
		}
		e2 := kt.combTL1[0].Construct.NamePR.BeautifulError(errSeeHere)
		return "", false, tlast.BeautifulError2(e1, e2)
	}
	// TODO - check TL1/TL2 references here
	//if kt.originTL2 {
	//	panic(fmt.Sprintf("canonical string tip %s not from TL1", tName))
	//}
	bare = tr.Bare
	if tr.Type != kt.tl1BoxedName {
		bare = true
	}
	if bare && !kt.CanBeBare() {
		// TODO - we could simply treat % as "bare if possible", which would allow writing it basically everywhere
		e1 := tr.PR.BeautifulError(fmt.Errorf("type reference to %s cannot be bare", tName))
		if kt.originTL2 {
			// TODO - beautiful
			return "", false, e1
		}
		e2 := kt.combTL1[0].TypeDecl.NamePR.BeautifulError(errSeeHere)
		return "", false, tlast.BeautifulError2(e1, e2)
	}
	if !bare && !kt.CanBeBoxed() { // TODO - impossible?
		e1 := tr.PR.BeautifulError(fmt.Errorf("type reference to %s cannot be boxed", tName))
		if kt.originTL2 {
			// TODO - beautiful
			return "", false, e1
		}
		e2 := kt.combTL1[0].Construct.NamePR.BeautifulError(errSeeHere)
		return "", false, tlast.BeautifulError2(e1, e2)
	}
	if !top && !bare && kt.CanBeBare() {
		s.WriteString("+")
	}
	s.WriteString(kt.canonicalName.String())
	if len(tr.Args) == 0 {
		return s.String(), bare, nil
	}
	s.WriteByte('<')
	for i, a := range tr.Args {
		if i != 0 {
			s.WriteByte(',')
		}
		if a.IsArith {
			s.WriteString(strconv.FormatUint(uint64(a.Arith.Res), 10))
		} else if a.T.Type.String() == "*" {
			s.WriteString("*")
		} else {
			str, _, err := k.canonicalStringTL1(a.T, false, false)
			if err != nil {
				return "", false, err
			}
			s.WriteString(str)
		}
	}
	s.WriteByte('>')
	return s.String(), bare, nil
}

func (k *Kernel) resolveTypeTL1(tr tlast.TypeRef, leftArgs []tlast.TemplateArgument,
	actualArgs []LocalArg) (tlast.TypeRef, []ActualNatArg, error) {
	ac, natArgs, err := k.resolveArgumentTL1(tlast.ArithmeticOrType{IsArith: false, T: tr}, leftArgs, actualArgs)
	if err != nil {
		return tr, nil, err
	}
	if ac.IsArith {
		// TODO - beautiful test case,
		return tr, nil, fmt.Errorf("type argument %s resolved to a number %d", tr, ac.Arith.Res)
	}
	if ac.T.String() == "*" {
		// TODO - beautiful test case,
		return tr, nil, tr.PR.BeautifulError(fmt.Errorf("type argument %s resolved to a nat argument %s", tr, ac.T))
	}
	return ac.T, natArgs, nil
}

func (k *Kernel) resolveArgumentTL1(tr tlast.ArithmeticOrType, leftArgs []tlast.TemplateArgument,
	actualArgs []LocalArg) (tlast.ArithmeticOrType, []ActualNatArg, error) {
	before := tr
	was := before.T.String()
	tr, natArgs, err := k.resolveArgumentTL1Impl(tr, leftArgs, actualArgs)
	after := before.T.String()
	if was != after {
		panic(fmt.Sprintf("tl2pure: internal error, resolveArgumentTL1 destroyed %s original value %s due to golang aliasing", after, was))
	}
	return tr, natArgs, err
}

func (k *Kernel) resolveArgumentTL1Impl(tr tlast.ArithmeticOrType, leftArgs []tlast.TemplateArgument,
	actualArgs []LocalArg) (tlast.ArithmeticOrType, []ActualNatArg, error) {
	if tr.IsArith {
		return tr, nil, nil
	}
	if tr.T.Type.String() == "*" {
		return tr, nil, fmt.Errorf("internal error: resolving type more than once")
	}
	// names found in local arguments have priority over global type names
	if tr.T.Type.Namespace == "" {
		for i, targ := range leftArgs {
			if targ.FieldName == tr.T.Type.Name {
				for _, arg := range tr.T.Args {
					e1 := arg.T.PR.BeautifulError(fmt.Errorf("reference to template argument %s cannot have arguments", targ.FieldName))
					e2 := targ.PR.BeautifulError(fmt.Errorf("declared here"))
					return tr, nil, tlast.BeautifulError2(e1, e2)
				}
				actualArg := actualArgs[i]
				if actualArg.wrongTypeErr != nil {
					e1 := tr.T.PR.BeautifulError(fmt.Errorf("reference %q should be to #-param or # field", targ.FieldName))
					return tr, nil, tlast.BeautifulError2(e1, actualArg.wrongTypeErr)
				}
				actualArg.arg.T.OriginalArgumentName = targ.FieldName // TODO - check if this is enough
				actualArg.arg.T.PR = tr.T.PR
				actualArg.arg.T.PRArgs = tr.T.PRArgs
				if actualArg.arg.IsArith || actualArg.arg.T.Type.String() == "*" {
					if tr.T.Bare {
						e1 := tr.T.PR.BeautifulError(fmt.Errorf("reference to #-param %q cannot be bare", targ.FieldName))
						e2 := targ.PR.BeautifulError(fmt.Errorf("declared here"))
						return tr, nil, tlast.BeautifulError2(e1, e2)
					}
					return actualArg.arg, actualArg.natArgs, nil
				}
				if tr.T.Bare {
					actualArg.arg.T.Bare = true
				}
				return actualArg.arg, actualArg.natArgs, nil
			}
		}
		// probably ref to global type or a typo
	}
	tr.T.Args = append([]tlast.ArithmeticOrType{}, tr.T.Args...) // preserve original
	var natArgs []ActualNatArg
	for i, arg := range tr.T.Args {
		rt, natArgs2, err := k.resolveArgumentTL1(arg, leftArgs, actualArgs)
		if err != nil {
			return tr, nil, err
		}
		tr.T.Args[i] = rt
		natArgs = append(natArgs, natArgs2...)
	}
	return tr, natArgs, nil
}

func (k *Kernel) resolveMaskTL1(mask tlast.FieldMask, leftArgs []tlast.TemplateArgument,
	actualArgs []LocalArg, combinatorField tlast.CombinatorField) (ActualNatArg, error) {
	for i, targ := range leftArgs {
		if targ.FieldName == mask.MaskName {
			actualArg := actualArgs[i]
			if actualArg.wrongTypeErr != nil {
				e1 := mask.PRName.BeautifulError(fmt.Errorf("reference %q should be to #-param or # field", targ.FieldName))
				return ActualNatArg{}, tlast.BeautifulError2(e1, actualArg.wrongTypeErr)
			}
			if !targ.IsNat {
				e1 := mask.PRName.BeautifulError(fmt.Errorf("fieldMask cannot reference Type-parameter %s", targ.FieldName))
				e2 := targ.PR.BeautifulError(fmt.Errorf("declared here"))
				return ActualNatArg{}, tlast.BeautifulError2(e1, e2)
			}
			if len(actualArg.natArgs) > 1 {
				return ActualNatArg{}, mask.PRName.BeautifulError(fmt.Errorf("internal error: fieldMask reference len(natArg) == %d for parameter %s", len(actualArg.natArgs), targ.FieldName))
			}
			if actualArg.arg.IsArith {
				return ActualNatArg{
					isNumber: true,
					number:   actualArg.arg.Arith.Res,
				}, nil
			}
			if sf := actualArg.arg.SourceField; sf != (tlast.CombinatorField{}) {
				field := &sf.Comb.Fields[sf.FieldIndex]
				if field.UsedAsSize {
					e3 := field.UsedAsSizePR.BeautifulError(fmt.Errorf("used as size here"))
					e3.PrintWarning(k.opts.ErrorWriter, nil)
					e1 := field.PRName.BeautifulError(fmt.Errorf("#-field %s is used as an field mask, while already being used as tuple size", field.FieldName))
					e2 := mask.PRName.BeautifulError(fmt.Errorf("used as mask here"))
					return ActualNatArg{}, tlast.BeautifulError2(e1, e2)
				}
				field.UsedAsMask = true
				field.UsedAsMaskPR = mask.PRName
				field.AffectedFields = append(field.AffectedFields, combinatorField)
			}
			return actualArg.natArgs[0], nil
		}
	}
	return ActualNatArg{}, mask.PRName.BeautifulError(errors.New("fieldMask reference not found"))
}

func (k *Kernel) GetInstanceTL1(tr tlast.TypeRef) (TypeInstance, bool, error) {
	ref, bare, err := k.getInstanceTL1(tr, false, true)
	if err != nil {
		return nil, false, err
	}
	return ref.ins, bare, nil
}

func (k *Kernel) getInstanceTL1(tr tlast.TypeRef, create bool, allowFunctions bool) (_ *TypeInstanceRef, bare bool, _ error) {
	canonicalName, bare, err := k.canonicalStringTL1(tr, true, allowFunctions)
	if err != nil {
		return nil, false, err
	}
	if ref, ok := k.instances[canonicalName]; ok {
		return ref, bare, nil
	}
	if !create {
		return nil, false, fmt.Errorf("internal error: instance %s must exist", canonicalName)
	}
	// log.Printf("creating an instance of type %s", canonicalName)
	tName := tr.Type.String()
	kt, ok := k.tips[tName]
	if !ok {
		return nil, false, fmt.Errorf("type %s does not exist", tr.Type)
	}
	// must store pointer before children GetInstance() terminates recursion
	// this instance stays not initialized in case of error, but kernel then is not consistent anyway
	ref := k.addInstance(canonicalName, kt)

	if kt.originTL2 {
		return nil, false, fmt.Errorf("TL1 combinator cannot reference TL2 combinator %s", tr.Type)
	}
	td := kt.combTL1[0]
	// checks below are redundant, but they catch type resolve errors early for beautiful errors
	if len(td.TemplateArguments) > len(tr.Args) {
		arg := td.TemplateArguments[len(tr.Args)]
		e1 := tr.PRArgs.CollapseToEnd().BeautifulError(fmt.Errorf("missing template argument %q here", arg.FieldName))
		e2 := arg.PR.BeautifulError(fmt.Errorf("declared here"))
		return nil, false, tlast.BeautifulError2(e1, e2)
	}
	if len(td.TemplateArguments) < len(tr.Args) {
		arg := tr.Args[len(td.TemplateArguments)]
		e1 := arg.T.PR.BeautifulError(errors.New("excess template argument here"))
		e2 := td.TemplateArgumentsPR.BeautifulError(fmt.Errorf("arguments declared here"))
		return nil, false, tlast.BeautifulError2(e1, e2)
	}
	for i, ta := range td.TemplateArguments {
		arg := tr.Args[i]
		argNat := arg.IsArith || arg.T.Type.String() == "*"
		if ta.IsNat && !argNat {
			e1 := arg.T.PR.BeautifulError(errors.New("template argument must be # here"))
			e2 := td.TemplateArgumentsPR.BeautifulError(fmt.Errorf("arguments declared here"))
			return nil, false, tlast.BeautifulError2(e1, e2)
		}
		if !ta.IsNat && argNat {
			e1 := arg.T.PR.BeautifulError(errors.New("template argument must be Type here"))
			e2 := td.TemplateArgumentsPR.BeautifulError(fmt.Errorf("arguments declared here"))
			return nil, false, tlast.BeautifulError2(e1, e2)
		}
	}
	isDict := false
	if k.opts.NewDicts &&
		strings.Contains(strings.ToLower(tName), "dictionary") && !kt.originTL2 &&
		len(kt.combTL1) == 1 {
		fieldT, fieldOk := k.tips[tName+"Field"]
		if fieldOk && !fieldT.originTL2 && len(fieldT.combTL1) == 1 &&
			len(fieldT.combTL1[0].TemplateArguments) != 0 &&
			len(fieldT.combTL1[0].TemplateArguments) == len(kt.combTL1[0].TemplateArguments) &&
			len(fieldT.combTL1[0].Fields) == 2 {
			// This only checks some type properties, they are enough for us for now
			isDict = true
			for i, targ := range kt.combTL1[0].TemplateArguments {
				farg := fieldT.combTL1[0].TemplateArguments[i]
				if targ.IsNat || farg.IsNat {
					isDict = false
				}
			}
			// log.Printf("creating an instance of dictionary type %s", canonicalName)
		}
	}
	switch {
	case tName == "__vector":
		ref.ins, err = k.createVectorTL1(canonicalName, tr, td.TemplateArguments, tr.Args)
	case tName == "__tuple":
		ref.ins, err = k.createTupleTL1(canonicalName, tr, td.TemplateArguments, tr.Args)
	case isDict:
		ref.ins, err = k.createDictTL1(canonicalName, kt, tr, td.TemplateArguments, tr.Args)
	case len(kt.combTL1) > 1:
		ref.ins, err = k.createUnionTL1FromTL1(canonicalName, kt, tr, kt.combTL1,
			td.TemplateArguments, tr.Args)
	case len(kt.combTL1) == 1:
		ref.ins, err = k.createStructTL1FromTL1(canonicalName, kt, tr,
			kt.combTL1[0],
			td.TemplateArguments, tr.Args,
			false, 0)
	default:
		return nil, false, fmt.Errorf("wrong type classification, internal error %s", canonicalName)
	}
	if err != nil {
		return nil, false, err
	}
	return ref, bare, nil
}

func (k *Kernel) getArgNamespace(rt tlast.TypeRef) string {
	argNamespaces := map[string]struct{}{}
	k.collectArgsNamespaces(tlast.ArithmeticOrType{T: rt}, argNamespaces)
	if len(argNamespaces) == 1 {
		for ns := range argNamespaces {
			return ns
		}
	}
	return rt.Type.Namespace
}

func (k *Kernel) collectArgsNamespaces(rt tlast.ArithmeticOrType, argNamespaces map[string]struct{}) {
	// This is policy. You can adjust it, so more or less templates instantiations
	// are moved into namespace of arguments. Code should compile anyway.
	if rt.IsArith || rt.T.Type.String() == "*" {
		return
	}
	if rt.T.Type.Namespace != "" {
		argNamespaces[rt.T.Type.Namespace] = struct{}{}
	}
	for _, arg := range rt.T.Args {
		k.collectArgsNamespaces(arg, argNamespaces)
	}
}

func (k *Kernel) createOrdinaryTypeTL1FromTL2(canonicalName string, definition []*tlast.Combinator,
	leftArgs []tlast.TemplateArgument, actualArgs []tlast.TL2TypeArgument) (TypeInstance, error) {

	switch {
	//case len(definition) > 1:
	//	return k.createUnion(canonicalName, definition.UnionType, leftArgs, actualArgs)
	//case definition[0].IsAlias():
	//	return k.createAlias(canonicalName, definition.TypeAlias, leftArgs, actualArgs)
	case len(definition) == 1:
		return k.createStructTL1FromTL2(canonicalName,
			definition[0].Fields,
			leftArgs, actualArgs,
			false, 0, nil)
	default:
		return nil, fmt.Errorf("wrong type classification, internal error %s", canonicalName)
	}
}
