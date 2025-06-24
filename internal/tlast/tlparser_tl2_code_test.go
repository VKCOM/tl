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
			assert.NoError(t, state.Error)
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
				assert.Equal(t, "", union.Variants[0].TypeDef.SomeType.Name.Namespace)
				assert.Equal(t, "int", union.Variants[0].TypeDef.SomeType.Name.Name)

				assert.Equal(t, "", union.Variants[1].TypeDef.SomeType.Name.Namespace)
				assert.Equal(t, "string", union.Variants[1].TypeDef.SomeType.Name.Name)
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
				assert.Equal(t, "", union.Variants[2].TypeDef.SomeType.Name.Namespace)
				assert.Equal(t, "string", union.Variants[2].TypeDef.SomeType.Name.Name)
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
	})
}
