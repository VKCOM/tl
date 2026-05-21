# Changelog

## 0.0.1

- Initial release.
- Syntax highlighting for `.tl2` files (TextMate grammar in JSON).
- Language icon for `.tl2` files.
- Document Symbol Provider (Outline / Go to Symbol): structs, unions with
  constructors as children, aliases, functions.
- Hover Provider: declaration signature + leading `//` comments.
- Definition Provider: Go to Definition / Ctrl+Click on a type reference,
  resolves cross-file within the workspace.
- Completion Provider: built-in primitives + every declared type in the
  workspace, with signature and doc-comment in the popup.
- Diagnostics: references to types not declared anywhere are marked as
  "Unknown type". Severity configurable via `tl2.diagnostics.unknownTypes`.
