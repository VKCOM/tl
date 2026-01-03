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

func (trw *TypeRWPrimitive) PhpReadMethodCall(targetName string, bare bool, initIfDefault bool, args *TypeArgumentsTree, supportSuffix string) []string {
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

func (trw *TypeRWPrimitive) PhpWriteMethodCall(targetName string, bare bool, args *TypeArgumentsTree, supportSuffix string) []string {
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

func (trw *TypeRWPrimitive) PhpReadTL2MethodCall(targetName string, bare bool, initIfDefault bool, args *TypeArgumentsTree, supportSuffix string, callLevel int, usedBytesPointer string, canDependOnLocalBit bool) []string {
	if trw.gen.options.UseBuiltinDataProviders {
		switch trw.goType {
		case "int32":
			return []string{
				fmt.Sprintf("%s = fetch_int();", targetName),
				fmt.Sprintf("%s += 4;", usedBytesPointer),
			}
		case "int64":
			return []string{
				fmt.Sprintf("%s = fetch_long();", targetName),
				fmt.Sprintf("%s += 8;", usedBytesPointer),
			}
		case "uint32":
			return []string{
				fmt.Sprintf("%s = fetch_int() & 0xFFFFFFFF;", targetName),
				fmt.Sprintf("%s += 4;", usedBytesPointer),
			}
		case "uint64":
			return []string{
				fmt.Sprintf("%s = fetch_long() & 0xFFFFFFFFFFFFFFFF;", targetName),
				fmt.Sprintf("%s += 8;", usedBytesPointer),
			}
		case "float32":
			return []string{
				fmt.Sprintf("%s = fetch_float();", targetName),
				fmt.Sprintf("%s += 4;", usedBytesPointer),
			}
		case "float64":
			return []string{
				fmt.Sprintf("%s = fetch_double();", targetName),
				fmt.Sprintf("%s += 8;", usedBytesPointer),
			}
		case "string":
			return []string{
				fmt.Sprintf("%s = fetch_string2();", targetName),
				fmt.Sprintf("%[1]s += strlen(%[2]s) + TL\\tl2_support::count_used_bytes(strlen(%[2]s));", usedBytesPointer, targetName),
			}
		default:
			panic(fmt.Sprintf("unsupported primitive generation for %[1]s in php", trw.goType))
		}
	} else {
		panic("unsupported generation for primitive in php")
	}
}

func (trw *TypeRWPrimitive) PhpWriteTL2MethodCall(targetName string, bare bool, args *TypeArgumentsTree, supportSuffix string, callLevel int, usedBytesPointer string, canDependOnLocalBit bool) []string {
	if trw.gen.options.UseBuiltinDataProviders {
		switch trw.goType {
		case "int32":
			return []string{
				fmt.Sprintf("store_int(%s);", targetName),
				fmt.Sprintf("%s += 4;", usedBytesPointer),
			}
		case "int64":
			return []string{
				fmt.Sprintf("store_long(%s);", targetName),
				fmt.Sprintf("%s += 8;", usedBytesPointer),
			}
		case "uint32":
			return []string{
				fmt.Sprintf("store_int(%s);", targetName),
				fmt.Sprintf("%s += 4;", usedBytesPointer),
			}
		case "uint64":
			return []string{
				fmt.Sprintf("store_long(%s);", targetName),
				fmt.Sprintf("%s += 8;", usedBytesPointer),
			}
		case "float32":
			return []string{
				fmt.Sprintf("store_float(%s);", targetName),
				fmt.Sprintf("%s += 4;", usedBytesPointer),
			}
		case "float64":
			return []string{
				fmt.Sprintf("store_double(%s);", targetName),
				fmt.Sprintf("%s += 8;", usedBytesPointer),
			}
		case "string":
			return []string{
				fmt.Sprintf("store_string2(%s);", targetName),
				fmt.Sprintf("%[1]s += strlen(%[2]s) + TL\\tl2_support::count_used_bytes(strlen(%[2]s));", usedBytesPointer, targetName),
			}
		}
	}
	panic("unsupported generation for primitive in php")
}

func (trw *TypeRWPrimitive) PhpCalculateSizesTL2MethodCall(targetName string, bare bool, args *TypeArgumentsTree, supportSuffix string, callLevel int, usedBytesPointer string) []string {
	if trw.gen.options.UseBuiltinDataProviders {
		switch trw.goType {
		case "int32":
			return []string{
				fmt.Sprintf("%s += 4;", usedBytesPointer),
			}
		case "int64":
			return []string{
				fmt.Sprintf("%s += 8;", usedBytesPointer),
			}
		case "uint32":
			return []string{
				fmt.Sprintf("%s += 4;", usedBytesPointer),
			}
		case "uint64":
			return []string{
				fmt.Sprintf("%s += 8;", usedBytesPointer),
			}
		case "float32":
			return []string{
				fmt.Sprintf("%s += 4;", usedBytesPointer),
			}
		case "float64":
			return []string{
				fmt.Sprintf("%s += 8;", usedBytesPointer),
			}
		case "string":
			return []string{
				fmt.Sprintf("%[1]s += strlen(%[2]s) + TL\\tl2_support::count_used_bytes(strlen(%[2]s));", usedBytesPointer, targetName),
			}
		}
	}
	panic("unsupported generation for primitive in php")
}

func (trw *TypeRWPrimitive) PhpDefaultInit() string {
	return trw.PhpDefaultValue()
}

func (trw *TypeRWWrapper) PhpTL2TrivialSize(targetObject string, canDependOnLocalBit bool, refObject bool) (isConstant bool, size string) {
	isConstant, size = trw.trw.tl2TrivialSize(targetObject, canDependOnLocalBit, refObject)
	if pr, isPr := trw.trw.(*TypeRWPrimitive); isPr && pr.goType == "string" {
		size = fmt.Sprintf("strlen(%[1]s)", targetObject)
	}
	return
}
