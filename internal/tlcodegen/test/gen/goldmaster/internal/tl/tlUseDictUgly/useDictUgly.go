// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlUseDictUgly

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlBuiltinVectorDictionaryElemIntPairIntInt"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlBuiltinVectorDictionaryElemLongPairIntInt"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlBuiltinVectorDictionaryElemPairBoolAColorInt"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlBuiltinVectorDictionaryElemPairFloatDoubleInt"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlBuiltinVectorDictionaryElemPairIntIntInt"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlBuiltinVectorDictionaryElemPairIntPairMultiPointStringInt"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlBuiltinVectorDictionaryElemStrangeString"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlBuiltinVectorDictionaryElemStringPairIntInt"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlBuiltinVectorDictionaryElemTupleStringInt"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlBuiltinVectorDictionaryElemUglyIntString"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlDictionaryElemPairBoolAColorInt"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlDictionaryElemPairFloatDoubleInt"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlDictionaryElemPairIntIntInt"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlDictionaryElemPairIntPairMultiPointStringInt"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlDictionaryElemTupleStringInt"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlDictionaryElemUglyIntString"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlPairIntInt"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type UseDictUgly struct {
	N uint32
	A []tlDictionaryElemUglyIntString.DictionaryElemUglyIntString
	B map[uint32]string
	C []tlDictionaryElemPairIntIntInt.DictionaryElemPairIntIntInt
	D []tlDictionaryElemTupleStringInt.DictionaryElemTupleStringInt
	E []tlDictionaryElemPairBoolAColorInt.DictionaryElemPairBoolAColorInt
	F []tlDictionaryElemPairFloatDoubleInt.DictionaryElemPairFloatDoubleInt
	G []tlDictionaryElemPairIntPairMultiPointStringInt.DictionaryElemPairIntPairMultiPointStringInt
	X map[int32]tlPairIntInt.PairIntInt
	Y map[int64]tlPairIntInt.PairIntInt
	Z map[string]tlPairIntInt.PairIntInt
}

func (UseDictUgly) TLName() string { return "useDictUgly" }
func (UseDictUgly) TLTag() uint32  { return 0xfb9ce817 }

func (item *UseDictUgly) Reset() {
	item.N = 0
	item.A = item.A[:0]
	tlBuiltinVectorDictionaryElemStrangeString.BuiltinVectorDictionaryElemStrangeStringReset(item.B)
	item.C = item.C[:0]
	item.D = item.D[:0]
	item.E = item.E[:0]
	item.F = item.F[:0]
	item.G = item.G[:0]
	tlBuiltinVectorDictionaryElemIntPairIntInt.BuiltinVectorDictionaryElemIntPairIntIntReset(item.X)
	tlBuiltinVectorDictionaryElemLongPairIntInt.BuiltinVectorDictionaryElemLongPairIntIntReset(item.Y)
	tlBuiltinVectorDictionaryElemStringPairIntInt.BuiltinVectorDictionaryElemStringPairIntIntReset(item.Z)
}

func (item *UseDictUgly) FillRandom(rg *basictl.RandGenerator) {
	var maskN uint32
	maskN = basictl.RandomUint(rg)
	maskN = rg.LimitValue(maskN)
	item.N = 0
	if maskN&(1<<0) != 0 {
		item.N |= (1 << 0)
	}
	if maskN&(1<<1) != 0 {
		item.N |= (1 << 1)
	}
	tlBuiltinVectorDictionaryElemUglyIntString.BuiltinVectorDictionaryElemUglyIntStringFillRandom(rg, &item.A, item.N)
	tlBuiltinVectorDictionaryElemStrangeString.BuiltinVectorDictionaryElemStrangeStringFillRandom(rg, &item.B)
	tlBuiltinVectorDictionaryElemPairIntIntInt.BuiltinVectorDictionaryElemPairIntIntIntFillRandom(rg, &item.C)
	tlBuiltinVectorDictionaryElemTupleStringInt.BuiltinVectorDictionaryElemTupleStringIntFillRandom(rg, &item.D, item.N)
	tlBuiltinVectorDictionaryElemPairBoolAColorInt.BuiltinVectorDictionaryElemPairBoolAColorIntFillRandom(rg, &item.E)
	tlBuiltinVectorDictionaryElemPairFloatDoubleInt.BuiltinVectorDictionaryElemPairFloatDoubleIntFillRandom(rg, &item.F)
	tlBuiltinVectorDictionaryElemPairIntPairMultiPointStringInt.BuiltinVectorDictionaryElemPairIntPairMultiPointStringIntFillRandom(rg, &item.G)
	tlBuiltinVectorDictionaryElemIntPairIntInt.BuiltinVectorDictionaryElemIntPairIntIntFillRandom(rg, &item.X)
	tlBuiltinVectorDictionaryElemLongPairIntInt.BuiltinVectorDictionaryElemLongPairIntIntFillRandom(rg, &item.Y)
	tlBuiltinVectorDictionaryElemStringPairIntInt.BuiltinVectorDictionaryElemStringPairIntIntFillRandom(rg, &item.Z)
}

func (item *UseDictUgly) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.N); err != nil {
		return w, err
	}
	if w, err = tlBuiltinVectorDictionaryElemUglyIntString.BuiltinVectorDictionaryElemUglyIntStringRead(w, &item.A, item.N); err != nil {
		return w, err
	}
	if w, err = tlBuiltinVectorDictionaryElemStrangeString.BuiltinVectorDictionaryElemStrangeStringRead(w, &item.B); err != nil {
		return w, err
	}
	if w, err = tlBuiltinVectorDictionaryElemPairIntIntInt.BuiltinVectorDictionaryElemPairIntIntIntRead(w, &item.C); err != nil {
		return w, err
	}
	if w, err = tlBuiltinVectorDictionaryElemTupleStringInt.BuiltinVectorDictionaryElemTupleStringIntRead(w, &item.D, item.N); err != nil {
		return w, err
	}
	if w, err = tlBuiltinVectorDictionaryElemPairBoolAColorInt.BuiltinVectorDictionaryElemPairBoolAColorIntRead(w, &item.E); err != nil {
		return w, err
	}
	if w, err = tlBuiltinVectorDictionaryElemPairFloatDoubleInt.BuiltinVectorDictionaryElemPairFloatDoubleIntRead(w, &item.F); err != nil {
		return w, err
	}
	if w, err = tlBuiltinVectorDictionaryElemPairIntPairMultiPointStringInt.BuiltinVectorDictionaryElemPairIntPairMultiPointStringIntRead(w, &item.G); err != nil {
		return w, err
	}
	if w, err = tlBuiltinVectorDictionaryElemIntPairIntInt.BuiltinVectorDictionaryElemIntPairIntIntRead(w, &item.X); err != nil {
		return w, err
	}
	if w, err = tlBuiltinVectorDictionaryElemLongPairIntInt.BuiltinVectorDictionaryElemLongPairIntIntRead(w, &item.Y); err != nil {
		return w, err
	}
	return tlBuiltinVectorDictionaryElemStringPairIntInt.BuiltinVectorDictionaryElemStringPairIntIntRead(w, &item.Z)
}

// This method is general version of Write, use it instead!
func (item *UseDictUgly) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w)
}

func (item *UseDictUgly) Write(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, item.N)
	w = tlBuiltinVectorDictionaryElemUglyIntString.BuiltinVectorDictionaryElemUglyIntStringWrite(w, item.A, item.N)
	w = tlBuiltinVectorDictionaryElemStrangeString.BuiltinVectorDictionaryElemStrangeStringWrite(w, item.B)
	w = tlBuiltinVectorDictionaryElemPairIntIntInt.BuiltinVectorDictionaryElemPairIntIntIntWrite(w, item.C)
	if w, err = tlBuiltinVectorDictionaryElemTupleStringInt.BuiltinVectorDictionaryElemTupleStringIntWrite(w, item.D, item.N); err != nil {
		return w, err
	}
	w = tlBuiltinVectorDictionaryElemPairBoolAColorInt.BuiltinVectorDictionaryElemPairBoolAColorIntWrite(w, item.E)
	w = tlBuiltinVectorDictionaryElemPairFloatDoubleInt.BuiltinVectorDictionaryElemPairFloatDoubleIntWrite(w, item.F)
	w = tlBuiltinVectorDictionaryElemPairIntPairMultiPointStringInt.BuiltinVectorDictionaryElemPairIntPairMultiPointStringIntWrite(w, item.G)
	w = tlBuiltinVectorDictionaryElemIntPairIntInt.BuiltinVectorDictionaryElemIntPairIntIntWrite(w, item.X)
	w = tlBuiltinVectorDictionaryElemLongPairIntInt.BuiltinVectorDictionaryElemLongPairIntIntWrite(w, item.Y)
	w = tlBuiltinVectorDictionaryElemStringPairIntInt.BuiltinVectorDictionaryElemStringPairIntIntWrite(w, item.Z)
	return w, nil
}

func (item *UseDictUgly) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xfb9ce817); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *UseDictUgly) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w)
}

func (item *UseDictUgly) WriteBoxed(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, 0xfb9ce817)
	return item.Write(w)
}

func (item UseDictUgly) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func (item *UseDictUgly) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propNPresented bool
	var rawA []byte
	var propBPresented bool
	var propCPresented bool
	var rawD []byte
	var propEPresented bool
	var propFPresented bool
	var propGPresented bool
	var propXPresented bool
	var propYPresented bool
	var propZPresented bool

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
					return internal.ErrorInvalidJSONWithDuplicatingKeys("useDictUgly", "n")
				}
				if err := internal.Json2ReadUint32(in, &item.N); err != nil {
					return err
				}
				propNPresented = true
			case "a":
				if rawA != nil {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("useDictUgly", "a")
				}
				rawA = in.Raw()
				if !in.Ok() {
					return in.Error()
				}
			case "b":
				if propBPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("useDictUgly", "b")
				}
				if err := tlBuiltinVectorDictionaryElemStrangeString.BuiltinVectorDictionaryElemStrangeStringReadJSON(legacyTypeNames, in, &item.B); err != nil {
					return err
				}
				propBPresented = true
			case "c":
				if propCPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("useDictUgly", "c")
				}
				if err := tlBuiltinVectorDictionaryElemPairIntIntInt.BuiltinVectorDictionaryElemPairIntIntIntReadJSON(legacyTypeNames, in, &item.C); err != nil {
					return err
				}
				propCPresented = true
			case "d":
				if rawD != nil {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("useDictUgly", "d")
				}
				rawD = in.Raw()
				if !in.Ok() {
					return in.Error()
				}
			case "e":
				if propEPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("useDictUgly", "e")
				}
				if err := tlBuiltinVectorDictionaryElemPairBoolAColorInt.BuiltinVectorDictionaryElemPairBoolAColorIntReadJSON(legacyTypeNames, in, &item.E); err != nil {
					return err
				}
				propEPresented = true
			case "f":
				if propFPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("useDictUgly", "f")
				}
				if err := tlBuiltinVectorDictionaryElemPairFloatDoubleInt.BuiltinVectorDictionaryElemPairFloatDoubleIntReadJSON(legacyTypeNames, in, &item.F); err != nil {
					return err
				}
				propFPresented = true
			case "g":
				if propGPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("useDictUgly", "g")
				}
				if err := tlBuiltinVectorDictionaryElemPairIntPairMultiPointStringInt.BuiltinVectorDictionaryElemPairIntPairMultiPointStringIntReadJSON(legacyTypeNames, in, &item.G); err != nil {
					return err
				}
				propGPresented = true
			case "x":
				if propXPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("useDictUgly", "x")
				}
				if err := tlBuiltinVectorDictionaryElemIntPairIntInt.BuiltinVectorDictionaryElemIntPairIntIntReadJSON(legacyTypeNames, in, &item.X); err != nil {
					return err
				}
				propXPresented = true
			case "y":
				if propYPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("useDictUgly", "y")
				}
				if err := tlBuiltinVectorDictionaryElemLongPairIntInt.BuiltinVectorDictionaryElemLongPairIntIntReadJSON(legacyTypeNames, in, &item.Y); err != nil {
					return err
				}
				propYPresented = true
			case "z":
				if propZPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("useDictUgly", "z")
				}
				if err := tlBuiltinVectorDictionaryElemStringPairIntInt.BuiltinVectorDictionaryElemStringPairIntIntReadJSON(legacyTypeNames, in, &item.Z); err != nil {
					return err
				}
				propZPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("useDictUgly", key)
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
	if !propBPresented {
		tlBuiltinVectorDictionaryElemStrangeString.BuiltinVectorDictionaryElemStrangeStringReset(item.B)
	}
	if !propCPresented {
		item.C = item.C[:0]
	}
	if !propEPresented {
		item.E = item.E[:0]
	}
	if !propFPresented {
		item.F = item.F[:0]
	}
	if !propGPresented {
		item.G = item.G[:0]
	}
	if !propXPresented {
		tlBuiltinVectorDictionaryElemIntPairIntInt.BuiltinVectorDictionaryElemIntPairIntIntReset(item.X)
	}
	if !propYPresented {
		tlBuiltinVectorDictionaryElemLongPairIntInt.BuiltinVectorDictionaryElemLongPairIntIntReset(item.Y)
	}
	if !propZPresented {
		tlBuiltinVectorDictionaryElemStringPairIntInt.BuiltinVectorDictionaryElemStringPairIntIntReset(item.Z)
	}
	var inAPointer *basictl.JsonLexer
	inA := basictl.JsonLexer{Data: rawA}
	if rawA != nil {
		inAPointer = &inA
	}
	if err := tlBuiltinVectorDictionaryElemUglyIntString.BuiltinVectorDictionaryElemUglyIntStringReadJSON(legacyTypeNames, inAPointer, &item.A, item.N); err != nil {
		return err
	}

	var inDPointer *basictl.JsonLexer
	inD := basictl.JsonLexer{Data: rawD}
	if rawD != nil {
		inDPointer = &inD
	}
	if err := tlBuiltinVectorDictionaryElemTupleStringInt.BuiltinVectorDictionaryElemTupleStringIntReadJSON(legacyTypeNames, inDPointer, &item.D, item.N); err != nil {
		return err
	}

	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *UseDictUgly) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w)
}

func (item *UseDictUgly) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w)
}
func (item *UseDictUgly) WriteJSONOpt(newTypeNames bool, short bool, w []byte) (_ []byte, err error) {
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
	w = tlBuiltinVectorDictionaryElemUglyIntString.BuiltinVectorDictionaryElemUglyIntStringWriteJSONOpt(newTypeNames, short, w, item.A, item.N)
	if (len(item.A) != 0) == false {
		w = w[:backupIndexA]
	}
	backupIndexB := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"b":`...)
	w = tlBuiltinVectorDictionaryElemStrangeString.BuiltinVectorDictionaryElemStrangeStringWriteJSONOpt(newTypeNames, short, w, item.B)
	if (len(item.B) != 0) == false {
		w = w[:backupIndexB]
	}
	backupIndexC := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"c":`...)
	w = tlBuiltinVectorDictionaryElemPairIntIntInt.BuiltinVectorDictionaryElemPairIntIntIntWriteJSONOpt(newTypeNames, short, w, item.C)
	if (len(item.C) != 0) == false {
		w = w[:backupIndexC]
	}
	backupIndexD := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"d":`...)
	if w, err = tlBuiltinVectorDictionaryElemTupleStringInt.BuiltinVectorDictionaryElemTupleStringIntWriteJSONOpt(newTypeNames, short, w, item.D, item.N); err != nil {
		return w, err
	}
	if (len(item.D) != 0) == false {
		w = w[:backupIndexD]
	}
	backupIndexE := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"e":`...)
	w = tlBuiltinVectorDictionaryElemPairBoolAColorInt.BuiltinVectorDictionaryElemPairBoolAColorIntWriteJSONOpt(newTypeNames, short, w, item.E)
	if (len(item.E) != 0) == false {
		w = w[:backupIndexE]
	}
	backupIndexF := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"f":`...)
	w = tlBuiltinVectorDictionaryElemPairFloatDoubleInt.BuiltinVectorDictionaryElemPairFloatDoubleIntWriteJSONOpt(newTypeNames, short, w, item.F)
	if (len(item.F) != 0) == false {
		w = w[:backupIndexF]
	}
	backupIndexG := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"g":`...)
	w = tlBuiltinVectorDictionaryElemPairIntPairMultiPointStringInt.BuiltinVectorDictionaryElemPairIntPairMultiPointStringIntWriteJSONOpt(newTypeNames, short, w, item.G)
	if (len(item.G) != 0) == false {
		w = w[:backupIndexG]
	}
	backupIndexX := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"x":`...)
	w = tlBuiltinVectorDictionaryElemIntPairIntInt.BuiltinVectorDictionaryElemIntPairIntIntWriteJSONOpt(newTypeNames, short, w, item.X)
	if (len(item.X) != 0) == false {
		w = w[:backupIndexX]
	}
	backupIndexY := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"y":`...)
	w = tlBuiltinVectorDictionaryElemLongPairIntInt.BuiltinVectorDictionaryElemLongPairIntIntWriteJSONOpt(newTypeNames, short, w, item.Y)
	if (len(item.Y) != 0) == false {
		w = w[:backupIndexY]
	}
	backupIndexZ := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"z":`...)
	w = tlBuiltinVectorDictionaryElemStringPairIntInt.BuiltinVectorDictionaryElemStringPairIntIntWriteJSONOpt(newTypeNames, short, w, item.Z)
	if (len(item.Z) != 0) == false {
		w = w[:backupIndexZ]
	}
	return append(w, '}'), nil
}

func (item *UseDictUgly) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *UseDictUgly) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("useDictUgly", err.Error())
	}
	return nil
}
