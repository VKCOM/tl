package genphp

import (
	"strings"

	"github.com/VKCOM/tl/internal/pure"
	"github.com/VKCOM/tl/internal/utils"
)

func PHPIsDict(tp *pure.KernelType) bool {
	// TODO!
	return tp.CanonicalName().String() == "__dict_field"
}

func PHPAddDollarSign(s []string) []string {
	return utils.MapSlice(s, func(a string) string {
		return "$" + a
	})
}

// TODO!
func PHPLegacyGoNameToCompare(canonicalName string) string {
	canonicalName = strings.Replace(canonicalName, "<", "", -1)
	canonicalName = strings.Replace(canonicalName, ">", "", -1)
	canonicalName = strings.Replace(canonicalName, "*", "", -1)
	canonicalName = strings.Replace(canonicalName, ".", "", -1)
	return canonicalName
}

func PHPRPCPrimitive(originalName string) (bool, string) {
	switch originalName {
	case "_":
		return true, "rpcResponseOk"
	case "reqError":
		return true, "rpcResponseError"
	case "reqResultHeader":
		return true, "rpcResponseHeader"
	case "ReqResult":
		return true, "TL\\RpcResponse"
	}
	return false, ""
}
