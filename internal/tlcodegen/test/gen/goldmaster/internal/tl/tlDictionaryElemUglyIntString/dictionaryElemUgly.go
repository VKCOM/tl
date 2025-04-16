// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlDictionaryElemUglyIntString

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type DictionaryElemUglyIntString struct {
	Key   int32  // Conditional: nat_f.0
	Value string // Conditional: nat_f.1
}

func (DictionaryElemUglyIntString) TLName() string { return "dictionaryElemUgly" }
func (DictionaryElemUglyIntString) TLTag() uint32  { return 0xe6790546 }

func (item *DictionaryElemUglyIntString) SetKey(v int32, nat_f *uint32) {
	item.Key = v
	if nat_f != nil {
		*nat_f |= 1 << 0
	}
}
func (item *DictionaryElemUglyIntString) ClearKey(nat_f *uint32) {
	item.Key = 0
	if nat_f != nil {
		*nat_f &^= 1 << 0
	}
}
func (item *DictionaryElemUglyIntString) IsSetKey(nat_f uint32) bool { return nat_f&(1<<0) != 0 }

func (item *DictionaryElemUglyIntString) SetValue(v string, nat_f *uint32) {
	item.Value = v
	if nat_f != nil {
		*nat_f |= 1 << 1
	}
}
func (item *DictionaryElemUglyIntString) ClearValue(nat_f *uint32) {
	item.Value = ""
	if nat_f != nil {
		*nat_f &^= 1 << 1
	}
}
func (item *DictionaryElemUglyIntString) IsSetValue(nat_f uint32) bool { return nat_f&(1<<1) != 0 }

func (item *DictionaryElemUglyIntString) Reset() {
	item.Key = 0
	item.Value = ""
}

func (item *DictionaryElemUglyIntString) FillRandom(rg *basictl.RandGenerator, nat_f uint32) {
	if nat_f&(1<<0) != 0 {
		item.Key = basictl.RandomInt(rg)
	} else {
		item.Key = 0
	}
	if nat_f&(1<<1) != 0 {
		item.Value = basictl.RandomString(rg)
	} else {
		item.Value = ""
	}
}

func (item *DictionaryElemUglyIntString) Read(w []byte, nat_f uint32) (_ []byte, err error) {
	if nat_f&(1<<0) != 0 {
		if w, err = basictl.IntRead(w, &item.Key); err != nil {
			return w, err
		}
	} else {
		item.Key = 0
	}
	if nat_f&(1<<1) != 0 {
		if w, err = basictl.StringRead(w, &item.Value); err != nil {
			return w, err
		}
	} else {
		item.Value = ""
	}
	return w, nil
}

// This method is general version of Write, use it instead!
func (item *DictionaryElemUglyIntString) WriteGeneral(w []byte, nat_f uint32) (_ []byte, err error) {
	return item.Write(w, nat_f), nil
}

func (item *DictionaryElemUglyIntString) Write(w []byte, nat_f uint32) []byte {
	if nat_f&(1<<0) != 0 {
		w = basictl.IntWrite(w, item.Key)
	}
	if nat_f&(1<<1) != 0 {
		w = basictl.StringWrite(w, item.Value)
	}
	return w
}

func (item *DictionaryElemUglyIntString) ReadBoxed(w []byte, nat_f uint32) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xe6790546); err != nil {
		return w, err
	}
	return item.Read(w, nat_f)
}

// This method is general version of WriteBoxed, use it instead!
func (item *DictionaryElemUglyIntString) WriteBoxedGeneral(w []byte, nat_f uint32) (_ []byte, err error) {
	return item.WriteBoxed(w, nat_f), nil
}

func (item *DictionaryElemUglyIntString) WriteBoxed(w []byte, nat_f uint32) []byte {
	w = basictl.NatWrite(w, 0xe6790546)
	return item.Write(w, nat_f)
}

func (item *DictionaryElemUglyIntString) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, nat_f uint32) error {
	var propKeyPresented bool
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
			case "key":
				if propKeyPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("dictionaryElemUgly", "key")
				}
				if nat_f&(1<<0) == 0 {
					return internal.ErrorInvalidJSON("dictionaryElemUgly", "field 'key' is defined, while corresponding implicit fieldmask bit is 0")
				}
				if err := internal.Json2ReadInt32(in, &item.Key); err != nil {
					return err
				}
				propKeyPresented = true
			case "value":
				if propValuePresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("dictionaryElemUgly", "value")
				}
				if nat_f&(1<<1) == 0 {
					return internal.ErrorInvalidJSON("dictionaryElemUgly", "field 'value' is defined, while corresponding implicit fieldmask bit is 0")
				}
				if err := internal.Json2ReadString(in, &item.Value); err != nil {
					return err
				}
				propValuePresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("dictionaryElemUgly", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propKeyPresented {
		item.Key = 0
	}
	if !propValuePresented {
		item.Value = ""
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *DictionaryElemUglyIntString) WriteJSONGeneral(w []byte, nat_f uint32) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w, nat_f), nil
}

func (item *DictionaryElemUglyIntString) WriteJSON(w []byte, nat_f uint32) []byte {
	return item.WriteJSONOpt(true, false, w, nat_f)
}
func (item *DictionaryElemUglyIntString) WriteJSONOpt(newTypeNames bool, short bool, w []byte, nat_f uint32) []byte {
	w = append(w, '{')
	if nat_f&(1<<0) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"key":`...)
		w = basictl.JSONWriteInt32(w, item.Key)
	}
	if nat_f&(1<<1) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"value":`...)
		w = basictl.JSONWriteString(w, item.Value)
	}
	return append(w, '}')
}
