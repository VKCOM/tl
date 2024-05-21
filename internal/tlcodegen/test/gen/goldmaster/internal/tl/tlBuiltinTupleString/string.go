// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlBuiltinTupleString

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

func BuiltinTupleStringFillRandom(rg *basictl.RandGenerator, vec *[]string, nat_n uint32) {
	rg.IncreaseDepth()
	*vec = make([]string, nat_n)
	for i := range *vec {
		(*vec)[i] = basictl.RandomString(rg)
	}
	rg.DecreaseDepth()
}

func BuiltinTupleStringRead(w []byte, vec *[]string, nat_n uint32) (_ []byte, err error) {
	if err = basictl.CheckLengthSanity(w, nat_n, 4); err != nil {
		return w, err
	}
	if uint32(cap(*vec)) < nat_n {
		*vec = make([]string, nat_n)
	} else {
		*vec = (*vec)[:nat_n]
	}
	for i := range *vec {
		if w, err = basictl.StringRead(w, &(*vec)[i]); err != nil {
			return w, err
		}
	}
	return w, nil
}

func BuiltinTupleStringWrite(w []byte, vec []string, nat_n uint32) (_ []byte, err error) {
	if uint32(len(vec)) != nat_n {
		return w, internal.ErrorWrongSequenceLength("[]string", len(vec), nat_n)
	}
	for _, elem := range vec {
		w = basictl.StringWrite(w, elem)
	}
	return w, nil
}

func BuiltinTupleStringReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, vec *[]string, nat_n uint32) error {
	if uint32(cap(*vec)) < nat_n {
		*vec = make([]string, nat_n)
	} else {
		*vec = (*vec)[:nat_n]
	}
	index := 0
	if in != nil {
		in.Delim('[')
		if !in.Ok() {
			return internal.ErrorInvalidJSON("[]string", "expected json array")
		}
		for ; !in.IsDelim(']'); index++ {
			if nat_n <= uint32(index) {
				return internal.ErrorInvalidJSON("[]string", "array is longer than expected")
			}
			if err := internal.Json2ReadString(in, &(*vec)[index]); err != nil {
				return err
			}
			in.WantComma()
		}
		in.Delim(']')
		if !in.Ok() {
			return internal.ErrorInvalidJSON("[]string", "expected json array's end")
		}
	}
	if uint32(index) != nat_n {
		return internal.ErrorWrongSequenceLength("[]string", index, nat_n)
	}
	return nil
}

func BuiltinTupleStringWriteJSON(w []byte, vec []string, nat_n uint32) (_ []byte, err error) {
	return BuiltinTupleStringWriteJSONOpt(true, false, w, vec, nat_n)
}
func BuiltinTupleStringWriteJSONOpt(newTypeNames bool, short bool, w []byte, vec []string, nat_n uint32) (_ []byte, err error) {
	if uint32(len(vec)) != nat_n {
		return w, internal.ErrorWrongSequenceLength("[]string", len(vec), nat_n)
	}
	w = append(w, '[')
	for _, elem := range vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = basictl.JSONWriteString(w, elem)
	}
	return append(w, ']'), nil
}
