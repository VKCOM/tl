// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package gengo

import "fmt"

func (trw *TypeRWPrimitive) calculateLayoutCall(
	directImports *DirectImports,
	bytesVersion bool,
	targetSizes string,
	targetObject string,
	zeroIfEmpty bool,
	ins *InternalNamespace,
	refObject bool,
) string {
	if trw.tlType == "string" {
		sz := fmt.Sprintf("currentSize += basictl.TL2CalculateSize(len(%[1]s)) + len(%[1]s)", addAsterisk(refObject, targetObject))
		if zeroIfEmpty {
			return fmt.Sprintf("if len(%s) != 0 {\n", addAsterisk(refObject, targetObject)) + sz
		}
		return sz
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
	ins *InternalNamespace,
	refObject bool,
) string {
	method := trw.writeValue
	if trw.tlType == "string" {
		method = addBytes("basictl.StringWriteTL2", bytesVersion)
	}
	sz := fmt.Sprintf(`%[2]s = %[1]s(%[2]s, %[3]s)`,
		method,
		targetBytes,
		addAsterisk(refObject, targetObject),
	)
	if trw.tlType == "string" {
		if zeroIfEmpty {
			return fmt.Sprintf("if len(%s) != 0 {\n", addAsterisk(refObject, targetObject)) + sz
		}
		return sz
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
	ins *InternalNamespace,
	refObject bool,
) string {
	method := ""
	switch trw.goType {
	case "int32":
		method = "basictl.IntRead"
	case "uint32":
		method = "basictl.NatRead"
	case "int64":
		method = "basictl.LongRead"
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
	ins *InternalNamespace,
	refObject bool,
) string {
	size := 0
	switch trw.goType {
	case "int32":
		size = 4
	case "uint32":
		size = 4
	case "int64":
		size = 8
	case "string":
		return fmt.Sprintf(`if %[2]s, err = basictl.SkipSizedValue(%[2]s); err != nil { return %[2]s, err }`,
			"",
			targetBytes)
	case "float32":
		size = 4
	case "float64":
		size = 8
	}
	return fmt.Sprintf(`if %[2]s, err = basictl.SkipFixedSizedValue(%[2]s, %[3]d); err != nil { return %[2]s, err }`,
		"",
		targetBytes,
		size)
}

func (trw *TypeRWPrimitive) doesReadTL2UseObject(canDependOnLocalBit bool) bool {
	return true
}

func (trw *TypeRWPrimitive) trivialSize() int {
	switch trw.goType {
	case "byte":
		return 1
	case "int32", "uint32", "float32":
		return 4
	case "int64", "float64":
		return 8
	}
	return 0
}
