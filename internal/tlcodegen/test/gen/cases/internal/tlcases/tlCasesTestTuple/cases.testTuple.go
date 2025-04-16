// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlCasesTestTuple

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tl/tlBuiltinTuple4Int"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type CasesTestTuple struct {
	Tpl [4]int32
}

func (CasesTestTuple) TLName() string { return "cases.testTuple" }
func (CasesTestTuple) TLTag() uint32  { return 0x4b9caf8f }

func (item *CasesTestTuple) Reset() {
	tlBuiltinTuple4Int.BuiltinTuple4IntReset(&item.Tpl)
}

func (item *CasesTestTuple) FillRandom(rg *basictl.RandGenerator) {
	tlBuiltinTuple4Int.BuiltinTuple4IntFillRandom(rg, &item.Tpl)
}

func (item *CasesTestTuple) Read(w []byte) (_ []byte, err error) {
	return tlBuiltinTuple4Int.BuiltinTuple4IntRead(w, &item.Tpl)
}

// This method is general version of Write, use it instead!
func (item *CasesTestTuple) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *CasesTestTuple) Write(w []byte) []byte {
	w = tlBuiltinTuple4Int.BuiltinTuple4IntWrite(w, &item.Tpl)
	return w
}

func (item *CasesTestTuple) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x4b9caf8f); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *CasesTestTuple) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *CasesTestTuple) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x4b9caf8f)
	return item.Write(w)
}

func (item *CasesTestTuple) String() string {
	return string(item.WriteJSON(nil))
}

func (item *CasesTestTuple) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propTplPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "tpl":
				if propTplPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("cases.testTuple", "tpl")
				}
				if err := tlBuiltinTuple4Int.BuiltinTuple4IntReadJSON(legacyTypeNames, in, &item.Tpl); err != nil {
					return err
				}
				propTplPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("cases.testTuple", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propTplPresented {
		tlBuiltinTuple4Int.BuiltinTuple4IntReset(&item.Tpl)
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *CasesTestTuple) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *CasesTestTuple) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *CasesTestTuple) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"tpl":`...)
	w = tlBuiltinTuple4Int.BuiltinTuple4IntWriteJSONOpt(newTypeNames, short, w, &item.Tpl)
	return append(w, '}')
}

func (item *CasesTestTuple) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *CasesTestTuple) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("cases.testTuple", err.Error())
	}
	return nil
}
