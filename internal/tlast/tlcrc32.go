// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlast

import (
	"hash/crc32"
)

// Crc32 returns tag of Combinator
// if tag not set explicitly Crc32 will calculate hash of Combinator in
// canonical form
// canonicalForm returns combinator in "canonical" form which is used for
//   - combinator is written in one line
//   - curly braces: '{', '}' are omitted
//   - all tokens are separates from each other with one space
//   - square brackets are also separated from content inside with one space
//     example: "[ T ]"
func (descriptor *Combinator) crc32() uint32 {
	return crc32.ChecksumIEEE([]byte(descriptor.canonicalForm()))
}

func (descriptor *Combinator) Crc32() uint32 {
	return descriptor.Construct.ID
}
