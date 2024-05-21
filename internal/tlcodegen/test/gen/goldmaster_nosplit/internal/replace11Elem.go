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

func BuiltinTupleReplace11ElemLongFillRandom(rg *basictl.RandGenerator, vec *[]Replace11ElemLong, nat_n uint32, nat_tn uint32, nat_tk uint32) {
	rg.IncreaseDepth()
	*vec = make([]Replace11ElemLong, nat_n)
	for i := range *vec {
		(*vec)[i].FillRandom(rg, nat_tn, nat_tk)
	}
	rg.DecreaseDepth()
}

func BuiltinTupleReplace11ElemLongRead(w []byte, vec *[]Replace11ElemLong, nat_n uint32, nat_tn uint32, nat_tk uint32) (_ []byte, err error) {
	if err = basictl.CheckLengthSanity(w, nat_n, 4); err != nil {
		return w, err
	}
	if uint32(cap(*vec)) < nat_n {
		*vec = make([]Replace11ElemLong, nat_n)
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

func BuiltinTupleReplace11ElemLongWrite(w []byte, vec []Replace11ElemLong, nat_n uint32, nat_tn uint32, nat_tk uint32) (_ []byte, err error) {
	if uint32(len(vec)) != nat_n {
		return w, ErrorWrongSequenceLength("[]Replace11ElemLong", len(vec), nat_n)
	}
	for _, elem := range vec {
		if w, err = elem.Write(w, nat_tn, nat_tk); err != nil {
			return w, err
		}
	}
	return w, nil
}

func BuiltinTupleReplace11ElemLongReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, vec *[]Replace11ElemLong, nat_n uint32, nat_tn uint32, nat_tk uint32) error {
	if uint32(cap(*vec)) < nat_n {
		*vec = make([]Replace11ElemLong, nat_n)
	} else {
		*vec = (*vec)[:nat_n]
	}
	index := 0
	if in != nil {
		in.Delim('[')
		if !in.Ok() {
			return ErrorInvalidJSON("[]Replace11ElemLong", "expected json array")
		}
		for ; !in.IsDelim(']'); index++ {
			if nat_n <= uint32(index) {
				return ErrorInvalidJSON("[]Replace11ElemLong", "array is longer than expected")
			}
			if err := (*vec)[index].ReadJSON(legacyTypeNames, in, nat_tn, nat_tk); err != nil {
				return err
			}
			in.WantComma()
		}
		in.Delim(']')
		if !in.Ok() {
			return ErrorInvalidJSON("[]Replace11ElemLong", "expected json array's end")
		}
	}
	if uint32(index) != nat_n {
		return ErrorWrongSequenceLength("[]Replace11ElemLong", index, nat_n)
	}
	return nil
}

func BuiltinTupleReplace11ElemLongWriteJSON(w []byte, vec []Replace11ElemLong, nat_n uint32, nat_tn uint32, nat_tk uint32) (_ []byte, err error) {
	return BuiltinTupleReplace11ElemLongWriteJSONOpt(true, false, w, vec, nat_n, nat_tn, nat_tk)
}
func BuiltinTupleReplace11ElemLongWriteJSONOpt(newTypeNames bool, short bool, w []byte, vec []Replace11ElemLong, nat_n uint32, nat_tn uint32, nat_tk uint32) (_ []byte, err error) {
	if uint32(len(vec)) != nat_n {
		return w, ErrorWrongSequenceLength("[]Replace11ElemLong", len(vec), nat_n)
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

type Replace11ElemLong struct {
	B [3]PairIntLong // Conditional: nat_n.0
	C []int32
}

func (item *Replace11ElemLong) SetB(v [3]PairIntLong, nat_n *uint32) {
	item.B = v
	if nat_n != nil {
		*nat_n |= 1 << 0
	}
}
func (item *Replace11ElemLong) ClearB(nat_n *uint32) {
	BuiltinTuple3PairBoxedIntLongReset(&item.B)
	if nat_n != nil {
		*nat_n &^= 1 << 0
	}
}
func (item Replace11ElemLong) IsSetB(nat_n uint32) bool { return nat_n&(1<<0) != 0 }

func (item *Replace11ElemLong) Reset() {
	BuiltinTuple3PairBoxedIntLongReset(&item.B)
	item.C = item.C[:0]
}

func (item *Replace11ElemLong) FillRandom(rg *basictl.RandGenerator, nat_n uint32, nat_k uint32) {
	if nat_n&(1<<0) != 0 {
		BuiltinTuple3PairBoxedIntLongFillRandom(rg, &item.B)
	} else {
		BuiltinTuple3PairBoxedIntLongReset(&item.B)
	}
	BuiltinTupleIntFillRandom(rg, &item.C, nat_n)
}

func (item *Replace11ElemLong) Read(w []byte, nat_n uint32, nat_k uint32) (_ []byte, err error) {
	if nat_n&(1<<0) != 0 {
		if w, err = BuiltinTuple3PairBoxedIntLongRead(w, &item.B); err != nil {
			return w, err
		}
	} else {
		BuiltinTuple3PairBoxedIntLongReset(&item.B)
	}
	return BuiltinTupleIntRead(w, &item.C, nat_n)
}

// This method is general version of Write, use it instead!
func (item *Replace11ElemLong) WriteGeneral(w []byte, nat_n uint32, nat_k uint32) (_ []byte, err error) {
	return item.Write(w, nat_n, nat_k)
}

func (item *Replace11ElemLong) Write(w []byte, nat_n uint32, nat_k uint32) (_ []byte, err error) {
	if nat_n&(1<<0) != 0 {
		w = BuiltinTuple3PairBoxedIntLongWrite(w, &item.B)
	}
	if w, err = BuiltinTupleIntWrite(w, item.C, nat_n); err != nil {
		return w, err
	}
	return w, nil
}

func (item *Replace11ElemLong) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, nat_n uint32, nat_k uint32) error {
	var propBPresented bool
	var rawC []byte

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
				if propBPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("replace11Elem", "b")
				}
				if nat_n&(1<<0) == 0 {
					return ErrorInvalidJSON("replace11Elem", "field 'b' is defined, while corresponding implicit fieldmask bit is 0")
				}
				if err := BuiltinTuple3PairBoxedIntLongReadJSON(legacyTypeNames, in, &item.B); err != nil {
					return err
				}
				propBPresented = true
			case "c":
				if rawC != nil {
					return ErrorInvalidJSONWithDuplicatingKeys("replace11Elem", "c")
				}
				rawC = in.Raw()
				if !in.Ok() {
					return in.Error()
				}
			default:
				return ErrorInvalidJSONExcessElement("replace11Elem", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propBPresented {
		BuiltinTuple3PairBoxedIntLongReset(&item.B)
	}
	var inCPointer *basictl.JsonLexer
	inC := basictl.JsonLexer{Data: rawC}
	if rawC != nil {
		inCPointer = &inC
	}
	if err := BuiltinTupleIntReadJSON(legacyTypeNames, inCPointer, &item.C, nat_n); err != nil {
		return err
	}

	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *Replace11ElemLong) WriteJSONGeneral(w []byte, nat_n uint32, nat_k uint32) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w, nat_n, nat_k)
}

func (item *Replace11ElemLong) WriteJSON(w []byte, nat_n uint32, nat_k uint32) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w, nat_n, nat_k)
}
func (item *Replace11ElemLong) WriteJSONOpt(newTypeNames bool, short bool, w []byte, nat_n uint32, nat_k uint32) (_ []byte, err error) {
	w = append(w, '{')
	if nat_n&(1<<0) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"b":`...)
		w = BuiltinTuple3PairBoxedIntLongWriteJSONOpt(newTypeNames, short, w, &item.B)
	}
	backupIndexC := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"c":`...)
	if w, err = BuiltinTupleIntWriteJSONOpt(newTypeNames, short, w, item.C, nat_n); err != nil {
		return w, err
	}
	if (len(item.C) != 0) == false {
		w = w[:backupIndexC]
	}
	return append(w, '}'), nil
}