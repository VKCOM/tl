// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlService5Params

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type Service5Params struct {
	FieldsMask             uint32
	MaxExecutionSpeed      int32 // Conditional: item.FieldsMask.0
	MaxExecutionSpeedBytes int32 // Conditional: item.FieldsMask.1
}

func (Service5Params) TLName() string { return "service5.params" }
func (Service5Params) TLTag() uint32  { return 0x12ae5cb5 }

func (item *Service5Params) SetMaxExecutionSpeed(v int32) {
	item.MaxExecutionSpeed = v
	item.FieldsMask |= 1 << 0
}
func (item *Service5Params) ClearMaxExecutionSpeed() {
	item.MaxExecutionSpeed = 0
	item.FieldsMask &^= 1 << 0
}
func (item Service5Params) IsSetMaxExecutionSpeed() bool { return item.FieldsMask&(1<<0) != 0 }

func (item *Service5Params) SetMaxExecutionSpeedBytes(v int32) {
	item.MaxExecutionSpeedBytes = v
	item.FieldsMask |= 1 << 1
}
func (item *Service5Params) ClearMaxExecutionSpeedBytes() {
	item.MaxExecutionSpeedBytes = 0
	item.FieldsMask &^= 1 << 1
}
func (item Service5Params) IsSetMaxExecutionSpeedBytes() bool { return item.FieldsMask&(1<<1) != 0 }

func (item *Service5Params) Reset() {
	item.FieldsMask = 0
	item.MaxExecutionSpeed = 0
	item.MaxExecutionSpeedBytes = 0
}

func (item *Service5Params) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.FieldsMask); err != nil {
		return w, err
	}
	if item.FieldsMask&(1<<0) != 0 {
		if w, err = basictl.IntRead(w, &item.MaxExecutionSpeed); err != nil {
			return w, err
		}
	} else {
		item.MaxExecutionSpeed = 0
	}
	if item.FieldsMask&(1<<1) != 0 {
		if w, err = basictl.IntRead(w, &item.MaxExecutionSpeedBytes); err != nil {
			return w, err
		}
	} else {
		item.MaxExecutionSpeedBytes = 0
	}
	return w, nil
}

// This method is general version of Write, use it instead!
func (item *Service5Params) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *Service5Params) Write(w []byte) []byte {
	w = basictl.NatWrite(w, item.FieldsMask)
	if item.FieldsMask&(1<<0) != 0 {
		w = basictl.IntWrite(w, item.MaxExecutionSpeed)
	}
	if item.FieldsMask&(1<<1) != 0 {
		w = basictl.IntWrite(w, item.MaxExecutionSpeedBytes)
	}
	return w
}

func (item *Service5Params) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x12ae5cb5); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *Service5Params) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *Service5Params) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x12ae5cb5)
	return item.Write(w)
}

func (item Service5Params) String() string {
	return string(item.WriteJSON(nil))
}

func (item *Service5Params) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propFieldsMaskPresented bool
	var propMaxExecutionSpeedPresented bool
	var propMaxExecutionSpeedBytesPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "fields_mask":
				if propFieldsMaskPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service5.params", "fields_mask")
				}
				if err := internal.Json2ReadUint32(in, &item.FieldsMask); err != nil {
					return err
				}
				propFieldsMaskPresented = true
			case "max_execution_speed":
				if propMaxExecutionSpeedPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service5.params", "max_execution_speed")
				}
				if err := internal.Json2ReadInt32(in, &item.MaxExecutionSpeed); err != nil {
					return err
				}
				propMaxExecutionSpeedPresented = true
			case "max_execution_speed_bytes":
				if propMaxExecutionSpeedBytesPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service5.params", "max_execution_speed_bytes")
				}
				if err := internal.Json2ReadInt32(in, &item.MaxExecutionSpeedBytes); err != nil {
					return err
				}
				propMaxExecutionSpeedBytesPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("service5.params", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propFieldsMaskPresented {
		item.FieldsMask = 0
	}
	if !propMaxExecutionSpeedPresented {
		item.MaxExecutionSpeed = 0
	}
	if !propMaxExecutionSpeedBytesPresented {
		item.MaxExecutionSpeedBytes = 0
	}
	if propMaxExecutionSpeedPresented {
		item.FieldsMask |= 1 << 0
	}
	if propMaxExecutionSpeedBytesPresented {
		item.FieldsMask |= 1 << 1
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *Service5Params) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *Service5Params) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *Service5Params) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexFieldsMask := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"fields_mask":`...)
	w = basictl.JSONWriteUint32(w, item.FieldsMask)
	if (item.FieldsMask != 0) == false {
		w = w[:backupIndexFieldsMask]
	}
	if item.FieldsMask&(1<<0) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"max_execution_speed":`...)
		w = basictl.JSONWriteInt32(w, item.MaxExecutionSpeed)
	}
	if item.FieldsMask&(1<<1) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"max_execution_speed_bytes":`...)
		w = basictl.JSONWriteInt32(w, item.MaxExecutionSpeedBytes)
	}
	return append(w, '}')
}

func (item *Service5Params) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *Service5Params) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("service5.params", err.Error())
	}
	return nil
}
