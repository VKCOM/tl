// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlBuiltinTuple1Int

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

func BuiltinTuple1IntReset(vec *[1]int32) {
	for i := range *vec {
		(*vec)[i] = 0
	}
}

func BuiltinTuple1IntFillRandom(rg *basictl.RandGenerator, vec *[1]int32) {
	rg.IncreaseDepth()
	for i := range *vec {
		(*vec)[i] = basictl.RandomInt(rg)
	}
	rg.DecreaseDepth()
}

func BuiltinTuple1IntRead(w []byte, vec *[1]int32) (_ []byte, err error) {
	for i := range *vec {
		if w, err = basictl.IntRead(w, &(*vec)[i]); err != nil {
			return w, err
		}
	}
	return w, nil
}

func BuiltinTuple1IntWrite(w []byte, vec *[1]int32) []byte {
	for _, elem := range *vec {
		w = basictl.IntWrite(w, elem)
	}
	return w
}

func BuiltinTuple1IntCalculateLayout(sizes []int, vec *[1]int32) []int {
	currentSize := 0
	sizePosition := len(sizes)
	sizes = append(sizes, 0)
	if 1 != 0 {
		currentSize += basictl.TL2CalculateSize(1)
	}

	for i := 0; i < 1; i++ {

		currentSize += 4
	}

	sizes[sizePosition] = currentSize
	return sizes
}

func BuiltinTuple1IntInternalWriteTL2(w []byte, sizes []int, vec *[1]int32) ([]byte, []int) {
	currentSize := sizes[0]
	sizes = sizes[1:]

	w = basictl.TL2WriteSize(w, currentSize)
	if 1 != 0 {
		w = basictl.TL2WriteSize(w, 1)
	}

	for i := 0; i < 1; i++ {
		w = basictl.IntWrite(w, (*vec)[i])
	}
	return w, sizes
}

func BuiltinTuple1IntInternalReadTL2(r []byte, vec *[1]int32) (_ []byte, err error) {
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

	lastIndex := elementCount
	if lastIndex > 1 {
		lastIndex = 1
	}
	for i := 0; i < lastIndex; i++ {
		if currentR, err = basictl.IntRead(currentR, &(*vec)[i]); err != nil {
			return currentR, err
		}
	}

	// reset elements if received less elements
	for i := lastIndex; i < 1; i++ {
		(*vec)[i] = 0
	}

	return r, nil
}

func BuiltinTuple1IntReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, vec *[1]int32) error {
	index := 0
	if in != nil {
		in.Delim('[')
		if !in.Ok() {
			return internal.ErrorInvalidJSON("[1]int32", "expected json array")
		}
		for ; !in.IsDelim(']'); index++ {
			if index == 1 {
				return internal.ErrorWrongSequenceLength("[1]int32", index+1, 1)
			}
			if err := internal.Json2ReadInt32(in, &(*vec)[index]); err != nil {
				return err
			}
			in.WantComma()
		}
		in.Delim(']')
		if !in.Ok() {
			return internal.ErrorInvalidJSON("[1]int32", "expected json array's end")
		}
	}
	if index != 1 {
		return internal.ErrorWrongSequenceLength("[1]int32", index+1, 1)
	}
	return nil
}

func BuiltinTuple1IntWriteJSON(w []byte, vec *[1]int32) []byte {
	tctx := basictl.JSONWriteContext{}
	return BuiltinTuple1IntWriteJSONOpt(&tctx, w, vec)
}
func BuiltinTuple1IntWriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte, vec *[1]int32) []byte {
	w = append(w, '[')
	for _, elem := range *vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = basictl.JSONWriteInt32(w, elem)
	}
	return append(w, ']')
}
