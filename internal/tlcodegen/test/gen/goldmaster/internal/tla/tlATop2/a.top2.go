// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlATop2

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tla/tlAMiddlePairAInnerAInnerAInnerBoxed3"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type ATop2 struct {
	N uint32
	M uint32
	C tlAMiddlePairAInnerAInnerAInnerBoxed3.AMiddlePairAInnerAInnerAInnerBoxed3
}

func (ATop2) TLName() string { return "a.top2" }
func (ATop2) TLTag() uint32  { return 0x7082d18f }

func (item *ATop2) Reset() {
	item.N = 0
	item.M = 0
	item.C.Reset()
}

func (item *ATop2) FillRandom(rg *basictl.RandGenerator) {
	item.C.FillRandom(rg, item.M, item.N, item.N)
}

func (item *ATop2) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.N); err != nil {
		return w, err
	}
	if w, err = basictl.NatRead(w, &item.M); err != nil {
		return w, err
	}
	return item.C.Read(w, item.M, item.N, item.N)
}

// This method is general version of Write, use it instead!
func (item *ATop2) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w)
}

func (item *ATop2) Write(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, item.N)
	w = basictl.NatWrite(w, item.M)
	if w, err = item.C.Write(w, item.M, item.N, item.N); err != nil {
		return w, err
	}
	return w, nil
}

func (item *ATop2) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x7082d18f); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *ATop2) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w)
}

func (item *ATop2) WriteBoxed(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, 0x7082d18f)
	return item.Write(w)
}

func (item ATop2) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func (item *ATop2) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propNPresented bool
	var propMPresented bool
	var rawC []byte

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "n":
				if propNPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("a.top2", "n")
				}
				if err := internal.Json2ReadUint32(in, &item.N); err != nil {
					return err
				}
				propNPresented = true
			case "m":
				if propMPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("a.top2", "m")
				}
				if err := internal.Json2ReadUint32(in, &item.M); err != nil {
					return err
				}
				propMPresented = true
			case "c":
				if rawC != nil {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("a.top2", "c")
				}
				rawC = in.Raw()
				if !in.Ok() {
					return in.Error()
				}
			default:
				return internal.ErrorInvalidJSONExcessElement("a.top2", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propNPresented {
		item.N = 0
	}
	if !propMPresented {
		item.M = 0
	}
	var inCPointer *basictl.JsonLexer
	inC := basictl.JsonLexer{Data: rawC}
	if rawC != nil {
		inCPointer = &inC
	}
	if err := item.C.ReadJSON(legacyTypeNames, inCPointer, item.M, item.N, item.N); err != nil {
		return err
	}

	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *ATop2) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w)
}

func (item *ATop2) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w)
}
func (item *ATop2) WriteJSONOpt(newTypeNames bool, short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	backupIndexN := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"n":`...)
	w = basictl.JSONWriteUint32(w, item.N)
	if (item.N != 0) == false {
		w = w[:backupIndexN]
	}
	backupIndexM := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"m":`...)
	w = basictl.JSONWriteUint32(w, item.M)
	if (item.M != 0) == false {
		w = w[:backupIndexM]
	}
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"c":`...)
	if w, err = item.C.WriteJSONOpt(newTypeNames, short, w, item.M, item.N, item.N); err != nil {
		return w, err
	}
	return append(w, '}'), nil
}

func (item *ATop2) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *ATop2) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("a.top2", err.Error())
	}
	return nil
}
