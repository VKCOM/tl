// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlCasesBytesTestVector

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tl/tlBuiltinVectorString"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type CasesBytesTestVector struct {
	Arr []string
}

func (CasesBytesTestVector) TLName() string { return "cases_bytes.testVector" }
func (CasesBytesTestVector) TLTag() uint32  { return 0x3647c8ae }

func (item *CasesBytesTestVector) Reset() {
	item.Arr = item.Arr[:0]
}

func (item *CasesBytesTestVector) FillRandom(rg *basictl.RandGenerator) {
	tlBuiltinVectorString.BuiltinVectorStringFillRandom(rg, &item.Arr)
}

func (item *CasesBytesTestVector) Read(w []byte) (_ []byte, err error) {
	return tlBuiltinVectorString.BuiltinVectorStringRead(w, &item.Arr)
}

// This method is general version of Write, use it instead!
func (item *CasesBytesTestVector) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *CasesBytesTestVector) Write(w []byte) []byte {
	w = tlBuiltinVectorString.BuiltinVectorStringWrite(w, item.Arr)
	return w
}

func (item *CasesBytesTestVector) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x3647c8ae); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *CasesBytesTestVector) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *CasesBytesTestVector) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x3647c8ae)
	return item.Write(w)
}

func (item CasesBytesTestVector) String() string {
	return string(item.WriteJSON(nil))
}

func (item *CasesBytesTestVector) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propArrPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "arr":
				if propArrPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("cases_bytes.testVector", "arr")
				}
				if err := tlBuiltinVectorString.BuiltinVectorStringReadJSON(legacyTypeNames, in, &item.Arr); err != nil {
					return err
				}
				propArrPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("cases_bytes.testVector", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propArrPresented {
		item.Arr = item.Arr[:0]
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *CasesBytesTestVector) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *CasesBytesTestVector) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *CasesBytesTestVector) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexArr := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"arr":`...)
	w = tlBuiltinVectorString.BuiltinVectorStringWriteJSONOpt(newTypeNames, short, w, item.Arr)
	if (len(item.Arr) != 0) == false {
		w = w[:backupIndexArr]
	}
	return append(w, '}')
}

func (item *CasesBytesTestVector) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *CasesBytesTestVector) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("cases_bytes.testVector", err.Error())
	}
	return nil
}

type CasesBytesTestVectorBytes struct {
	Arr [][]byte
}

func (CasesBytesTestVectorBytes) TLName() string { return "cases_bytes.testVector" }
func (CasesBytesTestVectorBytes) TLTag() uint32  { return 0x3647c8ae }

func (item *CasesBytesTestVectorBytes) Reset() {
	item.Arr = item.Arr[:0]
}

func (item *CasesBytesTestVectorBytes) FillRandom(rg *basictl.RandGenerator) {
	tlBuiltinVectorString.BuiltinVectorStringBytesFillRandom(rg, &item.Arr)
}

func (item *CasesBytesTestVectorBytes) Read(w []byte) (_ []byte, err error) {
	return tlBuiltinVectorString.BuiltinVectorStringBytesRead(w, &item.Arr)
}

// This method is general version of Write, use it instead!
func (item *CasesBytesTestVectorBytes) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *CasesBytesTestVectorBytes) Write(w []byte) []byte {
	w = tlBuiltinVectorString.BuiltinVectorStringBytesWrite(w, item.Arr)
	return w
}

func (item *CasesBytesTestVectorBytes) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x3647c8ae); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *CasesBytesTestVectorBytes) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *CasesBytesTestVectorBytes) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x3647c8ae)
	return item.Write(w)
}

func (item CasesBytesTestVectorBytes) String() string {
	return string(item.WriteJSON(nil))
}

func (item *CasesBytesTestVectorBytes) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propArrPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "arr":
				if propArrPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("cases_bytes.testVector", "arr")
				}
				if err := tlBuiltinVectorString.BuiltinVectorStringBytesReadJSON(legacyTypeNames, in, &item.Arr); err != nil {
					return err
				}
				propArrPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("cases_bytes.testVector", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propArrPresented {
		item.Arr = item.Arr[:0]
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *CasesBytesTestVectorBytes) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *CasesBytesTestVectorBytes) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *CasesBytesTestVectorBytes) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexArr := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"arr":`...)
	w = tlBuiltinVectorString.BuiltinVectorStringBytesWriteJSONOpt(newTypeNames, short, w, item.Arr)
	if (len(item.Arr) != 0) == false {
		w = w[:backupIndexArr]
	}
	return append(w, '}')
}

func (item *CasesBytesTestVectorBytes) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *CasesBytesTestVectorBytes) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("cases_bytes.testVector", err.Error())
	}
	return nil
}
