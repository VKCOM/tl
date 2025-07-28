package tlcodegen

type TypeRWTL2 interface {
	calculateLayoutCall(
		directImports *DirectImports,
		bytesVersion bool,
		targetSizes string,
		targetObject string,
		canDependOnLocalBit bool,
		ins *InternalNamespace,
		refObject bool,
	) string

	doesCalculateLayoutUseObject(allowInplace bool) bool
	isSizeWrittenInData() bool
	doesZeroSizeMeanEmpty(canDependOnLocalBit bool) bool

	writeTL2Call(
		directImports *DirectImports,
		bytesVersion bool,
		targetSizes string,
		targetBytes string,
		targetObject string,
		canDependOnLocalBit bool,
		ins *InternalNamespace,
		refObject bool,
	) string

	doesWriteTL2UseObject(canDependOnLocalBit bool) bool
	tl2TrivialSize(targetObject string, canDependOnLocalBit bool, refObject bool) (isConstant bool, size string)

	readTL2Call(
		directImports *DirectImports,
		bytesVersion bool,
		targetBytes string,
		targetObject string,
		canDependOnLocalBit bool,
		ins *InternalNamespace,
		refObject bool,
	) string

	doesReadTL2UseObject(canDependOnLocalBit bool) bool
	doesReadTL2UseBytes(canDependOnLocalBit bool) bool
}

func (w *TypeRWWrapper) CalculateLayoutCall(
	directImports *DirectImports,
	bytesVersion bool,
	targetSizes string,
	targetObject string,
	canDependOnLocalBit bool,
	ins *InternalNamespace,
	refObject bool,
) string {
	bytesVersion = bytesVersion && w.hasBytesVersion
	return w.trw.calculateLayoutCall(directImports, bytesVersion, targetSizes, targetObject, canDependOnLocalBit, ins, refObject)
}

func (w *TypeRWWrapper) WriteTL2Call(
	directImports *DirectImports,
	bytesVersion bool,
	targetSizes string,
	targetBytes string,
	targetObject string,
	canDependOnLocalBit bool,
	ins *InternalNamespace,
	refObject bool,
) string {
	bytesVersion = bytesVersion && w.hasBytesVersion
	return w.trw.writeTL2Call(directImports, bytesVersion, targetSizes, targetBytes, targetObject, canDependOnLocalBit, ins, refObject)
}

func (w *TypeRWWrapper) ReadTL2Call(
	directImports *DirectImports,
	bytesVersion bool,
	targetBytes string,
	targetObject string,
	canDependOnLocalBit bool,
	ins *InternalNamespace,
	refObject bool,
) string {
	bytesVersion = bytesVersion && w.hasBytesVersion
	return w.trw.readTL2Call(directImports, bytesVersion, targetBytes, targetObject, canDependOnLocalBit, ins, refObject)
}
