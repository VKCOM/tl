// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package genphp

type TypeRWMaybe struct {
	wr      *TypeRWWrapper
	element Field

	emptyTag uint32
	okTag    uint32
}

func (trw *TypeRWMaybe) fillRecursiveUnwrap(visitedNodes map[*TypeRWWrapper]bool) {
}

func (trw *TypeRWMaybe) BeforeCodeGenerationStep1() {
}

func (trw *TypeRWMaybe) IsDictKeySafe() (isSafe bool, isString bool) {
	return false, false
}

func (trw *TypeRWMaybe) CanBeBareBoxed() (canBare bool, canBoxed bool) {
	return false, true
}
