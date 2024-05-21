// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlBuiltinTuple3Int32Boxed

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlInt32"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

func BuiltinTuple3Int32BoxedReset(vec *[3]tlInt32.Int32) {
	for i := range *vec {
		(*vec)[i].Reset()
	}
}

func BuiltinTuple3Int32BoxedFillRandom(rg *basictl.RandGenerator, vec *[3]tlInt32.Int32) {
	rg.IncreaseDepth()
	for i := range *vec {
		(*vec)[i].FillRandom(rg)
	}
	rg.DecreaseDepth()
}

func BuiltinTuple3Int32BoxedRead(w []byte, vec *[3]tlInt32.Int32) (_ []byte, err error) {
	for i := range *vec {
		if w, err = (*vec)[i].ReadBoxed(w); err != nil {
			return w, err
		}
	}
	return w, nil
}

func BuiltinTuple3Int32BoxedWrite(w []byte, vec *[3]tlInt32.Int32) []byte {
	for _, elem := range *vec {
		w = elem.WriteBoxed(w)
	}
	return w
}

func BuiltinTuple3Int32BoxedReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, vec *[3]tlInt32.Int32) error {
	index := 0
	if in != nil {
		in.Delim('[')
		if !in.Ok() {
			return internal.ErrorInvalidJSON("[3]tlInt32.Int32", "expected json array")
		}
		for ; !in.IsDelim(']'); index++ {
			if index == 3 {
				return internal.ErrorWrongSequenceLength("[3]tlInt32.Int32", index+1, 3)
			}
			if err := (*vec)[index].ReadJSON(legacyTypeNames, in); err != nil {
				return err
			}
			in.WantComma()
		}
		in.Delim(']')
		if !in.Ok() {
			return internal.ErrorInvalidJSON("[3]tlInt32.Int32", "expected json array's end")
		}
	}
	if index != 3 {
		return internal.ErrorWrongSequenceLength("[3]tlInt32.Int32", index+1, 3)
	}
	return nil
}

func BuiltinTuple3Int32BoxedWriteJSON(w []byte, vec *[3]tlInt32.Int32) []byte {
	return BuiltinTuple3Int32BoxedWriteJSONOpt(true, false, w, vec)
}
func BuiltinTuple3Int32BoxedWriteJSONOpt(newTypeNames bool, short bool, w []byte, vec *[3]tlInt32.Int32) []byte {
	w = append(w, '[')
	for _, elem := range *vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = elem.WriteJSONOpt(newTypeNames, short, w)
	}
	return append(w, ']')
}