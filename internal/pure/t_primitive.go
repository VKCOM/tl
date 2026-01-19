// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package pure

import (
	"fmt"

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

func (k *Kernel) addPrimitive(name string, originTL2 bool, clone KernelValue, goodForMapKey bool) {
	// for the purpose of type check, this is object with no fields, like uint32 = ;
	combTL1 := &tlast.Combinator{
		Construct: tlast.Constructor{
			Name: tlast.Name{Name: name},
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
		},
		clone:         clone,
		goodForMapKey: goodForMapKey,
	}
	ref := &TypeInstanceRef{
		ins: &ins,
	}
	kt := &KernelType{
		originTL2: originTL2,
		combTL1:   []*tlast.Combinator{combTL1},
		combTL2:   combTL2,
		instances: map[string]*TypeInstanceRef{name: ref},
	}
	if _, ok := k.instances[name]; ok {
		panic(fmt.Sprintf("error adding primitive type %s: exist in global list", name))
	}
	if err := k.addTip(kt, name, ""); err != nil {
		panic(fmt.Sprintf("error adding primitive type %s: %v", name, err))
	}
	k.instances[name] = ref
	// k.instancesOrdered = append(k.instancesOrdered, ref) - we do not yet know if we need them here
}
