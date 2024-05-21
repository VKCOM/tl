// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlCasesTestVector

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tl/tlBuiltinVectorInt"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type CasesTestVector struct {
	Arr []int32
}

func (CasesTestVector) TLName() string { return "cases.testVector" }
func (CasesTestVector) TLTag() uint32  { return 0x4975695c }

func (item *CasesTestVector) Reset() {
	item.Arr = item.Arr[:0]
}

func (item *CasesTestVector) FillRandom(rg *basictl.RandGenerator) {
	tlBuiltinVectorInt.BuiltinVectorIntFillRandom(rg, &item.Arr)
}

func (item *CasesTestVector) Read(w []byte) (_ []byte, err error) {
	return tlBuiltinVectorInt.BuiltinVectorIntRead(w, &item.Arr)
}

// This method is general version of Write, use it instead!
func (item *CasesTestVector) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *CasesTestVector) Write(w []byte) []byte {
	w = tlBuiltinVectorInt.BuiltinVectorIntWrite(w, item.Arr)
	return w
}

func (item *CasesTestVector) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x4975695c); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *CasesTestVector) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *CasesTestVector) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x4975695c)
	return item.Write(w)
}

func (item CasesTestVector) String() string {
	return string(item.WriteJSON(nil))
}

func (item *CasesTestVector) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propArrPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "arr":
				if propArrPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("cases.testVector", "arr")
				}
				if err := tlBuiltinVectorInt.BuiltinVectorIntReadJSON(legacyTypeNames, in, &item.Arr); err != nil {
					return err
				}
				propArrPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("cases.testVector", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propArrPresented {
		item.Arr = item.Arr[:0]
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *CasesTestVector) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *CasesTestVector) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *CasesTestVector) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexArr := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"arr":`...)
	w = tlBuiltinVectorInt.BuiltinVectorIntWriteJSONOpt(newTypeNames, short, w, item.Arr)
	if (len(item.Arr) != 0) == false {
		w = w[:backupIndexArr]
	}
	return append(w, '}')
}

func (item *CasesTestVector) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *CasesTestVector) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("cases.testVector", err.Error())
	}
	return nil
}