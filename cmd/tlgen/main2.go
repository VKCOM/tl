// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

//go:build !tlgen1

package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"strings"

	"github.com/pkg/errors"

	"github.com/vkcom/tl/internal/tlast"
	"github.com/vkcom/tl/internal/tlcodegen"
	"github.com/vkcom/tl/internal/utils"
)

const tlExt = ".tl"

func addFlags(argv *arguments) {
	flag.StringVar(&argv.Language, "language", "go",
		`generation target language`)
	flag.StringVar(&argv.RootCPPNamespace, "cpp-namespace", "",
		`c++ root namespace, separated by '::' if more than 1 element`)
	flag.StringVar(&argv.BasicPackageNameFull, "basicPkgPath", "",
		"if empty, 'basictl' package will be generated in output dir, otherwise imports will be generated")
	flag.BoolVar(&argv.GenerateRandomCode, "generateRandomCode", false,
		"whether to generate methods for random filling structs")
	flag.BoolVar(&argv.GenerateLegacyJsonRead, "generateLegacyJsonRead", false,
		"whether to generate methods to read json in old way")
	flag.BoolVar(&argv.GenerateRPCCode, "generateRPCCode", true,
		"whether to generate *_server.go files")
	flag.StringVar(&argv.BasicRPCPath, "basicRPCPath", "",
		"path to rpc package")
	flag.StringVar(&argv.TLOPath, "tloPath", "",
		"whether to serialize TL schema in binary form")
	flag.StringVar(&argv.CanonicalFormPath, "canonicalFormPath", "",
		"generate file with combinators in canonical form")
	flag.BoolVar(&argv.SchemaDocumentation, "generateSchemaDocumentation", false,
		"whether to generate .html representation of schema in to tljson.html file")
	flag.StringVar(&argv.SchemaURLTemplate, "schemaURLTemplate", "",
		"template for url to current schema if documentation is generated")
	flag.BoolVar(&argv.SplitInternal, "split-internal", false,
		"generated code will be split into independent packages (in a simple word: speeds up compilation)")
	flag.StringVar(&argv.TypesWhileList, "typesWhiteList", "",
		"comma-separated list of fully-qualified top-level types or namespaces (if have trailing '.'), to generate code. Empty means 'all'")
	flag.StringVar(&argv.CopyrightFilePath, "copyrightPath", "",
		"path to file with copyright text")
	flag.BoolVar(&argv.WarningsAreErrors, "Werror", false,
		"treat all warnings as errors")
	flag.BoolVar(&argv.IgnoreGeneratedCode, "ignoreGeneratedCode", false,
		"ignores generated code, tlo and documentation will be generated with related flags")
}

func run(argv arguments) {
	var commit, version = func() (commit string, version string) {
		if info, ok := debug.ReadBuildInfo(); ok {
			for _, setting := range info.Settings {
				if setting.Key == "vcs.revision" {
					commit = setting.Value
					break
				}
			}
			version = info.Main.Version
		}
		return commit, version
	}()
	log.Printf("tlgen version: %s, commit: %s", version, commit)
	if err := runMain(&argv); err != nil {
		var parseError *tlast.ParseError
		if errors.As(err, &parseError) {
			parseError.ConsolePrint(argv.ErrorWriter, err, false)
		} else {
			log.Println(err.Error())
		}
		log.Printf("TL Generation Failed")
		os.Exit(1)
	} else {
		if argv.Language == "" {
			log.Printf("TL Linter Success")
		} else {
			log.Printf("TL Generation Success")
		}
	}
}

func runMain(argv *arguments) error {
	var ast tlast.TL
	var fullAst tlast.TL
	var args []string
	if argv.ErrorWriter == nil {
		argv.ErrorWriter = os.Stdout
	}
	if argv.IgnoreGeneratedCode {
		argv.Language = ""
	}
	if argv.SchemaFileName != "" {
		return fmt.Errorf("--schema argument is removed, specify 1 or more input TL schema filenames after flags")
	}
	args = append(args, flag.Args()...)
	if len(args) == 0 {
		return fmt.Errorf("specify 1 or more input TL schema filenames after flags")
	}
	paths, err := utils.WalkDeterministic(tlExt, args...)
	if err != nil {
		return fmt.Errorf("error while walkking through paths: %w", err)
	}
	for _, path := range paths {
		tl, err := parseTlFile(path)
		if err != nil {
			return err
		}
		fullTl, err := parseFullTlFile(path)
		if err != nil {
			return err
		}
		ast = append(ast, tl...)
		fullAst = append(fullAst, fullTl...)
	}
	gen, err := tlcodegen.GenerateCode(ast, argv.Gen2Options)
	if err != nil {
		return err // Do not add excess info to already long parse error
	}
	if argv.Language != "" {
		if argv.Outdir == "" {
			return fmt.Errorf("--outdir should not be empty")
		}
		if err = gen.WriteToDir(argv.Outdir); err != nil {
			return err // Context is already in err
		}
	}
	if argv.TLOPath != "" {
		if argv.Verbose {
			log.Print("generating tlo file")
		}
		s, err := fullAst.GenerateTLO()
		if err != nil {
			return fmt.Errorf("error on generating tlo: %w", err)
		}
		buf, err := s.WriteBoxed(nil)
		if err != nil {
			return fmt.Errorf("error writing boxed tlo: %w", err)
		}
		if err := os.WriteFile(argv.TLOPath, buf, 0644); err != nil {
			return fmt.Errorf("error writing tlo file: %w", err)
		}
		buf, err = s.WriteJSON(nil)
		if err != nil {
			return fmt.Errorf("error writing json tlo: %w", err)
		}
		var prettyJSON bytes.Buffer
		if err = json.Indent(&prettyJSON, buf, "", "  "); err != nil {
			return fmt.Errorf("error pretty printing json tlo: %w", err)
		}
		if err := os.WriteFile(argv.TLOPath+".json", prettyJSON.Bytes(), 0644); err != nil {
			return fmt.Errorf("error writing tlo json file: %w", err)
		}
	}
	if argv.CanonicalFormPath != "" {
		if argv.Verbose {
			log.Print("generating file with combinators in canonical form")
		}
		var buf bytes.Buffer
		fullAst.WriteGenerate2TL(&buf)
		if err := os.WriteFile(argv.CanonicalFormPath, buf.Bytes(), 0644); err != nil {
			return err
		}
	}
	return nil
}

func parseTlFile(file string) (tlast.TL, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("error reading schema file %q - %w", file, err)
	}
	// Exceptions we cannot fix upstream
	dataStr := strings.ReplaceAll(string(data), "_ {X:Type} result:X = ReqResult X;", "")
	dataStr = strings.ReplaceAll(dataStr, "engine.query {X:Type} query:!X = engine.Query;", "")
	dataStr = strings.ReplaceAll(dataStr, "engine.queryShortened query:%(VectorTotal int) = engine.Query;", "")

	tl, err := tlast.ParseTLFile(dataStr, file, false)
	if err != nil {
		return tl, err // Do not add excess info to already long parse error
	}
	return tl, nil
}

func parseFullTlFile(file string) (tlast.TL, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("error reading schema file %q - %w", file, err)
	}
	// Exceptions we cannot fix upstream
	tl, err := tlast.ParseTLFile(string(data), file, true)
	if err != nil {
		return tl, err // Do not add excess info to already long parse error
	}
	return tl, nil
}
