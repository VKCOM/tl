package codecreator

var cppLanguageBundle = LanguageBundle{
	ifPrefixTemplate:   "if (%[1]s) {",
	ifSuffixTemplate:   "}",
	elsePrefixTemplate: "} else {",
	forPrefixTemplate:  "for (%[1]s;%[2]s;%[3]s) {",
	forSuffixTemplate:  "}",
	commentPrefix:      "// ",
	varTemplate:        "%[1]s_%[2]s",
	allowIndexedFor:    true,
}

type CppCodeCreator struct {
	BasicCodeCreator
	CppHelper
}

func NewCppCodeCreator() CppCodeCreator {
	return CppCodeCreator{
		BasicCodeCreator: BasicCodeCreator{
			CodeCreator: CodeCreator{
				Shift: "\t",
			},
			LanguageBundle: cppLanguageBundle,
		},
	}
}

type CppHelper struct{}
