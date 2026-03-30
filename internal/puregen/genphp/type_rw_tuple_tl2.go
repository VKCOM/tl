package genphp

func (trw *TypeRWBrackets) doesZeroSizeMeanEmpty(canDependOnLocalBit bool) bool {
	return true
}

func (trw *TypeRWBrackets) doesCalculateLayoutUseObject(allowInplace bool) bool {
	return true
}

func (trw *TypeRWBrackets) isSizeWrittenInData() bool {
	return true
}

func (trw *TypeRWBrackets) doesWriteTL2UseObject(canDependOnLocalBit bool) bool {
	return true
}

func (trw *TypeRWBrackets) doesReadTL2UseObject(canDependOnLocalBit bool) bool {
	return true
}

func (trw *TypeRWBrackets) doesReadTL2UseBytes(canDependOnLocalBit bool) bool {
	return true
}

func (trw *TypeRWBrackets) tl2TrivialSize(targetObject string, canDependOnLocalBit bool, refObject bool) (isConstant bool, size string) {
	return false, ""
}
