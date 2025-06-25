package tlast

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTL2Lexer(t *testing.T) {
	t.Run("Check simple definition", func(t *testing.T) {
		testCorrectCase(t, `@read @write typeName = x?:int y:double; // some comment`)
	})
	t.Run("Check union definition", func(t *testing.T) {
		testCorrectCase(t, `@read @write typeName = int | Green | MyConstructor x:int; // some comment
`)
	})
	t.Run("Check multiline is banned", func(t *testing.T) {
		testIncorrectCase(t, `@read @write typeName = int | Green | MyConstructor x:int; /* some comment */
`)
	})
	t.Run("Check some symbols are banned", func(t *testing.T) {
		symbols := []string{"{", "}", "(", ")", "+", "*", "!", "%"}
		for _, symbol := range symbols {
			t.Run(symbol, func(t *testing.T) {
				testIncorrectCase(t, symbol)
			})
		}
	})
	t.Run("Check underscore", func(t *testing.T) {
		testCorrectCase(t, "_")
	})
}

func testCorrectCase(t *testing.T, str string) {
	lex := newLexer(str, "", LexerOptions{LexerLanguage: tl2})
	tokens, err := lex.generateTokens()
	require.NoError(t, err)
	require.Equal(t, 0, countToken(tokens, undefined))
	recombined := lex.recombineTokens()
	require.Equal(t, str, recombined)
}

func testIncorrectCase(t *testing.T, str string) {
	lex := newLexer(str, "", LexerOptions{LexerLanguage: tl2})
	_, err := lex.generateTokens()
	require.Error(t, err)
}
