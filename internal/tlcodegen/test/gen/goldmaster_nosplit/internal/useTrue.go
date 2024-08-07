// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package internal

import (
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite

type UseTrue struct {
	Fm uint32
	// A (TrueType) // Conditional: item.Fm.0
	// B (TrueType) // Conditional: item.Fm.1
	// C (TrueType)
	// D (TrueType)
	E bool // Conditional: item.Fm.2
}

func (UseTrue) TLName() string { return "useTrue" }
func (UseTrue) TLTag() uint32  { return 0xdfdd4180 }

func (item *UseTrue) SetA(v bool) {
	if v {
		item.Fm |= 1 << 0
	} else {
		item.Fm &^= 1 << 0
	}
}
func (item UseTrue) IsSetA() bool { return item.Fm&(1<<0) != 0 }

func (item *UseTrue) SetB(v bool) {
	if v {
		item.Fm |= 1 << 1
	} else {
		item.Fm &^= 1 << 1
	}
}
func (item UseTrue) IsSetB() bool { return item.Fm&(1<<1) != 0 }

func (item *UseTrue) SetE(v bool) {
	item.E = v
	item.Fm |= 1 << 2
}
func (item *UseTrue) ClearE() {
	item.E = false
	item.Fm &^= 1 << 2
}
func (item UseTrue) IsSetE() bool { return item.Fm&(1<<2) != 0 }

func (item *UseTrue) Reset() {
	item.Fm = 0
	item.E = false
}

func (item *UseTrue) FillRandom(rg *basictl.RandGenerator) {
	var maskFm uint32
	maskFm = basictl.RandomUint(rg)
	item.Fm = 0
	if maskFm&(1<<0) != 0 {
		item.Fm |= (1 << 0)
	}
	if maskFm&(1<<1) != 0 {
		item.Fm |= (1 << 1)
	}
	if maskFm&(1<<2) != 0 {
		item.Fm |= (1 << 2)
	}
	if item.Fm&(1<<2) != 0 {
		item.E = basictl.RandomUint(rg)&1 == 1
	} else {
		item.E = false
	}
}

func (item *UseTrue) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.Fm); err != nil {
		return w, err
	}
	if item.Fm&(1<<1) != 0 {
		if w, err = basictl.NatReadExactTag(w, 0x3fedd339); err != nil {
			return w, err
		}
	}
	if w, err = basictl.NatReadExactTag(w, 0x3fedd339); err != nil {
		return w, err
	}
	if item.Fm&(1<<2) != 0 {
		if w, err = BoolReadBoxed(w, &item.E); err != nil {
			return w, err
		}
	} else {
		item.E = false
	}
	return w, nil
}

// This method is general version of Write, use it instead!
func (item *UseTrue) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *UseTrue) Write(w []byte) []byte {
	w = basictl.NatWrite(w, item.Fm)
	if item.Fm&(1<<1) != 0 {
		w = basictl.NatWrite(w, 0x3fedd339)
	}
	w = basictl.NatWrite(w, 0x3fedd339)
	if item.Fm&(1<<2) != 0 {
		w = BoolWriteBoxed(w, item.E)
	}
	return w
}

func (item *UseTrue) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xdfdd4180); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *UseTrue) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *UseTrue) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0xdfdd4180)
	return item.Write(w)
}

func (item UseTrue) String() string {
	return string(item.WriteJSON(nil))
}

func (item *UseTrue) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propFmPresented bool
	var trueTypeAPresented bool
	var trueTypeAValue bool
	var trueTypeBPresented bool
	var trueTypeBValue bool
	var propEPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "fm":
				if propFmPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("useTrue", "fm")
				}
				if err := Json2ReadUint32(in, &item.Fm); err != nil {
					return err
				}
				propFmPresented = true
			case "a":
				if trueTypeAPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("useTrue", "a")
				}
				if err := Json2ReadBool(in, &trueTypeAValue); err != nil {
					return err
				}
				trueTypeAPresented = true
			case "b":
				if trueTypeBPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("useTrue", "b")
				}
				if err := Json2ReadBool(in, &trueTypeBValue); err != nil {
					return err
				}
				trueTypeBPresented = true
			case "c":
				var tmpC True
				if err := tmpC.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
			case "d":
				var tmpD True
				if err := tmpD.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
			case "e":
				if propEPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("useTrue", "e")
				}
				if err := Json2ReadBool(in, &item.E); err != nil {
					return err
				}
				propEPresented = true
			default:
				return ErrorInvalidJSONExcessElement("useTrue", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propFmPresented {
		item.Fm = 0
	}
	if !propEPresented {
		item.E = false
	}
	if trueTypeAPresented {
		if trueTypeAValue {
			item.Fm |= 1 << 0
		}
	}
	if trueTypeBPresented {
		if trueTypeBValue {
			item.Fm |= 1 << 1
		}
	}
	if propEPresented {
		item.Fm |= 1 << 2
	}
	// tries to set bit to zero if it is 1
	if trueTypeAPresented && !trueTypeAValue && (item.Fm&(1<<0) != 0) {
		return ErrorInvalidJSON("useTrue", "fieldmask bit fm.0 is indefinite because of the contradictions in values")
	}
	// tries to set bit to zero if it is 1
	if trueTypeBPresented && !trueTypeBValue && (item.Fm&(1<<1) != 0) {
		return ErrorInvalidJSON("useTrue", "fieldmask bit fm.0 is indefinite because of the contradictions in values")
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *UseTrue) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *UseTrue) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *UseTrue) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexFm := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"fm":`...)
	w = basictl.JSONWriteUint32(w, item.Fm)
	if (item.Fm != 0) == false {
		w = w[:backupIndexFm]
	}
	if item.Fm&(1<<0) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"a":true`...)
	}
	if item.Fm&(1<<1) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"b":true`...)
	}
	if item.Fm&(1<<2) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"e":`...)
		w = basictl.JSONWriteBool(w, item.E)
	}
	return append(w, '}')
}

func (item *UseTrue) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *UseTrue) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("useTrue", err.Error())
	}
	return nil
}
