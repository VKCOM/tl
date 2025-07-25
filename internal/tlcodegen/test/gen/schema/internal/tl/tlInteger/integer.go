// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlInteger

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type Integer struct {
	Value int32
}

func (Integer) TLName() string { return "integer" }
func (Integer) TLTag() uint32  { return 0x7e194796 }

func (item *Integer) Reset() {
	item.Value = 0
}

func (item *Integer) Read(w []byte) (_ []byte, err error) {
	return basictl.IntRead(w, &item.Value)
}

func (item *Integer) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *Integer) Write(w []byte) []byte {
	w = basictl.IntWrite(w, item.Value)
	return w
}

func (item *Integer) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x7e194796); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *Integer) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *Integer) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x7e194796)
	return item.Write(w)
}

func (item Integer) String() string {
	return string(item.WriteJSON(nil))
}

func (item *Integer) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propValuePresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "value":
				if propValuePresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("integer", "value")
				}
				if err := internal.Json2ReadInt32(in, &item.Value); err != nil {
					return err
				}
				propValuePresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("integer", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propValuePresented {
		item.Value = 0
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *Integer) WriteJSONGeneral(tctx *basictl.JSONWriteContext, w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(tctx, w), nil
}

func (item *Integer) WriteJSON(w []byte) []byte {
	tctx := basictl.JSONWriteContext{}
	return item.WriteJSONOpt(&tctx, w)
}
func (item *Integer) WriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte) []byte {
	w = append(w, '{')
	backupIndexValue := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"value":`...)
	w = basictl.JSONWriteInt32(w, item.Value)
	if (item.Value != 0) == false {
		w = w[:backupIndexValue]
	}
	return append(w, '}')
}

func (item *Integer) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *Integer) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("integer", err.Error())
	}
	return nil
}
