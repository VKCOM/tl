// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlService1ReplaceOrIncr

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/cycle_6ca945392bbf8b14f24e5653edc8b214"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type Service1ReplaceOrIncr struct {
	Key   string
	Flags int32
	Delay int32
	Value int64
}

func (Service1ReplaceOrIncr) TLName() string { return "service1.replaceOrIncr" }
func (Service1ReplaceOrIncr) TLTag() uint32  { return 0x9d1bdcfd }

func (item *Service1ReplaceOrIncr) Reset() {
	item.Key = ""
	item.Flags = 0
	item.Delay = 0
	item.Value = 0
}

func (item *Service1ReplaceOrIncr) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.StringRead(w, &item.Key); err != nil {
		return w, err
	}
	if w, err = basictl.IntRead(w, &item.Flags); err != nil {
		return w, err
	}
	if w, err = basictl.IntRead(w, &item.Delay); err != nil {
		return w, err
	}
	return basictl.LongRead(w, &item.Value)
}

// This method is general version of Write, use it instead!
func (item *Service1ReplaceOrIncr) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *Service1ReplaceOrIncr) Write(w []byte) []byte {
	w = basictl.StringWrite(w, item.Key)
	w = basictl.IntWrite(w, item.Flags)
	w = basictl.IntWrite(w, item.Delay)
	w = basictl.LongWrite(w, item.Value)
	return w
}

func (item *Service1ReplaceOrIncr) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x9d1bdcfd); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *Service1ReplaceOrIncr) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *Service1ReplaceOrIncr) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x9d1bdcfd)
	return item.Write(w)
}

func (item *Service1ReplaceOrIncr) ReadResult(w []byte, ret *cycle_6ca945392bbf8b14f24e5653edc8b214.Service1Value) (_ []byte, err error) {
	return ret.ReadBoxed(w)
}

func (item *Service1ReplaceOrIncr) WriteResult(w []byte, ret cycle_6ca945392bbf8b14f24e5653edc8b214.Service1Value) (_ []byte, err error) {
	w = ret.WriteBoxed(w)
	return w, nil
}

func (item *Service1ReplaceOrIncr) ReadResultJSON(legacyTypeNames bool, in *basictl.JsonLexer, ret *cycle_6ca945392bbf8b14f24e5653edc8b214.Service1Value) error {
	if err := ret.ReadJSON(legacyTypeNames, in); err != nil {
		return err
	}
	return nil
}

func (item *Service1ReplaceOrIncr) WriteResultJSON(w []byte, ret cycle_6ca945392bbf8b14f24e5653edc8b214.Service1Value) (_ []byte, err error) {
	return item.writeResultJSON(true, false, w, ret)
}

func (item *Service1ReplaceOrIncr) writeResultJSON(newTypeNames bool, short bool, w []byte, ret cycle_6ca945392bbf8b14f24e5653edc8b214.Service1Value) (_ []byte, err error) {
	w = ret.WriteJSONOpt(newTypeNames, short, w)
	return w, nil
}

func (item *Service1ReplaceOrIncr) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret cycle_6ca945392bbf8b14f24e5653edc8b214.Service1Value
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *Service1ReplaceOrIncr) ReadResultWriteResultJSONOpt(newTypeNames bool, short bool, r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret cycle_6ca945392bbf8b14f24e5653edc8b214.Service1Value
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(newTypeNames, short, w, ret)
	return r, w, err
}

func (item *Service1ReplaceOrIncr) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	var ret cycle_6ca945392bbf8b14f24e5653edc8b214.Service1Value
	err := item.ReadResultJSON(true, &basictl.JsonLexer{Data: r}, &ret)
	if err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item *Service1ReplaceOrIncr) String() string {
	return string(item.WriteJSON(nil))
}

func (item *Service1ReplaceOrIncr) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propKeyPresented bool
	var propFlagsPresented bool
	var propDelayPresented bool
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
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service1.replaceOrIncr", "key")
				}
				if err := internal.Json2ReadString(in, &item.Key); err != nil {
					return err
				}
				propKeyPresented = true
			case "flags":
				if propFlagsPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service1.replaceOrIncr", "flags")
				}
				if err := internal.Json2ReadInt32(in, &item.Flags); err != nil {
					return err
				}
				propFlagsPresented = true
			case "delay":
				if propDelayPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service1.replaceOrIncr", "delay")
				}
				if err := internal.Json2ReadInt32(in, &item.Delay); err != nil {
					return err
				}
				propDelayPresented = true
			case "value":
				if propValuePresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service1.replaceOrIncr", "value")
				}
				if err := internal.Json2ReadInt64(in, &item.Value); err != nil {
					return err
				}
				propValuePresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("service1.replaceOrIncr", key)
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
	if !propFlagsPresented {
		item.Flags = 0
	}
	if !propDelayPresented {
		item.Delay = 0
	}
	if !propValuePresented {
		item.Value = 0
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *Service1ReplaceOrIncr) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *Service1ReplaceOrIncr) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *Service1ReplaceOrIncr) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexKey := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"key":`...)
	w = basictl.JSONWriteString(w, item.Key)
	if (len(item.Key) != 0) == false {
		w = w[:backupIndexKey]
	}
	backupIndexFlags := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"flags":`...)
	w = basictl.JSONWriteInt32(w, item.Flags)
	if (item.Flags != 0) == false {
		w = w[:backupIndexFlags]
	}
	backupIndexDelay := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"delay":`...)
	w = basictl.JSONWriteInt32(w, item.Delay)
	if (item.Delay != 0) == false {
		w = w[:backupIndexDelay]
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

func (item *Service1ReplaceOrIncr) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *Service1ReplaceOrIncr) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("service1.replaceOrIncr", err.Error())
	}
	return nil
}
