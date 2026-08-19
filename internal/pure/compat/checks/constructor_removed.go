// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package checks

import "github.com/VKCOM/tl/internal/tlast"

// constructorRemoved forbids removing (or commenting out) a constructor when that renumbers
// constructors that survive. Variants of a union are numbered by position (see
// pure.TypeInstanceUnion, built with an ordinal index per variant), so dropping a constructor
// shifts every constructor declared after it, and existing peers decode the wrong variant.
//
// Two removals stay safe and are allowed:
//   - dropping a trailing run of constructors: nothing after them survives, so no surviving
//     variant changes its number;
//   - removing the type as a whole: this is normally part of a larger cleanup, and its only
//     wire-visible consequence is that fields can no longer reference the type -- which is
//     reported by field-type-changed, not here.
//
// To waive a flagged removal, tag the constructor in the previous schema (where the linter
// points).
func constructorRemoved(prev, cur *Schema, r *Reporter) *tlast.ParseError {
	for _, typeName := range prev.TypesOrder {
		survives := make(map[tlast.Name]bool, len(cur.Types[typeName]))
		for _, c := range cur.Types[typeName] {
			survives[c.Construct.Name] = true
		}

		// Only constructors before the last surviving one matter: removing them renumbers the
		// survivors. If nothing survives, the whole type is gone, which is allowed.
		prevConstructors := prev.Types[typeName]
		lastSurviving := -1
		for i, c := range prevConstructors {
			if survives[c.Construct.Name] {
				lastSurviving = i
			}
		}
		for i := 0; i <= lastSurviving; i++ {
			prevConstructor := prevConstructors[i]
			if survives[prevConstructor.Construct.Name] {
				continue
			}
			if e := r.flag([]string{prevConstructor.CommentRight}, prevConstructor.Construct.NamePR,
				"constructor %q cannot be removed: the constructors after it in union %q would be renumbered",
				prevConstructor.Construct.Name.String(), typeName.String()); e != nil {
				return e
			}
		}
	}
	return nil
}
