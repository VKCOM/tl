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

type AbUseCycle struct {
	A Cyc2MyCycle
	B []AColor
}

func (AbUseCycle) TLName() string { return "ab.useCycle" }
func (AbUseCycle) TLTag() uint32  { return 0x71687381 }

func (item *AbUseCycle) Reset() {
	item.A.Reset()
	item.B = item.B[:0]
}

func (item *AbUseCycle) FillRandom(rg *basictl.RandGenerator) {
	item.A.FillRandom(rg)
	BuiltinVectorAColorFillRandom(rg, &item.B)
}

func (item *AbUseCycle) Read(w []byte) (_ []byte, err error) {
	if w, err = item.A.Read(w); err != nil {
		return w, err
	}
	return BuiltinVectorAColorRead(w, &item.B)
}

func (item *AbUseCycle) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *AbUseCycle) Write(w []byte) []byte {
	w = item.A.Write(w)
	w = BuiltinVectorAColorWrite(w, item.B)
	return w
}

func (item *AbUseCycle) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x71687381); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *AbUseCycle) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *AbUseCycle) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x71687381)
	return item.Write(w)
}

func (item AbUseCycle) String() string {
	return string(item.WriteJSON(nil))
}

func (item *AbUseCycle) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propAPresented bool
	var propBPresented bool

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
					return ErrorInvalidJSONWithDuplicatingKeys("ab.useCycle", "a")
				}
				if err := item.A.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propAPresented = true
			case "b":
				if propBPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("ab.useCycle", "b")
				}
				if err := BuiltinVectorAColorReadJSON(legacyTypeNames, in, &item.B); err != nil {
					return err
				}
				propBPresented = true
			default:
				return ErrorInvalidJSONExcessElement("ab.useCycle", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propAPresented {
		item.A.Reset()
	}
	if !propBPresented {
		item.B = item.B[:0]
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *AbUseCycle) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *AbUseCycle) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *AbUseCycle) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"a":`...)
	w = item.A.WriteJSONOpt(newTypeNames, short, w)
	backupIndexB := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"b":`...)
	w = BuiltinVectorAColorWriteJSONOpt(newTypeNames, short, w, item.B)
	if (len(item.B) != 0) == false {
		w = w[:backupIndexB]
	}
	return append(w, '}')
}

func (item *AbUseCycle) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *AbUseCycle) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("ab.useCycle", err.Error())
	}
	return nil
}

func (item *AbUseCycle) CalculateLayout(sizes []int) []int {
	sizePosition := len(sizes)
	sizes = append(sizes, 0)

	currentSize := 0
	lastUsedByte := 0

	// calculate layout for item.A
	currentPosition := len(sizes)
	sizes = item.A.CalculateLayout(sizes)
	if sizes[currentPosition] != 0 {
		lastUsedByte = 1
		currentSize += sizes[currentPosition]
		currentSize += basictl.TL2CalculateSize(sizes[currentPosition])
	} else {
		sizes = sizes[:currentPosition+1]
	}

	// calculate layout for item.B
	currentPosition = len(sizes)
	if len(item.B) != 0 {
		sizes = BuiltinVectorAColorCalculateLayout(sizes, &item.B)
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

func (item *AbUseCycle) InternalWriteTL2(w []byte, sizes []int) ([]byte, []int) {
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
		w, sizes = item.A.InternalWriteTL2(w, sizes)
	} else {
		sizes = sizes[1:]
	}
	// write item.B
	if len(item.B) != 0 {
		serializedSize += sizes[0]
		if sizes[0] != 0 {
			serializedSize += basictl.TL2CalculateSize(sizes[0])
			currentBlock |= (1 << 2)
			w, sizes = BuiltinVectorAColorInternalWriteTL2(w, sizes, &item.B)
		} else {
			sizes = sizes[1:]
		}
	}
	w[currentBlockPosition] = currentBlock
	return w, sizes
}

func (item *AbUseCycle) WriteTL2(w []byte, sizes []int) ([]byte, []int) {
	sizes = item.CalculateLayout(sizes[:0])
	w, _ = item.InternalWriteTL2(w, sizes)
	return w, sizes[:0]
}

func (item *AbUseCycle) ReadTL2(r []byte) (_ []byte, err error) {
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
			if r, err = item.A.ReadTL2(r); err != nil {
				return r, err
			}
		} else {
			item.A.Reset()
		}

		// read item.B
		if block&(1<<2) != 0 {
			if r, err = BuiltinVectorAColorReadTL2(r, &item.B); err != nil {
				return r, err
			}
		} else {
			item.B = item.B[:0]
		}
	}

	if len(saveR) < len(r)+shift {
		r = saveR[shift:]
	}
	return r, nil
}
