package codecreator

var rustLanguageBundle = LanguageBundle{
	ifPrefixTemplate:   "if (%[1]s) {",
	ifSuffixTemplate:   "}",
	elsePrefixTemplate: "} else {",
	forPrefixTemplate:  "for (%[1]s;%[2]s;%[3]s) {",
	forSuffixTemplate:  "}",
	commentPrefix:      "// ",
	varTemplate:        "%[1]s_%[2]s",
	allowIndexedFor:    true,
}

type RustCodeCreator struct {
	BasicCodeCreator
	RustHelper
}

func NewRustCodeCreator() CppCodeCreator {
	return CppCodeCreator{
		BasicCodeCreator: BasicCodeCreator{
			CodeCreator: CodeCreator{
				Shift: "    ",
			},
			LanguageBundle: rustLanguageBundle,
		},
	}
}

type RustHelper struct{}
