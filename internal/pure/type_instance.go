// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package pure

import "github.com/vkcom/tl/internal/tlast"

type TypeInstance interface {
	CanonicalName() string
	KernelType() *KernelType
	Common() *TypeInstanceCommon

	GoodForMapKey() bool
	IsBit() bool // for vector/tuple special case
	FindCycle(c *cycleFinder)
	BoxedOnly() bool

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
	natParams     []string // external nat params (empty for TL2 types)
	tip           *KernelType
	isTopLevel    bool
	rt            tlast.TypeRef
	argNamespace  string // so vector<memcache.Value> is generated in memcache namespace
}

func (ins *TypeInstanceCommon) CanonicalName() string {
	return ins.canonicalName
}

func (ins *TypeInstanceCommon) NatParams() []string {
	return ins.natParams
}

func (ins *TypeInstanceCommon) KernelType() *KernelType {
	return ins.tip
}

func (ins *TypeInstanceCommon) IsTopLevel() bool {
	return ins.isTopLevel
}

func (ins *TypeInstanceCommon) ResolvedType() tlast.TypeRef {
	return ins.rt
}

func (ins *TypeInstanceCommon) ArgNamespace() string {
	return ins.argNamespace
}

func (ins *TypeInstanceCommon) BoxedOnly() bool {
	return false
}

func (ins *TypeInstanceCommon) Common() *TypeInstanceCommon {
	return ins
}

func (ins *TypeInstanceCommon) GoodForMapKey() bool {
	return false
}

func (ins *TypeInstanceCommon) IsBit() bool {
	return false
}
