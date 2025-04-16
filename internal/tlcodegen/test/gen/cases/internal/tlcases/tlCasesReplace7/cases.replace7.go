// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlCasesReplace7

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tl/tlBuiltinTupleTupleInt"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type CasesReplace7 struct {
	N uint32
	M uint32
	A [][]int32
}

func (CasesReplace7) TLName() string { return "cases.replace7" }
func (CasesReplace7) TLTag() uint32  { return 0x6ccce4be }

func (item *CasesReplace7) Reset() {
	item.N = 0
	item.M = 0
	item.A = item.A[:0]
}

func (item *CasesReplace7) FillRandom(rg *basictl.RandGenerator) {
	item.N = basictl.RandomUint(rg)
	item.N = rg.LimitValue(item.N)
	item.M = basictl.RandomUint(rg)
	item.M = rg.LimitValue(item.M)
	tlBuiltinTupleTupleInt.BuiltinTupleTupleIntFillRandom(rg, &item.A, item.N, item.M)
}

func (item *CasesReplace7) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.N); err != nil {
		return w, err
	}
	if w, err = basictl.NatRead(w, &item.M); err != nil {
		return w, err
	}
	return tlBuiltinTupleTupleInt.BuiltinTupleTupleIntRead(w, &item.A, item.N, item.M)
}

// This method is general version of Write, use it instead!
func (item *CasesReplace7) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w)
}

func (item *CasesReplace7) Write(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, item.N)
	w = basictl.NatWrite(w, item.M)
	if w, err = tlBuiltinTupleTupleInt.BuiltinTupleTupleIntWrite(w, item.A, item.N, item.M); err != nil {
		return w, err
	}
	return w, nil
}

func (item *CasesReplace7) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x6ccce4be); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *CasesReplace7) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w)
}

func (item *CasesReplace7) WriteBoxed(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, 0x6ccce4be)
	return item.Write(w)
}

func (item CasesReplace7) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func (item *CasesReplace7) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propNPresented bool
	var propMPresented bool
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
			case "n":
				if propNPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("cases.replace7", "n")
				}
				if err := internal.Json2ReadUint32(in, &item.N); err != nil {
					return err
				}
				propNPresented = true
			case "m":
				if propMPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("cases.replace7", "m")
				}
				if err := internal.Json2ReadUint32(in, &item.M); err != nil {
					return err
				}
				propMPresented = true
			case "a":
				if rawA != nil {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("cases.replace7", "a")
				}
				rawA = in.Raw()
				if !in.Ok() {
					return in.Error()
				}
			default:
				return internal.ErrorInvalidJSONExcessElement("cases.replace7", key)
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
	var inAPointer *basictl.JsonLexer
	inA := basictl.JsonLexer{Data: rawA}
	if rawA != nil {
		inAPointer = &inA
	}
	if err := tlBuiltinTupleTupleInt.BuiltinTupleTupleIntReadJSON(legacyTypeNames, inAPointer, &item.A, item.N, item.M); err != nil {
		return err
	}

	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *CasesReplace7) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w)
}

func (item *CasesReplace7) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w)
}
func (item *CasesReplace7) WriteJSONOpt(newTypeNames bool, short bool, w []byte) (_ []byte, err error) {
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
	backupIndexA := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"a":`...)
	if w, err = tlBuiltinTupleTupleInt.BuiltinTupleTupleIntWriteJSONOpt(newTypeNames, short, w, item.A, item.N, item.M); err != nil {
		return w, err
	}
	if (len(item.A) != 0) == false {
		w = w[:backupIndexA]
	}
	return append(w, '}'), nil
}

func (item *CasesReplace7) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *CasesReplace7) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("cases.replace7", err.Error())
	}
	return nil
}
