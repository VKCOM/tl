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

type Replace3 struct {
	A [3]int32
}

func (Replace3) TLName() string { return "replace3" }
func (Replace3) TLTag() uint32  { return 0x51e324e4 }

func (item *Replace3) Reset() {
	BuiltinTuple3IntReset(&item.A)
}

func (item *Replace3) FillRandom(rg *basictl.RandGenerator) {
	BuiltinTuple3IntFillRandom(rg, &item.A)
}

func (item *Replace3) Read(w []byte) (_ []byte, err error) {
	return BuiltinTuple3IntRead(w, &item.A)
}

// This method is general version of Write, use it instead!
func (item *Replace3) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *Replace3) Write(w []byte) []byte {
	w = BuiltinTuple3IntWrite(w, &item.A)
	return w
}

func (item *Replace3) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x51e324e4); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *Replace3) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *Replace3) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x51e324e4)
	return item.Write(w)
}

func (item Replace3) String() string {
	return string(item.WriteJSON(nil))
}

func (item *Replace3) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propAPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "a":
				if propAPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("replace3", "a")
				}
				if err := BuiltinTuple3IntReadJSON(legacyTypeNames, in, &item.A); err != nil {
					return err
				}
				propAPresented = true
			default:
				return ErrorInvalidJSONExcessElement("replace3", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propAPresented {
		BuiltinTuple3IntReset(&item.A)
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *Replace3) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *Replace3) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *Replace3) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"a":`...)
	w = BuiltinTuple3IntWriteJSONOpt(newTypeNames, short, w, &item.A)
	return append(w, '}')
}

func (item *Replace3) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *Replace3) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("replace3", err.Error())
	}
	return nil
}
