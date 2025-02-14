// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlBuiltinTupleTuplePairTupleIntTupleInt2

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tl/tlBuiltinTuple2PairTupleIntTupleInt"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tl/tlPairTupleIntTupleInt"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

func BuiltinTupleTuplePairTupleIntTupleInt2FillRandom(rg *basictl.RandGenerator, vec *[][2]tlPairTupleIntTupleInt.PairTupleIntTupleInt, nat_n uint32, nat_ttXn uint32, nat_ttYn uint32) {
	rg.IncreaseDepth()
	*vec = make([][2]tlPairTupleIntTupleInt.PairTupleIntTupleInt, nat_n)
	for i := range *vec {
		tlBuiltinTuple2PairTupleIntTupleInt.BuiltinTuple2PairTupleIntTupleIntFillRandom(rg, &(*vec)[i], nat_ttXn, nat_ttYn)
	}
	rg.DecreaseDepth()
}

func BuiltinTupleTuplePairTupleIntTupleInt2Read(w []byte, vec *[][2]tlPairTupleIntTupleInt.PairTupleIntTupleInt, nat_n uint32, nat_ttXn uint32, nat_ttYn uint32) (_ []byte, err error) {
	if uint32(cap(*vec)) < nat_n {
		*vec = make([][2]tlPairTupleIntTupleInt.PairTupleIntTupleInt, nat_n)
	} else {
		*vec = (*vec)[:nat_n]
	}
	for i := range *vec {
		if w, err = tlBuiltinTuple2PairTupleIntTupleInt.BuiltinTuple2PairTupleIntTupleIntRead(w, &(*vec)[i], nat_ttXn, nat_ttYn); err != nil {
			return w, err
		}
	}
	return w, nil
}

func BuiltinTupleTuplePairTupleIntTupleInt2Write(w []byte, vec [][2]tlPairTupleIntTupleInt.PairTupleIntTupleInt, nat_n uint32, nat_ttXn uint32, nat_ttYn uint32) (_ []byte, err error) {
	if uint32(len(vec)) != nat_n {
		return w, internal.ErrorWrongSequenceLength("[][2]tlPairTupleIntTupleInt.PairTupleIntTupleInt", len(vec), nat_n)
	}
	for _, elem := range vec {
		if w, err = tlBuiltinTuple2PairTupleIntTupleInt.BuiltinTuple2PairTupleIntTupleIntWrite(w, &elem, nat_ttXn, nat_ttYn); err != nil {
			return w, err
		}
	}
	return w, nil
}

func BuiltinTupleTuplePairTupleIntTupleInt2ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, vec *[][2]tlPairTupleIntTupleInt.PairTupleIntTupleInt, nat_n uint32, nat_ttXn uint32, nat_ttYn uint32) error {
	if uint32(cap(*vec)) < nat_n {
		*vec = make([][2]tlPairTupleIntTupleInt.PairTupleIntTupleInt, nat_n)
	} else {
		*vec = (*vec)[:nat_n]
	}
	index := 0
	if in != nil {
		in.Delim('[')
		if !in.Ok() {
			return internal.ErrorInvalidJSON("[][2]tlPairTupleIntTupleInt.PairTupleIntTupleInt", "expected json array")
		}
		for ; !in.IsDelim(']'); index++ {
			if nat_n <= uint32(index) {
				return internal.ErrorInvalidJSON("[][2]tlPairTupleIntTupleInt.PairTupleIntTupleInt", "array is longer than expected")
			}
			if err := tlBuiltinTuple2PairTupleIntTupleInt.BuiltinTuple2PairTupleIntTupleIntReadJSON(legacyTypeNames, in, &(*vec)[index], nat_ttXn, nat_ttYn); err != nil {
				return err
			}
			in.WantComma()
		}
		in.Delim(']')
		if !in.Ok() {
			return internal.ErrorInvalidJSON("[][2]tlPairTupleIntTupleInt.PairTupleIntTupleInt", "expected json array's end")
		}
	}
	if uint32(index) != nat_n {
		return internal.ErrorWrongSequenceLength("[][2]tlPairTupleIntTupleInt.PairTupleIntTupleInt", index, nat_n)
	}
	return nil
}

func BuiltinTupleTuplePairTupleIntTupleInt2WriteJSON(w []byte, vec [][2]tlPairTupleIntTupleInt.PairTupleIntTupleInt, nat_n uint32, nat_ttXn uint32, nat_ttYn uint32) (_ []byte, err error) {
	return BuiltinTupleTuplePairTupleIntTupleInt2WriteJSONOpt(true, false, w, vec, nat_n, nat_ttXn, nat_ttYn)
}
func BuiltinTupleTuplePairTupleIntTupleInt2WriteJSONOpt(newTypeNames bool, short bool, w []byte, vec [][2]tlPairTupleIntTupleInt.PairTupleIntTupleInt, nat_n uint32, nat_ttXn uint32, nat_ttYn uint32) (_ []byte, err error) {
	if uint32(len(vec)) != nat_n {
		return w, internal.ErrorWrongSequenceLength("[][2]tlPairTupleIntTupleInt.PairTupleIntTupleInt", len(vec), nat_n)
	}
	w = append(w, '[')
	for _, elem := range vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		if w, err = tlBuiltinTuple2PairTupleIntTupleInt.BuiltinTuple2PairTupleIntTupleIntWriteJSONOpt(newTypeNames, short, w, &elem, nat_ttXn, nat_ttYn); err != nil {
			return w, err
		}
	}
	return append(w, ']'), nil
}
