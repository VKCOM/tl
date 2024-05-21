// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlDictionaryElemStrangeString

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type DictionaryElemStrangeString struct {
	Key   uint32
	Value string // Conditional: item.Key.31
}

func (DictionaryElemStrangeString) TLName() string { return "dictionaryElemStrange" }
func (DictionaryElemStrangeString) TLTag() uint32  { return 0xe3b2385c }

func (item *DictionaryElemStrangeString) SetValue(v string) {
	item.Value = v
	item.Key |= 1 << 31
}
func (item *DictionaryElemStrangeString) ClearValue() {
	item.Value = ""
	item.Key &^= 1 << 31
}
func (item DictionaryElemStrangeString) IsSetValue() bool { return item.Key&(1<<31) != 0 }

func (item *DictionaryElemStrangeString) Reset() {
	item.Key = 0
	item.Value = ""
}

func (item *DictionaryElemStrangeString) FillRandom(rg *basictl.RandGenerator) {
	var maskKey uint32
	maskKey = basictl.RandomUint(rg)
	item.Key = 0
	if maskKey&(1<<0) != 0 {
		item.Key |= (1 << 31)
	}
	if item.Key&(1<<31) != 0 {
		item.Value = basictl.RandomString(rg)
	} else {
		item.Value = ""
	}
}

func (item *DictionaryElemStrangeString) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.Key); err != nil {
		return w, err
	}
	if item.Key&(1<<31) != 0 {
		if w, err = basictl.StringRead(w, &item.Value); err != nil {
			return w, err
		}
	} else {
		item.Value = ""
	}
	return w, nil
}

// This method is general version of Write, use it instead!
func (item *DictionaryElemStrangeString) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *DictionaryElemStrangeString) Write(w []byte) []byte {
	w = basictl.NatWrite(w, item.Key)
	if item.Key&(1<<31) != 0 {
		w = basictl.StringWrite(w, item.Value)
	}
	return w
}

func (item *DictionaryElemStrangeString) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xe3b2385c); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *DictionaryElemStrangeString) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *DictionaryElemStrangeString) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0xe3b2385c)
	return item.Write(w)
}

func (item DictionaryElemStrangeString) String() string {
	return string(item.WriteJSON(nil))
}

func (item *DictionaryElemStrangeString) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
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
					return internal.ErrorInvalidJSONWithDuplicatingKeys("dictionaryElemStrange", "key")
				}
				if err := internal.Json2ReadUint32(in, &item.Key); err != nil {
					return err
				}
				propKeyPresented = true
			case "value":
				if propValuePresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("dictionaryElemStrange", "value")
				}
				if err := internal.Json2ReadString(in, &item.Value); err != nil {
					return err
				}
				propValuePresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("dictionaryElemStrange", key)
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
	if propValuePresented {
		item.Key |= 1 << 31
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *DictionaryElemStrangeString) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *DictionaryElemStrangeString) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *DictionaryElemStrangeString) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexKey := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"key":`...)
	w = basictl.JSONWriteUint32(w, item.Key)
	if (item.Key != 0) == false {
		w = w[:backupIndexKey]
	}
	if item.Key&(1<<31) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"value":`...)
		w = basictl.JSONWriteString(w, item.Value)
	}
	return append(w, '}')
}

func (item *DictionaryElemStrangeString) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *DictionaryElemStrangeString) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("dictionaryElemStrange", err.Error())
	}
	return nil
}