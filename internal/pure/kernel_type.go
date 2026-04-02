// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package pure

import (
	"fmt"
	"strings"

	"github.com/VKCOM/tl/internal/tlast"
)

type KernelTypeTarg struct {
	// this is set during type resolution, so the information
	// about argument references not erased from the type
	usedAsNatVariable bool
	usedAsNatConst    map[uint32]struct{}
}

type KernelType struct {
	originTL2 bool
	builtin   bool
	combTL1   []*tlast.Combinator
	combTL2   tlast.TL2Combinator
	// index by canonical name
	instances map[string]*TypeInstanceRef
	// order of instantiation
	instancesOrdered []*TypeInstanceRef

	isTopLevel     bool
	isFunction     bool                    // to prohibit references
	exclamationArg *tlast.TemplateArgument // empty if no such arg
	annotations    []string

	tl1Names map[string]struct{}
	tl2Names map[string]struct{}

	builtinWrappedCanonicalName string // for 'int#XXXX ? = Int'; this is 'int'

	// for TL2-defined types, simply name of combinator
	// for TL2 dictionary element, __dict_elem
	// for TL1-defined types, if !function, TypeDecl.Name (right name)
	// for TL1-defined types, if function, Constructor.Name (left name)
	// for primitive types, TL2 name (int32, uint32, etc).
	canonicalName tlast.TL2TypeName
	// for TL1-defined types, if !function, TypeDecl.Name (right side of =)
	// for TL1 Bool, Bool
	tl1BoxedName tlast.TL2TypeName
	canBeBare    bool

	namePR tlast.PositionRange // for "see here" in beautiful errors

	// usage tracking for migration/compilation.
	// common for union types, so cannot be in combinator itself
	templateArguments []tlast.TL2TypeTemplate
	targs             []KernelTypeTarg
}

func (t *KernelType) OriginTL2() bool {
	return t.originTL2
}

func (t *KernelType) IsFunction() bool {
	return t.isFunction
}

// Very few functions wrap another function. This is for TL1 only, TL2 does not have this complexity.
// Currently only PHP generator needs them, because parsing of invokeReq is split between runtime and user code.
// All other generators should skip/ignore such functions.
func (t *KernelType) IsExclamationWrapper() bool {
	return t.exclamationArg != nil
}

func (t *KernelType) Annotations() []string {
	return t.annotations
}

func (t *KernelType) HasAnnotation(a string) bool {
	for _, an := range t.annotations {
		if an == a {
			return true
		}
	}
	return false
}

func (t *KernelType) CanonicalName() tlast.TL2TypeName {
	return t.canonicalName
}

func (t *KernelType) CanBeBare() bool {
	return t.canBeBare
}

func (t *KernelType) CanBeBoxed() bool {
	return t.tl1BoxedName != tlast.TL2TypeName{}
}

func (t *KernelType) functionCanNotBeReferencedError(PR tlast.PositionRange) error {
	e1 := PR.BeautifulError(fmt.Errorf("function %s cannot be referenced", t.canonicalName))
	var e2 *tlast.ParseError
	if t.originTL2 {
		e2 = t.combTL2.FuncDecl.PRName.BeautifulError(errSeeHere)
	} else {
		e2 = t.combTL1[0].Construct.NamePR.BeautifulError(errSeeHere)
	}
	return tlast.BeautifulError2(e1, e2)
}

// We do not want to give generators access to combinators directly.
// We want every piece of information to come through strict pure public interface.
//func (t *KernelType) TL1() []*tlast.Combinator {
//	return t.combTL1
//}

//
//func (t *KernelType) TL2() tlast.TL2Combinator {
//	return t.combTL2
//}

func (t *KernelType) CombinatorTexts() []string {
	var result []string
	if t.originTL2 {
		result = append(result, t.combTL2.String())
	} else {
		for _, comb := range t.combTL1 {
			result = append(result, comb.String())
		}
	}
	return result
}

func (t *KernelType) Templates() []tlast.TL2TypeTemplate {
	return t.templateArguments
}

func (t *KernelType) DifferConstructorAndTypeName() bool {
	if t.originTL2 {
		return false
	}
	return len(t.combTL1) == 1 && !t.builtin && !strings.EqualFold(t.combTL1[0].TypeDecl.Name.String(), t.combTL1[0].Construct.Name.String())
}

func (t *KernelType) LegacyTypeName() tlast.TL2TypeName {
	if t.originTL2 {
		return t.canonicalName
	}
	return tlast.TL2TypeName(t.combTL1[0].TypeDecl.Name)
}
