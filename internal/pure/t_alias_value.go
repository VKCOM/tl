// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package pure

import (
	"math/rand"
	"strings"
)

type KernelValueAlias struct {
	instance *TypeInstanceAlias
	value    KernelValue
}

var _ KernelValue = &KernelValueAlias{}

func (v *KernelValueAlias) Clone() KernelValue {
	return v.value.Clone()
}

func (v *KernelValueAlias) Reset() {
	v.value.Reset()
}

func (v *KernelValueAlias) Random(rg *rand.Rand) {
	v.value.Random(rg)
}

func (v *KernelValueAlias) WriteTL2(w *ByteBuilder, optimizeEmpty bool, onPath bool, level int, model *UIModel) {
	v.value.WriteTL2(w, optimizeEmpty, onPath, level, model)
}

func (v *KernelValueAlias) ReadTL2(r []byte, ctx *TL2Context) ([]byte, error) {
	return v.value.ReadTL2(r, ctx)
}

func (v *KernelValueAlias) WriteJSON(w []byte, ctx *TL2Context) []byte {
	return v.value.WriteJSON(w, ctx)
}

func (v *KernelValueAlias) UIWrite(sb *strings.Builder, onPath bool, level int, model *UIModel) {
	v.value.UIWrite(sb, onPath, level, model)
}

func (v *KernelValueAlias) UIFixPath(side int, level int, model *UIModel) int {
	return v.value.UIFixPath(side, level, model)
}

func (v *KernelValueAlias) UIStartEdit(level int, model *UIModel, createMode int) {
	v.value.UIStartEdit(level, model, createMode)
}

func (v *KernelValueAlias) UIKey(level int, model *UIModel, insert bool, delete bool, up bool, down bool) {
	v.value.UIKey(level, model, insert, delete, up, down)
}

func (v *KernelValueAlias) CompareForMapKey(other KernelValue) int {
	if v2, ok := other.(*KernelValueAlias); ok {
		return v.value.CompareForMapKey(v2.value)
	}
	return 0
}
