// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package pure

import (
	"fmt"
	"strings"

	"github.com/VKCOM/tl/internal/tlast"
)

type cyclePathElement struct {
	ins    TypeInstance
	prName tlast.PositionRange
}

type cycleFinder struct {
	visitedNodes map[TypeInstance]int // int: we find the shortest cycle to help with debugging
	path         []cyclePathElement
	found        []cyclePathElement
}

func (c *cycleFinder) printCycle() error {
	if len(c.found) < 2 {
		return nil
	}
	sb := strings.Builder{}
	for i, pe := range c.found {
		if i != 0 {
			sb.WriteString("->")
		}
		sb.WriteString(pe.ins.CanonicalName())
	}
	// first element has no PR because we start from type definition, not field
	return c.found[1].prName.BeautifulError(fmt.Errorf("found infinite cycle %s, use optional to break it", sb.String()))
}

func (c *cycleFinder) reset() {
	if c.visitedNodes == nil {
		c.visitedNodes = map[TypeInstance]int{}
	}
	clear(c.visitedNodes)
	c.path = c.path[:0]
	c.found = c.found[:0]
}

func (c *cycleFinder) push(ins TypeInstance, prName tlast.PositionRange) bool {
	depth := len(c.path)
	if depth == 0 { // source type
		c.path = append(c.path, cyclePathElement{ins: ins, prName: prName})
		return true
	}
	if wasDepth, ok := c.visitedNodes[ins]; ok && depth >= wasDepth {
		return false
	}
	c.visitedNodes[ins] = depth
	c.path = append(c.path, cyclePathElement{ins: ins, prName: prName})
	if len(c.path) > 1 && c.path[0].ins == ins && (len(c.found) == 0 || depth < len(c.found)) {
		c.found = append(c.found[:0], c.path...)
	}
	return true
}

func (c *cycleFinder) pop(ins TypeInstance) {
	if c.path[len(c.path)-1].ins != ins {
		panic("cycle finder wrong recursion")
	}
	c.path = c.path[:len(c.path)-1]
}
