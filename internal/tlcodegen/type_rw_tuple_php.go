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
