// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlcodegen

type TypeRWPrimitive struct {
	gen    *Gen2
	tlType string

	goType         string
	resetValue     string
	randomValue    string
	writeValue     string
	readValue      string
	writeJSONValue string
	readJSONValue  string
	readJSON2Value string

	cppFunctionSuffix string
	cppPrimitiveType  string
	cppDefaultInit    string
	cppResetValue     string
}

func (trw *TypeRWPrimitive) isFloat() bool {
	return trw.tlType == "float" || trw.tlType == "double"
}

func (trw *TypeRWPrimitive) fillRecursiveUnwrap(visitedNodes map[*TypeRWWrapper]bool) {
}

func (trw *TypeRWPrimitive) AllPossibleRecursionProducers() []*TypeRWWrapper {
	return nil
}

func (trw *TypeRWPrimitive) AllTypeDependencies(generic, countFunctions bool) []*TypeRWWrapper {
	return nil
}

func (trw *TypeRWPrimitive) IsWrappingType() bool {
	return true
}

func (trw *TypeRWPrimitive) ContainsUnion(visitedNodes map[*TypeRWWrapper]bool) bool {
	return false
}

func (trw *TypeRWPrimitive) FillRecursiveChildren(visitedNodes map[*TypeRWWrapper]int, generic bool) {
}

func (trw *TypeRWPrimitive) BeforeCodeGenerationStep1() {
}

func (trw *TypeRWPrimitive) BeforeCodeGenerationStep2() {
}

func (trw *TypeRWPrimitive) fillRecursiveChildren(visitedNodes map[*TypeRWWrapper]bool) {
}

func (trw *TypeRWPrimitive) IsDictKeySafe() (isSafe bool, isString bool) {
	return !trw.isFloat(), trw.tlType == "string"
}

func (trw *TypeRWPrimitive) CanBeBareBoxed() (canBare bool, canBoxed bool) {
	return true, false
}
