// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlMyInt32

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlInt32"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type MyInt32 tlInt32.Int32

func (MyInt32) TLName() string { return "myInt32" }
func (MyInt32) TLTag() uint32  { return 0xba59e151 }

func (item *MyInt32) Reset() {
	ptr := (*tlInt32.Int32)(item)
	ptr.Reset()
}

func (item *MyInt32) FillRandom(rg *basictl.RandGenerator) {
	ptr := (*tlInt32.Int32)(item)
	ptr.FillRandom(rg)
}

func (item *MyInt32) Read(w []byte) (_ []byte, err error) {
	ptr := (*tlInt32.Int32)(item)
	return ptr.Read(w)
}

// This method is general version of Write, use it instead!
func (item *MyInt32) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *MyInt32) Write(w []byte) []byte {
	ptr := (*tlInt32.Int32)(item)
	return ptr.Write(w)
}

func (item *MyInt32) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xba59e151); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *MyInt32) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *MyInt32) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0xba59e151)
	return item.Write(w)
}

func (item *MyInt32) String() string {
	return string(item.WriteJSON(nil))
}

func (item *MyInt32) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	ptr := (*tlInt32.Int32)(item)
	if err := ptr.ReadJSON(legacyTypeNames, in); err != nil {
		return err
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *MyInt32) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSON(w), nil
}

func (item *MyInt32) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}

func (item *MyInt32) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	ptr := (*tlInt32.Int32)(item)
	w = ptr.WriteJSONOpt(newTypeNames, short, w)
	return w
}
func (item *MyInt32) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *MyInt32) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("myInt32", err.Error())
	}
	return nil
}
