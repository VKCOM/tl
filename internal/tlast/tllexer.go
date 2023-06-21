// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlast

import (
	"fmt"
	"strings"
)

const (
	typesSection     = -1
	functionsSection = -2
	crc32hash        = -3
	functionModifier = -4
	numberSign       = -5
	number           = -6
	comment          = -7
	undefined        = -8
	lcIdent          = -9
	ucIdent          = -10
	lcIdentNS        = -11
	ucIdentNS        = -12
	eof              = -13 // so we always have next token in array
)

const (
	lRoundBracket  = '('
	rRoundBracket  = ')'
	lSquareBracket = '['
	rSquareBracket = ']'
	lCurlyBracket  = '{'
	rCurlyBracket  = '}'
	lAngleBracket  = '<'
	rAngleBracket  = '>'
	colon          = ':'
	semiColon      = ';'
	dotSign        = '.'
	commaSign      = ','
	percentSign    = '%'
	whiteSpace     = ' '
	tab            = '\t'
	newLine        = '\n'
	equalSign      = '='
	questionMark   = '?'
	asterisk       = '*'
	plus           = '+'
	exclamation    = '!'
)

// TODO - support windows line endings (by skipping '\r')

const (
	typesSectionString     = "---types---"
	functionsSectionString = "---functions---"
)

type Position struct {
	fileContent     string
	file            string
	line            int
	column          int
	startLineOffset int
	offset          int
}

type PositionRange struct {
	Outer Position // in lexer, it is empty, in parser set to start of combinator
	Begin Position
	End   Position
}

func (p Position) String() string {
	return fmt.Sprintf("%s %d:%d(%d..%d)", p.file, p.line, p.column, p.startLineOffset, p.offset)
}

type token struct {
	tokenType int
	val       string
	pos       Position
}

func lowerCase(c byte) bool { return 'a' <= c && c <= 'z' }

func upperCase(c byte) bool { return 'A' <= c && c <= 'Z' }

func digit(c byte) bool { return '0' <= c && c <= '9' }

func letter(c byte) bool { return lowerCase(c) || upperCase(c) }

func identChar(c byte) bool { return letter(c) || digit(c) || c == '_' }

func hex(c byte) bool { return digit(c) || ('a' <= c && c <= 'f') }

func nameIdent(s string) string {
	if len(s) == 0 || !letter(s[0]) {
		return ""
	}
	i := 1
	for ; i < len(s) && identChar(s[i]); i++ {
	}
	return s[:i]
}

func builtinIdent(s string) string {
	if len(s) == 0 || !(letter(s[0]) || s[0] == '_') {
		return ""
	}
	i := 1
	for ; i < len(s) && identChar(s[i]); i++ {
	}
	return s[:i]
}

func numberLexeme(s string) (string, bool) {
	i := 0
	allDigits := true
	for ; i < len(s) && identChar(s[i]); i++ {
		allDigits = allDigits && digit(s[i])
	}
	return s[:i], allDigits
}

type lexer struct {
	allowBuiltin bool
	str          string // iterator-like
	tokens       []token

	position Position
}

func newLexer(s, file string, allowBuiltin bool) lexer {
	return lexer{allowBuiltin, s, make([]token, 0, len(s)/3), Position{s, file, 1, 1, 0, 0}}
}

// when error is returned, undefined token is added to tokens
func (l *lexer) generateTokens() ([]token, error) {
	for l.str != "" {
		err := l.nextToken()
		if err != nil {
			return l.tokens, err
		}
	}
	l.advance(0, eof)
	return l.tokens, nil
}

func (l *lexer) recombineTokens() string {
	var b strings.Builder
	for _, tok := range l.tokens {
		b.WriteString(tok.val)
	}
	b.WriteString(l.str)
	return b.String()
}

func (l *lexer) advance(len int, tokenType int) token {
	tok := token{tokenType, l.str[:len], l.position}
	l.position.column += len
	l.position.offset += len
	l.str = l.str[len:]
	l.tokens = append(l.tokens, tok)
	return tok
}

func (l *lexer) checkPrimitive() bool {
	c := l.str[0]
	switch c {
	case lRoundBracket, rRoundBracket,
		lSquareBracket, rSquareBracket,
		lCurlyBracket, rCurlyBracket,
		lAngleBracket, rAngleBracket,
		dotSign, plus, asterisk, exclamation,
		colon, semiColon, whiteSpace,
		tab, equalSign, questionMark, percentSign,
		commaSign:
		l.advance(1, int(c))
		return true
	case newLine:
		l.advance(1, int(c))
		l.position.line++
		l.position.column = 1
		l.position.startLineOffset = l.position.offset
		return true
	default:
		return false
	}
}

func (l *lexer) nextToken() error {
	switch {
	case l.checkPrimitive():
		return nil
	case l.str[0] == '@':
		return l.lexFunctionModifier()
	case l.str[0] == '/':
		return l.lexComment()
	case l.str[0] == '-':
		return l.lexSection()
	case l.str[0] == '#':
		return l.lexNumberSign()
	case l.str[0] == '_':
		if l.allowBuiltin {
			w := builtinIdent(l.str)
			if w == "_" {
				l.advance(len(w), ucIdent) // for TypeDecls that do not exist
			} else {
				l.advance(len(w), lcIdent) // for constructor names
			}
			return nil
		}
		tok := l.advance(1, undefined)
		return parseErrToken(fmt.Errorf("identifier cannot start with underscore"), tok, tok.pos)
	case digit(l.str[0]):
		return l.lexNumber()
	case letter(l.str[0]):
		return l.lexLexeme()
	default:
		tok := l.advance(1, undefined)
		return parseErrToken(fmt.Errorf("undefined symbol: %q", tok.val), tok, tok.pos)
	}
}

func (l *lexer) lexFunctionModifier() error {
	w := nameIdent(l.str[1:])
	if w == "" || !lowerCase(w[0]) {
		tok := l.advance(1+len(w), undefined)
		return parseErrToken(fmt.Errorf("combinator modifier should start from lower case letter"), tok, tok.pos)
	}
	l.advance(1+len(w), functionModifier)
	return nil
}

func (l *lexer) lexComment() error {
	if len(l.str) <= 1 {
		tok := l.advance(len(l.str), undefined)
		return parseErrToken(fmt.Errorf("'//' expected as a comment start"), tok, tok.pos)
	}
	switch l.str[1] {
	case '\n':
		tok := l.advance(1, undefined)
		return parseErrToken(fmt.Errorf("'//' expected as a comment start"), tok, tok.pos)
	case '/': // separate case, otherwise beautiful error printing breaks, because token will contain a newline
		index := strings.IndexByte(l.str, '\n')
		if index < 0 {
			index = len(l.str)
		}
		l.advance(index, comment)
		return nil
	case '*':
		// from index 2 to not allow /*/
		closeBlockCommentIndex := strings.Index(l.str[2:], "*/")
		if closeBlockCommentIndex < 0 {
			tok := l.advance(2, undefined)
			return parseErrToken(fmt.Errorf("multiline comment start without closing '*/'"), tok, tok.pos)
		}
		commentBlock := l.advance(2+closeBlockCommentIndex+2, comment) // /* ... */
		nlIndex := strings.LastIndexByte(commentBlock.val, '\n')
		if nlIndex >= 0 {
			l.position.line += strings.Count(commentBlock.val, "\n")
			l.position.column = 1 + len(commentBlock.val) - (nlIndex + 1) // 1 - 1 cancel out
			l.position.startLineOffset = commentBlock.pos.offset + (nlIndex + 1)
		}
		return nil
	}
	tok := l.advance(2, undefined)
	return parseErrToken(fmt.Errorf("'//' expected as a comment start"), tok, tok.pos)
}

func (l *lexer) lexSection() error {
	switch {
	case strings.HasPrefix(l.str, typesSectionString):
		l.advance(len(typesSectionString), typesSection)
	case strings.HasPrefix(l.str, functionsSectionString):
		l.advance(len(functionsSectionString), functionsSection)
	default:
		l.advance(1, '-')
	}
	return nil
}

func (l *lexer) lexNumberSign() error {
	i := 1
	allDigits := true
	for ; i < len(l.str) && identChar(l.str[i]); i++ {
		allDigits = allDigits && hex(l.str[i])
	}
	if i == 1 {
		l.advance(i, numberSign)
	} else {
		if !allDigits || i != 1+8 {
			tok := l.advance(i, undefined)
			return parseErrToken(fmt.Errorf("expect tag with exactly 8 lowercase hex digits here"), tok, tok.pos)
		}
		l.advance(i, crc32hash)
	}
	return nil
}

func (l *lexer) lexNumber() error {
	n, allDigits := numberLexeme(l.str)
	if !allDigits {
		tok := l.advance(len(n), undefined)
		return parseErrToken(fmt.Errorf("expect only decimal digits in number here"), tok, tok.pos)
	}
	l.advance(len(n), number)
	return nil
}

func (l *lexer) lexLexeme() error {
	ns := ""
	w := nameIdent(l.str)
	if len(l.str) > len(w) && l.str[len(w)] == dotSign {
		w2 := nameIdent(l.str[len(w)+1:])
		if w2 != "" {
			ns = w + "."
			w = w2
			if !lowerCase(ns[0]) { // Hren.vam || Hren.Vam
				tok := l.advance(len(ns)+len(w), undefined)
				return parseErrToken(fmt.Errorf("namespace identifier should start from lower case letter"), tok, tok.pos)
			}
			// hren.vam ||  hren.Vam
			if lowerCase(w[0]) {
				l.advance(len(ns)+len(w), lcIdentNS)
			} else {
				l.advance(len(ns)+len(w), ucIdentNS)
			}
			return nil
		}
	}
	if lowerCase(w[0]) {
		l.advance(len(w), lcIdent)
	} else {
		l.advance(len(w), ucIdent)
	}
	return nil
}
