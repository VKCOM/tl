package tlcodegen

import "fmt"

func (trw *TypeRWBrackets) calculateLayoutCall(
	directImports *DirectImports,
	bytesVersion bool,
	targetSizes string,
	targetObject string,
	canDependOnLocalBit bool,
	ins *InternalNamespace,
	refObject bool,
) string {
	return fmt.Sprintf("%[1]s = %[4]s%[3]sCalculateLayout(%[1]s, %[2]s)",
		targetSizes,
		addAmpersand(refObject, targetObject),
		addBytes(trw.wr.goGlobalName, bytesVersion),
		trw.wr.ins.Prefix(directImports, ins),
	)
}

func (trw *TypeRWBrackets) writeTL2Call(
	directImports *DirectImports,
	bytesVersion bool,
	targetSizes string,
	targetBytes string,
	targetObject string,
	canDependOnLocalBit bool,
	ins *InternalNamespace,
	refObject bool,
) string {
	return fmt.Sprintf("%[5]s, %[1]s = %[4]s%[3]sInternalWriteTL2(%[5]s, %[1]s, %[2]s)",
		targetSizes,
		addAmpersand(refObject, targetObject),
		addBytes(trw.wr.goGlobalName, bytesVersion),
		trw.wr.ins.Prefix(directImports, ins),
		targetBytes,
	)
}

func (trw *TypeRWBrackets) readTL2Call(
	directImports *DirectImports,
	bytesVersion bool,
	targetBytes string,
	targetObject string,
	canDependOnLocalBit bool,
	ins *InternalNamespace,
	refObject bool,
) string {
	return fmt.Sprintf("if %[5]s, err = %[4]s%[3]sInternalReadTL2(%[5]s, %[2]s); err != nil { return %[5]s, err }",
		"",
		addAmpersand(refObject, targetObject),
		addBytes(trw.wr.goGlobalName, bytesVersion),
		trw.wr.ins.Prefix(directImports, ins),
		targetBytes,
	)
}

func (trw *TypeRWBrackets) skipTL2Call(
	directImports *DirectImports,
	bytesVersion bool,
	targetBytes string,
	canDependOnLocalBit bool,
	ins *InternalNamespace,
	refObject bool,
) string {
	return fmt.Sprintf(`if %[2]s, err = basictl.SkipSizedValue(%[2]s); err != nil { return %[2]s, err }`,
		"",
		targetBytes)
}

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
