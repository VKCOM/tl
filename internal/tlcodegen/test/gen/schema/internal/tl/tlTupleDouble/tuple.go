// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlTupleDouble

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tl/tlBuiltinTupleDouble"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type TupleDouble []float64

func (TupleDouble) TLName() string { return "tuple" }
func (TupleDouble) TLTag() uint32  { return 0x9770768a }

func (item *TupleDouble) Reset() {
	ptr := (*[]float64)(item)
	*ptr = (*ptr)[:0]
}

func (item *TupleDouble) Read(w []byte, nat_n uint32) (_ []byte, err error) {
	ptr := (*[]float64)(item)
	return tlBuiltinTupleDouble.BuiltinTupleDoubleRead(w, ptr, nat_n)
}

func (item *TupleDouble) WriteGeneral(w []byte, nat_n uint32) (_ []byte, err error) {
	return item.Write(w, nat_n)
}

func (item *TupleDouble) Write(w []byte, nat_n uint32) (_ []byte, err error) {
	ptr := (*[]float64)(item)
	return tlBuiltinTupleDouble.BuiltinTupleDoubleWrite(w, *ptr, nat_n)
}

func (item *TupleDouble) ReadBoxed(w []byte, nat_n uint32) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x9770768a); err != nil {
		return w, err
	}
	return item.Read(w, nat_n)
}

func (item *TupleDouble) WriteBoxedGeneral(w []byte, nat_n uint32) (_ []byte, err error) {
	return item.WriteBoxed(w, nat_n)
}

func (item *TupleDouble) WriteBoxed(w []byte, nat_n uint32) (_ []byte, err error) {
	w = basictl.NatWrite(w, 0x9770768a)
	return item.Write(w, nat_n)
}

func (item *TupleDouble) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, nat_n uint32) error {
	ptr := (*[]float64)(item)
	if err := tlBuiltinTupleDouble.BuiltinTupleDoubleReadJSON(legacyTypeNames, in, ptr, nat_n); err != nil {
		return err
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *TupleDouble) WriteJSONGeneral(tctx *basictl.JSONWriteContext, w []byte, nat_n uint32) (_ []byte, err error) {
	return item.WriteJSONOpt(tctx, w, nat_n)
}

func (item *TupleDouble) WriteJSON(w []byte, nat_n uint32) (_ []byte, err error) {
	tctx := basictl.JSONWriteContext{}
	return item.WriteJSONOpt(&tctx, w, nat_n)
}

func (item *TupleDouble) WriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte, nat_n uint32) (_ []byte, err error) {
	ptr := (*[]float64)(item)
	if w, err = tlBuiltinTupleDouble.BuiltinTupleDoubleWriteJSONOpt(tctx, w, *ptr, nat_n); err != nil {
		return w, err
	}
	return w, nil
}
