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
	//"net/http"
	//_ "net/http/pprof" // Import for side effects: registers pprof handlers
	"os"
	"runtime/pprof"
	"sort"
	"strings"

	"github.com/vkcom/tl/internal/pure"
	"github.com/vkcom/tl/internal/puregen"
	"github.com/vkcom/tl/internal/puregen/gencanonical"
	"github.com/vkcom/tl/internal/puregen/gengo"
	"github.com/vkcom/tl/internal/puregen/gentljsonhtml"
	"github.com/vkcom/tl/internal/puregen/gentlo"
	"github.com/vkcom/tl/internal/tlast"
	"github.com/vkcom/tl/internal/utils"
)

var languages = map[string]func(kernel *pure.Kernel, options *puregen.Options) error{
	"canonical":    gencanonical.Generate,
	"go":           gengo.Generate,
	"lint":         func(kernel *pure.Kernel, options *puregen.Options) error { return nil }, // nothing more than lint
	"tl2migration": func(kernel *pure.Kernel, options *puregen.Options) error { return kernel.Migration() },
	"tljson.html":  gentljsonhtml.Generate,
	"tlo":          gentlo.Generate,
}

func languagesString() string {
	var keys []string
	for k := range languages {
		keys = append(keys, fmt.Sprintf(`"%s"`, k))
	}
	sort.Strings(keys)
	return fmt.Sprintf("one of %s", strings.Join(keys, ", "))
}

func main() {
	fmt.Printf("tl2gen version: %s\n", utils.AppVersion())

	//runtime.SetMutexProfileFraction(1)
	//runtime.SetBlockProfileRate(1)
	//go func() {
	//	if err := http.ListenAndServe(":8821", nil); err != nil { // Use nil for default ServeMux
	//		fmt.Println(err)
	//	}
	//}()
	//log.SetFlags(0)

	opt := puregen.Options{
		ErrorWriter: os.Stdout,
	}
	opt.Bind(flag.CommandLine, languagesString())

	flag.Parse()

	if err := runMain(&opt); err != nil {
		var parseError *tlast.ParseError
		if errors.As(err, &parseError) {
			parseError.ConsolePrint(opt.ErrorWriter, err, false)
		} else {
			log.Println(err.Error())
		}
		if opt.Language == "lint" {
			log.Printf("TL Linter Failed") // do not check Verbose
		} else {
			log.Printf("TL Generation Failed") // do not check Verbose
		}
		os.Exit(1)
		return
	}
	if opt.Language == "lint" {
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

	if opt.ProfileCPU != "" {
		f, err := os.Create(opt.ProfileCPU)
		if err != nil {
			fmt.Printf("could not create CPU profile file %s: %v\n", opt.ProfileCPU, err)
		} else {
			defer f.Close() // error handling omitted for example
			if err := pprof.StartCPUProfile(f); err != nil {
				fmt.Printf("could not start CPU profile: %v\n", err)
			} else {
				fmt.Printf("starting CPU profiling, to ananlyze run 'go tool pprof -http=:8080 %s'\n", opt.ProfileCPU)
				defer pprof.StopCPUProfile()
			}
		}
	}

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
	if f, ok := languages[opt.Language]; ok {
		if err := f(kernel, opt); err != nil {
			return err
		}
		kernel.PrintUnusedWarnings()
		return opt.ReplaceStringInDir()
	}
	var keys []string
	for k := range languages {
		keys = append(keys, fmt.Sprintf("'%s'", k))
	}
	sort.Strings(keys)
	return fmt.Errorf("unsupported language %q, must be %s", opt.Language, languagesString())
}
