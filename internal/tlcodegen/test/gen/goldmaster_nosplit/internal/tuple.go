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

type TupleCycleTuple []CycleTuple

func (TupleCycleTuple) TLName() string { return "tuple" }
func (TupleCycleTuple) TLTag() uint32  { return 0x9770768a }

func (item *TupleCycleTuple) Reset() {
	ptr := (*[]CycleTuple)(item)
	*ptr = (*ptr)[:0]
}

func (item *TupleCycleTuple) FillRandom(rg *basictl.RandGenerator, nat_n uint32) {
	ptr := (*[]CycleTuple)(item)
	BuiltinTupleCycleTupleFillRandom(rg, ptr, nat_n)
}

func (item *TupleCycleTuple) Read(w []byte, nat_n uint32) (_ []byte, err error) {
	ptr := (*[]CycleTuple)(item)
	return BuiltinTupleCycleTupleRead(w, ptr, nat_n)
}

// This method is general version of Write, use it instead!
func (item *TupleCycleTuple) WriteGeneral(w []byte, nat_n uint32) (_ []byte, err error) {
	return item.Write(w, nat_n)
}

func (item *TupleCycleTuple) Write(w []byte, nat_n uint32) (_ []byte, err error) {
	ptr := (*[]CycleTuple)(item)
	return BuiltinTupleCycleTupleWrite(w, *ptr, nat_n)
}

func (item *TupleCycleTuple) ReadBoxed(w []byte, nat_n uint32) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x9770768a); err != nil {
		return w, err
	}
	return item.Read(w, nat_n)
}

// This method is general version of WriteBoxed, use it instead!
func (item *TupleCycleTuple) WriteBoxedGeneral(w []byte, nat_n uint32) (_ []byte, err error) {
	return item.WriteBoxed(w, nat_n)
}

func (item *TupleCycleTuple) WriteBoxed(w []byte, nat_n uint32) (_ []byte, err error) {
	w = basictl.NatWrite(w, 0x9770768a)
	return item.Write(w, nat_n)
}

func (item *TupleCycleTuple) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, nat_n uint32) error {
	ptr := (*[]CycleTuple)(item)
	if err := BuiltinTupleCycleTupleReadJSON(legacyTypeNames, in, ptr, nat_n); err != nil {
		return err
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *TupleCycleTuple) WriteJSONGeneral(w []byte, nat_n uint32) (_ []byte, err error) {
	return item.WriteJSON(w, nat_n)
}

func (item *TupleCycleTuple) WriteJSON(w []byte, nat_n uint32) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w, nat_n)
}

func (item *TupleCycleTuple) WriteJSONOpt(newTypeNames bool, short bool, w []byte, nat_n uint32) (_ []byte, err error) {
	ptr := (*[]CycleTuple)(item)
	if w, err = BuiltinTupleCycleTupleWriteJSONOpt(newTypeNames, short, w, *ptr, nat_n); err != nil {
		return w, err
	}
	return w, nil
}

type TupleCycleTuple2 [2]CycleTuple

func (TupleCycleTuple2) TLName() string { return "tuple" }
func (TupleCycleTuple2) TLTag() uint32  { return 0x9770768a }

func (item *TupleCycleTuple2) Reset() {
	ptr := (*[2]CycleTuple)(item)
	BuiltinTuple2CycleTupleReset(ptr)
}

func (item *TupleCycleTuple2) FillRandom(rg *basictl.RandGenerator) {
	ptr := (*[2]CycleTuple)(item)
	BuiltinTuple2CycleTupleFillRandom(rg, ptr)
}

func (item *TupleCycleTuple2) Read(w []byte) (_ []byte, err error) {
	ptr := (*[2]CycleTuple)(item)
	return BuiltinTuple2CycleTupleRead(w, ptr)
}

// This method is general version of Write, use it instead!
func (item *TupleCycleTuple2) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w)
}

func (item *TupleCycleTuple2) Write(w []byte) (_ []byte, err error) {
	ptr := (*[2]CycleTuple)(item)
	return BuiltinTuple2CycleTupleWrite(w, ptr)
}

func (item *TupleCycleTuple2) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x9770768a); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *TupleCycleTuple2) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w)
}

func (item *TupleCycleTuple2) WriteBoxed(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, 0x9770768a)
	return item.Write(w)
}

func (item TupleCycleTuple2) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func (item *TupleCycleTuple2) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	ptr := (*[2]CycleTuple)(item)
	if err := BuiltinTuple2CycleTupleReadJSON(legacyTypeNames, in, ptr); err != nil {
		return err
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *TupleCycleTuple2) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSON(w)
}

func (item *TupleCycleTuple2) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w)
}

func (item *TupleCycleTuple2) WriteJSONOpt(newTypeNames bool, short bool, w []byte) (_ []byte, err error) {
	ptr := (*[2]CycleTuple)(item)
	if w, err = BuiltinTuple2CycleTupleWriteJSONOpt(newTypeNames, short, w, ptr); err != nil {
		return w, err
	}
	return w, nil
}
func (item *TupleCycleTuple2) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *TupleCycleTuple2) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("tuple", err.Error())
	}
	return nil
}

type TupleInt []int32

func (TupleInt) TLName() string { return "tuple" }
func (TupleInt) TLTag() uint32  { return 0x9770768a }

func (item *TupleInt) Reset() {
	ptr := (*[]int32)(item)
	*ptr = (*ptr)[:0]
}

func (item *TupleInt) FillRandom(rg *basictl.RandGenerator, nat_n uint32) {
	ptr := (*[]int32)(item)
	BuiltinTupleIntFillRandom(rg, ptr, nat_n)
}

func (item *TupleInt) Read(w []byte, nat_n uint32) (_ []byte, err error) {
	ptr := (*[]int32)(item)
	return BuiltinTupleIntRead(w, ptr, nat_n)
}

// This method is general version of Write, use it instead!
func (item *TupleInt) WriteGeneral(w []byte, nat_n uint32) (_ []byte, err error) {
	return item.Write(w, nat_n)
}

func (item *TupleInt) Write(w []byte, nat_n uint32) (_ []byte, err error) {
	ptr := (*[]int32)(item)
	return BuiltinTupleIntWrite(w, *ptr, nat_n)
}

func (item *TupleInt) ReadBoxed(w []byte, nat_n uint32) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x9770768a); err != nil {
		return w, err
	}
	return item.Read(w, nat_n)
}

// This method is general version of WriteBoxed, use it instead!
func (item *TupleInt) WriteBoxedGeneral(w []byte, nat_n uint32) (_ []byte, err error) {
	return item.WriteBoxed(w, nat_n)
}

func (item *TupleInt) WriteBoxed(w []byte, nat_n uint32) (_ []byte, err error) {
	w = basictl.NatWrite(w, 0x9770768a)
	return item.Write(w, nat_n)
}

func (item *TupleInt) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, nat_n uint32) error {
	ptr := (*[]int32)(item)
	if err := BuiltinTupleIntReadJSON(legacyTypeNames, in, ptr, nat_n); err != nil {
		return err
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *TupleInt) WriteJSONGeneral(w []byte, nat_n uint32) (_ []byte, err error) {
	return item.WriteJSON(w, nat_n)
}

func (item *TupleInt) WriteJSON(w []byte, nat_n uint32) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w, nat_n)
}

func (item *TupleInt) WriteJSONOpt(newTypeNames bool, short bool, w []byte, nat_n uint32) (_ []byte, err error) {
	ptr := (*[]int32)(item)
	if w, err = BuiltinTupleIntWriteJSONOpt(newTypeNames, short, w, *ptr, nat_n); err != nil {
		return w, err
	}
	return w, nil
}

type TupleInt0 [0]int32

func (TupleInt0) TLName() string { return "tuple" }
func (TupleInt0) TLTag() uint32  { return 0x9770768a }

func (item *TupleInt0) Reset() {
	ptr := (*[0]int32)(item)
	BuiltinTuple0IntReset(ptr)
}

func (item *TupleInt0) FillRandom(rg *basictl.RandGenerator) {
	ptr := (*[0]int32)(item)
	BuiltinTuple0IntFillRandom(rg, ptr)
}

func (item *TupleInt0) Read(w []byte) (_ []byte, err error) {
	ptr := (*[0]int32)(item)
	return BuiltinTuple0IntRead(w, ptr)
}

// This method is general version of Write, use it instead!
func (item *TupleInt0) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *TupleInt0) Write(w []byte) []byte {
	ptr := (*[0]int32)(item)
	return BuiltinTuple0IntWrite(w, ptr)
}

func (item *TupleInt0) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x9770768a); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *TupleInt0) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *TupleInt0) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x9770768a)
	return item.Write(w)
}

func (item TupleInt0) String() string {
	return string(item.WriteJSON(nil))
}

func (item *TupleInt0) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	ptr := (*[0]int32)(item)
	if err := BuiltinTuple0IntReadJSON(legacyTypeNames, in, ptr); err != nil {
		return err
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *TupleInt0) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSON(w), nil
}

func (item *TupleInt0) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}

func (item *TupleInt0) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	ptr := (*[0]int32)(item)
	w = BuiltinTuple0IntWriteJSONOpt(newTypeNames, short, w, ptr)
	return w
}
func (item *TupleInt0) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *TupleInt0) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("tuple", err.Error())
	}
	return nil
}

type TupleInt0Maybe struct {
	Value [0]int32 // not deterministic if !Ok
	Ok    bool
}

func (item *TupleInt0Maybe) Reset() {
	item.Ok = false
}
func (item *TupleInt0Maybe) FillRandom(rg *basictl.RandGenerator) {
	if basictl.RandomUint(rg)&1 == 1 {
		item.Ok = true
		BuiltinTuple0IntFillRandom(rg, &item.Value)
	} else {
		item.Ok = false
	}
}

func (item *TupleInt0Maybe) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.ReadBool(w, &item.Ok, 0x27930a7b, 0x3f9c8ef8); err != nil {
		return w, err
	}
	if item.Ok {
		return BuiltinTuple0IntRead(w, &item.Value)
	}
	return w, nil
}

// This method is general version of WriteBoxed, use it instead!
func (item *TupleInt0Maybe) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *TupleInt0Maybe) WriteBoxed(w []byte) []byte {
	if item.Ok {
		w = basictl.NatWrite(w, 0x3f9c8ef8)
		return BuiltinTuple0IntWrite(w, &item.Value)
	}
	return basictl.NatWrite(w, 0x27930a7b)
}

func (item *TupleInt0Maybe) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	_ok, _jvalue, err := Json2ReadMaybe("Maybe", in)
	if err != nil {
		return err
	}
	item.Ok = _ok
	if _ok {
		var in2Pointer *basictl.JsonLexer
		if _jvalue != nil {
			in2 := basictl.JsonLexer{Data: _jvalue}
			in2Pointer = &in2
		}
		if err := BuiltinTuple0IntReadJSON(legacyTypeNames, in2Pointer, &item.Value); err != nil {
			return err
		}
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *TupleInt0Maybe) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *TupleInt0Maybe) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *TupleInt0Maybe) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	if !item.Ok {
		return append(w, "{}"...)
	}
	w = append(w, `{"ok":true`...)
	w = append(w, `,"value":`...)
	w = BuiltinTuple0IntWriteJSONOpt(newTypeNames, short, w, &item.Value)
	return append(w, '}')
}

func (item TupleInt0Maybe) String() string {
	return string(item.WriteJSON(nil))
}

type TupleInt3 [3]int32

func (TupleInt3) TLName() string { return "tuple" }
func (TupleInt3) TLTag() uint32  { return 0x9770768a }

func (item *TupleInt3) Reset() {
	ptr := (*[3]int32)(item)
	BuiltinTuple3IntReset(ptr)
}

func (item *TupleInt3) FillRandom(rg *basictl.RandGenerator) {
	ptr := (*[3]int32)(item)
	BuiltinTuple3IntFillRandom(rg, ptr)
}

func (item *TupleInt3) Read(w []byte) (_ []byte, err error) {
	ptr := (*[3]int32)(item)
	return BuiltinTuple3IntRead(w, ptr)
}

// This method is general version of Write, use it instead!
func (item *TupleInt3) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *TupleInt3) Write(w []byte) []byte {
	ptr := (*[3]int32)(item)
	return BuiltinTuple3IntWrite(w, ptr)
}

func (item *TupleInt3) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x9770768a); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *TupleInt3) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *TupleInt3) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x9770768a)
	return item.Write(w)
}

func (item TupleInt3) String() string {
	return string(item.WriteJSON(nil))
}

func (item *TupleInt3) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	ptr := (*[3]int32)(item)
	if err := BuiltinTuple3IntReadJSON(legacyTypeNames, in, ptr); err != nil {
		return err
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *TupleInt3) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSON(w), nil
}

func (item *TupleInt3) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}

func (item *TupleInt3) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	ptr := (*[3]int32)(item)
	w = BuiltinTuple3IntWriteJSONOpt(newTypeNames, short, w, ptr)
	return w
}
func (item *TupleInt3) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *TupleInt3) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("tuple", err.Error())
	}
	return nil
}

type TupleInt3BoxedMaybe struct {
	Value [3]int32 // not deterministic if !Ok
	Ok    bool
}

func (item *TupleInt3BoxedMaybe) Reset() {
	item.Ok = false
}
func (item *TupleInt3BoxedMaybe) FillRandom(rg *basictl.RandGenerator) {
	if basictl.RandomUint(rg)&1 == 1 {
		item.Ok = true
		BuiltinTuple3IntFillRandom(rg, &item.Value)
	} else {
		item.Ok = false
	}
}

func (item *TupleInt3BoxedMaybe) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.ReadBool(w, &item.Ok, 0x27930a7b, 0x3f9c8ef8); err != nil {
		return w, err
	}
	if item.Ok {
		if w, err = basictl.NatReadExactTag(w, 0x9770768a); err != nil {
			return w, err
		}
		return BuiltinTuple3IntRead(w, &item.Value)
	}
	return w, nil
}

// This method is general version of WriteBoxed, use it instead!
func (item *TupleInt3BoxedMaybe) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *TupleInt3BoxedMaybe) WriteBoxed(w []byte) []byte {
	if item.Ok {
		w = basictl.NatWrite(w, 0x3f9c8ef8)
		w = basictl.NatWrite(w, 0x9770768a)
		return BuiltinTuple3IntWrite(w, &item.Value)
	}
	return basictl.NatWrite(w, 0x27930a7b)
}

func (item *TupleInt3BoxedMaybe) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	_ok, _jvalue, err := Json2ReadMaybe("Maybe", in)
	if err != nil {
		return err
	}
	item.Ok = _ok
	if _ok {
		var in2Pointer *basictl.JsonLexer
		if _jvalue != nil {
			in2 := basictl.JsonLexer{Data: _jvalue}
			in2Pointer = &in2
		}
		if err := BuiltinTuple3IntReadJSON(legacyTypeNames, in2Pointer, &item.Value); err != nil {
			return err
		}
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *TupleInt3BoxedMaybe) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *TupleInt3BoxedMaybe) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *TupleInt3BoxedMaybe) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	if !item.Ok {
		return append(w, "{}"...)
	}
	w = append(w, `{"ok":true`...)
	w = append(w, `,"value":`...)
	w = BuiltinTuple3IntWriteJSONOpt(newTypeNames, short, w, &item.Value)
	return append(w, '}')
}

func (item TupleInt3BoxedMaybe) String() string {
	return string(item.WriteJSON(nil))
}

type TupleInt3Maybe struct {
	Value [3]int32 // not deterministic if !Ok
	Ok    bool
}

func (item *TupleInt3Maybe) Reset() {
	item.Ok = false
}
func (item *TupleInt3Maybe) FillRandom(rg *basictl.RandGenerator) {
	if basictl.RandomUint(rg)&1 == 1 {
		item.Ok = true
		BuiltinTuple3IntFillRandom(rg, &item.Value)
	} else {
		item.Ok = false
	}
}

func (item *TupleInt3Maybe) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.ReadBool(w, &item.Ok, 0x27930a7b, 0x3f9c8ef8); err != nil {
		return w, err
	}
	if item.Ok {
		return BuiltinTuple3IntRead(w, &item.Value)
	}
	return w, nil
}

// This method is general version of WriteBoxed, use it instead!
func (item *TupleInt3Maybe) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *TupleInt3Maybe) WriteBoxed(w []byte) []byte {
	if item.Ok {
		w = basictl.NatWrite(w, 0x3f9c8ef8)
		return BuiltinTuple3IntWrite(w, &item.Value)
	}
	return basictl.NatWrite(w, 0x27930a7b)
}

func (item *TupleInt3Maybe) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	_ok, _jvalue, err := Json2ReadMaybe("Maybe", in)
	if err != nil {
		return err
	}
	item.Ok = _ok
	if _ok {
		var in2Pointer *basictl.JsonLexer
		if _jvalue != nil {
			in2 := basictl.JsonLexer{Data: _jvalue}
			in2Pointer = &in2
		}
		if err := BuiltinTuple3IntReadJSON(legacyTypeNames, in2Pointer, &item.Value); err != nil {
			return err
		}
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *TupleInt3Maybe) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *TupleInt3Maybe) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *TupleInt3Maybe) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	if !item.Ok {
		return append(w, "{}"...)
	}
	w = append(w, `{"ok":true`...)
	w = append(w, `,"value":`...)
	w = BuiltinTuple3IntWriteJSONOpt(newTypeNames, short, w, &item.Value)
	return append(w, '}')
}

func (item TupleInt3Maybe) String() string {
	return string(item.WriteJSON(nil))
}

type TupleIntBoxed0 [0]int32

func (TupleIntBoxed0) TLName() string { return "tuple" }
func (TupleIntBoxed0) TLTag() uint32  { return 0x9770768a }

func (item *TupleIntBoxed0) Reset() {
	ptr := (*[0]int32)(item)
	BuiltinTuple0IntBoxedReset(ptr)
}

func (item *TupleIntBoxed0) FillRandom(rg *basictl.RandGenerator) {
	ptr := (*[0]int32)(item)
	BuiltinTuple0IntBoxedFillRandom(rg, ptr)
}

func (item *TupleIntBoxed0) Read(w []byte) (_ []byte, err error) {
	ptr := (*[0]int32)(item)
	return BuiltinTuple0IntBoxedRead(w, ptr)
}

// This method is general version of Write, use it instead!
func (item *TupleIntBoxed0) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *TupleIntBoxed0) Write(w []byte) []byte {
	ptr := (*[0]int32)(item)
	return BuiltinTuple0IntBoxedWrite(w, ptr)
}

func (item *TupleIntBoxed0) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x9770768a); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *TupleIntBoxed0) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *TupleIntBoxed0) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x9770768a)
	return item.Write(w)
}

func (item TupleIntBoxed0) String() string {
	return string(item.WriteJSON(nil))
}

func (item *TupleIntBoxed0) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	ptr := (*[0]int32)(item)
	if err := BuiltinTuple0IntBoxedReadJSON(legacyTypeNames, in, ptr); err != nil {
		return err
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *TupleIntBoxed0) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSON(w), nil
}

func (item *TupleIntBoxed0) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}

func (item *TupleIntBoxed0) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	ptr := (*[0]int32)(item)
	w = BuiltinTuple0IntBoxedWriteJSONOpt(newTypeNames, short, w, ptr)
	return w
}
func (item *TupleIntBoxed0) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *TupleIntBoxed0) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("tuple", err.Error())
	}
	return nil
}

type TupleIntBoxed0BoxedMaybe struct {
	Value [0]int32 // not deterministic if !Ok
	Ok    bool
}

func (item *TupleIntBoxed0BoxedMaybe) Reset() {
	item.Ok = false
}
func (item *TupleIntBoxed0BoxedMaybe) FillRandom(rg *basictl.RandGenerator) {
	if basictl.RandomUint(rg)&1 == 1 {
		item.Ok = true
		BuiltinTuple0IntBoxedFillRandom(rg, &item.Value)
	} else {
		item.Ok = false
	}
}

func (item *TupleIntBoxed0BoxedMaybe) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.ReadBool(w, &item.Ok, 0x27930a7b, 0x3f9c8ef8); err != nil {
		return w, err
	}
	if item.Ok {
		if w, err = basictl.NatReadExactTag(w, 0x9770768a); err != nil {
			return w, err
		}
		return BuiltinTuple0IntBoxedRead(w, &item.Value)
	}
	return w, nil
}

// This method is general version of WriteBoxed, use it instead!
func (item *TupleIntBoxed0BoxedMaybe) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *TupleIntBoxed0BoxedMaybe) WriteBoxed(w []byte) []byte {
	if item.Ok {
		w = basictl.NatWrite(w, 0x3f9c8ef8)
		w = basictl.NatWrite(w, 0x9770768a)
		return BuiltinTuple0IntBoxedWrite(w, &item.Value)
	}
	return basictl.NatWrite(w, 0x27930a7b)
}

func (item *TupleIntBoxed0BoxedMaybe) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	_ok, _jvalue, err := Json2ReadMaybe("Maybe", in)
	if err != nil {
		return err
	}
	item.Ok = _ok
	if _ok {
		var in2Pointer *basictl.JsonLexer
		if _jvalue != nil {
			in2 := basictl.JsonLexer{Data: _jvalue}
			in2Pointer = &in2
		}
		if err := BuiltinTuple0IntBoxedReadJSON(legacyTypeNames, in2Pointer, &item.Value); err != nil {
			return err
		}
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *TupleIntBoxed0BoxedMaybe) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *TupleIntBoxed0BoxedMaybe) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *TupleIntBoxed0BoxedMaybe) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	if !item.Ok {
		return append(w, "{}"...)
	}
	w = append(w, `{"ok":true`...)
	w = append(w, `,"value":`...)
	w = BuiltinTuple0IntBoxedWriteJSONOpt(newTypeNames, short, w, &item.Value)
	return append(w, '}')
}

func (item TupleIntBoxed0BoxedMaybe) String() string {
	return string(item.WriteJSON(nil))
}

type TupleIntBoxed3 [3]int32

func (TupleIntBoxed3) TLName() string { return "tuple" }
func (TupleIntBoxed3) TLTag() uint32  { return 0x9770768a }

func (item *TupleIntBoxed3) Reset() {
	ptr := (*[3]int32)(item)
	BuiltinTuple3IntBoxedReset(ptr)
}

func (item *TupleIntBoxed3) FillRandom(rg *basictl.RandGenerator) {
	ptr := (*[3]int32)(item)
	BuiltinTuple3IntBoxedFillRandom(rg, ptr)
}

func (item *TupleIntBoxed3) Read(w []byte) (_ []byte, err error) {
	ptr := (*[3]int32)(item)
	return BuiltinTuple3IntBoxedRead(w, ptr)
}

// This method is general version of Write, use it instead!
func (item *TupleIntBoxed3) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *TupleIntBoxed3) Write(w []byte) []byte {
	ptr := (*[3]int32)(item)
	return BuiltinTuple3IntBoxedWrite(w, ptr)
}

func (item *TupleIntBoxed3) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x9770768a); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *TupleIntBoxed3) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *TupleIntBoxed3) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x9770768a)
	return item.Write(w)
}

func (item TupleIntBoxed3) String() string {
	return string(item.WriteJSON(nil))
}

func (item *TupleIntBoxed3) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	ptr := (*[3]int32)(item)
	if err := BuiltinTuple3IntBoxedReadJSON(legacyTypeNames, in, ptr); err != nil {
		return err
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *TupleIntBoxed3) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSON(w), nil
}

func (item *TupleIntBoxed3) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}

func (item *TupleIntBoxed3) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	ptr := (*[3]int32)(item)
	w = BuiltinTuple3IntBoxedWriteJSONOpt(newTypeNames, short, w, ptr)
	return w
}
func (item *TupleIntBoxed3) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *TupleIntBoxed3) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("tuple", err.Error())
	}
	return nil
}

type TupleIntBoxed3Maybe struct {
	Value [3]int32 // not deterministic if !Ok
	Ok    bool
}

func (item *TupleIntBoxed3Maybe) Reset() {
	item.Ok = false
}
func (item *TupleIntBoxed3Maybe) FillRandom(rg *basictl.RandGenerator) {
	if basictl.RandomUint(rg)&1 == 1 {
		item.Ok = true
		BuiltinTuple3IntBoxedFillRandom(rg, &item.Value)
	} else {
		item.Ok = false
	}
}

func (item *TupleIntBoxed3Maybe) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.ReadBool(w, &item.Ok, 0x27930a7b, 0x3f9c8ef8); err != nil {
		return w, err
	}
	if item.Ok {
		return BuiltinTuple3IntBoxedRead(w, &item.Value)
	}
	return w, nil
}

// This method is general version of WriteBoxed, use it instead!
func (item *TupleIntBoxed3Maybe) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *TupleIntBoxed3Maybe) WriteBoxed(w []byte) []byte {
	if item.Ok {
		w = basictl.NatWrite(w, 0x3f9c8ef8)
		return BuiltinTuple3IntBoxedWrite(w, &item.Value)
	}
	return basictl.NatWrite(w, 0x27930a7b)
}

func (item *TupleIntBoxed3Maybe) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	_ok, _jvalue, err := Json2ReadMaybe("Maybe", in)
	if err != nil {
		return err
	}
	item.Ok = _ok
	if _ok {
		var in2Pointer *basictl.JsonLexer
		if _jvalue != nil {
			in2 := basictl.JsonLexer{Data: _jvalue}
			in2Pointer = &in2
		}
		if err := BuiltinTuple3IntBoxedReadJSON(legacyTypeNames, in2Pointer, &item.Value); err != nil {
			return err
		}
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *TupleIntBoxed3Maybe) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *TupleIntBoxed3Maybe) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *TupleIntBoxed3Maybe) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	if !item.Ok {
		return append(w, "{}"...)
	}
	w = append(w, `{"ok":true`...)
	w = append(w, `,"value":`...)
	w = BuiltinTuple3IntBoxedWriteJSONOpt(newTypeNames, short, w, &item.Value)
	return append(w, '}')
}

func (item TupleIntBoxed3Maybe) String() string {
	return string(item.WriteJSON(nil))
}

type TupleIntMaybe struct {
	Value []int32 // not deterministic if !Ok
	Ok    bool
}

func (item *TupleIntMaybe) Reset() {
	item.Ok = false
}
func (item *TupleIntMaybe) FillRandom(rg *basictl.RandGenerator, nat_t uint32) {
	if basictl.RandomUint(rg)&1 == 1 {
		item.Ok = true
		BuiltinTupleIntFillRandom(rg, &item.Value, nat_t)
	} else {
		item.Ok = false
	}
}

func (item *TupleIntMaybe) ReadBoxed(w []byte, nat_t uint32) (_ []byte, err error) {
	if w, err = basictl.ReadBool(w, &item.Ok, 0x27930a7b, 0x3f9c8ef8); err != nil {
		return w, err
	}
	if item.Ok {
		return BuiltinTupleIntRead(w, &item.Value, nat_t)
	}
	return w, nil
}

// This method is general version of WriteBoxed, use it instead!
func (item *TupleIntMaybe) WriteBoxedGeneral(w []byte, nat_t uint32) (_ []byte, err error) {
	return item.WriteBoxed(w, nat_t)
}

func (item *TupleIntMaybe) WriteBoxed(w []byte, nat_t uint32) (_ []byte, err error) {
	if item.Ok {
		w = basictl.NatWrite(w, 0x3f9c8ef8)
		return BuiltinTupleIntWrite(w, item.Value, nat_t)
	}
	return basictl.NatWrite(w, 0x27930a7b), nil
}

func (item *TupleIntMaybe) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, nat_t uint32) error {
	_ok, _jvalue, err := Json2ReadMaybe("Maybe", in)
	if err != nil {
		return err
	}
	item.Ok = _ok
	if _ok {
		var in2Pointer *basictl.JsonLexer
		if _jvalue != nil {
			in2 := basictl.JsonLexer{Data: _jvalue}
			in2Pointer = &in2
		}
		if err := BuiltinTupleIntReadJSON(legacyTypeNames, in2Pointer, &item.Value, nat_t); err != nil {
			return err
		}
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *TupleIntMaybe) WriteJSONGeneral(w []byte, nat_t uint32) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w, nat_t)
}

func (item *TupleIntMaybe) WriteJSON(w []byte, nat_t uint32) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w, nat_t)
}
func (item *TupleIntMaybe) WriteJSONOpt(newTypeNames bool, short bool, w []byte, nat_t uint32) (_ []byte, err error) {
	if !item.Ok {
		return append(w, "{}"...), nil
	}
	w = append(w, `{"ok":true`...)
	if len(item.Value) != 0 {
		w = append(w, `,"value":`...)
		if w, err = BuiltinTupleIntWriteJSONOpt(newTypeNames, short, w, item.Value, nat_t); err != nil {
			return w, err
		}
	}
	return append(w, '}'), nil
}

type TupleLong []int64

func (TupleLong) TLName() string { return "tuple" }
func (TupleLong) TLTag() uint32  { return 0x9770768a }

func (item *TupleLong) Reset() {
	ptr := (*[]int64)(item)
	*ptr = (*ptr)[:0]
}

func (item *TupleLong) FillRandom(rg *basictl.RandGenerator, nat_n uint32) {
	ptr := (*[]int64)(item)
	BuiltinTupleLongFillRandom(rg, ptr, nat_n)
}

func (item *TupleLong) Read(w []byte, nat_n uint32) (_ []byte, err error) {
	ptr := (*[]int64)(item)
	return BuiltinTupleLongRead(w, ptr, nat_n)
}

// This method is general version of Write, use it instead!
func (item *TupleLong) WriteGeneral(w []byte, nat_n uint32) (_ []byte, err error) {
	return item.Write(w, nat_n)
}

func (item *TupleLong) Write(w []byte, nat_n uint32) (_ []byte, err error) {
	ptr := (*[]int64)(item)
	return BuiltinTupleLongWrite(w, *ptr, nat_n)
}

func (item *TupleLong) ReadBoxed(w []byte, nat_n uint32) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x9770768a); err != nil {
		return w, err
	}
	return item.Read(w, nat_n)
}

// This method is general version of WriteBoxed, use it instead!
func (item *TupleLong) WriteBoxedGeneral(w []byte, nat_n uint32) (_ []byte, err error) {
	return item.WriteBoxed(w, nat_n)
}

func (item *TupleLong) WriteBoxed(w []byte, nat_n uint32) (_ []byte, err error) {
	w = basictl.NatWrite(w, 0x9770768a)
	return item.Write(w, nat_n)
}

func (item *TupleLong) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, nat_n uint32) error {
	ptr := (*[]int64)(item)
	if err := BuiltinTupleLongReadJSON(legacyTypeNames, in, ptr, nat_n); err != nil {
		return err
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *TupleLong) WriteJSONGeneral(w []byte, nat_n uint32) (_ []byte, err error) {
	return item.WriteJSON(w, nat_n)
}

func (item *TupleLong) WriteJSON(w []byte, nat_n uint32) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w, nat_n)
}

func (item *TupleLong) WriteJSONOpt(newTypeNames bool, short bool, w []byte, nat_n uint32) (_ []byte, err error) {
	ptr := (*[]int64)(item)
	if w, err = BuiltinTupleLongWriteJSONOpt(newTypeNames, short, w, *ptr, nat_n); err != nil {
		return w, err
	}
	return w, nil
}

type TupleString []string

func (TupleString) TLName() string { return "tuple" }
func (TupleString) TLTag() uint32  { return 0x9770768a }

func (item *TupleString) Reset() {
	ptr := (*[]string)(item)
	*ptr = (*ptr)[:0]
}

func (item *TupleString) FillRandom(rg *basictl.RandGenerator, nat_n uint32) {
	ptr := (*[]string)(item)
	BuiltinTupleStringFillRandom(rg, ptr, nat_n)
}

func (item *TupleString) Read(w []byte, nat_n uint32) (_ []byte, err error) {
	ptr := (*[]string)(item)
	return BuiltinTupleStringRead(w, ptr, nat_n)
}

// This method is general version of Write, use it instead!
func (item *TupleString) WriteGeneral(w []byte, nat_n uint32) (_ []byte, err error) {
	return item.Write(w, nat_n)
}

func (item *TupleString) Write(w []byte, nat_n uint32) (_ []byte, err error) {
	ptr := (*[]string)(item)
	return BuiltinTupleStringWrite(w, *ptr, nat_n)
}

func (item *TupleString) ReadBoxed(w []byte, nat_n uint32) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x9770768a); err != nil {
		return w, err
	}
	return item.Read(w, nat_n)
}

// This method is general version of WriteBoxed, use it instead!
func (item *TupleString) WriteBoxedGeneral(w []byte, nat_n uint32) (_ []byte, err error) {
	return item.WriteBoxed(w, nat_n)
}

func (item *TupleString) WriteBoxed(w []byte, nat_n uint32) (_ []byte, err error) {
	w = basictl.NatWrite(w, 0x9770768a)
	return item.Write(w, nat_n)
}

func (item *TupleString) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, nat_n uint32) error {
	ptr := (*[]string)(item)
	if err := BuiltinTupleStringReadJSON(legacyTypeNames, in, ptr, nat_n); err != nil {
		return err
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *TupleString) WriteJSONGeneral(w []byte, nat_n uint32) (_ []byte, err error) {
	return item.WriteJSON(w, nat_n)
}

func (item *TupleString) WriteJSON(w []byte, nat_n uint32) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w, nat_n)
}

func (item *TupleString) WriteJSONOpt(newTypeNames bool, short bool, w []byte, nat_n uint32) (_ []byte, err error) {
	ptr := (*[]string)(item)
	if w, err = BuiltinTupleStringWriteJSONOpt(newTypeNames, short, w, *ptr, nat_n); err != nil {
		return w, err
	}
	return w, nil
}
