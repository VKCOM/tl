// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlService3GetLastVisitTimestamp

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tl/tlIntMaybe"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type Service3GetLastVisitTimestamp struct {
	UserId int32
}

func (Service3GetLastVisitTimestamp) TLName() string { return "service3.getLastVisitTimestamp" }
func (Service3GetLastVisitTimestamp) TLTag() uint32  { return 0x9a4c788d }

func (item *Service3GetLastVisitTimestamp) Reset() {
	item.UserId = 0
}

func (item *Service3GetLastVisitTimestamp) Read(w []byte) (_ []byte, err error) {
	return basictl.IntRead(w, &item.UserId)
}

// This method is general version of Write, use it instead!
func (item *Service3GetLastVisitTimestamp) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *Service3GetLastVisitTimestamp) Write(w []byte) []byte {
	w = basictl.IntWrite(w, item.UserId)
	return w
}

func (item *Service3GetLastVisitTimestamp) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x9a4c788d); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *Service3GetLastVisitTimestamp) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *Service3GetLastVisitTimestamp) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x9a4c788d)
	return item.Write(w)
}

func (item *Service3GetLastVisitTimestamp) ReadResult(w []byte, ret *tlIntMaybe.IntMaybe) (_ []byte, err error) {
	return ret.ReadBoxed(w)
}

func (item *Service3GetLastVisitTimestamp) WriteResult(w []byte, ret tlIntMaybe.IntMaybe) (_ []byte, err error) {
	w = ret.WriteBoxed(w)
	return w, nil
}

func (item *Service3GetLastVisitTimestamp) ReadResultJSON(legacyTypeNames bool, in *basictl.JsonLexer, ret *tlIntMaybe.IntMaybe) error {
	if err := ret.ReadJSON(legacyTypeNames, in); err != nil {
		return err
	}
	return nil
}

func (item *Service3GetLastVisitTimestamp) WriteResultJSON(w []byte, ret tlIntMaybe.IntMaybe) (_ []byte, err error) {
	return item.writeResultJSON(true, false, w, ret)
}

func (item *Service3GetLastVisitTimestamp) writeResultJSON(newTypeNames bool, short bool, w []byte, ret tlIntMaybe.IntMaybe) (_ []byte, err error) {
	w = ret.WriteJSONOpt(newTypeNames, short, w)
	return w, nil
}

func (item *Service3GetLastVisitTimestamp) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret tlIntMaybe.IntMaybe
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *Service3GetLastVisitTimestamp) ReadResultWriteResultJSONOpt(newTypeNames bool, short bool, r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret tlIntMaybe.IntMaybe
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(newTypeNames, short, w, ret)
	return r, w, err
}

func (item *Service3GetLastVisitTimestamp) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	var ret tlIntMaybe.IntMaybe
	err := item.ReadResultJSON(true, &basictl.JsonLexer{Data: r}, &ret)
	if err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item *Service3GetLastVisitTimestamp) String() string {
	return string(item.WriteJSON(nil))
}

func (item *Service3GetLastVisitTimestamp) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propUserIdPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "user_id":
				if propUserIdPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service3.getLastVisitTimestamp", "user_id")
				}
				if err := internal.Json2ReadInt32(in, &item.UserId); err != nil {
					return err
				}
				propUserIdPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("service3.getLastVisitTimestamp", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propUserIdPresented {
		item.UserId = 0
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *Service3GetLastVisitTimestamp) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *Service3GetLastVisitTimestamp) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *Service3GetLastVisitTimestamp) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexUserId := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"user_id":`...)
	w = basictl.JSONWriteInt32(w, item.UserId)
	if (item.UserId != 0) == false {
		w = w[:backupIndexUserId]
	}
	return append(w, '}')
}

func (item *Service3GetLastVisitTimestamp) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *Service3GetLastVisitTimestamp) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("service3.getLastVisitTimestamp", err.Error())
	}
	return nil
}
