package tl2pure

import (
	"math/rand"
	"strings"
)

// common for read/write/json/etc... for simplicity
type TL2Context struct {
}

type KernelValue interface {
	Clone() KernelValue

	Reset()
	Random(rg *rand.Rand)
	WriteTL2(w *ByteBuilder, optimizeEmpty bool, onPath bool, level int, model *UIModel)
	ReadTL2(r []byte, ctx *TL2Context) ([]byte, error)
	WriteJSON(w []byte, ctx *TL2Context) []byte

	UIWrite(sb *strings.Builder, onPath bool, level int, model *UIModel)
	UIFixPath(side int, level int, model *UIModel) int // always called onPath
	UIStartEdit(level int, model *UIModel, createMode int)
	UIKey(level int, model *UIModel, insert bool, delete bool, up bool, down bool)

	CompareForMapKey(other KernelValue) int
}
