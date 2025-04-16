// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package internal

import (
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite

func BuiltinTupleLongFillRandom(rg *basictl.RandGenerator, vec *[]int64, nat_n uint32) {
	rg.IncreaseDepth()
	*vec = make([]int64, nat_n)
	for i := range *vec {
		(*vec)[i] = basictl.RandomLong(rg)
	}
	rg.DecreaseDepth()
}

func BuiltinTupleLongRead(w []byte, vec *[]int64, nat_n uint32) (_ []byte, err error) {
	if err = basictl.CheckLengthSanity(w, nat_n, 4); err != nil {
		return w, err
	}
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
		return w, ErrorWrongSequenceLength("[]int64", len(vec), nat_n)
	}
	for _, elem := range vec {
		w = basictl.LongWrite(w, elem)
	}
	return w, nil
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
			return ErrorInvalidJSON("[]int64", "expected json array")
		}
		for ; !in.IsDelim(']'); index++ {
			if nat_n <= uint32(index) {
				return ErrorInvalidJSON("[]int64", "array is longer than expected")
			}
			if err := Json2ReadInt64(in, &(*vec)[index]); err != nil {
				return err
			}
			in.WantComma()
		}
		in.Delim(']')
		if !in.Ok() {
			return ErrorInvalidJSON("[]int64", "expected json array's end")
		}
	}
	if uint32(index) != nat_n {
		return ErrorWrongSequenceLength("[]int64", index, nat_n)
	}
	return nil
}

func BuiltinTupleLongWriteJSON(w []byte, vec []int64, nat_n uint32) (_ []byte, err error) {
	return BuiltinTupleLongWriteJSONOpt(true, false, w, vec, nat_n)
}
func BuiltinTupleLongWriteJSONOpt(newTypeNames bool, short bool, w []byte, vec []int64, nat_n uint32) (_ []byte, err error) {
	if uint32(len(vec)) != nat_n {
		return w, ErrorWrongSequenceLength("[]int64", len(vec), nat_n)
	}
	w = append(w, '[')
	for _, elem := range vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = basictl.JSONWriteInt64(w, elem)
	}
	return append(w, ']'), nil
}

type Long int64

func (Long) TLName() string { return "long" }
func (Long) TLTag() uint32  { return 0x22076cba }

func (item *Long) Reset() {
	ptr := (*int64)(item)
	*ptr = 0
}

func (item *Long) FillRandom(rg *basictl.RandGenerator) {
	ptr := (*int64)(item)
	*ptr = basictl.RandomLong(rg)
}

func (item *Long) Read(w []byte) (_ []byte, err error) {
	ptr := (*int64)(item)
	return basictl.LongRead(w, ptr)
}

// This method is general version of Write, use it instead!
func (item *Long) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *Long) Write(w []byte) []byte {
	ptr := (*int64)(item)
	return basictl.LongWrite(w, *ptr)
}

func (item *Long) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x22076cba); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *Long) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *Long) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x22076cba)
	return item.Write(w)
}

func (item Long) String() string {
	return string(item.WriteJSON(nil))
}

func (item *Long) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	ptr := (*int64)(item)
	if err := Json2ReadInt64(in, ptr); err != nil {
		return err
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *Long) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSON(w), nil
}

func (item *Long) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}

func (item *Long) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	ptr := (*int64)(item)
	w = basictl.JSONWriteInt64(w, *ptr)
	return w
}
func (item *Long) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *Long) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("long", err.Error())
	}
	return nil
}
