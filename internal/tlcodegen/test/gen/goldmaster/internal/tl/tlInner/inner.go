// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlInner

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type Inner struct {
	A int32
}

func (Inner) TLName() string { return "inner" }
func (Inner) TLTag() uint32  { return 0x3b53db83 }

func (item *Inner) Reset() {
	item.A = 0
}

func (item *Inner) FillRandom(rg *basictl.RandGenerator, nat_X uint32) {
	item.A = basictl.RandomInt(rg)
}

func (item *Inner) Read(w []byte, nat_X uint32) (_ []byte, err error) {
	return basictl.IntRead(w, &item.A)
}

func (item *Inner) WriteGeneral(w []byte, nat_X uint32) (_ []byte, err error) {
	return item.Write(w, nat_X), nil
}

func (item *Inner) Write(w []byte, nat_X uint32) []byte {
	w = basictl.IntWrite(w, item.A)
	return w
}

func (item *Inner) ReadBoxed(w []byte, nat_X uint32) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x3b53db83); err != nil {
		return w, err
	}
	return item.Read(w, nat_X)
}

func (item *Inner) WriteBoxedGeneral(w []byte, nat_X uint32) (_ []byte, err error) {
	return item.WriteBoxed(w, nat_X), nil
}

func (item *Inner) WriteBoxed(w []byte, nat_X uint32) []byte {
	w = basictl.NatWrite(w, 0x3b53db83)
	return item.Write(w, nat_X)
}

func (item *Inner) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, nat_X uint32) error {
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
					return internal.ErrorInvalidJSONWithDuplicatingKeys("inner", "a")
				}
				if err := internal.Json2ReadInt32(in, &item.A); err != nil {
					return err
				}
				propAPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("inner", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propAPresented {
		item.A = 0
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *Inner) WriteJSONGeneral(w []byte, nat_X uint32) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w, nat_X), nil
}

func (item *Inner) WriteJSON(w []byte, nat_X uint32) []byte {
	return item.WriteJSONOpt(true, false, w, nat_X)
}
func (item *Inner) WriteJSONOpt(newTypeNames bool, short bool, w []byte, nat_X uint32) []byte {
	w = append(w, '{')
	backupIndexA := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"a":`...)
	w = basictl.JSONWriteInt32(w, item.A)
	if (item.A != 0) == false {
		w = w[:backupIndexA]
	}
	return append(w, '}')
}

func (item *Inner) CalculateLayout(sizes []int, nat_X uint32) []int {
	sizePosition := len(sizes)
	sizes = append(sizes, 0)

	currentSize := 0
	lastUsedByte := 0

	// calculate layout for item.A
	if item.A != 0 {

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

func (item *Inner) InternalWriteTL2(w []byte, sizes []int, nat_X uint32) ([]byte, []int) {
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
	if item.A != 0 {
		serializedSize += 4
		if 4 != 0 {
			currentBlock |= (1 << 1)
			w = basictl.IntWrite(w, item.A)
		}
	}
	w[currentBlockPosition] = currentBlock
	return w, sizes
}

func (item *Inner) ReadTL2(r []byte, nat_X uint32) (_ []byte, err error) {
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
			if r, err = basictl.IntRead(r, &item.A); err != nil {
				return r, err
			}
		} else {
			item.A = 0
		}
	}

	if len(saveR) < len(r)+shift {
		r = saveR[shift:]
	}
	return r, nil
}
