// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlInt

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type Int int32

func (Int) TLName() string { return "int" }
func (Int) TLTag() uint32  { return 0xa8509bda }

func (item *Int) Reset() {
	ptr := (*int32)(item)
	*ptr = 0
}

func (item *Int) FillRandom(rg *basictl.RandGenerator) {
	ptr := (*int32)(item)
	*ptr = basictl.RandomInt(rg)
}

func (item *Int) Read(w []byte) (_ []byte, err error) {
	ptr := (*int32)(item)
	return basictl.IntRead(w, ptr)
}

// This method is general version of Write, use it instead!
func (item *Int) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *Int) Write(w []byte) []byte {
	ptr := (*int32)(item)
	return basictl.IntWrite(w, *ptr)
}

func (item *Int) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xa8509bda); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *Int) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *Int) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0xa8509bda)
	return item.Write(w)
}

func (item Int) String() string {
	return string(item.WriteJSON(nil))
}

func (item *Int) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	ptr := (*int32)(item)
	if err := internal.Json2ReadInt32(in, ptr); err != nil {
		return err
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *Int) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSON(w), nil
}

func (item *Int) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}

func (item *Int) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	ptr := (*int32)(item)
	w = basictl.JSONWriteInt32(w, *ptr)
	return w
}
func (item *Int) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *Int) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("int", err.Error())
	}
	return nil
}
