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

	doesWriteTL2UseObject(canDependOnLocalBit bool) bool
	tl2TrivialSize(targetObject string, canDependOnLocalBit bool, refObject bool) (isConstant bool, size string)

	readTL2Call(
		bytesVersion bool,
		targetBytes string,
		targetObject string,
		canDependOnLocalBit bool,
		ins *InternalNamespace,
		refObject bool,
		natArgs []string,
	) string

	doesReadTL2UseObject(canDependOnLocalBit bool) bool
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

func (w *TypeRWWrapper) ReadTL2Call(
	bytesVersion bool,
	targetBytes string,
	targetObject string,
	canDependOnLocalBit bool,
	ins *InternalNamespace,
	refObject bool,
	natArgs []string,
) string {
	bytesVersion = bytesVersion && w.hasBytesVersion
	return w.trw.readTL2Call(bytesVersion, targetBytes, targetObject, canDependOnLocalBit, ins, refObject, natArgs)
}
