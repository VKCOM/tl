// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlAbCall6

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlTypeB"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tlcd/tlCdTypeA"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type AbCall6 struct {
	X tlCdTypeA.CdTypeA
}

func (AbCall6) TLName() string { return "ab.call6" }
func (AbCall6) TLTag() uint32  { return 0x84d815cb }

func (item *AbCall6) Reset() {
	item.X.Reset()
}

func (item *AbCall6) FillRandom(rg *basictl.RandGenerator) {
	item.X.FillRandom(rg)
}

func (item *AbCall6) Read(w []byte) (_ []byte, err error) {
	return item.X.Read(w)
}

// This method is general version of Write, use it instead!
func (item *AbCall6) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *AbCall6) Write(w []byte) []byte {
	w = item.X.Write(w)
	return w
}

func (item *AbCall6) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x84d815cb); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *AbCall6) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *AbCall6) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x84d815cb)
	return item.Write(w)
}

func (item *AbCall6) ReadResult(w []byte, ret *tlTypeB.TypeB) (_ []byte, err error) {
	return ret.ReadBoxed(w)
}

func (item *AbCall6) WriteResult(w []byte, ret tlTypeB.TypeB) (_ []byte, err error) {
	w = ret.WriteBoxed(w)
	return w, nil
}

func (item *AbCall6) ReadResultJSON(legacyTypeNames bool, in *basictl.JsonLexer, ret *tlTypeB.TypeB) error {
	if err := ret.ReadJSON(legacyTypeNames, in); err != nil {
		return err
	}
	return nil
}

func (item *AbCall6) WriteResultJSON(w []byte, ret tlTypeB.TypeB) (_ []byte, err error) {
	return item.writeResultJSON(true, false, w, ret)
}

func (item *AbCall6) writeResultJSON(newTypeNames bool, short bool, w []byte, ret tlTypeB.TypeB) (_ []byte, err error) {
	w = ret.WriteJSONOpt(newTypeNames, short, w)
	return w, nil
}

func (item *AbCall6) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret tlTypeB.TypeB
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *AbCall6) ReadResultWriteResultJSONOpt(newTypeNames bool, short bool, r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret tlTypeB.TypeB
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(newTypeNames, short, w, ret)
	return r, w, err
}

func (item *AbCall6) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	var ret tlTypeB.TypeB
	err := item.ReadResultJSON(true, &basictl.JsonLexer{Data: r}, &ret)
	if err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item *AbCall6) String() string {
	return string(item.WriteJSON(nil))
}

func (item *AbCall6) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
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
					return internal.ErrorInvalidJSONWithDuplicatingKeys("ab.call6", "x")
				}
				if err := item.X.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propXPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("ab.call6", key)
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
func (item *AbCall6) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *AbCall6) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *AbCall6) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"x":`...)
	w = item.X.WriteJSONOpt(newTypeNames, short, w)
	return append(w, '}')
}

func (item *AbCall6) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *AbCall6) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("ab.call6", err.Error())
	}
	return nil
}
