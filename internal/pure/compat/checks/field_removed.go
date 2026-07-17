// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package checks

import "github.com/VKCOM/tl/internal/tlast"

// fieldRemoved forbids dropping (or commenting out) a field from a combinator that survived into
// the new schema. Fields are positional on the wire, so removing one shifts every following field
// and breaks existing readers.
//
// Only the field count is checked here: reordering or retyping surviving fields is a separate
// scenario. Removing a whole combinator is handled by constructor-removed / function-removed, so
// this scenario looks only at combinators present in both versions.
//
// Adding new fields is not reported here (it is constrained by new-field-requires-mask).
func fieldRemoved(prev, cur *Schema, r *Reporter) *tlast.ParseError {
	for _, pair := range prev.MatchedCombinators(cur) {
		if len(pair.Cur.Fields) >= len(pair.Prev.Fields) {
			continue
		}
		// Point at the first previous field that no longer has a counterpart.
		idx := len(pair.Cur.Fields)
		missing := &pair.Prev.Fields[idx]
		if e := r.flag([]string{missing.CommentRight, pair.Prev.CommentRight}, missing.PR,
			"field %s cannot be removed", fieldLabel(missing, idx)); e != nil {
			return e
		}
	}
	return nil
}
