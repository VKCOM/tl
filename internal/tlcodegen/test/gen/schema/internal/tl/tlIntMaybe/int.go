// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlIntMaybe

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type IntMaybe struct {
	Value int32 // not deterministic if !Ok
	Ok    bool
}

func (item *IntMaybe) Reset() {
	item.Ok = false
}

func (item *IntMaybe) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.ReadBool(w, &item.Ok, 0x27930a7b, 0x3f9c8ef8); err != nil {
		return w, err
	}
	if item.Ok {
		return basictl.IntRead(w, &item.Value)
	}
	return w, nil
}

func (item *IntMaybe) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *IntMaybe) WriteBoxed(w []byte) []byte {
	if item.Ok {
		w = basictl.NatWrite(w, 0x3f9c8ef8)
		return basictl.IntWrite(w, item.Value)
	}
	return basictl.NatWrite(w, 0x27930a7b)
}

func (item *IntMaybe) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
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
		if err := internal.Json2ReadInt32(in2Pointer, &item.Value); err != nil {
			return err
		}
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *IntMaybe) WriteJSONGeneral(tctx *basictl.JSONWriteContext, w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(tctx, w), nil
}

func (item *IntMaybe) WriteJSON(w []byte) []byte {
	tctx := basictl.JSONWriteContext{}
	return item.WriteJSONOpt(&tctx, w)
}
func (item *IntMaybe) WriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte) []byte {
	if !item.Ok {
		return append(w, "{}"...)
	}
	w = append(w, `{"ok":true`...)
	if item.Value != 0 {
		w = append(w, `,"value":`...)
		w = basictl.JSONWriteInt32(w, item.Value)
	}
	return append(w, '}')
}

func (item IntMaybe) String() string {
	return string(item.WriteJSON(nil))
}
