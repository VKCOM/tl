// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlBuiltinVectorDictionaryElemIntPairIntInt

import (
	"sort"

	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlDictionaryElemIntPairIntInt"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlPairIntInt"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

func BuiltinVectorDictionaryElemIntPairIntIntReset(m map[int32]tlPairIntInt.PairIntInt) {
	for k := range m {
		delete(m, k)
	}
}

func BuiltinVectorDictionaryElemIntPairIntIntFillRandom(rg *basictl.RandGenerator, m *map[int32]tlPairIntInt.PairIntInt) {
	rg.IncreaseDepth()
	l := rg.LimitValue(basictl.RandomUint(rg))
	*m = make(map[int32]tlPairIntInt.PairIntInt, l)
	for i := 0; i < int(l); i++ {
		var elem tlDictionaryElemIntPairIntInt.DictionaryElemIntPairIntInt
		elem.FillRandom(rg)
		(*m)[elem.Key] = elem.Value
	}
	rg.DecreaseDepth()
}
func BuiltinVectorDictionaryElemIntPairIntIntRead(w []byte, m *map[int32]tlPairIntInt.PairIntInt) (_ []byte, err error) {
	var l uint32
	if w, err = basictl.NatRead(w, &l); err != nil {
		return w, err
	}
	if err = basictl.CheckLengthSanity(w, l, 4); err != nil {
		return w, err
	}
	var data map[int32]tlPairIntInt.PairIntInt
	if *m == nil {
		if l == 0 {
			return w, nil
		}
		data = make(map[int32]tlPairIntInt.PairIntInt, l)
		*m = data
	} else {
		data = *m
		for k := range data {
			delete(data, k)
		}
	}
	for i := 0; i < int(l); i++ {
		var elem tlDictionaryElemIntPairIntInt.DictionaryElemIntPairIntInt
		if w, err = elem.Read(w); err != nil {
			return w, err
		}
		data[elem.Key] = elem.Value
	}
	return w, nil
}

func BuiltinVectorDictionaryElemIntPairIntIntWrite(w []byte, m map[int32]tlPairIntInt.PairIntInt) []byte {
	w = basictl.NatWrite(w, uint32(len(m)))
	if len(m) == 0 {
		return w
	}
	keys := make([]int32, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	for _, key := range keys {
		val := m[key]
		elem := tlDictionaryElemIntPairIntInt.DictionaryElemIntPairIntInt{Key: key, Value: val}
		w = elem.Write(w)
	}
	return w
}

func BuiltinVectorDictionaryElemIntPairIntIntReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, m *map[int32]tlPairIntInt.PairIntInt) error {
	var data map[int32]tlPairIntInt.PairIntInt
	if *m == nil {
		*m = make(map[int32]tlPairIntInt.PairIntInt, 0)
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
			return internal.ErrorInvalidJSON("map[int32]tlPairIntInt.PairIntInt", "expected json object")
		}
		for !in.IsDelim('}') {
			keyBytes := []byte(in.UnsafeFieldName(false))
			in.WantColon()
			if !in.Ok() {
				return internal.ErrorInvalidJSON("map[int32]tlPairIntInt.PairIntInt", "expected correct json value in key")
			}
			in2 := basictl.JsonLexer{Data: keyBytes}
			var key int32
			if err := internal.Json2ReadInt32(&in2, &key); err != nil {
				return err
			}
			var value tlPairIntInt.PairIntInt
			if err := value.ReadJSON(legacyTypeNames, in); err != nil {
				return err
			}
			data[key] = value
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return internal.ErrorInvalidJSON("map[int32]tlPairIntInt.PairIntInt", "expected json object's end")
		}
	}
	return nil
}

func BuiltinVectorDictionaryElemIntPairIntIntWriteJSON(w []byte, m map[int32]tlPairIntInt.PairIntInt) []byte {
	return BuiltinVectorDictionaryElemIntPairIntIntWriteJSONOpt(true, false, w, m)
}
func BuiltinVectorDictionaryElemIntPairIntIntWriteJSONOpt(newTypeNames bool, short bool, w []byte, m map[int32]tlPairIntInt.PairIntInt) []byte {
	keys := make([]int32, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	w = append(w, '{')
	for _, key := range keys {
		value := m[key]
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"`...)
		w = basictl.JSONWriteInt32(w, key)
		w = append(w, `":`...)
		w = value.WriteJSONOpt(newTypeNames, short, w)
	}
	return append(w, '}')
}
