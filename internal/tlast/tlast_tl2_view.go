package tlast

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const OneLineConstructorSize = 120
const UnionConstructorSize = 80

type formatOptions struct {
	IgnoreComments         bool
	OneLineConstructorSize int
	UnionConstructorSize   int
}

func NewDefaultFormatOptions() formatOptions {
	return formatOptions{
		IgnoreComments:         false,
		OneLineConstructorSize: OneLineConstructorSize,
		UnionConstructorSize:   UnionConstructorSize,
	}
}

func NewCanonicalFormatOptions() formatOptions {
	return formatOptions{
		IgnoreComments:         true,
		OneLineConstructorSize: math.MaxInt32 - 10000,
		UnionConstructorSize:   math.MaxInt32 - 10000,
	}
}

func (t TL2File) String() string {
	sb := strings.Builder{}
	t.Print(&sb, NewDefaultFormatOptions())
	return sb.String()
}

func (t TL2Combinator) String() string {
	sb := strings.Builder{}
	t.Print(&sb, NewDefaultFormatOptions())
	return sb.String()
}

func (t TL2TypeRef) String() string {
	sb := strings.Builder{}
	t.Print(&sb)
	return sb.String()
}

func (t TL2File) Print(sb *strings.Builder, options formatOptions) {
	for _, combinator := range t.Combinators {
		combinator.Print(sb, options)
		sb.WriteString("\n")
	}
}

func (t *TL2Combinator) Print(sb *strings.Builder, options formatOptions) {
	start := sb.Len()
	if !options.IgnoreComments {
		if len(t.CommentBefore) != 0 {
			lines := strings.Split(t.CommentBefore, "\n")
			for _, line := range lines {
				sb.WriteString(strings.TrimSpace(line))
				sb.WriteString("\n")
			}
			start = sb.Len()
		}
	}
	for _, ann := range t.Annotations {
		sb.WriteString("@")
		sb.WriteString(ann.Name)
		sb.WriteString(" ")
	}
	if t.IsFunction {
		t.FuncDecl.print(sb, options, sb.Len()-start)
	} else {
		t.TypeDecl.print(sb, options, sb.Len()-start)
	}
	sb.WriteString(";")
}

func (t *TL2FuncDeclaration) print(sb *strings.Builder, options formatOptions, prefixSize int) {
	const defaultSep = " "
	const newLineSep = "\n\t"

	sep := defaultSep
	tmpSb := strings.Builder{}
	hasNewLines := t.printFunction(&tmpSb, options, sep, prefixSize)
	if hasNewLines {
		sep = newLineSep
	}
	t.printFunction(sb, options, sep, prefixSize)
}

func (t *TL2FuncDeclaration) printArguments(sb *strings.Builder, sep string) {
	for _, argument := range t.Arguments {
		sb.WriteString(sep)
		argument.Print(sb)
	}
}

func (t *TL2FuncDeclaration) printFunction(sb *strings.Builder, options formatOptions, sep string, prefixSize int) (hasNewLines bool) {
	startSize := sb.Len()
	sb.WriteString(t.Name.String())
	if t.ID != nil {
		sb.WriteString(fmt.Sprintf("#%08x", *t.ID))
	}
	t.printArguments(sb, sep)
	sb.WriteString(sep)
	sb.WriteString("=> ")
	hasNewLines = t.ReturnType.print(sb, options, sb.Len()-startSize+prefixSize)
	if hasNewLines || sb.Len()-startSize+prefixSize > options.OneLineConstructorSize {
		hasNewLines = true
	}
	return
}

func (t *TL2TypeDeclaration) print(sb *strings.Builder, options formatOptions, prefixSize int) {
	startSize := sb.Len()
	sb.WriteString(t.Name.String())
	if len(t.TemplateArguments) > 0 {
		sb.WriteString("<")
		for i, argument := range t.TemplateArguments {
			if i != 0 {
				sb.WriteString(",")
			}
			sb.WriteString(argument.Name)
			sb.WriteString(":")
			sb.WriteString(string(argument.Category))
		}
		sb.WriteString(">")
	}
	if t.ID != nil {
		sb.WriteString(fmt.Sprintf("#%08x", *t.ID))
	}
	sb.WriteString(" = ")
	t.Type.print(sb, options, prefixSize+sb.Len()-startSize)
}

func (t TL2TypeDefinition) print(sb *strings.Builder, options formatOptions, definitionPrefix int) (hasNewLines bool) {
	tmpSb := strings.Builder{}
	hasNewLines = t.printWithNewLineOption(&tmpSb, options, false)
	if hasNewLines || tmpSb.Len()+definitionPrefix > options.OneLineConstructorSize {
		hasNewLines = true
	}
	t.printWithNewLineOption(sb, options, hasNewLines)
	return hasNewLines
}

func (t TL2TypeDefinition) printWithNewLineOption(sb *strings.Builder, options formatOptions, forceNewline bool) (hasNewLines bool) {
	const defaultSep = " "
	const newLineSep = "\n\t"

	if t.IsUnionType {
		hasComments := false
		if !options.IgnoreComments {
			for _, variant := range t.UnionType.Variants {
				hasComments = hasComments || variant.HasBeforeCommentIn()
				if hasComments {
					break
				}
			}
		}
		forceNewline = forceNewline || hasComments
		sep := defaultSep + "| "
		if forceNewline {
			sep = newLineSep + "| "
		}
		for i, variant := range t.UnionType.Variants {
			if i != 0 || forceNewline {
				sb.WriteString(sep)
			}
			variantForceNewLine := variant.print(sb, options, len(sep))
			forceNewline = forceNewline || variantForceNewLine
		}
	} else if t.IsConstructorFields {
		hasComments := false
		if !options.IgnoreComments {
			for _, field := range t.ConstructorFields {
				hasComments = hasComments || field.CommentBefore != ""
				if hasComments {
					break
				}
			}
		}
		forceNewline = forceNewline || hasComments
		sep := defaultSep
		if forceNewline {
			sep = newLineSep
		}
		for i, field := range t.ConstructorFields {
			if i != 0 || forceNewline {
				sb.WriteString(sep)
			}
			if !options.IgnoreComments {
				if field.CommentBefore != "" {
					lines := strings.Split(field.CommentBefore, "\n")
					for _, line := range lines {
						sb.WriteString(strings.TrimSpace(line))
						sb.WriteString(sep)
					}
				}
			}
			field.Print(sb)
		}
	} else {
		t.TypeAlias.Print(sb)
	}
	return forceNewline
}

func (t *TL2UnionConstructor) print(sb *strings.Builder, options formatOptions, prefixSize int) (hasNewLine bool) {
	const defaultSep = " "
	const newLineSep = "\n\t\t"

	currentSize := sb.Len()
	sb.WriteString(t.Name)

	forceNewLine := false
	if !t.IsTypeAlias {
		if !options.IgnoreComments {
			if t.HasBeforeCommentIn() {
				forceNewLine = true
			}
		}
		sep := defaultSep
		if forceNewLine {
			sep = newLineSep
		}
		if !forceNewLine {
			tmp := strings.Builder{}
			t.printVariantFields(&tmp, options, sep)
			if prefixSize+(sb.Len()-currentSize)+tmp.Len() > options.UnionConstructorSize {
				t.printVariantFields(sb, options, newLineSep)
			} else {
				t.printVariantFields(sb, options, defaultSep)
			}
		} else {
			t.printVariantFields(sb, options, sep)
		}
	} else {
		sb.WriteString(" ")
		t.TypeAlias.Print(sb)
	}
	return forceNewLine
}

func (t *TL2UnionConstructor) printVariantFields(sb *strings.Builder, options formatOptions, sep string) {
	for _, field := range t.Fields {
		sb.WriteString(sep)
		if !options.IgnoreComments {
			if field.CommentBefore != "" {
				lines := strings.Split(field.CommentBefore, "\n")
				for _, line := range lines {
					sb.WriteString(strings.TrimSpace(line))
					sb.WriteString(sep)
				}
			}
		}
		field.Print(sb)
	}
}

func (t *TL2Field) Print(sb *strings.Builder) {
	if t.IsIgnored {
		sb.WriteString("_")
	} else {
		sb.WriteString(t.Name)
	}
	if t.IsOptional {
		sb.WriteString("?")
	}
	sb.WriteString(":")
	t.Type.Print(sb)
}

func (t *TL2TypeRef) Print(sb *strings.Builder) {
	if t.IsBracket {
		if t.BracketType != nil {
			t.BracketType.Print(sb)
		}
	} else {
		if t.SomeType != nil {
			t.SomeType.Print(sb)
		}
	}
}

func (t *TL2BracketType) Print(sb *strings.Builder) {
	sb.WriteString("[")
	if t.IndexType != nil {
		t.IndexType.Print(sb)
	}
	sb.WriteString("]")
	t.ArrayType.Print(sb)
}

func (t *TL2TypeApplication) Print(sb *strings.Builder) {
	sb.WriteString(t.Name.String())
	if len(t.Arguments) > 0 {
		sb.WriteString("<")
		for i, argument := range t.Arguments {
			if i != 0 {
				sb.WriteString(",")
			}
			argument.Print(sb)
		}
		sb.WriteString(">")
	}
}

func (t *TL2TypeArgument) Print(sb *strings.Builder) {
	if t.IsNumber {
		sb.WriteString(strconv.FormatUint(uint64(t.Number), 10))
	} else {
		t.Type.Print(sb)
	}
}
