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

func BuiltinTupleReplace15Elem2FillRandom(rg *basictl.RandGenerator, vec *[]Replace15Elem2, nat_n uint32, nat_t uint32) {
	rg.IncreaseDepth()
	*vec = make([]Replace15Elem2, nat_n)
	for i := range *vec {
		(*vec)[i].FillRandom(rg, nat_t)
	}
	rg.DecreaseDepth()
}

func BuiltinTupleReplace15Elem2Read(w []byte, vec *[]Replace15Elem2, nat_n uint32, nat_t uint32) (_ []byte, err error) {
	if err = basictl.CheckLengthSanity(w, nat_n, 4); err != nil {
		return w, err
	}
	if uint32(cap(*vec)) < nat_n {
		*vec = make([]Replace15Elem2, nat_n)
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

func BuiltinTupleReplace15Elem2Write(w []byte, vec []Replace15Elem2, nat_n uint32, nat_t uint32) (_ []byte, err error) {
	if uint32(len(vec)) != nat_n {
		return w, ErrorWrongSequenceLength("[]Replace15Elem2", len(vec), nat_n)
	}
	for _, elem := range vec {
		w = elem.Write(w, nat_t)
	}
	return w, nil
}

func BuiltinTupleReplace15Elem2CalculateLayout(sizes []int, vec *[]Replace15Elem2, nat_n uint32, nat_t uint32) []int {
	currentSize := 0
	sizePosition := len(sizes)
	sizes = append(sizes, 0)
	if nat_n != 0 {
		currentSize += basictl.TL2CalculateSize(int(nat_n))
	}

	lastIndex := uint32(len(*vec))
	if lastIndex > nat_n {
		lastIndex = nat_n
	}

	for i := uint32(0); i < lastIndex; i++ {
		currentPosition := len(sizes)
		sizes = (*vec)[i].CalculateLayout(sizes, nat_t)
		currentSize += sizes[currentPosition]
		currentSize += basictl.TL2CalculateSize(sizes[currentPosition])
	}

	// append empty objects if not enough
	for i := lastIndex; i < nat_n; i++ {
		var elem Replace15Elem2
		currentPosition := len(sizes)
		sizes = elem.CalculateLayout(sizes, nat_t)
		currentSize += sizes[currentPosition]
		currentSize += basictl.TL2CalculateSize(sizes[currentPosition])
	}

	sizes[sizePosition] = currentSize
	return sizes
}

func BuiltinTupleReplace15Elem2InternalWriteTL2(w []byte, sizes []int, vec *[]Replace15Elem2, nat_n uint32, nat_t uint32) ([]byte, []int) {
	currentSize := sizes[0]
	sizes = sizes[1:]

	w = basictl.TL2WriteSize(w, currentSize)
	if nat_n != 0 {
		w = basictl.TL2WriteSize(w, int(nat_n))
	}

	lastIndex := uint32(len(*vec))
	if lastIndex > nat_n {
		lastIndex = nat_n
	}

	for i := uint32(0); i < lastIndex; i++ {
		w, sizes = (*vec)[i].InternalWriteTL2(w, sizes, nat_t)
	}

	// append empty objects if not enough
	for i := lastIndex; i < nat_n; i++ {
		var elem Replace15Elem2
		w, sizes = elem.InternalWriteTL2(w, sizes, nat_t)
	}
	return w, sizes
}

func BuiltinTupleReplace15Elem2InternalReadTL2(r []byte, vec *[]Replace15Elem2, nat_n uint32, nat_t uint32) (_ []byte, err error) {
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

	if uint32(cap(*vec)) < nat_n {
		*vec = make([]Replace15Elem2, nat_n)
	} else {
		*vec = (*vec)[:nat_n]
	}

	lastIndex := uint32(elementCount)
	if lastIndex > nat_n {
		lastIndex = nat_n
	}

	for i := uint32(0); i < lastIndex; i++ {
		if currentR, err = (*vec)[i].InternalReadTL2(currentR, nat_t); err != nil {
			return currentR, err
		}
	}

	// reset elements if received less elements
	for i := lastIndex; i < nat_n; i++ {
		(*vec)[i].Reset()
	}

	return r, nil
}
func BuiltinTupleReplace15Elem2ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, vec *[]Replace15Elem2, nat_n uint32, nat_t uint32) error {
	if uint32(cap(*vec)) < nat_n {
		*vec = make([]Replace15Elem2, nat_n)
	} else {
		*vec = (*vec)[:nat_n]
	}
	index := 0
	if in != nil {
		in.Delim('[')
		if !in.Ok() {
			return ErrorInvalidJSON("[]Replace15Elem2", "expected json array")
		}
		for ; !in.IsDelim(']'); index++ {
			if nat_n <= uint32(index) {
				return ErrorInvalidJSON("[]Replace15Elem2", "array is longer than expected")
			}
			if err := (*vec)[index].ReadJSON(legacyTypeNames, in, nat_t); err != nil {
				return err
			}
			in.WantComma()
		}
		in.Delim(']')
		if !in.Ok() {
			return ErrorInvalidJSON("[]Replace15Elem2", "expected json array's end")
		}
	}
	if uint32(index) != nat_n {
		return ErrorWrongSequenceLength("[]Replace15Elem2", index, nat_n)
	}
	return nil
}

func BuiltinTupleReplace15Elem2WriteJSON(w []byte, vec []Replace15Elem2, nat_n uint32, nat_t uint32) (_ []byte, err error) {
	tctx := basictl.JSONWriteContext{}
	return BuiltinTupleReplace15Elem2WriteJSONOpt(&tctx, w, vec, nat_n, nat_t)
}
func BuiltinTupleReplace15Elem2WriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte, vec []Replace15Elem2, nat_n uint32, nat_t uint32) (_ []byte, err error) {
	if uint32(len(vec)) != nat_n {
		return w, ErrorWrongSequenceLength("[]Replace15Elem2", len(vec), nat_n)
	}
	w = append(w, '[')
	for _, elem := range vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = elem.WriteJSONOpt(tctx, w, nat_t)
	}
	return append(w, ']'), nil
}

type Replace15Elem2 struct {
	X int32
	Y int32
	Z int32
}

func (item *Replace15Elem2) Reset() {
	item.X = 0
	item.Y = 0
	item.Z = 0
}

func (item *Replace15Elem2) FillRandom(rg *basictl.RandGenerator, nat_n uint32) {
	item.X = basictl.RandomInt(rg)
	item.Y = basictl.RandomInt(rg)
	item.Z = basictl.RandomInt(rg)
}

func (item *Replace15Elem2) Read(w []byte, nat_n uint32) (_ []byte, err error) {
	if w, err = basictl.IntRead(w, &item.X); err != nil {
		return w, err
	}
	if w, err = basictl.IntRead(w, &item.Y); err != nil {
		return w, err
	}
	return basictl.IntRead(w, &item.Z)
}

func (item *Replace15Elem2) WriteGeneral(w []byte, nat_n uint32) (_ []byte, err error) {
	return item.Write(w, nat_n), nil
}

func (item *Replace15Elem2) Write(w []byte, nat_n uint32) []byte {
	w = basictl.IntWrite(w, item.X)
	w = basictl.IntWrite(w, item.Y)
	w = basictl.IntWrite(w, item.Z)
	return w
}

func (item *Replace15Elem2) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, nat_n uint32) error {
	var propXPresented bool
	var propYPresented bool
	var propZPresented bool

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
					return ErrorInvalidJSONWithDuplicatingKeys("replace15Elem2", "x")
				}
				if err := Json2ReadInt32(in, &item.X); err != nil {
					return err
				}
				propXPresented = true
			case "y":
				if propYPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("replace15Elem2", "y")
				}
				if err := Json2ReadInt32(in, &item.Y); err != nil {
					return err
				}
				propYPresented = true
			case "z":
				if propZPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("replace15Elem2", "z")
				}
				if err := Json2ReadInt32(in, &item.Z); err != nil {
					return err
				}
				propZPresented = true
			default:
				return ErrorInvalidJSONExcessElement("replace15Elem2", key)
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
	if !propZPresented {
		item.Z = 0
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *Replace15Elem2) WriteJSONGeneral(tctx *basictl.JSONWriteContext, w []byte, nat_n uint32) (_ []byte, err error) {
	return item.WriteJSONOpt(tctx, w, nat_n), nil
}

func (item *Replace15Elem2) WriteJSON(w []byte, nat_n uint32) []byte {
	tctx := basictl.JSONWriteContext{}
	return item.WriteJSONOpt(&tctx, w, nat_n)
}
func (item *Replace15Elem2) WriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte, nat_n uint32) []byte {
	w = append(w, '{')
	backupIndexX := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"x":`...)
	w = basictl.JSONWriteInt32(w, item.X)
	if (item.X != 0) == false {
		w = w[:backupIndexX]
	}
	backupIndexY := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"y":`...)
	w = basictl.JSONWriteInt32(w, item.Y)
	if (item.Y != 0) == false {
		w = w[:backupIndexY]
	}
	backupIndexZ := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"z":`...)
	w = basictl.JSONWriteInt32(w, item.Z)
	if (item.Z != 0) == false {
		w = w[:backupIndexZ]
	}
	return append(w, '}')
}

func (item *Replace15Elem2) CalculateLayout(sizes []int, nat_n uint32) []int {
	sizePosition := len(sizes)
	sizes = append(sizes, 0)

	currentSize := 0
	lastUsedByte := 0

	// calculate layout for item.X
	if item.X != 0 {

		lastUsedByte = 1
		currentSize += 4
	}

	// calculate layout for item.Y
	if item.Y != 0 {

		lastUsedByte = 1
		currentSize += 4
	}

	// calculate layout for item.Z
	if item.Z != 0 {

		lastUsedByte = 1
		currentSize += 4
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

func (item *Replace15Elem2) InternalWriteTL2(w []byte, sizes []int, nat_n uint32) ([]byte, []int) {
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
		serializedSize += 4
		if 4 != 0 {
			currentBlock |= (1 << 1)
			w = basictl.IntWrite(w, item.X)
		}
	}
	// write item.Y
	if item.Y != 0 {
		serializedSize += 4
		if 4 != 0 {
			currentBlock |= (1 << 2)
			w = basictl.IntWrite(w, item.Y)
		}
	}
	// write item.Z
	if item.Z != 0 {
		serializedSize += 4
		if 4 != 0 {
			currentBlock |= (1 << 3)
			w = basictl.IntWrite(w, item.Z)
		}
	}
	w[currentBlockPosition] = currentBlock
	return w, sizes
}

func (item *Replace15Elem2) WriteTL2(w []byte, ctx *basictl.TL2WriteContext, nat_n uint32) []byte {
	var sizes []int
	if ctx != nil {
		sizes = ctx.SizeBuffer
	}
	sizes = item.CalculateLayout(sizes[:0], nat_n)
	w, _ = item.InternalWriteTL2(w, sizes, nat_n)
	if ctx != nil {
		ctx.SizeBuffer = sizes[:0]
	}
	return w
}

func (item *Replace15Elem2) InternalReadTL2(r []byte, nat_n uint32) (_ []byte, err error) {
	currentSize := 0
	if r, currentSize, err = basictl.TL2ParseSize(r); err != nil {
		return r, err
	}
	if len(r) < currentSize {
		return r, basictl.TL2Error("not enough data: expected %d, got %d", currentSize, len(r))
	}

	currentR := r[:currentSize]
	r = r[currentSize:]

	if currentSize == 0 {
		item.Reset()
		return r, nil
	}
	var block byte
	if currentR, err = basictl.ByteReadTL2(currentR, &block); err != nil {
		return currentR, err
	}
	// read No of constructor
	if block&1 != 0 {
		var index int
		if currentR, err = basictl.TL2ReadSize(currentR, &index); err != nil {
			return currentR, err
		}
		if index != 0 {
			// unknown cases for current type
			item.Reset()
			return r, nil
		}
	}

	// read item.X
	if block&(1<<1) != 0 {
		if currentR, err = basictl.IntRead(currentR, &item.X); err != nil {
			return currentR, err
		}
	} else {
		item.X = 0
	}

	// read item.Y
	if block&(1<<2) != 0 {
		if currentR, err = basictl.IntRead(currentR, &item.Y); err != nil {
			return currentR, err
		}
	} else {
		item.Y = 0
	}

	// read item.Z
	if block&(1<<3) != 0 {
		if currentR, err = basictl.IntRead(currentR, &item.Z); err != nil {
			return currentR, err
		}
	} else {
		item.Z = 0
	}

	return r, nil
}

func (item *Replace15Elem2) ReadTL2(r []byte, ctx *basictl.TL2ReadContext, nat_n uint32) (_ []byte, err error) {
	return item.InternalReadTL2(r, nat_n)
}
