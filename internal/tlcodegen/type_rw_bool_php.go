package tlcodegen

import (
	"fmt"
	"strings"
)

func (trw *TypeRWBool) PhpClassName(withPath bool, bare bool) string {
	return "boolean"
}

func (trw *TypeRWBool) PhpClassNameReplaced() bool {
	return true
}

func (trw *TypeRWBool) PhpTypeName(withPath bool, bare bool) string {
	return trw.PhpClassName(withPath, true)
}

func (trw *TypeRWBool) PhpGenerateCode(code *strings.Builder, bytes bool) error {
	return fmt.Errorf("boolean doesn't have php code")
}

func (trw *TypeRWBool) PhpDefaultValue() string {
	return "false"
}

func (trw *TypeRWBool) PhpIterateReachableTypes(reachableTypes *map[*TypeRWWrapper]bool) {
}
