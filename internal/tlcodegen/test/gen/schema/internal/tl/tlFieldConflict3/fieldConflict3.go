// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlFieldConflict3

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type FieldConflict3 struct {
	X    int32
	SetX int32
}

func (FieldConflict3) TLName() string { return "fieldConflict3" }
func (FieldConflict3) TLTag() uint32  { return 0x2cf6e157 }

func (item *FieldConflict3) Reset() {
	item.X = 0
	item.SetX = 0
}

func (item *FieldConflict3) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.IntRead(w, &item.X); err != nil {
		return w, err
	}
	return basictl.IntRead(w, &item.SetX)
}

// This method is general version of Write, use it instead!
func (item *FieldConflict3) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *FieldConflict3) Write(w []byte) []byte {
	w = basictl.IntWrite(w, item.X)
	w = basictl.IntWrite(w, item.SetX)
	return w
}

func (item *FieldConflict3) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x2cf6e157); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *FieldConflict3) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *FieldConflict3) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x2cf6e157)
	return item.Write(w)
}

func (item FieldConflict3) String() string {
	return string(item.WriteJSON(nil))
}

func (item *FieldConflict3) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
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
			case "x":
				if propXPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("fieldConflict3", "x")
				}
				if err := internal.Json2ReadInt32(in, &item.X); err != nil {
					return err
				}
				propXPresented = true
			case "SetX":
				if propSetXPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("fieldConflict3", "SetX")
				}
				if err := internal.Json2ReadInt32(in, &item.SetX); err != nil {
					return err
				}
				propSetXPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("fieldConflict3", key)
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
func (item *FieldConflict3) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *FieldConflict3) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *FieldConflict3) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexX := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"x":`...)
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

func (item *FieldConflict3) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *FieldConflict3) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("fieldConflict3", err.Error())
	}
	return nil
}
