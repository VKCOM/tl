// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package cycle_e0e20d17984738a60e3bf741bd7db2e9

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tl/tlBuiltinVectorService6FindResultRow"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tlservice6/tlService6Error"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tlservice6/tlService6FindResultRow"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

var _EitherService6ErrorVectorService6FindResultRow = [2]internal.UnionElement{
	{TLTag: 0x0a29cd5d, TLName: "left", TLString: "left#0a29cd5d"},
	{TLTag: 0xdf3ecb3b, TLName: "right", TLString: "right#df3ecb3b"},
}

type EitherService6ErrorVectorService6FindResultRow struct {
	valueLeft  LeftService6ErrorVectorService6FindResultRow
	valueRight RightService6ErrorVectorService6FindResultRow
	index      int
}

func (item EitherService6ErrorVectorService6FindResultRow) TLName() string {
	return _EitherService6ErrorVectorService6FindResultRow[item.index].TLName
}
func (item EitherService6ErrorVectorService6FindResultRow) TLTag() uint32 {
	return _EitherService6ErrorVectorService6FindResultRow[item.index].TLTag
}

func (item *EitherService6ErrorVectorService6FindResultRow) Reset() { item.ResetToLeft() }

func (item *EitherService6ErrorVectorService6FindResultRow) IsLeft() bool { return item.index == 0 }

func (item *EitherService6ErrorVectorService6FindResultRow) AsLeft() (*LeftService6ErrorVectorService6FindResultRow, bool) {
	if item.index != 0 {
		return nil, false
	}
	return &item.valueLeft, true
}
func (item *EitherService6ErrorVectorService6FindResultRow) ResetToLeft() *LeftService6ErrorVectorService6FindResultRow {
	item.index = 0
	item.valueLeft.Reset()
	return &item.valueLeft
}
func (item *EitherService6ErrorVectorService6FindResultRow) SetLeft(value LeftService6ErrorVectorService6FindResultRow) {
	item.index = 0
	item.valueLeft = value
}

func (item *EitherService6ErrorVectorService6FindResultRow) IsRight() bool { return item.index == 1 }

func (item *EitherService6ErrorVectorService6FindResultRow) AsRight() (*RightService6ErrorVectorService6FindResultRow, bool) {
	if item.index != 1 {
		return nil, false
	}
	return &item.valueRight, true
}
func (item *EitherService6ErrorVectorService6FindResultRow) ResetToRight() *RightService6ErrorVectorService6FindResultRow {
	item.index = 1
	item.valueRight.Reset()
	return &item.valueRight
}
func (item *EitherService6ErrorVectorService6FindResultRow) SetRight(value RightService6ErrorVectorService6FindResultRow) {
	item.index = 1
	item.valueRight = value
}

func (item *EitherService6ErrorVectorService6FindResultRow) ReadBoxed(w []byte) (_ []byte, err error) {
	var tag uint32
	if w, err = basictl.NatRead(w, &tag); err != nil {
		return w, err
	}
	switch tag {
	case 0x0a29cd5d:
		item.index = 0
		return item.valueLeft.Read(w)
	case 0xdf3ecb3b:
		item.index = 1
		return item.valueRight.Read(w)
	default:
		return w, internal.ErrorInvalidUnionTag("Either", tag)
	}
}

func (item *EitherService6ErrorVectorService6FindResultRow) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *EitherService6ErrorVectorService6FindResultRow) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, _EitherService6ErrorVectorService6FindResultRow[item.index].TLTag)
	switch item.index {
	case 0:
		w = item.valueLeft.Write(w)
	case 1:
		w = item.valueRight.Write(w)
	}
	return w
}

func (item *EitherService6ErrorVectorService6FindResultRow) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	_tag, _value, err := internal.Json2ReadUnion("Either", in)
	if err != nil {
		return err
	}
	switch _tag {
	case "left#0a29cd5d", "left", "#0a29cd5d":
		if !legacyTypeNames && _tag == "left#0a29cd5d" {
			return internal.ErrorInvalidUnionLegacyTagJSON("Either", "left#0a29cd5d")
		}
		item.index = 0
		var in2Pointer *basictl.JsonLexer
		if _value != nil {
			in2 := basictl.JsonLexer{Data: _value}
			in2Pointer = &in2
		}
		if err := item.valueLeft.ReadJSON(legacyTypeNames, in2Pointer); err != nil {
			return err
		}
	case "right#df3ecb3b", "right", "#df3ecb3b":
		if !legacyTypeNames && _tag == "right#df3ecb3b" {
			return internal.ErrorInvalidUnionLegacyTagJSON("Either", "right#df3ecb3b")
		}
		item.index = 1
		var in2Pointer *basictl.JsonLexer
		if _value != nil {
			in2 := basictl.JsonLexer{Data: _value}
			in2Pointer = &in2
		}
		if err := item.valueRight.ReadJSON(legacyTypeNames, in2Pointer); err != nil {
			return err
		}
	default:
		return internal.ErrorInvalidUnionTagJSON("Either", _tag)
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *EitherService6ErrorVectorService6FindResultRow) WriteJSONGeneral(tctx *basictl.JSONWriteContext, w []byte) ([]byte, error) {
	return item.WriteJSONOpt(tctx, w), nil
}

func (item *EitherService6ErrorVectorService6FindResultRow) WriteJSON(w []byte) []byte {
	tctx := basictl.JSONWriteContext{}
	return item.WriteJSONOpt(&tctx, w)
}
func (item *EitherService6ErrorVectorService6FindResultRow) WriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte) []byte {
	switch item.index {
	case 0:
		if tctx.LegacyTypeNames {
			w = append(w, `{"type":"left#0a29cd5d"`...)
		} else {
			w = append(w, `{"type":"left"`...)
		}
		w = append(w, `,"value":`...)
		w = item.valueLeft.WriteJSONOpt(tctx, w)
		return append(w, '}')
	case 1:
		if tctx.LegacyTypeNames {
			w = append(w, `{"type":"right#df3ecb3b"`...)
		} else {
			w = append(w, `{"type":"right"`...)
		}
		w = append(w, `,"value":`...)
		w = item.valueRight.WriteJSONOpt(tctx, w)
		return append(w, '}')
	default: // Impossible due to panic above
		return w
	}
}

func (item EitherService6ErrorVectorService6FindResultRow) String() string {
	return string(item.WriteJSON(nil))
}

func (item *EitherService6ErrorVectorService6FindResultRow) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *EitherService6ErrorVectorService6FindResultRow) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("Either", err.Error())
	}
	return nil
}

func (item LeftService6ErrorVectorService6FindResultRow) AsUnion() EitherService6ErrorVectorService6FindResultRow {
	var ret EitherService6ErrorVectorService6FindResultRow
	ret.SetLeft(item)
	return ret
}

type LeftService6ErrorVectorService6FindResultRow struct {
	Value tlService6Error.Service6Error
}

func (LeftService6ErrorVectorService6FindResultRow) TLName() string { return "left" }
func (LeftService6ErrorVectorService6FindResultRow) TLTag() uint32  { return 0x0a29cd5d }

func (item *LeftService6ErrorVectorService6FindResultRow) Reset() {
	item.Value.Reset()
}

func (item *LeftService6ErrorVectorService6FindResultRow) Read(w []byte) (_ []byte, err error) {
	return item.Value.Read(w)
}

func (item *LeftService6ErrorVectorService6FindResultRow) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *LeftService6ErrorVectorService6FindResultRow) Write(w []byte) []byte {
	w = item.Value.Write(w)
	return w
}

func (item *LeftService6ErrorVectorService6FindResultRow) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x0a29cd5d); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *LeftService6ErrorVectorService6FindResultRow) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *LeftService6ErrorVectorService6FindResultRow) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x0a29cd5d)
	return item.Write(w)
}

func (item LeftService6ErrorVectorService6FindResultRow) String() string {
	return string(item.WriteJSON(nil))
}

func (item *LeftService6ErrorVectorService6FindResultRow) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propValuePresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "value":
				if propValuePresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("left", "value")
				}
				if err := item.Value.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propValuePresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("left", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propValuePresented {
		item.Value.Reset()
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *LeftService6ErrorVectorService6FindResultRow) WriteJSONGeneral(tctx *basictl.JSONWriteContext, w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(tctx, w), nil
}

func (item *LeftService6ErrorVectorService6FindResultRow) WriteJSON(w []byte) []byte {
	tctx := basictl.JSONWriteContext{}
	return item.WriteJSONOpt(&tctx, w)
}
func (item *LeftService6ErrorVectorService6FindResultRow) WriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte) []byte {
	w = append(w, '{')
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"value":`...)
	w = item.Value.WriteJSONOpt(tctx, w)
	return append(w, '}')
}

func (item *LeftService6ErrorVectorService6FindResultRow) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *LeftService6ErrorVectorService6FindResultRow) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("left", err.Error())
	}
	return nil
}

func (item RightService6ErrorVectorService6FindResultRow) AsUnion() EitherService6ErrorVectorService6FindResultRow {
	var ret EitherService6ErrorVectorService6FindResultRow
	ret.SetRight(item)
	return ret
}

type RightService6ErrorVectorService6FindResultRow struct {
	Value []tlService6FindResultRow.Service6FindResultRow
}

func (RightService6ErrorVectorService6FindResultRow) TLName() string { return "right" }
func (RightService6ErrorVectorService6FindResultRow) TLTag() uint32  { return 0xdf3ecb3b }

func (item *RightService6ErrorVectorService6FindResultRow) Reset() {
	item.Value = item.Value[:0]
}

func (item *RightService6ErrorVectorService6FindResultRow) Read(w []byte) (_ []byte, err error) {
	return tlBuiltinVectorService6FindResultRow.BuiltinVectorService6FindResultRowRead(w, &item.Value)
}

func (item *RightService6ErrorVectorService6FindResultRow) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *RightService6ErrorVectorService6FindResultRow) Write(w []byte) []byte {
	w = tlBuiltinVectorService6FindResultRow.BuiltinVectorService6FindResultRowWrite(w, item.Value)
	return w
}

func (item *RightService6ErrorVectorService6FindResultRow) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xdf3ecb3b); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *RightService6ErrorVectorService6FindResultRow) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *RightService6ErrorVectorService6FindResultRow) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0xdf3ecb3b)
	return item.Write(w)
}

func (item RightService6ErrorVectorService6FindResultRow) String() string {
	return string(item.WriteJSON(nil))
}

func (item *RightService6ErrorVectorService6FindResultRow) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propValuePresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "value":
				if propValuePresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("right", "value")
				}
				if err := tlBuiltinVectorService6FindResultRow.BuiltinVectorService6FindResultRowReadJSON(legacyTypeNames, in, &item.Value); err != nil {
					return err
				}
				propValuePresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("right", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propValuePresented {
		item.Value = item.Value[:0]
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *RightService6ErrorVectorService6FindResultRow) WriteJSONGeneral(tctx *basictl.JSONWriteContext, w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(tctx, w), nil
}

func (item *RightService6ErrorVectorService6FindResultRow) WriteJSON(w []byte) []byte {
	tctx := basictl.JSONWriteContext{}
	return item.WriteJSONOpt(&tctx, w)
}
func (item *RightService6ErrorVectorService6FindResultRow) WriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte) []byte {
	w = append(w, '{')
	backupIndexValue := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"value":`...)
	w = tlBuiltinVectorService6FindResultRow.BuiltinVectorService6FindResultRowWriteJSONOpt(tctx, w, item.Value)
	if (len(item.Value) != 0) == false {
		w = w[:backupIndexValue]
	}
	return append(w, '}')
}

func (item *RightService6ErrorVectorService6FindResultRow) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *RightService6ErrorVectorService6FindResultRow) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("right", err.Error())
	}
	return nil
}
