package codecreator

type CppCodeCreator struct {
	BasicCodeCreator
}

func NewCppCodeCreator() *CppCodeCreator {
	return &CppCodeCreator{
		BasicCodeCreator: BasicCodeCreator{
			shift:              "\t",
			ifPrefixTemplate:   "if (%[1]s) {",
			ifSuffixTemplate:   "}",
			elsePrefixTemplate: "} else {",
			forPrefixTemplate:  "for (%[1]s;%[2]s;%[3]s) {",
			forSuffixTemplate:  "}",
			commentPrefix:      "// ",
			varTemplate:        "%[1]s_%[2]s",
		},
	}
}

func (bcc *CppCodeCreator) ForIndexed(indexVar, startValue, upperBound, step string, block func()) {
	bcc.forIndexed(indexVar, startValue, upperBound, step, block)
}
