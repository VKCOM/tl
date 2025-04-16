// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlAbCall8

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlTypeA"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tlcd/tlCdTypeB"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type AbCall8 struct {
	X tlTypeA.TypeA
}

func (AbCall8) TLName() string { return "ab.call8" }
func (AbCall8) TLTag() uint32  { return 0x1b8652d9 }

func (item *AbCall8) Reset() {
	item.X.Reset()
}

func (item *AbCall8) FillRandom(rg *basictl.RandGenerator) {
	item.X.FillRandom(rg)
}

func (item *AbCall8) Read(w []byte) (_ []byte, err error) {
	return item.X.Read(w)
}

// This method is general version of Write, use it instead!
func (item *AbCall8) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *AbCall8) Write(w []byte) []byte {
	w = item.X.Write(w)
	return w
}

func (item *AbCall8) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x1b8652d9); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *AbCall8) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *AbCall8) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x1b8652d9)
	return item.Write(w)
}

func (item *AbCall8) ReadResult(w []byte, ret *tlCdTypeB.CdTypeB) (_ []byte, err error) {
	return ret.ReadBoxed(w)
}

func (item *AbCall8) WriteResult(w []byte, ret tlCdTypeB.CdTypeB) (_ []byte, err error) {
	w = ret.WriteBoxed(w)
	return w, nil
}

func (item *AbCall8) ReadResultJSON(legacyTypeNames bool, in *basictl.JsonLexer, ret *tlCdTypeB.CdTypeB) error {
	if err := ret.ReadJSON(legacyTypeNames, in); err != nil {
		return err
	}
	return nil
}

func (item *AbCall8) WriteResultJSON(w []byte, ret tlCdTypeB.CdTypeB) (_ []byte, err error) {
	return item.writeResultJSON(true, false, w, ret)
}

func (item *AbCall8) writeResultJSON(newTypeNames bool, short bool, w []byte, ret tlCdTypeB.CdTypeB) (_ []byte, err error) {
	w = ret.WriteJSONOpt(newTypeNames, short, w)
	return w, nil
}

func (item *AbCall8) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret tlCdTypeB.CdTypeB
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *AbCall8) ReadResultWriteResultJSONOpt(newTypeNames bool, short bool, r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret tlCdTypeB.CdTypeB
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(newTypeNames, short, w, ret)
	return r, w, err
}

func (item *AbCall8) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	var ret tlCdTypeB.CdTypeB
	err := item.ReadResultJSON(true, &basictl.JsonLexer{Data: r}, &ret)
	if err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item *AbCall8) String() string {
	return string(item.WriteJSON(nil))
}

func (item *AbCall8) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
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
					return internal.ErrorInvalidJSONWithDuplicatingKeys("ab.call8", "x")
				}
				if err := item.X.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propXPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("ab.call8", key)
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
func (item *AbCall8) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *AbCall8) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *AbCall8) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"x":`...)
	w = item.X.WriteJSONOpt(newTypeNames, short, w)
	return append(w, '}')
}

func (item *AbCall8) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *AbCall8) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("ab.call8", err.Error())
	}
	return nil
}
