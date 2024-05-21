// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlService1EnableExpiration

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tl/tlBool"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type Service1EnableExpiration struct {
	Prefix string
}

func (Service1EnableExpiration) TLName() string { return "service1.enableExpiration" }
func (Service1EnableExpiration) TLTag() uint32  { return 0x2b51ad67 }

func (item *Service1EnableExpiration) Reset() {
	item.Prefix = ""
}

func (item *Service1EnableExpiration) Read(w []byte) (_ []byte, err error) {
	return basictl.StringRead(w, &item.Prefix)
}

// This method is general version of Write, use it instead!
func (item *Service1EnableExpiration) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *Service1EnableExpiration) Write(w []byte) []byte {
	w = basictl.StringWrite(w, item.Prefix)
	return w
}

func (item *Service1EnableExpiration) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x2b51ad67); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *Service1EnableExpiration) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *Service1EnableExpiration) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x2b51ad67)
	return item.Write(w)
}

func (item *Service1EnableExpiration) ReadResult(w []byte, ret *bool) (_ []byte, err error) {
	return tlBool.BoolReadBoxed(w, ret)
}

func (item *Service1EnableExpiration) WriteResult(w []byte, ret bool) (_ []byte, err error) {
	w = tlBool.BoolWriteBoxed(w, ret)
	return w, nil
}

func (item *Service1EnableExpiration) ReadResultJSON(legacyTypeNames bool, in *basictl.JsonLexer, ret *bool) error {
	if err := internal.Json2ReadBool(in, ret); err != nil {
		return err
	}
	return nil
}

func (item *Service1EnableExpiration) WriteResultJSON(w []byte, ret bool) (_ []byte, err error) {
	return item.writeResultJSON(true, false, w, ret)
}

func (item *Service1EnableExpiration) writeResultJSON(newTypeNames bool, short bool, w []byte, ret bool) (_ []byte, err error) {
	w = basictl.JSONWriteBool(w, ret)
	return w, nil
}

func (item *Service1EnableExpiration) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret bool
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *Service1EnableExpiration) ReadResultWriteResultJSONOpt(newTypeNames bool, short bool, r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret bool
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(newTypeNames, short, w, ret)
	return r, w, err
}

func (item *Service1EnableExpiration) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	var ret bool
	err := item.ReadResultJSON(true, &basictl.JsonLexer{Data: r}, &ret)
	if err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item Service1EnableExpiration) String() string {
	return string(item.WriteJSON(nil))
}

func (item *Service1EnableExpiration) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propPrefixPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "prefix":
				if propPrefixPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service1.enableExpiration", "prefix")
				}
				if err := internal.Json2ReadString(in, &item.Prefix); err != nil {
					return err
				}
				propPrefixPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("service1.enableExpiration", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propPrefixPresented {
		item.Prefix = ""
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *Service1EnableExpiration) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *Service1EnableExpiration) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *Service1EnableExpiration) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexPrefix := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"prefix":`...)
	w = basictl.JSONWriteString(w, item.Prefix)
	if (len(item.Prefix) != 0) == false {
		w = w[:backupIndexPrefix]
	}
	return append(w, '}')
}

func (item *Service1EnableExpiration) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *Service1EnableExpiration) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("service1.enableExpiration", err.Error())
	}
	return nil
}
