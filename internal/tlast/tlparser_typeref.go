// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlast

import "fmt"

// parsing all kinds of type references

// can't be "apply" (without brackets) as name_1:type_1 name_2:type_2 => (type_1 name_2) will be parsed
// T := '(' apply ')'                 |
//
//	'(' T ')'                     |
//	fullName '<' aot ',' ... '>'  |
//	repeatType                    |
//	fullName                      |
//	bareType                      |
//	'#'                           |
//	'?'                               - we prohibit it, instead it is parsed as type body
//
// if applyFlag == true then it will try to parse a b c d type without brackets (used when there is outer brackets)
func parseTypeRef(tokens []token, applyFlag bool, outer Position) (*TypeRef, []token, error) {
	rest := tokens
	var pt *TypeRef
	var err error

	pr := skipWS(&rest, outer)

	if expect(&rest, numberSign) {
		pr.End = rest[0].pos
		return &TypeRef{Type: Name{Name: "#"}, PR: pr, PRArgs: pr.CollapseToEnd()}, rest, nil
	}

	bare := expect(&rest, percentSign)

	pt, rest, err = parseTypeRefInRoundBracketsOpt(rest, outer)
	if err != nil {
		return nil, tokens, err // fmt.Errorf("t in round brackets error: %w", err)
	}
	if pt != nil {
		pt.Bare = pt.Bare || bare
		pt.PR = pr
		pt.PR.End = rest[0].pos
		return pt, rest, nil
	}

	pt, rest, err = parseTypeRefWithAngleBracketsOpt(rest, outer)
	if err != nil {
		return nil, tokens, fmt.Errorf("angle brackets error: %w", err)
	}
	if pt != nil {
		pt.Bare = pt.Bare || bare
		pt.PR = pr
		pt.PR.End = rest[0].pos
		return pt, rest, nil
	}

	var name Name
	if name, rest, err = parseTypeRefAsName(rest, outer); err != nil {
		return nil, tokens, nil
	}
	res := TypeRef{PR: pr, Type: name}
	res.PR.End = rest[0].pos
	res.PRArgs = res.PR.CollapseToEnd() // skipWS(&rest, outer)
	if applyFlag {
		for {
			res.PRArgs.End = rest[0].pos
			var aot *ArithmeticOrType
			aot, rest, err = parseArithmeticOrTypeOpt(rest, false, outer)
			if err != nil {
				return nil, tokens, err
			}
			if aot == nil {
				break
			}
			res.Args = append(res.Args, *aot)
		}
	}
	res.Bare = res.Bare || bare
	return &res, rest, nil
}

// '(' T ')
// returns nil, nil if ( not found
// returns &t , nil if parsed
func parseTypeRefInRoundBracketsOpt(tokens []token, outer Position) (*TypeRef, []token, error) {
	rest := tokens
	if !expect(&rest, lRoundBracket) {
		return nil, tokens, nil
	}
	var res *TypeRef
	var err error
	res, rest, err = parseTypeRef(rest, true, outer)
	if err != nil {
		return nil, tokens, err // fmt.Errorf("bad t: %w", err)
	}
	if res == nil {
		return nil, tokens, parseErrToken(fmt.Errorf("')' or type is expected here"), rest[0], outer)
	}
	if !expect(&rest, rRoundBracket) {
		return nil, tokens, parseErrToken(fmt.Errorf("')' or type is expected here"), rest[0], outer)
	}
	return res, rest, nil
}

// fullName '<' aot ',' ... '>
// returns nil, nil if fullName or < not found after fullName
// returns &t , nil if parsed
func parseTypeRefWithAngleBracketsOpt(tokens []token, outer Position) (*TypeRef, []token, error) {
	var res TypeRef
	rest := tokens
	var err error
	if res.Type, rest, err = parseTypeRefAsName(rest, outer); err != nil {
		return nil, tokens, nil
	}
	if !expect(&rest, lAngleBracket) {
		return nil, tokens, nil
	}
	res.PRArgs = skipWS(&rest, outer)
	for {
		var aot *ArithmeticOrType
		aot, rest, err = parseArithmeticOrTypeOpt(rest, true, outer)
		if err != nil {
			return nil, tokens, err
		}
		if aot == nil {
			return nil, tokens, parseErrToken(fmt.Errorf("',', '>' or type expected here"), rest[0], outer)
		}
		res.Args = append(res.Args, *aot)
		if expect(&rest, commaSign) {
			continue
		}
		res.PRArgs.End = rest[0].pos
		if !expect(&rest, rAngleBracket) {
			return nil, tokens, parseErrToken(fmt.Errorf("'>' or type expected here"), rest[0], outer)
		}
		break
	}
	return &res, rest, nil
}

func parseTypeRefAsName(tokens []token, outer Position) (Name, []token, error) {
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
	return Name{}, tokens, parseErrToken(fmt.Errorf("name (with optional namespace) expected"), rest[0], outer)
}
