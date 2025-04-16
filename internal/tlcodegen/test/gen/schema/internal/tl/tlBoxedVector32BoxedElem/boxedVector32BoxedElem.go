// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlBoxedVector32BoxedElem

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tl/tlBuiltinVectorIntBoxed"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type BoxedVector32BoxedElem struct {
	X []int32
}

func (BoxedVector32BoxedElem) TLName() string { return "boxedVector32BoxedElem" }
func (BoxedVector32BoxedElem) TLTag() uint32  { return 0x591cecd4 }

func (item *BoxedVector32BoxedElem) Reset() {
	item.X = item.X[:0]
}

func (item *BoxedVector32BoxedElem) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x1cb5c415); err != nil {
		return w, err
	}
	return tlBuiltinVectorIntBoxed.BuiltinVectorIntBoxedRead(w, &item.X)
}

// This method is general version of Write, use it instead!
func (item *BoxedVector32BoxedElem) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *BoxedVector32BoxedElem) Write(w []byte) []byte {
	w = basictl.NatWrite(w, 0x1cb5c415)
	w = tlBuiltinVectorIntBoxed.BuiltinVectorIntBoxedWrite(w, item.X)
	return w
}

func (item *BoxedVector32BoxedElem) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x591cecd4); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *BoxedVector32BoxedElem) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *BoxedVector32BoxedElem) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x591cecd4)
	return item.Write(w)
}

func (item *BoxedVector32BoxedElem) ReadResult(w []byte, ret *[]int32) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x1cb5c415); err != nil {
		return w, err
	}
	return tlBuiltinVectorIntBoxed.BuiltinVectorIntBoxedRead(w, ret)
}

func (item *BoxedVector32BoxedElem) WriteResult(w []byte, ret []int32) (_ []byte, err error) {
	w = basictl.NatWrite(w, 0x1cb5c415)
	w = tlBuiltinVectorIntBoxed.BuiltinVectorIntBoxedWrite(w, ret)
	return w, nil
}

func (item *BoxedVector32BoxedElem) ReadResultJSON(legacyTypeNames bool, in *basictl.JsonLexer, ret *[]int32) error {
	if err := tlBuiltinVectorIntBoxed.BuiltinVectorIntBoxedReadJSON(legacyTypeNames, in, ret); err != nil {
		return err
	}
	return nil
}

func (item *BoxedVector32BoxedElem) WriteResultJSON(w []byte, ret []int32) (_ []byte, err error) {
	return item.writeResultJSON(true, false, w, ret)
}

func (item *BoxedVector32BoxedElem) writeResultJSON(newTypeNames bool, short bool, w []byte, ret []int32) (_ []byte, err error) {
	w = tlBuiltinVectorIntBoxed.BuiltinVectorIntBoxedWriteJSONOpt(newTypeNames, short, w, ret)
	return w, nil
}

func (item *BoxedVector32BoxedElem) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret []int32
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *BoxedVector32BoxedElem) ReadResultWriteResultJSONOpt(newTypeNames bool, short bool, r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret []int32
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(newTypeNames, short, w, ret)
	return r, w, err
}

func (item *BoxedVector32BoxedElem) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	var ret []int32
	err := item.ReadResultJSON(true, &basictl.JsonLexer{Data: r}, &ret)
	if err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item *BoxedVector32BoxedElem) String() string {
	return string(item.WriteJSON(nil))
}

func (item *BoxedVector32BoxedElem) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propXPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "x":
				if propXPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("boxedVector32BoxedElem", "x")
				}
				if err := tlBuiltinVectorIntBoxed.BuiltinVectorIntBoxedReadJSON(legacyTypeNames, in, &item.X); err != nil {
					return err
				}
				propXPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("boxedVector32BoxedElem", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propXPresented {
		item.X = item.X[:0]
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *BoxedVector32BoxedElem) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *BoxedVector32BoxedElem) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *BoxedVector32BoxedElem) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexX := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"x":`...)
	w = tlBuiltinVectorIntBoxed.BuiltinVectorIntBoxedWriteJSONOpt(newTypeNames, short, w, item.X)
	if (len(item.X) != 0) == false {
		w = w[:backupIndexX]
	}
	return append(w, '}')
}

func (item *BoxedVector32BoxedElem) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *BoxedVector32BoxedElem) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("boxedVector32BoxedElem", err.Error())
	}
	return nil
}
