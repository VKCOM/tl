package codecreator

var cppLanguageBundle = LanguageBundle{
	ifPrefixTemplate:   "if (%[1]s) {",
	ifSuffixTemplate:   "}",
	elsePrefixTemplate: "} else {",
	forPrefixTemplate:  "for (%[1]s;%[2]s;%[3]s) {",
	forSuffixTemplate:  "}",
	commentPrefix:      "// ",
	varTemplate:        "%[1]s_%[2]s",
}

type CppCodeCreator struct {
	BasicCodeCreator
}

func NewCppCodeCreator() *CppCodeCreator {
	return &CppCodeCreator{
		BasicCodeCreator: BasicCodeCreator{
			Shift:          "\t",
			LanguageBundle: cppLanguageBundle,
		},
	}
}

func (bcc *CppCodeCreator) ForIndexed(indexVar, startValue, upperBound, step string, block func()) {
	bcc.forIndexed(indexVar, startValue, upperBound, step, block)
}
