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

type Replace3 struct {
	A [3]int32
}

func (Replace3) TLName() string { return "replace3" }
func (Replace3) TLTag() uint32  { return 0x51e324e4 }

func (item *Replace3) Reset() {
	BuiltinTuple3IntReset(&item.A)
}

func (item *Replace3) FillRandom(rg *basictl.RandGenerator) {
	BuiltinTuple3IntFillRandom(rg, &item.A)
}

func (item *Replace3) Read(w []byte) (_ []byte, err error) {
	return BuiltinTuple3IntRead(w, &item.A)
}

func (item *Replace3) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *Replace3) Write(w []byte) []byte {
	w = BuiltinTuple3IntWrite(w, &item.A)
	return w
}

func (item *Replace3) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x51e324e4); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *Replace3) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *Replace3) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x51e324e4)
	return item.Write(w)
}

func (item Replace3) String() string {
	return string(item.WriteJSON(nil))
}

func (item *Replace3) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propAPresented bool

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
				if propAPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("replace3", "a")
				}
				if err := BuiltinTuple3IntReadJSON(legacyTypeNames, in, &item.A); err != nil {
					return err
				}
				propAPresented = true
			default:
				return ErrorInvalidJSONExcessElement("replace3", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propAPresented {
		BuiltinTuple3IntReset(&item.A)
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *Replace3) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *Replace3) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *Replace3) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"a":`...)
	w = BuiltinTuple3IntWriteJSONOpt(newTypeNames, short, w, &item.A)
	return append(w, '}')
}

func (item *Replace3) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *Replace3) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("replace3", err.Error())
	}
	return nil
}

func (item *Replace3) CalculateLayout(sizes []int) []int {
	sizePosition := len(sizes)
	sizes = append(sizes, 0)

	currentSize := 0
	lastUsedByte := 0

	// calculate layout for item.A
	currentPosition := len(sizes)
	sizes = BuiltinTuple3IntCalculateLayout(sizes, &item.A)
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

func (item *Replace3) InternalWriteTL2(w []byte, sizes []int) ([]byte, []int) {
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
	serializedSize += sizes[0]
	if sizes[0] != 0 {
		serializedSize += basictl.TL2CalculateSize(sizes[0])
		currentBlock |= (1 << 1)
		w, sizes = BuiltinTuple3IntInternalWriteTL2(w, sizes, &item.A)
	} else {
		sizes = sizes[1:]
	}
	w[currentBlockPosition] = currentBlock
	return w, sizes
}

func (item *Replace3) WriteTL2(w []byte, sizes []int) ([]byte, []int) {
	sizes = item.CalculateLayout(sizes[:0])
	w, _ = item.InternalWriteTL2(w, sizes)
	return w, sizes[:0]
}

func (item *Replace3) ReadTL2(r []byte) (_ []byte, err error) {
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

		// read item.A
		if block&(1<<1) != 0 {
			if r, err = BuiltinTuple3IntReadTL2(r, &item.A); err != nil {
				return r, err
			}
		} else {
			BuiltinTuple3IntReset(&item.A)
		}
	}

	if len(saveR) < len(r)+shift {
		r = saveR[shift:]
	}
	return r, nil
}
