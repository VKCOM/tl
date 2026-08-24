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
// Fields are matched by name first, so removing a field from the middle points at the field that
// is actually gone rather than at its positional counterpart. Anonymous fields carry no name, so
// if every named field is still present but the field count still shrank, the first previous
// field without a positional counterpart is reported.
//
// Reordering surviving fields is a separate scenario (field-order-changed). Removing a whole
// combinator is handled by constructor-removed (or allowed outright, for functions), so this
// scenario looks only at combinators present in both versions.
//
// Adding new fields is not reported here (it is constrained by new-field-requires-mask).
func fieldRemoved(prev, cur *Schema, r *Reporter) *tlast.ParseError {
	for _, pair := range prev.MatchedCombinators(cur) {
		if len(pair.Cur.Fields) >= len(pair.Prev.Fields) {
			continue
		}
		curNames := make(map[string]bool, len(pair.Cur.Fields))
		for i := range pair.Cur.Fields {
			if name := pair.Cur.Fields[i].FieldName; name != "" {
				curNames[name] = true
			}
		}
		foundNamed := false
		for i := range pair.Prev.Fields {
			prevField := &pair.Prev.Fields[i]
			if prevField.FieldName == "" || curNames[prevField.FieldName] {
				continue
			}
			foundNamed = true
			if e := r.flag([]string{prevField.CommentRight, pair.Prev.CommentRight}, prevField.PR,
				"field %s cannot be removed", fieldLabel(prevField, i)); e != nil {
				return e
			}
		}
		if foundNamed {
			continue
		}
		// No named field is missing, so an anonymous one was dropped: point at the first previous
		// field that no longer has a positional counterpart.
		idx := len(pair.Cur.Fields)
		missing := &pair.Prev.Fields[idx]
		if e := r.flag([]string{missing.CommentRight, pair.Prev.CommentRight}, missing.PR,
			"field %s cannot be removed", fieldLabel(missing, idx)); e != nil {
			return e
		}
	}
	return nil
}
