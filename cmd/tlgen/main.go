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
	"os"
	"strings"

	"github.com/pkg/errors"

	"github.com/vkcom/tl/internal/tlast"
	"github.com/vkcom/tl/internal/tlcodegen"
	"github.com/vkcom/tl/internal/tlcodegen/gen_tlo"
	"github.com/vkcom/tl/internal/utils"
)

type arguments struct {
	tlcodegen.Gen2Options
	outdir         string
	schemaFileName string
	strict         bool
	printTLOAsJSON string
}

type context struct {
	// argv хранит аргументы командной строки. Readonly после инициализации.
	argv arguments
}

func (ctx *context) parseFlags() error {
	flag.StringVar(&ctx.argv.Language, "language", "",
		`generation target language: go, cpp`)
	flag.StringVar(&ctx.argv.RootCPPNamespace, "cpp-namespace", "",
		`c++ root namespace, separated by '::' if more than 1 element`)
	flag.StringVar(&ctx.argv.outdir, "outdir", "",
		`where to write generated files; "" for stdout`)
	flag.StringVar(&ctx.argv.TLPackageNameFull, "pkgPath", "",
		"package path to be used inside generated code")
	flag.StringVar(&ctx.argv.BasicPackageNameFull, "basicPkgPath", "",
		"if empty, 'basictl' package will be generated in output dir, otherwise imports will be generated")
	flag.StringVar(&ctx.argv.schemaFileName, "schema", "",
		"input TL schema in binary format")
	flag.BoolVar(&ctx.argv.Verbose, "v", false,
		"verbose mode that prints debug info")
	flag.BoolVar(&ctx.argv.strict, "strict", false,
		"don't generate any code if can't cover 100% of the scheme")
	flag.BoolVar(&ctx.argv.GenerateRandomCode, "generateRandomCode", false,
		"whether to generate methods for random filling structs")
	flag.BoolVar(&ctx.argv.GenerateRPCCode, "generateRPCCode", true,
		"whether to generate *_server.go files")
	flag.StringVar(&ctx.argv.BasicRPCPath, "basicRPCPath", "",
		"path to rpc package")
	flag.StringVar(&ctx.argv.TLOPath, "tloPath", "",
		"whether to serialize TL schema in binary form")
	flag.BoolVar(&ctx.argv.SchemaDocumentation, "generateSchemaDocumentation", false,
		"whether to generate .html representation of schema in to tljson.html file")
	flag.BoolVar(&ctx.argv.SplitInternal, "split-internal", false,
		"generated code will be split into independent packages (in a simple word: speeds up compilation)")
	flag.StringVar(&ctx.argv.TypesWhileList, "typesWhiteList", "",
		"comma-separated list of fully-qualified top-level types or namespaces (if have trailing '.'), to generate code. Empty means 'all'")
	flag.StringVar(&ctx.argv.BytesVersions, "generateByteVersions", "",
		"comma-separated list of fully-qualified top-level types or namespaces (if have trailing '.'), to generate byte versions for. Empty means 'none'.")
	flag.StringVar(&ctx.argv.CopyrightFilePath, "copyrightPath", "",
		"path to file with copyright text")
	flag.BoolVar(&ctx.argv.IgnoreGeneratedCode, "ignoreGeneratedCode", false,
		"ignores generated code, tlo and documentation will be generated with related flags")
	flag.StringVar(&ctx.argv.printTLOAsJSON, "printTloAsJson", "",
		"accepts path to tlo file and writes equivalent json file nearby")
	flag.Parse()

	if ctx.argv.TLPackageNameFull == "" {
		return errors.New("can't use empty pkgName")
	}
	if ctx.argv.GenerateRPCCode && ctx.argv.BasicRPCPath == "" {
		return errors.New("flag '-generateRPCCode' is set but '-basicRPCPath' is empty")
	}
	ctx.argv.TLPackageNameFull = strings.TrimSuffix(ctx.argv.TLPackageNameFull, "/") // См. BACK-4267

	return nil
}

func main() {
	log.SetFlags(0)
	ctx := &context{}
	if err := ctx.parseFlags(); err != nil {
		log.Printf("parse flags: %v", err)
		os.Exit(1)
	}
	if err := runMain(ctx.argv); err != nil {
		var parseError *tlast.ParseError
		if errors.As(err, &parseError) {
			parseError.ConsolePrint(err)
		} else {
			log.Println(utils.ErrorStringWithStack(err))
		}
		log.Printf("TL Generation Failed")
		os.Exit(1)
	} else {
		log.Printf("TL Generation Success")
	}
}

const tlExt = ".tl"

func printTLOAsJSON(pathToFile string) error {
	file, err := os.ReadFile(pathToFile)
	if err != nil {
		return fmt.Errorf("error on reading file %s: %w", pathToFile, err)
	}
	var s4 gen_tlo.TlsSchemaV4
	if _, err = s4.ReadBoxed(file); err != nil {
		return fmt.Errorf("error on readinf TLO from %s: %w", pathToFile, err)
	}
	out, err := s4.WriteJSON(nil)
	if err != nil {
		return fmt.Errorf("error on creating json from %s: %w", pathToFile, err)
	}

	if err := os.WriteFile(
		pathToFile+".json",
		utils.JsonPrettyPrint(out),
		0644); err != nil {
		return fmt.Errorf("error on writing json to file %s: %w", pathToFile+".json", err)
	}
	return nil
}

func runMain(argv arguments) error {
	if argv.printTLOAsJSON != "" {
		if err := printTLOAsJSON(argv.printTLOAsJSON); err != nil {
			return err
		}
	}
	if argv.Verbose {
		log.Printf("No more awful TLO! Everyone happy!")
	}
	var ast tlast.TL
	var args []string
	if argv.schemaFileName != "" {
		args = append(args, argv.schemaFileName)
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
		ast = append(ast, tl...)
		if err != nil {
			return err
		}
	}

	if argv.Verbose {
		log.Printf("parsing TL...")
	}
	argv.TLPackageNameFull = strings.TrimSpace(argv.TLPackageNameFull)
	if argv.TLPackageNameFull == "" { // TODO - better validation
		return fmt.Errorf("--go-tl-package cannot be empty")
	}
	gen, err := tlcodegen.GenerateCode(ast, argv.Gen2Options)
	if err != nil {
		return err // Do not add excess info to already long parse error
	}
	if err = gen.WriteToDir(argv.outdir); err != nil {
		return err // Context is already in err
	}
	if err = gen.WriteTLO(); err != nil {
		return err
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

	tl, err := tlast.ParseTLFile(dataStr, file)
	if err != nil {
		return tl, err // Do not add excess info to already long parse error
	}
	return tl, nil
}
