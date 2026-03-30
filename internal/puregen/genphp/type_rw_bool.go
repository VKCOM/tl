// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package genphp

type TypeRWBool struct {
	wr          *TypeRWWrapper
	falseGoName string
	trueGoName  string
	falseTag    uint32
	trueTag     uint32

	isTL2 bool
	isBit bool
}

func (trw *TypeRWBool) fillRecursiveUnwrap(visitedNodes map[*TypeRWWrapper]bool) {
}

func (trw *TypeRWBool) AllPossibleRecursionProducers() []*TypeRWWrapper {
	return nil
}

func (trw *TypeRWBool) AllTypeDependencies(generic, countFunctions bool) []*TypeRWWrapper {
	return nil
}

func (trw *TypeRWBool) IsWrappingType() bool {
	return true
}

func (trw *TypeRWBool) ContainsUnion(visitedNodes map[*TypeRWWrapper]bool) bool {
	return false
}

func (trw *TypeRWBool) FillRecursiveChildren(visitedNodes map[*TypeRWWrapper]int, generic bool) {
}

func (trw *TypeRWBool) BeforeCodeGenerationStep1() {
}

func (trw *TypeRWBool) BeforeCodeGenerationStep2() {
}

func (trw *TypeRWBool) fillRecursiveChildren(visitedNodes map[*TypeRWWrapper]bool) {
}

func (trw *TypeRWBool) IsDictKeySafe() (isSafe bool, isString bool) {
	return false, false // TODO - low priority future
}

func (trw *TypeRWBool) CanBeBareBoxed() (canBare bool, canBoxed bool) {
	return false, true
}
