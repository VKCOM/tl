// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlBuiltinVectorDictionaryFieldInt

import (
	"sort"

	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tl/tlDictionaryFieldInt"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

func BuiltinVectorDictionaryFieldIntReset(m map[string]int32) {
	for k := range m {
		delete(m, k)
	}
}

func BuiltinVectorDictionaryFieldIntFillRandom(rg *basictl.RandGenerator, m *map[string]int32) {
	rg.IncreaseDepth()
	l := rg.LimitValue(basictl.RandomUint(rg))
	*m = make(map[string]int32, l)
	for i := 0; i < int(l); i++ {
		var elem tlDictionaryFieldInt.DictionaryFieldInt
		elem.FillRandom(rg)
		(*m)[elem.Key] = elem.Value
	}
	rg.DecreaseDepth()
}
func BuiltinVectorDictionaryFieldIntRead(w []byte, m *map[string]int32) (_ []byte, err error) {
	var l uint32
	if w, err = basictl.NatRead(w, &l); err != nil {
		return w, err
	}
	var data map[string]int32
	if *m == nil {
		if l == 0 {
			return w, nil
		}
		data = make(map[string]int32, l)
		*m = data
	} else {
		data = *m
		for k := range data {
			delete(data, k)
		}
	}
	for i := 0; i < int(l); i++ {
		var elem tlDictionaryFieldInt.DictionaryFieldInt
		if w, err = elem.Read(w); err != nil {
			return w, err
		}
		data[elem.Key] = elem.Value
	}
	return w, nil
}

func BuiltinVectorDictionaryFieldIntWrite(w []byte, m map[string]int32) []byte {
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
		elem := tlDictionaryFieldInt.DictionaryFieldInt{Key: key, Value: val}
		w = elem.Write(w)
	}
	return w
}

func BuiltinVectorDictionaryFieldIntCalculateLayout(sizes []int, m *map[string]int32) []int {
	sizePosition := len(sizes)
	sizes = append(sizes, 0)
	if len(*m) != 0 {
		sizes[sizePosition] += basictl.TL2CalculateSize(len(*m))
	}

	keys := make([]string, 0, len(*m))
	for k := range *m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for i := 0; i < len(keys); i++ {
		elem := tlDictionaryFieldInt.DictionaryFieldInt{Key: keys[i], Value: (*m)[keys[i]]}
		currentPosition := len(sizes)
		sizes = elem.CalculateLayout(sizes)
		sizes[sizePosition] += sizes[currentPosition]
		sizes[sizePosition] += basictl.TL2CalculateSize(sizes[currentPosition])
	}
	return sizes
}

func BuiltinVectorDictionaryFieldIntInternalWriteTL2(w []byte, sizes []int, m *map[string]int32) ([]byte, []int) {
	currentSize := sizes[0]
	sizes = sizes[1:]

	w = basictl.TL2WriteSize(w, currentSize)
	if len(*m) != 0 {
		w = basictl.TL2WriteSize(w, len(*m))
	}

	keys := make([]string, 0, len(*m))
	for k := range *m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for i := 0; i < len(keys); i++ {
		elem := tlDictionaryFieldInt.DictionaryFieldInt{Key: keys[i], Value: (*m)[keys[i]]}
		w, sizes = elem.InternalWriteTL2(w, sizes)
	}
	return w, sizes
}

func BuiltinVectorDictionaryFieldIntInternalReadTL2(r []byte, m *map[string]int32) (_ []byte, err error) {
	currentSize := 0
	if r, currentSize, err = basictl.TL2ParseSize(r); err != nil {
		return r, err
	}
	if len(r) < currentSize {
		return r, basictl.TL2Error("not enough data: expected %d, got %d", currentSize, len(r))
	}

	currentR := r[:currentSize]
	r = r[currentSize:]

	elementCount := 0
	if currentSize != 0 {
		if currentR, elementCount, err = basictl.TL2ParseSize(currentR); err != nil {
			return r, err
		}
	}

	if *m == nil {
		*m = make(map[string]int32)
	}

	for key := range *m {
		delete(*m, key)
	}

	data := *m

	for i := 0; i < elementCount; i++ {
		elem := tlDictionaryFieldInt.DictionaryFieldInt{}
		if currentR, err = elem.InternalReadTL2(currentR); err != nil {
			return currentR, err
		}
		data[elem.Key] = elem.Value
	}
	return r, nil
}

func BuiltinVectorDictionaryFieldIntReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, m *map[string]int32) error {
	var data map[string]int32
	if *m == nil {
		*m = make(map[string]int32, 0)
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
			return internal.ErrorInvalidJSON("map[string]int32", "expected json object")
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			var value int32
			if err := internal.Json2ReadInt32(in, &value); err != nil {
				return err
			}
			data[key] = value
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return internal.ErrorInvalidJSON("map[string]int32", "expected json object's end")
		}
	}
	return nil
}

func BuiltinVectorDictionaryFieldIntWriteJSON(w []byte, m map[string]int32) []byte {
	tctx := basictl.JSONWriteContext{}
	return BuiltinVectorDictionaryFieldIntWriteJSONOpt(&tctx, w, m)
}
func BuiltinVectorDictionaryFieldIntWriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte, m map[string]int32) []byte {
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
		w = basictl.JSONWriteInt32(w, value)
	}
	return append(w, '}')
}

func BuiltinVectorDictionaryFieldIntBytesFillRandom(rg *basictl.RandGenerator, vec *[]tlDictionaryFieldInt.DictionaryFieldIntBytes) {
	rg.IncreaseDepth()
	l := rg.LimitValue(basictl.RandomUint(rg))
	*vec = make([]tlDictionaryFieldInt.DictionaryFieldIntBytes, l)
	for i := range *vec {
		(*vec)[i].FillRandom(rg)
	}
	rg.DecreaseDepth()
}

func BuiltinVectorDictionaryFieldIntBytesRead(w []byte, vec *[]tlDictionaryFieldInt.DictionaryFieldIntBytes) (_ []byte, err error) {
	var l uint32
	if w, err = basictl.NatRead(w, &l); err != nil {
		return w, err
	}
	if uint32(cap(*vec)) < l {
		*vec = make([]tlDictionaryFieldInt.DictionaryFieldIntBytes, l)
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

func BuiltinVectorDictionaryFieldIntBytesWrite(w []byte, vec []tlDictionaryFieldInt.DictionaryFieldIntBytes) []byte {
	w = basictl.NatWrite(w, uint32(len(vec)))
	for _, elem := range vec {
		w = elem.Write(w)
	}
	return w
}

func BuiltinVectorDictionaryFieldIntBytesCalculateLayout(sizes []int, vec *[]tlDictionaryFieldInt.DictionaryFieldIntBytes) []int {
	sizePosition := len(sizes)
	sizes = append(sizes, 0)
	if len(*vec) != 0 {
		sizes[sizePosition] += basictl.TL2CalculateSize(len(*vec))
	}
	for i := 0; i < len(*vec); i++ {
		currentPosition := len(sizes)
		elem := (*vec)[i]
		sizes = elem.CalculateLayout(sizes)
		sizes[sizePosition] += sizes[currentPosition]
		sizes[sizePosition] += basictl.TL2CalculateSize(sizes[currentPosition])
	}
	return sizes
}

func BuiltinVectorDictionaryFieldIntBytesInternalWriteTL2(w []byte, sizes []int, vec *[]tlDictionaryFieldInt.DictionaryFieldIntBytes) ([]byte, []int) {
	currentSize := sizes[0]
	sizes = sizes[1:]

	w = basictl.TL2WriteSize(w, currentSize)
	if len(*vec) != 0 {
		w = basictl.TL2WriteSize(w, len(*vec))
	}

	for i := 0; i < len(*vec); i++ {
		elem := (*vec)[i]
		w, sizes = elem.InternalWriteTL2(w, sizes)
	}
	return w, sizes
}

func BuiltinVectorDictionaryFieldIntBytesInternalReadTL2(r []byte, vec *[]tlDictionaryFieldInt.DictionaryFieldIntBytes) (_ []byte, err error) {
	currentSize := 0
	if r, currentSize, err = basictl.TL2ParseSize(r); err != nil {
		return r, err
	}
	if len(r) < currentSize {
		return r, basictl.TL2Error("not enough data: expected %d, got %d", currentSize, len(r))
	}

	currentR := r[:currentSize]
	r = r[currentSize:]

	elementCount := 0
	if currentSize != 0 {
		if currentR, elementCount, err = basictl.TL2ParseSize(currentR); err != nil {
			return r, err
		}
	}

	if cap(*vec) < elementCount {
		*vec = make([]tlDictionaryFieldInt.DictionaryFieldIntBytes, elementCount)
	}
	*vec = (*vec)[:elementCount]
	for i := 0; i < elementCount; i++ {
		elem := (*vec)[i]
		if currentR, err = elem.InternalReadTL2(currentR); err != nil {
			return currentR, err
		}
	}
	return r, nil
}

func BuiltinVectorDictionaryFieldIntBytesReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, vec *[]tlDictionaryFieldInt.DictionaryFieldIntBytes) error {
	*vec = (*vec)[:cap(*vec)]
	index := 0
	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return internal.ErrorInvalidJSON("[]tlDictionaryFieldInt.DictionaryFieldIntBytes", "expected json object")
		}
		for ; !in.IsDelim('}'); index++ {
			if len(*vec) <= index {
				var newValue tlDictionaryFieldInt.DictionaryFieldIntBytes
				*vec = append(*vec, newValue)
				*vec = (*vec)[:cap(*vec)]
			}
			(*vec)[index].Key = append((*vec)[index].Key[:0], in.UnsafeFieldName(true)...)
			in.WantColon()
			if err := internal.Json2ReadInt32(in, &(*vec)[index].Value); err != nil {
				return err
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return internal.ErrorInvalidJSON("[]tlDictionaryFieldInt.DictionaryFieldIntBytes", "expected json object's end")
		}
	}
	*vec = (*vec)[:index]
	return nil
}

func BuiltinVectorDictionaryFieldIntBytesWriteJSON(w []byte, vec []tlDictionaryFieldInt.DictionaryFieldIntBytes) []byte {
	tctx := basictl.JSONWriteContext{}
	return BuiltinVectorDictionaryFieldIntBytesWriteJSONOpt(&tctx, w, vec)
}
func BuiltinVectorDictionaryFieldIntBytesWriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte, vec []tlDictionaryFieldInt.DictionaryFieldIntBytes) []byte {
	w = append(w, '{')
	for _, elem := range vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = basictl.JSONWriteStringBytes(w, elem.Key)
		w = append(w, ':')
		w = basictl.JSONWriteInt32(w, elem.Value)
	}
	return append(w, '}')
}
