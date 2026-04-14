package codecreator

import (
	"fmt"
	"strconv"
	"strings"
)

type BasicCodeCreator struct {
	shift string
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

	lastVarIdentifier uint32

	lines      []string
	linesShift []int

	currentShift int
}

func (bcc *BasicCodeCreator) AddLines(lines ...string) {
	bcc.lines = append(bcc.lines, lines...)
	for range lines {
		bcc.linesShift = append(bcc.linesShift, bcc.currentShift)
	}
}

func (bcc *BasicCodeCreator) AddEmptyLine() {
	bcc.lines = append(bcc.lines, "")
	bcc.linesShift = append(bcc.linesShift, bcc.currentShift)
}

func (bcc *BasicCodeCreator) AddFullEmptyLine() {
	bcc.lines = append(bcc.lines, "")
	bcc.linesShift = append(bcc.linesShift, 0)
}

func (bcc *BasicCodeCreator) AddLinef(format string, args ...any) {
	line := fmt.Sprintf(format, args...)
	if line == "" {
		return
	}
	bcc.lines = append(bcc.lines, line)
	bcc.linesShift = append(bcc.linesShift, bcc.currentShift)
}

func (bcc *BasicCodeCreator) AddShift(shift int) {
	bcc.currentShift += shift
}

// TODO - remove, use Lines()
func (bcc *BasicCodeCreator) Print() []string {
	return bcc.Lines()
}

func (bcc *BasicCodeCreator) Lines() []string {
	s := make([]string, len(bcc.lines))
	for i, line := range bcc.lines {
		s[i] = fmt.Sprintf("%s%s", strings.Repeat(bcc.shift, bcc.linesShift[i]), line)
	}
	return s
}

func (bcc *BasicCodeCreator) Text() string {
	b := strings.Builder{}
	for _, line := range bcc.Lines() {
		b.WriteString(line)
		b.WriteString("\n")
	}
	return b.String()
}

func (bcc *BasicCodeCreator) AddBlock(block func()) {
	bcc.AddShift(1)
	block()
	bcc.AddShift(-1)
}

func (bcc *BasicCodeCreator) addFullBlock(prefix, suffix string, block func()) {
	if prefix != "" {
		bcc.AddLines(prefix)
	}
	bcc.AddBlock(block)
	if suffix != "" {
		bcc.AddLines(suffix)
	}
}

func (bcc *BasicCodeCreator) Block(prefix string, block func(), suffix string) {
	bcc.addFullBlock(prefix, suffix, block)
}

func (bcc *BasicCodeCreator) FinishBlock(block func(), suffix string) {
	bcc.addFullBlock("", suffix, block)
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

func (bcc *BasicCodeCreator) forIndexed(indexVar, startValue, upperBound, step string, block func()) {
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
