// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package checks

import "github.com/VKCOM/tl/internal/tlast"

// functionRemoved forbids removing (or commenting out) an RPC function that existed before:
// existing callers would still send the request, and it would no longer be understood.
//
// Adding new functions is allowed and is not reported here.
func functionRemoved(prev, cur *Schema, r *Reporter) *tlast.ParseError {
	for _, functionName := range prev.FunctionsOrder {
		if _, ok := cur.Functions[functionName]; !ok {
			prevFunction := prev.Functions[functionName]
			if e := r.flag([]string{prevFunction.CommentRight}, prevFunction.Construct.NamePR,
				"function %q cannot be removed", prevFunction.Construct.Name.String()); e != nil {
				return e
			}
		}
	}
	return nil
}
