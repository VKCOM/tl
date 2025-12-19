package tlcodegen

import (
	"fmt"
)

func (trw *TypeRWStruct) calculateLayoutCall(
	directImports *DirectImports,
	bytesVersion bool,
	targetSizes string,
	targetObject string,
	zeroIfEmpty bool,
	ins *InternalNamespace,
	refObject bool,
) string {
	if trw.isUnwrapType() {
		return trw.Fields[0].t.CalculateLayoutCall(directImports, bytesVersion, targetSizes, targetObject, zeroIfEmpty, ins, refObject)
	}
	if trw.wr.IsTrueType() && trw.wr.unionParent == nil {
		if zeroIfEmpty {
			return "if false {"
		}
		return "currentSize += 1"
	}
	sz := fmt.Sprintf("%[1]s, sz = %[2]s.CalculateLayout(%[1]s, %[3]v)", targetSizes, addAsteriskAndBrackets(refObject, targetObject), zeroIfEmpty)
	if zeroIfEmpty {
		sz = fmt.Sprintf("if %s; sz != 0 {", sz)
	}
	return sz + "\ncurrentSize += sz"
}

func (trw *TypeRWStruct) writeTL2Call(
	directImports *DirectImports,
	bytesVersion bool,
	targetSizes string,
	targetBytes string,
	targetObject string,
	zeroIfEmpty bool,
	ins *InternalNamespace,
	refObject bool,
) string {
	if trw.isUnwrapType() {
		return trw.Fields[0].t.WriteTL2Call(directImports, bytesVersion, targetSizes, targetBytes, targetObject, zeroIfEmpty, ins, refObject)
	}
	if trw.wr.IsTrueType() && trw.wr.unionParent == nil {
		if zeroIfEmpty {
			return "if false {"
		}
		return fmt.Sprintf("%[1]s = append(%[1]s, 0)", targetBytes)
	}
	sz := fmt.Sprintf("%[3]s, %[1]s, sz = %[2]s.InternalWriteTL2(%[3]s, %[1]s, %[4]v)",
		targetSizes,
		targetObject,
		targetBytes,
		zeroIfEmpty,
	)
	if zeroIfEmpty {
		sz = fmt.Sprintf("if %s; sz != 0 {", sz)
	}
	return sz
}

func (trw *TypeRWStruct) readTL2Call(
	directImports *DirectImports,
	bytesVersion bool,
	targetBytes string,
	targetObject string,
	canDependOnLocalBit bool,
	ins *InternalNamespace,
	refObject bool,
) string {
	//if trw.wr.IsTrueType() && trw.wr.unionParent == nil {
	//	return ""
	//}
	if trw.isUnwrapType() {
		return trw.Fields[0].t.ReadTL2Call(directImports, bytesVersion, targetBytes, targetObject, canDependOnLocalBit, ins, refObject)
	}
	additionalSuffix := ""
	if trw.wr.unionParent != nil {
		additionalSuffix = ", block"
	}
	return fmt.Sprintf("if %[3]s, err = %[2]s.InternalReadTL2(%[3]s%[4]s); err != nil { return %[3]s, err }",
		"",
		targetObject,
		targetBytes,
		additionalSuffix,
	)
}

func (trw *TypeRWStruct) skipTL2Call(
	directImports *DirectImports,
	bytesVersion bool,
	targetBytes string,
	canDependOnLocalBit bool,
	ins *InternalNamespace,
	refObject bool,
) string {
	if trw.isUnwrapType() {
		return trw.Fields[0].t.SkipTL2Call(directImports, bytesVersion, targetBytes, canDependOnLocalBit, ins, refObject)
	}
	return fmt.Sprintf(`if %[2]s, err = basictl.SkipSizedValue(%[2]s); err != nil { return %[2]s, err }`,
		"",
		targetBytes)
}

func (trw *TypeRWStruct) doesZeroSizeMeanEmpty(canDependOnLocalBit bool) bool {
	//if trw.wr.IsTrueType() && trw.wr.unionParent == nil {
	//	return false
	//}
	if trw.isUnwrapType() {
		return trw.Fields[0].t.trw.doesZeroSizeMeanEmpty(canDependOnLocalBit)
	}
	return true
}

func (trw *TypeRWStruct) doesCalculateLayoutUseObject(allowInplace bool) bool {
	//if trw.wr.IsTrueType() && trw.wr.unionParent == nil {
	//	return false
	//}
	if trw.isUnwrapType() {
		return trw.Fields[0].t.trw.doesCalculateLayoutUseObject(allowInplace)
	}
	return true
}

func (trw *TypeRWStruct) isSizeWrittenInData() bool {
	//if trw.wr.IsTrueType() && trw.wr.unionParent == nil {
	//	return false
	//}
	if trw.isUnwrapType() || trw.isTypeDef() {
		return trw.Fields[0].t.trw.isSizeWrittenInData()
	}
	return true
}

func (trw *TypeRWStruct) doesWriteTL2UseObject(canDependOnLocalBit bool) bool {
	//if trw.wr.IsTrueType() && trw.wr.unionParent == nil {
	//	return false
	//}
	if trw.isUnwrapType() {
		return trw.Fields[0].t.trw.doesWriteTL2UseObject(canDependOnLocalBit)
	}
	return true
}

func (trw *TypeRWStruct) doesReadTL2UseObject(canDependOnLocalBit bool) bool {
	//if trw.wr.IsTrueType() && trw.wr.unionParent == nil {
	//	return false
	//}
	if trw.isUnwrapType() {
		return trw.Fields[0].t.trw.doesReadTL2UseObject(canDependOnLocalBit)
	}
	return true
}

func (trw *TypeRWStruct) doesReadTL2UseBytes(canDependOnLocalBit bool) bool {
	//if trw.wr.IsTrueType() && trw.wr.unionParent == nil {
	//	return false
	//}
	if trw.isUnwrapType() {
		return trw.Fields[0].t.trw.doesReadTL2UseObject(canDependOnLocalBit)
	}
	return true
}

func (trw *TypeRWStruct) tl2TrivialSize(targetObject string, canDependOnLocalBit bool, refObject bool) (isConstant bool, size string) {
	//if trw.wr.IsTrueType() && trw.wr.unionParent == nil {
	//	return true, "0"
	//}
	if trw.isUnwrapType() || trw.isTypeDef() {
		return trw.Fields[0].t.trw.tl2TrivialSize(targetObject, canDependOnLocalBit, refObject)
	}
	return false, ""
}

func (trw *TypeRWStruct) AllNewTL2Masks() []string {
	var result []string
	for _, field := range trw.Fields {
		if field.MaskTL2Bit != nil && *field.MaskTL2Bit%8 == 0 {
			result = append(result, fmt.Sprintf("tl2mask%d", *field.MaskTL2Bit/8))
		}
	}
	return result
}
