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

func (trw *TypeRWPrimitive) CPPFillRecursiveChildren(visitedNodes map[*TypeRWWrapper]bool) {
}

func (trw *TypeRWPrimitive) cppTypeStringInNamespace(bytesVersion bool, hppInc *DirectIncludesCPP) string {
	return trw.cppPrimitiveType
}

func (trw *TypeRWPrimitive) cppTypeStringInNamespaceHalfResolved(bytesVersion bool, hppInc *DirectIncludesCPP, halfResolved HalfResolvedArgument) string {
	return trw.cppPrimitiveType
}

func (trw *TypeRWPrimitive) cppTypeStringInNamespaceHalfResolved2(bytesVersion bool, typeReduction EvaluatedType) string {
	return trw.cppPrimitiveType
}

func (trw *TypeRWPrimitive) cppDefaultInitializer(halfResolved HalfResolvedArgument, halfResolve bool) string {
	return trw.cppDefaultInit
}

func (trw *TypeRWPrimitive) CPPHasBytesVersion() bool {
	return false
}

func (trw *TypeRWPrimitive) CPPTypeResettingCode(bytesVersion bool, val string) string {
	return "\t" + fmt.Sprintf(trw.cppResetValue, val)
}

func (trw *TypeRWPrimitive) CPPTypeWritingCode(bytesVersion bool, val string, bare bool, natArgs []string, last bool) string {
	return fmt.Sprintf("\ts.%s_write(%s);", trw.cppFunctionSuffix, val)
}

func (trw *TypeRWPrimitive) CPPTypeReadingCode(bytesVersion bool, val string, bare bool, natArgs []string, last bool) string {
	return fmt.Sprintf("\ts.%s_read(%s);", trw.cppFunctionSuffix, val)
}

func (trw *TypeRWPrimitive) CPPGenerateCode(hpp *strings.Builder, hppInc *DirectIncludesCPP, hppIncFwd *DirectIncludesCPP, hppDet *strings.Builder, hppDetInc *DirectIncludesCPP, cppDet *strings.Builder, cppDetInc *DirectIncludesCPP, bytesVersion bool, forwardDeclaration bool) {
}
