// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlMyTwoDicts

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tl/tlBuiltinVectorDictionaryFieldInt"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type MyTwoDicts struct {
	A map[string]int32
	B map[string]int32
}

func (MyTwoDicts) TLName() string { return "myTwoDicts" }
func (MyTwoDicts) TLTag() uint32  { return 0xa859581d }

func (item *MyTwoDicts) Reset() {
	tlBuiltinVectorDictionaryFieldInt.BuiltinVectorDictionaryFieldIntReset(item.A)
	tlBuiltinVectorDictionaryFieldInt.BuiltinVectorDictionaryFieldIntReset(item.B)
}

func (item *MyTwoDicts) Read(w []byte) (_ []byte, err error) {
	if w, err = tlBuiltinVectorDictionaryFieldInt.BuiltinVectorDictionaryFieldIntRead(w, &item.A); err != nil {
		return w, err
	}
	return tlBuiltinVectorDictionaryFieldInt.BuiltinVectorDictionaryFieldIntRead(w, &item.B)
}

// This method is general version of Write, use it instead!
func (item *MyTwoDicts) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *MyTwoDicts) Write(w []byte) []byte {
	w = tlBuiltinVectorDictionaryFieldInt.BuiltinVectorDictionaryFieldIntWrite(w, item.A)
	w = tlBuiltinVectorDictionaryFieldInt.BuiltinVectorDictionaryFieldIntWrite(w, item.B)
	return w
}

func (item *MyTwoDicts) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xa859581d); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *MyTwoDicts) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *MyTwoDicts) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0xa859581d)
	return item.Write(w)
}

func (item *MyTwoDicts) String() string {
	return string(item.WriteJSON(nil))
}

func (item *MyTwoDicts) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propAPresented bool
	var propBPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "a":
				if propAPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("myTwoDicts", "a")
				}
				if err := tlBuiltinVectorDictionaryFieldInt.BuiltinVectorDictionaryFieldIntReadJSON(legacyTypeNames, in, &item.A); err != nil {
					return err
				}
				propAPresented = true
			case "b":
				if propBPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("myTwoDicts", "b")
				}
				if err := tlBuiltinVectorDictionaryFieldInt.BuiltinVectorDictionaryFieldIntReadJSON(legacyTypeNames, in, &item.B); err != nil {
					return err
				}
				propBPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("myTwoDicts", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propAPresented {
		tlBuiltinVectorDictionaryFieldInt.BuiltinVectorDictionaryFieldIntReset(item.A)
	}
	if !propBPresented {
		tlBuiltinVectorDictionaryFieldInt.BuiltinVectorDictionaryFieldIntReset(item.B)
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *MyTwoDicts) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *MyTwoDicts) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *MyTwoDicts) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexA := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"a":`...)
	w = tlBuiltinVectorDictionaryFieldInt.BuiltinVectorDictionaryFieldIntWriteJSONOpt(newTypeNames, short, w, item.A)
	if (len(item.A) != 0) == false {
		w = w[:backupIndexA]
	}
	backupIndexB := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"b":`...)
	w = tlBuiltinVectorDictionaryFieldInt.BuiltinVectorDictionaryFieldIntWriteJSONOpt(newTypeNames, short, w, item.B)
	if (len(item.B) != 0) == false {
		w = w[:backupIndexB]
	}
	return append(w, '}')
}

func (item *MyTwoDicts) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *MyTwoDicts) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("myTwoDicts", err.Error())
	}
	return nil
}
