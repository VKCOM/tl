// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlTupleService2DeltaSet

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tl/tlBuiltinTupleService2DeltaSet"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tlservice2/tlService2DeltaSet"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type TupleService2DeltaSet []tlService2DeltaSet.Service2DeltaSet

func (TupleService2DeltaSet) TLName() string { return "tuple" }
func (TupleService2DeltaSet) TLTag() uint32  { return 0x9770768a }

func (item *TupleService2DeltaSet) Reset() {
	ptr := (*[]tlService2DeltaSet.Service2DeltaSet)(item)
	*ptr = (*ptr)[:0]
}

func (item *TupleService2DeltaSet) Read(w []byte, nat_tobjectIdLength uint32, nat_tintCountersNum uint32, nat_tfloatCountersNum uint32, nat_n uint32) (_ []byte, err error) {
	ptr := (*[]tlService2DeltaSet.Service2DeltaSet)(item)
	return tlBuiltinTupleService2DeltaSet.BuiltinTupleService2DeltaSetRead(w, ptr, nat_n, nat_tobjectIdLength, nat_tintCountersNum, nat_tfloatCountersNum)
}

func (item *TupleService2DeltaSet) WriteGeneral(w []byte, nat_tobjectIdLength uint32, nat_tintCountersNum uint32, nat_tfloatCountersNum uint32, nat_n uint32) (_ []byte, err error) {
	return item.Write(w, nat_tobjectIdLength, nat_tintCountersNum, nat_tfloatCountersNum, nat_n)
}

func (item *TupleService2DeltaSet) Write(w []byte, nat_tobjectIdLength uint32, nat_tintCountersNum uint32, nat_tfloatCountersNum uint32, nat_n uint32) (_ []byte, err error) {
	ptr := (*[]tlService2DeltaSet.Service2DeltaSet)(item)
	return tlBuiltinTupleService2DeltaSet.BuiltinTupleService2DeltaSetWrite(w, *ptr, nat_n, nat_tobjectIdLength, nat_tintCountersNum, nat_tfloatCountersNum)
}

func (item *TupleService2DeltaSet) ReadBoxed(w []byte, nat_tobjectIdLength uint32, nat_tintCountersNum uint32, nat_tfloatCountersNum uint32, nat_n uint32) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x9770768a); err != nil {
		return w, err
	}
	return item.Read(w, nat_tobjectIdLength, nat_tintCountersNum, nat_tfloatCountersNum, nat_n)
}

func (item *TupleService2DeltaSet) WriteBoxedGeneral(w []byte, nat_tobjectIdLength uint32, nat_tintCountersNum uint32, nat_tfloatCountersNum uint32, nat_n uint32) (_ []byte, err error) {
	return item.WriteBoxed(w, nat_tobjectIdLength, nat_tintCountersNum, nat_tfloatCountersNum, nat_n)
}

func (item *TupleService2DeltaSet) WriteBoxed(w []byte, nat_tobjectIdLength uint32, nat_tintCountersNum uint32, nat_tfloatCountersNum uint32, nat_n uint32) (_ []byte, err error) {
	w = basictl.NatWrite(w, 0x9770768a)
	return item.Write(w, nat_tobjectIdLength, nat_tintCountersNum, nat_tfloatCountersNum, nat_n)
}

func (item *TupleService2DeltaSet) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, nat_tobjectIdLength uint32, nat_tintCountersNum uint32, nat_tfloatCountersNum uint32, nat_n uint32) error {
	ptr := (*[]tlService2DeltaSet.Service2DeltaSet)(item)
	if err := tlBuiltinTupleService2DeltaSet.BuiltinTupleService2DeltaSetReadJSON(legacyTypeNames, in, ptr, nat_n, nat_tobjectIdLength, nat_tintCountersNum, nat_tfloatCountersNum); err != nil {
		return err
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *TupleService2DeltaSet) WriteJSONGeneral(tctx *basictl.JSONWriteContext, w []byte, nat_tobjectIdLength uint32, nat_tintCountersNum uint32, nat_tfloatCountersNum uint32, nat_n uint32) (_ []byte, err error) {
	return item.WriteJSONOpt(tctx, w, nat_tobjectIdLength, nat_tintCountersNum, nat_tfloatCountersNum, nat_n)
}

func (item *TupleService2DeltaSet) WriteJSON(w []byte, nat_tobjectIdLength uint32, nat_tintCountersNum uint32, nat_tfloatCountersNum uint32, nat_n uint32) (_ []byte, err error) {
	tctx := basictl.JSONWriteContext{}
	return item.WriteJSONOpt(&tctx, w, nat_tobjectIdLength, nat_tintCountersNum, nat_tfloatCountersNum, nat_n)
}

func (item *TupleService2DeltaSet) WriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte, nat_tobjectIdLength uint32, nat_tintCountersNum uint32, nat_tfloatCountersNum uint32, nat_n uint32) (_ []byte, err error) {
	ptr := (*[]tlService2DeltaSet.Service2DeltaSet)(item)
	if w, err = tlBuiltinTupleService2DeltaSet.BuiltinTupleService2DeltaSetWriteJSONOpt(tctx, w, *ptr, nat_n, nat_tobjectIdLength, nat_tintCountersNum, nat_tfloatCountersNum); err != nil {
		return w, err
	}
	return w, nil
}
