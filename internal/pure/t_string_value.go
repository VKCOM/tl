package pure

import (
	"cmp"
	"math/rand"
	"strings"

	"github.com/TwiN/go-color"
	"github.com/vkcom/tl/pkg/basictl"
)

type KernelValueString struct {
	value string
}

var _ KernelValue = &KernelValueString{}

func (v *KernelValueString) Reset() {
	v.value = ""
}

func (v *KernelValueString) Random(rg *rand.Rand) {
	count := 0
	if (rg.Uint32() & 3) != 0 { // many strings empty
		count = 1 + rg.Intn(8)
	}
	res := make([]byte, count)
	const letters = "abcdefghijklmnopqrstuvwxyz0123456789"
	for i := range res {
		res[i] = letters[rg.Int()%len(letters)]
	}
	v.value = string(res)
}

func (v *KernelValueString) WriteTL2(w *ByteBuilder, optimizeEmpty bool, onPath bool, level int, model *UIModel) {
	if optimizeEmpty && len(v.value) == 0 {
		return
	}
	if onPath {
		w.SetCursorStart()
	}
	w.WriteString(v.value)
	if onPath {
		w.SetCursorFinish()
	}
}

func (v *KernelValueString) ReadTL2(r []byte, ctx *TL2Context) ([]byte, error) {
	return basictl.StringReadTL2(r, &v.value)
}

func (v *KernelValueString) WriteJSON(w []byte, ctx *TL2Context) []byte {
	return basictl.JSONWriteString(w, v.value)
}

func (v *KernelValueString) UIWrite(sb *strings.Builder, onPath bool, level int, model *UIModel) {
	if model.CurrentEditor != nil && model.CurrentEditor.Value() == v {
		model.CurrentEditor.UIWrite(sb, model)
	} else {
		w := string(basictl.JSONWriteString(nil, v.value))
		if onPath {
			w = color.InBlue(w)
		}
		sb.WriteString(w)
	}
}

func (v *KernelValueString) UIFixPath(side int, level int, model *UIModel) int {
	model.Path = model.Path[:level]
	return 0
}

func (v *KernelValueString) UIStartEdit(level int, model *UIModel, createMode int) {
	if len(model.Path) != level {
		panic("unexpected path invariant")
	}
	model.EditorString.SetValue(v)
	model.SetCurrentEditor(&model.EditorString)
}

func (v *KernelValueString) UIKey(level int, model *UIModel, insert bool, delete bool, up bool, down bool) {
}

func (v *KernelValueString) Clone() KernelValue {
	return &KernelValueString{v.value}
}

func (v *KernelValueString) CompareForMapKey(other KernelValue) int {
	if v2, ok := other.(*KernelValueString); ok {
		return cmp.Compare(v.value, v2.value)
	}
	return 0
}
