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
	doesZeroSizeMeanEmpty(canDependOnLocalBit bool) bool

	writeTL2Call(
		bytesVersion bool,
		targetSizes string,
		targetBytes string,
		targetObject string,
		canDependOnLocalBit bool,
		ins *InternalNamespace,
		refObject bool,
		natArgs []string,
	) string
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

func (w *TypeRWWrapper) WriteTL2Call(
	bytesVersion bool,
	targetSizes string,
	targetBytes string,
	targetObject string,
	canDependOnLocalBit bool,
	ins *InternalNamespace,
	refObject bool,
	natArgs []string,
) string {
	bytesVersion = bytesVersion && w.hasBytesVersion
	return w.trw.writeTL2Call(bytesVersion, targetSizes, targetBytes, targetObject, canDependOnLocalBit, ins, refObject, natArgs)
}
