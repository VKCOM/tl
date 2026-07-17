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

// Check is a single backward-compatibility scenario. Func compares the previous (prev) and the
// current (cur) schema and reports incompatibilities through r, returning the first non-waived
// violation it finds, or nil if the current schema is compatible for this scenario. Violations
// waived with a comment tag are turned into warnings by r and do not stop the scan.
type Check struct {
	// Name is a short, kebab-case scenario identifier (matches the file name).
	Name string
	Func func(prev, cur *Schema, r *Reporter) *tlast.ParseError
}

// All returns every scenario check to run, in order. Each Func lives in its own file (named after
// the scenario); the Name/Func pairing is wired up here so the scenario files stay pure logic.
// Add new scenarios by writing the function in a new file and listing it below.
//
// TODO: remaining field-level scenarios still to port from
// tlcodegen.checkCombinatorsBackwardCompatibility (tlgen.go): a field's mask reference cannot be
// repointed to a different mask field, and a function's result type cannot change. The nat-usage
// scenarios (a new field's bit must be unused; an appended `#` must actually be used) are
// intentionally left out for now.
func All() []Check {
	return []Check{
		{Name: "constructor-removed", Func: constructorRemoved},
		{Name: "union-order-changed", Func: unionOrderChanged},
		{Name: "function-removed", Func: functionRemoved},
		{Name: "field-removed", Func: fieldRemoved},
		{Name: "template-arguments-removed", Func: templateArgumentsRemoved},
		{Name: "field-type-changed", Func: fieldTypeChanged},
		{Name: "field-bit-changed", Func: fieldBitChanged},
		{Name: "field-mask-added-or-removed", Func: fieldMaskAddedOrRemoved},
		{Name: "new-field-requires-mask", Func: newFieldRequiresMask},
	}
}

// fieldLabel describes a field for error messages: its name in quotes, or its positional index
// when the field is anonymous.
func fieldLabel(f *tlast.Field, idx int) string {
	if f.FieldName != "" {
		return fmt.Sprintf("%q", f.FieldName)
	}
	return fmt.Sprintf("#%d", idx)
}
