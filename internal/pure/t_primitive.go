// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package pure

import (
	"fmt"
	"io"

	"github.com/vkcom/tl/internal/tlast"
	"github.com/vkcom/tl/internal/utils"
	"github.com/vkcom/tl/pkg/basictl"
)

type TypeInstancePrimitive struct {
	TypeInstanceCommon
	goodForMapKey bool
	isString      bool
	fixedSize     int // bytes, 0 for bit and string
}

func (ins *TypeInstancePrimitive) GoodForMapKey() bool {
	return ins.goodForMapKey
}

func (ins *TypeInstancePrimitive) IsBit() bool {
	return ins.canonicalName == "bit"
}

func (ins *TypeInstancePrimitive) FindCycle(c *cycleFinder, prName tlast.PositionRange) {
}

func (ins *TypeInstancePrimitive) GetChildren(children []TypeInstance, withReturnType bool) []TypeInstance {
	return children
}

func (ins *TypeInstancePrimitive) SkipTL2(r []byte) ([]byte, error) {
	if ins.isString {
		return basictl.SkipSizedValue(r)
	}
	if len(r) < ins.fixedSize {
		return r, io.ErrUnexpectedEOF
	}
	return r[ins.fixedSize:], nil
}

func (ins *TypeInstancePrimitive) IsTL1Bool() (ok bool, falseTag uint32, trueTag uint32) {
	if ins.canonicalName == "bool" && len(ins.tip.combTL1) == 2 {
		return true, ins.tip.combTL1[0].Crc32(), ins.tip.combTL1[1].Crc32()
	}
	return false, 0, 0
}

func (k *Kernel) addPrimitive(name string, tl1name string, historicalName string, fixedSize int, goodForMapKey bool) *KernelType {
	// for the purpose of type check, this is object with no fields, like uint32 = ;
	combTL1 := &tlast.Combinator{
		Construct: tlast.Constructor{
			Name: tlast.Name{Name: tl1name},
		},
		TypeDecl: tlast.TypeDeclaration{
			Name: tlast.Name{Name: utils.ToUpperFirst(tl1name)},
		},
	}
	combTL2 := tlast.TL2Combinator{
		TypeDecl: tlast.TL2TypeDeclaration{
			Name: tlast.TL2TypeName{Name: name},
			Type: tlast.TL2TypeDefinition{},
		},
	}
	ins := TypeInstancePrimitive{
		TypeInstanceCommon: TypeInstanceCommon{
			canonicalName: name,
			tlName:        tlast.TL2TypeName{Name: historicalName},
		},
		isString:      name == "string",
		fixedSize:     fixedSize,
		goodForMapKey: goodForMapKey,
	}
	ref := &TypeInstanceRef{
		ins: &ins,
	}
	kt := &KernelType{
		originTL2:      tl1name == "",
		builtin:        true,
		combTL1:        []*tlast.Combinator{combTL1},
		combTL2:        combTL2,
		instances:      map[string]*TypeInstanceRef{name: ref},
		tl1Names:       map[string]struct{}{},
		tl2Names:       map[string]struct{}{},
		canonicalName:  tlast.TL2TypeName{Name: name},
		historicalName: tlast.TL2TypeName{Name: historicalName},
		canBeBare:      true,
	}
	kt.tl2Names[name] = struct{}{}
	if tl1name != "" {
		kt.tl1Names[tl1name] = struct{}{}
	}
	ins.tip = kt
	if _, ok := k.instances[name]; ok {
		panic(fmt.Sprintf("error adding primitive type %s: exist in global list", name))
	}
	if err := k.addTip(kt, name, tl1name); err != nil {
		panic(fmt.Sprintf("error adding primitive type %s: %v", name, err))
	}
	k.instances[name] = ref
	k.instancesOrdered = append(k.instancesOrdered, ref)
	return kt
}
