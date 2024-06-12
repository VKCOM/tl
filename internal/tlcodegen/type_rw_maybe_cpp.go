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

func (trw *TypeRWMaybe) CPPFillRecursiveChildren(visitedNodes map[*TypeRWWrapper]bool) {
	trw.element.t.CPPFillRecursiveChildren(visitedNodes)
}

func (trw *TypeRWMaybe) cppTypeStringInNamespace(bytesVersion bool, hppInc *DirectIncludesCPP) string {
	hppInc.ns[trw.wr.fileName] = CppIncludeInfo{componentId: trw.wr.typeComponent}
	return "::basictl::optional<" + trw.element.t.CPPTypeStringInNamespace(bytesVersion, hppInc) + ">"
}

func (trw *TypeRWMaybe) cppTypeStringInNamespaceHalfResolved(bytesVersion bool, hppInc *DirectIncludesCPP, halfResolved HalfResolvedArgument) string {
	hppInc.ns[trw.wr.fileName] = CppIncludeInfo{componentId: trw.wr.typeComponent}
	return "::basictl::optional<" + trw.element.t.CPPTypeStringInNamespaceHalfResolved(bytesVersion, hppInc, halfResolved.Args[0]) + ">"
}

func (trw *TypeRWMaybe) cppTypeStringInNamespaceHalfResolved2(bytesVersion bool, typeReduction EvaluatedType) string {
	return "::basictl::optional<" + trw.element.t.CPPTypeStringInNamespaceHalfResolved2(bytesVersion, typeReduction.Type.Arguments[0]) + ">"
}

func (trw *TypeRWMaybe) cppDefaultInitializer(halfResolved HalfResolvedArgument, halfResolve bool) string {
	return ""
}

func (trw *TypeRWMaybe) CPPHasBytesVersion() bool {
	return trw.element.t.trw.CPPHasBytesVersion()
}

func (trw *TypeRWMaybe) CPPTypeResettingCode(bytesVersion bool, val string) string {
	return fmt.Sprintf("\t%s.reset();", val)
}

func (trw *TypeRWMaybe) CPPTypeWritingCode(bytesVersion bool, val string, bare bool, natArgs []string, last bool) string {
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
	return fmt.Sprintf("\t::%s::%sWrite%s(s, %s%s);", trw.wr.gen.DetailsCPPNamespace, goGlobalName, addBare(bare), val, joinWithCommas(natArgs))
	//return wrapLast(last, fmt.Sprintf("\t%s.Write%s( w %s)", val, addBare(bare), formatNatArgsCallCPP(natArgs)))
}

func (trw *TypeRWMaybe) CPPTypeReadingCode(bytesVersion bool, val string, bare bool, natArgs []string, last bool) string {
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
	return fmt.Sprintf("\t::%s::%sRead%s(s, %s%s);", trw.wr.gen.DetailsCPPNamespace, goGlobalName, addBare(bare), val, joinWithCommas(natArgs))
	//return wrapLast(last, fmt.Sprintf("\t%s.Read%s( r %s)", val, addBare(bare), formatNatArgsCallCPP(natArgs)))
}

func (trw *TypeRWMaybe) CPPGenerateCode(hpp *strings.Builder, hppInc *DirectIncludesCPP, hppIncFwd *DirectIncludesCPP, hppDet *strings.Builder, hppDetInc *DirectIncludesCPP, cppDet *strings.Builder, cppDetInc *DirectIncludesCPP, bytesVersion bool, forwardDeclaration bool) {
	if forwardDeclaration {
		trw.element.t.trw.CPPGenerateCode(hpp, hppInc, hppIncFwd, hppDet, hppDetInc, cppDet, cppDetInc, bytesVersion, true)
		return
	}
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
	myFullType := trw.cppTypeStringInNamespace(bytesVersion, hppDetInc)

	cppStartNamespace(hppDet, trw.wr.gen.DetailsCPPNamespaceElements)

	hppDet.WriteString(fmt.Sprintf(`
void %[1]sReadBoxed(::basictl::tl_istream & s, %[2]s& item%[3]s);
void %[1]sWriteBoxed(::basictl::tl_ostream & s, const %[2]s& item%[3]s);

`, goGlobalName,
		myFullType,
		formatNatArgsDeclCPP(trw.wr.NatParams),
		formatNatArgsCallCPP(trw.wr.NatParams),
		trw.wr.tlTag))

	cppDet.WriteString(fmt.Sprintf(`
void %[6]s::%[1]sReadBoxed(::basictl::tl_istream & s, %[2]s& item%[3]s) {
	if (s.bool_read(0x%[4]x, 0x%[5]x)) {
		if (!item) {
			item.emplace();
		}
	%[7]s
		return;
	}
	item.reset();
}

void %[6]s::%[1]sWriteBoxed(::basictl::tl_ostream & s, const %[2]s& item%[3]s) {
	s.nat_write(item ? 0x%[5]x : 0x%[4]x);
	if (item) {
	%[8]s
	}
}
`,
		goGlobalName,
		myFullType,
		formatNatArgsDeclCPP(trw.wr.NatParams),
		trw.emptyTag,
		trw.okTag,
		trw.wr.gen.DetailsCPPNamespace,
		trw.element.t.trw.CPPTypeReadingCode(bytesVersion, "*item", trw.element.Bare(), formatNatArgs(nil, trw.element.natArgs), true),
		trw.element.t.trw.CPPTypeWritingCode(bytesVersion, "*item", trw.element.Bare(), formatNatArgs(nil, trw.element.natArgs), true),
	))

	cppFinishNamespace(hppDet, trw.wr.gen.DetailsCPPNamespaceElements)

	/*
			_ = fmt.Sprintf(`type %[1]s struct {
			Value %[2]s // Значение имеет смысл при Ok=true
			Ok    bool
		}

		func (item *%[1]s) Reset() {
			item.Ok = false
		}

		func (item *%[1]s) ReadBoxed(r *bytes.Buffer%[8]s) error {
			if err := readBool(r, &item.Ok, %#[6]x, %#[7]x); err != nil {
				return err
			}
			if item.Ok {
				%[3]s
			}
			%[5]s
			return nil
		}

		func (item *%[1]s) WriteBoxed(w *bytes.Buffer%[8]s) error {
			writeBool(w, item.Ok, %#[6]x, %#[7]x)
			if item.Ok {
				%[4]s
			}
			return nil
		}
		`,
				addBytes(trw.goGlobalName, bytesVersion),
				trw.element.t.TypeString(bytesVersion),
				trw.element.t.TypeReadingCode(bytesVersion, trw.wr.ins,
					"item.Value",
					trw.element.bare,
					formatNatArgsCPP(nil, trw.element.natArgs),
					false,
					true,
				),
				trw.element.t.TypeWritingCode(bytesVersion, trw.wr.ins,
					"item.Value",
					trw.element.bare,
					formatNatArgsCPP(nil, trw.element.natArgs),
					false,
					true,
				),
				trw.element.t.TypeResettingCode(bytesVersion, trw.wr.ins, "item.Value", false),
				trw.emptyTag,
				trw.okTag,
				formatNatArgsDeclCPP(trw.wr.NatParams),
			)
	*/
}
