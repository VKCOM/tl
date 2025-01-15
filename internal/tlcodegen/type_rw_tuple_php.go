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

func (trw *TypeRWBrackets) PhpReadMethodCall(targetName string, bare bool, args []string) []string {
	result := make([]string, 0)
	switch {
	// actual vector
	case trw.vectorLike && !trw.dictLike:
		elementName := fmt.Sprintf("$%s___element", trw.PhpClassName(false, true))
		elementRead := trw.element.t.trw.PhpReadMethodCall(elementName, trw.element.bare, args)
		for i := range elementRead {
			elementRead[i] = "  " + elementRead[i]
		}
		result = append(result,
			"[$vector_size, $success] = $stream->read_uint32();",
			"if (!$success) {",
			"  return false;",
			"}",
			// TODO MAKE MORE EFFICIENT
			fmt.Sprintf("%[1]s = [];", targetName),
			"for($i = 0; $i < $vector_size; $i++) {",
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
		elementName := fmt.Sprintf("$%s___element", trw.PhpClassName(false, true))
		tupleSize := args[len(args)-1]
		elementArgs := args[:len(args)-1]
		elementRead := trw.element.t.trw.PhpReadMethodCall(elementName, trw.element.bare, elementArgs)
		for i := range elementRead {
			elementRead[i] = "  " + elementRead[i]
		}
		result = append(result,
			// TODO MAKE MORE EFFICIENT
			fmt.Sprintf("%[1]s = [];", targetName),
			fmt.Sprintf("for($i = 0; $i < %[1]s; $i++) {", tupleSize),
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
		keyRead := trw.dictKeyField.t.trw.PhpReadMethodCall(keyElement, trw.dictKeyField.bare, args)
		for i := range keyRead {
			keyRead[i] = "  " + keyRead[i]
		}
		valueRead := trw.dictValueField.t.trw.PhpReadMethodCall(valueElement, trw.dictValueField.bare, args)
		for i := range valueRead {
			valueRead[i] = "  " + valueRead[i]
		}

		result = append(result,
			"[$dict_size, $success] = $stream->read_uint32();",
			"if (!$success) {",
			"  return false;",
			"}",
			// TODO MAKE MORE EFFICIENT
			fmt.Sprintf("%[1]s = [];", targetName),
			"for($i = 0; $i < $dict_size; $i++) {",
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

func (trw *TypeRWBrackets) PhpWriteMethodCall(targetName string, bare bool, args []string) []string {
	result := make([]string, 0)
	result = append(result,
		fmt.Sprintf("$success = $stream->write_uint32(count(%[1]s));", targetName),
		"if (!$success) {",
		"  return false;",
		"}",
	)
	switch {
	// actual vector
	case trw.vectorLike && !trw.dictLike:
		result = append(result,
			fmt.Sprintf("$vector_size = count(%[1]s);", targetName),
			// TODO MAKE MORE EFFICIENT
			"for($i = 0; $i < $vector_size; $i++) {",
		)
		{
			elementRead := trw.element.t.trw.PhpWriteMethodCall(fmt.Sprintf("%[1]s[$i]", targetName), trw.element.bare, args)
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
		tupleSize := args[len(args)-1]
		elementArgs := args[:len(args)-1]
		result = append(result,
			fmt.Sprintf("$tuple_size = %[1]s;", tupleSize),
			// TODO MAKE MORE EFFICIENT
			"for($i = 0; $i < $tuple_size; $i++) {",
		)
		{
			elementRead := trw.element.t.trw.PhpWriteMethodCall(fmt.Sprintf("%[1]s[$i]", targetName), trw.element.bare, elementArgs)
			for i := range elementRead {
				elementRead[i] = "  " + elementRead[i]
			}
			result = append(result, elementRead...)
		}
		result = append(result, "}")
		return result
	// actual map / dictionary
	case trw.dictLike:
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
