// similar to bnf definitions in grammar.TL2.txt

package tlast

// TL2TypeName := (lcName dot)? lcName;
type TL2TypeName struct {
	Namespace string
	Name      string
}

// TL2Annotation := at lcName;
type TL2Annotation struct {
	Name string
	PR   PositionRange
}

// TL2TypeArgument := TL2TypeRef | number;
type TL2TypeArgument struct {
	Type   TL2TypeRef
	Number uint32

	IsNumber bool

	PR PositionRange
}

// TL2TypeApplication := TL2TypeName (lts TL2TypeArgument (cm TL2TypeArgument)* gts)?;
type TL2TypeApplication struct {
	Name      TL2TypeName
	Arguments []TL2TypeArgument

	PR          PositionRange
	PRName      PositionRange
	PRArguments PositionRange
}

// TL2BracketType := lsb TL2TypeArgument? rsb TL2TypeRef;
type TL2BracketType struct {
	IndexType *TL2TypeArgument // nil means that it is vector
	ArrayType TL2TypeRef

	PR PositionRange
}

// TL2TypeRef := TL2TypeApplication | TL2BracketType;
type TL2TypeRef struct {
	SomeType    *TL2TypeApplication
	BracketType *TL2BracketType

	IsBracket bool

	PR PositionRange
}

// TL2Field := ((lcName qm?) | ucs) cl TL2TypeRef;
type TL2Field struct {
	Name       string
	IsOptional bool
	IsIgnored  bool
	Type       TL2TypeRef

	PR     PositionRange
	PRName PositionRange

	CommentBefore string
}

// TL2TypeDefinition = TL2TypeRef | TL2Field* | TL2UnionType;
type TL2TypeDefinition struct {
	TypeAlias         TL2TypeRef
	ConstructorFields []TL2Field
	UnionType         TL2UnionType

	IsConstructorFields bool
	IsUnionType         bool

	PR PositionRange
}

// TL2UnionConstructor := ucName (TL2TypeRef | TL2Field*);
type TL2UnionConstructor struct {
	Name string

	Fields    []TL2Field
	TypeAlias TL2TypeRef

	IsTypeAlias bool

	PR     PositionRange
	PRName PositionRange

	CommentBefore string
}

// TL2UnionType := vb? TL2UnionConstructor (vb TL2UnionConstructor)+;
type TL2UnionType struct {
	Variants []TL2UnionConstructor // at least 2

	PR PositionRange
}

type TL2TypeCategory string

// TL2TypeTemplate := lcName cl lcName;
type TL2TypeTemplate struct {
	Name     string
	Category TL2TypeCategory

	PR         PositionRange
	PRName     PositionRange
	PRCategory PositionRange
}

// TL2TypeDeclaration := TL2TypeName (lts TL2TypeArgumentDeclaration (cm TL2TypeArgumentDeclaration)* gts)? CRC32? eq TL2TypeDefinition?;
type TL2TypeDeclaration struct {
	Name              TL2TypeName
	ID                *uint32
	TemplateArguments []TL2TypeTemplate
	Type              TL2TypeDefinition

	PR     PositionRange
	PRName PositionRange
	PRID   PositionRange
}

// TL2FuncDeclaration := TL2TypeName CRC32 TL2Field* funEq TL2TypeDefinition?;
type TL2FuncDeclaration struct {
	Name       TL2TypeName
	ID         *uint32
	Arguments  []TL2Field
	ReturnType TL2TypeDefinition

	PR     PositionRange
	PRName PositionRange
	PRID   PositionRange
}

// TL2Combinator := TL2Annotation* (TL2TypeDeclaration | TL2FuncDeclaration) scl;
type TL2Combinator struct {
	Annotations []TL2Annotation

	TypeDecl TL2TypeDeclaration
	FuncDecl TL2FuncDeclaration

	IsFunction bool

	PR PositionRange

	CommentBefore string
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

func (t TL2UnionConstructor) HasBeforeCommentIn() bool {
	if t.CommentBefore != "" {
		return true
	}
	if !t.IsTypeAlias {
		for _, field := range t.Fields {
			if field.CommentBefore != "" {
				return true
			}
		}
	}
	return false
}

func (c TL2Combinator) HasAnnotation(value string) bool {
	for _, tl2Annotation := range c.Annotations {
		if tl2Annotation.Name == value {
			return true
		}
	}
	return false
}

func (c TL2Combinator) ReferenceName() TL2TypeName {
	if c.IsFunction {
		return c.FuncDecl.Name
	}
	return c.TypeDecl.Name
}

func (c TL2Combinator) ReferenceNamePR() PositionRange {
	if c.IsFunction {
		return c.FuncDecl.PRName
	}
	return c.TypeDecl.PRName
}
