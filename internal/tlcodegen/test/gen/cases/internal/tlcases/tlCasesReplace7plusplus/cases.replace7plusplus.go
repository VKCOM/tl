// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlCasesReplace7plusplus

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tl/tlBuiltinTupleTupleInt"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type CasesReplace7plusplus struct {
	N uint32
	M uint32
	A [][]int32 // Conditional: item.N.0
}

func (CasesReplace7plusplus) TLName() string { return "cases.replace7plusplus" }
func (CasesReplace7plusplus) TLTag() uint32  { return 0xabc39b68 }

func (item *CasesReplace7plusplus) SetA(v [][]int32) {
	item.A = v
	item.N |= 1 << 0
}
func (item *CasesReplace7plusplus) ClearA() {
	item.A = item.A[:0]
	item.N &^= 1 << 0
}
func (item CasesReplace7plusplus) IsSetA() bool { return item.N&(1<<0) != 0 }

func (item *CasesReplace7plusplus) Reset() {
	item.N = 0
	item.M = 0
	item.A = item.A[:0]
}

func (item *CasesReplace7plusplus) FillRandom(rg *basictl.RandGenerator) {
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

func (item *CasesReplace7plusplus) Read(w []byte) (_ []byte, err error) {
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
func (item *CasesReplace7plusplus) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w)
}

func (item *CasesReplace7plusplus) Write(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, item.N)
	w = basictl.NatWrite(w, item.M)
	if item.N&(1<<0) != 0 {
		if w, err = tlBuiltinTupleTupleInt.BuiltinTupleTupleIntWrite(w, item.A, item.N, item.M); err != nil {
			return w, err
		}
	}
	return w, nil
}

func (item *CasesReplace7plusplus) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xabc39b68); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *CasesReplace7plusplus) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w)
}

func (item *CasesReplace7plusplus) WriteBoxed(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, 0xabc39b68)
	return item.Write(w)
}

func (item CasesReplace7plusplus) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func (item *CasesReplace7plusplus) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
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
			case "N":
				if propNPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("cases.replace7plusplus", "N")
				}
				if err := internal.Json2ReadUint32(in, &item.N); err != nil {
					return err
				}
				propNPresented = true
			case "M":
				if propMPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("cases.replace7plusplus", "M")
				}
				if err := internal.Json2ReadUint32(in, &item.M); err != nil {
					return err
				}
				propMPresented = true
			case "A":
				if rawA != nil {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("cases.replace7plusplus", "A")
				}
				rawA = in.Raw()
				if !in.Ok() {
					return in.Error()
				}
			default:
				return internal.ErrorInvalidJSONExcessElement("cases.replace7plusplus", key)
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
func (item *CasesReplace7plusplus) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w)
}

func (item *CasesReplace7plusplus) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w)
}
func (item *CasesReplace7plusplus) WriteJSONOpt(newTypeNames bool, short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	backupIndexN := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"N":`...)
	w = basictl.JSONWriteUint32(w, item.N)
	if (item.N != 0) == false {
		w = w[:backupIndexN]
	}
	backupIndexM := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"M":`...)
	w = basictl.JSONWriteUint32(w, item.M)
	if (item.M != 0) == false {
		w = w[:backupIndexM]
	}
	if item.N&(1<<0) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"A":`...)
		if w, err = tlBuiltinTupleTupleInt.BuiltinTupleTupleIntWriteJSONOpt(newTypeNames, short, w, item.A, item.N, item.M); err != nil {
			return w, err
		}
	}
	return append(w, '}'), nil
}

func (item *CasesReplace7plusplus) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *CasesReplace7plusplus) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("cases.replace7plusplus", err.Error())
	}
	return nil
}
