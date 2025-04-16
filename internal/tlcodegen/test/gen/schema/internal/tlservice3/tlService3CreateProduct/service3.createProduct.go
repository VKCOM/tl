// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlService3CreateProduct

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tl/tlBool"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tl/tlBuiltinVectorInt"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type Service3CreateProduct struct {
	UserId         int32
	Type           int32
	Id             []int32
	Info           []int32
	Date           int32
	ExpirationDate int32
}

func (Service3CreateProduct) TLName() string { return "service3.createProduct" }
func (Service3CreateProduct) TLTag() uint32  { return 0xb7d92bd9 }

func (item *Service3CreateProduct) Reset() {
	item.UserId = 0
	item.Type = 0
	item.Id = item.Id[:0]
	item.Info = item.Info[:0]
	item.Date = 0
	item.ExpirationDate = 0
}

func (item *Service3CreateProduct) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.IntRead(w, &item.UserId); err != nil {
		return w, err
	}
	if w, err = basictl.IntRead(w, &item.Type); err != nil {
		return w, err
	}
	if w, err = tlBuiltinVectorInt.BuiltinVectorIntRead(w, &item.Id); err != nil {
		return w, err
	}
	if w, err = tlBuiltinVectorInt.BuiltinVectorIntRead(w, &item.Info); err != nil {
		return w, err
	}
	if w, err = basictl.IntRead(w, &item.Date); err != nil {
		return w, err
	}
	return basictl.IntRead(w, &item.ExpirationDate)
}

// This method is general version of Write, use it instead!
func (item *Service3CreateProduct) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *Service3CreateProduct) Write(w []byte) []byte {
	w = basictl.IntWrite(w, item.UserId)
	w = basictl.IntWrite(w, item.Type)
	w = tlBuiltinVectorInt.BuiltinVectorIntWrite(w, item.Id)
	w = tlBuiltinVectorInt.BuiltinVectorIntWrite(w, item.Info)
	w = basictl.IntWrite(w, item.Date)
	w = basictl.IntWrite(w, item.ExpirationDate)
	return w
}

func (item *Service3CreateProduct) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xb7d92bd9); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *Service3CreateProduct) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *Service3CreateProduct) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0xb7d92bd9)
	return item.Write(w)
}

func (item *Service3CreateProduct) ReadResult(w []byte, ret *bool) (_ []byte, err error) {
	return tlBool.BoolReadBoxed(w, ret)
}

func (item *Service3CreateProduct) WriteResult(w []byte, ret bool) (_ []byte, err error) {
	w = tlBool.BoolWriteBoxed(w, ret)
	return w, nil
}

func (item *Service3CreateProduct) ReadResultJSON(legacyTypeNames bool, in *basictl.JsonLexer, ret *bool) error {
	if err := internal.Json2ReadBool(in, ret); err != nil {
		return err
	}
	return nil
}

func (item *Service3CreateProduct) WriteResultJSON(w []byte, ret bool) (_ []byte, err error) {
	return item.writeResultJSON(true, false, w, ret)
}

func (item *Service3CreateProduct) writeResultJSON(newTypeNames bool, short bool, w []byte, ret bool) (_ []byte, err error) {
	w = basictl.JSONWriteBool(w, ret)
	return w, nil
}

func (item *Service3CreateProduct) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret bool
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *Service3CreateProduct) ReadResultWriteResultJSONOpt(newTypeNames bool, short bool, r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret bool
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(newTypeNames, short, w, ret)
	return r, w, err
}

func (item *Service3CreateProduct) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	var ret bool
	err := item.ReadResultJSON(true, &basictl.JsonLexer{Data: r}, &ret)
	if err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item *Service3CreateProduct) String() string {
	return string(item.WriteJSON(nil))
}

func (item *Service3CreateProduct) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propUserIdPresented bool
	var propTypePresented bool
	var propIdPresented bool
	var propInfoPresented bool
	var propDatePresented bool
	var propExpirationDatePresented bool

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
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service3.createProduct", "user_id")
				}
				if err := internal.Json2ReadInt32(in, &item.UserId); err != nil {
					return err
				}
				propUserIdPresented = true
			case "type":
				if propTypePresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service3.createProduct", "type")
				}
				if err := internal.Json2ReadInt32(in, &item.Type); err != nil {
					return err
				}
				propTypePresented = true
			case "id":
				if propIdPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service3.createProduct", "id")
				}
				if err := tlBuiltinVectorInt.BuiltinVectorIntReadJSON(legacyTypeNames, in, &item.Id); err != nil {
					return err
				}
				propIdPresented = true
			case "info":
				if propInfoPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service3.createProduct", "info")
				}
				if err := tlBuiltinVectorInt.BuiltinVectorIntReadJSON(legacyTypeNames, in, &item.Info); err != nil {
					return err
				}
				propInfoPresented = true
			case "date":
				if propDatePresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service3.createProduct", "date")
				}
				if err := internal.Json2ReadInt32(in, &item.Date); err != nil {
					return err
				}
				propDatePresented = true
			case "expiration_date":
				if propExpirationDatePresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service3.createProduct", "expiration_date")
				}
				if err := internal.Json2ReadInt32(in, &item.ExpirationDate); err != nil {
					return err
				}
				propExpirationDatePresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("service3.createProduct", key)
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
	if !propInfoPresented {
		item.Info = item.Info[:0]
	}
	if !propDatePresented {
		item.Date = 0
	}
	if !propExpirationDatePresented {
		item.ExpirationDate = 0
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *Service3CreateProduct) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *Service3CreateProduct) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *Service3CreateProduct) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
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
	backupIndexInfo := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"info":`...)
	w = tlBuiltinVectorInt.BuiltinVectorIntWriteJSONOpt(newTypeNames, short, w, item.Info)
	if (len(item.Info) != 0) == false {
		w = w[:backupIndexInfo]
	}
	backupIndexDate := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"date":`...)
	w = basictl.JSONWriteInt32(w, item.Date)
	if (item.Date != 0) == false {
		w = w[:backupIndexDate]
	}
	backupIndexExpirationDate := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"expiration_date":`...)
	w = basictl.JSONWriteInt32(w, item.ExpirationDate)
	if (item.ExpirationDate != 0) == false {
		w = w[:backupIndexExpirationDate]
	}
	return append(w, '}')
}

func (item *Service3CreateProduct) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *Service3CreateProduct) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("service3.createProduct", err.Error())
	}
	return nil
}
