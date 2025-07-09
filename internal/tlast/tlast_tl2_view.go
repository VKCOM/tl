package tlast

import (
	"fmt"
	"strconv"
	"strings"
)

const OneLineConstructorSize = 120
const UnionConstructorSize = 80

func (t TL2File) String() string {
	sb := strings.Builder{}
	t.Print(&sb)
	return sb.String()
}

func (t TL2Combinator) String() string {
	sb := strings.Builder{}
	t.Print(&sb)
	return sb.String()
}

func (t TL2TypeRef) String() string {
	sb := strings.Builder{}
	t.Print(&sb)
	return sb.String()
}

func (t TL2File) Print(sb *strings.Builder) {
	for _, combinator := range t.Combinators {
		combinator.Print(sb)
		sb.WriteString("\n")
	}
}

func (t *TL2Combinator) Print(sb *strings.Builder) {
	start := sb.Len()
	if len(t.CommentBefore) != 0 {
		lines := strings.Split(t.CommentBefore, "\n")
		for _, line := range lines {
			sb.WriteString(strings.TrimSpace(line))
			sb.WriteString("\n")
		}
		start = sb.Len()
	}
	for _, ann := range t.Annotations {
		sb.WriteString("@")
		sb.WriteString(ann.Name)
		sb.WriteString(" ")
	}
	if t.IsFunction {
		t.FuncDecl.print(sb, sb.Len()-start)
	} else {
		t.TypeDecl.print(sb, sb.Len()-start)
	}
	sb.WriteString(";")
}

func (t *TL2FuncDeclaration) print(sb *strings.Builder, prefixSize int) {
	if t.Name.String() == "s3meta.signText" {
		print("debug")
	}
	const defaultSep = " "
	const newLineSep = "\n\t"

	sep := defaultSep
	tmpSb := strings.Builder{}
	hasNewLines := t.printFunction(&tmpSb, sep, prefixSize)
	if hasNewLines {
		sep = newLineSep
	}
	t.printFunction(sb, sep, prefixSize)
}

func (t *TL2FuncDeclaration) printArguments(sb *strings.Builder, sep string) {
	for _, argument := range t.Arguments {
		sb.WriteString(sep)
		argument.Print(sb)
	}
}

func (t *TL2FuncDeclaration) printFunction(sb *strings.Builder, sep string, prefixSize int) (hasNewLines bool) {
	startSize := sb.Len()
	sb.WriteString(t.Name.String())
	if t.ID != nil {
		sb.WriteString(fmt.Sprintf("#%08x", *t.ID))
	}
	t.printArguments(sb, sep)
	sb.WriteString(sep)
	sb.WriteString("=> ")
	hasNewLines = t.ReturnType.print(sb, sb.Len()-startSize+prefixSize)
	if hasNewLines || sb.Len()-startSize+prefixSize > OneLineConstructorSize {
		hasNewLines = true
	}
	return
}

func (t *TL2TypeDeclaration) print(sb *strings.Builder, prefixSize int) {
	if t.Name.String() == "exactlyOnce.ackResponse" {
		print("debug")
	}
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
	t.Type.print(sb, prefixSize+sb.Len()-startSize)
}

func (t TL2TypeDefinition) print(sb *strings.Builder, definitionPrefix int) (hasNewLines bool) {
	tmpSb := strings.Builder{}
	hasNewLines = t.printWithNewLineOption(&tmpSb, false)
	if hasNewLines || tmpSb.Len()+definitionPrefix > OneLineConstructorSize {
		hasNewLines = true
	}
	t.printWithNewLineOption(sb, hasNewLines)
	return hasNewLines
}

func (t TL2TypeDefinition) printWithNewLineOption(sb *strings.Builder, forceNewline bool) (hasNewLines bool) {
	const defaultSep = " "
	const newLineSep = "\n\t"

	if t.IsUnionType {
		hasComments := false
		for _, variant := range t.UnionType.Variants {
			hasComments = hasComments || variant.HasBeforeCommentIn()
			if hasComments {
				break
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
			variantForceNewLine := variant.print(sb, len(sep))
			forceNewline = forceNewline || variantForceNewLine
		}
	} else if t.IsConstructorFields {
		hasComments := false
		for _, field := range t.ConstructorFields {
			hasComments = hasComments || field.CommentBefore != ""
			if hasComments {
				break
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
			if field.CommentBefore != "" {
				lines := strings.Split(field.CommentBefore, "\n")
				for _, line := range lines {
					sb.WriteString(strings.TrimSpace(line))
					sb.WriteString(sep)
				}
			}
			field.Print(sb)
		}
	} else {
		t.TypeAlias.Print(sb)
	}
	return forceNewline
}

func (t *TL2UnionTypeVariant) print(sb *strings.Builder, prefixSize int) (hasNewLine bool) {
	const defaultSep = " "
	const newLineSep = "\n\t\t"

	forceNewLine := false
	if t.IsConstructor {
		currentSize := sb.Len()
		sb.WriteString(t.Constructor.Name)
		if t.HasBeforeCommentIn() {
			forceNewLine = true
		}
		sep := defaultSep
		if forceNewLine {
			sep = newLineSep
		}
		if !forceNewLine {
			tmp := strings.Builder{}
			t.printVariantFields(&tmp, sep)
			if prefixSize+(sb.Len()-currentSize)+tmp.Len() > UnionConstructorSize {
				t.printVariantFields(sb, newLineSep)
			} else {
				t.printVariantFields(sb, defaultSep)
			}
		} else {
			t.printVariantFields(sb, sep)
		}
	} else {
		t.TypeAlias.Print(sb)
	}
	return forceNewLine
}

func (t *TL2UnionTypeVariant) printVariantFields(sb *strings.Builder, sep string) {
	for _, field := range t.Constructor.Fields {
		sb.WriteString(sep)
		if field.CommentBefore != "" {
			lines := strings.Split(field.CommentBefore, "\n")
			for _, line := range lines {
				sb.WriteString(strings.TrimSpace(line))
				sb.WriteString(sep)
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
