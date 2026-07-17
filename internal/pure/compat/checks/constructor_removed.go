// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package checks

import "github.com/VKCOM/tl/internal/tlast"

// constructorRemoved forbids removing (or commenting out) a constructor that existed before.
//
// This covers two user-visible cases at once:
//   - dropping one variant of a union (`a = T; b = T;` -> `a = T;`) changes variant numbering
//     and breaks readers of the removed variant;
//   - removing a whole type is just the case where all of its constructors disappear.
//
// Adding new constructors is allowed and is not reported here. To waive a removal, tag the
// constructor in the previous schema (where the linter points).
func constructorRemoved(prev, cur *Schema, r *Reporter) *tlast.ParseError {
	for _, typeName := range prev.TypesOrder {
		curByConstructor := make(map[tlast.Name]*tlast.Combinator, len(cur.Types[typeName]))
		for _, c := range cur.Types[typeName] {
			curByConstructor[c.Construct.Name] = c
		}
		for _, prevConstructor := range prev.Types[typeName] {
			if _, ok := curByConstructor[prevConstructor.Construct.Name]; !ok {
				if e := r.flag([]string{prevConstructor.CommentRight}, prevConstructor.Construct.NamePR,
					"constructor %q cannot be removed", prevConstructor.Construct.Name.String()); e != nil {
					return e
				}
			}
		}
	}
	return nil
}
