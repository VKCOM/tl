// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlBuiltinVectorInt

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

func BuiltinVectorIntFillRandom(rg *basictl.RandGenerator, vec *[]int32) {
	rg.IncreaseDepth()
	l := rg.LimitValue(basictl.RandomUint(rg))
	*vec = make([]int32, l)
	for i := range *vec {
		(*vec)[i] = basictl.RandomInt(rg)
	}
	rg.DecreaseDepth()
}
func BuiltinVectorIntRead(w []byte, vec *[]int32) (_ []byte, err error) {
	var l uint32
	if w, err = basictl.NatRead(w, &l); err != nil {
		return w, err
	}
	if uint32(cap(*vec)) < l {
		*vec = make([]int32, l)
	} else {
		*vec = (*vec)[:l]
	}
	for i := range *vec {
		if w, err = basictl.IntRead(w, &(*vec)[i]); err != nil {
			return w, err
		}
	}
	return w, nil
}

func BuiltinVectorIntWrite(w []byte, vec []int32) []byte {
	w = basictl.NatWrite(w, uint32(len(vec)))
	for _, elem := range vec {
		w = basictl.IntWrite(w, elem)
	}
	return w
}

func BuiltinVectorIntCalculateLayout(sizes []int, vec *[]int32) []int {
	currentSize := 0
	sizePosition := len(sizes)
	sizes = append(sizes, 0)
	if len(*vec) != 0 {
		currentSize += basictl.TL2CalculateSize(len(*vec))
	}
	for i := 0; i < len(*vec); i++ {

		currentSize += 4
	}
	sizes[sizePosition] = currentSize
	return sizes
}

func BuiltinVectorIntInternalWriteTL2(w []byte, sizes []int, vec *[]int32) ([]byte, []int) {
	currentSize := sizes[0]
	sizes = sizes[1:]

	w = basictl.TL2WriteSize(w, currentSize)
	if len(*vec) != 0 {
		w = basictl.TL2WriteSize(w, len(*vec))
	}

	for i := 0; i < len(*vec); i++ {
		elem := (*vec)[i]
		w = basictl.IntWrite(w, elem)
	}
	return w, sizes
}

func BuiltinVectorIntInternalReadTL2(r []byte, vec *[]int32) (_ []byte, err error) {
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
		*vec = make([]int32, elementCount)
	}
	*vec = (*vec)[:elementCount]
	for i := 0; i < elementCount; i++ {
		if currentR, err = basictl.IntRead(currentR, &(*vec)[i]); err != nil {
			return currentR, err
		}
	}
	return r, nil
}

func BuiltinVectorIntReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, vec *[]int32) error {
	*vec = (*vec)[:cap(*vec)]
	index := 0
	if in != nil {
		in.Delim('[')
		if !in.Ok() {
			return internal.ErrorInvalidJSON("[]int32", "expected json array")
		}
		for ; !in.IsDelim(']'); index++ {
			if len(*vec) <= index {
				var newValue int32
				*vec = append(*vec, newValue)
				*vec = (*vec)[:cap(*vec)]
			}
			if err := internal.Json2ReadInt32(in, &(*vec)[index]); err != nil {
				return err
			}
			in.WantComma()
		}
		in.Delim(']')
		if !in.Ok() {
			return internal.ErrorInvalidJSON("[]int32", "expected json array's end")
		}
	}
	*vec = (*vec)[:index]
	return nil
}

func BuiltinVectorIntWriteJSON(w []byte, vec []int32) []byte {
	tctx := basictl.JSONWriteContext{}
	return BuiltinVectorIntWriteJSONOpt(&tctx, w, vec)
}
func BuiltinVectorIntWriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte, vec []int32) []byte {
	w = append(w, '[')
	for _, elem := range vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = basictl.JSONWriteInt32(w, elem)
	}
	return append(w, ']')
}
