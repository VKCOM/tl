// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlMaybeTest1

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlInnerMaybe"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlInnerMaybe0"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlIntBoxedMaybe"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlIntMaybe"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlTupleInt3BoxedMaybe"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlTupleInt3Maybe"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlTupleIntBoxed0BoxedMaybe"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlTupleIntBoxed3Maybe"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlVectorIntBoxedMaybe"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlVectorIntMaybe"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type MaybeTest1 struct {
	N uint32
	A tlIntMaybe.IntMaybe
	B tlIntBoxedMaybe.IntBoxedMaybe
	C tlTupleInt3Maybe.TupleInt3Maybe
	D tlTupleIntBoxed3Maybe.TupleIntBoxed3Maybe
	E tlInnerMaybe.InnerMaybe
	F tlInnerMaybe0.InnerMaybe0
	G tlVectorIntMaybe.VectorIntMaybe
	H tlVectorIntBoxedMaybe.VectorIntBoxedMaybe
	I tlTupleInt3BoxedMaybe.TupleInt3BoxedMaybe
	J tlTupleIntBoxed0BoxedMaybe.TupleIntBoxed0BoxedMaybe
}

func (MaybeTest1) TLName() string { return "maybeTest1" }
func (MaybeTest1) TLTag() uint32  { return 0xc457763c }

func (item *MaybeTest1) Reset() {
	item.N = 0
	item.A.Reset()
	item.B.Reset()
	item.C.Reset()
	item.D.Reset()
	item.E.Reset()
	item.F.Reset()
	item.G.Reset()
	item.H.Reset()
	item.I.Reset()
	item.J.Reset()
}

func (item *MaybeTest1) FillRandom(rg *basictl.RandGenerator) {
	item.N = basictl.RandomUint(rg)
	item.N = rg.LimitValue(item.N)
	item.A.FillRandom(rg)
	item.B.FillRandom(rg)
	item.C.FillRandom(rg)
	item.D.FillRandom(rg)
	item.E.FillRandom(rg, item.N)
	item.F.FillRandom(rg)
	item.G.FillRandom(rg)
	item.H.FillRandom(rg)
	item.I.FillRandom(rg)
	item.J.FillRandom(rg)
}

func (item *MaybeTest1) Read(w []byte) (_ []byte, err error) {
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
	if w, err = item.D.ReadBoxed(w); err != nil {
		return w, err
	}
	if w, err = item.E.ReadBoxed(w, item.N); err != nil {
		return w, err
	}
	if w, err = item.F.ReadBoxed(w); err != nil {
		return w, err
	}
	if w, err = item.G.ReadBoxed(w); err != nil {
		return w, err
	}
	if w, err = item.H.ReadBoxed(w); err != nil {
		return w, err
	}
	if w, err = item.I.ReadBoxed(w); err != nil {
		return w, err
	}
	return item.J.ReadBoxed(w)
}

// This method is general version of Write, use it instead!
func (item *MaybeTest1) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w)
}

func (item *MaybeTest1) Write(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, item.N)
	w = item.A.WriteBoxed(w)
	w = item.B.WriteBoxed(w)
	w = item.C.WriteBoxed(w)
	w = item.D.WriteBoxed(w)
	if w, err = item.E.WriteBoxed(w, item.N); err != nil {
		return w, err
	}
	w = item.F.WriteBoxed(w)
	w = item.G.WriteBoxed(w)
	w = item.H.WriteBoxed(w)
	w = item.I.WriteBoxed(w)
	w = item.J.WriteBoxed(w)
	return w, nil
}

func (item *MaybeTest1) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xc457763c); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *MaybeTest1) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w)
}

func (item *MaybeTest1) WriteBoxed(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, 0xc457763c)
	return item.Write(w)
}

func (item *MaybeTest1) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func (item *MaybeTest1) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propNPresented bool
	var propAPresented bool
	var propBPresented bool
	var propCPresented bool
	var propDPresented bool
	var rawE []byte
	var propFPresented bool
	var propGPresented bool
	var propHPresented bool
	var propIPresented bool
	var propJPresented bool

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
					return internal.ErrorInvalidJSONWithDuplicatingKeys("maybeTest1", "n")
				}
				if err := internal.Json2ReadUint32(in, &item.N); err != nil {
					return err
				}
				propNPresented = true
			case "a":
				if propAPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("maybeTest1", "a")
				}
				if err := item.A.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propAPresented = true
			case "b":
				if propBPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("maybeTest1", "b")
				}
				if err := item.B.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propBPresented = true
			case "c":
				if propCPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("maybeTest1", "c")
				}
				if err := item.C.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propCPresented = true
			case "d":
				if propDPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("maybeTest1", "d")
				}
				if err := item.D.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propDPresented = true
			case "e":
				if rawE != nil {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("maybeTest1", "e")
				}
				rawE = in.Raw()
				if !in.Ok() {
					return in.Error()
				}
			case "f":
				if propFPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("maybeTest1", "f")
				}
				if err := item.F.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propFPresented = true
			case "g":
				if propGPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("maybeTest1", "g")
				}
				if err := item.G.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propGPresented = true
			case "h":
				if propHPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("maybeTest1", "h")
				}
				if err := item.H.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propHPresented = true
			case "i":
				if propIPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("maybeTest1", "i")
				}
				if err := item.I.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propIPresented = true
			case "j":
				if propJPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("maybeTest1", "j")
				}
				if err := item.J.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propJPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("maybeTest1", key)
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
	if !propFPresented {
		item.F.Reset()
	}
	if !propGPresented {
		item.G.Reset()
	}
	if !propHPresented {
		item.H.Reset()
	}
	if !propIPresented {
		item.I.Reset()
	}
	if !propJPresented {
		item.J.Reset()
	}
	var inEPointer *basictl.JsonLexer
	inE := basictl.JsonLexer{Data: rawE}
	if rawE != nil {
		inEPointer = &inE
	}
	if err := item.E.ReadJSON(legacyTypeNames, inEPointer, item.N); err != nil {
		return err
	}

	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *MaybeTest1) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w)
}

func (item *MaybeTest1) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w)
}
func (item *MaybeTest1) WriteJSONOpt(newTypeNames bool, short bool, w []byte) (_ []byte, err error) {
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
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"e":`...)
	if w, err = item.E.WriteJSONOpt(newTypeNames, short, w, item.N); err != nil {
		return w, err
	}
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"f":`...)
	w = item.F.WriteJSONOpt(newTypeNames, short, w)
	backupIndexG := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"g":`...)
	w = item.G.WriteJSONOpt(newTypeNames, short, w)
	if (item.G.Ok) == false {
		w = w[:backupIndexG]
	}
	backupIndexH := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"h":`...)
	w = item.H.WriteJSONOpt(newTypeNames, short, w)
	if (item.H.Ok) == false {
		w = w[:backupIndexH]
	}
	backupIndexI := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"i":`...)
	w = item.I.WriteJSONOpt(newTypeNames, short, w)
	if (item.I.Ok) == false {
		w = w[:backupIndexI]
	}
	backupIndexJ := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"j":`...)
	w = item.J.WriteJSONOpt(newTypeNames, short, w)
	if (item.J.Ok) == false {
		w = w[:backupIndexJ]
	}
	return append(w, '}'), nil
}

func (item *MaybeTest1) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *MaybeTest1) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("maybeTest1", err.Error())
	}
	return nil
}
