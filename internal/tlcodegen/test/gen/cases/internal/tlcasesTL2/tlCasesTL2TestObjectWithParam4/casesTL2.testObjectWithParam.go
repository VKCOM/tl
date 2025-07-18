// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlCasesTL2TestObjectWithParam4

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tl/tlBuiltinTuple4Int"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type CasesTL2TestObjectWithParam4 struct {
	X int32 // Conditional: 4.0
	Y [4]int32
}

func (CasesTL2TestObjectWithParam4) TLName() string { return "casesTL2.testObjectWithParam" }
func (CasesTL2TestObjectWithParam4) TLTag() uint32  { return 0xd0ce3a42 }

func (item *CasesTL2TestObjectWithParam4) IsSetX() bool { return 4&(1<<0) != 0 }

func (item *CasesTL2TestObjectWithParam4) Reset() {
	item.X = 0
	tlBuiltinTuple4Int.BuiltinTuple4IntReset(&item.Y)
}

func (item *CasesTL2TestObjectWithParam4) FillRandom(rg *basictl.RandGenerator) {
	if 4&(1<<0) != 0 {
		item.X = basictl.RandomInt(rg)
	} else {
		item.X = 0
	}
	tlBuiltinTuple4Int.BuiltinTuple4IntFillRandom(rg, &item.Y)
}

func (item *CasesTL2TestObjectWithParam4) Read(w []byte) (_ []byte, err error) {
	if 4&(1<<0) != 0 {
		if w, err = basictl.IntRead(w, &item.X); err != nil {
			return w, err
		}
	} else {
		item.X = 0
	}
	return tlBuiltinTuple4Int.BuiltinTuple4IntRead(w, &item.Y)
}

func (item *CasesTL2TestObjectWithParam4) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *CasesTL2TestObjectWithParam4) Write(w []byte) []byte {
	if 4&(1<<0) != 0 {
		w = basictl.IntWrite(w, item.X)
	}
	w = tlBuiltinTuple4Int.BuiltinTuple4IntWrite(w, &item.Y)
	return w
}

func (item *CasesTL2TestObjectWithParam4) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xd0ce3a42); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *CasesTL2TestObjectWithParam4) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *CasesTL2TestObjectWithParam4) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0xd0ce3a42)
	return item.Write(w)
}

func (item CasesTL2TestObjectWithParam4) String() string {
	return string(item.WriteJSON(nil))
}

func (item *CasesTL2TestObjectWithParam4) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
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
					return internal.ErrorInvalidJSONWithDuplicatingKeys("casesTL2.testObjectWithParam", "x")
				}
				if 4&(1<<0) == 0 {
					return internal.ErrorInvalidJSON("casesTL2.testObjectWithParam", "field 'x' is defined, while corresponding implicit fieldmask bit is 0")
				}
				if err := internal.Json2ReadInt32(in, &item.X); err != nil {
					return err
				}
				propXPresented = true
			case "y":
				if propYPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("casesTL2.testObjectWithParam", "y")
				}
				if err := tlBuiltinTuple4Int.BuiltinTuple4IntReadJSON(legacyTypeNames, in, &item.Y); err != nil {
					return err
				}
				propYPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("casesTL2.testObjectWithParam", key)
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
		tlBuiltinTuple4Int.BuiltinTuple4IntReset(&item.Y)
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *CasesTL2TestObjectWithParam4) WriteJSONGeneral(tctx *basictl.JSONWriteContext, w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(tctx, w), nil
}

func (item *CasesTL2TestObjectWithParam4) WriteJSON(w []byte) []byte {
	tctx := basictl.JSONWriteContext{}
	return item.WriteJSONOpt(&tctx, w)
}
func (item *CasesTL2TestObjectWithParam4) WriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte) []byte {
	w = append(w, '{')
	if 4&(1<<0) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"x":`...)
		w = basictl.JSONWriteInt32(w, item.X)
	}
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"y":`...)
	w = tlBuiltinTuple4Int.BuiltinTuple4IntWriteJSONOpt(tctx, w, &item.Y)
	return append(w, '}')
}

func (item *CasesTL2TestObjectWithParam4) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *CasesTL2TestObjectWithParam4) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("casesTL2.testObjectWithParam", err.Error())
	}
	return nil
}

func (item *CasesTL2TestObjectWithParam4) CalculateLayout(sizes []int) []int {
	sizePosition := len(sizes)
	sizes = append(sizes, 0)

	currentSize := 0
	lastUsedByte := 0

	// calculate layout for item.X
	if 4&(1<<0) != 0 {
		if item.X != 0 {

			lastUsedByte = 1
			currentSize += 4
		}
	}

	// calculate layout for item.Y
	currentPosition := len(sizes)
	sizes = tlBuiltinTuple4Int.BuiltinTuple4IntCalculateLayout(sizes, &item.Y)
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

func (item *CasesTL2TestObjectWithParam4) InternalWriteTL2(w []byte, sizes []int) ([]byte, []int) {
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
	if 4&(1<<0) != 0 {
		if item.X != 0 {
			serializedSize += 4
			if 4 != 0 {
				currentBlock |= (1 << 1)
				w = basictl.IntWrite(w, item.X)
			}
		}
	}
	// write item.Y
	serializedSize += sizes[0]
	if sizes[0] != 0 {
		serializedSize += basictl.TL2CalculateSize(sizes[0])
		currentBlock |= (1 << 2)
		w, sizes = tlBuiltinTuple4Int.BuiltinTuple4IntInternalWriteTL2(w, sizes, &item.Y)
	} else {
		sizes = sizes[1:]
	}
	w[currentBlockPosition] = currentBlock
	return w, sizes
}

func (item *CasesTL2TestObjectWithParam4) WriteTL2(w []byte, ctx *basictl.TL2WriteContext) []byte {
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

func (item *CasesTL2TestObjectWithParam4) InternalReadTL2(r []byte) (_ []byte, err error) {
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
		if 4&(1<<0) != 0 {
			if currentR, err = basictl.IntRead(currentR, &item.X); err != nil {
				return currentR, err
			}
		} else {
			return currentR, basictl.TL2Error("field mask contradiction: field item." + "X" + "is presented but depending bit is absent")
		}
	} else {
		item.X = 0
	}

	// read item.Y
	if block&(1<<2) != 0 {
		if currentR, err = tlBuiltinTuple4Int.BuiltinTuple4IntInternalReadTL2(currentR, &item.Y); err != nil {
			return currentR, err
		}
	} else {
		tlBuiltinTuple4Int.BuiltinTuple4IntReset(&item.Y)
	}

	return r, nil
}

func (item *CasesTL2TestObjectWithParam4) ReadTL2(r []byte, ctx *basictl.TL2ReadContext) (_ []byte, err error) {
	return item.InternalReadTL2(r)
}
