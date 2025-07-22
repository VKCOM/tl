// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package test

import (
	"errors"
	"github.com/TwiN/go-color"
	"strings"
	"testing"

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
			"a n:# data:" + color.InRed("n*[]") + " = A;\n           " + color.InWhite("^^^^--") + " replacing with canonical tuples: repetition with no fields is not allowed  (line 1 col 12)\n",
		},
		{
			"anonymous scale from previous field should be # PURPLE",
			"a n:string data:[int] = A;",
			"",
			"a n:" + color.InPurple("string") + " data:[int] = A;\n    " + color.InWhite("^^^^^^--") + " see here  (line 1 col 5)\na n:string data:" + color.InRed("") + "[int] = A;\n                " + color.InWhite("^--") + " replacing with canonical tuples: anonymous scale repeat implicitly references previous field \"n\" which should have type #  (line 1 col 17)\n",
		},
		{
			"tag duplication PURPLE",
			"a#12345678 = A;\nb#12345678 = B;",
			"",
			"a" + color.InPurple("#12345678") + " = A;\n " + color.InWhite("^^^^^^^^^--") + " see here  (line 1 col 2)\nb" + color.InRed("#12345678") + " = B;\n " + color.InWhite("^^^^^^^^^--") + " constructor tag #12345678 used by \"a\" is used again by \"b\"  (line 2 col 2)\n",
		},
		{
			"constructor tag 0 RED",
			"a#00000000 = A;",
			"",
			"a" + color.InRed("#00000000") + " = A;\n " + color.InWhite("^^^^^^^^^--") + " constructor tag 0 is prohibited, even if generated implicitly  (line 1 col 2)\n",
		},
		{
			"constructor name duplication PURPLE",
			"a = A;\na = B;",
			"",
			color.InPurple("a") + " = A;\n" + color.InWhite("^--") + " see here  (line 1 col 1)\n" + color.InRed("a") + " = B;\n" + color.InWhite("^--") + " constructor name \"a\" is used again here  (line 2 col 1)\n",
		},
		{
			"anonymous inner square bracket RED",
			"a {n:#} data:[[int]] = A n;\nb n:# x:(a n) = B;",
			"",
			"a {n:#} data:[" + color.InRed("") + "[int]] = A n;\n              " + color.InWhite("^--") + " replacing with canonical tuples: anonymous scale repeat can be used only in top-level square brackets  (line 1 col 15)\n",
		},
		{
			"wrong scale type from template args PURPLE",
			"a {T:Type} data:[int] = A;",
			"",
			"a " + color.InPurple("{T:Type}") + " data:[int] = A;\n  " + color.InWhite("^^^^^^^^--") + " see here  (line 1 col 3)\na {T:Type} data:" + color.InRed("") + "[int] = A;\n                " + color.InWhite("^--") + " replacing with canonical tuples: anonymous scale repeat implicitly references last template parameter \"T\" which should have type #  (line 1 col 17)\n",
		},
		{
			"local nat args name collision PURPLE",
			"a {n:#} {n:#} = A n n;\nb n:# a:(a n n) = B;",
			"",
			"a " + color.InPurple("{n:#}") + " {n:#} = A n n;\n  " + color.InWhite("^^^^^--") + " see here  (line 1 col 3)\na {n:#} " + color.InRed("{n:#}") + " = A n n;\n        " + color.InWhite("^^^^^--") + " nat-parametr name collision  (line 1 col 9)\n",
		},
		{

			"local Type args name collision PURPLE",
			"a {T:Type} {n:Type} = A n n;\nb a:(a int int) = B;",
			"",
			"a " + color.InPurple("{T:Type}") + " {n:Type} = A n n;\n  " + color.InWhite("^^^^^^^^--") + " see here  (line 1 col 3)\na {T:Type} {n:Type} = A " + color.InRed("n") + " n;\n                        " + color.InWhite("^--") + " type declaration \"A\" has wrong template argument name \"n\" here  (line 1 col 25)\n",
		},
		{
			"type declaration with missing template argument",
			"a {n:#} = A;\nb n:# a:(a n) = B;",
			"",
			"a {n:#} = A" + color.InRed("") + ";\n           " + color.InWhite("^--") + " type declaration \"A\" is missing template argument \"n\" here  (line 1 col 12)\n",
		},
		{
			"type declaration with excess template argument RED",
			"a {n:#} = A n m;\n",
			"",
			"a {n:#} = A n " + color.InRed("m") + ";\n              " + color.InWhite("^--") + " type declaration \"A\" has excess template argument \"m\" here  (line 1 col 15)\n",
		},
		{
			"type declaration with wrong template argument name PURPLE",
			"a {T:Type} = A n;\nb n:# a:(a n) = A;",
			"",
			"a " + color.InPurple("{T:Type}") + " = A n;\n  " + color.InWhite("^^^^^^^^--") + " see here  (line 1 col 3)\na {T:Type} = A " + color.InRed("n") + ";\n               " + color.InWhite("^--") + " type declaration \"A\" has wrong template argument name \"n\" here  (line 1 col 16)\n",
		},
		{
			"union constructor with missing argument PURPLE",
			"a {n:#} {m:#} = Union n m;\nb {n:#} = Union n;",
			"",
			"a {n:#} " + color.InPurple("{m:#}") + " = Union n m;\n        " + color.InWhite("^^^^^--") + " see here  (line 1 col 9)\nb {n:#}" + color.InRed("") + " = Union n;\n       " + color.InWhite("^--") + " union constructor \"b\" has missing argument \"m\" here  (line 2 col 8)\n",
		},
		{
			"union constructor with excess argument",
			"a {n:#} = Union n;\nb {n:#} {m:#} = Union n m;",
			"",
			"a {n:#}" + color.InPurple("") + " = Union n;\n       " + color.InWhite("^--") + " see here  (line 1 col 8)\nb {n:#} " + color.InRed("{m:#}") + " = Union n m;\n        " + color.InWhite("^^^^^--") + " union constructor \"b\" has excess argument \"m\" here  (line 2 col 9)\n",
		},
		{
			"union constructor with different argument name PURPLE",
			"a {n:#} = Union n;\nb {m:#} = Union m;",
			"",
			"a " + color.InPurple("{n:#}") + " = Union n;\n  " + color.InWhite("^^^^^--") + " see here  (line 1 col 3)\nb " + color.InRed("{m:#}") + " = Union m;\n  " + color.InWhite("^^^^^--") + " union constructor \"b\" has different argument name or type here \"m\"  (line 2 col 3)\n",
		},
		{
			"union constructor with different argument type PURPLE",
			"a {n:#} = Union n;\nb {T:Type} = Union T;",
			"",
			"a " + color.InPurple("{n:#}") + " = Union n;\n  " + color.InWhite("^^^^^--") + " see here  (line 1 col 3)\nb " + color.InRed("{T:Type}") + " = Union T;\n  " + color.InWhite("^^^^^^^^--") + " union constructor \"b\" has different argument name or type here \"T\"  (line 2 col 3)\n",
		},
		{
			"builtin wrapper cannot have template parameters RED",
			"int {n:#} n = Int n;",
			"",
			"int " + color.InRed("{n:#}") + " n = Int n;\n    " + color.InWhite("^^^^^--") + " builtin wrapper \"int\" cannot have template parameters  (line 1 col 5)\n",
		},
		{
			"builtin wrapper must have constructor name equal to some builtin type RED",
			"foo ? = Foo;",
			"",
			color.InRed("foo") + " ? = Foo;\n" + color.InWhite("^^^--") + " builtin wrapper \"foo\" must have constructor name equal to some builtin type  (line 1 col 1)\n",
		},
		{
			"builtin wrapper cannot have template parameter RED",
			"int {n:#} = Int n;",
			"",
			"int " + color.InRed("{n:#}") + " = Int n;\n    " + color.InWhite("^^^^^--") + " builtin wrapper \"int\" cannot have template parameters  (line 1 col 5)\n",
		},
		{
			"builtin wrapper must have exactly one field",
			"int = Int;",
			"",
			"int " + color.InRed("") + "= Int;\n    " + color.InWhite("^--") + " builtin wrapper \"int\" must have exactly 1 field  (line 1 col 5)\n",
		},
		{
			"builtin wrapper has excess must have exactly one field RED",
			"int a:int b:string = Int;",
			"",
			"int a:int b:" + color.InRed("string") + " = Int;\n            " + color.InWhite("^^^^^^--") + " builtin wrapper \"int\" has excess field, must have exactly 1  (line 1 col 13)\n",
		},
		{
			"builtin wrapper field must be anonymous RED",
			"int i:int = Int;",
			"",
			"int " + color.InRed("i") + ":int = Int;\n    " + color.InWhite("^--") + " builtin wrapper \"int\" field must be anonymous  (line 1 col 5)\n",
		},
		{
			"builtin wrapper field must not use field mask",
			"int field_mask.0?# = Int;",
			"",
			"int " + color.InRed("field_mask.0?#") + " = Int;\n    " + color.InWhite("^^^^^^^^^^^^^^--") + " anonymous fields are discouraged, except when used in '# a:[int]' pattern or when type has single anonymous field without fieldmask (typedef-like)  (line 1 col 5)\n",
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
			"b " + color.InPurple("{n:#}") + " a:(a n) = B n;\n  " + color.InWhite("^^^^^--") + " see here  (line 2 col 3)\nb {n:#} a:(a " + color.InRed("n") + ") = B n;\n             " + color.InWhite("^--") + " reference to #-param \"n\" where type is required  (line 2 col 14)\n",
		},
		{
			"reference to nat field where type is required",
			"a {T:Type} x:T = A T;\nb n:# a:(a n) = B;",
			"",
			"b " + color.InPurple("n") + ":# a:(a n) = B;\n  " + color.InWhite("^--") + " see here  (line 2 col 3)\nb n:# a:(a " + color.InRed("n") + ") = B;\n           " + color.InWhite("^--") + " reference to field \"n\" where type is required  (line 2 col 12)\n",
		},
		{
			"passing constant to Type-param is impossible",
			"a {T:Type} x:T = A T;\nb a:(a 3) = B;",
			"",
			"a " + color.InPurple("{T:Type}") + " x:T = A T;\n  " + color.InWhite("^^^^^^^^--") + " declared here  (line 1 col 3)\nb a:(a " + color.InRed("3") + ") = B;\n       " + color.InWhite("^--") + " passing constant \"3\" to Type-param \"T\" is impossible  (line 2 col 8)\n",
		},
		{
			"reference to template type arg cannot have arguments",
			"vector {T:Type} size:# data:size*[T] = Vector T;\na {T:Type} x:(T int) = A T;\nb a:(a (vector int)) = B;",
			"",
			"a " + color.InPurple("{T:Type}") + " x:(T int) = A T;\n  " + color.InWhite("^^^^^^^^--") + " defined here  (line 2 col 3)\na {T:Type} x:" + color.InRed("(T int)") + " = A T;\n             " + color.InWhite("^^^^^^^--") + " reference to template type arg \"T\" cannot have arguments  (line 2 col 14)\n",
		},
		{
			"field type is bare so union cannot be passed",
			"foo i:int s:string = Union;\nbar s:string i:int = Union;\nbareWrapper {X:Type} a:%X = BareWrapper X;\nbareWrapperTest a:(bareWrapper Union) = BareWrapperTest;",
			"",
			"bareWrapper " + color.InPurple("{X:Type}") + " a:%X = BareWrapper X;\n            " + color.InWhite("^^^^^^^^--") + " defined here  (line 3 col 13)\nbareWrapper {X:Type} a:" + color.InRed("%X") + " = BareWrapper X;\n                       " + color.InWhite("^^--") + " field type \"X\" is bare, so union \"Union\" cannot be passed  (line 3 col 24)\n",
		},
		{
			"reference to union cannot be bare",
			"a = Union;\nb b:int = Union;\nuse a:%Union = Use;",
			"",
			"a = " + color.InPurple("Union") + ";\n    " + color.InWhite("^^^^^--") + " see more  (line 1 col 5)\nuse a:" + color.InRed("%Union") + " = Use;\n      " + color.InWhite("^^^^^^--") + " reference to union \"Union\" cannot be bare  (line 3 col 7)\n",
		},
		{
			"reference to function constructor is not allowed",
			"---types---\nint int = Int;\nuseF f:f = UseF;\n---functions---\n@any f = Int;",
			"",
			"@any " + color.InPurple("f") + " = Int;\n     " + color.InWhite("^--") + " see more  (line 5 col 6)\nuseF f:" + color.InRed("f") + " = UseF;\n       " + color.InWhite("^--") + " reference to function constructor \"f\" is not allowed  (line 3 col 8)\n",
		},
		{
			"reference to union constructor is not allowed",
			"a = Union;\nb i:int = Union;\nuseConstructor a:a = UseConstructor;",
			"",
			color.InPurple("a") + " = Union;\n" + color.InWhite("^--") + " see more  (line 1 col 1)\nuseConstructor a:" + color.InRed("a") + " = UseConstructor;\n                 " + color.InWhite("^--") + " reference to union constructor \"a\" is not allowed  (line 3 col 18)\n",
		},
		{
			"anonymous fields are discouraged",
			"foo n:# # = Foo;",
			"",
			"foo n:# " + color.InRed("#") + " = Foo;\n        " + color.InWhite("^--") + " anonymous fields are discouraged, except when used in '# a:[int]' pattern or when type has single anonymous field without fieldmask (typedef-like)  (line 1 col 9)\n",
		},
		{
			"namespace must not differ by case only",
			`
service5.emptyOutput = service5.Output;
service5.hrenOutput = service5.Output;
`,
			"",
			color.InPurple("service5.hrenOutput") + " = ch_Proxy.Output;\n" + color.InWhite("^^^^^^^^^^^^^^^^^^^--") + " see here  (line 3 col 1)\nservice5.hrenOutput = " + color.InRed("ch_Proxy.Output") + ";\n                      " + color.InWhite("^^^^^^^^^^^^^^^--") + " namespaces must not differ by only case  (line 3 col 23)\n",
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
