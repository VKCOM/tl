// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlCdTopLevel3

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlHalfStr"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlUseStr"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type CdTopLevel3 struct {
	A tlUseStr.UseStr
	B tlHalfStr.HalfStr
}

func (CdTopLevel3) TLName() string { return "cd.topLevel3" }
func (CdTopLevel3) TLTag() uint32  { return 0x5cd1ca89 }

func (item *CdTopLevel3) Reset() {
	item.A.Reset()
	item.B.Reset()
}

func (item *CdTopLevel3) FillRandom(rg *basictl.RandGenerator) {
	item.A.FillRandom(rg)
	item.B.FillRandom(rg)
}

func (item *CdTopLevel3) Read(w []byte) (_ []byte, err error) {
	if w, err = item.A.Read(w); err != nil {
		return w, err
	}
	return item.B.Read(w)
}

// This method is general version of Write, use it instead!
func (item *CdTopLevel3) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *CdTopLevel3) Write(w []byte) []byte {
	w = item.A.Write(w)
	w = item.B.Write(w)
	return w
}

func (item *CdTopLevel3) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x5cd1ca89); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *CdTopLevel3) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *CdTopLevel3) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x5cd1ca89)
	return item.Write(w)
}

func (item CdTopLevel3) String() string {
	return string(item.WriteJSON(nil))
}

func (item *CdTopLevel3) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propAPresented bool
	var propBPresented bool

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
					return internal.ErrorInvalidJSONWithDuplicatingKeys("cd.topLevel3", "a")
				}
				if err := item.A.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propAPresented = true
			case "b":
				if propBPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("cd.topLevel3", "b")
				}
				if err := item.B.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propBPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("cd.topLevel3", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propAPresented {
		item.A.Reset()
	}
	if !propBPresented {
		item.B.Reset()
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *CdTopLevel3) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *CdTopLevel3) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *CdTopLevel3) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"a":`...)
	w = item.A.WriteJSONOpt(newTypeNames, short, w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"b":`...)
	w = item.B.WriteJSONOpt(newTypeNames, short, w)
	return append(w, '}')
}

func (item *CdTopLevel3) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *CdTopLevel3) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("cd.topLevel3", err.Error())
	}
	return nil
}
