// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlcodegen

import (
	"fmt"
	"log"
)

type TypeRWPrimitive struct {
	gen    *Gen2
	tlType string

	goType         string
	resetValue     string
	randomValue    string
	writeValue     string
	readValue      string
	writeJSONValue string
	readJSONValue  string
	readJSON2Value string
	writeHasError  bool // we keep this for future types

	cppFunctionSuffix string
	cppPrimitiveType  string
	cppDefaultInit    string
	cppResetValue     string
}

func (trw *TypeRWPrimitive) isFloat() bool {
	return trw.tlType == "float" || trw.tlType == "double"
}

func (trw *TypeRWPrimitive) typeString2(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, isLocal bool, skipAlias bool) string {
	if isLocal {
		return addBytes(trw.goType, bytesVersion)
	}
	if bytesVersion {
		return "[]byte"
	}
	return trw.goType
}

func (trw *TypeRWPrimitive) markHasBytesVersion(visitedNodes map[*TypeRWWrapper]bool) bool {
	return trw.tlType == "string"
}

func (trw *TypeRWPrimitive) markWriteHasError(visitedNodes map[*TypeRWWrapper]bool) bool {
	return false
}

func (trw *TypeRWPrimitive) fillRecursiveUnwrap(visitedNodes map[*TypeRWWrapper]bool) {
}

func (trw *TypeRWPrimitive) markWantsBytesVersion(visitedNodes map[*TypeRWWrapper]bool) {
}

func (trw *TypeRWPrimitive) markWantsTL2(visitedNodes map[*TypeRWWrapper]bool) {
}

func (trw *TypeRWPrimitive) AllPossibleRecursionProducers() []*TypeRWWrapper {
	return nil
}

func (trw *TypeRWPrimitive) AllTypeDependencies(generic, countFunctions bool) []*TypeRWWrapper {
	return nil
}

func (trw *TypeRWPrimitive) IsWrappingType() bool {
	return true
}

func (trw *TypeRWPrimitive) ContainsUnion(visitedNodes map[*TypeRWWrapper]bool) bool {
	return false
}

func (trw *TypeRWPrimitive) FillRecursiveChildren(visitedNodes map[*TypeRWWrapper]int, generic bool) {
}

func (trw *TypeRWPrimitive) BeforeCodeGenerationStep1() {
}

func (trw *TypeRWPrimitive) BeforeCodeGenerationStep2() {
}

func (trw *TypeRWPrimitive) fillRecursiveChildren(visitedNodes map[*TypeRWWrapper]bool) {
}

func (trw *TypeRWPrimitive) IsDictKeySafe() (isSafe bool, isString bool) {
	return !trw.isFloat(), trw.tlType == "string"
}

func (trw *TypeRWPrimitive) CanBeBareBoxed() (canBare bool, canBoxed bool) {
	return true, false
}

func (trw *TypeRWPrimitive) typeResettingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, ref bool) string {
	if bytesVersion {
		if ref {
			return fmt.Sprintf("*%[1]s = (*%[1]s)[:0]", val)
		}
		return fmt.Sprintf("%[1]s = %[1]s[:0]", val)
	}
	return fmt.Sprintf(trw.resetValue, addAsterisk(ref, val))
}

func (trw *TypeRWPrimitive) typeRandomCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, natArgs []string, ref bool) string {
	if bytesVersion {
		return fmt.Sprintf("%s = %sBytes(rg)", addAsterisk(ref, val), trw.randomValue)
	}
	return fmt.Sprintf("%s = %s(rg)", addAsterisk(ref, val), trw.randomValue)
}

func (trw *TypeRWPrimitive) typeWritingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, bare bool, natArgs []string, ref bool, last bool, needError bool) string {
	if !bare {
		log.Panicf("trw %q cannot be boxed", trw.tlType)
	}
	code := fmt.Sprintf("%s(w, %s)", addBytes(trw.writeValue, bytesVersion), addAsterisk(ref, val))
	if trw.writeHasError {
		return wrapLastW(last, code, needError)
	}
	if needError {
		return ifString(last, "return "+code+", nil", "w = "+code)
	} else {
		return ifString(last, "return "+code, "w = "+code)
	}
}

func (trw *TypeRWPrimitive) typeJSONEmptyCondition(bytesVersion bool, val string, ref bool) string {
	if trw.tlType == "string" {
		return fmt.Sprintf("len(%s) != 0", addAsterisk(ref, val))
	}
	return fmt.Sprintf("%s != 0", addAsterisk(ref, val))
}

func (trw *TypeRWPrimitive) typeReadingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, bare bool, natArgs []string, ref bool, last bool) string {
	if !bare {
		log.Panicf("trw %q cannot be boxed", trw.tlType)
	}
	if bytesVersion {
		return wrapLastW(last, fmt.Sprintf("basictl.StringReadBytes(w, %s )", addAmpersand(ref, val)), true)
	}
	return wrapLastW(last, fmt.Sprintf("%s(w, %s)", trw.readValue, addAmpersand(ref, val)), true)
}

func (trw *TypeRWPrimitive) typeJSONWritingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, natArgs []string, ref bool, needError bool) string {
	writeJSONValue := trw.writeJSONValue
	if bytesVersion {
		writeJSONValue += "Bytes"
	}
	return fmt.Sprintf("w = %s(w, %s)", writeJSONValue, addAsterisk(ref, val))
}

func (trw *TypeRWPrimitive) typeJSONReadingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, jvalue string, val string, natArgs []string, ref bool) string {
	readJSONValue := trw.readJSONValue
	if bytesVersion {
		readJSONValue += "Bytes"
	}
	return wrapLast(false, fmt.Sprintf("%s(%s, %s)", readJSONValue, jvalue, addAmpersand(ref, val)))
}

func (trw *TypeRWPrimitive) typeJSON2ReadingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, jvalue string, val string, natArgs []string, ref bool) string {
	readJSONValue := trw.readJSON2Value
	if bytesVersion {
		readJSONValue += "Bytes"
	}
	return wrapLast(false, fmt.Sprintf("%s(%s, %s)", readJSONValue, jvalue, addAmpersand(ref, val)))
}

func (trw *TypeRWPrimitive) typeJSON2ReadingRequiresContext() bool {
	return false
}

func (trw *TypeRWPrimitive) GenerateCode(byteVersion bool, directImports *DirectImports) string {
	return ""
}
