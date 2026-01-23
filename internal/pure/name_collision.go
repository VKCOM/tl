// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package pure

import (
	"fmt"
	"strings"

	"github.com/vkcom/tl/internal/tlast"
)

type NameCollision struct {
	namesNormalized map[string]*tlast.ParseError
	names           map[string]*tlast.ParseError
}

func (c *NameCollision) NormalizeName(s string) string {
	s = strings.ReplaceAll(s, "_", "")
	return strings.ToLower(s)
}

func (c *NameCollision) AddUniqueName(name string, pr tlast.PositionRange, reason string) error {
	return c.addName(name, pr, reason, false)
}

func (c *NameCollision) AddSameCaseName(name string, pr tlast.PositionRange, reason string) error {
	return c.addName(name, pr, reason, true)
}

// allows normalized checks between template arguments and fields, but strict checks between them
func (c *NameCollision) ResetNormalized() {
	clear(c.namesNormalized)
}

func (c *NameCollision) addName(name string, pr tlast.PositionRange, reason string, sameCase bool) error {
	nn := c.NormalizeName(name)
	errName, okName := c.names[name]
	errNorm, okNorm := c.namesNormalized[nn]
	if sameCase { // require unique case
		if okNorm && !okName {
			e1 := pr.BeautifulError(fmt.Errorf("%s %s must have unique normalized (lowercase without underscores) name", reason, name))
			return tlast.BeautifulError2(e1, errNorm)
		}
	} else { // require unique names
		if okName {
			e1 := pr.BeautifulError(fmt.Errorf("%s %s name collision", reason, name))
			return tlast.BeautifulError2(e1, errName)
		}
		if okNorm {
			e1 := pr.BeautifulError(fmt.Errorf("%s %s normalized (lowercase without underscores) name collision", reason, name))
			return tlast.BeautifulError2(e1, errNorm)
		}
	}
	err := pr.BeautifulError(fmt.Errorf("see here"))
	if c.names == nil {
		c.names = make(map[string]*tlast.ParseError)
	}
	if c.namesNormalized == nil {
		c.namesNormalized = make(map[string]*tlast.ParseError)
	}
	c.names[name] = err
	c.namesNormalized[nn] = err
	return nil
}
