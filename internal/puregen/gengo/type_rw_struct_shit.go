// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package gengo

import "sort"

// TODO - this must be rewritten in more clean style...

type FieldNatProperties = int

const (
	FieldIsNotNat        FieldNatProperties = 0
	FieldIsNat           FieldNatProperties = 1
	FieldUsedAsFieldMask FieldNatProperties = 2
	FieldUsedAsSize      FieldNatProperties = 4
)

func (trw *TypeRWStruct) GetFieldNatProperties(fieldId int) (FieldNatProperties, []uint32) {
	result, affectedIndexes := trw.GetFieldNatPropertiesAsUsageMap(fieldId, true, true)
	indexes := make([]uint32, 0)
	for i := range affectedIndexes {
		indexes = append(indexes, i)
	}
	// not necessary
	sort.Slice(indexes, func(i, j int) bool {
		return indexes[i] < indexes[j]
	})
	return result, indexes
}

func (trw *TypeRWStruct) GetFieldNatPropertiesAsUsageMap(fieldId int, inStruct, inReturnType bool) (FieldNatProperties, map[uint32]BitUsageInfo) {
	if fieldId < 0 || len(trw.Fields) <= fieldId {
		return FieldIsNotNat, nil
	}
	targetField := trw.Fields[fieldId]
	pr, isPr := targetField.t.trw.(*TypeRWPrimitive)
	if !isPr || pr.tlType != "#" {
		return FieldIsNotNat, nil
	}
	result := FieldIsNat
	affectedIndexes := make(map[uint32]BitUsageInfo)
	natParamUsageMap := make(map[VisitedTypeNatParam]VisitResult)
	if inStruct {
		for i, f := range trw.Fields {
			if i == fieldId {
				continue
			}
			if f.FieldMask() != nil &&
				f.FieldMask().IsField() &&
				f.FieldMask().FieldIndex() == fieldId {
				if _, hasBit := affectedIndexes[f.BitNumber()]; !hasBit {
					affectedIndexes[f.BitNumber()] = BitUsageInfo{AffectedFields: map[*TypeRWStruct][]int{}}
				}
				affectedIndexes[f.BitNumber()].AffectedFields[trw] = append(affectedIndexes[f.BitNumber()].AffectedFields[trw], i)

				result |= FieldUsedAsFieldMask
			}
			natIndexes := make([]int, 0)
			for j, natArg := range f.NatArgs() {
				if natArg.IsField() && natArg.FieldIndex() == fieldId {
					natIndexes = append(natIndexes, j)
				}
			}
			for _, j := range natIndexes {
				visit(f.t, j, &natParamUsageMap, &affectedIndexes, &result)
			}
		}
	}

	if inReturnType && (trw.ResultType != nil) {
		for j, natArg := range trw.ResultNatArgs {
			if natArg.IsField() && natArg.FieldIndex() == fieldId {
				visit(trw.ResultType, j, &natParamUsageMap, &affectedIndexes, &result)
			}
		}
	}

	return result, affectedIndexes
}

type BitUsageInfo struct {
	AffectedFields map[*TypeRWStruct][]int
}

type VisitedTypeNatParam struct {
	Type_    string
	NatIndex int
}

type VisitResult = int

const (
	VisitSuccess VisitResult = iota
	VisitFail
	VisitInProgress
)

func visit(
	t *TypeRWWrapper,
	natIndex int,
	visitResults *map[VisitedTypeNatParam]VisitResult,
	affectedIndexes *map[uint32]BitUsageInfo,
	natProps *FieldNatProperties,
) VisitResult {
	natParamName := t.NatParams[natIndex]
	typeName := t.goGlobalName
	key := VisitedTypeNatParam{typeName, natIndex}

	visitResult, isVisited := (*visitResults)[key]
	if isVisited {
		return visitResult
	}
	(*visitResults)[key] = VisitInProgress

	switch i := t.trw.(type) {
	case *TypeRWStruct:
		{
			for fId, f := range i.Fields {
				if f.FieldMask() != nil &&
					!f.FieldMask().IsField() &&
					!f.FieldMask().IsNumber() &&
					natParamName == f.FieldMask().Name() {
					*natProps |= FieldUsedAsFieldMask
					if _, hasBit := (*affectedIndexes)[f.BitNumber()]; !hasBit {
						(*affectedIndexes)[f.BitNumber()] = BitUsageInfo{AffectedFields: map[*TypeRWStruct][]int{}}
					}
					(*affectedIndexes)[f.BitNumber()].AffectedFields[i] = append((*affectedIndexes)[f.BitNumber()].AffectedFields[i], fId)
					(*visitResults)[key] = VisitSuccess
				}
				natIndexes := make([]int, 0)
				for i, natParam := range f.NatArgs() {
					if natParam.Name() == natParamName {
						natIndexes = append(natIndexes, i)
					}
				}
				for _, index := range natIndexes {
					res := visit(f.t, index, visitResults, affectedIndexes, natProps)
					if res == VisitSuccess {
						(*visitResults)[key] = VisitSuccess
					}
				}
			}
		}
	case *TypeRWUnion:
		{
			for _, f := range i.Fields {
				res := visit(f.t, natIndex, visitResults, affectedIndexes, natProps)
				if res == VisitSuccess {
					(*visitResults)[key] = VisitSuccess
				}
			}
		}
	case *TypeRWMaybe:
		{
			res := visit(i.element.t, natIndex, visitResults, affectedIndexes, natProps)
			if res == VisitSuccess {
				(*visitResults)[key] = VisitSuccess
			}
		}
	case *TypeRWBrackets:
		{
			// tuple
			if !i.vectorLike && i.dynamicSize && natIndex == 0 {
				*natProps |= FieldUsedAsSize
			} else {
				elementType := i.element.t
				natIndexes := make([]int, 0)
				for i, natParam := range i.element.NatArgs() {
					if !natParam.IsNumber() && natParam.Name() == natParamName {
						natIndexes = append(natIndexes, i)
					}
				}
				for _, index := range natIndexes {
					res := visit(elementType, index, visitResults, affectedIndexes, natProps)
					if res == VisitSuccess {
						(*visitResults)[key] = VisitSuccess
					}
				}
			}
		}
	case *TypeRWDict:
		{
			elementType := i.element.t
			natIndexes := make([]int, 0)
			for i, natParam := range i.element.NatArgs() {
				if !natParam.IsNumber() && natParam.Name() == natParamName {
					natIndexes = append(natIndexes, i)
				}
			}
			for _, index := range natIndexes {
				res := visit(elementType, index, visitResults, affectedIndexes, natProps)
				if res == VisitSuccess {
					(*visitResults)[key] = VisitSuccess
				}
			}
		}
	}

	if (*visitResults)[key] == VisitInProgress {
		(*visitResults)[key] = VisitFail
	}
	return (*visitResults)[key]
}
