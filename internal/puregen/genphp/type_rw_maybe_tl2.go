package genphp

func (trw *TypeRWMaybe) doesZeroSizeMeanEmpty(canDependOnLocalBit bool) bool {
	return true
}

func (trw *TypeRWMaybe) doesCalculateLayoutUseObject(allowInplace bool) bool {
	return true
}

func (trw *TypeRWMaybe) isSizeWrittenInData() bool {
	return true
}

func (trw *TypeRWMaybe) doesWriteTL2UseObject(canDependOnLocalBit bool) bool {
	return true
}

func (trw *TypeRWMaybe) doesReadTL2UseObject(canDependOnLocalBit bool) bool {
	return true
}

func (trw *TypeRWMaybe) doesReadTL2UseBytes(canDependOnLocalBit bool) bool {
	return true
}

func (trw *TypeRWMaybe) tl2TrivialSize(targetObject string, canDependOnLocalBit bool, refObject bool) (isConstant bool, size string) {
	return false, ""
}
