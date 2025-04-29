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
		size = fmt.Sprintf("basictl.TL2CalculateSize(len(%[1]s)) + len(%[1]s)", addAsterisk(refObject, targetObject))
	case "float32":
		size = "4"
	case "float64":
		size = "8"
	}
	return fmt.Sprintf("%[1]s = append(%[1]s, %[2]s)", targetSizes, size)
}

func (trw *TypeRWPrimitive) doesCalculateLayoutUseObject() bool {
	return trw.goType == "string"
}

func (trw *TypeRWPrimitive) isSizeWrittenInData() bool {
	return trw.goType == "string"
}
