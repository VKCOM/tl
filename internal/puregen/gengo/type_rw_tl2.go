// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package gengo

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

	skipTL2Call(
		directImports *DirectImports,
		bytesVersion bool,
		targetBytes string,
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
	zeroIfEmpty bool,
	ins *InternalNamespace,
	refObject bool,
) string {
	bytesVersion = bytesVersion && w.hasBytesVersion
	return w.trw.calculateLayoutCall(directImports, bytesVersion, targetSizes, targetObject, zeroIfEmpty, ins, refObject)
}

func (w *TypeRWWrapper) WriteTL2Call(
	directImports *DirectImports,
	bytesVersion bool,
	targetSizes string,
	targetBytes string,
	targetObject string,
	zeroIfEmpty bool,
	ins *InternalNamespace,
	refObject bool,
) string {
	bytesVersion = bytesVersion && w.hasBytesVersion
	return w.trw.writeTL2Call(directImports, bytesVersion, targetSizes, targetBytes, targetObject, zeroIfEmpty, ins, refObject)
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

func (w *TypeRWWrapper) SkipTL2Call(
	directImports *DirectImports,
	bytesVersion bool,
	targetBytes string,
	canDependOnLocalBit bool,
	ins *InternalNamespace,
	refObject bool,
) string {
	bytesVersion = bytesVersion && w.hasBytesVersion
	return w.trw.skipTL2Call(directImports, bytesVersion, targetBytes, canDependOnLocalBit, ins, refObject)
}
