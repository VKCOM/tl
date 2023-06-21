// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlcodegen

import (
	"fmt"
)

type TypeRWMaybe struct {
	wr           *TypeRWWrapper
	goGlobalName string
	goLocalName  string
	element      Field

	emptyTag uint32
	okTag    uint32
}

func (trw *TypeRWMaybe) canBeBareOrBoxed(bare bool) bool {
	return !bare
}

func (trw *TypeRWMaybe) typeStringGlobal(bytesVersion bool) string {
	return addBytes(trw.goGlobalName, bytesVersion)
}

func (trw *TypeRWMaybe) typeString2(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, isLocal bool, skipAlias bool) string {
	if isLocal {
		return addBytes(trw.goLocalName, bytesVersion)
	}
	return trw.wr.ins.Prefix(directImports, ins) + addBytes(trw.goGlobalName, bytesVersion)
}

func (trw *TypeRWMaybe) markHasBytesVersion(visitedNodes map[*TypeRWWrapper]bool) bool {
	return trw.element.t.MarkHasBytesVersion(visitedNodes)
}

func (trw *TypeRWMaybe) fillRecursiveUnwrap(visitedNodes map[*TypeRWWrapper]bool) {
}

func (trw *TypeRWMaybe) markWantsBytesVersion(visitedNodes map[*TypeRWWrapper]bool) {
	trw.element.t.MarkWantsBytesVersion(visitedNodes)
}

func (trw *TypeRWMaybe) BeforeCodeGenerationStep() error {
	return trw.element.checkBareBoxed()
}

func (trw *TypeRWMaybe) BeforeCodeGenerationStep2() {
}

func (trw *TypeRWMaybe) fillRecursiveChildren(visitedNodes map[*TypeRWWrapper]bool) {
	trw.element.t.FillRecursiveChildren(visitedNodes)
}

func (trw *TypeRWMaybe) IsDictKeySafe() (isSafe bool, isString bool) {
	return false, false // TODO - possible in future?
}

func (trw *TypeRWMaybe) typeResettingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, ref bool) string {
	return fmt.Sprintf("%s.Reset()", val)
}

func (trw *TypeRWMaybe) typeRandomCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, natArgs []string, ref bool) string {
	return fmt.Sprintf("%s.FillRandom(rand %s)", val, formatNatArgsCall(natArgs))
}

func (trw *TypeRWMaybe) typeWritingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, bare bool, natArgs []string, ref bool, last bool) string {
	return wrapLastW(last, fmt.Sprintf("%s.Write%s(w %s)", val, addBare(bare), formatNatArgsCall(natArgs)))
}

func (trw *TypeRWMaybe) typeReadingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, bare bool, natArgs []string, ref bool, last bool) string {
	return wrapLastW(last, fmt.Sprintf("%s.Read%s(w %s)", val, addBare(bare), formatNatArgsCall(natArgs)))
}

func (trw *TypeRWMaybe) typeJSONEmptyCondition(bytesVersion bool, val string, ref bool) string {
	// TODO - return "item.Ok"
	return val + ".Ok"
}

func (trw *TypeRWMaybe) typeJSONWritingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, natArgs []string, ref bool) string {
	return fmt.Sprintf("if w, err = %s.WriteJSON(w %s); err != nil { return w, err }", val, formatNatArgsCall(natArgs))
}

func (trw *TypeRWMaybe) typeJSONReadingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, jvalue string, val string, natArgs []string, ref bool) string {
	goName := addBytes(trw.goGlobalName, bytesVersion)
	return fmt.Sprintf("if err := %s__ReadJSON(%s, %s %s); err != nil { return err }", trw.wr.ins.Prefix(directImports, ins)+goName, addAmpersand(ref, val), jvalue, formatNatArgsCall(natArgs))
}
