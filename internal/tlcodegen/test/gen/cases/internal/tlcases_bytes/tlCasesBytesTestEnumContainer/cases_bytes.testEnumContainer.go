// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlCasesBytesTestEnumContainer

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tlcases/tlCasesTestEnum"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type CasesBytesTestEnumContainer struct {
	Value tlCasesTestEnum.CasesTestEnum
}

func (CasesBytesTestEnumContainer) TLName() string { return "cases_bytes.testEnumContainer" }
func (CasesBytesTestEnumContainer) TLTag() uint32  { return 0x32b92037 }

func (item *CasesBytesTestEnumContainer) Reset() {
	item.Value.Reset()
}

func (item *CasesBytesTestEnumContainer) FillRandom(rg *basictl.RandGenerator) {
	item.Value.FillRandom(rg)
}

func (item *CasesBytesTestEnumContainer) Read(w []byte) (_ []byte, err error) {
	return item.Value.ReadBoxed(w)
}

// This method is general version of Write, use it instead!
func (item *CasesBytesTestEnumContainer) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *CasesBytesTestEnumContainer) Write(w []byte) []byte {
	w = item.Value.WriteBoxed(w)
	return w
}

func (item *CasesBytesTestEnumContainer) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x32b92037); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *CasesBytesTestEnumContainer) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *CasesBytesTestEnumContainer) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x32b92037)
	return item.Write(w)
}

func (item CasesBytesTestEnumContainer) String() string {
	return string(item.WriteJSON(nil))
}

func (item *CasesBytesTestEnumContainer) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propValuePresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "value":
				if propValuePresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("cases_bytes.testEnumContainer", "value")
				}
				if err := item.Value.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propValuePresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("cases_bytes.testEnumContainer", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propValuePresented {
		item.Value.Reset()
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *CasesBytesTestEnumContainer) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *CasesBytesTestEnumContainer) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *CasesBytesTestEnumContainer) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"value":`...)
	w = item.Value.WriteJSONOpt(newTypeNames, short, w)
	return append(w, '}')
}

func (item *CasesBytesTestEnumContainer) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *CasesBytesTestEnumContainer) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("cases_bytes.testEnumContainer", err.Error())
	}
	return nil
}
