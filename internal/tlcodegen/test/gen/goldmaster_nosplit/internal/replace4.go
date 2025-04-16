// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package internal

import (
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite

type Replace4 struct {
	A []int32
}

func (Replace4) TLName() string { return "replace4" }
func (Replace4) TLTag() uint32  { return 0x87995fb4 }

func (item *Replace4) Reset() {
	item.A = item.A[:0]
}

func (item *Replace4) FillRandom(rg *basictl.RandGenerator, nat_n uint32) {
	BuiltinTupleIntFillRandom(rg, &item.A, nat_n)
}

func (item *Replace4) Read(w []byte, nat_n uint32) (_ []byte, err error) {
	return BuiltinTupleIntRead(w, &item.A, nat_n)
}

// This method is general version of Write, use it instead!
func (item *Replace4) WriteGeneral(w []byte, nat_n uint32) (_ []byte, err error) {
	return item.Write(w, nat_n)
}

func (item *Replace4) Write(w []byte, nat_n uint32) (_ []byte, err error) {
	if w, err = BuiltinTupleIntWrite(w, item.A, nat_n); err != nil {
		return w, err
	}
	return w, nil
}

func (item *Replace4) ReadBoxed(w []byte, nat_n uint32) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x87995fb4); err != nil {
		return w, err
	}
	return item.Read(w, nat_n)
}

// This method is general version of WriteBoxed, use it instead!
func (item *Replace4) WriteBoxedGeneral(w []byte, nat_n uint32) (_ []byte, err error) {
	return item.WriteBoxed(w, nat_n)
}

func (item *Replace4) WriteBoxed(w []byte, nat_n uint32) (_ []byte, err error) {
	w = basictl.NatWrite(w, 0x87995fb4)
	return item.Write(w, nat_n)
}

func (item *Replace4) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, nat_n uint32) error {
	var rawA []byte

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "a":
				if rawA != nil {
					return ErrorInvalidJSONWithDuplicatingKeys("replace4", "a")
				}
				rawA = in.Raw()
				if !in.Ok() {
					return in.Error()
				}
			default:
				return ErrorInvalidJSONExcessElement("replace4", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	var inAPointer *basictl.JsonLexer
	inA := basictl.JsonLexer{Data: rawA}
	if rawA != nil {
		inAPointer = &inA
	}
	if err := BuiltinTupleIntReadJSON(legacyTypeNames, inAPointer, &item.A, nat_n); err != nil {
		return err
	}

	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *Replace4) WriteJSONGeneral(w []byte, nat_n uint32) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w, nat_n)
}

func (item *Replace4) WriteJSON(w []byte, nat_n uint32) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w, nat_n)
}
func (item *Replace4) WriteJSONOpt(newTypeNames bool, short bool, w []byte, nat_n uint32) (_ []byte, err error) {
	w = append(w, '{')
	backupIndexA := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"a":`...)
	if w, err = BuiltinTupleIntWriteJSONOpt(newTypeNames, short, w, item.A, nat_n); err != nil {
		return w, err
	}
	if (len(item.A) != 0) == false {
		w = w[:backupIndexA]
	}
	return append(w, '}'), nil
}

type Replace43 struct {
	A [3]int32
}

func (Replace43) TLName() string { return "replace4" }
func (Replace43) TLTag() uint32  { return 0x87995fb4 }

func (item *Replace43) Reset() {
	BuiltinTuple3IntReset(&item.A)
}

func (item *Replace43) FillRandom(rg *basictl.RandGenerator) {
	BuiltinTuple3IntFillRandom(rg, &item.A)
}

func (item *Replace43) Read(w []byte) (_ []byte, err error) {
	return BuiltinTuple3IntRead(w, &item.A)
}

// This method is general version of Write, use it instead!
func (item *Replace43) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *Replace43) Write(w []byte) []byte {
	w = BuiltinTuple3IntWrite(w, &item.A)
	return w
}

func (item *Replace43) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x87995fb4); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *Replace43) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *Replace43) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x87995fb4)
	return item.Write(w)
}

func (item Replace43) String() string {
	return string(item.WriteJSON(nil))
}

func (item *Replace43) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propAPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "a":
				if propAPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("replace4", "a")
				}
				if err := BuiltinTuple3IntReadJSON(legacyTypeNames, in, &item.A); err != nil {
					return err
				}
				propAPresented = true
			default:
				return ErrorInvalidJSONExcessElement("replace4", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propAPresented {
		BuiltinTuple3IntReset(&item.A)
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *Replace43) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *Replace43) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *Replace43) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"a":`...)
	w = BuiltinTuple3IntWriteJSONOpt(newTypeNames, short, w, &item.A)
	return append(w, '}')
}

func (item *Replace43) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *Replace43) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("replace4", err.Error())
	}
	return nil
}
