// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package gentlo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/vkcom/tl/internal/pure"
	"github.com/vkcom/tl/internal/puregen"
	"github.com/vkcom/tl/internal/tlast"
)

func Generate(kernel *pure.Kernel, options *puregen.Options) error {
	if options.Verbose {
		log.Print("generating TLO file...")
	}
	if options.Outfile == "" {
		return fmt.Errorf("--outfile should not be empty")
	}

	s, err := tlast.TL(kernel.TL1()).GenerateTLO(uint32(options.SchemaTimestamp))
	if err != nil {
		return fmt.Errorf("error on generating tlo: %w", err)
	}
	buf, err := s.WriteBoxed(nil)
	if err != nil {
		return fmt.Errorf("error writing boxed tlo: %w", err)
	}
	if err := os.WriteFile(options.Outfile, buf, 0644); err != nil {
		return fmt.Errorf("error writing tlo file %q: %w", options.Outfile, err)
	}
	buf, err = s.WriteJSON(nil)
	if err != nil {
		return fmt.Errorf("error writing json tlo: %w", err)
	}
	var prettyJSON bytes.Buffer
	if err = json.Indent(&prettyJSON, buf, "", "  "); err != nil {
		return fmt.Errorf("error pretty printing json tlo: %w", err)
	}
	if err := os.WriteFile(options.Outfile+".json", prettyJSON.Bytes(), 0644); err != nil {
		return fmt.Errorf("error writing tlo json file %s: %w", options.Outfile+".json", err)
	}
	return nil
}
