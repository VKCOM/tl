package tlcodegen

import (
	"fmt"
	"github.com/vkcom/tl/internal/utils"
	"sort"
)

func (trw *TypeRWStruct) calculateLayoutCall(
	directImports *DirectImports,
	bytesVersion bool,
	targetSizes string,
	targetObject string,
	canDependOnLocalBit bool,
	ins *InternalNamespace,
	refObject bool,
) string {
	//if trw.wr.IsTrueType() && trw.wr.unionParent == nil {
	//	return ""
	//}
	if trw.isUnwrapType() {
		return trw.Fields[0].t.CalculateLayoutCall(directImports, bytesVersion, targetSizes, targetObject, canDependOnLocalBit, ins, refObject)
	}
	return fmt.Sprintf("%[1]s = %[2]s.CalculateLayout(%[1]s)", targetSizes, addAsteriskAndBrackets(refObject, targetObject))
}

func (trw *TypeRWStruct) writeTL2Call(
	directImports *DirectImports,
	bytesVersion bool,
	targetSizes string,
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
		return trw.Fields[0].t.WriteTL2Call(directImports, bytesVersion, targetSizes, targetBytes, targetObject, canDependOnLocalBit, ins, refObject)
	}
	return fmt.Sprintf("%[3]s, %[1]s = %[2]s.InternalWriteTL2(%[3]s, %[1]s)",
		targetSizes,
		targetObject,
		targetBytes,
	)
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

func (trw *TypeRWStruct) AllRequiredTL2Masks() []int {
	required := make(map[int]bool)
	for _, field := range trw.Fields {
		if field.fieldMask != nil && field.fieldMask.IsTL2() {
			required[-field.fieldMask.FieldIndex] = true
		}
	}
	keys := utils.Keys(required)
	sort.Ints(keys)
	return keys
}
