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

func (k *Kernel) resolveType(tr tlast.TL2TypeRef, leftArgs []tlast.TL2TypeTemplate,
	actualArgs []tlast.TL2TypeArgument) (tlast.TL2TypeRef, error) {
	ac, err := k.resolveArgument(tlast.TL2TypeArgument{Type: tr}, leftArgs, actualArgs)
	if err != nil {
		return tr, err
	}
	if ac.IsNumber {
		return tr, fmt.Errorf("type argument %s resolved to a number %d", tr, ac.Number)
	}
	return ac.Type, nil
}

func (k *Kernel) resolveArgument(tr tlast.TL2TypeArgument, leftArgs []tlast.TL2TypeTemplate,
	actualArgs []tlast.TL2TypeArgument) (tlast.TL2TypeArgument, error) {
	before := tr
	was := before.Type.String()
	tr, err := k.resolveArgumentImpl(tr, leftArgs, actualArgs)
	after := before.Type.String()
	if was != after {
		panic(fmt.Sprintf("tl2pure: internal error, resolveArgument destroyed %s original value %s due to golang aliasing", after, was))
	}
	return tr, err
}

func (k *Kernel) resolveArgumentImpl(tr tlast.TL2TypeArgument, leftArgs []tlast.TL2TypeTemplate,
	actualArgs []tlast.TL2TypeArgument) (tlast.TL2TypeArgument, error) {
	if tr.IsNumber {
		return tr, nil
	}
	if tr.Type.IsBracket {
		bracketType := *tr.Type.BracketType
		if bracketType.HasIndex {
			ic, err := k.resolveArgument(bracketType.IndexType, leftArgs, actualArgs)
			if err != nil {
				return tr, err
			}
			bracketType.IndexType = ic
		}
		ac, err := k.resolveType(bracketType.ArrayType, leftArgs, actualArgs)
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
				if len(someType.Arguments) != 0 {
					return tr, fmt.Errorf("reference to template argument %s cannot have arguments", targ.Name)
				}
				return actualArgs[i], nil
			}
		}
		// probably ref to global type or a typo
	}
	someType.Arguments = append([]tlast.TL2TypeArgument{}, someType.Arguments...) // preserve original
	for i, arg := range someType.Arguments {
		rt, err := k.resolveArgument(arg, leftArgs, actualArgs)
		if err != nil {
			return tr, err
		}
		someType.Arguments[i] = rt
	}
	tr.Type.SomeType = someType
	return tr, nil
}

func (k *Kernel) getInstance(tr tlast.TL2TypeRef) (*TypeInstanceRef, error) {
	canonicalName := tr.String()
	if ref, ok := k.instances[canonicalName]; ok {
		return ref, nil
	}
	if tr.IsBracket {
		log.Printf("creating a bracket instance of type %s", canonicalName)
		// must store pointer before children getInstance() terminates recursion
		// this instance stays not initialized in case of error, but kernel then is not consistent anyway
		ref := k.addInstance(canonicalName, k.brackets)

		elemInstance, err := k.getInstance(tr.BracketType.ArrayType)
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
			keyInstance, err := k.getInstance(tr.BracketType.IndexType.Type)
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
	log.Printf("creating an instance of type %s", canonicalName)
	// must store pointer before children getInstance() terminates recursion
	// this instance stays mpt initialized in case of error, but kernel then is not consistent anyway
	someType := tr.SomeType
	kt, ok := k.tips[someType.Name.String()]
	if !ok {
		return nil, fmt.Errorf("type %s does not exist", someType.Name)
	}
	// must store pointer before children getInstance() terminates recursion
	// this instance stays not initialized in case of error, but kernel then is not consistent anyway
	ref := k.addInstance(canonicalName, kt)

	var err error
	if kt.originTL2 {
		if !kt.combTL2.IsFunction {
			if len(kt.combTL2.TypeDecl.TemplateArguments) != len(someType.Arguments) {
				return nil, fmt.Errorf("typeref to %s must have %d template arguments, has %d", canonicalName, len(kt.combTL2.TypeDecl.TemplateArguments), len(someType.Arguments))
			}
			ref.ins, err = k.createOrdinaryType(canonicalName, kt.combTL2.TypeDecl.Type, kt.combTL2.TypeDecl.TemplateArguments, someType.Arguments)
			if err != nil {
				return nil, err
			}
			return ref, nil
		}
		funcDecl := kt.combTL2.FuncDecl
		resultType, err := k.createOrdinaryType(canonicalName, funcDecl.ReturnType, nil, nil)
		if err != nil {
			return nil, err
		}
		ref.ins, err = k.createStruct(canonicalName, true,
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
func (k *Kernel) createOrdinaryType(canonicalName string, definition tlast.TL2TypeDefinition,
	leftArgs []tlast.TL2TypeTemplate, actualArgs []tlast.TL2TypeArgument) (TypeInstance, error) {

	switch {
	case definition.IsUnionType:
		return k.createUnion(canonicalName, definition.UnionType, leftArgs, actualArgs)
	case definition.IsAlias():
		return k.createAlias(canonicalName, definition.TypeAlias, leftArgs, actualArgs)
	case definition.IsConstructorFields:
		return k.createStruct(canonicalName,
			true, definition.TypeAlias, definition.ConstructorFields,
			leftArgs, actualArgs,
			false, 0, nil)
	default:
		return nil, fmt.Errorf("wrong type classification, internal error %s", canonicalName)
	}
}
