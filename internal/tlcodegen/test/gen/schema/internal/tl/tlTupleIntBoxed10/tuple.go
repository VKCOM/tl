// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlTupleIntBoxed10

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tl/tlBuiltinTuple10IntBoxed"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type TupleIntBoxed10 [10]int32

func (TupleIntBoxed10) TLName() string { return "tuple" }
func (TupleIntBoxed10) TLTag() uint32  { return 0x9770768a }

func (item *TupleIntBoxed10) Reset() {
	ptr := (*[10]int32)(item)
	tlBuiltinTuple10IntBoxed.BuiltinTuple10IntBoxedReset(ptr)
}

func (item *TupleIntBoxed10) Read(w []byte) (_ []byte, err error) {
	ptr := (*[10]int32)(item)
	return tlBuiltinTuple10IntBoxed.BuiltinTuple10IntBoxedRead(w, ptr)
}

// This method is general version of Write, use it instead!
func (item *TupleIntBoxed10) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *TupleIntBoxed10) Write(w []byte) []byte {
	ptr := (*[10]int32)(item)
	return tlBuiltinTuple10IntBoxed.BuiltinTuple10IntBoxedWrite(w, ptr)
}

func (item *TupleIntBoxed10) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x9770768a); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *TupleIntBoxed10) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *TupleIntBoxed10) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x9770768a)
	return item.Write(w)
}

func (item TupleIntBoxed10) String() string {
	return string(item.WriteJSON(nil))
}

func (item *TupleIntBoxed10) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	ptr := (*[10]int32)(item)
	if err := tlBuiltinTuple10IntBoxed.BuiltinTuple10IntBoxedReadJSON(legacyTypeNames, in, ptr); err != nil {
		return err
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *TupleIntBoxed10) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSON(w), nil
}

func (item *TupleIntBoxed10) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}

func (item *TupleIntBoxed10) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	ptr := (*[10]int32)(item)
	w = tlBuiltinTuple10IntBoxed.BuiltinTuple10IntBoxedWriteJSONOpt(newTypeNames, short, w, ptr)
	return w
}
func (item *TupleIntBoxed10) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *TupleIntBoxed10) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("tuple", err.Error())
	}
	return nil
}
