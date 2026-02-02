// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlcodegen

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/vkcom/tl/internal/tlast"
	"github.com/vkcom/tl/internal/tlcodegen/codecreator"
	"github.com/vkcom/tl/internal/utils"
)

func (trw *TypeRWStruct) PHPFindNatByName(name string) (localNat bool, indexInDeps int) {
	for i, field := range trw.Fields {
		if field.originalName == name {
			return true, i
		}
	}
	for i, argument := range trw.wr.origTL[0].TemplateArguments {
		if argument.FieldName == name {
			return false, i
		}
	}
	panic(fmt.Sprintf("no such nat \"%s\"", name))
}

func (trw *TypeRWStruct) PHPGetFieldNatDependenciesValuesAsTypeTree(fieldIndex int, calculatedArgs *TypeArgumentsTree) TypeArgumentsTree {
	field := trw.Fields[fieldIndex]
	tree := TypeArgumentsTree{}
	localTree := TypeArgumentsTree{}

	if calculatedArgs == nil {
		trw.wr.PHPGetNatTypeDependenciesDecl(&tree)
		tree.FillAllLeafs()
	} else {
		tree = *calculatedArgs
	}

	genericsMap := make(map[string]*TypeArgumentsTree)
	for _, child := range tree.children {
		if child != nil {
			genericsMap[child.name] = child
		}
	}

	field.t.PHPGetNatTypeDependenciesDecl(&localTree)
	trw.phpGetFieldArgsTree(field.t, &field.origTL.FieldType, &localTree, &genericsMap)
	return localTree
}

func (trw *TypeRWStruct) PHPGetResultNatDependenciesValuesAsTypeTree() (TypeArgumentsTree, bool) {
	if trw.ResultType == nil {
		return TypeArgumentsTree{}, false
	}
	tree := TypeArgumentsTree{}
	localTree := TypeArgumentsTree{}
	trw.wr.PHPGetNatTypeDependenciesDecl(&tree)
	tree.FillAllLeafs()

	genericsMap := make(map[string]*TypeArgumentsTree)
	for _, child := range tree.children {
		if child != nil {
			genericsMap[child.name] = child
		}
	}

	trw.ResultType.PHPGetNatTypeDependenciesDecl(&localTree)
	trw.phpGetFieldArgsTree(trw.ResultType, &trw.wr.origTL[0].FuncDecl, &localTree, &genericsMap)
	return localTree, true
}

func phpFieldMaskNullCheck(value string) string {
	return fmt.Sprintf("(%[1]s ? not_null(%[1]s) : 0)", value)
}

func (trw *TypeRWStruct) PHPGetFieldMask(targetName string, calculatedArgs *TypeArgumentsTree, fieldIndex int) string {
	fieldMask := trw.Fields[fieldIndex].fieldMask
	if fieldMask != nil {
		if fieldMask.isField {
			fieldMaskOrigin := trw.Fields[fieldMask.FieldIndex]
			fieldUsage := fmt.Sprintf("%[1]s->%[2]s", targetName, fieldMaskOrigin.originalName)
			if fieldMaskOrigin.fieldMask != nil {
				return phpFieldMaskNullCheck(fieldUsage)
			} else {
				return fieldUsage
			}
		}
		if calculatedArgs == nil {
			return "$" + fieldMask.name
		} else {
			for _, child := range calculatedArgs.children {
				if child != nil && child.name == fieldMask.name {
					return *child.value
				}
			}
		}
	}

	return ""
}

func (trw *TypeRWStruct) phpGetFieldArgsTree(currentType *TypeRWWrapper, currentTypeRef *tlast.TypeRef, tree *TypeArgumentsTree, genericsToTrees *map[string]*TypeArgumentsTree) {
	if len(currentTypeRef.Args) != len(currentType.origTL[0].TemplateArguments) {
		generic := currentTypeRef.Type.String()
		tree.CloneValuesFrom((*genericsToTrees)[generic])
		return
	}
	for i := range currentType.origTL[0].TemplateArguments {
		actualArg := currentType.arguments[i]
		actualArgRef := currentTypeRef.Args[i]
		if actualArg.isNat {
			if actualArg.isArith {
				if actualArgRef.IsArith {
					value := strconv.FormatUint(uint64(actualArg.Arith.Res), 10)
					(*tree).children[i].value = &value
				} else {
					// argument resolving to constant by in definition it is outer nat
					_, index := trw.PHPFindNatByName(actualArgRef.T.String())
					(*tree).children[i].CloneValuesFrom((*genericsToTrees)[trw.wr.origTL[0].TemplateArguments[index].FieldName])
				}
			} else {
				isLocal, index := trw.PHPFindNatByName(actualArgRef.T.String())
				if isLocal {
					value := fmt.Sprintf("$this->%s", trw.Fields[index].originalName)
					if trw.Fields[index].fieldMask != nil {
						value = phpFieldMaskNullCheck(value)
					}
					(*tree).children[i].value = &value
				} else {
					(*tree).children[i].CloneValuesFrom((*genericsToTrees)[trw.wr.origTL[0].TemplateArguments[index].FieldName])
				}
			}
		} else {
			if tree != nil {
				trw.phpGetFieldArgsTree(actualArg.tip, &actualArgRef.T, tree.children[i], genericsToTrees)
			}
		}
	}
}

func (trw *TypeRWStruct) PhpClassNameReplaced() bool {
	unionParent := trw.PhpConstructorNeedsUnion()
	if unionParent == nil {
		if trw.PhpCanBeSimplify() {
			return true
		}

		if trw.ResultType == nil && trw.wr.PHPIsTrueType() {
			return true
		}

		if phpIsDictionary(trw.wr) {
			return true
		}

		if !trw.wr.gen.options.InplaceSimpleStructs &&
			strings.HasSuffix(trw.wr.tlName.String(), "dictionary") &&
			trw.wr.tlName.Namespace == "" {
			return true
		}
	}
	return false
}

func (trw *TypeRWStruct) PhpClassName(withPath bool, bare bool) string {
	if PHPSpecialMembersTypes(trw.wr) != "" {
		return ""
	}
	unionParent := trw.PhpConstructorNeedsUnion()
	if unionParent == nil {
		if trw.PhpCanBeSimplify() {
			return trw.Fields[0].t.trw.PhpClassName(withPath, bare)
		}

		if trw.ResultType == nil && trw.wr.PHPIsTrueType() {
			return "boolean"
		}

		if phpIsDictionary(trw.wr) { // TODO NOT A SOLUTION, BUT...
			_, _, _, valueType := isDictionaryElement(trw.wr)
			return valueType.t.trw.PhpClassName(withPath, bare)
		}

		if !trw.wr.gen.options.InplaceSimpleStructs &&
			strings.HasSuffix(trw.wr.tlName.String(), "dictionary") &&
			trw.wr.tlName.Namespace == "" {
			return trw.Fields[0].t.trw.PhpClassName(withPath, bare)
		}
	}

	name := trw.wr.tlName.Name
	if !bare {
		name = trw.wr.origTL[0].TypeDecl.Name.Name
	}
	if trw.wr.tlName.Namespace != "" {
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

func (trw *TypeRWStruct) PhpTypeName(withPath bool, bare bool) string {
	if specialCase := PHPSpecialMembersTypes(trw.wr); specialCase != "" {
		return specialCase
	}
	unionParent := trw.PhpConstructorNeedsUnion()
	if unionParent == nil {
		if trw.PhpCanBeSimplify() {
			return trw.Fields[0].t.trw.PhpTypeName(withPath, trw.Fields[0].bare)
		}

		if phpIsDictionary(trw.wr) { // TODO NOT A SOLUTION, BUT...
			_, _, _, valueType := isDictionaryElement(trw.wr)
			return valueType.t.trw.PhpTypeName(withPath, bare)
		}

		if !trw.wr.gen.options.InplaceSimpleStructs &&
			strings.HasSuffix(trw.wr.tlName.String(), "dictionary") &&
			trw.wr.tlName.Namespace == "" {
			return trw.Fields[0].t.trw.PhpTypeName(withPath, bare)
		}
	}
	return trw.PhpClassName(withPath, bare)
}

func (trw *TypeRWStruct) PhpGenerateCode(code *strings.Builder, bytes bool) error {
	trw.PHPStructHeader(code)
	trw.PHPStructFieldMasks(code)
	trw.PHPStructFields(code)
	trw.PHPStructResultType(code)

	necessaryFieldsInConstructor := make([]Field, 0)
	usedFieldMasksIndecies := make([]int, 0)
	usedFieldMasks := make(map[int][]Field)
	for _, f := range trw.Fields {
		if f.fieldMask == nil {
			necessaryFieldsInConstructor = append(necessaryFieldsInConstructor, f)
		} else {
			index := f.fieldMask.FieldIndex
			if !f.fieldMask.isField {
				for i, argument := range trw.wr.origTL[0].TemplateArguments {
					if argument.IsNat && argument.FieldName == f.fieldMask.name {
						index = -(i + 1)
						break
					}
				}
			}
			if usedFieldMasks[index] == nil {
				usedFieldMasksIndecies = append(usedFieldMasksIndecies, index)
			}
			usedFieldMasks[index] = append(usedFieldMasks[index], f)
		}
	}

	trw.PHPStructConstructor(code, necessaryFieldsInConstructor)
	trw.PHPStructRPCSpecialGetters(code)
	trw.PHPStructReadMethods(code)
	trw.PHPStructWriteMethods(code)
	trw.PHPStructFieldMaskCalculators(code, usedFieldMasksIndecies, usedFieldMasks)
	trw.PHPStructFunctionSpecificMethods(code)

	code.WriteString("\n}\n")

	trw.PHPStructFunctionSpecificTypes(code)
	return nil
}

func (trw *TypeRWStruct) PHPStructFunctionSpecificTypes(code *strings.Builder) {
	if trw.wr.gen.options.AddRPCTypes && trw.ResultType != nil {
		code.WriteString(
			fmt.Sprintf(
				`
/**
 * @kphp-tl-class
 */
class %[1]s_result implements TL\RpcFunctionReturnResult {

  /** @var %[2]s */
  public $value = %[3]s;

}
`,
				trw.PhpClassName(false, true),
				phpResultType(trw),
				trw.ResultType.trw.PhpDefaultValue(),
			),
		)

		if trw.wr.gen.options.AddFetchers && trw.wr.phpInfo.RequireFunctionBodies {
			args, _ := trw.PHPGetResultNatDependenciesValuesAsTypeTree()
			argsAsArray := args.EnumerateWithPrefixes()

			argAsArrayCopy := argsAsArray
			if trw.wr.wantsTL2 {
				argAsArrayCopy = append(argAsArrayCopy, "$use_tl2")
			}

			argsAsFields := strings.Join(
				utils.MapSlice(
					argAsArrayCopy,
					func(arg string) string {
						return fmt.Sprintf(
							`  /** @var int */
  public %[1]s = 0;
`,
							arg,
						)
					},
				),
				"\n",
			)

			if argsAsFields != "" {
				argsAsFields += "\n"
			}

			constructorComment := `  /**
   * @kphp-inline
   */`
			constructorArgs := ""
			constructorBody := ""

			if len(argsAsArray) > 0 {
				constructorComment = "  /**\n"
				for i, arg := range argsAsArray {
					suffix, _ := strings.CutPrefix(arg, "$")

					constructorComment += fmt.Sprintf("   * @param int $%s\n", suffix)

					if i != 0 {
						constructorArgs += ", "
					}
					constructorArgs += arg

					constructorBody += fmt.Sprintf("    $this->%[1]s = $%[1]s;\n", suffix)
				}
				constructorComment += "   */"
			}

			args.FillAllLeafsWithValues(utils.MapSlice(argsAsArray, func(a string) string {
				suffix, _ := strings.CutPrefix(a, "$")
				return fmt.Sprintf("$this->%s", suffix)
			}))

			readCall := strings.Builder{}
			writeCall := strings.Builder{}

			/** TODO make it better */
			if trw.wr.origTL[0].OriginalDescriptor != nil &&
				trw.wr.origTL[0].OriginalDescriptor.OriginalDescriptor != nil &&
				len(trw.wr.origTL[0].OriginalDescriptor.OriginalDescriptor.TemplateArguments) != 0 {
				readCall.WriteString(`    /** TODO FOR DIAGONAL */`)
				writeCall.WriteString(`    /** TODO FOR DIAGONAL */`)
			} else {
				readCallLines := trw.ResultType.trw.PhpReadMethodCall("$result->value", false, true, &args, "")
				if trw.wr.wantsTL2 {
					cc := codecreator.NewPhpCodeCreator()
					cc.IfElse("$this->use_tl2 == 0", func(cc *codecreator.PhpCodeCreator) {
						// tl1 case
						cc.AddLines(readCallLines...)
					}, func(cc *codecreator.PhpCodeCreator) {
						// tl2 case
						cc.AddLines("$used_bytes = 0;")
						cc.AddLines("$obj_size = TL\\tl2_support::fetch_size();")
						cc.If("$obj_size != 0", func(cc *codecreator.PhpCodeCreator) {
							cc.AddLines("$obj_block = fetch_byte();")
							cc.AddLines("$used_bytes += 1;")
							cc.If("$obj_block == (1 << 1)", func(cc *codecreator.PhpCodeCreator) {
								cc.AddLines(trw.ResultType.trw.PhpReadTL2MethodCall("$result->value", false, true, &args, "", 0, "$used_bytes", false)...)
							})
						})
						// skip rest
						cc.AddLines(fmt.Sprintf("TL\\tl2_support::skip_bytes(%[1]s - %[2]s);", "$obj_size", "$used_bytes"))
					})
					readCallLines = cc.Print()
				}
				for _, line := range readCallLines {
					targetLines := []string{line}
					if strings.Contains(line, "return false;") {
						prefix, _, _ := strings.Cut(line, "return false;")
						targetLines[0] = prefix + fmt.Sprintf("throw new \\Exception('can\\'t fetch %s_result');", trw.PhpClassName(false, true))
						targetLines = append(targetLines, prefix+"return null;")
					}
					for _, targetLine := range targetLines {
						readCall.WriteString(strings.Repeat(" ", 4))
						readCall.WriteString(targetLine)
						readCall.WriteString("\n")
					}
				}

				writeCallLines := trw.ResultType.trw.PhpWriteMethodCall("$result->value", false, &args, "")
				if trw.wr.wantsTL2 {
					cc := codecreator.NewPhpCodeCreator()
					cc.IfElse("$this->use_tl2 == 0", func(cc *codecreator.PhpCodeCreator) {
						cc.AddLines(writeCallLines...)
					}, func(cc *codecreator.PhpCodeCreator) {
						cc.AddLines(
							"$used_bytes = 0;",
							`$context_sizes = new TL\tl2_context();`,
							`$context_blocks = new TL\tl2_context();`,
						)
						cc.Comments("calculate sizes")
						cc.AddLines(trw.ResultType.trw.PhpCalculateSizesTL2MethodCall("$result->value", false, &args, "", 0, "$used_bytes", true)...)
						cc.Comments("write result")
						cc.IfElse("$used_bytes != 0", func(cc *codecreator.BasicCodeCreator[codecreator.PhpHelder]) {
							cc.AddLines("TL\\tl2_support::store_size(1 + $used_bytes);")
							cc.AddLines("store_byte(2);")
							cc.AddLines(trw.ResultType.trw.PhpWriteTL2MethodCall("$result->value", false, &args, "", 0, "$used_bytes", false)...)
						}, func(cc *codecreator.BasicCodeCreator[codecreator.PhpHelder]) {
							cc.AddLines("TL\\tl2_support::store_size(0);")
						})
					})
					writeCallLines = cc.Print()
				}
				for _, line := range writeCallLines {
					targetLines := []string{line}
					if strings.Contains(line, "return false;") {
						prefix, _, _ := strings.Cut(line, "return false;")
						targetLines[0] = prefix + fmt.Sprintf("throw new \\Exception('can\\'t store %s_result');", trw.PhpClassName(false, true))
						targetLines = append(targetLines, prefix+"return;")
					}
					for _, targetLine := range targetLines {
						writeCall.WriteString(strings.Repeat(" ", 6))
						writeCall.WriteString(targetLine)
						writeCall.WriteString("\n")
					}
				}
			}

			var fetchArgNames []string
			var fetchArgTypes []string

			var storeArgNames []string
			var storeArgTypes []string

			if !trw.wr.gen.options.UseBuiltinDataProviders {
				fetchArgNames = append(fetchArgNames, "stream")
				fetchArgTypes = append(fetchArgTypes, `TL\tl_input_stream`)

				storeArgNames = append(storeArgNames, "stream")
				storeArgTypes = append(storeArgTypes, `TL\tl_output_stream`)
			}

			storeArgNames = append(storeArgNames, "result")
			storeArgTypes = append(storeArgTypes, `TL\RpcFunctionReturnResult`)

			code.WriteString(
				fmt.Sprintf(
					`
class %[1]s_fetcher implements TL\RpcFunctionFetcher {
%[4]s%[6]s
  public function __construct(%[7]s) {
%[8]s  }

%[9]s
  public function typedFetch(%[11]s) {
    $result = new %[1]s_result();
%[3]s
    return $result;
  }

%[10]s
  public function typedStore(%[12]s) {
    if ($result instanceof %[1]s_result) {
%[5]s
    } else {
      throw new \Exception("can\'t store: %[1]s_result expected");
    }
  }
}
`,
					trw.PhpClassName(false, true),
					trw.ResultType.trw.PhpTypeName(true, true),
					readCall.String(),
					argsAsFields,
					writeCall.String(),
					constructorComment,
					constructorArgs,
					constructorBody,
					phpFunctionCommentFormat(
						fetchArgNames,
						fetchArgTypes,
						`TL\RpcFunctionReturnResult`,
						"  ",
					),
					phpFunctionCommentFormat(
						storeArgNames,
						storeArgTypes,
						``,
						"  ",
					),
					phpFunctionArgumentsFormat(fetchArgNames),
					phpFunctionArgumentsFormat(storeArgNames),
				),
			)
		}
	}
}

func (trw *TypeRWStruct) PHPStructFunctionSpecificMethods(code *strings.Builder) {
	// print function specific methods and types
	if trw.wr.gen.options.AddRPCTypes && trw.ResultType != nil {
		kphpSpecialCode := ""
		if trw.wr.HasAnnotation("kphp") {
			kphpSpecialCode = fmt.Sprintf(
				`

  /**
   * @param %[1]s $value
   * @return %[2]s_result
   */
  public static function createRpcServerResponse($value) {
    $response = new %[2]s_result();
    $response->value = $value;
    return $response;
  }`,
				trw.ResultType.trw.PhpTypeName(true, true),
				trw.PhpClassName(true, true),
			)
		}

		code.WriteString(
			fmt.Sprintf(
				`
  /**
   * @param TL\RpcFunctionReturnResult $function_return_result
   * @return %[4]s
   */
  public static function functionReturnValue($function_return_result) {
    if ($function_return_result instanceof %[1]s_result) {
      return $function_return_result->value;
    }
    warning('Unexpected result type in functionReturnValue: ' . ($function_return_result ? get_class($function_return_result) : 'null'));
    return (new %[1]s_result())->value;
  }

  /**
   * @kphp-inline
   *
   * @param TL\RpcResponse $response
   * @return %[4]s
   */
  public static function result(TL\RpcResponse $response) {
    return self::functionReturnValue($response->getResult());
  }%[5]s
`,
				trw.PhpClassName(false, true),
				trw.PhpClassName(true, true),
				trw.wr.tlName.String(),
				phpResultType(trw),
				kphpSpecialCode,
			),
		)

		code.WriteString(fmt.Sprintf(`
  /**
   * @kphp-inline
   *
   * @return int
   */
  public function getTLFunctionMagic() {
    return 0x%08[2]x;
  }

  /**
   * @kphp-inline
   *
   * @return string
   */
  public function getTLFunctionName() {
    return '%[1]s';
  }
`,
			trw.wr.tlName.String(),
			trw.wr.tlTag,
		))

		args, _ := trw.PHPGetResultNatDependenciesValuesAsTypeTree()
		argsArray := strings.Join(args.ListAllValues(), ", ")

		var fetchArgNames []string
		var fetchArgTypes []string

		var storeArgNames []string
		var storeArgTypes []string

		if !trw.wr.gen.options.UseBuiltinDataProviders {
			fetchArgNames = append(fetchArgNames, "stream")
			fetchArgTypes = append(fetchArgTypes, `TL\tl_input_stream`)

			storeArgNames = append(storeArgNames, "stream")
			storeArgTypes = append(storeArgTypes, `TL\tl_output_stream`)
		}

		if trw.wr.gen.options.AddFetchers &&
			// diagonal
			len(trw.wr.origTL[0].MostOriginalVersion().TemplateArguments) == 0 &&
			// don't have write / read
			trw.wr.phpInfo.RequireFunctionBodies {
			if !trw.wr.wantsTL2 {
				code.WriteString(
					fmt.Sprintf(`
%[6]s
  public function typedStore(%[8]s) {
%[10]s    %[9]sprint('%[1]s::typedStore()<br/>');
    $this->write_boxed(%[8]s);
    return new %[1]s_fetcher(%[4]s);
  }

%[5]s
  public function typedFetch(%[7]s) {
%[10]s    %[9]sprint('%[1]s::typedFetch()<br/>');
    $this->read(%[7]s);
    return new %[1]s_fetcher(%[4]s);
  }
`,
						trw.PhpClassName(false, true),
						trw.wr.tlName.String(),
						fmt.Sprintf("0x%08x", trw.wr.tlTag),
						argsArray,
						phpFunctionCommentFormat(
							fetchArgNames,
							fetchArgTypes,
							`TL\RpcFunctionFetcher`,
							"  ",
						),
						phpFunctionCommentFormat(
							storeArgNames,
							storeArgTypes,
							`TL\RpcFunctionFetcher`,
							"  ",
						),
						phpFunctionArgumentsFormat(fetchArgNames),
						phpFunctionArgumentsFormat(storeArgNames),
						ifString(trw.wr.gen.options.AddFetchersEchoComments, "", "//"),
						ifString(trw.wr.gen.options.AddSwitcher,
							fmt.Sprintf(`    if (TL\tl_switcher::tl_get_namespace_methods_mode("%[1]s") == 0) {
      return null;
    }
`,
								trw.wr.tlName.Namespace,
							),
							"",
						),
					),
				)
			} else {
				if !trw.wr.gen.options.AddSwitcher {
					// TODO
					panic("can't create tl2 call without switcher")
				}
				if !trw.wr.gen.options.UseBuiltinDataProviders {
					// TODO
					panic("can't create tl2 without builtin providers")
				}

				code.WriteString(
					fmt.Sprintf(`
%[6]s
  public function typedStore(%[8]s) {
    if (TL\tl_switcher::tl_get_namespace_methods_mode("%[10]s") == 1) {
      %[9]sprint('%[1]s::typedStore()<br/>');
      $this->write_boxed(%[8]s);
      return new %[1]s_fetcher(%[4]s);
    } else if (TL\tl_switcher::tl_get_namespace_methods_mode("%[10]s") == 2) {
      %[9]sprint('%[1]s::typedStore() in tl2<br/>');
      store_int(0x%08[11]x); 
      $this->write_tl2(%[8]s);
      $f = new %[1]s_fetcher(%[4]s);
      $f->use_tl2 = 1;
      return $f; 
    } else {
      return null;
    }
  }

%[5]s
  public function typedFetch(%[7]s) {
    if (TL\tl_switcher::tl_get_namespace_methods_mode("%[10]s") == 1) {
      %[9]sprint('%[1]s::typedFetch()<br/>');
      $this->read(%[7]s);
      return new %[1]s_fetcher(%[4]s);
    } else if (TL\tl_switcher::tl_get_namespace_methods_mode("%[10]s") == 2) {
      %[9]sprint('%[1]s::typedFetch() in tl2<br/>');
      $this->read_tl2(%[7]s);
      $f = new %[1]s_fetcher(%[4]s);
      $f->use_tl2 = 1;
      return $f;
    } else {
      return null;
    }
  }
`,
						trw.PhpClassName(false, true),
						trw.wr.tlName.String(),
						fmt.Sprintf("0x%08x", trw.wr.tlTag),
						argsArray,
						phpFunctionCommentFormat(
							fetchArgNames,
							fetchArgTypes,
							`TL\RpcFunctionFetcher`,
							"  ",
						),
						phpFunctionCommentFormat(
							storeArgNames,
							storeArgTypes,
							`TL\RpcFunctionFetcher`,
							"  ",
						),
						phpFunctionArgumentsFormat(fetchArgNames),
						phpFunctionArgumentsFormat(storeArgNames),
						ifString(trw.wr.gen.options.AddFetchersEchoComments, "", "//"),
						trw.wr.tlName.Namespace,
						trw.wr.tlTag,
					),
				)
			}
			// only for rpcDest*
		} else if trw.wr.gen.options.AddFetchers &&
			// don't have write / read
			trw.wr.phpInfo.RequireFunctionBodies &&
			// diagonal
			len(trw.wr.origTL[0].MostOriginalVersion().TemplateArguments) != 0 &&
			// from _common
			trw.wr.origTL[0].MostOriginalVersion().Construct.Name.Namespace == "" {
			code.WriteString(
				fmt.Sprintf(`
%[6]s
  public function typedStore(%[8]s) {
%[10]s    %[9]sprint('%[1]s::typedStore()<br/>');
    $this->write_boxed(%[8]s);
    $fetcher = $this->query->typedStore(%[8]s);
    if ($fetcher === null) {
      %[9]sprint('%[1]s rpc_clean()<br/>');
      rpc_clean();
    }
    return $fetcher;
  }

%[5]s
  public function typedFetch(%[7]s) {
    return null;
  }
`,
					trw.PhpClassName(false, true),
					trw.wr.tlName.String(),
					fmt.Sprintf("0x%08x", trw.wr.tlTag),
					argsArray,
					phpFunctionCommentFormat(
						fetchArgNames,
						fetchArgTypes,
						`TL\RpcFunctionFetcher`,
						"  ",
					),
					phpFunctionCommentFormat(
						storeArgNames,
						storeArgTypes,
						`TL\RpcFunctionFetcher`,
						"  ",
					),
					phpFunctionArgumentsFormat(fetchArgNames),
					phpFunctionArgumentsFormat(storeArgNames),
					ifString(trw.wr.gen.options.AddFetchersEchoComments, "", "//"),
					ifString(trw.wr.gen.options.AddSwitcher,
						fmt.Sprintf(`    if (TL\tl_switcher::tl_get_namespace_methods_mode("%[1]s") == 0) {
      return null;
    }
`,
							"_common",
						),
						"",
					),
				),
			)
		} else {
			code.WriteString(
				fmt.Sprintf(`
%[6]s
  public function typedStore(%[8]s) {
    return null;
  }

%[5]s
  public function typedFetch(%[7]s) {
    return null;
  }
`,
					trw.PhpClassName(false, true),
					trw.wr.tlName.String(),
					fmt.Sprintf("0x%08x", trw.wr.tlTag),
					argsArray,
					phpFunctionCommentFormat(
						fetchArgNames,
						fetchArgTypes,
						`TL\RpcFunctionFetcher`,
						"  ",
					),
					phpFunctionCommentFormat(
						storeArgNames,
						storeArgTypes,
						`TL\RpcFunctionFetcher`,
						"  ",
					),
					phpFunctionArgumentsFormat(fetchArgNames),
					phpFunctionArgumentsFormat(storeArgNames),
				),
			)
		}
	}
}

func (trw *TypeRWStruct) PHPStructReadMethods(code *strings.Builder) {
	useBuiltin := trw.wr.gen.options.UseBuiltinDataProviders
	if trw.wr.gen.options.AddFunctionBodies &&
		trw.wr.phpInfo.RequireFunctionBodies {
		natParams := trw.wr.PHPGetNatTypeDependenciesDeclAsArray()
		natParams = utils.MapSlice(natParams, func(a string) string {
			s, _ := strings.CutPrefix(a, "$")
			return s
		})
		argNames := make([]string, 0)
		argTypes := make([]string, 0)
		if !useBuiltin {
			argNames = append(argNames, "stream")
			argTypes = append(argTypes, `TL\tl_input_stream`)
		}
		for _, param := range natParams {
			argNames = append(argNames, param)
			argTypes = append(argTypes, "int")
		}

		magicRead := []string{
			"    [$magic, $success] = $stream->read_uint32();",
			fmt.Sprintf("    if (!$success || $magic != 0x%08[1]x) {", trw.wr.tlTag),
			"      return false;",
			"    }",
		}

		if useBuiltin {
			magicRead = []string{
				"    $magic = fetch_int() & 0xFFFFFFFF;",
				fmt.Sprintf("    if ($magic != 0x%08[1]x) {", trw.wr.tlTag),
				"      return false;",
				"    }",
			}
		}

		code.WriteString(fmt.Sprintf(`
%[1]s
  public function read_boxed(%[2]s) {
%[3]s
    return $this->read(%[2]s);
  }
`,
			phpFunctionCommentFormat(argNames, argTypes, "bool", "  "),
			phpFunctionArgumentsFormat(argNames),
			strings.Join(magicRead, "\n"),
		))

		code.WriteString(fmt.Sprintf(`
%[1]s
  public function read(%[2]s) {
`,
			phpFunctionCommentFormat(argNames, argTypes, "bool", "  "),
			phpFunctionArgumentsFormat(argNames),
		))

		for _, line := range trw.phpStructReadCode("$this", nil) {
			code.WriteString(fmt.Sprintf("%[1]s%[2]s\n", strings.Repeat("  ", 2), line))
		}

		code.WriteString("    return true;\n")
		code.WriteString("  }\n")

		if trw.wr.wantsTL2 {
			// TODO: add block calculated currentSize and block if union
			if trw.wr.PHPUnionParent() != nil {
				argNames = append(argNames, "block", "current_size")
				argTypes = append(argTypes, "int", "int")
			}
			code.WriteString(fmt.Sprintf(`
%[1]s
  public function read_tl2(%[2]s) {
`,
				phpFunctionCommentFormat(argNames, argTypes, "int", "  "),
				phpFunctionArgumentsFormat(argNames),
			))

			tab := strings.Repeat("  ", 2)
			usedBytes := "$used_bytes"
			code.WriteString(fmt.Sprintf("%[1]s%[2]s = 0;\n", tab, usedBytes))

			for _, line := range trw.phpStructReadTL2Code("$this", usedBytes, nil, true) {
				code.WriteString(fmt.Sprintf("%[1]s%[2]s\n", tab, line))
			}

			code.WriteString("    return $used_bytes;\n")
			code.WriteString("  }\n")
		}
	}
}

func (trw *TypeRWStruct) phpStructReadCode(targetName string, calculatedArgs *TypeArgumentsTree) []string {
	result := make([]string, 0)
	const tab = "  "
	for i, field := range trw.Fields {
		fieldMask := trw.PHPGetFieldMask(targetName, calculatedArgs, i)
		shift := 0
		textTab := func() string { return strings.Repeat(tab, shift) }
		if fieldMask != "" {
			result = append(result,
				fmt.Sprintf(
					"%[1]sif ((%[2]s & (1 << %[3]d)) != 0) {",
					textTab(),
					fieldMask,
					field.BitNumber,
				),
			)
			shift += 1
		}
		tree := trw.PHPGetFieldNatDependenciesValuesAsTypeTree(i, calculatedArgs)
		fieldRead := field.t.trw.PhpReadMethodCall(targetName+"->"+field.originalName, field.bare, true, &tree, strconv.Itoa(i))
		for _, line := range fieldRead {
			result = append(result, textTab()+line)
		}
		if fieldMask != "" {
			shift -= 1
			result = append(result, fmt.Sprintf("%[1]s} else {", textTab()))
			shift += 1
			_, defaultValue := fieldTypeAndDefaultValue(field)
			result = append(result, fmt.Sprintf(
				"%[1]s%[2]s = %[3]s;",
				textTab(),
				targetName+"->"+field.originalName,
				defaultValue,
			))
			shift -= 1
			result = append(result, fmt.Sprintf("%[1]s}", textTab()))
		}
	}

	return result
}

func (trw *TypeRWStruct) phpStructReadTL2Code(targetName string, usedBytesPointer string, calculatedArgs *TypeArgumentsTree, topLevel bool) []string {
	currentSize := "$current_size"
	block := "$block"

	cc := codecreator.NewPhpCodeCreator()

	subtractSize := func(value string) string {
		return fmt.Sprintf("%[1]s -= %[2]s;", currentSize, value)
	}

	if trw.wr.PHPUnionParent() == nil {
		// fetch currentSize and update usedBytesPointer
		cc.Comments("read size of object and update used bytes")
		cc.AddLines(
			fmt.Sprintf("%[1]s = TL\\tl2_support::fetch_size();", currentSize),
			fmt.Sprintf("%[1]s += TL\\tl2_support::count_used_bytes(%[2]s) + %[2]s;", usedBytesPointer, currentSize),
		)

		// if 0 return
		cc.If(fmt.Sprintf("%[1]s == 0", currentSize), func(cc *codecreator.PhpCodeCreator) {
			cc.AddLines(fmt.Sprintf("return %[1]s;", usedBytesPointer))
		})

		// fetch 1st block and start subtract fron size
		cc.Comments("read first block and check constructor id")
		cc.AddLines(
			fmt.Sprintf("%[1]s = fetch_byte();", block),
			subtractSize("1"),
		)

		cc.If(fmt.Sprintf("(%[1]s & 1) != 0", block), func(cc *codecreator.PhpCodeCreator) {
			cc.AddLines(
				"$index = TL\\tl2_support::fetch_size();",
				subtractSize("TL\\tl2_support::count_used_bytes($index)"),
			)
			cc.If("$index != 0", func(cc *codecreator.PhpCodeCreator) {
				cc.AddLines(
					fmt.Sprintf("TL\\tl2_support::skip_bytes(%[1]s);", currentSize),
					fmt.Sprintf("return %[1]s;", usedBytesPointer),
				)
			})
		})
	}

	for fieldIndex, field := range trw.Fields {
		isTrue := field.t.IsTrueType()
		fieldName := fmt.Sprintf("%[1]s->%[2]s", targetName, field.originalName)
		inBlockIndex := (fieldIndex + 1) % 8

		// add new block
		if inBlockIndex == 0 {
			cc.Comments(
				fmt.Sprintf("read new block with index %d", (fieldIndex+1)/8),
			)
			cc.IfElse(fmt.Sprintf("%[1]s > 0", currentSize),
				func(cc *codecreator.PhpCodeCreator) {
					cc.AddLines(
						fmt.Sprintf("%[1]s = fetch_byte();", block),
						subtractSize("1"),
					)
				},
				func(cc *codecreator.PhpCodeCreator) {
					cc.AddLines(fmt.Sprintf("%[1]s = 0;", block))
				},
			)
		}

		// read field
		cc.Comments(fmt.Sprintf("read field with index %d with name \"%s\"", fieldIndex, fieldName))
		if isTrue {
			cc.AddLines(fmt.Sprintf("%[1]s = false;", fieldName))
		}
		cc.If(fmt.Sprintf("($block & (1 << %d)) != 0", inBlockIndex), func(cc *codecreator.PhpCodeCreator) {
			tree := trw.PHPGetFieldNatDependenciesValuesAsTypeTree(fieldIndex, calculatedArgs)
			localUsedBytes := "$local_used_bytes"
			cc.AddLines(fmt.Sprintf("%[1]s = 0;", localUsedBytes))
			cc.AddLines(
				field.t.trw.PhpReadTL2MethodCall(fieldName, field.bare, true, &tree, strconv.Itoa(fieldIndex), 0, localUsedBytes, field.fieldMask == nil)...,
			)
			cc.AddLines(subtractSize(localUsedBytes))
			cc.If(fmt.Sprintf("%[1]s < 0", currentSize), func(cc *codecreator.PhpCodeCreator) {
				cc.AddLines(`throw new \Exception("read more bytes than passed in struct definition");`)
			})
		})
	}

	// skip tail
	cc.Comments("skip remaining bytes")
	cc.AddLines(fmt.Sprintf("TL\\tl2_support::skip_bytes(%[1]s);", currentSize))

	return cc.Print()
}

func (trw *TypeRWStruct) PHPStructWriteMethods(code *strings.Builder) {
	useBuiltin := trw.wr.gen.options.UseBuiltinDataProviders
	if trw.wr.gen.options.AddFunctionBodies &&
		trw.wr.phpInfo.RequireFunctionBodies {
		natParams := trw.wr.PHPGetNatTypeDependenciesDeclAsArray()
		natParams = utils.MapSlice(natParams, func(a string) string {
			s, _ := strings.CutPrefix(a, "$")
			return s
		})
		argNames := make([]string, 0)
		argTypes := make([]string, 0)
		if !useBuiltin {
			argNames = append(argNames, "stream")
			argTypes = append(argTypes, `TL\tl_output_stream`)
		}
		for _, param := range natParams {
			argNames = append(argNames, param)
			argTypes = append(argTypes, "int")
		}

		magicWrite := []string{
			fmt.Sprintf("    $success = $stream->write_uint32(0x%08[1]x)", trw.wr.tlTag),
			"    if (!$success) {",
			"      return false;",
			"    }",
		}

		if useBuiltin {
			magicWrite = []string{
				fmt.Sprintf("    store_int(0x%08[1]x);", trw.wr.tlTag),
			}
		}

		code.WriteString(fmt.Sprintf(`
%[1]s
  public function write_boxed(%[2]s) {
%[3]s
    return $this->write(%[2]s);
  }
`,
			phpFunctionCommentFormat(argNames, argTypes, "bool", "  "),
			phpFunctionArgumentsFormat(argNames),
			strings.Join(magicWrite, "\n"),
		))

		code.WriteString(fmt.Sprintf(`
%[1]s
  public function write(%[2]s) {
`,
			phpFunctionCommentFormat(argNames, argTypes, "bool", "  "),
			phpFunctionArgumentsFormat(argNames),
		))

		for _, line := range trw.phpStructWriteCode("$this", nil) {
			code.WriteString(fmt.Sprintf("%[1]s%[2]s\n", strings.Repeat("  ", 2), line))
		}

		code.WriteString("    return true;\n")
		code.WriteString("  }\n")

		if trw.wr.wantsTL2 {
			// TODO: add block calculated currentSize and block if union
			argNames = append(argNames, "context_sizes", "context_blocks")
			argTypes = append(argTypes, `TL\tl2_context`, `TL\tl2_context`)

			code.WriteString(fmt.Sprintf(`
%[1]s
  public function write_tl2(%[2]s) {
    $context_sizes = new TL\tl2_context();
    $context_blocks = new TL\tl2_context();
    $this->calculate_sizes_tl2(%[3]s);
    $this->internal_write_tl2(%[3]s);
  }
`,
				phpFunctionCommentFormat(argNames[:len(argNames)-2], argTypes[:len(argNames)-2], "", "  "),
				phpFunctionArgumentsFormat(argNames[:len(argNames)-2]),
				phpFunctionArgumentsFormat(argNames),
			))

			code.WriteString(fmt.Sprintf(`
%[1]s
  public function internal_write_tl2(%[2]s) {
`,
				phpFunctionCommentFormat(argNames, argTypes, "int", "  "),
				phpFunctionArgumentsFormat(argNames),
			))

			tab := strings.Repeat("  ", 2)

			code.WriteString(fmt.Sprintf("%s$used_bytes = 0;\n", tab))
			for _, line := range trw.phpStructWriteTL2Code("$this", nil, "", 0, "$used_bytes", false) {
				code.WriteString(fmt.Sprintf("%[1]s%[2]s\n", tab, line))
			}
			code.WriteString("    return $used_bytes;\n")
			code.WriteString("  }\n")

			code.WriteString(fmt.Sprintf(`
%[1]s
  public function calculate_sizes_tl2(%[2]s) {
`,
				phpFunctionCommentFormat(argNames, argTypes, "int", "  "),
				phpFunctionArgumentsFormat(argNames),
			))

			code.WriteString(fmt.Sprintf("%s$used_bytes = 0;\n", tab))
			for _, line := range trw.phpStructCalculateSizesTL2Code("$this", nil, "", 0, "$used_bytes", false, true) {
				code.WriteString(fmt.Sprintf("%[1]s%[2]s\n", tab, line))
			}
			code.WriteString("    return $used_bytes;\n")
			code.WriteString("  }\n")
		}
	}
}

func (trw *TypeRWStruct) phpStructWriteCode(targetName string, calculatedArgs *TypeArgumentsTree) []string {
	result := make([]string, 0)
	const tab = "  "
	for i, field := range trw.Fields {
		fieldMask := trw.PHPGetFieldMask(targetName, calculatedArgs, i)
		shift := 0
		textTab := func() string { return strings.Repeat(tab, shift) }
		tree := trw.PHPGetFieldNatDependenciesValuesAsTypeTree(i, calculatedArgs)
		fieldRead := field.t.trw.PhpWriteMethodCall(targetName+"->"+field.originalName, field.bare, &tree, strconv.Itoa(i))
		if fieldRead == nil {
			continue
		}
		if fieldMask != "" {
			result = append(result,
				fmt.Sprintf(
					"%[1]sif ((%[2]s & (1 << %[3]d)) != 0) {",
					textTab(),
					fieldMask,
					field.BitNumber,
				),
			)
			shift += 1
		}
		for _, line := range fieldRead {
			result = append(result, textTab()+line)
		}
		if fieldMask != "" {
			shift -= 1
			result = append(result, fmt.Sprintf("%[1]s}", textTab()))
		}
	}
	return result
}

func (trw *TypeRWStruct) phpStructWriteTL2Code(targetName string, args *TypeArgumentsTree, supportSuffix string, callLevel int, usedBytesPointer string, useItself bool) []string {
	if useItself && len(trw.Fields) != 1 {
		panic("can't use itself (for case similar to typedef)")
	}
	uniqueSuffix := fmt.Sprintf("_%s_%d", supportSuffix, callLevel)

	currentSize := fmt.Sprintf("$current_size%s", uniqueSuffix)
	writeSize := fmt.Sprintf("$write_size%s", uniqueSuffix)

	cc := codecreator.CodeCreator{Shift: "  "}

	// add size
	cc.AddLines(
		fmt.Sprintf("%s = $context_sizes->pop_front();", currentSize),
		fmt.Sprintf("TL\\tl2_support::store_size(%s);", currentSize),
		fmt.Sprintf("%[1]s += %[2]s + TL\\tl2_support::count_used_bytes(%[2]s);", usedBytesPointer, currentSize),
		fmt.Sprintf("%s = 0;", writeSize),
	)

	cc.AddLines(fmt.Sprintf("if (%s != 0) {", currentSize))
	cc.AddBlock(func(cc *codecreator.CodeCreator) {
		currentBlock := fmt.Sprintf("$block%s", uniqueSuffix)

		cc.AddLines(
			fmt.Sprintf("%s = $context_blocks->pop_front();", currentBlock),
			fmt.Sprintf("store_byte(%s & 0xFF);", currentBlock),
			fmt.Sprintf("%s += 1;", writeSize),
		)

		// add constructor id
		cc.AddLines(fmt.Sprintf("if ((%s & (1 << 0)) != 0) {", currentBlock))
		cc.AddBlock(func(cc *codecreator.CodeCreator) {
			index := 0
			if trw.wr.PHPUnionParent() != nil {
				index = trw.wr.unionIndex
			}
			cc.AddLines(
				fmt.Sprintf("TL\\tl2_support::store_size(%d);", index),
				fmt.Sprintf("%[1]s += TL\\tl2_support::count_used_bytes(%[2]d);", writeSize, index),
			)
		})
		cc.AddLines("}")

		for i, field := range trw.Fields {
			indexInBlock := (i + 1) % 8
			if indexInBlock == 0 {
				cc.AddLines("// write next block")
				cc.AddLines(fmt.Sprintf("%[1]s = 0;", currentBlock))
				cc.AddLines(fmt.Sprintf("if (%[1]s > %[2]s) {", currentSize, writeSize))
				cc.AddBlock(func(cc *codecreator.CodeCreator) {
					// read next block
					cc.AddLines(
						fmt.Sprintf("%s = $context_blocks->pop_front();", currentBlock),
						fmt.Sprintf("store_byte(%s & 0xFF);", currentBlock),
						fmt.Sprintf("%s += 1;", writeSize),
					)
				})
				cc.AddLines("}")
			}
			// skip if field is omitted
			if field.IsTL2Omitted() {
				continue
			}
			cc.AddLines(fmt.Sprintf("// write field %s", field.originalName))
			// skip bit since it is written in block
			if field.IsBit() {
				continue
			}

			fieldTarget := fmt.Sprintf("%s->%s", targetName, field.originalName)
			if useItself {
				fieldTarget = targetName
			}
			tree := trw.PHPGetFieldNatDependenciesValuesAsTypeTree(i, args)

			cc.AddLines(fmt.Sprintf("if ((%[1]s & (1 << %[2]d)) != 0) {", currentBlock, indexInBlock))
			cc.AddBlock(func(cc *codecreator.CodeCreator) {
				if field.MaskTL2Bit != nil {
					cc.AddLines(fmt.Sprintf("if (%[1]s !== null) {", fieldTarget))
					cc.AddShift(1)
				}
				cc.AddLines(field.t.trw.PhpWriteTL2MethodCall(fieldTarget, true, &tree, fmt.Sprintf("%s_%d", supportSuffix, i), callLevel+1, writeSize, true)...)
				if field.MaskTL2Bit != nil {
					cc.AddShift(-1)
					cc.AddLines("}")
				}
			})
			cc.AddLines("}")
		}
	})
	cc.AddLines("}")

	return cc.Print()
}

func (trw *TypeRWStruct) phpStructCalculateSizesTL2Code(targetName string, args *TypeArgumentsTree, supportSuffix string, callLevel int, usedBytesPointer string, useItself bool, canOmit bool) []string {
	if useItself && len(trw.Fields) != 1 {
		panic("can't use itself (for case similar to typedef)")
	}
	uniqueSuffix := fmt.Sprintf("_%s_%d", supportSuffix, callLevel)

	currentSize := fmt.Sprintf("$current_size%s", uniqueSuffix)
	currentSizeIndex := fmt.Sprintf("$current_size_index%s", uniqueSuffix)

	blocksUsed := fmt.Sprintf("$blocks_used%s", uniqueSuffix)
	nextBlockIndex := fmt.Sprintf("$next_block_index%s", uniqueSuffix)

	currentBlock := fmt.Sprintf("$current_block%s", uniqueSuffix)
	currentBlockIndex := fmt.Sprintf("$current_block_index%s", uniqueSuffix)

	cc := codecreator.NewPhpCodeCreator()

	cc.AddLines(
		fmt.Sprintf("%[1]s = 0;", currentSize),
		fmt.Sprintf("%[1]s = $context_sizes->push_back(0);", currentSizeIndex),
	)

	cc.AddLines(
		fmt.Sprintf("%[1]s = 0;", blocksUsed),
		fmt.Sprintf("%[1]s = 0;", currentBlock),
		fmt.Sprintf("%[1]s = $context_blocks->push_back(0);", currentBlockIndex),
		fmt.Sprintf("%[1]s = %[2]s;", nextBlockIndex, currentBlockIndex),
	)

	index := 0
	if trw.wr.PHPUnionParent() != nil {
		index = trw.wr.unionIndex
	}

	if index != 0 {
		cc.AddLines(
			fmt.Sprintf("%[1]s |= (1 << 0);", currentBlock),
			fmt.Sprintf("%[1]s += TL\\tl2_support::count_used_bytes(%[2]d);", currentSize, index),
			fmt.Sprintf("%[1]s = $context_blocks->get_current_size();", nextBlockIndex),
			fmt.Sprintf("%[1]s = 1;", blocksUsed),
		)
	}

	fieldSize := fmt.Sprintf("$field_size%s", uniqueSuffix)
	for i, field := range trw.Fields {
		indexInBlock := (i + 1) % 8
		if indexInBlock == 0 {
			cc.AddLines(
				"// add new block",
				fmt.Sprintf("$context_blocks->set_value(%[1]s, %[2]s);", currentBlockIndex, currentBlock),
				fmt.Sprintf("%[1]s = 0;", currentBlock),
				fmt.Sprintf("%[1]s = $context_blocks->push_back(0);", currentBlockIndex),
			)
		}
		if field.IsTL2Omitted() {
			continue
		}

		var fieldUsed string

		fieldTarget := fmt.Sprintf("%s->%s", targetName, field.originalName)
		if useItself {
			fieldTarget = targetName
		}

		tree := trw.PHPGetFieldNatDependenciesValuesAsTypeTree(i, args)

		cc.AddLines(fmt.Sprintf("// calculate field %s", field.originalName))
		cc.AddLines(fmt.Sprintf("%[1]s = 0;", fieldSize))
		if field.IsBit() {
			fieldUsed = fmt.Sprintf("%[1]s", fieldTarget)
		} else {
			fieldUsed = fmt.Sprintf("%[1]s != 0", fieldSize)
			if field.MaskTL2Bit != nil {
				cc.AddLines(fmt.Sprintf("if (%[1]s !== null) {", fieldTarget))
				cc.AddShift(1)
			}
			cc.AddLines(
				field.t.trw.PhpCalculateSizesTL2MethodCall(fieldTarget, false, &tree, fmt.Sprintf("%s_%d", supportSuffix, i), callLevel+1, fieldSize, true)...,
			)
			if field.MaskTL2Bit != nil {
				cc.AddShift(-1)
				cc.AddLines("}")
			}
		}

		cc.If(fieldUsed, func(cc *codecreator.PhpCodeCreator) {
			cc.AddLines(
				fmt.Sprintf("%[1]s |= (1 << %[2]d);", currentBlock, indexInBlock),
				fmt.Sprintf("%[1]s += %[2]s;", currentSize, fieldSize),
				fmt.Sprintf("%[1]s = $context_blocks->get_current_size();", nextBlockIndex),
				fmt.Sprintf("%[1]s = %[2]d;", blocksUsed, 1+(i+1)/8),
			)
		})
	}

	// add last block
	cc.AddLines(
		fmt.Sprintf("$context_blocks->set_value(%[1]s, %[2]s);", currentBlockIndex, currentBlock),
	)

	// add blocks
	cc.AddLines(fmt.Sprintf("%[1]s += %[2]s;", currentSize, blocksUsed))
	// remove tail of sizes if is zero
	cc.IfElse(
		fmt.Sprintf("%[1]s == 0", currentSize),
		func(cc *codecreator.PhpCodeCreator) {
			cc.AddLines(fmt.Sprintf("$context_sizes->cut_tail(%s + 1);", currentSizeIndex))
			if !canOmit {
				cc.AddLines(fmt.Sprintf("%[1]s = 1;", usedBytesPointer))
			}
		},
		func(cc *codecreator.PhpCodeCreator) {
			cc.AddLines(
				fmt.Sprintf("$context_sizes->set_value(%[1]s, %[2]s);", currentSizeIndex, currentSize),
				fmt.Sprintf("%[1]s += %[2]s + TL\\tl2_support::count_used_bytes(%[2]s);", usedBytesPointer, currentSize),
			)
		},
	)

	// remove unused blocks
	cc.AddLines(fmt.Sprintf("$context_blocks->cut_tail(%[1]s);", nextBlockIndex))
	// add used size
	return cc.Print()
}

func (trw *TypeRWStruct) PHPStructFieldMaskCalculators(code *strings.Builder, usedFieldMasksIndecies []int, usedFieldMasks map[int][]Field) {
	// print methods to calculate fieldmasks
	// fix order
	names := utils.MapSlice(usedFieldMasksIndecies, func(natIndex int) string {
		natName := ""
		if natIndex < 0 {
			natName = trw.wr.origTL[0].TemplateArguments[-(natIndex + 1)].FieldName
		} else {
			natName = trw.Fields[natIndex].originalName
		}
		return natName
	})

	namesToIndices := make(map[string]int)
	for i := range names {
		namesToIndices[names[i]] = usedFieldMasksIndecies[i]
	}
	sort.Strings(names)

	fieldNameToFieldOrder := make(map[string]int)
	for i := range trw.Fields {
		fieldNameToFieldOrder[trw.Fields[i].originalName] = i
	}

	for _, name := range names {
		natIndex := namesToIndices[name]
		natName := name
		code.WriteString(`
  /**`)
		additionalArgs := make([]string, 0)
		// arguments with ambiguous existence
		for _, dependentField := range usedFieldMasks[natIndex] {
			if _, isMaybe := dependentField.t.PHPGenCoreType().trw.(*TypeRWMaybe); isMaybe {
				additionalArgs = append(additionalArgs, fmt.Sprintf("$has_%s", dependentField.originalName))
				code.WriteString(fmt.Sprintf("\n   * @param bool $has_%s", dependentField.originalName))
			}
		}
		code.WriteString(`
   * @return int
   */
`,
		)
		code.WriteString(
			fmt.Sprintf(
				"  public function calculate%[1]s(%[2]s) {\n    $mask = 0;\n",
				toPhpFieldMaskName(natName),
				strings.Join(additionalArgs, ", "),
			),
		)

		fields := usedFieldMasks[natIndex]
		sort.Slice(fields, func(i, j int) bool {
			if fields[i].BitNumber == fields[j].BitNumber {
				return i < j
			}
			return fields[i].BitNumber < fields[j].BitNumber
		})

		fieldsGroupedByBitNumber := make([][]Field, 0)
		for _, dependentField := range fields {
			if len(fieldsGroupedByBitNumber) == 0 ||
				fieldsGroupedByBitNumber[len(fieldsGroupedByBitNumber)-1][0].BitNumber != dependentField.BitNumber {
				fieldsGroupedByBitNumber = append(fieldsGroupedByBitNumber, make([]Field, 0))
			}
			fieldsGroupedByBitNumber[len(fieldsGroupedByBitNumber)-1] = append(fieldsGroupedByBitNumber[len(fieldsGroupedByBitNumber)-1], dependentField)
		}

		for _, dependentFields := range fieldsGroupedByBitNumber {
			conditions := make([]string, 0)
			bitConstants := make([]string, 0)
			sort.Slice(dependentFields, func(i, j int) bool {
				return fieldNameToFieldOrder[dependentFields[i].originalName] < fieldNameToFieldOrder[dependentFields[j].originalName]
			})
			for _, dependentField := range dependentFields {
				condition := ""
				if dependentField.t.PHPIsTrueType() || dependentField.t.PHPGenCoreType().PHPNeedsCode() {
					condition = fmt.Sprintf(
						"$this->%[1]s",
						dependentField.originalName,
					)
				} else if _, isMaybe := dependentField.t.PHPGenCoreType().trw.(*TypeRWMaybe); isMaybe {
					condition = fmt.Sprintf("$has_%s", dependentField.originalName)
				} else {
					condition = fmt.Sprintf(
						"$this->%[1]s !== null",
						dependentField.originalName,
					)
				}
				conditions = append(conditions, condition)
				bitConstants = append(bitConstants, fmt.Sprintf(
					"self::BIT_%[1]s_%[2]d",
					strings.ToUpper(dependentField.originalName),
					dependentField.BitNumber))
			}

			finalCondition := conditions[0]
			finalMask := bitConstants[0]

			if len(conditions) > 1 {
				finalCondition = strings.Join(conditions, " && ")
				finalMask = "(" + strings.Join(bitConstants, " | ") + ")"
			}

			code.WriteString(
				fmt.Sprintf(
					`
    if (%[1]s) {
      $mask |= %[2]s;
    }
`,
					finalCondition,
					finalMask,
				),
			)
		}

		code.WriteString("\n    return $mask;\n")
		code.WriteString("  }\n")
	}
}

func (trw *TypeRWStruct) PHPStructConstructor(code *strings.Builder, necessaryFieldsInConstructor []Field) {
	// print constructor
	code.WriteString(`
  /**
`)
	for _, f := range necessaryFieldsInConstructor {
		fieldType, _ := fieldTypeAndDefaultValue(f)
		code.WriteString(fmt.Sprintf("   * @param %[1]s $%[2]s\n", fieldType, f.originalName))
	}
	if len(necessaryFieldsInConstructor) == 0 {
		code.WriteString("   * @kphp-inline\n")
	}

	code.WriteString(`   */
`)
	code.WriteString("  public function __construct(")

	for i, f := range necessaryFieldsInConstructor {
		_, defaultValue := fieldTypeAndDefaultValue(f)
		if i != 0 {
			code.WriteString(", ")
		}
		code.WriteString(fmt.Sprintf("$%[1]s = %[2]s", f.originalName, defaultValue))
	}

	code.WriteString(") {\n")
	for _, f := range necessaryFieldsInConstructor {
		code.WriteString(fmt.Sprintf("    $this->%[1]s = $%[1]s;\n", f.originalName))
	}
	code.WriteString("  }\n")
}

func (trw *TypeRWStruct) PHPStructRPCSpecialGetters(code *strings.Builder) {
	if !trw.wr.gen.options.AddRPCTypes {
		return
	}
	if unionParent := trw.wr.PHPUnionParent(); unionParent == nil || PHPSpecialMembersTypes(unionParent) == "" {
		return
	}

	const ThisType = "__this"
	type SpecialField struct {
		Name                string
		Type                string
		Default             string
		NullTypeIfNullValue bool
		AddHasMethod        bool
	}

	fields := []SpecialField{
		{
			"result",
			"TL\\RpcFunctionReturnResult",
			"null",
			true,
			false,
		},
		{
			"header",
			ThisType,
			"null",
			true,
			false,
		},
		{
			"error",
			ThisType,
			"null",
			true,
			true,
		},
	}

	containsSuchField := func(name, ifTrue, ifFalse string) string {
		for _, field := range trw.Fields {
			if field.originalName == name {
				return ifTrue
			}
		}
		return ifFalse
	}

	for _, field := range fields {
		returnObject := field.Default
		returnType := field.Default
		if field.Type == ThisType &&
			strings.Contains(strings.ToLower(trw.PhpClassName(false, true)), strings.ToLower(field.Name)) {
			returnObject = "$this"
			returnType = trw.PhpTypeName(true, true)
		} else {
			if field.Type != ThisType {
				returnObject = "$this->" + field.Name
				returnType = field.Type
			}
			if field.NullTypeIfNullValue {
				returnType = containsSuchField(field.Name, returnType, "null")
				returnObject = containsSuchField(field.Name, returnObject, "null")
			}
		}
		if field.AddHasMethod {
			code.WriteString(
				fmt.Sprintf(
					`
  /**
   * @return bool
   */
  public function is%[1]s() {
    return %[2]s;
  }
`,
					ToUpperFirst(field.Name),
					containsSuchField(field.Name, "true", "false"),
				),
			)
		}
		code.WriteString(
			fmt.Sprintf(
				`
  /**
   * @return %[3]s
   */
  public function get%[1]s() {
    return %[2]s;
  }
`,
				ToUpperFirst(field.Name),
				returnObject,
				returnType,
			),
		)
	}
}

func (trw *TypeRWStruct) PHPStructResultType(code *strings.Builder) {
	// print result type for function
	if trw.ResultType != nil {
		code.WriteString(
			fmt.Sprintf(
				`
  /** Allows kphp implicitly load function result class */
  private const RESULT = %s_result::class;
`,
				trw.PhpClassName(true, true),
			),
		)
	}
}

func (trw *TypeRWStruct) PHPStructFields(code *strings.Builder) {
	// print fields declarations
	for _, f := range trw.Fields {
		fieldType, defaultValue := fieldTypeAndDefaultValue(f)
		code.WriteString(
			fmt.Sprintf(
				`
  /** @var %[1]s */
  public $%[2]s = %[3]s;
`,
				fieldType,
				f.originalName,
				defaultValue,
			),
		)
	}
}

func (trw *TypeRWStruct) PHPStructFieldMasks(code *strings.Builder) {
	// print fieldmasks
	for _, f := range trw.Fields {
		if f.fieldMask == nil {
			continue
		}
		code.WriteString(
			fmt.Sprintf(
				`
  /** Field mask for $%[1]s field */
  const BIT_%[2]s_%[3]d = (1 << %[3]d);
`,
				f.originalName,
				strings.ToUpper(f.originalName),
				f.BitNumber,
			),
		)
	}
}

func (trw *TypeRWStruct) PHPStructHeader(code *strings.Builder) {
	unionParent := trw.PhpConstructorNeedsUnion()

	if isUsingTLImport(trw) ||
		trw.ResultType != nil ||
		unionParent != nil ||
		(trw.wr.gen.options.AddFunctionBodies && trw.wr.phpInfo.RequireFunctionBodies) {
		code.WriteString("\nuse VK\\TL;\n")
	}
	code.WriteString(`
/**
 * @kphp-tl-class
 */
`)
	code.WriteString(fmt.Sprintf("class %s ", trw.PhpClassName(false, true)))
	implementingInterfaces := make([]string, 0)

	if unionParent != nil {
		implementingInterfaces = append(implementingInterfaces, unionParent.trw.PhpClassName(true, false))
	}

	if trw.wr.gen.options.AddRPCTypes && trw.ResultType != nil {
		implementingInterfaces = append(implementingInterfaces, "TL\\RpcFunction")
	}

	if trw.wr.gen.options.AddFunctionBodies &&
		trw.wr.phpInfo.RequireFunctionBodies &&
		len(trw.wr.origTL[0].TemplateArguments) == 0 &&
		!trw.wr.gen.options.UseBuiltinDataProviders {
		implementingInterfaces = append(implementingInterfaces, "TL\\Readable")
		implementingInterfaces = append(implementingInterfaces, "TL\\Writeable")
	}

	if len(implementingInterfaces) != 0 {
		code.WriteString("implements ")
		code.WriteString(strings.Join(implementingInterfaces, ", "))
		code.WriteString(" ")
	}
	code.WriteString("{\n")
}

func phpResultType(trw *TypeRWStruct) string {
	return trw.ResultType.trw.PhpTypeName(true, trw.ResultType.PHPIsBare())
}

func toPhpFieldMaskName(natName string) string {
	parts := strings.Split(natName, "_")
	for i := range parts {
		parts[i] = ToUpperFirst(parts[i])
	}
	return strings.Join(parts, "")
}

func isUsingTLImport(trw *TypeRWStruct) bool {
	for _, field := range trw.Fields {
		fieldType, _ := fieldTypeAndDefaultValue(field)
		if strings.Contains(fieldType, "TL\\") {
			return true
		}
	}
	return false
}

func fieldTypeAndDefaultValue(f Field) (string, string) {
	fieldType := f.t.trw.PhpTypeName(true, f.t.PHPIsBare())
	defaultValue := f.t.trw.PhpDefaultValue()
	if f.t.PHPIsTrueType() {
		fieldType = "boolean"
		defaultValue = "true"
		if f.fieldMask != nil {
			defaultValue = "false"
		}
	} else {
		if f.fieldMask != nil {
			defaultValue = "null"
			if _, isMaybe := f.t.PHPGenCoreType().trw.(*TypeRWMaybe); !isMaybe {
				fieldType = fieldType + "|null"
			}
		}
	}
	return fieldType, defaultValue
}

func (trw *TypeRWStruct) PhpDefaultValue() string {
	core := trw.wr.PHPGenCoreType()
	if core != trw.wr {
		return core.PHPDefaultValue()
	}
	if core.PHPIsTrueType() {
		return "true"
	}
	if !trw.wr.gen.options.InplaceSimpleStructs &&
		strings.HasSuffix(trw.wr.tlName.String(), "dictionary") &&
		trw.wr.tlName.Namespace == "" {
		return trw.Fields[0].t.trw.PhpDefaultValue()
	}
	return "null"
}

func (trw *TypeRWStruct) PhpIterateReachableTypes(reachableTypes *map[*TypeRWWrapper]bool) {
	for _, field := range trw.Fields {
		field.t.PhpIterateReachableTypes(reachableTypes)
	}
	if trw.ResultType != nil {
		trw.ResultType.PhpIterateReachableTypes(reachableTypes)
	}
}

func (trw *TypeRWStruct) PhpConstructorNeedsUnion() (unionParent *TypeRWWrapper) {
	if trw.ResultType == nil {
		if trw.wr.unionParent != nil {
			return trw.wr.unionParent.wr
		} else if !strings.EqualFold(trw.wr.tlName.Name, trw.wr.origTL[0].TypeDecl.Name.Name) {
			// NOTE: constructor name is not same as type => type can become union in future?
			return trw.wr
		}
	}

	return nil
}

func (trw *TypeRWStruct) PhpReadMethodCall(targetName string, bare bool, initIfDefault bool, args *TypeArgumentsTree, supportSuffix string) []string {
	useBuiltIn := trw.wr.gen.options.UseBuiltinDataProviders
	if specialCase := PHPSpecialMembersTypes(trw.wr); specialCase != "" {
		return []string{
			"/** TODO */",
			fmt.Sprintf("/** $success = RPC_READ%s(%s%s); */",
				ifString(bare, "", "_boxed"),
				ifString(useBuiltIn, "", "$stream, "),
				targetName,
			),
		}
	}
	unionParent := trw.PhpConstructorNeedsUnion()
	if unionParent == nil {
		if trw.PhpCanBeSimplify() {
			var result []string
			if !bare {
				result = trw.phpStructReadMagic(useBuiltIn, result)
			}
			newArgs := trw.PHPGetFieldNatDependenciesValuesAsTypeTree(0, args)
			result = append(result, trw.Fields[0].t.trw.PhpReadMethodCall(targetName, trw.Fields[0].bare, initIfDefault, &newArgs, supportSuffix)...)
			return result
		}
		if trw.ResultType == nil && trw.wr.PHPIsTrueType() {
			var result []string
			if !bare {
				result = trw.phpStructReadMagic(useBuiltIn, result)
			}
			result = append(result, fmt.Sprintf("%[1]s = true;", targetName))
			return result
		}
		//isDict, _, _, valueType := isDictionaryElement(trw.wr)
		//if isDict && trw.wr.tlName.Namespace == "" { // TODO NOT A SOLUTION, BUT...
		//	return valueType.t.trw.PhpTypeName(withPath, bare)
		//}
		if !trw.wr.gen.options.InplaceSimpleStructs &&
			strings.HasSuffix(trw.wr.tlName.String(), "dictionary") &&
			trw.wr.tlName.Namespace == "" {
			var result []string
			if !bare {
				result = trw.phpStructReadMagic(useBuiltIn, result)
			}
			result = append(result, trw.Fields[0].t.trw.PhpReadMethodCall(targetName, bare, initIfDefault, args, supportSuffix)...)
			return result
		}
	}
	result := make([]string, 0)
	if initIfDefault {
		result = append(result,
			fmt.Sprintf("if (is_null(%[1]s)) {", targetName),
			fmt.Sprintf("  %[1]s = %[2]s;", targetName, trw.PhpDefaultInit()),
			"}",
		)
	}
	if trw.wr.phpInfo.IsDuplicate {
		if !bare {
			result = trw.phpStructReadMagic(useBuiltIn, result)
		}
		result = append(result, trw.phpStructReadCode(targetName, args)...)
	} else {
		result = append(result,
			fmt.Sprintf("$success = %[2]s->read%[1]s(%[4]s%[3]s);",
				ifString(bare, "", "_boxed"),
				targetName,
				phpFormatArgs(args.ListAllValues(), useBuiltIn),
				ifString(useBuiltIn, "", "$stream"),
			),
			"if (!$success) {",
			"  return false;",
			"}",
		)
	}
	return result
}

func (trw *TypeRWStruct) phpStructReadMagic(useBuiltIn bool, result []string) []string {
	if useBuiltIn {
		result = append(result,
			"$magic = fetch_int() & 0xFFFFFFFF;",
			fmt.Sprintf("if ($magic != 0x%08[1]x) {", trw.wr.tlTag),
			"  return false;",
			"}",
		)
	} else {
		result = append(result,
			"[$magic, $success] = $stream->read_uint32();",
			fmt.Sprintf("if (!$success || $magic != 0x%08[1]x) {", trw.wr.tlTag),
			"  return false;",
			"}",
		)
	}
	return result
}

func (trw *TypeRWStruct) PhpWriteMethodCall(targetName string, bare bool, args *TypeArgumentsTree, supportSuffix string) []string {
	useBuiltIn := trw.wr.gen.options.UseBuiltinDataProviders
	if specialCase := PHPSpecialMembersTypes(trw.wr); specialCase != "" {
		return []string{
			"/** TODO */",
			fmt.Sprintf("/** $success = RPC_WRITE%s(%s%s); */",
				ifString(bare, "", "_boxed"),
				ifString(useBuiltIn, "", "$stream, "),
				targetName,
			),
		}
	}
	unionParent := trw.PhpConstructorNeedsUnion()
	if unionParent == nil {
		if trw.PhpCanBeSimplify() {
			var result []string
			if !bare {
				result = trw.phpStructWriteMagic(useBuiltIn, result)
			}
			newArgs := trw.PHPGetFieldNatDependenciesValuesAsTypeTree(0, args)
			result = append(result, trw.Fields[0].t.trw.PhpWriteMethodCall(targetName, trw.Fields[0].bare, &newArgs, supportSuffix)...)
			return result
		}
		if trw.ResultType == nil && trw.wr.PHPIsTrueType() {
			var result []string
			if !bare {
				result = trw.phpStructWriteMagic(useBuiltIn, result)
			}
			return result
		}
		//isDict, _, _, valueType := isDictionaryElement(trw.wr)
		//if isDict && trw.wr.tlName.Namespace == "" { // TODO NOT A SOLUTION, BUT...
		//	return valueType.t.trw.PhpTypeName(withPath, bare)
		//}
		if !trw.wr.gen.options.InplaceSimpleStructs &&
			strings.HasSuffix(trw.wr.tlName.String(), "dictionary") &&
			trw.wr.tlName.Namespace == "" {
			var result []string
			if !bare {
				result = trw.phpStructWriteMagic(useBuiltIn, result)
			}
			result = append(result, trw.Fields[0].t.trw.PhpWriteMethodCall(targetName, bare, args, supportSuffix)...)
			return result
		}
	}
	result := []string{
		fmt.Sprintf("if (is_null(%[1]s)) {", targetName),
		fmt.Sprintf("  %[1]s = %[2]s;", targetName, trw.PhpDefaultInit()),
		"}",
	}
	if trw.wr.phpInfo.IsDuplicate {
		if !bare {
			result = trw.phpStructWriteMagic(useBuiltIn, result)
		}
		result = append(result, trw.phpStructWriteCode(targetName, args)...)
	} else {
		result = append(result,
			fmt.Sprintf(
				"$success = %[2]s->write%[1]s(%[4]s%[3]s);",
				ifString(bare, "", "_boxed"),
				targetName,
				phpFormatArgs(args.ListAllValues(), useBuiltIn),
				ifString(useBuiltIn, "", "$stream"),
			),
			"if (!$success) {",
			"  return false;",
			"}",
		)
	}
	return result
}

func (trw *TypeRWStruct) PhpReadTL2MethodCall(targetName string, bare bool, initIfDefault bool, args *TypeArgumentsTree, supportSuffix string, callLevel int, usedBytesPointer string, canDependOnLocalBit bool) []string {
	if !trw.wr.gen.options.UseBuiltinDataProviders {
		panic("generation tl2 for non builtin data providers is forbidden")
	}
	if specialCase := PHPSpecialMembersTypes(trw.wr); specialCase != "" {
		return []string{"// TODO"}
	}
	unionParent := trw.PhpConstructorNeedsUnion()
	if unionParent == nil {
		if trw.PhpCanBeSimplify() {
			newArgs := trw.PHPGetFieldNatDependenciesValuesAsTypeTree(0, args)
			if trw.isUnwrapType() {
				// inplace
				readText := trw.Fields[0].t.trw.PhpReadTL2MethodCall(targetName, trw.Fields[0].bare, initIfDefault, &newArgs, supportSuffix, callLevel+1, usedBytesPointer, canDependOnLocalBit)
				return readText
			} else {
				localUsedBytesPointer := fmt.Sprintf("$used_bytes_%[1]s_%[2]d", supportSuffix, callLevel)
				localCurrentSize := fmt.Sprintf("$current_size_%[1]s_%[2]d", supportSuffix, callLevel)
				localBlock := fmt.Sprintf("$block_%[1]s_%[2]d", supportSuffix, callLevel)
				localConstructor := fmt.Sprintf("$index_%[1]s_%[2]d", supportSuffix, callLevel)

				readText := trw.Fields[0].t.trw.PhpReadTL2MethodCall(targetName, trw.Fields[0].bare, initIfDefault, &newArgs, supportSuffix, callLevel+1, localUsedBytesPointer, canDependOnLocalBit)

				var result []string
				result = append(result,
					fmt.Sprintf("%[1]s = TL\\tl2_support::fetch_size();", localCurrentSize),
					fmt.Sprintf("%[1]s = 0;", localUsedBytesPointer),
					// add to global pointer
					fmt.Sprintf("%[1]s += %[2]s + TL\\tl2_support::count_used_bytes(%[2]s);", usedBytesPointer, localCurrentSize),
					// decide should we read body
					fmt.Sprintf("if (%[1]s != 0) {", localCurrentSize),
					fmt.Sprintf("  %[1]s = fetch_byte();", localBlock),
					fmt.Sprintf("  %[1]s += 1;", localUsedBytesPointer),
					fmt.Sprintf("  if ((%[1]s & 1) != 0) {", localBlock),
					fmt.Sprintf("    %[1]s = TL\\tl2_support::fetch_size();", localConstructor),
					fmt.Sprintf("    %[1]s += TL\\tl2_support::count_used_bytes(%[2]s);", localUsedBytesPointer, localConstructor),
					"  }",
					fmt.Sprintf("  if ((%[1]s & (1 << 1)) != 0) {", localBlock),
				)
				for _, line := range readText {
					result = append(result, "    "+line)
				}
				result = append(result,
					"  }",
					"}",
				)
				result = append(result, fmt.Sprintf("%[1]s += TL\\tl2_support::skip_bytes(%[2]s - %[1]s);", localUsedBytesPointer, localCurrentSize))
				return result
			}
		}
		if trw.ResultType == nil && trw.wr.PHPIsTrueType() {
			var result []string
			result = append(result, fmt.Sprintf("%[1]s = true;", targetName))
			return result
		}
		if !trw.wr.gen.options.InplaceSimpleStructs &&
			strings.HasSuffix(trw.wr.tlName.String(), "dictionary") &&
			trw.wr.tlName.Namespace == "" {
			var result []string
			result = append(result, trw.Fields[0].t.trw.PhpReadTL2MethodCall(targetName, bare, initIfDefault, args, supportSuffix, callLevel+1, usedBytesPointer, canDependOnLocalBit)...)
			return result
		}
	}
	result := make([]string, 0)
	additionalArguments := make([]string, 0)
	if unionParent != nil {
		currentSize := fmt.Sprintf("$current_size_%[1]s_%[2]d", supportSuffix, callLevel)
		currentBlock := fmt.Sprintf("$current_size_%[1]s_%[2]d", supportSuffix, callLevel)

		additionalArguments = append(additionalArguments, currentBlock, currentSize)

		cc := codecreator.CodeCreator{Shift: "  "}
		cc.AddLines(
			fmt.Sprintf("%[1]s = TL\\tl2_support::fetch_size();", currentSize),
			fmt.Sprintf("%[1]s += TL\\tl2_support::count_used_bytes(%[2]s);", usedBytesPointer, currentSize),
		)

		cc.AddLines(fmt.Sprintf("%[1]s = 0;", currentBlock))

		cc.AddLines(fmt.Sprintf("if (%[1]s != 0) {", currentSize))
		cc.AddBlock(func(cc *codecreator.CodeCreator) {
			cc.AddLines(fmt.Sprintf("%[1]s = fetch_byte();", currentBlock))
			cc.AddLines(fmt.Sprintf("%[1]s += 1;", usedBytesPointer))
		})
		cc.AddLines("}")
		result = append(result, cc.Print()...)
	}
	if initIfDefault {
		result = append(result,
			fmt.Sprintf("if (is_null(%[1]s)) {", targetName),
			fmt.Sprintf("  %[1]s = %[2]s;", targetName, trw.PhpDefaultInit()),
			"}",
		)
	}
	if trw.wr.phpInfo.IsDuplicate {
		result = append(result, trw.phpStructReadTL2Code(targetName, usedBytesPointer, args, false)...)
	} else {
		result = append(result,
			fmt.Sprintf("%[3]s += %[1]s->read_tl2(%[2]s);",
				targetName,
				phpFormatArgs(utils.Append(args.ListAllValues(), additionalArguments...), true),
				usedBytesPointer,
			),
		)
	}
	return result
}

func (trw *TypeRWStruct) PhpWriteTL2MethodCall(targetName string, bare bool, args *TypeArgumentsTree, supportSuffix string, callLevel int, usedBytesPointer string, canDependOnLocalBit bool) []string {
	if !trw.wr.gen.options.UseBuiltinDataProviders {
		panic("generation tl2 for non builtin data providers is forbidden")
	}
	if specialCase := PHPSpecialMembersTypes(trw.wr); specialCase != "" {
		return []string{"// TODO"}
	}
	unionParent := trw.PhpConstructorNeedsUnion()
	if unionParent == nil {
		if trw.PhpCanBeSimplify() {
			newArgs := trw.PHPGetFieldNatDependenciesValuesAsTypeTree(0, args)
			if trw.isUnwrapType() {
				calcText := trw.Fields[0].t.trw.PhpWriteTL2MethodCall(targetName, trw.Fields[0].bare, &newArgs, supportSuffix, callLevel+1, usedBytesPointer, canDependOnLocalBit)
				return calcText
			} else {
				calcText := trw.phpStructWriteTL2Code(targetName, args, supportSuffix, callLevel, usedBytesPointer, true)
				return calcText
			}
		}
		if trw.ResultType == nil && trw.wr.PHPIsTrueType() {
			var result []string
			return result
		}
		if !trw.wr.gen.options.InplaceSimpleStructs &&
			strings.HasSuffix(trw.wr.tlName.String(), "dictionary") &&
			trw.wr.tlName.Namespace == "" {
			newArgs := trw.PHPGetFieldNatDependenciesValuesAsTypeTree(0, args)
			calcText := trw.Fields[0].t.trw.PhpWriteTL2MethodCall(targetName, trw.Fields[0].bare, &newArgs, supportSuffix, callLevel+1, usedBytesPointer, canDependOnLocalBit)
			return calcText
		}
	}
	result := make([]string, 0)
	result = append(result,
		fmt.Sprintf("if (is_null(%[1]s)) {", targetName),
		fmt.Sprintf("  %[1]s = %[2]s;", targetName, trw.PhpDefaultInit()),
		"}",
	)
	if trw.wr.phpInfo.IsDuplicate {
		result = append(result, trw.phpStructWriteTL2Code(targetName, args, supportSuffix, callLevel, usedBytesPointer, false)...)
	} else {
		result = append(result,
			fmt.Sprintf("%[3]s += %[1]s->internal_write_tl2(%[2]s);",
				targetName,
				phpFormatArgs(utils.Append(args.ListAllValues(), "$context_sizes", "$context_blocks"), true),
				usedBytesPointer,
			),
		)
	}
	return result
}

func (trw *TypeRWStruct) PhpCalculateSizesTL2MethodCall(targetName string, bare bool, args *TypeArgumentsTree, supportSuffix string, callLevel int, usedBytesPointer string, canOmit bool) []string {
	if !trw.wr.gen.options.UseBuiltinDataProviders {
		panic("generation tl2 for non builtin data providers is forbidden")
	}
	if specialCase := PHPSpecialMembersTypes(trw.wr); specialCase != "" {
		return []string{"// TODO"}
	}
	unionParent := trw.PhpConstructorNeedsUnion()
	if unionParent == nil {
		if trw.PhpCanBeSimplify() {
			newArgs := trw.PHPGetFieldNatDependenciesValuesAsTypeTree(0, args)
			if trw.isUnwrapType() {
				calcText := trw.Fields[0].t.trw.PhpCalculateSizesTL2MethodCall(targetName, trw.Fields[0].bare, &newArgs, supportSuffix, callLevel+1, usedBytesPointer, canOmit)
				return calcText
			} else {
				calcText := trw.phpStructCalculateSizesTL2Code(targetName, args, supportSuffix, callLevel, usedBytesPointer, true, canOmit)
				return calcText
			}
		}
		if trw.ResultType == nil && trw.wr.PHPIsTrueType() {
			var result []string
			return result
		}
	}
	result := make([]string, 0)
	result = append(result,
		fmt.Sprintf("if (is_null(%[1]s)) {", targetName),
		fmt.Sprintf("  %[1]s = %[2]s;", targetName, trw.PhpDefaultInit()),
		"}",
	)
	if trw.wr.phpInfo.IsDuplicate {
		result = append(result, trw.phpStructCalculateSizesTL2Code(targetName, args, supportSuffix, callLevel, usedBytesPointer, false, canOmit)...)
	} else {
		localSize := fmt.Sprintf("$local_size_%[1]s_%[2]d", supportSuffix, callLevel)

		cc := codecreator.NewPhpCodeCreator()
		cc.AddLines(
			fmt.Sprintf("%[1]s = 0;", localSize),
			fmt.Sprintf("%[3]s += %[1]s->calculate_sizes_tl2(%[2]s);",
				targetName,
				phpFormatArgs(utils.Append(args.ListAllValues(), "$context_sizes", "$context_blocks"), true),
				localSize,
			),
		)
		if !canOmit {
			cc.If(fmt.Sprintf("%[1]s == 0", localSize), func(cc *codecreator.BasicCodeCreator[codecreator.PhpHelder]) {
				cc.AddLines(fmt.Sprintf("%[1]s = 1;", localSize))
			})
		} else {
			// remove itself
			cc.If(fmt.Sprintf("%[1]s == 0", localSize), func(cc *codecreator.BasicCodeCreator[codecreator.PhpHelder]) {
				cc.AddLines("$context_sizes->cut_tail($context_sizes->get_current_size() - 1);")
			})
		}
		cc.AddLines(
			fmt.Sprintf("%[1]s += %[2]s;", usedBytesPointer, localSize),
		)

		result = append(result, cc.Print()...)
	}
	return result
}

func (trw *TypeRWStruct) phpStructWriteMagic(useBuiltIn bool, result []string) []string {
	if useBuiltIn {
		result = append(result, fmt.Sprintf("store_int(0x%08[1]x);", trw.wr.tlTag))
	} else {
		result = append(result,
			fmt.Sprintf("$success = $stream->write_uint32(0x%08[1]x);", trw.wr.tlTag),
			"if (!$success) {",
			"  return false;",
			"}",
		)
	}
	return result
}

func (trw *TypeRWStruct) PhpDefaultInit() string {
	core := trw.wr.PHPGenCoreType()
	if core != trw.wr {
		return core.trw.PhpDefaultInit()
	}
	if core.PHPIsTrueType() {
		return "true"
	}
	if !trw.wr.gen.options.InplaceSimpleStructs &&
		strings.HasSuffix(trw.wr.tlName.String(), "dictionary") &&
		trw.wr.tlName.Namespace == "" {
		return trw.Fields[0].t.trw.PhpDefaultInit()
	}
	return fmt.Sprintf("new %s()", core.trw.PhpClassName(true, true))
}

func (trw *TypeRWStruct) PhpCanBeSimplify() bool {
	return len(trw.Fields) == 1 &&
		trw.ResultType == nil &&
		trw.Fields[0].fieldMask == nil &&
		(trw.wr.gen.options.InplaceSimpleStructs || trw.Fields[0].t.PHPIsPrimitiveType(false))
}
