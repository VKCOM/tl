package tl2pure

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/vkcom/tl/internal/tlast"
)

type TypeInstanceAlias struct {
	TypeInstanceCommon
	fieldType *TypeInstanceRef
}

type KernelValueAlias struct {
	instance *TypeInstanceAlias
	value    KernelValue
}

var _ KernelValue = &KernelValueAlias{}

func (ins *TypeInstanceAlias) GoodForMapKey() bool {
	return ins.fieldType.ins.GoodForMapKey()
}

func (ins *TypeInstanceAlias) IsBit() bool {
	return ins.fieldType.ins.IsBit()
}

func (ins *TypeInstanceAlias) FindCycle(c *cycleFinder) {
	if !c.push(ins) {
		return
	}
	defer c.pop(ins)
	ins.fieldType.ins.FindCycle(c)
}

func (ins *TypeInstanceAlias) CreateValue() KernelValue {
	value := &KernelValueAlias{
		instance: ins,
		value:    ins.fieldType.ins.CreateValue(),
	}
	return value
}

func (ins *TypeInstanceAlias) SkipTL2(r []byte) ([]byte, error) {
	return ins.fieldType.ins.SkipTL2(r)
}

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

func (k *Kernel) createAlias(canonicalName string, alias tlast.TL2TypeRef,
	leftArgs []tlast.TL2TypeTemplate, actualArgs []tlast.TL2TypeArgument) (TypeInstance, error) {
	rt, err := k.resolveType(alias, leftArgs, actualArgs)
	if err != nil {
		return nil, fmt.Errorf("fail to resolve type of alias %s to %s: %w", canonicalName, alias, err)
	}
	aliasBit := k.IsBit(alias) // we must not call anything on TypeInstance during recursive resolution
	if aliasBit {
		return nil, fmt.Errorf("type bit is not allowed as a type alias")
	}
	fieldType, err := k.getInstance(rt)
	if err != nil {
		return nil, fmt.Errorf("fail to instantiate alias %s to %s: %w", canonicalName, alias, err)
	}
	ins := &TypeInstanceAlias{
		TypeInstanceCommon: TypeInstanceCommon{
			canonicalName: canonicalName,
		},
		fieldType: fieldType,
	}
	return ins, nil
}
