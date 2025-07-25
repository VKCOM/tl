// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlVectorService6FindWithBoundsResult

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tl/tlBuiltinVectorService6FindWithBoundsResult"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tlservice6/tlService6FindWithBoundsResult"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type VectorService6FindWithBoundsResult []tlService6FindWithBoundsResult.Service6FindWithBoundsResult

func (VectorService6FindWithBoundsResult) TLName() string { return "vector" }
func (VectorService6FindWithBoundsResult) TLTag() uint32  { return 0x1cb5c415 }

func (item *VectorService6FindWithBoundsResult) Reset() {
	ptr := (*[]tlService6FindWithBoundsResult.Service6FindWithBoundsResult)(item)
	*ptr = (*ptr)[:0]
}

func (item *VectorService6FindWithBoundsResult) Read(w []byte) (_ []byte, err error) {
	ptr := (*[]tlService6FindWithBoundsResult.Service6FindWithBoundsResult)(item)
	return tlBuiltinVectorService6FindWithBoundsResult.BuiltinVectorService6FindWithBoundsResultRead(w, ptr)
}

func (item *VectorService6FindWithBoundsResult) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *VectorService6FindWithBoundsResult) Write(w []byte) []byte {
	ptr := (*[]tlService6FindWithBoundsResult.Service6FindWithBoundsResult)(item)
	return tlBuiltinVectorService6FindWithBoundsResult.BuiltinVectorService6FindWithBoundsResultWrite(w, *ptr)
}

func (item *VectorService6FindWithBoundsResult) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x1cb5c415); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *VectorService6FindWithBoundsResult) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *VectorService6FindWithBoundsResult) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x1cb5c415)
	return item.Write(w)
}

func (item VectorService6FindWithBoundsResult) String() string {
	return string(item.WriteJSON(nil))
}
func (item *VectorService6FindWithBoundsResult) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	ptr := (*[]tlService6FindWithBoundsResult.Service6FindWithBoundsResult)(item)
	if err := tlBuiltinVectorService6FindWithBoundsResult.BuiltinVectorService6FindWithBoundsResultReadJSON(legacyTypeNames, in, ptr); err != nil {
		return err
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *VectorService6FindWithBoundsResult) WriteJSONGeneral(tctx *basictl.JSONWriteContext, w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(tctx, w), nil
}

func (item *VectorService6FindWithBoundsResult) WriteJSON(w []byte) []byte {
	tctx := basictl.JSONWriteContext{}
	return item.WriteJSONOpt(&tctx, w)
}

func (item *VectorService6FindWithBoundsResult) WriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte) []byte {
	ptr := (*[]tlService6FindWithBoundsResult.Service6FindWithBoundsResult)(item)
	w = tlBuiltinVectorService6FindWithBoundsResult.BuiltinVectorService6FindWithBoundsResultWriteJSONOpt(tctx, w, *ptr)
	return w
}
func (item *VectorService6FindWithBoundsResult) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *VectorService6FindWithBoundsResult) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("vector", err.Error())
	}
	return nil
}
