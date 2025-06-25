package tlast

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func setupIterator(str string) (error, tokenIterator) {
	lex := newLexer(str, "", LexerOptions{LexerLanguage: tl2})
	allTokens, err := lex.generateTokens()

	return err, tokenIterator{tokens: allTokens}
}

func TestParseTL2(t *testing.T) {
	t.Run("optional parse annotations", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			_, it := setupIterator(`   @gigi    `)
			state, newIt, ann := parseTL2Annotation(it, Position{})
			assert.True(t, state.StartProcessing)
			assert.NoError(t, state.Error)
			assert.Equal(t, "gigi", ann.Name)
			assert.True(t, newIt.expect(eof))
		})

		t.Run("don't start processing", func(t *testing.T) {
			_, it := setupIterator(`   #11111111    `)
			state, newIt, ann := parseTL2Annotation(it, Position{})
			assert.False(t, state.StartProcessing)
			assert.NoError(t, state.Error)
			assert.Equal(t, "", ann.Name)
			assert.True(t, newIt.expect(crc32hash))
		})

		t.Run("parse several", func(t *testing.T) {
			_, it := setupIterator(` @gigi @gaga #12345678  `)
			state, newIt, anns := zeroOrMore(parseTL2Annotation)(it, Position{})
			assert.True(t, state.StartProcessing)
			assert.NoError(t, state.Error)
			assert.Equal(t, 2, len(anns))
			assert.Equal(t, "gigi", anns[0].Name)
			assert.Equal(t, "gaga", anns[1].Name)
			assert.True(t, newIt.expect(crc32hash))
		})

		t.Run("parse zero", func(t *testing.T) {
			_, it := setupIterator(` #12345678  `)
			state, newIt, anns := zeroOrMore(parseTL2Annotation)(it, Position{})
			assert.False(t, state.StartProcessing)
			assert.NoError(t, state.Error)
			assert.Equal(t, 0, len(anns))
			assert.True(t, newIt.expect(crc32hash))
		})
	})

	t.Run("optional parse type name", func(t *testing.T) {
		t.Run("success without namespace", func(t *testing.T) {
			_, it := setupIterator(` testName  `)
			state, newIt, testName := parseTL2TypeName(it, Position{})
			assert.True(t, state.StartProcessing)
			assert.NoError(t, state.Error)
			assert.Equal(t, "testName", testName.Name)
			assert.True(t, newIt.expect(eof))
		})
		t.Run("success with namespace", func(t *testing.T) {
			_, it := setupIterator(` testNS.testName  `)
			state, newIt, testName := parseTL2TypeName(it, Position{})
			assert.True(t, state.StartProcessing)
			assert.NoError(t, state.Error)
			assert.Equal(t, "testName", testName.Name)
			assert.Equal(t, "testNS", testName.Namespace)
			assert.True(t, newIt.expect(eof))
		})
		t.Run("don't start processing", func(t *testing.T) {
			_, it := setupIterator(` TestName  `)
			state, newIt, _ := parseTL2TypeName(it, Position{})
			assert.False(t, state.StartProcessing)
			assert.NoError(t, state.Error)
			assert.True(t, newIt.expect(ucIdent))
		})
		t.Run("fail because of ucName", func(t *testing.T) {
			_, it := setupIterator(` testNs.TestName  `)
			state, newIt, _ := parseTL2TypeName(it, Position{})
			assert.True(t, state.StartProcessing)
			assert.Error(t, state.Error)
			assert.True(t, newIt.expect(ucIdentNS))
		})
	})

	t.Run("combinator", func(t *testing.T) {
		t.Run("type declaration", func(t *testing.T) {
			t.Run("basic definition", func(t *testing.T) {
				_, it := setupIterator(` testNs.testName = x:int; `)
				comb, newIt, err := parseTL2Combinator(it)
				assert.NoError(t, err)
				assert.True(t, newIt.expect(eof))
				assert.False(t, comb.IsFunction)
				assert.Equal(t, "testNs", comb.TypeDecl.Name.Namespace)
				assert.Equal(t, "testName", comb.TypeDecl.Name.Name)
				assert.True(t, comb.TypeDecl.Type.IsConstructorFields)

				fields := comb.TypeDecl.Type.ConstructorFields

				assert.Equal(t, 1, len(fields))
				assert.Equal(t, "x", fields[0].Name)

				field0Type := fields[0].Type
				assert.False(t, field0Type.IsBracket)

				assert.Equal(t, "", field0Type.SomeType.Name.Namespace)
				assert.Equal(t, "int", field0Type.SomeType.Name.Name)
			})

			t.Run("basic definition with crc32", func(t *testing.T) {
				_, it := setupIterator(` testNs.testName #90abcdef = x:int; `)
				comb, newIt, err := parseTL2Combinator(it)
				assert.NoError(t, err)
				assert.True(t, newIt.expect(eof))
				assert.False(t, comb.IsFunction)
				assert.Equal(t, "testNs", comb.TypeDecl.Name.Namespace)
				assert.Equal(t, "testName", comb.TypeDecl.Name.Name)
				assert.NotNil(t, comb.TypeDecl.ID)
				assert.Equal(t, uint32(0x90abcdef), *comb.TypeDecl.ID)
			})

			t.Run("basic definition with optional field", func(t *testing.T) {
				_, it := setupIterator(` testNs.testName = x?:int; `)
				comb, newIt, err := parseTL2Combinator(it)
				assert.NoError(t, err)
				assert.True(t, newIt.expect(eof))
				assert.False(t, comb.IsFunction)
				assert.Equal(t, "testNs", comb.TypeDecl.Name.Namespace)
				assert.Equal(t, "testName", comb.TypeDecl.Name.Name)
				assert.True(t, comb.TypeDecl.Type.IsConstructorFields)

				fields := comb.TypeDecl.Type.ConstructorFields

				assert.Equal(t, 1, len(fields))
				assert.Equal(t, "x", fields[0].Name)
				assert.True(t, fields[0].IsOptional)

				field0Type := fields[0].Type
				assert.False(t, field0Type.IsBracket)

				assert.Equal(t, "", field0Type.SomeType.Name.Namespace)
				assert.Equal(t, "int", field0Type.SomeType.Name.Name)
			})

			t.Run("basic definition with ignored field", func(t *testing.T) {
				_, it := setupIterator(` testNs.testName = _:int; `)
				comb, newIt, err := parseTL2Combinator(it)
				assert.NoError(t, err)
				assert.True(t, newIt.expect(eof))
				assert.False(t, comb.IsFunction)
				assert.Equal(t, "testNs", comb.TypeDecl.Name.Namespace)
				assert.Equal(t, "testName", comb.TypeDecl.Name.Name)
				assert.True(t, comb.TypeDecl.Type.IsConstructorFields)

				fields := comb.TypeDecl.Type.ConstructorFields

				assert.Equal(t, 1, len(fields))
				assert.Equal(t, "_", fields[0].Name)
				assert.True(t, fields[0].IsIgnored)

				field0Type := fields[0].Type
				assert.False(t, field0Type.IsBracket)

				assert.Equal(t, "", field0Type.SomeType.Name.Namespace)
				assert.Equal(t, "int", field0Type.SomeType.Name.Name)
			})

			t.Run("basic definition with ignored optional field", func(t *testing.T) {
				_, it := setupIterator(` testNs.testName = _?:int; `)
				_, _, err := parseTL2Combinator(it)
				assert.Error(t, err)
			})

			t.Run("empty definition", func(t *testing.T) {
				_, it := setupIterator(` testNs.testName = ; `)
				comb, newIt, err := parseTL2Combinator(it)
				assert.NoError(t, err)
				assert.True(t, newIt.expect(eof))
				assert.False(t, comb.IsFunction)
				assert.Equal(t, "testNs", comb.TypeDecl.Name.Namespace)
				assert.Equal(t, "testName", comb.TypeDecl.Name.Name)
				assert.True(t, comb.TypeDecl.Type.IsConstructorFields)

				fields := comb.TypeDecl.Type.ConstructorFields

				assert.Equal(t, 0, len(fields))
			})

			t.Run("type alias", func(t *testing.T) {
				_, it := setupIterator(` testNs.testName = int; `)
				comb, newIt, err := parseTL2Combinator(it)
				assert.NoError(t, err)
				assert.True(t, newIt.expect(eof))
				assert.False(t, comb.IsFunction)
				assert.Equal(t, "testNs", comb.TypeDecl.Name.Namespace)
				assert.Equal(t, "testName", comb.TypeDecl.Name.Name)

				assert.False(t, comb.TypeDecl.Type.IsConstructorFields)
				assert.False(t, comb.TypeDecl.Type.IsUnionType)

				alias := comb.TypeDecl.Type.TypeAlias

				assert.False(t, alias.IsBracket)
				assert.Equal(t, "", alias.SomeType.Name.Namespace)
				assert.Equal(t, "int", alias.SomeType.Name.Name)
			})

			t.Run("union", func(t *testing.T) {
				_, it := setupIterator(` testNs.testName = int | string; `)
				comb, newIt, err := parseTL2Combinator(it)
				assert.NoError(t, err)
				assert.True(t, newIt.expect(eof))
				assert.False(t, comb.IsFunction)
				assert.Equal(t, "testNs", comb.TypeDecl.Name.Namespace)
				assert.Equal(t, "testName", comb.TypeDecl.Name.Name)

				assert.False(t, comb.TypeDecl.Type.IsConstructorFields)
				assert.True(t, comb.TypeDecl.Type.IsUnionType)

				union := comb.TypeDecl.Type.UnionType

				assert.Equal(t, 2, len(union.Variants))
				assert.Equal(t, "", union.Variants[0].TypeAlias.SomeType.Name.Namespace)
				assert.Equal(t, "int", union.Variants[0].TypeAlias.SomeType.Name.Name)

				assert.Equal(t, "", union.Variants[1].TypeAlias.SomeType.Name.Namespace)
				assert.Equal(t, "string", union.Variants[1].TypeAlias.SomeType.Name.Name)
			})

			t.Run("union with constructors", func(t *testing.T) {
				_, it := setupIterator(` testNs.testName = Green x:int | Red | string; `)
				comb, newIt, err := parseTL2Combinator(it)
				assert.NoError(t, err)
				assert.True(t, newIt.expect(eof))
				assert.False(t, comb.IsFunction)
				assert.Equal(t, "testNs", comb.TypeDecl.Name.Namespace)
				assert.Equal(t, "testName", comb.TypeDecl.Name.Name)

				assert.False(t, comb.TypeDecl.Type.IsConstructorFields)
				assert.True(t, comb.TypeDecl.Type.IsUnionType)

				union := comb.TypeDecl.Type.UnionType

				assert.Equal(t, 3, len(union.Variants))

				// variant 0
				assert.True(t, union.Variants[0].IsConstructor)
				assert.Equal(t, "Green", union.Variants[0].Constructor.Name)
				assert.Equal(t, 1, len(union.Variants[0].Constructor.Fields))

				assert.Equal(t, "", union.Variants[0].Constructor.Fields[0].Type.SomeType.Name.Namespace)
				assert.Equal(t, "int", union.Variants[0].Constructor.Fields[0].Type.SomeType.Name.Name)

				// variant 1
				assert.True(t, union.Variants[1].IsConstructor)
				assert.Equal(t, "Red", union.Variants[1].Constructor.Name)
				assert.Equal(t, 0, len(union.Variants[1].Constructor.Fields))

				// variant 2
				assert.Equal(t, "", union.Variants[2].TypeAlias.SomeType.Name.Namespace)
				assert.Equal(t, "string", union.Variants[2].TypeAlias.SomeType.Name.Name)
			})

			t.Run("basic definition with templates", func(t *testing.T) {
				_, it := setupIterator(` testNs.testName<x:type> = x:int; `)
				comb, newIt, err := parseTL2Combinator(it)
				assert.NoError(t, err)
				assert.True(t, newIt.expect(eof))

				assert.Equal(t, 1, len(comb.TypeDecl.TemplateArguments))

				arg := comb.TypeDecl.TemplateArguments[0]

				assert.Equal(t, "x", arg.Name)
				assert.Equal(t, "type", arg.Category)
			})

			t.Run("basic definition with multiple templates", func(t *testing.T) {
				_, it := setupIterator(` testNs.testName<x:type, y:int> = x:int; `)
				comb, newIt, err := parseTL2Combinator(it)
				assert.NoError(t, err)
				assert.True(t, newIt.expect(eof))

				assert.Equal(t, 2, len(comb.TypeDecl.TemplateArguments))

				arg := comb.TypeDecl.TemplateArguments[0]

				assert.Equal(t, "x", arg.Name)
				assert.Equal(t, "type", arg.Category)

				arg = comb.TypeDecl.TemplateArguments[1]

				assert.Equal(t, "y", arg.Name)
				assert.Equal(t, "int", arg.Category)
			})
		})
	})

	t.Run("function declration", func(t *testing.T) {
		t.Run("simple function", func(t *testing.T) {
			_, it := setupIterator(` testNs.testName#09123456 x:int => int; `)
			comb, newIt, err := parseTL2Combinator(it)
			assert.NoError(t, err)
			assert.True(t, newIt.expect(eof))
			assert.True(t, comb.IsFunction)
			assert.Equal(t, "testNs", comb.FuncDecl.Name.Namespace)
			assert.Equal(t, "testName", comb.FuncDecl.Name.Name)

			fields := comb.FuncDecl.Arguments

			assert.Equal(t, 1, len(fields))
			assert.Equal(t, "x", fields[0].Name)

			field0Type := fields[0].Type
			assert.False(t, field0Type.IsBracket)

			assert.Equal(t, "", field0Type.SomeType.Name.Namespace)
			assert.Equal(t, "int", field0Type.SomeType.Name.Name)

			assert.Equal(t, "int", comb.FuncDecl.ReturnType.TypeAlias.String())
		})

		t.Run("function without tag", func(t *testing.T) {
			_, it := setupIterator(` testNs.testName x:int => int; `)
			_, _, err := parseTL2Combinator(it)
			assert.Error(t, err)
		})

		t.Run("simple function with crc32", func(t *testing.T) {
			_, it := setupIterator(` testNs.testName#90abcdef x:int => int; `)
			comb, newIt, err := parseTL2Combinator(it)
			assert.NoError(t, err)
			assert.True(t, newIt.expect(eof))
			assert.True(t, comb.IsFunction)
			assert.Equal(t, "testNs", comb.FuncDecl.Name.Namespace)
			assert.Equal(t, "testName", comb.FuncDecl.Name.Name)
			assert.NotNil(t, comb.FuncDecl.ID)
			assert.Equal(t, uint32(0x90abcdef), *comb.FuncDecl.ID)
		})

		t.Run("simple function no arguments", func(t *testing.T) {
			_, it := setupIterator(` testNs.testName#09123456 => int; `)
			comb, newIt, err := parseTL2Combinator(it)
			assert.NoError(t, err)
			assert.True(t, newIt.expect(eof))
			assert.True(t, comb.IsFunction)
			assert.Equal(t, "testNs", comb.FuncDecl.Name.Namespace)
			assert.Equal(t, "testName", comb.FuncDecl.Name.Name)

			fields := comb.FuncDecl.Arguments

			assert.Equal(t, 0, len(fields))
		})

		t.Run("function with union type result", func(t *testing.T) {
			_, it := setupIterator(` testNs.testName#09123456 x:int => int | NonFound | Found name:string _:int; `)
			comb, newIt, err := parseTL2Combinator(it)
			assert.NoError(t, err)
			assert.True(t, newIt.expect(eof))
			assert.True(t, comb.IsFunction)
			assert.Equal(t, "testNs", comb.FuncDecl.Name.Namespace)
			assert.Equal(t, "testName", comb.FuncDecl.Name.Name)

			fields := comb.FuncDecl.Arguments

			assert.Equal(t, 1, len(fields))
			assert.Equal(t, "x", fields[0].Name)

			field0Type := fields[0].Type
			assert.False(t, field0Type.IsBracket)

			assert.Equal(t, "", field0Type.SomeType.Name.Namespace)
			assert.Equal(t, "int", field0Type.SomeType.Name.Name)

			assert.True(t, comb.FuncDecl.ReturnType.IsUnionType)
			assert.Equal(t, 3, len(comb.FuncDecl.ReturnType.UnionType.Variants))

			variants := comb.FuncDecl.ReturnType.UnionType.Variants

			// 0
			assert.False(t, variants[0].IsConstructor)
			assert.Equal(t, "int", variants[0].TypeAlias.String())

			// 1
			assert.True(t, variants[1].IsConstructor)
			assert.Equal(t, "NonFound", variants[1].Constructor.Name)
			assert.Equal(t, 0, len(variants[1].Constructor.Fields))

			// 2
			assert.True(t, variants[2].IsConstructor)
			assert.Equal(t, "Found", variants[2].Constructor.Name)
			assert.Equal(t, 2, len(variants[2].Constructor.Fields))
			assert.Equal(t, "name", variants[2].Constructor.Fields[0].Name)
			assert.Equal(t, "_", variants[2].Constructor.Fields[1].Name)

		})
	})

	t.Run("type reference", func(t *testing.T) {
		t.Run("simple type", func(t *testing.T) {
			_, it := setupIterator(` int `)
			state, newIt, typeRef := parseTL2Type(it, Position{})

			assert.True(t, state.StartProcessing)
			assert.NoError(t, state.Error)
			assert.True(t, newIt.expect(eof))

			assert.False(t, typeRef.IsBracket)
			assert.Equal(t, "", typeRef.SomeType.Name.Namespace)
			assert.Equal(t, "int", typeRef.SomeType.Name.Name)

			assert.Equal(t, 0, len(typeRef.SomeType.Arguments))
		})

		t.Run("simple type with namespace", func(t *testing.T) {
			_, it := setupIterator(` ns.int `)
			state, newIt, typeRef := parseTL2Type(it, Position{})

			assert.True(t, state.StartProcessing)
			assert.NoError(t, state.Error)
			assert.True(t, newIt.expect(eof))

			assert.False(t, typeRef.IsBracket)
			assert.Equal(t, "ns", typeRef.SomeType.Name.Namespace)
			assert.Equal(t, "int", typeRef.SomeType.Name.Name)

			assert.Equal(t, 0, len(typeRef.SomeType.Arguments))
		})

		t.Run("vector type", func(t *testing.T) {
			_, it := setupIterator(` []int `)
			state, newIt, typeRef := parseTL2Type(it, Position{})

			assert.True(t, state.StartProcessing)
			assert.NoError(t, state.Error)
			assert.True(t, newIt.expect(eof))

			assert.True(t, typeRef.IsBracket)

			bracket := typeRef.BracketType

			assert.Equal(t, "", bracket.ArrayType.SomeType.Name.Namespace)
			assert.Equal(t, "int", bracket.ArrayType.SomeType.Name.Name)

			assert.Nil(t, bracket.IndexType)
		})

		t.Run("fixed array type", func(t *testing.T) {
			_, it := setupIterator(` [4]int `)
			state, newIt, typeRef := parseTL2Type(it, Position{})

			assert.True(t, state.StartProcessing)
			assert.NoError(t, state.Error)
			assert.True(t, newIt.expect(eof))

			assert.True(t, typeRef.IsBracket)

			bracket := typeRef.BracketType

			assert.Equal(t, "", bracket.ArrayType.SomeType.Name.Namespace)
			assert.Equal(t, "int", bracket.ArrayType.SomeType.Name.Name)

			assert.NotNil(t, bracket.IndexType)
			assert.True(t, bracket.IndexType.IsNumber)
			assert.Equal(t, uint32(4), bracket.IndexType.Number)
		})

		t.Run("fixed array type with generic", func(t *testing.T) {
			_, it := setupIterator(` [n]int `)
			state, newIt, typeRef := parseTL2Type(it, Position{})

			assert.True(t, state.StartProcessing)
			assert.NoError(t, state.Error)
			assert.True(t, newIt.expect(eof))

			assert.True(t, typeRef.IsBracket)

			bracket := typeRef.BracketType

			assert.Equal(t, "", bracket.ArrayType.SomeType.Name.Namespace)
			assert.Equal(t, "int", bracket.ArrayType.SomeType.Name.Name)

			assert.NotNil(t, bracket.IndexType)
			assert.False(t, bracket.IndexType.IsNumber)
			assert.Equal(t, "n", bracket.IndexType.Type.SomeType.Name.Name)
		})

		t.Run("deep type", func(t *testing.T) {
			_, it := setupIterator(`[[]list<[]int>]array<2, [][]string> `)
			state, newIt, typeRef := parseTL2Type(it, Position{})

			assert.True(t, state.StartProcessing)
			assert.NoError(t, state.Error)
			assert.True(t, newIt.expect(eof))

			assert.True(t, typeRef.IsBracket)
			assert.NotNil(t, typeRef.BracketType.IndexType)

			indexType := typeRef.BracketType.IndexType
			//arrayType := typeRef.BracketType.ArrayType

			assert.False(t, indexType.IsNumber)
			assert.True(t, indexType.Type.IsBracket)
			assert.Nil(t, indexType.Type.BracketType.IndexType)
			assert.Equal(t, "list", indexType.Type.BracketType.ArrayType.SomeType.Name.String())
		})
	})

	t.Run("check print", func(t *testing.T) {
		t.Run("combinator", func(t *testing.T) {
			_, it := setupIterator(` testNs.testName <x:int,  y:type  >#09abcdef =Green x:int |   Red | string   ; `)
			comb, _, _ := parseTL2Combinator(it)
			assert.Equal(t,
				`testNs.testName<x:int,y:type>#09abcdef = Green x:int | Red | string;`,
				comb.String(),
			)
		})
		t.Run("combinators", func(t *testing.T) {
			str := ` 

testNs.testName <x:int,  y:type  > =
	  Green x:int 
	| Red _:int
	| string // my comment  
; 
@x@r testFunc#09123456 a:uint32 t:vector<vector < int>> 
	=> int;`
			combs, _ := ParseTL2(str)
			assert.Equal(t,
				`testNs.testName<x:int,y:type> = Green x:int | Red _:int | string;
@x @r testFunc#09123456 a:uint32 t:vector<vector<int>> => int;
`,
				combs.String(),
			)
		})
	})
}
