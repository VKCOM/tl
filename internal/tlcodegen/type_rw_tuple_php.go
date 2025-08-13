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

func (trw *TypeRWBrackets) PhpReadMethodCall(targetName string, bare bool, initIfDefault bool, args *TypeArgumentsTree) []string {
	useBuiltIn := trw.wr.gen.options.UseBuiltinDataProviders
	index := fmt.Sprintf("$i%d", len(trw.PhpClassName(false, true)))
	result := make([]string, 0)
	switch {
	// actual vector
	case trw.vectorLike && !trw.dictLike:
		elementName := fmt.Sprintf("$obj%d", len(trw.PhpClassName(false, true)))
		elementRead := trw.element.t.trw.PhpReadMethodCall(elementName, trw.element.bare, false, args.children[0])
		for i := range elementRead {
			elementRead[i] = "  " + elementRead[i]
		}
		if useBuiltIn {
			result = append(result, "$vector_size = fetch_int() & 0xFFFFFFFF;")
		} else {
			result = append(result,
				"[$vector_size, $success] = $stream->read_uint32();",
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
			fmt.Sprintf("for(%[1]s = 0; %[1]s < $vector_size; %[1]s++) {", index),
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
		elementName := fmt.Sprintf("$obj%d", len(trw.PhpClassName(false, true)))
		tupleSize := *args.children[0].value
		//elementArgs := args[1:]
		elementRead := trw.element.t.trw.PhpReadMethodCall(elementName, trw.element.bare, false, args.children[1])
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
		keyRead := trw.dictKeyField.t.trw.PhpReadMethodCall(keyElement, trw.dictKeyField.bare, true, args)
		for i := range keyRead {
			keyRead[i] = "  " + keyRead[i]
		}
		valueRead := trw.dictValueField.t.trw.PhpReadMethodCall(valueElement, trw.dictValueField.bare, true, args)
		for i := range valueRead {
			valueRead[i] = "  " + valueRead[i]
		}
		if useBuiltIn {
			result = append(result, "$dict_size = fetch_int() & 0xFFFFFFFF;")
		} else {
			result = append(result,
				"[$dict_size, $success] = $stream->read_uint32();",
				"if (!$success) {",
				"  return false;",
				"}",
			)
		}
		result = append(result,
			// TODO MAKE MORE EFFICIENT
			fmt.Sprintf("%[1]s = [];", targetName),
			fmt.Sprintf("for(%[1]s = 0; %[1]s < $dict_size; %[1]s++) {", index),
		)
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

func (trw *TypeRWBrackets) PhpWriteMethodCall(targetName string, bare bool, args *TypeArgumentsTree) []string {
	useBuiltIn := trw.wr.gen.options.UseBuiltinDataProviders
	index := fmt.Sprintf("$i%d", len(trw.PhpClassName(false, true)))
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
			elementRead := trw.element.t.trw.PhpWriteMethodCall(fmt.Sprintf("%[1]s[%[2]s]", targetName, index), trw.element.bare, args.children[0])
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
			elementRead := trw.element.t.trw.PhpWriteMethodCall(fmt.Sprintf("%[1]s[%[2]s]", targetName, index), trw.element.bare, elementArgs)
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
		keyElement := "$key"
		valueElement := "$value"
		result = append(result,
			fmt.Sprintf("ksort(%[1]s);", targetName),
			fmt.Sprintf("foreach(%[1]s as %[2]s => %[3]s) {", targetName, keyElement, valueElement),
		)
		{
			keyRead := trw.dictKeyField.t.trw.PhpWriteMethodCall(keyElement, trw.dictKeyField.bare, args)
			for i := range keyRead {
				keyRead[i] = "  " + keyRead[i]
			}
			valueRead := trw.dictValueField.t.trw.PhpWriteMethodCall(valueElement, trw.dictValueField.bare, args)
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

func (trw *TypeRWBrackets) PhpDefaultInit() string {
	return "[]"
}
