// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlCasesReplace7plus

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tl/tlBuiltinTupleTupleInt"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type CasesReplace7plus struct {
	N uint32
	M uint32
	A [][]int32 // Conditional: item.N.0
}

func (CasesReplace7plus) TLName() string { return "cases.replace7plus" }
func (CasesReplace7plus) TLTag() uint32  { return 0x197858f5 }

func (item *CasesReplace7plus) SetA(v [][]int32) {
	item.A = v
	item.N |= 1 << 0
}
func (item *CasesReplace7plus) ClearA() {
	item.A = item.A[:0]
	item.N &^= 1 << 0
}
func (item *CasesReplace7plus) IsSetA() bool { return item.N&(1<<0) != 0 }

func (item *CasesReplace7plus) Reset() {
	item.N = 0
	item.M = 0
	item.A = item.A[:0]
}

func (item *CasesReplace7plus) FillRandom(rg *basictl.RandGenerator) {
	var maskN uint32
	maskN = basictl.RandomUint(rg)
	maskN = rg.LimitValue(maskN)
	item.N = 0
	if maskN&(1<<0) != 0 {
		item.N |= (1 << 0)
	}
	item.M = basictl.RandomUint(rg)
	item.M = rg.LimitValue(item.M)
	if item.N&(1<<0) != 0 {
		tlBuiltinTupleTupleInt.BuiltinTupleTupleIntFillRandom(rg, &item.A, item.N, item.M)
	} else {
		item.A = item.A[:0]
	}
}

func (item *CasesReplace7plus) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.N); err != nil {
		return w, err
	}
	if w, err = basictl.NatRead(w, &item.M); err != nil {
		return w, err
	}
	if item.N&(1<<0) != 0 {
		if w, err = tlBuiltinTupleTupleInt.BuiltinTupleTupleIntRead(w, &item.A, item.N, item.M); err != nil {
			return w, err
		}
	} else {
		item.A = item.A[:0]
	}
	return w, nil
}

// This method is general version of Write, use it instead!
func (item *CasesReplace7plus) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w)
}

func (item *CasesReplace7plus) Write(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, item.N)
	w = basictl.NatWrite(w, item.M)
	if item.N&(1<<0) != 0 {
		if w, err = tlBuiltinTupleTupleInt.BuiltinTupleTupleIntWrite(w, item.A, item.N, item.M); err != nil {
			return w, err
		}
	}
	return w, nil
}

func (item *CasesReplace7plus) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x197858f5); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *CasesReplace7plus) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w)
}

func (item *CasesReplace7plus) WriteBoxed(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, 0x197858f5)
	return item.Write(w)
}

func (item CasesReplace7plus) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func (item *CasesReplace7plus) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
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
					return internal.ErrorInvalidJSONWithDuplicatingKeys("cases.replace7plus", "n")
				}
				if err := internal.Json2ReadUint32(in, &item.N); err != nil {
					return err
				}
				propNPresented = true
			case "m":
				if propMPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("cases.replace7plus", "m")
				}
				if err := internal.Json2ReadUint32(in, &item.M); err != nil {
					return err
				}
				propMPresented = true
			case "a":
				if rawA != nil {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("cases.replace7plus", "a")
				}
				rawA = in.Raw()
				if !in.Ok() {
					return in.Error()
				}
			default:
				return internal.ErrorInvalidJSONExcessElement("cases.replace7plus", key)
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
	if rawA != nil {
		item.N |= 1 << 0
	}
	if item.N&(1<<0) == 0 {
		item.A = item.A[:0]
	} else {
		var inAPointer *basictl.JsonLexer
		inA := basictl.JsonLexer{Data: rawA}
		if rawA != nil {
			inAPointer = &inA
		}
		if err := tlBuiltinTupleTupleInt.BuiltinTupleTupleIntReadJSON(legacyTypeNames, inAPointer, &item.A, item.N, item.M); err != nil {
			return err
		}

	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *CasesReplace7plus) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w)
}

func (item *CasesReplace7plus) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w)
}
func (item *CasesReplace7plus) WriteJSONOpt(newTypeNames bool, short bool, w []byte) (_ []byte, err error) {
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
	if item.N&(1<<0) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"a":`...)
		if w, err = tlBuiltinTupleTupleInt.BuiltinTupleTupleIntWriteJSONOpt(newTypeNames, short, w, item.A, item.N, item.M); err != nil {
			return w, err
		}
	}
	return append(w, '}'), nil
}

func (item *CasesReplace7plus) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *CasesReplace7plus) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("cases.replace7plus", err.Error())
	}
	return nil
}
