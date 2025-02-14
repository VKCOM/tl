// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlBuiltinTupleBenchmarksVruPosition

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tlbenchmarks/tlBenchmarksVruPosition"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

func BuiltinTupleBenchmarksVruPositionFillRandom(rg *basictl.RandGenerator, vec *[]tlBenchmarksVruPosition.BenchmarksVruPosition, nat_n uint32) {
	rg.IncreaseDepth()
	*vec = make([]tlBenchmarksVruPosition.BenchmarksVruPosition, nat_n)
	for i := range *vec {
		(*vec)[i].FillRandom(rg)
	}
	rg.DecreaseDepth()
}

func BuiltinTupleBenchmarksVruPositionRead(w []byte, vec *[]tlBenchmarksVruPosition.BenchmarksVruPosition, nat_n uint32) (_ []byte, err error) {
	if uint32(cap(*vec)) < nat_n {
		*vec = make([]tlBenchmarksVruPosition.BenchmarksVruPosition, nat_n)
	} else {
		*vec = (*vec)[:nat_n]
	}
	for i := range *vec {
		if w, err = (*vec)[i].Read(w); err != nil {
			return w, err
		}
	}
	return w, nil
}

func BuiltinTupleBenchmarksVruPositionWrite(w []byte, vec []tlBenchmarksVruPosition.BenchmarksVruPosition, nat_n uint32) (_ []byte, err error) {
	if uint32(len(vec)) != nat_n {
		return w, internal.ErrorWrongSequenceLength("[]tlBenchmarksVruPosition.BenchmarksVruPosition", len(vec), nat_n)
	}
	for _, elem := range vec {
		w = elem.Write(w)
	}
	return w, nil
}

func BuiltinTupleBenchmarksVruPositionReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, vec *[]tlBenchmarksVruPosition.BenchmarksVruPosition, nat_n uint32) error {
	if uint32(cap(*vec)) < nat_n {
		*vec = make([]tlBenchmarksVruPosition.BenchmarksVruPosition, nat_n)
	} else {
		*vec = (*vec)[:nat_n]
	}
	index := 0
	if in != nil {
		in.Delim('[')
		if !in.Ok() {
			return internal.ErrorInvalidJSON("[]tlBenchmarksVruPosition.BenchmarksVruPosition", "expected json array")
		}
		for ; !in.IsDelim(']'); index++ {
			if nat_n <= uint32(index) {
				return internal.ErrorInvalidJSON("[]tlBenchmarksVruPosition.BenchmarksVruPosition", "array is longer than expected")
			}
			if err := (*vec)[index].ReadJSON(legacyTypeNames, in); err != nil {
				return err
			}
			in.WantComma()
		}
		in.Delim(']')
		if !in.Ok() {
			return internal.ErrorInvalidJSON("[]tlBenchmarksVruPosition.BenchmarksVruPosition", "expected json array's end")
		}
	}
	if uint32(index) != nat_n {
		return internal.ErrorWrongSequenceLength("[]tlBenchmarksVruPosition.BenchmarksVruPosition", index, nat_n)
	}
	return nil
}

func BuiltinTupleBenchmarksVruPositionWriteJSON(w []byte, vec []tlBenchmarksVruPosition.BenchmarksVruPosition, nat_n uint32) (_ []byte, err error) {
	return BuiltinTupleBenchmarksVruPositionWriteJSONOpt(true, false, w, vec, nat_n)
}
func BuiltinTupleBenchmarksVruPositionWriteJSONOpt(newTypeNames bool, short bool, w []byte, vec []tlBenchmarksVruPosition.BenchmarksVruPosition, nat_n uint32) (_ []byte, err error) {
	if uint32(len(vec)) != nat_n {
		return w, internal.ErrorWrongSequenceLength("[]tlBenchmarksVruPosition.BenchmarksVruPosition", len(vec), nat_n)
	}
	w = append(w, '[')
	for _, elem := range vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = elem.WriteJSONOpt(newTypeNames, short, w)
	}
	return append(w, ']'), nil
}
