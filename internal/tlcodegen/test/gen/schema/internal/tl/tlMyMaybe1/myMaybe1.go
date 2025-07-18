// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlMyMaybe1

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tl/tlMyTuple10Maybe"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type MyMaybe1 tlMyTuple10Maybe.MyTuple10Maybe

func (MyMaybe1) TLName() string { return "myMaybe1" }
func (MyMaybe1) TLTag() uint32  { return 0x32c541fe }

func (item *MyMaybe1) Reset() {
	ptr := (*tlMyTuple10Maybe.MyTuple10Maybe)(item)
	ptr.Reset()
}

func (item *MyMaybe1) Read(w []byte) (_ []byte, err error) {
	ptr := (*tlMyTuple10Maybe.MyTuple10Maybe)(item)
	return ptr.ReadBoxed(w)
}

func (item *MyMaybe1) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *MyMaybe1) Write(w []byte) []byte {
	ptr := (*tlMyTuple10Maybe.MyTuple10Maybe)(item)
	return ptr.WriteBoxed(w)
}

func (item *MyMaybe1) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x32c541fe); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *MyMaybe1) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *MyMaybe1) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x32c541fe)
	return item.Write(w)
}

func (item MyMaybe1) String() string {
	return string(item.WriteJSON(nil))
}
func (item *MyMaybe1) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	ptr := (*tlMyTuple10Maybe.MyTuple10Maybe)(item)
	if err := ptr.ReadJSON(legacyTypeNames, in); err != nil {
		return err
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *MyMaybe1) WriteJSONGeneral(tctx *basictl.JSONWriteContext, w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(tctx, w), nil
}

func (item *MyMaybe1) WriteJSON(w []byte) []byte {
	tctx := basictl.JSONWriteContext{}
	return item.WriteJSONOpt(&tctx, w)
}

func (item *MyMaybe1) WriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte) []byte {
	ptr := (*tlMyTuple10Maybe.MyTuple10Maybe)(item)
	w = ptr.WriteJSONOpt(tctx, w)
	return w
}
func (item *MyMaybe1) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *MyMaybe1) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("myMaybe1", err.Error())
	}
	return nil
}
