// Copyright 2022 V Kontakte LLC
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
	primitiveType  string
	resetValue     string
	randomValue    string
	writeValue     string
	readValue      string
	writeJSONValue string
	readJSONValue  string
	writeHasError  bool
	isString       bool
	isFloat        bool

	cppFunctionSuffix string
	cppPrimitiveType  string
	cppDefaultInit    string
	cppResetValue     string
}

func (trw *TypeRWPrimitive) canBeBareOrBoxed(bare bool) bool {
	return bare
}

func (trw *TypeRWPrimitive) typeStringGlobal(bytesVersion bool) string {
	return addBytes(trw.primitiveType, bytesVersion)
}

func (trw *TypeRWPrimitive) typeString2(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, isLocal bool, skipAlias bool) string {
	if isLocal {
		return addBytes(trw.primitiveType, bytesVersion)
	}
	if bytesVersion {
		return "[]byte"
	}
	return trw.primitiveType
}

func (trw *TypeRWPrimitive) markHasBytesVersion(visitedNodes map[*TypeRWWrapper]bool) bool {
	return trw.isString
}

func (trw *TypeRWPrimitive) fillRecursiveUnwrap(visitedNodes map[*TypeRWWrapper]bool) {
}

func (trw *TypeRWPrimitive) markWantsBytesVersion(visitedNodes map[*TypeRWWrapper]bool) {
}

func (trw *TypeRWPrimitive) BeforeCodeGenerationStep() error {
	return nil
}

func (trw *TypeRWPrimitive) BeforeCodeGenerationStep2() {
}

func (trw *TypeRWPrimitive) fillRecursiveChildren(visitedNodes map[*TypeRWWrapper]bool) {
}

func (trw *TypeRWPrimitive) IsDictKeySafe() (isSafe bool, isString bool) {
	return !trw.isFloat, trw.isString
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
		return fmt.Sprintf("%s = %sBytes(rand)", addAsterisk(ref, val), trw.randomValue)
	}
	return fmt.Sprintf("%s = %s(rand)", addAsterisk(ref, val), trw.randomValue)
}

func (trw *TypeRWPrimitive) typeWritingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, bare bool, natArgs []string, ref bool, last bool) string {
	if !bare {
		log.Panicf("trw %q cannot be boxed", trw.primitiveType)
	}
	if trw.writeValue == "" {
		log.Panicf("trw %q cannot be bare", trw.primitiveType)
	}
	if bytesVersion {
		return wrapLastW(last, fmt.Sprintf("basictl.StringWriteBytes(w, %s )", addAsterisk(ref, val)))
	}
	code := fmt.Sprintf("%s(w, %s)", trw.writeValue, addAsterisk(ref, val))
	if trw.writeHasError {
		return wrapLastW(last, code)
	}
	return ifString(last, "return "+code+", nil", "w = "+code)
}

func (trw *TypeRWPrimitive) typeJSONEmptyCondition(bytesVersion bool, val string, ref bool) string {
	if trw.isString {
		return fmt.Sprintf("len(%s) != 0", addAsterisk(ref, val))
	}
	return fmt.Sprintf("%s != 0", addAsterisk(ref, val))
}

func (trw *TypeRWPrimitive) typeReadingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, bare bool, natArgs []string, ref bool, last bool) string {
	if !bare {
		log.Panicf("trw %q cannot be boxed", trw.primitiveType)
	}
	if bytesVersion {
		return wrapLastW(last, fmt.Sprintf("basictl.StringReadBytes(w, %s )", addAmpersand(ref, val)))
	}
	return wrapLastW(last, fmt.Sprintf("%s(w, %s)", trw.readValue, addAmpersand(ref, val)))
}

func (trw *TypeRWPrimitive) typeJSONWritingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, natArgs []string, ref bool) string {
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

func (trw *TypeRWPrimitive) GenerateCode(byteVersion bool, directImports *DirectImports) string {
	return ""
}
