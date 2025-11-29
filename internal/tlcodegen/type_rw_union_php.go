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

func (trw *TypeRWUnion) PhpClassNameReplaced() bool {
	return false
}

func (trw *TypeRWUnion) PhpClassName(withPath bool, bare bool) string {
	if specialCase := PHPSpecialMembersTypes(trw.wr); specialCase != "" {
		return specialCase
	}
	name := trw.wr.tlName.Name
	if len(trw.wr.tlName.Namespace) != 0 {
		name = fmt.Sprintf("%s_%s", trw.wr.tlName.Namespace, name)
	}

	elems := make([]string, 0, len(trw.wr.arguments))
	for _, arg := range trw.wr.arguments {
		if arg.tip != nil {
			argText := arg.tip.trw.PhpClassName(false, false)
			if argText != "" {
				elems = append(elems, "__", argText)
			}
		}
	}

	name += strings.Join(elems, "")
	if withPath {
		name = trw.wr.PHPTypePath() + name
	}
	return name
}

func (trw *TypeRWUnion) PhpTypeName(withPath bool, bare bool) string {
	return trw.PhpClassName(withPath, true)
}

func (trw *TypeRWUnion) PhpGenerateCode(code *strings.Builder, bytes bool) error {
	return PhpGenerateInterfaceCode(code, bytes, trw.wr, utils.MapSlice(trw.Fields, func(f Field) *TypeRWWrapper { return f.t }))
}

func PhpGenerateInterfaceCode(code *strings.Builder, bytes bool, targetType *TypeRWWrapper, itsConstructors []*TypeRWWrapper) error {
	constructors := make([]string, len(itsConstructors))
	for i, constructor := range itsConstructors {
		constructors[i] = fmt.Sprintf("%s::class", constructor.trw.PhpClassName(true, true))
	}

	code.WriteString(`
use VK\TL;

/**
 * @kphp-tl-class
 */
`)
	code.WriteString(fmt.Sprintf(
		`interface %[1]s {

  /** Allows kphp implicitly load all available constructors */
  const CONSTRUCTORS = [
    %[2]s
  ];
%[3]s
}
`,
		targetType.trw.PhpClassName(false, false),
		strings.Join(constructors, ",\n    "),
		phpGenerateIOBoxedMethodsForInterface(bytes, targetType),
	))

	return nil
}

func phpGenerateIOBoxedMethodsForInterface(bytes bool, targetType *TypeRWWrapper) string {
	useBuiltin := targetType.gen.options.UseBuiltinDataProviders
	if !(targetType.gen.options.AddFunctionBodies && targetType.phpInfo.RequireFunctionBodies) {
		return ""
	}

	readArgNames := make([]string, 0)
	readArgTypes := make([]string, 0)
	writeArgNames := make([]string, 0)
	writeArgTypes := make([]string, 0)

	if !useBuiltin {
		readArgNames = append(readArgNames, "stream")
		readArgTypes = append(readArgTypes, `TL\tl_input_stream`)

		writeArgNames = append(writeArgNames, "stream")
		writeArgTypes = append(writeArgTypes, `TL\tl_output_stream`)
	}

	for _, name := range targetType.PHPGetNatTypeDependenciesDeclAsArray() {
		name, _ = strings.CutPrefix(name, "$")
		readArgNames = append(readArgNames, name)
		readArgTypes = append(readArgTypes, "int")

		writeArgNames = append(writeArgNames, name)
		writeArgTypes = append(writeArgTypes, "int")
	}

	ioCode := ""
	ioCode += fmt.Sprintf(`
%[1]s
  public function read_boxed(%[2]s);

%[3]s
  public function write_boxed(%[4]s);`,
		phpFunctionCommentFormat(
			readArgNames,
			readArgTypes,
			"bool",
			"  ",
		),
		phpFunctionArgumentsFormat(readArgNames),
		phpFunctionCommentFormat(
			writeArgNames,
			writeArgTypes,
			"bool",
			"  ",
		),
		phpFunctionArgumentsFormat(writeArgNames),
	)

	if targetType.wantsTL2 {
		ioCode += fmt.Sprintf(`
%[5]s
  public function read_tl2(%[6]s);

%[1]s
  public function write_tl2(%[2]s);

%[3]s
  public function internal_write_tl2(%[4]s);

%[3]s
  public function calculate_sizes_tl2(%[4]s);`,
			phpFunctionCommentFormat(
				readArgNames,
				readArgTypes,
				"",
				"  ",
			),
			phpFunctionArgumentsFormat(readArgNames),
			phpFunctionCommentFormat(
				utils.Append(writeArgNames, "context_sizes", "context_blocks"),
				utils.Append(writeArgTypes, "TL\\tl2_context", "TL\\tl2_context"),
				"",
				"  ",
			),
			phpFunctionArgumentsFormat(utils.Append(writeArgNames, "context_sizes", "context_blocks")),
			phpFunctionCommentFormat(
				writeArgNames,
				writeArgTypes,
				"",
				"  ",
			),
			phpFunctionArgumentsFormat(writeArgNames),
		)
	}

	return ioCode
}

func (trw *TypeRWUnion) PhpDefaultValue() string {
	return "null"
}

func (trw *TypeRWUnion) PhpIterateReachableTypes(reachableTypes *map[*TypeRWWrapper]bool) {
	for _, field := range trw.Fields {
		field.t.PhpIterateReachableTypes(reachableTypes)
	}
}

func (trw *TypeRWUnion) PhpReadMethodCall(targetName string, bare bool, initIfDefault bool, args *TypeArgumentsTree, supportSuffix string) []string {
	if bare {
		panic("union can't be bare")
	}
	variantName := func(tag uint32, index int) string {
		return fmt.Sprintf("$variant0x%08x", tag)
	}
	var result []string
	if trw.wr.gen.options.UseBuiltinDataProviders {
		result = append(result,
			"$tag = fetch_int() & 0xFFFFFFFF;",
			"switch ($tag) {",
		)
		for i, field := range trw.Fields {
			curType := field.t
			name := variantName(field.t.tlTag, i)
			result = append(result,
				fmt.Sprintf("  case 0x%08[1]x:", curType.tlTag),
				fmt.Sprintf("    %[2]s = new %[1]s();", curType.trw.PhpTypeName(true, true), name),
				fmt.Sprintf("    $success = %[2]s->read(%[1]s);", phpFormatArgs(args.ListAllValues(), true), name),
				"    if (!$success) {",
				"      return false;",
				"    }",
				fmt.Sprintf("    %[1]s = %[2]s;", targetName, name),
				"    break;",
			)
		}
		result = append(result,
			"  default:",
			"    return false;",
			"}",
		)
	} else {
		result = append(result,
			"[$tag, $success] = $stream->read_uint32();",
			"if (!$success) {",
			"  return false;",
			"}",
			"switch ($tag) {",
		)
		for i, field := range trw.Fields {
			curType := field.t
			name := variantName(field.t.tlTag, i)
			result = append(result,
				fmt.Sprintf("  case 0x%08[1]x:", curType.tlTag),
				fmt.Sprintf("    %[2]s = new %[1]s();", curType.trw.PhpTypeName(true, true), name),
				fmt.Sprintf("    $success = %[2]s->read($stream%[1]s);", phpFormatArgs(args.ListAllValues(), false), name),
				"    if (!$success) {",
				"      return false;",
				"    }",
				fmt.Sprintf("    %[1]s = %[2]s;", targetName, name),
				"    break;",
			)
		}
		result = append(result,
			"  default:",
			"    return false;",
			"}",
		)
	}
	return result
}

func (trw *TypeRWUnion) PhpWriteMethodCall(targetName string, bare bool, args *TypeArgumentsTree, supportSuffix string) []string {
	if bare {
		panic("union can't be bare")
	}
	var result []string
	if trw.wr.gen.options.UseBuiltinDataProviders {
		result = append(result,
			fmt.Sprintf("if (is_null(%[1]s)) {", targetName),
			fmt.Sprintf("  %[1]s = %[2]s;", targetName, trw.PhpDefaultInit()),
			"}",
			fmt.Sprintf("$success = %[1]s->write_boxed(%[2]s);", targetName, phpFormatArgs(args.ListAllValues(), true)),
			"if (!$success) {",
			"  return false;",
			"}",
		)
	} else {
		result = append(result,
			fmt.Sprintf("if (is_null(%[1]s)) {", targetName),
			fmt.Sprintf("  %[1]s = %[2]s;", targetName, trw.PhpDefaultInit()),
			"}",
			fmt.Sprintf("$success = %[1]s->write_boxed($stream%[2]s);", targetName, phpFormatArgs(args.ListAllValues(), false)),
			"if (!$success) {",
			"  return false;",
			"}",
		)
	}

	return result
}

func (trw *TypeRWUnion) PhpReadTL2MethodCall(targetName string, bare bool, initIfDefault bool, args *TypeArgumentsTree, supportSuffix string, callLevel int, usedBytesPointer string, canDependOnLocalBit bool) []string {
	localUsedBytesPointer := fmt.Sprintf("$used_bytes_%[1]s_%[2]d", supportSuffix, callLevel)
	localCurrentSize := fmt.Sprintf("$current_size_%[1]s_%[2]d", supportSuffix, callLevel)
	localBlock := fmt.Sprintf("$block_%[1]s_%[2]d", supportSuffix, callLevel)
	localConstructor := fmt.Sprintf("$index_%[1]s_%[2]d", supportSuffix, callLevel)

	variantName := func(index int) string {
		return fmt.Sprintf("$variant_%[1]s_%[2]d_%[3]d", supportSuffix, callLevel, index)
	}
	var result []string
	if trw.wr.gen.options.UseBuiltinDataProviders {
		result = append(result,
			fmt.Sprintf("%[1]s = 0;", localConstructor),
			fmt.Sprintf("%[1]s = TL\\tl2_support::fetch_size();", localCurrentSize),
			fmt.Sprintf("%[1]s = 0;", localUsedBytesPointer),
			// add to global pointer
			fmt.Sprintf("%[1]s += %[2]s + TL\\tl2_support::count_used_bytes(%[2]s);", usedBytesPointer, localCurrentSize),
			// decide should we read body
			fmt.Sprintf("if (%[1]s != 0) {", localCurrentSize),
			fmt.Sprintf("  %[1]s = fetch_byte();", localBlock),
			fmt.Sprintf("  %[1]s += 1;", localUsedBytesPointer),
			fmt.Sprintf("  if (%[1]s & 1 != 0) {", localBlock),
			fmt.Sprintf("    %[1]s = TL\\tl2_support::fetch_size();", localConstructor),
			fmt.Sprintf("    %[1]s += TL\\tl2_support::count_used_bytes(%[2]s);", localUsedBytesPointer, localConstructor),
			"  }",
			"}",
			"// check variants",
			fmt.Sprintf("switch (%[1]s) {", localConstructor),
		)
		// iterate variants
		for i, field := range trw.Fields {
			curType := field.t
			name := variantName(i)
			result = append(result,
				fmt.Sprintf("  case %[1]d:", i),
				fmt.Sprintf("    %[2]s = new %[1]s();", curType.trw.PhpTypeName(true, true), name),
				fmt.Sprintf("    %[2]s->read_tl2(%[1]s);", phpFormatArgs(append(args.ListAllValues(), localBlock, fmt.Sprintf("%[1]s - %[2]s", localCurrentSize, localUsedBytesPointer)), true), name),
				fmt.Sprintf("    %[1]s = %[2]s;", targetName, name),
				"    break;",
			)
		}
		// end
		result = append(result,
			"}",
		)
	} else {
		panic("unsupported generation for union in php")
	}
	return result
}

func (trw *TypeRWUnion) PhpWriteTL2MethodCall(targetName string, bare bool, args *TypeArgumentsTree, supportSuffix string, callLevel int, usedBytesPointer string, canDependOnLocalBit bool) []string {
	panic("TODO UNION")
}

func (trw *TypeRWUnion) PhpDefaultInit() string {
	return trw.Fields[0].t.trw.PhpDefaultInit()
}
