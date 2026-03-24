// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package pure

import (
	"fmt"
	"slices"

	"github.com/VKCOM/tl/internal/tlast"
)

type NatFieldUsage struct {
	UsedAsMask     bool
	usedAsMaskPR   tlast.PositionRange
	UsedAsSize     bool
	usedAsSizePR   tlast.PositionRange
	AffectedFields [32]map[*TypeInstanceStruct][]int // bit->type->fieldIndexes
}

func (f *NatFieldUsage) UsedBits() []uint32 {
	bits := make([]uint32, 0)
	for bit, aff := range f.AffectedFields {
		if len(aff) == 0 {
			continue
		}
		bits = append(bits, uint32(bit))
	}
	return bits
}

func (f *NatFieldUsage) appendUsage(bit uint32, Ins *TypeInstanceStruct, FieldIndex int) {
	if f.AffectedFields[bit] == nil {
		f.AffectedFields[bit] = map[*TypeInstanceStruct][]int{}
	}
	was := f.AffectedFields[bit][Ins]
	// keep ordered unique state
	for _, u := range was {
		if u == FieldIndex {
			return
		}
	}
	was = append(was, FieldIndex)
	slices.Sort(was)
	f.AffectedFields[bit][Ins] = was
}

func (ins *TypeInstanceStruct) GetNatFieldUsage(fieldIndex int, inStructFields bool, inReturnType bool) NatFieldUsage {
	var natFieldUsage NatFieldUsage

	if inStructFields {
		for i, field := range ins.fields {
			if field.FieldMask() != nil &&
				field.FieldMask().IsField() && field.FieldMask().FieldIndex() == fieldIndex {
				if !natFieldUsage.UsedAsMask {
					natFieldUsage.UsedAsMask = true
					natFieldUsage.usedAsMaskPR = field.pr
				}
				natFieldUsage.appendUsage(field.bitNumber, ins, i)
			}
			for argIndex, natArg := range field.natArgs {
				if natArg.IsField() && natArg.FieldIndex() == fieldIndex {
					visitedNodes := map[TypeInstance]struct{}{}
					markAffectedFields(field.ins.ins, visitedNodes, &natFieldUsage, argIndex)
				}
			}
		}
	}
	if inReturnType && ins.resultType != nil {
		for argIndex, natArg := range ins.resultNatArgs {
			if natArg.IsField() && natArg.FieldIndex() == fieldIndex {
				visitedNodes := map[TypeInstance]struct{}{}
				markAffectedFields(ins.resultType, visitedNodes, &natFieldUsage, argIndex)
			}
		}
	}
	return natFieldUsage
}

func markAffectedFields(node TypeInstance, visitedNodes map[TypeInstance]struct{}, natFieldUsage *NatFieldUsage, natIndex int) {
	if _, ok := visitedNodes[node]; ok {
		return
	}
	visitedNodes[node] = struct{}{}
	if natIndex > len(node.Common().natParams) {
		fmt.Printf("natIndex %s\n", node.CanonicalName())
	}
	natParamName := node.Common().natParams[natIndex]
	switch ins := node.(type) {
	case *TypeInstanceStruct:
		for fieldIndex, field := range ins.fields {
			if field.FieldMask() != nil && !field.FieldMask().IsField() && !field.FieldMask().IsNumber() && field.FieldMask().name == natParamName {
				if !natFieldUsage.UsedAsMask {
					natFieldUsage.UsedAsMask = true
					natFieldUsage.usedAsMaskPR = field.pr
				}
				natFieldUsage.appendUsage(field.bitNumber, ins, fieldIndex)
			}
			for argIndex, natArg := range field.NatArgs() {
				if natArg.IsNatParam() && natArg.name == natParamName {
					markAffectedFields(field.ins.ins, visitedNodes, natFieldUsage, argIndex)
				}
			}
		}
	case *TypeInstanceUnion:
		for _, variant := range ins.variantTypes {
			markAffectedFields(variant, visitedNodes, natFieldUsage, natIndex)
		}
	case *TypeInstanceArray:
		// tuple
		if ins.dynamicSize && natIndex == 0 {
			if !natFieldUsage.UsedAsSize {
				natFieldUsage.UsedAsSize = true
				natFieldUsage.usedAsSizePR = ins.resolvedType.BracketType.IndexType.PR // TODO - check
			}
		} else {
			for argIndex, natArg := range ins.field.NatArgs() {
				if natArg.IsNatParam() && natArg.name == natParamName {
					markAffectedFields(ins.field.ins.ins, visitedNodes, natFieldUsage, argIndex)
				}
			}
		}
	case *TypeInstanceDict:
		for argIndex, natArg := range ins.field.NatArgs() {
			if natArg.IsNatParam() && natArg.name == natParamName {
				markAffectedFields(ins.field.ins.ins, visitedNodes, natFieldUsage, argIndex)
			}
		}
	}
}
