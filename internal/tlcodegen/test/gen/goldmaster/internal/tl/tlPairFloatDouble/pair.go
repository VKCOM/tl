// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlPairFloatDouble

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type PairFloatDouble struct {
	A float32
	B float64
}

func (PairFloatDouble) TLName() string { return "pair" }
func (PairFloatDouble) TLTag() uint32  { return 0xf3c47ab }

func (item *PairFloatDouble) Reset() {
	item.A = 0
	item.B = 0
}

func (item *PairFloatDouble) FillRandom(rg *basictl.RandGenerator) {
	item.A = basictl.RandomFloat(rg)
	item.B = basictl.RandomDouble(rg)
}

func (item *PairFloatDouble) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.FloatRead(w, &item.A); err != nil {
		return w, err
	}
	return basictl.DoubleRead(w, &item.B)
}

// This method is general version of Write, use it instead!
func (item *PairFloatDouble) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *PairFloatDouble) Write(w []byte) []byte {
	w = basictl.FloatWrite(w, item.A)
	w = basictl.DoubleWrite(w, item.B)
	return w
}

func (item *PairFloatDouble) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xf3c47ab); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *PairFloatDouble) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *PairFloatDouble) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0xf3c47ab)
	return item.Write(w)
}

func (item PairFloatDouble) String() string {
	return string(item.WriteJSON(nil))
}

func (item *PairFloatDouble) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
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
				if err := internal.Json2ReadFloat32(in, &item.A); err != nil {
					return err
				}
				propAPresented = true
			case "b":
				if propBPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("pair", "b")
				}
				if err := internal.Json2ReadFloat64(in, &item.B); err != nil {
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
func (item *PairFloatDouble) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *PairFloatDouble) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *PairFloatDouble) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexA := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"a":`...)
	w = basictl.JSONWriteFloat32(w, item.A)
	if (item.A != 0) == false {
		w = w[:backupIndexA]
	}
	backupIndexB := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"b":`...)
	w = basictl.JSONWriteFloat64(w, item.B)
	if (item.B != 0) == false {
		w = w[:backupIndexB]
	}
	return append(w, '}')
}

func (item *PairFloatDouble) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *PairFloatDouble) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("pair", err.Error())
	}
	return nil
}
