// similar to bnf definitions in grammar.tl2.txt

package tlast

type TL2TypeName struct {
	Namespace string
	Name      string
}

type TL2Annotation struct {
	Name string
	PR   PositionRange
}

type TL2TypeArgument struct {
	Type   TL2TypeRef
	Number uint32

	IsNumber bool

	PR PositionRange
}

type TL2TypeApplication struct {
	Name      TL2TypeName
	Arguments []TL2TypeArgument

	PR          PositionRange
	PRName      PositionRange
	PRArguments PositionRange
}

type TL2BracketType struct {
	IndexType *TL2TypeArgument // nil means that it is vector
	ArrayType TL2TypeRef

	PR PositionRange
}

type TL2TypeRef struct {
	SomeType    *TL2TypeApplication
	BracketType *TL2BracketType

	IsBracket bool

	PR PositionRange
}

type TL2Field struct {
	Name       string
	IsOptional bool
	IsIgnored  bool
	Type       TL2TypeRef

	PR     PositionRange
	PRName PositionRange
}

type TL2TypeDefinition struct {
	TypeAlias         TL2TypeRef
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
	TypeAlias   TL2TypeRef
	Constructor TL2UnionConstructor

	IsConstructor bool

	PR PositionRange
}

type TL2UnionType struct {
	Variants []TL2UnionTypeVariant // at least 2

	PR PositionRange
}

type TL2TypeCategory string

type TL2TypeTemplate struct {
	Name     string
	Category TL2TypeCategory

	PR         PositionRange
	PRName     PositionRange
	PRCategory PositionRange
}

type TL2TypeDeclaration struct {
	Name              TL2TypeName
	ID                *uint32
	TemplateArguments []TL2TypeTemplate
	Type              TL2TypeDefinition

	PR     PositionRange
	PRName PositionRange
	PRID   PositionRange
}

type TL2FuncDeclaration struct {
	Name       TL2TypeName
	ID         *uint32
	Arguments  []TL2Field
	ReturnType TL2TypeDefinition

	PR     PositionRange
	PRName PositionRange
	PRID   PositionRange
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

func (c TL2TypeCategory) IsType() bool {
	return c == "type"
}

func (c TL2TypeCategory) IsUint32() bool {
	return c == "uint32"
}

func (c TL2TypeCategory) IsLegalCategory() bool {
	return c.IsUint32() || c.IsType()
}
