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

func (trw *TypeRWMaybe) doesCalculateLayoutUseObject() bool {
	return true
}

func (trw *TypeRWMaybe) isSizeWrittenInData() bool {
	return trw.element.t.trw.isSizeWrittenInData()
}
