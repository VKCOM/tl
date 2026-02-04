// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package pure

import "github.com/vkcom/tl/internal/tlast"

type KernelType struct {
	originTL2 bool
	builtin   bool
	combTL1   []*tlast.Combinator
	combTL2   tlast.TL2Combinator
	// index by canonical name
	instances map[string]*TypeInstanceRef
	// order of instantiation
	instancesOrdered []*TypeInstanceRef

	isFunction bool // to prohibit references

	tl1Names map[string]struct{}
	tl2Names map[string]struct{}

	canonicalName tlast.Name
	tl1BoxedName  tlast.Name
	canBeBare     bool
	tl1name       string // if !empty, go generator will use it for template names (VectorInt, not VectorInt32)
}

func (t *KernelType) OriginTL2() bool {
	return t.originTL2
}

func (t *KernelType) IsFunction() bool {
	return t.isFunction
}

func (t *KernelType) CanonicalName() tlast.Name {
	return t.canonicalName
}

func (t *KernelType) CanBeBare() bool {
	return t.canBeBare
}

func (t *KernelType) CanBeBoxed() bool {
	return t.tl1BoxedName != tlast.Name{}
}

func (t *KernelType) TL1() []*tlast.Combinator {
	return t.combTL1
}

func (t *KernelType) TL2() tlast.TL2Combinator {
	return t.combTL2
}
