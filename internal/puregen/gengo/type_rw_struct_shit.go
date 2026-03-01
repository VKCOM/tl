// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package gengo

import (
	"github.com/vkcom/tl/internal/pure"
)

func (trw *TypeRWStruct) GetFieldNatProperties(fieldId int, inStruct bool, inReturnType bool) (pure.NatFieldUsage, []uint32) {
	linear2 := trw.pureTypeStruct.GetNatFieldUsage(fieldId, inStruct, inReturnType)

	bits := make([]uint32, 0)
	for bit, aff := range linear2.AffectedFields {
		if len(aff) == 0 {
			continue
		}
		bits = append(bits, uint32(bit))
	}
	return linear2, bits
}
