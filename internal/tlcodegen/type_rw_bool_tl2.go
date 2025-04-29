package tlcodegen

import "fmt"

/**
BOOL = BYTE or BIT in STRUCT
*/

func (trw *TypeRWBool) calculateLayout(
	bytesVersion bool,
	targetSizes string,
	targetObject string,
	canDependOnLocalBit bool,
	ins *InternalNamespace,
	refObject bool,
	natArgs []string) string {
	if canDependOnLocalBit {
		return fmt.Sprintf("%[1]s = append(%[1]s, 0)", targetSizes)
	}
	return fmt.Sprintf("%[1]s = append(%[1]s, 1)", targetSizes)
}

func (trw *TypeRWBool) doesCalculateLayoutUseObject() bool {
	return false
}

func (trw *TypeRWBool) isSizeWrittenInData() bool {
	return false
}
