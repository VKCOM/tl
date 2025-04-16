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

type Replace6 struct {
	A []int32
}

func (Replace6) TLName() string { return "replace6" }
func (Replace6) TLTag() uint32  { return 0xabd49d06 }

func (item *Replace6) Reset() {
	item.A = item.A[:0]
}

func (item *Replace6) FillRandom(rg *basictl.RandGenerator) {
	BuiltinVectorIntFillRandom(rg, &item.A)
}

func (item *Replace6) Read(w []byte) (_ []byte, err error) {
	return BuiltinVectorIntRead(w, &item.A)
}

// This method is general version of Write, use it instead!
func (item *Replace6) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *Replace6) Write(w []byte) []byte {
	w = BuiltinVectorIntWrite(w, item.A)
	return w
}

func (item *Replace6) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xabd49d06); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *Replace6) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *Replace6) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0xabd49d06)
	return item.Write(w)
}

func (item *Replace6) String() string {
	return string(item.WriteJSON(nil))
}

func (item *Replace6) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
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
			case "a":
				if propAPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("replace6", "a")
				}
				if err := BuiltinVectorIntReadJSON(legacyTypeNames, in, &item.A); err != nil {
					return err
				}
				propAPresented = true
			default:
				return ErrorInvalidJSONExcessElement("replace6", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propAPresented {
		item.A = item.A[:0]
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *Replace6) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *Replace6) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *Replace6) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexA := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"a":`...)
	w = BuiltinVectorIntWriteJSONOpt(newTypeNames, short, w, item.A)
	if (len(item.A) != 0) == false {
		w = w[:backupIndexA]
	}
	return append(w, '}')
}

func (item *Replace6) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *Replace6) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("replace6", err.Error())
	}
	return nil
}
