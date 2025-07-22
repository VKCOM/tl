package tlcodegen

import "fmt"

func (trw *TypeRWUnion) calculateLayoutCall(
	directImports *DirectImports,
	bytesVersion bool,
	targetSizes string,
	targetObject string,
	canDependOnLocalBit bool,
	ins *InternalNamespace,
	refObject bool,
) string {
	return fmt.Sprintf("%[1]s = %[2]s.CalculateLayout(%[1]s)", targetSizes, addAsteriskAndBrackets(refObject, targetObject))
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
) string {
	return fmt.Sprintf("%[3]s, %[1]s = %[2]s.InternalWriteTL2(%[3]s, %[1]s)",
		targetSizes,
		targetObject,
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
) string {
	return fmt.Sprintf("if %[2]s, err = %[1]s.InternalReadTL2(%[2]s); err != nil { return %[2]s, err }",
		targetObject,
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
