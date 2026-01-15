package pure

import (
	"math/rand"
	"strings"

	"github.com/TwiN/go-color"
	"github.com/vkcom/tl/pkg/basictl"
)

type KernelValueArray struct {
	instance *TypeInstanceArray
	elements []KernelValue
}

var _ KernelValue = &KernelValueArray{}

func (v *KernelValueArray) resize(count int) {
	v.elements = v.elements[:min(count, cap(v.elements))]
	for len(v.elements) < count {
		v.elements = append(v.elements, v.instance.elemType.ins.CreateValue())
	}
	if len(v.elements) > count {
		v.elements = v.elements[:count]
	}
}

func (v *KernelValueArray) Clone() KernelValue {
	clone := *v // TODO - copy slice
	for i, el := range clone.elements {
		clone.elements[i] = el.Clone()
	}
	return &clone
}

func (v *KernelValueArray) Reset() {
	if !v.instance.isTuple {
		v.elements = v.elements[:0]
		return
	}
	for _, el := range v.elements {
		el.Reset()
	}
}

func (v *KernelValueArray) Random(rg *rand.Rand) {
	if !v.instance.isTuple {
		count := 0
		if (rg.Uint32() & 3) != 0 { // many vectors empty
			count = 1 + rg.Intn(4)
		}
		v.resize(count)
	}
	for _, el := range v.elements {
		el.Random(rg)
	}
}

func (v *KernelValueArray) WriteTL2(w *ByteBuilder, optimizeEmpty bool, onPath bool, level int, model *UIModel) {
	if len(v.elements) == 0 && optimizeEmpty {
		return
	}

	firstUsedByte := w.ReserveSpaceForSize()
	w.WriteElementCount(len(v.elements))

	for i, elem := range v.elements {
		fieldOnPath := onPath && len(model.Path) > level && model.Path[level] == i
		if fieldOnPath {
			elem.WriteTL2(w, false, true, level+1, model)
		} else {
			elem.WriteTL2(w, false, false, 0, model)
		}
	}

	lastUsedByte := w.Len()
	w.FinishSize(firstUsedByte, lastUsedByte, optimizeEmpty)
}

func (v *KernelValueArray) ReadTL2(r []byte, ctx *TL2Context) (_ []byte, err error) {
	currentSize := 0
	if r, currentSize, err = basictl.TL2ParseSize(r); err != nil {
		return r, err
	}
	if len(r) < currentSize {
		return r, basictl.TL2Error("not enough data: expected %d, got %d", currentSize, len(r))
	}

	currentR := r[:currentSize]
	r = r[currentSize:]

	elementCount := 0
	if currentSize != 0 {
		if currentR, elementCount, err = basictl.TL2ParseSize(currentR); err != nil {
			return r, err
		}
		if !v.instance.isTuple && elementCount > len(currentR) {
			return r, basictl.TL2ElementCountError(elementCount, currentR)
		}
	}
	if !v.instance.isTuple {
		v.resize(elementCount)
	}
	lastIndex := min(elementCount, elementCount)
	for i := 0; i < lastIndex; i++ {
		if currentR, err = v.elements[i].ReadTL2(currentR, ctx); err != nil {
			return r, err
		}
	}
	for i := lastIndex; i < len(v.elements); i++ {
		v.elements[i].Reset()
	}
	// we skip excess element all at once. not one by one
	return r, nil
}

func (v *KernelValueArray) WriteJSON(w []byte, ctx *TL2Context) []byte {
	w = append(w, '[')
	for i, el := range v.elements {
		if i != 0 {
			w = append(w, ',')
		}
		w = el.WriteJSON(w, ctx)
	}
	w = append(w, ']')
	return w
}

func (v *KernelValueArray) UIWrite(sb *strings.Builder, onPath bool, level int, model *UIModel) {
	// selectedWhole := onPath && len(path) == level
	if onPath {
		sb.WriteString(color.InBlue("["))
	} else {
		sb.WriteString("[")
	}
	for i, el := range v.elements {
		fieldOnPath := onPath && len(model.Path) > level && model.Path[level] == i
		if i != 0 {
			sb.WriteString(",")
		}
		if fieldOnPath {
			el.UIWrite(sb, true, level+1, model)
		} else {
			el.UIWrite(sb, false, 0, model)
		}
	}
	if onPath && len(model.Path) > level && model.Path[level] == len(v.elements) { // insert placeholder
		if len(v.elements) != 0 {
			sb.WriteString(",")
		}
		sb.WriteString(color.InBlue("+"))
	}
	if onPath {
		sb.WriteString(color.InBlue("]"))
	} else {
		sb.WriteString("]")
	}
}

func (v *KernelValueArray) UIFixPath(side int, level int, model *UIModel) int {
	if len(model.Path) < level {
		panic("unexpected path invariant")
	}
	maximumIndex := len(v.elements) - 1
	if !v.instance.isTuple {
		maximumIndex++
	}
	if len(model.Path) == level {
		if side >= 0 {
			model.Path = append(model.Path[:level], maximumIndex)
		} else {
			model.Path = append(model.Path[:level], 0)
		}
	} else {
		selectedIndex := model.Path[level]
		if selectedIndex > maximumIndex {
			return 1
		} else if selectedIndex < 0 {
			return -1
		}
		if selectedIndex == maximumIndex {
			model.Path = model.Path[:level+1]
			return 0
		}
		childWantsSide := v.elements[selectedIndex].UIFixPath(side, level+1, model)
		if childWantsSide == 0 {
			return 0
		}
		if childWantsSide < 0 {
			if selectedIndex <= 0 {
				return -1
			}
			model.Path = append(model.Path[:level], selectedIndex-1)
		} else {
			if selectedIndex >= maximumIndex {
				return 1
			}
			model.Path = append(model.Path[:level], selectedIndex+1)
		}
	}
	if model.Path[level] == maximumIndex {
		model.Path = model.Path[:level+1]
		return 0
	}
	childWantsSide := v.elements[model.Path[level]].UIFixPath(side, level+1, model)
	if childWantsSide != 0 {
		panic("unexpected path invariant")
	}
	return 0
}

func (v *KernelValueArray) UIStartEdit(level int, model *UIModel, createMode int) {
	if len(model.Path) < level {
		panic("unexpected path invariant")
	}
	if len(model.Path) == level {
		model.Path = append(model.Path[:level], 0)
	}
	selectedIndex := model.Path[level]
	if selectedIndex == len(v.elements) {
		if v.instance.isTuple {
			panic("unexpected path invariant for tuple")
		}
		if createMode == 0 { // require Enter or Rune to insert element
			return
		}
		v.elements = append(v.elements, v.instance.elemType.ins.CreateValue())
		if createMode == 1 {
			createMode = 0
		}
	}
	v.elements[selectedIndex].UIStartEdit(level+1, model, createMode)
}

func (v *KernelValueArray) UIKey(level int, model *UIModel, insert bool, delete bool, up bool, down bool) {
	if len(model.Path) < level+1 {
		return
	}
	selectedIndex := model.Path[level]
	if v.instance.isTuple || selectedIndex == len(v.elements) {
		return
	}
	if len(model.Path) == level+1 {
		v.elements = append(v.elements[:selectedIndex], v.elements[selectedIndex+1:]...)
		return
	}
	v.elements[selectedIndex].UIKey(level+1, model, insert, delete, up, down)
}

func (v *KernelValueArray) CompareForMapKey(other KernelValue) int {
	return 0
}
