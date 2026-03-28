// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package genrust

import "fmt"

func (trw *TypeRWPrimitive) calculateLayoutCall(
	directImports *DirectImports,
	bytesVersion bool,
	targetSizes string,
	targetObject string,
	zeroIfEmpty bool,
	refObject bool,
) string {
	if trw.canonicalType == "string" {
		sz := fmt.Sprintf("currentSize += basictl.TL2CalculateSize(len(%[1]s)) + len(%[1]s)", addAsterisk(refObject, targetObject))
		if zeroIfEmpty {
			return fmt.Sprintf("if len(%s) != 0 {\n", addAsterisk(refObject, targetObject)) + sz
		}
		return sz
	}
	if trw.canonicalType == "__function" || trw.canonicalType == "__function_result" {
		//sz := "currentSize += 0"
		if zeroIfEmpty {
			return "{\n"
		}
		return ""
	}
	sz := fmt.Sprintf("currentSize += %d", trw.trivialSize())
	if zeroIfEmpty {
		return fmt.Sprintf("if %s != 0 {\n", addAsterisk(refObject, targetObject)) + sz
	}
	return sz
}

func (trw *TypeRWPrimitive) writeTL2Call(
	directImports *DirectImports,
	bytesVersion bool,
	targetSizes string,
	targetBytes string,
	targetObject string,
	zeroIfEmpty bool,

	refObject bool,
) string {
	method := trw.writeValue
	if trw.canonicalType == "string" {
		method = addBytes("basictl.StringWriteTL2", bytesVersion)
	}
	sz := fmt.Sprintf(`%[2]s = %[1]s(%[2]s, %[3]s)`,
		method,
		targetBytes,
		addAsterisk(refObject, targetObject),
	)
	if trw.canonicalType == "string" {
		if zeroIfEmpty {
			return fmt.Sprintf("if len(%s) != 0 {\n", addAsterisk(refObject, targetObject)) + sz
		}
		return sz
	}
	if trw.canonicalType == "__function" || trw.canonicalType == "__function_result" {
		if zeroIfEmpty {
			return "{\n"
		}
		return ""
	}
	if zeroIfEmpty {
		return fmt.Sprintf("if %s != 0 {\n", addAsterisk(refObject, targetObject)) + sz
	}
	return sz
}

func (trw *TypeRWPrimitive) readTL2Call(
	directImports *DirectImports,
	bytesVersion bool,
	targetBytes string,
	targetObject string,
	canDependOnLocalBit bool,

	refObject bool,
) string {
	if trw.canonicalType == "__function" || trw.canonicalType == "__function_result" {
		return ""
	}
	method := ""
	switch trw.canonicalType {
	case "byte":
		method = "basictl.ByteRead"
	case "int32":
		method = "basictl.IntRead"
	case "uint32":
		method = "basictl.NatRead"
	case "int64":
		method = "basictl.LongRead"
	case "uint64":
		method = "basictl.Uint64Read"
	case "string":
		if bytesVersion {
			method = "basictl.StringReadTL2Bytes"
		} else {
			method = "basictl.StringReadTL2"
		}
	case "float32":
		method = "basictl.FloatRead"
	case "float64":
		method = "basictl.DoubleRead"
	}
	return fmt.Sprintf(`if %[2]s, err = %[1]s(%[2]s, %[3]s); err != nil { return %[2]s, err }`,
		method,
		targetBytes,
		addAmpersand(refObject, targetObject),
	)
}

func (trw *TypeRWPrimitive) skipTL2Call(
	directImports *DirectImports,
	bytesVersion bool,
	targetBytes string,
	canDependOnLocalBit bool,

	refObject bool,
) string {
	if trw.canonicalType == "__function" || trw.canonicalType == "__function_result" {
		return ""
	}
	size := 0
	switch trw.canonicalType {
	case "byte":
		size = 1
	case "int32", "uint32", "float32":
		size = 4
	case "int64", "uint64", "float64":
		size = 8
	case "string":
		return fmt.Sprintf(`if %[2]s, err = basictl.SkipSizedValue(%[2]s); err != nil { return %[2]s, err }`,
			"",
			targetBytes)
	}
	return fmt.Sprintf(`if %[2]s, err = basictl.SkipFixedSizedValue(%[2]s, %[3]d); err != nil { return %[2]s, err }`,
		"",
		targetBytes,
		size)
}

func (trw *TypeRWPrimitive) trivialSize() int {
	switch trw.canonicalType {
	case "byte":
		return 1
	case "int32", "uint32", "float32":
		return 4
	case "int64", "uint64", "float64":
		return 8
	}
	return 0
}
