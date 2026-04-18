// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package vkext

import (
	"fmt"
	"math/rand/v2"
	"strings"

	"github.com/TwiN/go-color"
	"github.com/VKCOM/tl/internal/pure"
	"github.com/VKCOM/tl/pkg/basictl"
)

type KernelValueUnion struct {
	instance *pure.TypeInstanceUnion
	index    int
	variants []KernelValueStruct // we remember state of all variants to improve editing experience (but cannot always create them all due to possible recursion)
}

var _ KernelValue = &KernelValueUnion{}

func (v *KernelValueUnion) setIndex(index int) *KernelValueStruct {
	v.index = index
	if v.variants[v.index].instance == nil {
		v.variants[v.index] = CreateValueStruct(v.instance.VariantTypes()[v.index])
	}
	return &v.variants[v.index]
}

func (v *KernelValueUnion) Reset() {
	v.setIndex(0).Reset()
}

func (v *KernelValueUnion) Random(rg *rand.Rand) {
	v.setIndex(rg.IntN(len(v.variants))).Random(rg)
}

func (v *KernelValueUnion) ReadTL1(r []byte, ctx *TLContext, bare bool, natArgs []uint32) ([]byte, []uint32, error) {
	if bare {
		panic(fmt.Errorf("trying to read TL1 bare union %s, please report TL which caused this", v.instance.CanonicalName()))
	}
	var tag uint32
	r, err := basictl.NatRead(r, &tag)
	if err != nil {
		return r, natArgs, err
	}
	for i, variant := range v.instance.VariantTypes() {
		if tag == variant.TLTag() {
			return v.setIndex(i).ReadTL1(r, ctx, true, natArgs)
			//return v.variants[v.index].ReadTL1(r, ctx, true, natArgs)
		}
	}
	return r, natArgs, fmt.Errorf("no TL1 union variant found for tag %d in %s", tag, v.instance.CanonicalName())
}

func (v *KernelValueUnion) WriteTL1(w *ByteBuilder, bare bool, natArgs []uint32, onPath bool, level int, model *UIModel) []uint32 {
	if bare {
		panic(fmt.Errorf("trying to write TL1 bare union %s, please report TL which caused this", v.instance.CanonicalName()))
	}
	return v.variants[v.index].WriteTL1(w, false, natArgs, onPath, level, model)
}

func (v *KernelValueUnion) WriteTL2(w *ByteBuilder, optimizeEmpty bool, onPath bool, level int, model *UIModel) {
	v.variants[v.index].WriteTL2(w, optimizeEmpty, onPath, level, model)
}

func (v *KernelValueUnion) ReadTL2(r []byte, ctx *TLContext) (_ []byte, err error) {
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
	if currentR, err = basictl.ByteRead(currentR, &block); err != nil {
		return currentR, err
	}
	// read No of constructor
	var index int
	if block&1 != 0 {
		if currentR, index, err = basictl.TL2ParseSize(currentR); err != nil {
			return currentR, err
		}
		if index < 0 || index >= len(v.variants) {
			return currentR, basictl.TL2Error("unexpected variant index %d, must be [0..%d)", index, len(v.variants))
		}
	}
	return r, v.setIndex(index).ReadFieldsTL2(block, currentR, ctx)
}

func (v *KernelValueUnion) WriteJSON(w []byte, ctx *TLContext) []byte {
	w = append(w, `{"type":"`...)
	w = append(w, v.instance.VariantNames()[v.index]...)
	if len(v.instance.VariantTypes()[v.index].Fields()) == 0 {
		return append(w, `"}`...)
	}
	w = append(w, `","value":`...)
	w = v.variants[v.index].WriteJSON(w, ctx)
	w = append(w, '}')
	return w
}

func (v *KernelValueUnion) UIWrite(sb *strings.Builder, onPath bool, level int, model *UIModel) {
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
		sb.WriteString(v.instance.VariantNames()[v.index])
		sb.WriteString(`"`)
	}
	if len(v.instance.VariantTypes()[v.index].Fields()) == 0 {
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
		if va.instance != nil {
			clone.variants[i] = va.CloneObject()
		}
	}
	return &clone
}

func (v *KernelValueUnion) CompareForMapKey(other KernelValue) int {
	return 0
}
