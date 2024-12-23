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

}
`,
		targetType.trw.PhpClassName(false, false),
		strings.Join(constructors, ",\n    "),
	))
	return nil
}

func (trw *TypeRWUnion) PhpDefaultValue() string {
	return "null"
}

func (trw *TypeRWUnion) PhpIterateReachableTypes(reachableTypes *map[*TypeRWWrapper]bool) {
	for _, field := range trw.Fields {
		field.t.PhpIterateReachableTypes(reachableTypes)
	}
}
