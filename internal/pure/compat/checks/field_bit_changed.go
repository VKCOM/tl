// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package checks

import "github.com/VKCOM/tl/internal/tlast"

// fieldBitChanged forbids changing which field-mask bit a field is gated on. A field written as
// `flags.3?int` is present on the wire only when bit 3 of its mask is set; moving it to another
// bit means old and new peers disagree about when the field is present, corrupting the stream.
//
// Only fields that are gated on a mask in both versions are compared. Adding or removing the mask
// entirely, or repointing it to a different mask field, are separate scenarios.
func fieldBitChanged(prev, cur *Schema, r *Reporter) *tlast.ParseError {
	for _, pair := range prev.MatchedCombinators(cur) {
		n := len(pair.Prev.Fields)
		if len(pair.Cur.Fields) < n {
			n = len(pair.Cur.Fields)
		}
		for i := 0; i < n; i++ {
			prevField := &pair.Prev.Fields[i]
			curField := &pair.Cur.Fields[i]
			if prevField.Mask == nil || curField.Mask == nil {
				continue
			}
			if prevField.Mask.BitNumber == curField.Mask.BitNumber {
				continue
			}
			if e := r.flag([]string{curField.CommentRight, pair.Cur.CommentRight}, curField.Mask.PRBits,
				"field-mask bit of field %s cannot change from %d to %d",
				fieldLabel(curField, i), prevField.Mask.BitNumber, curField.Mask.BitNumber); e != nil {
				return e
			}
		}
	}
	return nil
}
