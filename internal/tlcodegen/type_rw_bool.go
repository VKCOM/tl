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

type TypeRWBool struct {
	wr          *TypeRWWrapper
	falseGoName string
	trueGoName  string
	falseTag    uint32
	trueTag     uint32
}

func (trw *TypeRWBool) typeString2(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, isLocal bool, skipAlias bool) string {
	if !skipAlias {
		return "bool"
	}
	if isLocal {
		return addBytes(trw.wr.goLocalName, bytesVersion)
	}
	return trw.wr.ins.Prefix(directImports, ins) + addBytes(trw.wr.goGlobalName, bytesVersion)
}

func (trw *TypeRWBool) markHasBytesVersion(visitedNodes map[*TypeRWWrapper]bool) bool {
	return false
}

func (trw *TypeRWBool) markWriteHasError(visitedNodes map[*TypeRWWrapper]bool) bool {
	return false
}

func (trw *TypeRWBool) fillRecursiveUnwrap(visitedNodes map[*TypeRWWrapper]bool) {
}

func (trw *TypeRWBool) markWantsBytesVersion(visitedNodes map[*TypeRWWrapper]bool) {
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

func (trw *TypeRWBool) typeResettingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, ref bool) string {
	return fmt.Sprintf("%s = false", addAsterisk(ref, val))
}

func (trw *TypeRWBool) typeRandomCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, natArgs []string, ref bool) string {
	return fmt.Sprintf("%s = basictl.RandomUint(rg) & 1 == 1", addAsterisk(ref, val))
}

func (trw *TypeRWBool) typeWritingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, bare bool, natArgs []string, ref bool, last bool, needError bool) string {
	return wrapLastW(last, trw.wr.ins.Prefix(directImports, ins)+fmt.Sprintf("%sWrite%s(w, %s%s)", trw.wr.goGlobalName, addBare(bare), addAsterisk(ref, val), joinWithCommas(natArgs)), needError)
}

func (trw *TypeRWBool) typeReadingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, bare bool, natArgs []string, ref bool, last bool) string {
	return wrapLastW(last, trw.wr.ins.Prefix(directImports, ins)+fmt.Sprintf("%sRead%s(w, %s%s)", trw.wr.goGlobalName, addBare(bare), addAmpersand(ref, val), joinWithCommas(natArgs)), true)
}

func (trw *TypeRWBool) typeJSONEmptyCondition(bytesVersion bool, val string, ref bool) string {
	return addAsterisk(ref, val)
}

func (trw *TypeRWBool) typeJSONWritingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, natArgs []string, ref bool, needError bool) string {
	return fmt.Sprintf("w = basictl.JSONWriteBool(w, %s)", addAsterisk(ref, val))
}

func (trw *TypeRWBool) typeJSONReadingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, jvalue string, val string, natArgs []string, ref bool) string {
	return wrapLast(false, fmt.Sprintf("%sJsonReadBool(%s, %s)", trw.wr.gen.InternalPrefix(), jvalue, addAmpersand(ref, val)))
}

func (trw *TypeRWBool) typeJSON2ReadingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, jvalue string, val string, natArgs []string, ref bool) string {
	return wrapLast(false, fmt.Sprintf("%sJson2ReadBool(%s, %s)", trw.wr.gen.InternalPrefix(), jvalue, addAmpersand(ref, val)))
}

func (trw *TypeRWBool) PhpClassName(withPath bool, bare bool) string {
	return "boolean"
}

func (trw *TypeRWBool) PhpTypeName(withPath bool) string {
	return trw.PhpClassName(withPath, true)
}

func (trw *TypeRWBool) PhpGenerateCode(code *strings.Builder, bytes bool) error {
	return fmt.Errorf("boolean doesn't have php code")
}

func (trw *TypeRWBool) PhpDefaultValue() string {
	return "false"
}
