// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlCasesBytesTestDictString

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tl/tlBuiltinVectorDictionaryFieldInt"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tl/tlDictionaryFieldInt"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type CasesBytesTestDictString struct {
	Dict map[string]int32
}

func (CasesBytesTestDictString) TLName() string { return "cases_bytes.testDictString" }
func (CasesBytesTestDictString) TLTag() uint32  { return 0x6c04d6ce }

func (item *CasesBytesTestDictString) Reset() {
	tlBuiltinVectorDictionaryFieldInt.BuiltinVectorDictionaryFieldIntReset(item.Dict)
}

func (item *CasesBytesTestDictString) FillRandom(rg *basictl.RandGenerator) {
	tlBuiltinVectorDictionaryFieldInt.BuiltinVectorDictionaryFieldIntFillRandom(rg, &item.Dict)
}

func (item *CasesBytesTestDictString) Read(w []byte) (_ []byte, err error) {
	return tlBuiltinVectorDictionaryFieldInt.BuiltinVectorDictionaryFieldIntRead(w, &item.Dict)
}

// This method is general version of Write, use it instead!
func (item *CasesBytesTestDictString) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *CasesBytesTestDictString) Write(w []byte) []byte {
	w = tlBuiltinVectorDictionaryFieldInt.BuiltinVectorDictionaryFieldIntWrite(w, item.Dict)
	return w
}

func (item *CasesBytesTestDictString) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x6c04d6ce); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *CasesBytesTestDictString) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *CasesBytesTestDictString) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x6c04d6ce)
	return item.Write(w)
}

func (item CasesBytesTestDictString) String() string {
	return string(item.WriteJSON(nil))
}

func (item *CasesBytesTestDictString) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
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
					return internal.ErrorInvalidJSONWithDuplicatingKeys("cases_bytes.testDictString", "dict")
				}
				if err := tlBuiltinVectorDictionaryFieldInt.BuiltinVectorDictionaryFieldIntReadJSON(legacyTypeNames, in, &item.Dict); err != nil {
					return err
				}
				propDictPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("cases_bytes.testDictString", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propDictPresented {
		tlBuiltinVectorDictionaryFieldInt.BuiltinVectorDictionaryFieldIntReset(item.Dict)
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *CasesBytesTestDictString) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *CasesBytesTestDictString) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *CasesBytesTestDictString) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexDict := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"dict":`...)
	w = tlBuiltinVectorDictionaryFieldInt.BuiltinVectorDictionaryFieldIntWriteJSONOpt(newTypeNames, short, w, item.Dict)
	if (len(item.Dict) != 0) == false {
		w = w[:backupIndexDict]
	}
	return append(w, '}')
}

func (item *CasesBytesTestDictString) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *CasesBytesTestDictString) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("cases_bytes.testDictString", err.Error())
	}
	return nil
}

type CasesBytesTestDictStringBytes struct {
	Dict []tlDictionaryFieldInt.DictionaryFieldIntBytes
}

func (CasesBytesTestDictStringBytes) TLName() string { return "cases_bytes.testDictString" }
func (CasesBytesTestDictStringBytes) TLTag() uint32  { return 0x6c04d6ce }

func (item *CasesBytesTestDictStringBytes) Reset() {
	item.Dict = item.Dict[:0]
}

func (item *CasesBytesTestDictStringBytes) FillRandom(rg *basictl.RandGenerator) {
	tlBuiltinVectorDictionaryFieldInt.BuiltinVectorDictionaryFieldIntBytesFillRandom(rg, &item.Dict)
}

func (item *CasesBytesTestDictStringBytes) Read(w []byte) (_ []byte, err error) {
	return tlBuiltinVectorDictionaryFieldInt.BuiltinVectorDictionaryFieldIntBytesRead(w, &item.Dict)
}

// This method is general version of Write, use it instead!
func (item *CasesBytesTestDictStringBytes) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *CasesBytesTestDictStringBytes) Write(w []byte) []byte {
	w = tlBuiltinVectorDictionaryFieldInt.BuiltinVectorDictionaryFieldIntBytesWrite(w, item.Dict)
	return w
}

func (item *CasesBytesTestDictStringBytes) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x6c04d6ce); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *CasesBytesTestDictStringBytes) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *CasesBytesTestDictStringBytes) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x6c04d6ce)
	return item.Write(w)
}

func (item CasesBytesTestDictStringBytes) String() string {
	return string(item.WriteJSON(nil))
}

func (item *CasesBytesTestDictStringBytes) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
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
					return internal.ErrorInvalidJSONWithDuplicatingKeys("cases_bytes.testDictString", "dict")
				}
				if err := tlBuiltinVectorDictionaryFieldInt.BuiltinVectorDictionaryFieldIntBytesReadJSON(legacyTypeNames, in, &item.Dict); err != nil {
					return err
				}
				propDictPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("cases_bytes.testDictString", key)
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
func (item *CasesBytesTestDictStringBytes) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *CasesBytesTestDictStringBytes) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *CasesBytesTestDictStringBytes) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexDict := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"dict":`...)
	w = tlBuiltinVectorDictionaryFieldInt.BuiltinVectorDictionaryFieldIntBytesWriteJSONOpt(newTypeNames, short, w, item.Dict)
	if (len(item.Dict) != 0) == false {
		w = w[:backupIndexDict]
	}
	return append(w, '}')
}

func (item *CasesBytesTestDictStringBytes) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *CasesBytesTestDictStringBytes) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("cases_bytes.testDictString", err.Error())
	}
	return nil
}
