// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlBuiltinTuple0IntBoxed

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

func BuiltinTuple0IntBoxedReset(vec *[0]int32) {
	for i := range *vec {
		(*vec)[i] = 0
	}
}

func BuiltinTuple0IntBoxedFillRandom(rg *basictl.RandGenerator, vec *[0]int32) {
	rg.IncreaseDepth()
	for i := range *vec {
		(*vec)[i] = basictl.RandomInt(rg)
	}
	rg.DecreaseDepth()
}

func BuiltinTuple0IntBoxedRead(w []byte, vec *[0]int32) (_ []byte, err error) {
	for i := range *vec {
		if w, err = basictl.NatReadExactTag(w, 0xa8509bda); err != nil {
			return w, err
		}
		if w, err = basictl.IntRead(w, &(*vec)[i]); err != nil {
			return w, err
		}
	}
	return w, nil
}

func BuiltinTuple0IntBoxedWrite(w []byte, vec *[0]int32) []byte {
	for _, elem := range *vec {
		w = basictl.NatWrite(w, 0xa8509bda)
		w = basictl.IntWrite(w, elem)
	}
	return w
}

func BuiltinTuple0IntBoxedReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, vec *[0]int32) error {
	index := 0
	if in != nil {
		in.Delim('[')
		if !in.Ok() {
			return internal.ErrorInvalidJSON("[0]int32", "expected json array")
		}
		for ; !in.IsDelim(']'); index++ {
			if index == 0 {
				return internal.ErrorWrongSequenceLength("[0]int32", index+1, 0)
			}
			if err := internal.Json2ReadInt32(in, &(*vec)[index]); err != nil {
				return err
			}
			in.WantComma()
		}
		in.Delim(']')
		if !in.Ok() {
			return internal.ErrorInvalidJSON("[0]int32", "expected json array's end")
		}
	}
	if index != 0 {
		return internal.ErrorWrongSequenceLength("[0]int32", index+1, 0)
	}
	return nil
}

func BuiltinTuple0IntBoxedWriteJSON(w []byte, vec *[0]int32) []byte {
	return BuiltinTuple0IntBoxedWriteJSONOpt(true, false, w, vec)
}
func BuiltinTuple0IntBoxedWriteJSONOpt(newTypeNames bool, short bool, w []byte, vec *[0]int32) []byte {
	w = append(w, '[')
	for _, elem := range *vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = basictl.JSONWriteInt32(w, elem)
	}
	return append(w, ']')
}
