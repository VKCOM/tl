package tlast

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupIterator(str string) (tokenIterator, error) {
	lex := newLexer(str, "", LexerOptions{LexerLanguage: TL2})
	allTokens, err := lex.generateTokens()

	return tokenIterator{tokens: allTokens}, err
}

func TestParseTL2(t *testing.T) {
	t.Run("optional parse annotations", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			it, _ := setupIterator(`   @gigi    `)
			state, newIt, ann := parseTL2Annotation(it, Position{})
			assert.True(t, state.StartProcessing)
			assert.NoError(t, state.Error)
			assert.Equal(t, "gigi", ann.Name)
			assert.True(t, newIt.expect(eof))
		})

		t.Run("don't start processing", func(t *testing.T) {
			it, _ := setupIterator(`   #11111111    `)
			state, newIt, ann := parseTL2Annotation(it, Position{})
			assert.False(t, state.StartProcessing)
			assert.NoError(t, state.Error)
			assert.Equal(t, "", ann.Name)
			assert.True(t, newIt.expect(crc32hash))
		})

		t.Run("parse several", func(t *testing.T) {
			it, _ := setupIterator(` @gigi @gaga #12345678  `)
			state, newIt, anns := zeroOrMore(parseTL2Annotation)(it, Position{})
			assert.True(t, state.StartProcessing)
			assert.NoError(t, state.Error)
			assert.Equal(t, 2, len(anns))
			assert.Equal(t, "gigi", anns[0].Name)
			assert.Equal(t, "gaga", anns[1].Name)
			assert.True(t, newIt.expect(crc32hash))
		})

		t.Run("parse zero", func(t *testing.T) {
			it, _ := setupIterator(` #12345678  `)
			state, newIt, anns := zeroOrMore(parseTL2Annotation)(it, Position{})
			assert.False(t, state.StartProcessing)
			assert.NoError(t, state.Error)
			assert.Equal(t, 0, len(anns))
			assert.True(t, newIt.expect(crc32hash))
		})
	})

	t.Run("optional parse type name", func(t *testing.T) {
		t.Run("success without namespace", func(t *testing.T) {
			it, _ := setupIterator(` testName  `)
			state, newIt, testName := parseTL2TypeName(it, Position{})
			assert.True(t, state.StartProcessing)
			assert.NoError(t, state.Error)
			assert.Equal(t, "testName", testName.Name)
			assert.True(t, newIt.expect(eof))
		})
		t.Run("success with namespace", func(t *testing.T) {
			it, _ := setupIterator(` testNS.testName  `)
			state, newIt, testName := parseTL2TypeName(it, Position{})
			assert.True(t, state.StartProcessing)
			assert.NoError(t, state.Error)
			assert.Equal(t, "testName", testName.Name)
			assert.Equal(t, "testNS", testName.Namespace)
			assert.True(t, newIt.expect(eof))
		})
		t.Run("success with ucName", func(t *testing.T) {
			it, _ := setupIterator(` TestName  `)
			state, newIt, _ := parseTL2TypeName(it, Position{})
			assert.True(t, state.StartProcessing)
			assert.NoError(t, state.Error)
			assert.True(t, newIt.expect(eof))
		})
		t.Run("success with ucNameNs", func(t *testing.T) {
			it, _ := setupIterator(` testNs.TestName  `)
			state, newIt, _ := parseTL2TypeName(it, Position{})
			assert.True(t, state.StartProcessing)
			assert.NoError(t, state.Error)
			assert.True(t, newIt.expect(eof))
		})
	})

	t.Run("combinator", func(t *testing.T) {
		t.Run("type declaration", func(t *testing.T) {
			t.Run("basic definition", func(t *testing.T) {
				it, _ := setupIterator(` testNs.testName = x:int; `)
				comb, newIt, err := parseTL2Combinator(it)
				assert.NoError(t, err)
				assert.True(t, newIt.expect(eof))
				assert.False(t, comb.IsFunction)
				assert.Equal(t, "testNs", comb.TypeDecl.Name.Namespace)
				assert.Equal(t, "testName", comb.TypeDecl.Name.Name)
				assert.False(t, comb.TypeDecl.Type.IsTypeAlias)
				assert.False(t, comb.TypeDecl.Type.StructType.IsUnionType)

				fields := comb.TypeDecl.Type.StructType.ConstructorFields

				assert.Equal(t, 1, len(fields))
				assert.Equal(t, "x", fields[0].Name)

				field0Type := fields[0].Type
				assert.False(t, field0Type.IsBracket)

				assert.Equal(t, "", field0Type.SomeType.Name.Namespace)
				assert.Equal(t, "int", field0Type.SomeType.Name.Name)
			})

			t.Run("basic definition with crc32", func(t *testing.T) {
				it, _ := setupIterator(` testNs.testName #90abcdef = x:int; `)
				comb, newIt, err := parseTL2Combinator(it)
				assert.NoError(t, err)
				assert.True(t, newIt.expect(eof))
				assert.False(t, comb.IsFunction)
				assert.Equal(t, "testNs", comb.TypeDecl.Name.Namespace)
				assert.Equal(t, "testName", comb.TypeDecl.Name.Name)
				assert.Equal(t, uint32(0x90abcdef), comb.TypeDecl.Magic)
			})

			t.Run("basic definition with optional field", func(t *testing.T) {
				it, _ := setupIterator(` testNs.testName = x?:int; `)
				comb, newIt, err := parseTL2Combinator(it)
				assert.NoError(t, err)
				assert.True(t, newIt.expect(eof))
				assert.False(t, comb.IsFunction)
				assert.Equal(t, "testNs", comb.TypeDecl.Name.Namespace)
				assert.Equal(t, "testName", comb.TypeDecl.Name.Name)
				assert.False(t, comb.TypeDecl.Type.IsTypeAlias)
				assert.False(t, comb.TypeDecl.Type.StructType.IsUnionType)

				fields := comb.TypeDecl.Type.StructType.ConstructorFields

				assert.Equal(t, 1, len(fields))
				assert.Equal(t, "x", fields[0].Name)
				assert.True(t, fields[0].IsOptional)

				field0Type := fields[0].Type
				assert.False(t, field0Type.IsBracket)

				assert.Equal(t, "", field0Type.SomeType.Name.Namespace)
				assert.Equal(t, "int", field0Type.SomeType.Name.Name)
			})

			t.Run("basic definition with ignored field", func(t *testing.T) {
				it, _ := setupIterator(` testNs.testName = _:int; `)
				comb, newIt, err := parseTL2Combinator(it)
				assert.NoError(t, err)
				assert.True(t, newIt.expect(eof))
				assert.False(t, comb.IsFunction)
				assert.Equal(t, "testNs", comb.TypeDecl.Name.Namespace)
				assert.Equal(t, "testName", comb.TypeDecl.Name.Name)
				assert.False(t, comb.TypeDecl.Type.IsTypeAlias)
				assert.False(t, comb.TypeDecl.Type.StructType.IsUnionType)

				fields := comb.TypeDecl.Type.StructType.ConstructorFields

				assert.Equal(t, 1, len(fields))
				assert.Equal(t, "_", fields[0].Name)
				assert.True(t, fields[0].IsIgnored)

				field0Type := fields[0].Type
				assert.False(t, field0Type.IsBracket)

				assert.Equal(t, "", field0Type.SomeType.Name.Namespace)
				assert.Equal(t, "int", field0Type.SomeType.Name.Name)
			})

			t.Run("basic definition with ignored optional field", func(t *testing.T) {
				it, _ := setupIterator(` testNs.testName = _?:int; `)
				_, _, err := parseTL2Combinator(it)
				assert.Error(t, err)
			})

			t.Run("empty definition", func(t *testing.T) {
				it, _ := setupIterator(` testNs.testName = ; `)
				comb, newIt, err := parseTL2Combinator(it)
				assert.NoError(t, err)
				assert.True(t, newIt.expect(eof))
				assert.False(t, comb.IsFunction)
				assert.Equal(t, "testNs", comb.TypeDecl.Name.Namespace)
				assert.Equal(t, "testName", comb.TypeDecl.Name.Name)
				assert.False(t, comb.TypeDecl.Type.IsTypeAlias)
				assert.False(t, comb.TypeDecl.Type.StructType.IsUnionType)

				fields := comb.TypeDecl.Type.StructType.ConstructorFields

				assert.Equal(t, 0, len(fields))
			})

			t.Run("type alias", func(t *testing.T) {
				it, _ := setupIterator(` testNs.testName <=> int; `)
				comb, newIt, err := parseTL2Combinator(it)
				assert.NoError(t, err)
				assert.True(t, newIt.expect(eof))
				assert.False(t, comb.IsFunction)
				assert.Equal(t, "testNs", comb.TypeDecl.Name.Namespace)
				assert.Equal(t, "testName", comb.TypeDecl.Name.Name)

				assert.True(t, comb.TypeDecl.Type.IsTypeAlias)

				alias := comb.TypeDecl.Type.TypeAlias

				assert.False(t, alias.IsBracket)
				assert.Equal(t, "", alias.SomeType.Name.Namespace)
				assert.Equal(t, "int", alias.SomeType.Name.Name)
			})

			t.Run("union", func(t *testing.T) {
				it, _ := setupIterator(` testNs.testName = SomeInt int | SomeString string; `)
				comb, newIt, err := parseTL2Combinator(it)
				assert.NoError(t, err)
				assert.True(t, newIt.expect(eof))
				assert.False(t, comb.IsFunction)
				assert.Equal(t, "testNs", comb.TypeDecl.Name.Namespace)
				assert.Equal(t, "testName", comb.TypeDecl.Name.Name)

				assert.False(t, comb.TypeDecl.Type.IsTypeAlias)
				assert.True(t, comb.TypeDecl.Type.StructType.IsUnionType)

				union := comb.TypeDecl.Type.StructType.UnionType

				assert.Equal(t, 2, len(union.Variants))
				assert.Equal(t, "", union.Variants[0].TypeAlias.SomeType.Name.Namespace)
				assert.Equal(t, "int", union.Variants[0].TypeAlias.SomeType.Name.Name)

				assert.Equal(t, "", union.Variants[1].TypeAlias.SomeType.Name.Namespace)
				assert.Equal(t, "string", union.Variants[1].TypeAlias.SomeType.Name.Name)
			})

			t.Run("union with constructors", func(t *testing.T) {
				it, _ := setupIterator(` testNs.testName = Green x:int | Red | SomeString string; `)
				comb, newIt, err := parseTL2Combinator(it)
				assert.NoError(t, err)
				assert.True(t, newIt.expect(eof))
				assert.False(t, comb.IsFunction)
				assert.Equal(t, "testNs", comb.TypeDecl.Name.Namespace)
				assert.Equal(t, "testName", comb.TypeDecl.Name.Name)

				assert.False(t, comb.TypeDecl.Type.IsTypeAlias)
				assert.True(t, comb.TypeDecl.Type.StructType.IsUnionType)

				union := comb.TypeDecl.Type.StructType.UnionType

				assert.Equal(t, 3, len(union.Variants))

				// variant 0
				assert.False(t, union.Variants[0].IsTypeAlias)
				assert.Equal(t, "Green", union.Variants[0].Name)
				assert.Equal(t, 1, len(union.Variants[0].Fields))

				assert.Equal(t, "", union.Variants[0].Fields[0].Type.SomeType.Name.Namespace)
				assert.Equal(t, "int", union.Variants[0].Fields[0].Type.SomeType.Name.Name)

				// variant 1
				assert.False(t, union.Variants[1].IsTypeAlias)
				assert.Equal(t, "Red", union.Variants[1].Name)
				assert.Equal(t, 0, len(union.Variants[1].Fields))

				// variant 2
				assert.Equal(t, "", union.Variants[2].TypeAlias.SomeType.Name.Namespace)
				assert.Equal(t, "string", union.Variants[2].TypeAlias.SomeType.Name.Name)
			})

			t.Run("union with constructors and trailing vb", func(t *testing.T) {
				it, _ := setupIterator(`
testNs.testName = 
	| Green x:int 
	| Red 
	| Text string
	;`)
				comb, newIt, err := parseTL2Combinator(it)
				assert.NoError(t, err)
				assert.True(t, newIt.expect(eof))
				assert.False(t, comb.IsFunction)
				assert.Equal(t, "testNs", comb.TypeDecl.Name.Namespace)
				assert.Equal(t, "testName", comb.TypeDecl.Name.Name)

				assert.False(t, comb.TypeDecl.Type.IsTypeAlias)
				assert.True(t, comb.TypeDecl.Type.StructType.IsUnionType)

				union := comb.TypeDecl.Type.StructType.UnionType

				assert.Equal(t, 3, len(union.Variants))

				// variant 0
				assert.False(t, union.Variants[0].IsTypeAlias)
				assert.Equal(t, "Green", union.Variants[0].Name)
				assert.Equal(t, 1, len(union.Variants[0].Fields))

				assert.Equal(t, "", union.Variants[0].Fields[0].Type.SomeType.Name.Namespace)
				assert.Equal(t, "int", union.Variants[0].Fields[0].Type.SomeType.Name.Name)

				// variant 1
				assert.False(t, union.Variants[1].IsTypeAlias)
				assert.Equal(t, "Red", union.Variants[1].Name)
				assert.Equal(t, 0, len(union.Variants[1].Fields))

				// variant 2
				assert.True(t, union.Variants[2].IsTypeAlias)
				assert.Equal(t, "", union.Variants[2].TypeAlias.SomeType.Name.Namespace)
				assert.Equal(t, "string", union.Variants[2].TypeAlias.SomeType.Name.Name)
			})

			t.Run("union with constructors with one constructor", func(t *testing.T) {
				it, _ := setupIterator(`
testNs.testName = Green x:int
	;`)
				_, _, err := parseTL2Combinator(it)
				assert.Error(t, err)
			})

			t.Run("union with constructors with one constructor with leading vb", func(t *testing.T) {
				it, _ := setupIterator(`
testNs.testName = | Green x:int
	;`)
				_, _, err := parseTL2Combinator(it)
				assert.NoError(t, err)
			})

			t.Run("union with constructors with one constructor with leading vb and lcName", func(t *testing.T) {
				it, _ := setupIterator(`
testNs.testName = | green x:int
	;`)
				_, _, err := parseTL2Combinator(it)
				assert.NoError(t, err)
			})

			t.Run("union with constructors with zero constructor", func(t *testing.T) {
				it, _ := setupIterator(`
testNs.testName = |
	;`)
				_, _, err := parseTL2Combinator(it)
				assert.Error(t, err)
			})

			t.Run("union with constructors with one constructor and empty alternative", func(t *testing.T) {
				it, _ := setupIterator(`
testNs.testName = Green x:int |
	;`)
				_, _, err := parseTL2Combinator(it)
				assert.Error(t, err)
			})

			t.Run("basic definition with templates", func(t *testing.T) {
				it, _ := setupIterator(` testNs.testName<x:Type> = x:int; `)
				comb, newIt, err := parseTL2Combinator(it)
				assert.NoError(t, err)
				assert.True(t, newIt.expect(eof))

				assert.Equal(t, 1, len(comb.TypeDecl.TemplateArguments))

				arg := comb.TypeDecl.TemplateArguments[0]

				assert.Equal(t, "x", arg.Name)
				assert.Equal(t, TL2TypeCategory{IsNatValue: false}, arg.Category)
			})

			t.Run("union without variant declarations", func(t *testing.T) {
				it, _ := setupIterator(` testNs.testName<x:Type> = |; `)
				_, _, err := parseTL2Combinator(it)
				fmt.Println(err)
				assert.Error(t, err)
			})

			t.Run("correct categories", func(t *testing.T) {
				it, _ := setupIterator(` testNs.testName<x:Type, y:#> = ; `)
				_, _, err := parseTL2Combinator(it)
				fmt.Println(err)
				assert.NoError(t, err)
			})

			t.Run("incorrect categories", func(t *testing.T) {
				it, _ := setupIterator(` testNs.testName<x:Type, y:int> = ; `)
				_, _, err := parseTL2Combinator(it)
				fmt.Println(err)
				assert.Error(t, err)
			})

			t.Run("basic definition with multiple templates", func(t *testing.T) {
				it, _ := setupIterator(` testNs.testName<x:Type, y:#> = x:int; `)
				comb, newIt, err := parseTL2Combinator(it)
				assert.NoError(t, err)
				assert.True(t, newIt.expect(eof))

				assert.Equal(t, 2, len(comb.TypeDecl.TemplateArguments))

				arg := comb.TypeDecl.TemplateArguments[0]

				assert.Equal(t, "x", arg.Name)
				assert.Equal(t, TL2TypeCategory{IsNatValue: false}, arg.Category)

				arg = comb.TypeDecl.TemplateArguments[1]

				assert.Equal(t, "y", arg.Name)
				assert.Equal(t, TL2TypeCategory{IsNatValue: true}, arg.Category)
			})
		})
	})

	t.Run("function declaration", func(t *testing.T) {
		t.Run("simple function", func(t *testing.T) {
			it, _ := setupIterator(` testNs.testName#09123456 x:int => int; `)
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

			assert.False(t, comb.FuncDecl.ReturnType.IsTypeAlias)
			assert.False(t, comb.FuncDecl.ReturnType.StructType.IsUnionType)
			assert.Equal(t, "int", comb.FuncDecl.ReturnType.StructType.ConstructorFields[0].Type.String())
		})

		t.Run("function without tag", func(t *testing.T) {
			it, _ := setupIterator(` testNs.testName x:int => int; `)
			_, _, err := parseTL2Combinator(it)
			assert.Error(t, err)
		})

		t.Run("simple function with crc32", func(t *testing.T) {
			it, _ := setupIterator(` testNs.testName#90abcdef x:int => int; `)
			comb, newIt, err := parseTL2Combinator(it)
			assert.NoError(t, err)
			assert.True(t, newIt.expect(eof))
			assert.True(t, comb.IsFunction)
			assert.Equal(t, "testNs", comb.FuncDecl.Name.Namespace)
			assert.Equal(t, "testName", comb.FuncDecl.Name.Name)
			assert.Equal(t, uint32(0x90abcdef), comb.FuncDecl.Magic)
		})

		t.Run("simple function no arguments", func(t *testing.T) {
			it, _ := setupIterator(` testNs.testName#09123456 => int; `)
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
			it, _ := setupIterator(` testNs.testName#09123456 x:int => SomeInt int | NonFound | Found name:string _:int; `)
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

			assert.False(t, comb.FuncDecl.ReturnType.IsTypeAlias)
			assert.True(t, comb.FuncDecl.ReturnType.StructType.IsUnionType)
			assert.Equal(t, 3, len(comb.FuncDecl.ReturnType.StructType.UnionType.Variants))

			variants := comb.FuncDecl.ReturnType.StructType.UnionType.Variants

			// 0
			assert.True(t, variants[0].IsTypeAlias)
			assert.Equal(t, "int", variants[0].TypeAlias.String())

			// 1
			assert.False(t, variants[1].IsTypeAlias)
			assert.Equal(t, "NonFound", variants[1].Name)
			assert.Equal(t, 0, len(variants[1].Fields))

			// 2
			assert.False(t, variants[2].IsTypeAlias)
			assert.Equal(t, "Found", variants[2].Name)
			assert.Equal(t, 2, len(variants[2].Fields))
			assert.Equal(t, "name", variants[2].Fields[0].Name)
			assert.Equal(t, "_", variants[2].Fields[1].Name)

		})

		t.Run("func with alias return type", func(t *testing.T) {
			it, _ := setupIterator(`testNs.testName#09123456 x:int => <=>int; `)
			comb, newIt, err := parseTL2Combinator(it)
			assert.NoError(t, err)
			assert.True(t, newIt.expect(eof))

			assert.True(t, comb.IsFunction)
			assert.True(t, comb.FuncDecl.ReturnType.IsTypeAlias)
			assert.Equal(t, "int", comb.FuncDecl.ReturnType.TypeAlias.String())
		})

		t.Run("func with anonymous return type", func(t *testing.T) {
			it, _ := setupIterator(`testNs.testName#09123456 x:int => int; `)
			comb, newIt, err := parseTL2Combinator(it)
			assert.NoError(t, err)
			assert.True(t, newIt.expect(eof))

			assert.True(t, comb.IsFunction)
			assert.False(t, comb.FuncDecl.ReturnType.IsTypeAlias)

			typ := comb.FuncDecl.ReturnType.StructType

			assert.False(t, typ.IsUnionType)
			assert.Equal(t, 1, len(typ.ConstructorFields))
			assert.Equal(t, "", typ.ConstructorFields[0].Name)
			assert.Equal(t, "int", typ.ConstructorFields[0].Type.String())

		})
	})

	t.Run("type reference", func(t *testing.T) {
		t.Run("simple type", func(t *testing.T) {
			it, _ := setupIterator(` int `)
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
			it, _ := setupIterator(` ns.int `)
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
			it, _ := setupIterator(` []int `)
			state, newIt, typeRef := parseTL2Type(it, Position{})

			assert.True(t, state.StartProcessing)
			assert.NoError(t, state.Error)
			assert.True(t, newIt.expect(eof))

			assert.True(t, typeRef.IsBracket)

			bracket := typeRef.BracketType

			assert.Equal(t, "", bracket.ArrayType.SomeType.Name.Namespace)
			assert.Equal(t, "int", bracket.ArrayType.SomeType.Name.Name)

			assert.False(t, bracket.HasIndex)
		})

		t.Run("fixed array type", func(t *testing.T) {
			it, _ := setupIterator(` [4]int `)
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
			it, _ := setupIterator(` [n]int `)
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
			it, _ := setupIterator(`[[]list<[]int>]array<2, [][]string> `)
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
			assert.False(t, indexType.Type.BracketType.HasIndex)
			assert.Equal(t, "list", indexType.Type.BracketType.ArrayType.SomeType.Name.String())
		})
	})

	t.Run("check print", func(t *testing.T) {
		t.Run("combinator", func(t *testing.T) {
			it, _ := setupIterator(` testNs.testName#09abcdef <x:#,  y:Type  > =Green x:int |   Red | SomeStr string   ; `)
			comb, _, err := parseTL2Combinator(it)
			assert.NoError(t, err)
			assert.Equal(t,
				`testNs.testName#09abcdef<x:#,y:Type> = Green x:int | Red | SomeStr string;`,
				comb.String(),
			)
		})
		t.Run("combinators", func(t *testing.T) {
			str := ` 

testNs.testName <x:#,  y:Type  > =
	  Green x:int 
	| Red _:int
	| SomeStr     string // my comment  
; 
@x@r testFunc#09123456 a:uint32 t:vector<vector < int>> 
	=> int;`
			combs, err := ParseTL2(str)
			assert.NoError(t, err)
			assert.Equal(t,
				`testNs.testName<x:#,y:Type> = Green x:int | Red _:int | SomeStr string;
@x @r testFunc#09123456 a:uint32 t:vector<vector<int>> => int;
`,
				combs.String(),
			)
		})
		t.Run("combinator with comment", func(t *testing.T) {
			it, _ := setupIterator(`// comment 
testNs.testName  #09abcdef <x:#,  y:Type  > =Green x:int |   Red | SomeStr    string   ; `)
			comb, _, err := parseTL2Combinator(it)
			assert.NoError(t, err)
			assert.Equal(t,
				`// comment
testNs.testName#09abcdef<x:#,y:Type> = Green x:int | Red | SomeStr string;`,
				comb.String(),
			)
		})

		t.Run("fields with comments", func(t *testing.T) {
			it, _ := setupIterator(`testNs.testName #09abcdef <x:#,  y:Type  > = // a
x:int

// b
// c
y:int

//

z:int; `)
			comb, _, err := parseTL2Combinator(it)
			assert.NoError(t, err)
			assert.Equal(t,
				`testNs.testName#09abcdef<x:#,y:Type> = 
	// a
	x:int
	// b
	// c
	y:int
	z:int;`,
				comb.String(),
			)
		})

		t.Run("variant fields with comments", func(t *testing.T) {
			it, _ := setupIterator(`testNs.testName #09abcdef <x:#,  y:Type  > = SomeInt int | SomeStr string | Green

// a
x:int

// b
// c
y:int

//

z:int; `)
			comb, _, err := parseTL2Combinator(it)
			assert.NoError(t, err)
			assert.Equal(t,
				`testNs.testName#09abcdef<x:#,y:Type> = 
	| SomeInt int
	| SomeStr string
	| Green
		// a
		x:int
		// b
		// c
		y:int
		z:int;`,
				comb.String(),
			)
		})

		t.Run("variant with comment", func(t *testing.T) {
			it, _ := setupIterator(`testNs.testName#09abcdef <x:#,  y:Type  > = 
// SomeInt
SomeInt int 

// SomeStr
| SomeStr string // Green 
| Green; `)
			comb, _, err := parseTL2Combinator(it)
			assert.NoError(t, err)
			assert.Equal(t,
				`testNs.testName#09abcdef<x:#,y:Type> = 
	// SomeInt
	| SomeInt int
	// SomeStr
	| SomeStr string
	// Green
	| Green;`,
				comb.String(),
			)
		})
	})

	t.Run("check comments", func(t *testing.T) {
		t.Run("comments before", func(t *testing.T) {
			t.Run("combinators", func(t *testing.T) {
				t.Run("one line", func(t *testing.T) {
					it, _ := setupIterator(`// a

// b
name = x:int;`)
					comb, _, err := parseTL2Combinator(it)
					assert.NoError(t, err)
					assert.Equal(t, "// b", comb.CommentBefore)
				})

				t.Run("few lines", func(t *testing.T) {
					it, _ := setupIterator(`// a

// l1
// l2
name = x:int;`)
					comb, _, err := parseTL2Combinator(it)
					assert.NoError(t, err)
					assert.Equal(t, "// l1\n// l2", comb.CommentBefore)
				})

				t.Run("no lines", func(t *testing.T) {
					it, _ := setupIterator(`// a

// l1
// l2

name = x:int;`)
					comb, _, err := parseTL2Combinator(it)
					assert.NoError(t, err)
					assert.Equal(t, "", comb.CommentBefore)
				})
			})

			t.Run("fields", func(t *testing.T) {
				t.Run("one line for first", func(t *testing.T) {
					it, _ := setupIterator(`name = // a
														x:int;`)
					comb, _, err := parseTL2Combinator(it)
					assert.NoError(t, err)
					assert.Equal(t, "// a", comb.TypeDecl.Type.StructType.ConstructorFields[0].CommentBefore)
				})

				t.Run("one line for second", func(t *testing.T) {
					it, _ := setupIterator(`name = y: int
						// a
						x:int;`)
					comb, _, err := parseTL2Combinator(it)
					assert.NoError(t, err)
					assert.Equal(t, "", comb.TypeDecl.Type.StructType.ConstructorFields[0].CommentBefore)
					assert.Equal(t, "// a", comb.TypeDecl.Type.StructType.ConstructorFields[1].CommentBefore)
				})

				t.Run("few lines for first", func(t *testing.T) {
					it, _ := setupIterator(`// a
name = // l1
// l2
x:int;`)
					comb, _, err := parseTL2Combinator(it)
					assert.NoError(t, err)
					assert.Equal(t, "// l1\n// l2", comb.TypeDecl.Type.StructType.ConstructorFields[0].CommentBefore)
				})

				t.Run("few lines for second", func(t *testing.T) {
					it, _ := setupIterator(`name = y: int
// a
// b
						x:int;`)
					comb, _, err := parseTL2Combinator(it)
					assert.NoError(t, err)
					assert.Equal(t, "", comb.TypeDecl.Type.StructType.ConstructorFields[0].CommentBefore)
					assert.Equal(t, "// a\n// b", comb.TypeDecl.Type.StructType.ConstructorFields[1].CommentBefore)
				})

				t.Run("no lines for first", func(t *testing.T) {
					it, _ := setupIterator(`// a
name = // l1
// l2

x:int;`)
					comb, _, err := parseTL2Combinator(it)
					assert.NoError(t, err)
					assert.Equal(t, "", comb.TypeDecl.Type.StructType.ConstructorFields[0].CommentBefore)
				})

				t.Run("no lines for second", func(t *testing.T) {
					it, _ := setupIterator(`name = y: int
// a
// b

						x:int;`)
					comb, _, err := parseTL2Combinator(it)
					assert.NoError(t, err)
					assert.Equal(t, "", comb.TypeDecl.Type.StructType.ConstructorFields[0].CommentBefore)
					assert.Equal(t, "", comb.TypeDecl.Type.StructType.ConstructorFields[1].CommentBefore)
				})
			})

			t.Run("union", func(t *testing.T) {
				t.Run("before first constructor", func(t *testing.T) {
					it, _ := setupIterator(`
testNs.testName<x:Type, y:#> = 
	// IntValue
	IntValue x:int
	| StrValue string; `)
					comb, newIt, err := parseTL2Combinator(it)
					assert.NoError(t, err)
					assert.True(t, newIt.expect(eof))

					assert.Equal(t, `// IntValue`, comb.TypeDecl.Type.StructType.UnionType.Variants[0].CommentBefore)
				})

				t.Run("before first constructor with trailing vb", func(t *testing.T) {
					it, _ := setupIterator(`
testNs.testName<x:Type, y:#> = 
	// IntValue
	| IntValue x:int
	| StrValue string; `)
					comb, newIt, err := parseTL2Combinator(it)
					assert.NoError(t, err)
					assert.True(t, newIt.expect(eof))

					assert.Equal(t, `// IntValue`, comb.TypeDecl.Type.StructType.UnionType.Variants[0].CommentBefore)
				})

				t.Run("before second constructor", func(t *testing.T) {
					it, _ := setupIterator(`
testNs.testName<x:Type, y:#> = IntValue x:int 
	// StrValue
	| StrValue string; `)
					comb, newIt, err := parseTL2Combinator(it)
					assert.NoError(t, err)
					assert.True(t, newIt.expect(eof))

					assert.Equal(t, `// StrValue`, comb.TypeDecl.Type.StructType.UnionType.Variants[1].CommentBefore)
				})
			})
		})
	})

	t.Run("beautiful errors", func(t *testing.T) {

	})
}
