# TL2 for VS Code

Syntax highlighting and basic navigation for [TL2](https://github.com/VKCOM/tl/blob/master/docs/TL2Primer.pdf) — Type Language 2.

## Features

- Syntax highlighting for `.tl2` files (TextMate grammar)
- Comments toggling (`//`, `/* */`)
- Bracket matching and auto-closing for `[]`, `()`, `<>`
- Document outline / Go to Symbol (`Ctrl+Shift+O`) — lists every declaration:
  - structs (`name = ...`)
  - unions (`name = | A | B ...`) with constructors as children
  - aliases (`name <=> ...`)
  - functions (`name#crc32 ... => ...`)
- **Go to Definition** (`F12` / `Ctrl+Click`) on any type reference —
  works across files in the workspace
- **Hover** shows declaration signature and leading `//` comments
- **Autocompletion** of all types declared in the workspace + built-in primitives
- **Diagnostics**: references to types that don't exist anywhere are flagged
  as "Unknown type". Configurable severity (see below).

## Built-in types highlighted as primitives

`int32`, `uint32`, `int64`, `uint64`, `float32`, `float64`, `byte`, `string`, `bool`, `bit`, `true`.

## Settings

- `tl2.diagnostics.unknownTypes` (default: `"warning"`) — severity for
  references to types not declared in the workspace. Allowed values:
  `"error"`, `"warning"`, `"information"`, `"hint"`, `"off"`. Set to `"off"`
  if your project mixes TL2 files with external references the extension
  can't see.

## Running locally (development)

1. Open this folder (`TL2VSCode/`) in VS Code:

   ```sh
   code /path/to/tl/TL2VSCode
   ```

2. Install dependencies and compile (only needed once, and after `src/` changes):

   ```sh
   npm install
   npm run compile
   ```

3. Press **F5** (or *Run → Start Debugging*). A second VS Code window opens
   with `[Extension Development Host]` in the title — your extension is loaded there.

4. In that second window, open any `.tl2` file (e.g. `internal/tlcodegen/test/tls/cases.tl2`).
   You should see syntax highlighting, Outline (`Ctrl+Shift+O`), and hover/Go-to-Definition.

While iterating: after editing `src/extension.ts`, run `npm run compile` (or `npm run watch` in
a separate terminal for auto-rebuild), then `Ctrl+R` in the Extension Development Host
window to reload it. Grammar changes (`syntaxes/*.json`) and `package.json` changes also need
the window reloaded.

## Installing persistently (across all VS Code windows)

Symlink the folder into your VS Code extensions directory:

```sh
ln -s /absolute/path/to/tl/TL2VSCode ~/.vscode/extensions/tl2-0.0.1
```

Restart VS Code. To update: pull repo, `npm run compile`, `Developer: Reload Window`.

## Packaging a `.vsix`

```sh
npm install -g @vscode/vsce
vsce package
code --install-extension tl2-0.0.1.vsix
```

## Replacing the icon

The icons live in [`icons/`](icons/):

- [`tl2-language-light.png`](icons/tl2-language-light.png) — dark glyph, used on light VS Code themes.
- [`tl2-language-dark.png`](icons/tl2-language-dark.png) — light glyph, used on dark VS Code themes.

To swap them, replace those files (or change the `icon.light`/`icon.dark`
paths in [`package.json`](package.json)). SVG also works and is preferred for
sharp rendering at all sizes.

The icon appears in file tabs and the Outline view. In the Explorer tree it
shows up only if your selected File Icon Theme doesn't define a `.tl2` icon
(most don't), so usually you'll see it.

## Customizing colors

The grammar only tags tokens with **TextMate scope names** (e.g.
`entity.name.type.tl2`, `variable.parameter.field.tl2`). Colors are picked by
your current color theme. Different themes color the same scope very
differently, so if anything looks off — it's the theme, not the extension.

To inspect what is actually under the cursor: `Ctrl+Shift+P` →
**Developer: Inspect Editor Tokens and Scopes**. You'll see both the TextMate
scope and which theme rule produced the color.

If your theme paints fields and types the same color, override them just for
TL2 in your User Settings (`settings.json`):

```jsonc
"editor.tokenColorCustomizations": {
  "textMateRules": [
    {
      "scope": "variable.parameter.field.tl2",
      "settings": { "foreground": "#9CDCFE" }
    },
    {
      "scope": "entity.name.type.tl2",
      "settings": { "foreground": "#4EC9B0" }
    },
    {
      "scope": "entity.name.namespace.tl2",
      "settings": { "foreground": "#DCDCAA", "fontStyle": "italic" }
    },
    {
      "scope": "comment.metadata.crc32.tl2",
      "settings": { "foreground": "#6A6A6A" }
    }
  ]
}
```

### Scopes used by the grammar

| Token | Scope |
| --- | --- |
| Comments | `comment.line.double-slash.tl2`, `comment.block.tl2` |
| Annotation `@read` | `entity.name.tag.annotation.tl2` |
| Built-in types (`int32`, `string`, …) | `support.type.primitive.tl2` |
| Type name in declaration | `entity.name.type.declaration.tl2` |
| Namespace in declaration (left of `.`) | `entity.name.namespace.declaration.tl2` |
| CRC32 magic (`#1b8b9feb`) | `comment.metadata.crc32.tl2` |
| Template parameter (`X` in `<X:Type>`) | `variable.parameter.template.tl2` |
| Template category (`Type`, `#`) | `storage.type.template.category.tl2` |
| Operators `=`, `<=>`, `=>`, `\|`, `?` | `keyword.operator.{assignment,alias,return,union,optional}.tl2` |
| Field name (`name:`) | `variable.parameter.field.tl2` |
| Anonymous field (`_:`) | `variable.language.anonymous.tl2` |
| Type reference (RHS) | `entity.name.type.tl2` |
| Namespace reference (RHS) | `entity.name.namespace.tl2` |
| Union constructor (`\| ctor`) | `entity.name.function.constructor.tl2` |
