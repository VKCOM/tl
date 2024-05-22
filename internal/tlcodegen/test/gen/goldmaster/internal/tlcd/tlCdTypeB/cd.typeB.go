// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlCdTypeB

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tlab/tlAbTypeA"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type CdTypeB struct {
	X tlAbTypeA.AbTypeA
}

func (CdTypeB) TLName() string { return "cd.typeB" }
func (CdTypeB) TLTag() uint32  { return 0x377b4996 }

func (item *CdTypeB) Reset() {
	item.X.Reset()
}

func (item *CdTypeB) FillRandom(rg *basictl.RandGenerator) {
	item.X.FillRandom(rg)
}

func (item *CdTypeB) Read(w []byte) (_ []byte, err error) {
	return item.X.Read(w)
}

// This method is general version of Write, use it instead!
func (item *CdTypeB) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *CdTypeB) Write(w []byte) []byte {
	w = item.X.Write(w)
	return w
}

func (item *CdTypeB) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x377b4996); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *CdTypeB) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *CdTypeB) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x377b4996)
	return item.Write(w)
}

func (item CdTypeB) String() string {
	return string(item.WriteJSON(nil))
}

func (item *CdTypeB) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propXPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "x":
				if propXPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("cd.typeB", "x")
				}
				if err := item.X.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propXPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("cd.typeB", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propXPresented {
		item.X.Reset()
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *CdTypeB) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *CdTypeB) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *CdTypeB) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"x":`...)
	w = item.X.WriteJSONOpt(newTypeNames, short, w)
	return append(w, '}')
}

func (item *CdTypeB) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *CdTypeB) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("cd.typeB", err.Error())
	}
	return nil
}
