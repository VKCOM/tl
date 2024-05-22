// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlReplace10

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlBuiltinTupleTuple3Replace10Elem"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlReplace10Elem"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type Replace10 struct {
	N uint32
	A [][3]tlReplace10Elem.Replace10Elem
}

func (Replace10) TLName() string { return "replace10" }
func (Replace10) TLTag() uint32  { return 0xfc81f008 }

func (item *Replace10) Reset() {
	item.N = 0
	item.A = item.A[:0]
}

func (item *Replace10) FillRandom(rg *basictl.RandGenerator) {
	item.N = basictl.RandomUint(rg)
	item.N = rg.LimitValue(item.N)
	tlBuiltinTupleTuple3Replace10Elem.BuiltinTupleTuple3Replace10ElemFillRandom(rg, &item.A, item.N, item.N)
}

func (item *Replace10) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.N); err != nil {
		return w, err
	}
	return tlBuiltinTupleTuple3Replace10Elem.BuiltinTupleTuple3Replace10ElemRead(w, &item.A, item.N, item.N)
}

// This method is general version of Write, use it instead!
func (item *Replace10) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w)
}

func (item *Replace10) Write(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, item.N)
	if w, err = tlBuiltinTupleTuple3Replace10Elem.BuiltinTupleTuple3Replace10ElemWrite(w, item.A, item.N, item.N); err != nil {
		return w, err
	}
	return w, nil
}

func (item *Replace10) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xfc81f008); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *Replace10) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w)
}

func (item *Replace10) WriteBoxed(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, 0xfc81f008)
	return item.Write(w)
}

func (item Replace10) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func (item *Replace10) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propNPresented bool
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
			case "n":
				if propNPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("replace10", "n")
				}
				if err := internal.Json2ReadUint32(in, &item.N); err != nil {
					return err
				}
				propNPresented = true
			case "a":
				if rawA != nil {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("replace10", "a")
				}
				rawA = in.Raw()
				if !in.Ok() {
					return in.Error()
				}
			default:
				return internal.ErrorInvalidJSONExcessElement("replace10", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propNPresented {
		item.N = 0
	}
	var inAPointer *basictl.JsonLexer
	inA := basictl.JsonLexer{Data: rawA}
	if rawA != nil {
		inAPointer = &inA
	}
	if err := tlBuiltinTupleTuple3Replace10Elem.BuiltinTupleTuple3Replace10ElemReadJSON(legacyTypeNames, inAPointer, &item.A, item.N, item.N); err != nil {
		return err
	}

	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *Replace10) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w)
}

func (item *Replace10) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w)
}
func (item *Replace10) WriteJSONOpt(newTypeNames bool, short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	backupIndexN := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"n":`...)
	w = basictl.JSONWriteUint32(w, item.N)
	if (item.N != 0) == false {
		w = w[:backupIndexN]
	}
	backupIndexA := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"a":`...)
	if w, err = tlBuiltinTupleTuple3Replace10Elem.BuiltinTupleTuple3Replace10ElemWriteJSONOpt(newTypeNames, short, w, item.A, item.N, item.N); err != nil {
		return w, err
	}
	if (len(item.A) != 0) == false {
		w = w[:backupIndexA]
	}
	return append(w, '}'), nil
}

func (item *Replace10) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *Replace10) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("replace10", err.Error())
	}
	return nil
}
