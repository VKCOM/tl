// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package gengo

import (
	"fmt"
)

func (trw *TypeRWStruct) calculateLayoutCall(
	directImports *DirectImports,
	bytesVersion bool,
	targetSizes string,
	targetObject string,
	zeroIfEmpty bool,
	ins *InternalNamespace,
	refObject bool,
) string {
	if trw.isUnwrapType() {
		return trw.Fields[0].t.CalculateLayoutCall(directImports, bytesVersion, targetSizes, targetObject, zeroIfEmpty, ins, refObject)
	}
	if trw.wr.IsTrueType() && trw.wr.unionParent == nil {
		if zeroIfEmpty {
			return "if false {"
		}
		return "currentSize += 1"
	}
	if trw.isTypeDef() && trw.wr.unionParent == nil {
		actualType := trw.Fields[0].t.TypeString2(bytesVersion, directImports, ins, false, false)
		if refObject {
			targetObject = fmt.Sprintf("(*%s)(%s)", actualType, targetObject)
		} else {
			targetObject = fmt.Sprintf("(*%s)(&%s)", actualType, targetObject)
			refObject = true
		}
		return trw.Fields[0].t.CalculateLayoutCall(directImports, bytesVersion, targetSizes, targetObject, zeroIfEmpty, ins, refObject)
	}
	sz := fmt.Sprintf("%[1]s, sz = %[2]s.CalculateLayout(%[1]s, %[3]v)", targetSizes, addAsteriskAndBrackets(refObject, targetObject), zeroIfEmpty)
	if zeroIfEmpty {
		sz = fmt.Sprintf("if %s; sz != 0 {", sz)
	}
	return sz + "\ncurrentSize += sz"
}

func (trw *TypeRWStruct) writeTL2Call(
	directImports *DirectImports,
	bytesVersion bool,
	targetSizes string,
	targetBytes string,
	targetObject string,
	zeroIfEmpty bool,
	ins *InternalNamespace,
	refObject bool,
) string {
	if trw.isUnwrapType() {
		return trw.Fields[0].t.WriteTL2Call(directImports, bytesVersion, targetSizes, targetBytes, targetObject, zeroIfEmpty, ins, refObject)
	}
	if trw.wr.IsTrueType() && trw.wr.unionParent == nil {
		if zeroIfEmpty {
			return "if false {"
		}
		return fmt.Sprintf("%[1]s = append(%[1]s, 0)", targetBytes)
	}
	if trw.isTypeDef() && trw.wr.unionParent == nil {
		actualType := trw.Fields[0].t.TypeString2(bytesVersion, directImports, ins, false, false)
		if refObject {
			targetObject = fmt.Sprintf("(*%s)(%s)", actualType, targetObject)
		} else {
			targetObject = fmt.Sprintf("(*%s)(&%s)", actualType, targetObject)
			refObject = true
		}
		return trw.Fields[0].t.WriteTL2Call(directImports, bytesVersion, targetSizes, targetBytes, targetObject, zeroIfEmpty, ins, refObject)
	}
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

func (trw *TypeRWStruct) readTL2Call(
	directImports *DirectImports,
	bytesVersion bool,
	targetBytes string,
	targetObject string,
	canDependOnLocalBit bool,
	ins *InternalNamespace,
	refObject bool,
) string {
	//if trw.wr.IsTrueType() && trw.wr.unionParent == nil {
	//	return ""
	//}
	if trw.isUnwrapType() {
		return trw.Fields[0].t.ReadTL2Call(directImports, bytesVersion, targetBytes, targetObject, canDependOnLocalBit, ins, refObject)
	}
	additionalSuffix := ""
	if trw.wr.unionParent != nil {
		additionalSuffix = ", block"
	}
	return fmt.Sprintf("if %[3]s, err = %[2]s.InternalReadTL2(%[3]s%[4]s); err != nil { return %[3]s, err }",
		"",
		targetObject,
		targetBytes,
		additionalSuffix,
	)
}

func (trw *TypeRWStruct) skipTL2Call(
	directImports *DirectImports,
	bytesVersion bool,
	targetBytes string,
	canDependOnLocalBit bool,
	ins *InternalNamespace,
	refObject bool,
) string {
	if trw.isUnwrapType() {
		return trw.Fields[0].t.SkipTL2Call(directImports, bytesVersion, targetBytes, canDependOnLocalBit, ins, refObject)
	}
	return fmt.Sprintf(`if %[2]s, err = basictl.SkipSizedValue(%[2]s); err != nil { return %[2]s, err }`,
		"",
		targetBytes)
}

func (trw *TypeRWStruct) AllNewTL2Masks() []string {
	var result []string
	for _, field := range trw.Fields {
		if field.MaskTL2Bit() != nil && (*field.MaskTL2Bit())%8 == 0 {
			result = append(result, fmt.Sprintf("tl2mask%d", *field.MaskTL2Bit()/8))
		}
	}
	return result
}
