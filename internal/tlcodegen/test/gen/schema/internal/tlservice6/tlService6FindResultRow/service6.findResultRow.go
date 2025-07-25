// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlService6FindResultRow

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type Service6FindResultRow struct {
	X int32
}

func (Service6FindResultRow) TLName() string { return "service6.findResultRow" }
func (Service6FindResultRow) TLTag() uint32  { return 0xbd3946e3 }

func (item *Service6FindResultRow) Reset() {
	item.X = 0
}

func (item *Service6FindResultRow) Read(w []byte) (_ []byte, err error) {
	return basictl.IntRead(w, &item.X)
}

func (item *Service6FindResultRow) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *Service6FindResultRow) Write(w []byte) []byte {
	w = basictl.IntWrite(w, item.X)
	return w
}

func (item *Service6FindResultRow) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xbd3946e3); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *Service6FindResultRow) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *Service6FindResultRow) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0xbd3946e3)
	return item.Write(w)
}

func (item Service6FindResultRow) String() string {
	return string(item.WriteJSON(nil))
}

func (item *Service6FindResultRow) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
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
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service6.findResultRow", "x")
				}
				if err := internal.Json2ReadInt32(in, &item.X); err != nil {
					return err
				}
				propXPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("service6.findResultRow", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propXPresented {
		item.X = 0
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *Service6FindResultRow) WriteJSONGeneral(tctx *basictl.JSONWriteContext, w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(tctx, w), nil
}

func (item *Service6FindResultRow) WriteJSON(w []byte) []byte {
	tctx := basictl.JSONWriteContext{}
	return item.WriteJSONOpt(&tctx, w)
}
func (item *Service6FindResultRow) WriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte) []byte {
	w = append(w, '{')
	backupIndexX := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"x":`...)
	w = basictl.JSONWriteInt32(w, item.X)
	if (item.X != 0) == false {
		w = w[:backupIndexX]
	}
	return append(w, '}')
}

func (item *Service6FindResultRow) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *Service6FindResultRow) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("service6.findResultRow", err.Error())
	}
	return nil
}
