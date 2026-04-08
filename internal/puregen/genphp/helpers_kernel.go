package genphp

import (
	"github.com/VKCOM/tl/internal/pure"
	"github.com/VKCOM/tl/internal/tlast"
	"github.com/VKCOM/tl/internal/utils"
)

func PHPIsDict(tp *pure.KernelType) bool {
	// TODO!
	return tp.CanonicalName().String() == "__dict_field"
}

func PHPIsArgumentNumber(arg tlast.TL2TypeArgument) bool {
	return arg.IsNumber || arg.Type.SomeType.Name.String() == "*"
}

func PHPAddDollarSign(s []string) []string {
	return utils.MapSlice(s, func(a string) string {
		return "$" + a
	})
}

func PHPAddThisSign(s []string) []string {
	return utils.MapSlice(s, func(a string) string {
		return "this->" + a
	})
}
