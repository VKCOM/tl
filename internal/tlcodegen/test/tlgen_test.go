// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package test

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestMigrationOneStep(t *testing.T) {
	var nextState []tlcodegen.FileToWrite
	var stepSuccess bool
	stepSuccess, nextState = assertMigrationStep(
		t,
		nextState,
		`
vector#1cb5c415 {t:Type} # [t] = Vector t;
tuple#9770768a {t:Type} {n:#} [t] = Tuple t n;
pair {X:Type} {Y:Type} x:X y:Y = Pair X Y;

a.t1 f1:int f2:vector<int> = a.T1;
a.t2 n:# f1:tuple<a.t1, n> f2:tuple<a.t1, 2> = a.T2;

b.t1 x:vector<a.t2> = b.T1;

c.t1 n:# f1:n.0?int = c.T1;
c.t2 f1:pair<c.t1, double> = c.T2;
`,
		"b.",
		map[string]string{
			"": `@tl1 vector<t:type> = ;
@tl1 tuple<t:type> = ;
@tl1 @tl2ext tuple_N<t:type,n:uint32> = ;
`,
			"a": `@tl1 a.t1 = ;
@tl1 a.t2 = ;
`,
			"b": `b.t1 = x:vector<a.t2>;
`,
		},
	)
	assert.True(t, stepSuccess)
}

func TestMigrationMerge(t *testing.T) {
	var nextState []tlcodegen.FileToWrite
	var stepSuccess bool
	stepSuccess, nextState = assertMigrationStep(
		t,
		nextState,
		`
vector#1cb5c415 {t:Type} # [t] = Vector t;
tuple#9770768a {t:Type} {n:#} [t] = Tuple t n;

a.t1 f1:int f2:vector<int> = a.T1;
a.t2 n:# f1:tuple<a.t1, n> = a.T2;

b.t1 x:tuple<int, 2> = b.T1;
`,
		"b.",
		map[string]string{
			"": `@tl1 tuple<t:type> = ;
@tl1 @tl2ext tuple_N<t:type,n:uint32> = ;
`,
			"b": `b.t1 = x:tuple_N<int,2>;
`,
		},
	)
	assert.True(t, stepSuccess)

	stepSuccess, nextState = assertMigrationStep(
		t,
		nextState,
		`
vector#1cb5c415 {t:Type} # [t] = Vector t;
tuple#9770768a {t:Type} {n:#} [t] = Tuple t n;

a.t1 f1:int f2:vector<int> = a.T1;
a.t2 n:# f1:tuple<a.t1, n> = a.T2;
`,
		"b.,a.,.",
		map[string]string{
			"": `tuple<t:type> = []t;
@tl2ext tuple_N<t:type,n:uint32> = [n]t;
vector<t:type> = []t;
`,
			"a": `a.t1 = f1:int f2:vector<int>;
a.t2 = n:uint32 f1:tuple<a.t1>;
`,
			"b": `b.t1 = x:tuple_N<int,2>;
`,
		},
	)
	assert.True(t, stepSuccess)
}

func TestFullMigrationStepByStep(t *testing.T) {
	var nextState []tlcodegen.FileToWrite
	var stepSuccess bool
	// STEP 1: remove b.
	stepSuccess, nextState = assertMigrationStep(
		t,
		nextState,
		`
vector#1cb5c415 {t:Type} # [t] = Vector t;
tuple#9770768a {t:Type} {n:#} [t] = Tuple t n;
pair {X:Type} {Y:Type} x:X y:Y = Pair X Y;

a.t1 f1:int f2:vector<int> = a.T1;
a.t2 n:# f1:tuple<a.t1, n> f2:tuple<a.t1, 2> = a.T2;

b.t1 x:vector<a.t2> = b.T1;

c.t1 n:# f1:n.0?int = c.T1;
c.t2 f1:pair<c.t1, double> = c.T2;
`,
		"b.",
		map[string]string{
			"": `@tl1 vector<t:type> = ;
@tl1 tuple<t:type> = ;
@tl1 @tl2ext tuple_N<t:type,n:uint32> = ;
`,
			"a": `@tl1 a.t1 = ;
@tl1 a.t2 = ;
`,
			"b": `b.t1 = x:vector<a.t2>;
`,
		},
	)
	assert.True(t, stepSuccess)

	// STEP 2: remove b. and decide to move a.
	stepSuccess, nextState = assertMigrationStep(
		t,
		nextState,
		`
vector#1cb5c415 {t:Type} # [t] = Vector t;
tuple#9770768a {t:Type} {n:#} [t] = Tuple t n;
pair {X:Type} {Y:Type} x:X y:Y = Pair X Y;

a.t1 f1:int f2:vector<int> = a.T1;
a.t2 n:# f1:tuple<a.t1, n> f2:tuple<a.t1, 2> = a.T2;

c.t1 n:# f1:n.0?int = c.T1;
c.t2 f1:pair<c.t1, double> = c.T2;
`,
		"b.,a.",
		map[string]string{
			"": `@tl1 vector<t:type> = ;
@tl1 tuple<t:type> = ;
@tl1 @tl2ext tuple_N<t:type,n:uint32> = ;
`,
			"a": `a.t1 = f1:int f2:vector<int>;
a.t2 = n:uint32 f1:tuple<a.t1> f2:tuple_N<a.t1,2>;
`,
			"b": `b.t1 = x:vector<a.t2>;
`,
		},
	)
	assert.True(t, stepSuccess)

	// STEP 3: remove b. and decide to move c.
	stepSuccess, nextState = assertMigrationStep(
		t,
		nextState,
		`
vector#1cb5c415 {t:Type} # [t] = Vector t;
tuple#9770768a {t:Type} {n:#} [t] = Tuple t n;
pair {X:Type} {Y:Type} x:X y:Y = Pair X Y;

c.t1 n:# f1:n.0?int = c.T1;
c.t2 f1:pair<c.t1, double> = c.T2;
`,
		"b.,a.,c.",
		map[string]string{
			"": `@tl1 vector<t:type> = ;
@tl1 tuple<t:type> = ;
@tl1 @tl2ext tuple_N<t:type,n:uint32> = ;
@tl1 pair<x:type,y:type> = ;
`,
			"a": `a.t1 = f1:int f2:vector<int>;
a.t2 = n:uint32 f1:tuple<a.t1> f2:tuple_N<a.t1,2>;
`,
			"b": `b.t1 = x:vector<a.t2>;
`,
			"c": `c.t1 = n:uint32 f1?:int;
c.t2 = f1:pair<c.t1,double>;
`,
		},
	)
	assert.True(t, stepSuccess)

	// STEP 4: remove c. and decide to move __common_namespace.
	stepSuccess, nextState = assertMigrationStep(
		t,
		nextState,
		`
vector#1cb5c415 {t:Type} # [t] = Vector t;
tuple#9770768a {t:Type} {n:#} [t] = Tuple t n;
pair {X:Type} {Y:Type} x:X y:Y = Pair X Y;
`,
		"b.,a.,c.,.",
		map[string]string{
			"": `vector<t:type> = []t;
tuple<t:type> = []t;
@tl2ext tuple_N<t:type,n:uint32> = [n]t;
pair<x:type,y:type> = x:x y:y;
`,
			"a": `a.t1 = f1:int f2:vector<int>;
a.t2 = n:uint32 f1:tuple<a.t1> f2:tuple_N<a.t1,2>;
`,
			"b": `b.t1 = x:vector<a.t2>;
`,
			"c": `c.t1 = n:uint32 f1?:int;
c.t2 = f1:pair<c.t1,double>;
`,
		},
	)
	assert.True(t, stepSuccess)
}

func assertMigrationStep(t *testing.T, prevState []tlcodegen.FileToWrite, tl1 string, filter string, allExpectedNamespaces map[string]string) (success bool, nextState []tlcodegen.FileToWrite) {
	ast, err := tlast.ParseTL(tl1)
	assert.NoError(t, err)
	gen, err := tlcodegen.GenerateCode(ast, tlcodegen.Gen2Options{
		ErrorWriter: io.Discard,
		Verbose:     true,

		TL2MigrationFile:       "tmp",
		TL2MigrateByNamespaces: true,
		TL2MigratingWhitelist:  filter,
		TL2ContinuousMigration: true,
	})

	nextState, err = gen.MigrateToTL2(prevState)
	assert.NoError(t, err)
	assert.Equal(t, len(allExpectedNamespaces), len(nextState))

	for ns, value := range allExpectedNamespaces {
		assertNamespaceTranslation(
			t,
			nextState,
			ns,
			value,
		)
	}

	return !t.Failed(), nextState
}

func assertNamespaceTranslation(t *testing.T, files []tlcodegen.FileToWrite, ns string, expectedValue string) {
	if ns == "" {
		ns = "__common_namespace"
	}
	targetFile := tlcodegen.FileToWrite{}
	found := false

	for _, file := range files {
		if strings.HasSuffix(file.Path, fmt.Sprintf("%s.tl2", ns)) {
			found = true
			targetFile = file
			break
		}
	}

	if !found {
		t.Error(fmt.Errorf("can't find file for namespace \"%s\"", ns))
		return
	}

	str := strings.Builder{}
	targetFile.Ast.Print(&str, tlast.NewCanonicalFormatOptions())

	assert.Equal(t, expectedValue, str.String())
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
				ErrorWriter:       io.Discard,
				Verbose:           true,
				LinterPHPCheck:    true,
				WarningsAreErrors: true,
			})

			assert.Error(t, err)
			assert.Equal(t, "can't have boxed reference in field to flat type due to php generator issues (instance: myTestFunction)", err.Error())
		})

		t.Run("no fail when all flat types are bare", func(t *testing.T) {
			data := `vector2#1cb5c415 {t:Type} n:# m:n*[t] = Vector2 t;
myType#12345678 x:vector2<int> = MyType;
---functions---
@read myTestFunction x:myType = MyType;`

			ast, err := tlast.ParseTL(data)
			if err != nil {
				t.Error(err)
				return
			}

			_, err = tlcodegen.GenerateCode(ast, tlcodegen.Gen2Options{
				ErrorWriter:       io.Discard,
				Verbose:           true,
				LinterPHPCheck:    true,
				WarningsAreErrors: true,
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
						ErrorWriter:       io.Discard,
						Verbose:           true,
						LinterPHPCheck:    true,
						WarningsAreErrors: true,
					})

					if s == "True" {
						assert.Error(t, err)
					} else {
						assert.NoError(t, err)
					}
				})
			}
		})
	})

	t.Run("Issue template flat types", func(t *testing.T) {
		t.Run("fail on non-special type with template", func(t *testing.T) {
			data := `int#a8509bda ? = Int;
vector#12345679 {t:Type} # [t] = Vector t;
myType#12345678 {T:Type} x:vector<T> = MyType T;
---functions---
@read myTestFunction x:myType<int> = MyType<int>;`

			ast, err := tlast.ParseTL(data)
			if err != nil {
				t.Error(err)
				return
			}

			_, err = tlcodegen.GenerateCode(ast, tlcodegen.Gen2Options{
				ErrorWriter:       io.Discard,
				Verbose:           true,
				LinterPHPCheck:    true,
				WarningsAreErrors: true,
			})

			assert.Error(t, err)
			assert.Equal(t, "flat types can't have type templates due to php generator issues", err.Error())
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
				ErrorWriter:       io.Discard,
				Verbose:           true,
				LinterPHPCheck:    true,
				WarningsAreErrors: true,
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
						ErrorWriter:       io.Discard,
						Verbose:           true,
						LinterPHPCheck:    true,
						WarningsAreErrors: true,
					})

					if s == "True" {
						assert.Error(t, err)
					} else {
						assert.NoError(t, err)
					}
				})
			}
		})
	})

	t.Run("Issue with non-polymorphic types", func(t *testing.T) {
		t.Run("fail case", func(t *testing.T) {
			data := `int#a8509bda ? = Int;
vector#12345679 {t:Type} # [t] = Vector t;
myType x:vector<int> y:vector<int> = MyType;
---functions---
@read myTestFunction x:MyType = MyType;`

			ast, err := tlast.ParseTL(data)
			if err != nil {
				t.Error(err)
				return
			}

			_, err = tlcodegen.GenerateCode(ast, tlcodegen.Gen2Options{
				ErrorWriter:                     io.Discard,
				Verbose:                         true,
				LinterPHPCheck:                  true,
				LinterPHPNonPolymorphicBoxedRef: true,
				WarningsAreErrors:               true,
			})

			assert.Error(t, err)
			assert.Equal(t, "can't boxed reference type with a single constructor with the same name in field due to php generator issues", err.Error())
		})

		t.Run("correct case", func(t *testing.T) {
			data := `int#a8509bda ? = Int;
vector#12345679 {t:Type} # [t] = Vector t;
myType2 x:vector<int> y:vector<int> = MyType; // tlgen:nolint
---functions---
@read myTestFunction x:MyType = MyType;`

			ast, err := tlast.ParseTL(data)
			if err != nil {
				t.Error(err)
				return
			}

			_, err = tlcodegen.GenerateCode(ast, tlcodegen.Gen2Options{
				ErrorWriter:                     io.Discard,
				Verbose:                         true,
				LinterPHPCheck:                  true,
				LinterPHPNonPolymorphicBoxedRef: true,
				WarningsAreErrors:               true,
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
					if s == "#" {
						return
					}
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

---functions---
@read myTestFunction x:[[[TYPE_HERE]]] = [[[TYPE_HERE]]];`

					ast, err := tlast.ParseTL(strings.Replace(data, "[[[TYPE_HERE]]]", s+argsToAdd[s], -1))
					if err != nil {
						t.Error(err)
						return
					}

					_, err = tlcodegen.GenerateCode(ast, tlcodegen.Gen2Options{
						ErrorWriter:       io.Discard,
						Verbose:           true,
						LinterPHPCheck:    true,
						WarningsAreErrors: true,
					})

					if s == "True" {
						assert.Error(t, err)
					} else {
						assert.NoError(t, err)
					}
				})
			}
		})
	})
}
