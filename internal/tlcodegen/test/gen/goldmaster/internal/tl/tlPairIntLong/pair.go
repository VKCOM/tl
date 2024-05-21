// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlPairIntLong

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type PairIntLong struct {
	A int32
	B int64
}

func (PairIntLong) TLName() string { return "pair" }
func (PairIntLong) TLTag() uint32  { return 0xf3c47ab }

func (item *PairIntLong) Reset() {
	item.A = 0
	item.B = 0
}

func (item *PairIntLong) FillRandom(rg *basictl.RandGenerator) {
	item.A = basictl.RandomInt(rg)
	item.B = basictl.RandomLong(rg)
}

func (item *PairIntLong) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.IntRead(w, &item.A); err != nil {
		return w, err
	}
	return basictl.LongRead(w, &item.B)
}

// This method is general version of Write, use it instead!
func (item *PairIntLong) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *PairIntLong) Write(w []byte) []byte {
	w = basictl.IntWrite(w, item.A)
	w = basictl.LongWrite(w, item.B)
	return w
}

func (item *PairIntLong) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xf3c47ab); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *PairIntLong) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *PairIntLong) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0xf3c47ab)
	return item.Write(w)
}

func (item PairIntLong) String() string {
	return string(item.WriteJSON(nil))
}

func (item *PairIntLong) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
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
				if err := internal.Json2ReadInt64(in, &item.B); err != nil {
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
func (item *PairIntLong) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *PairIntLong) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *PairIntLong) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
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
	w = basictl.JSONWriteInt64(w, item.B)
	if (item.B != 0) == false {
		w = w[:backupIndexB]
	}
	return append(w, '}')
}

func (item *PairIntLong) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *PairIntLong) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("pair", err.Error())
	}
	return nil
}
