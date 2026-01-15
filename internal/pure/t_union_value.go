// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package pure

import (
	"math/rand"
	"strings"

	"github.com/TwiN/go-color"
	"github.com/vkcom/tl/pkg/basictl"
)

type KernelValueUnion struct {
	instance *TypeInstanceUnion
	index    int
	variants []KernelValueStruct // we remember state of all variants to improve editing experience
}

var _ KernelValue = &KernelValueUnion{}

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
	if len(v.instance.variantTypes[v.index].fields) == 0 {
		return append(w, `"}`...)
	}
	w = append(w, `","value":`...)
	w = v.variants[v.index].WriteJSON(w, ctx)
	w = append(w, '}')
	return w
}

func (v *KernelValueUnion) UIWrite(sb *strings.Builder, onPath bool, level int, model *UIModel) {
	defVariant := v.instance.def.Variants[v.index]
	if onPath {
		sb.WriteString(color.InBlue("{"))
	} else {
		sb.WriteString("{")
	}
	sb.WriteString(`"type":`)
	if model.CurrentEditor != nil && model.CurrentEditor.Value() == v {
		model.CurrentEditor.UIWrite(sb, model)
	} else {
		sb.WriteString(`"`)
		sb.WriteString(defVariant.Name)
		sb.WriteString(`"`)
	}
	if len(v.instance.variantTypes[v.index].fields) == 0 {
		sb.WriteString(`}`)
		return
	}
	sb.WriteString(`,"value":`)
	v.variants[v.index].UIWrite(sb, onPath, level, model)
	if onPath {
		sb.WriteString(color.InBlue("}"))
	} else {
		sb.WriteString("}")
	}
}

func (v *KernelValueUnion) UIFixPath(side int, level int, model *UIModel) int {
	return v.variants[v.index].UIFixPath(side, level, model)
}

func (v *KernelValueUnion) UIStartEdit(level int, model *UIModel, createMode int) {
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
	v.variants[v.index].UIStartEdit(level, model, createMode)
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
