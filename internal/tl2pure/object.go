package tl2pure

import (
	"fmt"
	"math/rand"

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
}

type KernelValueObject struct {
	instance *TypeInstanceObject
	fields   []KernelValue // nil if optional field not set, to break recursion
}

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

func (v *KernelValueObject) WriteTL2(w []byte, optimizeEmpty bool, ctx *TL2Context) []byte {
	oldLen := len(w)
	w = append(w, make([]byte, 16)...) // reserve space for

	firstUsedByte := len(w)
	lastUsedByte := firstUsedByte
	var currentBlock byte
	currentBlockPosition := len(w)
	w = append(w, 0)

	if v.instance.isUnionElement {
		w = basictl.TL2WriteSize(w, v.instance.unionIndex)
		lastUsedByte = len(w)
		currentBlock |= 1
	}

	for i, field := range v.fields {
		fieldDef := v.instance.constructorFields[i]
		if (i+1)%8 == 0 {
			w[currentBlockPosition] = currentBlock
			currentBlock = 0
			// start the next block
			currentBlockPosition = len(w)
			w = append(w, 0)
		}
		if fieldDef.IsOmitted() {
			continue
		}
		if fieldDef.IsOptional {
			if field != nil {
				w = append(w, 1)
				w = field.WriteTL2(w, false, ctx)
				lastUsedByte = len(w)
				currentBlock |= 1 << ((i + 1) % 8)
			}
			continue
		}
		wasLen := len(w)
		w = field.WriteTL2(w, true, ctx)
		if len(w) != wasLen {
			lastUsedByte = len(w)
			currentBlock |= 1 << ((i + 1) % 8)
		}
	}
	w[currentBlockPosition] = currentBlock
	if optimizeEmpty && firstUsedByte == lastUsedByte {
		return w[:oldLen]
	}
	offset := basictl.TL2PutSize(w[oldLen:], lastUsedByte-firstUsedByte)
	offset += copy(w[oldLen+offset:], w[firstUsedByte:lastUsedByte])
	return w[:oldLen+offset]
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
		if currentR, err = basictl.TL2ReadSize(currentR, &index); err != nil {
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
		if fieldDef.IsOptional {
			if v.fields[i] == nil {
				continue
			}
		}
		if !first {
			w = append(w, ',')
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

func (k *Kernel) createObject(canonicalName string, declaration tlast.TL2TypeDeclaration,
	isConstructorFields bool, alias tlast.TL2TypeRef, constructorFields []tlast.TL2Field,
	templateArguments []tlast.TL2TypeTemplate, lrc []tlast.TL2TypeArgument,
	isUnionElement bool, unionIndex int) (*TypeInstanceObject, error) {

	ins := &TypeInstanceObject{
		TypeInstanceCommon: TypeInstanceCommon{
			canonicalName: canonicalName,
			declaration:   declaration,
		},
		isConstructorFields: isConstructorFields,
		constructorFields:   constructorFields,
		isUnionElement:      isUnionElement,
		unionIndex:          unionIndex,
	}
	if !isConstructorFields { // if we are here, this is union variant or function result, where alias is field 1
		ins.constructorFields = append(ins.constructorFields, tlast.TL2Field{Type: alias})
	}

	for _, fieldDef := range ins.constructorFields {
		rt, err := k.resolveType(fieldDef.Type, templateArguments, lrc)
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
