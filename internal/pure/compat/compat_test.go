// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package compat

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/VKCOM/tl/internal/pure"
)

// buildKernel parses an inline TL1 schema into a fresh kernel. We only parse (no Compile):
// the backward-compatibility check reads combinators via Kernel.TL1(), which is populated at
// parse time. This keeps the test schemas minimal (unresolved primitive references are fine).
func buildKernel(t *testing.T, schema string) *pure.Kernel {
	t.Helper()
	k := pure.NewKernel(&pure.OptionsKernel{ErrorWriter: io.Discard})
	if err := k.AddTL1FromString("test.tl", schema); err != nil {
		t.Fatalf("failed to parse schema:\n%s\nerror: %v", schema, err)
	}
	return k
}

func check(t *testing.T, prev, cur string) error {
	t.Helper()
	err := CheckBackwardCompatibility(buildKernel(t, prev), buildKernel(t, cur), io.Discard)
	if err == nil {
		return nil
	}
	return err
}

func requireIncompatible(t *testing.T, prev, cur string) {
	t.Helper()
	if err := check(t, prev, cur); err == nil {
		t.Fatalf("expected a backward-compatibility error, got none\nprev:\n%s\ncur:\n%s", prev, cur)
	}
}

func requireCompatible(t *testing.T, prev, cur string) {
	t.Helper()
	if err := check(t, prev, cur); err != nil {
		t.Fatalf("expected no backward-compatibility error, got: %v\nprev:\n%s\ncur:\n%s", err, prev, cur)
	}
}

func TestUnionConstructorRemoved(t *testing.T) {
	prev := `
valueInt value:int = Value;
valueStr value:string = Value;
`
	cur := `
valueInt value:int = Value;
`
	requireIncompatible(t, prev, cur)
}

func TestUnionConstructorCommentedOut(t *testing.T) {
	prev := `
valueInt value:int = Value;
valueStr value:string = Value;
`
	cur := `
valueInt value:int = Value;
// valueStr value:string = Value;
`
	requireIncompatible(t, prev, cur)
}

func TestUnionConstructorAdded(t *testing.T) {
	prev := `
valueInt value:int = Value;
valueStr value:string = Value;
`
	cur := `
valueInt value:int = Value;
valueStr value:string = Value;
valueLong value:long = Value;
`
	requireCompatible(t, prev, cur)
}

func TestTypeRemoved(t *testing.T) {
	prev := `
foo x:int = Foo;
bar y:int = Bar;
`
	cur := `
bar y:int = Bar;
`
	requireIncompatible(t, prev, cur)
}

func TestFunctionRemoved(t *testing.T) {
	prev := `
---functions---
@any getValue key:string = Value;
@any setValue key:string value:string = Value;
`
	cur := `
---functions---
@any getValue key:string = Value;
`
	requireIncompatible(t, prev, cur)
}

func TestFunctionAdded(t *testing.T) {
	prev := `
---functions---
@any getValue key:string = Value;
`
	cur := `
---functions---
@any getValue key:string = Value;
@any setValue key:string value:string = Value;
`
	requireCompatible(t, prev, cur)
}

func TestNoChanges(t *testing.T) {
	schema := `
valueInt value:int = Value;
valueStr value:string = Value;

---functions---
@any getValue key:string = Value;
`
	requireCompatible(t, schema, schema)
}

// TestUnionOrderChanged: swapping two union constructors renumbers the variants and is rejected.
func TestUnionOrderChanged(t *testing.T) {
	prev := `
valueInt value:int = Value;
valueStr value:string = Value;
`
	cur := `
valueStr value:string = Value;
valueInt value:int = Value;
`
	requireIncompatible(t, prev, cur)
}

// TestUnionOrderPreservedWithInsertion: inserting a new constructor while keeping the relative
// order of the existing ones is fine.
func TestUnionOrderPreservedWithInsertion(t *testing.T) {
	prev := `
valueInt value:int = Value;
valueStr value:string = Value;
`
	cur := `
valueInt value:int = Value;
valueLong value:long = Value;
valueStr value:string = Value;
`
	requireCompatible(t, prev, cur)
}

func TestConstructorFieldRemoved(t *testing.T) {
	prev := `
point x:int y:int = Point;
`
	cur := `
point x:int = Point;
`
	requireIncompatible(t, prev, cur)
}

func TestFunctionFieldRemoved(t *testing.T) {
	prev := `
---functions---
@any setValue key:string value:string = Value;
`
	cur := `
---functions---
@any setValue key:string = Value;
`
	requireIncompatible(t, prev, cur)
}

func TestFieldTypeChanged(t *testing.T) {
	prev := `
point x:int y:int = Point;
`
	cur := `
point x:int y:long = Point;
`
	requireIncompatible(t, prev, cur)
}

func TestFieldBitChanged(t *testing.T) {
	prev := `
user flags:# name:flags.0?string = User;
`
	cur := `
user flags:# name:flags.1?string = User;
`
	requireIncompatible(t, prev, cur)
}

// TestTemplateArgumentRemoved: dropping a template argument shifts the rest and is rejected.
func TestTemplateArgumentRemoved(t *testing.T) {
	prev := `
pair {x:Type} {y:Type} a:x = Pair x y;
`
	cur := `
pair {x:Type} a:x = Pair x;
`
	requireIncompatible(t, prev, cur)
}

// TestFieldMaskAdded: an existing unconditional field cannot become conditional.
func TestFieldMaskAdded(t *testing.T) {
	prev := `
user flags:# name:string = User;
`
	cur := `
user flags:# name:flags.0?string = User;
`
	requireIncompatible(t, prev, cur)
}

// TestFieldMaskRemoved: an existing conditional field cannot become unconditional.
func TestFieldMaskRemoved(t *testing.T) {
	prev := `
user flags:# name:flags.0?string = User;
`
	cur := `
user flags:# name:string = User;
`
	requireIncompatible(t, prev, cur)
}

// TestMaskedFieldAppended: appending a field gated on an existing mask is fine.
func TestMaskedFieldAppended(t *testing.T) {
	prev := `
user flags:# name:flags.0?string = User;
`
	cur := `
user flags:# name:flags.0?string age:flags.1?int = User;
`
	requireCompatible(t, prev, cur)
}

// TestUnmaskedFieldAppendedToExisting: appending a plain field to a combinator that already
// existed breaks old readers and is rejected.
func TestUnmaskedFieldAppendedToExisting(t *testing.T) {
	prev := `
point x:int y:int = Point;
`
	cur := `
point x:int y:int z:int = Point;
`
	requireIncompatible(t, prev, cur)
}

// TestNewCombinatorUnmaskedFieldsAllowed: a brand-new type/function is unconstrained and may
// declare plain (unmasked) fields.
func TestNewCombinatorUnmaskedFieldsAllowed(t *testing.T) {
	prev := `
foo x:int = Foo;

---functions---
@any ping x:int = Pong;
`
	cur := `
foo x:int = Foo;
bar a:int b:string = Bar;

---functions---
@any ping x:int = Pong;
@any getUser id:int name:string = User;
`
	requireCompatible(t, prev, cur)
}

// TestFunctionAppendWithNewMask: a function that had no mask may append a leading `#` field and
// then gate the rest of the appended fields on it.
func TestFunctionAppendWithNewMask(t *testing.T) {
	prev := `
---functions---
@any getUser id:int = User;
`
	cur := `
---functions---
@any getUser id:int flags:# name:flags.0?string = User;
`
	requireCompatible(t, prev, cur)
}

// TestFunctionAppendWithoutMask: appending a plain field to an existing function, without leading
// it with a `#` mask, is rejected.
func TestFunctionAppendWithoutMask(t *testing.T) {
	prev := `
---functions---
@any getUser id:int = User;
`
	cur := `
---functions---
@any getUser id:int name:string = User;
`
	requireIncompatible(t, prev, cur)
}

// --- suppression via comment tags ---

// TestConstructorRemovalWaivedBlanket: a blanket `// tlgen:nolint` on the removed constructor
// (in the previous schema, where the linter points) waives the removal.
func TestConstructorRemovalWaivedBlanket(t *testing.T) {
	prev := `
valueInt value:int = Value;
valueStr value:string = Value; // tlgen:nolint
`
	cur := `
valueInt value:int = Value;
`
	requireCompatible(t, prev, cur)
}

// TestConstructorRemovalWaivedIgnoreCompat: the compat-scoped tag waives the removal too.
func TestConstructorRemovalWaivedIgnoreCompat(t *testing.T) {
	prev := `
valueInt value:int = Value;
valueStr value:string = Value; // tlgen:nolint:ignore-compatibility
`
	cur := `
valueInt value:int = Value;
`
	requireCompatible(t, prev, cur)
}

// TestFieldTypeChangeWaived: a trailing tag on a (single-line) combinator waives a field change,
// since the combinator's comment is the fallback source for its fields.
func TestFieldTypeChangeWaived(t *testing.T) {
	prev := `
point x:int y:int = Point;
`
	cur := `
point x:int y:long = Point; // tlgen:nolint:ignore-compatibility
`
	requireCompatible(t, prev, cur)
}

// TestUnwaivedChangeStillFails: an unrelated tag must not waive a real change.
func TestUnwaivedChangeStillFails(t *testing.T) {
	prev := `
point x:int y:int = Point;
`
	cur := `
point x:int y:long = Point;
`
	requireIncompatible(t, prev, cur)
}

// TestWaivedChangeEmitsWarning: a waived change is reported as a warning, not swallowed silently.
func TestWaivedChangeEmitsWarning(t *testing.T) {
	prev := `
valueInt value:int = Value;
valueStr value:string = Value; // tlgen:nolint:ignore-compatibility
`
	cur := `
valueInt value:int = Value;
`
	var buf bytes.Buffer
	if err := CheckBackwardCompatibility(buildKernel(t, prev), buildKernel(t, cur), &buf); err != nil {
		t.Fatalf("expected waived change (nil error), got: %v", err)
	}
	if !strings.Contains(buf.String(), "waived") {
		t.Fatalf("expected a warning mentioning the waived change, got: %q", buf.String())
	}
}
