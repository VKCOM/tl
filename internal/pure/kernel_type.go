// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package pure

import (
	"fmt"

	"github.com/vkcom/tl/internal/tlast"
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

	isTopLevel  bool
	isFunction  bool // to prohibit references
	annotations []string

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
	tl1BoxedName   tlast.TL2TypeName
	canBeBare      bool
	historicalName tlast.TL2TypeName // go gen historically uses TL1-style names for builtin types

	namePR tlast.PositionRange // for "see here" in beautiful errors

	// usage tracking for migration/compilcation.
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

func (t *KernelType) HistoricalName() tlast.TL2TypeName {
	return t.historicalName
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
func (t *KernelType) TL1() []*tlast.Combinator {
	return t.combTL1
}

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
