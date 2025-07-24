// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlast

import (
	"fmt"
	"io"
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
		fmt.Printf("beautiful error created without context, this is a %s, please report with TL file\n", color.InRed("bug"))
	}
}

func (pr PositionRange) BeautifulError(err error) *ParseError {
	pr.CheckValidity()
	return &ParseError{
		Err: err,
		Pos: pr,
	}
}

func BeautifulError2(original *ParseError, compare *ParseError) *ParseError {
	original.compare = compare
	return original
}

type ParseError struct {
	Err error

	Pos PositionRange

	compare *ParseError
}

func (e ParseError) Error() string {
	return e.Err.Error()
}

func (e ParseError) Unwrap() error {
	return e.Err
}

func (e *ParseError) PrintWarning(out io.Writer, outmostError error) {
	e.ConsolePrint(out, outmostError, true)
	_, _ = fmt.Fprintf(out, "\n")
}

func (e *ParseError) ConsolePrint(out io.Writer, outmostError error, isWarning bool) {
	if e.compare != nil {
		e.compare.consolePrint(out, e.compare.Err, color.Purple, false)
	}
	c := color.Red
	if isWarning {
		c = color.Yellow
	}
	if outmostError == nil {
		outmostError = e.Err
	}
	e.consolePrint(out, outmostError, c, isWarning)
}

func safeRange(s string, b int, e int, anyCorrupted *bool) string {
	if b < 0 || b > len(s) || e < b || e > len(s) {
		*anyCorrupted = true
		return ""
	}
	return s[b:e]
}

func (e *ParseError) consolePrint(out io.Writer, err error, c string, isWarning bool) {
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
	arrowText := strings.Repeat("^", len(ourLineRed))
	if len(arrowText) == 0 {
		arrowText = "^"
	}
	arrowText += "--" // arrow handle
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
		_, _ = fmt.Fprintf(out, "%s%s%s", beforeBegin, c, beforeEndLine)
	} else {
		_, _ = fmt.Fprintf(out, "%s", beforeBegin)
	}
	warnText := ""
	if isWarning {
		warnText = color.InYellow("warning: ")
	}
	if anyCorrupted {
		_, _ = fmt.Fprintf(out, "%v\n", err.Error())
		_, _ = fmt.Fprintf(out, "beautiful error context corrupted, %s, please report with TL file\n", color.InRed("internal error"))
	} else {
		_, _ = fmt.Fprintf(out, "%s%s%s\n", ourLineBeforeBegin, color.Colorize(c, ourLineRed), after1)
		// color.InWhite() keeps lines same length in bytes, so arrow points to correct text even if printed into file
		_, _ = fmt.Fprintf(out, "%s%s %s%s\n", errLineBeforeBegin, color.InWhite(arrowText), warnText, after2)
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
