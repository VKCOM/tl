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

	"github.com/vkcom/tl/internal/tlast"
)

func (k *Kernel) resolveTypeTL2(tr tlast.TL2TypeRef, leftArgs []tlast.TL2TypeTemplate,
	actualArgs []tlast.TL2TypeArgument) (tlast.TL2TypeRef, error) {
	ac, err := k.resolveArgumentTL2(tlast.TL2TypeArgument{Type: tr}, leftArgs, actualArgs)
	if err != nil {
		return tr, err
	}
	if ac.IsNumber {
		return tr, fmt.Errorf("type argument %s resolved to a number %d", tr, ac.Number)
	}
	return ac.Type, nil
}

func (k *Kernel) resolveArgumentTL2(tr tlast.TL2TypeArgument, leftArgs []tlast.TL2TypeTemplate,
	actualArgs []tlast.TL2TypeArgument) (tlast.TL2TypeArgument, error) {
	before := tr
	was := before.Type.String()
	tr, err := k.resolveArgumentTL2Impl(tr, leftArgs, actualArgs)
	after := before.Type.String()
	if was != after {
		panic(fmt.Sprintf("tl2pure: internal error, resolveArgumentTL2 destroyed %s original value %s due to golang aliasing", after, was))
	}
	return tr, err
}

func (k *Kernel) resolveArgumentTL2Impl(tr tlast.TL2TypeArgument, leftArgs []tlast.TL2TypeTemplate,
	actualArgs []tlast.TL2TypeArgument) (tlast.TL2TypeArgument, error) {
	if tr.IsNumber {
		return tr, nil
	}
	if tr.Type.IsBracket() {
		bracketType := *tr.Type.BracketType
		if bracketType.HasIndex {
			ic, err := k.resolveArgumentTL2(bracketType.IndexType, leftArgs, actualArgs)
			if err != nil {
				return tr, err
			}
			bracketType.IndexType = ic
		}
		ac, err := k.resolveTypeTL2(bracketType.ArrayType, leftArgs, actualArgs)
		if err != nil {
			return tr, err
		}
		bracketType.ArrayType = ac
		tr.Type.BracketType = &bracketType
		return tr, nil
	}
	// names found in local arguments have priority over global type names
	someType := tr.Type.SomeType
	if someType.Name.Namespace == "" {
		for i, targ := range leftArgs {
			if targ.Name == someType.Name.Name {
				for _, arg := range someType.Arguments {
					e1 := arg.PR.BeautifulError(fmt.Errorf("reference to template argument %s cannot have arguments", targ.Name))
					e2 := targ.PR.BeautifulError(fmt.Errorf("declared here"))
					return tr, tlast.BeautifulError2(e1, e2)
				}
				return actualArgs[i], nil
			}
		}
		// probably ref to global type or a typo
	}
	tName := someType.Name.String()
	kt, ok := k.tips[tName]
	if !ok {
		return tr, someType.PRName.BeautifulError(fmt.Errorf("type or argument reference %s not found", tName))
	}
	if kt.isFunction {
		return tr, kt.functionCanNotBeReferencedError(someType.PRName)
	}
	if kt.canonicalName != tlast.Name(someType.Name) {
		return tr, someType.PRName.BeautifulError(fmt.Errorf("TL2 type reference must be to canonical name %s", kt.canonicalName))
	}
	someType.Arguments = append([]tlast.TL2TypeArgument{}, someType.Arguments...) // preserve original
	if kt.originTL2 {
		td := kt.combTL2
		if len(td.TypeDecl.TemplateArguments) > len(someType.Arguments) {
			arg := td.TypeDecl.TemplateArguments[len(someType.Arguments)]
			e1 := someType.PRArguments.CollapseToEnd().BeautifulError(fmt.Errorf("missing template argument %q here", arg.Name))
			e2 := arg.PR.BeautifulError(fmt.Errorf("declared here"))
			return tr, tlast.BeautifulError2(e1, e2)
		}
		if len(td.TypeDecl.TemplateArguments) < len(someType.Arguments) {
			arg := someType.Arguments[len(td.TypeDecl.TemplateArguments)]
			e1 := arg.PR.BeautifulError(errors.New("excess template argument here"))
			// TODO - we use TemplateArgumentsPR in TL1, may be need to make one also
			e2 := td.TypeDecl.PRName.CollapseToEnd().BeautifulError(fmt.Errorf("arguments declared here"))
			return tr, tlast.BeautifulError2(e1, e2)
		}
		for i, arg := range someType.Arguments {
			ta := td.TypeDecl.TemplateArguments[i]
			rt, err := k.resolveArgumentTL2(arg, leftArgs, actualArgs)
			if err != nil {
				return tr, err
			}
			someType.Arguments[i] = rt

			if ta.Category.IsNatValue && !rt.IsNumber {
				e1 := arg.PR.BeautifulError(errors.New("template argument must be # here"))
				e2 := ta.PR.BeautifulError(fmt.Errorf("argument declared here"))
				return tr, tlast.BeautifulError2(e1, e2)
			}
			if !ta.Category.IsNatValue && rt.IsNumber {
				e1 := arg.PR.BeautifulError(errors.New("template argument must be Type here"))
				e2 := ta.PR.BeautifulError(fmt.Errorf("argument declared here"))
				return tr, tlast.BeautifulError2(e1, e2)
			}
			if rt.IsNumber {
				if kt.targs[i].usedAsNatConst == nil {
					kt.targs[i].usedAsNatConst = map[uint32]struct{}{}
				}
				kt.targs[i].usedAsNatConst[rt.Number] = struct{}{}
			}
		}
	} else {
		td := kt.combTL1[0]
		// checks below are redundant, but they catch type resolve errors early for beautiful errors
		// if modifying this code, modify also code in func (k *Kernel) resolveArgumentTL1Impl()
		if len(td.TemplateArguments) > len(someType.Arguments) {
			arg := td.TemplateArguments[len(someType.Arguments)]
			e1 := someType.PRArguments.CollapseToEnd().BeautifulError(fmt.Errorf("missing template argument %q here", arg.FieldName))
			e2 := arg.PR.BeautifulError(fmt.Errorf("declared here"))
			return tr, tlast.BeautifulError2(e1, e2)
		}
		if len(td.TemplateArguments) < len(someType.Arguments) {
			arg := someType.Arguments[len(td.TemplateArguments)]
			e1 := arg.PR.BeautifulError(errors.New("excess template argument here"))
			e2 := td.TemplateArgumentsPR.BeautifulError(fmt.Errorf("arguments declared here"))
			return tr, tlast.BeautifulError2(e1, e2)
		}
		for i, arg := range someType.Arguments {
			ta := td.TemplateArguments[i]
			rt, err := k.resolveArgumentTL2(arg, leftArgs, actualArgs)
			if err != nil {
				return tr, err
			}
			someType.Arguments[i] = rt

			if ta.IsNat && !rt.IsNumber {
				e1 := arg.PR.BeautifulError(errors.New("template argument must be # here"))
				e2 := ta.PR.BeautifulError(fmt.Errorf("argument declared here"))
				return tr, tlast.BeautifulError2(e1, e2)
			}
			if !ta.IsNat && rt.IsNumber {
				e1 := arg.PR.BeautifulError(errors.New("template argument must be Type here"))
				e2 := ta.PR.BeautifulError(fmt.Errorf("argument declared here"))
				return tr, tlast.BeautifulError2(e1, e2)
			}
			if rt.IsNumber {
				if kt.targs[i].usedAsNatConst == nil {
					kt.targs[i].usedAsNatConst = map[uint32]struct{}{}
				}
				kt.targs[i].usedAsNatConst[rt.Number] = struct{}{}
			}
		}
	}
	tr.Type.SomeType = someType
	return tr, nil
}

func (k *Kernel) GetInstance(tr tlast.TL2TypeRef) (*TypeInstanceRef, error) {
	canonicalName := tr.String()
	if ref, ok := k.instances[canonicalName]; ok {
		return ref, nil
	}
	if tr.IsBracket() {
		log.Printf("creating a bracket instance of type %s", canonicalName)
		// must store pointer before children GetInstance() terminates recursion
		// this instance stays not initialized in case of error, but kernel then is not consistent anyway
		ref := k.addInstance(canonicalName, k.brackets)

		elemInstance, err := k.GetInstance(tr.BracketType.ArrayType)
		if err != nil {
			return nil, err
		}
		if tr.BracketType.HasIndex {
			if tr.BracketType.IndexType.IsNumber {
				// tuple
				ref.ins = k.createArray(canonicalName, true, tr.BracketType.IndexType.Number, elemInstance)
				return ref, nil
			}
			// dict
			keyInstance, err := k.GetInstance(tr.BracketType.IndexType.Type)
			if err != nil {
				return nil, err
			}
			if !keyInstance.ins.GoodForMapKey() {
				return nil, fmt.Errorf("type %s is not allowed as a map key (only 'bool', integers and 'string' allowed)", keyInstance.ins.CanonicalName())
			}
			ref.ins = k.createDict(canonicalName, keyInstance, elemInstance)
			return ref, nil
		}
		// vector
		ref.ins = k.createArray(canonicalName, false, 0, elemInstance)
		return ref, nil
	}
	// log.Printf("creating an instance of type %s", canonicalName)
	// must store pointer before children GetInstance() terminates recursion
	// this instance stays mpt initialized in case of error, but kernel then is not consistent anyway
	someType := tr.SomeType
	kt, ok := k.tips[someType.Name.String()]
	if !ok {
		return nil, fmt.Errorf("type %s does not exist", someType.Name)
	}
	// must store pointer before children GetInstance() terminates recursion
	// this instance stays not initialized in case of error, but kernel then is not consistent anyway
	ref := k.addInstance(canonicalName, kt)

	var err error
	if kt.originTL2 {
		if !kt.combTL2.IsFunction {
			if len(kt.combTL2.TypeDecl.TemplateArguments) != len(someType.Arguments) {
				return nil, fmt.Errorf("typeref to %s must have %d template arguments, has %d", canonicalName, len(kt.combTL2.TypeDecl.TemplateArguments), len(someType.Arguments))
			}
			ref.ins, err = k.createOrdinaryType(canonicalName, kt, kt.combTL2.TypeDecl.Type, kt.combTL2.TypeDecl.TemplateArguments, someType.Arguments)
			if err != nil {
				return nil, err
			}
			return ref, nil
		}
		// TODO - should we check template arguments, as above?
		funcDecl := kt.combTL2.FuncDecl
		resultType, err := k.createOrdinaryType(canonicalName, kt, funcDecl.ReturnType, nil, nil)
		if err != nil {
			return nil, err
		}
		ref.ins, err = k.createStruct(canonicalName, kt, true,
			tlast.TL2TypeRef{}, funcDecl.Arguments, nil, nil, false, 0,
			resultType)
		if err != nil {
			return nil, err
		}
		return ref, nil
	}
	comb := kt.combTL1[0]
	if !comb.IsFunction {
		if len(comb.TemplateArguments) != len(someType.Arguments) {
			return nil, fmt.Errorf("typeref to %s must have %d template arguments, has %d", canonicalName, len(comb.TemplateArguments), len(someType.Arguments))
		}
		ref.ins, err = k.createOrdinaryTypeTL1FromTL2(canonicalName, kt.combTL1, comb.TemplateArguments, someType.Arguments)
		if err != nil {
			return nil, err
		}
		return ref, nil
	}
	return nil, fmt.Errorf("TODO - function from TL1 not yet supported")
}

// alias || fields || union
func (k *Kernel) createOrdinaryType(canonicalName string, tip *KernelType, definition tlast.TL2TypeDefinition,
	leftArgs []tlast.TL2TypeTemplate, actualArgs []tlast.TL2TypeArgument) (TypeInstance, error) {

	switch {
	case definition.IsAlias():
		return k.createAlias(canonicalName, tip, definition.TypeAlias, leftArgs, actualArgs)
	case definition.StructType.IsUnionType:
		return k.createUnion(canonicalName, tip, definition.StructType.UnionType, leftArgs, actualArgs)
	default:
		return k.createStruct(canonicalName, tip,
			true, definition.TypeAlias, definition.StructType.ConstructorFields,
			leftArgs, actualArgs,
			false, 0, nil)
	}
}
