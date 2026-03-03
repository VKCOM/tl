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

	"github.com/vkcom/tl/internal/pure"
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

var primitiveValues = map[string]KernelValue{
	"uint32":  &KernelValueUint32{},
	"int32":   &KernelValueInt32{},
	"float32": &KernelValueInt32{},
	"uint64":  &KernelValueUint64{},
	"int64":   &KernelValueInt64{},
	"float64": &KernelValueInt64{},
	"byte":    &KernelValueByte{},
	"bool":    &KernelValueBool{},
	"bit":     &KernelValueBit{},
	"string":  &KernelValueString{},
}

func CreateValueStruct(ins *pure.TypeInstanceStruct) KernelValueStruct {
	value := KernelValueStruct{
		instance: ins,
		fields:   make([]KernelValue, len(ins.Fields())),
	}
	for i, ft := range ins.Fields() {
		if ft.FieldMask() == nil {
			value.fields[i] = CreateValue(ft.TypeInstance())
		}
	}
	return value
}

func CreateValue(ins pure.TypeInstance) KernelValue {
	switch ins := ins.(type) {
	case *pure.TypeInstanceArray:
		if ins.Field().IsBit() {
			value := &KernelValueArrayBit{
				instance: ins,
			}
			if ins.IsTuple() {
				value.resize(int(ins.Count()))
			}
			return value
		}
		value := &KernelValueArray{
			instance: ins,
		}
		if ins.IsTuple() {
			value.resize(int(ins.Count()))
		}
		return value
	case *pure.TypeInstanceDict:
		value := &KernelValueDict{
			instance: ins,
		}
		return value
	case *pure.TypeInstancePrimitive:
		return primitiveValues[ins.CanonicalName()].Clone()
	case *pure.TypeInstanceStruct:
		value := CreateValueStruct(ins)
		return &value
	case *pure.TypeInstanceUnion:
		value := &KernelValueUnion{
			instance: ins,
			index:    0,
			variants: make([]KernelValueStruct, len(ins.VariantTypes())),
		}
		for i, vt := range ins.VariantTypes() {
			value.variants[i] = CreateValueStruct(vt)
		}
		return value
	}
	panic(fmt.Errorf("type instance %s not implemented in vkext", ins.CanonicalName()))
}
