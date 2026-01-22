package tlcodegen

import (
	"fmt"
	"strings"
)

type CodeCreator struct {
	Shift string

	lines      []string
	linesShift []int

	currentShift int
}

func (cc *CodeCreator) AddLines(lines ...string) {
	cc.lines = append(cc.lines, lines...)
	for range lines {
		cc.linesShift = append(cc.linesShift, cc.currentShift)
	}
}

func (cc *CodeCreator) AddShift(shift int) {
	cc.currentShift += shift
}

func (cc *CodeCreator) Print() []string {
	s := make([]string, len(cc.lines))
	for i, line := range cc.lines {
		s[i] = fmt.Sprintf("%s%s", strings.Repeat(cc.Shift, cc.linesShift[i]), line)
	}
	return s
}

func (cc *CodeCreator) AddBlock(block func(cc *CodeCreator)) {
	cc.AddShift(1)
	block(cc)
	cc.AddShift(-1)
}

func (cc *CodeCreator) addFullBlock(prefix, suffix string, block func(cc *CodeCreator)) {
	if prefix != "" {
		cc.AddLines(prefix)
	}
	cc.AddBlock(block)
	if suffix != "" {
		cc.AddLines(suffix)
	}
}

type LanguageBundle struct {
	// if
	ifPrefixTemplate string
	ifSuffixTemplate string
	// for
	forPrefixTemplate string
	forSuffixTemplate string
	// else
	elsePrefixTemplate string
	// comment
	commentPrefix string

	// allow Indexed For
	allowIndexedFor bool
}

var phpLanguageBundle = LanguageBundle{
	ifPrefixTemplate:   "if (%s) {",
	ifSuffixTemplate:   "}",
	elsePrefixTemplate: "} else {",
	forPrefixTemplate:  "for (%[1]s;%[2]s;%[3]s) {",
	forSuffixTemplate:  "}",
	commentPrefix:      "// ",
	allowIndexedFor:    true,
}

type BasicCodeCreator struct {
	CodeCreator
	LanguageBundle
}

func (bcc *BasicCodeCreator) If(condition string, block func(cc *CodeCreator)) {
	bcc.addFullBlock(fmt.Sprintf(bcc.ifPrefixTemplate, condition), bcc.ifSuffixTemplate, block)
}

func (bcc *BasicCodeCreator) IfElse(condition string, block func(cc *CodeCreator), elseBlock func(cc *CodeCreator)) {
	bcc.addFullBlock(
		fmt.Sprintf(bcc.ifPrefixTemplate, condition),
		bcc.elsePrefixTemplate,
		block,
	)
	bcc.AddBlock(elseBlock)
	bcc.AddLines(bcc.ifSuffixTemplate)
}

func (bcc *BasicCodeCreator) For(initState, condition, iterationStep string, block func(cc *CodeCreator)) {
	bcc.addFullBlock(
		fmt.Sprintf(bcc.forPrefixTemplate, initState, condition, iterationStep),
		bcc.forSuffixTemplate,
		block,
	)
}

func (bcc *BasicCodeCreator) ForIndexed(indexVar, startValue, upperBound, step string, block func(cc *CodeCreator)) {
	if !bcc.allowIndexedFor {
		panic("can't use for by index")
	}
	bcc.For(
		fmt.Sprintf("%[1]s = %[2]s", indexVar, startValue),
		fmt.Sprintf("%[1]s < %[2]s", indexVar, upperBound),
		fmt.Sprintf("%[1]s += %[2]s", indexVar, step),
		block,
	)
}

func (bcc *BasicCodeCreator) Comment(lines string) {
	for _, s := range strings.Split(lines, "\n") {
		bcc.AddLines(bcc.commentPrefix + s)
	}
}

func NewPhpCodeCreator() BasicCodeCreator {
	return BasicCodeCreator{
		CodeCreator: CodeCreator{
			Shift: "  ",
		},
		LanguageBundle: phpLanguageBundle,
	}
}
