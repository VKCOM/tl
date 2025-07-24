// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package test

import (
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vkcom/tl/internal/tlast"
	"github.com/vkcom/tl/internal/tlcodegen"
)

func TestGen(t *testing.T) {
	outputDir := "output"
	defer func() { require.NoError(t, os.RemoveAll(outputDir)) }()
	data, err := os.ReadFile("./tls/goldmaster.tl")

	require.NoError(t, err)

	ast, err := tlast.ParseTL(string(data))
	if err != nil {
		t.Error(err)
	}

	gen, err := tlcodegen.GenerateCode(ast, tlcodegen.Gen2Options{
		ErrorWriter: io.Discard,
		Verbose:     true,
	})

	require.NoError(t, err)
	require.NoError(t, os.RemoveAll(outputDir))

	err = gen.WriteToDir(outputDir)

	require.NoError(t, err)
}

func TestPHPLinterCheck(t *testing.T) {
	t.Run("Issue with boxed reference to flat type", func(t *testing.T) {
		t.Run("fail on boxed ref to flat type", func(t *testing.T) {
			data := `vector#12345679 {t:Type} # [t] = Vector t;
myType#12345678 x:Vector<int> = MyType;
---functions---
@read myTestFunction x:MyType = MyType;`

			ast, err := tlast.ParseTL(data)
			if err != nil {
				t.Error(err)
				return
			}

			_, err = tlcodegen.GenerateCode(ast, tlcodegen.Gen2Options{
				ErrorWriter:    io.Discard,
				Verbose:        true,
				LinterPHPCheck: true,
			})

			assert.Error(t, err)
		})

		t.Run("no fail when all flat types are bare", func(t *testing.T) {
			data := `vector2#1cb5c415 {t:Type} n:# m:n*[t] = Vector t;
myType#12345678 x:vector2<int> = MyType;
---functions---
@read myTestFunction x:myType = MyType;`

			ast, err := tlast.ParseTL(data)
			if err != nil {
				t.Error(err)
				return
			}

			_, err = tlcodegen.GenerateCode(ast, tlcodegen.Gen2Options{
				ErrorWriter:    io.Discard,
				Verbose:        true,
				LinterPHPCheck: true,
			})

			assert.NoError(t, err)
		})

		t.Run("not fail for special cases", func(t *testing.T) {
			argsToAdd := map[string]string{
				"Vector":            "<int>",
				"Tuple":             "<int, 2>",
				"Dictionary":        "<int>",
				"IntKeyDictionary":  "<int>",
				"LongKeyDictionary": "<int>",
				"Maybe":             "<int>",
			}
			for _, s := range tlcodegen.PHPNamesToIgnoreForLinterCheck {
				t.Run(s, func(t *testing.T) {
					data := `// Builtin types
int#a8509bda ? = Int;
long#22076cba ? = Long;
float#824dab22 ? = Float;    // 4 bytes -- single precision
double#2210c154 ? = Double;  // 8 bytes -- double precision
string#b5286e24 ? = String;

// Common aliases with checks
positiveInt int = PositiveInt;
positiveLong long = PositiveLong;
nonNegativeInt int = NonNegativeInt;
nonNegativeLong long = NonNegativeLong;

// Boolean emulation
boolFalse#bc799737 = Bool;
boolTrue#997275b5 = Bool;

// Boolean for diagonal queries
boolStat statTrue:int statFalse:int statUnknown:int = BoolStat;

// Vector
vector#1cb5c415 {t:Type} # [t] = Vector t;
tuple#9770768a {t:Type} {n:#} [t] = Tuple t n;
vectorTotal {t:Type} total_count:int vector:%(Vector t) = VectorTotal t;


statOne key:string value:string = StatOne;
//stat vector:%(Vector %statOne) = Stat;
stat#9d56e6b2 %(Dictionary string) = Stat;

dictionaryField {t:Type} key:string value:t = DictionaryField t;
dictionary#1f4c618f {t:Type} %(Vector %(DictionaryField t)) = Dictionary t;

intKeyDictionaryField {t:Type}
    key:int
    value:t
    = IntKeyDictionaryField t;

intKeyDictionary#07bafc42 {t:Type}
    %(Vector %(intKeyDictionaryField t))
    = IntKeyDictionary t;

longKeyDictionaryField {t:Type}
    key:long
    value:t
    = LongKeyDictionaryField t;

longKeyDictionary#b424d8f1 {t:Type}
    %(Vector %(longKeyDictionaryField t))
    = LongKeyDictionary t;

resultFalse#27930a7b {t:Type} = Maybe t;
resultTrue#3f9c8ef8 {t:Type} result:t = Maybe t;


pair {X:Type} {Y:Type} a:X b:Y = Pair X Y;

map {X:Type} {Y:Type} key:X value:Y = Map X Y;

true = True; // this can be used as void type and serialized to empty array in PHP

myType#12345678 x:[[[TYPE_HERE]]] = MyType;
---functions---
@read myTestFunction x:myType = MyType;`

					ast, err := tlast.ParseTL(strings.Replace(data, "[[[TYPE_HERE]]]", s+argsToAdd[s], 1))
					if err != nil {
						t.Error(err)
						return
					}

					_, err = tlcodegen.GenerateCode(ast, tlcodegen.Gen2Options{
						ErrorWriter:    io.Discard,
						Verbose:        true,
						LinterPHPCheck: true,
					})

					assert.NoError(t, err)
				})
			}
		})
	})

	t.Run("Issue template flat types", func(t *testing.T) {
		t.Run("fail on non-special type with template", func(t *testing.T) {
			data := `int#a8509bda ? = Int;
vector#12345679 {t:Type} # [t] = Vector t;
myType#12345678 {t:Type} x:vector<T> = MyType T;
---functions---
@read myTestFunction x:myType<int> = MyType<int>;`

			ast, err := tlast.ParseTL(data)
			if err != nil {
				t.Error(err)
				return
			}

			_, err = tlcodegen.GenerateCode(ast, tlcodegen.Gen2Options{
				ErrorWriter:    io.Discard,
				Verbose:        true,
				LinterPHPCheck: true,
			})

			assert.Error(t, err)
		})

		// similar reference for original issue
		t.Run("fail for exact case for sure", func(t *testing.T) {
			data := `int#a8509bda ? = Int;
vector#12345679 {t:Type} # [t] = Vector t;
myType#eadc11aa {T:Type}
	vector<vector<T>>
	= MyType T;
---functions---
@read myTestFunction x:myType<int> = MyType<int>;`

			ast, err := tlast.ParseTL(data)
			if err != nil {
				t.Error(err)
				return
			}

			_, err = tlcodegen.GenerateCode(ast, tlcodegen.Gen2Options{
				ErrorWriter:    io.Discard,
				Verbose:        true,
				LinterPHPCheck: true,
			})

			assert.Error(t, err)

		})

		t.Run("not fail for special cases", func(t *testing.T) {
			argsToAdd := map[string]string{
				"Vector":            "<int>",
				"Tuple":             "<int, 2>",
				"Dictionary":        "<int>",
				"IntKeyDictionary":  "<int>",
				"LongKeyDictionary": "<int>",
				"Maybe":             "<int>",
			}
			for _, s := range tlcodegen.PHPNamesToIgnoreForLinterCheck {
				t.Run(s, func(t *testing.T) {
					data := `// Builtin types
int#a8509bda ? = Int;
long#22076cba ? = Long;
float#824dab22 ? = Float;    // 4 bytes -- single precision
double#2210c154 ? = Double;  // 8 bytes -- double precision
string#b5286e24 ? = String;

// Common aliases with checks
positiveInt int = PositiveInt;
positiveLong long = PositiveLong;
nonNegativeInt int = NonNegativeInt;
nonNegativeLong long = NonNegativeLong;

// Boolean emulation
boolFalse#bc799737 = Bool;
boolTrue#997275b5 = Bool;

// Boolean for diagonal queries
boolStat statTrue:int statFalse:int statUnknown:int = BoolStat;

// Vector
vector#1cb5c415 {t:Type} # [t] = Vector t;
tuple#9770768a {t:Type} {n:#} [t] = Tuple t n;
vectorTotal {t:Type} total_count:int vector:%(Vector t) = VectorTotal t;


statOne key:string value:string = StatOne;
//stat vector:%(Vector %statOne) = Stat;
stat#9d56e6b2 %(Dictionary string) = Stat;

dictionaryField {t:Type} key:string value:t = DictionaryField t;
dictionary#1f4c618f {t:Type} %(Vector %(DictionaryField t)) = Dictionary t;

intKeyDictionaryField {t:Type}
    key:int
    value:t
    = IntKeyDictionaryField t;

intKeyDictionary#07bafc42 {t:Type}
    %(Vector %(intKeyDictionaryField t))
    = IntKeyDictionary t;

longKeyDictionaryField {t:Type}
    key:long
    value:t
    = LongKeyDictionaryField t;

longKeyDictionary#b424d8f1 {t:Type}
    %(Vector %(longKeyDictionaryField t))
    = LongKeyDictionary t;

resultFalse#27930a7b {t:Type} = Maybe t;
resultTrue#3f9c8ef8 {t:Type} result:t = Maybe t;


pair {X:Type} {Y:Type} a:X b:Y = Pair X Y;

map {X:Type} {Y:Type} key:X value:Y = Map X Y;

true = True; // this can be used as void type and serialized to empty array in PHP

myType#12345678 x:[[[TYPE_HERE]]] = MyType;
---functions---
@read myTestFunction x:myType = MyType;`

					ast, err := tlast.ParseTL(strings.Replace(data, "[[[TYPE_HERE]]]", s+argsToAdd[s], 1))
					if err != nil {
						t.Error(err)
						return
					}

					_, err = tlcodegen.GenerateCode(ast, tlcodegen.Gen2Options{
						ErrorWriter:    io.Discard,
						Verbose:        true,
						LinterPHPCheck: true,
					})

					assert.NoError(t, err)
				})
			}
		})
	})
}
