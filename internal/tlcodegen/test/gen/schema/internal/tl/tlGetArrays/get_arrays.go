// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlGetArrays

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tl/tlBuiltinTuple5Int"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tl/tlBuiltinTupleInt"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type GetArrays struct {
	N uint32
	A []int32
	B [5]int32
}

func (GetArrays) TLName() string { return "get_arrays" }
func (GetArrays) TLTag() uint32  { return 0x90658cdb }

func (item *GetArrays) Reset() {
	item.N = 0
	item.A = item.A[:0]
	tlBuiltinTuple5Int.BuiltinTuple5IntReset(&item.B)
}

func (item *GetArrays) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.N); err != nil {
		return w, err
	}
	if w, err = tlBuiltinTupleInt.BuiltinTupleIntRead(w, &item.A, item.N); err != nil {
		return w, err
	}
	return tlBuiltinTuple5Int.BuiltinTuple5IntRead(w, &item.B)
}

func (item *GetArrays) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w)
}

func (item *GetArrays) Write(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, item.N)
	if w, err = tlBuiltinTupleInt.BuiltinTupleIntWrite(w, item.A, item.N); err != nil {
		return w, err
	}
	w = tlBuiltinTuple5Int.BuiltinTuple5IntWrite(w, &item.B)
	return w, nil
}

func (item *GetArrays) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x90658cdb); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *GetArrays) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w)
}

func (item *GetArrays) WriteBoxed(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, 0x90658cdb)
	return item.Write(w)
}

func (item *GetArrays) ReadResult(w []byte, ret *[5]int32) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x9770768a); err != nil {
		return w, err
	}
	return tlBuiltinTuple5Int.BuiltinTuple5IntRead(w, ret)
}

func (item *GetArrays) WriteResult(w []byte, ret [5]int32) (_ []byte, err error) {
	w = basictl.NatWrite(w, 0x9770768a)
	w = tlBuiltinTuple5Int.BuiltinTuple5IntWrite(w, &ret)
	return w, nil
}

func (item *GetArrays) ReadResultJSON(legacyTypeNames bool, in *basictl.JsonLexer, ret *[5]int32) error {
	if err := tlBuiltinTuple5Int.BuiltinTuple5IntReadJSON(legacyTypeNames, in, ret); err != nil {
		return err
	}
	return nil
}

func (item *GetArrays) WriteResultJSON(w []byte, ret [5]int32) (_ []byte, err error) {
	tctx := basictl.JSONWriteContext{}
	return item.writeResultJSON(&tctx, w, ret)
}

func (item *GetArrays) writeResultJSON(tctx *basictl.JSONWriteContext, w []byte, ret [5]int32) (_ []byte, err error) {
	w = tlBuiltinTuple5Int.BuiltinTuple5IntWriteJSONOpt(tctx, w, &ret)
	return w, nil
}

func (item *GetArrays) ReadResultWriteResultJSON(tctx *basictl.JSONWriteContext, r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret [5]int32
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(tctx, w, ret)
	return r, w, err
}

func (item *GetArrays) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	var ret [5]int32
	err := item.ReadResultJSON(true, &basictl.JsonLexer{Data: r}, &ret)
	if err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item GetArrays) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func (item *GetArrays) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propNPresented bool
	var rawA []byte
	var propBPresented bool

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
					return internal.ErrorInvalidJSONWithDuplicatingKeys("get_arrays", "n")
				}
				if err := internal.Json2ReadUint32(in, &item.N); err != nil {
					return err
				}
				propNPresented = true
			case "a":
				if rawA != nil {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("get_arrays", "a")
				}
				rawA = in.Raw()
				if !in.Ok() {
					return in.Error()
				}
			case "b":
				if propBPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("get_arrays", "b")
				}
				if err := tlBuiltinTuple5Int.BuiltinTuple5IntReadJSON(legacyTypeNames, in, &item.B); err != nil {
					return err
				}
				propBPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("get_arrays", key)
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
	if !propBPresented {
		tlBuiltinTuple5Int.BuiltinTuple5IntReset(&item.B)
	}
	var inAPointer *basictl.JsonLexer
	inA := basictl.JsonLexer{Data: rawA}
	if rawA != nil {
		inAPointer = &inA
	}
	if err := tlBuiltinTupleInt.BuiltinTupleIntReadJSON(legacyTypeNames, inAPointer, &item.A, item.N); err != nil {
		return err
	}

	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *GetArrays) WriteJSONGeneral(tctx *basictl.JSONWriteContext, w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(tctx, w)
}

func (item *GetArrays) WriteJSON(w []byte) (_ []byte, err error) {
	tctx := basictl.JSONWriteContext{}
	return item.WriteJSONOpt(&tctx, w)
}
func (item *GetArrays) WriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte) (_ []byte, err error) {
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
	if w, err = tlBuiltinTupleInt.BuiltinTupleIntWriteJSONOpt(tctx, w, item.A, item.N); err != nil {
		return w, err
	}
	if (len(item.A) != 0) == false {
		w = w[:backupIndexA]
	}
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"b":`...)
	w = tlBuiltinTuple5Int.BuiltinTuple5IntWriteJSONOpt(tctx, w, &item.B)
	return append(w, '}'), nil
}

func (item *GetArrays) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *GetArrays) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("get_arrays", err.Error())
	}
	return nil
}
