// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlTupleInt3

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tl/tlBuiltinTuple3Int"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type TupleInt3 [3]int32

func (TupleInt3) TLName() string { return "tuple" }
func (TupleInt3) TLTag() uint32  { return 0x9770768a }

func (item *TupleInt3) Reset() {
	ptr := (*[3]int32)(item)
	tlBuiltinTuple3Int.BuiltinTuple3IntReset(ptr)
}

func (item *TupleInt3) Read(w []byte) (_ []byte, err error) {
	ptr := (*[3]int32)(item)
	return tlBuiltinTuple3Int.BuiltinTuple3IntRead(w, ptr)
}

// This method is general version of Write, use it instead!
func (item *TupleInt3) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *TupleInt3) Write(w []byte) []byte {
	ptr := (*[3]int32)(item)
	return tlBuiltinTuple3Int.BuiltinTuple3IntWrite(w, ptr)
}

func (item *TupleInt3) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x9770768a); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *TupleInt3) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *TupleInt3) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x9770768a)
	return item.Write(w)
}

func (item TupleInt3) String() string {
	return string(item.WriteJSON(nil))
}

func (item *TupleInt3) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	ptr := (*[3]int32)(item)
	if err := tlBuiltinTuple3Int.BuiltinTuple3IntReadJSON(legacyTypeNames, in, ptr); err != nil {
		return err
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *TupleInt3) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSON(w), nil
}

func (item *TupleInt3) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}

func (item *TupleInt3) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	ptr := (*[3]int32)(item)
	w = tlBuiltinTuple3Int.BuiltinTuple3IntWriteJSONOpt(newTypeNames, short, w, ptr)
	return w
}
func (item *TupleInt3) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *TupleInt3) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("tuple", err.Error())
	}
	return nil
}