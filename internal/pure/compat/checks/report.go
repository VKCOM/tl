// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package checks

import (
	"fmt"
	"io"
	"strings"

	"github.com/VKCOM/tl/internal/tlast"
	"github.com/VKCOM/tl/internal/utils"
)

// ignoreCompatTag is the backward-compatibility-scoped waiver. Unlike the blanket `tlgen:nolint`,
// it silences only the compatibility linter, so unrelated lints on the same element keep firing.
const ignoreCompatTag = "tlgen:nolint:ignore-compatibility"

// Reporter decides what happens at a detected incompatibility: a hard error that stops the linter,
// or -- when the flagged element opts out via a comment tag -- a printed warning that lets the run
// continue. The waiving comment stays in the schema as an obvious, greppable marker for reviewers.
type Reporter struct {
	warn io.Writer
}

// NewReporter builds a Reporter that prints waived-change warnings to warn (nil -> discarded).
func NewReporter(warn io.Writer) *Reporter {
	if warn == nil {
		warn = io.Discard
	}
	return &Reporter{warn: warn}
}

// flag reports an incompatibility found at pr. comments lists the right-hand comments that may
// waive it -- typically the flagged element's own comment plus its enclosing combinator's, since
// a single-line combinator's trailing comment belongs to the combinator, not the field. If any of
// them waives the check (see suppressed), flag prints a warning and returns nil, so the caller
// keeps scanning for further, non-waived violations. Otherwise it returns the error the scenario
// should propagate.
func (r *Reporter) flag(comments []string, pr tlast.PositionRange, format string, args ...any) *tlast.ParseError {
	for _, c := range comments {
		if suppressed(c) {
			pr.BeautifulError(fmt.Errorf("backward-compatibility change waived by tlgen:nolint: "+format, args...)).
				PrintWarning(r.warn, nil)
			return nil
		}
	}
	return pr.BeautifulError(fmt.Errorf("breaks backward compatibility: "+format, args...))
}

// suppressed reports whether a backward-compatibility violation on an element carrying commentRight
// has been explicitly waived. Two forms are honored:
//   - `// tlgen:nolint` -- the repo-wide blanket waiver (utils.DoLint), silences every linter;
//   - `// tlgen:nolint:ignore-compatibility` -- waives only the backward-compatibility checks.
func suppressed(commentRight string) bool {
	if !utils.DoLint(commentRight) { // blanket tlgen:nolint
		return true
	}
	if len(commentRight) < 2 {
		return false
	}
	for _, f := range strings.Fields(commentRight[2:]) {
		if f == ignoreCompatTag {
			return true
		}
	}
	return false
}
