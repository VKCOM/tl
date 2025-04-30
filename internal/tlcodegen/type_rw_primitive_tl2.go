package tlcodegen

import "fmt"

func (trw *TypeRWPrimitive) calculateLayout(
	bytesVersion bool,
	targetSizes string,
	targetObject string,
	canDependOnLocalBit bool,
	ins *InternalNamespace,
	refObject bool,
	natArgs []string) string {
	size := "4"
	switch trw.goType {
	case "int32", "uint32":
		size = "4"
	case "int64":
		size = "8"
	case "string":
		size = fmt.Sprintf("len(%[1]s)", addAsterisk(refObject, targetObject))
	case "float32":
		size = "4"
	case "float64":
		size = "8"
	}
	return fmt.Sprintf("%[1]s = append(%[1]s, %[2]s)", targetSizes, size)
}

func (trw *TypeRWPrimitive) writeTL2Call(
	bytesVersion bool,
	targetSizes string,
	targetBytes string,
	targetObject string,
	canDependOnLocalBit bool,
	ins *InternalNamespace,
	refObject bool,
	natArgs []string) string {
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
	return fmt.Sprintf(`%[1]s = %[1]s[1:]
%[3]s = %[2]s(%[3]s, %[4]s)`,
		targetSizes,
		method,
		targetBytes,
		addAsterisk(refObject, targetObject),
	)
}

func (trw *TypeRWPrimitive) readTL2Call(
	bytesVersion bool,
	targetBytes string,
	targetObject string,
	canDependOnLocalBit bool,
	ins *InternalNamespace,
	refObject bool,
	natArgs []string,
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
			method = "basictl.StringBytesRead"
		} else {
			method = "basictl.StringRead"
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
		addAsterisk(refObject, targetObject),
	)
}

func (trw *TypeRWPrimitive) doesZeroSizeMeanEmpty(canDependOnLocalBit bool) bool {
	return true
}

func (trw *TypeRWPrimitive) doesCalculateLayoutUseObject() bool {
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
