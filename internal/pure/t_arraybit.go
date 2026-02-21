// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package pure

import (
	"github.com/vkcom/tl/pkg/basictl"
)

type TypeInstanceArrayBit struct {
	TypeInstanceCommon
	isTuple bool
	count   int
}

func (ins *TypeInstanceArrayBit) FindCycle(c *cycleFinder) {
}

func (ins *TypeInstanceArrayBit) GetChildren(children []TypeInstance, withReturnType bool) []TypeInstance {
	return children
}

func (ins *TypeInstanceArrayBit) CreateValue() KernelValue {
	value := &KernelValueArrayBit{
		instance: ins,
	}
	if ins.isTuple {
		value.resize(ins.count)
	}
	return value
}

func (ins *TypeInstanceArrayBit) SkipTL2(r []byte) ([]byte, error) {
	return basictl.SkipSizedValue(r)
}
