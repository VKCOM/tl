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
	if kt.builtinWrappedCanonicalName != "" {
		return tr, someType.PRName.BeautifulError(fmt.Errorf("TL2 type reference must be to built-in canonical name %s", kt.builtinWrappedCanonicalName))
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

// TODO - we can decide later to convert TL1 type refs to TL2, should weight pro and cons
func (k *Kernel) convertTL2TypeRefToTL1(tr tlast.TL2TypeRef) (tlast.TypeRef, error) {
	if tr.IsBracket() {
		elemType, err := k.convertTL2TypeRefToTL1(tr.BracketType.ArrayType)
		if err != nil {
			return tlast.TypeRef{}, err
		}
		if tr.BracketType.HasIndex {
			if tr.BracketType.IndexType.IsNumber {
				// tuple
				rt := tlast.TypeRef{
					Type:   tlast.Name{Name: "__tuple"},
					Args:   nil,
					Bare:   false,
					PR:     tr.BracketType.PR,
					PRArgs: tr.BracketType.PR, // TODO
					// OriginalArgumentName: , TODO
				}
				rt.Args = append(rt.Args, tlast.ArithmeticOrType{
					IsArith: true,
					Arith:   tlast.Arithmetic{Res: tr.BracketType.IndexType.Number},
					T:       tlast.TypeRef{PR: tr.BracketType.IndexType.PR},
					// SourceField: tlast.CombinatorField{}, // TODO
				})
				rt.Args = append(rt.Args, tlast.ArithmeticOrType{
					IsArith: false,
					T:       elemType,
					// SourceField: tlast.CombinatorField{}, // TODO
				})
				return rt, nil
			}
			// dict
			panic("TODO - dict")
		}
		// vector
		rt := tlast.TypeRef{
			Type:   tlast.Name{Name: "__vector"},
			Args:   nil,
			Bare:   false,
			PR:     tr.BracketType.PR,
			PRArgs: tr.BracketType.PR, // TODO
			// OriginalArgumentName: , TODO
		}
		rt.Args = append(rt.Args, tlast.ArithmeticOrType{
			IsArith: false,
			T:       elemType,
			// SourceField: tlast.CombinatorField{}, // TODO
		})
		return rt, nil
	}
	someType := tr.SomeType
	tName := someType.Name.String()
	kt, ok := k.tips[tName]
	if !ok {
		return tlast.TypeRef{}, someType.PRName.BeautifulError(fmt.Errorf("type or argument reference %s not found", tName))
	}
	rt := tlast.TypeRef{
		Type:   tlast.Name(someType.Name),
		Args:   nil,
		Bare:   kt.canBeBare,
		PR:     someType.PR,
		PRArgs: someType.PR, // TODO
		// OriginalArgumentName: , TODO
	}
	for _, arg := range someType.Arguments {
		if arg.IsNumber {
			rt.Args = append(rt.Args, tlast.ArithmeticOrType{
				IsArith: true,
				Arith:   tlast.Arithmetic{Res: arg.Number},
				T:       tlast.TypeRef{PR: arg.PR},
				// SourceField: tlast.CombinatorField{}, // TODO
			})
			continue
		}
		argType, err := k.convertTL2TypeRefToTL1(arg.Type)
		if err != nil {
			return tlast.TypeRef{}, err
		}
		rt.Args = append(rt.Args, tlast.ArithmeticOrType{
			IsArith: false,
			T:       argType,
			// SourceField: tlast.CombinatorField{}, // TODO
		})
	}
	return rt, nil
}

func (k *Kernel) getInstanceTL2(tr tlast.TL2TypeRef, create bool) (*TypeInstanceRef, error) {
	trTL1, err := k.convertTL2TypeRefToTL1(tr)
	if err != nil {
		return nil, err
	}
	canonicalName, _, err := k.canonicalStringTL1(trTL1, true)
	if err != nil {
		return nil, err
	}
	if ref, ok := k.instances[canonicalName]; ok {
		return ref, nil
	}
	if !create {
		return nil, fmt.Errorf("internal error: instance %s must exist", canonicalName)
	}
	if tr.IsBracket() {
		log.Printf("creating a bracket instance of type %s", canonicalName)
		// must store pointer before children GetInstanceTL2() terminates recursion
		// this instance stays not initialized in case of error, but kernel then is not consistent anyway
		ref := k.addInstance(canonicalName, k.brackets)

		elemInstance, err := k.getInstanceTL2(tr.BracketType.ArrayType, true)
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
			keyInstance, err := k.getInstanceTL2(tr.BracketType.IndexType.Type, true)
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
	// must store pointer before children GetInstanceTL2() terminates recursion
	// this instance stays mpt initialized in case of error, but kernel then is not consistent anyway
	someType := tr.SomeType
	kt, ok := k.tips[someType.Name.String()]
	if !ok {
		return nil, someType.PRName.BeautifulError(fmt.Errorf("type %s does not exist", someType.Name))
	}
	if kt.originTL2 {
		// must store pointer before children GetInstanceTL2() terminates recursion
		// this instance stays not initialized in case of error, but kernel then is not consistent anyway
		ref := k.addInstance(canonicalName, kt)
		if !kt.combTL2.IsFunction {
			if len(kt.combTL2.TypeDecl.TemplateArguments) != len(someType.Arguments) {
				return nil, fmt.Errorf("typeref to %s must have %d template arguments, has %d", canonicalName, len(kt.combTL2.TypeDecl.TemplateArguments), len(someType.Arguments))
			}
			ref.ins, err = k.createOrdinaryType(canonicalName, kt, trTL1, kt.canonicalName, kt.combTL2.TypeDecl.Magic,
				kt.combTL2.TypeDecl.Type, kt.combTL2.TypeDecl.TemplateArguments, someType.Arguments)
			if err != nil {
				return nil, err
			}
			return ref, nil
		}
		// TODO - should we check template arguments, as above?
		funcDecl := kt.combTL2.FuncDecl
		resultTlName := kt.canonicalName
		resultTlName.Name += "__Result"
		resultType, err := k.createOrdinaryType(canonicalName, kt, trTL1, resultTlName, 0, funcDecl.ReturnType, nil, nil)
		if err != nil {
			return nil, err
		}
		ref.ins, err = k.createStruct(canonicalName, kt, trTL1, kt.canonicalName, funcDecl.Magic, true,
			tlast.TL2TypeRef{}, funcDecl.Arguments, nil, nil, false, 0,
			resultType)
		if err != nil {
			return nil, err
		}
		return ref, nil
	}
	ref, _, err := k.getInstanceTL1(trTL1, create)
	return ref, err
}

// alias || fields || union
func (k *Kernel) createOrdinaryType(canonicalName string, tip *KernelType, trTL1 tlast.TypeRef,
	tlName tlast.Name, tlTag uint32,
	definition tlast.TL2TypeDefinition,
	leftArgs []tlast.TL2TypeTemplate, actualArgs []tlast.TL2TypeArgument) (TypeInstance, error) {

	switch {
	case definition.IsAlias():
		return k.createAlias(canonicalName, tip, trTL1, definition.TypeAlias, leftArgs, actualArgs)
	case definition.StructType.IsUnionType:
		return k.createUnion(canonicalName, tip, trTL1, definition.StructType.UnionType, leftArgs, actualArgs)
	default:
		return k.createStruct(canonicalName, tip, trTL1,
			tlName, tlTag,
			true, definition.TypeAlias, definition.StructType.ConstructorFields,
			leftArgs, actualArgs,
			false, 0, nil)
	}
}
