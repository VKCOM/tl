// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package cycle_2383fe3e154dfb1e44c2ac7759547e8a

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

func (item MyInt) AsUnion() MyValue {
	var ret MyValue
	ret.SetMyInt(item)
	return ret
}

type MyInt struct {
	Val1 int32
}

func (MyInt) TLName() string { return "myInt" }
func (MyInt) TLTag() uint32  { return 0xc12375b7 }

func (item *MyInt) Reset() {
	item.Val1 = 0
}

func (item *MyInt) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xa8509bda); err != nil {
		return w, err
	}
	return basictl.IntRead(w, &item.Val1)
}

// This method is general version of Write, use it instead!
func (item *MyInt) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *MyInt) Write(w []byte) []byte {
	w = basictl.NatWrite(w, 0xa8509bda)
	w = basictl.IntWrite(w, item.Val1)
	return w
}

func (item *MyInt) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xc12375b7); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *MyInt) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *MyInt) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0xc12375b7)
	return item.Write(w)
}

func (item MyInt) String() string {
	return string(item.WriteJSON(nil))
}

func (item *MyInt) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propVal1Presented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "val1":
				if propVal1Presented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("myInt", "val1")
				}
				if err := internal.Json2ReadInt32(in, &item.Val1); err != nil {
					return err
				}
				propVal1Presented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("myInt", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propVal1Presented {
		item.Val1 = 0
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *MyInt) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *MyInt) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *MyInt) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexVal1 := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"val1":`...)
	w = basictl.JSONWriteInt32(w, item.Val1)
	if (item.Val1 != 0) == false {
		w = w[:backupIndexVal1]
	}
	return append(w, '}')
}

func (item *MyInt) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *MyInt) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("myInt", err.Error())
	}
	return nil
}

func (item MyString) AsUnion() MyValue {
	var ret MyValue
	ret.SetMyString(item)
	return ret
}

type MyString struct {
	Val2 string
}

func (MyString) TLName() string { return "myString" }
func (MyString) TLTag() uint32  { return 0xc8bfa969 }

func (item *MyString) Reset() {
	item.Val2 = ""
}

func (item *MyString) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xb5286e24); err != nil {
		return w, err
	}
	return basictl.StringRead(w, &item.Val2)
}

// This method is general version of Write, use it instead!
func (item *MyString) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *MyString) Write(w []byte) []byte {
	w = basictl.NatWrite(w, 0xb5286e24)
	w = basictl.StringWrite(w, item.Val2)
	return w
}

func (item *MyString) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xc8bfa969); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *MyString) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *MyString) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0xc8bfa969)
	return item.Write(w)
}

func (item MyString) String() string {
	return string(item.WriteJSON(nil))
}

func (item *MyString) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propVal2Presented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "val2":
				if propVal2Presented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("myString", "val2")
				}
				if err := internal.Json2ReadString(in, &item.Val2); err != nil {
					return err
				}
				propVal2Presented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("myString", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propVal2Presented {
		item.Val2 = ""
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *MyString) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *MyString) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *MyString) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexVal2 := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"val2":`...)
	w = basictl.JSONWriteString(w, item.Val2)
	if (len(item.Val2) != 0) == false {
		w = w[:backupIndexVal2]
	}
	return append(w, '}')
}

func (item *MyString) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *MyString) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("myString", err.Error())
	}
	return nil
}

var _MyValue = [2]internal.UnionElement{
	{TLTag: 0xc12375b7, TLName: "myInt", TLString: "myInt#c12375b7"},
	{TLTag: 0xc8bfa969, TLName: "myString", TLString: "myString#c8bfa969"},
}

type MyValue struct {
	valueMyInt    MyInt
	valueMyString MyString
	index         int
}

func (item MyValue) TLName() string { return _MyValue[item.index].TLName }
func (item MyValue) TLTag() uint32  { return _MyValue[item.index].TLTag }

func (item *MyValue) Reset() { item.ResetToMyInt() }

func (item *MyValue) IsMyInt() bool { return item.index == 0 }

func (item *MyValue) AsMyInt() (*MyInt, bool) {
	if item.index != 0 {
		return nil, false
	}
	return &item.valueMyInt, true
}
func (item *MyValue) ResetToMyInt() *MyInt {
	item.index = 0
	item.valueMyInt.Reset()
	return &item.valueMyInt
}
func (item *MyValue) SetMyInt(value MyInt) {
	item.index = 0
	item.valueMyInt = value
}

func (item *MyValue) IsMyString() bool { return item.index == 1 }

func (item *MyValue) AsMyString() (*MyString, bool) {
	if item.index != 1 {
		return nil, false
	}
	return &item.valueMyString, true
}
func (item *MyValue) ResetToMyString() *MyString {
	item.index = 1
	item.valueMyString.Reset()
	return &item.valueMyString
}
func (item *MyValue) SetMyString(value MyString) {
	item.index = 1
	item.valueMyString = value
}

func (item *MyValue) ReadBoxed(w []byte) (_ []byte, err error) {
	var tag uint32
	if w, err = basictl.NatRead(w, &tag); err != nil {
		return w, err
	}
	switch tag {
	case 0xc12375b7:
		item.index = 0
		return item.valueMyInt.Read(w)
	case 0xc8bfa969:
		item.index = 1
		return item.valueMyString.Read(w)
	default:
		return w, internal.ErrorInvalidUnionTag("MyValue", tag)
	}
}

// This method is general version of WriteBoxed, use it instead!
func (item *MyValue) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *MyValue) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, _MyValue[item.index].TLTag)
	switch item.index {
	case 0:
		w = item.valueMyInt.Write(w)
	case 1:
		w = item.valueMyString.Write(w)
	}
	return w
}

func (item *MyValue) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	_tag, _value, err := internal.Json2ReadUnion("MyValue", in)
	if err != nil {
		return err
	}
	switch _tag {
	case "myInt#c12375b7", "myInt", "#c12375b7":
		if !legacyTypeNames && _tag == "myInt#c12375b7" {
			return internal.ErrorInvalidUnionLegacyTagJSON("MyValue", "myInt#c12375b7")
		}
		item.index = 0
		var in2Pointer *basictl.JsonLexer
		if _value != nil {
			in2 := basictl.JsonLexer{Data: _value}
			in2Pointer = &in2
		}
		if err := item.valueMyInt.ReadJSON(legacyTypeNames, in2Pointer); err != nil {
			return err
		}
	case "myString#c8bfa969", "myString", "#c8bfa969":
		if !legacyTypeNames && _tag == "myString#c8bfa969" {
			return internal.ErrorInvalidUnionLegacyTagJSON("MyValue", "myString#c8bfa969")
		}
		item.index = 1
		var in2Pointer *basictl.JsonLexer
		if _value != nil {
			in2 := basictl.JsonLexer{Data: _value}
			in2Pointer = &in2
		}
		if err := item.valueMyString.ReadJSON(legacyTypeNames, in2Pointer); err != nil {
			return err
		}
	default:
		return internal.ErrorInvalidUnionTagJSON("MyValue", _tag)
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *MyValue) WriteJSONGeneral(w []byte) ([]byte, error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *MyValue) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *MyValue) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	switch item.index {
	case 0:
		if newTypeNames {
			w = append(w, `{"type":"myInt"`...)
		} else {
			w = append(w, `{"type":"myInt#c12375b7"`...)
		}
		w = append(w, `,"value":`...)
		w = item.valueMyInt.WriteJSONOpt(newTypeNames, short, w)
		return append(w, '}')
	case 1:
		if newTypeNames {
			w = append(w, `{"type":"myString"`...)
		} else {
			w = append(w, `{"type":"myString#c8bfa969"`...)
		}
		w = append(w, `,"value":`...)
		w = item.valueMyString.WriteJSONOpt(newTypeNames, short, w)
		return append(w, '}')
	default: // Impossible due to panic above
		return w
	}
}

func (item MyValue) String() string {
	return string(item.WriteJSON(nil))
}

func (item *MyValue) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *MyValue) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("MyValue", err.Error())
	}
	return nil
}
