package codecreator

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunctions(t *testing.T) {
	t.Run("one argument, no result", func(t *testing.T) {
		cc := NewPhpCodeCreator()
		cc.Function([]string{"public", "static"}, "name", []FunctionArgument{{Name: "a", TypeName: "int"}}, "", func() {
			cc.Comments("empty body")
		})
		assert.Equal(t, `/**
 * @param int $a
 */
public static function name($a) {
  // empty body
}`,
			strings.Join(cc.Print(), "\n"))
	})

	t.Run("one argument, some result", func(t *testing.T) {
		cc := NewPhpCodeCreator()
		cc.Function([]string{"public", "static"}, "name", []FunctionArgument{{Name: "a", TypeName: "int"}}, "int", func() {
			cc.Comments("empty body")
		})
		assert.Equal(t, `/**
 * @param int $a
 *
 * @return int
 */
public static function name($a) {
  // empty body
}`,
			strings.Join(cc.Print(), "\n"))
	})

	t.Run("no arguments, some result", func(t *testing.T) {
		cc := NewPhpCodeCreator()
		cc.Function([]string{"public", "static"}, "name", []FunctionArgument{}, "int", func() {
			cc.Comments("empty body")
		})
		assert.Equal(t, `/**
 * @kphp-inline
 *
 * @return int
 */
public static function name() {
  // empty body
}`,
			strings.Join(cc.Print(), "\n"))
	})

	t.Run("no arguments, no result", func(t *testing.T) {
		cc := NewPhpCodeCreator()
		cc.Function([]string{"public", "static"}, "name", []FunctionArgument{}, "", func() {
			cc.Comments("empty body")
		})
		assert.Equal(t, `/**
 * @kphp-inline
 */
public static function name() {
  // empty body
}`,
			strings.Join(cc.Print(), "\n"))
	})

}
