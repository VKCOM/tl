// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package internal

import (
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite

type UseDictUgly struct {
	N uint32
	A []DictionaryElemUglyIntString
	B map[uint32]string
	C []DictionaryElemPairIntIntInt
	D []DictionaryElemTupleStringInt
	E []DictionaryElemPairBoolAColorInt
	F []DictionaryElemPairFloatDoubleInt
	G []DictionaryElemPairIntPairMultiPointStringInt
	X map[int32]PairIntInt
	Y map[int64]PairIntInt
	Z map[string]PairIntInt
}

func (UseDictUgly) TLName() string { return "useDictUgly" }
func (UseDictUgly) TLTag() uint32  { return 0xfb9ce817 }

func (item *UseDictUgly) Reset() {
	item.N = 0
	item.A = item.A[:0]
	BuiltinVectorDictionaryElemStrangeStringReset(item.B)
	item.C = item.C[:0]
	item.D = item.D[:0]
	item.E = item.E[:0]
	item.F = item.F[:0]
	item.G = item.G[:0]
	BuiltinVectorDictionaryElemIntPairIntIntReset(item.X)
	BuiltinVectorDictionaryElemLongPairIntIntReset(item.Y)
	BuiltinVectorDictionaryElemStringPairIntIntReset(item.Z)
}

func (item *UseDictUgly) FillRandom(rg *basictl.RandGenerator) {
	item.N = basictl.RandomUint(rg)
	item.N = rg.LimitValue(item.N)
	BuiltinVectorDictionaryElemUglyIntStringFillRandom(rg, &item.A, item.N)
	BuiltinVectorDictionaryElemStrangeStringFillRandom(rg, &item.B)
	BuiltinVectorDictionaryElemPairIntIntIntFillRandom(rg, &item.C)
	BuiltinVectorDictionaryElemTupleStringIntFillRandom(rg, &item.D, item.N)
	BuiltinVectorDictionaryElemPairBoolAColorIntFillRandom(rg, &item.E)
	BuiltinVectorDictionaryElemPairFloatDoubleIntFillRandom(rg, &item.F)
	BuiltinVectorDictionaryElemPairIntPairMultiPointStringIntFillRandom(rg, &item.G)
	BuiltinVectorDictionaryElemIntPairIntIntFillRandom(rg, &item.X)
	BuiltinVectorDictionaryElemLongPairIntIntFillRandom(rg, &item.Y)
	BuiltinVectorDictionaryElemStringPairIntIntFillRandom(rg, &item.Z)
}

func (item *UseDictUgly) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.N); err != nil {
		return w, err
	}
	if w, err = BuiltinVectorDictionaryElemUglyIntStringRead(w, &item.A, item.N); err != nil {
		return w, err
	}
	if w, err = BuiltinVectorDictionaryElemStrangeStringRead(w, &item.B); err != nil {
		return w, err
	}
	if w, err = BuiltinVectorDictionaryElemPairIntIntIntRead(w, &item.C); err != nil {
		return w, err
	}
	if w, err = BuiltinVectorDictionaryElemTupleStringIntRead(w, &item.D, item.N); err != nil {
		return w, err
	}
	if w, err = BuiltinVectorDictionaryElemPairBoolAColorIntRead(w, &item.E); err != nil {
		return w, err
	}
	if w, err = BuiltinVectorDictionaryElemPairFloatDoubleIntRead(w, &item.F); err != nil {
		return w, err
	}
	if w, err = BuiltinVectorDictionaryElemPairIntPairMultiPointStringIntRead(w, &item.G); err != nil {
		return w, err
	}
	if w, err = BuiltinVectorDictionaryElemIntPairIntIntRead(w, &item.X); err != nil {
		return w, err
	}
	if w, err = BuiltinVectorDictionaryElemLongPairIntIntRead(w, &item.Y); err != nil {
		return w, err
	}
	return BuiltinVectorDictionaryElemStringPairIntIntRead(w, &item.Z)
}

// This method is general version of Write, use it instead!
func (item *UseDictUgly) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w)
}

func (item *UseDictUgly) Write(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, item.N)
	w = BuiltinVectorDictionaryElemUglyIntStringWrite(w, item.A, item.N)
	w = BuiltinVectorDictionaryElemStrangeStringWrite(w, item.B)
	w = BuiltinVectorDictionaryElemPairIntIntIntWrite(w, item.C)
	if w, err = BuiltinVectorDictionaryElemTupleStringIntWrite(w, item.D, item.N); err != nil {
		return w, err
	}
	w = BuiltinVectorDictionaryElemPairBoolAColorIntWrite(w, item.E)
	w = BuiltinVectorDictionaryElemPairFloatDoubleIntWrite(w, item.F)
	w = BuiltinVectorDictionaryElemPairIntPairMultiPointStringIntWrite(w, item.G)
	w = BuiltinVectorDictionaryElemIntPairIntIntWrite(w, item.X)
	w = BuiltinVectorDictionaryElemLongPairIntIntWrite(w, item.Y)
	w = BuiltinVectorDictionaryElemStringPairIntIntWrite(w, item.Z)
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
					return ErrorInvalidJSONWithDuplicatingKeys("useDictUgly", "n")
				}
				if err := Json2ReadUint32(in, &item.N); err != nil {
					return err
				}
				propNPresented = true
			case "a":
				if rawA != nil {
					return ErrorInvalidJSONWithDuplicatingKeys("useDictUgly", "a")
				}
				rawA = in.Raw()
				if !in.Ok() {
					return in.Error()
				}
			case "b":
				if propBPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("useDictUgly", "b")
				}
				if err := BuiltinVectorDictionaryElemStrangeStringReadJSON(legacyTypeNames, in, &item.B); err != nil {
					return err
				}
				propBPresented = true
			case "c":
				if propCPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("useDictUgly", "c")
				}
				if err := BuiltinVectorDictionaryElemPairIntIntIntReadJSON(legacyTypeNames, in, &item.C); err != nil {
					return err
				}
				propCPresented = true
			case "d":
				if rawD != nil {
					return ErrorInvalidJSONWithDuplicatingKeys("useDictUgly", "d")
				}
				rawD = in.Raw()
				if !in.Ok() {
					return in.Error()
				}
			case "e":
				if propEPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("useDictUgly", "e")
				}
				if err := BuiltinVectorDictionaryElemPairBoolAColorIntReadJSON(legacyTypeNames, in, &item.E); err != nil {
					return err
				}
				propEPresented = true
			case "f":
				if propFPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("useDictUgly", "f")
				}
				if err := BuiltinVectorDictionaryElemPairFloatDoubleIntReadJSON(legacyTypeNames, in, &item.F); err != nil {
					return err
				}
				propFPresented = true
			case "g":
				if propGPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("useDictUgly", "g")
				}
				if err := BuiltinVectorDictionaryElemPairIntPairMultiPointStringIntReadJSON(legacyTypeNames, in, &item.G); err != nil {
					return err
				}
				propGPresented = true
			case "x":
				if propXPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("useDictUgly", "x")
				}
				if err := BuiltinVectorDictionaryElemIntPairIntIntReadJSON(legacyTypeNames, in, &item.X); err != nil {
					return err
				}
				propXPresented = true
			case "y":
				if propYPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("useDictUgly", "y")
				}
				if err := BuiltinVectorDictionaryElemLongPairIntIntReadJSON(legacyTypeNames, in, &item.Y); err != nil {
					return err
				}
				propYPresented = true
			case "z":
				if propZPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("useDictUgly", "z")
				}
				if err := BuiltinVectorDictionaryElemStringPairIntIntReadJSON(legacyTypeNames, in, &item.Z); err != nil {
					return err
				}
				propZPresented = true
			default:
				return ErrorInvalidJSONExcessElement("useDictUgly", key)
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
		BuiltinVectorDictionaryElemStrangeStringReset(item.B)
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
		BuiltinVectorDictionaryElemIntPairIntIntReset(item.X)
	}
	if !propYPresented {
		BuiltinVectorDictionaryElemLongPairIntIntReset(item.Y)
	}
	if !propZPresented {
		BuiltinVectorDictionaryElemStringPairIntIntReset(item.Z)
	}
	var inAPointer *basictl.JsonLexer
	inA := basictl.JsonLexer{Data: rawA}
	if rawA != nil {
		inAPointer = &inA
	}
	if err := BuiltinVectorDictionaryElemUglyIntStringReadJSON(legacyTypeNames, inAPointer, &item.A, item.N); err != nil {
		return err
	}

	var inDPointer *basictl.JsonLexer
	inD := basictl.JsonLexer{Data: rawD}
	if rawD != nil {
		inDPointer = &inD
	}
	if err := BuiltinVectorDictionaryElemTupleStringIntReadJSON(legacyTypeNames, inDPointer, &item.D, item.N); err != nil {
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
	w = BuiltinVectorDictionaryElemUglyIntStringWriteJSONOpt(newTypeNames, short, w, item.A, item.N)
	if (len(item.A) != 0) == false {
		w = w[:backupIndexA]
	}
	backupIndexB := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"b":`...)
	w = BuiltinVectorDictionaryElemStrangeStringWriteJSONOpt(newTypeNames, short, w, item.B)
	if (len(item.B) != 0) == false {
		w = w[:backupIndexB]
	}
	backupIndexC := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"c":`...)
	w = BuiltinVectorDictionaryElemPairIntIntIntWriteJSONOpt(newTypeNames, short, w, item.C)
	if (len(item.C) != 0) == false {
		w = w[:backupIndexC]
	}
	backupIndexD := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"d":`...)
	if w, err = BuiltinVectorDictionaryElemTupleStringIntWriteJSONOpt(newTypeNames, short, w, item.D, item.N); err != nil {
		return w, err
	}
	if (len(item.D) != 0) == false {
		w = w[:backupIndexD]
	}
	backupIndexE := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"e":`...)
	w = BuiltinVectorDictionaryElemPairBoolAColorIntWriteJSONOpt(newTypeNames, short, w, item.E)
	if (len(item.E) != 0) == false {
		w = w[:backupIndexE]
	}
	backupIndexF := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"f":`...)
	w = BuiltinVectorDictionaryElemPairFloatDoubleIntWriteJSONOpt(newTypeNames, short, w, item.F)
	if (len(item.F) != 0) == false {
		w = w[:backupIndexF]
	}
	backupIndexG := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"g":`...)
	w = BuiltinVectorDictionaryElemPairIntPairMultiPointStringIntWriteJSONOpt(newTypeNames, short, w, item.G)
	if (len(item.G) != 0) == false {
		w = w[:backupIndexG]
	}
	backupIndexX := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"x":`...)
	w = BuiltinVectorDictionaryElemIntPairIntIntWriteJSONOpt(newTypeNames, short, w, item.X)
	if (len(item.X) != 0) == false {
		w = w[:backupIndexX]
	}
	backupIndexY := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"y":`...)
	w = BuiltinVectorDictionaryElemLongPairIntIntWriteJSONOpt(newTypeNames, short, w, item.Y)
	if (len(item.Y) != 0) == false {
		w = w[:backupIndexY]
	}
	backupIndexZ := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"z":`...)
	w = BuiltinVectorDictionaryElemStringPairIntIntWriteJSONOpt(newTypeNames, short, w, item.Z)
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
		return ErrorInvalidJSON("useDictUgly", err.Error())
	}
	return nil
}
