// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package checks

import "github.com/VKCOM/tl/internal/tlast"

// fieldOrderChanged forbids changing where the existing fields of a combinator sit. Fields are
// positional on the wire, so every field that existed before must keep its position, and new
// fields may only be appended at the end (their constraints are checked by new-field-requires-mask).
//
// Fields are matched by name, which lets the scenario tell apart the ways a position can change
// hands and report the real culprit instead of a shifted positional counterpart:
//   - a new field inserted before an existing one,
//   - an existing field moved to an earlier position, and
//   - an existing field renamed -- indistinguishable from removing it and adding another field in
//     its place, so it is rejected as well (waive it if the rename is deliberate).
//
// Anonymous fields carry no name and cannot be matched this way; pairs involving them are skipped
// here and stay guarded by the positional checks (field-type-changed and friends).
func fieldOrderChanged(prev, cur *Schema, r *Reporter) *tlast.ParseError {
	for _, pair := range prev.MatchedCombinators(cur) {
		n := len(pair.Prev.Fields)
		if len(pair.Cur.Fields) < n {
			n = len(pair.Cur.Fields)
		}
		prevIndex := make(map[string]int, len(pair.Prev.Fields))
		for i := range pair.Prev.Fields {
			if name := pair.Prev.Fields[i].FieldName; name != "" {
				prevIndex[name] = i
			}
		}
		curIndex := make(map[string]int, len(pair.Cur.Fields))
		for i := range pair.Cur.Fields {
			if name := pair.Cur.Fields[i].FieldName; name != "" {
				curIndex[name] = i
			}
		}
		for i := 0; i < n; i++ {
			prevName := pair.Prev.Fields[i].FieldName
			curField := &pair.Cur.Fields[i]
			curName := curField.FieldName
			if prevName == "" || curName == "" || prevName == curName {
				continue
			}
			comments := []string{curField.CommentRight, pair.Cur.CommentRight}
			if j, ok := curIndex[prevName]; ok {
				if j < i {
					continue // the move was already reported (or waived) at position j
				}
				if _, isExisting := prevIndex[curName]; isExisting {
					if e := r.flag(comments, curField.PRName,
						"field %q cannot move before field %q: the order of existing fields must stay stable",
						curName, prevName); e != nil {
						return e
					}
				} else {
					if e := r.flag(comments, curField.PRName,
						"new field %q cannot be inserted before existing field %q: fields are positional on the wire, so new fields can only be appended at the end",
						curName, prevName); e != nil {
						return e
					}
				}
				continue
			}
			// The previous field is gone from cur while the field count did not shrink: either it
			// was renamed, or it was removed and another field took its place. The two are
			// indistinguishable, and a compatibility linter must not produce false negatives.
			if e := r.flag(comments, curField.PRName,
				"field %q cannot be renamed to %q: existing fields must keep their names and positions",
				prevName, curName); e != nil {
				return e
			}
		}
	}
	return nil
}
