package codecreator

import (
	"fmt"
	"strings"

	"github.com/VKCOM/tl/internal/utils"
)

type PhpCodeCreator struct {
	BasicCodeCreator
}

func NewPhpCodeCreator() *PhpCodeCreator {
	return &PhpCodeCreator{
		BasicCodeCreator: BasicCodeCreator{
			shift:              "  ",
			ifPrefixTemplate:   "if (%[1]s) {",
			ifSuffixTemplate:   "}",
			elsePrefixTemplate: "} else {",
			forPrefixTemplate:  "for (%[1]s;%[2]s;%[3]s) {",
			forSuffixTemplate:  "}",
			commentPrefix:      "// ",
			varTemplate:        "$%[1]s_%[2]s",
		},
	}
}

func (ph *PhpCodeCreator) ForIndexed(indexVar, startValue, upperBound, step string, block func()) {
	ph.forIndexed(indexVar, startValue, upperBound, step, block)
}

func (ph *PhpCodeCreator) Assign(name, value string) string {
	return fmt.Sprintf("%[1]s = %[2]s;", name, value)
}

func (ph *PhpCodeCreator) AddAssign(name, value string) string {
	return fmt.Sprintf("%[1]s += %[2]s;", name, value)
}

func (ph *PhpCodeCreator) SubAssign(name, value string) string {
	return fmt.Sprintf("%[1]s -= %[2]s;", name, value)
}

func (ph *PhpCodeCreator) OrAssign(name, value string) string {
	return fmt.Sprintf("%[1]s |= %[2]s;", name, value)
}

func (ph *PhpCodeCreator) TL2CountBytes(value string) string {
	return fmt.Sprintf("TL\\tl2_support::count_used_bytes(%[1]s)", value)
}

func (ph *PhpCodeCreator) TL2FetchSize() string {
	return "TL\\tl2_support::fetch_size()"
}

func (ph *PhpCodeCreator) TL2FetchSizeTo(name string) string {
	return ph.Assign(name, ph.TL2FetchSize())
}

func (ph *PhpCodeCreator) TL2StoreSize(value string) string {
	return fmt.Sprintf("TL\\tl2_support::store_size(%[1]s);", value)
}

func (ph *PhpCodeCreator) TL2SkipBytes(value string) string {
	return fmt.Sprintf("TL\\tl2_support::skip_bytes(%[1]s);", value)
}

func (ph *PhpCodeCreator) CheckBit(target string, bit int) string {
	return fmt.Sprintf("(%[1]s & (1 << %[2]d)) != 0", target, bit)
}

func (ph *PhpCodeCreator) Equal(left, right string) string {
	return fmt.Sprintf("%[1]s == %[2]s", left, right)
}

func (ph *PhpCodeCreator) NotEqual(left, right string) string {
	return fmt.Sprintf("%[1]s != %[2]s", left, right)
}

func (ph *PhpCodeCreator) StrongEqual(left, right string) string {
	return fmt.Sprintf("%[1]s === %[2]s", left, right)
}

func (ph *PhpCodeCreator) NotStrongEqual(left, right string) string {
	return fmt.Sprintf("%[1]s !== %[2]s", left, right)
}

func (ph *PhpCodeCreator) TLFetchUint32To(target string) string {
	return ph.Assign(target, ph.TLFetchUint32())
}

func (ph *PhpCodeCreator) TLFetchUint32() string {
	return "fetch_int() & 0xFFFFFFFF"
}

func (ph *PhpCodeCreator) TLStoreUint32(target string) string {
	return fmt.Sprintf("store_int(%[1]s);", target)
}

type FunctionArgument struct {
	Name     string
	TypeName string
}

func (ph *PhpCodeCreator) FunctionComment(args []FunctionArgument, returnType string) {
	ph.AddLines("/**")

	if len(args) == 0 {
		ph.AddLines(" * @kphp-inline")
	} else {
		for i := range args {
			ph.AddLines(fmt.Sprintf(" * @param %[1]s $%[2]s", args[i].TypeName, args[i].Name))
		}
	}
	if returnType != "" {
		ph.AddLines(" *")
		ph.AddLines(" * @return " + returnType)
	}
	ph.AddLines(" */")
}

func (ph *PhpCodeCreator) Function(modifiers []string, name string, args []FunctionArgument, returnType string, block func()) {
	ph.FunctionComment(args, returnType)
	mds_ := append(modifiers, "function", name)
	s := strings.Join(mds_, " ")
	s += "("
	s += strings.Join(utils.MapSlice(args, func(a FunctionArgument) string {
		return fmt.Sprintf("$%[1]s", a.Name)
	}), ", ")
	s += ") {"
	ph.AddLines(s)
	ph.AddBlock(block)
	ph.AddLines("}")
}
