// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package internal

import (
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite

type NoStr struct {
	X int32
}

func (NoStr) TLName() string { return "noStr" }
func (NoStr) TLTag() uint32  { return 0x3a728324 }

func (item *NoStr) Reset() {
	item.X = 0
}

func (item *NoStr) FillRandom(rg *basictl.RandGenerator) {
	item.X = basictl.RandomInt(rg)
}

func (item *NoStr) Read(w []byte) (_ []byte, err error) {
	return basictl.IntRead(w, &item.X)
}

// This method is general version of Write, use it instead!
func (item *NoStr) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *NoStr) Write(w []byte) []byte {
	w = basictl.IntWrite(w, item.X)
	return w
}

func (item *NoStr) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x3a728324); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *NoStr) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *NoStr) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x3a728324)
	return item.Write(w)
}

func (item *NoStr) String() string {
	return string(item.WriteJSON(nil))
}

func (item *NoStr) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
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
					return ErrorInvalidJSONWithDuplicatingKeys("noStr", "x")
				}
				if err := Json2ReadInt32(in, &item.X); err != nil {
					return err
				}
				propXPresented = true
			default:
				return ErrorInvalidJSONExcessElement("noStr", key)
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
func (item *NoStr) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *NoStr) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *NoStr) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
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

func (item *NoStr) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *NoStr) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("noStr", err.Error())
	}
	return nil
}
