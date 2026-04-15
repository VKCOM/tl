// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package genphp

import (
	"fmt"
	"io"
	//"log"
	"os"
	"path/filepath"

	"strings"
	"time"

	"github.com/TwiN/go-color"
	"github.com/VKCOM/tl/internal/pure"
	"github.com/VKCOM/tl/internal/puregen"
	"github.com/VKCOM/tl/internal/tlast"
	"github.com/VKCOM/tl/internal/utils"
	"github.com/google/go-cmp/cmp"
	//"golang.org/x/exp/slices"
)

const BuiltinTupleName = "__tuple"
const BuiltinVectorName = "__vector"
const markerFile = "tlgen2_version.txt"
const EnableWarningsUnionNamespace = true
const EnableWarningsUnionNamePrefix = true
const EnableWarningsUnionNameExact = true
const EnableWarningsSimpleTypeName = true

const buildVersionFormat = `tlgen version: %s
schema url: %s
schema commit: %s
schema version: %d (%v)
`

// For debugging
var DEBUG = true

func Debugf(format string, args ...interface{}) {
	if DEBUG {
		fmt.Printf(format, args...)
	}
}

type Namespace struct {
	types []*TypeRWWrapper
	decGo Deconflicter
}

type Gen2Options struct {
	// General
	Language          string
	Outdir            string
	CopyrightFilePath string
	WarningsAreErrors bool
	Verbose           bool
	PrintDiff         bool
	ErrorWriter       io.Writer // all Errors and warnings should be redirected to this io.Writer, by default it is os.Stderr
	SplitInternal     bool
	AddMetaData       bool
	AddFactoryData    bool

	// TL2
	TL2WhiteList string

	// Linter
	Schema2Compare string

	// Linter php
	LinterPHPCheck                  bool
	LinterPHPNonPolymorphicBoxedRef bool

	// Go
	GenerateRPCCode      bool
	BytesWhiteList       string
	TypesWhiteList       string
	GenerateRandomCode   bool
	SchemaDocumentation  bool
	SchemaURL            string
	SchemaTimestamp      uint // for TLO version/date
	SchemaCommit         string
	UseCheckLengthSanity bool

	// C++
	RootCPP                string
	RootCPPNamespace       string
	SeparateFiles          bool
	GenerateCommonMakefile bool
	DeleteUnrelatedFiles   bool
	BasicTLNamespace       string
	GenerateFieldMasks     bool

	// PHP
	AddFunctionBodies            bool
	FunctionsBodiesWhiteList     string
	IgnoreUnusedInFunctionsTypes bool
	AddRPCTypes                  bool
	AddFetchers                  bool
	AddSwitcher                  bool
	AddFetchersEchoComments      bool
	InplaceSimpleStructs         bool
	UseBuiltinDataProviders      bool
	AddTypeComments              bool

	// PHP Unique actions
	CreateTLFilesWithAllTypesInReturn          bool
	CreateTLSplitedFilesForEachNamespace       bool
	CreateTLSplitedFilesForEachNamespaceFolder string

	// .tlo
	TLOPath           string
	CanonicalFormPath string // combinators in canonical form, with comment of source schema file path

	// Other modes
	PrintVersion bool
}

func (opt *Gen2Options) GenerateTL2() bool {
	return opt.TL2WhiteList != ""
}

type Gen2 struct {
	kernel *pure.Kernel

	// options
	options *puregen.Options // pointer so code modifying options in GenerateCode refers to the same structure

	// parsed TL
	allAnnotations []string // position is bit

	// generation
	builtinTypes       map[string]*TypeRWWrapper
	generatedTypes     map[string]*TypeRWWrapper
	generatedTypesList []*TypeRWWrapper // we need more deterministic order than sort predicate can establish

	globalDec  Deconflicter
	Namespaces map[string]*Namespace // Handlers Code is inside

	Code          map[string]string // fileName->Content, split by file names relative to output dir
	copyrightText string
}

func canonicalGoName(name tlast.Name, insideNamespace string) string {
	if name.Namespace == insideNamespace {
		return utils.CNameToCamelName(name.Name)
	}
	return utils.CNameToCamelName(name.Namespace) + utils.CNameToCamelName(name.Name)
}

func (gen *Gen2) getNamespace(n string) *Namespace {
	na, ok := gen.Namespaces[n]
	if !ok {
		na = &Namespace{}
		gen.Namespaces[n] = na
		// TODO - ALL golang-specific names
		na.decGo.deconflictName("Handler")
		na.decGo.deconflictName("Handle")
		na.decGo.deconflictName("Client")
		// TODO - if we want lowercase C++ identifiers, we need to add ~100 reserved keywords here
		// na.decCpp.deconflictName("double")
		// na.decCpp.deconflictName("int")
		// etc...
	}
	return na
}

func prepareNameFilter(filter string) []string {
	var result []string
	for _, str := range strings.Split(filter, ",") {
		str = strings.TrimSpace(str)
		if str == "" {
			continue
		}
		result = append(result, str)
	}
	return result
}

func inNameFilter(name tlast.Name, filters []string) bool {
	for _, filter := range filters {
		if inNameFilterElement(name, filter) {
			return true
		}
	}
	return false
}

func inNameFilterElement(name tlast.Name, filter string) bool {
	if filter == "*" {
		return true
	}
	if !strings.HasSuffix(filter, ".") {
		return name.String() == filter
	}
	return name.Namespace == strings.TrimSuffix(filter, ".")
}

func collectRelativePaths(absDirName string, relDirName string, relativeFiles map[string]bool, relativeDirs *[]string) error {
	fis, err := os.ReadDir(absDirName)
	if err != nil {
		return err
	}
	for _, fi := range fis { // try all snapshots, loading the latest
		relFilename := filepath.Join(relDirName, fi.Name())
		absFilename := filepath.Join(absDirName, fi.Name())
		if fi.IsDir() {
			*relativeDirs = append(*relativeDirs, relFilename)
			if err = collectRelativePaths(absFilename, relFilename, relativeFiles, relativeDirs); err != nil {
				return err
			}
			continue
		}
		relativeFiles[relFilename] = true
	}
	return nil
}

// WriteToDir Most common action with generated code, so clients do not repeat it
func (gen *Gen2) WriteToDir(outdir string) error {
	if err := os.Mkdir(outdir, 0755); err != nil && !os.IsExist(err) { // we thus require parent directory to exist
		return fmt.Errorf("error creating outdir %q: %w", outdir, err)
	}
	// We do not want to touch files which did not change at all.
	relativeFiles := map[string]bool{}
	var relativeDirs []string
	if err := collectRelativePaths(outdir, "", relativeFiles, &relativeDirs); err != nil {
		return fmt.Errorf("error reading outdir content %q: %w", outdir, err)
	}
	if len(relativeFiles) != 0 && !relativeFiles[markerFile] {
		return fmt.Errorf("outdir %q not empty and has no %q marker file, please clean manually", outdir, markerFile)
	}
	markerContent := fmt.Sprintf(buildVersionFormat,
		strings.TrimSpace(utils.AppVersion()),
		strings.TrimSpace(gen.options.SchemaURL),
		strings.TrimSpace(gen.options.SchemaCommit),
		gen.options.SchemaTimestamp, time.Unix(int64(gen.options.SchemaTimestamp), 0).UTC())
	if err := gen.addCodeFile(markerFile, markerContent); err != nil {
		return err
	}
	notTouched := 0
	written := 0
	deleted := 0
	for filepathName, code := range gen.Code {
		d := filepath.Join(outdir, filepath.Dir(filepathName))
		f := filepath.Join(outdir, filepathName)
		if !strings.HasPrefix(filepathName, "..") {
			// we allow relative paths outside gen folder for basictl*
			if err := os.MkdirAll(d, 0755); err != nil && !os.IsExist(err) {
				return fmt.Errorf("error creating dir %q: %w", d, err)
			}
		}
		if relativeFiles[filepathName] {
			delete(relativeFiles, filepathName)
			was, err := os.ReadFile(f)
			if err != nil {
				return fmt.Errorf("error reading previous file %q: %w", f, err)
			}
			if string(was) == code {
				notTouched++
				continue
			} else {
				Debugf("File \"%s\":\n", f)
				Debugf("%s\n", cmp.Diff(string(was), code))
			}
		}
		written++
		if err := os.WriteFile(f, []byte(code), 0644); err != nil {
			return fmt.Errorf("error writing file %q: %w", f, err)
		}
	}
	for filepathName := range relativeFiles {
		f := filepath.Join(outdir, filepathName)

		deleted++
		if err := os.Remove(f); err != nil {
			return fmt.Errorf("error deleting previous file %q: %w", f, err)
		}
	}
	for i := len(relativeDirs) - 1; i >= 0; i-- {
		f := filepath.Join(outdir, relativeDirs[i])
		_ = os.Remove(f) // non-empty dirs simply will not remove. This is good enough for us
	}
	// do not check Verbose
	fmt.Printf("%d target files did not change so were not touched, %d written, %d deleted\n", notTouched, written, deleted)
	return nil
}

func (gen *Gen2) addCodeFile(filepathName string, code string) error {
	if _, ok := gen.Code[filepathName]; ok {
		return fmt.Errorf("generator %s: source file %q is generated twice", color.InRed("internal error"), filepathName)
	}
	gen.Code[filepathName] = code
	return nil
}

func Generate(kernel *pure.Kernel, options *puregen.Options) error {
	options.Kernel.InstantiateConstants = false
	options.Kernel.InstantiateExclamationWrappers = true
	options.Kernel.NatArgsDelimiter = "_"
	options.Kernel.NotSimplifyNatArgs = true

	if !options.PHP.UseBuiltinDataProviders {
		panic("usage of \"UseBuiltinDataProviders\" is currently mandatory")
	}

	if err := kernel.Compile(); err != nil {
		return err
	}

	gen, err := generateCode(kernel, options)
	if gen == nil {
		return err
	}

	if err = gen.WriteToDir(options.Outdir); err != nil {
		return err // Context is already in err
	}
	return err
}

func generateCode(kernel *pure.Kernel, options *puregen.Options) (*Gen2, error) {
	if options.Kernel.Verbose || DEBUG {
		if DEBUG {
			Debugf(">>> [WARNING] DEBUG = true <<<\n")
		}
		DEBUG = true
		Debugf(">>> [WARNING] ENABLED DEBUG MODE <<<\n")
	}

	gen := &Gen2{
		kernel: kernel,

		options:    options,
		Code:       map[string]string{},
		Namespaces: map[string]*Namespace{},
		// Files:                 map[string][]*TypeRWWrapper{},
		builtinTypes:   map[string]*TypeRWWrapper{},
		generatedTypes: map[string]*TypeRWWrapper{},
	}

	if err := gen.compile(); err != nil {
		return nil, err
	}

	switch options.Language {
	case "php":
		{
			// TODO ADD FEATURE TO CHANGE IT
			gen.copyrightText = `/**
 * AUTOGENERATED, DO NOT EDIT! If you want to modify it, check tl schema.
 *
 * This autogenerated code represents tl class for typed RPC API.
 */

`
		}
		if err := gen.generateCodePHP(); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unsupported language: %s", options.Language)
	}

	return gen, nil
}
