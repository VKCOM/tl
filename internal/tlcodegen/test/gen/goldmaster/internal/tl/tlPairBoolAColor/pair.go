// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlPairBoolAColor

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlBool"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tla/tlAColor"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type PairBoolAColor struct {
	A bool
	B tlAColor.AColor
}

func (PairBoolAColor) TLName() string { return "pair" }
func (PairBoolAColor) TLTag() uint32  { return 0x0f3c47ab }

func (item *PairBoolAColor) Reset() {
	item.A = false
	item.B.Reset()
}

func (item *PairBoolAColor) FillRandom(rg *basictl.RandGenerator) {
	item.A = basictl.RandomUint(rg)&1 == 1
	item.B.FillRandom(rg)
}

func (item *PairBoolAColor) Read(w []byte) (_ []byte, err error) {
	if w, err = tlBool.BoolReadBoxed(w, &item.A); err != nil {
		return w, err
	}
	return item.B.ReadBoxed(w)
}

// This method is general version of Write, use it instead!
func (item *PairBoolAColor) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *PairBoolAColor) Write(w []byte) []byte {
	w = tlBool.BoolWriteBoxed(w, item.A)
	w = item.B.WriteBoxed(w)
	return w
}

func (item *PairBoolAColor) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x0f3c47ab); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *PairBoolAColor) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *PairBoolAColor) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x0f3c47ab)
	return item.Write(w)
}

func (item *PairBoolAColor) String() string {
	return string(item.WriteJSON(nil))
}

func (item *PairBoolAColor) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propAPresented bool
	var propBPresented bool

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
					return internal.ErrorInvalidJSONWithDuplicatingKeys("pair", "a")
				}
				if err := internal.Json2ReadBool(in, &item.A); err != nil {
					return err
				}
				propAPresented = true
			case "b":
				if propBPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("pair", "b")
				}
				if err := item.B.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propBPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("pair", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propAPresented {
		item.A = false
	}
	if !propBPresented {
		item.B.Reset()
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *PairBoolAColor) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *PairBoolAColor) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *PairBoolAColor) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexA := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"a":`...)
	w = basictl.JSONWriteBool(w, item.A)
	if (item.A) == false {
		w = w[:backupIndexA]
	}
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"b":`...)
	w = item.B.WriteJSONOpt(newTypeNames, short, w)
	return append(w, '}')
}

func (item *PairBoolAColor) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *PairBoolAColor) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("pair", err.Error())
	}
	return nil
}
