// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/vkcom/tl/internal/pure"
	"github.com/vkcom/tl/internal/puregen"
	"github.com/vkcom/tl/internal/puregen/gengo"
	"github.com/vkcom/tl/internal/tlast"
	"github.com/vkcom/tl/internal/utils"
)

func main() {
	log.Printf("tl2gen version: %s", utils.AppVersion())

	log.SetFlags(0)

	opt := puregen.Options{
		ErrorWriter: os.Stdout,
	}
	opt.Bind(flag.CommandLine)

	flag.Parse()

	if err := runMain(&opt); err != nil {
		var parseError *tlast.ParseError
		if errors.As(err, &parseError) {
			parseError.ConsolePrint(opt.ErrorWriter, err, false)
		} else {
			log.Println(err.Error())
		}
		if opt.Language == "" {
			log.Printf("TL Linter Failed") // do not check Verbose
		} else {
			log.Printf("TL Generation Failed") // do not check Verbose
		}
		os.Exit(1)
		return
	}
	if opt.Language == "" {
		log.Printf("TL Linter Success") // do not check Verbose
	} else {
		log.Printf("TL Generation Success") // do not check Verbose
	}
}

func runMain(opt *puregen.Options) error {
	if err := opt.Validate(); err != nil {
		return err
	}

	args := flag.Args()
	if len(args) == 0 {
		return fmt.Errorf("specify 1 or more input TL schema filenames after flags")
	}

	kernel := pure.NewKernel(&opt.Kernel)

	// parse tl1
	pathsTL1, err := utils.WalkDeterministic(".tl", args...)
	if err != nil {
		return fmt.Errorf("error while walking through paths: %w", err)
	}
	for _, path := range pathsTL1 {
		if err := kernel.AddFileTL1(path); err != nil {
			return err
		}
	}
	// parse tl2
	pathsTL2, err := utils.WalkDeterministic(".tl2", args...)
	if err != nil {
		return fmt.Errorf("error while walking through tl2 paths: %w", err)
	}
	for _, path := range pathsTL2 {
		if err := kernel.AddFileTL2(path); err != nil {
			return err
		}
	}
	if err := kernel.Compile(); err != nil {
		return err
	}
	switch opt.Language {
	case "":
		return nil
	case "go":
		return gengo.Generate(kernel, opt)
	case "tlo":
		return fmt.Errorf("TODO generate TLO here")
	case "tljson.html":
		return fmt.Errorf("TODO generate tljson.html")
	default:
		return fmt.Errorf("unsupported language, must be 'go', 'tlo', 'htmldoc' or empty for linter: %q", opt.Language)
	}
}
