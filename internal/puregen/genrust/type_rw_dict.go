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

type TypeRWDict struct {
	wr *TypeRWWrapper

	structElement *TypeRWStruct
	element       Field

	dictKeyString  bool
	dictKeyField   Field
	dictValueField Field
}

var _ TypeRW = &TypeRWDict{}

//lint:ignore U1000 will be used later
func (trw *TypeRWDict) formatValueNatArgs() []string {
	return trw.structElement.wr.formatNatArgs(nil, trw.dictValueField.NatArgs())
}

func (trw *TypeRWDict) typeString2(bytesVersion bool, directImports *DirectImports, isLocal bool, skipAlias bool) string {
	//if !bytesVersion {
	//	return fmt.Sprintf("map[%s]%s",
	//		trw.dictKeyField.t.TypeString2(bytesVersion, directImports, isLocal, skipAlias),
	//		trw.dictValueField.t.TypeString2(bytesVersion, directImports, isLocal, skipAlias))
	//}
	return fmt.Sprintf("Vec<%s>", trw.element.t.TypeString2(bytesVersion, directImports, isLocal, skipAlias))
}

func (trw *TypeRWDict) markHasBytesVersion(visitedNodes map[*TypeRWWrapper]bool) bool {
	return true
}

func (trw *TypeRWDict) markHasRepairMasks(visitedNodes map[*TypeRWWrapper]bool) bool {
	return trw.element.t.MarkHasRepairMasks(visitedNodes)
}

func (trw *TypeRWDict) markWriteHasError(visitedNodes map[*TypeRWWrapper]bool) bool {
	return trw.element.t.MarkWriteHasError(visitedNodes)
}

func (trw *TypeRWDict) markWantsBytesVersion(visitedNodes map[*TypeRWWrapper]bool) {
	trw.element.t.MarkWantsBytesVersion(visitedNodes)
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

func (trw *TypeRWDict) typeResettingCode(cc *codecreator.RustCodeCreator, bytesVersion bool, directImports *DirectImports, val string, ref bool) {
	cc.AddLinef("%s.clear();", val)
}

func (trw *TypeRWDict) typeRandomCode(bytesVersion bool, directImports *DirectImports, val string, natArgs []string, ref bool) string {
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
	return fmt.Sprintf("%sFillRandom(rg, %s%s)", goGlobalName, addAmpersand(ref, val), joinWithCommas(natArgs))
}

func (trw *TypeRWDict) typeRepairMasksCode(bytesVersion bool, directImports *DirectImports, val string, natArgs []string, ref bool) string {
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
	return fmt.Sprintf("%sRepairMasks(%s%s)", goGlobalName, addAmpersand(ref, val), joinWithCommas(natArgs))
}

func (trw *TypeRWDict) typeWritingCode(bytesVersion bool, directImports *DirectImports, val string, bare bool, natArgs []string, ref bool, last bool, needError bool) string {
	//prefix := ""
	//if !bare {
	//	prefix = fmt.Sprintf("w = basictl.NatWrite(w, 0x%x)\n", trw.wr.tlTag)
	//}
	refVal := addAsterisk(ref, val)
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
	//return prefix + wrapLastW(last, fmt.Sprintf("%sWrite(w, %s%s)", goGlobalName, refVal, joinWithCommas(natArgs)), needError)
	return wrapLastW(last, fmt.Sprintf("%sWriteTL1%s(w, %s%s)", goGlobalName, addBare(bare), refVal, joinWithCommas(natArgs)), needError)
}

func (trw *TypeRWDict) typeReadingCode(cc *codecreator.RustCodeCreator, bytesVersion bool, directImports *DirectImports, val string, bare bool, natArgs []string, ref bool) {
	cc.AddLinef("crate::types::%s::read_tl1(&mut %s, buf%s%s)?;", trw.wr.goGlobalName, addAsterisk(ref, val), joinWithCommas(natArgs), trw.wr.fetcherCall())
}

func (trw *TypeRWDict) typeJSONEmptyCondition(bytesVersion bool, val string, ref bool) string {
	return fmt.Sprintf("len(%s) != 0", addAsterisk(ref, val))
}

func (trw *TypeRWDict) typeJSONWritingCode(bytesVersion bool, directImports *DirectImports, val string, natArgs []string, ref bool, needError bool) string {
	//refVal := addAmpersand(ref, val)
	//if !bytesVersion {
	refVal := addAsterisk(ref, val) // those version pass to Write method by pointer
	//}
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
	// Code which depends on serialization location (skipping empty array if object property) is generated in that location.
	if needError {
		return fmt.Sprintf("if w, err = %sWriteJSONOpt(jctx, w, %s%s); err != nil { return w, err }", goGlobalName, refVal, joinWithCommas(natArgs))
	} else {
		return fmt.Sprintf("w = %sWriteJSONOpt(jctx, w, %s%s)", goGlobalName, refVal, joinWithCommas(natArgs))
	}
}

func (trw *TypeRWDict) typeJSON2ReadingCode(bytesVersion bool, directImports *DirectImports, jvalue string, val string, natArgs []string, ref bool) string {
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
	return fmt.Sprintf("if err := %sReadJSONGeneral(jctx, %s, %s%s); err != nil { return err }", goGlobalName, jvalue, addAmpersand(ref, val), joinWithCommas(natArgs))
}

func (trw *TypeRWDict) GenerateCode(bytesVersion bool, directImports *DirectImports) string {
	cc := codecreator.NewRustCodeCreator()

	//goName := addBytes(trw.wr.goGlobalName, bytesVersion)
	natDecl := trw.wr.formatNatArgsDecl()
	//natCall := trw.wr.formatNatArgsDeclCall()
	typeString := trw.wr.TypeString2(bytesVersion, directImports, false, false)
	//elementTypeString := trw.element.t.TypeString2(bytesVersion, directImports, false, false)

	cc.AddLinef("use basictl::TLRead as _;")
	cc.AddEmptyLine()

	cc.AddLinef("pub(crate) fn read_tl1<B: bytes::Buf + Copy>(value: &mut %s, buf: &mut B%s) -> basictl::Result<()> {", typeString, natDecl)
	cc.AddBlock(func() {
		if trw.wr.OriginTL2() {
			cc.AddLinef(`Err(basictl::Error::NoTL1("%s")`, trw.wr.pureType.CanonicalName())
			return
		}
		cc.AddLinef("let l = buf.read_u32()?;")
		// TODO - use length sanity check
		cc.AddLinef("value.resize_with(l as usize, Default::default);")
		cc.For("element", "value.iter_mut()", "", func() {
			trw.element.t.TypeReadingCode(cc, bytesVersion, directImports, "*element", trw.element.Bare(), trw.wr.formatNatArgs(nil, trw.element.NatArgs()), false)
		})
		cc.AddLinef("Ok(())")
	})
	cc.AddLinef("}")
	return cc.Text()
}
