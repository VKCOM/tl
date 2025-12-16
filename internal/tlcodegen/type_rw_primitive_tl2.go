package tlcodegen

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
		sz := fmt.Sprintf("sz = 1 + basictl.TL2CalculateSize(%s)", addAsterisk(refObject, targetObject))
		if zeroIfEmpty {
			sz += fmt.Sprintf("; len(%s) != 0", addAsterisk(refObject, targetObject))
		}
		return sz
	}
	sz := fmt.Sprintf("sz = %d", trw.trivialSize())
	if zeroIfEmpty {
		sz += fmt.Sprintf("; %s != 0", addAsterisk(refObject, targetObject))
	}
	return sz
}

func (trw *TypeRWPrimitive) writeTL2Call(
	directImports *DirectImports,
	bytesVersion bool,
	targetSizes string,
	targetBytes string,
	targetObject string,
	canDependOnLocalBit bool,
	ins *InternalNamespace,
	refObject bool,
) string {
	method := ""
	switch trw.goType {
	case "int32":
		method = "basictl.IntWrite"
	case "uint32":
		method = "basictl.NatWrite"
	case "int64":
		method = "basictl.LongWrite"
	case "string":
		if bytesVersion {
			method = "basictl.StringBytesWriteTL2"
		} else {
			method = "basictl.StringWriteTL2"
		}
	case "float32":
		method = "basictl.FloatWrite"
	case "float64":
		method = "basictl.DoubleWrite"
	}
	return fmt.Sprintf(`%[3]s = %[2]s(%[3]s, %[4]s)`,
		targetSizes,
		method,
		targetBytes,
		addAsterisk(refObject, targetObject),
	)
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
			method = "basictl.StringReadBytesTL2"
		} else {
			method = "basictl.StringReadTL2"
		}
	case "float32":
		method = "basictl.FloatRead"
	case "float64":
		method = "basictl.DoubleRead"
	}
	return fmt.Sprintf(`if %[3]s, err = %[2]s(%[3]s, %[4]s); err != nil { return %[3]s, err }`,
		"",
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

func (trw *TypeRWPrimitive) doesZeroSizeMeanEmpty(canDependOnLocalBit bool) bool {
	return true
}

func (trw *TypeRWPrimitive) doesCalculateLayoutUseObject(allowInplace bool) bool {
	if allowInplace {
		return false
	}
	return trw.goType == "string"
}

func (trw *TypeRWPrimitive) isSizeWrittenInData() bool {
	return trw.goType == "string"
}

func (trw *TypeRWPrimitive) doesWriteTL2UseObject(canDependOnLocalBit bool) bool {
	return true
}

func (trw *TypeRWPrimitive) doesReadTL2UseObject(canDependOnLocalBit bool) bool {
	return true
}

func (trw *TypeRWPrimitive) doesReadTL2UseBytes(canDependOnLocalBit bool) bool {
	return true
}

func (trw *TypeRWPrimitive) tl2TrivialSize(targetObject string, canDependOnLocalBit bool, refObject bool) (isConstant bool, size string) {
	size = "4"
	isConstant = true
	switch trw.goType {
	case "int32", "uint32":
		size = "4"
	case "int64":
		size = "8"
	case "string":
		size = fmt.Sprintf("len(%[1]s)", addAsterisk(refObject, targetObject))
		isConstant = false
	case "float32":
		size = "4"
	case "float64":
		size = "8"
	}
	return isConstant, size
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
