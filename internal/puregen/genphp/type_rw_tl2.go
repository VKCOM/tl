package genphp

type TypeRWTL2 interface {
	doesCalculateLayoutUseObject(allowInplace bool) bool
	isSizeWrittenInData() bool
	doesZeroSizeMeanEmpty(canDependOnLocalBit bool) bool

	doesWriteTL2UseObject(canDependOnLocalBit bool) bool
	tl2TrivialSize(targetObject string, canDependOnLocalBit bool, refObject bool) (isConstant bool, size string)

	doesReadTL2UseObject(canDependOnLocalBit bool) bool
	doesReadTL2UseBytes(canDependOnLocalBit bool) bool
}
