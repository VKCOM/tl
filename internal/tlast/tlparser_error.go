// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlast

import (
	"fmt"
	"strings"

	"github.com/TwiN/go-color"
)

const tabSpaces = "    " // otherwise our error messages wil be shifted

func (pr PositionRange) CollapseToEnd() PositionRange {
	return PositionRange{Outer: pr.Outer, Begin: pr.End, End: pr.End}
}

func (pr PositionRange) CollapseToBegin() PositionRange {
	return PositionRange{Outer: pr.Outer, Begin: pr.Begin, End: pr.Begin}
}

func (pr PositionRange) CheckValidity() {
	if pr.Outer.column == 0 || pr.Begin.column == 0 || pr.End.column == 0 {
		fmt.Printf("beautiful error created without context, this is a %sbug%s, please report with TL file\n", color.Red, color.Reset)
	}
}

func (pr PositionRange) BeautifulError(err error) *ParseError {
	pr.CheckValidity()
	return &ParseError{
		Err: err,
		Pos: pr,
	}
}

func (pr PositionRange) BeautifulWarning(err error) *ParseError {
	pr.CheckValidity()
	return &ParseError{
		Err:       err,
		Pos:       pr,
		isWarning: true,
	}
}

func (pr PositionRange) LogicError(err error) *ParseError {
	pr.CheckValidity()
	return &ParseError{
		Err:        err,
		Pos:        pr,
		LogicError: true,
	}
}

func BeautifulError2(original *ParseError, compare *ParseError) *ParseError {
	original.compare = compare
	return original
}

type ParseError struct {
	Err error

	Pos       PositionRange
	isWarning bool

	LogicError bool // write file dump and ask to send to developers

	compare *ParseError
}

func (e ParseError) Error() string {
	return e.Err.Error()
}

func (e ParseError) Unwrap() error {
	return e.Err
}

func (e *ParseError) ConsolePrint(outmostError error) {
	if e.compare != nil {
		e.compare.consolePrint(e.compare.Err, color.Purple)
	}
	c := color.Red
	if e.isWarning {
		c = color.Yellow
	}
	if outmostError == nil {
		outmostError = e.Err
	}
	e.consolePrint(outmostError, c)
}

func safeRange(s string, b int, e int, anyCorrupted *bool) string {
	if b < 0 || b > len(s) || e < b || e > len(s) {
		*anyCorrupted = true
		return ""
	}
	return s[b:e]
}

func (e *ParseError) consolePrint(err error, c string) {
	// we check all ranges so that we do not crash in case of convoluted TL token replacements
	fc := e.Pos.Outer.fileContent
	anyCorrupted := e.Pos.Begin.fileContent != fc || e.Pos.End.fileContent != fc // combinator must be in one file

	beforeBegin := safeRange(fc, e.Pos.Outer.startLineOffset, e.Pos.Begin.startLineOffset, &anyCorrupted)
	beforeEndLine := safeRange(fc, e.Pos.Begin.startLineOffset, e.Pos.End.startLineOffset, &anyCorrupted)

	ourLineBeforeBegin := ""
	ourLineRed := safeRange(fc, e.Pos.End.startLineOffset, e.Pos.End.offset, &anyCorrupted)
	ourLineRed = strings.ReplaceAll(ourLineRed, "\t", tabSpaces)
	if e.Pos.Begin.startLineOffset == e.Pos.End.startLineOffset {
		ourLineBeforeBegin = safeRange(fc, e.Pos.Begin.startLineOffset, e.Pos.Begin.offset, &anyCorrupted)
		ourLineBeforeBegin = strings.ReplaceAll(ourLineBeforeBegin, "\t", tabSpaces)
		ourLineRed = safeRange(fc, e.Pos.Begin.offset, e.Pos.End.offset, &anyCorrupted)
		ourLineRed = strings.ReplaceAll(ourLineRed, "\t", tabSpaces)
	}
	errLineBeforeBegin := strings.Repeat(" ", len(ourLineBeforeBegin))
	errLineRed := strings.Repeat("^", len(ourLineRed))
	if len(errLineRed) == 0 {
		errLineRed = "^"
	}
	tail := ""
	if e.Pos.End.offset < 0 || e.Pos.End.offset > len(fc) {
		anyCorrupted = true
	} else {
		tail = fc[e.Pos.End.offset:]
	}
	after1 := tail
	after2 := fmt.Sprintf("%s %s (line %d col %d)", err.Error(), e.Pos.Begin.file, e.Pos.Begin.line, e.Pos.Begin.column)
	i := strings.IndexAny(tail, "\r\n")
	if i != -1 {
		after1 = tail[:i]
	}
	after1 = strings.ReplaceAll(after1, "\t", tabSpaces)
	if beforeEndLine != "" {
		fmt.Printf("%s%s%s", beforeBegin, c, beforeEndLine)
	} else {
		fmt.Printf("%s", beforeBegin)
	}
	warnText := ""
	if e.isWarning {
		warnText = color.Yellow + "warning: " + color.Reset
	}
	if anyCorrupted {
		fmt.Printf("%v\n", err.Error())
		fmt.Printf("beautiful error context corrupted, %sinternal error%s, please report with TL file\n", color.Red, color.Reset)
	} else {
		fmt.Printf("%s%s%s%s%s\n", ourLineBeforeBegin, c, ourLineRed, color.Reset, after1)
		fmt.Printf("%s%s%s%s-- %s%s\n", errLineBeforeBegin, color.Reset, errLineRed, color.Reset, warnText, after2) // keep lines same length in bytes
	}
}

func parseErrToken(err error, tok token, outer Position) *ParseError {
	end := tok.pos
	end.offset += len(tok.val)
	end.column += len(tok.val)
	pr := PositionRange{
		Outer: outer,
		Begin: tok.pos,
		End:   end,
	}
	return &ParseError{
		Err: err,
		Pos: pr,
	}
}
