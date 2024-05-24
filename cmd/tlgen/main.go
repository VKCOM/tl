// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package main

import (
	"flag"
	"fmt"
	"log"
	"runtime/debug"

	"github.com/vkcom/tl/internal/tlcodegen"
)

type arguments struct {
	tlcodegen.Gen2Options

	TLOPath           string
	CanonicalFormPath string // combinators in canonical form, with comment of source schema file path
	Outdir            string
}

func commonFlags(argv *arguments) {
	flag.StringVar(&argv.Outdir, "outdir", "",
		`where to write generated files`)
	flag.StringVar(&argv.TLPackageNameFull, "pkgPath", "",
		"package path to be used inside generated code")
	flag.StringVar(&argv.BytesVersions, "generateByteVersions", "",
		"comma-separated list of fully-qualified top-level types or namespaces (if have trailing '.'), to generate byte versions for. Empty means 'none'.")
	flag.BoolVar(&argv.Verbose, "v", false,
		"verbose mode that prints debug info")
}

func main() {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		log.Printf("Failed to read build info")
		return
	}
	fmt.Printf("tlgen version: %+v\n", bi.Main.Version)

	log.SetFlags(0)

	var argv arguments

	commonFlags(&argv)
	addFlags(&argv)

	flag.Parse()

	run(argv)
}
