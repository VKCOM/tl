// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlVectorString

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tl/tlBuiltinVectorString"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type VectorString []string

func (VectorString) TLName() string { return "vector" }
func (VectorString) TLTag() uint32  { return 0x1cb5c415 }

func (item *VectorString) Reset() {
	ptr := (*[]string)(item)
	*ptr = (*ptr)[:0]
}

func (item *VectorString) FillRandom(rg *basictl.RandGenerator) {
	ptr := (*[]string)(item)
	tlBuiltinVectorString.BuiltinVectorStringFillRandom(rg, ptr)
}

func (item *VectorString) Read(w []byte) (_ []byte, err error) {
	ptr := (*[]string)(item)
	return tlBuiltinVectorString.BuiltinVectorStringRead(w, ptr)
}

func (item *VectorString) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *VectorString) Write(w []byte) []byte {
	ptr := (*[]string)(item)
	return tlBuiltinVectorString.BuiltinVectorStringWrite(w, *ptr)
}

func (item *VectorString) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x1cb5c415); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *VectorString) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *VectorString) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x1cb5c415)
	return item.Write(w)
}

func (item VectorString) String() string {
	return string(item.WriteJSON(nil))
}
func (item *VectorString) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	ptr := (*[]string)(item)
	if err := tlBuiltinVectorString.BuiltinVectorStringReadJSON(legacyTypeNames, in, ptr); err != nil {
		return err
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *VectorString) WriteJSONGeneral(tctx *basictl.JSONWriteContext, w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(tctx, w), nil
}

func (item *VectorString) WriteJSON(w []byte) []byte {
	tctx := basictl.JSONWriteContext{}
	return item.WriteJSONOpt(&tctx, w)
}

func (item *VectorString) WriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte) []byte {
	ptr := (*[]string)(item)
	w = tlBuiltinVectorString.BuiltinVectorStringWriteJSONOpt(tctx, w, *ptr)
	return w
}
func (item *VectorString) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *VectorString) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("vector", err.Error())
	}
	return nil
}

func (item *VectorString) CalculateLayout(sizes []int) []int {
	ptr := (*[]string)(item)
	sizes = tlBuiltinVectorString.BuiltinVectorStringCalculateLayout(sizes, ptr)
	return sizes
}

func (item *VectorString) InternalWriteTL2(w []byte, sizes []int) ([]byte, []int) {
	ptr := (*[]string)(item)
	w, sizes = tlBuiltinVectorString.BuiltinVectorStringInternalWriteTL2(w, sizes, ptr)
	return w, sizes
}

func (item *VectorString) WriteTL2(w []byte, ctx *basictl.TL2WriteContext) []byte {
	var sizes []int
	if ctx != nil {
		sizes = ctx.SizeBuffer
	}
	sizes = item.CalculateLayout(sizes[:0])
	w, _ = item.InternalWriteTL2(w, sizes)
	if ctx != nil {
		ctx.SizeBuffer = sizes[:0]
	}
	return w
}

func (item *VectorString) InternalReadTL2(r []byte) (_ []byte, err error) {
	ptr := (*[]string)(item)
	if r, err = tlBuiltinVectorString.BuiltinVectorStringInternalReadTL2(r, ptr); err != nil {
		return r, err
	}
	return r, nil
}

func (item *VectorString) ReadTL2(r []byte, ctx *basictl.TL2ReadContext) (_ []byte, err error) {
	return item.InternalReadTL2(r)
}

type VectorStringBytes [][]byte

func (VectorStringBytes) TLName() string { return "vector" }
func (VectorStringBytes) TLTag() uint32  { return 0x1cb5c415 }

func (item *VectorStringBytes) Reset() {
	ptr := (*[][]byte)(item)
	*ptr = (*ptr)[:0]
}

func (item *VectorStringBytes) FillRandom(rg *basictl.RandGenerator) {
	ptr := (*[][]byte)(item)
	tlBuiltinVectorString.BuiltinVectorStringBytesFillRandom(rg, ptr)
}

func (item *VectorStringBytes) Read(w []byte) (_ []byte, err error) {
	ptr := (*[][]byte)(item)
	return tlBuiltinVectorString.BuiltinVectorStringBytesRead(w, ptr)
}

func (item *VectorStringBytes) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *VectorStringBytes) Write(w []byte) []byte {
	ptr := (*[][]byte)(item)
	return tlBuiltinVectorString.BuiltinVectorStringBytesWrite(w, *ptr)
}

func (item *VectorStringBytes) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x1cb5c415); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *VectorStringBytes) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *VectorStringBytes) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x1cb5c415)
	return item.Write(w)
}

func (item VectorStringBytes) String() string {
	return string(item.WriteJSON(nil))
}
func (item *VectorStringBytes) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	ptr := (*[][]byte)(item)
	if err := tlBuiltinVectorString.BuiltinVectorStringBytesReadJSON(legacyTypeNames, in, ptr); err != nil {
		return err
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *VectorStringBytes) WriteJSONGeneral(tctx *basictl.JSONWriteContext, w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(tctx, w), nil
}

func (item *VectorStringBytes) WriteJSON(w []byte) []byte {
	tctx := basictl.JSONWriteContext{}
	return item.WriteJSONOpt(&tctx, w)
}

func (item *VectorStringBytes) WriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte) []byte {
	ptr := (*[][]byte)(item)
	w = tlBuiltinVectorString.BuiltinVectorStringBytesWriteJSONOpt(tctx, w, *ptr)
	return w
}
func (item *VectorStringBytes) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *VectorStringBytes) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("vector", err.Error())
	}
	return nil
}

func (item *VectorStringBytes) CalculateLayout(sizes []int) []int {
	ptr := (*[][]byte)(item)
	sizes = tlBuiltinVectorString.BuiltinVectorStringBytesCalculateLayout(sizes, ptr)
	return sizes
}

func (item *VectorStringBytes) InternalWriteTL2(w []byte, sizes []int) ([]byte, []int) {
	ptr := (*[][]byte)(item)
	w, sizes = tlBuiltinVectorString.BuiltinVectorStringBytesInternalWriteTL2(w, sizes, ptr)
	return w, sizes
}

func (item *VectorStringBytes) WriteTL2(w []byte, ctx *basictl.TL2WriteContext) []byte {
	var sizes []int
	if ctx != nil {
		sizes = ctx.SizeBuffer
	}
	sizes = item.CalculateLayout(sizes[:0])
	w, _ = item.InternalWriteTL2(w, sizes)
	if ctx != nil {
		ctx.SizeBuffer = sizes[:0]
	}
	return w
}

func (item *VectorStringBytes) InternalReadTL2(r []byte) (_ []byte, err error) {
	ptr := (*[][]byte)(item)
	if r, err = tlBuiltinVectorString.BuiltinVectorStringBytesInternalReadTL2(r, ptr); err != nil {
		return r, err
	}
	return r, nil
}

func (item *VectorStringBytes) ReadTL2(r []byte, ctx *basictl.TL2ReadContext) (_ []byte, err error) {
	return item.InternalReadTL2(r)
}
