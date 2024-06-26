// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlBuiltinTupleTuple3Int

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlBuiltinTuple3Int"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

func BuiltinTupleTuple3IntFillRandom(rg *basictl.RandGenerator, vec *[][3]int32, nat_n uint32) {
	rg.IncreaseDepth()
	*vec = make([][3]int32, nat_n)
	for i := range *vec {
		tlBuiltinTuple3Int.BuiltinTuple3IntFillRandom(rg, &(*vec)[i])
	}
	rg.DecreaseDepth()
}

func BuiltinTupleTuple3IntRead(w []byte, vec *[][3]int32, nat_n uint32) (_ []byte, err error) {
	if err = basictl.CheckLengthSanity(w, nat_n, 4); err != nil {
		return w, err
	}
	if uint32(cap(*vec)) < nat_n {
		*vec = make([][3]int32, nat_n)
	} else {
		*vec = (*vec)[:nat_n]
	}
	for i := range *vec {
		if w, err = tlBuiltinTuple3Int.BuiltinTuple3IntRead(w, &(*vec)[i]); err != nil {
			return w, err
		}
	}
	return w, nil
}

func BuiltinTupleTuple3IntWrite(w []byte, vec [][3]int32, nat_n uint32) (_ []byte, err error) {
	if uint32(len(vec)) != nat_n {
		return w, internal.ErrorWrongSequenceLength("[][3]int32", len(vec), nat_n)
	}
	for _, elem := range vec {
		w = tlBuiltinTuple3Int.BuiltinTuple3IntWrite(w, &elem)
	}
	return w, nil
}

func BuiltinTupleTuple3IntReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, vec *[][3]int32, nat_n uint32) error {
	if uint32(cap(*vec)) < nat_n {
		*vec = make([][3]int32, nat_n)
	} else {
		*vec = (*vec)[:nat_n]
	}
	index := 0
	if in != nil {
		in.Delim('[')
		if !in.Ok() {
			return internal.ErrorInvalidJSON("[][3]int32", "expected json array")
		}
		for ; !in.IsDelim(']'); index++ {
			if nat_n <= uint32(index) {
				return internal.ErrorInvalidJSON("[][3]int32", "array is longer than expected")
			}
			if err := tlBuiltinTuple3Int.BuiltinTuple3IntReadJSON(legacyTypeNames, in, &(*vec)[index]); err != nil {
				return err
			}
			in.WantComma()
		}
		in.Delim(']')
		if !in.Ok() {
			return internal.ErrorInvalidJSON("[][3]int32", "expected json array's end")
		}
	}
	if uint32(index) != nat_n {
		return internal.ErrorWrongSequenceLength("[][3]int32", index, nat_n)
	}
	return nil
}

func BuiltinTupleTuple3IntWriteJSON(w []byte, vec [][3]int32, nat_n uint32) (_ []byte, err error) {
	return BuiltinTupleTuple3IntWriteJSONOpt(true, false, w, vec, nat_n)
}
func BuiltinTupleTuple3IntWriteJSONOpt(newTypeNames bool, short bool, w []byte, vec [][3]int32, nat_n uint32) (_ []byte, err error) {
	if uint32(len(vec)) != nat_n {
		return w, internal.ErrorWrongSequenceLength("[][3]int32", len(vec), nat_n)
	}
	w = append(w, '[')
	for _, elem := range vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = tlBuiltinTuple3Int.BuiltinTuple3IntWriteJSONOpt(newTypeNames, short, w, &elem)
	}
	return append(w, ']'), nil
}
