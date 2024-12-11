// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlast

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// Please add methods in the same order types defined in tlparser.go
// All parsing, except TypeRef, which is in separate file tlparser_typeref.go

type tokenIterator struct {
	tokens []token
	offset int
}

func (it *tokenIterator) front() token {
	return it.tokens[it.offset]
}

func (it *tokenIterator) count() int {
	return len(it.tokens) - it.offset
}

func (it *tokenIterator) popFront() token {
	it.offset++
	return it.tokens[it.offset-1]
}

func (it *tokenIterator) skipWS(outer Position) PositionRange {
	for ; it.count() != 0; it.popFront() {
		switch tok := it.front(); tok.tokenType {
		case comment, whiteSpace, tab, newLine:
			continue
		default:
			return PositionRange{Outer: outer, Begin: tok.pos, End: tok.pos}
		}
	}
	if it.count() == 0 {
		log.Panicf("tokenizer invariant failed, no eof token, %s", ContactAuthorsString)
	}
	return PositionRange{}
}

// returns if newline was found
func (it *tokenIterator) skipToNewline() bool {
	for ; it.count() != 0; it.popFront() {
		switch tok := it.front(); tok.tokenType {
		case comment, whiteSpace, tab:
			continue
		case newLine, eof:
			return true
		default:
			return false
		}
	}
	return false
}

func (it *tokenIterator) checkToken(i int) bool {
	it.skipWS(Position{})
	return it.front().tokenType == i
}

func (it *tokenIterator) expect(i int) bool {
	if it.checkToken(i) {
		it.popFront()
		return true
	}
	return false
}

func (it *tokenIterator) expectOrPanic(i int) {
	if !it.expect(i) {
		log.Panicf("unexpected token during parsing, %s", ContactAuthorsString)
	}
}

func splitIdenNSFromToken(s string) Name {
	i := strings.Index(s, ".")
	if i < 0 {
		log.Panicf("identifier with namespace has no dot, %s", ContactAuthorsString)
	}
	return Name{
		Namespace: s[:i],
		Name:      s[i+1:],
	}
}

func parseLCIdentNS(tokens tokenIterator, outer Position) (Name, tokenIterator, error) {
	rest := tokens
	var res Name
	if rest.checkToken(lcIdentNS) {
		res = splitIdenNSFromToken(rest.front().val)
		rest.expectOrPanic(lcIdentNS)
		return res, rest, nil
	}
	if rest.checkToken(lcIdent) {
		res.Name = rest.front().val
		rest.expectOrPanic(lcIdent)
		return res, rest, nil
	}
	return Name{}, tokens, parseErrToken(fmt.Errorf("low-case name (with optional namespace) expected"), rest.front(), outer)
}

func parseUCIdentNS(tokens tokenIterator, outer Position) (Name, tokenIterator, error) {
	rest := tokens
	var res Name
	if rest.checkToken(ucIdentNS) {
		res = splitIdenNSFromToken(rest.front().val)
		rest.expectOrPanic(ucIdentNS)
		return res, rest, nil
	}
	if rest.checkToken(ucIdent) {
		res.Name = rest.front().val
		rest.expectOrPanic(ucIdent)
		return res, rest, nil
	}
	return Name{}, tokens, parseErrToken(fmt.Errorf("upper-case name (with optional namespace) expected"), rest.front(), outer)
}

func parseVarIdent(tokens tokenIterator, outer Position) (string, tokenIterator, error) {
	rest := tokens
	if rest.checkToken(lcIdent) {
		res := rest.front().val
		rest.expectOrPanic(lcIdent)
		return res, rest, nil
	}
	if rest.checkToken(ucIdent) {
		res := rest.front().val
		rest.expectOrPanic(ucIdent)
		return res, rest, nil
	}
	return "", tokens, parseErrToken(fmt.Errorf("name without namespace expected"), rest.front(), outer)
}

// functionModifier := @any  | @internal  | @kphp | @read | @readwrite | @write
func parseModifiers(tokens tokenIterator, outer Position) ([]Modifier, tokenIterator) {
	var res []Modifier
	rest := tokens
	for {
		mod := Modifier{PR: rest.skipWS(outer)}
		if !rest.checkToken(functionModifier) {
			break
		}
		mod.Name = rest.front().val[1:] // always safe
		rest.expectOrPanic(functionModifier)
		mod.PR.End = rest.front().pos
		res = append(res, mod)
	}
	return res, rest
}

// constructor := fullName [CRC32]
func parseConstructor(tokens tokenIterator, outer Position, allowBuiltin bool) (Constructor, tokenIterator, error) {
	rest := tokens
	res := Constructor{NamePR: rest.skipWS(outer)}
	var err error
	if allowBuiltin && rest.checkToken(numberSign) {
		res.Name.Name = rest.front().val
		rest.expectOrPanic(numberSign)
		res.NamePR.End = rest.front().pos
		res.IDPR = res.NamePR // wish to highlight name, if tag absent
	} else {
		if res.Name, rest, err = parseLCIdentNS(rest, outer); err != nil {
			return Constructor{}, tokens, err // parseErrToken(fmt.Errorf("constructor name expected"), rest.front())
		}
		res.NamePR.End = rest.front().pos
		res.IDPR = res.NamePR // wish to highlight name, if tag absent
	}
	if rest.checkToken(crc32hash) {
		res.IDPR.Begin = rest.front().pos
		i, err := strconv.ParseUint(rest.front().val[1:], 16, 32)
		if err != nil {
			return Constructor{}, tokens, parseErrToken(fmt.Errorf("error converting constructor tag to uint32: %w", err), rest.front(), outer)
		}
		x := uint32(i)
		res.ID = &x
		rest.expectOrPanic(crc32hash)
		res.IDPR.End = rest.front().pos
	}
	return res, rest, nil
}

func parseTemplateArguments(tokens tokenIterator, outer Position) ([]TemplateArgument, tokenIterator, error) {
	var res []TemplateArgument
	rest := tokens
	var err error
	for {
		var of *TemplateArgument
		of, rest, err = parseTemplateArgument(rest, outer)
		if err != nil {
			return nil, tokens, err
		}
		if of == nil {
			break
		}
		res = append(res, *of)
	}
	return res, rest, nil
}

// templateArgument := '{' fieldName T '}'
func parseTemplateArgument(tokens tokenIterator, outer Position) (*TemplateArgument, tokenIterator, error) {
	rest := tokens
	var err error
	of := &TemplateArgument{PR: rest.skipWS(outer)}
	if !rest.expect(lCurlyBracket) {
		return nil, tokens, nil
	}
	of.FieldName, rest, err = parseVarIdent(rest, outer)
	if err != nil {
		return nil, tokens, fmt.Errorf("template argument name expected: %w", err)
	}
	if !rest.expect(colon) {
		return nil, tokens, parseErrToken(fmt.Errorf("':' after template argument name expected"), rest.front(), outer)
	}
	switch {
	case rest.checkToken(ucIdent) && rest.front().val == "Type":
		of.IsNat = false
		rest.expectOrPanic(ucIdent)
	case rest.checkToken(numberSign):
		of.IsNat = true
		rest.expectOrPanic(numberSign)
	default:
		return nil, tokens, parseErrToken(fmt.Errorf("template argument type can be either 'Type' or '#'"), rest.front(), outer)
	}
	if !rest.expect(rCurlyBracket) {
		return nil, tokens, parseErrToken(fmt.Errorf("'}' after template argument type expected"), rest.front(), outer)
	}
	of.PR.End = rest.front().pos
	return of, rest, nil
}

// fullName [word] ...
func parseTypeDeclaration(tokens tokenIterator, outer Position) (TypeDeclaration, tokenIterator, error) {
	rest := tokens
	res := TypeDeclaration{PR: rest.skipWS(outer)}
	res.NamePR = res.PR
	var name Name
	var err error
	if name, rest, err = parseUCIdentNS(rest, outer); err != nil {
		return TypeDeclaration{}, tokens, err
	}
	res.Name = name
	res.NamePR.End = rest.front().pos
	for {
		var argName string
		argPR := rest.skipWS(outer)
		argName, rest, err = parseVarIdent(rest, outer)
		if err != nil {
			break
		}
		argPR.End = rest.front().pos
		res.Arguments = append(res.Arguments, argName)
		res.ArgumentsPR = append(res.ArgumentsPR, argPR)
	}
	res.PR.End = rest.front().pos
	return res, rest, nil
}

// TODO: make this grammar correct
// arithmetic :=  arithmetic '+' arithmetic | number | '(' arithmetic ')'
func parseArithmetic(tokens tokenIterator, outer Position, force bool) (*Arithmetic, tokenIterator, error) {
	rest := tokens
	var err error
	var res *Arithmetic
	switch {
	case rest.expect(lRoundBracket):
		res, rest, err = parseArithmetic(rest, outer, force)
		if err != nil {
			return nil, tokens, err
		}
		if res == nil {
			return nil, tokens, nil
		}
		if !rest.expect(rRoundBracket) {
			return nil, tokens, parseErrToken(fmt.Errorf("')' expected"), rest.front(), outer)
		}
	case rest.checkToken(number):
		var i uint64
		i, err = strconv.ParseUint(rest.front().val, 10, 32)
		if err != nil {
			return nil, tokens, parseErrToken(fmt.Errorf("constant overflows uint32: %w", err), rest.front(), outer)
		}
		rest.expectOrPanic(number)
		res = &Arithmetic{
			Nums: []uint32{uint32(i)},
			Res:  uint32(i),
		}
	default:
		if force {
			return nil, tokens, parseErrToken(fmt.Errorf("arithmetic expression expected after '+'"), rest.front(), outer)
		}
		return nil, tokens, nil
	}
	for rest.expect(plus) {
		var res2 *Arithmetic
		res2, rest, err = parseArithmetic(rest, outer, true)
		if err != nil {
			return nil, tokens, err
		}
		sum := uint64(res.Res) + uint64(res2.Res)
		if sum >= math.MaxUint32 {
			return nil, tokens, parseErrToken(fmt.Errorf("arithmetic expression overflows uint32"), rest.front(), outer)
		}
		res.Res = uint32(sum)
		res.Nums = append(res.Nums, res2.Nums...)
	}
	return res, rest, err
}

// aot := [T | arithmetic]
func parseArithmeticOrTypeOpt(tokens tokenIterator, applyFlag bool, outer Position) (*ArithmeticOrType, tokenIterator, error) {
	rest := tokens
	var a *Arithmetic
	pr := rest.skipWS(outer)
	var err error
	a, rest, err = parseArithmetic(rest, outer, false)
	if err != nil {
		return nil, tokens, err
	}
	if a != nil {
		t := TypeRef{PR: pr}        // TODO - move into parsing of Arithmetic
		t.PR.End = rest.front().pos // t stores PR in arithmeticOrType
		return &ArithmeticOrType{T: t, IsArith: true, Arith: *a}, rest, nil
	}
	var t *TypeRef
	t, rest, err = parseTypeRef(rest, applyFlag, true, outer)
	if err != nil {
		return nil, tokens, err
	}
	if t != nil {
		return &ArithmeticOrType{T: *t}, rest, nil
	}
	return nil, tokens, nil
}

// repeatType := [ scale '*' ]
func parseScaleFactorOpt(tokens tokenIterator, outer Position) (*ScaleFactor, tokenIterator, error) {
	rest := tokens
	res := ScaleFactor{PR: rest.skipWS(outer)}
	var err error
	res.Scale, rest, err = parseVarIdent(rest, outer)
	if err == nil {
		res.PR.End = rest.front().pos
		return &res, rest, nil
		// return ScaleFactor{}, tokens, nil // fmt.Errorf("scale factor error: %w", err) // TODO - better error
	}
	var arith *Arithmetic
	arith, rest, err = parseArithmetic(rest, outer, false)
	// can't parse arithmetic without error
	if err != nil {
		return nil, tokens, err
	}
	if arith != nil {
		res = ScaleFactor{IsArith: true, Arith: *arith}
		return &res, rest, nil
	}
	res.PR.End = rest.front().pos
	return nil, rest, nil // TODO - check
}

// repeatType := [ scale '*' ] '[' field ... ']'
func parseRepeatWithScaleOpt(tokens tokenIterator, outer Position) (*RepeatWithScale, tokenIterator, error) {
	rest := tokens
	res := RepeatWithScale{PR: rest.skipWS(outer)}
	var err error
	var scale *ScaleFactor
	scale, rest, err = parseScaleFactorOpt(rest, outer)
	if err != nil {
		return nil, tokens, fmt.Errorf("error parsing multiplier: %w", err)
	}
	if scale != nil {
		res.Scale = *scale
		if !rest.expect(asterisk) {
			return nil, tokens, nil
		}
		res.ExplicitScale = true
		if !rest.expect(lSquareBracket) {
			return nil, tokens, parseErrToken(fmt.Errorf("'[' is expected after '*'"), rest.front(), outer)
		}
	} else if !rest.expect(lSquareBracket) {
		return nil, tokens, nil
	}
	// rBracketToken := rest.front()
	_, res.Rep, rest, err = parseFields(rest, rSquareBracket, rSquareBracket, outer)
	if err != nil {
		return nil, tokens, err // fmt.Errorf("error parsing fields in square brackets: %w", err)
	}
	// if res.Rep == nil {
	//	return RepeatWithScale{}, nil, parseErrToken(fmt.Errorf("empty square brackets not allowed: "), rBracketToken, outer)
	// }
	res.PR.End = rest.front().pos
	return &res, rest, nil
}

// fieldMask := word '.' number '?'
func parseFieldMask(tokens tokenIterator, outer Position) (*FieldMask, tokenIterator, error) {
	rest := tokens
	res := &FieldMask{PRName: rest.skipWS(outer)}
	var name string
	var err error
	if name, rest, err = parseVarIdent(rest, outer); err != nil {
		return nil, tokens, nil
	}
	res.MaskName = name
	res.PRName.End = rest.front().pos
	if !rest.expect(dotSign) {
		return nil, tokens, nil
	}
	res.PRBits = rest.skipWS(outer)
	if !rest.checkToken(number) {
		return nil, tokens, parseErrToken(fmt.Errorf("expecting decimal bitmask bit number"), rest.front(), outer)
	}
	i, err := strconv.ParseUint(rest.front().val, 10, 32)
	if err != nil {
		return nil, tokens, parseErrToken(fmt.Errorf("error converting bitmask to uint32: %w", err), rest.front(), outer)
	}
	res.BitNumber = uint32(i)
	rest.expectOrPanic(number)
	res.PRBits.End = rest.front().pos
	if !rest.expect(questionMark) {
		return nil, tokens, parseErrToken(fmt.Errorf("'?' expected after field bitmask "), rest.front(), outer)
	}
	return res, rest, nil
}

// fieldName := name ':'
func parseFieldName(tokens tokenIterator, outer Position) (string, tokenIterator, Position) {
	rest := tokens
	var name string
	var err error
	if name, rest, err = parseVarIdent(rest, outer); err != nil {
		return "", rest, Position{}
	}
	end := rest.front().pos
	if !rest.expect(colon) {
		return "", tokens, end
	}
	return name, rest, end
}

// field := [ fieldName ] [ ! ]     [ fieldMask ] T
func parseField(commentStart tokenIterator, tokens tokenIterator, outer Position) (Field, tokenIterator, error) {
	rest := tokens
	res := Field{PR: rest.skipWS(outer)}
	res.CommentBefore = parseCommentBefore(commentStart, rest)
	res.PRName = res.PR
	var err error
	res.FieldName, rest, res.PRName.End = parseFieldName(rest, outer)
	var mask *FieldMask
	mask, rest, err = parseFieldMask(rest, outer)
	if err != nil {
		return Field{}, rest, err
	}
	if mask != nil {
		res.Mask = mask
	}
	if rest.expect(exclamation) {
		res.Excl = true
	}
	var rws *RepeatWithScale
	rws, rest, err = parseRepeatWithScaleOpt(rest, outer)
	if err != nil {
		return Field{}, tokens, err
	}
	if rws != nil {
		res.IsRepeated = true
		res.ScaleRepeat = *rws
		res.FieldType.PR = rws.PR // after type replacing, it will become new field
		res.PR.End = rest.front().pos
		return res, rest, nil
	}
	// typeStartPR := rest.skipWS(outer) - experimental beautiful error code below
	var t *TypeRef
	t, rest, err = parseTypeRef(rest, false, true, outer)
	if err != nil {
		// berr := err.(*ParseError)
		// typeStartPR.End = berr.Pos.Begin
		// return Field{}, tokens, BeautifulError2(berr, typeStartPR.BeautifulError(fmt.Errorf("field type here")))
		return Field{}, tokens, fmt.Errorf("error parsing field type: %w", err)
	}
	if t == nil {
		return Field{}, tokens, parseErrToken(fmt.Errorf("field type is expected here (missed '()' around complex type?)"), rest.front(), outer)
	}
	res.FieldType = *t
	res.PR.End = rest.front().pos
	return res, rest, nil
}

func parseFields(tokens tokenIterator, finishToken1 int, finishToken2 int, outer Position) (int, []Field, tokenIterator, error) {
	var res []Field
	rest := tokens
	var err error
	commentStart := tokens
	for {
		switch {
		case rest.checkToken(finishToken1):
			rest.popFront()
			return finishToken1, res, rest, nil
		case rest.checkToken(finishToken2):
			rest.popFront()
			return finishToken2, res, rest, nil
		}
		var field Field
		field, rest, err = parseField(commentStart, rest, outer)
		if err != nil {
			return 0, nil, tokens, err // fmt.Errorf("'%c' or field declaration expected: %w", finishToken, err) // parseErrToken(fmt.Errorf("field declaration expected:"), rest.front())
		}
		commentStart = rest
		if rest.skipToNewline() {
			field.CommentRight = parseCommentRight(commentStart, rest)
		}
		res = append(res, field)
		commentStart = rest
	}
}

// funcDecl := apply | '(' apply ')' | '%' T | fullName '<' aot ',' ... '>
func parseFuncDecl(tokens tokenIterator, outer Position) (TypeRef, tokenIterator, error) {
	// Separate function for better documentation
	rest := tokens
	t, rest, err := parseTypeRef(rest, true, false, outer)
	if err != nil {
		return TypeRef{}, tokens, err
	}
	if t == nil {
		return TypeRef{}, tokens, parseErrToken(fmt.Errorf("return type is expected here"), rest.front(), outer)
	}
	return *t, rest, nil
}

// type := constructor [templateArgument] ... [field] ...  '=' typeDecl ';'
// or
// function := [ functionModifier ] constructor [templateArgument] ... [field] ... '=' apply;'
// but also => marks function independent of section
func parseCombinator(commentStart tokenIterator, tokens tokenIterator, isFunction bool, allowBuiltin bool) (Combinator, tokenIterator, error) {
	rest := tokens

	td := Combinator{PR: rest.skipWS(Position{})}
	td.CommentBefore = parseCommentBefore(commentStart, rest)

	outer := td.PR.Begin
	td.PR.Outer = outer // Set outer context for all parsing
	var err error
	td.Modifiers, rest = parseModifiers(rest, outer) // we support modifiers for normal types also
	// list of modifiers is checked in generator, where they are used for some purposes
	td.Construct, rest, err = parseConstructor(rest, outer, allowBuiltin)
	if err != nil {
		return Combinator{}, tokens, fmt.Errorf("constructor declaration error: %w", err)
	}
	td.TemplateArgumentsPR = rest.skipWS(outer)
	td.TemplateArguments, rest, err = parseTemplateArguments(rest, outer)
	if err != nil {
		return Combinator{}, tokens, err
	}
	if len(td.TemplateArguments) == 0 {
		td.TemplateArgumentsPR.Begin = td.Construct.IDPR.End // highlight empty space when no arguments
	}
	td.TemplateArgumentsPR.End = rest.front().pos
	if rest.checkToken(questionMark) {
		if isFunction {
			return Combinator{}, tokens, parseErrToken(fmt.Errorf("'?' (legacy builtin type body) is not allowed in functions"), rest.front(), outer)
		}
		rest.expectOrPanic(questionMark)
		td.Builtin = true
		if !rest.expect(equalSign) {
			return Combinator{}, tokens, parseErrToken(fmt.Errorf("'=' expected after '?' (legacy builtin type body)"), rest.front(), outer)
		}
	} else {
		var actualModifier int
		actualModifier, td.Fields, rest, err = parseFields(rest, equalSign, functionSign, outer)
		if actualModifier == functionSign {
			isFunction = true
		}
		if err != nil {
			return Combinator{}, tokens, err
		}
	}
	if isFunction {
		td.FuncDecl, rest, err = parseFuncDecl(rest, outer)
		if err != nil {
			return Combinator{}, tokens, err // TODO - check
		}
		td.IsFunction = true
	} else {
		if td.TypeDecl, rest, err = parseTypeDeclaration(rest, outer); err != nil {
			return Combinator{}, tokens, fmt.Errorf("type declaration expected: %w", err)
		}
		if td.TypeDecl.Name.String() == "_" {
			td.Builtin = true
		}
	}
	if !rest.expect(semiColon) {
		return Combinator{}, tokens, parseErrToken(fmt.Errorf("';' or type argument expected"), rest.front(), outer)
	}
	td.PR.End = rest.front().pos
	commentStart = rest
	if rest.skipToNewline() {
		td.CommentRight = parseCommentRight(commentStart, rest)
	}

	return td, rest, nil
}

func ParseTL(str string) (TL, error) {
	return ParseTLFile(str, "", LexerOptions{AllowMLC: true}, os.Stdout)
}

// ParseTL2 TL := TypesSection [ type ... ] FunctionSection [ function ... ]
func ParseTLFile(str, file string, opts LexerOptions, errorWriter io.Writer) (TL, error) {
	lex := newLexer(str, file, opts)
	allTokens, err := lex.generateTokens()
	if err != nil {
		return TL{}, fmt.Errorf("tokenizer error: %w", err)
	}

	recombined := lex.recombineTokens()

	if str != recombined { // We test on all user files forever
		log.Panicf("invariant violation in tokenizer, %s", ContactAuthorsString)
	}

	it := tokenIterator{tokens: allTokens}
	for ; it.count() != 0; it.popFront() {
		tok := it.front()
		if tok.tokenType == comment && strings.HasPrefix(tok.val, "/*") {
			tok.val = tok.val[:2] // do not print the whole comment, but only the first line
			e1 := parseErrToken(fmt.Errorf("multiline comments are not part of language"), tok, tok.pos)
			if !opts.AllowMLC {
				return TL{}, e1
			}
			e1.PrintWarning(errorWriter, nil)
		}
	}

	functionSection := false
	var res TL

	orderIndex := 0
	rest := tokenIterator{tokens: allTokens}
	commentStart := rest
	for !rest.expect(eof) {
		switch rest.front().tokenType {
		case typesSection:
			functionSection = false
			rest.popFront()
			commentStart = rest
			continue
		case functionsSection:
			functionSection = true
			rest.popFront()
			commentStart = rest
			continue
		}
		var td Combinator
		td, rest, err = parseCombinator(commentStart, rest, functionSection, opts.AllowBuiltin)
		if err != nil {
			if functionSection {
				return nil, fmt.Errorf("function declaration error: %w", err)
			}
			return nil, fmt.Errorf("type declaration error: %w", err)
		}
		td.OriginalOrderIndex = orderIndex
		orderIndex++
		res = append(res, &td)
		commentStart = rest
	}
	return res, nil
}
