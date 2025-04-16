// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlVectorDictionaryFieldDictionaryInt

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tl/tlBuiltinVectorDictionaryFieldDictionaryInt"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type VectorDictionaryFieldDictionaryInt map[string]map[string]int32

func (VectorDictionaryFieldDictionaryInt) TLName() string { return "vector" }
func (VectorDictionaryFieldDictionaryInt) TLTag() uint32  { return 0x1cb5c415 }

func (item *VectorDictionaryFieldDictionaryInt) Reset() {
	ptr := (*map[string]map[string]int32)(item)
	tlBuiltinVectorDictionaryFieldDictionaryInt.BuiltinVectorDictionaryFieldDictionaryIntReset(*ptr)
}

func (item *VectorDictionaryFieldDictionaryInt) Read(w []byte) (_ []byte, err error) {
	ptr := (*map[string]map[string]int32)(item)
	return tlBuiltinVectorDictionaryFieldDictionaryInt.BuiltinVectorDictionaryFieldDictionaryIntRead(w, ptr)
}

// This method is general version of Write, use it instead!
func (item *VectorDictionaryFieldDictionaryInt) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *VectorDictionaryFieldDictionaryInt) Write(w []byte) []byte {
	ptr := (*map[string]map[string]int32)(item)
	return tlBuiltinVectorDictionaryFieldDictionaryInt.BuiltinVectorDictionaryFieldDictionaryIntWrite(w, *ptr)
}

func (item *VectorDictionaryFieldDictionaryInt) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x1cb5c415); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *VectorDictionaryFieldDictionaryInt) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *VectorDictionaryFieldDictionaryInt) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x1cb5c415)
	return item.Write(w)
}

func (item *VectorDictionaryFieldDictionaryInt) String() string {
	return string(item.WriteJSON(nil))
}

func (item *VectorDictionaryFieldDictionaryInt) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	ptr := (*map[string]map[string]int32)(item)
	if err := tlBuiltinVectorDictionaryFieldDictionaryInt.BuiltinVectorDictionaryFieldDictionaryIntReadJSON(legacyTypeNames, in, ptr); err != nil {
		return err
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *VectorDictionaryFieldDictionaryInt) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSON(w), nil
}

func (item *VectorDictionaryFieldDictionaryInt) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}

func (item *VectorDictionaryFieldDictionaryInt) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	ptr := (*map[string]map[string]int32)(item)
	w = tlBuiltinVectorDictionaryFieldDictionaryInt.BuiltinVectorDictionaryFieldDictionaryIntWriteJSONOpt(newTypeNames, short, w, *ptr)
	return w
}
func (item *VectorDictionaryFieldDictionaryInt) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *VectorDictionaryFieldDictionaryInt) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("vector", err.Error())
	}
	return nil
}
