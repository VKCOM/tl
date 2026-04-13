package codecreator

type RustCodeCreator struct {
	BasicCodeCreator
}

func NewRustCodeCreator() *RustCodeCreator {
	return &RustCodeCreator{
		BasicCodeCreator: BasicCodeCreator{
			shift:              "    ",
			ifPrefixTemplate:   "if %[1]s {",
			ifSuffixTemplate:   "}",
			elsePrefixTemplate: "} else {",
			forPrefixTemplate:  "for %[1]s in %[2]s {",
			forSuffixTemplate:  "}",
			commentPrefix:      "// ",
			varTemplate:        "%[1]s_%[2]s",
		},
	}
}
