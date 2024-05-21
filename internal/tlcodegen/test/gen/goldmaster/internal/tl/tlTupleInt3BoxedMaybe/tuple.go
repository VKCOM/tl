// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlTupleInt3BoxedMaybe

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlBuiltinTuple3Int"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type TupleInt3BoxedMaybe struct {
	Value [3]int32 // not deterministic if !Ok
	Ok    bool
}

func (item *TupleInt3BoxedMaybe) Reset() {
	item.Ok = false
}
func (item *TupleInt3BoxedMaybe) FillRandom(rg *basictl.RandGenerator) {
	if basictl.RandomUint(rg)&1 == 1 {
		item.Ok = true
		tlBuiltinTuple3Int.BuiltinTuple3IntFillRandom(rg, &item.Value)
	} else {
		item.Ok = false
	}
}

func (item *TupleInt3BoxedMaybe) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.ReadBool(w, &item.Ok, 0x27930a7b, 0x3f9c8ef8); err != nil {
		return w, err
	}
	if item.Ok {
		if w, err = basictl.NatReadExactTag(w, 0x9770768a); err != nil {
			return w, err
		}
		return tlBuiltinTuple3Int.BuiltinTuple3IntRead(w, &item.Value)
	}
	return w, nil
}

// This method is general version of WriteBoxed, use it instead!
func (item *TupleInt3BoxedMaybe) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *TupleInt3BoxedMaybe) WriteBoxed(w []byte) []byte {
	if item.Ok {
		w = basictl.NatWrite(w, 0x3f9c8ef8)
		w = basictl.NatWrite(w, 0x9770768a)
		return tlBuiltinTuple3Int.BuiltinTuple3IntWrite(w, &item.Value)
	}
	return basictl.NatWrite(w, 0x27930a7b)
}

func (item *TupleInt3BoxedMaybe) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	_ok, _jvalue, err := internal.Json2ReadMaybe("Maybe", in)
	if err != nil {
		return err
	}
	item.Ok = _ok
	if _ok {
		var in2Pointer *basictl.JsonLexer
		if _jvalue != nil {
			in2 := basictl.JsonLexer{Data: _jvalue}
			in2Pointer = &in2
		}
		if err := tlBuiltinTuple3Int.BuiltinTuple3IntReadJSON(legacyTypeNames, in2Pointer, &item.Value); err != nil {
			return err
		}
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *TupleInt3BoxedMaybe) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *TupleInt3BoxedMaybe) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *TupleInt3BoxedMaybe) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	if !item.Ok {
		return append(w, "{}"...)
	}
	w = append(w, `{"ok":true`...)
	w = append(w, `,"value":`...)
	w = tlBuiltinTuple3Int.BuiltinTuple3IntWriteJSONOpt(newTypeNames, short, w, &item.Value)
	return append(w, '}')
}

func (item TupleInt3BoxedMaybe) String() string {
	return string(item.WriteJSON(nil))
}
