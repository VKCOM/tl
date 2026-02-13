// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package gengo

import "fmt"

func (trw *TypeRWMaybe) calculateLayoutCall(
	directImports *DirectImports,
	bytesVersion bool,
	targetSizes string,
	targetObject string,
	zeroIfEmpty bool,
	ins *InternalNamespace,
	refObject bool,
) string {
	sz := fmt.Sprintf("%[1]s, sz = %[2]s.CalculateLayout(%[1]s, %[3]v)", targetSizes, addAsteriskAndBrackets(refObject, targetObject), zeroIfEmpty)
	if zeroIfEmpty {
		sz = fmt.Sprintf("if %s; sz != 0 {", sz)
	}
	return sz + "\ncurrentSize += sz"
}

func (trw *TypeRWMaybe) writeTL2Call(
	directImports *DirectImports,
	bytesVersion bool,
	targetSizes string,
	targetBytes string,
	targetObject string,
	zeroIfEmpty bool,
	ins *InternalNamespace,
	refObject bool,
) string {
	sz := fmt.Sprintf("%[3]s, %[1]s, %[5]s = %[2]s.InternalWriteTL2(%[3]s, %[1]s, %[4]v)",
		targetSizes,
		targetObject,
		targetBytes,
		zeroIfEmpty,
		ifString(zeroIfEmpty, "sz", "_"),
	)
	if zeroIfEmpty {
		sz = fmt.Sprintf("if %s; sz != 0 {", sz)
	}
	return sz
}

func (trw *TypeRWMaybe) readTL2Call(
	directImports *DirectImports,
	bytesVersion bool,
	targetBytes string,
	targetObject string,
	canDependOnLocalBit bool,
	ins *InternalNamespace,
	refObject bool,
) string {
	return fmt.Sprintf("if %[3]s, err = %[2]s.InternalReadTL2(%[3]s); err != nil { return %[3]s, err }",
		"",
		addAsteriskAndBrackets(refObject, targetObject),
		targetBytes,
	)
}

func (trw *TypeRWMaybe) skipTL2Call(
	directImports *DirectImports,
	bytesVersion bool,
	targetBytes string,
	canDependOnLocalBit bool,
	ins *InternalNamespace,
	refObject bool,
) string {
	return fmt.Sprintf(`if %[2]s, err = basictl.SkipSizedValue(%[2]s); err != nil { return %[2]s, err }`,
		"",
		targetBytes)
}

func (trw *TypeRWMaybe) doesReadTL2UseObject(canDependOnLocalBit bool) bool {
	return true
}
