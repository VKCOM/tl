// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlHalfStr

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlNoStr"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlUseStr"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type HalfStr struct {
	X tlNoStr.NoStr
	Y tlUseStr.UseStr
}

func (HalfStr) TLName() string { return "halfStr" }
func (HalfStr) TLTag() uint32  { return 0x647ddaf5 }

func (item *HalfStr) Reset() {
	item.X.Reset()
	item.Y.Reset()
}

func (item *HalfStr) FillRandom(rg *basictl.RandGenerator) {
	item.X.FillRandom(rg)
	item.Y.FillRandom(rg)
}

func (item *HalfStr) Read(w []byte) (_ []byte, err error) {
	if w, err = item.X.Read(w); err != nil {
		return w, err
	}
	return item.Y.Read(w)
}

// This method is general version of Write, use it instead!
func (item *HalfStr) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *HalfStr) Write(w []byte) []byte {
	w = item.X.Write(w)
	w = item.Y.Write(w)
	return w
}

func (item *HalfStr) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x647ddaf5); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *HalfStr) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *HalfStr) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x647ddaf5)
	return item.Write(w)
}

func (item *HalfStr) String() string {
	return string(item.WriteJSON(nil))
}

func (item *HalfStr) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propXPresented bool
	var propYPresented bool

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
					return internal.ErrorInvalidJSONWithDuplicatingKeys("halfStr", "x")
				}
				if err := item.X.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propXPresented = true
			case "y":
				if propYPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("halfStr", "y")
				}
				if err := item.Y.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propYPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("halfStr", key)
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
	if !propYPresented {
		item.Y.Reset()
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *HalfStr) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *HalfStr) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *HalfStr) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"x":`...)
	w = item.X.WriteJSONOpt(newTypeNames, short, w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"y":`...)
	w = item.Y.WriteJSONOpt(newTypeNames, short, w)
	return append(w, '}')
}

func (item *HalfStr) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *HalfStr) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("halfStr", err.Error())
	}
	return nil
}
