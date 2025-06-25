// similar to bnf definitions in grammar.tl2.txt

package tlast

type TL2TypeName struct {
	Namespace string
	Name      string

	PR PositionRange
}

type TL2Annotation struct {
	Name string
	PR   PositionRange
}

type TL2TypeArgument struct {
	Type   TL2Type
	Number uint32

	IsNumber bool

	PR PositionRange
}

type TL2TypeApplication struct {
	Name      TL2TypeName
	Arguments []TL2TypeArgument

	PR          PositionRange
	PRArguments PositionRange
}

type TL2BracketType struct {
	IndexType *TL2TypeArgument // nil means that it is vector
	ArrayType TL2Type

	PR PositionRange
}

type TL2Type struct {
	SomeType    *TL2TypeApplication
	BracketType *TL2BracketType

	IsBracket bool

	PR PositionRange
}

type TL2Field struct {
	Name       string
	IsOptional bool
	IsIgnored  bool
	Type       TL2Type

	PR     PositionRange
	PRName PositionRange
}

type TL2TypeDefinition struct {
	TypeAlias         TL2Type
	ConstructorFields []TL2Field
	UnionType         TL2UnionType

	IsConstructorFields bool
	IsUnionType         bool

	PR PositionRange
}

type TL2UnionConstructor struct {
	Name   string
	Fields []TL2Field

	PR     PositionRange
	PRName PositionRange
}

type TL2UnionTypeVariant struct {
	TypeAlias   TL2Type
	Constructor TL2UnionConstructor

	IsConstructor bool

	PR PositionRange
}

type TL2UnionType struct {
	Variants []TL2UnionTypeVariant // at least 2

	PR PositionRange
}

type TL2TypeTemplate struct {
	Name     string
	Category string

	PR         PositionRange
	PRName     PositionRange
	PRCategory PositionRange
}

type TL2TypeDeclaration struct {
	Name              TL2TypeName
	ID                *uint32
	TemplateArguments []TL2TypeTemplate
	Type              TL2TypeDefinition

	PR   PositionRange
	PRID PositionRange
}

type TL2FuncDeclaration struct {
	Name       TL2TypeName
	ID         *uint32
	Arguments  []TL2Field
	ReturnType TL2TypeDefinition

	PR   PositionRange
	PRID PositionRange
}

type TL2Combinator struct {
	Annotations []TL2Annotation

	TypeDecl TL2TypeDeclaration
	FuncDecl TL2FuncDeclaration

	IsFunction bool

	PR PositionRange
}

// TL2File := TL2Combinator*;
type TL2File struct {
	Combinators []TL2Combinator
}

// trivial methods

func (t TL2TypeName) String() string {
	prefix := ""
	if t.Namespace != "" {
		prefix = t.Namespace + "."
	}
	return prefix + t.Name
}

func (t TL2TypeDefinition) IsAlias() bool {
	return !t.IsUnionType && !t.IsConstructorFields
}
