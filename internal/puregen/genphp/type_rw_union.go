// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package genphp

import "github.com/VKCOM/tl/internal/pure"

type TypeRWUnion struct {
	pureType *pure.TypeInstanceUnion

	wr     *TypeRWWrapper
	Fields []Field
	IsEnum bool

	fieldsDec Deconflicter // TODO - add all generated methods here
}
