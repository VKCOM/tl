package genphp

import (
	"github.com/VKCOM/tl/internal/pure"
	"github.com/VKCOM/tl/internal/tlast"
)

func PHPIsDict(tp *pure.KernelType) bool {
	// TODO!
	return tp.CanonicalName().String() == "__dict_field"
}

func PHPIsArgumentNumber(arg tlast.TL2TypeArgument) bool {
	return arg.IsNumber || arg.Type.SomeType.Name.String() == "*"
}
