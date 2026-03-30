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

func (trw *TypeRWMaybe) AllPossibleRecursionProducers() []*TypeRWWrapper {
	return trw.element.t.trw.AllPossibleRecursionProducers()
}

func (trw *TypeRWMaybe) AllTypeDependencies(generic, countFunctions bool) (res []*TypeRWWrapper) {
	if !generic {
		res = append(res, trw.element.t)
	}
	return
}

func (trw *TypeRWMaybe) IsWrappingType() bool {
	return true
}

func (trw *TypeRWMaybe) ContainsUnion(visitedNodes map[*TypeRWWrapper]bool) bool {
	return trw.element.t.containsUnion(visitedNodes)
}

func (trw *TypeRWMaybe) FillRecursiveChildren(visitedNodes map[*TypeRWWrapper]int, generic bool) {
	visitedNodes[trw.wr] = 1
	trw.element.t.trw.FillRecursiveChildren(visitedNodes, generic)
	visitedNodes[trw.wr] = 2
}

func (trw *TypeRWMaybe) BeforeCodeGenerationStep1() {
}

func (trw *TypeRWMaybe) BeforeCodeGenerationStep2() {
}

func (trw *TypeRWMaybe) fillRecursiveChildren(visitedNodes map[*TypeRWWrapper]bool) {
	trw.element.t.FillRecursiveChildren(visitedNodes)
}

func (trw *TypeRWMaybe) IsDictKeySafe() (isSafe bool, isString bool) {
	return false, false
}

func (trw *TypeRWMaybe) CanBeBareBoxed() (canBare bool, canBoxed bool) {
	return false, true
}
