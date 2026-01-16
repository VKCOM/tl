// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package pure

import "github.com/vkcom/tl/internal/tlast"

type KernelType struct {
	originTL2 bool
	combTL1   []*tlast.Combinator
	combTL2   tlast.TL2Combinator
	// index by canonical name
	instances map[string]*TypeInstanceRef
	// order of instantiation
	instancesOrdered []*TypeInstanceRef
}

func (t *KernelType) OriginTL2() bool {
	return t.originTL2
}

func (t *KernelType) TL1() []*tlast.Combinator {
	return t.combTL1
}

func (t *KernelType) TL2() tlast.TL2Combinator {
	return t.combTL2
}
