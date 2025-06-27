package tlast

import (
	"fmt"
	"strconv"
	"strings"
)

func (t *TL2File) String() string {
	sb := strings.Builder{}
	t.Print(&sb)
	return sb.String()
}

func (t *TL2Combinator) String() string {
	sb := strings.Builder{}
	t.Print(&sb)
	return sb.String()
}

func (t *TL2TypeRef) String() string {
	sb := strings.Builder{}
	t.Print(&sb)
	return sb.String()
}

func (t *TL2File) Print(sb *strings.Builder) {
	for _, combinator := range t.Combinators {
		combinator.Print(sb)
		sb.WriteString("\n")
	}
}

func (t *TL2Combinator) Print(sb *strings.Builder) {
	for _, ann := range t.Annotations {
		sb.WriteString("@")
		sb.WriteString(ann.Name)
		sb.WriteString(" ")
	}
	if t.IsFunction {
		t.FuncDecl.Print(sb)
	} else {
		t.TypeDecl.Print(sb)
	}
	sb.WriteString(";")
}

func (t *TL2FuncDeclaration) Print(sb *strings.Builder) {
	sb.WriteString(t.Name.String())
	if t.ID != nil {
		sb.WriteString(fmt.Sprintf("#%08x", *t.ID))
	}
	for _, argument := range t.Arguments {
		sb.WriteString(" ")
		argument.Print(sb)
	}
	sb.WriteString(" => ")
	t.ReturnType.Print(sb)
}

func (t *TL2TypeDeclaration) Print(sb *strings.Builder) {
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
	t.Type.Print(sb)
}

func (t *TL2TypeDefinition) Print(sb *strings.Builder) {
	if t.IsUnionType {
		for i, variant := range t.UnionType.Variants {
			if i != 0 {
				sb.WriteString(" | ")
			}
			variant.Print(sb)
		}
	} else if t.IsConstructorFields {
		for i, field := range t.ConstructorFields {
			if i != 0 {
				sb.WriteString(" ")
			}
			field.Print(sb)
		}
	} else {
		t.TypeAlias.Print(sb)
	}
}

func (t *TL2UnionTypeVariant) Print(sb *strings.Builder) {
	if t.IsConstructor {
		sb.WriteString(t.Constructor.Name)
		for _, field := range t.Constructor.Fields {
			sb.WriteString(" ")
			field.Print(sb)
		}
	} else {
		t.TypeAlias.Print(sb)
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
		t.BracketType.Print(sb)
	} else {
		t.SomeType.Print(sb)
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
