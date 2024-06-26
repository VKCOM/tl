// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlAMiddlePairAInnerAInnerAInnerBoxed3

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlPairPairAInnerAInnerAInnerBoxed3"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tla/tlAInner"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type AMiddlePairAInnerAInnerAInnerBoxed3 struct {
	A tlAInner.AInner
	B tlPairPairAInnerAInnerAInnerBoxed3.PairPairAInnerAInnerAInnerBoxed3
}

func (AMiddlePairAInnerAInnerAInnerBoxed3) TLName() string { return "a.middle" }
func (AMiddlePairAInnerAInnerAInnerBoxed3) TLTag() uint32  { return 0xaf5e2b14 }

func (item *AMiddlePairAInnerAInnerAInnerBoxed3) Reset() {
	item.A.Reset()
	item.B.Reset()
}

func (item *AMiddlePairAInnerAInnerAInnerBoxed3) FillRandom(rg *basictl.RandGenerator, nat_W uint32, nat_PXI uint32, nat_PYI uint32) {
	item.A.FillRandom(rg, nat_W)
	item.B.FillRandom(rg, nat_PXI, nat_PYI)
}

func (item *AMiddlePairAInnerAInnerAInnerBoxed3) Read(w []byte, nat_W uint32, nat_PXI uint32, nat_PYI uint32) (_ []byte, err error) {
	if w, err = item.A.Read(w, nat_W); err != nil {
		return w, err
	}
	return item.B.Read(w, nat_PXI, nat_PYI)
}

// This method is general version of Write, use it instead!
func (item *AMiddlePairAInnerAInnerAInnerBoxed3) WriteGeneral(w []byte, nat_W uint32, nat_PXI uint32, nat_PYI uint32) (_ []byte, err error) {
	return item.Write(w, nat_W, nat_PXI, nat_PYI)
}

func (item *AMiddlePairAInnerAInnerAInnerBoxed3) Write(w []byte, nat_W uint32, nat_PXI uint32, nat_PYI uint32) (_ []byte, err error) {
	if w, err = item.A.Write(w, nat_W); err != nil {
		return w, err
	}
	if w, err = item.B.Write(w, nat_PXI, nat_PYI); err != nil {
		return w, err
	}
	return w, nil
}

func (item *AMiddlePairAInnerAInnerAInnerBoxed3) ReadBoxed(w []byte, nat_W uint32, nat_PXI uint32, nat_PYI uint32) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xaf5e2b14); err != nil {
		return w, err
	}
	return item.Read(w, nat_W, nat_PXI, nat_PYI)
}

// This method is general version of WriteBoxed, use it instead!
func (item *AMiddlePairAInnerAInnerAInnerBoxed3) WriteBoxedGeneral(w []byte, nat_W uint32, nat_PXI uint32, nat_PYI uint32) (_ []byte, err error) {
	return item.WriteBoxed(w, nat_W, nat_PXI, nat_PYI)
}

func (item *AMiddlePairAInnerAInnerAInnerBoxed3) WriteBoxed(w []byte, nat_W uint32, nat_PXI uint32, nat_PYI uint32) (_ []byte, err error) {
	w = basictl.NatWrite(w, 0xaf5e2b14)
	return item.Write(w, nat_W, nat_PXI, nat_PYI)
}

func (item *AMiddlePairAInnerAInnerAInnerBoxed3) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, nat_W uint32, nat_PXI uint32, nat_PYI uint32) error {
	var rawA []byte
	var rawB []byte

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
					return internal.ErrorInvalidJSONWithDuplicatingKeys("a.middle", "a")
				}
				rawA = in.Raw()
				if !in.Ok() {
					return in.Error()
				}
			case "b":
				if rawB != nil {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("a.middle", "b")
				}
				rawB = in.Raw()
				if !in.Ok() {
					return in.Error()
				}
			default:
				return internal.ErrorInvalidJSONExcessElement("a.middle", key)
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
	if err := item.A.ReadJSON(legacyTypeNames, inAPointer, nat_W); err != nil {
		return err
	}

	var inBPointer *basictl.JsonLexer
	inB := basictl.JsonLexer{Data: rawB}
	if rawB != nil {
		inBPointer = &inB
	}
	if err := item.B.ReadJSON(legacyTypeNames, inBPointer, nat_PXI, nat_PYI); err != nil {
		return err
	}

	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *AMiddlePairAInnerAInnerAInnerBoxed3) WriteJSONGeneral(w []byte, nat_W uint32, nat_PXI uint32, nat_PYI uint32) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w, nat_W, nat_PXI, nat_PYI)
}

func (item *AMiddlePairAInnerAInnerAInnerBoxed3) WriteJSON(w []byte, nat_W uint32, nat_PXI uint32, nat_PYI uint32) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w, nat_W, nat_PXI, nat_PYI)
}
func (item *AMiddlePairAInnerAInnerAInnerBoxed3) WriteJSONOpt(newTypeNames bool, short bool, w []byte, nat_W uint32, nat_PXI uint32, nat_PYI uint32) (_ []byte, err error) {
	w = append(w, '{')
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"a":`...)
	if w, err = item.A.WriteJSONOpt(newTypeNames, short, w, nat_W); err != nil {
		return w, err
	}
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"b":`...)
	if w, err = item.B.WriteJSONOpt(newTypeNames, short, w, nat_PXI, nat_PYI); err != nil {
		return w, err
	}
	return append(w, '}'), nil
}
