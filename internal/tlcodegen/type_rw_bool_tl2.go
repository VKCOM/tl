package tlcodegen

import "fmt"

/**
BOOL = BYTE or BIT in STRUCT
*/

func (trw *TypeRWBool) calculateLayout(
	directImports *DirectImports,
	bytesVersion bool,
	targetSizes string,
	targetObject string,
	canDependOnLocalBit bool,
	ins *InternalNamespace,
	refObject bool,
	natArgs []string) string {
	return ""
	//if canDependOnLocalBit {
	//	return fmt.Sprintf("%[1]s = append(%[1]s, 0)", targetSizes)
	//}
	//return fmt.Sprintf("%[1]s = append(%[1]s, 1)", targetSizes)
}

func (trw *TypeRWBool) writeTL2Call(
	directImports *DirectImports,
	bytesVersion bool,
	targetSizes string,
	targetBytes string,
	targetObject string,
	canDependOnLocalBit bool,
	ins *InternalNamespace,
	refObject bool,
	natArgs []string,
) string {
	if canDependOnLocalBit {
		return "" // fmt.Sprintf("%[1]s = %[1]s[1:]", targetSizes)
	}
	return fmt.Sprintf(`if %[2]s {
	%[3]s = append(%[3]s, 1)
} else {
	%[3]s = append(%[3]s, 0)
}`,
		targetSizes,
		targetObject,
		targetBytes)
}

func (trw *TypeRWBool) readTL2Call(
	directImports *DirectImports,
	bytesVersion bool,
	targetBytes string,
	targetObject string,
	canDependOnLocalBit bool,
	ins *InternalNamespace,
	refObject bool,
	natArgs []string,
) string {
	if canDependOnLocalBit {
		return fmt.Sprintf("%[1]s = true", targetObject)
	}
	return fmt.Sprintf(`if %[2]s, err = basictl.BoolReadTL2(%[2]s, %[1]s); err != nil { return %[2]s, err }`,
		addAmpersand(refObject, targetObject),
		targetBytes)
}

func (trw *TypeRWBool) doesZeroSizeMeanEmpty(canDependOnLocalBit bool) bool {
	return !canDependOnLocalBit
}

func (trw *TypeRWBool) doesCalculateLayoutUseObject(allowInplace bool) bool {
	return false
}

func (trw *TypeRWBool) isSizeWrittenInData() bool {
	return false
}

func (trw *TypeRWBool) doesWriteTL2UseObject(canDependOnLocalBit bool) bool {
	return !canDependOnLocalBit
}

func (trw *TypeRWBool) doesReadTL2UseObject(canDependOnLocalBit bool) bool {
	return true
}

func (trw *TypeRWBool) doesReadTL2UseBytes(canDependOnLocalBit bool) bool {
	return !canDependOnLocalBit
}

func (trw *TypeRWBool) tl2TrivialSize(targetObject string, canDependOnLocalBit bool, refObject bool) (isConstant bool, size string) {
	if canDependOnLocalBit {
		return true, "0"
	}
	return true, "1"
}
