package tl2pure

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/TwiN/go-color"
	"github.com/vkcom/tl/internal/tlast"
	"github.com/vkcom/tl/pkg/basictl"
)

type TypeInstanceUnion struct {
	TypeInstanceCommon
	def          tlast.TL2UnionType
	variantTypes []*TypeInstanceObject
}

type KernelValueUnion struct {
	instance *TypeInstanceUnion
	index    int
	variants []KernelValueObject // we remember state of all variants to improve editing experience
}

var _ KernelValue = &KernelValueUnion{}

func (ins *TypeInstanceUnion) FindCycle(c *cycleFinder) {
	if !c.push(ins) {
		return
	}
	defer c.pop(ins)
	// any variant with a cycle is prohibited, because it could be set active
	for _, variant := range ins.variantTypes {
		variant.FindCycle(c)
	}
}

func (ins *TypeInstanceUnion) CreateValue() KernelValue {
	value := &KernelValueUnion{
		instance: ins,
		index:    0,
		variants: make([]KernelValueObject, len(ins.variantTypes)),
	}
	for i, vt := range ins.variantTypes {
		value.variants[i] = vt.CreateValueObject()
	}
	return value
}

func (ins *TypeInstanceUnion) SkipTL2(r []byte) ([]byte, error) {
	return basictl.SkipSizedValue(r)
}

func (v *KernelValueUnion) Reset() {
	v.index = 0
	v.variants[0].Reset()
}

func (v *KernelValueUnion) Random(rg *rand.Rand) {
	v.index = rg.Intn(len(v.variants))
	v.variants[v.index].Random(rg)
}

func (v *KernelValueUnion) WriteTL2(w *ByteBuilder, optimizeEmpty bool, onPath bool, level int, model *UIModel) {
	v.variants[v.index].WriteTL2(w, optimizeEmpty, onPath, level, model)
}

func (v *KernelValueUnion) ReadTL2(r []byte, ctx *TL2Context) (_ []byte, err error) {
	currentSize := 0
	if r, currentSize, err = basictl.TL2ParseSize(r); err != nil {
		return r, err
	}
	if len(r) < currentSize {
		return r, basictl.TL2Error("not enough data: expected %d, got %d", currentSize, len(r))
	}
	if currentSize == 0 {
		v.Reset()
		return r, nil
	}
	currentR := r[:currentSize]
	r = r[currentSize:]

	var block byte
	if currentR, err = basictl.ByteReadTL2(currentR, &block); err != nil {
		return currentR, err
	}
	// read No of constructor
	if block&1 != 0 {
		var index int
		if currentR, index, err = basictl.TL2ParseSize(currentR); err != nil {
			return currentR, err
		}
		if index < 0 || index >= len(v.variants) {
			return currentR, basictl.TL2Error("unexpected variant index %d, must be [0..%d)", index, len(v.variants))
		}
		v.index = index
	} else {
		v.index = 0
	}
	return r, v.variants[v.index].ReadFieldsTL2(block, currentR, ctx)
}

func (v *KernelValueUnion) WriteJSON(w []byte, ctx *TL2Context) []byte {
	defVariant := v.instance.def.Variants[v.index]
	w = append(w, `{"type":"`...)
	w = append(w, defVariant.Name...)
	if len(v.instance.variantTypes[v.index].constructorFields) == 0 {
		return append(w, `"}`...)
	}
	w = append(w, `","value":`...)
	w = v.variants[v.index].WriteJSON(w, ctx)
	w = append(w, '}')
	return w
}

func (v *KernelValueUnion) UIWrite(sb *strings.Builder, onPath bool, level int, path []int, model *UIModel) {
	defVariant := v.instance.def.Variants[v.index]
	if onPath {
		sb.WriteString(color.InBlue("{"))
	} else {
		sb.WriteString("{")
	}
	sb.WriteString(`"type":`)
	if model.CurrentEditor != nil && model.CurrentEditor.Value() == v {
		model.CurrentEditor.UIWrite(sb)
	} else {
		sb.WriteString(`"`)
		sb.WriteString(defVariant.Name)
		sb.WriteString(`"`)
	}
	if len(v.instance.variantTypes[v.index].constructorFields) == 0 {
		sb.WriteString(`}`)
		return
	}
	sb.WriteString(`,"value":`)
	v.variants[v.index].UIWrite(sb, onPath, level, path, model)
	if onPath {
		sb.WriteString(color.InBlue("}"))
	} else {
		sb.WriteString("}")
	}
}

func (v *KernelValueUnion) UIFixPath(side int, level int, model *UIModel) int {
	return v.variants[v.index].UIFixPath(side, level, model)
}

func (v *KernelValueUnion) UIStartEdit(level int, model *UIModel, fromTab bool) {
	if len(model.Path) < level {
		panic("unexpected path invariant")
	}
	if len(model.Path) == level {
		model.Path = append(model.Path[:level], -1)
	}
	selectedIndex := model.Path[level]

	if selectedIndex == -1 {
		model.EditorUnion.SetValue(v)
		model.SetCurrentEditor(&model.EditorUnion)
		return
	}
	v.variants[v.index].UIStartEdit(level, model, fromTab)
}

func (v *KernelValueUnion) UIKey(level int, model *UIModel, insert bool, delete bool, up bool, down bool) {
}

func (v *KernelValueUnion) Clone() KernelValue {
	clone := *v
	for i, va := range clone.variants {
		clone.variants[i] = va.CloneObject()
	}
	return &clone
}

func (v *KernelValueUnion) CompareForMapKey(other KernelValue) int {
	return 0
}

func (k *Kernel) createUnion(canonicalName string, def tlast.TL2UnionType,
	leftArgs []tlast.TL2TypeTemplate, actualArgs []tlast.TL2TypeArgument) (TypeInstance, error) {
	ins := &TypeInstanceUnion{
		TypeInstanceCommon: TypeInstanceCommon{
			canonicalName: canonicalName,
		},
		def:          def,
		variantTypes: make([]*TypeInstanceObject, len(def.Variants)),
	}
	for i, variantDef := range def.Variants {
		element, err := k.createObject(canonicalName+"__"+variantDef.Name,
			!variantDef.IsTypeAlias, variantDef.TypeAlias, variantDef.Fields, leftArgs, actualArgs, true, i, nil)
		if err != nil {
			return nil, fmt.Errorf("fail to resolve type of union %s element %d: %w", canonicalName, i, err)
		}
		ins.variantTypes[i] = element
	}
	return ins, nil
}
