package tlcodegen

import "fmt"

func (trw *TypeRWUnion) calculateLayout(
	bytesVersion bool,
	targetSizes string,
	targetObject string,
	canDependOnLocalBit bool,
	ins *InternalNamespace,
	refObject bool,
	natArgs []string) string {
	return fmt.Sprintf("%[1]s = %[2]s.CalculateLayout(%[1]s%[3]s)", targetSizes, addAsterisk(refObject, targetObject), joinWithCommas(natArgs))
}

func (trw *TypeRWUnion) writeTL2Call(
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
		addAsterisk(refObject, targetObject),
		joinWithCommas(natArgs),
		targetBytes,
	)
}

func (trw *TypeRWUnion) readTL2Call(
	bytesVersion bool,
	targetBytes string,
	targetObject string,
	canDependOnLocalBit bool,
	ins *InternalNamespace,
	refObject bool,
	natArgs []string,
) string {
	return fmt.Sprintf("if %[3]s, err = %[1]s.ReadTL2(%[3]s, %[1]s%[2]s); err != nil { return %[3]s, err }",
		addAsterisk(refObject, targetObject),
		joinWithCommas(natArgs),
		targetBytes,
	)
}

func (trw *TypeRWUnion) doesCalculateLayoutUseObject() bool {
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
