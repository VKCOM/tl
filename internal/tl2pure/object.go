package tl2pure

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/TwiN/go-color"
	"github.com/vkcom/tl/internal/tlast"
	"github.com/vkcom/tl/pkg/basictl"
)

type TypeInstanceObject struct {
	TypeInstanceCommon
	isConstructorFields bool
	constructorFields   []tlast.TL2Field
	fieldTypes          []*TypeInstanceRef
	isUnionElement      bool
	unionIndex          int

	// if function
	resultType TypeInstance
}

type KernelValueObject struct {
	instance *TypeInstanceObject
	fields   []KernelValue // nil if optional field not set, to break recursion
}

var _ KernelValue = &KernelValueObject{}

func (ins *TypeInstanceObject) FindCycle(c *cycleFinder) {
	if !c.push(ins) {
		return
	}
	defer c.pop(ins)
	for i, fieldDef := range ins.constructorFields {
		ft := ins.fieldTypes[i]
		if fieldDef.IsOptional {
			ft.ins.FindCycle(c)
		}
	}
}

func (ins *TypeInstanceObject) CreateValue() KernelValue {
	v := ins.CreateValueObject()
	return &v
}

func (ins *TypeInstanceObject) SkipTL2(r []byte) ([]byte, error) {
	return basictl.SkipSizedValue(r)
}

func (ins *TypeInstanceObject) CreateValueObject() KernelValueObject {
	value := KernelValueObject{
		instance: ins,
		fields:   make([]KernelValue, len(ins.fieldTypes)),
	}
	for i, fieldDef := range ins.constructorFields {
		ft := ins.fieldTypes[i]
		if !fieldDef.IsOptional {
			value.fields[i] = ft.ins.CreateValue()
		}
	}
	return value
}

func (v *KernelValueObject) Reset() {
	for i, fieldDef := range v.instance.constructorFields {
		if fieldDef.IsOptional {
			v.fields[i] = nil
			continue
		}
		v.fields[i].Reset()
	}
}

func (v *KernelValueObject) Random(rg *rand.Rand) {
	for i, fieldDef := range v.instance.constructorFields {
		ft := v.instance.fieldTypes[i]
		if fieldDef.IsOptional {
			set := rg.Uint32()&1 != 0
			if !set {
				v.fields[i] = nil
				continue
			}
			if v.fields[i] == nil {
				v.fields[i] = ft.ins.CreateValue()
			}
		}
		v.fields[i].Random(rg)
	}
}

func (v *KernelValueObject) WriteTL2(w *ByteBuilder, optimizeEmpty bool, onPath bool, level int, model *UIModel) {
	firstUsedByte := w.ReserveSpaceForSize()

	lastUsedByte := firstUsedByte
	var currentBlock byte
	currentBlockPosition := w.Len()
	w.WriteFieldmask()

	if v.instance.isUnionElement && v.instance.unionIndex != 0 {
		w.WriteVariantIndex(v.instance.unionIndex)
		lastUsedByte = w.Len()
		currentBlock |= 1
	}

	for i, field := range v.fields {
		fieldOnPath := onPath && len(model.Path) > level && model.Path[level] == i
		fieldDef := v.instance.constructorFields[i]
		if (i+1)%8 == 0 {
			w.buf[currentBlockPosition] = currentBlock
			currentBlock = 0
			// start the next block
			currentBlockPosition = w.Len()
			w.WriteFieldmask()
		}
		if fieldDef.IsOmitted() {
			continue
		}
		if fieldDef.IsOptional {
			if field != nil {
				if fieldOnPath {
					field.WriteTL2(w, false, true, level+1, model)
				} else {
					field.WriteTL2(w, false, false, 0, model)
				}
				lastUsedByte = w.Len()
				currentBlock |= 1 << ((i + 1) % 8)
			}
			continue
		}
		wasLen := w.Len()
		if fieldOnPath {
			field.WriteTL2(w, true, true, level+1, model)
		} else {
			field.WriteTL2(w, true, false, 0, model)
		}
		if w.Len() != wasLen {
			lastUsedByte = w.Len()
			currentBlock |= 1 << ((i + 1) % 8)
		}
	}
	w.buf[currentBlockPosition] = currentBlock
	w.FinishSize(firstUsedByte, lastUsedByte, optimizeEmpty)
}

func (v *KernelValueObject) ReadFieldsTL2(block byte, currentR []byte, ctx *TL2Context) (err error) {
	for i, field := range v.fields {
		if (i+1)%8 == 0 {
			// start the next block
			if len(currentR) == 0 {
				for ; i < len(v.instance.fieldTypes); i++ {
					v.fields[i].Reset()
				}
				return nil
			}
			if currentR, err = basictl.ByteReadTL2(currentR, &block); err != nil {
				return err
			}
		}
		if block&(1<<((i+1)%8)) != 0 {
			// we also read omitted fields for simplicity
			if currentR, err = field.ReadTL2(currentR, ctx); err != nil {
				return err
			}
		} else {
			if v.instance.constructorFields[i].IsOptional {
				v.fields[i] = nil
			} else {
				field.Reset()
			}
		}
	}
	return nil
}

func (v *KernelValueObject) ReadTL2(r []byte, ctx *TL2Context) (_ []byte, err error) {
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
		if index != 0 {
			return currentR, basictl.TL2Error("unexpected variant index %d, expected %d", index, v.instance.unionIndex)
		}
	}
	return r, v.ReadFieldsTL2(block, currentR, ctx)
}

func (v *KernelValueObject) WriteJSON(w []byte, ctx *TL2Context) []byte {
	if !v.instance.isConstructorFields {
		return v.fields[0].WriteJSON(w, ctx)
	}
	w = append(w, '{')
	first := true
	for i, fieldDef := range v.instance.constructorFields {
		if !first {
			w = append(w, ',')
		}
		if fieldDef.IsOptional {
			if v.fields[i] == nil {
				continue
			}
		}
		first = false
		w = append(w, `"`...)
		w = append(w, fieldDef.Name...)
		w = append(w, `":`...)
		w = v.fields[i].WriteJSON(w, ctx)
	}
	w = append(w, '}')
	return w
}

func (v *KernelValueObject) UIWrite(sb *strings.Builder, onPath bool, level int, model *UIModel) {
	if onPath {
		sb.WriteString(color.InBlue("{"))
	} else {
		sb.WriteString("{")
	}
	first := true
	for i, fieldDef := range v.instance.constructorFields {
		fieldOnPath := onPath && len(model.Path) > level && model.Path[level] == i
		if fieldDef.IsOptional {
			if v.fields[i] == nil {
				if onPath && len(model.Path) > level {
					if !first {
						sb.WriteString(",")
					}
					sb.WriteString(color.InGray(fieldDef.Name))
					if fieldOnPath {
						sb.WriteString(color.InBlue("?"))
					}
					first = false
				}
				continue
			}
		}
		if !first {
			sb.WriteString(",")
		}
		first = false
		sb.WriteString(`"`)
		sb.WriteString(fieldDef.Name)
		sb.WriteString(`":`)
		if fieldDef.IsOptional {
			if v.fields[i] == nil {
				sb.WriteString("_")
				continue
			}
		}
		if fieldOnPath {
			v.fields[i].UIWrite(sb, true, level+1, model)
		} else {
			v.fields[i].UIWrite(sb, false, 0, model)
		}
	}
	if onPath {
		sb.WriteString(color.InBlue("}"))
	} else {
		sb.WriteString("}")
	}
}

func (v *KernelValueObject) UIFixPath(side int, level int, model *UIModel) int {
	if len(model.Path) < level {
		panic("unexpected path invariant")
	}
	minimalIndex := 0
	if v.instance.isUnionElement {
		minimalIndex = -1
	}
	if len(model.Path) == level {
		if side >= 0 {
			model.Path = append(model.Path[:level], len(v.fields)-1)
		} else {
			model.Path = append(model.Path[:level], minimalIndex)
		}
	} else {
		selectedIndex := model.Path[level]
		if selectedIndex >= len(v.fields) {
			return 1
		} else if selectedIndex < minimalIndex {
			return -1
		}
		if selectedIndex == -1 || v.fields[selectedIndex] == nil {
			model.Path = model.Path[:level+1]
			return 0
		}
		childWantsSide := v.fields[selectedIndex].UIFixPath(side, level+1, model)
		if childWantsSide == 0 {
			return 0
		}
		if childWantsSide < 0 {
			if selectedIndex <= minimalIndex {
				return -1
			}
			model.Path = append(model.Path[:level], selectedIndex-1)
		} else {
			if selectedIndex >= len(v.fields)-1 {
				return 1
			}
			model.Path = append(model.Path[:level], selectedIndex+1)
		}
	}
	selectedIndex := model.Path[level]
	if selectedIndex == -1 || v.fields[selectedIndex] == nil {
		model.Path = model.Path[:level+1]
		return 0
	}
	childWantsSide := v.fields[selectedIndex].UIFixPath(side, level+1, model)
	if childWantsSide != 0 {
		panic("unexpected path invariant")
	}
	return 0
}

func (v *KernelValueObject) UIStartEdit(level int, model *UIModel, createMode int) {
	if len(model.Path) < level {
		panic("unexpected path invariant")
	}
	if len(model.Path) == level {
		model.Path = append(model.Path[:level], 0)
	}
	selectedIndex := model.Path[level]
	if v.fields[selectedIndex] == nil {
		if createMode == 0 { // require Enter to insert element
			return
		}
		v.fields[selectedIndex] = v.instance.fieldTypes[selectedIndex].ins.CreateValue()
		if createMode == 1 {
			createMode = 0
		}
	}
	v.fields[selectedIndex].UIStartEdit(level+1, model, createMode)
}

func (v *KernelValueObject) UIKey(level int, model *UIModel, insert bool, delete bool, up bool, down bool) {
	if len(model.Path) < level+1 {
		return
	}
	selectedIndex := model.Path[level]
	if len(model.Path) == level+1 {
		fieldDef := v.instance.constructorFields[selectedIndex]
		if fieldDef.IsOptional && v.fields[selectedIndex] != nil {
			v.fields[selectedIndex] = nil
		}
		return
	}
	v.fields[selectedIndex].UIKey(level+1, model, insert, delete, up, down)
}

func (v *KernelValueObject) Clone() KernelValue {
	clone := v.CloneObject()
	return &clone
}

func (v *KernelValueObject) CloneObject() KernelValueObject {
	clone := *v
	for i, v := range clone.fields {
		if v != nil { // skip not set optional fields
			clone.fields[i] = v.Clone()
		}
	}
	return clone
}

func (v *KernelValueObject) CompareForMapKey(other KernelValue) int {
	return 0
}

func (k *Kernel) createObject(canonicalName string,
	isConstructorFields bool, alias tlast.TL2TypeRef, constructorFields []tlast.TL2Field,
	leftArgs []tlast.TL2TypeTemplate, actualArgs []tlast.TL2TypeArgument,
	isUnionElement bool, unionIndex int, resultType TypeInstance) (*TypeInstanceObject, error) {

	ins := &TypeInstanceObject{
		TypeInstanceCommon: TypeInstanceCommon{
			canonicalName: canonicalName,
		},
		isConstructorFields: isConstructorFields,
		constructorFields:   constructorFields,
		isUnionElement:      isUnionElement,
		unionIndex:          unionIndex,
		resultType:          resultType,
	}
	if !isConstructorFields { // if we are here, this is union variant or function result, where alias is field 1
		ins.constructorFields = append(ins.constructorFields, tlast.TL2Field{Type: alias})
	}

	for _, fieldDef := range ins.constructorFields {
		rt, err := k.resolveType(fieldDef.Type, leftArgs, actualArgs)
		if err != nil {
			return nil, fmt.Errorf("fail to resolve type of object %s field %s: %w", canonicalName, fieldDef.Name, err)
		}
		fieldIns, err := k.getInstance(rt)
		if err != nil {
			return nil, fmt.Errorf("fail to insantiate type of object %s field %s: %w", canonicalName, fieldDef.Name, err)
		}
		ins.fieldTypes = append(ins.fieldTypes, fieldIns)
	}
	return ins, nil
}
