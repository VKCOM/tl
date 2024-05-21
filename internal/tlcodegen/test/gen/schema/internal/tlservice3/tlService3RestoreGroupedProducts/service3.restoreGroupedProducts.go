// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlService3RestoreGroupedProducts

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tl/tlBool"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tl/tlBuiltinVectorInt"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type Service3RestoreGroupedProducts struct {
	UserId    int32
	Type      int32
	Id        []int32
	StartDate int32
	EndDate   int32
}

func (Service3RestoreGroupedProducts) TLName() string { return "service3.restoreGroupedProducts" }
func (Service3RestoreGroupedProducts) TLTag() uint32  { return 0x1f17bfac }

func (item *Service3RestoreGroupedProducts) Reset() {
	item.UserId = 0
	item.Type = 0
	item.Id = item.Id[:0]
	item.StartDate = 0
	item.EndDate = 0
}

func (item *Service3RestoreGroupedProducts) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.IntRead(w, &item.UserId); err != nil {
		return w, err
	}
	if w, err = basictl.IntRead(w, &item.Type); err != nil {
		return w, err
	}
	if w, err = tlBuiltinVectorInt.BuiltinVectorIntRead(w, &item.Id); err != nil {
		return w, err
	}
	if w, err = basictl.IntRead(w, &item.StartDate); err != nil {
		return w, err
	}
	return basictl.IntRead(w, &item.EndDate)
}

// This method is general version of Write, use it instead!
func (item *Service3RestoreGroupedProducts) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *Service3RestoreGroupedProducts) Write(w []byte) []byte {
	w = basictl.IntWrite(w, item.UserId)
	w = basictl.IntWrite(w, item.Type)
	w = tlBuiltinVectorInt.BuiltinVectorIntWrite(w, item.Id)
	w = basictl.IntWrite(w, item.StartDate)
	w = basictl.IntWrite(w, item.EndDate)
	return w
}

func (item *Service3RestoreGroupedProducts) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x1f17bfac); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *Service3RestoreGroupedProducts) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *Service3RestoreGroupedProducts) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x1f17bfac)
	return item.Write(w)
}

func (item *Service3RestoreGroupedProducts) ReadResult(w []byte, ret *bool) (_ []byte, err error) {
	return tlBool.BoolReadBoxed(w, ret)
}

func (item *Service3RestoreGroupedProducts) WriteResult(w []byte, ret bool) (_ []byte, err error) {
	w = tlBool.BoolWriteBoxed(w, ret)
	return w, nil
}

func (item *Service3RestoreGroupedProducts) ReadResultJSON(legacyTypeNames bool, in *basictl.JsonLexer, ret *bool) error {
	if err := internal.Json2ReadBool(in, ret); err != nil {
		return err
	}
	return nil
}

func (item *Service3RestoreGroupedProducts) WriteResultJSON(w []byte, ret bool) (_ []byte, err error) {
	return item.writeResultJSON(true, false, w, ret)
}

func (item *Service3RestoreGroupedProducts) writeResultJSON(newTypeNames bool, short bool, w []byte, ret bool) (_ []byte, err error) {
	w = basictl.JSONWriteBool(w, ret)
	return w, nil
}

func (item *Service3RestoreGroupedProducts) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret bool
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *Service3RestoreGroupedProducts) ReadResultWriteResultJSONOpt(newTypeNames bool, short bool, r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret bool
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(newTypeNames, short, w, ret)
	return r, w, err
}

func (item *Service3RestoreGroupedProducts) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	var ret bool
	err := item.ReadResultJSON(true, &basictl.JsonLexer{Data: r}, &ret)
	if err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item Service3RestoreGroupedProducts) String() string {
	return string(item.WriteJSON(nil))
}

func (item *Service3RestoreGroupedProducts) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propUserIdPresented bool
	var propTypePresented bool
	var propIdPresented bool
	var propStartDatePresented bool
	var propEndDatePresented bool

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
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service3.restoreGroupedProducts", "user_id")
				}
				if err := internal.Json2ReadInt32(in, &item.UserId); err != nil {
					return err
				}
				propUserIdPresented = true
			case "type":
				if propTypePresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service3.restoreGroupedProducts", "type")
				}
				if err := internal.Json2ReadInt32(in, &item.Type); err != nil {
					return err
				}
				propTypePresented = true
			case "id":
				if propIdPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service3.restoreGroupedProducts", "id")
				}
				if err := tlBuiltinVectorInt.BuiltinVectorIntReadJSON(legacyTypeNames, in, &item.Id); err != nil {
					return err
				}
				propIdPresented = true
			case "start_date":
				if propStartDatePresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service3.restoreGroupedProducts", "start_date")
				}
				if err := internal.Json2ReadInt32(in, &item.StartDate); err != nil {
					return err
				}
				propStartDatePresented = true
			case "end_date":
				if propEndDatePresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service3.restoreGroupedProducts", "end_date")
				}
				if err := internal.Json2ReadInt32(in, &item.EndDate); err != nil {
					return err
				}
				propEndDatePresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("service3.restoreGroupedProducts", key)
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
	if !propTypePresented {
		item.Type = 0
	}
	if !propIdPresented {
		item.Id = item.Id[:0]
	}
	if !propStartDatePresented {
		item.StartDate = 0
	}
	if !propEndDatePresented {
		item.EndDate = 0
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *Service3RestoreGroupedProducts) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *Service3RestoreGroupedProducts) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *Service3RestoreGroupedProducts) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexUserId := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"user_id":`...)
	w = basictl.JSONWriteInt32(w, item.UserId)
	if (item.UserId != 0) == false {
		w = w[:backupIndexUserId]
	}
	backupIndexType := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"type":`...)
	w = basictl.JSONWriteInt32(w, item.Type)
	if (item.Type != 0) == false {
		w = w[:backupIndexType]
	}
	backupIndexId := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"id":`...)
	w = tlBuiltinVectorInt.BuiltinVectorIntWriteJSONOpt(newTypeNames, short, w, item.Id)
	if (len(item.Id) != 0) == false {
		w = w[:backupIndexId]
	}
	backupIndexStartDate := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"start_date":`...)
	w = basictl.JSONWriteInt32(w, item.StartDate)
	if (item.StartDate != 0) == false {
		w = w[:backupIndexStartDate]
	}
	backupIndexEndDate := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"end_date":`...)
	w = basictl.JSONWriteInt32(w, item.EndDate)
	if (item.EndDate != 0) == false {
		w = w[:backupIndexEndDate]
	}
	return append(w, '}')
}

func (item *Service3RestoreGroupedProducts) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *Service3RestoreGroupedProducts) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("service3.restoreGroupedProducts", err.Error())
	}
	return nil
}