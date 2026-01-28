package tlcodegen

import (
	"fmt"
	"strconv"
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
	// condition -> if
	ifPrefixTemplate string
	ifSuffixTemplate string
	// (init, condition, step) -> for
	forPrefixTemplate string
	forSuffixTemplate string
	// else
	elsePrefixTemplate string
	// comment
	commentPrefix string

	// (name, suffixSupport) -> variable initialization
	varTemplate string

	// allow Indexed For
	allowIndexedFor bool
}

var phpLanguageBundle = LanguageBundle{
	ifPrefixTemplate:   "if (%[1]s) {",
	ifSuffixTemplate:   "}",
	elsePrefixTemplate: "} else {",
	forPrefixTemplate:  "for (%[1]s;%[2]s;%[3]s) {",
	forSuffixTemplate:  "}",
	commentPrefix:      "// ",
	varTemplate:        "$%[1]s_%[2]s",
	allowIndexedFor:    true,
}

type PhpHelder struct{}

type BasicCodeCreator[T any] struct {
	CodeCreator
	LanguageBundle
	lastVarIdentifier uint32
	langHelp          T
}

type PhpCodeCreator = BasicCodeCreator[PhpHelder]

func NewPhpCodeCreator() PhpCodeCreator {
	return BasicCodeCreator[PhpHelder]{
		CodeCreator: CodeCreator{
			Shift: "  ",
		},
		LanguageBundle: phpLanguageBundle,
	}
}

func (bcc *BasicCodeCreator[T]) addFullBlock(prefix, suffix string, block func(cc *BasicCodeCreator[T])) {
	bcc.CodeCreator.addFullBlock(prefix, suffix, func(cc *CodeCreator) {
		block(bcc)
	})
}

func (bcc *BasicCodeCreator[T]) If(condition string, block func(cc *BasicCodeCreator[T])) {
	bcc.addFullBlock(fmt.Sprintf(bcc.ifPrefixTemplate, condition), bcc.ifSuffixTemplate, block)
}

func (bcc *BasicCodeCreator[T]) IfElse(condition string, block func(cc *BasicCodeCreator[T]), elseBlock func(cc *BasicCodeCreator[T])) {
	bcc.addFullBlock(
		fmt.Sprintf(bcc.ifPrefixTemplate, condition),
		bcc.elsePrefixTemplate,
		block,
	)
	bcc.addFullBlock("", "", elseBlock)
	bcc.AddLines(bcc.ifSuffixTemplate)
}

func (bcc *BasicCodeCreator[T]) For(initState, condition, iterationStep string, block func(cc *BasicCodeCreator[T])) {
	bcc.addFullBlock(
		fmt.Sprintf(bcc.forPrefixTemplate, initState, condition, iterationStep),
		bcc.forSuffixTemplate,
		block,
	)
}

func (bcc *BasicCodeCreator[T]) ForIndexed(indexVar, startValue, upperBound, step string, block func(cc *BasicCodeCreator[T])) {
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

func (bcc *BasicCodeCreator[T]) Comments(lines ...string) {
	for _, s := range lines {
		bcc.AddLines(bcc.commentPrefix + s)
	}
}

func (bcc *BasicCodeCreator[T]) Comment(lines string) {
	bcc.Comments(strings.Split(lines, "\n")...)
}

func (bcc *BasicCodeCreator[T]) NewVariable(name string) string {
	if name == "" {
		name = "var"
	}
	bcc.lastVarIdentifier += 1
	return fmt.Sprintf(bcc.varTemplate, name, strconv.FormatUint(uint64(bcc.lastVarIdentifier), 10))
}
