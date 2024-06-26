// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlCasesBytesTestDictStringString

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tl/tlBuiltinVectorDictionaryFieldString"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tl/tlDictionaryFieldString"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type CasesBytesTestDictStringString struct {
	Dict map[string]string
}

func (CasesBytesTestDictStringString) TLName() string { return "cases_bytes.testDictStringString" }
func (CasesBytesTestDictStringString) TLTag() uint32  { return 0xad69c772 }

func (item *CasesBytesTestDictStringString) Reset() {
	tlBuiltinVectorDictionaryFieldString.BuiltinVectorDictionaryFieldStringReset(item.Dict)
}

func (item *CasesBytesTestDictStringString) FillRandom(rg *basictl.RandGenerator) {
	tlBuiltinVectorDictionaryFieldString.BuiltinVectorDictionaryFieldStringFillRandom(rg, &item.Dict)
}

func (item *CasesBytesTestDictStringString) Read(w []byte) (_ []byte, err error) {
	return tlBuiltinVectorDictionaryFieldString.BuiltinVectorDictionaryFieldStringRead(w, &item.Dict)
}

// This method is general version of Write, use it instead!
func (item *CasesBytesTestDictStringString) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *CasesBytesTestDictStringString) Write(w []byte) []byte {
	w = tlBuiltinVectorDictionaryFieldString.BuiltinVectorDictionaryFieldStringWrite(w, item.Dict)
	return w
}

func (item *CasesBytesTestDictStringString) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xad69c772); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *CasesBytesTestDictStringString) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *CasesBytesTestDictStringString) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0xad69c772)
	return item.Write(w)
}

func (item CasesBytesTestDictStringString) String() string {
	return string(item.WriteJSON(nil))
}

func (item *CasesBytesTestDictStringString) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propDictPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "dict":
				if propDictPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("cases_bytes.testDictStringString", "dict")
				}
				if err := tlBuiltinVectorDictionaryFieldString.BuiltinVectorDictionaryFieldStringReadJSON(legacyTypeNames, in, &item.Dict); err != nil {
					return err
				}
				propDictPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("cases_bytes.testDictStringString", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propDictPresented {
		tlBuiltinVectorDictionaryFieldString.BuiltinVectorDictionaryFieldStringReset(item.Dict)
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *CasesBytesTestDictStringString) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *CasesBytesTestDictStringString) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *CasesBytesTestDictStringString) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexDict := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"dict":`...)
	w = tlBuiltinVectorDictionaryFieldString.BuiltinVectorDictionaryFieldStringWriteJSONOpt(newTypeNames, short, w, item.Dict)
	if (len(item.Dict) != 0) == false {
		w = w[:backupIndexDict]
	}
	return append(w, '}')
}

func (item *CasesBytesTestDictStringString) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *CasesBytesTestDictStringString) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("cases_bytes.testDictStringString", err.Error())
	}
	return nil
}

type CasesBytesTestDictStringStringBytes struct {
	Dict []tlDictionaryFieldString.DictionaryFieldStringBytes
}

func (CasesBytesTestDictStringStringBytes) TLName() string { return "cases_bytes.testDictStringString" }
func (CasesBytesTestDictStringStringBytes) TLTag() uint32  { return 0xad69c772 }

func (item *CasesBytesTestDictStringStringBytes) Reset() {
	item.Dict = item.Dict[:0]
}

func (item *CasesBytesTestDictStringStringBytes) FillRandom(rg *basictl.RandGenerator) {
	tlBuiltinVectorDictionaryFieldString.BuiltinVectorDictionaryFieldStringBytesFillRandom(rg, &item.Dict)
}

func (item *CasesBytesTestDictStringStringBytes) Read(w []byte) (_ []byte, err error) {
	return tlBuiltinVectorDictionaryFieldString.BuiltinVectorDictionaryFieldStringBytesRead(w, &item.Dict)
}

// This method is general version of Write, use it instead!
func (item *CasesBytesTestDictStringStringBytes) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *CasesBytesTestDictStringStringBytes) Write(w []byte) []byte {
	w = tlBuiltinVectorDictionaryFieldString.BuiltinVectorDictionaryFieldStringBytesWrite(w, item.Dict)
	return w
}

func (item *CasesBytesTestDictStringStringBytes) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xad69c772); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *CasesBytesTestDictStringStringBytes) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *CasesBytesTestDictStringStringBytes) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0xad69c772)
	return item.Write(w)
}

func (item CasesBytesTestDictStringStringBytes) String() string {
	return string(item.WriteJSON(nil))
}

func (item *CasesBytesTestDictStringStringBytes) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propDictPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "dict":
				if propDictPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("cases_bytes.testDictStringString", "dict")
				}
				if err := tlBuiltinVectorDictionaryFieldString.BuiltinVectorDictionaryFieldStringBytesReadJSON(legacyTypeNames, in, &item.Dict); err != nil {
					return err
				}
				propDictPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("cases_bytes.testDictStringString", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propDictPresented {
		item.Dict = item.Dict[:0]
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *CasesBytesTestDictStringStringBytes) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *CasesBytesTestDictStringStringBytes) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *CasesBytesTestDictStringStringBytes) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexDict := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"dict":`...)
	w = tlBuiltinVectorDictionaryFieldString.BuiltinVectorDictionaryFieldStringBytesWriteJSONOpt(newTypeNames, short, w, item.Dict)
	if (len(item.Dict) != 0) == false {
		w = w[:backupIndexDict]
	}
	return append(w, '}')
}

func (item *CasesBytesTestDictStringStringBytes) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *CasesBytesTestDictStringStringBytes) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("cases_bytes.testDictStringString", err.Error())
	}
	return nil
}
