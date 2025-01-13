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

func (trw *TypeRWBool) PhpReadMethodCall(targetName string, bare bool, args []string) []string {
	if !bare {
		return []string{
			fmt.Sprintf(
				"[%[1]s, $success] = $stream->read_bool(0x%08[2]x, 0x%08[3]x);",
				targetName,
				trw.falseTag,
				trw.trueTag,
			),
			"if (!$success) {",
			"  return false;",
			"}",
		}
	}
	return nil
}

func (trw *TypeRWBool) PhpWriteMethodCall(targetName string, bare bool, args []string) []string {
	if !bare {
		return []string{
			fmt.Sprintf(
				"$success = $stream->write_bool(%[1]s, 0x%08[2]x, 0x%08[3]x);",
				targetName,
				trw.falseTag,
				trw.trueTag,
			),
			"if (!$success) {",
			"  return false;",
			"}",
		}
	}
	return nil
}

func (trw *TypeRWBool) PhpDefaultInit() string {
	return "false"
}
