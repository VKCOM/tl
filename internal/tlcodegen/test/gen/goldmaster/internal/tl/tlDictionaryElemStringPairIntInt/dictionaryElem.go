// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlDictionaryElemStringPairIntInt

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlPairIntInt"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type DictionaryElemStringPairIntInt struct {
	Key   string
	Value tlPairIntInt.PairIntInt
}

func (DictionaryElemStringPairIntInt) TLName() string { return "dictionaryElem" }
func (DictionaryElemStringPairIntInt) TLTag() uint32  { return 0xa69d7dd0 }

func (item *DictionaryElemStringPairIntInt) Reset() {
	item.Key = ""
	item.Value.Reset()
}

func (item *DictionaryElemStringPairIntInt) FillRandom(rg *basictl.RandGenerator) {
	item.Key = basictl.RandomString(rg)
	item.Value.FillRandom(rg)
}

func (item *DictionaryElemStringPairIntInt) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.StringRead(w, &item.Key); err != nil {
		return w, err
	}
	return item.Value.Read(w)
}

// This method is general version of Write, use it instead!
func (item *DictionaryElemStringPairIntInt) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *DictionaryElemStringPairIntInt) Write(w []byte) []byte {
	w = basictl.StringWrite(w, item.Key)
	w = item.Value.Write(w)
	return w
}

func (item *DictionaryElemStringPairIntInt) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xa69d7dd0); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *DictionaryElemStringPairIntInt) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *DictionaryElemStringPairIntInt) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0xa69d7dd0)
	return item.Write(w)
}

func (item *DictionaryElemStringPairIntInt) String() string {
	return string(item.WriteJSON(nil))
}

func (item *DictionaryElemStringPairIntInt) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
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
					return internal.ErrorInvalidJSONWithDuplicatingKeys("dictionaryElem", "key")
				}
				if err := internal.Json2ReadString(in, &item.Key); err != nil {
					return err
				}
				propKeyPresented = true
			case "value":
				if propValuePresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("dictionaryElem", "value")
				}
				if err := item.Value.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propValuePresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("dictionaryElem", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propKeyPresented {
		item.Key = ""
	}
	if !propValuePresented {
		item.Value.Reset()
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *DictionaryElemStringPairIntInt) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *DictionaryElemStringPairIntInt) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *DictionaryElemStringPairIntInt) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexKey := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"key":`...)
	w = basictl.JSONWriteString(w, item.Key)
	if (len(item.Key) != 0) == false {
		w = w[:backupIndexKey]
	}
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"value":`...)
	w = item.Value.WriteJSONOpt(newTypeNames, short, w)
	return append(w, '}')
}

func (item *DictionaryElemStringPairIntInt) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *DictionaryElemStringPairIntInt) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("dictionaryElem", err.Error())
	}
	return nil
}
