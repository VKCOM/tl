// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package pure

import (
	"github.com/VKCOM/tl/internal/tlast"
)

type ActualNatArg struct {
	isNumber bool
	number   uint32
	isField  bool
	index    int
	name     string
}

func (arg *ActualNatArg) IsNumber() bool {
	return arg.isNumber
}

func (arg *ActualNatArg) Number() uint32 {
	return arg.number
}

func (arg *ActualNatArg) IsField() bool {
	return !arg.isNumber && arg.isField
}

func (arg *ActualNatArg) IsNatParam() bool {
	return !arg.isNumber && !arg.isField
}

// also NatParam index in parent type, useful for vkext dynamic parsing,
// so we can use stack(slice) instead of map, which is faster and allocates less
func (arg *ActualNatArg) FieldIndex() int {
	return arg.index
}

func (arg *ActualNatArg) NatParamName() string {
	return arg.name
}

type Field struct {
	owner TypeInstance
	name  string
	ins   *TypeInstanceRef

	commentBefore string
	commentRight  string

	// though all TL2 types are bare, we still set Boxed for unions, because we want
	// vector<Union> and []Union to reference the same generated type
	bare bool

	fieldMask *ActualNatArg
	bitNumber uint32 // only used when fieldMask != nil

	maskTL2Bit *int

	natArgs []ActualNatArg // for TL1 only, empty for TL2

	pr tlast.PositionRange
}

func (f Field) OwnerTypeInstance() TypeInstance { return f.owner }

func (f Field) Bare() bool                 { return f.bare }
func (f Field) Name() string               { return f.name }
func (f Field) CommentBefore() string      { return f.commentBefore }
func (f Field) CommentRight() string       { return f.commentRight }
func (f Field) TypeInstance() TypeInstance { return f.ins.ins }
func (f Field) FieldMask() *ActualNatArg   { return f.fieldMask }
func (f Field) BitNumber() uint32          { return f.bitNumber }
func (f Field) NatArgs() []ActualNatArg    { return f.natArgs }

// we do not know if this object is used by some other TL2 object when we generate this,
// so we return nil if owner does not marked as one needing TL2
func (f Field) MaskTL2Bit() *int {
	if !f.owner.Common().HasTL2() {
		return nil
	}
	return f.maskTL2Bit
}

func (f Field) IsBit() bool {
	if f.ins.ins == nil {
		// recursive type, this check allows IsBit to be called even during type resolution
		// bit is never recursive, so this case always works correctly
		return false
	}
	if f.ins.ins.IsBit() {
		return true
	}
	return f.fieldMask != nil && f.ins.ins.CanonicalName() == "True"
}
