// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlBoolStat

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type BoolStat struct {
	StatTrue    int32
	StatFalse   int32
	StatUnknown int32
}

func (BoolStat) TLName() string { return "boolStat" }
func (BoolStat) TLTag() uint32  { return 0x92cbcbfa }

func (item *BoolStat) Reset() {
	item.StatTrue = 0
	item.StatFalse = 0
	item.StatUnknown = 0
}

func (item *BoolStat) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.IntRead(w, &item.StatTrue); err != nil {
		return w, err
	}
	if w, err = basictl.IntRead(w, &item.StatFalse); err != nil {
		return w, err
	}
	return basictl.IntRead(w, &item.StatUnknown)
}

// This method is general version of Write, use it instead!
func (item *BoolStat) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *BoolStat) Write(w []byte) []byte {
	w = basictl.IntWrite(w, item.StatTrue)
	w = basictl.IntWrite(w, item.StatFalse)
	w = basictl.IntWrite(w, item.StatUnknown)
	return w
}

func (item *BoolStat) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x92cbcbfa); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *BoolStat) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *BoolStat) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x92cbcbfa)
	return item.Write(w)
}

func (item *BoolStat) String() string {
	return string(item.WriteJSON(nil))
}

func (item *BoolStat) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propStatTruePresented bool
	var propStatFalsePresented bool
	var propStatUnknownPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "statTrue":
				if propStatTruePresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("boolStat", "statTrue")
				}
				if err := internal.Json2ReadInt32(in, &item.StatTrue); err != nil {
					return err
				}
				propStatTruePresented = true
			case "statFalse":
				if propStatFalsePresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("boolStat", "statFalse")
				}
				if err := internal.Json2ReadInt32(in, &item.StatFalse); err != nil {
					return err
				}
				propStatFalsePresented = true
			case "statUnknown":
				if propStatUnknownPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("boolStat", "statUnknown")
				}
				if err := internal.Json2ReadInt32(in, &item.StatUnknown); err != nil {
					return err
				}
				propStatUnknownPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("boolStat", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propStatTruePresented {
		item.StatTrue = 0
	}
	if !propStatFalsePresented {
		item.StatFalse = 0
	}
	if !propStatUnknownPresented {
		item.StatUnknown = 0
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *BoolStat) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *BoolStat) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *BoolStat) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexStatTrue := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"statTrue":`...)
	w = basictl.JSONWriteInt32(w, item.StatTrue)
	if (item.StatTrue != 0) == false {
		w = w[:backupIndexStatTrue]
	}
	backupIndexStatFalse := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"statFalse":`...)
	w = basictl.JSONWriteInt32(w, item.StatFalse)
	if (item.StatFalse != 0) == false {
		w = w[:backupIndexStatFalse]
	}
	backupIndexStatUnknown := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"statUnknown":`...)
	w = basictl.JSONWriteInt32(w, item.StatUnknown)
	if (item.StatUnknown != 0) == false {
		w = w[:backupIndexStatUnknown]
	}
	return append(w, '}')
}

func (item *BoolStat) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *BoolStat) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("boolStat", err.Error())
	}
	return nil
}
