// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package gengo

import (
	"sort"

	"github.com/vkcom/tl/internal/pure"
)

func (trw *TypeRWStruct) GetFieldNatProperties(fieldId int) (pure.NatFieldUsage, []uint32) {
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

func (trw *TypeRWStruct) GetFieldNatPropertiesAsUsageMap(fieldId int, inStruct bool, inReturnType bool) (pure.NatFieldUsage, map[uint32]BitUsageInfo) {
	linear2 := trw.pureTypeStruct.GetNatFieldUsage(fieldId, inStruct, inReturnType)

	u2 := map[uint32]BitUsageInfo{}
	for bit, aff := range linear2.AffectedFields {
		if len(aff) == 0 {
			continue
		}
		if u2[uint32(bit)].AffectedFields == nil {
			u2[uint32(bit)] = BitUsageInfo{map[*TypeRWStruct][]int{}}
		}
		af := u2[uint32(bit)].AffectedFields
		for pt, fieldIndexes := range aff {
			wr := trw.wr.gen.getTypeMust(pt).trw.(*TypeRWStruct)
			for _, fi := range fieldIndexes {
				af[wr] = append(af[wr], fi)
			}
		}
	}
	return linear2, u2
}

type BitUsageInfo struct {
	AffectedFields map[*TypeRWStruct][]int
}
