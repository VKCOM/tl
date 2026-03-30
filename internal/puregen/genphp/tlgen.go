// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package genphp

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/TwiN/go-color"
	"github.com/VKCOM/tl/internal/pure"
	"github.com/VKCOM/tl/internal/puregen"
	"github.com/VKCOM/tl/internal/purelegacy"
	"github.com/VKCOM/tl/internal/tlast"
	"github.com/VKCOM/tl/internal/utils"
	"github.com/google/go-cmp/cmp"
	"golang.org/x/exp/slices"
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

var (
	errSeeHere                = fmt.Errorf("see here")
	errFieldNameCollision     = fmt.Errorf("field name collision")
	errNatParamNameCollision  = fmt.Errorf("nat-parameter name collision")
	errTypeParamNameCollision = fmt.Errorf("type-parameter name collision ")
)

// For debugging
var DEBUG = false

func Debugf(format string, args ...interface{}) {
	if DEBUG {
		fmt.Printf(format, args...)
	}
}

type LocalResolveContext struct {
	localTypeArgs map[string]LocalTypeArg
	localNatArgs  map[string]LocalNatArg

	allowAnyConstructor bool   // we can reference all constructors (functions, union elements) directly internally
	overrideFileName    string // used for unions and built-in vectors and tuples, so they are defined in the file of argument
}

// checkArgsCollision checks if passed name is already used in local context.
// pr: PR of the name we want to check
// err: will be returned (wrapped in beautiful error) if collision was NOT in type-parameter,
// must be defined depending on call context
func (lrc *LocalResolveContext) checkArgsCollision(name string, pr tlast.PositionRange, err error) error {
	if nat, ok := lrc.localNatArgs[name]; ok {
		e1 := pr.BeautifulError(err)
		e2 := nat.NamePR.BeautifulError(errSeeHere)
		return tlast.BeautifulError2(e1, e2)
	}
	if typ, ok := lrc.localTypeArgs[name]; ok {
		e1 := pr.BeautifulError(errTypeParamNameCollision)
		e2 := typ.PR.BeautifulError(errSeeHere)
		return tlast.BeautifulError2(e1, e2)
	}
	return nil
}

type LocalNatArg struct {
	wrongTypeErr error // we must add all field names to local context, because they must correctly shadow names outside, but we check the type

	NamePR tlast.PositionRange
	TypePR tlast.PositionRange
	natArg ActualNatArg
}

type LocalTypeArg struct {
	arg     ResolvedArgument
	PR      tlast.PositionRange // original template arg reference
	natArgs []ActualNatArg      // nat args associated with this type argument, if type argument itself has some nat args
}

type ResolvedArgument struct {
	isNat   bool
	isArith bool
	Arith   tlast.Arithmetic
	tip     *TypeRWWrapper
	bare    bool // vector Int is not the same as vector int, we must capture the difference somewhere
}

type ActualNatArg struct {
	isArith    bool
	Arith      tlast.Arithmetic
	isField    bool // otherwise it is # param with name
	FieldIndex int
	name       string // param name
}

func (arg *ActualNatArg) IsNumber() bool {
	return arg.isArith
}

func (arg *ActualNatArg) Number() uint32 {
	return arg.Arith.Res
}

func (arg *ActualNatArg) IsField() bool {
	return arg.isField
}

func (arg *ActualNatArg) IsTL2() bool {
	return arg.isField && arg.FieldIndex < 0
}

type HalfResolvedArgument struct { // TODO - better name
	Name string                 // if empty, this is not argument position
	Args []HalfResolvedArgument // recursion
}

type Namespace struct {
	types        []*TypeRWWrapper
	cppTemplates map[string]*TypeRWWrapper // canonical C++ template names like cross, cross<I,J>, cross<i>, cross<J> to avoid duplicates
	decGo        Deconflicter
	decCpp       Deconflicter
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
	// options
	options *puregen.Options // pointer so code modifying options in GenerateCode refers to the same structure

	// parsed TL
	supportedAnnotations map[string]struct{}
	typeDescriptors      map[string][]*tlast.Combinator
	singleConstructors   map[string]*tlast.Combinator // only types with 1 constructor, no functions
	allConstructors      map[string]*tlast.Combinator // for printing beautiful errors
	allAnnotations       []string                     // position is bit

	// generation
	builtinTypes       map[string]*TypeRWWrapper
	generatedTypes     map[string]*TypeRWWrapper
	generatedTypesList []*TypeRWWrapper // we need more deterministic order than sort predicate can establish

	globalDec  Deconflicter
	Namespaces map[string]*Namespace // Handlers Code is inside

	Code          map[string]string // fileName->Content, split by file names relative to output dir
	copyrightText string

	// new options
	typesInfo       *TypesInfo
	componentsOrder []int
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
		na = &Namespace{cppTemplates: map[string]*TypeRWWrapper{}}
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

func checkTagCollisions(tl tlast.TL) error {
	constructorTags := map[uint32]*tlast.Combinator{}
	for _, typ := range tl {
		crc32 := typ.Crc32()
		if crc32 == 0 {
			// typeA#00000000 = TypeA;
			return typ.Construct.IDPR.BeautifulError(fmt.Errorf("constructor tag 0 is prohibited, even if generated implicitly"))
		}
		if col, ok := constructorTags[crc32]; ok {
			// typeA#dfc15abf = TypeA;
			// typeB#dfc15abf = TypeB;
			e1 := typ.Construct.IDPR.BeautifulError(fmt.Errorf("constructor tag #%08x used by %q is used again by %q", crc32, col.Construct.Name.String(), typ.Construct.Name.String()))
			e2 := col.Construct.IDPR.BeautifulError(errSeeHere)
			return tlast.BeautifulError2(e1, e2)
		}
		constructorTags[crc32] = typ
	}
	return nil
}

func checkNamespaceCollisions(tl tlast.TL) error {
	namespaces := map[string]struct {
		s  string
		pr tlast.PositionRange
	}{}
	for _, typ := range tl {
		ns := strings.ToLower(typ.Construct.Name.Namespace)
		if col, ok := namespaces[ns]; ok && col.s != typ.Construct.Name.Namespace {
			e1 := typ.Construct.NamePR.BeautifulError(fmt.Errorf("namespaces must not differ by only case"))
			e2 := col.pr.BeautifulError(errSeeHere)
			return tlast.BeautifulError2(e1, e2)
		}
		namespaces[ns] = struct {
			s  string
			pr tlast.PositionRange
		}{typ.Construct.Name.Namespace, typ.Construct.NamePR}
		ns = strings.ToLower(typ.TypeDecl.Name.Namespace)
		if col, ok := namespaces[ns]; ok && col.s != typ.TypeDecl.Name.Namespace {
			e1 := typ.TypeDecl.NamePR.BeautifulError(fmt.Errorf("namespaces must not differ by only case"))
			e2 := col.pr.BeautifulError(errSeeHere)
			return tlast.BeautifulError2(e1, e2)
		}
		namespaces[ns] = struct {
			s  string
			pr tlast.PositionRange
		}{typ.TypeDecl.Name.Namespace, typ.TypeDecl.NamePR}
	}
	return nil
}

func (gen *Gen2) buildMapDescriptors(tl tlast.TL) error {
	// Collect unions, check that functions cannot form a union with each other or with normal singleConstructors
	for _, typ := range tl {
		for _, f := range typ.Fields {
			if f.FieldName == "" && (len(typ.Fields) != 1 || f.Mask != nil) {
				return f.PR.BeautifulError(fmt.Errorf("anonymous fields are discouraged, except when used in '# a:[int]' pattern or when type has single anonymous field without fieldmask (typedef-like)"))
			}
		}
		conName := typ.Construct.Name.String()
		if col, ok := gen.allConstructors[conName]; ok {
			// typeA = TypeA;
			// typeA = TypeB;
			e1 := typ.Construct.NamePR.BeautifulError(fmt.Errorf("constructor name %q is used again here", conName))
			e2 := col.Construct.NamePR.BeautifulError(errSeeHere)
			return tlast.BeautifulError2(e1, e2)
		}
		gen.allConstructors[conName] = typ
		if !typ.IsFunction {
			typeName := typ.TypeDecl.Name.String()
			if len(typ.TemplateArguments) > len(typ.TypeDecl.Arguments) {
				// rightLeftArgs {X:Type} {Y:#} = RightLeftArgs X; <- bad
				arg := typ.TemplateArguments[len(typ.TypeDecl.Arguments)]
				return typ.TypeDecl.PR.CollapseToEnd().BeautifulError(fmt.Errorf("type declaration %q is missing template argument %q here", typeName, arg.FieldName))
			}
			if len(typ.TemplateArguments) < len(typ.TypeDecl.Arguments) {
				// rightLeftArgs {X:Type} {Y:#} = RightLeftArgs X Y Y; <- bad
				arg := typ.TypeDecl.Arguments[len(typ.TemplateArguments)]
				pr := typ.TypeDecl.ArgumentsPR[len(typ.TemplateArguments)]
				return pr.BeautifulError(fmt.Errorf("type declaration %q has excess template argument %q here", typeName, arg))
			}
			for j, t := range typ.TemplateArguments {
				if t.FieldName != typ.TypeDecl.Arguments[j] {
					// rightLeftArgs {X:Type} {Y:#} = RightLeftArgs Y X;   <- bad
					pr := typ.TypeDecl.ArgumentsPR[j]
					typArg := typ.TypeDecl.Arguments[j]
					e1 := pr.BeautifulError(fmt.Errorf("type declaration %q has wrong template argument name %q here", typeName, typArg))
					e2 := t.PR.BeautifulError(errSeeHere)
					return tlast.BeautifulError2(e1, e2)
				}
			}
			if typeName != "_" { // prohibit boxed type
				gen.typeDescriptors[typeName] = append(gen.typeDescriptors[typeName], typ)
			}
		} else {
			for _, t := range typ.TemplateArguments {
				if t.IsNat {
					// @read funWithArg {fields_mask: #} => True;
					return t.PR.BeautifulError(fmt.Errorf("function declaration %q cannot have template arguments", conName))
				}
				// TODO - sort out things with rpc wrapping later which has a form
				// @readwrite tree_stats.preferMaster {X:Type} query:!X = X;
			}
			if len(typ.Modifiers) == 0 && utils.DoLint(typ.CommentRight) {
				e1 := typ.Construct.NamePR.CollapseToBegin().BeautifulError(fmt.Errorf("function constructor %q without modifier (identifier starting with '@') not recommended", typ.Construct.Name.String()))
				if gen.options.Kernel.WarningsAreErrors {
					return e1
				}
				e1.PrintWarning(gen.options.ErrorWriter, nil)
			}
		}
	}
	for _, typ := range gen.typeDescriptors {
		if len(typ) == 1 { // here there is no functions
			cName := typ[0].Construct.Name
			tName := typ[0].TypeDecl.Name
			typePrefix := strings.ToLower(ToLowerFirst(tName.Name))

			if cName.Namespace != tName.Namespace {
				e1 := typ[0].Construct.NamePR.BeautifulError(fmt.Errorf("simple type constructor namespace should exactly match type namespace"))
				e2 := typ[0].TypeDecl.NamePR.BeautifulError(errSeeHere)
				return tlast.BeautifulError2(e1, e2)
			}
			// We temporarily allow relaxed case match. To use strict match, remove strings.ToLower() calls below
			if EnableWarningsSimpleTypeName && strings.ToLower(cName.Name) != typePrefix &&
				!purelegacy.EnableWarningsSimpleTypeNameSkip(cName.String()) && utils.DoLint(typ[0].CommentRight) {
				e1 := typ[0].Construct.NamePR.BeautifulError(fmt.Errorf("simple type constructor name should differ from type name by case only"))
				e2 := typ[0].TypeDecl.NamePR.BeautifulError(errSeeHere)
				if gen.options.Kernel.WarningsAreErrors {
					return tlast.BeautifulError2(e1, e2)
				}
				tlast.BeautifulError2(e1, e2).PrintWarning(gen.options.ErrorWriter, nil)
			}
			conName := cName.String()
			if col, ok := gen.singleConstructors[conName]; ok {
				e1 := typ[0].Construct.NamePR.BeautifulError(fmt.Errorf("constructor name %q is used again here", conName))
				e2 := col.Construct.NamePR.BeautifulError(errSeeHere)
				// TODO: not sure that this code ever runs, as we have same check on gen.allConstructors
				//       in Gen2.buildMapDescriptors method
				//       use panic to catch such case, remove it later
				err := tlast.BeautifulError2(e1, e2)
				var s strings.Builder
				err.ConsolePrint(&s, err, false)
				panic(s.String())
				// return tlast.BeautifulError2(e1, e2)
			}
			gen.singleConstructors[conName] = typ[0]
			continue
		}
		if err := checkUnionElementsCompatibility(typ, gen.options); err != nil {
			return err
		}
	}
	return nil
}

func checkUnionElementsCompatibility(types []*tlast.Combinator, options *puregen.Options) error {
	// We temporarily allow relaxed case match. To use strict match, remove strings.ToLower() calls below
	typePrefix := strings.ToLower(ToLowerFirst(types[0].TypeDecl.Name.Name))
	typeSuffix := strings.ToLower(types[0].TypeDecl.Name.Name)
	for _, typ := range types {
		conName := strings.ToLower(typ.Construct.Name.Name)
		if EnableWarningsUnionNamespace && typ.Construct.Name.Namespace != typ.TypeDecl.Name.Namespace &&
			!purelegacy.EnableWarningsUnionNamespaceSkip(typ.Construct.Name.Namespace, typ.TypeDecl.Name.Namespace) &&
			utils.DoLint(typ.CommentRight) {
			e1 := typ.Construct.NamePR.BeautifulError(fmt.Errorf("union constructor namespace %q should match type namespace %q", typ.Construct.Name.Namespace, typ.TypeDecl.Name.Namespace))
			e2 := typ.TypeDecl.NamePR.BeautifulError(errSeeHere)
			if options.Kernel.WarningsAreErrors {
				return tlast.BeautifulError2(e1, e2)
			}
			tlast.BeautifulError2(e1, e2).PrintWarning(options.ErrorWriter, nil)
		}
		if EnableWarningsUnionNamePrefix &&
			!strings.HasPrefix(conName, typePrefix) &&
			!strings.HasSuffix(conName, typeSuffix) &&
			!purelegacy.EnableWarningsUnionNamePrefixSkip(typ.Construct.Name.Name, typePrefix, typeSuffix) &&
			utils.DoLint(typ.CommentRight) { // same check as in generateType
			e1 := typ.Construct.NamePR.BeautifulError(fmt.Errorf("union constructor should have type name prefix or suffix %q", typePrefix))
			e2 := typ.TypeDecl.NamePR.BeautifulError(errSeeHere)
			if options.Kernel.WarningsAreErrors {
				return tlast.BeautifulError2(e1, e2)
			}
			tlast.BeautifulError2(e1, e2).PrintWarning(options.ErrorWriter, nil)
			continue
		}
		if EnableWarningsUnionNameExact && conName == typePrefix &&
			!purelegacy.EnableWarningsUnionNameExactSkip(typ.Construct.Name.String()) &&
			utils.DoLint(typ.CommentRight) {
			e1 := typ.Construct.NamePR.BeautifulError(fmt.Errorf("union constructor name should not exactly match type name %q", typePrefix))
			e2 := typ.TypeDecl.PR.BeautifulError(errSeeHere)
			if options.Kernel.WarningsAreErrors {
				return tlast.BeautifulError2(e1, e2)
			}
			tlast.BeautifulError2(e1, e2).PrintWarning(options.ErrorWriter, nil)
		}
	}
	base := types[0]
	for _, typ := range types[1:] {
		cur := typ.Construct.Name.String()
		if len(typ.TemplateArguments) < len(base.TemplateArguments) {
			baseArg := base.TemplateArguments[len(typ.TemplateArguments)]
			// unionArgs2 {A:Type} {B:#} a:A = UnionArgs A B;
			// unionArgs1 {X:Type} a:X = UnionArgs X;
			e1 := typ.TemplateArgumentsPR.CollapseToEnd().BeautifulError(fmt.Errorf("union constructor %q has missing argument %q here", cur, baseArg.FieldName))
			e2 := baseArg.PR.BeautifulError(errSeeHere)
			return tlast.BeautifulError2(e1, e2)
		}
		if len(typ.TemplateArguments) > len(base.TemplateArguments) {
			typArg := typ.TemplateArguments[len(base.TemplateArguments)]
			// unionArgs1 {X:Type} a:X = UnionArgs X;
			// unionArgs2 {A:Type} {B:#} a:A = UnionArgs A B;
			e1 := typArg.PR.BeautifulError(fmt.Errorf("union constructor %q has excess argument %q here", cur, typArg.FieldName))
			e2 := base.TemplateArgumentsPR.CollapseToEnd().BeautifulError(errSeeHere)
			return tlast.BeautifulError2(e1, e2)
		}
		for i, typArg := range typ.TemplateArguments {
			baseArg := base.TemplateArguments[i]
			// unionArgs1 {X:Type} {Y:#} a:X = UnionArgs X Y;
			// unionArgs2 {A:Type} {B:Type} a:A = UnionArgs A B;
			// We cannot support this, because resolveType replaces parameter names into names of first union field
			if baseArg.IsNat != typArg.IsNat || baseArg.FieldName != typArg.FieldName {
				e1 := typArg.PR.BeautifulError(fmt.Errorf("union constructor %q has different argument name or type here %q", cur, typArg.FieldName))
				e2 := baseArg.PR.BeautifulError(errSeeHere)
				return tlast.BeautifulError2(e1, e2)
			}
		}
	}
	return nil
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
				Debugf(cmp.Diff(string(was), code))
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

func (gen *Gen2) cppFilterFile(file string, filters []string) bool {
	if strings.HasSuffix(file, ".o") {
		return true
	}
	// for future?
	//if !gen.options.DeleteUnrelatedFiles {
	//	cleanFile := filepath.Clean(file)
	//	folders := make([]string, 0)
	//	for cleanFile != "." {
	//		cleanFile = filepath.Dir(cleanFile)
	//		folders = append(folders, cleanFile)
	//	}
	//	if len(folders) < 2 {
	//		return true
	//	}
	//	folder := folders[len(folders)-2]
	//	if folder == CommonGroup {
	//		folder = ""
	//	}
	//	return !inNameFilter(tlast.Name{Namespace: folder}, filters)
	//}

	// TODO change for future development
	return false
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
	if err := kernel.Compile(); err != nil {
		return err
	}

	gen, err := generateCode(kernel.TL1(), options)
	if gen == nil {
		return err
	}
	if err = gen.WriteToDir(options.Outdir); err != nil {
		return err // Context is already in err
	}
	return err
}

func generateCode(tl tlast.TL, options *puregen.Options) (*Gen2, error) {
	if options.Kernel.Verbose {
		DEBUG = true
		Debugf(">>> ENABLE DEBUG MODE <<<\n")
	}
	gen := &Gen2{
		options:    options,
		Code:       map[string]string{},
		Namespaces: map[string]*Namespace{},
		// Files:                 map[string][]*TypeRWWrapper{},
		typeDescriptors:    map[string][]*tlast.Combinator{},
		singleConstructors: map[string]*tlast.Combinator{},
		allConstructors:    map[string]*tlast.Combinator{},
		builtinTypes:       map[string]*TypeRWWrapper{},
		generatedTypes:     map[string]*TypeRWWrapper{},
	}
	switch options.Language {
	case "php":
	default:
		return nil, fmt.Errorf("unsupported language %q, only 'go' and 'cpp' are supported, plus '' for linting", options.Language)
	}
	typesWhiteList := prepareNameFilter(options.Kernel.TypesWhiteList)
	bytesWhiteList := prepareNameFilter(options.BytesWhiteList)
	tl2WhiteList := prepareNameFilter(options.Kernel.TL2WhiteList)
	gen.supportedAnnotations = map[string]struct{}{"read": {}, "any": {}, "internal": {}, "write": {}, "readwrite": {}, "kphp": {}}
	rootNamespace := gen.getNamespace("")
	primitiveTypesList := []*TypeRWPrimitive{
		{
			gen:               gen,
			tlType:            "#",
			goType:            "uint32",
			cppPrimitiveType:  "uint32_t",
			cppDefaultInit:    " = 0",
			cppFunctionSuffix: "nat",
			cppResetValue:     "%s = 0;",
			writeJSONValue:    "basictl.JSONWriteUint32",
			resetValue:        "%s = 0",
			randomValue:       "basictl.RandomUint",
			writeValue:        "basictl.NatWrite",
			readValue:         "basictl.NatRead",
		}, {
			gen:               gen,
			tlType:            "int",
			goType:            "int32",
			cppPrimitiveType:  "int32_t",
			cppDefaultInit:    " = 0",
			cppFunctionSuffix: "int",
			cppResetValue:     "%s = 0;",
			writeJSONValue:    "basictl.JSONWriteInt32",
			resetValue:        "%s = 0",
			randomValue:       "basictl.RandomInt",
			writeValue:        "basictl.IntWrite",
			readValue:         "basictl.IntRead",
		}, {
			gen:               gen,
			tlType:            "long",
			goType:            "int64",
			cppPrimitiveType:  "int64_t",
			cppDefaultInit:    " = 0",
			cppFunctionSuffix: "long",
			cppResetValue:     "%s = 0;",
			writeJSONValue:    "basictl.JSONWriteInt64",
			resetValue:        "%s = 0",
			randomValue:       "basictl.RandomLong",
			writeValue:        "basictl.LongWrite",
			readValue:         "basictl.LongRead",
		}, {
			gen:               gen,
			tlType:            "float",
			goType:            "float32",
			cppPrimitiveType:  "float",
			cppDefaultInit:    " = 0",
			cppFunctionSuffix: "float",
			cppResetValue:     "%s = 0;",
			writeJSONValue:    "basictl.JSONWriteFloat32",
			resetValue:        "%s = 0",
			randomValue:       "basictl.RandomFloat",
			writeValue:        "basictl.FloatWrite",
			readValue:         "basictl.FloatRead",
		}, {
			gen:               gen,
			tlType:            "double",
			goType:            "float64",
			cppPrimitiveType:  "double",
			cppDefaultInit:    " = 0",
			cppFunctionSuffix: "double",
			cppResetValue:     "%s = 0;",
			writeJSONValue:    "basictl.JSONWriteFloat64",
			resetValue:        "%s = 0",
			randomValue:       "basictl.RandomDouble",
			writeValue:        "basictl.DoubleWrite",
			readValue:         "basictl.DoubleRead",
		}, {
			gen:               gen,
			tlType:            "string",
			goType:            "string",
			cppPrimitiveType:  "std::string",
			cppDefaultInit:    "",
			cppFunctionSuffix: "string",
			cppResetValue:     "%s.clear();",
			writeJSONValue:    "basictl.JSONWriteString",
			resetValue:        "%s = \"\"",
			randomValue:       "basictl.RandomString",
			writeValue:        "basictl.StringWrite",
			readValue:         "basictl.StringRead",
		},
	}
	if gen.options.GenerateTL2() {
		primitiveTypesList = append(primitiveTypesList, &TypeRWPrimitive{
			gen:               gen,
			tlType:            "uint32",
			goType:            "uint32",
			cppPrimitiveType:  "uint32_t",
			cppDefaultInit:    " = 0",
			cppFunctionSuffix: "nat",
			cppResetValue:     "%s = 0;",
			writeJSONValue:    "basictl.JSONWriteUint32",
			resetValue:        "%s = 0",
			randomValue:       "basictl.RandomUint",
			writeValue:        "basictl.NatWrite",
			readValue:         "basictl.NatRead",
		})
	}
	builtinBeautifulText := fmt.Sprintf(`
%s {n:#} {t:Type} n*[t] = _ n t; // builtin tuple
%s {t:Type} # [t] = _ t; // builtin vector
`, BuiltinTupleName, BuiltinVectorName)
	primitiveTypes := map[string]*TypeRWPrimitive{}
	for _, cn := range primitiveTypesList {
		builtinBeautifulText += fmt.Sprintf("%s ? = _; // builtin primitive type\n", cn.tlType)
		primitiveTypes[cn.tlType] = cn
	}

	btl, err := tlast.ParseTLFile(builtinBeautifulText, "<builtin>", tlast.LexerOptions{
		AllowBuiltin: true,
		AllowDirty:   false,
	}) // We need references to token positions for beautification, so we decided to parse as a TL file
	if err != nil {
		return nil, fmt.Errorf("failed to parse internal builtin type representation for beautification: %w", err)
	}

	if gen.options.Language == "php" {
		// RPC SPECIAL CHANGES
		if gen.options.PHP.AddRPCTypes {
			const rpcRequestResultName = "ReqResult"
			rpcResultsMapping := map[string]string{
				"reqError":        "rpcResponseError",
				"reqResultHeader": "rpcResponseHeader",
				"_":               "rpcResponseOk",
			}
			rpcRemovedTypes := map[string]bool{
				"rpcReqResult": true,
				"rpcReqError":  true,
				"rpcInvokeReq": true,
			}
			rpcFunctionTypeRef := tlast.TypeRef{
				Type: tlast.Name{
					Name: PHPRPCFunctionMock,
				},
			}
			rpcFunctionResultTypeRef := tlast.TypeRef{
				Type: tlast.Name{
					Name: PHPRPCFunctionResultMock,
				},
			}
			rpcResponseTypeRef := tlast.TypeRef{
				Type: tlast.Name{
					Name: PHPRPCResponseMock,
				},
			}
			//hasRPCCombinators := false
			//for _, typ := range tl {
			//	if !typ.IsFunction &&
			//		rpcResultsMapping[typ.Construct.Name.String()] != "" &&
			//		typ.TypeDecl.Name.String() == rpcRequestResultName {
			//		hasRPCCombinators = true
			//	}
			//}
			//// TODO: RETURN ORIGINAL COMBINATOR
			//if hasRPCCombinators {
			//	tl = append(tl, &tlast.Combinator{
			//		CommentRight: "// tlgen:nolint",
			//		TypeDecl: tlast.TypeDeclaration{
			//			Name: tlast.Name{
			//				Name: "ReqResult",
			//			},
			//			Arguments: []string{"X"},
			//		},
			//		Construct: tlast.Constructor{
			//			Name: tlast.Name{Name: "_"},
			//		},
			//		TemplateArguments: []tlast.TemplateArgument{
			//			{
			//				FieldName: "X",
			//				IsNat:     false,
			//			},
			//		},
			//		Fields: []tlast.Field{
			//			{
			//				FieldName: "result",
			//				FieldType: tlast.TypeRef{
			//					Type: tlast.Name{
			//						Name: "X",
			//					},
			//				},
			//			},
			//		},
			//	})
			//	// TODO make normal version for this
			//	tl[len(tl)-1].Construct.ID = tl[len(tl)-1].GenCrc32()
			//}
			tl = append(tl, &tlast.Combinator{
				CommentRight: "// tlgen:nolint",
				TypeDecl: tlast.TypeDeclaration{
					Name: rpcFunctionTypeRef.Type,
				},
				Construct: tlast.Constructor{Name: rpcFunctionTypeRef.Type},
			})
			// TODO make normal version for this
			tl[len(tl)-1].Construct.ID = tl[len(tl)-1].GenCrc32()
			tl = append(tl, &tlast.Combinator{
				CommentRight: "// tlgen:nolint",
				TypeDecl: tlast.TypeDeclaration{
					Name: rpcFunctionResultTypeRef.Type,
				},
				Construct: tlast.Constructor{Name: rpcFunctionResultTypeRef.Type},
			})
			// TODO make normal version for this
			tl[len(tl)-1].Construct.ID = tl[len(tl)-1].GenCrc32()
			for _, typ := range tl {
				if typ.IsFunction && len(typ.TemplateArguments) == 1 {
					// copy original
					copyOriginal := *typ
					copyOriginal.Fields = make([]tlast.Field, len(copyOriginal.Fields))
					for i := range typ.Fields {
						copyOriginal.Fields[i] = typ.Fields[i]
						copyOriginal.Fields[i].FieldType = typ.Fields[i].FieldType.DeepCopy()
					}
					copyOriginal.FuncDecl = typ.FuncDecl.DeepCopy()
					// modify for generation
					phpRemoveTemplateFromGeneric(typ, &rpcFunctionTypeRef, &rpcFunctionResultTypeRef)
					typ.OriginalDescriptor = &copyOriginal
				} else if !typ.IsFunction &&
					rpcResultsMapping[typ.Construct.Name.String()] != "" &&
					typ.TypeDecl.Name.String() == rpcRequestResultName {
					typ.Construct.ID = typ.GenCrc32()

					typ.Crc32()
					typ.Construct.Name.Name = rpcResultsMapping[typ.Construct.Name.String()]
					// copy original
					copyOriginal := *typ
					copyOriginal.Fields = make([]tlast.Field, len(copyOriginal.Fields))
					for i := range typ.Fields {
						copyOriginal.Fields[i] = typ.Fields[i]
						copyOriginal.Fields[i].FieldType = typ.Fields[i].FieldType.DeepCopy()
					}
					copyOriginal.TypeDecl = typ.TypeDecl
					copy(copyOriginal.TypeDecl.Arguments, typ.TypeDecl.Arguments)
					// modify for generation
					typ.TypeDecl = tlast.TypeDeclaration{Name: rpcResponseTypeRef.Type}
					phpRemoveTemplateFromGeneric(typ, &rpcFunctionResultTypeRef, &rpcFunctionResultTypeRef)

					typ.OriginalDescriptor = &copyOriginal
				}
			}
			// TODO DELETE AS NORMAL PEOPLE
			var removedTypesIndecies []int
			for i, typ := range tl {
				if rpcRemovedTypes[typ.Construct.Name.String()] {
					removedTypesIndecies = append(removedTypesIndecies, i)
				}
			}
			sort.Ints(removedTypesIndecies)
			for i, index := range removedTypesIndecies {
				tl = append(tl[:index-i], tl[index-i+1:]...)
			}
		}
	}

	for i, typ := range tl { // replace built-in
		tName := typ.Construct.Name.String()
		// convert that old syntax to new syntax.
		if !typ.Builtin {
			continue
		}
		if len(typ.TemplateArguments) != 0 {
			return nil, typ.TemplateArgumentsPR.BeautifulError(fmt.Errorf("builtin wrapper %q cannot have template parameters", tName))
		}
		if _, ok := primitiveTypes[tName]; !ok {
			return nil, typ.Construct.NamePR.BeautifulError(fmt.Errorf("builtin wrapper %q must have constructor name equal to some builtin type", tName))
		}
		newDesc := &tlast.Combinator{}
		*newDesc = *tl[i]
		newDesc.OriginalDescriptor = typ
		newDesc.Fields = append(newDesc.Fields, tlast.Field{
			FieldType: tlast.TypeRef{
				Type: tlast.Name{Name: tName},
				Bare: true,
			},
		})
		newDesc.Builtin = false
		tl[i] = newDesc
	}
	if err := checkTagCollisions(tl); err != nil {
		return nil, err
	}
	if err := checkNamespaceCollisions(tl); err != nil {
		return nil, err
	}

	// ReplaceSquareBrackets will generate types with id 0, we will not generate boxed methods for such types
	if tl, err = gen.ReplaceSquareBracketsElem(tl); err != nil {
		return nil, fmt.Errorf("replacing with canonical tuples: %w", err)
	}
	err = gen.buildMapDescriptors(tl)
	if err != nil {
		return nil, err
	}

	// Now we replace all builtin legitimate builtin wrapper constructors to constructors of builtins
	// Int and %Int will reference wrappers, while int will reference builtin constructor.
	// To avoid 2 canonical forms, resolveType will replace %Int to int for wrappers
	for _, bt := range btl {
		bt.Construct.ID = 0
		bt.Construct.IDExplicit = true
		tName := bt.Construct.Name.String()
		if tName == BuiltinTupleName || tName == BuiltinVectorName {
			gen.allConstructors[tName] = bt
			gen.singleConstructors[tName] = bt
			continue
		}
		cn, ok := primitiveTypes[tName]
		if !ok {
			panic("broken primitive types list")
		}
		wrapper := &TypeRWWrapper{
			gen:          gen,
			ns:           rootNamespace,
			trw:          cn,
			goGlobalName: cn.goType,
			goLocalName:  "",
			fileName:     cn.tlType,
			origTL:       []*tlast.Combinator{bt},
		}
		if cn.tlType == "#" {
			wrapper.fileName = "nat"
		}
		gen.builtinTypes[cn.tlType] = wrapper
		gen.generatedTypesList = append(gen.generatedTypesList, wrapper)
		typ, ok := gen.allConstructors[tName]
		if ok {
			// vasya ? = Int;
			// vasya {T:Type} {N:#} ? = Int T N;
			// int {T:Type} {N:#} int = Int T N;
			// int = Int;
			// int int int = Int;
			// int vasya:int = Int;
			// int n.0?int = Int;
			// int (int 5) = Int;
			// int (%Int) = Int;
			// int (Int) = Int;
			// int ? = Int;   <-- allowed, but deprecated shape. TODO - prohibit in TL2
			// int int = Int; <-- allowed shape
			if len(typ.TemplateArguments) != 0 {
				return nil, typ.TemplateArgumentsPR.BeautifulError(fmt.Errorf("builtin wrapper %q cannot have template parameters", tName))
			}
			if len(typ.Fields) == 0 {
				return nil, typ.TemplateArgumentsPR.CollapseToEnd().BeautifulError(fmt.Errorf("builtin wrapper %q must have exactly 1 field", tName))
			}
			if len(typ.Fields) > 1 {
				return nil, typ.Fields[1].FieldType.PR.BeautifulError(fmt.Errorf("builtin wrapper %q has excess field, must have exactly 1", tName))
			}
			if typ.Fields[0].FieldName != "" {
				return nil, typ.Fields[0].PRName.BeautifulError(fmt.Errorf("builtin wrapper %q field must be anonymous", tName))
			}
			if typ.Fields[0].Mask != nil {
				return nil, typ.Fields[0].Mask.PRName.BeautifulError(fmt.Errorf("builtin wrapper %q field must not use field mask", tName))
			}
			if typ.Fields[0].FieldType.Type.String() != tName || len(typ.Fields[0].FieldType.Args) != 0 { // do not check Bare, because int === %int
				return nil, typ.Fields[0].FieldType.PR.BeautifulError(fmt.Errorf("builtin wrapper %q field type must match constructor name", tName))
			}
			if _, ok := gen.builtinTypes[tName]; !ok {
				return nil, typ.Construct.NamePR.BeautifulError(fmt.Errorf("builtin wrapper %q must have constructor name equal to some builtin type", tName))
			}
		} else {
			gen.allConstructors[tName] = bt
			gen.singleConstructors[tName] = bt
		}
	}

	// tupleDesc := btl[0]
	// vectorDesc := btl[1]
	// gen.singleConstructors[tupleDesc.Construct.Name.String()] = tupleDesc
	// gen.allConstructors[tupleDesc.Construct.Name.String()] = tupleDesc
	// gen.singleConstructors[vectorDesc.Construct.Name.String()] = vectorDesc
	// gen.allConstructors[vectorDesc.Construct.Name.String()] = vectorDesc

	{
		allAnnotations := map[string]struct{}{}
		// generated RPC code can depend on those annotations, even
		// if none present in current tl file.
		// so we add all supported annotations always.
		for m := range gen.supportedAnnotations {
			allAnnotations[m] = struct{}{}
			gen.allAnnotations = append(gen.allAnnotations, m)
		}
		for _, typ := range tl {
			for _, m := range typ.Modifiers {
				if strings.ToLower(m.Name) != m.Name { // TODO - move into lexer
					return nil, m.PR.BeautifulError(fmt.Errorf("annotations must be lower case"))
				}
				if _, ok := allAnnotations[m.Name]; !ok {
					if _, ok := gen.supportedAnnotations[m.Name]; !ok && utils.DoLint(typ.CommentRight) {
						e1 := m.PR.BeautifulError(fmt.Errorf("annotation %q not known to tlgen", m.Name))
						if gen.options.Kernel.WarningsAreErrors {
							return nil, e1
						}
						e1.PrintWarning(options.ErrorWriter, nil)
					}
					allAnnotations[m.Name] = struct{}{}
					gen.allAnnotations = append(gen.allAnnotations, m.Name)
				}
			}
		}
		if len(gen.allAnnotations) > 32 {
			return nil, fmt.Errorf("too many (%d) different annotations, max is 32 for now", len(gen.allAnnotations))
		}
		sort.Strings(gen.allAnnotations)
	}
	skippedDueToWhitelist := 0

	for _, typ := range tl {
		if purelegacy.GenerateUnusedNatTemplates(typ.Construct.Name.String()) && len(typ.TemplateArguments) == 1 && typ.TemplateArguments[0].IsNat {
			t := tlast.TypeRef{Type: typ.TypeDecl.Name, PR: typ.TypeDecl.PR}
			argT := tlast.TypeRef{Type: tlast.Name{
				Namespace: "",
				Name:      "ArgumentN",
			}}
			t.Args = append(t.Args, tlast.ArithmeticOrType{
				IsArith: false,
				T:       argT,
			})
			lrc := LocalResolveContext{allowAnyConstructor: true, localNatArgs: map[string]LocalNatArg{}}
			lrc.localNatArgs["ArgumentN"] = LocalNatArg{
				natArg: ActualNatArg{isField: true, FieldIndex: 0},
			}
			_, _, _, _, err = gen.getType(lrc, t, nil)
			if err != nil {
				return nil, err
			}
		}
		shouldGenerate := options.Kernel.TypesWhiteList == ""
		whiteListName := typ.Construct.Name
		if !typ.IsFunction {
			whiteListName = typ.TypeDecl.Name
		}
		if inNameFilter(whiteListName, bytesWhiteList) {
			shouldGenerate = true
		}
		if inNameFilter(whiteListName, typesWhiteList) {
			shouldGenerate = true
		}
		if !shouldGenerate {
			skippedDueToWhitelist++
			continue
		}
		if len(typ.TemplateArguments) == 0 {
			t := tlast.TypeRef{Type: typ.Construct.Name, PR: typ.Construct.NamePR}
			if !typ.IsFunction {
				t = tlast.TypeRef{Type: typ.TypeDecl.Name, PR: typ.TypeDecl.PR}
			}
			_, _, _, _, err = gen.getType(LocalResolveContext{allowAnyConstructor: true}, t, nil)
			if err != nil {
				return nil, err
			}
		}
	}

	purelegacy.PrintGlobalMap()

	typesCounterMarkBytes := 0
	typesCounterMarkTL2 := 0
	for _, v := range gen.generatedTypesList { // we do not need tl2masks in this case
		if str, ok := v.trw.(*TypeRWStruct); ok && !v.wantsTL2 {
			for i := range str.Fields {
				str.Fields[i].MaskTL2Bit = nil
			}
		}
	}
	slices.SortStableFunc(gen.generatedTypesList, func(a, b *TypeRWWrapper) int { //  TODO - better idea?
		return TypeRWWrapperLessGlobal(a, b)
	})
	sortedTypes := gen.generatedTypesList
	// for _, st := range sortedTypes {
	//	fmt.Printf("sorted type %q\n", st.localTypeArg.rt.String())
	// }
	for _, v := range sortedTypes {
		// fmt.Printf("type %s names %s %s %d\n", v.CanonicalStringTop(), v.goGlobalName, v.tlName.String(), v.tlTag)
		// if len(v.origTL) <= 1 {
		//	fmt.Printf("     %s\n", v.CanonicalString(true))
		// } else {
		//	fmt.Printf("     %s\n", v.CanonicalString(false))
		// }
		// r # [r] = S;
		visitedNodes := map[*TypeRWWrapper]bool{}
		v.trw.fillRecursiveUnwrap(visitedNodes)
		v.preventUnwrap = visitedNodes[v]
		if v.preventUnwrap {
			fmt.Printf("prevented unwrap of %v\n", v.tlName)
		}
	}

	// in BeforeCodeGenerationStep we split recursion. Which links will be broken depends on order of nodes visited
	for _, v := range sortedTypes {
		v.trw.BeforeCodeGenerationStep1()
	}

	// we link normal and long types for VK int->long conversion. This code is VK-specific and will be removed after full migration
	for _, v := range sortedTypes {
		// @readwrite queueLong.getQueueKey id:long ip:int timeout:int queue:string = queueLong.TimestampKey;
		// @readwrite queue.getQueueKey id:int ip:int timeout:int queue:string = queue.TimestampKey;
		longName := v.CanonicalStringTop()
		argsStart := strings.Index(longName, "<")
		if argsStart < 0 {
			argsStart = len(longName)
		}
		if i := strings.Index(longName[:argsStart], "."); i >= 0 {
			longName = longName[:i] + "Long" + longName[i:]

			if tt, ok := gen.generatedTypes[longName]; ok {
				// log.Printf("long name %s discovered for %s", longName, v.CanonicalStringTop())
				v.WrLong = tt
				tt.WrWithoutLong = v
			}
		}
	}

	//// additional php check
	//if gen.options.LinterPHPCheck {
	//	storedOption := gen.options.PHP.InplaceSimpleStructs
	//	gen.options.PHP.InplaceSimpleStructs = true
	//
	//	namesToIgnoreMap := make(map[string]bool)
	//	for _, s := range PHPNamesToIgnoreForLinterCheck {
	//		namesToIgnoreMap[strings.ToLower(s)] = true
	//	}
	//
	//	isFlatType := func(t *TypeRWWrapper) bool {
	//		struct_, isStruct := t.trw.(*TypeRWStruct)
	//		return isStruct &&
	//			struct_.PhpCanBeSimplify() &&
	//			strings.EqualFold(t.origTL[0].TypeDecl.Name.String(), t.origTL[0].Construct.Name.String())
	//	}
	//
	//	isNotPolymorphicType := func(t *TypeRWWrapper) bool {
	//		_, isStruct := t.trw.(*TypeRWStruct)
	//		return isStruct &&
	//			t.unionParent == nil &&
	//			len(t.origTL) == 1 &&
	//			strings.EqualFold(t.origTL[0].TypeDecl.Name.String(), t.origTL[0].Construct.Name.String())
	//	}
	//
	//	// tmp
	//	namespacesToIgnore := []string{"tls"}
	//	namespacesToIgnoreMap := utils.SliceToSet(namespacesToIgnore)
	//
	//	// issue 1
	//	for _, wrapper := range sortedTypes {
	//		if strct_, ok := wrapper.trw.(*TypeRWStruct); ok {
	//			for _, field := range strct_.Fields {
	//				if field.t.originateFromTL2 {
	//					continue
	//				}
	//				if PHPSpecialMembersTypes(field.t) != "" {
	//					continue
	//				}
	//				if !field.bare && isFlatType(field.t) && !namesToIgnoreMap[strings.ToLower(field.t.tlName.String())] {
	//					linterErr := field.origTL.FieldType.PR.BeautifulError(fmt.Errorf("can't have boxed reference in field to flat type due to php generator issues (instance: %s)", wrapper.CanonicalStringTop()))
	//					if gen.options.Kernel.WarningsAreErrors && !namespacesToIgnoreMap[field.t.tlName.Namespace] {
	//						return nil, linterErr
	//					}
	//					linterErr.PrintWarning(gen.options.ErrorWriter, nil)
	//				}
	//			}
	//		}
	//	}
	//
	//	// issue 2
	//	for _, wrapper := range sortedTypes {
	//		if isFlatType(wrapper) && !namesToIgnoreMap[strings.ToLower(wrapper.tlName.String())] {
	//			for _, argument := range wrapper.origTL[0].TemplateArguments {
	//				if !argument.IsNat {
	//					linterErr := argument.PR.BeautifulError(fmt.Errorf("flat types can't have type templates due to php generator issues"))
	//					if gen.options.Kernel.WarningsAreErrors {
	//						return nil, linterErr
	//					}
	//					linterErr.PrintWarning(gen.options.ErrorWriter, nil)
	//				}
	//			}
	//		}
	//	}
	//
	//	if gen.options.LinterPHPNonPolymorphicBoxedRef {
	//		// issue 3
	//		for _, wrapper := range sortedTypes {
	//			if strct_, ok := wrapper.trw.(*TypeRWStruct); ok {
	//				for _, field := range strct_.Fields {
	//					if PHPSpecialMembersTypes(field.t) != "" {
	//						continue
	//					}
	//					if !field.bare && isNotPolymorphicType(field.t) && !namesToIgnoreMap[strings.ToLower(field.t.tlName.String())] {
	//						linterErr := field.origTL.FieldType.PR.BeautifulError(fmt.Errorf("can't boxed reference type with a single constructor with the same name in field due to php generator issues"))
	//						if gen.options.Kernel.WarningsAreErrors {
	//							return nil, linterErr
	//						}
	//						linterErr.PrintWarning(gen.options.ErrorWriter, nil)
	//					}
	//				}
	//			}
	//		}
	//	}
	//
	//	gen.options.PHP.InplaceSimpleStructs = storedOption
	//}

	// detect recursion loops first
	if options.Kernel.Verbose {
		if skippedDueToWhitelist != 0 {
			log.Printf("skipped %d object roots by the whitelist filter: %s", skippedDueToWhitelist, strings.Join(typesWhiteList, ", "))
		}
		if filter := strings.Join(bytesWhiteList, ", "); filter != "" {
			log.Printf("found %d object roots for byte-optimized versions of types by the following filter: %s", typesCounterMarkBytes, filter)
		}
		if filter := strings.Join(tl2WhiteList, ", "); filter != "" {
			log.Printf("found %d object roots for TL2 versions of types by the following filter: %s", typesCounterMarkTL2, filter)
		}
	}
	if gen.options.CopyrightFilePath != "" {
		buf, err := os.ReadFile(gen.options.CopyrightFilePath)
		if err != nil {
			return nil, fmt.Errorf("failed to open copyright text file: %w", err)
		}
		gen.copyrightText = string(buf)
	}
	switch options.Language {
	case "go":
		return nil, fmt.Errorf("go generator is moved into tl2gen")
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
		if err := gen.generateCodePHP(bytesWhiteList); err != nil {
			return nil, err
		}

	}

	return gen, nil
}

func phpRemoveTemplateFromGeneric(combinator *tlast.Combinator, newTypeRef, newTypeResultRef *tlast.TypeRef) {
	template := combinator.TemplateArguments[0].FieldName
	combinator.TemplateArguments = nil
	for i := range combinator.Fields {
		phpRemoveTemplateFromTypeDecl(&combinator.Fields[i].FieldType, template, newTypeRef)
	}
	phpRemoveTemplateFromTypeDecl(&combinator.FuncDecl, template, newTypeResultRef)
}

func phpRemoveTemplateFromTypeDecl(declaration *tlast.TypeRef, template string, newTypeRef *tlast.TypeRef) {
	if declaration.Type.String() == template {
		*declaration = *newTypeRef
	} else {
		for i := range declaration.Args {
			if !declaration.Args[i].IsArith {
				phpRemoveTemplateFromTypeDecl(&declaration.Args[i].T, template, newTypeRef)
			}
		}
	}
}

var TypeComparator = func(a, b *TypeRWWrapper) int {
	return strings.Compare(a.goGlobalName, b.goGlobalName)
}

type LanguageTypeSystemInfo struct {
	AllowedTypeGenerics                bool
	AllowedCompileTimeConstantGenerics bool
}

func (lti LanguageTypeSystemInfo) OnlyFullyResolvedTypes() bool {
	return !(lti.AllowedCompileTimeConstantGenerics && lti.AllowedTypeGenerics)
}

type ConstructorName = tlast.Name
type Constructor struct {
	Type   *TypeDefinition
	Id     uint
	Name   ConstructorName
	Fields []tlast.Field
	Result *tlast.TypeRef
}

type TypeName = tlast.Name
type TypeDefinition struct {
	IsBasic       bool
	Name          TypeName
	TypeArguments []tlast.TemplateArgument
	Constructors  []*Constructor
}

type EvaluatedTypeVariant = int

const (
	NumberConstant EvaluatedTypeVariant = 0
	NumberVariable EvaluatedTypeVariant = 1
	TypeConstant   EvaluatedTypeVariant = 2
	TypeVariable   EvaluatedTypeVariant = 3
)

type EvaluatedType struct {
	Index EvaluatedTypeVariant

	// union variants
	Constant     uint32         // 0
	Variable     string         // 1
	Type         *TypeReduction // 2
	TypeVariable string         // 3

	VariableActsAsConstant bool // only if Index == 1 and only for type declarations
}

type TypeReduction struct {
	IsType bool

	Type        *TypeDefinition
	Constructor *Constructor

	Arguments []EvaluatedType
}

func (tr *TypeReduction) EraseTypes() TypeReduction {
	erasedReduction := *tr
	erasedReduction.Arguments = make([]EvaluatedType, 0)

	var typeIndex int

	for _, arg := range tr.Arguments {
		switch arg.Index {
		case TypeConstant:
			erasedReduction.Arguments = append(
				erasedReduction.Arguments,
				EvaluatedType{Index: TypeVariable, TypeVariable: fmt.Sprintf("X%d", typeIndex)},
			)
			typeIndex++
		default:
			erasedReduction.Arguments = append(erasedReduction.Arguments, arg)
		}
	}
	return erasedReduction
}

func (tr *TypeReduction) EraseConstants() TypeReduction {
	erasedReduction := *tr
	erasedReduction.Arguments = make([]EvaluatedType, 0)

	var conIndex int

	for _, arg := range tr.Arguments {
		switch arg.Index {
		case NumberConstant:
			erasedReduction.Arguments = append(
				erasedReduction.Arguments,
				EvaluatedType{Index: NumberVariable, Variable: fmt.Sprintf("n%d", conIndex), VariableActsAsConstant: true},
			)
			conIndex++
		default:
			erasedReduction.Arguments = append(erasedReduction.Arguments, arg)
		}
	}
	return erasedReduction
}

func (tr *TypeReduction) EraseConstantsAndTypes() TypeReduction {
	erasedReduction := (*tr).EraseTypes()
	erasedReduction = erasedReduction.EraseConstants()
	return erasedReduction
}

func (tr *TypeReduction) AncestorTemplateInTypeSystem(info LanguageTypeSystemInfo) TypeReduction {
	ancestor := *tr
	if info.AllowedTypeGenerics {
		ancestor = ancestor.EraseTypes()
	}
	if info.AllowedCompileTimeConstantGenerics {
		ancestor = ancestor.EraseConstants()
	}
	return ancestor
}

func (tr *TypeReduction) IsFullyResolved() bool {
	for _, arg := range tr.Arguments {
		switch arg.Index {
		case NumberConstant:
			continue
		case NumberVariable:
			continue
		case TypeConstant:
			if !arg.Type.IsFullyResolved() {
				return false
			}
		case TypeVariable:
			return false
		default:
			panic("not all variants of reduction considered")
		}
	}
	return true
}

func (tr *TypeReduction) ReferenceName() (name tlast.Name) {
	if tr.IsType {
		name = tr.Type.Name
	} else {
		name = tr.Constructor.Name
	}
	return
}

func (tr *TypeReduction) ReferenceType() *TypeDefinition {
	if tr.IsType {
		return tr.Type
	}
	return tr.Constructor.Type
}

func (tr *TypeReduction) String() string {
	s := ""
	if tr.Type != nil || tr.Constructor != nil {
		s += tr.ReferenceName().String()
		if len(tr.Arguments) != 0 {
			s += "<"
			for i, arg := range tr.Arguments {
				switch arg.Index {
				case 0:
					s += "Con" + strconv.FormatInt(int64(arg.Constant), 10)
				case 1:
					if arg.VariableActsAsConstant {
						s += "ConstVar"
					} else {
						s += "Var"
					}
				case 2:
					s += arg.Type.String()
				case 3:
					s += "TypeVar"
				}
				if len(tr.Arguments) != i+1 {
					s += ","
				}
			}
			s += ">"
		}
	}
	return s
}

func processCombinators(types map[string]*tlast.Combinator) *TypesInfo {
	existingTypes := make(map[TypeName]*TypeDefinition)
	existingConstructors := make(map[ConstructorName]*Constructor)

	for _, comb := range types {
		declaredType := comb.TypeDecl.Name
		if comb.Builtin || comb.IsFunction {
			declaredType = comb.Construct.Name
		}
		currentConstructor := comb.Construct.Name
		if _, presented := existingTypes[declaredType]; !presented {
			existingTypes[declaredType] = &TypeDefinition{
				IsBasic:       comb.Builtin,
				Name:          declaredType,
				Constructors:  []*Constructor{},
				TypeArguments: comb.TemplateArguments,
			}
		}
		targetType := existingTypes[declaredType]
		constructor := Constructor{
			Name:   currentConstructor,
			Fields: comb.Fields,
			Id:     uint(len(targetType.Constructors)),
			Type:   targetType,
			Result: &comb.FuncDecl,
		}
		if !comb.IsFunction {
			constructor.Result = nil
		}
		targetType.Constructors = append(targetType.Constructors, &constructor)
		existingConstructors[currentConstructor] = &constructor
	}

	typeReductions := make(map[string]*TypeReduction)
	visitedTypes := make(map[TypeName]bool)

	for _, comb := range existingTypes {
		if len(comb.TypeArguments) != 0 {
			continue
		}
		reduce(
			TypeReduction{
				IsType: true,
				Type:   existingTypes[comb.Name],
			},
			&typeReductions,
			&existingTypes,
			&existingConstructors,
			&visitedTypes,
		)
	}

	ti := TypesInfo{
		Types:        existingTypes,
		Constructors: existingConstructors,
		//TypeReductions: typeReductions,
	}

	//printResults(ti)

	return &ti
}

func reduce(
	targetTypeReduction TypeReduction,
	visitedReductions *map[string]*TypeReduction,
	types *map[TypeName]*TypeDefinition,
	constructors *map[ConstructorName]*Constructor,
	visitedTypes *map[TypeName]bool,
) {
	for _, arg := range targetTypeReduction.Arguments {
		if arg.Index == TypeConstant {
			reduce(*arg.Type, visitedReductions, types, constructors, visitedTypes)
		}
	}

	var visitingConstructors []*Constructor

	if targetTypeReduction.IsType {
		visitingConstructors = targetTypeReduction.Type.Constructors
	} else {
		visitingConstructors = append(visitingConstructors, targetTypeReduction.Constructor)
	}

	for _, constr := range visitingConstructors {
		reduceConstructor(
			constr,
			targetTypeReduction.Arguments,
			visitedReductions,
			types,
			constructors,
			visitedTypes,
		)
	}
}

func reduceConstructor(
	constructor *Constructor,
	args []EvaluatedType,
	visitedReductions *map[string]*TypeReduction,
	types *map[TypeName]*TypeDefinition,
	constructors *map[ConstructorName]*Constructor,
	visitedTypes *map[TypeName]bool,
) {
	if constructor == nil {
		return
	}

	if (*visitedTypes)[constructor.Name] {
		return
	} else {
		(*visitedTypes)[constructor.Name] = true
		defer func() {
			(*visitedTypes)[constructor.Name] = false
		}()
	}

	constructorReduction := TypeReduction{IsType: false, Constructor: constructor, Arguments: args}

	if _, ok := (*visitedReductions)[constructorReduction.String()]; !ok {
		(*visitedReductions)[constructorReduction.String()] = &constructorReduction
	} else {
		return
	}

	defaultFields := calculateDefaultFields(constructor, args)

	for _, field := range constructor.Fields {
		fieldType := toTypeReduction(field.FieldType, types, constructors)
		if fieldType != nil {
			fillTypeReduction(fieldType, args, constructor.Type, &defaultFields)
			reduce(*fieldType, visitedReductions, types, constructors, visitedTypes)
		}
	}
}

func calculateDefaultFields(
	constructor *Constructor,
	args []EvaluatedType,
) map[string]bool {
	defaults := make(map[string]bool)

	for _, field := range constructor.Fields {
		if field.Mask != nil {
			name := field.Mask.MaskName
			bit := field.Mask.BitNumber
			if _, ok := defaults[name]; ok {
				defaults[field.FieldName] = true
				continue
			}
			if argIndex := findArgByName(name, constructor.Type.TypeArguments); argIndex != -1 {
				arg := args[argIndex]
				if arg.Index == 0 && (arg.Constant&(1<<bit) == 0) {
					defaults[field.FieldName] = true
					continue
				}
			}
		}
	}

	return defaults
}

func findArgByName(targetArg string, args []tlast.TemplateArgument) int {
	for i, arg := range args {
		if arg.FieldName == targetArg {
			return i
		}
	}
	return -1
}

func fillTypeReduction(
	typeReduction *TypeReduction,
	args []EvaluatedType,
	originalType *TypeDefinition,
	defaultFields *map[string]bool,
) {
	for argI, arg := range typeReduction.Arguments {
		switch arg.Index {
		// nat var
		case NumberVariable:
			j := findArgByName(arg.Variable, originalType.TypeArguments)
			if j != -1 && args[j].Index == NumberConstant {
				typeReduction.Arguments[argI] = args[j]
			} else if j != -1 && args[j].Index == NumberVariable && args[j].VariableActsAsConstant {
				typeReduction.Arguments[argI] = args[j]
				typeReduction.Arguments[argI].VariableActsAsConstant = true
			} else if _, ok := (*defaultFields)[arg.Variable]; ok {
				typeReduction.Arguments[argI] = EvaluatedType{Index: NumberConstant, Constant: 0}
			}
		// type
		case TypeConstant:
			fillTypeReduction(arg.Type, args, originalType, defaultFields)
		// type var
		case TypeVariable:
			j := findArgByName(arg.TypeVariable, originalType.TypeArguments)
			if j != -1 {
				index := args[j].Index
				if index == TypeConstant || index == TypeVariable {
					typeReduction.Arguments[argI] = args[j]
				}
			}
		}
	}
}

func toTypeReduction(
	typeRef tlast.TypeRef,
	types *map[TypeName]*TypeDefinition,
	constructors *map[ConstructorName]*Constructor,
) *TypeReduction {
	var reduction TypeReduction

	typeName := typeRef.Type
	var targetType *TypeDefinition

	if constr, isConstructor := (*constructors)[typeName]; isConstructor {
		reduction.IsType = false
		reduction.Constructor = constr
		targetType = constr.Type
	} else if typ, isType := (*types)[typeName]; isType {
		reduction.IsType = true
		reduction.Type = typ
		targetType = typ
	} else {
		return nil
	}

	evalArguments := make([]EvaluatedType, 0)

	for i, actualArg := range typeRef.Args {
		typeArg := targetType.TypeArguments[i]
		if actualArg.IsArith {
			evalArguments = append(evalArguments, EvaluatedType{
				Index:    NumberConstant,
				Constant: actualArg.Arith.Res,
			})
		} else if typeArg.IsNat {
			evalArguments = append(evalArguments, EvaluatedType{
				Index:    NumberVariable,
				Variable: actualArg.T.Type.Name,
			})
		} else {
			typ := toTypeReduction(actualArg.T, types, constructors)
			if typ != nil {
				evalArguments = append(evalArguments, EvaluatedType{
					Index: TypeConstant,
					Type:  typ,
				})
			} else {
				evalArguments = append(evalArguments, EvaluatedType{
					Index:        TypeVariable,
					TypeVariable: actualArg.T.Type.Name,
				})
			}
		}
	}

	reduction.Arguments = evalArguments

	return &reduction
}

type TypesInfo struct {
	Types        map[TypeName]*TypeDefinition
	Constructors map[ConstructorName]*Constructor
	//TypeReductions map[string]*TypeReduction
}

// works for given constructor or for 1-st
func (ti *TypesInfo) FieldTypeReduction(tr *TypeReduction, fieldId int) EvaluatedType {
	constructor := tr.Constructor
	if tr.IsType {
		constructor = tr.ReferenceType().Constructors[0]
	}
	field := constructor.Fields[fieldId]
	fieldType := toTypeReduction(field.FieldType, &ti.Types, &ti.Constructors)
	if fieldType == nil {
		genericName := field.FieldType.Type.Name
		i := findArgByName(genericName, tr.ReferenceType().TypeArguments)
		if i == -1 || tr.Arguments[i].Type == nil {
			return EvaluatedType{Index: TypeVariable, TypeVariable: genericName}
		}
		return EvaluatedType{Index: TypeConstant, Type: tr.Arguments[i].Type}
	}
	defaultValues := calculateDefaultFields(constructor, tr.Arguments)
	fillTypeReduction(fieldType, tr.Arguments, tr.ReferenceType(), &defaultValues)
	return EvaluatedType{Index: TypeConstant, Type: fieldType}
}

func (ti *TypesInfo) ResultTypeReduction(tr *TypeReduction) EvaluatedType {
	constructor := tr.Constructor
	if tr.IsType {
		constructor = tr.ReferenceType().Constructors[0]
	}
	resultRef := constructor.Result
	if resultRef == nil {
		panic("not a function")
	}

	resultType := toTypeReduction(*resultRef, &ti.Types, &ti.Constructors)
	defaultValues := calculateDefaultFields(constructor, tr.Arguments)
	fillTypeReduction(resultType, tr.Arguments, tr.ReferenceType(), &defaultValues)
	return EvaluatedType{Index: TypeConstant, Type: resultType}
}

func (ti *TypesInfo) TypeNameToGenericTypeReduction(t TypeName) TypeReduction {
	var rd TypeReduction
	rd.Type, rd.IsType = ti.Types[t]
	rd.Constructor = ti.Constructors[t]

	refType := rd.ReferenceType()
	for i, arg := range refType.TypeArguments {
		var evalType EvaluatedType
		if arg.IsNat {
			evalType = EvaluatedType{Index: NumberVariable, Variable: fmt.Sprintf("_Nat%d", i)}
		} else {
			evalType = EvaluatedType{Index: TypeVariable, TypeVariable: fmt.Sprintf("_Type%d", i)}
		}
		rd.Arguments = append(rd.Arguments, evalType)
	}

	return rd
}

func (ti *TypesInfo) TypeRWWrapperToTypeReduction(t *TypeRWWrapper) TypeReduction {
	tr := ti.TypeNameToGenericTypeReduction(t.tlName)
	for i, arg := range t.arguments {
		if arg.tip != nil {
			evalArg := ti.TypeRWWrapperToTypeReduction(arg.tip)
			tr.Arguments[i] = EvaluatedType{Index: TypeConstant, Type: &evalArg}
		} else {
			if arg.isArith {
				tr.Arguments[i] = EvaluatedType{Index: NumberConstant, Constant: arg.Arith.Res}
			}
		}
	}
	return tr
}
