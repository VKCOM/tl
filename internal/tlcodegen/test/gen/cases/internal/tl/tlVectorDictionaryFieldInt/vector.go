// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlVectorDictionaryFieldInt

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tl/tlBuiltinVectorDictionaryFieldInt"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tl/tlDictionaryFieldInt"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type VectorDictionaryFieldInt map[string]int32

func (VectorDictionaryFieldInt) TLName() string { return "vector" }
func (VectorDictionaryFieldInt) TLTag() uint32  { return 0x1cb5c415 }

func (item *VectorDictionaryFieldInt) Reset() {
	ptr := (*map[string]int32)(item)
	tlBuiltinVectorDictionaryFieldInt.BuiltinVectorDictionaryFieldIntReset(*ptr)
}

func (item *VectorDictionaryFieldInt) FillRandom(rg *basictl.RandGenerator) {
	ptr := (*map[string]int32)(item)
	tlBuiltinVectorDictionaryFieldInt.BuiltinVectorDictionaryFieldIntFillRandom(rg, ptr)
}

func (item *VectorDictionaryFieldInt) Read(w []byte) (_ []byte, err error) {
	ptr := (*map[string]int32)(item)
	return tlBuiltinVectorDictionaryFieldInt.BuiltinVectorDictionaryFieldIntRead(w, ptr)
}

// This method is general version of Write, use it instead!
func (item *VectorDictionaryFieldInt) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *VectorDictionaryFieldInt) Write(w []byte) []byte {
	ptr := (*map[string]int32)(item)
	return tlBuiltinVectorDictionaryFieldInt.BuiltinVectorDictionaryFieldIntWrite(w, *ptr)
}

func (item *VectorDictionaryFieldInt) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x1cb5c415); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *VectorDictionaryFieldInt) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *VectorDictionaryFieldInt) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x1cb5c415)
	return item.Write(w)
}

func (item *VectorDictionaryFieldInt) String() string {
	return string(item.WriteJSON(nil))
}

func (item *VectorDictionaryFieldInt) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	ptr := (*map[string]int32)(item)
	if err := tlBuiltinVectorDictionaryFieldInt.BuiltinVectorDictionaryFieldIntReadJSON(legacyTypeNames, in, ptr); err != nil {
		return err
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *VectorDictionaryFieldInt) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSON(w), nil
}

func (item *VectorDictionaryFieldInt) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}

func (item *VectorDictionaryFieldInt) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	ptr := (*map[string]int32)(item)
	w = tlBuiltinVectorDictionaryFieldInt.BuiltinVectorDictionaryFieldIntWriteJSONOpt(newTypeNames, short, w, *ptr)
	return w
}
func (item *VectorDictionaryFieldInt) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *VectorDictionaryFieldInt) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("vector", err.Error())
	}
	return nil
}

type VectorDictionaryFieldIntBytes []tlDictionaryFieldInt.DictionaryFieldIntBytes

func (VectorDictionaryFieldIntBytes) TLName() string { return "vector" }
func (VectorDictionaryFieldIntBytes) TLTag() uint32  { return 0x1cb5c415 }

func (item *VectorDictionaryFieldIntBytes) Reset() {
	ptr := (*[]tlDictionaryFieldInt.DictionaryFieldIntBytes)(item)
	*ptr = (*ptr)[:0]
}

func (item *VectorDictionaryFieldIntBytes) FillRandom(rg *basictl.RandGenerator) {
	ptr := (*[]tlDictionaryFieldInt.DictionaryFieldIntBytes)(item)
	tlBuiltinVectorDictionaryFieldInt.BuiltinVectorDictionaryFieldIntBytesFillRandom(rg, ptr)
}

func (item *VectorDictionaryFieldIntBytes) Read(w []byte) (_ []byte, err error) {
	ptr := (*[]tlDictionaryFieldInt.DictionaryFieldIntBytes)(item)
	return tlBuiltinVectorDictionaryFieldInt.BuiltinVectorDictionaryFieldIntBytesRead(w, ptr)
}

// This method is general version of Write, use it instead!
func (item *VectorDictionaryFieldIntBytes) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *VectorDictionaryFieldIntBytes) Write(w []byte) []byte {
	ptr := (*[]tlDictionaryFieldInt.DictionaryFieldIntBytes)(item)
	return tlBuiltinVectorDictionaryFieldInt.BuiltinVectorDictionaryFieldIntBytesWrite(w, *ptr)
}

func (item *VectorDictionaryFieldIntBytes) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x1cb5c415); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *VectorDictionaryFieldIntBytes) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *VectorDictionaryFieldIntBytes) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x1cb5c415)
	return item.Write(w)
}

func (item *VectorDictionaryFieldIntBytes) String() string {
	return string(item.WriteJSON(nil))
}

func (item *VectorDictionaryFieldIntBytes) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	ptr := (*[]tlDictionaryFieldInt.DictionaryFieldIntBytes)(item)
	if err := tlBuiltinVectorDictionaryFieldInt.BuiltinVectorDictionaryFieldIntBytesReadJSON(legacyTypeNames, in, ptr); err != nil {
		return err
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *VectorDictionaryFieldIntBytes) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSON(w), nil
}

func (item *VectorDictionaryFieldIntBytes) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}

func (item *VectorDictionaryFieldIntBytes) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	ptr := (*[]tlDictionaryFieldInt.DictionaryFieldIntBytes)(item)
	w = tlBuiltinVectorDictionaryFieldInt.BuiltinVectorDictionaryFieldIntBytesWriteJSONOpt(newTypeNames, short, w, *ptr)
	return w
}
func (item *VectorDictionaryFieldIntBytes) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *VectorDictionaryFieldIntBytes) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("vector", err.Error())
	}
	return nil
}
