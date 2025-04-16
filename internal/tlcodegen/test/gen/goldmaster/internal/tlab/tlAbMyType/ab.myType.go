// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlAbMyType

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type AbMyType struct {
	X int32
}

func (AbMyType) TLName() string { return "ab.myType" }
func (AbMyType) TLTag() uint32  { return 0xe0e96c86 }

func (item *AbMyType) Reset() {
	item.X = 0
}

func (item *AbMyType) FillRandom(rg *basictl.RandGenerator) {
	item.X = basictl.RandomInt(rg)
}

func (item *AbMyType) Read(w []byte) (_ []byte, err error) {
	return basictl.IntRead(w, &item.X)
}

// This method is general version of Write, use it instead!
func (item *AbMyType) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *AbMyType) Write(w []byte) []byte {
	w = basictl.IntWrite(w, item.X)
	return w
}

func (item *AbMyType) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xe0e96c86); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *AbMyType) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *AbMyType) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0xe0e96c86)
	return item.Write(w)
}

func (item *AbMyType) String() string {
	return string(item.WriteJSON(nil))
}

func (item *AbMyType) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propXPresented bool

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
					return internal.ErrorInvalidJSONWithDuplicatingKeys("ab.myType", "x")
				}
				if err := internal.Json2ReadInt32(in, &item.X); err != nil {
					return err
				}
				propXPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("ab.myType", key)
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
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *AbMyType) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *AbMyType) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *AbMyType) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexX := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"x":`...)
	w = basictl.JSONWriteInt32(w, item.X)
	if (item.X != 0) == false {
		w = w[:backupIndexX]
	}
	return append(w, '}')
}

func (item *AbMyType) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *AbMyType) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("ab.myType", err.Error())
	}
	return nil
}
