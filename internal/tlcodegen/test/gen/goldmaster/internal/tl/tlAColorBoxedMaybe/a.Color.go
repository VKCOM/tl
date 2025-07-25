// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlAColorBoxedMaybe

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tla/tlAColor"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type AColorBoxedMaybe struct {
	Value tlAColor.AColor // not deterministic if !Ok
	Ok    bool
}

func (item *AColorBoxedMaybe) Reset() {
	item.Ok = false
}
func (item *AColorBoxedMaybe) FillRandom(rg *basictl.RandGenerator) {
	if basictl.RandomUint(rg)&1 == 1 {
		item.Ok = true
		item.Value.FillRandom(rg)
	} else {
		item.Ok = false
	}
}

func (item *AColorBoxedMaybe) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.ReadBool(w, &item.Ok, 0x27930a7b, 0x3f9c8ef8); err != nil {
		return w, err
	}
	if item.Ok {
		return item.Value.ReadBoxed(w)
	}
	return w, nil
}

func (item *AColorBoxedMaybe) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *AColorBoxedMaybe) WriteBoxed(w []byte) []byte {
	if item.Ok {
		w = basictl.NatWrite(w, 0x3f9c8ef8)
		return item.Value.WriteBoxed(w)
	}
	return basictl.NatWrite(w, 0x27930a7b)
}

func (item *AColorBoxedMaybe) CalculateLayout(sizes []int) []int {
	sizePosition := len(sizes)
	sizes = append(sizes, 0)
	if item.Ok {
		sizes[sizePosition] += 1
		sizes[sizePosition] += basictl.TL2CalculateSize(1)
		currentPosition := len(sizes)
		sizes = item.Value.CalculateLayout(sizes)
		if sizes[currentPosition] != 0 {
			sizes[sizePosition] += sizes[currentPosition]
			sizes[sizePosition] += basictl.TL2CalculateSize(sizes[currentPosition])
		}
	}
	return sizes
}

func (item *AColorBoxedMaybe) InternalWriteTL2(w []byte, sizes []int) ([]byte, []int) {
	currentSize := sizes[0]
	sizes = sizes[1:]

	w = basictl.TL2WriteSize(w, currentSize)
	if currentSize == 0 {
		return w, sizes
	}

	if item.Ok {
		currentPosition := len(w)
		w = append(w, 1)
		w = basictl.TL2WriteSize(w, 1)
		if sizes[0] != 0 {
			w[currentPosition] |= (1 << 1)
			w, sizes = item.Value.InternalWriteTL2(w, sizes)
		} else {
			sizes = sizes[1:]
		}
	}
	return w, sizes
}

func (item *AColorBoxedMaybe) InternalReadTL2(r []byte) (_ []byte, err error) {
	saveR := r
	currentSize := 0
	if r, currentSize, err = basictl.TL2ParseSize(r); err != nil {
		return r, err
	}
	shift := currentSize + basictl.TL2CalculateSize(currentSize)

	if currentSize == 0 {
		item.Ok = false
	} else {
		var block byte
		if r, err = basictl.ByteReadTL2(r, &block); err != nil {
			return r, err
		}
		if block&1 == 0 {
			return r, basictl.TL2Error("must have constructor bytes")
		}
		var index int
		if r, index, err = basictl.TL2ParseSize(r); err != nil {
			return r, err
		}
		if index != 1 {
			return r, basictl.TL2Error("expected 1")
		}
		item.Ok = true
		if block&(1<<1) != 0 {
			if r, err = item.Value.InternalReadTL2(r); err != nil {
				return r, err
			}
		} else {
			item.Value.Reset()
		}
	}
	if len(saveR) < len(r)+shift {
		r = saveR[shift:]
	}
	return r, nil
}

func (item *AColorBoxedMaybe) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	_ok, _jvalue, err := internal.Json2ReadMaybe("Maybe", in)
	if err != nil {
		return err
	}
	item.Ok = _ok
	if _ok {
		var in2Pointer *basictl.JsonLexer
		if _jvalue != nil {
			in2 := basictl.JsonLexer{Data: _jvalue}
			in2Pointer = &in2
		}
		if err := item.Value.ReadJSON(legacyTypeNames, in2Pointer); err != nil {
			return err
		}
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *AColorBoxedMaybe) WriteJSONGeneral(tctx *basictl.JSONWriteContext, w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(tctx, w), nil
}

func (item *AColorBoxedMaybe) WriteJSON(w []byte) []byte {
	tctx := basictl.JSONWriteContext{}
	return item.WriteJSONOpt(&tctx, w)
}
func (item *AColorBoxedMaybe) WriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte) []byte {
	if !item.Ok {
		return append(w, "{}"...)
	}
	w = append(w, `{"ok":true`...)
	w = append(w, `,"value":`...)
	w = item.Value.WriteJSONOpt(tctx, w)
	return append(w, '}')
}

func (item AColorBoxedMaybe) String() string {
	return string(item.WriteJSON(nil))
}
