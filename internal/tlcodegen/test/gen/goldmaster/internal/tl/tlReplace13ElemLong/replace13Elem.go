// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlReplace13ElemLong

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlBuiltinTupleInt"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlBuiltinTuplePairBoxedIntLong"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlPairIntLong"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type Replace13ElemLong struct {
	B []tlPairIntLong.PairIntLong // Conditional: nat_n.0
	C []int32                     // Conditional: nat_k.0
}

func (item *Replace13ElemLong) SetB(v []tlPairIntLong.PairIntLong, nat_n *uint32) {
	item.B = v
	if nat_n != nil {
		*nat_n |= 1 << 0
	}
}
func (item *Replace13ElemLong) ClearB(nat_n *uint32) {
	item.B = item.B[:0]
	if nat_n != nil {
		*nat_n &^= 1 << 0
	}
}
func (item *Replace13ElemLong) IsSetB(nat_n uint32) bool { return nat_n&(1<<0) != 0 }

func (item *Replace13ElemLong) SetC(v []int32, nat_k *uint32) {
	item.C = v
	if nat_k != nil {
		*nat_k |= 1 << 0
	}
}
func (item *Replace13ElemLong) ClearC(nat_k *uint32) {
	item.C = item.C[:0]
	if nat_k != nil {
		*nat_k &^= 1 << 0
	}
}
func (item *Replace13ElemLong) IsSetC(nat_k uint32) bool { return nat_k&(1<<0) != 0 }

func (item *Replace13ElemLong) Reset() {
	item.B = item.B[:0]
	item.C = item.C[:0]
}

func (item *Replace13ElemLong) FillRandom(rg *basictl.RandGenerator, nat_n uint32, nat_k uint32) {
	if nat_n&(1<<0) != 0 {
		tlBuiltinTuplePairBoxedIntLong.BuiltinTuplePairBoxedIntLongFillRandom(rg, &item.B, nat_k)
	} else {
		item.B = item.B[:0]
	}
	if nat_k&(1<<0) != 0 {
		tlBuiltinTupleInt.BuiltinTupleIntFillRandom(rg, &item.C, nat_n)
	} else {
		item.C = item.C[:0]
	}
}

func (item *Replace13ElemLong) Read(w []byte, nat_n uint32, nat_k uint32) (_ []byte, err error) {
	if nat_n&(1<<0) != 0 {
		if w, err = tlBuiltinTuplePairBoxedIntLong.BuiltinTuplePairBoxedIntLongRead(w, &item.B, nat_k); err != nil {
			return w, err
		}
	} else {
		item.B = item.B[:0]
	}
	if nat_k&(1<<0) != 0 {
		if w, err = tlBuiltinTupleInt.BuiltinTupleIntRead(w, &item.C, nat_n); err != nil {
			return w, err
		}
	} else {
		item.C = item.C[:0]
	}
	return w, nil
}

// This method is general version of Write, use it instead!
func (item *Replace13ElemLong) WriteGeneral(w []byte, nat_n uint32, nat_k uint32) (_ []byte, err error) {
	return item.Write(w, nat_n, nat_k)
}

func (item *Replace13ElemLong) Write(w []byte, nat_n uint32, nat_k uint32) (_ []byte, err error) {
	if nat_n&(1<<0) != 0 {
		if w, err = tlBuiltinTuplePairBoxedIntLong.BuiltinTuplePairBoxedIntLongWrite(w, item.B, nat_k); err != nil {
			return w, err
		}
	}
	if nat_k&(1<<0) != 0 {
		if w, err = tlBuiltinTupleInt.BuiltinTupleIntWrite(w, item.C, nat_n); err != nil {
			return w, err
		}
	}
	return w, nil
}

func (item *Replace13ElemLong) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, nat_n uint32, nat_k uint32) error {
	var rawB []byte
	var rawC []byte

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "b":
				if rawB != nil {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("replace13Elem", "b")
				}
				rawB = in.Raw()
				if !in.Ok() {
					return in.Error()
				}
			case "c":
				if rawC != nil {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("replace13Elem", "c")
				}
				rawC = in.Raw()
				if !in.Ok() {
					return in.Error()
				}
			default:
				return internal.ErrorInvalidJSONExcessElement("replace13Elem", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if nat_n&(1<<0) == 0 {
		if rawB != nil {
			return internal.ErrorInvalidJSON("replace13Elem", "field 'b' is defined, while corresponding implicit fieldmask bit is 0")
		}
		item.B = item.B[:0]
	} else {
		var inBPointer *basictl.JsonLexer
		inB := basictl.JsonLexer{Data: rawB}
		if rawB != nil {
			inBPointer = &inB
		}
		if err := tlBuiltinTuplePairBoxedIntLong.BuiltinTuplePairBoxedIntLongReadJSON(legacyTypeNames, inBPointer, &item.B, nat_k); err != nil {
			return err
		}

	}
	if nat_k&(1<<0) == 0 {
		if rawC != nil {
			return internal.ErrorInvalidJSON("replace13Elem", "field 'c' is defined, while corresponding implicit fieldmask bit is 0")
		}
		item.C = item.C[:0]
	} else {
		var inCPointer *basictl.JsonLexer
		inC := basictl.JsonLexer{Data: rawC}
		if rawC != nil {
			inCPointer = &inC
		}
		if err := tlBuiltinTupleInt.BuiltinTupleIntReadJSON(legacyTypeNames, inCPointer, &item.C, nat_n); err != nil {
			return err
		}

	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *Replace13ElemLong) WriteJSONGeneral(w []byte, nat_n uint32, nat_k uint32) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w, nat_n, nat_k)
}

func (item *Replace13ElemLong) WriteJSON(w []byte, nat_n uint32, nat_k uint32) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w, nat_n, nat_k)
}
func (item *Replace13ElemLong) WriteJSONOpt(newTypeNames bool, short bool, w []byte, nat_n uint32, nat_k uint32) (_ []byte, err error) {
	w = append(w, '{')
	if nat_n&(1<<0) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"b":`...)
		if w, err = tlBuiltinTuplePairBoxedIntLong.BuiltinTuplePairBoxedIntLongWriteJSONOpt(newTypeNames, short, w, item.B, nat_k); err != nil {
			return w, err
		}
	}
	if nat_k&(1<<0) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"c":`...)
		if w, err = tlBuiltinTupleInt.BuiltinTupleIntWriteJSONOpt(newTypeNames, short, w, item.C, nat_n); err != nil {
			return w, err
		}
	}
	return append(w, '}'), nil
}
