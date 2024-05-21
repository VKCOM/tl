// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlBuiltinVectorDictionaryFieldString

import (
	"sort"

	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tl/tlDictionaryFieldString"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

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
		var elem tlDictionaryFieldString.DictionaryFieldString
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
		var elem tlDictionaryFieldString.DictionaryFieldString
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
		elem := tlDictionaryFieldString.DictionaryFieldString{Key: key, Value: val}
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
			return internal.ErrorInvalidJSON("map[string]string", "expected json object")
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			var value string
			if err := internal.Json2ReadString(in, &value); err != nil {
				return err
			}
			data[key] = value
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return internal.ErrorInvalidJSON("map[string]string", "expected json object's end")
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

func BuiltinVectorDictionaryFieldStringBytesFillRandom(rg *basictl.RandGenerator, vec *[]tlDictionaryFieldString.DictionaryFieldStringBytes) {
	rg.IncreaseDepth()
	l := rg.LimitValue(basictl.RandomUint(rg))
	*vec = make([]tlDictionaryFieldString.DictionaryFieldStringBytes, l)
	for i := range *vec {
		(*vec)[i].FillRandom(rg)
	}
	rg.DecreaseDepth()
}

func BuiltinVectorDictionaryFieldStringBytesRead(w []byte, vec *[]tlDictionaryFieldString.DictionaryFieldStringBytes) (_ []byte, err error) {
	var l uint32
	if w, err = basictl.NatRead(w, &l); err != nil {
		return w, err
	}
	if err = basictl.CheckLengthSanity(w, l, 4); err != nil {
		return w, err
	}
	if uint32(cap(*vec)) < l {
		*vec = make([]tlDictionaryFieldString.DictionaryFieldStringBytes, l)
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

func BuiltinVectorDictionaryFieldStringBytesWrite(w []byte, vec []tlDictionaryFieldString.DictionaryFieldStringBytes) []byte {
	w = basictl.NatWrite(w, uint32(len(vec)))
	for _, elem := range vec {
		w = elem.Write(w)
	}
	return w
}

func BuiltinVectorDictionaryFieldStringBytesReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, vec *[]tlDictionaryFieldString.DictionaryFieldStringBytes) error {
	*vec = (*vec)[:cap(*vec)]
	index := 0
	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return internal.ErrorInvalidJSON("[]tlDictionaryFieldString.DictionaryFieldStringBytes", "expected json object")
		}
		for ; !in.IsDelim('}'); index++ {
			if len(*vec) <= index {
				var newValue tlDictionaryFieldString.DictionaryFieldStringBytes
				*vec = append(*vec, newValue)
				*vec = (*vec)[:cap(*vec)]
			}
			(*vec)[index].Key = append((*vec)[index].Key[:0], in.UnsafeFieldName(true)...)
			in.WantColon()
			if err := internal.Json2ReadStringBytes(in, &(*vec)[index].Value); err != nil {
				return err
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return internal.ErrorInvalidJSON("[]tlDictionaryFieldString.DictionaryFieldStringBytes", "expected json object's end")
		}
	}
	*vec = (*vec)[:index]
	return nil
}

func BuiltinVectorDictionaryFieldStringBytesWriteJSON(w []byte, vec []tlDictionaryFieldString.DictionaryFieldStringBytes) []byte {
	return BuiltinVectorDictionaryFieldStringBytesWriteJSONOpt(true, false, w, vec)
}
func BuiltinVectorDictionaryFieldStringBytesWriteJSONOpt(newTypeNames bool, short bool, w []byte, vec []tlDictionaryFieldString.DictionaryFieldStringBytes) []byte {
	w = append(w, '{')
	for _, elem := range vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = basictl.JSONWriteStringBytes(w, elem.Key)
		w = append(w, ':')
		w = basictl.JSONWriteStringBytes(w, elem.Value)
	}
	return append(w, '}')
}