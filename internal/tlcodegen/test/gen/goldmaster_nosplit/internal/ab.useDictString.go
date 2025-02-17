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

type AbUseDictString struct {
	FieldsMask uint32
	Tags       map[string]string
}

func (AbUseDictString) TLName() string { return "ab.useDictString" }
func (AbUseDictString) TLTag() uint32  { return 0x3325d884 }

func (item *AbUseDictString) Reset() {
	item.FieldsMask = 0
	BuiltinVectorDictionaryFieldStringReset(item.Tags)
}

func (item *AbUseDictString) FillRandom(rg *basictl.RandGenerator) {
	item.FieldsMask = basictl.RandomUint(rg)
	BuiltinVectorDictionaryFieldStringFillRandom(rg, &item.Tags)
}

func (item *AbUseDictString) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.FieldsMask); err != nil {
		return w, err
	}
	return BuiltinVectorDictionaryFieldStringRead(w, &item.Tags)
}

// This method is general version of Write, use it instead!
func (item *AbUseDictString) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *AbUseDictString) Write(w []byte) []byte {
	w = basictl.NatWrite(w, item.FieldsMask)
	w = BuiltinVectorDictionaryFieldStringWrite(w, item.Tags)
	return w
}

func (item *AbUseDictString) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x3325d884); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *AbUseDictString) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *AbUseDictString) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x3325d884)
	return item.Write(w)
}

func (item AbUseDictString) String() string {
	return string(item.WriteJSON(nil))
}

func (item *AbUseDictString) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propFieldsMaskPresented bool
	var propTagsPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "fields_mask":
				if propFieldsMaskPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("ab.useDictString", "fields_mask")
				}
				if err := Json2ReadUint32(in, &item.FieldsMask); err != nil {
					return err
				}
				propFieldsMaskPresented = true
			case "tags":
				if propTagsPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("ab.useDictString", "tags")
				}
				if err := BuiltinVectorDictionaryFieldStringReadJSON(legacyTypeNames, in, &item.Tags); err != nil {
					return err
				}
				propTagsPresented = true
			default:
				return ErrorInvalidJSONExcessElement("ab.useDictString", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propFieldsMaskPresented {
		item.FieldsMask = 0
	}
	if !propTagsPresented {
		BuiltinVectorDictionaryFieldStringReset(item.Tags)
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *AbUseDictString) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *AbUseDictString) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *AbUseDictString) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexFieldsMask := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"fields_mask":`...)
	w = basictl.JSONWriteUint32(w, item.FieldsMask)
	if (item.FieldsMask != 0) == false {
		w = w[:backupIndexFieldsMask]
	}
	backupIndexTags := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"tags":`...)
	w = BuiltinVectorDictionaryFieldStringWriteJSONOpt(newTypeNames, short, w, item.Tags)
	if (len(item.Tags) != 0) == false {
		w = w[:backupIndexTags]
	}
	return append(w, '}')
}

func (item *AbUseDictString) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *AbUseDictString) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("ab.useDictString", err.Error())
	}
	return nil
}

type AbUseDictStringBytes struct {
	FieldsMask uint32
	Tags       []DictionaryFieldStringBytes
}

func (AbUseDictStringBytes) TLName() string { return "ab.useDictString" }
func (AbUseDictStringBytes) TLTag() uint32  { return 0x3325d884 }

func (item *AbUseDictStringBytes) Reset() {
	item.FieldsMask = 0
	item.Tags = item.Tags[:0]
}

func (item *AbUseDictStringBytes) FillRandom(rg *basictl.RandGenerator) {
	item.FieldsMask = basictl.RandomUint(rg)
	BuiltinVectorDictionaryFieldStringBytesFillRandom(rg, &item.Tags)
}

func (item *AbUseDictStringBytes) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.FieldsMask); err != nil {
		return w, err
	}
	return BuiltinVectorDictionaryFieldStringBytesRead(w, &item.Tags)
}

// This method is general version of Write, use it instead!
func (item *AbUseDictStringBytes) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *AbUseDictStringBytes) Write(w []byte) []byte {
	w = basictl.NatWrite(w, item.FieldsMask)
	w = BuiltinVectorDictionaryFieldStringBytesWrite(w, item.Tags)
	return w
}

func (item *AbUseDictStringBytes) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x3325d884); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *AbUseDictStringBytes) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *AbUseDictStringBytes) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x3325d884)
	return item.Write(w)
}

func (item AbUseDictStringBytes) String() string {
	return string(item.WriteJSON(nil))
}

func (item *AbUseDictStringBytes) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propFieldsMaskPresented bool
	var propTagsPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "fields_mask":
				if propFieldsMaskPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("ab.useDictString", "fields_mask")
				}
				if err := Json2ReadUint32(in, &item.FieldsMask); err != nil {
					return err
				}
				propFieldsMaskPresented = true
			case "tags":
				if propTagsPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("ab.useDictString", "tags")
				}
				if err := BuiltinVectorDictionaryFieldStringBytesReadJSON(legacyTypeNames, in, &item.Tags); err != nil {
					return err
				}
				propTagsPresented = true
			default:
				return ErrorInvalidJSONExcessElement("ab.useDictString", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propFieldsMaskPresented {
		item.FieldsMask = 0
	}
	if !propTagsPresented {
		item.Tags = item.Tags[:0]
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *AbUseDictStringBytes) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *AbUseDictStringBytes) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *AbUseDictStringBytes) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexFieldsMask := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"fields_mask":`...)
	w = basictl.JSONWriteUint32(w, item.FieldsMask)
	if (item.FieldsMask != 0) == false {
		w = w[:backupIndexFieldsMask]
	}
	backupIndexTags := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"tags":`...)
	w = BuiltinVectorDictionaryFieldStringBytesWriteJSONOpt(newTypeNames, short, w, item.Tags)
	if (len(item.Tags) != 0) == false {
		w = w[:backupIndexTags]
	}
	return append(w, '}')
}

func (item *AbUseDictStringBytes) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *AbUseDictStringBytes) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("ab.useDictString", err.Error())
	}
	return nil
}
