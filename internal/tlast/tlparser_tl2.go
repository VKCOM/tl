// similar to bnf definitions in grammar.tl2.txt

package tlast

// TL2TypeName := (lcName dot)? lcName;
type TL2TypeName struct {
	Namespace string
	Name      string

	PR PositionRange
}

// TL2Annotation := at lcName;
type TL2Annotation struct {
	Name string
	PR   PositionRange
}

// TL2TypeArgument := TL2Type | number;
type TL2TypeArgument struct {
	Type   TL2Type
	Number uint32

	IsNumber bool

	PR PositionRange
}

// TL2TypeApplication := TL2TypeName (lts TL2TypeArgument (cm TL2TypeArgument)* gts)?;
type TL2TypeApplication struct {
	Name      TL2TypeName
	Arguments []TL2TypeArgument

	PR          PositionRange
	PRArguments PositionRange
}

// TL2BracketType := lsb TL2TypeArgument? rsb TL2Type;
type TL2BracketType struct {
	IndexType TL2TypeArgument
	ArrayType TL2Type

	PR PositionRange
}

// TL2Type := TL2TypeApplication | TL2BracketType;
type TL2Type struct {
	SomeType    *TL2TypeApplication
	BracketType *TL2BracketType

	IsBracket bool

	PR PositionRange
}

// TL2Field := lcName qm? cl TL2Type;
type TL2Field struct {
	Name       string
	IsOptional bool
	Type       TL2Type

	PR     PositionRange
	PRName PositionRange
}

// TL2TypeDefinition = TL2Type | TL2Field* | TL2UnionType;
type TL2TypeDefinition struct {
	TypeDef           TL2Type
	ConstructorFields []TL2Field
	UnionType         TL2UnionType

	IsConstructorFields bool
	IsUnionType         bool

	PR PositionRange
}

// TL2UnionConstructor := ucName TL2Field*;
type TL2UnionConstructor struct {
	Name   string
	Fields []TL2Field

	PR     PositionRange
	PRName PositionRange
}

// TL2UnionTypeVariant := TL2Type | TL2UnionConstructor;
type TL2UnionTypeVariant struct {
	TypeDef     TL2Type
	Constructor TL2UnionConstructor

	IsConstructor bool

	PR PositionRange
}

// TL2UnionType := TL2UnionTypeVariant (vb TL2UnionTypeVariant)+;
type TL2UnionType struct {
	Variants []TL2UnionTypeVariant // at least 2

	PR PositionRange
}

// TL2TypeArgumentDeclaration := lcName cl lcName;
type TL2TypeArgumentDeclaration struct {
	Name     string
	Category string

	PR         PositionRange
	PRName     PositionRange
	PRCategory PositionRange
}

// TL2TypeDeclaration := TL2TypeName (lts TL2TypeArgumentDeclaration (cm TL2TypeArgumentDeclaration)* gts)? eq TL2TypeDefinition;
type TL2TypeDeclaration struct {
	Name              TL2TypeName
	TemplateArguments []TL2TypeArgumentDeclaration
	Type              TL2TypeDefinition

	PR PositionRange
}

// TL2FuncDeclaration := TL2TypeName TL2Field* funEq TL2TypeDefinition;
type TL2FuncDeclaration struct {
	Name       TL2TypeName
	Arguments  []TL2Field
	ReturnType TL2TypeDefinition

	PR PositionRange
}

// TL2Combinator := TL2Annotation* (TL2TypeDeclaration | TL2FuncDeclaration) scl;
type TL2Combinator struct {
	Annotations []TL2Annotation

	TypeDecl TL2TypeDeclaration
	FuncDecl TL2FuncDeclaration

	IsFunction bool

	PR PositionRange
}
