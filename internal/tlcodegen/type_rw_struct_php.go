// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlcodegen

import (
	"fmt"
	"github.com/vkcom/tl/internal/tlast"
	"github.com/vkcom/tl/internal/utils"
	"sort"
	"strconv"
	"strings"
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

func (trw *TypeRWStruct) PHPGetFieldMask(fieldIndex int) string {
	fieldMask := trw.Fields[fieldIndex].fieldMask
	if fieldMask != nil {
		if fieldMask.isField {
			return fmt.Sprintf("$this->%s", trw.Fields[fieldMask.FieldIndex].originalName)
		}
		return "$" + fieldMask.name
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

  /**
   * @kphp-inline
   *
   * @return string
   */
  public function getTLFunctionName() {
    return '%[3]s';
  }
`,
				trw.PhpClassName(false, true),
				trw.PhpClassName(true, true),
				trw.wr.tlName.String(),
				phpResultType(trw),
				kphpSpecialCode,
			),
		)

	}
}

func (trw *TypeRWStruct) PHPStructReadMethods(code *strings.Builder) {
	if trw.wr.gen.options.AddFunctionBodies {
		natParamsComment := strings.Join(
			utils.MapSlice(
				trw.wr.PHPGetNatTypeDependenciesDeclAsArray(),
				func(s string) string { return fmt.Sprintf("\n   * @param int %s", s) }),
			"",
		)
		natParamsDecl := strings.Join(
			utils.MapSlice(
				trw.wr.PHPGetNatTypeDependenciesDeclAsArray(),
				func(s string) string { return ", " + s }),
			"",
		)
		code.WriteString(fmt.Sprintf(`
  /**
   * @param TL\tl_input_stream $stream%[1]s
   * @return bool 
   */
  public function read_boxed($stream%[2]s) {
    [$magic, $success] = $stream->read_uint32();
    if (!$success || $magic != 0x%08[3]x) {
      return false;
    }
    return $this->read($stream%[2]s);
  }
`,
			natParamsComment,
			natParamsDecl,
			trw.wr.tlTag,
		))

		code.WriteString(fmt.Sprintf(`
  /**
   * @param TL\tl_input_stream $stream%[1]s
   * @return bool 
   */
  public function read($stream%[2]s) {
`,
			natParamsComment,
			natParamsDecl,
			trw.wr.tlTag,
		))
		const tab = "  "
		for i, field := range trw.Fields {
			fieldMask := trw.PHPGetFieldMask(i)
			shift := 2
			textTab := func() string { return strings.Repeat(tab, shift) }
			if fieldMask != "" {
				code.WriteString(
					fmt.Sprintf(
						"%[1]sif ((%[2]s & (1 << %[3]d)) != 0) {\n",
						textTab(),
						fieldMask,
						field.BitNumber,
					),
				)
				shift += 1
			}
			tree := trw.PHPGetFieldNatDependenciesValuesAsTypeTree(i, nil)
			fieldRead := field.t.trw.PhpReadMethodCall("$this->"+field.originalName, field.bare, true, &tree)
			for _, line := range fieldRead {
				code.WriteString(textTab() + line + "\n")
			}
			if fieldMask != "" {
				shift -= 1
				code.WriteString(fmt.Sprintf("%[1]s} else {\n", textTab()))
				shift += 1
				_, defaultValue := fieldTypeAndDefaultValue(field)
				code.WriteString(fmt.Sprintf(
					"%[1]s%[2]s = %[3]s;\n",
					textTab(),
					"$this->"+field.originalName,
					defaultValue,
				))
				shift -= 1
				code.WriteString(fmt.Sprintf("%[1]s}\n", textTab()))
			}
		}

		code.WriteString("    return true;\n")
		code.WriteString("  }\n")
	}
}

func (trw *TypeRWStruct) PHPStructWriteMethods(code *strings.Builder) {
	if trw.wr.gen.options.AddFunctionBodies {
		natParamsComment := strings.Join(
			utils.MapSlice(
				trw.wr.PHPGetNatTypeDependenciesDeclAsArray(),
				func(s string) string { return fmt.Sprintf("\n   * @param int %s", s) }),
			"",
		)
		natParamsDecl := strings.Join(
			utils.MapSlice(
				trw.wr.PHPGetNatTypeDependenciesDeclAsArray(),
				func(s string) string { return ", " + s }),
			"",
		)
		code.WriteString(fmt.Sprintf(`
  /**
   * @param TL\tl_output_stream $stream%[1]s
   * @return bool 
   */
  public function write_boxed($stream%[2]s) {
    $success = $stream->write_uint32(0x%08[3]x);
    if (!$success) {
      return false;
    }
    return $this->write($stream%[2]s);
  }
`,
			natParamsComment,
			natParamsDecl,
			trw.wr.tlTag,
		))

		code.WriteString(fmt.Sprintf(`
  /**
   * @param TL\tl_output_stream $stream%[1]s
   * @return bool 
   */
  public function write($stream%[2]s) {
`,
			natParamsComment,
			natParamsDecl,
			trw.wr.tlTag,
		))
		const tab = "  "
		for i, field := range trw.Fields {
			fieldMask := trw.PHPGetFieldMask(i)
			shift := 2
			textTab := func() string { return strings.Repeat(tab, shift) }
			tree := trw.PHPGetFieldNatDependenciesValuesAsTypeTree(i, nil)
			fieldRead := field.t.trw.PhpWriteMethodCall("$this->"+field.originalName, field.bare, &tree)
			if fieldRead == nil {
				continue
			}
			if fieldMask != "" {
				code.WriteString(
					fmt.Sprintf(
						"%[1]sif ((%[2]s & (1 << %[3]d)) != 0) {\n",
						textTab(),
						fieldMask,
						field.BitNumber,
					),
				)
				shift += 1
			}
			for _, line := range fieldRead {
				code.WriteString(textTab() + line + "\n")
			}
			if fieldMask != "" {
				shift -= 1
				code.WriteString(fmt.Sprintf("%[1]s}\n", textTab()))
			}
		}

		code.WriteString("    return true;\n")
		code.WriteString("  }\n")
	}
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
				if dependentField.t.PHPIsTrueType() || dependentField.t.PHPNeedsCode() {
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
		trw.wr.gen.options.AddFunctionBodies {
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

	if trw.wr.gen.options.AddFunctionBodies && len(trw.wr.origTL[0].TemplateArguments) == 0 {
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

func (trw *TypeRWStruct) PhpReadMethodCall(targetName string, bare bool, initIfDefault bool, args *TypeArgumentsTree) []string {
	if specialCase := PHPSpecialMembersTypes(trw.wr); specialCase != "" {
		return []string{fmt.Sprintf("$success = RPC_READ%s($stream, %s);", ifString(bare, "", "_boxed"), targetName)}
	}
	unionParent := trw.PhpConstructorNeedsUnion()
	if unionParent == nil {
		if trw.PhpCanBeSimplify() {
			var result []string
			if !bare {
				result = append(result,
					"[$magic, $success] = $stream->read_uint32();",
					fmt.Sprintf("if (!$success || $magic != 0x%08[1]x) {", trw.wr.tlTag),
					"  return false;",
					"}",
				)
			}
			newArgs := trw.PHPGetFieldNatDependenciesValuesAsTypeTree(0, args)
			result = append(result, trw.Fields[0].t.trw.PhpReadMethodCall(targetName, trw.Fields[0].bare, initIfDefault, &newArgs)...)
			return result
		}
		if trw.ResultType == nil && trw.wr.PHPIsTrueType() {
			var result []string
			if !bare {
				result = append(result,
					"[$magic, $success] = $stream->read_uint32();",
					fmt.Sprintf("if (!$success || $magic != 0x%08[1]x) {", trw.wr.tlTag),
					"  return false;",
					"}",
				)
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
				result = append(result,
					"[$magic, $success] = $stream->read_uint32();",
					fmt.Sprintf("if (!$success || $magic != 0x%08[1]x) {", trw.wr.tlTag),
					"  return false;",
					"}",
				)
			}
			result = append(result, trw.Fields[0].t.trw.PhpReadMethodCall(targetName, bare, initIfDefault, args)...)
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
	result = append(result,
		fmt.Sprintf("$success = %[2]s->read%[1]s($stream%[3]s);", ifString(bare, "", "_boxed"), targetName, phpFormatArgs(args.ListAllValues())),
		"if (!$success) {",
		"  return false;",
		"}",
	)
	return result
}

func (trw *TypeRWStruct) PhpWriteMethodCall(targetName string, bare bool, args *TypeArgumentsTree) []string {
	if specialCase := PHPSpecialMembersTypes(trw.wr); specialCase != "" {
		return []string{fmt.Sprintf("$success = RPC_WRITE%s($stream, %s);", ifString(bare, "", "_boxed"), targetName)}
	}
	unionParent := trw.PhpConstructorNeedsUnion()
	if unionParent == nil {
		if trw.PhpCanBeSimplify() {
			var result []string
			if !bare {
				result = append(result,
					fmt.Sprintf("$success = $stream->write_uint32(0x%08[1]x);", trw.wr.tlTag),
					"if (!$success) {",
					"  return false;",
					"}",
				)
			}
			newArgs := trw.PHPGetFieldNatDependenciesValuesAsTypeTree(0, args)
			result = append(result, trw.Fields[0].t.trw.PhpWriteMethodCall(targetName, trw.Fields[0].bare, &newArgs)...)
			return result
		}
		if trw.ResultType == nil && trw.wr.PHPIsTrueType() {
			var result []string
			if !bare {
				result = append(result,
					fmt.Sprintf("$success = $stream->write_uint32(0x%08[1]x);", trw.wr.tlTag),
					"if (!$success) {",
					"  return false;",
					"}",
				)
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
				result = append(result,
					fmt.Sprintf("$success = $stream->write_uint32(0x%08[1]x);", trw.wr.tlTag),
					"if (!$success) {",
					"  return false;",
					"}",
				)
			}
			result = append(result, trw.Fields[0].t.trw.PhpWriteMethodCall(targetName, bare, args)...)
			return result
		}
	}
	return []string{
		fmt.Sprintf("if (is_null(%[1]s)) {", targetName),
		fmt.Sprintf("  %[1]s = %[2]s;", targetName, trw.PhpDefaultInit()),
		"}",
		fmt.Sprintf("$success = %[2]s->write%[1]s($stream%[3]s);", ifString(bare, "", "_boxed"), targetName, phpFormatArgs(args.ListAllValues())),
		"if (!$success) {",
		"  return false;",
		"}",
	}
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
