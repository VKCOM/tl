// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlBuiltinVectorService3GroupCountLimit

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tlservice3/tlService3GroupCountLimit"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

func BuiltinVectorService3GroupCountLimitRead(w []byte, vec *[]tlService3GroupCountLimit.Service3GroupCountLimit) (_ []byte, err error) {
	var l uint32
	if w, err = basictl.NatRead(w, &l); err != nil {
		return w, err
	}
	if err = basictl.CheckLengthSanity(w, l, 4); err != nil {
		return w, err
	}
	if uint32(cap(*vec)) < l {
		*vec = make([]tlService3GroupCountLimit.Service3GroupCountLimit, l)
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

func BuiltinVectorService3GroupCountLimitWrite(w []byte, vec []tlService3GroupCountLimit.Service3GroupCountLimit) []byte {
	w = basictl.NatWrite(w, uint32(len(vec)))
	for _, elem := range vec {
		w = elem.Write(w)
	}
	return w
}

func BuiltinVectorService3GroupCountLimitReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, vec *[]tlService3GroupCountLimit.Service3GroupCountLimit) error {
	*vec = (*vec)[:cap(*vec)]
	index := 0
	if in != nil {
		in.Delim('[')
		if !in.Ok() {
			return internal.ErrorInvalidJSON("[]tlService3GroupCountLimit.Service3GroupCountLimit", "expected json array")
		}
		for ; !in.IsDelim(']'); index++ {
			if len(*vec) <= index {
				var newValue tlService3GroupCountLimit.Service3GroupCountLimit
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
			return internal.ErrorInvalidJSON("[]tlService3GroupCountLimit.Service3GroupCountLimit", "expected json array's end")
		}
	}
	*vec = (*vec)[:index]
	return nil
}

func BuiltinVectorService3GroupCountLimitWriteJSON(w []byte, vec []tlService3GroupCountLimit.Service3GroupCountLimit) []byte {
	return BuiltinVectorService3GroupCountLimitWriteJSONOpt(true, false, w, vec)
}
func BuiltinVectorService3GroupCountLimitWriteJSONOpt(newTypeNames bool, short bool, w []byte, vec []tlService3GroupCountLimit.Service3GroupCountLimit) []byte {
	w = append(w, '[')
	for _, elem := range vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = elem.WriteJSONOpt(newTypeNames, short, w)
	}
	return append(w, ']')
}