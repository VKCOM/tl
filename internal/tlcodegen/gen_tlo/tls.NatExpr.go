// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package gen_tlo

var _ = NatWrite

func (item TlsNatConst) AsUnion() TlsNatExprUnion {
	var ret TlsNatExprUnion
	ret.SetNatConst(item)
	return ret
}

// AsUnion will be here
type TlsNatConst struct {
	Value int32
}

func (TlsNatConst) TLName() string { return "tls.natConst" }

// TODO - replace wrong const dcb49bd8 with correct 8ce940b1
func (TlsNatConst) TLTag() uint32 { return 0xdcb49bd8 }

func (item *TlsNatConst) Reset() {
	item.Value = 0
}

func (item *TlsNatConst) Read(w []byte) (_ []byte, err error) {
	return IntRead(w, &item.Value)
}

func (item *TlsNatConst) Write(w []byte) (_ []byte, err error) {
	return IntWrite(w, item.Value), nil
}

func (item *TlsNatConst) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = NatReadExactTag(w, 0xdcb49bd8); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *TlsNatConst) WriteBoxed(w []byte) ([]byte, error) {
	w = NatWrite(w, 0xdcb49bd8)
	return item.Write(w)
}

func (item TlsNatConst) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func TlsNatConst__ReadJSON(item *TlsNatConst, j interface{}) error { return item.readJSON(j) }
func (item *TlsNatConst) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("tls.natConst", "expected json object")
	}
	_jValue := _jm["value"]
	delete(_jm, "value")
	if err := JsonReadInt32(_jValue, &item.Value); err != nil {
		return err
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("tls.natConst", k)
	}
	return nil
}

func (item *TlsNatConst) WriteJSON(w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if item.Value != 0 {
		w = JSONAddCommaIfNeeded(w)
		w = append(w, `"value":`...)
		w = JSONWriteInt32(w, item.Value)
	}
	return append(w, '}'), nil
}

func (item *TlsNatConst) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *TlsNatConst) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("tls.natConst", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("tls.natConst", err.Error())
	}
	return nil
}

var _TlsNatExprUnion = [2]UnionElement{
	{TLTag: 0xdcb49bd8, TLName: "tls.natConst", TLString: "tls.natConst#dcb49bd8"},
	{TLTag: 0x4e8a14f0, TLName: "tls.natVar", TLString: "tls.natVar#4e8a14f0"},
}

type TlsNatExprUnion struct {
	valueNatConst TlsNatConst
	valueNatVar   TlsNatVar
	index         int
}

func (item TlsNatExprUnion) TLName() string { return _TlsNatExprUnion[item.index].TLName }
func (item TlsNatExprUnion) TLTag() uint32  { return _TlsNatExprUnion[item.index].TLTag }

func (item *TlsNatExprUnion) Reset() { item.ResetToNatConst() }

func (item *TlsNatExprUnion) IsNatConst() bool { return item.index == 0 }

func (item *TlsNatExprUnion) AsNatConst() (*TlsNatConst, bool) {
	if item.index != 0 {
		return nil, false
	}
	return &item.valueNatConst, true
}
func (item *TlsNatExprUnion) ResetToNatConst() *TlsNatConst {
	item.index = 0
	item.valueNatConst.Reset()
	return &item.valueNatConst
}
func (item *TlsNatExprUnion) SetNatConst(value TlsNatConst) {
	item.index = 0
	item.valueNatConst = value
}

func (item *TlsNatExprUnion) IsNatVar() bool { return item.index == 1 }

func (item *TlsNatExprUnion) AsNatVar() (*TlsNatVar, bool) {
	if item.index != 1 {
		return nil, false
	}
	return &item.valueNatVar, true
}
func (item *TlsNatExprUnion) ResetToNatVar() *TlsNatVar {
	item.index = 1
	item.valueNatVar.Reset()
	return &item.valueNatVar
}
func (item *TlsNatExprUnion) SetNatVar(value TlsNatVar) {
	item.index = 1
	item.valueNatVar = value
}

func (item *TlsNatExprUnion) ReadBoxed(w []byte) (_ []byte, err error) {
	var tag uint32
	if w, err = NatRead(w, &tag); err != nil {
		return w, err
	}
	switch tag {
	case 0xdcb49bd8:
		item.index = 0
		return item.valueNatConst.Read(w)
	case 0x4e8a14f0:
		item.index = 1
		return item.valueNatVar.Read(w)
	default:
		return w, ErrorInvalidUnionTag("tls.NatExpr", tag)
	}
}

func (item *TlsNatExprUnion) WriteBoxed(w []byte) (_ []byte, err error) {
	w = NatWrite(w, _TlsNatExprUnion[item.index].TLTag)
	switch item.index {
	case 0:
		return item.valueNatConst.Write(w)
	case 1:
		return item.valueNatVar.Write(w)
	default: // Impossible due to panic above
		return w, nil
	}
}

func TlsNatExprUnion__ReadJSON(item *TlsNatExprUnion, j interface{}) error { return item.readJSON(j) }
func (item *TlsNatExprUnion) readJSON(j interface{}) error {
	_jm, _tag, err := JsonReadUnionType("tls.NatExpr", j)
	if err != nil {
		return err
	}
	jvalue := _jm["value"]
	switch _tag {
	case "tls.natConst#0xdcb49bd8", "tls.natConst", "#dcb49bd8":
		item.index = 0
		if err := TlsNatConst__ReadJSON(&item.valueNatConst, jvalue); err != nil {
			return err
		}
		delete(_jm, "value")
	case "tls.natVar#4e8a14f0", "tls.natVar", "#4e8a14f0":
		item.index = 1
		if err := TlsNatVar__ReadJSON(&item.valueNatVar, jvalue); err != nil {
			return err
		}
		delete(_jm, "value")
	default:
		return ErrorInvalidUnionTagJSON("tls.NatExpr", _tag)
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("tls.NatExpr", k)
	}
	return nil
}

func (item *TlsNatExprUnion) WriteJSON(w []byte) (_ []byte, err error) {
	switch item.index {
	case 0:
		w = append(w, `{"type":"tls.natConst#dcb49bd8","value":`...)
		if w, err = item.valueNatConst.WriteJSON(w); err != nil {
			return w, err
		}
		return append(w, '}'), nil
	case 1:
		w = append(w, `{"type":"tls.natVar#4e8a14f0","value":`...)
		if w, err = item.valueNatVar.WriteJSON(w); err != nil {
			return w, err
		}
		return append(w, '}'), nil
	default: // Impossible due to panic above
		return w, nil
	}
}

func (item TlsNatExprUnion) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func (item TlsNatVar) AsUnion() TlsNatExprUnion {
	var ret TlsNatExprUnion
	ret.SetNatVar(item)
	return ret
}

// AsUnion will be here
type TlsNatVar struct {
	Dif    int32
	VarNum int32
}

func (TlsNatVar) TLName() string { return "tls.natVar" }
func (TlsNatVar) TLTag() uint32  { return 0x4e8a14f0 }

func (item *TlsNatVar) Reset() {
	item.Dif = 0
	item.VarNum = 0
}

func (item *TlsNatVar) Read(w []byte) (_ []byte, err error) {
	if w, err = IntRead(w, &item.Dif); err != nil {
		return w, err
	}
	return IntRead(w, &item.VarNum)
}

func (item *TlsNatVar) Write(w []byte) (_ []byte, err error) {
	w = IntWrite(w, item.Dif)
	return IntWrite(w, item.VarNum), nil
}

func (item *TlsNatVar) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = NatReadExactTag(w, 0x4e8a14f0); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *TlsNatVar) WriteBoxed(w []byte) ([]byte, error) {
	w = NatWrite(w, 0x4e8a14f0)
	return item.Write(w)
}

func (item TlsNatVar) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func TlsNatVar__ReadJSON(item *TlsNatVar, j interface{}) error { return item.readJSON(j) }
func (item *TlsNatVar) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("tls.natVar", "expected json object")
	}
	_jDif := _jm["dif"]
	delete(_jm, "dif")
	if err := JsonReadInt32(_jDif, &item.Dif); err != nil {
		return err
	}
	_jVarNum := _jm["var_num"]
	delete(_jm, "var_num")
	if err := JsonReadInt32(_jVarNum, &item.VarNum); err != nil {
		return err
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("tls.natVar", k)
	}
	return nil
}

func (item *TlsNatVar) WriteJSON(w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if item.Dif != 0 {
		w = JSONAddCommaIfNeeded(w)
		w = append(w, `"dif":`...)
		w = JSONWriteInt32(w, item.Dif)
	}
	if item.VarNum != 0 {
		w = JSONAddCommaIfNeeded(w)
		w = append(w, `"var_num":`...)
		w = JSONWriteInt32(w, item.VarNum)
	}
	return append(w, '}'), nil
}

func (item *TlsNatVar) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *TlsNatVar) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("tls.natVar", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("tls.natVar", err.Error())
	}
	return nil
}
