// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlBenchmarksVrutoyTopLevelContainerWithDependency

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tlbenchmarks/tlBenchmarksVrutoyPositions"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type BenchmarksVrutoyTopLevelContainerWithDependency struct {
	N     uint32
	Value tlBenchmarksVrutoyPositions.BenchmarksVrutoyPositions
}

func (BenchmarksVrutoyTopLevelContainerWithDependency) TLName() string {
	return "benchmarks.vrutoyTopLevelContainerWithDependency"
}
func (BenchmarksVrutoyTopLevelContainerWithDependency) TLTag() uint32 { return 0xc176008e }

func (item *BenchmarksVrutoyTopLevelContainerWithDependency) Reset() {
	item.N = 0
	item.Value.Reset()
}

func (item *BenchmarksVrutoyTopLevelContainerWithDependency) FillRandom(rg *basictl.RandGenerator) {
	item.N = basictl.RandomUint(rg)
	item.N = rg.LimitValue(item.N)
	item.Value.FillRandom(rg, item.N)
}

func (item *BenchmarksVrutoyTopLevelContainerWithDependency) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.N); err != nil {
		return w, err
	}
	return item.Value.Read(w, item.N)
}

// This method is general version of Write, use it instead!
func (item *BenchmarksVrutoyTopLevelContainerWithDependency) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w)
}

func (item *BenchmarksVrutoyTopLevelContainerWithDependency) Write(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, item.N)
	if w, err = item.Value.Write(w, item.N); err != nil {
		return w, err
	}
	return w, nil
}

func (item *BenchmarksVrutoyTopLevelContainerWithDependency) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xc176008e); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *BenchmarksVrutoyTopLevelContainerWithDependency) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w)
}

func (item *BenchmarksVrutoyTopLevelContainerWithDependency) WriteBoxed(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, 0xc176008e)
	return item.Write(w)
}

func (item BenchmarksVrutoyTopLevelContainerWithDependency) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func (item *BenchmarksVrutoyTopLevelContainerWithDependency) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propNPresented bool
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
			case "n":
				if propNPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("benchmarks.vrutoyTopLevelContainerWithDependency", "n")
				}
				if err := internal.Json2ReadUint32(in, &item.N); err != nil {
					return err
				}
				propNPresented = true
			case "value":
				if rawValue != nil {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("benchmarks.vrutoyTopLevelContainerWithDependency", "value")
				}
				rawValue = in.Raw()
				if !in.Ok() {
					return in.Error()
				}
			default:
				return internal.ErrorInvalidJSONExcessElement("benchmarks.vrutoyTopLevelContainerWithDependency", key)
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
	var inValuePointer *basictl.JsonLexer
	inValue := basictl.JsonLexer{Data: rawValue}
	if rawValue != nil {
		inValuePointer = &inValue
	}
	if err := item.Value.ReadJSON(legacyTypeNames, inValuePointer, item.N); err != nil {
		return err
	}

	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *BenchmarksVrutoyTopLevelContainerWithDependency) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w)
}

func (item *BenchmarksVrutoyTopLevelContainerWithDependency) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w)
}
func (item *BenchmarksVrutoyTopLevelContainerWithDependency) WriteJSONOpt(newTypeNames bool, short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	backupIndexN := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"n":`...)
	w = basictl.JSONWriteUint32(w, item.N)
	if (item.N != 0) == false {
		w = w[:backupIndexN]
	}
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"value":`...)
	if w, err = item.Value.WriteJSONOpt(newTypeNames, short, w, item.N); err != nil {
		return w, err
	}
	return append(w, '}'), nil
}

func (item *BenchmarksVrutoyTopLevelContainerWithDependency) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *BenchmarksVrutoyTopLevelContainerWithDependency) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("benchmarks.vrutoyTopLevelContainerWithDependency", err.Error())
	}
	return nil
}
