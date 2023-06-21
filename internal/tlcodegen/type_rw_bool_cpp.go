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

func (trw *TypeRWBool) CPPFillRecursiveChildren(visitedNodes map[*TypeRWWrapper]bool) {
}

func (trw *TypeRWBool) cppTypeStringInNamespace(bytesVersion bool, hppInc *DirectIncludesCPP, resolvedType ResolvedType, halfResolve bool) string {
	hppInc.ns[trw.wr.fileName] = struct{}{}
	return "bool"
}

func (trw *TypeRWBool) cppDefaultInitializer(resolvedType ResolvedType, halfResolve bool) string {
	return " = false"
}

func (trw *TypeRWBool) CPPHasBytesVersion() bool {
	return false
}

func (trw *TypeRWBool) CPPTypeResettingCode(bytesVersion bool, val string) string {
	return fmt.Sprintf("\t%s = false;", val)
}

func (trw *TypeRWBool) CPPTypeWritingCode(bytesVersion bool, val string, bare bool, natArgs []string, last bool) string {
	goGlobalName := addBytes(trw.goGlobalName, bytesVersion)
	return fmt.Sprintf("\t::%s::%sWrite%s(s, %s%s);", trw.wr.gen.DetailsCPPNamespace, goGlobalName, addBare(bare), val, formatNatArgsCallCPP(natArgs))
}

func (trw *TypeRWBool) CPPTypeReadingCode(bytesVersion bool, val string, bare bool, natArgs []string, last bool) string {
	goGlobalName := addBytes(trw.goGlobalName, bytesVersion)
	return fmt.Sprintf("\t::%s::%sRead%s(s, %s%s);", trw.wr.gen.DetailsCPPNamespace, goGlobalName, addBare(bare), val, formatNatArgsCallCPP(natArgs))
}

func (trw *TypeRWBool) CPPGenerateCode(hpp *strings.Builder, hppInc *DirectIncludesCPP, hppIncFwd *DirectIncludesCPP, hppDet *strings.Builder, hppDetInc *DirectIncludesCPP, cppDet *strings.Builder, cppDetInc *DirectIncludesCPP, bytesVersion bool, forwardDeclaration bool) {
	typeNamespace := trw.wr.gen.RootCPPNamespaceElements
	if trw.wr.tlName.Namespace != "" {
		typeNamespace = append(typeNamespace, trw.wr.tlName.Namespace)
	}
	cppStartNamespace(hpp, typeNamespace)
	// TODO - better names of enums
	hpp.WriteString(fmt.Sprintf(`
	enum { %[4]s = 0x%[2]x, %[5]s = 0x%[3]x };
`, addBytes(trw.goGlobalName, bytesVersion), trw.falseTag, trw.trueTag, trw.falseGoName, trw.trueGoName))
	cppFinishNamespace(hpp, typeNamespace)

	cppStartNamespace(hppDet, trw.wr.gen.DetailsCPPNamespaceElements)

	hppDet.WriteString(fmt.Sprintf(`
void %[1]sReadBoxed(::basictl::tl_istream & s, bool& item);
void %[1]sWriteBoxed(::basictl::tl_ostream & s, bool item);
`, addBytes(trw.goGlobalName, bytesVersion)))

	cppDet.WriteString(fmt.Sprintf(`
void %[6]s::%[1]sReadBoxed(::basictl::tl_istream & s, bool& item) {
	item = s.bool_read(0x%[2]x, 0x%[3]x);
}

void %[6]s::%[1]sWriteBoxed(::basictl::tl_ostream & s, bool item) {
	s.nat_write(item ? 0x%[3]x : 0x%[2]x);
}
`, addBytes(trw.goGlobalName, bytesVersion), trw.falseTag, trw.trueTag, trw.falseGoName, trw.trueGoName, trw.wr.gen.DetailsCPPNamespace))

	cppFinishNamespace(hppDet, trw.wr.gen.DetailsCPPNamespaceElements)
}
