// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package cycle_44515dca4b2e76ca676b13645e716786

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type CasesMyCycle2 struct {
	FieldsMask uint32
	A          CasesMyCycle3 // Conditional: item.FieldsMask.0
}

func (CasesMyCycle2) TLName() string { return "cases.myCycle2" }
func (CasesMyCycle2) TLTag() uint32  { return 0x5444c9a2 }

func (item *CasesMyCycle2) SetA(v CasesMyCycle3) {
	item.A = v
	item.FieldsMask |= 1 << 0
}
func (item *CasesMyCycle2) ClearA() {
	item.A.Reset()
	item.FieldsMask &^= 1 << 0
}
func (item *CasesMyCycle2) IsSetA() bool { return item.FieldsMask&(1<<0) != 0 }

func (item *CasesMyCycle2) Reset() {
	item.FieldsMask = 0
	item.A.Reset()
}

func (item *CasesMyCycle2) FillRandom(rg *basictl.RandGenerator) {
	var maskFieldsMask uint32
	maskFieldsMask = basictl.RandomUint(rg)
	item.FieldsMask = 0
	if maskFieldsMask&(1<<0) != 0 {
		item.FieldsMask |= (1 << 0)
	}
	if item.FieldsMask&(1<<0) != 0 {
		item.A.FillRandom(rg)
	} else {
		item.A.Reset()
	}
}

func (item *CasesMyCycle2) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.FieldsMask); err != nil {
		return w, err
	}
	if item.FieldsMask&(1<<0) != 0 {
		if w, err = item.A.Read(w); err != nil {
			return w, err
		}
	} else {
		item.A.Reset()
	}
	return w, nil
}

// This method is general version of Write, use it instead!
func (item *CasesMyCycle2) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *CasesMyCycle2) Write(w []byte) []byte {
	w = basictl.NatWrite(w, item.FieldsMask)
	if item.FieldsMask&(1<<0) != 0 {
		w = item.A.Write(w)
	}
	return w
}

func (item *CasesMyCycle2) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x5444c9a2); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *CasesMyCycle2) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *CasesMyCycle2) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x5444c9a2)
	return item.Write(w)
}

func (item CasesMyCycle2) String() string {
	return string(item.WriteJSON(nil))
}

func (item *CasesMyCycle2) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
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
					return internal.ErrorInvalidJSONWithDuplicatingKeys("cases.myCycle2", "fields_mask")
				}
				if err := internal.Json2ReadUint32(in, &item.FieldsMask); err != nil {
					return err
				}
				propFieldsMaskPresented = true
			case "a":
				if propAPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("cases.myCycle2", "a")
				}
				if err := item.A.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propAPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("cases.myCycle2", key)
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
		item.A.Reset()
	}
	if propAPresented {
		item.FieldsMask |= 1 << 0
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *CasesMyCycle2) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *CasesMyCycle2) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *CasesMyCycle2) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
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

func (item *CasesMyCycle2) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *CasesMyCycle2) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("cases.myCycle2", err.Error())
	}
	return nil
}
