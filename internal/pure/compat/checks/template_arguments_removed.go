// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package checks

import (
	"fmt"

	"github.com/VKCOM/tl/internal/tlast"
)

// templateArgumentsRemoved forbids dropping a template argument from a combinator that survived
// into the new schema. Template arguments are positional, so removing one shifts the rest and
// changes how every instantiation is interpreted.
//
// Template arguments carry no right-hand comment of their own, so a waiver is read from the
// combinator's comment (in the previous schema).
//
// Adding new template arguments is not reported here.
func templateArgumentsRemoved(prev, cur *Schema, r *Reporter) *tlast.ParseError {
	for _, pair := range prev.MatchedCombinators(cur) {
		if len(pair.Cur.TemplateArguments) >= len(pair.Prev.TemplateArguments) {
			continue
		}
		idx := len(pair.Cur.TemplateArguments)
		missing := pair.Prev.TemplateArguments[idx]
		name := fmt.Sprintf("#%d", idx)
		if missing.FieldName != "" {
			name = fmt.Sprintf("%q", missing.FieldName)
		}
		if e := r.flag([]string{pair.Prev.CommentRight}, missing.PR,
			"template argument %s cannot be removed", name); e != nil {
			return e
		}
	}
	return nil
}
