// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlService6Error

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type Service6Error struct {
	Code int32
}

func (Service6Error) TLName() string { return "service6.error" }
func (Service6Error) TLTag() uint32  { return 0x738553ef }

func (item *Service6Error) Reset() {
	item.Code = 0
}

func (item *Service6Error) Read(w []byte) (_ []byte, err error) {
	return basictl.IntRead(w, &item.Code)
}

// This method is general version of Write, use it instead!
func (item *Service6Error) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *Service6Error) Write(w []byte) []byte {
	w = basictl.IntWrite(w, item.Code)
	return w
}

func (item *Service6Error) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x738553ef); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *Service6Error) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *Service6Error) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x738553ef)
	return item.Write(w)
}

func (item *Service6Error) String() string {
	return string(item.WriteJSON(nil))
}

func (item *Service6Error) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propCodePresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "code":
				if propCodePresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service6.error", "code")
				}
				if err := internal.Json2ReadInt32(in, &item.Code); err != nil {
					return err
				}
				propCodePresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("service6.error", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propCodePresented {
		item.Code = 0
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *Service6Error) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *Service6Error) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *Service6Error) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexCode := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"code":`...)
	w = basictl.JSONWriteInt32(w, item.Code)
	if (item.Code != 0) == false {
		w = w[:backupIndexCode]
	}
	return append(w, '}')
}

func (item *Service6Error) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *Service6Error) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("service6.error", err.Error())
	}
	return nil
}
