package tlcodegen

import "fmt"

func (trw *TypeRWMaybe) calculateLayout(
	bytesVersion bool,
	targetSizes string,
	targetObject string,
	canDependOnLocalBit bool,
	ins *InternalNamespace,
	refObject bool,
	natArgs []string) string {
	return fmt.Sprintf("%[1]s = %[2]s.CalculateLayout(%[1]s%[3]s)", targetSizes, addAsterisk(refObject, targetObject), joinWithCommas(natArgs))
}

func (trw *TypeRWMaybe) writeTL2Call(
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

func (trw *TypeRWMaybe) readTL2Call(
	bytesVersion bool,
	targetBytes string,
	targetObject string,
	canDependOnLocalBit bool,
	ins *InternalNamespace,
	refObject bool,
	natArgs []string,
) string {
	return fmt.Sprintf("if %[4]s, err = %[2]s.ReadTL2(%[4]s%[3]s); err != nil { return %[4]s, err }",
		"",
		addAsterisk(refObject, targetObject),
		joinWithCommas(natArgs),
		targetBytes,
	)
}

func (trw *TypeRWMaybe) doesZeroSizeMeanEmpty(canDependOnLocalBit bool) bool {
	return true
}

func (trw *TypeRWMaybe) doesCalculateLayoutUseObject() bool {
	return true
}

func (trw *TypeRWMaybe) isSizeWrittenInData() bool {
	return trw.element.t.trw.isSizeWrittenInData()
}

func (trw *TypeRWMaybe) doesWriteTL2UseObject(canDependOnLocalBit bool) bool {
	return true
}

func (trw *TypeRWMaybe) doesReadTL2UseObject(canDependOnLocalBit bool) bool {
	return true
}
