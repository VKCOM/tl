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
	return fmt.Sprintf("%[1]s = %[2]s.CalculateLayout(%[1]s%[3]s)", targetSizes, addAsterisk(refObject, targetObject), joinWithCommas(natArgs), trw.wr.ins.AddPrefix(ins))
}

func (trw *TypeRWUnion) doesCalculateLayoutUseObject() bool {
	return true
}

func (trw *TypeRWUnion) isSizeWrittenInData() bool {
	return true
}
