// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlast

import (
	"strings"
)

// comment is everything from the last empty line up to combinator
func parseCommentBefore(begin tokenIterator, end tokenIterator) string {
	nonWSSinceStartOfLine := false
	it := begin
	for it.offset < end.offset {
		switch tok := it.popFront(); tok.tokenType {
		case comment:
			nonWSSinceStartOfLine = true
			continue
		case newLine:
			if !nonWSSinceStartOfLine { // fully empty line
				begin = it
			}
			nonWSSinceStartOfLine = false
		case whiteSpace, tab:
			continue
		default:
			panic("unexpected token in whitespace")
		}
	}
	commentBefore := strings.TrimSpace(begin.front().pos.fileContent[begin.front().pos.offset:end.front().pos.offset])
	//if commentBefore != "" {
	//	fmt.Printf("<%s>\n", commentBefore)
	//}
	return commentBefore
}

// comment is everything from the last empty line up to combinator
func parseCommentRight(begin tokenIterator, end tokenIterator) string {
	commentRight := strings.TrimSpace(begin.front().pos.fileContent[begin.front().pos.offset:end.front().pos.offset])
	//if commentRight != "" {
	//	fmt.Printf("|%s\n", commentRight)
	//}
	return commentRight
}
