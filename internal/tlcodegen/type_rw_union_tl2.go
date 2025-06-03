package tlcodegen

import "fmt"

func (trw *TypeRWUnion) calculateLayout(
	directImports *DirectImports,
	bytesVersion bool,
	targetSizes string,
	targetObject string,
	canDependOnLocalBit bool,
	ins *InternalNamespace,
	refObject bool,
	natArgs []string) string {
	return fmt.Sprintf("%[1]s = %[2]s.CalculateLayout(%[1]s%[3]s)", targetSizes, addAsteriskAndBrackets(refObject, targetObject), joinWithCommas(natArgs))
}

func (trw *TypeRWUnion) writeTL2Call(
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
	return fmt.Sprintf("%[4]s, %[1]s = %[2]s.InternalWriteTL2(%[4]s, %[1]s%[3]s)",
		targetSizes,
		targetObject,
		joinWithCommas(natArgs),
		targetBytes,
	)
}

func (trw *TypeRWUnion) readTL2Call(
	directImports *DirectImports,
	bytesVersion bool,
	targetBytes string,
	targetObject string,
	canDependOnLocalBit bool,
	ins *InternalNamespace,
	refObject bool,
	natArgs []string,
) string {
	return fmt.Sprintf("if %[3]s, err = %[1]s.InternalReadTL2(%[3]s%[2]s); err != nil { return %[3]s, err }",
		targetObject,
		joinWithCommas(natArgs),
		targetBytes,
	)
}

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
