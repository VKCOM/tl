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
	"strings"

	"github.com/pkg/errors"

	"github.com/vkcom/tl/internal/tlast"
	"github.com/vkcom/tl/internal/tlcodegen"
	"github.com/vkcom/tl/internal/utils"
)

const tlExt = ".tl"

func parseFlags(opt *tlcodegen.Gen2Options) {
	// General
	flag.StringVar(&opt.Language, "language", "",
		`generation target language (go, cpp). Empty for linter.`)
	ignoreGeneratedCode := false
	flag.BoolVar(&ignoreGeneratedCode, "ignoreGeneratedCode", false,
		"flag is ignored, because default generator is linting now")
	flag.StringVar(&opt.Outdir, "outdir", "",
		`where to write generated files`)
	flag.StringVar(&opt.CopyrightFilePath, "copyrightPath", "",
		"path to file with copyright text")
	flag.BoolVar(&opt.WarningsAreErrors, "Werror", false,
		"treat all warnings as errors")
	flag.BoolVar(&opt.Verbose, "v", false,
		"verbose mode that prints debug info")
	flag.BoolVar(&opt.PrintDiff, "print-diff", false,
		"prints diff of outdir contents before and after generating")
	flag.BoolVar(&opt.SplitInternal, "split-internal", false,
		"generated code will be split into independent packages (in a simple word: speeds up compilation)")

	// Go
	flag.StringVar(&opt.BasicPackageNameFull, "basicPkgPath", "",
		"if empty, 'basictl' package will be generated in output dir, otherwise imports will be generated")
	flag.StringVar(&opt.TLPackageNameFull, "pkgPath", "",
		"package path to be used inside generated code")
	flag.BoolVar(&opt.GenerateRPCCode, "generateRPCCode", false,
		"whether to generate *_server.go files")
	flag.StringVar(&opt.BasicRPCPath, "basicRPCPath", "",
		"path to rpc package")
	flag.StringVar(&opt.BytesVersions, "generateByteVersions", "",
		"comma-separated list of fully-qualified top-level types or namespaces (if have trailing '.'), to generate byte versions for. Empty means 'none'.")
	flag.StringVar(&opt.TypesWhileList, "typesWhiteList", "",
		"comma-separated list of fully-qualified top-level types or namespaces (if have trailing '.'), to generate code. Empty means 'all'")
	flag.BoolVar(&opt.GenerateRandomCode, "generateRandomCode", false,
		"whether to generate methods for random filling structs")
	flag.BoolVar(&opt.GenerateLegacyJsonRead, "generateLegacyJsonRead", false,
		"whether to generate methods to read json in old way")
	flag.BoolVar(&opt.SchemaDocumentation, "generateSchemaDocumentation", false,
		"whether to generate .html representation of schema in to tljson.html file")
	flag.StringVar(&opt.SchemaURL, "schemaURL", "",
		"url of schema (for documentation)")
	flag.UintVar(&opt.SchemaTimestamp, "schemaTimestamp", 0,
		"timestamp of schema (for documentation, TLO version)")
	flag.StringVar(&opt.SchemaCommit, "schemaCommit", "",
		"commit of schema (for documentation)")

	// C++
	flag.StringVar(&opt.RootCPPNamespace, "cpp-namespace", "",
		`c++ root namespace, separated by '::' if more than 1 element`)

	// PHP
	flag.BoolVar(&opt.AddFunctionBodies, "php-serialization-bodies", false,
		`whether to generate body to write/read generated structs and functions`)
	flag.BoolVar(&opt.AddMetaData, "php-generate-meta", false,
		`whether to generate methods to get meta information about tl objects`)
	flag.BoolVar(&opt.AddFactoryData, "php-generate-factory", false,
		`whether to generate factory of tl objects`)

	if opt.AddFactoryData {
		opt.AddFunctionBodies = true
	}
	// .tlo
	flag.StringVar(&opt.TLOPath, "tloPath", "",
		"whether to serialize TL schema in binary form")
	flag.StringVar(&opt.CanonicalFormPath, "canonicalFormPath", "",
		"generate file with combinators in canonical form")

	flag.Parse()
}

func run(opt tlcodegen.Gen2Options) {
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

func runMain(opt *tlcodegen.Gen2Options) error {
	var ast tlast.TL
	var fullAst tlast.TL
	if opt.ErrorWriter == nil {
		opt.ErrorWriter = os.Stdout
	}
	args := flag.Args()
	if len(args) == 0 {
		return fmt.Errorf("specify 1 or more input TL schema filenames after flags")
	}
	paths, err := utils.WalkDeterministic(tlExt, args...)
	if err != nil {
		return fmt.Errorf("error while walkking through paths: %w", err)
	}
	for _, path := range paths {
		tl, err := parseTlFile(path, true, opt)
		if err != nil {
			return err
		}
		fullTl, err := parseTlFile(path, false, opt)
		if err != nil {
			return err
		}
		ast = append(ast, tl...)
		fullAst = append(fullAst, fullTl...)
	}
	gen, err := tlcodegen.GenerateCode(ast, *opt)
	if err != nil {
		return err // Do not add excess info to already long parse error
	}
	if opt.Language != "" {
		if opt.Outdir == "" {
			return fmt.Errorf("--outdir should not be empty")
		}
		if err = gen.WriteToDir(opt.Outdir); err != nil {
			return err // Context is already in err
		}
	}
	if opt.TLOPath != "" {
		if opt.Verbose {
			log.Print("generating TLO file...")
		}
		s, err := fullAst.GenerateTLO(uint32(opt.SchemaTimestamp))
		if err != nil {
			return fmt.Errorf("error on generating tlo: %w", err)
		}
		buf, err := s.WriteBoxed(nil)
		if err != nil {
			return fmt.Errorf("error writing boxed tlo: %w", err)
		}
		if err := os.WriteFile(opt.TLOPath, buf, 0644); err != nil {
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
		if err := os.WriteFile(opt.TLOPath+".json", prettyJSON.Bytes(), 0644); err != nil {
			return fmt.Errorf("error writing tlo json file: %w", err)
		}
	}
	if opt.CanonicalFormPath != "" {
		if opt.Verbose {
			log.Print("generating file with combinators in canonical form,,,")
		}
		var buf bytes.Buffer
		fullAst.WriteGenerate2TL(&buf)
		if err := os.WriteFile(opt.CanonicalFormPath, buf.Bytes(), 0644); err != nil {
			return err
		}
	}
	return nil
}

func parseTlFile(file string, replaceStrange bool, opt *tlcodegen.Gen2Options) (tlast.TL, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("error reading schema file %q - %w", file, err)
	}
	dataStr := string(data)
	// Exceptions we cannot fix upstream
	if replaceStrange {
		dataStr = strings.ReplaceAll(dataStr, "_ {X:Type} result:X = ReqResult X;", "")
		dataStr = strings.ReplaceAll(dataStr, "engine.query {X:Type} query:!X = engine.Query;", "")
		dataStr = strings.ReplaceAll(dataStr, "engine.queryShortened query:%(VectorTotal int) = engine.Query;", "")
		//		dataStr = strings.ReplaceAll(dataStr, "engine.queryShortened query:%(VectorTotal int) = engine.Query;", "")
		//		dataStr = strings.ReplaceAll(dataStr, "@any @internal engine.sendResponseTo {X:Type} pid:%net.Pid query:!X = Bool;", "")
		//		dataStr = strings.ReplaceAll(dataStr, "@internal @write messages.responseQuery {X:Type} response_query_id:long query:!X = X;", "")
		//		dataStr = strings.ReplaceAll(dataStr, "@any rpcDestActor#7568aabd {X:Type} actor_id:long query:!X = X;", "")
		//		dataStr = strings.ReplaceAll(dataStr, "@any rpcDestActorFlags#f0a5acf7 {X:Type} actor_id:long flags:# extra:%(RpcInvokeReqExtra flags) query:!X = X;", "")
		//		dataStr = strings.ReplaceAll(dataStr, "@any rpcDestFlags#e352035e {X:Type} flags:# extra:%(RpcInvokeReqExtra flags) query:!X = X;", "")
		//		dataStr = strings.ReplaceAll(dataStr, "@any rpcInvokeReq#2374df3d {X:Type} query_id:long query:!X = RpcReqResult X;", "")
	}
	tl, err := tlast.ParseTLFile(dataStr, file, tlast.LexerOptions{
		AllowBuiltin: false,
		AllowDirty:   !replaceStrange,
		AllowMLC:     !opt.WarningsAreErrors,
	}, opt.ErrorWriter)
	if err != nil {
		return tl, err // Do not add excess info to already long parse error
	}
	return tl, nil
}
