package tl2pure

import (
	"cmp"
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"strings"

	"github.com/TwiN/go-color"
	"github.com/vkcom/tl/internal/tlast"
)

type TypeInstancePrimitive struct {
	TypeInstanceCommon
	goodForMapKey bool
	clone         KernelValue
}

func (ins *TypeInstancePrimitive) GoodForMapKey() bool {
	return ins.goodForMapKey
}

func (ins *TypeInstancePrimitive) IsBit() bool {
	return ins.canonicalName == "bit"
}

func (ins *TypeInstancePrimitive) FindCycle(c *cycleFinder) {
}

func (ins *TypeInstancePrimitive) CreateValue() KernelValue {
	return ins.clone.Clone()
}

func (ins *TypeInstancePrimitive) SkipTL2(r []byte) ([]byte, error) {
	return ins.clone.ReadTL2(r, nil)
}

type KernelValuePrimitive interface {
	KernelValue

	SetFromEditor(str string) error
}

type KernelValueUint32 struct {
	value uint32
}

var _ KernelValuePrimitive = &KernelValueUint32{}

func (v *KernelValueUint32) Reset() {
	v.value = 0
}

func (v *KernelValueUint32) Random(rg *rand.Rand) {
	v.value = rg.Uint32()
}

func (v *KernelValueUint32) WriteTL2(w *ByteBuilder, optimizeEmpty bool, onPath bool, level int, model *UIModel) {
	if optimizeEmpty && v.value == 0 {
		return
	}
	if onPath {
		w.SetCursorStart()
	}
	w.buf = binary.LittleEndian.AppendUint32(w.buf, v.value)
	if onPath {
		w.SetCursorFinish()
	}
}

func (v *KernelValueUint32) ReadTL2(r []byte, ctx *TL2Context) ([]byte, error) {
	if len(r) < 4 {
		return r, io.ErrUnexpectedEOF
	}
	v.value = binary.LittleEndian.Uint32(r)
	return r[4:], nil
}

func (v *KernelValueUint32) WriteJSON(w []byte, ctx *TL2Context) []byte {
	return strconv.AppendUint(w, uint64(v.value), 10)
}

func (v *KernelValueUint32) UIWrite(sb *strings.Builder, onPath bool, level int, model *UIModel) {
	if model.CurrentEditor != nil && model.CurrentEditor.Value() == v {
		model.CurrentEditor.UIWrite(sb, model)
	} else {
		w := string(strconv.AppendUint(nil, uint64(v.value), 10))
		if onPath {
			w = color.InBlue(w)
		}
		sb.WriteString(w)
	}
}

func (v *KernelValueUint32) UIFixPath(side int, level int, model *UIModel) int {
	model.Path = model.Path[:level]
	return 0
}

func (v *KernelValueUint32) UIStartEdit(level int, model *UIModel, createMode int) {
	if len(model.Path) != level {
		panic("unexpected path invariant")
	}
	model.EditorPrimitive.SetValue(v)
	model.SetCurrentEditor(&model.EditorPrimitive)
}

func (v *KernelValueUint32) UIKey(level int, model *UIModel, insert bool, delete bool, up bool, down bool) {
}

func (v *KernelValueUint32) SetFromEditor(str string) error {
	value, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return err
	}
	v.value = uint32(value)
	return nil
}

func (v *KernelValueUint32) Clone() KernelValue {
	return &KernelValueUint32{value: v.value}
}

func (v *KernelValueUint32) CompareForMapKey(other KernelValue) int {
	if v2, ok := other.(*KernelValueUint32); ok {
		return cmp.Compare(v.value, v2.value)
	}
	return 0
}

type KernelValueInt32 struct {
	value int32
}

var _ KernelValuePrimitive = &KernelValueInt32{}

func (v *KernelValueInt32) Reset() {
	v.value = 0
}

func (v *KernelValueInt32) Random(rg *rand.Rand) {
	v.value = int32(rg.Uint32())
}

func (v *KernelValueInt32) WriteTL2(w *ByteBuilder, optimizeEmpty bool, onPath bool, level int, model *UIModel) {
	if optimizeEmpty && v.value == 0 {
		return
	}
	if onPath {
		w.SetCursorStart()
	}
	w.buf = binary.LittleEndian.AppendUint32(w.buf, uint32(v.value))
	if onPath {
		w.SetCursorFinish()
	}
}

func (v *KernelValueInt32) ReadTL2(r []byte, ctx *TL2Context) ([]byte, error) {
	if len(r) < 4 {
		return r, io.ErrUnexpectedEOF
	}
	v.value = int32(binary.LittleEndian.Uint32(r))
	return r[4:], nil
}

func (v *KernelValueInt32) WriteJSON(w []byte, ctx *TL2Context) []byte {
	return strconv.AppendInt(w, int64(v.value), 10)
}

func (v *KernelValueInt32) UIWrite(sb *strings.Builder, onPath bool, level int, model *UIModel) {
	if model.CurrentEditor != nil && model.CurrentEditor.Value() == v {
		model.CurrentEditor.UIWrite(sb, model)
	} else {
		w := string(strconv.AppendInt(nil, int64(v.value), 10))
		if onPath {
			w = color.InBlue(w)
		}
		sb.WriteString(w)
	}
}

func (v *KernelValueInt32) UIFixPath(side int, level int, model *UIModel) int {
	model.Path = model.Path[:level]
	return 0
}

func (v *KernelValueInt32) UIStartEdit(level int, model *UIModel, createMode int) {
	if len(model.Path) != level {
		panic("unexpected path invariant")
	}
	model.EditorPrimitive.SetValue(v)
	model.SetCurrentEditor(&model.EditorPrimitive)
}

func (v *KernelValueInt32) UIKey(level int, model *UIModel, insert bool, delete bool, up bool, down bool) {
}

func (v *KernelValueInt32) SetFromEditor(str string) error {
	value, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return err
	}
	v.value = int32(value)
	return nil
}

func (v *KernelValueInt32) Clone() KernelValue {
	return &KernelValueInt32{value: v.value}
}

func (v *KernelValueInt32) CompareForMapKey(other KernelValue) int {
	if v2, ok := other.(*KernelValueInt32); ok {
		return cmp.Compare(v.value, v2.value)
	}
	return 0
}

type KernelValueUint64 struct {
	value uint64
}

var _ KernelValuePrimitive = &KernelValueUint64{}

func (v *KernelValueUint64) Reset() {
	v.value = 0
}

func (v *KernelValueUint64) Random(rg *rand.Rand) {
	v.value = rg.Uint64()
}

func (v *KernelValueUint64) WriteTL2(w *ByteBuilder, optimizeEmpty bool, onPath bool, level int, model *UIModel) {
	if optimizeEmpty && v.value == 0 {
		return
	}
	if onPath {
		w.SetCursorStart()
	}
	w.buf = binary.LittleEndian.AppendUint64(w.buf, v.value)
	if onPath {
		w.SetCursorFinish()
	}
}

func (v *KernelValueUint64) ReadTL2(r []byte, ctx *TL2Context) ([]byte, error) {
	if len(r) < 8 {
		return r, io.ErrUnexpectedEOF
	}
	v.value = binary.LittleEndian.Uint64(r)
	return r[8:], nil
}

func (v *KernelValueUint64) WriteJSON(w []byte, ctx *TL2Context) []byte {
	return strconv.AppendUint(w, v.value, 10)
}

func (v *KernelValueUint64) UIWrite(sb *strings.Builder, onPath bool, level int, model *UIModel) {
	if model.CurrentEditor != nil && model.CurrentEditor.Value() == v {
		model.CurrentEditor.UIWrite(sb, model)
	} else {
		w := string(strconv.AppendUint(nil, v.value, 10))
		if onPath {
			w = color.InBlue(w)
		}
		sb.WriteString(w)
	}
}

func (v *KernelValueUint64) UIFixPath(side int, level int, model *UIModel) int {
	model.Path = model.Path[:level]
	return 0
}

func (v *KernelValueUint64) UIStartEdit(level int, model *UIModel, createMode int) {
	if len(model.Path) != level {
		panic("unexpected path invariant")
	}
	model.EditorPrimitive.SetValue(v)
	model.SetCurrentEditor(&model.EditorPrimitive)
}

func (v *KernelValueUint64) UIKey(level int, model *UIModel, insert bool, delete bool, up bool, down bool) {
}

func (v *KernelValueUint64) SetFromEditor(str string) error {
	value, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return err
	}
	v.value = value
	return nil
}

func (v *KernelValueUint64) Clone() KernelValue {
	return &KernelValueUint64{value: v.value}
}

func (v *KernelValueUint64) CompareForMapKey(other KernelValue) int {
	if v2, ok := other.(*KernelValueUint64); ok {
		return cmp.Compare(v.value, v2.value)
	}
	return 0
}

type KernelValueInt64 struct {
	value int64
}

var _ KernelValuePrimitive = &KernelValueInt64{}

func (v *KernelValueInt64) Reset() {
	v.value = 0
}

func (v *KernelValueInt64) Random(rg *rand.Rand) {
	v.value = int64(rg.Uint64())
}

func (v *KernelValueInt64) WriteTL2(w *ByteBuilder, optimizeEmpty bool, onPath bool, level int, model *UIModel) {
	if optimizeEmpty && v.value == 0 {
		return
	}
	if onPath {
		w.SetCursorStart()
	}
	w.buf = binary.LittleEndian.AppendUint64(w.buf, uint64(v.value))
	if onPath {
		w.SetCursorFinish()
	}
}

func (v *KernelValueInt64) ReadTL2(r []byte, ctx *TL2Context) ([]byte, error) {
	if len(r) < 8 {
		return r, io.ErrUnexpectedEOF
	}
	v.value = int64(binary.LittleEndian.Uint64(r))
	return r[8:], nil
}

func (v *KernelValueInt64) WriteJSON(w []byte, ctx *TL2Context) []byte {
	return strconv.AppendInt(w, v.value, 10)
}

func (v *KernelValueInt64) UIWrite(sb *strings.Builder, onPath bool, level int, model *UIModel) {
	if model.CurrentEditor != nil && model.CurrentEditor.Value() == v {
		model.CurrentEditor.UIWrite(sb, model)
	} else {
		w := string(strconv.AppendInt(nil, v.value, 10))
		if onPath {
			w = color.InBlue(w)
		}
		sb.WriteString(w)
	}
}

func (v *KernelValueInt64) UIFixPath(side int, level int, model *UIModel) int {
	model.Path = model.Path[:level]
	return 0
}

func (v *KernelValueInt64) UIStartEdit(level int, model *UIModel, createMode int) {
	if len(model.Path) != level {
		panic("unexpected path invariant")
	}
	model.EditorPrimitive.SetValue(v)
	model.SetCurrentEditor(&model.EditorPrimitive)
}

func (v *KernelValueInt64) UIKey(level int, model *UIModel, insert bool, delete bool, up bool, down bool) {
}

func (v *KernelValueInt64) SetFromEditor(str string) error {
	value, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return err
	}
	v.value = value
	return nil
}

func (v *KernelValueInt64) Clone() KernelValue {
	return &KernelValueInt64{value: v.value}
}

func (v *KernelValueInt64) CompareForMapKey(other KernelValue) int {
	if v2, ok := other.(*KernelValueInt64); ok {
		return cmp.Compare(v.value, v2.value)
	}
	return 0
}

type KernelValueByte struct {
	value byte
}

var _ KernelValuePrimitive = &KernelValueByte{}

func (v *KernelValueByte) Reset() {
	v.value = 0
}

func (v *KernelValueByte) Random(rg *rand.Rand) {
	v.value = byte(rg.Uint32())
}

func (v *KernelValueByte) WriteTL2(w *ByteBuilder, optimizeEmpty bool, onPath bool, level int, model *UIModel) {
	if optimizeEmpty && v.value == 0 {
		return
	}
	if onPath {
		w.SetCursorStart()
	}
	w.buf = append(w.buf, v.value)
	if onPath {
		w.SetCursorFinish()
	}
}

func (v *KernelValueByte) ReadTL2(r []byte, ctx *TL2Context) ([]byte, error) {
	if len(r) < 1 {
		return r, io.ErrUnexpectedEOF
	}
	v.value = r[0]
	return r[1:], nil
}

func (v *KernelValueByte) WriteJSON(w []byte, ctx *TL2Context) []byte {
	return strconv.AppendUint(w, uint64(v.value), 10)
}

func (v *KernelValueByte) UIWrite(sb *strings.Builder, onPath bool, level int, model *UIModel) {
	if model.CurrentEditor != nil && model.CurrentEditor.Value() == v {
		model.CurrentEditor.UIWrite(sb, model)
	} else {
		w := string(strconv.AppendUint(nil, uint64(v.value), 10))
		if onPath {
			w = color.InBlue(w)
		}
		sb.WriteString(w)
	}
}

func (v *KernelValueByte) UIFixPath(side int, level int, model *UIModel) int {
	model.Path = model.Path[:level]
	return 0
}

func (v *KernelValueByte) UIStartEdit(level int, model *UIModel, createMode int) {
	if len(model.Path) != level {
		panic("unexpected path invariant")
	}
	model.EditorPrimitive.SetValue(v)
	model.SetCurrentEditor(&model.EditorPrimitive)
}

func (v *KernelValueByte) UIKey(level int, model *UIModel, insert bool, delete bool, up bool, down bool) {
}

func (v *KernelValueByte) SetFromEditor(str string) error {
	value, err := strconv.ParseUint(str, 10, 8)
	if err != nil {
		return err
	}
	v.value = byte(value)
	return nil
}

func (v *KernelValueByte) Clone() KernelValue {
	return &KernelValueByte{value: v.value}
}

func (v *KernelValueByte) CompareForMapKey(other KernelValue) int {
	if v2, ok := other.(*KernelValueByte); ok {
		return cmp.Compare(v.value, v2.value)
	}
	return 0
}

type KernelValueBool struct {
	value bool
}

var _ KernelValue = &KernelValueBool{}

func (v *KernelValueBool) Reset() {
	v.value = false
}

func (v *KernelValueBool) Random(rg *rand.Rand) {
	v.value = (rg.Uint32() & 1) != 0
}

func (v *KernelValueBool) WriteTL2(w *ByteBuilder, optimizeEmpty bool, onPath bool, level int, model *UIModel) {
	if optimizeEmpty && !v.value {
		return
	}
	if onPath {
		w.SetCursorStart()
	}
	if v.value {
		w.buf = append(w.buf, 1)
	} else {
		w.buf = append(w.buf, 0)
	}
	if onPath {
		w.SetCursorFinish()
	}
}

func (v *KernelValueBool) ReadTL2(r []byte, ctx *TL2Context) ([]byte, error) {
	if len(r) < 1 {
		return r, io.ErrUnexpectedEOF
	}
	v.value = r[0] != 0
	return r[1:], nil
}

func (v *KernelValueBool) WriteJSON(w []byte, ctx *TL2Context) []byte {
	if v.value {
		return append(w, "true"...)
	}
	return append(w, "false"...)
}

func (v *KernelValueBool) UIWrite(sb *strings.Builder, onPath bool, level int, model *UIModel) {
	if model.CurrentEditor != nil && model.CurrentEditor.Value() == v {
		model.CurrentEditor.UIWrite(sb, model)
	} else {
		w := "false"
		if v.value {
			w = "true"
		}
		if onPath {
			w = color.InBlue(w)
		}
		sb.WriteString(w)
	}
}

func (v *KernelValueBool) UIFixPath(side int, level int, model *UIModel) int {
	model.Path = model.Path[:level]
	return 0
}

func (v *KernelValueBool) UIStartEdit(level int, model *UIModel, createMode int) {
}

func (v *KernelValueBool) UIKey(level int, model *UIModel, insert bool, delete bool, up bool, down bool) {
}

func (v *KernelValueBool) Clone() KernelValue {
	return &KernelValueBool{value: v.value}
}

func (v *KernelValueBool) CompareForMapKey(other KernelValue) int {
	if v2, ok := other.(*KernelValueBool); ok {
		if !v.value && v2.value {
			return -1
		}
		if v.value && !v2.value {
			return 1
		}
	}
	return 0
}

type KernelValueBit struct {
}

var _ KernelValue = &KernelValueBit{}

func (v *KernelValueBit) Reset() {
}

func (v *KernelValueBit) Random(rg *rand.Rand) {
}

func (v *KernelValueBit) WriteTL2(w *ByteBuilder, optimizeEmpty bool, onPath bool, level int, model *UIModel) {
}

func (v *KernelValueBit) ReadTL2(r []byte, ctx *TL2Context) ([]byte, error) {
	return r, nil
}

func (v *KernelValueBit) WriteJSON(w []byte, ctx *TL2Context) []byte {
	return append(w, "true"...)
}

func (v *KernelValueBit) UIWrite(sb *strings.Builder, onPath bool, level int, model *UIModel) {
	w := "bit"
	if onPath {
		w = color.InBlue(w)
	}
	sb.WriteString(w)
}

func (v *KernelValueBit) UIFixPath(side int, level int, model *UIModel) int {
	model.Path = model.Path[:level]
	return 0
}

func (v *KernelValueBit) UIStartEdit(level int, model *UIModel, createMode int) {
}

func (v *KernelValueBit) UIKey(level int, model *UIModel, insert bool, delete bool, up bool, down bool) {
}

func (v *KernelValueBit) Clone() KernelValue {
	return v
}

func (v *KernelValueBit) CompareForMapKey(other KernelValue) int {
	return 0
}

func (k *Kernel) addPrimitive(name string, clone KernelValue, goodForMapKey bool) {
	// for the purpose of type check, this is object with no fields, like uint32 = ;
	comb := tlast.TL2Combinator{
		TypeDecl: tlast.TL2TypeDeclaration{
			Name: tlast.TL2TypeName{Name: name},
			Type: tlast.TL2TypeDefinition{IsConstructorFields: true},
		},
	}
	ins := TypeInstancePrimitive{
		TypeInstanceCommon: TypeInstanceCommon{
			canonicalName: name,
		},
		clone:         clone,
		goodForMapKey: goodForMapKey,
	}
	ref := &TypeInstanceRef{
		ins: &ins,
	}
	kt := &KernelType{
		comb:      comb,
		instances: map[string]*TypeInstanceRef{name: ref},
	}
	if _, ok := k.instances[name]; ok {
		panic(fmt.Sprintf("error adding primitive type %s: exist in global list", name))
	}
	if err := k.addTip(kt); err != nil {
		panic(fmt.Sprintf("error adding primitive type %s: %v", name, err))
	}
	k.instances[name] = ref
	// k.instancesOrdered = append(k.instancesOrdered, ref) - we do not yet know if we need them here
}
