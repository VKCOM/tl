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

var _MyNat3 = [2]UnionElement{
	{TLTag: 0x103a40cf, TLName: "myZero3", TLString: "myZero3#103a40cf"},
	{TLTag: 0x692c291b, TLName: "myPlus3", TLString: "myPlus3#692c291b"},
}

type MyNat3 struct {
	valueMyPlus3 *MyPlus3
	index        int
}

func (item MyNat3) TLName() string { return _MyNat3[item.index].TLName }
func (item MyNat3) TLTag() uint32  { return _MyNat3[item.index].TLTag }

func (item *MyNat3) Reset() { item.index = 0 }
func (item *MyNat3) FillRandom(rg *basictl.RandGenerator) {
	index := basictl.RandomUint(rg) % 2
	switch index {
	case 0:
		item.index = 0
	case 1:
		item.index = 1
		if item.valueMyPlus3 == nil {
			var value MyPlus3
			value.FillRandom(rg)
			item.valueMyPlus3 = &value
		}
	default:
	}
}

func (item *MyNat3) IsMyZero3() bool { return item.index == 0 }

func (item *MyNat3) AsMyZero3() (MyZero3, bool) {
	var value MyZero3
	return value, item.index == 0
}
func (item *MyNat3) ResetToMyZero3() { item.index = 0 }
func (item *MyNat3) SetMyZero3()     { item.index = 0 }

func (item *MyNat3) IsMyPlus3() bool { return item.index == 1 }

func (item *MyNat3) AsMyPlus3() (*MyPlus3, bool) {
	if item.index != 1 {
		return nil, false
	}
	return item.valueMyPlus3, true
}
func (item *MyNat3) ResetToMyPlus3() *MyPlus3 {
	item.index = 1
	if item.valueMyPlus3 == nil {
		var value MyPlus3
		item.valueMyPlus3 = &value
	} else {
		item.valueMyPlus3.Reset()
	}
	return item.valueMyPlus3
}
func (item *MyNat3) SetMyPlus3(value MyPlus3) {
	item.index = 1
	if item.valueMyPlus3 == nil {
		item.valueMyPlus3 = &value
	} else {
		*item.valueMyPlus3 = value
	}
}

func (item *MyNat3) ReadBoxed(w []byte) (_ []byte, err error) {
	var tag uint32
	if w, err = basictl.NatRead(w, &tag); err != nil {
		return w, err
	}
	switch tag {
	case 0x103a40cf:
		item.index = 0
		return w, nil
	case 0x692c291b:
		item.index = 1
		if item.valueMyPlus3 == nil {
			var value MyPlus3
			item.valueMyPlus3 = &value
		}
		return item.valueMyPlus3.Read(w)
	default:
		return w, ErrorInvalidUnionTag("MyNat3", tag)
	}
}

func (item *MyNat3) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *MyNat3) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, _MyNat3[item.index].TLTag)
	switch item.index {
	case 0:
		return w
	case 1:
		w = item.valueMyPlus3.Write(w)
	}
	return w
}

func (item *MyNat3) CalculateLayout(sizes []int) []int {
	switch item.index {
	case 0:
		sizes = append(sizes, 0)
	case 1:
		sizes = (*item.valueMyPlus3).CalculateLayout(sizes)
	}
	return sizes
}

func (item *MyNat3) InternalWriteTL2(w []byte, sizes []int) ([]byte, []int) {
	switch item.index {
	case 0:
		sizes = sizes[1:]
		w = basictl.TL2WriteSize(w, 0)
	case 1:
		w, sizes = item.valueMyPlus3.InternalWriteTL2(w, sizes)
	}
	return w, sizes
}

func (item *MyNat3) InternalReadTL2(r []byte) (_ []byte, err error) {
	currentSize := 0
	if r, currentSize, err = basictl.TL2ParseSize(r); err != nil {
		return r, err
	}

	currentR := r[:currentSize]
	r = r[currentSize:]

	var block byte
	if currentSize == 0 {
		item.index = 0
	} else {
		if currentR, err = basictl.ByteReadTL2(currentR, &block); err != nil {
			return r, err
		}
		if (block & 1) != 0 {
			if currentR, item.index, err = basictl.TL2ParseSize(currentR); err != nil {
				return r, err
			}
		} else {
			item.index = 0
		}
	}
	switch item.index {
	case 0:
		break
	case 1:
		if item.valueMyPlus3 == nil {
			var newValue MyPlus3
			item.valueMyPlus3 = &newValue
		}
		if currentR, err = item.valueMyPlus3.InternalReadTL2(currentR, block); err != nil {
			return currentR, err
		}
	}
	return r, nil
}
func (item *MyNat3) WriteTL2(w []byte, ctx *basictl.TL2WriteContext) []byte {
	var sizes []int
	if ctx != nil {
		sizes = ctx.SizeBuffer
	}
	sizes = item.CalculateLayout(sizes[:0])
	w, _ = item.InternalWriteTL2(w, sizes)
	if ctx != nil {
		ctx.SizeBuffer = sizes[:0]
	}
	return w
}

func (item *MyNat3) ReadTL2(r []byte, ctx *basictl.TL2ReadContext) ([]byte, error) {
	return item.InternalReadTL2(r)
}

func (item *MyNat3) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	_tag, _value, err := Json2ReadUnion("MyNat3", in)
	if err != nil {
		return err
	}
	switch _tag {
	case "myZero3#103a40cf", "myZero3", "#103a40cf":
		if !legacyTypeNames && _tag == "myZero3#103a40cf" {
			return ErrorInvalidUnionLegacyTagJSON("MyNat3", "myZero3#103a40cf")
		}
		item.index = 0
	case "myPlus3#692c291b", "myPlus3", "#692c291b":
		if !legacyTypeNames && _tag == "myPlus3#692c291b" {
			return ErrorInvalidUnionLegacyTagJSON("MyNat3", "myPlus3#692c291b")
		}
		item.index = 1
		if item.valueMyPlus3 == nil {
			var value MyPlus3
			item.valueMyPlus3 = &value
		}
		var in2Pointer *basictl.JsonLexer
		if _value != nil {
			in2 := basictl.JsonLexer{Data: _value}
			in2Pointer = &in2
		}
		if err := item.valueMyPlus3.ReadJSON(legacyTypeNames, in2Pointer); err != nil {
			return err
		}
	default:
		return ErrorInvalidUnionTagJSON("MyNat3", _tag)
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *MyNat3) WriteJSONGeneral(tctx *basictl.JSONWriteContext, w []byte) ([]byte, error) {
	return item.WriteJSONOpt(tctx, w), nil
}

func (item *MyNat3) WriteJSON(w []byte) []byte {
	tctx := basictl.JSONWriteContext{}
	return item.WriteJSONOpt(&tctx, w)
}
func (item *MyNat3) WriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte) []byte {
	switch item.index {
	case 0:
		if tctx.LegacyTypeNames {
			w = append(w, `{"type":"myZero3#103a40cf"`...)
		} else {
			w = append(w, `{"type":"myZero3"`...)
		}
		return append(w, '}')
	case 1:
		if tctx.LegacyTypeNames {
			w = append(w, `{"type":"myPlus3#692c291b"`...)
		} else {
			w = append(w, `{"type":"myPlus3"`...)
		}
		w = append(w, `,"value":`...)
		w = item.valueMyPlus3.WriteJSONOpt(tctx, w)
		return append(w, '}')
	default: // Impossible due to panic above
		return w
	}
}

func (item MyNat3) String() string {
	return string(item.WriteJSON(nil))
}

func (item *MyNat3) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *MyNat3) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("MyNat3", err.Error())
	}
	return nil
}

func (item MyPlus3) AsUnion() MyNat3 {
	var ret MyNat3
	ret.SetMyPlus3(item)
	return ret
}

type MyPlus3 MyNat3

func (MyPlus3) TLName() string { return "myPlus3" }
func (MyPlus3) TLTag() uint32  { return 0x692c291b }

func (item *MyPlus3) Reset() {
	ptr := (*MyNat3)(item)
	ptr.Reset()
}

func (item *MyPlus3) FillRandom(rg *basictl.RandGenerator) {
	ptr := (*MyNat3)(item)
	ptr.FillRandom(rg)
}

func (item *MyPlus3) Read(w []byte) (_ []byte, err error) {
	ptr := (*MyNat3)(item)
	return ptr.ReadBoxed(w)
}

func (item *MyPlus3) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *MyPlus3) Write(w []byte) []byte {
	ptr := (*MyNat3)(item)
	return ptr.WriteBoxed(w)
}

func (item *MyPlus3) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x692c291b); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *MyPlus3) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *MyPlus3) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x692c291b)
	return item.Write(w)
}

func (item MyPlus3) String() string {
	return string(item.WriteJSON(nil))
}
func (item *MyPlus3) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	ptr := (*MyNat3)(item)
	if err := ptr.ReadJSON(legacyTypeNames, in); err != nil {
		return err
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *MyPlus3) WriteJSONGeneral(tctx *basictl.JSONWriteContext, w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(tctx, w), nil
}

func (item *MyPlus3) WriteJSON(w []byte) []byte {
	tctx := basictl.JSONWriteContext{}
	return item.WriteJSONOpt(&tctx, w)
}

func (item *MyPlus3) WriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte) []byte {
	ptr := (*MyNat3)(item)
	w = ptr.WriteJSONOpt(tctx, w)
	return w
}
func (item *MyPlus3) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *MyPlus3) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("myPlus3", err.Error())
	}
	return nil
}

func (item *MyPlus3) CalculateLayout(sizes []int) []int {
	ptr := (*MyNat3)(item)
	sizePosition := len(sizes)
	sizes = append(sizes, 0)

	currentSize := 0
	lastUsedByte := 0

	// add constructor No for union type in case of non first option
	lastUsedByte = 1
	currentSize += basictl.TL2CalculateSize(1)

	// calculate layout for ptr
	currentPosition := len(sizes)
	sizes = (*ptr).CalculateLayout(sizes)
	if sizes[currentPosition] != 0 {
		lastUsedByte = 1
		currentSize += sizes[currentPosition]
		currentSize += basictl.TL2CalculateSize(sizes[currentPosition])
	} else {
		sizes = sizes[:currentPosition+1]
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

func (item *MyPlus3) InternalWriteTL2(w []byte, sizes []int) ([]byte, []int) {
	ptr := (*MyNat3)(item)
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

	// add constructor No for union type in case of non first option
	currentBlock |= (1 << 0)

	w = basictl.TL2WriteSize(w, 1)
	serializedSize += basictl.TL2CalculateSize(1)
	// write ptr
	serializedSize += sizes[0]
	if sizes[0] != 0 {
		serializedSize += basictl.TL2CalculateSize(sizes[0])
		currentBlock |= (1 << 1)
		w, sizes = ptr.InternalWriteTL2(w, sizes)
	} else {
		sizes = sizes[1:]
	}
	w[currentBlockPosition] = currentBlock
	return w, sizes
}

func (item *MyPlus3) WriteTL2(w []byte, ctx *basictl.TL2WriteContext) []byte {
	var sizes []int
	if ctx != nil {
		sizes = ctx.SizeBuffer
	}
	sizes = item.CalculateLayout(sizes[:0])
	w, _ = item.InternalWriteTL2(w, sizes)
	if ctx != nil {
		ctx.SizeBuffer = sizes[:0]
	}
	return w
}

func (item *MyPlus3) InternalReadTL2(r []byte, block byte) (_ []byte, err error) {
	currentR := r
	ptr := (*MyNat3)(item)

	// read ptr
	if block&(1<<1) != 0 {
		if ptr == nil {
			var newValue MyNat3
			ptr = &newValue
		}
		if currentR, err = ptr.InternalReadTL2(currentR); err != nil {
			return currentR, err
		}
	} else {
		if ptr == nil {
			var newValue MyNat3
			ptr = &newValue
		}
		ptr.Reset()
	}

	return r, nil
}

func (item *MyPlus3) ReadTL2(r []byte, ctx *basictl.TL2ReadContext) (_ []byte, err error) {
	currentSize := 0
	if r, currentSize, err = basictl.TL2ParseSize(r); err != nil {
		return r, err
	}

	currentR := r[:currentSize]
	r = r[currentSize:]

	var block byte
	var index int
	if currentSize == 0 {
		index = 0
	} else {
		if currentR, err = basictl.ByteReadTL2(currentR, &block); err != nil {
			return r, err
		}
		if (block & 1) != 0 {
			if currentR, index, err = basictl.TL2ParseSize(currentR); err != nil {
				return r, err
			}
		} else {
			index = 0
		}
	}
	if index != 1 {
		return r, basictl.TL2Error("unexpected constructor number %d, instead of %d", index, 1)
	}
	_, err = item.InternalReadTL2(currentR, block)
	return r, err
}

func (item MyZero3) AsUnion() MyNat3 {
	var ret MyNat3
	ret.SetMyZero3()
	return ret
}

type MyZero3 struct {
}

func (MyZero3) TLName() string { return "myZero3" }
func (MyZero3) TLTag() uint32  { return 0x103a40cf }

func (item *MyZero3) Reset() {}

func (item *MyZero3) FillRandom(rg *basictl.RandGenerator) {}

func (item *MyZero3) Read(w []byte) (_ []byte, err error) { return w, nil }

func (item *MyZero3) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *MyZero3) Write(w []byte) []byte {
	return w
}

func (item *MyZero3) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x103a40cf); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *MyZero3) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *MyZero3) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x103a40cf)
	return item.Write(w)
}

func (item MyZero3) String() string {
	return string(item.WriteJSON(nil))
}

func (item *MyZero3) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			return ErrorInvalidJSON("myZero3", "this object can't have properties")
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *MyZero3) WriteJSONGeneral(tctx *basictl.JSONWriteContext, w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(tctx, w), nil
}

func (item *MyZero3) WriteJSON(w []byte) []byte {
	tctx := basictl.JSONWriteContext{}
	return item.WriteJSONOpt(&tctx, w)
}
func (item *MyZero3) WriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte) []byte {
	w = append(w, '{')
	return append(w, '}')
}

func (item *MyZero3) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *MyZero3) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("myZero3", err.Error())
	}
	return nil
}

func (item *MyZero3) CalculateLayout(sizes []int) []int {
	sizePosition := len(sizes)
	sizes = append(sizes, 0)

	currentSize := 0
	lastUsedByte := 0

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

func (item *MyZero3) InternalWriteTL2(w []byte, sizes []int) ([]byte, []int) {
	currentSize := sizes[0]
	sizes = sizes[1:]

	serializedSize := 0

	w = basictl.TL2WriteSize(w, currentSize)
	if currentSize == 0 {
		return w, sizes
	}

	w = append(w, 0)
	serializedSize += 1
	return w, sizes
}

func (item *MyZero3) WriteTL2(w []byte, ctx *basictl.TL2WriteContext) []byte {
	var sizes []int
	if ctx != nil {
		sizes = ctx.SizeBuffer
	}
	sizes = item.CalculateLayout(sizes[:0])
	w, _ = item.InternalWriteTL2(w, sizes)
	if ctx != nil {
		ctx.SizeBuffer = sizes[:0]
	}
	return w
}

func (item *MyZero3) InternalReadTL2(r []byte, block byte) (_ []byte, err error) {

	return r, nil
}

func (item *MyZero3) ReadTL2(r []byte, ctx *basictl.TL2ReadContext) (_ []byte, err error) {
	currentSize := 0
	if r, currentSize, err = basictl.TL2ParseSize(r); err != nil {
		return r, err
	}

	currentR := r[:currentSize]
	r = r[currentSize:]

	var block byte
	var index int
	if currentSize == 0 {
		index = 0
	} else {
		if currentR, err = basictl.ByteReadTL2(currentR, &block); err != nil {
			return r, err
		}
		if (block & 1) != 0 {
			if currentR, index, err = basictl.TL2ParseSize(currentR); err != nil {
				return r, err
			}
		} else {
			index = 0
		}
	}
	if index != 0 {
		return r, basictl.TL2Error("unexpected constructor number %d, instead of %d", index, 0)
	}
	_, err = item.InternalReadTL2(currentR, block)
	return r, err
}
