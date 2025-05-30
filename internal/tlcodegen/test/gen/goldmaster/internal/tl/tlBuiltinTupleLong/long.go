// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlBuiltinTupleLong

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

func BuiltinTupleLongFillRandom(rg *basictl.RandGenerator, vec *[]int64, nat_n uint32) {
	rg.IncreaseDepth()
	*vec = make([]int64, nat_n)
	for i := range *vec {
		(*vec)[i] = basictl.RandomLong(rg)
	}
	rg.DecreaseDepth()
}

func BuiltinTupleLongRead(w []byte, vec *[]int64, nat_n uint32) (_ []byte, err error) {
	if uint32(cap(*vec)) < nat_n {
		*vec = make([]int64, nat_n)
	} else {
		*vec = (*vec)[:nat_n]
	}
	for i := range *vec {
		if w, err = basictl.LongRead(w, &(*vec)[i]); err != nil {
			return w, err
		}
	}
	return w, nil
}

func BuiltinTupleLongWrite(w []byte, vec []int64, nat_n uint32) (_ []byte, err error) {
	if uint32(len(vec)) != nat_n {
		return w, internal.ErrorWrongSequenceLength("[]int64", len(vec), nat_n)
	}
	for _, elem := range vec {
		w = basictl.LongWrite(w, elem)
	}
	return w, nil
}

func BuiltinTupleLongCalculateLayout(sizes []int, vec *[]int64, nat_n uint32) []int {
	sizePosition := len(sizes)
	sizes = append(sizes, 0)

	for i := 0; i < len(*vec); i++ {

		sizes[sizePosition] += 8
	}
	return sizes
}

func BuiltinTupleLongInternalWriteTL2(w []byte, sizes []int, vec *[]int64, nat_n uint32) ([]byte, []int) {
	currentSize := sizes[0]
	sizes = sizes[1:]

	w = basictl.TL2WriteSize(w, currentSize)
	if currentSize == 0 {
		return w, sizes
	}

	for i := 0; i < len(*vec); i++ {
		w = basictl.LongWrite(w, (*vec)[i])
	}
	return w, sizes
}

func BuiltinTupleLongReadTL2(r []byte, vec *[]int64, nat_n uint32) (_ []byte, err error) {
	saveR := r
	currentSize := 0
	if r, currentSize, err = basictl.TL2ParseSize(r); err != nil {
		return r, err
	}
	shift := currentSize + basictl.TL2CalculateSize(currentSize)

	if uint32(cap(*vec)) < nat_n {
		*vec = make([]int64, nat_n)
	} else {
		*vec = (*vec)[:nat_n]
	}
	i := 0
	for len(saveR) < len(r)+shift {
		if uint32(i) == nat_n {
			return r, basictl.TL2Error("more elements than expected")
		}
		if r, err = basictl.LongRead(r, &(*vec)[i]); err != nil {
			return r, err
		}
		i += 1
	}
	if uint32(i) != nat_n {
		return r, basictl.TL2Error("less elements than expected")
	}
	return r, nil
}
func BuiltinTupleLongReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, vec *[]int64, nat_n uint32) error {
	if uint32(cap(*vec)) < nat_n {
		*vec = make([]int64, nat_n)
	} else {
		*vec = (*vec)[:nat_n]
	}
	index := 0
	if in != nil {
		in.Delim('[')
		if !in.Ok() {
			return internal.ErrorInvalidJSON("[]int64", "expected json array")
		}
		for ; !in.IsDelim(']'); index++ {
			if nat_n <= uint32(index) {
				return internal.ErrorInvalidJSON("[]int64", "array is longer than expected")
			}
			if err := internal.Json2ReadInt64(in, &(*vec)[index]); err != nil {
				return err
			}
			in.WantComma()
		}
		in.Delim(']')
		if !in.Ok() {
			return internal.ErrorInvalidJSON("[]int64", "expected json array's end")
		}
	}
	if uint32(index) != nat_n {
		return internal.ErrorWrongSequenceLength("[]int64", index, nat_n)
	}
	return nil
}

func BuiltinTupleLongWriteJSON(w []byte, vec []int64, nat_n uint32) (_ []byte, err error) {
	return BuiltinTupleLongWriteJSONOpt(true, false, w, vec, nat_n)
}
func BuiltinTupleLongWriteJSONOpt(newTypeNames bool, short bool, w []byte, vec []int64, nat_n uint32) (_ []byte, err error) {
	if uint32(len(vec)) != nat_n {
		return w, internal.ErrorWrongSequenceLength("[]int64", len(vec), nat_n)
	}
	w = append(w, '[')
	for _, elem := range vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = basictl.JSONWriteInt64(w, elem)
	}
	return append(w, ']'), nil
}
