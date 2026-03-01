// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package gengo

import (
	"fmt"
	"strings"

	"github.com/vkcom/tl/internal/pure"
)

type Field struct {
	pureField pure.Field
	t         *TypeRWWrapper
	goName    string
	recursive bool
}

func (f *Field) OriginalName() string {
	return f.pureField.Name()
}

func (f *Field) Bare() bool {
	return f.pureField.Bare()
}

func (f *Field) FieldMask() *pure.ActualNatArg {
	return f.pureField.FieldMask()
}

func (f *Field) BitNumber() uint32 {
	return f.pureField.BitNumber()
}

func (f *Field) MaskTL2Bit() *int {
	return f.pureField.MaskTL2Bit()
}

func (f *Field) NatArgs() []pure.ActualNatArg {
	return f.pureField.NatArgs()
}

func (f *Field) IsAffectingLocalFieldMasks() bool {
	return f.FieldMask() != nil && f.FieldMask().IsField()
}

func (f *Field) IsAffectedByExternalFieldMask() bool {
	return f.FieldMask() != nil && !f.FieldMask().IsField()
}

func (f *Field) IsTypeDependsFromLocalFields() bool {
	for _, natArg := range f.NatArgs() {
		if natArg.IsField() {
			return true
		}
	}
	return false
}

func (f *Field) HasNatArguments() bool {
	return len(f.NatArgs()) != 0
}

// do not generate fields, but affect block position and skip during reading
// TL1: never
// TL2: _:X
func (f *Field) IsTL2Omitted() bool {
	return strings.HasPrefix(f.OriginalName(), "_")
}

// generate Set/IsSet with external (TL1) or internal (TL1 & TL2) mask/
// must exactly correspond to migrator logic
// TL1: x:fm.b?true x:fm.b?True
// TL2: x:bit
func (f *Field) IsBit() bool {
	if b, ok := f.t.trw.(*TypeRWBool); ok {
		return b.isBit
	}
	return f.FieldMask() != nil && (f.t.IsTrueType() && (f.t.tlName.String() == "true" || f.t.tlName.String() == "True"))
}

func (f *Field) TL2MaskForOP(op string) string {
	return fmt.Sprintf("tl2mask%d %s %d", *f.MaskTL2Bit()/8, op, 1<<(*f.MaskTL2Bit()%8))
}

func (f *Field) EnsureRecursive(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace) string {
	if !f.recursive {
		return ""
	}
	myType := f.t.TypeString2(bytesVersion, directImports, ins, false, false)
	return fmt.Sprintf(`	if item.%s == nil { item.%s = new(%s) }
`, f.goName, f.goName, myType)
}

func (f *Field) TypeResettingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace) string {
	resetCode := f.t.TypeResettingCode(bytesVersion, directImports, ins, fmt.Sprintf("item.%s", f.goName), f.recursive)
	if f.recursive {
		return fmt.Sprintf(`	if item.%s != nil {
		%s
	}`, f.goName, resetCode)
	}
	return resetCode
}

func (f *Field) FieldAccess(trw *TypeRWStruct, bytesVersion bool, directImports *DirectImports, ins *InternalNamespace) (string, bool) {
	// presumably, the only case when we have an empty name, is single first field of alias-like defs
	if f.pureField.Name() != "" {
		return fmt.Sprintf("item.%s", f.goName), f.recursive
	}
	myType := f.t.TypeString2(bytesVersion, directImports, ins, false, false)
	return fmt.Sprintf("(*%s)(item)", myType), true
}
