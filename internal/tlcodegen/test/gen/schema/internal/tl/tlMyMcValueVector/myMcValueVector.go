// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlMyMcValueVector

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/cycle_6ca945392bbf8b14f24e5653edc8b214"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tl/tlBuiltinVectorService1Value"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type MyMcValueVector struct {
	Xs []cycle_6ca945392bbf8b14f24e5653edc8b214.Service1Value
}

func (MyMcValueVector) TLName() string { return "myMcValueVector" }
func (MyMcValueVector) TLTag() uint32  { return 0x761d6d58 }

func (item *MyMcValueVector) Reset() {
	item.Xs = item.Xs[:0]
}

func (item *MyMcValueVector) Read(w []byte) (_ []byte, err error) {
	return tlBuiltinVectorService1Value.BuiltinVectorService1ValueRead(w, &item.Xs)
}

func (item *MyMcValueVector) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *MyMcValueVector) Write(w []byte) []byte {
	w = tlBuiltinVectorService1Value.BuiltinVectorService1ValueWrite(w, item.Xs)
	return w
}

func (item *MyMcValueVector) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x761d6d58); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *MyMcValueVector) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *MyMcValueVector) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x761d6d58)
	return item.Write(w)
}

func (item MyMcValueVector) String() string {
	return string(item.WriteJSON(nil))
}

func (item *MyMcValueVector) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propXsPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "xs":
				if propXsPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("myMcValueVector", "xs")
				}
				if err := tlBuiltinVectorService1Value.BuiltinVectorService1ValueReadJSON(legacyTypeNames, in, &item.Xs); err != nil {
					return err
				}
				propXsPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("myMcValueVector", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propXsPresented {
		item.Xs = item.Xs[:0]
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *MyMcValueVector) WriteJSONGeneral(tctx *basictl.JSONWriteContext, w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(tctx, w), nil
}

func (item *MyMcValueVector) WriteJSON(w []byte) []byte {
	tctx := basictl.JSONWriteContext{}
	return item.WriteJSONOpt(&tctx, w)
}
func (item *MyMcValueVector) WriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte) []byte {
	w = append(w, '{')
	backupIndexXs := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"xs":`...)
	w = tlBuiltinVectorService1Value.BuiltinVectorService1ValueWriteJSONOpt(tctx, w, item.Xs)
	if (len(item.Xs) != 0) == false {
		w = w[:backupIndexXs]
	}
	return append(w, '}')
}

func (item *MyMcValueVector) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *MyMcValueVector) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("myMcValueVector", err.Error())
	}
	return nil
}
