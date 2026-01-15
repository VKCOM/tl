// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package pure

type TypeInstance interface {
	CanonicalName() string
	// Combinator() tlast.TL2Combinator

	GoodForMapKey() bool
	IsBit() bool // for vector/tuple special case
	FindCycle(c *cycleFinder)

	CreateValue() KernelValue
	SkipTL2(r []byte) ([]byte, error)
}

// during recursive type resolution, we store pointer to this type,
// later type instance is instantiated and ins is set
type TypeInstanceRef struct {
	ins TypeInstance
}

type TypeInstanceCommon struct {
	canonicalName string
	NatParams     []string // external nat params (empty for TL2 types)
	// tip           *KernelType
}

func (ins *TypeInstanceCommon) CanonicalName() string {
	return ins.canonicalName
}

//func (ins *TypeInstanceCommon) Combinator() tlast.TL2Combinator {
//	return ins.comb
//}

func (ins *TypeInstanceCommon) GoodForMapKey() bool {
	return false
}

func (ins *TypeInstanceCommon) IsBit() bool {
	return false
}
