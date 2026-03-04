package tlcodegen

func (trw *TypeRWUnion) doesCalculateLayoutUseObject(allowInplace bool) bool {
	return true
}

func (trw *TypeRWUnion) doesZeroSizeMeanEmpty(canDependOnLocalBit bool) bool {
	return true
}

func (trw *TypeRWUnion) isSizeWrittenInData() bool {
	return true
}

func (trw *TypeRWUnion) doesWriteTL2UseObject(canDependOnLocalBit bool) bool {
	return true
}

func (trw *TypeRWUnion) doesReadTL2UseObject(canDependOnLocalBit bool) bool {
	return true
}

func (trw *TypeRWUnion) doesReadTL2UseBytes(canDependOnLocalBit bool) bool {
	return true
}

func (trw *TypeRWUnion) tl2TrivialSize(targetObject string, canDependOnLocalBit bool, refObject bool) (isConstant bool, size string) {
	return false, ""
}
