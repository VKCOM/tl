// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package pure

import (
	"github.com/vkcom/tl/pkg/basictl"
)

type TypeInstanceDict struct {
	TypeInstanceCommon

	fieldType TypeInstanceStruct
}

func (ins *TypeInstanceDict) FindCycle(c *cycleFinder) {
}

func (ins *TypeInstanceDict) CreateValue() KernelValue {
	value := &KernelValueDict{
		instance: ins,
	}
	return value
}

func (ins *TypeInstanceDict) SkipTL2(r []byte) ([]byte, error) {
	return basictl.SkipSizedValue(r)
}

func (k *Kernel) createDict(canonicalName string, keyType *TypeInstanceRef, fieldType *TypeInstanceRef) TypeInstance {
	ins := &TypeInstanceDict{
		TypeInstanceCommon: TypeInstanceCommon{
			canonicalName: canonicalName,
			tip:           nil, // TODO - dicts have no corresponding type
		},
		fieldType: TypeInstanceStruct{
			TypeInstanceCommon: TypeInstanceCommon{
				canonicalName: canonicalName + "__elem",
				tip:           nil, //  TODO - TL2 dict elements have no corresponding type
			},
			isConstructorFields: true,
			fields: []Field{{
				name: "k",
				ins:  keyType,
			}, {
				name: "v",
				ins:  fieldType,
			}},
		},
	}
	return ins
}
