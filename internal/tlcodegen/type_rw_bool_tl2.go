package tlcodegen

/**
BOOL = BYTE or BIT in STRUCT
*/

func (trw *TypeRWBool) doesZeroSizeMeanEmpty(canDependOnLocalBit bool) bool {
	return !(trw.isBit && canDependOnLocalBit)
}

func (trw *TypeRWBool) doesCalculateLayoutUseObject(allowInplace bool) bool {
	return false
}

func (trw *TypeRWBool) isSizeWrittenInData() bool {
	return false
}

func (trw *TypeRWBool) doesWriteTL2UseObject(canDependOnLocalBit bool) bool {
	return !(trw.isBit && canDependOnLocalBit)
}

func (trw *TypeRWBool) doesReadTL2UseObject(canDependOnLocalBit bool) bool {
	return true
}

func (trw *TypeRWBool) doesReadTL2UseBytes(canDependOnLocalBit bool) bool {
	return !(trw.isBit && canDependOnLocalBit)
}

func (trw *TypeRWBool) tl2TrivialSize(targetObject string, canDependOnLocalBit bool, refObject bool) (isConstant bool, size string) {
	if trw.isBit && canDependOnLocalBit {
		return true, "0"
	}
	return true, "1"
}
