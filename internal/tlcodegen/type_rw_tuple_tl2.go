package tlcodegen

import "fmt"

func (trw *TypeRWBrackets) calculateLayout(
	bytesVersion bool,
	targetSizes string,
	targetObject string,
	canDependOnLocalBit bool,
	ins *InternalNamespace,
	refObject bool,
	natArgs []string) string {
	return fmt.Sprintf("%[1]s = %[5]s%[4]sCalculateLayout(%[1]s, %[2]s%[3]s)", targetSizes, addAmpersand(refObject, targetObject), joinWithCommas(natArgs), addBytes(trw.wr.goGlobalName, bytesVersion), trw.wr.ins.AddPrefix(ins))
}

func (trw *TypeRWBrackets) writeTL2Call(
	bytesVersion bool,
	targetSizes string,
	targetBytes string,
	targetObject string,
	canDependOnLocalBit bool,
	ins *InternalNamespace,
	refObject bool,
	natArgs []string,
) string {
	return fmt.Sprintf("%[6]s, %[1]s = %[5]s%[4]sInternalWriteTL2(%[6]s, %[1]s, %[2]s%[3]s)",
		targetSizes,
		addAmpersand(refObject, targetObject),
		joinWithCommas(natArgs),
		addBytes(trw.wr.goGlobalName, bytesVersion),
		trw.wr.ins.AddPrefix(ins),
		targetBytes,
	)
}

func (trw *TypeRWBrackets) readTL2Call(
	bytesVersion bool,
	targetBytes string,
	targetObject string,
	canDependOnLocalBit bool,
	ins *InternalNamespace,
	refObject bool,
	natArgs []string,
) string {
	return fmt.Sprintf("if %[6]s, err = %[5]s%[4]sReadTL2(%[6]s, %[2]s%[3]s); err != nil { return %[6]s, err }",
		"",
		addAmpersand(refObject, targetObject),
		joinWithCommas(natArgs),
		addBytes(trw.wr.goGlobalName, bytesVersion),
		trw.wr.ins.AddPrefix(ins),
		targetBytes,
	)
}

func (trw *TypeRWBrackets) doesZeroSizeMeanEmpty(canDependOnLocalBit bool) bool {
	return true
}

func (trw *TypeRWBrackets) doesCalculateLayoutUseObject() bool {
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

func (trw *TypeRWBrackets) tl2TrivialSize(targetObject string, canDependOnLocalBit bool, refObject bool) (isConstant bool, size string) {
	return false, ""
}
