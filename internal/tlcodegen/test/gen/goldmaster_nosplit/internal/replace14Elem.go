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

func BuiltinTupleReplace14ElemLongFillRandom(rg *basictl.RandGenerator, vec *[]Replace14ElemLong, nat_n uint32, nat_tn uint32, nat_tk uint32) {
	rg.IncreaseDepth()
	*vec = make([]Replace14ElemLong, nat_n)
	for i := range *vec {
		(*vec)[i].FillRandom(rg, nat_tn, nat_tk)
	}
	rg.DecreaseDepth()
}

func BuiltinTupleReplace14ElemLongRead(w []byte, vec *[]Replace14ElemLong, nat_n uint32, nat_tn uint32, nat_tk uint32) (_ []byte, err error) {
	if err = basictl.CheckLengthSanity(w, nat_n, 4); err != nil {
		return w, err
	}
	if uint32(cap(*vec)) < nat_n {
		*vec = make([]Replace14ElemLong, nat_n)
	} else {
		*vec = (*vec)[:nat_n]
	}
	for i := range *vec {
		if w, err = (*vec)[i].Read(w, nat_tn, nat_tk); err != nil {
			return w, err
		}
	}
	return w, nil
}

func BuiltinTupleReplace14ElemLongWrite(w []byte, vec []Replace14ElemLong, nat_n uint32, nat_tn uint32, nat_tk uint32) (_ []byte, err error) {
	if uint32(len(vec)) != nat_n {
		return w, ErrorWrongSequenceLength("[]Replace14ElemLong", len(vec), nat_n)
	}
	for _, elem := range vec {
		if w, err = elem.Write(w, nat_tn, nat_tk); err != nil {
			return w, err
		}
	}
	return w, nil
}

func BuiltinTupleReplace14ElemLongReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, vec *[]Replace14ElemLong, nat_n uint32, nat_tn uint32, nat_tk uint32) error {
	if uint32(cap(*vec)) < nat_n {
		*vec = make([]Replace14ElemLong, nat_n)
	} else {
		*vec = (*vec)[:nat_n]
	}
	index := 0
	if in != nil {
		in.Delim('[')
		if !in.Ok() {
			return ErrorInvalidJSON("[]Replace14ElemLong", "expected json array")
		}
		for ; !in.IsDelim(']'); index++ {
			if nat_n <= uint32(index) {
				return ErrorInvalidJSON("[]Replace14ElemLong", "array is longer than expected")
			}
			if err := (*vec)[index].ReadJSON(legacyTypeNames, in, nat_tn, nat_tk); err != nil {
				return err
			}
			in.WantComma()
		}
		in.Delim(']')
		if !in.Ok() {
			return ErrorInvalidJSON("[]Replace14ElemLong", "expected json array's end")
		}
	}
	if uint32(index) != nat_n {
		return ErrorWrongSequenceLength("[]Replace14ElemLong", index, nat_n)
	}
	return nil
}

func BuiltinTupleReplace14ElemLongWriteJSON(w []byte, vec []Replace14ElemLong, nat_n uint32, nat_tn uint32, nat_tk uint32) (_ []byte, err error) {
	return BuiltinTupleReplace14ElemLongWriteJSONOpt(true, false, w, vec, nat_n, nat_tn, nat_tk)
}
func BuiltinTupleReplace14ElemLongWriteJSONOpt(newTypeNames bool, short bool, w []byte, vec []Replace14ElemLong, nat_n uint32, nat_tn uint32, nat_tk uint32) (_ []byte, err error) {
	if uint32(len(vec)) != nat_n {
		return w, ErrorWrongSequenceLength("[]Replace14ElemLong", len(vec), nat_n)
	}
	w = append(w, '[')
	for _, elem := range vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		if w, err = elem.WriteJSONOpt(newTypeNames, short, w, nat_tn, nat_tk); err != nil {
			return w, err
		}
	}
	return append(w, ']'), nil
}

type Replace14ElemLong struct {
	B [3]Replace14ElemElemLong // Conditional: nat_n.0
	C int32
}

func (item *Replace14ElemLong) SetB(v [3]Replace14ElemElemLong, nat_n *uint32) {
	item.B = v
	if nat_n != nil {
		*nat_n |= 1 << 0
	}
}
func (item *Replace14ElemLong) ClearB(nat_n *uint32) {
	BuiltinTuple3Replace14ElemElemLongReset(&item.B)
	if nat_n != nil {
		*nat_n &^= 1 << 0
	}
}
func (item *Replace14ElemLong) IsSetB(nat_n uint32) bool { return nat_n&(1<<0) != 0 }

func (item *Replace14ElemLong) Reset() {
	BuiltinTuple3Replace14ElemElemLongReset(&item.B)
	item.C = 0
}

func (item *Replace14ElemLong) FillRandom(rg *basictl.RandGenerator, nat_n uint32, nat_k uint32) {
	if nat_n&(1<<0) != 0 {
		BuiltinTuple3Replace14ElemElemLongFillRandom(rg, &item.B, nat_n, nat_k)
	} else {
		BuiltinTuple3Replace14ElemElemLongReset(&item.B)
	}
	item.C = basictl.RandomInt(rg)
}

func (item *Replace14ElemLong) Read(w []byte, nat_n uint32, nat_k uint32) (_ []byte, err error) {
	if nat_n&(1<<0) != 0 {
		if w, err = BuiltinTuple3Replace14ElemElemLongRead(w, &item.B, nat_n, nat_k); err != nil {
			return w, err
		}
	} else {
		BuiltinTuple3Replace14ElemElemLongReset(&item.B)
	}
	return basictl.IntRead(w, &item.C)
}

// This method is general version of Write, use it instead!
func (item *Replace14ElemLong) WriteGeneral(w []byte, nat_n uint32, nat_k uint32) (_ []byte, err error) {
	return item.Write(w, nat_n, nat_k)
}

func (item *Replace14ElemLong) Write(w []byte, nat_n uint32, nat_k uint32) (_ []byte, err error) {
	if nat_n&(1<<0) != 0 {
		if w, err = BuiltinTuple3Replace14ElemElemLongWrite(w, &item.B, nat_n, nat_k); err != nil {
			return w, err
		}
	}
	w = basictl.IntWrite(w, item.C)
	return w, nil
}

func (item *Replace14ElemLong) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, nat_n uint32, nat_k uint32) error {
	var rawB []byte
	var propCPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "b":
				if rawB != nil {
					return ErrorInvalidJSONWithDuplicatingKeys("replace14Elem", "b")
				}
				rawB = in.Raw()
				if !in.Ok() {
					return in.Error()
				}
			case "c":
				if propCPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("replace14Elem", "c")
				}
				if err := Json2ReadInt32(in, &item.C); err != nil {
					return err
				}
				propCPresented = true
			default:
				return ErrorInvalidJSONExcessElement("replace14Elem", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propCPresented {
		item.C = 0
	}
	if nat_n&(1<<0) == 0 {
		if rawB != nil {
			return ErrorInvalidJSON("replace14Elem", "field 'b' is defined, while corresponding implicit fieldmask bit is 0")
		}
		BuiltinTuple3Replace14ElemElemLongReset(&item.B)
	} else {
		var inBPointer *basictl.JsonLexer
		inB := basictl.JsonLexer{Data: rawB}
		if rawB != nil {
			inBPointer = &inB
		}
		if err := BuiltinTuple3Replace14ElemElemLongReadJSON(legacyTypeNames, inBPointer, &item.B, nat_n, nat_k); err != nil {
			return err
		}

	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *Replace14ElemLong) WriteJSONGeneral(w []byte, nat_n uint32, nat_k uint32) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w, nat_n, nat_k)
}

func (item *Replace14ElemLong) WriteJSON(w []byte, nat_n uint32, nat_k uint32) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w, nat_n, nat_k)
}
func (item *Replace14ElemLong) WriteJSONOpt(newTypeNames bool, short bool, w []byte, nat_n uint32, nat_k uint32) (_ []byte, err error) {
	w = append(w, '{')
	if nat_n&(1<<0) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"b":`...)
		if w, err = BuiltinTuple3Replace14ElemElemLongWriteJSONOpt(newTypeNames, short, w, &item.B, nat_n, nat_k); err != nil {
			return w, err
		}
	}
	backupIndexC := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"c":`...)
	w = basictl.JSONWriteInt32(w, item.C)
	if (item.C != 0) == false {
		w = w[:backupIndexC]
	}
	return append(w, '}'), nil
}
