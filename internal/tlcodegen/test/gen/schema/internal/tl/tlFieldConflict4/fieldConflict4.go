// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlFieldConflict4

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type FieldConflict4 struct {
	X    int32
	SetX int32
}

func (FieldConflict4) TLName() string { return "fieldConflict4" }
func (FieldConflict4) TLTag() uint32  { return 0xd93c186a }

func (item *FieldConflict4) Reset() {
	item.X = 0
	item.SetX = 0
}

func (item *FieldConflict4) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.IntRead(w, &item.X); err != nil {
		return w, err
	}
	return basictl.IntRead(w, &item.SetX)
}

func (item *FieldConflict4) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *FieldConflict4) Write(w []byte) []byte {
	w = basictl.IntWrite(w, item.X)
	w = basictl.IntWrite(w, item.SetX)
	return w
}

func (item *FieldConflict4) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xd93c186a); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *FieldConflict4) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *FieldConflict4) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0xd93c186a)
	return item.Write(w)
}

func (item FieldConflict4) String() string {
	return string(item.WriteJSON(nil))
}

func (item *FieldConflict4) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propXPresented bool
	var propSetXPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "X":
				if propXPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("fieldConflict4", "X")
				}
				if err := internal.Json2ReadInt32(in, &item.X); err != nil {
					return err
				}
				propXPresented = true
			case "SetX":
				if propSetXPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("fieldConflict4", "SetX")
				}
				if err := internal.Json2ReadInt32(in, &item.SetX); err != nil {
					return err
				}
				propSetXPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("fieldConflict4", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propXPresented {
		item.X = 0
	}
	if !propSetXPresented {
		item.SetX = 0
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *FieldConflict4) WriteJSONGeneral(tctx *basictl.JSONWriteContext, w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(tctx, w), nil
}

func (item *FieldConflict4) WriteJSON(w []byte) []byte {
	tctx := basictl.JSONWriteContext{}
	return item.WriteJSONOpt(&tctx, w)
}
func (item *FieldConflict4) WriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte) []byte {
	w = append(w, '{')
	backupIndexX := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"X":`...)
	w = basictl.JSONWriteInt32(w, item.X)
	if (item.X != 0) == false {
		w = w[:backupIndexX]
	}
	backupIndexSetX := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"SetX":`...)
	w = basictl.JSONWriteInt32(w, item.SetX)
	if (item.SetX != 0) == false {
		w = w[:backupIndexSetX]
	}
	return append(w, '}')
}

func (item *FieldConflict4) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *FieldConflict4) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("fieldConflict4", err.Error())
	}
	return nil
}
