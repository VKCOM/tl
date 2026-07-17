// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Package compat implements a backward-compatibility linter for TL schemas on top of the
// new pure kernel. It is the successor of tlcodegen.CheckBackwardCompatibility, which was
// written for the old kernel.
//
// "Backward compatibility" here is conditional: a newer client can always craft data that an
// older client will fail to parse. The linter only catches the common, mechanical mistakes
// that break the wire format for existing peers.
//
// This is a prototype. Only TL1 combinators are checked, and only a base set of scenarios is
// implemented. Each scenario is a separate checks.Check in the checks subpackage; more can be
// added there incrementally.
package compat

import (
	"io"

	"github.com/VKCOM/tl/internal/pure"
	"github.com/VKCOM/tl/internal/pure/compat/checks"
	"github.com/VKCOM/tl/internal/tlast"
)

// CheckBackwardCompatibility compares two parsed kernels and reports the first change in cur
// that breaks backward compatibility with prev, or nil if none is found. It runs every
// scenario returned by checks.All against the two schemas.
//
// Violations explicitly waived in the schema (via a `// tlgen:nolint` or
// `// tlgen:nolint:ignore-compatibility` comment on the flagged element) are printed as warnings
// to warn instead of stopping the check.
func CheckBackwardCompatibility(prev, cur *pure.Kernel, warn io.Writer) *tlast.ParseError {
	prevSchema := checks.NewSchema(prev)
	curSchema := checks.NewSchema(cur)
	reporter := checks.NewReporter(warn)
	for _, c := range checks.All() {
		if err := c.Func(prevSchema, curSchema, reporter); err != nil {
			return err
		}
	}
	return nil
}
