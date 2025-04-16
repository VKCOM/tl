// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlVectorMapStringString

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tl/tlBuiltinVectorMapStringString"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tl/tlMapStringString"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type VectorMapStringString []tlMapStringString.MapStringString

func (VectorMapStringString) TLName() string { return "vector" }
func (VectorMapStringString) TLTag() uint32  { return 0x1cb5c415 }

func (item *VectorMapStringString) Reset() {
	ptr := (*[]tlMapStringString.MapStringString)(item)
	*ptr = (*ptr)[:0]
}

func (item *VectorMapStringString) Read(w []byte) (_ []byte, err error) {
	ptr := (*[]tlMapStringString.MapStringString)(item)
	return tlBuiltinVectorMapStringString.BuiltinVectorMapStringStringRead(w, ptr)
}

// This method is general version of Write, use it instead!
func (item *VectorMapStringString) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *VectorMapStringString) Write(w []byte) []byte {
	ptr := (*[]tlMapStringString.MapStringString)(item)
	return tlBuiltinVectorMapStringString.BuiltinVectorMapStringStringWrite(w, *ptr)
}

func (item *VectorMapStringString) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x1cb5c415); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *VectorMapStringString) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *VectorMapStringString) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x1cb5c415)
	return item.Write(w)
}

func (item *VectorMapStringString) String() string {
	return string(item.WriteJSON(nil))
}

func (item *VectorMapStringString) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	ptr := (*[]tlMapStringString.MapStringString)(item)
	if err := tlBuiltinVectorMapStringString.BuiltinVectorMapStringStringReadJSON(legacyTypeNames, in, ptr); err != nil {
		return err
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *VectorMapStringString) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSON(w), nil
}

func (item *VectorMapStringString) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}

func (item *VectorMapStringString) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	ptr := (*[]tlMapStringString.MapStringString)(item)
	w = tlBuiltinVectorMapStringString.BuiltinVectorMapStringStringWriteJSONOpt(newTypeNames, short, w, *ptr)
	return w
}
func (item *VectorMapStringString) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *VectorMapStringString) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("vector", err.Error())
	}
	return nil
}
