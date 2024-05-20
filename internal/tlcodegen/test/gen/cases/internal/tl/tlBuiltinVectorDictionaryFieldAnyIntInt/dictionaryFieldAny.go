// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlBuiltinVectorDictionaryFieldAnyIntInt

import (
	"sort"

	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tl/tlDictionaryFieldAnyIntInt"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

func BuiltinVectorDictionaryFieldAnyIntIntReset(m map[int32]int32) {
	for k := range m {
		delete(m, k)
	}
}

func BuiltinVectorDictionaryFieldAnyIntIntFillRandom(rg *basictl.RandGenerator, m *map[int32]int32) {
	rg.IncreaseDepth()
	l := rg.LimitValue(basictl.RandomUint(rg))
	*m = make(map[int32]int32, l)
	for i := 0; i < int(l); i++ {
		var elem tlDictionaryFieldAnyIntInt.DictionaryFieldAnyIntInt
		elem.FillRandom(rg)
		(*m)[elem.Key] = elem.Value
	}
	rg.DecreaseDepth()
}
func BuiltinVectorDictionaryFieldAnyIntIntRead(w []byte, m *map[int32]int32) (_ []byte, err error) {
	var l uint32
	if w, err = basictl.NatRead(w, &l); err != nil {
		return w, err
	}
	if err = basictl.CheckLengthSanity(w, l, 4); err != nil {
		return w, err
	}
	var data map[int32]int32
	if *m == nil {
		if l == 0 {
			return w, nil
		}
		data = make(map[int32]int32, l)
		*m = data
	} else {
		data = *m
		for k := range data {
			delete(data, k)
		}
	}
	for i := 0; i < int(l); i++ {
		var elem tlDictionaryFieldAnyIntInt.DictionaryFieldAnyIntInt
		if w, err = elem.Read(w); err != nil {
			return w, err
		}
		data[elem.Key] = elem.Value
	}
	return w, nil
}

func BuiltinVectorDictionaryFieldAnyIntIntWrite(w []byte, m map[int32]int32) []byte {
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
		elem := tlDictionaryFieldAnyIntInt.DictionaryFieldAnyIntInt{Key: key, Value: val}
		w = elem.Write(w)
	}
	return w
}

func BuiltinVectorDictionaryFieldAnyIntIntReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, m *map[int32]int32) error {
	var data map[int32]int32
	if *m == nil {
		*m = make(map[int32]int32, 0)
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
			return internal.ErrorInvalidJSON("map[int32]int32", "expected json object")
		}
		for !in.IsDelim('}') {
			keyBytes := []byte(in.UnsafeFieldName(false))
			in.WantColon()
			if !in.Ok() {
				return internal.ErrorInvalidJSON("map[int32]int32", "expected correct json value in key")
			}
			in2 := basictl.JsonLexer{Data: keyBytes}
			var key int32
			if err := internal.Json2ReadInt32(&in2, &key); err != nil {
				return err
			}
			var value int32
			if err := internal.Json2ReadInt32(in, &value); err != nil {
				return err
			}
			data[key] = value
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return internal.ErrorInvalidJSON("map[int32]int32", "expected json object's end")
		}
	}
	return nil
}

func BuiltinVectorDictionaryFieldAnyIntIntWriteJSON(w []byte, m map[int32]int32) []byte {
	return BuiltinVectorDictionaryFieldAnyIntIntWriteJSONOpt(true, false, w, m)
}
func BuiltinVectorDictionaryFieldAnyIntIntWriteJSONOpt(newTypeNames bool, short bool, w []byte, m map[int32]int32) []byte {
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
		w = basictl.JSONWriteInt32(w, value)
	}
	return append(w, '}')
}
