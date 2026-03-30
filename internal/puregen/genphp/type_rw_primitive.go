// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package genphp

type TypeRWPrimitive struct {
	gen    *Gen2
	tlType string

	goType         string
	resetValue     string
	randomValue    string
	writeValue     string
	readValue      string
	writeJSONValue string
}

func (trw *TypeRWPrimitive) isFloat() bool {
	return trw.tlType == "float" || trw.tlType == "double"
}

func (trw *TypeRWPrimitive) fillRecursiveUnwrap(visitedNodes map[*TypeRWWrapper]bool) {
}

func (trw *TypeRWPrimitive) BeforeCodeGenerationStep1() {
}

func (trw *TypeRWPrimitive) IsDictKeySafe() (isSafe bool, isString bool) {
	return !trw.isFloat(), trw.tlType == "string"
}

func (trw *TypeRWPrimitive) CanBeBareBoxed() (canBare bool, canBoxed bool) {
	return true, false
}
