// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package gencanonical

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/vkcom/tl/internal/pure"
	"github.com/vkcom/tl/internal/puregen"
	"github.com/vkcom/tl/internal/tlast"
)

func Generate(kernel *pure.Kernel, options *puregen.Options) error {
	if options.Verbose {
		log.Print("generating file with combinators in canonical form...")
	}
	if options.Outfile == "" {
		return fmt.Errorf("--outfile should not be empty")
	}
	var buf bytes.Buffer
	tlast.TL(kernel.TL1()).WriteGenerate2TL(&buf)
	if err := os.WriteFile(options.Outfile, buf.Bytes(), 0644); err != nil {
		return fmt.Errorf("error writing canonical form file %q: %w", options.Outfile, err)
	}
	return nil
}
