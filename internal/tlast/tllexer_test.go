// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlast

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func countToken(tokens []token, tokenType int) int {
	c := 0
	for _, t := range tokens {
		if t.tokenType == tokenType {
			c++
		}
	}
	return c
}

func TestLexer(t *testing.T) {
	//combinedBytes, err := os.ReadFile("test_combined.tl")
	combinedBytes := `
---types---

// Builtin types
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
`
	var err error
	require.NoError(t, err)
	t.Run("Full file", func(t *testing.T) {
		str := combinedBytes
		lex := newLexer(str, "", false)
		tokens, _ := lex.generateTokens(false) // TODO - what if err?
		require.Equal(t, 0, countToken(tokens, undefined))
		recombined := lex.recombineTokens()
		require.Equal(t, str, recombined)
	})

	t.Run("Empty file", func(t *testing.T) {
		str := ""
		lex := newLexer(str, "", false)
		_, _ = lex.generateTokens(false)
		recombined := lex.recombineTokens()
		require.Equal(t, str, recombined)
	})

	t.Run("Upper case in tag", func(t *testing.T) {
		str := "foo#1234567F = Foo;"
		lex := newLexer(str, "", false)
		_, err = lex.generateTokens(false)
		require.EqualError(t, err, "expect tag with exactly 8 lowercase hex digits here")
	})

	t.Run("From random to random byte", func(t *testing.T) {
		for j := 1; j <= 400; j++ {
			from := rand.Intn(len(combinedBytes))
			to := from + rand.Intn(len(combinedBytes)-from)
			str := combinedBytes[from:to]
			lex := newLexer(str, "", false)
			_, _ = lex.generateTokens(false)
			recombined := lex.recombineTokens()
			require.Equal(t, str, recombined)
		}
	})
}
