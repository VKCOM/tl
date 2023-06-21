// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlcodegen

import (
	"fmt"
)

type TypeRWBool struct {
	wr           *TypeRWWrapper
	goGlobalName string
	falseGoName  string
	trueGoName   string
	falseTag     uint32
	trueTag      uint32
}

func (trw *TypeRWBool) canBeBareOrBoxed(bare bool) bool {
	return !bare
}

func (trw *TypeRWBool) typeStringGlobal(bytesVersion bool) string {
	return addBytes(trw.goGlobalName, bytesVersion)
}

func (trw *TypeRWBool) typeString2(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, isLocal bool, skipAlias bool) string {
	return "bool"
}

func (trw *TypeRWBool) markHasBytesVersion(visitedNodes map[*TypeRWWrapper]bool) bool {
	return false
}

func (trw *TypeRWBool) fillRecursiveUnwrap(visitedNodes map[*TypeRWWrapper]bool) {
}

func (trw *TypeRWBool) markWantsBytesVersion(visitedNodes map[*TypeRWWrapper]bool) {
}

func (trw *TypeRWBool) BeforeCodeGenerationStep() error {
	return nil
}

func (trw *TypeRWBool) BeforeCodeGenerationStep2() {
}

func (trw *TypeRWBool) fillRecursiveChildren(visitedNodes map[*TypeRWWrapper]bool) {
}

func (trw *TypeRWBool) IsDictKeySafe() (isSafe bool, isString bool) {
	return false, false // TODO - maybe in future
}

func (trw *TypeRWBool) typeResettingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, ref bool) string {
	return fmt.Sprintf("%s = false", addAsterisk(ref, val))
}

func (trw *TypeRWBool) typeRandomCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, natArgs []string, ref bool) string {
	return fmt.Sprintf("%s = basictl.RandomNat(rand) & 1 == 1", addAsterisk(ref, val))
}

func (trw *TypeRWBool) typeWritingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, bare bool, natArgs []string, ref bool, last bool) string {
	return wrapLastW(last, trw.wr.ins.Prefix(directImports, ins)+fmt.Sprintf("%sWrite%s(w, %s%s)", trw.goGlobalName, addBare(bare), addAsterisk(ref, val), formatNatArgsCall(natArgs)))
}

func (trw *TypeRWBool) typeReadingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, bare bool, natArgs []string, ref bool, last bool) string {
	return wrapLastW(last, trw.wr.ins.Prefix(directImports, ins)+fmt.Sprintf("%sRead%s(w, %s%s)", trw.goGlobalName, addBare(bare), addAmpersand(ref, val), formatNatArgsCall(natArgs)))
}

func (trw *TypeRWBool) typeJSONEmptyCondition(bytesVersion bool, val string, ref bool) string {
	return addAsterisk(ref, val)
}

func (trw *TypeRWBool) typeJSONWritingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, natArgs []string, ref bool) string {
	return fmt.Sprintf("w = basictl.JSONWriteBool(w, %s)", addAsterisk(ref, val))
}

func (trw *TypeRWBool) typeJSONReadingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, jvalue string, val string, natArgs []string, ref bool) string {
	return wrapLast(false, fmt.Sprintf("%sJsonReadBool(%s, %s)", trw.wr.gen.InternalPrefix(), jvalue, addAmpersand(ref, val)))
}
