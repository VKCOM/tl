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

func (trw *TypeRWBool) cppTypeStringInNamespace(bytesVersion bool, hppInc *DirectIncludesCPP) string {
	hppInc.ns[trw.wr] = CppIncludeInfo{componentId: trw.wr.typeComponent, namespace: trw.wr.tlName.Namespace}
	return "bool"
}

func (trw *TypeRWBool) cppTypeStringInNamespaceHalfResolved(bytesVersion bool, hppInc *DirectIncludesCPP, halfResolved HalfResolvedArgument) string {
	hppInc.ns[trw.wr] = CppIncludeInfo{componentId: trw.wr.typeComponent, namespace: trw.wr.tlName.Namespace}
	return "bool"
}

func (trw *TypeRWBool) cppTypeStringInNamespaceHalfResolved2(bytesVersion bool, typeReduction EvaluatedType) string {
	return "bool"
}

func (trw *TypeRWBool) cppDefaultInitializer(halfResolved HalfResolvedArgument, halfResolve bool) string {
	return " = false"
}

func (trw *TypeRWBool) CPPHasBytesVersion() bool {
	return false
}

func (trw *TypeRWBool) CPPTypeResettingCode(bytesVersion bool, val string) string {
	return fmt.Sprintf("\t%s = false;", val)
}

func (trw *TypeRWBool) CPPTypeWritingCode(bytesVersion bool, val string, bare bool, natArgs []string, last bool) string {
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
	return fmt.Sprintf("\tif (!::%s::%sWrite%s(s, %s%s)) { return false; }", trw.wr.gen.DetailsCPPNamespace, goGlobalName, addBare(bare), val, joinWithCommas(natArgs))
}

func (trw *TypeRWBool) CPPTypeWritingJsonCode(bytesVersion bool, val string, bare bool, natArgs []string, last bool) string {
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
	return fmt.Sprintf("\tif (!::%s::%sWriteJSON(s, %s%s)) { return false; }", trw.wr.gen.DetailsCPPNamespace, goGlobalName, val, joinWithCommas(natArgs))
}

func (trw *TypeRWBool) CPPTypeReadingCode(bytesVersion bool, val string, bare bool, natArgs []string, last bool) string {
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
	return fmt.Sprintf("\tif (!::%s::%sRead%s(s, %s%s)) { return false; }", trw.wr.gen.DetailsCPPNamespace, goGlobalName, addBare(bare), val, joinWithCommas(natArgs))
}

func (trw *TypeRWBool) CPPGenerateCode(hpp *strings.Builder, hppInc *DirectIncludesCPP, hppIncFwd *DirectIncludesCPP, hppDet *strings.Builder, hppDetInc *DirectIncludesCPP, cppDet *strings.Builder, cppDetInc *DirectIncludesCPP, bytesVersion bool, forwardDeclaration bool) {
	typeNamespace := trw.wr.gen.RootCPPNamespaceElements
	if trw.wr.tlName.Namespace != "" {
		typeNamespace = append(typeNamespace, trw.wr.tlName.Namespace)
	}
	if hpp != nil {
		cppStartNamespace(hpp, typeNamespace)
		// TODO - better names of enums
		hpp.WriteString(fmt.Sprintf(`
	enum { %[4]s = 0x%[2]x, %[5]s = 0x%[3]x };
`, addBytes(trw.wr.goGlobalName, bytesVersion), trw.falseTag, trw.trueTag, trw.falseGoName, trw.trueGoName))
		cppFinishNamespace(hpp, typeNamespace)
	}
	if hppDet != nil {
		cppStartNamespace(hppDet, trw.wr.gen.DetailsCPPNamespaceElements)

		hppDet.WriteString(fmt.Sprintf(`
bool %[1]sWriteJSON(std::ostream & s, bool item);
bool %[1]sReadBoxed(::basictl::tl_istream & s, bool& item);
bool %[1]sWriteBoxed(::basictl::tl_ostream & s, bool item);
`, addBytes(trw.wr.goGlobalName, bytesVersion)))

		cppFinishNamespace(hppDet, trw.wr.gen.DetailsCPPNamespaceElements)
	}

	if cppDet != nil {
		cppDet.WriteString(fmt.Sprintf(`
bool %[6]s::%[1]sWriteJSON(std::ostream & s, bool item) {
	if (item) {
		s << "true";
	} else {
		s << "false";
	}
	return true;
}

bool %[6]s::%[1]sReadBoxed(::basictl::tl_istream & s, bool& item) {
	return s.bool_read(item, 0x%[2]x, 0x%[3]x);
}

bool %[6]s::%[1]sWriteBoxed(::basictl::tl_ostream & s, bool item) {
	return s.nat_write(item ? 0x%[3]x : 0x%[2]x);
}
`, addBytes(trw.wr.goGlobalName, bytesVersion), trw.falseTag, trw.trueTag, trw.falseGoName, trw.trueGoName, trw.wr.gen.DetailsCPPNamespace))

	}
}
