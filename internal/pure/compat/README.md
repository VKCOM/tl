# compat — TL backward-compatibility linter

`compat` checks that a new version of a TL schema stays wire-compatible with the previous one.
It is the successor of `tlcodegen.CheckBackwardCompatibility`, rebuilt on top of the new `pure`
kernel.

"Backward compatibility" here is **conditional**: a newer client can always craft data an older
client fails to parse. The linter only catches the common, mechanical mistakes that break the wire
format for existing peers — removing things, renumbering them, or changing their layout.

Only **TL1** combinators are inspected.

## How it works

Two parsed kernels are compared: `prev` (the old, committed schema) and `cur` (the new one). Each
version is turned into an indexed [`checks.Schema`](checks/schema.go), and every scenario in
[`checks.All()`](checks/check.go) runs against the pair. The first non-waived violation is returned
as a `*tlast.ParseError` pointing at the offending line.

```
prev, cur *pure.Kernel
        │
        ▼
checks.NewSchema(prev), checks.NewSchema(cur)   // grouped by type / function, MatchedCombinators cached
        │
        ▼
for _, c := range checks.All(): c.Func(prev, cur, reporter)   // stops at first hard violation
```

## Usage

### From the CLI (`cmd/tl2gen`)

The current schema is passed as a positional argument; the previous version via
`--backwardCompatibilityWith` (comma-separated paths are allowed):

```sh
go run ./cmd/tl2gen \
    --language=compat \
    --backwardCompatibilityWith=schema.old.tl \
    schema.tl
```

On success it prints `TL Backward Compatibility Check Success`; on a violation it prints the
beautiful error, `TL Backward Compatibility Check Failed`, and exits with code 1.

### Programmatically

```go
import (
    "os"

    "github.com/VKCOM/tl/internal/pure"
    "github.com/VKCOM/tl/internal/pure/compat"
)

func check(oldPath, newPath string) error {
    prev := pure.NewKernel(&pure.OptionsKernel{ErrorWriter: os.Stderr})
    if err := prev.AddFilesFromPaths([]string{oldPath}); err != nil {
        return err
    }

    cur := pure.NewKernel(&pure.OptionsKernel{ErrorWriter: os.Stderr})
    if err := cur.AddFilesFromPaths([]string{newPath}); err != nil {
        return err
    }

    // Waived changes (see below) are printed to the io.Writer as warnings; hard violations are
    // returned as a *tlast.ParseError.
    if pe := compat.CheckBackwardCompatibility(prev, cur, os.Stderr); pe != nil {
        return pe
    }
    return nil
}
```

## Example

Say the previous schema is:

```tl
user flags:# name:flags.0?string = User;
```

The following edit is rejected, because field `name` changes its type (`string` → `int`) and old
readers would decode the same bytes as the wrong type:

```tl
user flags:# name:flags.0?int = User;
//                        ^ breaks backward compatibility: type of field "name" cannot change from "string" to "int"
```

Appending a new field is fine **only** if it is gated on a field mask, so peers that never set the
bit stay byte-compatible:

```tl
user flags:# name:flags.0?string age:flags.1?int = User;   // OK: age is optional via flags.1
user flags:# name:flags.0?string age:int         = User;   // rejected: new field must be masked
```

## Waiving a change (`tlgen:nolint`)

Sometimes a change is genuinely acceptable and you want to allow it deliberately. Mark the flagged
element with a comment; the waiver stays in the schema as an obvious, greppable marker for review,
and the linter downgrades the violation to a printed warning instead of failing.

Two forms are honored (same `utils.DoLint` convention used across the repo):

| Comment | Effect |
|---------|--------|
| `// tlgen:nolint` | blanket waiver — silences **every** linter on this element |
| `// tlgen:nolint:ignore-compatibility` | waives **only** the backward-compatibility checks |

```tl
user flags:# name:flags.0?int = User; // tlgen:nolint:ignore-compatibility
```

Notes:
- Put the comment **on the line the linter points at**. For additions/changes that element lives in
  the new schema; for removals it lives in the previous schema (the element is gone from the new one).
- TL combinators are usually single-line, so a trailing comment belongs to the *combinator*, not to
  an individual field. Field-level scenarios therefore accept a waiver from either the field's own
  comment or its enclosing combinator's comment.

## Scenarios

Each scenario lives in its own file under [`checks/`](checks/), named after the scenario, and is
exposed as a `checks.Check{Name, Func}` from `All()`.

| Name | Rejects |
|------|---------|
| `constructor-removed` | removing a union constructor or a whole type |
| `union-order-changed` | reordering the constructors of a union (they are numbered by position) |
| `function-removed` | removing an RPC function |
| `field-removed` | dropping a field from a surviving combinator |
| `template-arguments-removed` | dropping a template argument |
| `field-type-changed` | changing the type (or args, or bare/boxed) of an existing field |
| `field-bit-changed` | moving an existing field to a different field-mask bit |
| `field-mask-added-or-removed` | adding or removing the `mask.bit?` gate of an existing field |
| `new-field-requires-mask` | appending a field to an existing combinator without a field mask (functions may lead the appended fields with a bare `#`) |

Brand-new types and functions are unconstrained — they may declare plain, unmasked fields freely.

Intentionally **not** implemented yet: a field's mask reference cannot be repointed, a function's
result type cannot change, and the nat-usage rules (a new field's bit must be unused; an appended
`#` must actually be used).

## Adding a scenario

1. Create `checks/<scenario_name>.go` with a function
   `func scenarioName(prev, cur *Schema, r *Reporter) *tlast.ParseError`.
2. Reuse `prev.MatchedCombinators(cur)` to iterate combinators present in both versions (its result
   is cached; set `Schema.Force` to bypass the cache if a schema is mutated).
3. At each violation call `r.flag(comments, pos, format, args...)`: pass the candidate waiver
   comments (element + enclosing combinator), and `return` the result if it is non-nil; a `nil`
   means the violation was waived, so keep scanning.
4. Register it in `checks.All()`.
