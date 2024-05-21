// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlVectorDictionaryFieldString

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tl/tlBuiltinVectorDictionaryFieldString"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type VectorDictionaryFieldString map[string]string

func (VectorDictionaryFieldString) TLName() string { return "vector" }
func (VectorDictionaryFieldString) TLTag() uint32  { return 0x1cb5c415 }

func (item *VectorDictionaryFieldString) Reset() {
	ptr := (*map[string]string)(item)
	tlBuiltinVectorDictionaryFieldString.BuiltinVectorDictionaryFieldStringReset(*ptr)
}

func (item *VectorDictionaryFieldString) FillRandom(rg *basictl.RandGenerator) {
	ptr := (*map[string]string)(item)
	tlBuiltinVectorDictionaryFieldString.BuiltinVectorDictionaryFieldStringFillRandom(rg, ptr)
}

func (item *VectorDictionaryFieldString) Read(w []byte) (_ []byte, err error) {
	ptr := (*map[string]string)(item)
	return tlBuiltinVectorDictionaryFieldString.BuiltinVectorDictionaryFieldStringRead(w, ptr)
}

// This method is general version of Write, use it instead!
func (item *VectorDictionaryFieldString) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *VectorDictionaryFieldString) Write(w []byte) []byte {
	ptr := (*map[string]string)(item)
	return tlBuiltinVectorDictionaryFieldString.BuiltinVectorDictionaryFieldStringWrite(w, *ptr)
}

func (item *VectorDictionaryFieldString) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x1cb5c415); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *VectorDictionaryFieldString) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *VectorDictionaryFieldString) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x1cb5c415)
	return item.Write(w)
}

func (item VectorDictionaryFieldString) String() string {
	return string(item.WriteJSON(nil))
}

func (item *VectorDictionaryFieldString) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	ptr := (*map[string]string)(item)
	if err := tlBuiltinVectorDictionaryFieldString.BuiltinVectorDictionaryFieldStringReadJSON(legacyTypeNames, in, ptr); err != nil {
		return err
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *VectorDictionaryFieldString) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSON(w), nil
}

func (item *VectorDictionaryFieldString) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}

func (item *VectorDictionaryFieldString) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	ptr := (*map[string]string)(item)
	w = tlBuiltinVectorDictionaryFieldString.BuiltinVectorDictionaryFieldStringWriteJSONOpt(newTypeNames, short, w, *ptr)
	return w
}
func (item *VectorDictionaryFieldString) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *VectorDictionaryFieldString) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("vector", err.Error())
	}
	return nil
}
