// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlast

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

// Please add methods in the same order types defined in tlparser.go
// All parsing, except TypeRef, which is in separate file tlparser_typeref.go

func skipWS(rest *[]token, outer Position) PositionRange {
	for ; len(*rest) != 0; *rest = (*rest)[1:] {
		switch tok := (*rest)[0]; tok.tokenType {
		case comment, whiteSpace, tab, newLine:
			continue
		default:
			return PositionRange{Outer: outer, Begin: tok.pos, End: tok.pos}
		}
	}
	if len(*rest) == 0 {
		log.Panicf("tokenizer invariant failed, no eof token, %s", ContactAuthorsString)
	}
	return PositionRange{}
}

func checkToken(rest *[]token, i int) bool {
	skipWS(rest, Position{})
	return (*rest)[0].tokenType == i
}

func expect(rest *[]token, i int) bool {
	if checkToken(rest, i) {
		*rest = (*rest)[1:]
		return true
	}
	return false
}

func expectOrPanic(tokens *[]token, i int) {
	if !expect(tokens, i) {
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

func parseLCIdentNS(tokens []token, outer Position) (Name, []token, error) {
	rest := tokens
	var res Name
	if checkToken(&rest, lcIdentNS) {
		res = splitIdenNSFromToken(rest[0].val)
		expectOrPanic(&rest, lcIdentNS)
		return res, rest, nil
	}
	if checkToken(&rest, lcIdent) {
		res.Name = rest[0].val
		expectOrPanic(&rest, lcIdent)
		return res, rest, nil
	}
	return Name{}, tokens, parseErrToken(fmt.Errorf("low-case name (with optional namespace) expected"), rest[0], outer)
}

func parseUCIdentNS(tokens []token, outer Position) (Name, []token, error) {
	rest := tokens
	var res Name
	if checkToken(&rest, ucIdentNS) {
		res = splitIdenNSFromToken(rest[0].val)
		expectOrPanic(&rest, ucIdentNS)
		return res, rest, nil
	}
	if checkToken(&rest, ucIdent) {
		res.Name = rest[0].val
		expectOrPanic(&rest, ucIdent)
		return res, rest, nil
	}
	return Name{}, tokens, parseErrToken(fmt.Errorf("upper-case name (with optional namespace) expected"), rest[0], outer)
}

func parseVarIdent(tokens []token, outer Position) (string, []token, error) {
	rest := tokens
	if checkToken(&rest, lcIdent) {
		res := rest[0].val
		expectOrPanic(&rest, lcIdent)
		return res, rest, nil
	}
	if checkToken(&rest, ucIdent) {
		res := rest[0].val
		expectOrPanic(&rest, ucIdent)
		return res, rest, nil
	}
	return "", tokens, parseErrToken(fmt.Errorf("name without namespace expected"), rest[0], outer)
}

// functionModifier := @any  | @internal  | @kphp | @read | @readwrite | @write
func parseModifiers(tokens []token, outer Position) ([]Modifier, []token) {
	var res []Modifier
	rest := tokens
	for {
		mod := Modifier{PR: skipWS(&rest, outer)}
		if !checkToken(&rest, functionModifier) {
			break
		}
		mod.Name = rest[0].val
		expectOrPanic(&rest, functionModifier)
		mod.PR.End = rest[0].pos
		res = append(res, mod)
	}
	return res, rest
}

// constructor := fullName [CRC32]
func parseConstructor(tokens []token, outer Position, allowBuiltin bool) (Constructor, []token, error) {
	rest := tokens
	res := Constructor{NamePR: skipWS(&rest, outer)}
	var err error
	if allowBuiltin {
		if checkToken(&rest, numberSign) {
			res.Name.Name = rest[0].val
			expectOrPanic(&rest, numberSign)
			res.NamePR.End = rest[0].pos
			res.IDPR = res.NamePR // wish to highlight name, if tag absent
			return res, rest, nil
		}
	}
	if res.Name, rest, err = parseLCIdentNS(rest, outer); err != nil {
		return Constructor{}, tokens, err // parseErrToken(fmt.Errorf("constructor name expected"), rest[0])
	}
	res.NamePR.End = rest[0].pos
	res.IDPR = res.NamePR // wish to highlight name, if tag absent
	if checkToken(&rest, crc32hash) {
		res.IDPR.Begin = rest[0].pos
		i, err := strconv.ParseUint(rest[0].val[1:], 16, 32)
		if err != nil {
			return Constructor{}, tokens, parseErrToken(fmt.Errorf("error converting constructor tag to uint32: %w", err), rest[0], outer)
		}
		x := uint32(i)
		res.ID = &x
		expectOrPanic(&rest, crc32hash)
		res.IDPR.End = rest[0].pos
	}
	return res, rest, nil
}

func parseTemplateArguments(tokens []token, outer Position) ([]TemplateArgument, []token, error) {
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
func parseTemplateArgument(tokens []token, outer Position) (*TemplateArgument, []token, error) {
	rest := tokens
	var err error
	of := &TemplateArgument{PR: skipWS(&rest, outer)}
	if !expect(&rest, lCurlyBracket) {
		return nil, tokens, nil
	}
	of.FieldName, rest, err = parseVarIdent(rest, outer)
	if err != nil {
		return nil, tokens, fmt.Errorf("template argument name expected: %w", err)
	}
	if !expect(&rest, colon) {
		return nil, tokens, parseErrToken(fmt.Errorf("':' after template argument name expected"), rest[0], outer)
	}
	switch {
	case checkToken(&rest, ucIdent) && rest[0].val == "Type":
		of.IsNat = false
		expectOrPanic(&rest, ucIdent)
	case checkToken(&rest, numberSign):
		of.IsNat = true
		expectOrPanic(&rest, numberSign)
	default:
		return nil, tokens, parseErrToken(fmt.Errorf("template argument type can be either 'Type' or '#'"), rest[0], outer)
	}
	if !expect(&rest, rCurlyBracket) {
		return nil, tokens, parseErrToken(fmt.Errorf("'}' after template argument type expected"), rest[0], outer)
	}
	of.PR.End = rest[0].pos
	return of, rest, nil
}

// fullName [word] ...
func parseTypeDeclaration(tokens []token, outer Position) (TypeDeclaration, []token, error) {
	rest := tokens
	res := TypeDeclaration{PR: skipWS(&rest, outer)}
	res.NamePR = res.PR
	var name Name
	var err error
	if name, rest, err = parseUCIdentNS(rest, outer); err != nil {
		return TypeDeclaration{}, tokens, err
	}
	res.Name = name
	res.NamePR.End = rest[0].pos
	for {
		var argName string
		argPR := skipWS(&rest, outer)
		argName, rest, err = parseVarIdent(rest, outer)
		if err != nil {
			break
		}
		argPR.End = rest[0].pos
		res.Arguments = append(res.Arguments, argName)
		res.ArgumentsPR = append(res.ArgumentsPR, argPR)
	}
	res.PR.End = rest[0].pos
	return res, rest, nil
}

// TODO: make this grammar correct
// arithmetic :=  arithmetic '+' arithmetic | number | '(' arithmetic ')'
func parseArithmetic(tokens []token, outer Position, force bool) (*Arithmetic, []token, error) {
	rest := tokens
	var err error
	var res *Arithmetic
	switch {
	case expect(&rest, lRoundBracket):
		res, rest, err = parseArithmetic(rest, outer, force)
		if err != nil {
			return nil, tokens, err
		}
		if res == nil {
			return nil, tokens, nil
		}
		if !expect(&rest, rRoundBracket) {
			return nil, tokens, parseErrToken(fmt.Errorf("')' expected"), rest[0], outer)
		}
	case checkToken(&rest, number):
		var i uint64
		i, err = strconv.ParseUint(rest[0].val, 10, 32)
		if err != nil {
			return nil, tokens, parseErrToken(fmt.Errorf("constant overflows uint32: %w", err), rest[0], outer)
		}
		expectOrPanic(&rest, number)
		res = &Arithmetic{
			Nums: []uint32{uint32(i)},
			Res:  uint32(i),
		}
	default:
		if force {
			return nil, tokens, parseErrToken(fmt.Errorf("arithmetic expression expected after '+'"), rest[0], outer)
		}
		return nil, tokens, nil
	}
	for expect(&rest, plus) {
		var res2 *Arithmetic
		res2, rest, err = parseArithmetic(rest, outer, true)
		if err != nil {
			return nil, tokens, err
		}
		sum := uint64(res.Res) + uint64(res2.Res)
		if sum >= math.MaxUint32 {
			return nil, tokens, parseErrToken(fmt.Errorf("arithmetic expression overflows uint32"), rest[0], outer)
		}
		res.Res = uint32(sum)
		res.Nums = append(res.Nums, res2.Nums...)
	}
	return res, rest, err
}

// aot := [T | arithmetic]
func parseArithmeticOrTypeOpt(tokens []token, applyFlag bool, outer Position) (*ArithmeticOrType, []token, error) {
	rest := tokens
	var a *Arithmetic
	pr := skipWS(&rest, outer)
	var err error
	a, rest, err = parseArithmetic(rest, outer, false)
	if err != nil {
		return nil, tokens, err
	}
	if a != nil {
		t := TypeRef{PR: pr}   // TODO - move into parsing of Arithmetic
		t.PR.End = rest[0].pos // t stores PR in arithmeticOrType
		return &ArithmeticOrType{T: t, IsArith: true, Arith: *a}, rest, nil
	}
	var t *TypeRef
	t, rest, err = parseTypeRef(rest, applyFlag, outer)
	if err != nil {
		return nil, tokens, err
	}
	if t != nil {
		return &ArithmeticOrType{T: *t}, rest, nil
	}
	return nil, tokens, nil
}

// repeatType := [ scale '*' ]
func parseScaleFactorOpt(tokens []token, outer Position) (*ScaleFactor, []token, error) {
	rest := tokens
	res := ScaleFactor{PR: skipWS(&rest, outer)}
	var err error
	res.Scale, rest, err = parseVarIdent(rest, outer)
	if err == nil {
		res.PR.End = rest[0].pos
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
	res.PR.End = rest[0].pos
	return nil, rest, nil // TODO - check
}

// repeatType := [ scale '*' ] '[' field ... ']'
func parseRepeatWithScaleOpt(tokens []token, outer Position) (*RepeatWithScale, []token, error) {
	rest := tokens
	res := RepeatWithScale{PR: skipWS(&rest, outer)}
	var err error
	var scale *ScaleFactor
	scale, rest, err = parseScaleFactorOpt(rest, outer)
	if err != nil {
		return nil, tokens, fmt.Errorf("error parsing multiplier: %w", err)
	}
	if scale != nil {
		res.Scale = *scale
		if !expect(&rest, asterisk) {
			return nil, tokens, nil
		}
		res.ExplicitScale = true
		if !expect(&rest, lSquareBracket) {
			return nil, tokens, parseErrToken(fmt.Errorf("'[' is expected after '*'"), rest[0], outer)
		}
	} else if !expect(&rest, lSquareBracket) {
		return nil, tokens, nil
	}
	// rBracketToken := rest[0]
	res.Rep, rest, err = parseFields(rest, rSquareBracket, outer)
	if err != nil {
		return nil, tokens, err // fmt.Errorf("error parsing fields in square brackets: %w", err)
	}
	// if res.Rep == nil {
	//	return RepeatWithScale{}, nil, parseErrToken(fmt.Errorf("empty square brackets not allowed: "), rBracketToken, outer)
	// }
	res.PR.End = rest[0].pos
	return &res, rest, nil
}

// fieldMask := word '.' number '?'
func parseFieldMask(tokens []token, outer Position) (*FieldMask, []token, error) {
	rest := tokens
	res := &FieldMask{PRName: skipWS(&rest, outer)}
	var name string
	var err error
	if name, rest, err = parseVarIdent(rest, outer); err != nil {
		return nil, tokens, nil
	}
	res.MaskName = name
	res.PRName.End = rest[0].pos
	if !expect(&rest, dotSign) {
		return nil, tokens, nil
	}
	res.PRBits = skipWS(&rest, outer)
	if !checkToken(&rest, number) {
		return nil, tokens, parseErrToken(fmt.Errorf("expecting decimal bitmask bit number"), rest[0], outer)
	}
	i, err := strconv.ParseUint(rest[0].val, 10, 32)
	if err != nil {
		return nil, tokens, parseErrToken(fmt.Errorf("error converting bitmask to uint32: %w", err), rest[0], outer)
	}
	res.BitNumber = uint32(i)
	expectOrPanic(&rest, number)
	res.PRBits.End = rest[0].pos
	if !expect(&rest, questionMark) {
		return nil, tokens, parseErrToken(fmt.Errorf("'?' expected after field bitmask "), rest[0], outer)
	}
	return res, rest, nil
}

// fieldName := name ':'
func parseFieldName(tokens []token, outer Position) (string, []token, Position) {
	rest := tokens
	var name string
	var err error
	if name, rest, err = parseVarIdent(rest, outer); err != nil {
		return "", rest, Position{}
	}
	end := rest[0].pos
	if !expect(&rest, colon) {
		return "", tokens, end
	}
	return name, rest, end
}

// field := [ fieldName ] [ ! ]     [ fieldMask ] T
func parseField(tokens []token, outer Position) (Field, []token, error) {
	rest := tokens
	res := Field{PR: skipWS(&rest, outer)}
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
	if expect(&rest, exclamation) {
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
		res.PR.End = rest[0].pos
		return res, rest, nil
	}
	// typeStartPR := skipWS(&rest, outer) - experimental beautiful error code below
	var t *TypeRef
	t, rest, err = parseTypeRef(rest, false, outer)
	if err != nil {
		// berr := err.(*ParseError)
		// typeStartPR.End = berr.Pos.Begin
		// return Field{}, tokens, BeautifulError2(berr, typeStartPR.BeautifulError(fmt.Errorf("field type here")))
		return Field{}, tokens, fmt.Errorf("error parsing field type: %w", err)
	}
	if t == nil {
		return Field{}, tokens, parseErrToken(fmt.Errorf("field type is expected here (missed '()' around complex type?)"), rest[0], outer)
	}
	res.FieldType = *t
	res.PR.End = rest[0].pos
	return res, rest, nil
}

func parseFields(tokens []token, finishToken int, outer Position) ([]Field, []token, error) {
	var res []Field
	rest := tokens
	var err error
	for !expect(&rest, finishToken) {
		var field Field
		field, rest, err = parseField(rest, outer)
		if err != nil {
			return nil, tokens, err // fmt.Errorf("'%c' or field declaration expected: %w", finishToken, err) // parseErrToken(fmt.Errorf("field declaration expected:"), rest[0])
		}
		res = append(res, field)
	}
	return res, rest, nil
}

// funcDecl := apply | '(' apply ')' | '%' T | fullName '<' aot ',' ... '>
func parseFuncDecl(tokens []token, outer Position) (TypeRef, []token, error) {
	// Separate function for better documentation
	rest := tokens
	t, rest, err := parseTypeRef(rest, true, outer)
	if err != nil {
		return TypeRef{}, tokens, err
	}
	if t == nil {
		return TypeRef{}, tokens, parseErrToken(fmt.Errorf("return type is expected here"), rest[0], outer)
	}
	return *t, rest, nil
}

// type := constructor [templateArgument] ... [field] ...  '=' typeDecl ';'
// or
// function := [ functionModifier ] constructor [templateArgument] ... [field] ... '=' apply;'
func parseCombinator(tokens []token, isFunction bool, allowBuiltin bool) (Combinator, []token, error) {
	rest := tokens

	td := Combinator{PR: skipWS(&rest, Position{})}
	outer := td.PR.Begin
	td.PR.Outer = outer // Set outer context for all parsing
	var err error
	td.Modifiers, rest = parseModifiers(rest, outer) // we support modifiers for normal types also
	// list of modifiers is checked in generator, where they are used for some purposes
	td.Construct, rest, err = parseConstructor(rest, outer, allowBuiltin)
	if err != nil {
		return Combinator{}, tokens, fmt.Errorf("constructor declaration error: %w", err)
	}
	td.TemplateArgumentsPR = skipWS(&rest, outer)
	td.TemplateArguments, rest, err = parseTemplateArguments(rest, outer)
	if err != nil {
		return Combinator{}, tokens, err
	}
	if len(td.TemplateArguments) == 0 {
		td.TemplateArgumentsPR.Begin = td.Construct.IDPR.End // highlight empty space when no arguments
	}
	td.TemplateArgumentsPR.End = rest[0].pos
	if checkToken(&rest, questionMark) {
		if isFunction {
			return Combinator{}, tokens, parseErrToken(fmt.Errorf("'?' (legacy builtin type body) is not allowed in functions"), rest[0], outer)
		}
		expectOrPanic(&rest, questionMark)
		td.Builtin = true
		if !expect(&rest, equalSign) {
			return Combinator{}, tokens, parseErrToken(fmt.Errorf("'=' expected after '?' (legacy builtin type body)"), rest[0], outer)
		}
	} else {
		td.Fields, rest, err = parseFields(rest, equalSign, outer)
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
	}
	if !expect(&rest, semiColon) {
		return Combinator{}, tokens, parseErrToken(fmt.Errorf("';' or type argument expected"), rest[0], outer)
	}
	td.PR.End = rest[0].pos
	return td, rest, nil
}

func ParseTL(str string) (TL, error) {
	return ParseTLFile(str, "")
}

func ParseTLFile(str, file string) (TL, error) {
	return ParseTL2(str, file, false)
}

// ParseTL TL := TypesSection [ type ... ] FunctionSection [ function ... ]
func ParseTL2(str, file string, allowBuiltin bool) (TL, error) {
	lex := newLexer(str, file, allowBuiltin)
	rest, err := lex.generateTokens()
	if err != nil {
		return TL{}, fmt.Errorf("tokenizer error: %w", err)
	}

	recombined := lex.recombineTokens()

	if str != recombined { // We test on all user files forever
		log.Panicf("invariant violation in tokenizer, %s", ContactAuthorsString)
	}

	functionSection := false
	var res TL

	orderIndex := 0
	for !expect(&rest, eof) {
		switch rest[0].tokenType {
		case typesSection:
			functionSection = false
			rest = rest[1:]
			continue
		case functionsSection:
			functionSection = true
			rest = rest[1:]
			continue
		}
		var td Combinator
		td, rest, err = parseCombinator(rest, functionSection, allowBuiltin)
		if err != nil {
			if functionSection {
				return nil, fmt.Errorf("function declaration error: %w", err)
			}
			return nil, fmt.Errorf("type declaration error: %w", err)
		}
		td.OriginalOrderIndex = orderIndex
		orderIndex++
		res = append(res, &td)
	}
	return res, nil
}
