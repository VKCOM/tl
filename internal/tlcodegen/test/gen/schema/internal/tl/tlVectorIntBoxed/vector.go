// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlVectorIntBoxed

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tl/tlBuiltinVectorIntBoxed"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type VectorIntBoxed []int32

func (VectorIntBoxed) TLName() string { return "vector" }
func (VectorIntBoxed) TLTag() uint32  { return 0x1cb5c415 }

func (item *VectorIntBoxed) Reset() {
	ptr := (*[]int32)(item)
	*ptr = (*ptr)[:0]
}

func (item *VectorIntBoxed) Read(w []byte) (_ []byte, err error) {
	ptr := (*[]int32)(item)
	return tlBuiltinVectorIntBoxed.BuiltinVectorIntBoxedRead(w, ptr)
}

// This method is general version of Write, use it instead!
func (item *VectorIntBoxed) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *VectorIntBoxed) Write(w []byte) []byte {
	ptr := (*[]int32)(item)
	return tlBuiltinVectorIntBoxed.BuiltinVectorIntBoxedWrite(w, *ptr)
}

func (item *VectorIntBoxed) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x1cb5c415); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *VectorIntBoxed) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *VectorIntBoxed) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x1cb5c415)
	return item.Write(w)
}

func (item *VectorIntBoxed) String() string {
	return string(item.WriteJSON(nil))
}

func (item *VectorIntBoxed) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	ptr := (*[]int32)(item)
	if err := tlBuiltinVectorIntBoxed.BuiltinVectorIntBoxedReadJSON(legacyTypeNames, in, ptr); err != nil {
		return err
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *VectorIntBoxed) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSON(w), nil
}

func (item *VectorIntBoxed) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}

func (item *VectorIntBoxed) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	ptr := (*[]int32)(item)
	w = tlBuiltinVectorIntBoxed.BuiltinVectorIntBoxedWriteJSONOpt(newTypeNames, short, w, *ptr)
	return w
}
func (item *VectorIntBoxed) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *VectorIntBoxed) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("vector", err.Error())
	}
	return nil
}
