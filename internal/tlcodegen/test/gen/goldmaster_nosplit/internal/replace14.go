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

type Replace14Long struct {
	K uint32
	A []Replace14ElemLong
}

func (Replace14Long) TLName() string { return "replace14" }
func (Replace14Long) TLTag() uint32  { return 0xb9801f9 }

func (item *Replace14Long) Reset() {
	item.K = 0
	item.A = item.A[:0]
}

func (item *Replace14Long) FillRandom(rg *basictl.RandGenerator, nat_n uint32) {
	var maskK uint32
	maskK = basictl.RandomUint(rg)
	maskK = rg.LimitValue(maskK)
	item.K = 0
	if maskK&(1<<0) != 0 {
		item.K |= (1 << 0)
	}
	BuiltinTupleReplace14ElemLongFillRandom(rg, &item.A, item.K, nat_n, item.K)
}

func (item *Replace14Long) Read(w []byte, nat_n uint32) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.K); err != nil {
		return w, err
	}
	return BuiltinTupleReplace14ElemLongRead(w, &item.A, item.K, nat_n, item.K)
}

// This method is general version of Write, use it instead!
func (item *Replace14Long) WriteGeneral(w []byte, nat_n uint32) (_ []byte, err error) {
	return item.Write(w, nat_n)
}

func (item *Replace14Long) Write(w []byte, nat_n uint32) (_ []byte, err error) {
	w = basictl.NatWrite(w, item.K)
	if w, err = BuiltinTupleReplace14ElemLongWrite(w, item.A, item.K, nat_n, item.K); err != nil {
		return w, err
	}
	return w, nil
}

func (item *Replace14Long) ReadBoxed(w []byte, nat_n uint32) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xb9801f9); err != nil {
		return w, err
	}
	return item.Read(w, nat_n)
}

// This method is general version of WriteBoxed, use it instead!
func (item *Replace14Long) WriteBoxedGeneral(w []byte, nat_n uint32) (_ []byte, err error) {
	return item.WriteBoxed(w, nat_n)
}

func (item *Replace14Long) WriteBoxed(w []byte, nat_n uint32) (_ []byte, err error) {
	w = basictl.NatWrite(w, 0xb9801f9)
	return item.Write(w, nat_n)
}

func (item *Replace14Long) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, nat_n uint32) error {
	var propKPresented bool
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
			case "k":
				if propKPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("replace14", "k")
				}
				if err := Json2ReadUint32(in, &item.K); err != nil {
					return err
				}
				propKPresented = true
			case "a":
				if rawA != nil {
					return ErrorInvalidJSONWithDuplicatingKeys("replace14", "a")
				}
				rawA = in.Raw()
				if !in.Ok() {
					return in.Error()
				}
			default:
				return ErrorInvalidJSONExcessElement("replace14", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propKPresented {
		item.K = 0
	}
	var inAPointer *basictl.JsonLexer
	inA := basictl.JsonLexer{Data: rawA}
	if rawA != nil {
		inAPointer = &inA
	}
	if err := BuiltinTupleReplace14ElemLongReadJSON(legacyTypeNames, inAPointer, &item.A, item.K, nat_n, item.K); err != nil {
		return err
	}

	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *Replace14Long) WriteJSONGeneral(w []byte, nat_n uint32) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w, nat_n)
}

func (item *Replace14Long) WriteJSON(w []byte, nat_n uint32) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w, nat_n)
}
func (item *Replace14Long) WriteJSONOpt(newTypeNames bool, short bool, w []byte, nat_n uint32) (_ []byte, err error) {
	w = append(w, '{')
	backupIndexK := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"k":`...)
	w = basictl.JSONWriteUint32(w, item.K)
	if (item.K != 0) == false {
		w = w[:backupIndexK]
	}
	backupIndexA := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"a":`...)
	if w, err = BuiltinTupleReplace14ElemLongWriteJSONOpt(newTypeNames, short, w, item.A, item.K, nat_n, item.K); err != nil {
		return w, err
	}
	if (len(item.A) != 0) == false {
		w = w[:backupIndexA]
	}
	return append(w, '}'), nil
}
