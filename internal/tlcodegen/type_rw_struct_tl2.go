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
