// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlService1Get

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/cycle_6ca945392bbf8b14f24e5653edc8b214"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type Service1Get struct {
	Key string
}

func (Service1Get) TLName() string { return "service1.get" }
func (Service1Get) TLTag() uint32  { return 0x29099b19 }

func (item *Service1Get) Reset() {
	item.Key = ""
}

func (item *Service1Get) Read(w []byte) (_ []byte, err error) {
	return basictl.StringRead(w, &item.Key)
}

// This method is general version of Write, use it instead!
func (item *Service1Get) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *Service1Get) Write(w []byte) []byte {
	w = basictl.StringWrite(w, item.Key)
	return w
}

func (item *Service1Get) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x29099b19); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *Service1Get) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *Service1Get) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x29099b19)
	return item.Write(w)
}

func (item *Service1Get) ReadResult(w []byte, ret *cycle_6ca945392bbf8b14f24e5653edc8b214.Service1Value) (_ []byte, err error) {
	return ret.ReadBoxed(w)
}

func (item *Service1Get) WriteResult(w []byte, ret cycle_6ca945392bbf8b14f24e5653edc8b214.Service1Value) (_ []byte, err error) {
	w = ret.WriteBoxed(w)
	return w, nil
}

func (item *Service1Get) ReadResultJSON(legacyTypeNames bool, in *basictl.JsonLexer, ret *cycle_6ca945392bbf8b14f24e5653edc8b214.Service1Value) error {
	if err := ret.ReadJSON(legacyTypeNames, in); err != nil {
		return err
	}
	return nil
}

func (item *Service1Get) WriteResultJSON(w []byte, ret cycle_6ca945392bbf8b14f24e5653edc8b214.Service1Value) (_ []byte, err error) {
	return item.writeResultJSON(true, false, w, ret)
}

func (item *Service1Get) writeResultJSON(newTypeNames bool, short bool, w []byte, ret cycle_6ca945392bbf8b14f24e5653edc8b214.Service1Value) (_ []byte, err error) {
	w = ret.WriteJSONOpt(newTypeNames, short, w)
	return w, nil
}

func (item *Service1Get) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret cycle_6ca945392bbf8b14f24e5653edc8b214.Service1Value
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *Service1Get) ReadResultWriteResultJSONOpt(newTypeNames bool, short bool, r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret cycle_6ca945392bbf8b14f24e5653edc8b214.Service1Value
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(newTypeNames, short, w, ret)
	return r, w, err
}

func (item *Service1Get) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	var ret cycle_6ca945392bbf8b14f24e5653edc8b214.Service1Value
	err := item.ReadResultJSON(true, &basictl.JsonLexer{Data: r}, &ret)
	if err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item *Service1Get) String() string {
	return string(item.WriteJSON(nil))
}

func (item *Service1Get) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propKeyPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "key":
				if propKeyPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service1.get", "key")
				}
				if err := internal.Json2ReadString(in, &item.Key); err != nil {
					return err
				}
				propKeyPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("service1.get", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propKeyPresented {
		item.Key = ""
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *Service1Get) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *Service1Get) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *Service1Get) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexKey := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"key":`...)
	w = basictl.JSONWriteString(w, item.Key)
	if (len(item.Key) != 0) == false {
		w = w[:backupIndexKey]
	}
	return append(w, '}')
}

func (item *Service1Get) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *Service1Get) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("service1.get", err.Error())
	}
	return nil
}
