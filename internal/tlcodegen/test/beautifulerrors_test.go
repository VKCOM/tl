// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package test

import (
	"strings"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/vkcom/tl/internal/tlast"
	"github.com/vkcom/tl/internal/tlcodegen"
)

func TestBeautiful(t *testing.T) {
	var errorBuffer strings.Builder
	errorTests := []struct {
		name             string
		sourceTL         string
		parsingError     string
		compilationError string
	}{
		{
			"empty square bracket RED",
			"a n:# data:n*[] = A;\nb n:# n*[] = B;",
			"",
			"a n:# data:\x1b[31mn*[]\x1b[0m = A;\n           \x1b[0m^^^^\x1b[0m-- replacing with canonical tuples: repetition with no fields is not allowed  (line 1 col 12)\n",
		},
		{
			"anonymous scale from previous field should be # PURPLE",
			"a n:string data:[int] = A;",
			"",
			"a n:\x1b[35mstring\x1b[0m data:[int] = A;\n    \x1b[0m^^^^^^\x1b[0m-- see here  (line 1 col 5)\na n:string data:\x1b[31m\x1b[0m[int] = A;\n                \x1b[0m^\x1b[0m-- replacing with canonical tuples: anonymous scale repeat implicitly references previous field \"n\" which should have type #  (line 1 col 17)\n",
		},
		{
			"tag duplication PURPLE",
			"a#12345678 = A;\nb#12345678 = B;",
			"",
			"a\x1b[35m#12345678\x1b[0m = A;\n \x1b[0m^^^^^^^^^\x1b[0m-- see here  (line 1 col 2)\nb\x1b[31m#12345678\x1b[0m = B;\n \x1b[0m^^^^^^^^^\x1b[0m-- constructor tag #12345678 used by \"a\" is used again by \"b\"  (line 2 col 2)\n",
		},
		{
			"constructor tag 0 RED",
			"a#00000000 = A;",
			"",
			"a\x1b[31m#00000000\x1b[0m = A;\n \x1b[0m^^^^^^^^^\x1b[0m-- constructor tag 0 is prohibited, even if generated implicitly  (line 1 col 2)\n",
		},
		{
			"constructor name duplication PURPLE",
			"a = A;\na = B;",
			"",
			"\x1b[35ma\x1b[0m = A;\n\x1b[0m^\x1b[0m-- see here  (line 1 col 1)\n\x1b[31ma\x1b[0m = B;\n\x1b[0m^\x1b[0m-- constructor name \"a\" is used again here  (line 2 col 1)\n",
		},
		{
			"anonymous inner square bracket RED",
			"a {n:#} data:[[int]] = A n;\nb n:# x:(a n) = B;",
			"",
			"a {n:#} data:[\x1b[31m\x1b[0m[int]] = A n;\n              \x1b[0m^\x1b[0m-- replacing with canonical tuples: anonymous scale repeat can be used only in top-level square brackets  (line 1 col 15)\n",
		},
		{
			"wrong scale type from template args PURPLE",
			"a {T:Type} data:[int] = A;",
			"",
			"a \x1b[35m{T:Type}\x1b[0m data:[int] = A;\n  \x1b[0m^^^^^^^^\x1b[0m-- see here  (line 1 col 3)\na {T:Type} data:\x1b[31m\x1b[0m[int] = A;\n                \x1b[0m^\x1b[0m-- replacing with canonical tuples: anonymous scale repeat implicitly references last template parameter \"T\" which should have type #  (line 1 col 17)\n",
		},
		{
			"local nat args name collision PURPLE",
			"a {n:#} {n:#} = A n n;\nb n:# a:(a n n) = B;",
			"",
			"a \x1b[35m{n:#}\x1b[0m {n:#} = A n n;\n  \x1b[0m^^^^^\x1b[0m-- see here  (line 1 col 3)\na {n:#} \x1b[31m{n:#}\x1b[0m = A n n;\n        \x1b[0m^^^^^\x1b[0m-- nat-parametr name collision  (line 1 col 9)\n",
		},
		{

			"local Type args name collision PURPLE",
			"a {T:Type} {n:Type} = A n n;\nb a:(a int int) = B;",
			"",
			"a \x1b[35m{T:Type}\x1b[0m {n:Type} = A n n;\n  \x1b[0m^^^^^^^^\x1b[0m-- see here  (line 1 col 3)\na {T:Type} {n:Type} = A \x1b[31mn\x1b[0m n;\n                        \x1b[0m^\x1b[0m-- type declaration \"A\" has wrong template argument name \"n\" here  (line 1 col 25)\n",
		},
		{
			"type declaration with missing template argument",
			"a {n:#} = A;\nb n:# a:(a n) = B;",
			"",
			"a {n:#} = A\x1b[31m\x1b[0m;\n           \x1b[0m^\x1b[0m-- type declaration \"A\" is missing template argument \"n\" here  (line 1 col 12)\n",
		},
		{
			"type declaration with excess template argument RED",
			"a {n:#} = A n m;\n",
			"",
			"a {n:#} = A n \x1b[31mm\x1b[0m;\n              \x1b[0m^\x1b[0m-- type declaration \"A\" has excess template argument \"m\" here  (line 1 col 15)\n",
		},
		{
			"type declaration with wrong template argument name PURPLE",
			"a {T:Type} = A n;\nb n:# a:(a n) = A;",
			"",
			"a \x1b[35m{T:Type}\x1b[0m = A n;\n  \x1b[0m^^^^^^^^\x1b[0m-- see here  (line 1 col 3)\na {T:Type} = A \x1b[31mn\x1b[0m;\n               \x1b[0m^\x1b[0m-- type declaration \"A\" has wrong template argument name \"n\" here  (line 1 col 16)\n",
		},
		{
			"union constructor with missing argument PURPLE",
			"a {n:#} {m:#} = Union n m;\nb {n:#} = Union n;",
			"",
			"a {n:#} \x1b[35m{m:#}\x1b[0m = Union n m;\n        \x1b[0m^^^^^\x1b[0m-- see here  (line 1 col 9)\nb {n:#}\x1b[31m\x1b[0m = Union n;\n       \x1b[0m^\x1b[0m-- union constructor \"b\" has missing argument \"m\" here  (line 2 col 8)\n",
		},
		{
			"union constructor with excess argument",
			"a {n:#} = Union n;\nb {n:#} {m:#} = Union n m;",
			"",
			"a {n:#}\x1b[35m\x1b[0m = Union n;\n       \x1b[0m^\x1b[0m-- see here  (line 1 col 8)\nb {n:#} \x1b[31m{m:#}\x1b[0m = Union n m;\n        \x1b[0m^^^^^\x1b[0m-- union constructor \"b\" has excess argument \"m\" here  (line 2 col 9)\n",
		},
		{
			"union constructor with different argument name PURPLE",
			"a {n:#} = Union n;\nb {m:#} = Union m;",
			"",
			"a \x1b[35m{n:#}\x1b[0m = Union n;\n  \x1b[0m^^^^^\x1b[0m-- see here  (line 1 col 3)\nb \x1b[31m{m:#}\x1b[0m = Union m;\n  \x1b[0m^^^^^\x1b[0m-- union constructor \"b\" has different argument name or type here \"m\"  (line 2 col 3)\n",
		},
		{
			"union constructor with different argument type PURPLE",
			"a {n:#} = Union n;\nb {T:Type} = Union T;",
			"",
			"a \x1b[35m{n:#}\x1b[0m = Union n;\n  \x1b[0m^^^^^\x1b[0m-- see here  (line 1 col 3)\nb \x1b[31m{T:Type}\x1b[0m = Union T;\n  \x1b[0m^^^^^^^^\x1b[0m-- union constructor \"b\" has different argument name or type here \"T\"  (line 2 col 3)\n",
		},
		{
			"builtin wrapper cannot have template parameters RED",
			"int {n:#} n = Int n;",
			"",
			"int \x1b[31m{n:#}\x1b[0m n = Int n;\n    \x1b[0m^^^^^\x1b[0m-- builtin wrapper \"int\" cannot have template parameters  (line 1 col 5)\n",
		},
		{
			"builtin wrapper must have constructor name equal to some builtin type RED",
			"foo ? = Foo;",
			"",
			"\x1b[31mfoo\x1b[0m ? = Foo;\n\x1b[0m^^^\x1b[0m-- builtin wrapper \"foo\" must have constructor name equal to some builtin type  (line 1 col 1)\n",
		},
		{
			"builtin wrapper cannot have template parameter RED",
			"int {n:#} = Int n;",
			"",
			"int \x1b[31m{n:#}\x1b[0m = Int n;\n    \x1b[0m^^^^^\x1b[0m-- builtin wrapper \"int\" cannot have template parameters  (line 1 col 5)\n",
		},
		{
			"builtin wrapper must have exactly one field",
			"int = Int;",
			"",
			"int \x1b[31m\x1b[0m= Int;\n    \x1b[0m^\x1b[0m-- builtin wrapper \"int\" must have exactly 1 field  (line 1 col 5)\n",
		},
		{
			"builtin wrapper has excess must have exactly one field RED",
			"int a:int b:string = Int;",
			"",
			"int a:int b:\x1b[31mstring\x1b[0m = Int;\n            \x1b[0m^^^^^^\x1b[0m-- builtin wrapper \"int\" has excess field, must have exactly 1  (line 1 col 13)\n",
		},
		{
			"builtin wrapper field must be anonymous RED",
			"int i:int = Int;",
			"",
			"int \x1b[31mi\x1b[0m:int = Int;\n    \x1b[0m^\x1b[0m-- builtin wrapper \"int\" field must be anonymous  (line 1 col 5)\n",
		},
		{
			"builtin wrapper field must not use field mask",
			"int field_mask.0?# = Int;",
			"",
			"int \x1b[31mfield_mask.0?#\x1b[0m = Int;\n    \x1b[0m^^^^^^^^^^^^^^\x1b[0m-- anonymous fields are discouraged, except when used in '# a:[int]' pattern or when type has single anonymous field without fieldmask (typedef-like)  (line 1 col 5)\n",
		},
		{
			"builtin wrapper field type must match constructor name", // TODO: didn't find case
			"a int = A;",
			"",
			"",
		},
		{
			"builtin wrapper must have constructor name equal to some builtin type", // TODO: didn't find case
			"a int = A;",
			"",
			"",
		},
		{
			"field mask error",
			"foo x?field_mask.0",
			"",
			"",
		},
		{
			"field mask error",
			"foo x.field_mask?0",
			"",
			"",
		},
		{
			"field mask error",
			"foo field_mask.0?x",
			"",
			"",
		},
		{
			"field mask error",
			"foo field_mask?0.x",
			"",
			"",
		},
		{
			"reference to nat param where type is required",
			"a {T:Type} x:T = A T;\nb {n:#} a:(a n) = B n;\nc n:# b:(b n) = C;",
			"",
			"b \x1b[35m{n:#}\x1b[0m a:(a n) = B n;\n  \x1b[0m^^^^^\x1b[0m-- see here  (line 2 col 3)\nb {n:#} a:(a \x1b[31mn\x1b[0m) = B n;\n             \x1b[0m^\x1b[0m-- reference to #-param \"n\" where type is required  (line 2 col 14)\n",
		},
		{
			"reference to nat field where type is required",
			"a {T:Type} x:T = A T;\nb n:# a:(a n) = B;",
			"",
			"b \x1b[35mn\x1b[0m:# a:(a n) = B;\n  \x1b[0m^\x1b[0m-- see here  (line 2 col 3)\nb n:# a:(a \x1b[31mn\x1b[0m) = B;\n           \x1b[0m^\x1b[0m-- reference to field \"n\" where type is required  (line 2 col 12)\n",
		},
		{
			"passing constant to Type-param is impossible",
			"a {T:Type} x:T = A T;\nb a:(a 3) = B;",
			"",
			"a \x1b[35m{T:Type}\x1b[0m x:T = A T;\n  \x1b[0m^^^^^^^^\x1b[0m-- declared here  (line 1 col 3)\nb a:(a \x1b[31m3\x1b[0m) = B;\n       \x1b[0m^\x1b[0m-- passing constant \"3\" to Type-param \"T\" is impossible  (line 2 col 8)\n",
		},
		{
			"reference to template type arg cannot have arguments",
			"vector {T:Type} size:# data:size*[T] = Vector T;\na {T:Type} x:(T int) = A T;\nb a:(a (vector int)) = B;",
			"",
			"a \x1b[35m{T:Type}\x1b[0m x:(T int) = A T;\n  \x1b[0m^^^^^^^^\x1b[0m-- defined here  (line 2 col 3)\na {T:Type} x:\x1b[31m(T int)\x1b[0m = A T;\n             \x1b[0m^^^^^^^\x1b[0m-- reference to template type arg \"T\" cannot have arguments  (line 2 col 14)\n",
		},
		{
			"field type is bare so union cannot be passed",
			"foo i:int s:string = Union;\nbar s:string i:int = Union;\nbareWrapper {X:Type} a:%X = BareWrapper X;\nbareWrapperTest a:(bareWrapper Union) = BareWrapperTest;",
			"",
			"bareWrapper \x1b[35m{X:Type}\x1b[0m a:%X = BareWrapper X;\n            \x1b[0m^^^^^^^^\x1b[0m-- defined here  (line 3 col 13)\nbareWrapper {X:Type} a:\x1b[31m%X\x1b[0m = BareWrapper X;\n                       \x1b[0m^^\x1b[0m-- field type \"X\" is bare, so union \"Union\" cannot be passed  (line 3 col 24)\n",
		},
		{
			"reference to union cannot be bare",
			"a = Union;\nb b:int = Union;\nuse a:%Union = Use;",
			"",
			"a = \x1b[35mUnion\x1b[0m;\n    \x1b[0m^^^^^\x1b[0m-- see more  (line 1 col 5)\nuse a:\x1b[31m%Union\x1b[0m = Use;\n      \x1b[0m^^^^^^\x1b[0m-- reference to union \"Union\" cannot be bare  (line 3 col 7)\n",
		},
		{
			"reference to function constructor is not allowed",
			"---types---\nint int = Int;\nuseF f:f = UseF;\n---functions---\n@any f = Int;",
			"",
			"@any \x1b[35mf\x1b[0m = Int;\n     \x1b[0m^\x1b[0m-- see more  (line 5 col 6)\nuseF f:\x1b[31mf\x1b[0m = UseF;\n       \x1b[0m^\x1b[0m-- reference to function constructor \"f\" is not allowed  (line 3 col 8)\n",
		},
		{
			"reference to union constructor is not allowed",
			"a = Union;\nb i:int = Union;\nuseConstructor a:a = UseConstructor;",
			"",
			"\x1b[35ma\x1b[0m = Union;\n\x1b[0m^\x1b[0m-- see more  (line 1 col 1)\nuseConstructor a:\x1b[31ma\x1b[0m = UseConstructor;\n                 \x1b[0m^\x1b[0m-- reference to union constructor \"a\" is not allowed  (line 3 col 18)\n",
		},
		{
			"anonymous fields are discouraged",
			"foo n:# # = Foo;",
			"",
			"foo n:# \x1b[31m#\x1b[0m = Foo;\n        \x1b[0m^\x1b[0m-- anonymous fields are discouraged, except when used in '# a:[int]' pattern or when type has single anonymous field without fieldmask (typedef-like)  (line 1 col 9)\n",
		},
		{
			"namespace must not differ by case only",
			`
service5.emptyOutput = service5.Output;
service5.hrenOutput = ch_Proxy.Output;
`,
			"",
			"\x1b[35mservice5.hrenOutput\x1b[0m = ch_Proxy.Output;\n\x1b[0m^^^^^^^^^^^^^^^^^^^\x1b[0m-- see here  (line 3 col 1)\nservice5.hrenOutput = \x1b[31mch_Proxy.Output\x1b[0m;\n                      \x1b[0m^^^^^^^^^^^^^^^\x1b[0m-- namespaces must not differ by only case  (line 3 col 23)\n",
		},
	}
	for i, test := range errorTests {
		errorBuffer.Reset()
		t.Run(test.name, func(t *testing.T) {
			if test.name == "field mask error" {
				t.Skip()
			}
			var parseError *tlast.ParseError
			if ast, err := tlast.ParseTL(test.sourceTL); err != nil && errors.As(err, &parseError) {
				parseError.ConsolePrint(&errorBuffer, err, false)
				require.Equalf(t, test.parsingError, errorBuffer.String(), "errors for parsing stage didn't match for test %d", i)

			} else if _, err := tlcodegen.GenerateCode(ast, tlcodegen.Gen2Options{ErrorWriter: &errorBuffer}); err != nil && errors.As(err, &parseError) {
				parseError.ConsolePrint(&errorBuffer, err, false)
				require.Equalf(t, test.compilationError, errorBuffer.String(), "errors for compiling stage didn't match for test %d", i)
			}

		})
	}
}
