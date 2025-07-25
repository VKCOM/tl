// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlMyDictOfInt

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tl/tlBuiltinVectorDictionaryFieldInt"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type MyDictOfInt map[string]int32

func (MyDictOfInt) TLName() string { return "myDictOfInt" }
func (MyDictOfInt) TLTag() uint32  { return 0xb8019a3d }

func (item *MyDictOfInt) Reset() {
	ptr := (*map[string]int32)(item)
	tlBuiltinVectorDictionaryFieldInt.BuiltinVectorDictionaryFieldIntReset(*ptr)
}

func (item *MyDictOfInt) Read(w []byte) (_ []byte, err error) {
	ptr := (*map[string]int32)(item)
	return tlBuiltinVectorDictionaryFieldInt.BuiltinVectorDictionaryFieldIntRead(w, ptr)
}

func (item *MyDictOfInt) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *MyDictOfInt) Write(w []byte) []byte {
	ptr := (*map[string]int32)(item)
	return tlBuiltinVectorDictionaryFieldInt.BuiltinVectorDictionaryFieldIntWrite(w, *ptr)
}

func (item *MyDictOfInt) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xb8019a3d); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *MyDictOfInt) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *MyDictOfInt) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0xb8019a3d)
	return item.Write(w)
}

func (item MyDictOfInt) String() string {
	return string(item.WriteJSON(nil))
}
func (item *MyDictOfInt) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	ptr := (*map[string]int32)(item)
	if err := tlBuiltinVectorDictionaryFieldInt.BuiltinVectorDictionaryFieldIntReadJSON(legacyTypeNames, in, ptr); err != nil {
		return err
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *MyDictOfInt) WriteJSONGeneral(tctx *basictl.JSONWriteContext, w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(tctx, w), nil
}

func (item *MyDictOfInt) WriteJSON(w []byte) []byte {
	tctx := basictl.JSONWriteContext{}
	return item.WriteJSONOpt(&tctx, w)
}

func (item *MyDictOfInt) WriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte) []byte {
	ptr := (*map[string]int32)(item)
	w = tlBuiltinVectorDictionaryFieldInt.BuiltinVectorDictionaryFieldIntWriteJSONOpt(tctx, w, *ptr)
	return w
}
func (item *MyDictOfInt) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *MyDictOfInt) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("myDictOfInt", err.Error())
	}
	return nil
}
