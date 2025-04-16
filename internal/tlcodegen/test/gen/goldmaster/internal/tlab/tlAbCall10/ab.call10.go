// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlAbCall10

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tla/tlAColor"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type AbCall10 struct {
}

func (AbCall10) TLName() string { return "ab.call10" }
func (AbCall10) TLTag() uint32  { return 0x8db2a4f8 }

func (item *AbCall10) Reset() {}

func (item *AbCall10) FillRandom(rg *basictl.RandGenerator) {}

func (item *AbCall10) Read(w []byte) (_ []byte, err error) { return w, nil }

// This method is general version of Write, use it instead!
func (item *AbCall10) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *AbCall10) Write(w []byte) []byte {
	return w
}

func (item *AbCall10) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x8db2a4f8); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *AbCall10) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *AbCall10) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x8db2a4f8)
	return item.Write(w)
}

func (item *AbCall10) ReadResult(w []byte, ret *tlAColor.AColor) (_ []byte, err error) {
	return ret.ReadBoxed(w)
}

func (item *AbCall10) WriteResult(w []byte, ret tlAColor.AColor) (_ []byte, err error) {
	w = ret.WriteBoxed(w)
	return w, nil
}

func (item *AbCall10) ReadResultJSON(legacyTypeNames bool, in *basictl.JsonLexer, ret *tlAColor.AColor) error {
	if err := ret.ReadJSON(legacyTypeNames, in); err != nil {
		return err
	}
	return nil
}

func (item *AbCall10) WriteResultJSON(w []byte, ret tlAColor.AColor) (_ []byte, err error) {
	return item.writeResultJSON(true, false, w, ret)
}

func (item *AbCall10) writeResultJSON(newTypeNames bool, short bool, w []byte, ret tlAColor.AColor) (_ []byte, err error) {
	w = ret.WriteJSONOpt(newTypeNames, short, w)
	return w, nil
}

func (item *AbCall10) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret tlAColor.AColor
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *AbCall10) ReadResultWriteResultJSONOpt(newTypeNames bool, short bool, r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret tlAColor.AColor
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(newTypeNames, short, w, ret)
	return r, w, err
}

func (item *AbCall10) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	var ret tlAColor.AColor
	err := item.ReadResultJSON(true, &basictl.JsonLexer{Data: r}, &ret)
	if err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item *AbCall10) String() string {
	return string(item.WriteJSON(nil))
}

func (item *AbCall10) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			return internal.ErrorInvalidJSON("ab.call10", "this object can't have properties")
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *AbCall10) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *AbCall10) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *AbCall10) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	return append(w, '}')
}

func (item *AbCall10) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *AbCall10) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("ab.call10", err.Error())
	}
	return nil
}
