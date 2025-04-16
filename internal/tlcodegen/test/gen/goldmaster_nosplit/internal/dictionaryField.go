// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package internal

import (
	"sort"

	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite

func BuiltinVectorDictionaryFieldStringReset(m map[string]string) {
	for k := range m {
		delete(m, k)
	}
}

func BuiltinVectorDictionaryFieldStringFillRandom(rg *basictl.RandGenerator, m *map[string]string) {
	rg.IncreaseDepth()
	l := rg.LimitValue(basictl.RandomUint(rg))
	*m = make(map[string]string, l)
	for i := 0; i < int(l); i++ {
		var elem DictionaryFieldString
		elem.FillRandom(rg)
		(*m)[elem.Key] = elem.Value
	}
	rg.DecreaseDepth()
}
func BuiltinVectorDictionaryFieldStringRead(w []byte, m *map[string]string) (_ []byte, err error) {
	var l uint32
	if w, err = basictl.NatRead(w, &l); err != nil {
		return w, err
	}
	if err = basictl.CheckLengthSanity(w, l, 4); err != nil {
		return w, err
	}
	var data map[string]string
	if *m == nil {
		if l == 0 {
			return w, nil
		}
		data = make(map[string]string, l)
		*m = data
	} else {
		data = *m
		for k := range data {
			delete(data, k)
		}
	}
	for i := 0; i < int(l); i++ {
		var elem DictionaryFieldString
		if w, err = elem.Read(w); err != nil {
			return w, err
		}
		data[elem.Key] = elem.Value
	}
	return w, nil
}

func BuiltinVectorDictionaryFieldStringWrite(w []byte, m map[string]string) []byte {
	w = basictl.NatWrite(w, uint32(len(m)))
	if len(m) == 0 {
		return w
	}
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		val := m[key]
		elem := DictionaryFieldString{Key: key, Value: val}
		w = elem.Write(w)
	}
	return w
}

func BuiltinVectorDictionaryFieldStringReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, m *map[string]string) error {
	var data map[string]string
	if *m == nil {
		*m = make(map[string]string, 0)
		data = *m
	} else {
		data = *m
		for k := range data {
			delete(data, k)
		}
	}
	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return ErrorInvalidJSON("map[string]string", "expected json object")
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			var value string
			if err := Json2ReadString(in, &value); err != nil {
				return err
			}
			data[key] = value
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return ErrorInvalidJSON("map[string]string", "expected json object's end")
		}
	}
	return nil
}

func BuiltinVectorDictionaryFieldStringWriteJSON(w []byte, m map[string]string) []byte {
	return BuiltinVectorDictionaryFieldStringWriteJSONOpt(true, false, w, m)
}
func BuiltinVectorDictionaryFieldStringWriteJSONOpt(newTypeNames bool, short bool, w []byte, m map[string]string) []byte {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	w = append(w, '{')
	for _, key := range keys {
		value := m[key]
		w = basictl.JSONAddCommaIfNeeded(w)
		w = basictl.JSONWriteString(w, key)
		w = append(w, ':')
		w = basictl.JSONWriteString(w, value)
	}
	return append(w, '}')
}

func BuiltinVectorDictionaryFieldStringBytesFillRandom(rg *basictl.RandGenerator, vec *[]DictionaryFieldStringBytes) {
	rg.IncreaseDepth()
	l := rg.LimitValue(basictl.RandomUint(rg))
	*vec = make([]DictionaryFieldStringBytes, l)
	for i := range *vec {
		(*vec)[i].FillRandom(rg)
	}
	rg.DecreaseDepth()
}

func BuiltinVectorDictionaryFieldStringBytesRead(w []byte, vec *[]DictionaryFieldStringBytes) (_ []byte, err error) {
	var l uint32
	if w, err = basictl.NatRead(w, &l); err != nil {
		return w, err
	}
	if err = basictl.CheckLengthSanity(w, l, 4); err != nil {
		return w, err
	}
	if uint32(cap(*vec)) < l {
		*vec = make([]DictionaryFieldStringBytes, l)
	} else {
		*vec = (*vec)[:l]
	}
	for i := range *vec {
		if w, err = (*vec)[i].Read(w); err != nil {
			return w, err
		}
	}
	return w, nil
}

func BuiltinVectorDictionaryFieldStringBytesWrite(w []byte, vec []DictionaryFieldStringBytes) []byte {
	w = basictl.NatWrite(w, uint32(len(vec)))
	for _, elem := range vec {
		w = elem.Write(w)
	}
	return w
}

func BuiltinVectorDictionaryFieldStringBytesReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, vec *[]DictionaryFieldStringBytes) error {
	*vec = (*vec)[:cap(*vec)]
	index := 0
	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return ErrorInvalidJSON("[]DictionaryFieldStringBytes", "expected json object")
		}
		for ; !in.IsDelim('}'); index++ {
			if len(*vec) <= index {
				var newValue DictionaryFieldStringBytes
				*vec = append(*vec, newValue)
				*vec = (*vec)[:cap(*vec)]
			}
			(*vec)[index].Key = append((*vec)[index].Key[:0], in.UnsafeFieldName(true)...)
			in.WantColon()
			if err := Json2ReadStringBytes(in, &(*vec)[index].Value); err != nil {
				return err
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return ErrorInvalidJSON("[]DictionaryFieldStringBytes", "expected json object's end")
		}
	}
	*vec = (*vec)[:index]
	return nil
}

func BuiltinVectorDictionaryFieldStringBytesWriteJSON(w []byte, vec []DictionaryFieldStringBytes) []byte {
	return BuiltinVectorDictionaryFieldStringBytesWriteJSONOpt(true, false, w, vec)
}
func BuiltinVectorDictionaryFieldStringBytesWriteJSONOpt(newTypeNames bool, short bool, w []byte, vec []DictionaryFieldStringBytes) []byte {
	w = append(w, '{')
	for _, elem := range vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = basictl.JSONWriteStringBytes(w, elem.Key)
		w = append(w, ':')
		w = basictl.JSONWriteStringBytes(w, elem.Value)
	}
	return append(w, '}')
}

func BuiltinVectorDictionaryFieldUsefulServiceUserEntityPaymentItemBoxedReset(m map[string]UsefulServiceUserEntityPaymentItem) {
	for k := range m {
		delete(m, k)
	}
}

func BuiltinVectorDictionaryFieldUsefulServiceUserEntityPaymentItemBoxedFillRandom(rg *basictl.RandGenerator, m *map[string]UsefulServiceUserEntityPaymentItem, nat_t uint32) {
	rg.IncreaseDepth()
	l := rg.LimitValue(basictl.RandomUint(rg))
	*m = make(map[string]UsefulServiceUserEntityPaymentItem, l)
	for i := 0; i < int(l); i++ {
		var elem DictionaryFieldUsefulServiceUserEntityPaymentItemBoxed
		elem.FillRandom(rg, nat_t)
		(*m)[elem.Key] = elem.Value
	}
	rg.DecreaseDepth()
}
func BuiltinVectorDictionaryFieldUsefulServiceUserEntityPaymentItemBoxedRead(w []byte, m *map[string]UsefulServiceUserEntityPaymentItem, nat_t uint32) (_ []byte, err error) {
	var l uint32
	if w, err = basictl.NatRead(w, &l); err != nil {
		return w, err
	}
	if err = basictl.CheckLengthSanity(w, l, 4); err != nil {
		return w, err
	}
	var data map[string]UsefulServiceUserEntityPaymentItem
	if *m == nil {
		if l == 0 {
			return w, nil
		}
		data = make(map[string]UsefulServiceUserEntityPaymentItem, l)
		*m = data
	} else {
		data = *m
		for k := range data {
			delete(data, k)
		}
	}
	for i := 0; i < int(l); i++ {
		var elem DictionaryFieldUsefulServiceUserEntityPaymentItemBoxed
		if w, err = elem.Read(w, nat_t); err != nil {
			return w, err
		}
		data[elem.Key] = elem.Value
	}
	return w, nil
}

func BuiltinVectorDictionaryFieldUsefulServiceUserEntityPaymentItemBoxedWrite(w []byte, m map[string]UsefulServiceUserEntityPaymentItem, nat_t uint32) []byte {
	w = basictl.NatWrite(w, uint32(len(m)))
	if len(m) == 0 {
		return w
	}
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		val := m[key]
		elem := DictionaryFieldUsefulServiceUserEntityPaymentItemBoxed{Key: key, Value: val}
		w = elem.Write(w, nat_t)
	}
	return w
}

func BuiltinVectorDictionaryFieldUsefulServiceUserEntityPaymentItemBoxedReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, m *map[string]UsefulServiceUserEntityPaymentItem, nat_t uint32) error {
	var data map[string]UsefulServiceUserEntityPaymentItem
	if *m == nil {
		*m = make(map[string]UsefulServiceUserEntityPaymentItem, 0)
		data = *m
	} else {
		data = *m
		for k := range data {
			delete(data, k)
		}
	}
	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return ErrorInvalidJSON("map[string]UsefulServiceUserEntityPaymentItem", "expected json object")
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			var value UsefulServiceUserEntityPaymentItem
			if err := value.ReadJSON(legacyTypeNames, in, nat_t); err != nil {
				return err
			}
			data[key] = value
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return ErrorInvalidJSON("map[string]UsefulServiceUserEntityPaymentItem", "expected json object's end")
		}
	}
	return nil
}

func BuiltinVectorDictionaryFieldUsefulServiceUserEntityPaymentItemBoxedWriteJSON(w []byte, m map[string]UsefulServiceUserEntityPaymentItem, nat_t uint32) []byte {
	return BuiltinVectorDictionaryFieldUsefulServiceUserEntityPaymentItemBoxedWriteJSONOpt(true, false, w, m, nat_t)
}
func BuiltinVectorDictionaryFieldUsefulServiceUserEntityPaymentItemBoxedWriteJSONOpt(newTypeNames bool, short bool, w []byte, m map[string]UsefulServiceUserEntityPaymentItem, nat_t uint32) []byte {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	w = append(w, '{')
	for _, key := range keys {
		value := m[key]
		w = basictl.JSONAddCommaIfNeeded(w)
		w = basictl.JSONWriteString(w, key)
		w = append(w, ':')
		w = value.WriteJSONOpt(newTypeNames, short, w, nat_t)
	}
	return append(w, '}')
}

type DictionaryFieldString struct {
	Key   string
	Value string
}

func (DictionaryFieldString) TLName() string { return "dictionaryField" }
func (DictionaryFieldString) TLTag() uint32  { return 0x239c1b62 }

func (item *DictionaryFieldString) Reset() {
	item.Key = ""
	item.Value = ""
}

func (item *DictionaryFieldString) FillRandom(rg *basictl.RandGenerator) {
	item.Key = basictl.RandomString(rg)
	item.Value = basictl.RandomString(rg)
}

func (item *DictionaryFieldString) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.StringRead(w, &item.Key); err != nil {
		return w, err
	}
	return basictl.StringRead(w, &item.Value)
}

// This method is general version of Write, use it instead!
func (item *DictionaryFieldString) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *DictionaryFieldString) Write(w []byte) []byte {
	w = basictl.StringWrite(w, item.Key)
	w = basictl.StringWrite(w, item.Value)
	return w
}

func (item *DictionaryFieldString) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x239c1b62); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *DictionaryFieldString) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *DictionaryFieldString) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x239c1b62)
	return item.Write(w)
}

func (item *DictionaryFieldString) String() string {
	return string(item.WriteJSON(nil))
}

func (item *DictionaryFieldString) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
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
					return ErrorInvalidJSONWithDuplicatingKeys("dictionaryField", "key")
				}
				if err := Json2ReadString(in, &item.Key); err != nil {
					return err
				}
				propKeyPresented = true
			case "value":
				if propValuePresented {
					return ErrorInvalidJSONWithDuplicatingKeys("dictionaryField", "value")
				}
				if err := Json2ReadString(in, &item.Value); err != nil {
					return err
				}
				propValuePresented = true
			default:
				return ErrorInvalidJSONExcessElement("dictionaryField", key)
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
		item.Value = ""
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *DictionaryFieldString) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *DictionaryFieldString) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *DictionaryFieldString) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexKey := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"key":`...)
	w = basictl.JSONWriteString(w, item.Key)
	if (len(item.Key) != 0) == false {
		w = w[:backupIndexKey]
	}
	backupIndexValue := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"value":`...)
	w = basictl.JSONWriteString(w, item.Value)
	if (len(item.Value) != 0) == false {
		w = w[:backupIndexValue]
	}
	return append(w, '}')
}

func (item *DictionaryFieldString) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *DictionaryFieldString) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("dictionaryField", err.Error())
	}
	return nil
}

type DictionaryFieldStringBytes struct {
	Key   []byte
	Value []byte
}

func (DictionaryFieldStringBytes) TLName() string { return "dictionaryField" }
func (DictionaryFieldStringBytes) TLTag() uint32  { return 0x239c1b62 }

func (item *DictionaryFieldStringBytes) Reset() {
	item.Key = item.Key[:0]
	item.Value = item.Value[:0]
}

func (item *DictionaryFieldStringBytes) FillRandom(rg *basictl.RandGenerator) {
	item.Key = basictl.RandomStringBytes(rg)
	item.Value = basictl.RandomStringBytes(rg)
}

func (item *DictionaryFieldStringBytes) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.StringReadBytes(w, &item.Key); err != nil {
		return w, err
	}
	return basictl.StringReadBytes(w, &item.Value)
}

// This method is general version of Write, use it instead!
func (item *DictionaryFieldStringBytes) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *DictionaryFieldStringBytes) Write(w []byte) []byte {
	w = basictl.StringWriteBytes(w, item.Key)
	w = basictl.StringWriteBytes(w, item.Value)
	return w
}

func (item *DictionaryFieldStringBytes) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x239c1b62); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *DictionaryFieldStringBytes) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *DictionaryFieldStringBytes) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x239c1b62)
	return item.Write(w)
}

func (item *DictionaryFieldStringBytes) String() string {
	return string(item.WriteJSON(nil))
}

func (item *DictionaryFieldStringBytes) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
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
					return ErrorInvalidJSONWithDuplicatingKeys("dictionaryField", "key")
				}
				if err := Json2ReadStringBytes(in, &item.Key); err != nil {
					return err
				}
				propKeyPresented = true
			case "value":
				if propValuePresented {
					return ErrorInvalidJSONWithDuplicatingKeys("dictionaryField", "value")
				}
				if err := Json2ReadStringBytes(in, &item.Value); err != nil {
					return err
				}
				propValuePresented = true
			default:
				return ErrorInvalidJSONExcessElement("dictionaryField", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propKeyPresented {
		item.Key = item.Key[:0]
	}
	if !propValuePresented {
		item.Value = item.Value[:0]
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *DictionaryFieldStringBytes) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *DictionaryFieldStringBytes) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *DictionaryFieldStringBytes) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexKey := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"key":`...)
	w = basictl.JSONWriteStringBytes(w, item.Key)
	if (len(item.Key) != 0) == false {
		w = w[:backupIndexKey]
	}
	backupIndexValue := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"value":`...)
	w = basictl.JSONWriteStringBytes(w, item.Value)
	if (len(item.Value) != 0) == false {
		w = w[:backupIndexValue]
	}
	return append(w, '}')
}

func (item *DictionaryFieldStringBytes) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *DictionaryFieldStringBytes) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("dictionaryField", err.Error())
	}
	return nil
}

type DictionaryFieldUsefulServiceUserEntityPaymentItemBoxed struct {
	Key   string
	Value UsefulServiceUserEntityPaymentItem
}

func (DictionaryFieldUsefulServiceUserEntityPaymentItemBoxed) TLName() string {
	return "dictionaryField"
}
func (DictionaryFieldUsefulServiceUserEntityPaymentItemBoxed) TLTag() uint32 { return 0x239c1b62 }

func (item *DictionaryFieldUsefulServiceUserEntityPaymentItemBoxed) Reset() {
	item.Key = ""
	item.Value.Reset()
}

func (item *DictionaryFieldUsefulServiceUserEntityPaymentItemBoxed) FillRandom(rg *basictl.RandGenerator, nat_t uint32) {
	item.Key = basictl.RandomString(rg)
	item.Value.FillRandom(rg, nat_t)
}

func (item *DictionaryFieldUsefulServiceUserEntityPaymentItemBoxed) Read(w []byte, nat_t uint32) (_ []byte, err error) {
	if w, err = basictl.StringRead(w, &item.Key); err != nil {
		return w, err
	}
	return item.Value.ReadBoxed(w, nat_t)
}

// This method is general version of Write, use it instead!
func (item *DictionaryFieldUsefulServiceUserEntityPaymentItemBoxed) WriteGeneral(w []byte, nat_t uint32) (_ []byte, err error) {
	return item.Write(w, nat_t), nil
}

func (item *DictionaryFieldUsefulServiceUserEntityPaymentItemBoxed) Write(w []byte, nat_t uint32) []byte {
	w = basictl.StringWrite(w, item.Key)
	w = item.Value.WriteBoxed(w, nat_t)
	return w
}

func (item *DictionaryFieldUsefulServiceUserEntityPaymentItemBoxed) ReadBoxed(w []byte, nat_t uint32) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x239c1b62); err != nil {
		return w, err
	}
	return item.Read(w, nat_t)
}

// This method is general version of WriteBoxed, use it instead!
func (item *DictionaryFieldUsefulServiceUserEntityPaymentItemBoxed) WriteBoxedGeneral(w []byte, nat_t uint32) (_ []byte, err error) {
	return item.WriteBoxed(w, nat_t), nil
}

func (item *DictionaryFieldUsefulServiceUserEntityPaymentItemBoxed) WriteBoxed(w []byte, nat_t uint32) []byte {
	w = basictl.NatWrite(w, 0x239c1b62)
	return item.Write(w, nat_t)
}

func (item *DictionaryFieldUsefulServiceUserEntityPaymentItemBoxed) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, nat_t uint32) error {
	var propKeyPresented bool
	var rawValue []byte

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
					return ErrorInvalidJSONWithDuplicatingKeys("dictionaryField", "key")
				}
				if err := Json2ReadString(in, &item.Key); err != nil {
					return err
				}
				propKeyPresented = true
			case "value":
				if rawValue != nil {
					return ErrorInvalidJSONWithDuplicatingKeys("dictionaryField", "value")
				}
				rawValue = in.Raw()
				if !in.Ok() {
					return in.Error()
				}
			default:
				return ErrorInvalidJSONExcessElement("dictionaryField", key)
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
	var inValuePointer *basictl.JsonLexer
	inValue := basictl.JsonLexer{Data: rawValue}
	if rawValue != nil {
		inValuePointer = &inValue
	}
	if err := item.Value.ReadJSON(legacyTypeNames, inValuePointer, nat_t); err != nil {
		return err
	}

	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *DictionaryFieldUsefulServiceUserEntityPaymentItemBoxed) WriteJSONGeneral(w []byte, nat_t uint32) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w, nat_t), nil
}

func (item *DictionaryFieldUsefulServiceUserEntityPaymentItemBoxed) WriteJSON(w []byte, nat_t uint32) []byte {
	return item.WriteJSONOpt(true, false, w, nat_t)
}
func (item *DictionaryFieldUsefulServiceUserEntityPaymentItemBoxed) WriteJSONOpt(newTypeNames bool, short bool, w []byte, nat_t uint32) []byte {
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
	w = item.Value.WriteJSONOpt(newTypeNames, short, w, nat_t)
	return append(w, '}')
}
