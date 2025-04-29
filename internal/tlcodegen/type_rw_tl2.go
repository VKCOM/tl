package tlcodegen

type TypeRWTL2 interface {
	calculateLayout(
		bytesVersion bool,
		targetSizes string,
		targetObject string,
		canDependOnLocalBit bool,
		ins *InternalNamespace,
		refObject bool,
		natArgs []string,
	) string
	doesCalculateLayoutUseObject() bool
	isSizeWrittenInData() bool
}

func (w *TypeRWWrapper) CalculateLayout(
	bytesVersion bool,
	targetSizes string,
	targetObject string,
	canDependOnLocalBit bool,
	ins *InternalNamespace,
	refObject bool,
	natArgs []string,
) string {
	bytesVersion = bytesVersion && w.hasBytesVersion
	return w.trw.calculateLayout(bytesVersion, targetSizes, targetObject, canDependOnLocalBit, ins, refObject, natArgs)
}
