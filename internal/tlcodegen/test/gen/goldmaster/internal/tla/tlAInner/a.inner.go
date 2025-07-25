// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlAInner

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlBuiltinTupleInt"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type AInner struct {
	A []int32
}

func (AInner) TLName() string { return "a.inner" }
func (AInner) TLTag() uint32  { return 0xec5089b9 }

func (item *AInner) Reset() {
	item.A = item.A[:0]
}

func (item *AInner) FillRandom(rg *basictl.RandGenerator, nat_I uint32) {
	tlBuiltinTupleInt.BuiltinTupleIntFillRandom(rg, &item.A, nat_I)
}

func (item *AInner) Read(w []byte, nat_I uint32) (_ []byte, err error) {
	return tlBuiltinTupleInt.BuiltinTupleIntRead(w, &item.A, nat_I)
}

func (item *AInner) WriteGeneral(w []byte, nat_I uint32) (_ []byte, err error) {
	return item.Write(w, nat_I)
}

func (item *AInner) Write(w []byte, nat_I uint32) (_ []byte, err error) {
	if w, err = tlBuiltinTupleInt.BuiltinTupleIntWrite(w, item.A, nat_I); err != nil {
		return w, err
	}
	return w, nil
}

func (item *AInner) ReadBoxed(w []byte, nat_I uint32) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xec5089b9); err != nil {
		return w, err
	}
	return item.Read(w, nat_I)
}

func (item *AInner) WriteBoxedGeneral(w []byte, nat_I uint32) (_ []byte, err error) {
	return item.WriteBoxed(w, nat_I)
}

func (item *AInner) WriteBoxed(w []byte, nat_I uint32) (_ []byte, err error) {
	w = basictl.NatWrite(w, 0xec5089b9)
	return item.Write(w, nat_I)
}

func (item *AInner) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, nat_I uint32) error {
	var rawA []byte

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "a":
				if rawA != nil {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("a.inner", "a")
				}
				rawA = in.Raw()
				if !in.Ok() {
					return in.Error()
				}
			default:
				return internal.ErrorInvalidJSONExcessElement("a.inner", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	var inAPointer *basictl.JsonLexer
	inA := basictl.JsonLexer{Data: rawA}
	if rawA != nil {
		inAPointer = &inA
	}
	if err := tlBuiltinTupleInt.BuiltinTupleIntReadJSON(legacyTypeNames, inAPointer, &item.A, nat_I); err != nil {
		return err
	}

	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *AInner) WriteJSONGeneral(tctx *basictl.JSONWriteContext, w []byte, nat_I uint32) (_ []byte, err error) {
	return item.WriteJSONOpt(tctx, w, nat_I)
}

func (item *AInner) WriteJSON(w []byte, nat_I uint32) (_ []byte, err error) {
	tctx := basictl.JSONWriteContext{}
	return item.WriteJSONOpt(&tctx, w, nat_I)
}
func (item *AInner) WriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte, nat_I uint32) (_ []byte, err error) {
	w = append(w, '{')
	backupIndexA := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"a":`...)
	if w, err = tlBuiltinTupleInt.BuiltinTupleIntWriteJSONOpt(tctx, w, item.A, nat_I); err != nil {
		return w, err
	}
	if (len(item.A) != 0) == false {
		w = w[:backupIndexA]
	}
	return append(w, '}'), nil
}

func (item *AInner) CalculateLayout(sizes []int, nat_I uint32) []int {
	sizePosition := len(sizes)
	sizes = append(sizes, 0)

	currentSize := 0
	lastUsedByte := 0

	// calculate layout for item.A
	currentPosition := len(sizes)
	if len(item.A) != 0 {
		sizes = tlBuiltinTupleInt.BuiltinTupleIntCalculateLayout(sizes, &item.A, nat_I)
		if sizes[currentPosition] != 0 {
			lastUsedByte = 1
			currentSize += sizes[currentPosition]
			currentSize += basictl.TL2CalculateSize(sizes[currentPosition])
		} else {
			sizes = sizes[:currentPosition+1]
		}
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

func (item *AInner) InternalWriteTL2(w []byte, sizes []int, nat_I uint32) ([]byte, []int) {
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
	// write item.A
	if len(item.A) != 0 {
		serializedSize += sizes[0]
		if sizes[0] != 0 {
			serializedSize += basictl.TL2CalculateSize(sizes[0])
			currentBlock |= (1 << 1)
			w, sizes = tlBuiltinTupleInt.BuiltinTupleIntInternalWriteTL2(w, sizes, &item.A, nat_I)
		} else {
			sizes = sizes[1:]
		}
	}
	w[currentBlockPosition] = currentBlock
	return w, sizes
}

func (item *AInner) WriteTL2(w []byte, ctx *basictl.TL2WriteContext, nat_I uint32) []byte {
	var sizes []int
	if ctx != nil {
		sizes = ctx.SizeBuffer
	}
	sizes = item.CalculateLayout(sizes[:0], nat_I)
	w, _ = item.InternalWriteTL2(w, sizes, nat_I)
	if ctx != nil {
		ctx.SizeBuffer = sizes[:0]
	}
	return w
}

func (item *AInner) InternalReadTL2(r []byte, nat_I uint32) (_ []byte, err error) {
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

	// read item.A
	if block&(1<<1) != 0 {
		if currentR, err = tlBuiltinTupleInt.BuiltinTupleIntInternalReadTL2(currentR, &item.A, nat_I); err != nil {
			return currentR, err
		}
	} else {
		item.A = item.A[:0]
	}

	return r, nil
}

func (item *AInner) ReadTL2(r []byte, ctx *basictl.TL2ReadContext, nat_I uint32) (_ []byte, err error) {
	return item.InternalReadTL2(r, nat_I)
}
