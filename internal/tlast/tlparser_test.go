// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlast

import (
	"errors"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func ptr(i uint32) *uint32 {
	var x = i
	return &x
}

func requireEqualTL(t *testing.T, a TL, b TL) {
	tlText1 := a.String()
	tlText2 := b.String()
	//require.Equal(t, tl1, tl2)
	require.Equal(t, tlText1, tlText2)
}

func TestParser(t *testing.T) {
	t.Run("double square bracket", func(t *testing.T) {
		str := `z n:# a:n*[inner1:3*[int] inner2:int] = Z;`
		ast, err := ParseTL(str)
		require.NoError(t, err)
		log.Printf("%s", ast.String())
	})
	t.Run("square bracket variants", func(t *testing.T) {
		str := `
replace1 {n:#} a:[int] = Replace1;
replace2 n:# a:[int] = Replace2;
replace3 a:3*[int] = Replace3;
replace4 {n:#} a:n*[int] = Replace4;
replace5 n:# a:n*[int] = Replace5;`
		ast, err := ParseTL(str)
		require.NoError(t, err)
		log.Printf("%s", ast.String())
	})
	t.Run("hard square bracket", func(t *testing.T) {
		str := `replace6 {n:#} {X:Type} k:# a:k*[b:n.0?3*[Pair int X] c:k.0?n*[int]] = Replace6;`
		ast, err := ParseTL(str)
		require.NoError(t, err)
		log.Printf("%s", ast.String())
	})

	t.Run("Primitive", func(t *testing.T) {
		primitives := "int#a8509bda ? = Int;\n" +
			"positiveInt int = PositiveInt;\n" +
			"boolFalse#bc799737 = Bool;\n"
		tl1, err := ParseTL(primitives)
		require.NoError(t, err)
		tl2 := TL{
			{
				Builtin:           true,
				Construct:         Constructor{Name: Name{Name: "int"}, ID: ptr(0xa8509bda)},
				TemplateArguments: nil,
				TypeDecl:          TypeDeclaration{Name: Name{Name: "Int"}},
			},
			{
				Construct:         Constructor{Name: Name{Name: "positiveInt"}},
				TemplateArguments: nil,
				Fields:            []Field{{FieldType: TypeRef{Type: Name{Name: "int"}}}},
				TypeDecl:          TypeDeclaration{Name: Name{Name: "PositiveInt"}},
			},
			{
				Construct:         Constructor{Name: Name{Name: "boolFalse"}, ID: ptr(0xbc799737)},
				TemplateArguments: nil,
				Fields:            nil,
				TypeDecl:          TypeDeclaration{Name: Name{Name: "Bool"}},
			},
		}
		requireEqualTL(t, tl1, tl2)
	})

	t.Run("Bare %(...)", func(t *testing.T) {
		bare := "string ? = String;\n" +
			"vector#1cb5c415 {t:Type} # [t] = Vector t;\n" +
			"dictionaryField {t:Type} key:string value:t = DictionaryField t;\n" +
			"dictionary#1f4c618f {t:Type} %(Vector %(DictionaryField t)) = Dictionary t;\n" +
			"stat#9d56e6b2 %(Dictionary string) = Stat;\n"
		tl1, err := ParseTL(bare)
		require.NoError(t, err)
		tl2 := TL{
			{
				Construct:         Constructor{Name: Name{Name: "stat"}, ID: ptr(0x9d56e6b2)},
				TemplateArguments: nil,
				Fields: []Field{
					{
						FieldType: TypeRef{
							Type: Name{
								Name: "Dictionary",
							},
							Args: []ArithmeticOrType{
								{
									T: TypeRef{
										Type: Name{
											Name: "string",
										}}}},
							Bare: true,
						}}},
				TypeDecl: TypeDeclaration{Name: Name{Name: "Stat"}},
			}}
		requireEqualTL(t, TL{tl1[len(tl1)-1]}, TL{tl2[0]})
		// require.Equal(t, tl1[len(tl1)-1], tl2[0])
	})

	t.Run("Bare (% ... )", func(t *testing.T) {
		bare := "int ? = Int;\n" +
			"foo.bar {n:#} = foo.Bar n;\n" +
			"foo.biba {m:#} id:(%foo.Bar m) i:int = foo.Biba m;\n"
		tl1, err := ParseTL(bare)
		require.NoError(t, err)
		tl2 := TL{
			{
				Construct: Constructor{
					Name: Name{
						Namespace: "foo",
						Name:      "bar",
					},
				},
				TemplateArguments: []TemplateArgument{
					{"n", true, PositionRange{}},
				},
				TypeDecl: TypeDeclaration{
					Name: Name{
						Namespace: "foo", Name: "Bar",
					},
					Arguments: []string{
						"n",
					}},
			},
			{
				Construct: Constructor{
					Name: Name{
						Namespace: "foo",
						Name:      "biba",
					},
				},
				TemplateArguments: []TemplateArgument{
					{"m", true, PositionRange{}},
				},
				Fields: []Field{
					{
						FieldName: "id",
						FieldType: TypeRef{
							Type: Name{
								Namespace: "foo",
								Name:      "Bar",
							},
							Args: []ArithmeticOrType{
								{
									T: TypeRef{
										Type: Name{Name: "m"},
									}}},
							Bare: true,
						}},
					{
						FieldName: "i",
						FieldType: TypeRef{
							Type: Name{
								Name: "int",
							}}}},
				TypeDecl: TypeDeclaration{
					Name: Name{
						Namespace: "foo", Name: "Biba",
					},
					Arguments: []string{
						"m",
					}}}}
		requireEqualTL(t, tl2, tl1[1:])

	})

	t.Run("Bare (%(vector t))", func(t *testing.T) {
		str := "vector#1cb5c415 {t:Type} # [t] = Vector t;\n" +
			"a {t:Type} (%(vector t)) = A t;"
		tl1, err := ParseTL(str)
		require.NoError(t, err)
		require.NotNil(t, tl1)
	})

	t.Run("Fail at (%(Vector t) ...)", func(t *testing.T) {
		str := "a {t:Type} n:(%(Vector t) m) = A t;"
		_, err := ParseTL(str)
		require.Error(t, err)
	})

	t.Run("Bare (%vector<t> ...)", func(t *testing.T) {
		str := "vector#1cb5c415 {t:Type} # [t] = Vector t;\n" +
			"a {t:Type} (%vector<t>) = A t;"
		tl1, err := ParseTL(str)
		require.NoError(t, err)
		tl2 := TL{
			{
				Construct: Constructor{Name: Name{Name: "a"}},
				TemplateArguments: []TemplateArgument{
					{FieldName: "t", IsNat: false},
				},
				Fields: []Field{
					{
						FieldType: TypeRef{
							Type: Name{Name: "vector"},
							Args: []ArithmeticOrType{
								{
									T: TypeRef{
										Type: Name{Name: "t"},
									}}},
							Bare: true}}},
				TypeDecl: TypeDeclaration{
					Name:      Name{Name: "A"},
					Arguments: []string{"t"},
				}}}
		requireEqualTL(t, TL{tl1[len(tl1)-1]}, TL{tl2[0]})
		// require.Equal(t, tl1[len(tl1)-1], tl2[0])
	})
	t.Run("Arithmetic", func(t *testing.T) {
		str := `
foo {X:#} = Foo X;
hren1 a:(foo  (1 + (3 + (4 + 5) + 2) + 0)) = Hren;
hren2 a:(foo  (((((2))))) )= Hren;
`
		ast, err := ParseTL(str)
		require.NoError(t, err)
		require.Equal(t, uint32(1+3+4+5+2), ast[1].Fields[0].FieldType.Args[0].Arith.Res)
		require.Equal(t, uint32(2), ast[2].Fields[0].FieldType.Args[0].Arith.Res)
		tests := []struct {
			tlText     string
			epectedErr error
		}{
			{
				`foo {X:#} = Foo X; hren1 a:(foo  1 + + 2) = Hren;`,
				errors.New("arithmetic expression expected after '+'"),
			},
			{
				`foo {X:#} = Foo X; hren2 a:(foo  (2 2))= Hren;`,
				errors.New("')' expected"),
			},
			{
				`foo {X:#} = Foo X; hren3 a:(foo  (+))= Hren;`,
				errors.New("')' or type is expected here"),
			},
			{
				`foo {X:#} = Foo X; hren3 a:(foo  (4294967295+1))= Hren;`,
				errors.New("arithmetic expression overflows uint32"),
			},
		}
		for i, t_ := range tests {
			_, err := ParseTL(t_.tlText)
			require.ErrorAs(t, err, &t_.epectedErr, "test %d", i)
		}
	})
}

func TestSplitMultilineComment(t *testing.T) {
	lines := SplitMultilineComment("hren\r\nvam\npopolam\n\r\n\r\n")
	result := []string{"hren", "vam", "popolam", "", "", ""}
	require.Equal(t, lines, result)
}
