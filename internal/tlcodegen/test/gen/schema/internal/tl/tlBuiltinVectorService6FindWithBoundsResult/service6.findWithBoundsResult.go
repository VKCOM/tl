// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlBuiltinVectorService6FindWithBoundsResult

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tlservice6/tlService6FindWithBoundsResult"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

func BuiltinVectorService6FindWithBoundsResultRead(w []byte, vec *[]tlService6FindWithBoundsResult.Service6FindWithBoundsResult) (_ []byte, err error) {
	var l uint32
	if w, err = basictl.NatRead(w, &l); err != nil {
		return w, err
	}
	if err = basictl.CheckLengthSanity(w, l, 4); err != nil {
		return w, err
	}
	if uint32(cap(*vec)) < l {
		*vec = make([]tlService6FindWithBoundsResult.Service6FindWithBoundsResult, l)
	} else {
		*vec = (*vec)[:l]
	}
	for i := range *vec {
		if w, err = (*vec)[i].Read(w); err != nil {
			return w, err
		}
	}
	return w, nil
}

func BuiltinVectorService6FindWithBoundsResultWrite(w []byte, vec []tlService6FindWithBoundsResult.Service6FindWithBoundsResult) []byte {
	w = basictl.NatWrite(w, uint32(len(vec)))
	for _, elem := range vec {
		w = elem.Write(w)
	}
	return w
}

func BuiltinVectorService6FindWithBoundsResultReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, vec *[]tlService6FindWithBoundsResult.Service6FindWithBoundsResult) error {
	*vec = (*vec)[:cap(*vec)]
	index := 0
	if in != nil {
		in.Delim('[')
		if !in.Ok() {
			return internal.ErrorInvalidJSON("[]tlService6FindWithBoundsResult.Service6FindWithBoundsResult", "expected json array")
		}
		for ; !in.IsDelim(']'); index++ {
			if len(*vec) <= index {
				var newValue tlService6FindWithBoundsResult.Service6FindWithBoundsResult
				*vec = append(*vec, newValue)
				*vec = (*vec)[:cap(*vec)]
			}
			if err := (*vec)[index].ReadJSON(legacyTypeNames, in); err != nil {
				return err
			}
			in.WantComma()
		}
		in.Delim(']')
		if !in.Ok() {
			return internal.ErrorInvalidJSON("[]tlService6FindWithBoundsResult.Service6FindWithBoundsResult", "expected json array's end")
		}
	}
	*vec = (*vec)[:index]
	return nil
}

func BuiltinVectorService6FindWithBoundsResultWriteJSON(w []byte, vec []tlService6FindWithBoundsResult.Service6FindWithBoundsResult) []byte {
	return BuiltinVectorService6FindWithBoundsResultWriteJSONOpt(true, false, w, vec)
}
func BuiltinVectorService6FindWithBoundsResultWriteJSONOpt(newTypeNames bool, short bool, w []byte, vec []tlService6FindWithBoundsResult.Service6FindWithBoundsResult) []byte {
	w = append(w, '[')
	for _, elem := range vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = elem.WriteJSONOpt(newTypeNames, short, w)
	}
	return append(w, ']')
}
