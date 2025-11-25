// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlcodegen

import (
	"fmt"
	"strings"
)

func (trw *TypeRWBrackets) PhpClassName(withPath bool, bare bool) string {
	if strings.HasPrefix(trw.wr.tlName.String(), BuiltinTupleName) ||
		strings.HasPrefix(trw.wr.tlName.String(), BuiltinVectorName) {
		return "array_" + trw.element.t.trw.PhpClassName(false, false)
	}
	return fmt.Sprintf("<? %s>", trw.wr.goGlobalName)
}

func (trw *TypeRWBrackets) PhpClassNameReplaced() bool {
	return true
}

func (trw *TypeRWBrackets) PhpTypeName(withPath bool, bare bool) string {
	if strings.HasPrefix(trw.wr.tlName.String(), BuiltinTupleName) ||
		strings.HasPrefix(trw.wr.tlName.String(), BuiltinVectorName) {
		elementText := trw.element.t.trw.PhpTypeName(withPath, trw.element.t.PHPIsBare())
		if _, ok := trw.element.t.trw.(*TypeRWMaybe); ok {
			elementText = "(" + elementText + ")"
		}
		return elementText + "[]"
	}
	return fmt.Sprintf("<? %s>", trw.wr.goGlobalName)
}

func (trw *TypeRWBrackets) PhpGenerateCode(code *strings.Builder, bytes bool) error {
	return fmt.Errorf("tuples don't have php code")
}

func (trw *TypeRWBrackets) PhpDefaultValue() string {
	return "[]"
}

func (trw *TypeRWBrackets) PhpIterateReachableTypes(reachableTypes *map[*TypeRWWrapper]bool) {
	trw.element.t.PhpIterateReachableTypes(reachableTypes)
}

func (trw *TypeRWBrackets) PhpHasPrimitiveInnerElement() bool {
	if vec, ok := trw.element.t.trw.(*TypeRWBrackets); ok {
		return vec.PhpHasPrimitiveInnerElement()
	}
	core := trw.element.t.PHPGenCoreType()
	_, ok := core.trw.(*TypeRWPrimitive)
	return ok
}

func (trw *TypeRWBrackets) PhpReadMethodCall(targetName string, bare bool, initIfDefault bool, args *TypeArgumentsTree, supportSuffix string) []string {
	useBuiltIn := trw.wr.gen.options.UseBuiltinDataProviders
	index := fmt.Sprintf("$i%d", len(trw.PhpClassName(false, true)))
	result := make([]string, 0)
	switch {
	// actual vector
	case trw.vectorLike && !trw.dictLike:
		elementName := fmt.Sprintf("$obj%s_%d", supportSuffix, len(trw.PhpClassName(false, true)))
		vectorSizeName := fmt.Sprintf("$vector_size_%d", len(trw.PhpClassName(false, true)))
		elementRead := trw.element.t.trw.PhpReadMethodCall(elementName, trw.element.bare, false, args.children[0], supportSuffix)
		for i := range elementRead {
			elementRead[i] = "  " + elementRead[i]
		}
		if useBuiltIn {
			result = append(result, fmt.Sprintf("%s = fetch_int() & 0xFFFFFFFF;", vectorSizeName))
		} else {
			result = append(result,
				fmt.Sprintf("[%s, $success] = $stream->read_uint32();", vectorSizeName),
				"if (!$success) {",
				"  return false;",
				"}",
			)
		}
		if initIfDefault {
			result = append(result,
				// TODO MAKE MORE EFFICIENT
				fmt.Sprintf("%[1]s = [];", targetName),
			)
		}
		result = append(result,
			fmt.Sprintf("for(%[1]s = 0; %[1]s < %[2]s; %[1]s++) {", index, vectorSizeName),
			fmt.Sprintf("  /** @var %[1]s */", trw.element.t.trw.PhpTypeName(true, true)),
			fmt.Sprintf("  %[2]s = %[1]s;", trw.element.t.trw.PhpDefaultInit(), elementName),
		)
		result = append(result, elementRead...)
		result = append(result,
			fmt.Sprintf("  %[1]s[] = %[2]s;", targetName, elementName),
			"}",
		)
		return result
	// tuple with size as last argument
	case !trw.vectorLike && !trw.dictLike:
		elementName := fmt.Sprintf("$obj%s_%d", supportSuffix, len(trw.PhpClassName(false, true)))
		tupleSize := *args.children[0].value
		//elementArgs := args[1:]
		elementRead := trw.element.t.trw.PhpReadMethodCall(elementName, trw.element.bare, false, args.children[1], supportSuffix)
		for i := range elementRead {
			elementRead[i] = "  " + elementRead[i]
		}
		if initIfDefault {
			result = append(result,
				// TODO MAKE MORE EFFICIENT
				fmt.Sprintf("%[1]s = [];", targetName),
			)
		}
		result = append(result,
			fmt.Sprintf("for(%[1]s = 0; %[1]s < %[2]s; %[1]s++) {", index, tupleSize),
			fmt.Sprintf("  /** @var %[1]s */", trw.element.t.trw.PhpTypeName(true, true)),
			fmt.Sprintf("  %[2]s = %[1]s;", trw.element.t.trw.PhpDefaultInit(), elementName),
		)
		result = append(result, elementRead...)
		result = append(result,
			fmt.Sprintf("  %[1]s[] = %[2]s;", targetName, elementName),
			"}",
		)
		return result
	// actual map / dictionary
	case trw.dictLike:
		keyElement := fmt.Sprintf("$%s___key", trw.PhpClassName(false, true))
		valueElement := fmt.Sprintf("$%s___value", trw.PhpClassName(false, true))
		dictSizeName := fmt.Sprintf("$dict_size_%d", len(trw.PhpClassName(false, true)))
		keyRead := trw.dictKeyField.t.trw.PhpReadMethodCall(keyElement, trw.dictKeyField.bare, true, args.children[0], supportSuffix)
		for i := range keyRead {
			keyRead[i] = "  " + keyRead[i]
		}
		valueRead := trw.dictValueField.t.trw.PhpReadMethodCall(valueElement, trw.dictValueField.bare, true, args.children[0], supportSuffix)
		for i := range valueRead {
			valueRead[i] = "  " + valueRead[i]
		}
		if useBuiltIn {
			result = append(result, fmt.Sprintf("%[1]s = fetch_int() & 0xFFFFFFFF;", dictSizeName))
		} else {
			result = append(result,
				fmt.Sprintf("[%[1]s, $success] = $stream->read_uint32();", dictSizeName),
				"if (!$success) {",
				"  return false;",
				"}",
			)
		}
		result = append(result,
			// TODO MAKE MORE EFFICIENT
			fmt.Sprintf("%[1]s = [];", targetName),
			fmt.Sprintf("for(%[1]s = 0; %[1]s < %[2]s; %[1]s++) {", index, dictSizeName),
		)
		result = append(result, fmt.Sprintf("  /** @var %[1]s */", trw.dictKeyField.t.trw.PhpTypeName(true, true)))
		result = append(result, fmt.Sprintf("  %[1]s = %[2]s;", keyElement, trw.dictKeyField.t.trw.PhpDefaultInit()))
		result = append(result, fmt.Sprintf("  /** @var %[1]s */", trw.dictValueField.t.trw.PhpTypeName(true, true)))
		result = append(result, fmt.Sprintf("  %[1]s = %[2]s;", valueElement, trw.dictValueField.t.trw.PhpDefaultInit()))
		result = append(result, keyRead...)
		result = append(result, valueRead...)
		result = append(result,
			fmt.Sprintf("  %[1]s[%[2]s] = %[3]s;", targetName, keyElement, valueElement),
			"}",
		)
		return result
	}
	return []string{fmt.Sprintf("<??? %s read>", trw.wr.goGlobalName)}
}

func (trw *TypeRWBrackets) PhpWriteMethodCall(targetName string, bare bool, args *TypeArgumentsTree, supportSuffix string) []string {
	useBuiltIn := trw.wr.gen.options.UseBuiltinDataProviders
	layerIndex := len(trw.PhpClassName(false, true))
	index := fmt.Sprintf("$i%d", layerIndex)
	elementObj := fmt.Sprintf("$obj%s_%d", supportSuffix, layerIndex)
	result := make([]string, 0)
	switch {
	// actual vector
	case trw.vectorLike && !trw.dictLike:
		if useBuiltIn {
			result = append(result, fmt.Sprintf("store_int(count(%[1]s));", targetName))
		} else {
			result = append(result,
				fmt.Sprintf("$success = $stream->write_uint32(count(%[1]s));", targetName),
				"if (!$success) {",
				"  return false;",
				"}",
			)
		}
		result = append(result,
			// TODO MAKE MORE EFFICIENT
			fmt.Sprintf("for(%[1]s = 0; %[1]s < count(%[2]s); %[1]s++) {", index, targetName),
		)
		{
			result = append(result, fmt.Sprintf("  %[1]s = %[2]s;", elementObj, fmt.Sprintf("%[1]s[%[2]s]", targetName, index)))
			elementRead := trw.element.t.trw.PhpWriteMethodCall(elementObj, trw.element.bare, args.children[0], supportSuffix)
			for i := range elementRead {
				elementRead[i] = "  " + elementRead[i]
			}
			result = append(result, elementRead...)
		}
		result = append(result,
			"}",
		)
		return result
	// tuple with size as last argument
	case !trw.vectorLike && !trw.dictLike:
		tupleSize := *args.children[0].value
		elementArgs := args.children[1]
		result = append(result,
			// TODO MAKE MORE EFFICIENT
			fmt.Sprintf("for(%[1]s = 0; %[1]s < %[2]s; %[1]s++) {", index, tupleSize),
		)
		{
			result = append(result, fmt.Sprintf("  %[1]s = %[2]s;", elementObj, fmt.Sprintf("%[1]s[%[2]s]", targetName, index)))
			elementRead := trw.element.t.trw.PhpWriteMethodCall(elementObj, trw.element.bare, elementArgs, supportSuffix)
			for i := range elementRead {
				elementRead[i] = "  " + elementRead[i]
			}
			result = append(result, elementRead...)
		}
		result = append(result, "}")
		return result
	// actual map / dictionary
	case trw.dictLike:
		if useBuiltIn {
			result = append(result, fmt.Sprintf("store_int(count(%[1]s));", targetName))
		} else {
			result = append(result,
				fmt.Sprintf("$success = $stream->write_uint32(count(%[1]s));", targetName),
				"if (!$success) {",
				"  return false;",
				"}",
			)
		}
		keyElement := fmt.Sprintf("$%s___key", trw.PhpClassName(false, true))
		valueElement := fmt.Sprintf("$%s___value", trw.PhpClassName(false, true))
		flags := ""
		if trw.dictKeyString {
			flags = ", SORT_STRING"
		}
		result = append(result,
			fmt.Sprintf("ksort(%[1]s%[2]s);", targetName, flags),
			fmt.Sprintf("foreach(%[1]s as %[2]s => %[3]s) {", targetName, keyElement, valueElement),
		)
		{
			keyRead := trw.dictKeyField.t.trw.PhpWriteMethodCall(keyElement, trw.dictKeyField.bare, args, supportSuffix)
			for i := range keyRead {
				keyRead[i] = "  " + keyRead[i]
			}
			valueRead := trw.dictValueField.t.trw.PhpWriteMethodCall(valueElement, trw.dictValueField.bare, args, supportSuffix)
			for i := range valueRead {
				valueRead[i] = "  " + valueRead[i]
			}
			result = append(result, keyRead...)
			result = append(result, valueRead...)
		}
		result = append(result, "}")
		return result
	}
	return []string{fmt.Sprintf("<??? %s write>", trw.wr.goGlobalName)}
}

func (trw *TypeRWBrackets) PhpReadTL2MethodCall(targetName string, bare bool, initIfDefault bool, args *TypeArgumentsTree, supportSuffix string, callLevel int, usedBytesPointer string, canDependOnLocalBit bool) []string {
	//panic("not implemented")
	return []string{"// TODO FOR BRACKETS"}
}

func (trw *TypeRWBrackets) PhpDefaultInit() string {
	return "[]"
}
