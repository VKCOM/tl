package tlcodegen

import "fmt"

func (trw *TypeRWStruct) calculateLayout(
	bytesVersion bool,
	targetSizes string,
	targetObject string,
	canDependOnLocalBit bool,
	ins *InternalNamespace,
	refObject bool,
	natArgs []string) string {
	if trw.wr.IsTrueType() && trw.wr.unionParent == nil {
		return fmt.Sprintf("%[1]s = append(%[1]s, 0)", targetSizes)
	}
	if trw.isUnwrapType() {
		return trw.Fields[0].t.CalculateLayout(bytesVersion, targetSizes, targetObject, canDependOnLocalBit, ins, refObject, natArgs)
	}
	return fmt.Sprintf("%[1]s = %[2]s.CalculateLayout(%[1]s%[3]s)", targetSizes, addAsteriskAndBrackets(refObject, targetObject), joinWithCommas(natArgs))
}

func (trw *TypeRWStruct) writeTL2Call(
	bytesVersion bool,
	targetSizes string,
	targetBytes string,
	targetObject string,
	canDependOnLocalBit bool,
	ins *InternalNamespace,
	refObject bool,
	natArgs []string) string {
	if trw.wr.IsTrueType() && trw.wr.unionParent == nil {
		return fmt.Sprintf("%[1]s = %[1]s[1:]", targetSizes)
	}
	if trw.isUnwrapType() {
		return trw.Fields[0].t.WriteTL2Call(bytesVersion, targetSizes, targetBytes, targetObject, canDependOnLocalBit, ins, refObject, natArgs)
	}
	return fmt.Sprintf("%[4]s, %[1]s = %[2]s.InternalWriteTL2(%[4]s, %[1]s%[3]s)",
		targetSizes,
		addAsteriskAndBrackets(refObject, targetObject),
		joinWithCommas(natArgs),
		targetBytes,
	)
}

func (trw *TypeRWStruct) readTL2Call(
	bytesVersion bool,
	targetBytes string,
	targetObject string,
	canDependOnLocalBit bool,
	ins *InternalNamespace,
	refObject bool,
	natArgs []string,
) string {
	if trw.wr.IsTrueType() && trw.wr.unionParent == nil {
		return ""
	}
	if trw.isUnwrapType() {
		return trw.Fields[0].t.ReadTL2Call(bytesVersion, targetBytes, targetObject, canDependOnLocalBit, ins, refObject, natArgs)
	}
	return fmt.Sprintf("if %[4]s, err = %[2]s.ReadTL2(%[4]s%[3]s); err != nil { return %[4]s, err }",
		"",
		addAsteriskAndBrackets(refObject, targetObject),
		joinWithCommas(natArgs),
		targetBytes,
	)
}

func (trw *TypeRWStruct) doesZeroSizeMeanEmpty(canDependOnLocalBit bool) bool {
	if trw.wr.IsTrueType() && trw.wr.unionParent == nil {
		return false
	}
	if trw.isUnwrapType() {
		return trw.Fields[0].t.trw.doesZeroSizeMeanEmpty(canDependOnLocalBit)
	}
	return true
}

func (trw *TypeRWStruct) doesCalculateLayoutUseObject() bool {
	if trw.wr.IsTrueType() && trw.wr.unionParent == nil {
		return false
	}
	if trw.isUnwrapType() {
		return trw.Fields[0].t.trw.doesCalculateLayoutUseObject()
	}
	return true
}

func (trw *TypeRWStruct) isSizeWrittenInData() bool {
	return !(trw.wr.IsTrueType() && trw.wr.unionParent == nil)
}

func (trw *TypeRWStruct) doesWriteTL2UseObject(canDependOnLocalBit bool) bool {
	if trw.wr.IsTrueType() && trw.wr.unionParent == nil {
		return false
	}
	if trw.isUnwrapType() {
		return trw.Fields[0].t.trw.doesWriteTL2UseObject(canDependOnLocalBit)
	}
	return true
}

func (trw *TypeRWStruct) doesReadTL2UseObject(canDependOnLocalBit bool) bool {
	if trw.wr.IsTrueType() && trw.wr.unionParent == nil {
		return false
	}
	if trw.isUnwrapType() {
		return trw.Fields[0].t.trw.doesReadTL2UseObject(canDependOnLocalBit)
	}
	return true
}
