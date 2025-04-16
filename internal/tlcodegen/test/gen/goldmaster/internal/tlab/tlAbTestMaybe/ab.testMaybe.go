// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlAbTestMaybe

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlAbMyTypeBoxedMaybe"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlAbMyTypeMaybe"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlCdMyTypeMaybe"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlIntMaybe"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type AbTestMaybe struct {
	N uint32
	A tlIntMaybe.IntMaybe
	B tlAbMyTypeMaybe.AbMyTypeMaybe
	C tlCdMyTypeMaybe.CdMyTypeMaybe
	D tlAbMyTypeBoxedMaybe.AbMyTypeBoxedMaybe
}

func (AbTestMaybe) TLName() string { return "ab.testMaybe" }
func (AbTestMaybe) TLTag() uint32  { return 0x4dac492a }

func (item *AbTestMaybe) Reset() {
	item.N = 0
	item.A.Reset()
	item.B.Reset()
	item.C.Reset()
	item.D.Reset()
}

func (item *AbTestMaybe) FillRandom(rg *basictl.RandGenerator) {
	item.N = basictl.RandomUint(rg)
	item.A.FillRandom(rg)
	item.B.FillRandom(rg)
	item.C.FillRandom(rg)
	item.D.FillRandom(rg)
}

func (item *AbTestMaybe) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.N); err != nil {
		return w, err
	}
	if w, err = item.A.ReadBoxed(w); err != nil {
		return w, err
	}
	if w, err = item.B.ReadBoxed(w); err != nil {
		return w, err
	}
	if w, err = item.C.ReadBoxed(w); err != nil {
		return w, err
	}
	return item.D.ReadBoxed(w)
}

// This method is general version of Write, use it instead!
func (item *AbTestMaybe) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *AbTestMaybe) Write(w []byte) []byte {
	w = basictl.NatWrite(w, item.N)
	w = item.A.WriteBoxed(w)
	w = item.B.WriteBoxed(w)
	w = item.C.WriteBoxed(w)
	w = item.D.WriteBoxed(w)
	return w
}

func (item *AbTestMaybe) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x4dac492a); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *AbTestMaybe) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *AbTestMaybe) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x4dac492a)
	return item.Write(w)
}

func (item *AbTestMaybe) String() string {
	return string(item.WriteJSON(nil))
}

func (item *AbTestMaybe) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propNPresented bool
	var propAPresented bool
	var propBPresented bool
	var propCPresented bool
	var propDPresented bool

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
					return internal.ErrorInvalidJSONWithDuplicatingKeys("ab.testMaybe", "n")
				}
				if err := internal.Json2ReadUint32(in, &item.N); err != nil {
					return err
				}
				propNPresented = true
			case "a":
				if propAPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("ab.testMaybe", "a")
				}
				if err := item.A.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propAPresented = true
			case "b":
				if propBPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("ab.testMaybe", "b")
				}
				if err := item.B.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propBPresented = true
			case "c":
				if propCPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("ab.testMaybe", "c")
				}
				if err := item.C.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propCPresented = true
			case "d":
				if propDPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("ab.testMaybe", "d")
				}
				if err := item.D.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propDPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("ab.testMaybe", key)
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
	if !propAPresented {
		item.A.Reset()
	}
	if !propBPresented {
		item.B.Reset()
	}
	if !propCPresented {
		item.C.Reset()
	}
	if !propDPresented {
		item.D.Reset()
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *AbTestMaybe) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *AbTestMaybe) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *AbTestMaybe) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexN := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"n":`...)
	w = basictl.JSONWriteUint32(w, item.N)
	if (item.N != 0) == false {
		w = w[:backupIndexN]
	}
	backupIndexA := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"a":`...)
	w = item.A.WriteJSONOpt(newTypeNames, short, w)
	if (item.A.Ok) == false {
		w = w[:backupIndexA]
	}
	backupIndexB := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"b":`...)
	w = item.B.WriteJSONOpt(newTypeNames, short, w)
	if (item.B.Ok) == false {
		w = w[:backupIndexB]
	}
	backupIndexC := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"c":`...)
	w = item.C.WriteJSONOpt(newTypeNames, short, w)
	if (item.C.Ok) == false {
		w = w[:backupIndexC]
	}
	backupIndexD := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"d":`...)
	w = item.D.WriteJSONOpt(newTypeNames, short, w)
	if (item.D.Ok) == false {
		w = w[:backupIndexD]
	}
	return append(w, '}')
}

func (item *AbTestMaybe) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *AbTestMaybe) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("ab.testMaybe", err.Error())
	}
	return nil
}
