// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package main

import (
	"fmt"
	"log"

	"github.com/VKCOM/tl/internal/tlcodegen"
	"github.com/VKCOM/tl/internal/utils"
)

func main() {
	var options tlcodegen.Gen2Options
	parseFlags(&options)

	// mode to print version
	if options.PrintVersion {
		fmt.Println(utils.AppVersion())
		return
	}

	log.Printf("tlgen version: %s", utils.AppVersion())
	log.SetFlags(0)
	run(options)
}
