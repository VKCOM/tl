package genphp

import "github.com/VKCOM/tl/internal/puregen"

const tl1Diagonal = "tl1_diagonal"
const tl1Ref = "tl1"
const tl2Ext = "tl2ext"
const tl2Maybe = "tl2Maybe"

func GenerateTL2(opt *puregen.Options) bool {
	return opt.Kernel.TL2WhiteList != ""
}
