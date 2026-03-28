// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package genrust

import "fmt"

func (trw *TypeRWBrackets) calculateLayoutCall(
	directImports *DirectImports,
	bytesVersion bool,
	targetSizes string,
	targetObject string,
	zeroIfEmpty bool,
	refObject bool,
) string {
	sz := fmt.Sprintf("%[1]s, sz = %[4]s%[3]sCalculateLayout(%[1]s, %[5]v, %[2]s)",
		targetSizes,
		addAmpersand(refObject, targetObject),
		addBytes(trw.wr.goGlobalName, bytesVersion),
		"",
		zeroIfEmpty,
	)
	if zeroIfEmpty {
		sz = fmt.Sprintf("if %s; sz != 0 {", sz)
	}
	return sz + "\ncurrentSize += sz"
}

func (trw *TypeRWBrackets) writeTL2Call(
	directImports *DirectImports,
	bytesVersion bool,
	targetSizes string,
	targetBytes string,
	targetObject string,
	zeroIfEmpty bool,

	refObject bool,
) string {
	sz := fmt.Sprintf("%[5]s, %[1]s, %[7]s = %[4]s%[3]sInternalWriteTL2(%[5]s, %[1]s, %[6]v, %[2]s)",
		targetSizes,
		addAmpersand(refObject, targetObject),
		addBytes(trw.wr.goGlobalName, bytesVersion),
		"",
		targetBytes,
		zeroIfEmpty,
		ifString(zeroIfEmpty, "sz", "_"),
	)
	if zeroIfEmpty {
		sz = fmt.Sprintf("if %s; sz != 0 {", sz)
	}
	return sz
}

func (trw *TypeRWBrackets) readTL2Call(
	directImports *DirectImports,
	bytesVersion bool,
	targetBytes string,
	targetObject string,
	canDependOnLocalBit bool,

	refObject bool,
) string {
	return fmt.Sprintf("if %[5]s, err = %[4]s%[3]sInternalReadTL2(%[5]s, %[2]s); err != nil { return %[5]s, err }",
		"",
		addAmpersand(refObject, targetObject),
		addBytes(trw.wr.goGlobalName, bytesVersion),
		"",
		targetBytes,
	)
}

func (trw *TypeRWBrackets) skipTL2Call(
	directImports *DirectImports,
	bytesVersion bool,
	targetBytes string,
	canDependOnLocalBit bool,

	refObject bool,
) string {
	return fmt.Sprintf(`if %[2]s, err = basictl.SkipSizedValue(%[2]s); err != nil { return %[2]s, err }`,
		"",
		targetBytes)
}
