// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlBoxedTupleSlice3

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tl/tlBuiltinTupleInt"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type BoxedTupleSlice3 struct {
	N uint32
	X []int32
}

func (BoxedTupleSlice3) TLName() string { return "boxedTupleSlice3" }
func (BoxedTupleSlice3) TLTag() uint32  { return 0xa19b8106 }

func (item *BoxedTupleSlice3) Reset() {
	item.N = 0
	item.X = item.X[:0]
}

func (item *BoxedTupleSlice3) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.N); err != nil {
		return w, err
	}
	if w, err = basictl.NatReadExactTag(w, 0x9770768a); err != nil {
		return w, err
	}
	return tlBuiltinTupleInt.BuiltinTupleIntRead(w, &item.X, item.N)
}

// This method is general version of Write, use it instead!
func (item *BoxedTupleSlice3) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w)
}

func (item *BoxedTupleSlice3) Write(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, item.N)
	w = basictl.NatWrite(w, 0x9770768a)
	if w, err = tlBuiltinTupleInt.BuiltinTupleIntWrite(w, item.X, item.N); err != nil {
		return w, err
	}
	return w, nil
}

func (item *BoxedTupleSlice3) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xa19b8106); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *BoxedTupleSlice3) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w)
}

func (item *BoxedTupleSlice3) WriteBoxed(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, 0xa19b8106)
	return item.Write(w)
}

func (item *BoxedTupleSlice3) ReadResult(w []byte, ret *[]int32) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x9770768a); err != nil {
		return w, err
	}
	return tlBuiltinTupleInt.BuiltinTupleIntRead(w, ret, item.N)
}

func (item *BoxedTupleSlice3) WriteResult(w []byte, ret []int32) (_ []byte, err error) {
	w = basictl.NatWrite(w, 0x9770768a)
	if w, err = tlBuiltinTupleInt.BuiltinTupleIntWrite(w, ret, item.N); err != nil {
		return w, err
	}
	return w, nil
}

func (item *BoxedTupleSlice3) ReadResultJSON(legacyTypeNames bool, in *basictl.JsonLexer, ret *[]int32) error {
	if err := tlBuiltinTupleInt.BuiltinTupleIntReadJSON(legacyTypeNames, in, ret, item.N); err != nil {
		return err
	}
	return nil
}

func (item *BoxedTupleSlice3) WriteResultJSON(w []byte, ret []int32) (_ []byte, err error) {
	return item.writeResultJSON(true, false, w, ret)
}

func (item *BoxedTupleSlice3) writeResultJSON(newTypeNames bool, short bool, w []byte, ret []int32) (_ []byte, err error) {
	if w, err = tlBuiltinTupleInt.BuiltinTupleIntWriteJSONOpt(newTypeNames, short, w, ret, item.N); err != nil {
		return w, err
	}
	return w, nil
}

func (item *BoxedTupleSlice3) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret []int32
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *BoxedTupleSlice3) ReadResultWriteResultJSONOpt(newTypeNames bool, short bool, r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret []int32
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(newTypeNames, short, w, ret)
	return r, w, err
}

func (item *BoxedTupleSlice3) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	var ret []int32
	err := item.ReadResultJSON(true, &basictl.JsonLexer{Data: r}, &ret)
	if err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item *BoxedTupleSlice3) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func (item *BoxedTupleSlice3) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propNPresented bool
	var rawX []byte

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "n":
				if propNPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("boxedTupleSlice3", "n")
				}
				if err := internal.Json2ReadUint32(in, &item.N); err != nil {
					return err
				}
				propNPresented = true
			case "x":
				if rawX != nil {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("boxedTupleSlice3", "x")
				}
				rawX = in.Raw()
				if !in.Ok() {
					return in.Error()
				}
			default:
				return internal.ErrorInvalidJSONExcessElement("boxedTupleSlice3", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propNPresented {
		item.N = 0
	}
	var inXPointer *basictl.JsonLexer
	inX := basictl.JsonLexer{Data: rawX}
	if rawX != nil {
		inXPointer = &inX
	}
	if err := tlBuiltinTupleInt.BuiltinTupleIntReadJSON(legacyTypeNames, inXPointer, &item.X, item.N); err != nil {
		return err
	}

	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *BoxedTupleSlice3) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w)
}

func (item *BoxedTupleSlice3) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w)
}
func (item *BoxedTupleSlice3) WriteJSONOpt(newTypeNames bool, short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	backupIndexN := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"n":`...)
	w = basictl.JSONWriteUint32(w, item.N)
	if (item.N != 0) == false {
		w = w[:backupIndexN]
	}
	backupIndexX := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"x":`...)
	if w, err = tlBuiltinTupleInt.BuiltinTupleIntWriteJSONOpt(newTypeNames, short, w, item.X, item.N); err != nil {
		return w, err
	}
	if (len(item.X) != 0) == false {
		w = w[:backupIndexX]
	}
	return append(w, '}'), nil
}

func (item *BoxedTupleSlice3) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *BoxedTupleSlice3) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("boxedTupleSlice3", err.Error())
	}
	return nil
}
