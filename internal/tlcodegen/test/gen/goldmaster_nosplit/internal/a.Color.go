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

func ABlue() AColor { return AColor__MakeEnum(4) }

var _AColor = [5]UnionElement{
	{TLTag: 0xf35d7a69, TLName: "a.color", TLString: "a.color#f35d7a69"},
	{TLTag: 0xb83a723d, TLName: "a.red", TLString: "a.red#b83a723d"},
	{TLTag: 0x6127e7b8, TLName: "a.green", TLString: "a.green#6127e7b8"},
	{TLTag: 0xa9471844, TLName: "b.red", TLString: "b.red#a9471844"},
	{TLTag: 0x623360f3, TLName: "a.blue", TLString: "a.blue#623360f3"},
}

func AColor__MakeEnum(i int) AColor { return AColor{index: i} }

type AColor struct {
	index int
}

func (item AColor) TLName() string { return _AColor[item.index].TLName }
func (item AColor) TLTag() uint32  { return _AColor[item.index].TLTag }

func (item *AColor) Reset() { item.index = 0 }
func (item *AColor) FillRandom(rg *basictl.RandGenerator) {
	index := basictl.RandomUint(rg) % 5
	switch index {
	case 0:
		item.index = 0
	case 1:
		item.index = 1
	case 2:
		item.index = 2
	case 3:
		item.index = 3
	case 4:
		item.index = 4
	default:
	}
}

func (item *AColor) IsColor() bool { return item.index == 0 }
func (item *AColor) SetColor()     { item.index = 0 }

func (item *AColor) IsRed() bool { return item.index == 1 }
func (item *AColor) SetRed()     { item.index = 1 }

func (item *AColor) IsGreen() bool { return item.index == 2 }
func (item *AColor) SetGreen()     { item.index = 2 }

func (item *AColor) IsBRed() bool { return item.index == 3 }
func (item *AColor) SetBRed()     { item.index = 3 }

func (item *AColor) IsBlue() bool { return item.index == 4 }
func (item *AColor) SetBlue()     { item.index = 4 }

func (item *AColor) ReadBoxed(w []byte) (_ []byte, err error) {
	var tag uint32
	if w, err = basictl.NatRead(w, &tag); err != nil {
		return w, err
	}
	switch tag {
	case 0xf35d7a69:
		item.index = 0
		return w, nil
	case 0xb83a723d:
		item.index = 1
		return w, nil
	case 0x6127e7b8:
		item.index = 2
		return w, nil
	case 0xa9471844:
		item.index = 3
		return w, nil
	case 0x623360f3:
		item.index = 4
		return w, nil
	default:
		return w, ErrorInvalidUnionTag("a.Color", tag)
	}
}

// This method is general version of WriteBoxed, use it instead!
func (item *AColor) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item AColor) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, _AColor[item.index].TLTag)
	return w
}

func (item *AColor) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	_jtype := in.UnsafeString()
	if !in.Ok() {
		return ErrorInvalidJSON("a.Color", "expected string")
	}
	switch _jtype {
	case "a.color#f35d7a69", "a.color", "#f35d7a69":
		if !legacyTypeNames && _jtype == "a.color#f35d7a69" {
			return ErrorInvalidUnionLegacyTagJSON("a.Color", "a.color#f35d7a69")
		}
		item.index = 0
		return nil
	case "a.red#b83a723d", "a.red", "#b83a723d":
		if !legacyTypeNames && _jtype == "a.red#b83a723d" {
			return ErrorInvalidUnionLegacyTagJSON("a.Color", "a.red#b83a723d")
		}
		item.index = 1
		return nil
	case "a.green#6127e7b8", "a.green", "#6127e7b8":
		if !legacyTypeNames && _jtype == "a.green#6127e7b8" {
			return ErrorInvalidUnionLegacyTagJSON("a.Color", "a.green#6127e7b8")
		}
		item.index = 2
		return nil
	case "b.red#a9471844", "b.red", "#a9471844":
		if !legacyTypeNames && _jtype == "b.red#a9471844" {
			return ErrorInvalidUnionLegacyTagJSON("a.Color", "b.red#a9471844")
		}
		item.index = 3
		return nil
	case "a.blue#623360f3", "a.blue", "#623360f3":
		if !legacyTypeNames && _jtype == "a.blue#623360f3" {
			return ErrorInvalidUnionLegacyTagJSON("a.Color", "a.blue#623360f3")
		}
		item.index = 4
		return nil
	default:
		return ErrorInvalidEnumTagJSON("a.Color", _jtype)
	}
}

// This method is general version of WriteJSON, use it instead!
func (item AColor) WriteJSONGeneral(w []byte) ([]byte, error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item AColor) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item AColor) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '"')
	if newTypeNames {
		w = append(w, _AColor[item.index].TLName...)
	} else {
		w = append(w, _AColor[item.index].TLString...)
	}
	return append(w, '"')
}

func (item *AColor) String() string {
	return string(item.WriteJSON(nil))
}

func (item *AColor) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *AColor) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("a.Color", err.Error())
	}
	return nil
}

func AColor0() AColor { return AColor__MakeEnum(0) }

type AColorBoxedMaybe struct {
	Value AColor // not deterministic if !Ok
	Ok    bool
}

func (item *AColorBoxedMaybe) Reset() {
	item.Ok = false
}
func (item *AColorBoxedMaybe) FillRandom(rg *basictl.RandGenerator) {
	if basictl.RandomUint(rg)&1 == 1 {
		item.Ok = true
		item.Value.FillRandom(rg)
	} else {
		item.Ok = false
	}
}

func (item *AColorBoxedMaybe) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.ReadBool(w, &item.Ok, 0x27930a7b, 0x3f9c8ef8); err != nil {
		return w, err
	}
	if item.Ok {
		return item.Value.ReadBoxed(w)
	}
	return w, nil
}

// This method is general version of WriteBoxed, use it instead!
func (item *AColorBoxedMaybe) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *AColorBoxedMaybe) WriteBoxed(w []byte) []byte {
	if item.Ok {
		w = basictl.NatWrite(w, 0x3f9c8ef8)
		return item.Value.WriteBoxed(w)
	}
	return basictl.NatWrite(w, 0x27930a7b)
}

func (item *AColorBoxedMaybe) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	_ok, _jvalue, err := Json2ReadMaybe("Maybe", in)
	if err != nil {
		return err
	}
	item.Ok = _ok
	if _ok {
		var in2Pointer *basictl.JsonLexer
		if _jvalue != nil {
			in2 := basictl.JsonLexer{Data: _jvalue}
			in2Pointer = &in2
		}
		if err := item.Value.ReadJSON(legacyTypeNames, in2Pointer); err != nil {
			return err
		}
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *AColorBoxedMaybe) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *AColorBoxedMaybe) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *AColorBoxedMaybe) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	if !item.Ok {
		return append(w, "{}"...)
	}
	w = append(w, `{"ok":true`...)
	w = append(w, `,"value":`...)
	w = item.Value.WriteJSONOpt(newTypeNames, short, w)
	return append(w, '}')
}

func (item *AColorBoxedMaybe) String() string {
	return string(item.WriteJSON(nil))
}

func AGreen() AColor { return AColor__MakeEnum(2) }

func ARed() AColor { return AColor__MakeEnum(1) }

func BRed() AColor { return AColor__MakeEnum(3) }

func BuiltinVectorAColorFillRandom(rg *basictl.RandGenerator, vec *[]AColor) {
	rg.IncreaseDepth()
	l := rg.LimitValue(basictl.RandomUint(rg))
	*vec = make([]AColor, l)
	for i := range *vec {
		(*vec)[i].FillRandom(rg)
	}
	rg.DecreaseDepth()
}
func BuiltinVectorAColorRead(w []byte, vec *[]AColor) (_ []byte, err error) {
	var l uint32
	if w, err = basictl.NatRead(w, &l); err != nil {
		return w, err
	}
	if err = basictl.CheckLengthSanity(w, l, 4); err != nil {
		return w, err
	}
	if uint32(cap(*vec)) < l {
		*vec = make([]AColor, l)
	} else {
		*vec = (*vec)[:l]
	}
	for i := range *vec {
		if w, err = (*vec)[i].ReadBoxed(w); err != nil {
			return w, err
		}
	}
	return w, nil
}

func BuiltinVectorAColorWrite(w []byte, vec []AColor) []byte {
	w = basictl.NatWrite(w, uint32(len(vec)))
	for _, elem := range vec {
		w = elem.WriteBoxed(w)
	}
	return w
}

func BuiltinVectorAColorReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, vec *[]AColor) error {
	*vec = (*vec)[:cap(*vec)]
	index := 0
	if in != nil {
		in.Delim('[')
		if !in.Ok() {
			return ErrorInvalidJSON("[]AColor", "expected json array")
		}
		for ; !in.IsDelim(']'); index++ {
			if len(*vec) <= index {
				var newValue AColor
				*vec = append(*vec, newValue)
				*vec = (*vec)[:cap(*vec)]
			}
			if err := (*vec)[index].ReadJSON(legacyTypeNames, in); err != nil {
				return err
			}
			in.WantComma()
		}
		in.Delim(']')
		if !in.Ok() {
			return ErrorInvalidJSON("[]AColor", "expected json array's end")
		}
	}
	*vec = (*vec)[:index]
	return nil
}

func BuiltinVectorAColorWriteJSON(w []byte, vec []AColor) []byte {
	return BuiltinVectorAColorWriteJSONOpt(true, false, w, vec)
}
func BuiltinVectorAColorWriteJSONOpt(newTypeNames bool, short bool, w []byte, vec []AColor) []byte {
	w = append(w, '[')
	for _, elem := range vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = elem.WriteJSONOpt(newTypeNames, short, w)
	}
	return append(w, ']')
}
