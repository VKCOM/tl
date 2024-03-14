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

func BuiltinTupleTlsCombinatorBoxedRead(w []byte, vec *[]TlsCombinator, nat_n uint32) (_ []byte, err error) {
	if err = basictl.CheckLengthSanity(w, nat_n, 4); err != nil {
		return w, err
	}
	if uint32(cap(*vec)) < nat_n {
		*vec = make([]TlsCombinator, nat_n)
	} else {
		*vec = (*vec)[:nat_n]
	}
	for i := range *vec {
		if w, err = (*vec)[i].ReadBoxed(w); err != nil {
			return w, err
		}
	}
	return w, nil
}

func BuiltinTupleTlsCombinatorBoxedWrite(w []byte, vec []TlsCombinator, nat_n uint32) (_ []byte, err error) {
	if uint32(len(vec)) != nat_n {
		return w, ErrorWrongSequenceLength("[]TlsCombinator", len(vec), nat_n)
	}
	for _, elem := range vec {
		if w, err = elem.WriteBoxed(w); err != nil {
			return w, err
		}
	}
	return w, nil
}

func BuiltinTupleTlsCombinatorBoxedReadJSON(j interface{}, vec *[]TlsCombinator, nat_n uint32) error {
	_, _arr, err := JsonReadArrayFixedSize("[]TlsCombinator", j, nat_n)
	if err != nil {
		return err
	}
	if uint32(cap(*vec)) < nat_n {
		*vec = make([]TlsCombinator, nat_n)
	} else {
		*vec = (*vec)[:nat_n]
	}
	for i := range *vec {
		if err := TlsCombinator__ReadJSON(&(*vec)[i], _arr[i]); err != nil {
			return err
		}
	}
	return nil
}

func BuiltinTupleTlsCombinatorBoxedWriteJSON(w []byte, vec []TlsCombinator, nat_n uint32) (_ []byte, err error) {
	return BuiltinTupleTlsCombinatorBoxedWriteJSONOpt(false, w, vec, nat_n)
}
func BuiltinTupleTlsCombinatorBoxedWriteJSONOpt(short bool, w []byte, vec []TlsCombinator, nat_n uint32) (_ []byte, err error) {
	if uint32(len(vec)) != nat_n {
		return w, ErrorWrongSequenceLength("[]TlsCombinator", len(vec), nat_n)
	}
	w = append(w, '[')
	for _, elem := range vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		if w, err = elem.WriteJSONOpt(short, w); err != nil {
			return w, err
		}
	}
	return append(w, ']'), nil
}

var _TlsCombinator = [2]UnionElement{
	{TLTag: 0x5c0a1ed5, TLName: "tls.combinator", TLString: "tls.combinator#5c0a1ed5"},
	{TLTag: 0xe91692d5, TLName: "tls.combinator_v4", TLString: "tls.combinator_v4#e91692d5"},
}

type TlsCombinator struct {
	valueCombinator TlsCombinator0
	valueV4         TlsCombinatorV4
	index           int
}

func (item TlsCombinator) TLName() string { return _TlsCombinator[item.index].TLName }
func (item TlsCombinator) TLTag() uint32  { return _TlsCombinator[item.index].TLTag }

func (item *TlsCombinator) Reset() { item.ResetToCombinator() }

func (item *TlsCombinator) IsCombinator() bool { return item.index == 0 }

func (item *TlsCombinator) AsCombinator() (*TlsCombinator0, bool) {
	if item.index != 0 {
		return nil, false
	}
	return &item.valueCombinator, true
}
func (item *TlsCombinator) ResetToCombinator() *TlsCombinator0 {
	item.index = 0
	item.valueCombinator.Reset()
	return &item.valueCombinator
}
func (item *TlsCombinator) SetCombinator(value TlsCombinator0) {
	item.index = 0
	item.valueCombinator = value
}

func (item *TlsCombinator) IsV4() bool { return item.index == 1 }

func (item *TlsCombinator) AsV4() (*TlsCombinatorV4, bool) {
	if item.index != 1 {
		return nil, false
	}
	return &item.valueV4, true
}
func (item *TlsCombinator) ResetToV4() *TlsCombinatorV4 {
	item.index = 1
	item.valueV4.Reset()
	return &item.valueV4
}
func (item *TlsCombinator) SetV4(value TlsCombinatorV4) {
	item.index = 1
	item.valueV4 = value
}

func (item *TlsCombinator) ReadBoxed(w []byte) (_ []byte, err error) {
	var tag uint32
	if w, err = basictl.NatRead(w, &tag); err != nil {
		return w, err
	}
	switch tag {
	case 0x5c0a1ed5:
		item.index = 0
		return item.valueCombinator.Read(w)
	case 0xe91692d5:
		item.index = 1
		return item.valueV4.Read(w)
	default:
		return w, ErrorInvalidUnionTag("tls.Combinator", tag)
	}
}

func (item *TlsCombinator) WriteBoxed(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, _TlsCombinator[item.index].TLTag)
	switch item.index {
	case 0:
		return item.valueCombinator.Write(w)
	case 1:
		return item.valueV4.Write(w)
	default: // Impossible due to panic above
		return w, nil
	}
}

func TlsCombinator__ReadJSON(item *TlsCombinator, j interface{}) error { return item.readJSON(j) }
func (item *TlsCombinator) readJSON(j interface{}) error {
	_jm, _tag, err := JsonReadUnionType("tls.Combinator", j)
	if err != nil {
		return err
	}
	jvalue := _jm["value"]
	switch _tag {
	case "tls.combinator#5c0a1ed5", "tls.combinator", "#5c0a1ed5":
		item.index = 0
		if err := TlsCombinator0__ReadJSON(&item.valueCombinator, jvalue); err != nil {
			return err
		}
		delete(_jm, "value")
	case "tls.combinator_v4#e91692d5", "tls.combinator_v4", "#e91692d5":
		item.index = 1
		if err := TlsCombinatorV4__ReadJSON(&item.valueV4, jvalue); err != nil {
			return err
		}
		delete(_jm, "value")
	default:
		return ErrorInvalidUnionTagJSON("tls.Combinator", _tag)
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("tls.Combinator", k)
	}
	return nil
}

func (item *TlsCombinator) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(false, w)
}
func (item *TlsCombinator) WriteJSONOpt(short bool, w []byte) (_ []byte, err error) {
	switch item.index {
	case 0:
		w = append(w, `{"type":"tls.combinator#5c0a1ed5","value":`...)
		if w, err = item.valueCombinator.WriteJSONOpt(short, w); err != nil {
			return w, err
		}
		return append(w, '}'), nil
	case 1:
		w = append(w, `{"type":"tls.combinator_v4#e91692d5","value":`...)
		if w, err = item.valueV4.WriteJSONOpt(short, w); err != nil {
			return w, err
		}
		return append(w, '}'), nil
	default: // Impossible due to panic above
		return w, nil
	}
}

func (item TlsCombinator) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func (item *TlsCombinator) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *TlsCombinator) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("tls.Combinator", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("tls.Combinator", err.Error())
	}
	return nil
}

func (item TlsCombinator0) AsUnion() TlsCombinator {
	var ret TlsCombinator
	ret.SetCombinator(item)
	return ret
}

type TlsCombinator0 struct {
	Name     int32
	Id       string
	TypeName int32
	Left     TlsCombinatorLeft
	Right    TlsCombinatorRight
}

func (TlsCombinator0) TLName() string { return "tls.combinator" }
func (TlsCombinator0) TLTag() uint32  { return 0x5c0a1ed5 }

func (item *TlsCombinator0) Reset() {
	item.Name = 0
	item.Id = ""
	item.TypeName = 0
	item.Left.Reset()
	item.Right.Reset()
}

func (item *TlsCombinator0) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.IntRead(w, &item.Name); err != nil {
		return w, err
	}
	if w, err = basictl.StringRead(w, &item.Id); err != nil {
		return w, err
	}
	if w, err = basictl.IntRead(w, &item.TypeName); err != nil {
		return w, err
	}
	if w, err = item.Left.ReadBoxed(w); err != nil {
		return w, err
	}
	return item.Right.ReadBoxed(w)
}

func (item *TlsCombinator0) Write(w []byte) (_ []byte, err error) {
	w = basictl.IntWrite(w, item.Name)
	if w, err = basictl.StringWrite(w, item.Id); err != nil {
		return w, err
	}
	w = basictl.IntWrite(w, item.TypeName)
	if w, err = item.Left.WriteBoxed(w); err != nil {
		return w, err
	}
	return item.Right.WriteBoxed(w)
}

func (item *TlsCombinator0) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x5c0a1ed5); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *TlsCombinator0) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x5c0a1ed5)
	return item.Write(w)
}

func (item TlsCombinator0) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func TlsCombinator0__ReadJSON(item *TlsCombinator0, j interface{}) error { return item.readJSON(j) }
func (item *TlsCombinator0) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("tls.combinator", "expected json object")
	}
	_jName := _jm["name"]
	delete(_jm, "name")
	if err := JsonReadInt32(_jName, &item.Name); err != nil {
		return err
	}
	_jId := _jm["id"]
	delete(_jm, "id")
	if err := JsonReadString(_jId, &item.Id); err != nil {
		return err
	}
	_jTypeName := _jm["type_name"]
	delete(_jm, "type_name")
	if err := JsonReadInt32(_jTypeName, &item.TypeName); err != nil {
		return err
	}
	_jLeft := _jm["left"]
	delete(_jm, "left")
	_jRight := _jm["right"]
	delete(_jm, "right")
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("tls.combinator", k)
	}
	if err := TlsCombinatorLeft__ReadJSON(&item.Left, _jLeft); err != nil {
		return err
	}
	if err := TlsCombinatorRight__ReadJSON(&item.Right, _jRight); err != nil {
		return err
	}
	return nil
}

func (item *TlsCombinator0) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(false, w)
}
func (item *TlsCombinator0) WriteJSONOpt(short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if item.Name != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"name":`...)
		w = basictl.JSONWriteInt32(w, item.Name)
	}
	if len(item.Id) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"id":`...)
		w = basictl.JSONWriteString(w, item.Id)
	}
	if item.TypeName != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"type_name":`...)
		w = basictl.JSONWriteInt32(w, item.TypeName)
	}
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"left":`...)
	if w, err = item.Left.WriteJSONOpt(short, w); err != nil {
		return w, err
	}
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"right":`...)
	if w, err = item.Right.WriteJSONOpt(short, w); err != nil {
		return w, err
	}
	return append(w, '}'), nil
}

func (item *TlsCombinator0) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *TlsCombinator0) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("tls.combinator", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("tls.combinator", err.Error())
	}
	return nil
}

func (item TlsCombinatorV4) AsUnion() TlsCombinator {
	var ret TlsCombinator
	ret.SetV4(item)
	return ret
}

type TlsCombinatorV4 struct {
	Name     int32
	Id       string
	TypeName int32
	Left     TlsCombinatorLeft
	Right    TlsCombinatorRight
	Flags    int32
}

func (TlsCombinatorV4) TLName() string { return "tls.combinator_v4" }
func (TlsCombinatorV4) TLTag() uint32  { return 0xe91692d5 }

func (item *TlsCombinatorV4) Reset() {
	item.Name = 0
	item.Id = ""
	item.TypeName = 0
	item.Left.Reset()
	item.Right.Reset()
	item.Flags = 0
}

func (item *TlsCombinatorV4) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.IntRead(w, &item.Name); err != nil {
		return w, err
	}
	if w, err = basictl.StringRead(w, &item.Id); err != nil {
		return w, err
	}
	if w, err = basictl.IntRead(w, &item.TypeName); err != nil {
		return w, err
	}
	if w, err = item.Left.ReadBoxed(w); err != nil {
		return w, err
	}
	if w, err = item.Right.ReadBoxed(w); err != nil {
		return w, err
	}
	return basictl.IntRead(w, &item.Flags)
}

func (item *TlsCombinatorV4) Write(w []byte) (_ []byte, err error) {
	w = basictl.IntWrite(w, item.Name)
	if w, err = basictl.StringWrite(w, item.Id); err != nil {
		return w, err
	}
	w = basictl.IntWrite(w, item.TypeName)
	if w, err = item.Left.WriteBoxed(w); err != nil {
		return w, err
	}
	if w, err = item.Right.WriteBoxed(w); err != nil {
		return w, err
	}
	return basictl.IntWrite(w, item.Flags), nil
}

func (item *TlsCombinatorV4) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xe91692d5); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *TlsCombinatorV4) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0xe91692d5)
	return item.Write(w)
}

func (item TlsCombinatorV4) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func TlsCombinatorV4__ReadJSON(item *TlsCombinatorV4, j interface{}) error { return item.readJSON(j) }
func (item *TlsCombinatorV4) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("tls.combinator_v4", "expected json object")
	}
	_jName := _jm["name"]
	delete(_jm, "name")
	if err := JsonReadInt32(_jName, &item.Name); err != nil {
		return err
	}
	_jId := _jm["id"]
	delete(_jm, "id")
	if err := JsonReadString(_jId, &item.Id); err != nil {
		return err
	}
	_jTypeName := _jm["type_name"]
	delete(_jm, "type_name")
	if err := JsonReadInt32(_jTypeName, &item.TypeName); err != nil {
		return err
	}
	_jLeft := _jm["left"]
	delete(_jm, "left")
	_jRight := _jm["right"]
	delete(_jm, "right")
	_jFlags := _jm["flags"]
	delete(_jm, "flags")
	if err := JsonReadInt32(_jFlags, &item.Flags); err != nil {
		return err
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("tls.combinator_v4", k)
	}
	if err := TlsCombinatorLeft__ReadJSON(&item.Left, _jLeft); err != nil {
		return err
	}
	if err := TlsCombinatorRight__ReadJSON(&item.Right, _jRight); err != nil {
		return err
	}
	return nil
}

func (item *TlsCombinatorV4) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(false, w)
}
func (item *TlsCombinatorV4) WriteJSONOpt(short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if item.Name != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"name":`...)
		w = basictl.JSONWriteInt32(w, item.Name)
	}
	if len(item.Id) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"id":`...)
		w = basictl.JSONWriteString(w, item.Id)
	}
	if item.TypeName != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"type_name":`...)
		w = basictl.JSONWriteInt32(w, item.TypeName)
	}
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"left":`...)
	if w, err = item.Left.WriteJSONOpt(short, w); err != nil {
		return w, err
	}
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"right":`...)
	if w, err = item.Right.WriteJSONOpt(short, w); err != nil {
		return w, err
	}
	if item.Flags != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"flags":`...)
		w = basictl.JSONWriteInt32(w, item.Flags)
	}
	return append(w, '}'), nil
}

func (item *TlsCombinatorV4) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *TlsCombinatorV4) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("tls.combinator_v4", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("tls.combinator_v4", err.Error())
	}
	return nil
}
