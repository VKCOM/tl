// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

//go:build !tlgen1

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/vkcom/tl/internal/tlast"
	"github.com/vkcom/tl/internal/tlcodegen"
	"github.com/vkcom/tl/internal/utils"
)

const tlExt = ".tl"

func parseFlags(opt *tlcodegen.Gen2Options) {
	// General
	flag.StringVar(&opt.Language, "language", "",
		`generation target language (go, cpp, php). Empty for linter.`)
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

	// General TL2
	flag.BoolVar(&opt.GenerateTL2, "tl2-generate", false,
		"generate code for tl2 methods (currently work only for golang)")

	// Linter
	flag.StringVar(&opt.Schema2Compare, "schema-to-compare", "",
		`path to old version TL schema to compare on backward compatibility`)

	// Go
	flag.StringVar(&opt.BasicPackageNameFull, "basicPkgPath", "",
		"if empty, 'basictl' package will be generated in output dir, otherwise imports will be generated")
	flag.StringVar(&opt.TLPackageNameFull, "pkgPath", "",
		"package path to be used inside generated code")
	flag.BoolVar(&opt.GenerateRPCCode, "generateRPCCode", false,
		"whether to generate *_server.go files")
	flag.StringVar(&opt.BasicRPCPath, "basicRPCPath", "",
		"path to rpc package")
	flag.StringVar(&opt.BytesWhiteList, "generateByteVersions", "",
		"comma-separated list of fully-qualified top-level types or namespaces (if have trailing '.'), to generate byte versions for. Empty means none, '.' means all.")
	flag.StringVar(&opt.TypesWhileList, "typesWhiteList", ".",
		"comma-separated list of fully-qualified top-level types or namespaces (if have trailing '.'), to generate code. Empty means none, '.' means all")
	flag.StringVar(&opt.TL2WhiteList, "tl2WhiteList", ".",
		"comma-separated list of fully-qualified top-level types or namespaces (if have trailing '.'), to generate TL2 code. Empty means none, '.' means all")
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
	flag.BoolVar(&opt.UseCheckLengthSanity, "checkLengthSanity", true,
		"enable feature to generate code to check length sanity of arrays (default:true)")

	// C++
	flag.StringVar(&opt.RootCPPNamespace, "cpp-namespace", "",
		`c++ root namespace, separated by '::' if more than 1 element`)
	flag.StringVar(&opt.RootCPP, "cpp-root", "",
		`c++ root package`)
	flag.BoolVar(&opt.AddMetaData, "cpp-generate-meta", false,
		`whether to generate methods to get meta information about tl objects`)
	flag.BoolVar(&opt.AddFactoryData, "cpp-generate-factory", false,
		`whether to generate factory of tl objects`)
	flag.BoolVar(&opt.GenerateCommonMakefile, "cpp-generate-common-makefile", true,
		`whether to generate Makefile in a root with all generated namespaces targets (default:true)`)
	flag.BoolVar(&opt.DeleteUnrelatedFiles, "cpp-delete-unrelated-files", true,
		`whether to delete files that are already in the target directory, but will not be affected by the new generation (default:true)`)

	// PHP
	flag.BoolVar(&opt.AddFunctionBodies, "php-serialization-bodies", false,
		`whether to generate body to write/read generated structs and functions`)
	flag.BoolVar(&opt.AddMetaData, "php-generate-meta", false,
		`whether to generate methods to get meta information about tl objects`)
	flag.BoolVar(&opt.AddFactoryData, "php-generate-factory", false,
		`whether to generate factory of tl objects`)
	flag.BoolVar(&opt.IgnoreUnusedInFunctionsTypes, "php-ignore-unused-types", true,
		`whether to not generate types without usages in functions (default:true)`)
	flag.BoolVar(&opt.AddRPCTypes, "php-rpc-support", true,
		`whether to generate special rpc types (default:true)`)
	flag.BoolVar(&opt.InplaceSimpleStructs, "php-inplace-simple-structs", true,
		`whether to avoid generation of structs with no arguments and only 1 field (default:true)`)

	if opt.AddFactoryData {
		opt.AddMetaData = true
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
	if opt.Language == "" && opt.Schema2Compare != "" {
		compErr := runCompatibilityCheck(opt, &ast)
		if compErr != nil {
			var parseError *tlast.ParseError
			if errors.As(compErr, &parseError) {
				if opt.WarningsAreErrors {
					return compErr
				}
				parseError.ConsolePrint(opt.ErrorWriter, compErr, true)
			} else {
				return compErr
			}
		}
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

func runCompatibilityCheck(opt *tlcodegen.Gen2Options, ast *tlast.TL) error {
	clonedOpt := *opt
	clonedOpt.ErrorWriter = io.Discard

	log.Print("STEP: Load old tl schema to compare...")

	parsedPaths, err := utils.WalkDeterministic(tlExt, clonedOpt.Schema2Compare)
	if err != nil {
		return err
	}
	oldTLPath := parsedPaths[0]
	oldTL, err := parseTlFile(oldTLPath, true, &clonedOpt)
	if err != nil {
		return err
	}

	log.Print("STEP: Compare old tl schema with passed...")

	if compErr := tlcodegen.CheckBackwardCompatibility(ast, &oldTL); compErr != nil {
		return compErr
	}
	log.Print("RESULT: New version is backward compatible with passed schema")
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
