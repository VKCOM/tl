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

func (trw *TypeRWPrimitive) PhpClassName(withPath bool, bare bool) string {
	switch trw.goType {
	case "int32", "int64", "uint32":
		return "int"
	case "string":
		return "string"
	case "float32", "float64":
		return "float"
	default:
		return fmt.Sprintf("<? %s>", trw.tlType)
	}
}

func (trw *TypeRWPrimitive) PhpClassNameReplaced() bool {
	return true
}

func (trw *TypeRWPrimitive) PhpTypeName(withPath bool, bare bool) string {
	return trw.PhpClassName(withPath, true)
}

func (trw *TypeRWPrimitive) PhpGenerateCode(code *strings.Builder, bytes bool) error {
	return fmt.Errorf("primitives don't have php code")
}

func (trw *TypeRWPrimitive) PhpDefaultValue() string {
	switch trw.goType {
	case "int32", "int64", "uint32":
		return "0"
	case "string":
		return "''"
	case "float32", "float64":
		return "0.0"
	default:
		return fmt.Sprintf("<? %s>", trw.tlType)
	}
}

func (trw *TypeRWPrimitive) PhpIterateReachableTypes(reachableTypes *map[*TypeRWWrapper]bool) {}

func (trw *TypeRWPrimitive) phpIOMethodsSuffix() string {
	switch trw.goType {
	case "int32", "int64", "uint32", "string":
		return trw.goType
	case "float32":
		return "float"
	case "float64":
		return "double"
	default:
		return fmt.Sprintf("<? %s>", trw.tlType)
	}
}

func (trw *TypeRWPrimitive) PhpReadMethodCall(targetName string, bare bool, initIfDefault bool, args *TypeArgumentsTree) []string {
	if !bare {
		panic("can't be boxed")
	}
	if trw.gen.options.UseBuiltinDataProviders {
		switch trw.goType {
		case "int32":
			return []string{fmt.Sprintf("%s = fetch_int();", targetName)}
		case "int64":
			return []string{fmt.Sprintf("%s = fetch_long();", targetName)}
		case "uint32":
			return []string{fmt.Sprintf("%s = fetch_int() & 0xFFFFFFFF;", targetName)}
		case "uint64":
			return []string{fmt.Sprintf("%s = fetch_long() & 0xFFFFFFFFFFFFFFFF;", targetName)}
		case "float32":
			return []string{fmt.Sprintf("%s = fetch_float();", targetName)}
		case "float64":
			return []string{fmt.Sprintf("%s = fetch_double();", targetName)}
		case "string":
			return []string{fmt.Sprintf("%s = fetch_string();", targetName)}
		default:
			return []string{fmt.Sprintf("<? %s>", trw.tlType)}
		}
	}
	return []string{
		fmt.Sprintf("[%[1]s, $success] = $stream->read_%[2]s();", targetName, trw.phpIOMethodsSuffix()),
		"if (!$success) {",
		"  return false;",
		"}",
	}
}

func (trw *TypeRWPrimitive) PhpWriteMethodCall(targetName string, bare bool, args *TypeArgumentsTree) []string {
	if !bare {
		panic("can't be boxed")
	}
	if trw.gen.options.UseBuiltinDataProviders {
		switch trw.goType {
		case "int32":
			return []string{fmt.Sprintf("store_int(%s);", targetName)}
		case "int64":
			return []string{fmt.Sprintf("store_long(%s);", targetName)}
		case "uint32":
			return []string{fmt.Sprintf("store_int(%s);", targetName)}
		case "uint64":
			return []string{fmt.Sprintf("store_long(%s);", targetName)}
		case "float32":
			return []string{fmt.Sprintf("store_float(%s);", targetName)}
		case "float64":
			return []string{fmt.Sprintf("store_double(%s);", targetName)}
		case "string":
			return []string{fmt.Sprintf("store_string(%s);", targetName)}
		default:
			return []string{fmt.Sprintf("<? %s>", trw.tlType)}
		}
	}
	return []string{
		fmt.Sprintf("$success = $stream->write_%[2]s(%[1]s);", targetName, trw.phpIOMethodsSuffix()),
		"if (!$success) {",
		"  return false;",
		"}",
	}
}

func (trw *TypeRWPrimitive) PhpDefaultInit() string {
	return trw.PhpDefaultValue()
}
