package tlast

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type OptionalParse[T any] func(tokens tokenIterator, position Position) (state OptionalState, restTokens tokenIterator, result T)

type OptionalState struct {
	StartProcessing bool
	Error           error
}

func (state *OptionalState) Inherit(otherState OptionalState) {
	state.StartProcessing = state.StartProcessing || otherState.StartProcessing
	if state.Error == nil {
		state.Error = otherState.Error
	}
}

func (state *OptionalState) ExpectProgress(orError string) bool {
	if state.HasProgress() {
		return true
	}
	if state.IsOmitted() {
		state.Fail(orError)
	}
	return false
}

func (state *OptionalState) Fail(error string) {
	state.StartProcessing = true
	if state.Error == nil {
		state.Error = fmt.Errorf("%s", error)
	}
}

func (state *OptionalState) FailWithError(error error) {
	state.StartProcessing = true
	if state.Error == nil {
		state.Error = error
	}
}

func (state *OptionalState) HasProgress() bool {
	return state.StartProcessing && state.Error == nil
}

func (state *OptionalState) IsFailed() bool {
	return state.StartProcessing && state.Error != nil
}

func (state *OptionalState) IsOmitted() bool {
	return !state.StartProcessing
}

func zeroOrMore[T any](parser OptionalParse[T]) OptionalParse[[]T] {
	return func(tokens tokenIterator, position Position) (state OptionalState, restTokens tokenIterator, result []T) {
		restTokens = tokens
		for {
			var localState OptionalState
			var element T
			localState, restTokens, element = parser(restTokens, position)
			state.StartProcessing = state.StartProcessing || localState.StartProcessing
			if !localState.HasProgress() {
				state.Error = localState.Error
				break
			}
			result = append(result, element)
		}
		return
	}
}

func ParseTL2(str string) (TL2File, error) {
	return ParseTL2File(str, "", LexerOptions{LexerLanguage: tl2}, os.Stdout)
}

// TL2File := TL2Combinator* EOF;
func ParseTL2File(str, file string, opts LexerOptions, errorWriter io.Writer) (tl2File TL2File, err error) {
	if opts.LexerLanguage != tl2 {
		return TL2File{}, fmt.Errorf("ParseTL2File can't parse nothing rather than tl2")
	}
	lex := newLexer(str, file, opts)
	allTokens, err := lex.generateTokens()
	if err != nil {
		return TL2File{}, fmt.Errorf("tokenizer error: %w", err)
	}

	recombined := lex.recombineTokens()

	if str != recombined { // We test on all user files forever
		log.Panicf("invariant violation in tokenizer, %s", ContactAuthorsString)
	}

	it := tokenIterator{tokens: allTokens}
	for !it.expect(eof) {
		var combinator TL2Combinator
		combinator, it, err = parseTL2Combinator(it)
		if err != nil {
			return TL2File{}, err
		}
		tl2File.Combinators = append(tl2File.Combinators, combinator)
	}

	return
}

// TL2Combinator := TL2Annotation* (TL2TypeDeclaration | TL2FuncDeclaration) scl;
func parseTL2Combinator(it tokenIterator) (TL2Combinator, tokenIterator, error) {
	var state OptionalState

	rest := it
	combinator := TL2Combinator{PR: rest.skipWS(Position{})}
	outer := combinator.PR.Begin

	combinator.CommentBefore = parseCommentBefore(it, rest)

	state, rest, combinator.Annotations = zeroOrMore(parseTL2Annotation)(rest, outer)
	if state.Error != nil {
		return TL2Combinator{}, tokenIterator{}, state.Error
	}

	var name TL2TypeName
	prName := rest.skipWS(outer)
	state, rest, name = parseTL2TypeName(rest, outer)
	prName.End = rest.front().pos
	if !state.ExpectProgress("expected type name") {
		return TL2Combinator{}, tokenIterator{}, state.Error
	}

	var typeDeclState OptionalState
	typeDeclState, rest, combinator.TypeDecl = parseTL2TypeDeclarationWithoutName(rest, outer, name)
	state.Inherit(typeDeclState)
	if !typeDeclState.StartProcessing {
		var funcDeclState OptionalState
		funcDeclState, rest, combinator.FuncDecl = parseTL2FuncDeclarationWithoutName(rest, outer, name)
		state.Inherit(funcDeclState)
		if funcDeclState.StartProcessing {
			combinator.IsFunction = true
			combinator.FuncDecl.PRName = prName
		} else {
			state.Fail("expected either function or type declaration")
		}
	} else {
		combinator.TypeDecl.PRName = prName
	}

	if !rest.expect(semiColon) {
		state.Fail("expected semicolon in the end of combinator definition")
	}

	combinator.PR.End = rest.front().pos
	return combinator, rest, state.Error
}

// TL2Annotation := at lcName;
func parseTL2Annotation(tokens tokenIterator, position Position) (state OptionalState, restTokens tokenIterator, result TL2Annotation) {
	restTokens = tokens
	state.StartProcessing = true
	result.PR = restTokens.skipWS(position)
	if !restTokens.checkToken(annotation) {
		state.StartProcessing = false
	} else {
		curToken := restTokens.popFront()
		result.Name = curToken.val[1:]
		result.PR.End = restTokens.front().pos
	}
	result.PR.End = restTokens.front().pos
	return
}

// TL2TypeName := (lcName dot)? lcName;
func parseTL2TypeName(tokens tokenIterator, position Position) (state OptionalState, restTokens tokenIterator, result TL2TypeName) {
	restTokens = tokens
	state.StartProcessing = true
	switch {
	case restTokens.checkToken(lcIdent):
		result.Name = restTokens.popFront().val
	case restTokens.checkToken(lcIdentNS):
		value := restTokens.popFront().val
		dotIndex := strings.Index(value, ".")

		result.Namespace = value[:dotIndex]
		result.Name = value[dotIndex+1:]
	case restTokens.checkToken(ucIdentNS):
		state.Fail("tl2 type names can't start from uppercase")
	default:
		state.StartProcessing = false
	}
	return
}

// TL2FuncDeclaration := TL2TypeName CRC32 TL2Field* funEq TL2TypeDefinition?;
func parseTL2FuncDeclarationWithoutName(tokens tokenIterator, position Position, name TL2TypeName) (state OptionalState, restTokens tokenIterator, result TL2FuncDeclaration) {
	restTokens = tokens
	result.PR = restTokens.skipWS(position)
	result.Name = name

	if restTokens.checkToken(crc32hash) {
		result.PRID = restTokens.skipWS(position)
		crcToken := restTokens.popFront()
		value, err := strconv.ParseUint(crcToken.val[1:], 16, 32)
		if err != nil {
			state.FailWithError(err)
			return
		}
		result.ID = new(uint32)
		*result.ID = uint32(value)
		result.PRID.End = restTokens.front().pos
	} else {
		state.StartProcessing = false
		return
	}

	state, restTokens, result.Arguments = zeroOrMore(parseTL2Field)(restTokens, position)
	if state.IsFailed() {
		return
	}

	if restTokens.expect(functionSign) {
		state.StartProcessing = true

		var localState OptionalState
		localState, restTokens, result.ReturnType = parseTL2TypeDefinition(restTokens, position)
		state.Inherit(localState)
		if localState.IsOmitted() {
			result.ReturnType.IsUnionType = false
			result.ReturnType.IsConstructorFields = true
			return
		}
	}

	result.PR.End = restTokens.front().pos
	return
}

// TL2TypeDeclaration := TL2TypeName (lts TL2TypeArgumentDeclaration (cm TL2TypeArgumentDeclaration)* gts)? CRC32? eq TL2TypeDefinition?;
func parseTL2TypeDeclarationWithoutName(tokens tokenIterator, position Position, name TL2TypeName) (state OptionalState, restTokens tokenIterator, result TL2TypeDeclaration) {
	restTokens = tokens
	result.PR = restTokens.skipWS(position)
	result.Name = name
	switch {
	case restTokens.expect(lAngleBracket):
		state.StartProcessing = true
		result.TemplateArguments = make([]TL2TypeTemplate, 1)
		var localState OptionalState
		localState, restTokens, result.TemplateArguments[0] = parseTL2TypeArgumentDeclaration(restTokens, position)
		state.Inherit(localState)
		if !state.ExpectProgress("expected type argument declaration") {
			return
		}
		for {
			if restTokens.expect(commaSign) {
				var template TL2TypeTemplate
				var localState OptionalState
				localState, restTokens, template = parseTL2TypeArgumentDeclaration(restTokens, position)
				state.Inherit(localState)
				if !state.ExpectProgress("expected type argument declaration") {
					return
				}
				result.TemplateArguments = append(result.TemplateArguments, template)
			} else {
				break
			}
		}
		if !restTokens.expect(rAngleBracket) {
			state.Fail("can't stop parse template arguments without closing brackets")
			return
		}
	}

	if restTokens.checkToken(crc32hash) {
		result.PRID = restTokens.skipWS(position)
		crcToken := restTokens.popFront()
		value, err := strconv.ParseUint(crcToken.val[1:], 16, 32)
		if err != nil {
			state.FailWithError(err)
			return
		}
		result.ID = new(uint32)
		*result.ID = uint32(value)
		result.PRID.End = restTokens.front().pos
	}

	if restTokens.expect(equalSign) {
		state.StartProcessing = true
	} else {
		restTokens = tokens
		return
	}

	var localState OptionalState
	localState, restTokens, result.Type = parseTL2TypeDefinition(restTokens, position)
	state.Inherit(localState)
	if localState.IsOmitted() {
		result.Type.IsUnionType = false
		result.Type.IsConstructorFields = true
		return
	}
	result.PR.End = restTokens.front().pos
	return
}

// TL2TypeDefinition = TL2TypeRef | TL2Field* | TL2UnionType;
func parseTL2TypeDefinition(tokens tokenIterator, position Position) (state OptionalState, restTokens tokenIterator, result TL2TypeDefinition) {
	restTokens = tokens
	result.PR = restTokens.currentPositionRange(position)

	state, restTokens, result.ConstructorFields = zeroOrMore(parseTL2Field)(restTokens, position)
	if state.StartProcessing {
		result.IsConstructorFields = true
	} else {
		state, restTokens, result.UnionType = parseTL2UnionType(restTokens, position)
		if state.StartProcessing {
			result.IsUnionType = true
		} else {
			state, restTokens, result.TypeAlias = parseTL2Type(restTokens, position)
		}
	}

	result.PR.End = restTokens.front().pos
	return
}

// TL2UnionType := vb? TL2UnionTypeVariant (vb TL2UnionTypeVariant)+;
func parseTL2UnionType(tokens tokenIterator, position Position) (state OptionalState, restTokens tokenIterator, result TL2UnionType) {
	defer func() {
		if !state.StartProcessing {
			restTokens = tokens
		}
	}()

	restTokens = tokens
	result.PR = restTokens.skipWS(position)

	if restTokens.expect(verticalBar) {
		state.StartProcessing = true
	}

	result.Variants = make([]TL2UnionTypeVariant, 1)
	var localState OptionalState
	localState, restTokens, result.Variants[0] = parseTL2UnionTypeVariant(restTokens, position)
	if localState.IsFailed() {
		state.Fail("can't parse first variant")
		return
	} else if state.StartProcessing && localState.IsOmitted() {
		state.Fail("expected first variant of union")
		return
	}

	state.Inherit(localState)

	if !state.StartProcessing && !result.Variants[0].IsConstructor {
		// maybe typedef
		state.StartProcessing = false
	}

	for {
		if !restTokens.expect(verticalBar) {
			break
		}
		state.StartProcessing = true
		var localState OptionalState
		var variant TL2UnionTypeVariant
		localState, restTokens, variant = parseTL2UnionTypeVariant(restTokens, position)
		result.Variants = append(result.Variants, variant)
		state.Inherit(localState)
		if !localState.ExpectProgress("expected union variant definition after vertical var") {
			state.FailWithError(localState.Error)
			return
		}
	}

	if state.IsFailed() {
		return
	}

	if len(result.Variants) < 2 {
		if len(result.Variants) == 1 && !result.Variants[0].IsConstructor {
			state.StartProcessing = false
			return
		}
		state.Fail("expected at least 2 variants on union type")
	}

	result.PR.End = restTokens.front().pos
	return
}

// TL2UnionTypeVariant := TL2TypeRef | TL2UnionConstructor;
func parseTL2UnionTypeVariant(tokens tokenIterator, position Position) (state OptionalState, restTokens tokenIterator, result TL2UnionTypeVariant) {
	restTokens = tokens
	result.PR = restTokens.skipWS(position)

	state, restTokens, result.TypeAlias = parseTL2Type(restTokens, position)
	if !state.StartProcessing {
		state, restTokens, result.Constructor = parseTL2UnionConstructor(restTokens, position)
		result.IsConstructor = state.StartProcessing
	}

	result.PR.End = restTokens.front().pos
	return
}

// TL2UnionConstructor := ucName TL2Field*;
func parseTL2UnionConstructor(tokens tokenIterator, position Position) (state OptionalState, restTokens tokenIterator, result TL2UnionConstructor) {
	restTokens = tokens
	result.PR = restTokens.skipWS(position)

	if restTokens.checkToken(ucIdent) {
		state.StartProcessing = true

		result.PRName = restTokens.skipWS(position)
		result.Name = restTokens.popFront().val
		result.PRName.End = restTokens.front().pos

		var fieldsState OptionalState
		fieldsState, restTokens, result.Fields = zeroOrMore(parseTL2Field)(restTokens, position)
		state.Inherit(fieldsState)
	}

	result.PR.End = restTokens.front().pos
	return
}

// TL2Field := ((lcName qm?) | ucs) cl TL2TypeRef;
func parseTL2Field(tokens tokenIterator, position Position) (state OptionalState, restTokens tokenIterator, result TL2Field) {
	defer func() {
		if !state.StartProcessing {
			restTokens = tokens
		}
	}()

	restTokens = tokens
	result.PR = restTokens.skipWS(position)

	result.CommentBefore = parseCommentBefore(tokens, restTokens)

	switch {
	case restTokens.checkToken(lcIdent) || restTokens.checkToken(underscore):
		result.PRName = restTokens.skipWS(position)

		nameToken := restTokens.popFront()
		result.Name = nameToken.val
		if nameToken.tokenType == underscore {
			result.IsIgnored = true
		}
		result.PRName.End = restTokens.front().pos

		if restTokens.expect(questionMark) {
			if result.IsIgnored {
				state.Fail("ignored field can't be optional")
				break
			} else {
				result.IsOptional = true
				state.StartProcessing = true
			}
		}
		if !restTokens.expect(colon) {
			if state.StartProcessing {
				state.Fail("can't parse field since there is no colon after field name declaration")
			} else {
				state.StartProcessing = false
			}
			break
		}

		state.StartProcessing = true

		var localState OptionalState
		localState, restTokens, result.Type = parseTL2Type(restTokens, position)
		state.Inherit(localState)
		if !state.ExpectProgress("expected type of field") {
			return
		}
	}
	result.PR.End = restTokens.front().pos
	return
}

// TL2TypeRef := TL2TypeApplication | TL2BracketType;
func parseTL2Type(tokens tokenIterator, position Position) (state OptionalState, restTokens tokenIterator, result TL2TypeRef) {
	restTokens = tokens
	result.PR = restTokens.skipWS(position)

	var someType TL2TypeApplication
	state, restTokens, someType = parseTL2TypeApplication(restTokens, position)
	if !state.StartProcessing {
		var bracketType TL2BracketType
		state, restTokens, bracketType = parseTL2BracketType(restTokens, position)
		if state.StartProcessing {
			result.BracketType = &bracketType
			result.IsBracket = true
		}
	} else {
		result.SomeType = &someType
	}
	result.PR.End = restTokens.front().pos
	return
}

// TL2TypeApplication := TL2TypeName (lts TL2TypeArgument (cm TL2TypeArgument)* gts)?;
func parseTL2TypeApplication(tokens tokenIterator, position Position) (state OptionalState, restTokens tokenIterator, result TL2TypeApplication) {
	restTokens = tokens
	result.PR = restTokens.skipWS(position)

	result.PRName = restTokens.skipWS(position)
	state, restTokens, result.Name = parseTL2TypeName(tokens, position)
	if !state.HasProgress() {
		return
	}
	result.PRName.End = restTokens.front().pos

	if restTokens.expectLazy(lAngleBracket) {
		result.Arguments = make([]TL2TypeArgument, 1)

		var argState OptionalState
		argState, restTokens, result.Arguments[0] = parseTL2TypeArgument(restTokens, position)
		state.Inherit(argState)
		if !state.ExpectProgress("expected type argument") {
			return
		}

		for {
			if restTokens.expect(commaSign) {
				var arg TL2TypeArgument
				argState, restTokens, result.Arguments[len(result.Arguments)-1] = parseTL2TypeArgument(restTokens, position)
				result.Arguments = append(result.Arguments, arg)
				state.Inherit(argState)
				if !state.ExpectProgress("expected type argument") {
					return
				}
			} else {
				break
			}
		}

		if !restTokens.expect(rAngleBracket) {
			state.Fail("can't parse type arguments without closing bracket in the end")
			return
		}
	}

	result.PR.End = restTokens.front().pos
	return
}

// TL2BracketType := lsb TL2TypeArgument? rsb TL2TypeRef;
func parseTL2BracketType(tokens tokenIterator, position Position) (state OptionalState, restTokens tokenIterator, result TL2BracketType) {
	restTokens = tokens
	result.PR = restTokens.skipWS(position)

	if restTokens.expectLazy(lSquareBracket) {
		state.StartProcessing = true

		var indexState OptionalState
		var indexType TL2TypeArgument
		indexState, restTokens, indexType = parseTL2TypeArgument(restTokens, position)
		if indexState.IsOmitted() {
			result.IndexType = nil
		} else {
			result.IndexType = &indexType
		}

		state.Inherit(indexState)
		if !state.HasProgress() {
			return
		}

		if !restTokens.expect(rSquareBracket) {
			state.Fail("expected closing square bracket")
			return
		}

		var argState OptionalState
		argState, restTokens, result.ArrayType = parseTL2Type(restTokens, position)
		state.Inherit(argState)
		if !state.ExpectProgress("expected array type argument") {
			return
		}
	}

	result.PR.End = restTokens.front().pos
	return
}

// TL2TypeArgument := TL2TypeRef | number;
func parseTL2TypeArgument(tokens tokenIterator, position Position) (state OptionalState, restTokens tokenIterator, result TL2TypeArgument) {
	restTokens = tokens
	result.PR = restTokens.skipWS(position)

	if restTokens.checkToken(number) {
		state.StartProcessing = true
		result.IsNumber = true

		value, parseErr := strconv.ParseUint(restTokens.popFront().val, 10, 32)
		state.Error = parseErr
		result.Number = uint32(value)
	} else {
		state, restTokens, result.Type = parseTL2Type(restTokens, position)
	}

	result.PR.End = restTokens.front().pos
	return
}

// TL2TypeArgumentDeclaration := lcName cl lcName;
func parseTL2TypeArgumentDeclaration(tokens tokenIterator, position Position) (state OptionalState, restTokens tokenIterator, result TL2TypeTemplate) {
	restTokens = tokens
	state.StartProcessing = true
	result.PR = restTokens.skipWS(position)
	switch {
	case restTokens.checkToken(lcIdent):
		result.Name = restTokens.front().val
		result.PRName = restTokens.skipWS(position)

		restTokens.popFront()
		result.PRName.End = restTokens.front().pos

		if !restTokens.expect(colon) {
			state.Fail("unexpected token during type argument declaration")
			break
		}

		if !restTokens.checkToken(lcIdent) {
			state.Fail("can't parse category of type argument (it can't either \"uint32\" or \"type\")")
			break
		}

		result.Category = TL2TypeCategory(restTokens.front().val)
		result.PRCategory = restTokens.skipWS(position)
		result.PRCategory.End = restTokens.front().pos
		restTokens.popFront()
		if !result.Category.IsLegalCategory() {
			state.FailWithError(fmt.Errorf("unknown category of template: %s (allowed options: \"type\", \"uint32\")", result.Category))
		}
	default:
		state.StartProcessing = false
	}
	result.PR.End = restTokens.front().pos
	return
}
