// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlService1Decr

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/cycle_6ca945392bbf8b14f24e5653edc8b214"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type Service1Decr struct {
	Key   string
	Value int64
}

func (Service1Decr) TLName() string { return "service1.decr" }
func (Service1Decr) TLTag() uint32  { return 0xeb179ce7 }

func (item *Service1Decr) Reset() {
	item.Key = ""
	item.Value = 0
}

func (item *Service1Decr) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.StringRead(w, &item.Key); err != nil {
		return w, err
	}
	return basictl.LongRead(w, &item.Value)
}

// This method is general version of Write, use it instead!
func (item *Service1Decr) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *Service1Decr) Write(w []byte) []byte {
	w = basictl.StringWrite(w, item.Key)
	w = basictl.LongWrite(w, item.Value)
	return w
}

func (item *Service1Decr) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xeb179ce7); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *Service1Decr) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *Service1Decr) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0xeb179ce7)
	return item.Write(w)
}

func (item *Service1Decr) ReadResult(w []byte, ret *cycle_6ca945392bbf8b14f24e5653edc8b214.Service1Value) (_ []byte, err error) {
	return ret.ReadBoxed(w)
}

func (item *Service1Decr) WriteResult(w []byte, ret cycle_6ca945392bbf8b14f24e5653edc8b214.Service1Value) (_ []byte, err error) {
	w = ret.WriteBoxed(w)
	return w, nil
}

func (item *Service1Decr) ReadResultJSON(legacyTypeNames bool, in *basictl.JsonLexer, ret *cycle_6ca945392bbf8b14f24e5653edc8b214.Service1Value) error {
	if err := ret.ReadJSON(legacyTypeNames, in); err != nil {
		return err
	}
	return nil
}

func (item *Service1Decr) WriteResultJSON(w []byte, ret cycle_6ca945392bbf8b14f24e5653edc8b214.Service1Value) (_ []byte, err error) {
	return item.writeResultJSON(true, false, w, ret)
}

func (item *Service1Decr) writeResultJSON(newTypeNames bool, short bool, w []byte, ret cycle_6ca945392bbf8b14f24e5653edc8b214.Service1Value) (_ []byte, err error) {
	w = ret.WriteJSONOpt(newTypeNames, short, w)
	return w, nil
}

func (item *Service1Decr) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret cycle_6ca945392bbf8b14f24e5653edc8b214.Service1Value
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *Service1Decr) ReadResultWriteResultJSONOpt(newTypeNames bool, short bool, r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret cycle_6ca945392bbf8b14f24e5653edc8b214.Service1Value
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(newTypeNames, short, w, ret)
	return r, w, err
}

func (item *Service1Decr) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	var ret cycle_6ca945392bbf8b14f24e5653edc8b214.Service1Value
	err := item.ReadResultJSON(true, &basictl.JsonLexer{Data: r}, &ret)
	if err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item Service1Decr) String() string {
	return string(item.WriteJSON(nil))
}

func (item *Service1Decr) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propKeyPresented bool
	var propValuePresented bool

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
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service1.decr", "key")
				}
				if err := internal.Json2ReadString(in, &item.Key); err != nil {
					return err
				}
				propKeyPresented = true
			case "value":
				if propValuePresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service1.decr", "value")
				}
				if err := internal.Json2ReadInt64(in, &item.Value); err != nil {
					return err
				}
				propValuePresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("service1.decr", key)
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
	if !propValuePresented {
		item.Value = 0
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *Service1Decr) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *Service1Decr) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *Service1Decr) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexKey := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"key":`...)
	w = basictl.JSONWriteString(w, item.Key)
	if (len(item.Key) != 0) == false {
		w = w[:backupIndexKey]
	}
	backupIndexValue := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"value":`...)
	w = basictl.JSONWriteInt64(w, item.Value)
	if (item.Value != 0) == false {
		w = w[:backupIndexValue]
	}
	return append(w, '}')
}

func (item *Service1Decr) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *Service1Decr) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("service1.decr", err.Error())
	}
	return nil
}
