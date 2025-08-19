// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlcodegen

import (
	"fmt"
	"github.com/vkcom/tl/internal/utils"
	"strings"
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
	if !targetType.gen.options.AddFunctionBodies {
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

func (trw *TypeRWUnion) PhpDefaultInit() string {
	return trw.Fields[0].t.trw.PhpDefaultInit()
}
