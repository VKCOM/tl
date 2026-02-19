// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package gengo

import (
	"fmt"
)

// check that brackets cannot be function return type

type TypeRWBrackets struct {
	wr          *TypeRWWrapper
	vectorLike  bool   // # [T], because # has no reference name
	dynamicSize bool   // with passed nat param
	size        uint32 // if !dynamicSize
	element     Field
}

var _ TypeRW = &TypeRWBrackets{}

func (trw *TypeRWBrackets) typeString2(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, isLocal bool, skipAlias bool) string {
	if trw.vectorLike || trw.dynamicSize {
		return fmt.Sprintf("[]%s", trw.element.t.TypeString2(bytesVersion, directImports, ins, isLocal, skipAlias))
	}
	return fmt.Sprintf("[%d]%s", trw.size, trw.element.t.TypeString2(bytesVersion, directImports, ins, isLocal, skipAlias))
}

func (trw *TypeRWBrackets) markHasBytesVersion(visitedNodes map[*TypeRWWrapper]bool) bool {
	return trw.element.t.MarkHasBytesVersion(visitedNodes)
}

func (trw *TypeRWBrackets) markWriteHasError(visitedNodes map[*TypeRWWrapper]bool) bool {
	if trw.dynamicSize {
		return true
	}
	return trw.element.t.MarkWriteHasError(visitedNodes)
}

func (trw *TypeRWBrackets) fillRecursiveUnwrap(visitedNodes map[*TypeRWWrapper]bool) {
	trw.element.t.FillRecursiveUnwrap(visitedNodes)
}

func (trw *TypeRWBrackets) markWantsBytesVersion(visitedNodes map[*TypeRWWrapper]bool) {
	trw.element.t.MarkWantsBytesVersion(visitedNodes)
}

func (trw *TypeRWBrackets) markWantsTL2(visitedNodes map[*TypeRWWrapper]bool) {
	trw.element.t.MarkWantsTL2(visitedNodes)
}

func (trw *TypeRWBrackets) FillRecursiveChildren(visitedNodes map[*TypeRWWrapper]int, generic bool) {
}

func (trw *TypeRWBrackets) ContainsUnion(visitedNodes map[*TypeRWWrapper]bool) bool {
	return trw.element.t.containsUnion(visitedNodes)
}

func (trw *TypeRWBrackets) BeforeCodeGenerationStep1() {
}

func (trw *TypeRWBrackets) BeforeCodeGenerationStep2() {
}

func (trw *TypeRWBrackets) fillRecursiveChildren(visitedNodes map[*TypeRWWrapper]bool) {
	if trw.vectorLike || trw.dynamicSize {
		return
	}
	trw.element.t.FillRecursiveChildren(visitedNodes)
}

func (trw *TypeRWBrackets) typeResettingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, ref bool) string {
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
	if trw.dynamicSize || trw.vectorLike {
		if ref {
			return fmt.Sprintf("*%[1]s = (*%[1]s)[:0]", val)
		}
		return fmt.Sprintf("%[1]s = %[1]s[:0]", val)
	}
	return trw.wr.ins.Prefix(directImports, ins) + fmt.Sprintf("%[1]sReset(%[2]s)", goGlobalName, addAmpersand(ref, val))
}

func (trw *TypeRWBrackets) typeRandomCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, natArgs []string, ref bool) string {
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
	return trw.wr.ins.Prefix(directImports, ins) + fmt.Sprintf("%sFillRandom(rg, %s%s)", goGlobalName, addAmpersand(ref, val), joinWithCommas(natArgs))
}

func (trw *TypeRWBrackets) typeWritingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, bare bool, natArgs []string, ref bool, last bool, needError bool) string {
	refVal := addAmpersand(ref, val)
	if trw.vectorLike || trw.dynamicSize {
		refVal = addAsterisk(ref, val) // those version pass to Write method by pointer
	}
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
	return wrapLastW(last, trw.wr.ins.Prefix(directImports, ins)+fmt.Sprintf("%sWrite%s(w, %s%s)", goGlobalName, addBare(bare), refVal, joinWithCommas(natArgs)), needError)
}

func (trw *TypeRWBrackets) typeReadingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, bare bool, natArgs []string, ref bool, last bool) string {
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
	return wrapLastW(last, trw.wr.ins.Prefix(directImports, ins)+fmt.Sprintf("%sRead%s(w, %s%s)", goGlobalName, addBare(bare), addAmpersand(ref, val), joinWithCommas(natArgs)), true)
}

func (trw *TypeRWBrackets) typeJSONEmptyCondition(bytesVersion bool, val string, ref bool) string {
	if trw.vectorLike || trw.dynamicSize {
		return fmt.Sprintf("len(%s) != 0", addAsterisk(ref, val))
	}
	return ""
}

func (trw *TypeRWBrackets) typeJSONWritingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, natArgs []string, ref bool, needError bool) string {
	refVal := addAmpersand(ref, val)
	if trw.vectorLike || trw.dynamicSize {
		refVal = addAsterisk(ref, val) // those version pass to Write method by pointer
	}
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
	// Code which depends on serialization location (skipping empty array if object property) is generated in that location.
	if needError {
		return fmt.Sprintf("if w, err = %sWriteJSONOpt(tctx, w, %s%s); err != nil { return w, err }", trw.wr.ins.Prefix(directImports, ins)+goGlobalName, refVal, joinWithCommas(natArgs))
	} else {
		return fmt.Sprintf("w = %sWriteJSONOpt(tctx, w, %s%s)", trw.wr.ins.Prefix(directImports, ins)+goGlobalName, refVal, joinWithCommas(natArgs))
	}
}

func (trw *TypeRWBrackets) typeJSONReadingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, jvalue string, val string, natArgs []string, ref bool) string {
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
	return fmt.Sprintf("if err := %sReadJSONLegacy(legacyTypeNames, %s, %s%s); err != nil { return err }", trw.wr.ins.Prefix(directImports, ins)+goGlobalName, jvalue, addAmpersand(ref, val), joinWithCommas(natArgs))
}

func (trw *TypeRWBrackets) typeJSON2ReadingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, jvalue string, val string, natArgs []string, ref bool) string {
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
	return fmt.Sprintf("if err := %sReadJSONGeneral(tctx, %s, %s%s); err != nil { return err }", trw.wr.ins.Prefix(directImports, ins)+goGlobalName, jvalue, addAmpersand(ref, val), joinWithCommas(natArgs))
}

func (trw *TypeRWBrackets) typeJSON2ReadingRequiresContext() bool {
	return true
}
