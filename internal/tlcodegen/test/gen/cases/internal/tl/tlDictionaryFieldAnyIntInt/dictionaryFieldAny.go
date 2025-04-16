// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlDictionaryFieldAnyIntInt

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type DictionaryFieldAnyIntInt struct {
	Key   int32
	Value int32
}

func (DictionaryFieldAnyIntInt) TLName() string { return "dictionaryFieldAny" }
func (DictionaryFieldAnyIntInt) TLTag() uint32  { return 0x2c43a65b }

func (item *DictionaryFieldAnyIntInt) Reset() {
	item.Key = 0
	item.Value = 0
}

func (item *DictionaryFieldAnyIntInt) FillRandom(rg *basictl.RandGenerator) {
	item.Key = basictl.RandomInt(rg)
	item.Value = basictl.RandomInt(rg)
}

func (item *DictionaryFieldAnyIntInt) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.IntRead(w, &item.Key); err != nil {
		return w, err
	}
	return basictl.IntRead(w, &item.Value)
}

// This method is general version of Write, use it instead!
func (item *DictionaryFieldAnyIntInt) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *DictionaryFieldAnyIntInt) Write(w []byte) []byte {
	w = basictl.IntWrite(w, item.Key)
	w = basictl.IntWrite(w, item.Value)
	return w
}

func (item *DictionaryFieldAnyIntInt) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x2c43a65b); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *DictionaryFieldAnyIntInt) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *DictionaryFieldAnyIntInt) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x2c43a65b)
	return item.Write(w)
}

func (item *DictionaryFieldAnyIntInt) String() string {
	return string(item.WriteJSON(nil))
}

func (item *DictionaryFieldAnyIntInt) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
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
					return internal.ErrorInvalidJSONWithDuplicatingKeys("dictionaryFieldAny", "key")
				}
				if err := internal.Json2ReadInt32(in, &item.Key); err != nil {
					return err
				}
				propKeyPresented = true
			case "value":
				if propValuePresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("dictionaryFieldAny", "value")
				}
				if err := internal.Json2ReadInt32(in, &item.Value); err != nil {
					return err
				}
				propValuePresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("dictionaryFieldAny", key)
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
		item.Value = 0
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *DictionaryFieldAnyIntInt) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *DictionaryFieldAnyIntInt) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *DictionaryFieldAnyIntInt) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexKey := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"key":`...)
	w = basictl.JSONWriteInt32(w, item.Key)
	if (item.Key != 0) == false {
		w = w[:backupIndexKey]
	}
	backupIndexValue := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"value":`...)
	w = basictl.JSONWriteInt32(w, item.Value)
	if (item.Value != 0) == false {
		w = w[:backupIndexValue]
	}
	return append(w, '}')
}

func (item *DictionaryFieldAnyIntInt) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *DictionaryFieldAnyIntInt) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("dictionaryFieldAny", err.Error())
	}
	return nil
}
