// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package genrust

type TypeRWTL2 interface {
	calculateLayoutCall(
		directImports *DirectImports,
		bytesVersion bool,
		targetSizes string,
		targetObject string,
		canDependOnLocalBit bool,
		refObject bool,
	) string

	writeTL2Call(
		directImports *DirectImports,
		bytesVersion bool,
		targetSizes string,
		targetBytes string,
		targetObject string,
		canDependOnLocalBit bool,
		refObject bool,
	) string

	readTL2Call(
		directImports *DirectImports,
		bytesVersion bool,
		targetBytes string,
		targetObject string,
		canDependOnLocalBit bool,

		refObject bool,
	) string

	skipTL2Call(
		directImports *DirectImports,
		bytesVersion bool,
		targetBytes string,
		canDependOnLocalBit bool,

		refObject bool,
	) string
}

func (w *TypeRWWrapper) CalculateLayoutCall(
	directImports *DirectImports,
	bytesVersion bool,
	targetSizes string,
	targetObject string,
	zeroIfEmpty bool,

	refObject bool,
) string {
	bytesVersion = bytesVersion && w.hasBytesVersion
	return w.trw.calculateLayoutCall(directImports, bytesVersion, targetSizes, targetObject, zeroIfEmpty, refObject)
}

func (w *TypeRWWrapper) WriteTL2Call(
	directImports *DirectImports,
	bytesVersion bool,
	targetSizes string,
	targetBytes string,
	targetObject string,
	zeroIfEmpty bool,

	refObject bool,
) string {
	bytesVersion = bytesVersion && w.hasBytesVersion
	return w.trw.writeTL2Call(directImports, bytesVersion, targetSizes, targetBytes, targetObject, zeroIfEmpty, refObject)
}

func (w *TypeRWWrapper) ReadTL2Call(
	directImports *DirectImports,
	bytesVersion bool,
	targetBytes string,
	targetObject string,
	canDependOnLocalBit bool,

	refObject bool,
) string {
	bytesVersion = bytesVersion && w.hasBytesVersion
	return w.trw.readTL2Call(directImports, bytesVersion, targetBytes, targetObject, canDependOnLocalBit, refObject)
}

func (w *TypeRWWrapper) SkipTL2Call(
	directImports *DirectImports,
	bytesVersion bool,
	targetBytes string,
	canDependOnLocalBit bool,

	refObject bool,
) string {
	bytesVersion = bytesVersion && w.hasBytesVersion
	return w.trw.skipTL2Call(directImports, bytesVersion, targetBytes, canDependOnLocalBit, refObject)
}
