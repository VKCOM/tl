// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlService2ObjectId

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tl/tlBuiltinTupleInt"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type Service2ObjectId struct {
	Id []int32
}

func (Service2ObjectId) TLName() string { return "service2.objectId" }
func (Service2ObjectId) TLTag() uint32  { return 0xaa0af282 }

func (item *Service2ObjectId) Reset() {
	item.Id = item.Id[:0]
}

func (item *Service2ObjectId) Read(w []byte, nat_objectIdLength uint32) (_ []byte, err error) {
	return tlBuiltinTupleInt.BuiltinTupleIntRead(w, &item.Id, nat_objectIdLength)
}

// This method is general version of Write, use it instead!
func (item *Service2ObjectId) WriteGeneral(w []byte, nat_objectIdLength uint32) (_ []byte, err error) {
	return item.Write(w, nat_objectIdLength)
}

func (item *Service2ObjectId) Write(w []byte, nat_objectIdLength uint32) (_ []byte, err error) {
	if w, err = tlBuiltinTupleInt.BuiltinTupleIntWrite(w, item.Id, nat_objectIdLength); err != nil {
		return w, err
	}
	return w, nil
}

func (item *Service2ObjectId) ReadBoxed(w []byte, nat_objectIdLength uint32) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xaa0af282); err != nil {
		return w, err
	}
	return item.Read(w, nat_objectIdLength)
}

// This method is general version of WriteBoxed, use it instead!
func (item *Service2ObjectId) WriteBoxedGeneral(w []byte, nat_objectIdLength uint32) (_ []byte, err error) {
	return item.WriteBoxed(w, nat_objectIdLength)
}

func (item *Service2ObjectId) WriteBoxed(w []byte, nat_objectIdLength uint32) (_ []byte, err error) {
	w = basictl.NatWrite(w, 0xaa0af282)
	return item.Write(w, nat_objectIdLength)
}

func (item *Service2ObjectId) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, nat_objectIdLength uint32) error {
	var rawId []byte

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "id":
				if rawId != nil {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service2.objectId", "id")
				}
				rawId = in.Raw()
				if !in.Ok() {
					return in.Error()
				}
			default:
				return internal.ErrorInvalidJSONExcessElement("service2.objectId", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	var inIdPointer *basictl.JsonLexer
	inId := basictl.JsonLexer{Data: rawId}
	if rawId != nil {
		inIdPointer = &inId
	}
	if err := tlBuiltinTupleInt.BuiltinTupleIntReadJSON(legacyTypeNames, inIdPointer, &item.Id, nat_objectIdLength); err != nil {
		return err
	}

	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *Service2ObjectId) WriteJSONGeneral(w []byte, nat_objectIdLength uint32) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w, nat_objectIdLength)
}

func (item *Service2ObjectId) WriteJSON(w []byte, nat_objectIdLength uint32) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w, nat_objectIdLength)
}
func (item *Service2ObjectId) WriteJSONOpt(newTypeNames bool, short bool, w []byte, nat_objectIdLength uint32) (_ []byte, err error) {
	w = append(w, '{')
	backupIndexId := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"id":`...)
	if w, err = tlBuiltinTupleInt.BuiltinTupleIntWriteJSONOpt(newTypeNames, short, w, item.Id, nat_objectIdLength); err != nil {
		return w, err
	}
	if (len(item.Id) != 0) == false {
		w = w[:backupIndexId]
	}
	return append(w, '}'), nil
}