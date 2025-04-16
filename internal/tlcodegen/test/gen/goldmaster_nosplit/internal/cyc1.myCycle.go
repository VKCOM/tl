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

func BuiltinVectorCyc1MyCycleFillRandom(rg *basictl.RandGenerator, vec *[]Cyc1MyCycle) {
	rg.IncreaseDepth()
	l := rg.LimitValue(basictl.RandomUint(rg))
	*vec = make([]Cyc1MyCycle, l)
	for i := range *vec {
		(*vec)[i].FillRandom(rg)
	}
	rg.DecreaseDepth()
}
func BuiltinVectorCyc1MyCycleRead(w []byte, vec *[]Cyc1MyCycle) (_ []byte, err error) {
	var l uint32
	if w, err = basictl.NatRead(w, &l); err != nil {
		return w, err
	}
	if err = basictl.CheckLengthSanity(w, l, 4); err != nil {
		return w, err
	}
	if uint32(cap(*vec)) < l {
		*vec = make([]Cyc1MyCycle, l)
	} else {
		*vec = (*vec)[:l]
	}
	for i := range *vec {
		if w, err = (*vec)[i].Read(w); err != nil {
			return w, err
		}
	}
	return w, nil
}

func BuiltinVectorCyc1MyCycleWrite(w []byte, vec []Cyc1MyCycle) []byte {
	w = basictl.NatWrite(w, uint32(len(vec)))
	for _, elem := range vec {
		w = elem.Write(w)
	}
	return w
}

func BuiltinVectorCyc1MyCycleReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, vec *[]Cyc1MyCycle) error {
	*vec = (*vec)[:cap(*vec)]
	index := 0
	if in != nil {
		in.Delim('[')
		if !in.Ok() {
			return ErrorInvalidJSON("[]Cyc1MyCycle", "expected json array")
		}
		for ; !in.IsDelim(']'); index++ {
			if len(*vec) <= index {
				var newValue Cyc1MyCycle
				*vec = append(*vec, newValue)
				*vec = (*vec)[:cap(*vec)]
			}
			if err := (*vec)[index].ReadJSON(legacyTypeNames, in); err != nil {
				return err
			}
			in.WantComma()
		}
		in.Delim(']')
		if !in.Ok() {
			return ErrorInvalidJSON("[]Cyc1MyCycle", "expected json array's end")
		}
	}
	*vec = (*vec)[:index]
	return nil
}

func BuiltinVectorCyc1MyCycleWriteJSON(w []byte, vec []Cyc1MyCycle) []byte {
	return BuiltinVectorCyc1MyCycleWriteJSONOpt(true, false, w, vec)
}
func BuiltinVectorCyc1MyCycleWriteJSONOpt(newTypeNames bool, short bool, w []byte, vec []Cyc1MyCycle) []byte {
	w = append(w, '[')
	for _, elem := range vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = elem.WriteJSONOpt(newTypeNames, short, w)
	}
	return append(w, ']')
}

type Cyc1MyCycle struct {
	FieldsMask uint32
	A          Cyc2MyCycle // Conditional: item.FieldsMask.0
}

func (Cyc1MyCycle) TLName() string { return "cyc1.myCycle" }
func (Cyc1MyCycle) TLTag() uint32  { return 0x136ecc9e }

func (item *Cyc1MyCycle) SetA(v Cyc2MyCycle) {
	item.A = v
	item.FieldsMask |= 1 << 0
}
func (item *Cyc1MyCycle) ClearA() {
	item.A.Reset()
	item.FieldsMask &^= 1 << 0
}
func (item *Cyc1MyCycle) IsSetA() bool { return item.FieldsMask&(1<<0) != 0 }

func (item *Cyc1MyCycle) Reset() {
	item.FieldsMask = 0
	item.A.Reset()
}

func (item *Cyc1MyCycle) FillRandom(rg *basictl.RandGenerator) {
	var maskFieldsMask uint32
	maskFieldsMask = basictl.RandomUint(rg)
	item.FieldsMask = 0
	if maskFieldsMask&(1<<0) != 0 {
		item.FieldsMask |= (1 << 0)
	}
	if item.FieldsMask&(1<<0) != 0 {
		item.A.FillRandom(rg)
	} else {
		item.A.Reset()
	}
}

func (item *Cyc1MyCycle) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.FieldsMask); err != nil {
		return w, err
	}
	if item.FieldsMask&(1<<0) != 0 {
		if w, err = item.A.Read(w); err != nil {
			return w, err
		}
	} else {
		item.A.Reset()
	}
	return w, nil
}

// This method is general version of Write, use it instead!
func (item *Cyc1MyCycle) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *Cyc1MyCycle) Write(w []byte) []byte {
	w = basictl.NatWrite(w, item.FieldsMask)
	if item.FieldsMask&(1<<0) != 0 {
		w = item.A.Write(w)
	}
	return w
}

func (item *Cyc1MyCycle) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x136ecc9e); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *Cyc1MyCycle) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *Cyc1MyCycle) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x136ecc9e)
	return item.Write(w)
}

func (item Cyc1MyCycle) String() string {
	return string(item.WriteJSON(nil))
}

func (item *Cyc1MyCycle) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propFieldsMaskPresented bool
	var propAPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "fields_mask":
				if propFieldsMaskPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("cyc1.myCycle", "fields_mask")
				}
				if err := Json2ReadUint32(in, &item.FieldsMask); err != nil {
					return err
				}
				propFieldsMaskPresented = true
			case "a":
				if propAPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("cyc1.myCycle", "a")
				}
				if err := item.A.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propAPresented = true
			default:
				return ErrorInvalidJSONExcessElement("cyc1.myCycle", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propFieldsMaskPresented {
		item.FieldsMask = 0
	}
	if !propAPresented {
		item.A.Reset()
	}
	if propAPresented {
		item.FieldsMask |= 1 << 0
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *Cyc1MyCycle) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *Cyc1MyCycle) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *Cyc1MyCycle) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexFieldsMask := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"fields_mask":`...)
	w = basictl.JSONWriteUint32(w, item.FieldsMask)
	if (item.FieldsMask != 0) == false {
		w = w[:backupIndexFieldsMask]
	}
	if item.FieldsMask&(1<<0) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"a":`...)
		w = item.A.WriteJSONOpt(newTypeNames, short, w)
	}
	return append(w, '}')
}

func (item *Cyc1MyCycle) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *Cyc1MyCycle) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("cyc1.myCycle", err.Error())
	}
	return nil
}
