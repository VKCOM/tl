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

func (trw *TypeRWBrackets) doesCalculateLayoutUseObject() bool {
	return true
}

func (trw *TypeRWBrackets) isSizeWrittenInData() bool {
	return true
}
