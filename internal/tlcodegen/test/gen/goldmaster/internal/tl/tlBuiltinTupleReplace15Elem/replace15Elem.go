// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlBuiltinTupleReplace15Elem

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlReplace15Elem"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

func BuiltinTupleReplace15ElemFillRandom(rg *basictl.RandGenerator, vec *[]tlReplace15Elem.Replace15Elem, nat_n uint32, nat_t uint32) {
	rg.IncreaseDepth()
	*vec = make([]tlReplace15Elem.Replace15Elem, nat_n)
	for i := range *vec {
		(*vec)[i].FillRandom(rg, nat_t)
	}
	rg.DecreaseDepth()
}

func BuiltinTupleReplace15ElemRead(w []byte, vec *[]tlReplace15Elem.Replace15Elem, nat_n uint32, nat_t uint32) (_ []byte, err error) {
	if uint32(cap(*vec)) < nat_n {
		*vec = make([]tlReplace15Elem.Replace15Elem, nat_n)
	} else {
		*vec = (*vec)[:nat_n]
	}
	for i := range *vec {
		if w, err = (*vec)[i].Read(w, nat_t); err != nil {
			return w, err
		}
	}
	return w, nil
}

func BuiltinTupleReplace15ElemWrite(w []byte, vec []tlReplace15Elem.Replace15Elem, nat_n uint32, nat_t uint32) (_ []byte, err error) {
	if uint32(len(vec)) != nat_n {
		return w, internal.ErrorWrongSequenceLength("[]tlReplace15Elem.Replace15Elem", len(vec), nat_n)
	}
	for _, elem := range vec {
		w = elem.Write(w, nat_t)
	}
	return w, nil
}

func BuiltinTupleReplace15ElemCalculateLayout(sizes []int, vec *[]tlReplace15Elem.Replace15Elem, nat_n uint32, nat_t uint32) []int {
	sizePosition := len(sizes)
	sizes = append(sizes, 0)

	for i := 0; i < len(*vec); i++ {
		currentPosition := len(sizes)
		sizes = (*vec)[i].CalculateLayout(sizes, nat_t)
		sizes[sizePosition] += sizes[currentPosition]
		sizes[sizePosition] += basictl.TL2CalculateSize(sizes[currentPosition])
	}
	return sizes
}

func BuiltinTupleReplace15ElemInternalWriteTL2(w []byte, sizes []int, vec *[]tlReplace15Elem.Replace15Elem, nat_n uint32, nat_t uint32) ([]byte, []int) {
	currentSize := sizes[0]
	sizes = sizes[1:]

	w = basictl.TL2WriteSize(w, currentSize)
	if currentSize == 0 {
		return w, sizes
	}

	for i := 0; i < len(*vec); i++ {
		w, sizes = (*vec)[i].InternalWriteTL2(w, sizes, nat_t)
	}
	return w, sizes
}

func BuiltinTupleReplace15ElemReadTL2(r []byte, vec *[]tlReplace15Elem.Replace15Elem, nat_n uint32, nat_t uint32) (_ []byte, err error) {
	saveR := r
	currentSize := 0
	if r, currentSize, err = basictl.TL2ParseSize(r); err != nil {
		return r, err
	}
	shift := currentSize + basictl.TL2CalculateSize(currentSize)

	if uint32(cap(*vec)) < nat_n {
		*vec = make([]tlReplace15Elem.Replace15Elem, nat_n)
	} else {
		*vec = (*vec)[:nat_n]
	}
	i := 0
	for len(saveR) < len(r)+shift {
		if uint32(i) == nat_n {
			return r, basictl.TL2Error("more elements than expected")
		}
		if r, err = (*vec)[i].ReadTL2(r, nat_t); err != nil {
			return r, err
		}
		i += 1
	}
	if uint32(i) != nat_n {
		return r, basictl.TL2Error("less elements than expected")
	}
	return r, nil
}
func BuiltinTupleReplace15ElemReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, vec *[]tlReplace15Elem.Replace15Elem, nat_n uint32, nat_t uint32) error {
	if uint32(cap(*vec)) < nat_n {
		*vec = make([]tlReplace15Elem.Replace15Elem, nat_n)
	} else {
		*vec = (*vec)[:nat_n]
	}
	index := 0
	if in != nil {
		in.Delim('[')
		if !in.Ok() {
			return internal.ErrorInvalidJSON("[]tlReplace15Elem.Replace15Elem", "expected json array")
		}
		for ; !in.IsDelim(']'); index++ {
			if nat_n <= uint32(index) {
				return internal.ErrorInvalidJSON("[]tlReplace15Elem.Replace15Elem", "array is longer than expected")
			}
			if err := (*vec)[index].ReadJSON(legacyTypeNames, in, nat_t); err != nil {
				return err
			}
			in.WantComma()
		}
		in.Delim(']')
		if !in.Ok() {
			return internal.ErrorInvalidJSON("[]tlReplace15Elem.Replace15Elem", "expected json array's end")
		}
	}
	if uint32(index) != nat_n {
		return internal.ErrorWrongSequenceLength("[]tlReplace15Elem.Replace15Elem", index, nat_n)
	}
	return nil
}

func BuiltinTupleReplace15ElemWriteJSON(w []byte, vec []tlReplace15Elem.Replace15Elem, nat_n uint32, nat_t uint32) (_ []byte, err error) {
	return BuiltinTupleReplace15ElemWriteJSONOpt(true, false, w, vec, nat_n, nat_t)
}
func BuiltinTupleReplace15ElemWriteJSONOpt(newTypeNames bool, short bool, w []byte, vec []tlReplace15Elem.Replace15Elem, nat_n uint32, nat_t uint32) (_ []byte, err error) {
	if uint32(len(vec)) != nat_n {
		return w, internal.ErrorWrongSequenceLength("[]tlReplace15Elem.Replace15Elem", len(vec), nat_n)
	}
	w = append(w, '[')
	for _, elem := range vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = elem.WriteJSONOpt(newTypeNames, short, w, nat_t)
	}
	return append(w, ']'), nil
}
