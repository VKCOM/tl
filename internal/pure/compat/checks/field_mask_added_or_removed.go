// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package checks

import "github.com/VKCOM/tl/internal/tlast"

// fieldMaskAddedOrRemoved forbids toggling whether an existing field is gated on a field mask.
// A field written plainly is always present on the wire; a field written as `flags.3?int` is
// present only conditionally. Switching between the two for a field that already existed changes
// the byte layout old peers expect, so both directions are rejected:
//
//   - adding `mask.bit?` to a previously unconditional field, and
//   - dropping `mask.bit?` from a previously conditional field.
//
// Only fields present in both versions are compared; newly appended fields are the concern of
// new-field-requires-mask.
func fieldMaskAddedOrRemoved(prev, cur *Schema, r *Reporter) *tlast.ParseError {
	for _, pair := range prev.MatchedCombinators(cur) {
		n := len(pair.Prev.Fields)
		if len(pair.Cur.Fields) < n {
			n = len(pair.Cur.Fields)
		}
		for i := 0; i < n; i++ {
			prevField := &pair.Prev.Fields[i]
			curField := &pair.Cur.Fields[i]
			if (prevField.Mask == nil) == (curField.Mask == nil) {
				continue
			}

			var e *tlast.ParseError
			if curField.Mask != nil {
				e = r.flag([]string{curField.CommentRight, pair.Cur.CommentRight}, curField.Mask.PRName,
					"field %s cannot gain a field mask: it was unconditional before", fieldLabel(curField, i))
			} else {
				e = r.flag([]string{curField.CommentRight, pair.Cur.CommentRight}, curField.PR,
					"field %s cannot lose its field mask: it was conditional before", fieldLabel(curField, i))
			}
			if e != nil {
				return e
			}
		}
	}
	return nil
}
