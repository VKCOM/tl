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

type TypeRWBool struct {
	wr          *TypeRWWrapper
	falseGoName string
	trueGoName  string
	falseTag    uint32
	trueTag     uint32

	isBit bool
}

var _ TypeRW = &TypeRWBool{}

func (trw *TypeRWBool) typeString2(bytesVersion bool, directImports *DirectImports, isLocal bool, skipAlias bool) string {
	if !skipAlias {
		return "bool"
	}
	if isLocal {
		return addBytes(trw.wr.goLocalName, bytesVersion)
	}
	return addBytes(trw.wr.goGlobalName, bytesVersion)
}

func (trw *TypeRWBool) markHasBytesVersion(visitedNodes map[*TypeRWWrapper]bool) bool {
	return false
}

func (trw *TypeRWBool) markHasRepairMasks(visitedNodes map[*TypeRWWrapper]bool) bool {
	return false
}

func (trw *TypeRWBool) markWriteHasError(visitedNodes map[*TypeRWWrapper]bool) bool {
	return false
}

func (trw *TypeRWBool) markWantsBytesVersion(visitedNodes map[*TypeRWWrapper]bool) {
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

func (trw *TypeRWBool) typeResettingCode(cc *codecreator.RustCodeCreator, bytesVersion bool, directImports *DirectImports, val string, ref bool) {
	cc.AddLinef("%s = false", addAsterisk(ref, val))
}

func (trw *TypeRWBool) typeRandomCode(bytesVersion bool, directImports *DirectImports, val string, natArgs []string, ref bool) string {
	return fmt.Sprintf("%s = basictl.RandomUint(rg) & 1 == 1", addAsterisk(ref, val))
}

func (trw *TypeRWBool) typeRepairMasksCode(bytesVersion bool, directImports *DirectImports, val string, natArgs []string, ref bool) string {
	return ""
}

func (trw *TypeRWBool) typeWritingCode(bytesVersion bool, directImports *DirectImports, val string, bare bool, natArgs []string, ref bool, last bool, needError bool) string {
	return wrapLastW(last, fmt.Sprintf("%sWriteTL1%s(w, %s%s)", trw.wr.goGlobalName, addBare(bare), addAsterisk(ref, val), joinWithCommas(natArgs)), needError)
}

func (trw *TypeRWBool) typeReadingCode(cc *codecreator.RustCodeCreator, bytesVersion bool, directImports *DirectImports, val string, bare bool, natArgs []string, ref bool) {
	cc.AddLinef("%s = crate::types::%s::read_tl1(buf%s%s)?;", addAsterisk(ref, val), trw.wr.goGlobalName, joinWithCommas(natArgs), trw.wr.fetcherCall())
}

func (trw *TypeRWBool) typeJSONEmptyCondition(bytesVersion bool, val string, ref bool) string {
	return addAsterisk(ref, val)
}

func (trw *TypeRWBool) typeJSONWritingCode(bytesVersion bool, directImports *DirectImports, val string, natArgs []string, ref bool, needError bool) string {
	return fmt.Sprintf("w = basictl.JSONWriteBool(w, %s)", addAsterisk(ref, val))
}

func (trw *TypeRWBool) typeJSON2ReadingCode(bytesVersion bool, directImports *DirectImports, jvalue string, val string, natArgs []string, ref bool) string {
	return wrapLast(false, fmt.Sprintf("%sJson2ReadBool(%s, %s)", trw.wr.gen.InternalPrefix(), jvalue, addAmpersand(ref, val)))
}

func (trw *TypeRWBool) GenerateCode(bytesVersion bool, directImports *DirectImports) string {
	if trw.isBit {
		return ""
	}
	cc := codecreator.NewRustCodeCreator()

	cc.AddLinef("use basictl::TLRead as _;")
	cc.AddEmptyLine()
	cc.AddLinef("pub(crate) fn read_tl1<B: bytes::Buf + Copy>(buf: &mut B) -> basictl::Result<bool> {")
	cc.AddBlock(func() {
		if trw.wr.OriginTL2() {
			cc.AddLinef(`Err(basictl::Error::NoTL1("%s")`, trw.wr.pureType.CanonicalName())
			return
		}
		cc.AddLinef("match buf.read_u32()? {")
		cc.AddBlock(func() {
			cc.AddLinef("0x%08x => Ok(true),", trw.trueTag)
			cc.AddLinef("0x%08x => Ok(false),", trw.falseTag)
			cc.AddLinef("other => Err(basictl::Error::UnexpectedMagic(other)),")
		})
		cc.AddLinef("}")
	})
	cc.AddLinef("}")
	return cc.Text()
}
