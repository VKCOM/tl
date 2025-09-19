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

func (trw *TypeRWBool) PhpClassName(withPath bool, bare bool) string {
	return "boolean"
}

func (trw *TypeRWBool) PhpClassNameReplaced() bool {
	return true
}

func (trw *TypeRWBool) PhpTypeName(withPath bool, bare bool) string {
	return trw.PhpClassName(withPath, true)
}

func (trw *TypeRWBool) PhpGenerateCode(code *strings.Builder, bytes bool) error {
	return fmt.Errorf("boolean doesn't have php code")
}

func (trw *TypeRWBool) PhpDefaultValue() string {
	return "false"
}

func (trw *TypeRWBool) PhpIterateReachableTypes(reachableTypes *map[*TypeRWWrapper]bool) {
}

func (trw *TypeRWBool) PhpReadMethodCall(targetName string, bare bool, initIfDefault bool, args *TypeArgumentsTree, supportSuffix string) []string {
	if !bare {
		if trw.wr.gen.options.UseBuiltinDataProviders {
			return []string{
				"$magic = fetch_int() & 0xFFFFFFFF;",
				fmt.Sprintf("if ($magic == 0x%08[1]x) {", trw.falseTag),
				fmt.Sprintf("  %[1]s = false;", targetName),
				fmt.Sprintf("} elseif ($magic == 0x%08[1]x) {", trw.trueTag),
				fmt.Sprintf("  %[1]s = true;", targetName),
				"} else {",
				"  return false;",
				"}",
			}
		} else {
			return []string{
				fmt.Sprintf(
					"[%[1]s, $success] = $stream->read_bool(0x%08[2]x, 0x%08[3]x);",
					targetName,
					trw.falseTag,
					trw.trueTag,
				),
				"if (!$success) {",
				"  return false;",
				"}",
			}
		}
	}
	return nil
}

func (trw *TypeRWBool) PhpWriteMethodCall(targetName string, bare bool, args *TypeArgumentsTree, supportSuffix string) []string {
	if !bare {
		if trw.wr.gen.options.UseBuiltinDataProviders {
			return []string{
				fmt.Sprintf("if (%[1]s) {", targetName),
				fmt.Sprintf("  store_int(0x%08[1]x);", trw.trueTag),
				"} else {",
				fmt.Sprintf("  store_int(0x%08[1]x);", trw.falseTag),
				"}",
			}
		} else {
			return []string{
				fmt.Sprintf(
					"$success = $stream->write_bool(%[1]s, 0x%08[2]x, 0x%08[3]x);",
					targetName,
					trw.falseTag,
					trw.trueTag,
				),
				"if (!$success) {",
				"  return false;",
				"}",
			}
		}
	}
	return nil
}

func (trw *TypeRWBool) PhpDefaultInit() string {
	return "false"
}
