// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlCasesTestArray

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tl/tlBuiltinTupleInt"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type CasesTestArray struct {
	N   uint32
	Arr []int32
}

func (CasesTestArray) TLName() string { return "cases.testArray" }
func (CasesTestArray) TLTag() uint32  { return 0xa888030d }

func (item *CasesTestArray) Reset() {
	item.N = 0
	item.Arr = item.Arr[:0]
}

func (item *CasesTestArray) FillRandom(rg *basictl.RandGenerator) {
	item.N = basictl.RandomUint(rg)
	item.N = rg.LimitValue(item.N)
	tlBuiltinTupleInt.BuiltinTupleIntFillRandom(rg, &item.Arr, item.N)
}

func (item *CasesTestArray) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.N); err != nil {
		return w, err
	}
	return tlBuiltinTupleInt.BuiltinTupleIntRead(w, &item.Arr, item.N)
}

// This method is general version of Write, use it instead!
func (item *CasesTestArray) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w)
}

func (item *CasesTestArray) Write(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, item.N)
	if w, err = tlBuiltinTupleInt.BuiltinTupleIntWrite(w, item.Arr, item.N); err != nil {
		return w, err
	}
	return w, nil
}

func (item *CasesTestArray) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xa888030d); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *CasesTestArray) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w)
}

func (item *CasesTestArray) WriteBoxed(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, 0xa888030d)
	return item.Write(w)
}

func (item CasesTestArray) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func (item *CasesTestArray) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propNPresented bool
	var rawArr []byte

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
					return internal.ErrorInvalidJSONWithDuplicatingKeys("cases.testArray", "n")
				}
				if err := internal.Json2ReadUint32(in, &item.N); err != nil {
					return err
				}
				propNPresented = true
			case "arr":
				if rawArr != nil {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("cases.testArray", "arr")
				}
				rawArr = in.Raw()
				if !in.Ok() {
					return in.Error()
				}
			default:
				return internal.ErrorInvalidJSONExcessElement("cases.testArray", key)
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
	var inArrPointer *basictl.JsonLexer
	inArr := basictl.JsonLexer{Data: rawArr}
	if rawArr != nil {
		inArrPointer = &inArr
	}
	if err := tlBuiltinTupleInt.BuiltinTupleIntReadJSON(legacyTypeNames, inArrPointer, &item.Arr, item.N); err != nil {
		return err
	}

	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *CasesTestArray) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w)
}

func (item *CasesTestArray) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w)
}
func (item *CasesTestArray) WriteJSONOpt(newTypeNames bool, short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	backupIndexN := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"n":`...)
	w = basictl.JSONWriteUint32(w, item.N)
	if (item.N != 0) == false {
		w = w[:backupIndexN]
	}
	backupIndexArr := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"arr":`...)
	if w, err = tlBuiltinTupleInt.BuiltinTupleIntWriteJSONOpt(newTypeNames, short, w, item.Arr, item.N); err != nil {
		return w, err
	}
	if (len(item.Arr) != 0) == false {
		w = w[:backupIndexArr]
	}
	return append(w, '}'), nil
}

func (item *CasesTestArray) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *CasesTestArray) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("cases.testArray", err.Error())
	}
	return nil
}
