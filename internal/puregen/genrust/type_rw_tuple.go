// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package genrust

import (
	"fmt"

	"github.com/VKCOM/tl/internal/tlcodegen/codecreator"
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

func (trw *TypeRWBrackets) typeString2(bytesVersion bool, directImports *DirectImports, isLocal bool, skipAlias bool) string {
	if trw.vectorLike || trw.dynamicSize {
		return fmt.Sprintf("Vec<%s>", trw.element.t.TypeString2(bytesVersion, directImports, isLocal, skipAlias))
	}
	return fmt.Sprintf("[%s; %d]", trw.element.t.TypeString2(bytesVersion, directImports, isLocal, skipAlias), trw.size)
}

func (trw *TypeRWBrackets) markHasBytesVersion(visitedNodes map[*TypeRWWrapper]bool) bool {
	return trw.element.t.MarkHasBytesVersion(visitedNodes)
}

func (trw *TypeRWBrackets) markHasRepairMasks(visitedNodes map[*TypeRWWrapper]bool) bool {
	return trw.element.t.MarkHasRepairMasks(visitedNodes)
}

func (trw *TypeRWBrackets) markWriteHasError(visitedNodes map[*TypeRWWrapper]bool) bool {
	if trw.dynamicSize {
		return true
	}
	return trw.element.t.MarkWriteHasError(visitedNodes)
}

func (trw *TypeRWBrackets) markWantsBytesVersion(visitedNodes map[*TypeRWWrapper]bool) {
	trw.element.t.MarkWantsBytesVersion(visitedNodes)
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

func (trw *TypeRWBrackets) typeResettingCode(bytesVersion bool, directImports *DirectImports, val string, ref bool) string {
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
	if trw.dynamicSize || trw.vectorLike {
		if ref {
			return fmt.Sprintf("*%[1]s = (*%[1]s)[:0]", val)
		}
		return fmt.Sprintf("%[1]s = %[1]s[:0]", val)
	}
	return fmt.Sprintf("%[1]sReset(%[2]s)", goGlobalName, addAmpersand(ref, val))
}

func (trw *TypeRWBrackets) typeRandomCode(bytesVersion bool, directImports *DirectImports, val string, natArgs []string, ref bool) string {
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
	return fmt.Sprintf("%sFillRandom(rg, %s%s)", goGlobalName, addAmpersand(ref, val), joinWithCommas(natArgs))
}

func (trw *TypeRWBrackets) typeRepairMasksCode(bytesVersion bool, directImports *DirectImports, val string, natArgs []string, ref bool) string {
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
	return fmt.Sprintf("%sRepairMasks(%s%s)", goGlobalName, addAmpersand(ref, val), joinWithCommas(natArgs))
}

func (trw *TypeRWBrackets) typeWritingCode(bytesVersion bool, directImports *DirectImports, val string, bare bool, natArgs []string, ref bool, last bool, needError bool) string {
	refVal := addAmpersand(ref, val)
	if trw.vectorLike || trw.dynamicSize {
		refVal = addAsterisk(ref, val) // those version pass to Write method by pointer
	}
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
	return wrapLastW(last, fmt.Sprintf("%sWriteTL1%s(w, %s%s %s)", goGlobalName, addBare(bare), refVal, joinWithCommas(natArgs), trw.wr.fetcherCall()), needError)
}

func (trw *TypeRWBrackets) typeReadingCode(cc *codecreator.RustCodeCreator, bytesVersion bool, directImports *DirectImports, val string, bare bool, natArgs []string, ref bool) {
	//goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
	//return wrapLastW(last, fmt.Sprintf("%sReadTL1%s(w, %s%s %s)", goGlobalName, addBare(bare), addAmpersand(ref, val), joinWithCommas(natArgs), trw.wr.fetcherCall()), true)
	cc.AddLinef("crate::types::%s::read_tl1(&mut %s, buf%s%s)?;", trw.wr.goGlobalName, addAsterisk(ref, val), joinWithCommas(natArgs), trw.wr.fetcherCall())
}

func (trw *TypeRWBrackets) typeJSONEmptyCondition(bytesVersion bool, val string, ref bool) string {
	if trw.vectorLike || trw.dynamicSize {
		return fmt.Sprintf("len(%s) != 0", addAsterisk(ref, val))
	}
	return ""
}

func (trw *TypeRWBrackets) typeJSONWritingCode(bytesVersion bool, directImports *DirectImports, val string, natArgs []string, ref bool, needError bool) string {
	refVal := addAmpersand(ref, val)
	if trw.vectorLike || trw.dynamicSize {
		refVal = addAsterisk(ref, val) // those version pass to Write method by pointer
	}
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
	// Code which depends on serialization location (skipping empty array if object property) is generated in that location.
	if needError {
		return fmt.Sprintf("if w, err = %sWriteJSONOpt(jctx, w, %s%s); err != nil { return w, err }", goGlobalName, refVal, joinWithCommas(natArgs))
	} else {
		return fmt.Sprintf("w = %sWriteJSONOpt(jctx, w, %s%s)", goGlobalName, refVal, joinWithCommas(natArgs))
	}
}

func (trw *TypeRWBrackets) typeJSON2ReadingCode(bytesVersion bool, directImports *DirectImports, jvalue string, val string, natArgs []string, ref bool) string {
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
	return fmt.Sprintf("if err := %sReadJSONGeneral(jctx, %s, %s%s); err != nil { return err }", goGlobalName, jvalue, addAmpersand(ref, val), joinWithCommas(natArgs))
}

func (trw *TypeRWBrackets) GenerateCode(bytesVersion bool, directImports *DirectImports) string {
	cc := codecreator.NewRustCodeCreator()

	//goName := addBytes(trw.wr.goGlobalName, bytesVersion)
	natDecl := trw.wr.formatNatArgsDecl()
	//natCall := trw.wr.formatNatArgsDeclCall()
	typeString := trw.wr.TypeString2(bytesVersion, directImports, false, false)
	elementTypeString := trw.element.t.TypeString2(bytesVersion, directImports, false, false)

	cc.AddLinef("use basictl::TLRead as _;")
	cc.AddEmptyLine()
	cc.AddLinef("pub(crate) fn read_tl1<B: bytes::Buf + Copy>(value: &mut %s, buf: &mut B%s) -> basictl::Result<()> {", typeString, natDecl)
	cc.AddBlock(func() {
		if trw.wr.OriginTL2() {
			cc.AddLinef(`Err(basictl::Error::NoTL1("%s")`, trw.wr.pureType.CanonicalName())
			return
		}
		switch {
		case trw.vectorLike:
			cc.AddLinef("let l = buf.read_u32()?;")
			// TODO - use length sanity check
			cc.AddLinef("value.resize_with(l as usize, %s::default);", elementTypeString)
			cc.For("element", "value.iter_mut()", "", func() {
				trw.element.t.TypeReadingCode(cc, bytesVersion, directImports, "*element", trw.element.Bare(), trw.wr.formatNatArgs(nil, trw.element.NatArgs()), false)
			})
		case trw.dynamicSize:
			cc.AddLinef("value.resize_with(nat_n as usize, %s::default);", elementTypeString)
			cc.For("element", "value.iter_mut()", "", func() {
				trw.element.t.TypeReadingCode(cc, bytesVersion, directImports, "*element", trw.element.Bare(), trw.wr.formatNatArgs(nil, trw.element.NatArgs()), false)
			})
		default:
			cc.For("element", "value.iter_mut()", "", func() {
				trw.element.t.TypeReadingCode(cc, bytesVersion, directImports, "*element", trw.element.Bare(), trw.wr.formatNatArgs(nil, trw.element.NatArgs()), false)
			})
		}
		cc.AddLinef("Ok(())")
	})
	cc.AddLinef("}")
	return cc.Text()
}
