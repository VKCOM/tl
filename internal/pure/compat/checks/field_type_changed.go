// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package checks

import "github.com/VKCOM/tl/internal/tlast"

// fieldTypeChanged forbids changing the type of a field that exists in both schema versions.
// Fields are matched positionally, so a reader of the old schema decodes field i as the old type;
// giving it a different type (or different type arguments, or flipping bare/boxed) reinterprets
// the same bytes and corrupts every peer.
//
// The comparison is the canonical string form of the field's type reference, which also forbids
// reordering type arguments. Note this is stricter than the old linter for the rare case of a
// renamed template argument used positionally; that is acceptable for a compatibility linter,
// which must not produce false negatives.
//
// The field-mask condition (`bit?`) is not part of the type here -- changing it is reported by
// field-bit-changed instead.
func fieldTypeChanged(prev, cur *Schema, r *Reporter) *tlast.ParseError {
	for _, pair := range prev.MatchedCombinators(cur) {
		n := len(pair.Prev.Fields)
		if len(pair.Cur.Fields) < n {
			n = len(pair.Cur.Fields)
		}
		for i := 0; i < n; i++ {
			prevField := &pair.Prev.Fields[i]
			curField := &pair.Cur.Fields[i]
			if prevField.FieldType.String() == curField.FieldType.String() {
				continue
			}
			if e := r.flag([]string{curField.CommentRight, pair.Cur.CommentRight}, curField.FieldType.PR,
				"type of field %s cannot change from %q to %q",
				fieldLabel(curField, i), prevField.FieldType.String(), curField.FieldType.String()); e != nil {
				return e
			}
		}
	}
	return nil
}
