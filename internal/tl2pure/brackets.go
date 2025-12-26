package tl2pure

import (
	"math/rand"
	"strings"

	"github.com/TwiN/go-color"
	"github.com/vkcom/tl/pkg/basictl"
)

type TypeInstanceTupleVector struct {
	TypeInstanceCommon
	isTuple   bool
	count     int
	fieldType *TypeInstanceRef // TODO rename to elemType
}

type KernelValueTuple struct {
	instance *TypeInstanceTupleVector
	elements []KernelValue
}

func (ins *TypeInstanceTupleVector) FindCycle(c *cycleFinder) {
	if !c.push(ins) {
		return
	}
	defer c.pop(ins)
	if ins.isTuple {
		ins.fieldType.ins.FindCycle(c)
	}
}

func (ins *TypeInstanceTupleVector) CreateValue() KernelValue {
	value := &KernelValueTuple{
		instance: ins,
	}
	if ins.isTuple {
		value.resize(ins.count)
	}
	return value
}

func (ins *TypeInstanceTupleVector) SkipTL2(r []byte) ([]byte, error) {
	return basictl.SkipSizedValue(r)
}

func (v *KernelValueTuple) resize(count int) {
	v.elements = v.elements[:min(count, cap(v.elements))]
	for len(v.elements) < count {
		v.elements = append(v.elements, v.instance.fieldType.ins.CreateValue())
	}
	if len(v.elements) > count {
		v.elements = v.elements[:count]
	}
}

func (v *KernelValueTuple) Clone() KernelValue {
	clone := *v // TODO - copy slice
	for i, el := range clone.elements {
		clone.elements[i] = el.Clone()
	}
	return &clone
}

func (v *KernelValueTuple) Reset() {
	if !v.instance.isTuple {
		v.elements = v.elements[:0]
		return
	}
	for _, el := range v.elements {
		el.Reset()
	}
}

func (v *KernelValueTuple) Random(rg *rand.Rand) {
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

func (v *KernelValueTuple) WriteTL2(w []byte, optimizeEmpty bool, ctx *TL2Context) []byte {
	if len(v.elements) == 0 && optimizeEmpty {
		return w
	}

	oldLen := len(w)
	w = append(w, make([]byte, 16)...) // reserve space for

	firstUsedByte := len(w)

	w = basictl.TL2WriteSize(w, len(v.elements))

	for _, elem := range v.elements {
		w = elem.WriteTL2(w, false, ctx)
	}

	lastUsedByte := len(w)
	offset := basictl.TL2PutSize(w[oldLen:], lastUsedByte-firstUsedByte)
	offset += copy(w[oldLen+offset:], w[firstUsedByte:lastUsedByte])
	return w[:oldLen+offset]
}

func (v *KernelValueTuple) ReadTL2(r []byte, ctx *TL2Context) (_ []byte, err error) {
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

func (v *KernelValueTuple) WriteJSON(w []byte, ctx *TL2Context) []byte {
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

func (v *KernelValueTuple) UIWrite(sb *strings.Builder, onPath bool, level int, path []int, model *UIModel) {
	// selectedWhole := onPath && len(path) == level
	if onPath && len(path) > level && path[level] == 0 {
		sb.WriteString(color.InBlue("["))
	} else {
		sb.WriteString("[")
	}
	for i, el := range v.elements {
		fieldOnPath := onPath && len(path) > level && path[level] == i
		if i != 0 {
			if fieldOnPath {
				sb.WriteString(color.InBlue(","))
			} else {
				sb.WriteString(",")
			}
		}
		if fieldOnPath {
			el.UIWrite(sb, true, level+1, path, model)
			continue
		}
		el.UIWrite(sb, false, 0, nil, model)
	}
	if onPath && len(path) > level && path[level] == len(v.elements) { // insert placeholder
		if len(v.elements) != 0 {
			sb.WriteString(",")
		}
		sb.WriteString("_")
	}
	sb.WriteString("]")
}

func (v *KernelValueTuple) UIFixPath(side int, level int, model *UIModel) int {
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

func (v *KernelValueTuple) UIStartEdit(level int, model *UIModel, fromTab bool) {
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
		if fromTab { // require Enter to insert element
			return
		}
		fromTab = true // do not recursively create first field
		v.elements = append(v.elements, v.instance.fieldType.ins.CreateValue())
	}
	v.elements[selectedIndex].UIStartEdit(level+1, model, fromTab)
}

func (v *KernelValueTuple) CompareForMapKey(other KernelValue) int {
	return 0
}

func (k *Kernel) createTupleVector(canonicalName string, isTuple bool, count uint32, fieldType *TypeInstanceRef) TypeInstance {
	if fieldType.ins.IsBit() {
		ins := &TypeInstanceTupleVectorBit{
			TypeInstanceCommon: TypeInstanceCommon{
				canonicalName: canonicalName,
			},
			isTuple: isTuple,
			count:   int(count),
		}
		return ins
	}
	ins := &TypeInstanceTupleVector{
		TypeInstanceCommon: TypeInstanceCommon{
			canonicalName: canonicalName,
		},
		isTuple:   isTuple,
		count:     int(count),
		fieldType: fieldType,
	}
	return ins
}
