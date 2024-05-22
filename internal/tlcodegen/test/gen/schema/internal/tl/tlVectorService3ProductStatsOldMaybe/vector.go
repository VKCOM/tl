// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlVectorService3ProductStatsOldMaybe

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tl/tlBuiltinVectorService3ProductStatsOld"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tlservice3/tlService3ProductStatsOld"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type VectorService3ProductStatsOldMaybe struct {
	Value []tlService3ProductStatsOld.Service3ProductStatsOld // not deterministic if !Ok
	Ok    bool
}

func (item *VectorService3ProductStatsOldMaybe) Reset() {
	item.Ok = false
}

func (item *VectorService3ProductStatsOldMaybe) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.ReadBool(w, &item.Ok, 0x27930a7b, 0x3f9c8ef8); err != nil {
		return w, err
	}
	if item.Ok {
		return tlBuiltinVectorService3ProductStatsOld.BuiltinVectorService3ProductStatsOldRead(w, &item.Value)
	}
	return w, nil
}

// This method is general version of WriteBoxed, use it instead!
func (item *VectorService3ProductStatsOldMaybe) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *VectorService3ProductStatsOldMaybe) WriteBoxed(w []byte) []byte {
	if item.Ok {
		w = basictl.NatWrite(w, 0x3f9c8ef8)
		return tlBuiltinVectorService3ProductStatsOld.BuiltinVectorService3ProductStatsOldWrite(w, item.Value)
	}
	return basictl.NatWrite(w, 0x27930a7b)
}

func (item *VectorService3ProductStatsOldMaybe) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
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
		if err := tlBuiltinVectorService3ProductStatsOld.BuiltinVectorService3ProductStatsOldReadJSON(legacyTypeNames, in2Pointer, &item.Value); err != nil {
			return err
		}
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *VectorService3ProductStatsOldMaybe) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *VectorService3ProductStatsOldMaybe) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *VectorService3ProductStatsOldMaybe) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	if !item.Ok {
		return append(w, "{}"...)
	}
	w = append(w, `{"ok":true`...)
	if len(item.Value) != 0 {
		w = append(w, `,"value":`...)
		w = tlBuiltinVectorService3ProductStatsOld.BuiltinVectorService3ProductStatsOldWriteJSONOpt(newTypeNames, short, w, item.Value)
	}
	return append(w, '}')
}

func (item VectorService3ProductStatsOldMaybe) String() string {
	return string(item.WriteJSON(nil))
}
