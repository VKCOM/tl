package genphp

import "fmt"

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
