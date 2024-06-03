// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package main

import (
	"fmt"
	"log"
	"runtime/debug"

	"github.com/vkcom/tl/internal/tlcodegen"
)

func main() {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		log.Printf("Failed to read build info")
		return
	}
	fmt.Printf("tlgen version: %+v\n", bi.Main.Version)

	log.SetFlags(0)

	var options tlcodegen.Gen2Options

	parseFlags(&options)
	run(options)
}
