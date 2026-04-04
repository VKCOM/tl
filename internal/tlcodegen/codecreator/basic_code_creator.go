package codecreator

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

func (cc *CodeCreator) AddLinef(format string, args ...any) {
	cc.lines = append(cc.lines, fmt.Sprintf(format, args...))
	cc.linesShift = append(cc.linesShift, cc.currentShift)
}

func (cc *CodeCreator) AddShift(shift int) {
	cc.currentShift += shift
}

// TODO - remove, use Lines()
func (cc *CodeCreator) Print() []string {
	s := make([]string, len(cc.lines))
	for i, line := range cc.lines {
		s[i] = fmt.Sprintf("%s%s", strings.Repeat(cc.Shift, cc.linesShift[i]), line)
	}
	return s
}

func (cc *CodeCreator) Lines() []string {
	s := make([]string, len(cc.lines))
	for i, line := range cc.lines {
		s[i] = fmt.Sprintf("%s%s", strings.Repeat(cc.Shift, cc.linesShift[i]), line)
	}
	return s
}

func (cc *CodeCreator) Text() string {
	return strings.Join(cc.Lines(), "\n")
}

func (cc *CodeCreator) AddBlock(block func()) {
	cc.AddShift(1)
	block()
	cc.AddShift(-1)
}

func (cc *CodeCreator) addFullBlock(prefix, suffix string, block func()) {
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

type BasicCodeCreator struct {
	CodeCreator
	LanguageBundle
	lastVarIdentifier uint32
}

func (bcc *BasicCodeCreator) Block(prefix string, block func(), suffix string) {
	bcc.addFullBlock(prefix, suffix, block)
}

func (bcc *BasicCodeCreator) If(condition string, block func()) {
	bcc.addFullBlock(fmt.Sprintf(bcc.ifPrefixTemplate, condition), bcc.ifSuffixTemplate, block)
}

func (bcc *BasicCodeCreator) IfElse(condition string, block func(), elseBlock func()) {
	bcc.addFullBlock(
		fmt.Sprintf(bcc.ifPrefixTemplate, condition),
		bcc.elsePrefixTemplate,
		block,
	)
	bcc.addFullBlock("", "", elseBlock)
	bcc.AddLines(bcc.ifSuffixTemplate)
}

func (bcc *BasicCodeCreator) For(initState, condition, iterationStep string, block func()) {
	bcc.addFullBlock(
		fmt.Sprintf(bcc.forPrefixTemplate, initState, condition, iterationStep),
		bcc.forSuffixTemplate,
		block,
	)
}

func (bcc *BasicCodeCreator) ForIndexed(indexVar, startValue, upperBound, step string, block func()) {
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

func (bcc *BasicCodeCreator) Comments(lines ...string) {
	for _, s := range lines {
		bcc.AddLines(bcc.commentPrefix + s)
	}
}

func (bcc *BasicCodeCreator) Comment(lines string) {
	bcc.Comments(strings.Split(lines, "\n")...)
}

func (bcc *BasicCodeCreator) NewVariable(name string) string {
	if name == "" {
		name = "var"
	}
	bcc.lastVarIdentifier += 1
	return fmt.Sprintf(bcc.varTemplate, name, strconv.FormatUint(uint64(bcc.lastVarIdentifier), 10))
}
