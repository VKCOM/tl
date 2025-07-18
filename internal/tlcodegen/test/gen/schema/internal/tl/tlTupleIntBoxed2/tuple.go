// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlTupleIntBoxed2

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tl/tlBuiltinTuple2IntBoxed"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type TupleIntBoxed2 [2]int32

func (TupleIntBoxed2) TLName() string { return "tuple" }
func (TupleIntBoxed2) TLTag() uint32  { return 0x9770768a }

func (item *TupleIntBoxed2) Reset() {
	ptr := (*[2]int32)(item)
	tlBuiltinTuple2IntBoxed.BuiltinTuple2IntBoxedReset(ptr)
}

func (item *TupleIntBoxed2) Read(w []byte) (_ []byte, err error) {
	ptr := (*[2]int32)(item)
	return tlBuiltinTuple2IntBoxed.BuiltinTuple2IntBoxedRead(w, ptr)
}

func (item *TupleIntBoxed2) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *TupleIntBoxed2) Write(w []byte) []byte {
	ptr := (*[2]int32)(item)
	return tlBuiltinTuple2IntBoxed.BuiltinTuple2IntBoxedWrite(w, ptr)
}

func (item *TupleIntBoxed2) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x9770768a); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *TupleIntBoxed2) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *TupleIntBoxed2) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x9770768a)
	return item.Write(w)
}

func (item TupleIntBoxed2) String() string {
	return string(item.WriteJSON(nil))
}
func (item *TupleIntBoxed2) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	ptr := (*[2]int32)(item)
	if err := tlBuiltinTuple2IntBoxed.BuiltinTuple2IntBoxedReadJSON(legacyTypeNames, in, ptr); err != nil {
		return err
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *TupleIntBoxed2) WriteJSONGeneral(tctx *basictl.JSONWriteContext, w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(tctx, w), nil
}

func (item *TupleIntBoxed2) WriteJSON(w []byte) []byte {
	tctx := basictl.JSONWriteContext{}
	return item.WriteJSONOpt(&tctx, w)
}

func (item *TupleIntBoxed2) WriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte) []byte {
	ptr := (*[2]int32)(item)
	w = tlBuiltinTuple2IntBoxed.BuiltinTuple2IntBoxedWriteJSONOpt(tctx, w, ptr)
	return w
}
func (item *TupleIntBoxed2) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *TupleIntBoxed2) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("tuple", err.Error())
	}
	return nil
}
