// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlcodegen

import (
	"fmt"
	"strings"
)

type TypeRWMaybe struct {
	wr      *TypeRWWrapper
	element Field

	emptyTag uint32
	okTag    uint32
}

func (trw *TypeRWMaybe) typeString2(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, isLocal bool, skipAlias bool) string {
	if isLocal {
		return addBytes(trw.wr.goLocalName, bytesVersion)
	}
	return trw.wr.ins.Prefix(directImports, ins) + addBytes(trw.wr.goGlobalName, bytesVersion)
}

func (trw *TypeRWMaybe) markHasBytesVersion(visitedNodes map[*TypeRWWrapper]bool) bool {
	return trw.element.t.MarkHasBytesVersion(visitedNodes)
}

func (trw *TypeRWMaybe) markWriteHasError(visitedNodes map[*TypeRWWrapper]bool) bool {
	return trw.element.t.MarkWriteHasError(visitedNodes)
}

func (trw *TypeRWMaybe) fillRecursiveUnwrap(visitedNodes map[*TypeRWWrapper]bool) {
}

func (trw *TypeRWMaybe) markWantsBytesVersion(visitedNodes map[*TypeRWWrapper]bool) {
	trw.element.t.MarkWantsBytesVersion(visitedNodes)
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

func (trw *TypeRWMaybe) typeResettingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, ref bool) string {
	return fmt.Sprintf("%s.Reset()", val)
}

func (trw *TypeRWMaybe) typeRandomCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, natArgs []string, ref bool) string {
	return fmt.Sprintf("%s.FillRandom(rg %s)", val, joinWithCommas(natArgs))
}

func (trw *TypeRWMaybe) typeWritingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, bare bool, natArgs []string, ref bool, last bool, needError bool) string {
	return wrapLastW(last, fmt.Sprintf("%s.Write%s(w %s)", val, addBare(bare), joinWithCommas(natArgs)), needError)
}

func (trw *TypeRWMaybe) typeReadingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, bare bool, natArgs []string, ref bool, last bool) string {
	return wrapLastW(last, fmt.Sprintf("%s.Read%s(w %s)", val, addBare(bare), joinWithCommas(natArgs)), true)
}

func (trw *TypeRWMaybe) typeJSONEmptyCondition(bytesVersion bool, val string, ref bool) string {
	return val + ".Ok"
}

func (trw *TypeRWMaybe) typeJSONWritingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, natArgs []string, ref bool, needError bool) string {
	if needError {
		return fmt.Sprintf("if w, err = %s.WriteJSONOpt(newTypeNames, short, w %s); err != nil { return w, err }", val, joinWithCommas(natArgs))
	} else {
		return fmt.Sprintf("w = %s.WriteJSONOpt(newTypeNames, short, w %s)", val, joinWithCommas(natArgs))
	}
}

func (trw *TypeRWMaybe) typeJSONReadingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, jvalue string, val string, natArgs []string, ref bool) string {
	return fmt.Sprintf("if err := %s.ReadJSONLegacy(legacyTypeNames, %s %s); err != nil { return err }", val, jvalue, joinWithCommas(natArgs))
}

func (trw *TypeRWMaybe) typeJSON2ReadingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, jvalue string, val string, natArgs []string, ref bool) string {
	return fmt.Sprintf("if err := %s.ReadJSON(legacyTypeNames, %s %s); err != nil { return err }", val, jvalue, joinWithCommas(natArgs))
}

func (trw *TypeRWMaybe) PhpClassName(withPath bool, bare bool) string {
	target := trw.getInnerTarget()
	return "maybe_" + target.t.trw.PhpClassName(withPath, target.bare)
}

func (trw *TypeRWMaybe) PhpClassNameReplaced() bool {
	return true
}

func (trw *TypeRWMaybe) PhpTypeName(withPath bool, bare bool) string {
	target := trw.getInnerTarget()
	return target.t.trw.PhpTypeName(withPath, target.t.PHPIsBare()) + "|null"
}

func (trw *TypeRWMaybe) getInnerTarget() Field {
	if inner, ok := trw.element.t.trw.(*TypeRWMaybe); ok {
		return inner.getInnerTarget()
	} else {
		return trw.element
	}
}

func (trw *TypeRWMaybe) PhpGenerateCode(code *strings.Builder, bytes bool) error {
	return fmt.Errorf("maybe doesn't have php code")
}

func (trw *TypeRWMaybe) PhpDefaultValue() string {
	return "null"
}

func (trw *TypeRWMaybe) PhpIterateReachableTypes(reachableTypes *map[*TypeRWWrapper]bool) {
	trw.element.t.PhpIterateReachableTypes(reachableTypes)
}
