package tlcodegen

import (
	"fmt"
)

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
