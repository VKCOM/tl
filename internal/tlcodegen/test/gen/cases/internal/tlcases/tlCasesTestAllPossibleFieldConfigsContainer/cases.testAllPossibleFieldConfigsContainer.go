// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlCasesTestAllPossibleFieldConfigsContainer

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tlcases/tlCasesTestAllPossibleFieldConfigs"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type CasesTestAllPossibleFieldConfigsContainer struct {
	Outer uint32
	Value tlCasesTestAllPossibleFieldConfigs.CasesTestAllPossibleFieldConfigs
}

func (CasesTestAllPossibleFieldConfigsContainer) TLName() string {
	return "cases.testAllPossibleFieldConfigsContainer"
}
func (CasesTestAllPossibleFieldConfigsContainer) TLTag() uint32 { return 0xe3fae936 }

func (item *CasesTestAllPossibleFieldConfigsContainer) Reset() {
	item.Outer = 0
	item.Value.Reset()
}

func (item *CasesTestAllPossibleFieldConfigsContainer) FillRandom(rg *basictl.RandGenerator) {
	var maskOuter uint32
	maskOuter = basictl.RandomUint(rg)
	item.Outer = 0
	if maskOuter&(1<<0) != 0 {
		item.Outer |= (1 << 0)
	}
	if maskOuter&(1<<1) != 0 {
		item.Outer |= (1 << 1)
	}
	if maskOuter&(1<<2) != 0 {
		item.Outer |= (1 << 2)
	}
	if maskOuter&(1<<3) != 0 {
		item.Outer |= (1 << 3)
	}
	item.Value.FillRandom(rg, item.Outer)
}

func (item *CasesTestAllPossibleFieldConfigsContainer) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.Outer); err != nil {
		return w, err
	}
	return item.Value.Read(w, item.Outer)
}

// This method is general version of Write, use it instead!
func (item *CasesTestAllPossibleFieldConfigsContainer) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w)
}

func (item *CasesTestAllPossibleFieldConfigsContainer) Write(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, item.Outer)
	if w, err = item.Value.Write(w, item.Outer); err != nil {
		return w, err
	}
	return w, nil
}

func (item *CasesTestAllPossibleFieldConfigsContainer) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xe3fae936); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *CasesTestAllPossibleFieldConfigsContainer) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w)
}

func (item *CasesTestAllPossibleFieldConfigsContainer) WriteBoxed(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, 0xe3fae936)
	return item.Write(w)
}

func (item CasesTestAllPossibleFieldConfigsContainer) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func (item *CasesTestAllPossibleFieldConfigsContainer) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propOuterPresented bool
	var rawValue []byte

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "outer":
				if propOuterPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("cases.testAllPossibleFieldConfigsContainer", "outer")
				}
				if err := internal.Json2ReadUint32(in, &item.Outer); err != nil {
					return err
				}
				propOuterPresented = true
			case "value":
				if rawValue != nil {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("cases.testAllPossibleFieldConfigsContainer", "value")
				}
				rawValue = in.Raw()
				if !in.Ok() {
					return in.Error()
				}
			default:
				return internal.ErrorInvalidJSONExcessElement("cases.testAllPossibleFieldConfigsContainer", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propOuterPresented {
		item.Outer = 0
	}
	var inValuePointer *basictl.JsonLexer
	inValue := basictl.JsonLexer{Data: rawValue}
	if rawValue != nil {
		inValuePointer = &inValue
	}
	if err := item.Value.ReadJSON(legacyTypeNames, inValuePointer, item.Outer); err != nil {
		return err
	}

	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *CasesTestAllPossibleFieldConfigsContainer) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w)
}

func (item *CasesTestAllPossibleFieldConfigsContainer) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w)
}
func (item *CasesTestAllPossibleFieldConfigsContainer) WriteJSONOpt(newTypeNames bool, short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	backupIndexOuter := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"outer":`...)
	w = basictl.JSONWriteUint32(w, item.Outer)
	if (item.Outer != 0) == false {
		w = w[:backupIndexOuter]
	}
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"value":`...)
	if w, err = item.Value.WriteJSONOpt(newTypeNames, short, w, item.Outer); err != nil {
		return w, err
	}
	return append(w, '}'), nil
}

func (item *CasesTestAllPossibleFieldConfigsContainer) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *CasesTestAllPossibleFieldConfigsContainer) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("cases.testAllPossibleFieldConfigsContainer", err.Error())
	}
	return nil
}
