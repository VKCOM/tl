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

func (trw *TypeRWMaybe) PhpReadMethodCall(targetName string, bare bool, args []string) []string {
	if !bare {
		result := []string{
			fmt.Sprintf(
				"[$maybeContainsValue, $success] = $stream->read_bool(0x%08[1]x, 0x%08[2]x)",
				trw.emptyTag,
				trw.okTag,
			),
			"if (!$success) {",
			"  return false;",
			"}",
			"if ($maybeContainsValue) {",
		}
		if trw.element.t == trw.getInnerTarget().t {
			result = append(result,
				fmt.Sprintf("  if (%[1]s == null) {", targetName),
				fmt.Sprintf("    %[1]s = %[2]s;", targetName, trw.element.t.trw.PhpDefaultInit()),
				"  }",
			)
		}
		bodyReader := trw.element.t.trw.PhpReadMethodCall(targetName, trw.element.bare, args)
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

func (trw *TypeRWMaybe) PhpDefaultInit() string {
	return trw.element.t.trw.PhpDefaultInit()
}
