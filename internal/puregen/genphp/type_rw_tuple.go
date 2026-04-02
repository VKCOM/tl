// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package genphp

import (
	"fmt"
	"strings"
)

// check that brackets cannot be function return type

type TypeRWBrackets struct {
	wr          *TypeRWWrapper
	vectorLike  bool   // # [T], because # has no reference name
	dynamicSize bool   // with passed nat param
	size        uint32 // if !dynamicSize
	element     Field

	dictLike       bool // for now, can be true only if vectorLike is true. But should work for dynamicSize tuples, so TODO
	dictKeyString  bool
	dictKeyField   Field
	dictValueField Field
}

func (trw *TypeRWBrackets) fillRecursiveUnwrap(visitedNodes map[*TypeRWWrapper]bool) {
	trw.element.t.FillRecursiveUnwrap(visitedNodes)
}

func isDictionaryElement(wr *TypeRWWrapper) (bool, bool, Field, Field) {
	// it is hard to mark Dictionary constructor as dictionary,
	// because it is typedef to Vector or built-in brackets.
	// TODO: FIX IT, because len(structElement.Fields) != 2 is true
	structElement, ok := wr.trw.(*TypeRWStruct)
	if !ok || len(structElement.Fields) != 2 || !strings.Contains(strings.ToLower(wr.TLName().Name), "dictionary") {
		return false, false, Field{}, Field{}
	}
	if structElement.Fields[0].pureField.FieldMask() != nil { // TODO - allowing this complicates json serialization
		return false, false, Field{}, Field{}
	}
	ok, isString := structElement.Fields[0].t.trw.IsDictKeySafe()
	return ok, isString, structElement.Fields[0], structElement.Fields[1]
}

func phpDictionaryElement(wr *TypeRWWrapper) Field {
	if !phpIsDictionary(wr) {
		panic(fmt.Sprintf("not a dict: %s", wr.TLName()))
	}
	structElement, _ := wr.trw.(*TypeRWStruct)
	return structElement.Fields[1]
}

func phpIsDictionary(wr *TypeRWWrapper) bool {
	if PHPIsDict(wr.pureType.KernelType()) {
		return true
	}
	isDict, _, _, _ := isDictionaryElement(wr)
	if isDict && wr.TLName().Namespace == "" { // TODO NOT A SOLUTION, BUT...
		return true
	}
	return false
}

func (trw *TypeRWBrackets) BeforeCodeGenerationStep1() {
	if trw.vectorLike {
		if ok, isString, kf, vf := isDictionaryElement(trw.element.t); ok {
			trw.dictLike = true
			trw.dictKeyString = isString
			trw.dictKeyField = kf
			trw.dictValueField = vf
		}
	}
}

func (trw *TypeRWBrackets) IsDictKeySafe() (isSafe bool, isString bool) {
	return false, false
}

func (trw *TypeRWBrackets) CanBeBareBoxed() (canBare bool, canBoxed bool) {
	return true, false
}
