// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlPairAInnerAInner

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tla/tlAInner"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type PairAInnerAInner struct {
	A tlAInner.AInner
	B tlAInner.AInner
}

func (PairAInnerAInner) TLName() string { return "pair" }
func (PairAInnerAInner) TLTag() uint32  { return 0xf3c47ab }

func (item *PairAInnerAInner) Reset() {
	item.A.Reset()
	item.B.Reset()
}

func (item *PairAInnerAInner) FillRandom(rg *basictl.RandGenerator, nat_X uint32, nat_Y uint32) {
	item.A.FillRandom(rg, nat_X)
	item.B.FillRandom(rg, nat_Y)
}

func (item *PairAInnerAInner) Read(w []byte, nat_X uint32, nat_Y uint32) (_ []byte, err error) {
	if w, err = item.A.Read(w, nat_X); err != nil {
		return w, err
	}
	return item.B.Read(w, nat_Y)
}

// This method is general version of Write, use it instead!
func (item *PairAInnerAInner) WriteGeneral(w []byte, nat_X uint32, nat_Y uint32) (_ []byte, err error) {
	return item.Write(w, nat_X, nat_Y)
}

func (item *PairAInnerAInner) Write(w []byte, nat_X uint32, nat_Y uint32) (_ []byte, err error) {
	if w, err = item.A.Write(w, nat_X); err != nil {
		return w, err
	}
	if w, err = item.B.Write(w, nat_Y); err != nil {
		return w, err
	}
	return w, nil
}

func (item *PairAInnerAInner) ReadBoxed(w []byte, nat_X uint32, nat_Y uint32) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xf3c47ab); err != nil {
		return w, err
	}
	return item.Read(w, nat_X, nat_Y)
}

// This method is general version of WriteBoxed, use it instead!
func (item *PairAInnerAInner) WriteBoxedGeneral(w []byte, nat_X uint32, nat_Y uint32) (_ []byte, err error) {
	return item.WriteBoxed(w, nat_X, nat_Y)
}

func (item *PairAInnerAInner) WriteBoxed(w []byte, nat_X uint32, nat_Y uint32) (_ []byte, err error) {
	w = basictl.NatWrite(w, 0xf3c47ab)
	return item.Write(w, nat_X, nat_Y)
}

func (item *PairAInnerAInner) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, nat_X uint32, nat_Y uint32) error {
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
					return internal.ErrorInvalidJSONWithDuplicatingKeys("pair", "a")
				}
				rawA = in.Raw()
				if !in.Ok() {
					return in.Error()
				}
			case "b":
				if rawB != nil {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("pair", "b")
				}
				rawB = in.Raw()
				if !in.Ok() {
					return in.Error()
				}
			default:
				return internal.ErrorInvalidJSONExcessElement("pair", key)
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
	if err := item.A.ReadJSON(legacyTypeNames, inAPointer, nat_X); err != nil {
		return err
	}

	var inBPointer *basictl.JsonLexer
	inB := basictl.JsonLexer{Data: rawB}
	if rawB != nil {
		inBPointer = &inB
	}
	if err := item.B.ReadJSON(legacyTypeNames, inBPointer, nat_Y); err != nil {
		return err
	}

	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *PairAInnerAInner) WriteJSONGeneral(w []byte, nat_X uint32, nat_Y uint32) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w, nat_X, nat_Y)
}

func (item *PairAInnerAInner) WriteJSON(w []byte, nat_X uint32, nat_Y uint32) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w, nat_X, nat_Y)
}
func (item *PairAInnerAInner) WriteJSONOpt(newTypeNames bool, short bool, w []byte, nat_X uint32, nat_Y uint32) (_ []byte, err error) {
	w = append(w, '{')
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"a":`...)
	if w, err = item.A.WriteJSONOpt(newTypeNames, short, w, nat_X); err != nil {
		return w, err
	}
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"b":`...)
	if w, err = item.B.WriteJSONOpt(newTypeNames, short, w, nat_Y); err != nil {
		return w, err
	}
	return append(w, '}'), nil
}
