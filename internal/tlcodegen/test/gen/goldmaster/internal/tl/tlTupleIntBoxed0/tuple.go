// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlTupleIntBoxed0

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlBuiltinTuple0IntBoxed"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type TupleIntBoxed0 [0]int32

func (TupleIntBoxed0) TLName() string { return "tuple" }
func (TupleIntBoxed0) TLTag() uint32  { return 0x9770768a }

func (item *TupleIntBoxed0) Reset() {
	ptr := (*[0]int32)(item)
	tlBuiltinTuple0IntBoxed.BuiltinTuple0IntBoxedReset(ptr)
}

func (item *TupleIntBoxed0) FillRandom(rg *basictl.RandGenerator) {
	ptr := (*[0]int32)(item)
	tlBuiltinTuple0IntBoxed.BuiltinTuple0IntBoxedFillRandom(rg, ptr)
}

func (item *TupleIntBoxed0) Read(w []byte) (_ []byte, err error) {
	ptr := (*[0]int32)(item)
	return tlBuiltinTuple0IntBoxed.BuiltinTuple0IntBoxedRead(w, ptr)
}

// This method is general version of Write, use it instead!
func (item *TupleIntBoxed0) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *TupleIntBoxed0) Write(w []byte) []byte {
	ptr := (*[0]int32)(item)
	return tlBuiltinTuple0IntBoxed.BuiltinTuple0IntBoxedWrite(w, ptr)
}

func (item *TupleIntBoxed0) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x9770768a); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *TupleIntBoxed0) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *TupleIntBoxed0) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x9770768a)
	return item.Write(w)
}

func (item TupleIntBoxed0) String() string {
	return string(item.WriteJSON(nil))
}

func (item *TupleIntBoxed0) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	ptr := (*[0]int32)(item)
	if err := tlBuiltinTuple0IntBoxed.BuiltinTuple0IntBoxedReadJSON(legacyTypeNames, in, ptr); err != nil {
		return err
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *TupleIntBoxed0) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSON(w), nil
}

func (item *TupleIntBoxed0) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}

func (item *TupleIntBoxed0) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	ptr := (*[0]int32)(item)
	w = tlBuiltinTuple0IntBoxed.BuiltinTuple0IntBoxedWriteJSONOpt(newTypeNames, short, w, ptr)
	return w
}
func (item *TupleIntBoxed0) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *TupleIntBoxed0) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("tuple", err.Error())
	}
	return nil
}
