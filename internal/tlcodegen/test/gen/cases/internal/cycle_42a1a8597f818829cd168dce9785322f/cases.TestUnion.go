// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package cycle_42a1a8597f818829cd168dce9785322f

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

var _CasesTestUnion = [2]internal.UnionElement{
	{TLTag: 0x4b4f09b1, TLName: "cases.testUnion1", TLString: "cases.testUnion1#4b4f09b1"},
	{TLTag: 0x464f96c4, TLName: "cases.testUnion2", TLString: "cases.testUnion2#464f96c4"},
}

type CasesTestUnion struct {
	value1 CasesTestUnion1
	value2 CasesTestUnion2
	index  int
}

func (item CasesTestUnion) TLName() string { return _CasesTestUnion[item.index].TLName }
func (item CasesTestUnion) TLTag() uint32  { return _CasesTestUnion[item.index].TLTag }

func (item *CasesTestUnion) Reset() { item.ResetTo1() }
func (item *CasesTestUnion) FillRandom(rg *basictl.RandGenerator) {
	index := basictl.RandomUint(rg) % 2
	switch index {
	case 0:
		item.index = 0
		item.value1.FillRandom(rg)
	case 1:
		item.index = 1
		item.value2.FillRandom(rg)
	default:
	}
}

func (item *CasesTestUnion) Is1() bool { return item.index == 0 }

func (item *CasesTestUnion) As1() (*CasesTestUnion1, bool) {
	if item.index != 0 {
		return nil, false
	}
	return &item.value1, true
}
func (item *CasesTestUnion) ResetTo1() *CasesTestUnion1 {
	item.index = 0
	item.value1.Reset()
	return &item.value1
}
func (item *CasesTestUnion) Set1(value CasesTestUnion1) {
	item.index = 0
	item.value1 = value
}

func (item *CasesTestUnion) Is2() bool { return item.index == 1 }

func (item *CasesTestUnion) As2() (*CasesTestUnion2, bool) {
	if item.index != 1 {
		return nil, false
	}
	return &item.value2, true
}
func (item *CasesTestUnion) ResetTo2() *CasesTestUnion2 {
	item.index = 1
	item.value2.Reset()
	return &item.value2
}
func (item *CasesTestUnion) Set2(value CasesTestUnion2) {
	item.index = 1
	item.value2 = value
}

func (item *CasesTestUnion) ReadBoxed(w []byte) (_ []byte, err error) {
	var tag uint32
	if w, err = basictl.NatRead(w, &tag); err != nil {
		return w, err
	}
	switch tag {
	case 0x4b4f09b1:
		item.index = 0
		return item.value1.Read(w)
	case 0x464f96c4:
		item.index = 1
		return item.value2.Read(w)
	default:
		return w, internal.ErrorInvalidUnionTag("cases.TestUnion", tag)
	}
}

// This method is general version of WriteBoxed, use it instead!
func (item *CasesTestUnion) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *CasesTestUnion) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, _CasesTestUnion[item.index].TLTag)
	switch item.index {
	case 0:
		w = item.value1.Write(w)
	case 1:
		w = item.value2.Write(w)
	}
	return w
}

func (item *CasesTestUnion) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	_tag, _value, err := internal.Json2ReadUnion("cases.TestUnion", in)
	if err != nil {
		return err
	}
	switch _tag {
	case "cases.testUnion1#4b4f09b1", "cases.testUnion1", "#4b4f09b1":
		if !legacyTypeNames && _tag == "cases.testUnion1#4b4f09b1" {
			return internal.ErrorInvalidUnionLegacyTagJSON("cases.TestUnion", "cases.testUnion1#4b4f09b1")
		}
		item.index = 0
		var in2Pointer *basictl.JsonLexer
		if _value != nil {
			in2 := basictl.JsonLexer{Data: _value}
			in2Pointer = &in2
		}
		if err := item.value1.ReadJSON(legacyTypeNames, in2Pointer); err != nil {
			return err
		}
	case "cases.testUnion2#464f96c4", "cases.testUnion2", "#464f96c4":
		if !legacyTypeNames && _tag == "cases.testUnion2#464f96c4" {
			return internal.ErrorInvalidUnionLegacyTagJSON("cases.TestUnion", "cases.testUnion2#464f96c4")
		}
		item.index = 1
		var in2Pointer *basictl.JsonLexer
		if _value != nil {
			in2 := basictl.JsonLexer{Data: _value}
			in2Pointer = &in2
		}
		if err := item.value2.ReadJSON(legacyTypeNames, in2Pointer); err != nil {
			return err
		}
	default:
		return internal.ErrorInvalidUnionTagJSON("cases.TestUnion", _tag)
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *CasesTestUnion) WriteJSONGeneral(w []byte) ([]byte, error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *CasesTestUnion) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *CasesTestUnion) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	switch item.index {
	case 0:
		if newTypeNames {
			w = append(w, `{"type":"cases.testUnion1"`...)
		} else {
			w = append(w, `{"type":"cases.testUnion1#4b4f09b1"`...)
		}
		w = append(w, `,"value":`...)
		w = item.value1.WriteJSONOpt(newTypeNames, short, w)
		return append(w, '}')
	case 1:
		if newTypeNames {
			w = append(w, `{"type":"cases.testUnion2"`...)
		} else {
			w = append(w, `{"type":"cases.testUnion2#464f96c4"`...)
		}
		w = append(w, `,"value":`...)
		w = item.value2.WriteJSONOpt(newTypeNames, short, w)
		return append(w, '}')
	default: // Impossible due to panic above
		return w
	}
}

func (item CasesTestUnion) String() string {
	return string(item.WriteJSON(nil))
}

func (item *CasesTestUnion) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *CasesTestUnion) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("cases.TestUnion", err.Error())
	}
	return nil
}

func (item CasesTestUnion1) AsUnion() CasesTestUnion {
	var ret CasesTestUnion
	ret.Set1(item)
	return ret
}

type CasesTestUnion1 struct {
	Value int32
}

func (CasesTestUnion1) TLName() string { return "cases.testUnion1" }
func (CasesTestUnion1) TLTag() uint32  { return 0x4b4f09b1 }

func (item *CasesTestUnion1) Reset() {
	item.Value = 0
}

func (item *CasesTestUnion1) FillRandom(rg *basictl.RandGenerator) {
	item.Value = basictl.RandomInt(rg)
}

func (item *CasesTestUnion1) Read(w []byte) (_ []byte, err error) {
	return basictl.IntRead(w, &item.Value)
}

// This method is general version of Write, use it instead!
func (item *CasesTestUnion1) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *CasesTestUnion1) Write(w []byte) []byte {
	w = basictl.IntWrite(w, item.Value)
	return w
}

func (item *CasesTestUnion1) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x4b4f09b1); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *CasesTestUnion1) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *CasesTestUnion1) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x4b4f09b1)
	return item.Write(w)
}

func (item CasesTestUnion1) String() string {
	return string(item.WriteJSON(nil))
}

func (item *CasesTestUnion1) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
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
					return internal.ErrorInvalidJSONWithDuplicatingKeys("cases.testUnion1", "value")
				}
				if err := internal.Json2ReadInt32(in, &item.Value); err != nil {
					return err
				}
				propValuePresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("cases.testUnion1", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propValuePresented {
		item.Value = 0
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *CasesTestUnion1) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *CasesTestUnion1) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *CasesTestUnion1) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexValue := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"value":`...)
	w = basictl.JSONWriteInt32(w, item.Value)
	if (item.Value != 0) == false {
		w = w[:backupIndexValue]
	}
	return append(w, '}')
}

func (item *CasesTestUnion1) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *CasesTestUnion1) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("cases.testUnion1", err.Error())
	}
	return nil
}

func (item CasesTestUnion2) AsUnion() CasesTestUnion {
	var ret CasesTestUnion
	ret.Set2(item)
	return ret
}

type CasesTestUnion2 struct {
	Value string
}

func (CasesTestUnion2) TLName() string { return "cases.testUnion2" }
func (CasesTestUnion2) TLTag() uint32  { return 0x464f96c4 }

func (item *CasesTestUnion2) Reset() {
	item.Value = ""
}

func (item *CasesTestUnion2) FillRandom(rg *basictl.RandGenerator) {
	item.Value = basictl.RandomString(rg)
}

func (item *CasesTestUnion2) Read(w []byte) (_ []byte, err error) {
	return basictl.StringRead(w, &item.Value)
}

// This method is general version of Write, use it instead!
func (item *CasesTestUnion2) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *CasesTestUnion2) Write(w []byte) []byte {
	w = basictl.StringWrite(w, item.Value)
	return w
}

func (item *CasesTestUnion2) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x464f96c4); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *CasesTestUnion2) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *CasesTestUnion2) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x464f96c4)
	return item.Write(w)
}

func (item CasesTestUnion2) String() string {
	return string(item.WriteJSON(nil))
}

func (item *CasesTestUnion2) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
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
					return internal.ErrorInvalidJSONWithDuplicatingKeys("cases.testUnion2", "value")
				}
				if err := internal.Json2ReadString(in, &item.Value); err != nil {
					return err
				}
				propValuePresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("cases.testUnion2", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propValuePresented {
		item.Value = ""
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *CasesTestUnion2) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *CasesTestUnion2) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *CasesTestUnion2) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexValue := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"value":`...)
	w = basictl.JSONWriteString(w, item.Value)
	if (len(item.Value) != 0) == false {
		w = w[:backupIndexValue]
	}
	return append(w, '}')
}

func (item *CasesTestUnion2) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *CasesTestUnion2) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("cases.testUnion2", err.Error())
	}
	return nil
}
