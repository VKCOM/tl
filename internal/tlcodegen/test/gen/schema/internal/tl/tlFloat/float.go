// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlFloat

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type Float float32

func (Float) TLName() string { return "float" }
func (Float) TLTag() uint32  { return 0x824dab22 }

func (item *Float) Reset() {
	ptr := (*float32)(item)
	*ptr = 0
}

func (item *Float) Read(w []byte) (_ []byte, err error) {
	ptr := (*float32)(item)
	return basictl.FloatRead(w, ptr)
}

// This method is general version of Write, use it instead!
func (item *Float) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *Float) Write(w []byte) []byte {
	ptr := (*float32)(item)
	return basictl.FloatWrite(w, *ptr)
}

func (item *Float) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x824dab22); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *Float) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *Float) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x824dab22)
	return item.Write(w)
}

func (item Float) String() string {
	return string(item.WriteJSON(nil))
}

func (item *Float) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	ptr := (*float32)(item)
	if err := internal.Json2ReadFloat32(in, ptr); err != nil {
		return err
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *Float) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSON(w), nil
}

func (item *Float) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}

func (item *Float) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	ptr := (*float32)(item)
	w = basictl.JSONWriteFloat32(w, *ptr)
	return w
}
func (item *Float) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *Float) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("float", err.Error())
	}
	return nil
}