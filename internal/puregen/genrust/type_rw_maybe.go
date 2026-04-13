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
	elementTypeString := trw.element.t.TypeString2(bytesVersion, directImports, false, false)
	//if isLocal {
	//	return addBytes(trw.wr.goLocalName, bytesVersion)
	//}
	return fmt.Sprintf("Option<%s>", elementTypeString)
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
	cc.AddLinef("%s = None;", val)
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
	cc.AddLinef("crate::types::%s::read_tl1_boxed(&mut %s, buf%s%s)?;", trw.wr.goGlobalName, addAsterisk(ref, val), joinWithCommas(natArgs), trw.wr.fetcherCall())
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
	cc := codecreator.NewRustCodeCreator()

	//goName := addBytes(trw.wr.goGlobalName, bytesVersion)
	natDecl := trw.wr.formatNatArgsDecl()
	//natCall := trw.wr.formatNatArgsDeclCall()
	typeString := trw.wr.TypeString2(bytesVersion, directImports, false, false)
	//elementTypeString := trw.element.t.TypeString2(bytesVersion, directImports, false, false)

	cc.AddLinef("use basictl::TLRead as _;")
	cc.AddEmptyLine()

	cc.AddLinef("pub(crate) fn read_tl1_boxed<B: bytes::Buf + Copy>(value: &mut %s, buf: &mut B%s) -> basictl::Result<()> {", typeString, natDecl)
	cc.FinishBlock(func() {
		if trw.wr.OriginTL2() {
			cc.AddLinef(`Err(basictl::Error::NoTL1("%s")`, trw.wr.pureType.CanonicalName())
			return
		}
		cc.AddLinef("let ok = buf.read_bool(0x%08x, 0x%08x)?;", trw.emptyTag, trw.okTag)
		cc.IfElse("ok", func() {
			cc.If("value.is_none()", func() {
				cc.AddLinef("*value = Some(Default::default())")
			})
			cc.If("let Some(subValue) = value", func() {
				natArgs := trw.wr.formatNatArgs(nil, trw.element.NatArgs())
				trw.element.t.TypeReadingCode(cc, bytesVersion, directImports, "subValue", trw.element.Bare(), natArgs, true)
			})
		}, func() {
			cc.AddLinef("*value = None")
		})
		cc.AddLinef("Ok(())")
	}, "}")
	return cc.Text()
}
