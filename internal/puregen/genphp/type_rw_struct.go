// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package genphp

import "github.com/VKCOM/tl/internal/pure"

type TypeRWStruct struct {
	pureType *pure.TypeInstanceStruct

	wr     *TypeRWWrapper
	Fields []Field

	ResultType    *TypeRWWrapper
	ResultNatArgs []pure.ActualNatArg

	fieldsDec Deconflicter // TODO - add all generated methods here
}

func (trw *TypeRWStruct) isTypeDef() bool {
	return len(trw.Fields) == 1 && trw.Fields[0].pureField.Name() == "" && trw.Fields[0].pureField.FieldMask() == nil && !trw.Fields[0].recursive
}

func (trw *TypeRWStruct) isUnwrapType() bool {
	if !trw.isTypeDef() {
		return false
	}
	// Motivation - we want default wrappers for primitive types, vector and tuple to generate primitive language types
	primitive, isPrimitive := trw.Fields[0].t.trw.(*TypeRWPrimitive)
	if isPrimitive && primitive.tlType == trw.wr.TLName().String() {
		return true
	}
	brackets, isBuiltinBrackets := trw.Fields[0].t.trw.(*TypeRWBrackets)
	if isBuiltinBrackets && (brackets.dictLike || trw.wr.TLName().String() == "vector" || trw.wr.TLName().String() == "tuple") {
		return true
	}
	//if trw.wr.gen.options.Language != "cpp" {
	//in combined TL Dictionary is defined via Vector.
	//dictionaryField {t:Type} key:string value:t = DictionaryField t;
	//dictionary#1f4c618f {t:Type} %(Vector %(DictionaryField t)) = Dictionary t;
	//TODO - change combined.tl to use # [] after we fully control generation of C++ & (k)PHP and remove code below
	str, isStruct := trw.Fields[0].t.trw.(*TypeRWStruct)
	if isStruct && str.wr.TLName().String() == "vector" {
		// repeat check above 1 level deeper
		brackets, isBuiltinBrackets := str.Fields[0].t.trw.(*TypeRWBrackets)
		if isBuiltinBrackets && brackets.dictLike {
			return true
		}
	}
	//}
	return false
}
