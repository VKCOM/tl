package tl2pure

import (
	"cmp"
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"strconv"

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

func (ins *TypeInstancePrimitive) FindCycle(c *cycleFinder) {
}

func (ins *TypeInstancePrimitive) CreateValue() KernelValue {
	return ins.clone.Clone()
}

type KernelValueUint32 struct {
	value uint32
}

func (v *KernelValueUint32) Random(rg *rand.Rand) {
	v.value = rg.Uint32()
}

func (v *KernelValueUint32) WriteTL2(w []byte) []byte {
	return binary.LittleEndian.AppendUint32(w, v.value)
}

func (v *KernelValueUint32) ReadTL2(w []byte) ([]byte, error) {
	if len(w) < 4 {
		return w, io.ErrUnexpectedEOF
	}
	v.value = binary.LittleEndian.Uint32(w)
	return w[4:], nil
}

func (v *KernelValueUint32) WriteJSON(w []byte) []byte {
	return strconv.AppendUint(w, uint64(v.value), 10)
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

func (v *KernelValueInt32) Random(rg *rand.Rand) {
	v.value = int32(rg.Uint32())
}

func (v *KernelValueInt32) WriteTL2(w []byte) []byte {
	return binary.LittleEndian.AppendUint32(w, uint32(v.value))
}

func (v *KernelValueInt32) ReadTL2(w []byte) ([]byte, error) {
	if len(w) < 4 {
		return w, io.ErrUnexpectedEOF
	}
	v.value = int32(binary.LittleEndian.Uint32(w))
	return w[4:], nil
}

func (v *KernelValueInt32) WriteJSON(w []byte) []byte {
	return strconv.AppendInt(w, int64(v.value), 10)
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

func (v *KernelValueUint64) Random(rg *rand.Rand) {
	v.value = rg.Uint64()
}

func (v *KernelValueUint64) WriteTL2(w []byte) []byte {
	return binary.LittleEndian.AppendUint64(w, v.value)
}

func (v *KernelValueUint64) ReadTL2(w []byte) ([]byte, error) {
	if len(w) < 8 {
		return w, io.ErrUnexpectedEOF
	}
	v.value = binary.LittleEndian.Uint64(w)
	return w[8:], nil
}

func (v *KernelValueUint64) WriteJSON(w []byte) []byte {
	return strconv.AppendUint(w, v.value, 10)
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

func (v *KernelValueInt64) Random(rg *rand.Rand) {
	v.value = int64(rg.Uint64())
}

func (v *KernelValueInt64) WriteTL2(w []byte) []byte {
	return binary.LittleEndian.AppendUint64(w, uint64(v.value))
}

func (v *KernelValueInt64) ReadTL2(w []byte) ([]byte, error) {
	if len(w) < 8 {
		return w, io.ErrUnexpectedEOF
	}
	v.value = int64(binary.LittleEndian.Uint64(w))
	return w[8:], nil
}

func (v *KernelValueInt64) WriteJSON(w []byte) []byte {
	return strconv.AppendInt(w, v.value, 10)
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

func (v *KernelValueByte) Random(rg *rand.Rand) {
	v.value = byte(rg.Uint32())
}

func (v *KernelValueByte) WriteTL2(w []byte) []byte {
	return append(w, v.value)
}

func (v *KernelValueByte) ReadTL2(w []byte) ([]byte, error) {
	if len(w) < 1 {
		return w, io.ErrUnexpectedEOF
	}
	v.value = w[0]
	return w[1:], nil
}

func (v *KernelValueByte) WriteJSON(w []byte) []byte {
	return strconv.AppendUint(w, uint64(v.value), 10)

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

func (v *KernelValueBool) Random(rg *rand.Rand) {
	v.value = (rg.Uint32() & 1) != 0
}

func (v *KernelValueBool) WriteTL2(w []byte) []byte {
	if v.value {
		return append(w, 1)
	}
	return append(w, 0)
}

func (v *KernelValueBool) ReadTL2(w []byte) ([]byte, error) {
	if len(w) < 1 {
		return w, io.ErrUnexpectedEOF
	}
	v.value = w[0] != 0
	return w[1:], nil
}

func (v *KernelValueBool) WriteJSON(w []byte) []byte {
	if v.value {
		return append(w, "true"...)
	}
	return append(w, "false"...)
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

func (v *KernelValueBit) Random(rg *rand.Rand) {
}

func (v *KernelValueBit) WriteTL2(w []byte) []byte {
	return w
}

func (v *KernelValueBit) ReadTL2(w []byte) ([]byte, error) {
	return w, nil
}

func (v *KernelValueBit) WriteJSON(w []byte) []byte {
	return append(w, "true"...)
}

func (v *KernelValueBit) Clone() KernelValue {
	return v
}

func (v *KernelValueBit) CompareForMapKey(other KernelValue) int {
	return 0
}

func (k *Kernel) addPrimitive(name string, clone KernelValue, goodForMapKey bool) {
	decl := tlast.TL2TypeDeclaration{
		Name: tlast.TL2TypeName{Name: name},
		Type: tlast.TL2TypeDefinition{IsConstructorFields: true}, // for the purpose of type check, this is object with no fields
	}
	ins := TypeInstancePrimitive{
		TypeInstanceCommon: TypeInstanceCommon{
			canonicalName: name,
			declaration:   decl,
		},
		clone:         clone,
		goodForMapKey: goodForMapKey,
	}
	ref := &TypeInstanceRef{
		ins: &ins,
	}
	kt := &KernelType{
		tip:       decl,
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
