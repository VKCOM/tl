// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package main

import (
	"log"

	"github.com/vkcom/tl/internal/tlcodegen"
)

func main() {
	log.Printf("tlgen version: %s", tlcodegen.TLGenVersion())

	log.SetFlags(0)

	var options tlcodegen.Gen2Options

	parseFlags(&options)
	run(options)
}
