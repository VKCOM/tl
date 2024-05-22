// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlAntispamGetPattern

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/cycle_4174bfee82ee7ea4902a121c2642c5ff"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type AntispamGetPattern struct {
	Id int32
}

func (AntispamGetPattern) TLName() string { return "antispam.getPattern" }
func (AntispamGetPattern) TLTag() uint32  { return 0x3de14136 }

func (item *AntispamGetPattern) Reset() {
	item.Id = 0
}

func (item *AntispamGetPattern) Read(w []byte) (_ []byte, err error) {
	return basictl.IntRead(w, &item.Id)
}

// This method is general version of Write, use it instead!
func (item *AntispamGetPattern) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *AntispamGetPattern) Write(w []byte) []byte {
	w = basictl.IntWrite(w, item.Id)
	return w
}

func (item *AntispamGetPattern) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x3de14136); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *AntispamGetPattern) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *AntispamGetPattern) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x3de14136)
	return item.Write(w)
}

func (item *AntispamGetPattern) ReadResult(w []byte, ret *cycle_4174bfee82ee7ea4902a121c2642c5ff.AntispamPatternFull) (_ []byte, err error) {
	return ret.ReadBoxed(w)
}

func (item *AntispamGetPattern) WriteResult(w []byte, ret cycle_4174bfee82ee7ea4902a121c2642c5ff.AntispamPatternFull) (_ []byte, err error) {
	w = ret.WriteBoxed(w)
	return w, nil
}

func (item *AntispamGetPattern) ReadResultJSON(legacyTypeNames bool, in *basictl.JsonLexer, ret *cycle_4174bfee82ee7ea4902a121c2642c5ff.AntispamPatternFull) error {
	if err := ret.ReadJSON(legacyTypeNames, in); err != nil {
		return err
	}
	return nil
}

func (item *AntispamGetPattern) WriteResultJSON(w []byte, ret cycle_4174bfee82ee7ea4902a121c2642c5ff.AntispamPatternFull) (_ []byte, err error) {
	return item.writeResultJSON(true, false, w, ret)
}

func (item *AntispamGetPattern) writeResultJSON(newTypeNames bool, short bool, w []byte, ret cycle_4174bfee82ee7ea4902a121c2642c5ff.AntispamPatternFull) (_ []byte, err error) {
	w = ret.WriteJSONOpt(newTypeNames, short, w)
	return w, nil
}

func (item *AntispamGetPattern) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret cycle_4174bfee82ee7ea4902a121c2642c5ff.AntispamPatternFull
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *AntispamGetPattern) ReadResultWriteResultJSONOpt(newTypeNames bool, short bool, r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret cycle_4174bfee82ee7ea4902a121c2642c5ff.AntispamPatternFull
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(newTypeNames, short, w, ret)
	return r, w, err
}

func (item *AntispamGetPattern) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	var ret cycle_4174bfee82ee7ea4902a121c2642c5ff.AntispamPatternFull
	err := item.ReadResultJSON(true, &basictl.JsonLexer{Data: r}, &ret)
	if err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item AntispamGetPattern) String() string {
	return string(item.WriteJSON(nil))
}

func (item *AntispamGetPattern) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propIdPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "id":
				if propIdPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("antispam.getPattern", "id")
				}
				if err := internal.Json2ReadInt32(in, &item.Id); err != nil {
					return err
				}
				propIdPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("antispam.getPattern", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propIdPresented {
		item.Id = 0
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *AntispamGetPattern) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *AntispamGetPattern) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *AntispamGetPattern) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexId := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"id":`...)
	w = basictl.JSONWriteInt32(w, item.Id)
	if (item.Id != 0) == false {
		w = w[:backupIndexId]
	}
	return append(w, '}')
}

func (item *AntispamGetPattern) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *AntispamGetPattern) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("antispam.getPattern", err.Error())
	}
	return nil
}
