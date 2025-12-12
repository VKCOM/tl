// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlcodegen

import (
	"fmt"
	"strings"

	"github.com/vkcom/tl/internal/utils"
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
	result := make([]string, 0)

	//panic("not implemented")
	//useBuiltIn := trw.wr.gen.options.UseBuiltinDataProviders

	index := fmt.Sprintf("$i_%s_%d", supportSuffix, callLevel)
	localUsedBytesPointer := fmt.Sprintf("$used_bytes_%[1]s_%[2]d", supportSuffix, callLevel)
	localCurrentSize := fmt.Sprintf("$current_size_%[1]s_%[2]d", supportSuffix, callLevel)
	localElementCount := fmt.Sprintf("$element_count_%[1]s_%[2]d", supportSuffix, callLevel)

	switch {
	// actual vector
	case trw.vectorLike && !trw.dictLike:
		elementName := fmt.Sprintf("$obj%s_%d", supportSuffix, len(trw.PhpClassName(false, true)))
		elementRead := trw.element.t.trw.PhpReadTL2MethodCall(elementName, trw.element.bare, false, args.children[0], supportSuffix, callLevel+1, localUsedBytesPointer, false)
		for i := range elementRead {
			elementRead[i] = "  " + elementRead[i]
		}
		if initIfDefault {
			result = append(result,
				// TODO MAKE MORE EFFICIENT
				fmt.Sprintf("%[1]s = [];", targetName),
			)
		}
		if _, ok := trw.element.t.trw.(*TypeRWBool); ok {
			localBytesCount := fmt.Sprintf("$byte_blocks_count_%[1]s_%[2]d", supportSuffix, callLevel)
			localByteIndex := fmt.Sprintf("$byte_index_%[1]s_%[2]d", supportSuffix, callLevel)
			localBitIndex := fmt.Sprintf("$bit_index_%[1]s_%[2]d", supportSuffix, callLevel)

			localByte := fmt.Sprintf("$byte_%[1]s_%[2]d", supportSuffix, callLevel)
			localBooleanPointer := fmt.Sprintf("$boolean_index_%[1]s_%[2]d", supportSuffix, callLevel)
			localBooleanCount := fmt.Sprintf("$boolean_count_%[1]s_%[2]d", supportSuffix, callLevel)

			bytesRead := []string{
				fmt.Sprintf("%[1]s = 0;", localBooleanPointer),
				fmt.Sprintf("%[1]s = (%[2]s + 7) / 8;", localBytesCount, localElementCount),
				fmt.Sprintf("for (%[1]s = 0; %[1]s < %[2]s; %[1]s++) {", localByteIndex, localBytesCount),
				fmt.Sprintf("  %[1]s = fetch_byte();", localByte),
				fmt.Sprintf("  %[1]s += 1;", localUsedBytesPointer),
				fmt.Sprintf("  %[1]s = 8;", localBooleanCount),
				fmt.Sprintf("  if (%[1]s + %[2]s > %[3]s) {", localBooleanPointer, localBooleanCount, localElementCount),
				fmt.Sprintf("    %[1]s = %[2]s - %[3]s;", localBooleanCount, localElementCount, localBooleanPointer),
				"  }",
				fmt.Sprintf("  for (%[1]s = 0; %[1]s < %[2]s; %[1]s++) {", localBitIndex, localBooleanCount),
				fmt.Sprintf("    %[1]s[] = (%[2]s & (1 << %[3]s)) != 0;", targetName, localByte, localBitIndex),
				fmt.Sprintf("    %[1]s += 1;", localBooleanPointer),
				"  }",
				"}",
			}
			result = append(result, utils.ShiftAll(bytesRead, "  ")...)
		} else {
			result = append(result,
				fmt.Sprintf("for(%[1]s = 0; %[1]s < %[2]s; %[1]s++) {", index, localElementCount),
				fmt.Sprintf("  /** @var %[1]s */", trw.element.t.trw.PhpTypeName(true, true)),
				fmt.Sprintf("  %[2]s = %[1]s;", trw.element.t.trw.PhpDefaultInit(), elementName),
			)
			result = append(result, elementRead...)
			result = append(result,
				fmt.Sprintf("  %[1]s[] = %[2]s;", targetName, elementName),
				"}",
			)
		}
		break
	// tuple with size as last argument
	case !trw.vectorLike && !trw.dictLike:
		elementName := fmt.Sprintf("$obj%s_%d", supportSuffix, len(trw.PhpClassName(false, true)))
		tupleSize := *args.children[0].value
		//elementArgs := args[1:]
		elementRead := trw.element.t.trw.PhpReadTL2MethodCall(elementName, trw.element.bare, false, args.children[1], supportSuffix, callLevel+1, localUsedBytesPointer, false)
		for i := range elementRead {
			elementRead[i] = "  " + elementRead[i]
		}
		if initIfDefault {
			result = append(result,
				// TODO MAKE MORE EFFICIENT
				fmt.Sprintf("%[1]s = [];", targetName),
			)
		}
		if _, ok := trw.element.t.trw.(*TypeRWBool); ok {
			localBytesCount := fmt.Sprintf("$byte_blocks_count_%[1]s_%[2]d", supportSuffix, callLevel)
			localByteIndex := fmt.Sprintf("$byte_index_%[1]s_%[2]d", supportSuffix, callLevel)
			localBitIndex := fmt.Sprintf("$bit_index_%[1]s_%[2]d", supportSuffix, callLevel)

			localByte := fmt.Sprintf("$byte_%[1]s_%[2]d", supportSuffix, callLevel)
			localBooleanPointer := fmt.Sprintf("$boolean_index_%[1]s_%[2]d", supportSuffix, callLevel)
			localBooleanCount := fmt.Sprintf("$boolean_count_%[1]s_%[2]d", supportSuffix, callLevel)

			bytesRead := []string{
				fmt.Sprintf("%[1]s = 0;", localBooleanPointer),
				fmt.Sprintf("if (%[1]s > %[2]s) {", localElementCount, tupleSize),
				fmt.Sprintf("  %[1]s = %[2]s;", localElementCount, tupleSize),
				"}",
				fmt.Sprintf("%[1]s = (%[2]s + 7) / 8;", localBytesCount, localElementCount),
				fmt.Sprintf("for (%[1]s = 0; %[1]s < %[2]s; %[1]s++) {", localByteIndex, localBytesCount),
				fmt.Sprintf("  %[1]s = fetch_byte();", localByte),
				fmt.Sprintf("  %[1]s += 1;", localUsedBytesPointer),
				fmt.Sprintf("  %[1]s = 8;", localBooleanCount),
				fmt.Sprintf("  if (%[1]s + %[2]s > %[3]s) {", localBooleanPointer, localBooleanCount, localElementCount),
				fmt.Sprintf("    %[1]s = %[2]s - %[3]s;", localBooleanCount, localElementCount, localBooleanPointer),
				"  }",
				fmt.Sprintf("  for (%[1]s = 0; %[1]s < %[2]s; %[1]s++) {", localBitIndex, localBooleanCount),
				fmt.Sprintf("    %[1]s[] = (%[2]s & (1 << %[3]s)) != 0;", targetName, localByte, localBitIndex),
				fmt.Sprintf("    %[1]s += 1;", localBooleanPointer),
				"  }",
				"}",
			}
			result = append(result, utils.ShiftAll(bytesRead, "  ")...)
		} else {
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
		}
		break
	// actual map / dictionary
	case trw.dictLike:
		keyElement := fmt.Sprintf("$%s___key", trw.PhpClassName(false, true))
		valueElement := fmt.Sprintf("$%s___value", trw.PhpClassName(false, true))

		keyRead := trw.dictKeyField.t.trw.PhpReadTL2MethodCall(keyElement, trw.dictKeyField.bare, true, args.children[0], supportSuffix, callLevel+1, localUsedBytesPointer, false)
		valueRead := trw.dictValueField.t.trw.PhpReadTL2MethodCall(valueElement, trw.dictValueField.bare, true, args.children[0], supportSuffix, callLevel+1, localUsedBytesPointer, false)

		result = append(result,
			// TODO MAKE MORE EFFICIENT
			fmt.Sprintf("%[1]s = [];", targetName),
			fmt.Sprintf("for(%[1]s = 0; %[1]s < %[2]s; %[1]s++) {", index, localElementCount),
		)
		// read pair prefix
		pairCurrentSize := fmt.Sprintf("$pair_current_size_%[1]s_%[2]d", supportSuffix, callLevel)
		pairUsedBytes := fmt.Sprintf("$pair_used_bytes_%[1]s_%[2]d", supportSuffix, callLevel)
		pairBlock := fmt.Sprintf("$pair_block_%[1]s_%[2]d", supportSuffix, callLevel)

		result = append(result, fmt.Sprintf("  %[1]s = TL\\tl2_support::fetch_size();", pairCurrentSize))
		result = append(result, fmt.Sprintf("  %[1]s += TL\\tl2_support::count_used_bytes(%[2]s);", pairUsedBytes, pairCurrentSize))
		result = append(result,
			fmt.Sprintf("  if (%[1]s == 0) {", pairCurrentSize),
			"    continue;",
			"  }",
		)

		result = append(result, fmt.Sprintf("  %[1]s += %[2]s;", pairCurrentSize, pairUsedBytes))
		result = append(result, fmt.Sprintf("  %[1]s = fetch_byte();", pairBlock))
		result = append(result,
			fmt.Sprintf("  if ((%[1]s & 1) != 0) {", pairBlock),
			"    $index = TL\\tl2_support::fetch_size();",
			fmt.Sprintf("    %[1]s += TL\\tl2_support::count_used_bytes($index);", pairUsedBytes),
			"    if ($index != 0) {",
			fmt.Sprintf("      %[1]s += TL\\tl2_support::skip_bytes(%[2]s - %[1]s);", pairUsedBytes, pairCurrentSize),
			"      continue;",
			"    }",
			"  }",
		)
		// init key and value
		result = append(result, fmt.Sprintf("  /** @var %[1]s */", trw.dictKeyField.t.trw.PhpTypeName(true, true)))
		result = append(result, fmt.Sprintf("  %[1]s = %[2]s;", keyElement, trw.dictKeyField.t.trw.PhpDefaultInit()))
		result = append(result, fmt.Sprintf("  /** @var %[1]s */", trw.dictValueField.t.trw.PhpTypeName(true, true)))
		result = append(result, fmt.Sprintf("  %[1]s = %[2]s;", valueElement, trw.dictValueField.t.trw.PhpDefaultInit()))
		// read key and value
		result = append(result, fmt.Sprintf("  if ((%[1]s & (1 << 1)) != 0) {", pairBlock))
		result = append(result, utils.ShiftAll(keyRead, "    ")...)
		result = append(result, "  }")
		result = append(result, fmt.Sprintf("  if ((%[1]s & (1 << 2)) != 0) {", pairBlock))
		result = append(result, utils.ShiftAll(valueRead, "    ")...)
		result = append(result, "  }")
		result = append(result,
			fmt.Sprintf("  %[1]s[%[2]s] = %[3]s;", targetName, keyElement, valueElement),
			"}",
		)
		break
	}

	totalRead := make([]string, 0)
	totalRead = append(totalRead,
		// read byte size
		fmt.Sprintf("%[1]s = TL\\tl2_support::fetch_size();", localCurrentSize),
		fmt.Sprintf("%[1]s = 0;", localUsedBytesPointer),
		// add to global pointer
		fmt.Sprintf("%[1]s += %[2]s + TL\\tl2_support::count_used_bytes(%[2]s);", usedBytesPointer, localCurrentSize),
		fmt.Sprintf("if (%[1]s != 0) {", localCurrentSize),
		fmt.Sprintf("  %[1]s = TL\\tl2_support::fetch_size();", localElementCount),
		fmt.Sprintf("  %[1]s += TL\\tl2_support::count_used_bytes(%[2]s);", localUsedBytesPointer, localElementCount),
		// after actual read
	)
	for _, line := range result {
		totalRead = append(totalRead, "  "+line)
	}
	totalRead = append(totalRead,
		fmt.Sprintf("  %[1]s += TL\\tl2_support::skip_bytes(%[2]s - %[1]s);", localUsedBytesPointer, localCurrentSize),
		"}",
	)
	return totalRead
}

func (trw *TypeRWBrackets) readTL2Array(targetName string, supportSuffix string, callLevel int, localElementCount string, localUsedBytesPointer string, result []string, index string, elementName string, elementRead []string) []string {
	if _, ok := trw.element.t.trw.(*TypeRWBool); ok {
		localBytesCount := fmt.Sprintf("$byte_blocks_count_%[1]s_%[2]d", supportSuffix, callLevel)
		localByteIndex := fmt.Sprintf("$byte_index_%[1]s_%[2]d", supportSuffix, callLevel)
		localBitIndex := fmt.Sprintf("$bit_index_%[1]s_%[2]d", supportSuffix, callLevel)

		localByte := fmt.Sprintf("$byte_%[1]s_%[2]d", supportSuffix, callLevel)
		localBooleanPointer := fmt.Sprintf("$boolean_index_%[1]s_%[2]d", supportSuffix, callLevel)
		localBooleanCount := fmt.Sprintf("$boolean_count_%[1]s_%[2]d", supportSuffix, callLevel)

		bytesRead := []string{
			fmt.Sprintf("%[1]s = 0;", localBooleanPointer),
			fmt.Sprintf("%[1]s = (%[2]s + 7) / 8;", localBytesCount, localElementCount),
			fmt.Sprintf("for (%[1]s = 0; %[1]s < %[2]s; %[1]s++) {", localByteIndex, localBytesCount),
			fmt.Sprintf("  %[1]s = fetch_byte();", localByte),
			fmt.Sprintf("  %[1]s += 1;", localUsedBytesPointer),
			fmt.Sprintf("  %[1]s = 8;", localBooleanCount),
			fmt.Sprintf("  if (%[1]s + %[2]s > %[3]s) {", localBooleanPointer, localBooleanCount, localElementCount),
			fmt.Sprintf("    %[1]s = %[2]s - %[3]s;", localBooleanCount, localElementCount, localBooleanPointer),
			"  }",
			fmt.Sprintf("  for (%[1]s = 0; %[1]s < %[2]s; %[1]s++) {", localBitIndex, localBooleanCount),
			fmt.Sprintf("    %[1]s[] = (%[2]s & (1 << %[3]s)) != 0;", targetName, localByte, localBitIndex),
			fmt.Sprintf("    %[1]s += 1;", localBooleanPointer),
			"  }",
			"}",
		}
		result = append(result, utils.ShiftAll(bytesRead, "  ")...)
	} else {
		result = append(result,
			fmt.Sprintf("for(%[1]s = 0; %[1]s < %[2]s; %[1]s++) {", index, localElementCount),
			fmt.Sprintf("  /** @var %[1]s */", trw.element.t.trw.PhpTypeName(true, true)),
			fmt.Sprintf("  %[2]s = %[1]s;", trw.element.t.trw.PhpDefaultInit(), elementName),
		)
		result = append(result, elementRead...)
		result = append(result,
			fmt.Sprintf("  %[1]s[] = %[2]s;", targetName, elementName),
			"}",
		)
	}
	return result
}

func (trw *TypeRWBrackets) PhpWriteTL2MethodCall(targetName string, bare bool, args *TypeArgumentsTree, supportSuffix string, callLevel int, usedBytesPointer string, canDependOnLocalBit bool) []string {
	panic("TODO TUPLES")
}

func (trw *TypeRWBrackets) PhpCalculateSizesTL2MethodCall(targetName string, bare bool, args *TypeArgumentsTree, supportSuffix string, callLevel int, canDependOnLocalBit bool) []string {
	panic("TODO TUPLES")
}

func (trw *TypeRWBrackets) PhpDefaultInit() string {
	return "[]"
}
