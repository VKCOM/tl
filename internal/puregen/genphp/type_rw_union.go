// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package genphp

type TypeRWUnion struct {
	wr     *TypeRWWrapper
	Fields []Field
	IsEnum bool

	fieldsDec    Deconflicter // TODO - add all generated methods here
	fieldsDecCPP Deconflicter // TODO - add all generated methods here
}

func (trw *TypeRWUnion) fillRecursiveUnwrap(visitedNodes map[*TypeRWWrapper]bool) {
}

func (trw *TypeRWUnion) BeforeCodeGenerationStep1() {
}

func (trw *TypeRWUnion) IsDictKeySafe() (isSafe bool, isString bool) {
	return false, false // trw.IsEnum - TODO - in the future?
}

func (trw *TypeRWUnion) CanBeBareBoxed() (canBare bool, canBoxed bool) {
	return false, true
}
