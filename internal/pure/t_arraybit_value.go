// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package pure

import (
	"math/rand"
	"strings"

	"github.com/vkcom/tl/pkg/basictl"
)

type KernelValueArrayBit struct {
	instance *TypeInstanceArrayBit
	elements []bool
}

var _ KernelValue = &KernelValueArrayBit{}

func (v *KernelValueArrayBit) resize(count int) {
	v.elements = v.elements[:min(count, cap(v.elements))]
	if len(v.elements) < count {
		v.elements = append(v.elements, make([]bool, count-len(v.elements))...)
	}
	if len(v.elements) > count {
		v.elements = v.elements[:count]
	}
}

func (v *KernelValueArrayBit) Reset() {
	if !v.instance.isTuple {
		v.elements = v.elements[:0]
		return
	}
	clear(v.elements)
}

func (v *KernelValueArrayBit) Random(rg *rand.Rand) {
	if !v.instance.isTuple {
		count := 0
		if (rg.Uint32() & 3) != 0 { // many vectors empty
			count = 1 + rg.Intn(4)
		}
		v.resize(count)
	}
	for i := range v.elements {
		v.elements[i] = rg.Uint32()&1 != 0
	}
}

func (v *KernelValueArrayBit) WriteTL2(w *ByteBuilder, optimizeEmpty bool, onPath bool, level int, model *UIModel) {
	if len(v.elements) == 0 && optimizeEmpty {
		return
	}

	firstUsedByte := w.ReserveSpaceForSize()

	w.WriteElementCount(len(v.elements))

	w.buf = basictl.VectorBoolContentWriteTL2(w.buf, v.elements)

	lastUsedByte := w.Len()
	w.FinishSize(firstUsedByte, lastUsedByte, optimizeEmpty)
}

func (v *KernelValueArrayBit) ReadTL2(r []byte, ctx *TL2Context) (_ []byte, err error) {
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
		if !v.instance.isTuple && elementCount/8 > len(currentR) { // this is relaxed check, +7 could overflow
			return r, basictl.TL2ElementCountError(elementCount, currentR)
		}
	}
	if !v.instance.isTuple {
		v.resize(elementCount)
	}
	lastIndex := min(elementCount, elementCount)
	if _, err = basictl.VectorBoolContentReadTL2(currentR, v.elements[:lastIndex]); err != nil {
		return r, err
	}
	clear(v.elements[lastIndex:])
	// we skip excess element all at once. not one by one
	return r, nil
}

func (v *KernelValueArrayBit) WriteJSON(w []byte, ctx *TL2Context) []byte {
	w = append(w, '[')
	first := true
	for _, el := range v.elements {
		if !first {
			w = append(w, ',')
		}
		first = false
		value := KernelValueBool{el}
		w = value.WriteJSON(w, ctx)
	}
	w = append(w, ']')
	return w
}

func (v *KernelValueArrayBit) UIWrite(sb *strings.Builder, onPath bool, level int, model *UIModel) {
	sb.WriteString("<KernelValueArrayBit>")
}

func (v *KernelValueArrayBit) UIFixPath(side int, level int, model *UIModel) int {
	model.Path = model.Path[:level]
	return 0
}

func (v *KernelValueArrayBit) UIStartEdit(level int, model *UIModel, createMode int) {
}

func (v *KernelValueArrayBit) UIKey(level int, model *UIModel, insert bool, delete bool, up bool, down bool) {
}

func (v *KernelValueArrayBit) Clone() KernelValue {
	return &KernelValueArrayBit{
		instance: v.instance,
		elements: append([]bool{}, v.elements...),
	}
}

func (v *KernelValueArrayBit) CompareForMapKey(other KernelValue) int {
	return 0
}
