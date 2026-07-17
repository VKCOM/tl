// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package compat

import (
	"fmt"
	"strings"

	"github.com/VKCOM/tl/internal/pure"
	"github.com/VKCOM/tl/internal/puregen"
)

// Run is the entry point for the --language=compat mode of cmd/tl2gen. It has the same
// signature as the language generators (e.g. gengo.Generate), so it plugs into the languages
// map directly, keeping main.go free of wiring logic.
//
// kernel holds the current (new) schema. The previous (old) schema is taken from
// opt.BackwardCompatibilityWith. Prototype: only TL1 combinators are checked.
func Run(kernel *pure.Kernel, opt *puregen.Options) error {
	if opt.BackwardCompatibilityWith == "" {
		return fmt.Errorf("--language=compat requires --backwardCompatibilityWith=<path(s) to previous schema version>")
	}
	// Validate the current (new) schema fully.
	if err := kernel.Compile(); err != nil {
		return err
	}
	// Parse the previous (old) schema. We do not Compile it: it is historical and only needs
	// to be readable as a set of combinators to compare against.
	prev := pure.NewKernel(&opt.Kernel)
	prevPaths := strings.Split(opt.BackwardCompatibilityWith, ",")
	for i := range prevPaths {
		prevPaths[i] = strings.TrimSpace(prevPaths[i])
	}
	if err := prev.AddFilesFromPaths(prevPaths); err != nil {
		return err
	}
	if pe := CheckBackwardCompatibility(prev, kernel, opt.Kernel.ErrorWriter); pe != nil {
		return pe
	}
	return nil
}
