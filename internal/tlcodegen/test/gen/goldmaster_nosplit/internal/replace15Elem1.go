// Copyright 2025 V Kontakte LLC
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

func BuiltinTupleReplace15Elem1FillRandom(rg *basictl.RandGenerator, vec *[]Replace15Elem1, nat_n uint32, nat_t uint32) {
	rg.IncreaseDepth()
	*vec = make([]Replace15Elem1, nat_n)
	for i := range *vec {
		(*vec)[i].FillRandom(rg, nat_t)
	}
	rg.DecreaseDepth()
}

func BuiltinTupleReplace15Elem1Read(w []byte, vec *[]Replace15Elem1, nat_n uint32, nat_t uint32) (_ []byte, err error) {
	if err = basictl.CheckLengthSanity(w, nat_n, 4); err != nil {
		return w, err
	}
	if uint32(cap(*vec)) < nat_n {
		*vec = make([]Replace15Elem1, nat_n)
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

func BuiltinTupleReplace15Elem1Write(w []byte, vec []Replace15Elem1, nat_n uint32, nat_t uint32) (_ []byte, err error) {
	if uint32(len(vec)) != nat_n {
		return w, ErrorWrongSequenceLength("[]Replace15Elem1", len(vec), nat_n)
	}
	for _, elem := range vec {
		w = elem.Write(w, nat_t)
	}
	return w, nil
}

func BuiltinTupleReplace15Elem1CalculateLayout(sizes []int, vec *[]Replace15Elem1, nat_n uint32, nat_t uint32) []int {
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

func BuiltinTupleReplace15Elem1InternalWriteTL2(w []byte, sizes []int, vec *[]Replace15Elem1, nat_n uint32, nat_t uint32) ([]byte, []int) {
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

func BuiltinTupleReplace15Elem1ReadTL2(r []byte, vec *[]Replace15Elem1, nat_n uint32, nat_t uint32) (_ []byte, err error) {
	saveR := r
	currentSize := 0
	if r, currentSize, err = basictl.TL2ParseSize(r); err != nil {
		return r, err
	}
	shift := currentSize + basictl.TL2CalculateSize(currentSize)

	if uint32(cap(*vec)) < nat_n {
		*vec = make([]Replace15Elem1, nat_n)
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
func BuiltinTupleReplace15Elem1ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, vec *[]Replace15Elem1, nat_n uint32, nat_t uint32) error {
	if uint32(cap(*vec)) < nat_n {
		*vec = make([]Replace15Elem1, nat_n)
	} else {
		*vec = (*vec)[:nat_n]
	}
	index := 0
	if in != nil {
		in.Delim('[')
		if !in.Ok() {
			return ErrorInvalidJSON("[]Replace15Elem1", "expected json array")
		}
		for ; !in.IsDelim(']'); index++ {
			if nat_n <= uint32(index) {
				return ErrorInvalidJSON("[]Replace15Elem1", "array is longer than expected")
			}
			if err := (*vec)[index].ReadJSON(legacyTypeNames, in, nat_t); err != nil {
				return err
			}
			in.WantComma()
		}
		in.Delim(']')
		if !in.Ok() {
			return ErrorInvalidJSON("[]Replace15Elem1", "expected json array's end")
		}
	}
	if uint32(index) != nat_n {
		return ErrorWrongSequenceLength("[]Replace15Elem1", index, nat_n)
	}
	return nil
}

func BuiltinTupleReplace15Elem1WriteJSON(w []byte, vec []Replace15Elem1, nat_n uint32, nat_t uint32) (_ []byte, err error) {
	return BuiltinTupleReplace15Elem1WriteJSONOpt(true, false, w, vec, nat_n, nat_t)
}
func BuiltinTupleReplace15Elem1WriteJSONOpt(newTypeNames bool, short bool, w []byte, vec []Replace15Elem1, nat_n uint32, nat_t uint32) (_ []byte, err error) {
	if uint32(len(vec)) != nat_n {
		return w, ErrorWrongSequenceLength("[]Replace15Elem1", len(vec), nat_n)
	}
	w = append(w, '[')
	for _, elem := range vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = elem.WriteJSONOpt(newTypeNames, short, w, nat_t)
	}
	return append(w, ']'), nil
}

type Replace15Elem1 struct {
	X int64
	Y int64
}

func (item *Replace15Elem1) Reset() {
	item.X = 0
	item.Y = 0
}

func (item *Replace15Elem1) FillRandom(rg *basictl.RandGenerator, nat_n uint32) {
	item.X = basictl.RandomLong(rg)
	item.Y = basictl.RandomLong(rg)
}

func (item *Replace15Elem1) Read(w []byte, nat_n uint32) (_ []byte, err error) {
	if w, err = basictl.LongRead(w, &item.X); err != nil {
		return w, err
	}
	return basictl.LongRead(w, &item.Y)
}

func (item *Replace15Elem1) WriteGeneral(w []byte, nat_n uint32) (_ []byte, err error) {
	return item.Write(w, nat_n), nil
}

func (item *Replace15Elem1) Write(w []byte, nat_n uint32) []byte {
	w = basictl.LongWrite(w, item.X)
	w = basictl.LongWrite(w, item.Y)
	return w
}

func (item *Replace15Elem1) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, nat_n uint32) error {
	var propXPresented bool
	var propYPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "x":
				if propXPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("replace15Elem1", "x")
				}
				if err := Json2ReadInt64(in, &item.X); err != nil {
					return err
				}
				propXPresented = true
			case "y":
				if propYPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("replace15Elem1", "y")
				}
				if err := Json2ReadInt64(in, &item.Y); err != nil {
					return err
				}
				propYPresented = true
			default:
				return ErrorInvalidJSONExcessElement("replace15Elem1", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propXPresented {
		item.X = 0
	}
	if !propYPresented {
		item.Y = 0
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *Replace15Elem1) WriteJSONGeneral(w []byte, nat_n uint32) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w, nat_n), nil
}

func (item *Replace15Elem1) WriteJSON(w []byte, nat_n uint32) []byte {
	return item.WriteJSONOpt(true, false, w, nat_n)
}
func (item *Replace15Elem1) WriteJSONOpt(newTypeNames bool, short bool, w []byte, nat_n uint32) []byte {
	w = append(w, '{')
	backupIndexX := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"x":`...)
	w = basictl.JSONWriteInt64(w, item.X)
	if (item.X != 0) == false {
		w = w[:backupIndexX]
	}
	backupIndexY := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"y":`...)
	w = basictl.JSONWriteInt64(w, item.Y)
	if (item.Y != 0) == false {
		w = w[:backupIndexY]
	}
	return append(w, '}')
}

func (item *Replace15Elem1) CalculateLayout(sizes []int, nat_n uint32) []int {
	sizePosition := len(sizes)
	sizes = append(sizes, 0)

	currentSize := 0
	lastUsedByte := 0

	// calculate layout for item.X
	if item.X != 0 {

		lastUsedByte = 1
		currentSize += 8
	}

	// calculate layout for item.Y
	if item.Y != 0 {

		lastUsedByte = 1
		currentSize += 8
	}

	// append byte for each section until last mentioned field
	if lastUsedByte != 0 {
		currentSize += lastUsedByte
	} else {
		// remove unused values
		sizes = sizes[:sizePosition+1]
	}
	sizes[sizePosition] = currentSize
	return sizes
}

func (item *Replace15Elem1) InternalWriteTL2(w []byte, sizes []int, nat_n uint32) ([]byte, []int) {
	currentSize := sizes[0]
	sizes = sizes[1:]

	serializedSize := 0

	w = basictl.TL2WriteSize(w, currentSize)
	if currentSize == 0 {
		return w, sizes
	}

	var currentBlock byte
	currentBlockPosition := len(w)
	w = append(w, 0)
	serializedSize += 1
	// write item.X
	if item.X != 0 {
		serializedSize += 8
		if 8 != 0 {
			currentBlock |= (1 << 1)
			w = basictl.LongWrite(w, item.X)
		}
	}
	// write item.Y
	if item.Y != 0 {
		serializedSize += 8
		if 8 != 0 {
			currentBlock |= (1 << 2)
			w = basictl.LongWrite(w, item.Y)
		}
	}
	w[currentBlockPosition] = currentBlock
	return w, sizes
}

func (item *Replace15Elem1) ReadTL2(r []byte, nat_n uint32) (_ []byte, err error) {
	saveR := r
	currentSize := 0
	if r, currentSize, err = basictl.TL2ParseSize(r); err != nil {
		return r, err
	}
	shift := currentSize + basictl.TL2CalculateSize(currentSize)

	if currentSize == 0 {
		item.Reset()
	} else {
		var block byte
		if r, err = basictl.ByteReadTL2(r, &block); err != nil {
			return r, err
		}
		// read No of constructor
		if block&1 != 0 {
			var _skip int
			if r, err = basictl.TL2ReadSize(r, &_skip); err != nil {
				return r, err
			}
		}

		// read item.X
		if block&(1<<1) != 0 {
			if r, err = basictl.LongRead(r, &item.X); err != nil {
				return r, err
			}
		} else {
			item.X = 0
		}

		// read item.Y
		if block&(1<<2) != 0 {
			if r, err = basictl.LongRead(r, &item.Y); err != nil {
				return r, err
			}
		} else {
			item.Y = 0
		}
	}

	if len(saveR) < len(r)+shift {
		r = saveR[shift:]
	}
	return r, nil
}
