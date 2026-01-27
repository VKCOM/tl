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

func (trw *TypeRWMaybe) PhpClassName(withPath bool, bare bool) string {
	target := trw.getInnerTarget()
	return "maybe_" + target.t.trw.PhpClassName(withPath, target.bare)
}

func (trw *TypeRWMaybe) PhpClassNameReplaced() bool {
	return true
}

func (trw *TypeRWMaybe) PhpTypeName(withPath bool, bare bool) string {
	target := trw.getInnerTarget()
	return target.t.trw.PhpTypeName(withPath, target.t.PHPIsBare()) + "|null"
}

func (trw *TypeRWMaybe) getInnerTarget() Field {
	if inner, ok := trw.element.t.trw.(*TypeRWMaybe); ok {
		return inner.getInnerTarget()
	} else {
		return trw.element
	}
}

func (trw *TypeRWMaybe) PhpGenerateCode(code *strings.Builder, bytes bool) error {
	return fmt.Errorf("maybe doesn't have php code")
}

func (trw *TypeRWMaybe) PhpDefaultValue() string {
	return "null"
}

func (trw *TypeRWMaybe) PhpIterateReachableTypes(reachableTypes *map[*TypeRWWrapper]bool) {
	trw.element.t.PhpIterateReachableTypes(reachableTypes)
}

func (trw *TypeRWMaybe) PhpReadMethodCall(targetName string, bare bool, initIfDefault bool, args *TypeArgumentsTree, supportSuffix string) []string {
	if !bare {
		maybeContainsValueName := fmt.Sprintf("$maybeContainsValue_%[1]s", supportSuffix)
		var result []string
		if trw.wr.gen.options.UseBuiltinDataProviders {
			result = append(result,
				"/** @var bool */",
				fmt.Sprintf("%[1]s = false;", maybeContainsValueName),
				"$magic = fetch_int() & 0xFFFFFFFF;",
				fmt.Sprintf("if ($magic == 0x%08[1]x) {", trw.emptyTag),
				fmt.Sprintf("  %[1]s = false;", maybeContainsValueName),
				fmt.Sprintf("} elseif ($magic == 0x%08[1]x) {", trw.okTag),
				fmt.Sprintf("  %[1]s = true;", maybeContainsValueName),
				"} else {",
				"  return false;",
				"}",
			)
		} else {
			result = append(result,
				fmt.Sprintf(
					"[%[3]s, $success] = $stream->read_bool(0x%08[1]x, 0x%08[2]x);",
					trw.emptyTag,
					trw.okTag,
					maybeContainsValueName,
				),
				"if (!$success) {",
				"  return false;",
				"}",
			)
		}
		result = append(result, fmt.Sprintf("if (%[1]s) {", maybeContainsValueName))

		if trw.element.t == trw.getInnerTarget().t && initIfDefault {
			result = append(result,
				fmt.Sprintf("  if (is_null(%[1]s)) {", targetName),
				fmt.Sprintf("    %[1]s = %[2]s;", targetName, trw.element.t.trw.PhpDefaultInit()),
				"  }",
			)
			initIfDefault = false
		}
		var newArgs *TypeArgumentsTree
		if args != nil {
			newArgs = args.children[0]
		}
		bodyReader := trw.element.t.trw.PhpReadMethodCall(targetName, trw.element.bare, initIfDefault, newArgs, supportSuffix)
		for i := range bodyReader {
			bodyReader[i] = "  " + bodyReader[i]
		}
		result = append(result, bodyReader...)
		result = append(result,
			"} else {",
			fmt.Sprintf("  %[1]s = null;", targetName),
			"}",
		)
		return result
	}
	return nil
}

func (trw *TypeRWMaybe) PhpWriteMethodCall(targetName string, bare bool, args *TypeArgumentsTree, supportSuffix string) []string {
	if !bare {
		var result []string
		if trw.wr.gen.options.UseBuiltinDataProviders {
			result = append(result,
				fmt.Sprintf("if (is_null(%[1]s)) {", targetName),
				fmt.Sprintf("  store_int(0x%08[1]x);", trw.emptyTag),
				"} else {",
				fmt.Sprintf("  store_int(0x%08[1]x);", trw.okTag),
				"}",
			)
		} else {
			result = append(result,
				fmt.Sprintf(
					"$success = $stream->write_bool(!is_null(%[1]s), 0x%08[2]x, 0x%08[3]x);",
					targetName,
					trw.emptyTag,
					trw.okTag,
				),
				"if (!$success) {",
				"  return false;",
				"}",
			)
		}
		result = append(result,
			fmt.Sprintf("if (!is_null(%[1]s)) {", targetName),
		)
		{
			var newArgs *TypeArgumentsTree
			if args != nil {
				newArgs = args.children[0]
			}
			bodyWriter := trw.element.t.trw.PhpWriteMethodCall(targetName, trw.element.bare, newArgs, supportSuffix)
			for i := range bodyWriter {
				bodyWriter[i] = "  " + bodyWriter[i]
			}
			result = append(result, bodyWriter...)
		}
		result = append(result,
			"}",
		)
		return result
	}
	return nil
}

func (trw *TypeRWMaybe) PhpReadTL2MethodCall(targetName string, bare bool, initIfDefault bool, args *TypeArgumentsTree, supportSuffix string, callLevel int, usedBytesPointer string, canDependOnLocalBit bool) []string {
	maybeContainsValueName := fmt.Sprintf("$maybe_contains_value_%[1]s_%[2]d", supportSuffix, callLevel)
	localUsedBytesPointer := fmt.Sprintf("$used_bytes_%[1]s_%[2]d", supportSuffix, callLevel)
	localCurrentSize := fmt.Sprintf("$current_size_%[1]s_%[2]d", supportSuffix, callLevel)
	localBlock := fmt.Sprintf("$block_%[1]s_%[2]d", supportSuffix, callLevel)

	var result []string
	// investigate read necessity
	if trw.wr.gen.options.UseBuiltinDataProviders {
		result = append(result,
			"/** @var bool */",
			fmt.Sprintf("%[1]s = false;", maybeContainsValueName),
			fmt.Sprintf("%[1]s = TL\\tl2_support::fetch_size();", localCurrentSize),
			fmt.Sprintf("%[1]s = 0;", localUsedBytesPointer),
			// add to global pointer
			fmt.Sprintf("%[1]s += %[2]s + TL\\tl2_support::count_used_bytes(%[2]s);", usedBytesPointer, localCurrentSize),
			// decide should we read body
			fmt.Sprintf("if (%[1]s != 0) {", localCurrentSize),
			fmt.Sprintf("  %[1]s = fetch_byte();", localBlock),
			fmt.Sprintf("  %[1]s += 1;", localUsedBytesPointer),
			fmt.Sprintf("  if (%[1]s & 1 != 0) {", localBlock),
			fmt.Sprintf("    %[1]s = (fetch_byte() == 1);", maybeContainsValueName),
			fmt.Sprintf("    %[1]s += 1;", localUsedBytesPointer),
			"  }",
			fmt.Sprintf("  if (%[1]s) {", maybeContainsValueName),
			fmt.Sprintf("    %[1]s = ((%[2]s & (1 << 1)) != 0);", maybeContainsValueName, localBlock),
			"  }",
			"}",
		)
	} else {
		panic("unsupported generation for maybe in php")
	}
	// read inner
	result = append(result, fmt.Sprintf("if (%[1]s) {", maybeContainsValueName))
	if trw.element.t == trw.getInnerTarget().t && initIfDefault {
		result = append(result,
			fmt.Sprintf("  if (is_null(%[1]s)) {", targetName),
			fmt.Sprintf("    %[1]s = %[2]s;", targetName, trw.element.t.trw.PhpDefaultInit()),
			"  }",
		)
		initIfDefault = false
	}
	var newArgs *TypeArgumentsTree
	if args != nil {
		newArgs = args.children[0]
	}
	bodyReader := trw.element.t.trw.PhpReadTL2MethodCall(targetName, trw.element.bare, initIfDefault, newArgs, supportSuffix, callLevel+1, localUsedBytesPointer, false)
	for i := range bodyReader {
		bodyReader[i] = "  " + bodyReader[i]
	}
	result = append(result, bodyReader...)
	result = append(result,
		"} else {",
		fmt.Sprintf("  %[1]s = null;", targetName),
		"}",
	)
	// skip rest
	result = append(result,
		fmt.Sprintf("%[1]s += TL\\tl2_support::skip_bytes(%[2]s - %[1]s);", localUsedBytesPointer, localCurrentSize),
	)
	return result

	return nil
}

func (trw *TypeRWMaybe) PhpWriteTL2MethodCall(targetName string, bare bool, args *TypeArgumentsTree, supportSuffix string, callLevel int, usedBytesPointer string, canDependOnLocalBit bool) []string {
	localCurrentSize := fmt.Sprintf("$current_size_%[1]s_%[2]d", supportSuffix, callLevel)
	localBlock := fmt.Sprintf("$block_%[1]s_%[2]d", supportSuffix, callLevel)
	localUsedSizePointer := fmt.Sprintf("$used_size_%[1]s_%[2]d", supportSuffix, callLevel)

	result := make([]string, 0)
	result = append(result,
		fmt.Sprintf("%[1]s = $context_sizes->pop_front();", localCurrentSize),
	)
	if usedBytesPointer != "" {
		result = append(result,
			fmt.Sprintf("%[1]s += %[2]s;", usedBytesPointer, localCurrentSize),
		)
	}
	result = append(result,
		fmt.Sprintf("TL\\tl2_support::store_size(%[1]s);", localCurrentSize),
		fmt.Sprintf("if (%[1]s != 0) {", localCurrentSize),
	)

	// write inner part
	innerPart := make([]string, 0)
	innerPart = append(innerPart,
		fmt.Sprintf("if (is_null(%[1]s)) {", targetName),
		`  throw new \Exception("inner element is null but object size != 0");`,
		"}",
		fmt.Sprintf("%[1]s = $context_blocks->pop_front();", localBlock),
		fmt.Sprintf("store_byte(%[1]s);", localBlock),
		fmt.Sprintf("if ((%[1]s & (1 << 0)) != 0) {", localBlock),
		"  store_byte(1);",
		"}",
		fmt.Sprintf("if ((%[1]s & (1 << 1)) != 0) {", localBlock),
	)

	var newArgs *TypeArgumentsTree
	if args != nil {
		newArgs = args.children[0]
	}

	innerPart = append(innerPart, fmt.Sprintf("%[1]s = 0;", localUsedSizePointer))
	innerPart = append(innerPart, utils.ShiftAll(trw.element.t.trw.PhpWriteTL2MethodCall(targetName, bare, newArgs, supportSuffix, callLevel+1, localUsedSizePointer, false), "  ")...)
	innerPart = append(innerPart, "}")

	// add it with shift
	result = append(result, utils.ShiftAll(innerPart, "  ")...)
	result = append(result, "}")
	return result
}

func (trw *TypeRWMaybe) PhpCalculateSizesTL2MethodCall(targetName string, bare bool, args *TypeArgumentsTree, supportSuffix string, callLevel int, usedBytesPointer string) []string {
	localCurrentSize := fmt.Sprintf("$current_size_%[1]s_%[2]d", supportSuffix, callLevel)
	localBlock := fmt.Sprintf("$block_%[1]s_%[2]d", supportSuffix, callLevel)

	result := make([]string, 0)
	result = append(result,
		fmt.Sprintf("%s = 0;", localCurrentSize),
		fmt.Sprintf("%s_index = $context_sizes->push_back(0);", localCurrentSize),
		fmt.Sprintf("if (!is_null(%[1]s)) {", targetName),
	)

	// write inner part
	innerPart := make([]string, 0)
	innerPart = append(innerPart,
		fmt.Sprintf("%[1]s = (1 << 0) + (1 << 1);", localBlock),
		fmt.Sprintf("%[1]s_index = $context_blocks->push_back(%[1]s);", localBlock),
		fmt.Sprintf("%[1]s += 1 + 1;", localCurrentSize),
	)

	var newArgs *TypeArgumentsTree
	if args != nil {
		newArgs = args.children[0]
	}
	innerPart = append(innerPart, trw.element.t.trw.PhpCalculateSizesTL2MethodCall(targetName, bare, newArgs, supportSuffix, callLevel+1, usedBytesPointer)...)
	isSizeConstant, trivialSize := trw.element.t.PhpTL2TrivialSize(targetName, false, trw.element.recursive)
	sizeValue := fmt.Sprintf("$context_sizes->get_value(%[1]s_index)", localCurrentSize)
	if len(trivialSize) != 0 {
		sizeValue = trivialSize
	}

	innerTab := ""
	if !isSizeConstant {
		innerPart = append(innerPart,
			fmt.Sprintf("if (%[1]s != 0) {", sizeValue),
		)
		innerTab = "  "
	}
	innerPart = append(innerPart,
		utils.ShiftAll([]string{
			fmt.Sprintf("%[1]s += %[2]s;", localCurrentSize, sizeValue),
		}, innerTab)...,
	)
	if trw.element.t.trw.isSizeWrittenInData() {
		innerPart = append(innerPart,
			utils.ShiftAll([]string{
				fmt.Sprintf("%[1]s += TL\\tl2_support::count_used_bytes(%[2]s);", localCurrentSize, sizeValue),
			}, innerTab)...,
		)
	}
	if !isSizeConstant {
		innerPart = append(innerPart,
			"}",
		)
	}
	// add it with shift
	result = append(result, utils.ShiftAll(innerPart, "  ")...)

	result = append(result, "}")
	// add actual size
	result = append(result, fmt.Sprintf("$context_sizes->set_value(%[1]s_index, %[1]s);", localCurrentSize))

	return result
}

func (trw *TypeRWMaybe) PhpDefaultInit() string {
	return trw.element.t.trw.PhpDefaultInit()
}
