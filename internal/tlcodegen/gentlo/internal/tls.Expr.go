// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package internal

import (
	"github.com/vkcom/tl/internal/tlcodegen/gentlo/basictl"
)

var _ = basictl.NatWrite

func BuiltinTupleTlsExprBoxedRead(w []byte, vec *[]TlsExprUnion, nat_n uint32) (_ []byte, err error) {
	if err = basictl.CheckLengthSanity(w, nat_n, 4); err != nil {
		return w, err
	}
	if uint32(cap(*vec)) < nat_n {
		*vec = make([]TlsExprUnion, nat_n)
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

func BuiltinTupleTlsExprBoxedWrite(w []byte, vec []TlsExprUnion, nat_n uint32) (_ []byte, err error) {
	if uint32(len(vec)) != nat_n {
		return w, ErrorWrongSequenceLength("[]TlsExprUnion", len(vec), nat_n)
	}
	for _, elem := range vec {
		if w, err = elem.WriteBoxed(w); err != nil {
			return w, err
		}
	}
	return w, nil
}

func BuiltinTupleTlsExprBoxedReadJSON(j interface{}, vec *[]TlsExprUnion, nat_n uint32) error {
	_, _arr, err := JsonReadArrayFixedSize("[]TlsExprUnion", j, nat_n)
	if err != nil {
		return err
	}
	if uint32(cap(*vec)) < nat_n {
		*vec = make([]TlsExprUnion, nat_n)
	} else {
		*vec = (*vec)[:nat_n]
	}
	for i := range *vec {
		if err := TlsExprUnion__ReadJSON(&(*vec)[i], _arr[i]); err != nil {
			return err
		}
	}
	return nil
}

func BuiltinTupleTlsExprBoxedWriteJSON(w []byte, vec []TlsExprUnion, nat_n uint32) (_ []byte, err error) {
	return BuiltinTupleTlsExprBoxedWriteJSONOpt(false, w, vec, nat_n)
}
func BuiltinTupleTlsExprBoxedWriteJSONOpt(short bool, w []byte, vec []TlsExprUnion, nat_n uint32) (_ []byte, err error) {
	if uint32(len(vec)) != nat_n {
		return w, ErrorWrongSequenceLength("[]TlsExprUnion", len(vec), nat_n)
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

func (item TlsExprNat) AsUnion() TlsExprUnion {
	var ret TlsExprUnion
	ret.SetNat(item)
	return ret
}

// AsUnion will be here
type TlsExprNat struct {
	Expr TlsNatExprUnion
}

func (TlsExprNat) TLName() string { return "tls.exprNat" }
func (TlsExprNat) TLTag() uint32  { return 0xdcb49bd8 }

func (item *TlsExprNat) Reset() {
	item.Expr.Reset()
}

func (item *TlsExprNat) Read(w []byte) (_ []byte, err error) {
	return item.Expr.ReadBoxed(w)
}

func (item *TlsExprNat) Write(w []byte) (_ []byte, err error) {
	return item.Expr.WriteBoxed(w)
}

func (item *TlsExprNat) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xdcb49bd8); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *TlsExprNat) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0xdcb49bd8)
	return item.Write(w)
}

func (item TlsExprNat) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func TlsExprNat__ReadJSON(item *TlsExprNat, j interface{}) error { return item.readJSON(j) }
func (item *TlsExprNat) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("tls.exprNat", "expected json object")
	}
	_jExpr := _jm["expr"]
	delete(_jm, "expr")
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("tls.exprNat", k)
	}
	if err := TlsNatExprUnion__ReadJSON(&item.Expr, _jExpr); err != nil {
		return err
	}
	return nil
}

func (item *TlsExprNat) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(false, w)
}
func (item *TlsExprNat) WriteJSONOpt(short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"expr":`...)
	if w, err = item.Expr.WriteJSONOpt(short, w); err != nil {
		return w, err
	}
	return append(w, '}'), nil
}

func (item *TlsExprNat) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *TlsExprNat) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("tls.exprNat", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("tls.exprNat", err.Error())
	}
	return nil
}

func (item TlsExprType) AsUnion() TlsExprUnion {
	var ret TlsExprUnion
	ret.SetType(item)
	return ret
}

// AsUnion will be here
type TlsExprType struct {
	Expr TlsTypeExprUnion
}

func (TlsExprType) TLName() string { return "tls.exprType" }
func (TlsExprType) TLTag() uint32  { return 0xecc9da78 }

func (item *TlsExprType) Reset() {
	item.Expr.Reset()
}

func (item *TlsExprType) Read(w []byte) (_ []byte, err error) {
	return item.Expr.ReadBoxed(w)
}

func (item *TlsExprType) Write(w []byte) (_ []byte, err error) {
	return item.Expr.WriteBoxed(w)
}

func (item *TlsExprType) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xecc9da78); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *TlsExprType) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0xecc9da78)
	return item.Write(w)
}

func (item TlsExprType) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func TlsExprType__ReadJSON(item *TlsExprType, j interface{}) error { return item.readJSON(j) }
func (item *TlsExprType) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("tls.exprType", "expected json object")
	}
	_jExpr := _jm["expr"]
	delete(_jm, "expr")
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("tls.exprType", k)
	}
	if err := TlsTypeExprUnion__ReadJSON(&item.Expr, _jExpr); err != nil {
		return err
	}
	return nil
}

func (item *TlsExprType) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(false, w)
}
func (item *TlsExprType) WriteJSONOpt(short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"expr":`...)
	if w, err = item.Expr.WriteJSONOpt(short, w); err != nil {
		return w, err
	}
	return append(w, '}'), nil
}

func (item *TlsExprType) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *TlsExprType) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("tls.exprType", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("tls.exprType", err.Error())
	}
	return nil
}

var _TlsExprUnion = [2]UnionElement{
	{TLTag: 0xecc9da78, TLName: "tls.exprType", TLString: "tls.exprType#ecc9da78"},
	{TLTag: 0xdcb49bd8, TLName: "tls.exprNat", TLString: "tls.exprNat#dcb49bd8"},
}

type TlsExprUnion struct {
	valueType TlsExprType
	valueNat  TlsExprNat
	index     int
}

func (item TlsExprUnion) TLName() string { return _TlsExprUnion[item.index].TLName }
func (item TlsExprUnion) TLTag() uint32  { return _TlsExprUnion[item.index].TLTag }

func (item *TlsExprUnion) Reset() { item.ResetToType() }

func (item *TlsExprUnion) IsType() bool { return item.index == 0 }

func (item *TlsExprUnion) AsType() (*TlsExprType, bool) {
	if item.index != 0 {
		return nil, false
	}
	return &item.valueType, true
}
func (item *TlsExprUnion) ResetToType() *TlsExprType {
	item.index = 0
	item.valueType.Reset()
	return &item.valueType
}
func (item *TlsExprUnion) SetType(value TlsExprType) {
	item.index = 0
	item.valueType = value
}

func (item *TlsExprUnion) IsNat() bool { return item.index == 1 }

func (item *TlsExprUnion) AsNat() (*TlsExprNat, bool) {
	if item.index != 1 {
		return nil, false
	}
	return &item.valueNat, true
}
func (item *TlsExprUnion) ResetToNat() *TlsExprNat {
	item.index = 1
	item.valueNat.Reset()
	return &item.valueNat
}
func (item *TlsExprUnion) SetNat(value TlsExprNat) {
	item.index = 1
	item.valueNat = value
}

func (item *TlsExprUnion) ReadBoxed(w []byte) (_ []byte, err error) {
	var tag uint32
	if w, err = basictl.NatRead(w, &tag); err != nil {
		return w, err
	}
	switch tag {
	case 0xecc9da78:
		item.index = 0
		return item.valueType.Read(w)
	case 0xdcb49bd8:
		item.index = 1
		return item.valueNat.Read(w)
	default:
		return w, ErrorInvalidUnionTag("tls.Expr", tag)
	}
}

func (item *TlsExprUnion) WriteBoxed(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, _TlsExprUnion[item.index].TLTag)
	switch item.index {
	case 0:
		return item.valueType.Write(w)
	case 1:
		return item.valueNat.Write(w)
	default: // Impossible due to panic above
		return w, nil
	}
}

func TlsExprUnion__ReadJSON(item *TlsExprUnion, j interface{}) error { return item.readJSON(j) }
func (item *TlsExprUnion) readJSON(j interface{}) error {
	_jm, _tag, err := JsonReadUnionType("tls.Expr", j)
	if err != nil {
		return err
	}
	jvalue := _jm["value"]
	switch _tag {
	case "tls.exprType#ecc9da78", "tls.exprType", "#ecc9da78":
		item.index = 0
		if err := TlsExprType__ReadJSON(&item.valueType, jvalue); err != nil {
			return err
		}
		delete(_jm, "value")
	case "tls.exprNat#dcb49bd8", "tls.exprNat", "#dcb49bd8":
		item.index = 1
		if err := TlsExprNat__ReadJSON(&item.valueNat, jvalue); err != nil {
			return err
		}
		delete(_jm, "value")
	default:
		return ErrorInvalidUnionTagJSON("tls.Expr", _tag)
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("tls.Expr", k)
	}
	return nil
}

func (item *TlsExprUnion) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(false, w)
}
func (item *TlsExprUnion) WriteJSONOpt(short bool, w []byte) (_ []byte, err error) {
	switch item.index {
	case 0:
		w = append(w, `{"type":"tls.exprType#ecc9da78","value":`...)
		if w, err = item.valueType.WriteJSONOpt(short, w); err != nil {
			return w, err
		}
		return append(w, '}'), nil
	case 1:
		w = append(w, `{"type":"tls.exprNat#dcb49bd8","value":`...)
		if w, err = item.valueNat.WriteJSONOpt(short, w); err != nil {
			return w, err
		}
		return append(w, '}'), nil
	default: // Impossible due to panic above
		return w, nil
	}
}

func (item TlsExprUnion) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}
