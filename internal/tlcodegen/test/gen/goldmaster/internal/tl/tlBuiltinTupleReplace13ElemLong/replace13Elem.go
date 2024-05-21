// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlBuiltinTupleReplace13ElemLong

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlReplace13ElemLong"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

func BuiltinTupleReplace13ElemLongFillRandom(rg *basictl.RandGenerator, vec *[]tlReplace13ElemLong.Replace13ElemLong, nat_n uint32, nat_tn uint32, nat_tk uint32) {
	rg.IncreaseDepth()
	*vec = make([]tlReplace13ElemLong.Replace13ElemLong, nat_n)
	for i := range *vec {
		(*vec)[i].FillRandom(rg, nat_tn, nat_tk)
	}
	rg.DecreaseDepth()
}

func BuiltinTupleReplace13ElemLongRead(w []byte, vec *[]tlReplace13ElemLong.Replace13ElemLong, nat_n uint32, nat_tn uint32, nat_tk uint32) (_ []byte, err error) {
	if err = basictl.CheckLengthSanity(w, nat_n, 4); err != nil {
		return w, err
	}
	if uint32(cap(*vec)) < nat_n {
		*vec = make([]tlReplace13ElemLong.Replace13ElemLong, nat_n)
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

func BuiltinTupleReplace13ElemLongWrite(w []byte, vec []tlReplace13ElemLong.Replace13ElemLong, nat_n uint32, nat_tn uint32, nat_tk uint32) (_ []byte, err error) {
	if uint32(len(vec)) != nat_n {
		return w, internal.ErrorWrongSequenceLength("[]tlReplace13ElemLong.Replace13ElemLong", len(vec), nat_n)
	}
	for _, elem := range vec {
		if w, err = elem.Write(w, nat_tn, nat_tk); err != nil {
			return w, err
		}
	}
	return w, nil
}

func BuiltinTupleReplace13ElemLongReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, vec *[]tlReplace13ElemLong.Replace13ElemLong, nat_n uint32, nat_tn uint32, nat_tk uint32) error {
	if uint32(cap(*vec)) < nat_n {
		*vec = make([]tlReplace13ElemLong.Replace13ElemLong, nat_n)
	} else {
		*vec = (*vec)[:nat_n]
	}
	index := 0
	if in != nil {
		in.Delim('[')
		if !in.Ok() {
			return internal.ErrorInvalidJSON("[]tlReplace13ElemLong.Replace13ElemLong", "expected json array")
		}
		for ; !in.IsDelim(']'); index++ {
			if nat_n <= uint32(index) {
				return internal.ErrorInvalidJSON("[]tlReplace13ElemLong.Replace13ElemLong", "array is longer than expected")
			}
			if err := (*vec)[index].ReadJSON(legacyTypeNames, in, nat_tn, nat_tk); err != nil {
				return err
			}
			in.WantComma()
		}
		in.Delim(']')
		if !in.Ok() {
			return internal.ErrorInvalidJSON("[]tlReplace13ElemLong.Replace13ElemLong", "expected json array's end")
		}
	}
	if uint32(index) != nat_n {
		return internal.ErrorWrongSequenceLength("[]tlReplace13ElemLong.Replace13ElemLong", index, nat_n)
	}
	return nil
}

func BuiltinTupleReplace13ElemLongWriteJSON(w []byte, vec []tlReplace13ElemLong.Replace13ElemLong, nat_n uint32, nat_tn uint32, nat_tk uint32) (_ []byte, err error) {
	return BuiltinTupleReplace13ElemLongWriteJSONOpt(true, false, w, vec, nat_n, nat_tn, nat_tk)
}
func BuiltinTupleReplace13ElemLongWriteJSONOpt(newTypeNames bool, short bool, w []byte, vec []tlReplace13ElemLong.Replace13ElemLong, nat_n uint32, nat_tn uint32, nat_tk uint32) (_ []byte, err error) {
	if uint32(len(vec)) != nat_n {
		return w, internal.ErrorWrongSequenceLength("[]tlReplace13ElemLong.Replace13ElemLong", len(vec), nat_n)
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
