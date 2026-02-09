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

type TypeRWDict struct {
	wr      *TypeRWWrapper
	element Field

	dictKeyString  bool
	dictKeyField   Field
	dictValueField Field
}

var _ TypeRW = &TypeRWDict{}

func (trw *TypeRWDict) typeString2(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, isLocal bool, skipAlias bool) string {
	if !bytesVersion {
		return fmt.Sprintf("map[%s]%s",
			trw.dictKeyField.t.TypeString2(bytesVersion, directImports, ins, isLocal, skipAlias),
			trw.dictValueField.t.TypeString2(bytesVersion, directImports, ins, isLocal, skipAlias))
	}
	return fmt.Sprintf("[]%s", trw.element.t.TypeString2(bytesVersion, directImports, ins, isLocal, skipAlias))
}

func (trw *TypeRWDict) markHasBytesVersion(visitedNodes map[*TypeRWWrapper]bool) bool {
	return true
}

func (trw *TypeRWDict) markWriteHasError(visitedNodes map[*TypeRWWrapper]bool) bool {
	return trw.element.t.MarkWriteHasError(visitedNodes)
}

func (trw *TypeRWDict) fillRecursiveUnwrap(visitedNodes map[*TypeRWWrapper]bool) {
	trw.element.t.FillRecursiveUnwrap(visitedNodes)
}

func (trw *TypeRWDict) markWantsBytesVersion(visitedNodes map[*TypeRWWrapper]bool) {
	trw.element.t.MarkWantsBytesVersion(visitedNodes)
}

func (trw *TypeRWDict) markWantsTL2(visitedNodes map[*TypeRWWrapper]bool) {
	trw.element.t.MarkWantsTL2(visitedNodes)
}

func (trw *TypeRWDict) FillRecursiveChildren(visitedNodes map[*TypeRWWrapper]int, generic bool) {
}

func (trw *TypeRWDict) ContainsUnion(visitedNodes map[*TypeRWWrapper]bool) bool {
	return trw.element.t.containsUnion(visitedNodes)
}

func (trw *TypeRWDict) BeforeCodeGenerationStep1() {
}

func (trw *TypeRWDict) BeforeCodeGenerationStep2() {
}

func (trw *TypeRWDict) fillRecursiveChildren(visitedNodes map[*TypeRWWrapper]bool) {
}

func (trw *TypeRWDict) IsDictKeySafe() (isSafe bool, isString bool) {
	return false, false
}

func (trw *TypeRWDict) CanBeBareBoxed() (canBare bool, canBoxed bool) {
	return true, true
}

func (trw *TypeRWDict) typeResettingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, ref bool) string {
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
	if !bytesVersion {
		return trw.wr.ins.Prefix(directImports, ins) + fmt.Sprintf("%[1]sReset(%s)", goGlobalName, addAsterisk(ref, val))
	}
	if ref {
		return fmt.Sprintf("*%[1]s = (*%[1]s)[:0]", val)
	}
	return fmt.Sprintf("%[1]s = %[1]s[:0]", val)
}

func (trw *TypeRWDict) typeRandomCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, natArgs []string, ref bool) string {
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
	return trw.wr.ins.Prefix(directImports, ins) + fmt.Sprintf("%sFillRandom(rg, %s%s)", goGlobalName, addAmpersand(ref, val), joinWithCommas(natArgs))
}

func (trw *TypeRWDict) typeWritingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, bare bool, natArgs []string, ref bool, last bool, needError bool) string {
	//prefix := ""
	//if !bare {
	//	prefix = fmt.Sprintf("w = basictl.NatWrite(w, 0x%x)\n", trw.wr.tlTag)
	//}
	refVal := addAsterisk(ref, val)
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
	//return prefix + wrapLastW(last, trw.wr.ins.Prefix(directImports, ins)+fmt.Sprintf("%sWrite(w, %s%s)", goGlobalName, refVal, joinWithCommas(natArgs)), needError)
	return wrapLastW(last, trw.wr.ins.Prefix(directImports, ins)+fmt.Sprintf("%sWrite%s(w, %s%s)", goGlobalName, addBare(bare), refVal, joinWithCommas(natArgs)), needError)
}

func (trw *TypeRWDict) typeReadingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, bare bool, natArgs []string, ref bool, last bool) string {
	//prefix := ""
	//if !bare {
	//	prefix = fmt.Sprintf("if w, err = basictl.NatReadExactTag(w, 0x%x); err != nil {\nreturn w, err\n}\n", trw.wr.tlTag)
	//}
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
	//return prefix + wrapLastW(last, trw.wr.ins.Prefix(directImports, ins)+fmt.Sprintf("%sRead(w, %s%s)", goGlobalName, addAmpersand(ref, val), joinWithCommas(natArgs)), true)
	return wrapLastW(last, trw.wr.ins.Prefix(directImports, ins)+fmt.Sprintf("%sRead%s(w, %s%s)", goGlobalName, addBare(bare), addAmpersand(ref, val), joinWithCommas(natArgs)), true)
}

func (trw *TypeRWDict) typeJSONEmptyCondition(bytesVersion bool, val string, ref bool) string {
	return fmt.Sprintf("len(%s) != 0", addAsterisk(ref, val))
}

func (trw *TypeRWDict) typeJSONWritingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, natArgs []string, ref bool, needError bool) string {
	//refVal := addAmpersand(ref, val)
	//if !bytesVersion {
	refVal := addAsterisk(ref, val) // those version pass to Write method by pointer
	//}
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
	// Code which depends on serialization location (skipping empty array if object property) is generated in that location.
	if needError {
		return fmt.Sprintf("if w, err = %sWriteJSONOpt(tctx, w, %s%s); err != nil { return w, err }", trw.wr.ins.Prefix(directImports, ins)+goGlobalName, refVal, joinWithCommas(natArgs))
	} else {
		return fmt.Sprintf("w = %sWriteJSONOpt(tctx, w, %s%s)", trw.wr.ins.Prefix(directImports, ins)+goGlobalName, refVal, joinWithCommas(natArgs))
	}
}

func (trw *TypeRWDict) typeJSONReadingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, jvalue string, val string, natArgs []string, ref bool) string {
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
	return fmt.Sprintf("if err := %sReadJSONLegacy(legacyTypeNames, %s, %s%s); err != nil { return err }", trw.wr.ins.Prefix(directImports, ins)+goGlobalName, jvalue, addAmpersand(ref, val), joinWithCommas(natArgs))
}

func (trw *TypeRWDict) typeJSON2ReadingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, jvalue string, val string, natArgs []string, ref bool) string {
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
	return fmt.Sprintf("if err := %sReadJSONGeneral(tctx, %s, %s%s); err != nil { return err }", trw.wr.ins.Prefix(directImports, ins)+goGlobalName, jvalue, addAmpersand(ref, val), joinWithCommas(natArgs))
}

func (trw *TypeRWDict) typeJSON2ReadingRequiresContext() bool {
	return true
}
