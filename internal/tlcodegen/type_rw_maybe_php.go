package tlcodegen

import (
	"fmt"
	"strings"
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

func (trw *TypeRWMaybe) PhpReadMethodCall(targetName string, bare bool, initIfDefault bool, args *TypeArgumentsTree) []string {
	if !bare {
		result := []string{
			fmt.Sprintf(
				"[$maybeContainsValue, $success] = $stream->read_bool(0x%08[1]x, 0x%08[2]x);",
				trw.emptyTag,
				trw.okTag,
			),
			"if (!$success) {",
			"  return false;",
			"}",
			"if ($maybeContainsValue) {",
		}
		if trw.element.t == trw.getInnerTarget().t && initIfDefault {
			result = append(result,
				fmt.Sprintf("  if (is_null(%[1]s)) {", targetName),
				fmt.Sprintf("    %[1]s = %[2]s;", targetName, trw.element.t.trw.PhpDefaultInit()),
				"  }",
			)
			initIfDefault = false
		}
		bodyReader := trw.element.t.trw.PhpReadMethodCall(targetName, trw.element.bare, initIfDefault, args)
		for i, _ := range bodyReader {
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

func (trw *TypeRWMaybe) PhpWriteMethodCall(targetName string, bare bool, args *TypeArgumentsTree) []string {
	if !bare {
		result := []string{
			fmt.Sprintf(
				"$success = $stream->write_bool(!is_null(%[1]s), 0x%08[2]x, 0x%08[3]x);",
				targetName,
				trw.emptyTag,
				trw.okTag,
			),
			"if (!$success) {",
			"  return false;",
			"}",
			fmt.Sprintf("if (!is_null(%[1]s)) {", targetName),
		}
		{
			bodyWriter := trw.element.t.trw.PhpWriteMethodCall(targetName, trw.element.bare, args)
			for i, _ := range bodyWriter {
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

func (trw *TypeRWMaybe) PhpDefaultInit() string {
	return trw.element.t.trw.PhpDefaultInit()
}
