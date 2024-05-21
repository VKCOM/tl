// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlReplace1

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlBuiltinTupleInt"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type Replace1 struct {
	A []int32
}

func (Replace1) TLName() string { return "replace1" }
func (Replace1) TLTag() uint32  { return 0x89eac43a }

func (item *Replace1) Reset() {
	item.A = item.A[:0]
}

func (item *Replace1) FillRandom(rg *basictl.RandGenerator, nat_n uint32) {
	tlBuiltinTupleInt.BuiltinTupleIntFillRandom(rg, &item.A, nat_n)
}

func (item *Replace1) Read(w []byte, nat_n uint32) (_ []byte, err error) {
	return tlBuiltinTupleInt.BuiltinTupleIntRead(w, &item.A, nat_n)
}

// This method is general version of Write, use it instead!
func (item *Replace1) WriteGeneral(w []byte, nat_n uint32) (_ []byte, err error) {
	return item.Write(w, nat_n)
}

func (item *Replace1) Write(w []byte, nat_n uint32) (_ []byte, err error) {
	if w, err = tlBuiltinTupleInt.BuiltinTupleIntWrite(w, item.A, nat_n); err != nil {
		return w, err
	}
	return w, nil
}

func (item *Replace1) ReadBoxed(w []byte, nat_n uint32) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x89eac43a); err != nil {
		return w, err
	}
	return item.Read(w, nat_n)
}

// This method is general version of WriteBoxed, use it instead!
func (item *Replace1) WriteBoxedGeneral(w []byte, nat_n uint32) (_ []byte, err error) {
	return item.WriteBoxed(w, nat_n)
}

func (item *Replace1) WriteBoxed(w []byte, nat_n uint32) (_ []byte, err error) {
	w = basictl.NatWrite(w, 0x89eac43a)
	return item.Write(w, nat_n)
}

func (item *Replace1) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, nat_n uint32) error {
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
					return internal.ErrorInvalidJSONWithDuplicatingKeys("replace1", "a")
				}
				rawA = in.Raw()
				if !in.Ok() {
					return in.Error()
				}
			default:
				return internal.ErrorInvalidJSONExcessElement("replace1", key)
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
	if err := tlBuiltinTupleInt.BuiltinTupleIntReadJSON(legacyTypeNames, inAPointer, &item.A, nat_n); err != nil {
		return err
	}

	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *Replace1) WriteJSONGeneral(w []byte, nat_n uint32) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w, nat_n)
}

func (item *Replace1) WriteJSON(w []byte, nat_n uint32) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w, nat_n)
}
func (item *Replace1) WriteJSONOpt(newTypeNames bool, short bool, w []byte, nat_n uint32) (_ []byte, err error) {
	w = append(w, '{')
	backupIndexA := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"a":`...)
	if w, err = tlBuiltinTupleInt.BuiltinTupleIntWriteJSONOpt(newTypeNames, short, w, item.A, nat_n); err != nil {
		return w, err
	}
	if (len(item.A) != 0) == false {
		w = w[:backupIndexA]
	}
	return append(w, '}'), nil
}