// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package checks

import "github.com/VKCOM/tl/internal/tlast"

// newFieldRequiresMask constrains fields appended to a combinator that already existed. Old peers
// stop reading after the fields they know about, so an appended field is only safe if it is gated
// on a field mask (`mask.bit?type`): peers that never set the bit stay byte-compatible.
//
// Brand-new combinators (absent from prev) are unconstrained and may add plain fields freely --
// they are simply not returned by MatchedCombinators, so this scenario never sees them.
//
// Exception for functions: a function that had no field mask at all cannot suddenly gate a field
// on one (there is no mask field to reference yet). Such a function may instead append a single
// bare `#` field as the new mask, provided every field appended after it is gated on some mask
// (not necessarily the one just introduced).
func newFieldRequiresMask(prev, cur *Schema, r *Reporter) *tlast.ParseError {
	for _, pair := range prev.MatchedCombinators(cur) {
		start := len(pair.Prev.Fields)
		if start >= len(pair.Cur.Fields) {
			continue // nothing appended
		}

		// Functions without a prior field mask may lead the appended fields with a bare `#`.
		if pair.Prev.IsFunction && pair.Cur.IsFunction && !hasFieldMaskField(pair.Prev) {
			first := &pair.Cur.Fields[start]
			if first.FieldType.Type.String() == "#" && first.Mask == nil {
				start++
			}
		}

		for i := start; i < len(pair.Cur.Fields); i++ {
			field := &pair.Cur.Fields[i]
			if field.Mask != nil {
				continue
			}
			if e := r.flag([]string{field.CommentRight, pair.Cur.CommentRight}, field.PR,
				"new field %s must be gated on a field mask (write it as `mask.bit?%s`); "+
					"a function with no mask yet may instead append a `#` field first and put the masked fields after it",
				fieldLabel(field, i), field.FieldType.String()); e != nil {
				return e
			}
		}
	}
	return nil
}

// hasFieldMaskField reports whether the combinator already declares a `#` (nat) field that can
// serve as a field mask.
func hasFieldMaskField(c *tlast.Combinator) bool {
	for i := range c.Fields {
		if c.Fields[i].FieldType.Type.String() == "#" {
			return true
		}
	}
	return false
}
