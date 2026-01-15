// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package pure

import "strings"

type cycleFinder struct {
	visitedNodes map[TypeInstance]int
	path         []TypeInstance
	found        []TypeInstance
}

func (c *cycleFinder) printCycle() string {
	if len(c.found) == 0 {
		return ""
	}
	sb := strings.Builder{}
	for i, ins := range c.found {
		if i != 0 {
			sb.WriteString("->")
		}
		sb.WriteString(ins.CanonicalName())
	}
	return sb.String()
}

func (c *cycleFinder) reset() {
	if c.visitedNodes == nil {
		c.visitedNodes = map[TypeInstance]int{}
	}
	clear(c.visitedNodes)
	c.path = c.path[:0]
	c.found = c.found[:0]
}

func (c *cycleFinder) push(ins TypeInstance) bool {
	depth := len(c.path)
	if depth == 0 { // source type
		c.path = append(c.path, ins)
		return true
	}
	if wasDepth, ok := c.visitedNodes[ins]; ok && depth >= wasDepth {
		return false
	}
	c.visitedNodes[ins] = depth
	c.path = append(c.path, ins)
	if len(c.path) > 1 && c.path[0] == ins && (len(c.found) == 0 || depth < len(c.found)) {
		c.found = append(c.found[:0], c.path...)
	}
	return true
}

func (c *cycleFinder) pop(ins TypeInstance) {
	if c.path[len(c.path)-1] != ins {
		panic("cycle finder wrong recursion")
	}
	c.path = c.path[:len(c.path)-1]
}
