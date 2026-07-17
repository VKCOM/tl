// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package checks

import "github.com/VKCOM/tl/internal/tlast"

// unionOrderChanged forbids reordering the constructors of a union. Variants are numbered by their
// position in the union (see pure.TypeInstanceUnion, built with an ordinal index per variant), so
// swapping two constructors reassigns their numbers and every peer decodes the wrong variant.
//
// Only constructors present in both versions are compared: a removed one is reported by
// constructor-removed, and a newly added one may appear among the others without, on its own,
// changing the relative order of the pre-existing constructors. A waived reorder is skipped and
// does not shift the reference point for the constructors after it.
func unionOrderChanged(prev, cur *Schema, r *Reporter) *tlast.ParseError {
	for _, typeName := range prev.TypesOrder {
		curByName := make(map[tlast.Name]*tlast.Combinator, len(cur.Types[typeName]))
		curPos := make(map[tlast.Name]int, len(cur.Types[typeName]))
		for i, c := range cur.Types[typeName] {
			curByName[c.Construct.Name] = c
			curPos[c.Construct.Name] = i
		}

		lastPos := -1
		var lastName tlast.Name
		for _, c := range prev.Types[typeName] {
			curConstructor, ok := curByName[c.Construct.Name]
			if !ok {
				continue // removed: constructor-removed reports it
			}
			pos := curPos[c.Construct.Name]
			if pos < lastPos {
				if e := r.flag([]string{curConstructor.CommentRight}, curConstructor.Construct.NamePR,
					"constructor %q cannot move before %q in union %q: the order of union constructors must stay stable",
					c.Construct.Name.String(), lastName.String(), typeName.String()); e != nil {
					return e
				}
				continue // waived: keep the previous reference point
			}
			lastPos = pos
			lastName = c.Construct.Name
		}
	}
	return nil
}
