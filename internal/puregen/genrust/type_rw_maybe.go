// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package genrust

import (
	"fmt"
	strings "strings"

	"github.com/VKCOM/tl/internal/tlcodegen/codecreator"
)

type TypeRWMaybe struct {
	wr      *TypeRWWrapper
	element Field

	emptyTag uint32
	okTag    uint32
}

var _ TypeRW = &TypeRWMaybe{}

func (trw *TypeRWMaybe) typeString2(bytesVersion bool, directImports *DirectImports, isLocal bool, skipAlias bool) string {
	if isLocal {
		return addBytes(trw.wr.goLocalName, bytesVersion)
	}
	return addBytes(trw.wr.goGlobalName, bytesVersion)
}

func (trw *TypeRWMaybe) markHasBytesVersion(visitedNodes map[*TypeRWWrapper]bool) bool {
	return trw.element.t.MarkHasBytesVersion(visitedNodes)
}

func (trw *TypeRWMaybe) markHasRepairMasks(visitedNodes map[*TypeRWWrapper]bool) bool {
	return trw.element.t.MarkHasRepairMasks(visitedNodes)
}

func (trw *TypeRWMaybe) markWriteHasError(visitedNodes map[*TypeRWWrapper]bool) bool {
	return trw.element.t.MarkWriteHasError(visitedNodes)
}

func (trw *TypeRWMaybe) markWantsBytesVersion(visitedNodes map[*TypeRWWrapper]bool) {
	trw.element.t.MarkWantsBytesVersion(visitedNodes)
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

func (trw *TypeRWMaybe) typeResettingCode(cc *codecreator.RustCodeCreator, bytesVersion bool, directImports *DirectImports, val string, ref bool) {
	cc.AddLinef("%s.Reset()", val)
}

func (trw *TypeRWMaybe) typeRandomCode(bytesVersion bool, directImports *DirectImports, val string, natArgs []string, ref bool) string {
	return fmt.Sprintf("%s.FillRandom(rg %s)", val, joinWithCommas(natArgs))
}

func (trw *TypeRWMaybe) typeRepairMasksCode(bytesVersion bool, directImports *DirectImports, val string, natArgs []string, ref bool) string {
	return fmt.Sprintf("%s.RepairMasks(%s)", val, strings.Join(natArgs, ","))
}

func (trw *TypeRWMaybe) typeWritingCode(bytesVersion bool, directImports *DirectImports, val string, bare bool, natArgs []string, ref bool, last bool, needError bool) string {
	return wrapLastW(last, fmt.Sprintf("%s.WriteTL1%s(w %s %s)", val, addBare(bare), joinWithCommas(natArgs), trw.wr.fetcherCall()), needError)
}

func (trw *TypeRWMaybe) typeReadingCode(cc *codecreator.RustCodeCreator, bytesVersion bool, directImports *DirectImports, val string, bare bool, natArgs []string, ref bool) {
	cc.AddLines("TypeRWMaybe::typeReadingCode")
	//	return wrapLastW(last, fmt.Sprintf("%s.ReadTL1%s(w %s %s)", val, addBare(bare), joinWithCommas(natArgs), trw.wr.fetcherCall()), true)
}

func (trw *TypeRWMaybe) typeJSONEmptyCondition(bytesVersion bool, val string, ref bool) string {
	if ref {
		return val + "!= nil && " + val + ".Ok"
	}
	return val + ".Ok"
}

func (trw *TypeRWMaybe) typeJSONWritingCode(bytesVersion bool, directImports *DirectImports, val string, natArgs []string, ref bool, needError bool) string {
	if needError {
		return fmt.Sprintf("if w, err = %s.WriteJSONOpt(jctx, w %s); err != nil { return w, err }", val, joinWithCommas(natArgs))
	} else {
		return fmt.Sprintf("w = %s.WriteJSONOpt(jctx, w %s)", val, joinWithCommas(natArgs))
	}
}

func (trw *TypeRWMaybe) typeJSON2ReadingCode(bytesVersion bool, directImports *DirectImports, jvalue string, val string, natArgs []string, ref bool) string {
	return fmt.Sprintf("if err := %s.ReadJSONGeneral(jctx, %s %s); err != nil { return err }", val, jvalue, joinWithCommas(natArgs))
}

func (trw *TypeRWMaybe) GenerateCode(bytesVersion bool, directImports *DirectImports) string {
	return ""
}
