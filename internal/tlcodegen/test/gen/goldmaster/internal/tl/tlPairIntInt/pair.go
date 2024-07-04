// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlPairIntInt

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type PairIntInt struct {
	A int32
	B int32
}

func (PairIntInt) TLName() string { return "pair" }
func (PairIntInt) TLTag() uint32  { return 0x0f3c47ab }

func (item *PairIntInt) Reset() {
	item.A = 0
	item.B = 0
}

func (item *PairIntInt) FillRandom(rg *basictl.RandGenerator) {
	item.A = basictl.RandomInt(rg)
	item.B = basictl.RandomInt(rg)
}

func (item *PairIntInt) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.IntRead(w, &item.A); err != nil {
		return w, err
	}
	return basictl.IntRead(w, &item.B)
}

// This method is general version of Write, use it instead!
func (item *PairIntInt) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *PairIntInt) Write(w []byte) []byte {
	w = basictl.IntWrite(w, item.A)
	w = basictl.IntWrite(w, item.B)
	return w
}

func (item *PairIntInt) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x0f3c47ab); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *PairIntInt) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *PairIntInt) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x0f3c47ab)
	return item.Write(w)
}

func (item PairIntInt) String() string {
	return string(item.WriteJSON(nil))
}

func (item *PairIntInt) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
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
				if err := internal.Json2ReadInt32(in, &item.A); err != nil {
					return err
				}
				propAPresented = true
			case "b":
				if propBPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("pair", "b")
				}
				if err := internal.Json2ReadInt32(in, &item.B); err != nil {
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
		item.A = 0
	}
	if !propBPresented {
		item.B = 0
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *PairIntInt) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *PairIntInt) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *PairIntInt) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexA := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"a":`...)
	w = basictl.JSONWriteInt32(w, item.A)
	if (item.A != 0) == false {
		w = w[:backupIndexA]
	}
	backupIndexB := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"b":`...)
	w = basictl.JSONWriteInt32(w, item.B)
	if (item.B != 0) == false {
		w = w[:backupIndexB]
	}
	return append(w, '}')
}

func (item *PairIntInt) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *PairIntInt) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("pair", err.Error())
	}
	return nil
}
