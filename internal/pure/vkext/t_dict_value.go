// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package vkext

import (
	"fmt"
	"math/rand/v2"
	"slices"
	"strings"

	"github.com/VKCOM/tl/internal/pure"
	"github.com/VKCOM/tl/pkg/basictl"
)

type KernelValueDict struct {
	instance *pure.TypeInstanceDict
	elements []KernelValueStruct // cap contains created elements
}

var _ KernelValue = &KernelValueDict{}

func (v *KernelValueDict) resize(count int) {
	// v.elements = v.elements[:min(count, cap(v.elements))]
	for len(v.elements) < count {
		v.elements = append(v.elements, CreateValueStruct(v.instance.FieldType()))
	}
	if len(v.elements) > count {
		v.elements = v.elements[:count]
	}
}

func (v *KernelValueDict) sort() {
	slices.SortFunc(v.elements, func(a KernelValueStruct, b KernelValueStruct) int {
		return a.fields[0].CompareForMapKey(b.fields[0])
	})
	v.elements = slices.CompactFunc(v.elements, func(a KernelValueStruct, b KernelValueStruct) bool {
		return a.fields[0].CompareForMapKey(b.fields[0]) == 0
	})
}

func (v *KernelValueDict) Reset() {
	v.elements = v.elements[:0]
}

func (v *KernelValueDict) Random(rg *rand.Rand) {
	count := 0
	if (rg.Uint32() & 3) != 0 { // many vectors empty
		count = 1 + rg.IntN(4)
	}
	v.resize(count)
	for _, el := range v.elements {
		el.Random(rg)
	}
	v.sort()
}

func (v *KernelValueDict) ReadTL1(r []byte, ctx *TLContext, bare bool, natArgs []uint32) ([]byte, []uint32, error) {
	if !bare {
		panic(fmt.Errorf("trying to read TL1 boxed dict %s, please report TL which caused this", v.instance.CanonicalName()))
	}

	natArgsFinish := len(natArgs)
	myNatArgs := natArgs[natArgsFinish-len(v.instance.NatParams()):]

	var err error
	var count uint32
	if r, err = basictl.NatRead(r, &count); err != nil {
		return r, natArgs, err
	}
	v.resize(int(count))

	for i, elem := range v.elements {
		natArgs = formatNatArgs(natArgs[:natArgsFinish], myNatArgs, v.instance.Field().NatArgs())
		if r, natArgs, err = elem.ReadTL1(r, ctx, v.instance.Field().Bare(), natArgs); err != nil {
			// leave container in good shape
			v.resize(i)
			v.sort()
			return r, natArgs, err
		}
	}
	v.sort()
	return r, natArgs, nil
}

func (v *KernelValueDict) WriteTL1(w *ByteBuilder, bare bool, natArgs []uint32, onPath bool, level int, model *UIModel) []uint32 {
	if !bare {
		panic(fmt.Errorf("trying to write TL1 boxed dict %s, please report TL which caused this", v.instance.CanonicalName()))
	}
	v.sort() // TODO - remove, sort when container changes

	natArgsFinish := len(natArgs)
	myNatArgs := natArgs[natArgsFinish-len(v.instance.NatParams()):]

	w.WriteElementCountTL1(uint32(len(v.elements)))

	for _, elem := range v.elements {
		// TODO - onPath
		natArgs = formatNatArgs(natArgs[:natArgsFinish], myNatArgs, v.instance.Field().NatArgs())
		natArgs = elem.WriteTL1(w, v.instance.Field().Bare(), natArgs, false, 0, model)
	}
	return natArgs
}

func (v *KernelValueDict) WriteTL2(w *ByteBuilder, optimizeEmpty bool, onPath bool, level int, model *UIModel) {
	if len(v.elements) == 0 && optimizeEmpty {
		return
	}
	v.sort() // TODO - remove, sort when container changes

	firstUsedByte := w.ReserveSpaceForSize()
	w.WriteElementCountTL2(len(v.elements))

	for _, elem := range v.elements {
		// TODO - onPath
		elem.WriteTL2(w, false, false, 0, model)
	}

	lastUsedByte := w.Len()
	w.FinishSize(firstUsedByte, lastUsedByte, optimizeEmpty)
}

func (v *KernelValueDict) ReadTL2(r []byte, ctx *TLContext) (_ []byte, err error) {
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
		if elementCount > len(currentR) {
			return r, basictl.TL2ElementCountError(elementCount, currentR)
		}
	}

	v.resize(elementCount)
	for i, elem := range v.elements {
		if currentR, err = elem.ReadTL2(currentR, ctx); err != nil {
			// leave container in good shape
			v.resize(i)
			v.sort()
			return r, err
		}
	}
	v.sort()
	return r, nil
}

func (v *KernelValueDict) WriteJSON(w []byte, ctx *TLContext) []byte {
	v.sort() // TODO - remove, sort when container changes
	w = append(w, '[')
	first := true
	for _, el := range v.elements {
		if !first {
			w = append(w, ',')
		}
		first = false
		w = el.WriteJSON(w, ctx)
	}
	w = append(w, ']')
	return w
}

func (v *KernelValueDict) UIWrite(sb *strings.Builder, onPath bool, level int, model *UIModel) {
	sb.WriteString("<KernelValueDict>") // TODO
}

func (v *KernelValueDict) UIFixPath(side int, level int, model *UIModel) int {
	return 0
}

func (v *KernelValueDict) UIStartEdit(level int, model *UIModel, createMode int) {
}

func (v *KernelValueDict) UIKey(level int, model *UIModel, insert bool, delete bool, up bool, down bool) {
}

func (v *KernelValueDict) Clone() KernelValue {
	clone := *v
	for i, el := range clone.elements {
		clone.elements[i] = el.CloneObject()
	}
	return &clone
}

func (v *KernelValueDict) CompareForMapKey(other KernelValue) int {
	return 0
}
