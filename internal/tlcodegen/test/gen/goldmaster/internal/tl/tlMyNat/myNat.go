// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlMyNat

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type MyNat struct {
	FieldsMask uint32
	A          *MyNat // Conditional: item.FieldsMask.0
}

func (MyNat) TLName() string { return "myNat" }
func (MyNat) TLTag() uint32  { return 0xc60c1b41 }

func (item *MyNat) SetA(v MyNat) {
	if item.A == nil {
		var value MyNat
		item.A = &value
	}
	*item.A = v
	item.FieldsMask |= 1 << 0
}
func (item *MyNat) ClearA() {
	if item.A != nil {
		item.A.Reset()
	}
	item.FieldsMask &^= 1 << 0
}
func (item *MyNat) IsSetA() bool { return item.FieldsMask&(1<<0) != 0 }

func (item *MyNat) Reset() {
	item.FieldsMask = 0
	if item.A != nil {
		item.A.Reset()
	}
}

func (item *MyNat) FillRandom(rg *basictl.RandGenerator) {
	var maskFieldsMask uint32
	maskFieldsMask = basictl.RandomUint(rg)
	item.FieldsMask = 0
	if maskFieldsMask&(1<<0) != 0 {
		item.FieldsMask |= (1 << 0)
	}
	if item.FieldsMask&(1<<0) != 0 {
		rg.IncreaseDepth()
		if item.A == nil {
			var value MyNat
			item.A = &value
		}
		item.A.FillRandom(rg)
		rg.DecreaseDepth()
	} else {
		if item.A != nil {
			item.A.Reset()
		}
	}
}

func (item *MyNat) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.FieldsMask); err != nil {
		return w, err
	}
	if item.FieldsMask&(1<<0) != 0 {
		if item.A == nil {
			var value MyNat
			item.A = &value
		}
		if w, err = item.A.Read(w); err != nil {
			return w, err
		}
	} else {
		if item.A != nil {
			item.A.Reset()
		}
	}
	return w, nil
}

func (item *MyNat) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *MyNat) Write(w []byte) []byte {
	w = basictl.NatWrite(w, item.FieldsMask)
	if item.FieldsMask&(1<<0) != 0 {
		if item.A == nil {
			var tmpValue MyNat
			w = (&tmpValue).Write(w)
		} else {
			w = item.A.Write(w)
		}
	}
	return w
}

func (item *MyNat) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xc60c1b41); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *MyNat) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *MyNat) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0xc60c1b41)
	return item.Write(w)
}

func (item MyNat) String() string {
	return string(item.WriteJSON(nil))
}

func (item *MyNat) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propFieldsMaskPresented bool
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
			case "fields_mask":
				if propFieldsMaskPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("myNat", "fields_mask")
				}
				if err := internal.Json2ReadUint32(in, &item.FieldsMask); err != nil {
					return err
				}
				propFieldsMaskPresented = true
			case "a":
				if propAPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("myNat", "a")
				}
				if item.A == nil {
					var value MyNat
					item.A = &value
				}
				if err := item.A.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propAPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("myNat", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propFieldsMaskPresented {
		item.FieldsMask = 0
	}
	if !propAPresented {
		if item.A != nil {
			item.A.Reset()
		}
	}
	if propAPresented {
		item.FieldsMask |= 1 << 0
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *MyNat) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *MyNat) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *MyNat) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexFieldsMask := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"fields_mask":`...)
	w = basictl.JSONWriteUint32(w, item.FieldsMask)
	if (item.FieldsMask != 0) == false {
		w = w[:backupIndexFieldsMask]
	}
	if item.FieldsMask&(1<<0) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"a":`...)
		w = item.A.WriteJSONOpt(newTypeNames, short, w)
	}
	return append(w, '}')
}

func (item *MyNat) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *MyNat) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("myNat", err.Error())
	}
	return nil
}

func (item *MyNat) CalculateLayout(sizes []int) []int {
	sizePosition := len(sizes)
	sizes = append(sizes, 0)

	currentSize := 0
	lastUsedByte := 0

	// calculate layout for item.FieldsMask
	if item.FieldsMask != 0 {

		lastUsedByte = 1
		currentSize += 4
	}

	// calculate layout for item.A
	currentPosition := len(sizes)
	if item.FieldsMask&(1<<0) != 0 {
		sizes = (*item.A).CalculateLayout(sizes)
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

func (item *MyNat) InternalWriteTL2(w []byte, sizes []int) ([]byte, []int) {
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
	// write item.FieldsMask
	if item.FieldsMask != 0 {
		serializedSize += 4
		if 4 != 0 {
			currentBlock |= (1 << 1)
			w = basictl.NatWrite(w, item.FieldsMask)
		}
	}
	// write item.A
	if item.FieldsMask&(1<<0) != 0 {
		serializedSize += sizes[0]
		if sizes[0] != 0 {
			serializedSize += basictl.TL2CalculateSize(sizes[0])
			currentBlock |= (1 << 2)
			w, sizes = item.A.InternalWriteTL2(w, sizes)
		} else {
			sizes = sizes[1:]
		}
	}
	w[currentBlockPosition] = currentBlock
	return w, sizes
}

func (item *MyNat) WriteTL2(w []byte, sizes []int) ([]byte, []int) {
	sizes = item.CalculateLayout(sizes[:0])
	w, _ = item.InternalWriteTL2(w, sizes)
	return w, sizes[:0]
}

func (item *MyNat) ReadTL2(r []byte) (_ []byte, err error) {
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

		// read item.FieldsMask
		if block&(1<<1) != 0 {
			if r, err = basictl.NatRead(r, &item.FieldsMask); err != nil {
				return r, err
			}
		} else {
			item.FieldsMask = 0
		}

		// read item.A
		if block&(1<<2) != 0 {
			if item.A == nil {
				var newValue MyNat
				item.A = &newValue
			}
			if item.FieldsMask&(1<<0) != 0 {
				if r, err = item.A.ReadTL2(r); err != nil {
					return r, err
				}
			} else {
				return r, basictl.TL2Error("field mask contradiction: field item." + "A" + "is presented but depending bit is absent")
			}
		} else {
			if item.A == nil {
				var newValue MyNat
				item.A = &newValue
			}
			item.A.Reset()
		}
	}

	if len(saveR) < len(r)+shift {
		r = saveR[shift:]
	}
	return r, nil
}
