// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package internal

import (
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite

var _MyNat2 = [2]UnionElement{
	{TLTag: 0x8d868379, TLName: "myZero", TLString: "myZero#8d868379"},
	{TLTag: 0x79e0c6df, TLName: "myPlus", TLString: "myPlus#79e0c6df"},
}

type MyNat2 struct {
	valueMyPlus *MyPlus
	index       int
}

func (item MyNat2) TLName() string { return _MyNat2[item.index].TLName }
func (item MyNat2) TLTag() uint32  { return _MyNat2[item.index].TLTag }

func (item *MyNat2) Reset() { item.index = 0 }
func (item *MyNat2) FillRandom(rg *basictl.RandGenerator) {
	index := basictl.RandomUint(rg) % 2
	switch index {
	case 0:
		item.index = 0
	case 1:
		item.index = 1
		if item.valueMyPlus == nil {
			var value MyPlus
			value.FillRandom(rg)
			item.valueMyPlus = &value
		}
	default:
	}
}

func (item *MyNat2) IsMyZero() bool { return item.index == 0 }

func (item *MyNat2) AsMyZero() (MyZero, bool) {
	var value MyZero
	return value, item.index == 0
}
func (item *MyNat2) ResetToMyZero() { item.index = 0 }
func (item *MyNat2) SetMyZero()     { item.index = 0 }

func (item *MyNat2) IsMyPlus() bool { return item.index == 1 }

func (item *MyNat2) AsMyPlus() (*MyPlus, bool) {
	if item.index != 1 {
		return nil, false
	}
	return item.valueMyPlus, true
}
func (item *MyNat2) ResetToMyPlus() *MyPlus {
	item.index = 1
	if item.valueMyPlus == nil {
		var value MyPlus
		item.valueMyPlus = &value
	} else {
		item.valueMyPlus.Reset()
	}
	return item.valueMyPlus
}
func (item *MyNat2) SetMyPlus(value MyPlus) {
	item.index = 1
	if item.valueMyPlus == nil {
		item.valueMyPlus = &value
	} else {
		*item.valueMyPlus = value
	}
}

func (item *MyNat2) ReadBoxed(w []byte) (_ []byte, err error) {
	var tag uint32
	if w, err = basictl.NatRead(w, &tag); err != nil {
		return w, err
	}
	switch tag {
	case 0x8d868379:
		item.index = 0
		return w, nil
	case 0x79e0c6df:
		item.index = 1
		if item.valueMyPlus == nil {
			var value MyPlus
			item.valueMyPlus = &value
		}
		return item.valueMyPlus.Read(w)
	default:
		return w, ErrorInvalidUnionTag("MyNat2", tag)
	}
}

// This method is general version of WriteBoxed, use it instead!
func (item *MyNat2) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *MyNat2) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, _MyNat2[item.index].TLTag)
	switch item.index {
	case 0:
		return w
	case 1:
		w = item.valueMyPlus.Write(w)
	}
	return w
}

func (item *MyNat2) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	_tag, _value, err := Json2ReadUnion("MyNat2", in)
	if err != nil {
		return err
	}
	switch _tag {
	case "myZero#8d868379", "myZero", "#8d868379":
		if !legacyTypeNames && _tag == "myZero#8d868379" {
			return ErrorInvalidUnionLegacyTagJSON("MyNat2", "myZero#8d868379")
		}
		item.index = 0
	case "myPlus#79e0c6df", "myPlus", "#79e0c6df":
		if !legacyTypeNames && _tag == "myPlus#79e0c6df" {
			return ErrorInvalidUnionLegacyTagJSON("MyNat2", "myPlus#79e0c6df")
		}
		item.index = 1
		if item.valueMyPlus == nil {
			var value MyPlus
			item.valueMyPlus = &value
		}
		var in2Pointer *basictl.JsonLexer
		if _value != nil {
			in2 := basictl.JsonLexer{Data: _value}
			in2Pointer = &in2
		}
		if err := item.valueMyPlus.ReadJSON(legacyTypeNames, in2Pointer); err != nil {
			return err
		}
	default:
		return ErrorInvalidUnionTagJSON("MyNat2", _tag)
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *MyNat2) WriteJSONGeneral(w []byte) ([]byte, error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *MyNat2) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *MyNat2) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	switch item.index {
	case 0:
		if newTypeNames {
			w = append(w, `{"type":"myZero"`...)
		} else {
			w = append(w, `{"type":"myZero#8d868379"`...)
		}
		return append(w, '}')
	case 1:
		if newTypeNames {
			w = append(w, `{"type":"myPlus"`...)
		} else {
			w = append(w, `{"type":"myPlus#79e0c6df"`...)
		}
		w = append(w, `,"value":`...)
		w = item.valueMyPlus.WriteJSONOpt(newTypeNames, short, w)
		return append(w, '}')
	default: // Impossible due to panic above
		return w
	}
}

func (item *MyNat2) String() string {
	return string(item.WriteJSON(nil))
}

func (item *MyNat2) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *MyNat2) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("MyNat2", err.Error())
	}
	return nil
}

func (item MyPlus) AsUnion() MyNat2 {
	var ret MyNat2
	ret.SetMyPlus(item)
	return ret
}

type MyPlus struct {
	A MyNat2
}

func (MyPlus) TLName() string { return "myPlus" }
func (MyPlus) TLTag() uint32  { return 0x79e0c6df }

func (item *MyPlus) Reset() {
	item.A.Reset()
}

func (item *MyPlus) FillRandom(rg *basictl.RandGenerator) {
	item.A.FillRandom(rg)
}

func (item *MyPlus) Read(w []byte) (_ []byte, err error) {
	return item.A.ReadBoxed(w)
}

// This method is general version of Write, use it instead!
func (item *MyPlus) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *MyPlus) Write(w []byte) []byte {
	w = item.A.WriteBoxed(w)
	return w
}

func (item *MyPlus) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x79e0c6df); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *MyPlus) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *MyPlus) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x79e0c6df)
	return item.Write(w)
}

func (item *MyPlus) String() string {
	return string(item.WriteJSON(nil))
}

func (item *MyPlus) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propAPresented bool

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
					return ErrorInvalidJSONWithDuplicatingKeys("myPlus", "a")
				}
				if err := item.A.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propAPresented = true
			default:
				return ErrorInvalidJSONExcessElement("myPlus", key)
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
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *MyPlus) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *MyPlus) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *MyPlus) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"a":`...)
	w = item.A.WriteJSONOpt(newTypeNames, short, w)
	return append(w, '}')
}

func (item *MyPlus) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *MyPlus) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("myPlus", err.Error())
	}
	return nil
}

func (item MyZero) AsUnion() MyNat2 {
	var ret MyNat2
	ret.SetMyZero()
	return ret
}

type MyZero struct {
}

func (MyZero) TLName() string { return "myZero" }
func (MyZero) TLTag() uint32  { return 0x8d868379 }

func (item *MyZero) Reset() {}

func (item *MyZero) FillRandom(rg *basictl.RandGenerator) {}

func (item *MyZero) Read(w []byte) (_ []byte, err error) { return w, nil }

// This method is general version of Write, use it instead!
func (item *MyZero) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *MyZero) Write(w []byte) []byte {
	return w
}

func (item *MyZero) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x8d868379); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *MyZero) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *MyZero) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x8d868379)
	return item.Write(w)
}

func (item *MyZero) String() string {
	return string(item.WriteJSON(nil))
}

func (item *MyZero) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			return ErrorInvalidJSON("myZero", "this object can't have properties")
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *MyZero) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *MyZero) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *MyZero) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	return append(w, '}')
}

func (item *MyZero) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *MyZero) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("myZero", err.Error())
	}
	return nil
}
