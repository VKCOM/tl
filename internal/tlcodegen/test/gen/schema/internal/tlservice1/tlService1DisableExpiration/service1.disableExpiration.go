// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlService1DisableExpiration

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tl/tlBool"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type Service1DisableExpiration struct {
	Prefix string
}

func (Service1DisableExpiration) TLName() string { return "service1.disableExpiration" }
func (Service1DisableExpiration) TLTag() uint32  { return 0xf1c39c2d }

func (item *Service1DisableExpiration) Reset() {
	item.Prefix = ""
}

func (item *Service1DisableExpiration) Read(w []byte) (_ []byte, err error) {
	return basictl.StringRead(w, &item.Prefix)
}

func (item *Service1DisableExpiration) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *Service1DisableExpiration) Write(w []byte) []byte {
	w = basictl.StringWrite(w, item.Prefix)
	return w
}

func (item *Service1DisableExpiration) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xf1c39c2d); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *Service1DisableExpiration) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *Service1DisableExpiration) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0xf1c39c2d)
	return item.Write(w)
}

func (item *Service1DisableExpiration) ReadResult(w []byte, ret *bool) (_ []byte, err error) {
	return tlBool.BoolReadBoxed(w, ret)
}

func (item *Service1DisableExpiration) WriteResult(w []byte, ret bool) (_ []byte, err error) {
	w = tlBool.BoolWriteBoxed(w, ret)
	return w, nil
}

func (item *Service1DisableExpiration) ReadResultJSON(legacyTypeNames bool, in *basictl.JsonLexer, ret *bool) error {
	if err := internal.Json2ReadBool(in, ret); err != nil {
		return err
	}
	return nil
}

func (item *Service1DisableExpiration) WriteResultJSON(w []byte, ret bool) (_ []byte, err error) {
	tctx := basictl.JSONWriteContext{}
	return item.writeResultJSON(&tctx, w, ret)
}

func (item *Service1DisableExpiration) writeResultJSON(tctx *basictl.JSONWriteContext, w []byte, ret bool) (_ []byte, err error) {
	w = basictl.JSONWriteBool(w, ret)
	return w, nil
}

func (item *Service1DisableExpiration) ReadResultWriteResultJSON(tctx *basictl.JSONWriteContext, r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret bool
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(tctx, w, ret)
	return r, w, err
}

func (item *Service1DisableExpiration) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	var ret bool
	err := item.ReadResultJSON(true, &basictl.JsonLexer{Data: r}, &ret)
	if err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item Service1DisableExpiration) String() string {
	return string(item.WriteJSON(nil))
}

func (item *Service1DisableExpiration) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
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
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service1.disableExpiration", "prefix")
				}
				if err := internal.Json2ReadString(in, &item.Prefix); err != nil {
					return err
				}
				propPrefixPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("service1.disableExpiration", key)
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
func (item *Service1DisableExpiration) WriteJSONGeneral(tctx *basictl.JSONWriteContext, w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(tctx, w), nil
}

func (item *Service1DisableExpiration) WriteJSON(w []byte) []byte {
	tctx := basictl.JSONWriteContext{}
	return item.WriteJSONOpt(&tctx, w)
}
func (item *Service1DisableExpiration) WriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte) []byte {
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

func (item *Service1DisableExpiration) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *Service1DisableExpiration) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("service1.disableExpiration", err.Error())
	}
	return nil
}
