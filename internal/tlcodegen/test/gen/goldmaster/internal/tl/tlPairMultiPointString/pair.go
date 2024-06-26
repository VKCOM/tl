// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlPairMultiPointString

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlMultiPoint"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type PairMultiPointString struct {
	A tlMultiPoint.MultiPoint
	B string
}

func (PairMultiPointString) TLName() string { return "pair" }
func (PairMultiPointString) TLTag() uint32  { return 0x0f3c47ab }

func (item *PairMultiPointString) Reset() {
	item.A.Reset()
	item.B = ""
}

func (item *PairMultiPointString) FillRandom(rg *basictl.RandGenerator) {
	item.A.FillRandom(rg)
	item.B = basictl.RandomString(rg)
}

func (item *PairMultiPointString) Read(w []byte) (_ []byte, err error) {
	if w, err = item.A.Read(w); err != nil {
		return w, err
	}
	return basictl.StringRead(w, &item.B)
}

// This method is general version of Write, use it instead!
func (item *PairMultiPointString) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *PairMultiPointString) Write(w []byte) []byte {
	w = item.A.Write(w)
	w = basictl.StringWrite(w, item.B)
	return w
}

func (item *PairMultiPointString) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x0f3c47ab); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *PairMultiPointString) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *PairMultiPointString) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x0f3c47ab)
	return item.Write(w)
}

func (item PairMultiPointString) String() string {
	return string(item.WriteJSON(nil))
}

func (item *PairMultiPointString) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
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
				if err := item.A.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propAPresented = true
			case "b":
				if propBPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("pair", "b")
				}
				if err := internal.Json2ReadString(in, &item.B); err != nil {
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
		item.A.Reset()
	}
	if !propBPresented {
		item.B = ""
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *PairMultiPointString) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *PairMultiPointString) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *PairMultiPointString) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"a":`...)
	w = item.A.WriteJSONOpt(newTypeNames, short, w)
	backupIndexB := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"b":`...)
	w = basictl.JSONWriteString(w, item.B)
	if (len(item.B) != 0) == false {
		w = w[:backupIndexB]
	}
	return append(w, '}')
}

func (item *PairMultiPointString) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *PairMultiPointString) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("pair", err.Error())
	}
	return nil
}
