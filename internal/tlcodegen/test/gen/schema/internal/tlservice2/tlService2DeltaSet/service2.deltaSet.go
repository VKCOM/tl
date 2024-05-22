// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlService2DeltaSet

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tlservice2/tlService2CounterSet"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tlservice2/tlService2ObjectId"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type Service2DeltaSet struct {
	Id       tlService2ObjectId.Service2ObjectId
	Counters tlService2CounterSet.Service2CounterSet
}

func (Service2DeltaSet) TLName() string { return "service2.deltaSet" }
func (Service2DeltaSet) TLTag() uint32  { return 0xbf49abc2 }

func (item *Service2DeltaSet) Reset() {
	item.Id.Reset()
	item.Counters.Reset()
}

func (item *Service2DeltaSet) Read(w []byte, nat_objectIdLength uint32, nat_intCountersNum uint32, nat_floatCountersNum uint32) (_ []byte, err error) {
	if w, err = item.Id.Read(w, nat_objectIdLength); err != nil {
		return w, err
	}
	return item.Counters.Read(w, nat_intCountersNum, nat_floatCountersNum)
}

// This method is general version of Write, use it instead!
func (item *Service2DeltaSet) WriteGeneral(w []byte, nat_objectIdLength uint32, nat_intCountersNum uint32, nat_floatCountersNum uint32) (_ []byte, err error) {
	return item.Write(w, nat_objectIdLength, nat_intCountersNum, nat_floatCountersNum)
}

func (item *Service2DeltaSet) Write(w []byte, nat_objectIdLength uint32, nat_intCountersNum uint32, nat_floatCountersNum uint32) (_ []byte, err error) {
	if w, err = item.Id.Write(w, nat_objectIdLength); err != nil {
		return w, err
	}
	if w, err = item.Counters.Write(w, nat_intCountersNum, nat_floatCountersNum); err != nil {
		return w, err
	}
	return w, nil
}

func (item *Service2DeltaSet) ReadBoxed(w []byte, nat_objectIdLength uint32, nat_intCountersNum uint32, nat_floatCountersNum uint32) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xbf49abc2); err != nil {
		return w, err
	}
	return item.Read(w, nat_objectIdLength, nat_intCountersNum, nat_floatCountersNum)
}

// This method is general version of WriteBoxed, use it instead!
func (item *Service2DeltaSet) WriteBoxedGeneral(w []byte, nat_objectIdLength uint32, nat_intCountersNum uint32, nat_floatCountersNum uint32) (_ []byte, err error) {
	return item.WriteBoxed(w, nat_objectIdLength, nat_intCountersNum, nat_floatCountersNum)
}

func (item *Service2DeltaSet) WriteBoxed(w []byte, nat_objectIdLength uint32, nat_intCountersNum uint32, nat_floatCountersNum uint32) (_ []byte, err error) {
	w = basictl.NatWrite(w, 0xbf49abc2)
	return item.Write(w, nat_objectIdLength, nat_intCountersNum, nat_floatCountersNum)
}

func (item *Service2DeltaSet) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, nat_objectIdLength uint32, nat_intCountersNum uint32, nat_floatCountersNum uint32) error {
	var rawId []byte
	var rawCounters []byte

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
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service2.deltaSet", "id")
				}
				rawId = in.Raw()
				if !in.Ok() {
					return in.Error()
				}
			case "counters":
				if rawCounters != nil {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service2.deltaSet", "counters")
				}
				rawCounters = in.Raw()
				if !in.Ok() {
					return in.Error()
				}
			default:
				return internal.ErrorInvalidJSONExcessElement("service2.deltaSet", key)
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
	if err := item.Id.ReadJSON(legacyTypeNames, inIdPointer, nat_objectIdLength); err != nil {
		return err
	}

	var inCountersPointer *basictl.JsonLexer
	inCounters := basictl.JsonLexer{Data: rawCounters}
	if rawCounters != nil {
		inCountersPointer = &inCounters
	}
	if err := item.Counters.ReadJSON(legacyTypeNames, inCountersPointer, nat_intCountersNum, nat_floatCountersNum); err != nil {
		return err
	}

	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *Service2DeltaSet) WriteJSONGeneral(w []byte, nat_objectIdLength uint32, nat_intCountersNum uint32, nat_floatCountersNum uint32) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w, nat_objectIdLength, nat_intCountersNum, nat_floatCountersNum)
}

func (item *Service2DeltaSet) WriteJSON(w []byte, nat_objectIdLength uint32, nat_intCountersNum uint32, nat_floatCountersNum uint32) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w, nat_objectIdLength, nat_intCountersNum, nat_floatCountersNum)
}
func (item *Service2DeltaSet) WriteJSONOpt(newTypeNames bool, short bool, w []byte, nat_objectIdLength uint32, nat_intCountersNum uint32, nat_floatCountersNum uint32) (_ []byte, err error) {
	w = append(w, '{')
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"id":`...)
	if w, err = item.Id.WriteJSONOpt(newTypeNames, short, w, nat_objectIdLength); err != nil {
		return w, err
	}
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"counters":`...)
	if w, err = item.Counters.WriteJSONOpt(newTypeNames, short, w, nat_intCountersNum, nat_floatCountersNum); err != nil {
		return w, err
	}
	return append(w, '}'), nil
}
