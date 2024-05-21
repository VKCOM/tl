// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlDictionaryAnyIntInt

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tl/tlBuiltinVectorDictionaryFieldAnyIntInt"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tl/tlDictionaryFieldAnyIntInt"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type DictionaryAnyIntInt map[int32]int32

func (DictionaryAnyIntInt) TLName() string { return "dictionaryAny" }
func (DictionaryAnyIntInt) TLTag() uint32  { return 0x1f4c6190 }

func (item *DictionaryAnyIntInt) Reset() {
	ptr := (*map[int32]int32)(item)
	tlBuiltinVectorDictionaryFieldAnyIntInt.BuiltinVectorDictionaryFieldAnyIntIntReset(*ptr)
}

func (item *DictionaryAnyIntInt) FillRandom(rg *basictl.RandGenerator) {
	ptr := (*map[int32]int32)(item)
	tlBuiltinVectorDictionaryFieldAnyIntInt.BuiltinVectorDictionaryFieldAnyIntIntFillRandom(rg, ptr)
}

func (item *DictionaryAnyIntInt) Read(w []byte) (_ []byte, err error) {
	ptr := (*map[int32]int32)(item)
	return tlBuiltinVectorDictionaryFieldAnyIntInt.BuiltinVectorDictionaryFieldAnyIntIntRead(w, ptr)
}

// This method is general version of Write, use it instead!
func (item *DictionaryAnyIntInt) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *DictionaryAnyIntInt) Write(w []byte) []byte {
	ptr := (*map[int32]int32)(item)
	return tlBuiltinVectorDictionaryFieldAnyIntInt.BuiltinVectorDictionaryFieldAnyIntIntWrite(w, *ptr)
}

func (item *DictionaryAnyIntInt) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x1f4c6190); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *DictionaryAnyIntInt) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *DictionaryAnyIntInt) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x1f4c6190)
	return item.Write(w)
}

func (item DictionaryAnyIntInt) String() string {
	return string(item.WriteJSON(nil))
}

func (item *DictionaryAnyIntInt) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	ptr := (*map[int32]int32)(item)
	if err := tlBuiltinVectorDictionaryFieldAnyIntInt.BuiltinVectorDictionaryFieldAnyIntIntReadJSON(legacyTypeNames, in, ptr); err != nil {
		return err
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *DictionaryAnyIntInt) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSON(w), nil
}

func (item *DictionaryAnyIntInt) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}

func (item *DictionaryAnyIntInt) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	ptr := (*map[int32]int32)(item)
	w = tlBuiltinVectorDictionaryFieldAnyIntInt.BuiltinVectorDictionaryFieldAnyIntIntWriteJSONOpt(newTypeNames, short, w, *ptr)
	return w
}
func (item *DictionaryAnyIntInt) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *DictionaryAnyIntInt) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("dictionaryAny", err.Error())
	}
	return nil
}

type DictionaryAnyIntIntBytes []tlDictionaryFieldAnyIntInt.DictionaryFieldAnyIntInt

func (DictionaryAnyIntIntBytes) TLName() string { return "dictionaryAny" }
func (DictionaryAnyIntIntBytes) TLTag() uint32  { return 0x1f4c6190 }

func (item *DictionaryAnyIntIntBytes) Reset() {
	ptr := (*[]tlDictionaryFieldAnyIntInt.DictionaryFieldAnyIntInt)(item)
	*ptr = (*ptr)[:0]
}

func (item *DictionaryAnyIntIntBytes) FillRandom(rg *basictl.RandGenerator) {
	ptr := (*[]tlDictionaryFieldAnyIntInt.DictionaryFieldAnyIntInt)(item)
	tlBuiltinVectorDictionaryFieldAnyIntInt.BuiltinVectorDictionaryFieldAnyIntIntBytesFillRandom(rg, ptr)
}

func (item *DictionaryAnyIntIntBytes) Read(w []byte) (_ []byte, err error) {
	ptr := (*[]tlDictionaryFieldAnyIntInt.DictionaryFieldAnyIntInt)(item)
	return tlBuiltinVectorDictionaryFieldAnyIntInt.BuiltinVectorDictionaryFieldAnyIntIntBytesRead(w, ptr)
}

// This method is general version of Write, use it instead!
func (item *DictionaryAnyIntIntBytes) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *DictionaryAnyIntIntBytes) Write(w []byte) []byte {
	ptr := (*[]tlDictionaryFieldAnyIntInt.DictionaryFieldAnyIntInt)(item)
	return tlBuiltinVectorDictionaryFieldAnyIntInt.BuiltinVectorDictionaryFieldAnyIntIntBytesWrite(w, *ptr)
}

func (item *DictionaryAnyIntIntBytes) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x1f4c6190); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *DictionaryAnyIntIntBytes) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *DictionaryAnyIntIntBytes) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x1f4c6190)
	return item.Write(w)
}

func (item DictionaryAnyIntIntBytes) String() string {
	return string(item.WriteJSON(nil))
}

func (item *DictionaryAnyIntIntBytes) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	ptr := (*[]tlDictionaryFieldAnyIntInt.DictionaryFieldAnyIntInt)(item)
	if err := tlBuiltinVectorDictionaryFieldAnyIntInt.BuiltinVectorDictionaryFieldAnyIntIntBytesReadJSON(legacyTypeNames, in, ptr); err != nil {
		return err
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *DictionaryAnyIntIntBytes) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSON(w), nil
}

func (item *DictionaryAnyIntIntBytes) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}

func (item *DictionaryAnyIntIntBytes) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	ptr := (*[]tlDictionaryFieldAnyIntInt.DictionaryFieldAnyIntInt)(item)
	w = tlBuiltinVectorDictionaryFieldAnyIntInt.BuiltinVectorDictionaryFieldAnyIntIntBytesWriteJSONOpt(newTypeNames, short, w, *ptr)
	return w
}
func (item *DictionaryAnyIntIntBytes) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *DictionaryAnyIntIntBytes) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("dictionaryAny", err.Error())
	}
	return nil
}
