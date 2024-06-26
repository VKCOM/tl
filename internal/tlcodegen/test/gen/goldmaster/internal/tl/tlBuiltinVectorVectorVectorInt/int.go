// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlBuiltinVectorVectorVectorInt

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlBuiltinVectorVectorInt"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

func BuiltinVectorVectorVectorIntFillRandom(rg *basictl.RandGenerator, vec *[][][]int32) {
	rg.IncreaseDepth()
	l := rg.LimitValue(basictl.RandomUint(rg))
	*vec = make([][][]int32, l)
	for i := range *vec {
		tlBuiltinVectorVectorInt.BuiltinVectorVectorIntFillRandom(rg, &(*vec)[i])
	}
	rg.DecreaseDepth()
}
func BuiltinVectorVectorVectorIntRead(w []byte, vec *[][][]int32) (_ []byte, err error) {
	var l uint32
	if w, err = basictl.NatRead(w, &l); err != nil {
		return w, err
	}
	if err = basictl.CheckLengthSanity(w, l, 4); err != nil {
		return w, err
	}
	if uint32(cap(*vec)) < l {
		*vec = make([][][]int32, l)
	} else {
		*vec = (*vec)[:l]
	}
	for i := range *vec {
		if w, err = tlBuiltinVectorVectorInt.BuiltinVectorVectorIntRead(w, &(*vec)[i]); err != nil {
			return w, err
		}
	}
	return w, nil
}

func BuiltinVectorVectorVectorIntWrite(w []byte, vec [][][]int32) []byte {
	w = basictl.NatWrite(w, uint32(len(vec)))
	for _, elem := range vec {
		w = tlBuiltinVectorVectorInt.BuiltinVectorVectorIntWrite(w, elem)
	}
	return w
}

func BuiltinVectorVectorVectorIntReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, vec *[][][]int32) error {
	*vec = (*vec)[:cap(*vec)]
	index := 0
	if in != nil {
		in.Delim('[')
		if !in.Ok() {
			return internal.ErrorInvalidJSON("[][][]int32", "expected json array")
		}
		for ; !in.IsDelim(']'); index++ {
			if len(*vec) <= index {
				var newValue [][]int32
				*vec = append(*vec, newValue)
				*vec = (*vec)[:cap(*vec)]
			}
			if err := tlBuiltinVectorVectorInt.BuiltinVectorVectorIntReadJSON(legacyTypeNames, in, &(*vec)[index]); err != nil {
				return err
			}
			in.WantComma()
		}
		in.Delim(']')
		if !in.Ok() {
			return internal.ErrorInvalidJSON("[][][]int32", "expected json array's end")
		}
	}
	*vec = (*vec)[:index]
	return nil
}

func BuiltinVectorVectorVectorIntWriteJSON(w []byte, vec [][][]int32) []byte {
	return BuiltinVectorVectorVectorIntWriteJSONOpt(true, false, w, vec)
}
func BuiltinVectorVectorVectorIntWriteJSONOpt(newTypeNames bool, short bool, w []byte, vec [][][]int32) []byte {
	w = append(w, '[')
	for _, elem := range vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = tlBuiltinVectorVectorInt.BuiltinVectorVectorIntWriteJSONOpt(newTypeNames, short, w, elem)
	}
	return append(w, ']')
}
