// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlcodegen

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"

	"github.com/vkcom/tl/internal/tlast"
)

// During recursive generation, we store wrappers to type when they are needed, so that
// we can generate actual types later, when all references to wrappers are set
// also wrapper stores common information

type Deconflicter struct {
	usedNames map[string]bool
}

func (d *Deconflicter) hasConflict(s string) bool {
	_, ok := d.usedNames[s]
	return ok
}

func (d *Deconflicter) deconflictName(s string) string {
	if d.usedNames == nil {
		d.usedNames = map[string]bool{}
	}
	var suffix string
	for i := 0; d.usedNames[s+suffix]; i++ {
		suffix = strconv.Itoa(i)
	}
	s += suffix
	d.usedNames[s] = true
	return s
}

func (d *Deconflicter) fillCPPIdentifiers() { // TODO - full list
	d.deconflictName("int")
	d.deconflictName("double")
	d.deconflictName("float")
	d.deconflictName("long")
	d.deconflictName("else")
	d.deconflictName("inline")
	d.deconflictName("namespace")
}

type TypeRWWrapper struct {
	gen *Gen2 // options and packages are here

	ns        *Namespace
	ins       *InternalNamespace
	trw       TypeRW
	NatParams []string // external params of type Read/Write method, with nat_ prefix

	arguments []ResolvedArgument

	goGlobalName string // globally unique, so could be used also in html anchors, internal C++ function names, etc.
	goLocalName  string
	cppLocalName string

	wantsBytesVersion bool
	preventUnwrap     bool // we can have infinite typedef loop in rare cases

	hasBytesVersion bool

	fileName string
	tlTag    uint32     // TODO - turn into function
	tlName   tlast.Name // TODO - turn into function constructor name or union name for code generation
	origTL   []*tlast.Combinator

	isTopLevel bool

	unionParent *TypeRWWrapper // a bit hackish, but simple
	unionField  Field
	unionIndex  int
	unionIsEnum bool

	WrLong        *TypeRWWrapper // long transitioning code
	WrWithoutLong *TypeRWWrapper // long transitioning code
}

func (w *TypeRWWrapper) CanonicalStringTop() string {
	return w.CanonicalString(len(w.origTL) <= 1) // single constructors, arrays and primitives are naturally bare, unions are naturally boxed
}

func (w *TypeRWWrapper) CanonicalString(bare bool) string {
	var s strings.Builder
	if len(w.origTL) > 1 {
		if bare {
			panic("CanonicalString of bare union")
		}
		w.origTL[0].TypeDecl.Name.WriteString(&s)
	} else if len(w.origTL) == 1 {
		if bare {
			w.origTL[0].Construct.Name.WriteString(&s)
		} else {
			w.origTL[0].TypeDecl.Name.WriteString(&s)
		}
	} else {
		panic("all builtins are parsed from TL text, so must have exactly one constructor")
	}
	if len(w.arguments) == 0 {
		return s.String()
	}
	s.WriteByte('<')
	for i, a := range w.arguments {
		// fieldName := t.origTL[0].TemplateArguments[i].FieldName // arguments must be the same for all union elements
		if i != 0 {
			s.WriteByte(',')
		}
		if a.isNat {
			if a.isArith {
				s.WriteString(strconv.FormatUint(uint64(a.Arith.Res), 10))
			} else {
				s.WriteString("#") // TODO - write fieldName here if special argument to function is set
			}
		} else {
			s.WriteString(a.tip.CanonicalString(a.bare))
		}
	}
	s.WriteByte('>')
	return s.String()
}

func (w *TypeRWWrapper) HasAnnotation(str string) bool {
	for _, m := range w.origTL[0].Modifiers {
		if m.Name == str {
			return true
		}
	}
	return false
}

func (w *TypeRWWrapper) AnnotationsMask() uint32 {
	var mask uint32
	for bit, v := range w.gen.allAnnotations {
		if w.HasAnnotation(v) {
			mask |= (1 << bit)
		}
	}
	return mask
}

// Assign structural names to external arguments
func (w *TypeRWWrapper) NatArgs(result []ActualNatArg, prefix string) []ActualNatArg {
	for i, a := range w.arguments {
		fieldName := w.origTL[0].TemplateArguments[i].FieldName // arguments must be the same for all union elements
		if a.isNat {
			if !a.isArith {
				result = append(result, ActualNatArg{
					isArith: a.isArith,
					Arith:   a.Arith,
					name:    prefix + fieldName,
				})
			}
		} else {
			result = a.tip.NatArgs(result, prefix+fieldName)
		}
	}
	return result
}

func (w *TypeRWWrapper) resolvedT2GoName(insideNamespace string) (head, tail string) {
	b := strings.Builder{}
	for _, a := range w.arguments {
		if a.isNat {
			if a.isArith {
				b.WriteString(strconv.FormatUint(uint64(a.Arith.Res), 10))
			}
		} else {
			head, tail := a.tip.resolvedT2GoName(insideNamespace)
			b.WriteString(head)
			b.WriteString(tail)
			if !a.bare {
				b.WriteString("Boxed")
			}
		}
	}
	// We keep compatibility with legacy golang naming
	// This is customization point, generated code should work with whatever naming strategy is selected here
	if len(w.origTL) == 1 && (w.origTL[0].TypeDecl.Name.String() == "_" || w.origTL[0].IsFunction || w.unionParent != nil) {
		return canonicalGoName(w.origTL[0].Construct.Name, insideNamespace), b.String()
	}
	return canonicalGoName(w.origTL[0].TypeDecl.Name, insideNamespace), b.String()
}

// for golang cycle detection
type DirectImports struct {
	ns         map[*InternalNamespace]struct{}
	importSort bool
}

// for C++ includes
type DirectIncludesCPP struct {
	ns map[string]struct{}
}

func (d DirectIncludesCPP) sortedNames() []string {
	var sortedNames []string
	for im := range d.ns { // Imports of this file.
		sortedNames = append(sortedNames, im)
	}
	slices.Sort(sortedNames)
	return sortedNames
}

func stringCompare(a string, b string) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

func TypeRWWrapperLessLocal(a *TypeRWWrapper, b *TypeRWWrapper) int {
	an := a.TypeString2(false, nil, nil, true, true)
	bn := b.TypeString2(false, nil, nil, true, true)
	return stringCompare(an, bn)
}

func TypeRWWrapperLessGlobal(a *TypeRWWrapper, b *TypeRWWrapper) int {
	// return stringCompare(a.CanonicalString(), b.CanonicalString()) TODO - better idea after everything is stabilized
	return stringCompare(a.goGlobalName, b.goGlobalName)
}

func (w *TypeRWWrapper) ShouldWriteTypeAlias() bool { // TODO - interface method
	if _, ok := w.trw.(*TypeRWStruct); ok {
		if w.unionParent == nil || !w.unionIsEnum {
			return true
		}
	}
	if _, ok := w.trw.(*TypeRWUnion); ok {
		return true
	}
	if _, ok := w.trw.(*TypeRWMaybe); ok {
		return true
	}
	return false
}

func (w *TypeRWWrapper) ShouldWriteEnumElementAlias() bool {
	_, ok := w.trw.(*TypeRWStruct)
	return ok && w.unionParent != nil && w.unionIsEnum
}

func (w *TypeRWWrapper) MarkHasBytesVersion(visitedNodes map[*TypeRWWrapper]bool) bool {
	if visitedNodes[w] {
		return false // We OR results of fields, so if we visited field, and it returned true, this true is already recorded
	}
	visitedNodes[w] = true
	return w.trw.markHasBytesVersion(visitedNodes)
}

func (w *TypeRWWrapper) FillRecursiveUnwrap(visitedNodes map[*TypeRWWrapper]bool) {
	if visitedNodes[w] {
		return
	}
	visitedNodes[w] = true
	w.trw.fillRecursiveUnwrap(visitedNodes)
}

func (w *TypeRWWrapper) FillRecursiveChildren(visitedNodes map[*TypeRWWrapper]bool) {
	if visitedNodes[w] {
		return
	}
	visitedNodes[w] = true
	w.trw.fillRecursiveChildren(visitedNodes)
}

func (w *TypeRWWrapper) MarkWantsBytesVersion(visitedNodes map[*TypeRWWrapper]bool) {
	if visitedNodes[w] {
		return
	}
	w.wantsBytesVersion = true
	visitedNodes[w] = true
	w.trw.markWantsBytesVersion(visitedNodes)
}

func (w *TypeRWWrapper) TypeString2(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, isLocal bool, skipAlias bool) string {
	bytesVersion = bytesVersion && w.hasBytesVersion
	return w.trw.typeString2(bytesVersion, directImports, ins, isLocal, skipAlias)
}
func (w *TypeRWWrapper) TypeResettingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, ref bool) string {
	bytesVersion = bytesVersion && w.hasBytesVersion
	return w.trw.typeResettingCode(bytesVersion, directImports, ins, val, ref)
}
func (w *TypeRWWrapper) TypeRandomCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, natArgs []string, ref bool) string {
	bytesVersion = bytesVersion && w.hasBytesVersion
	return w.trw.typeRandomCode(bytesVersion, directImports, ins, val, natArgs, ref)
}
func (w *TypeRWWrapper) TypeWritingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, bare bool, natArgs []string, ref bool, last bool) string {
	bytesVersion = bytesVersion && w.hasBytesVersion
	return w.trw.typeWritingCode(bytesVersion, directImports, ins, val, bare, natArgs, ref, last)
}
func (w *TypeRWWrapper) TypeReadingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, bare bool, natArgs []string, ref bool, last bool) string {
	bytesVersion = bytesVersion && w.hasBytesVersion
	return w.trw.typeReadingCode(bytesVersion, directImports, ins, val, bare, natArgs, ref, last)
}
func (w *TypeRWWrapper) TypeJSONEmptyCondition(bytesVersion bool, val string, ref bool) string {
	bytesVersion = bytesVersion && w.hasBytesVersion
	return w.trw.typeJSONEmptyCondition(bytesVersion, val, ref)
}
func (w *TypeRWWrapper) TypeJSONWritingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, natArgs []string, ref bool) string {
	bytesVersion = bytesVersion && w.hasBytesVersion
	return w.trw.typeJSONWritingCode(bytesVersion, directImports, ins, val, natArgs, ref)
}
func (w *TypeRWWrapper) TypeJSONReadingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, jvalue string, val string, natArgs []string, ref bool) string {
	bytesVersion = bytesVersion && w.hasBytesVersion
	return w.trw.typeJSONReadingCode(bytesVersion, directImports, ins, jvalue, val, natArgs, ref)
}

func (w *TypeRWWrapper) IsTrueType() bool {
	structElement, ok := w.trw.(*TypeRWStruct)
	if !ok {
		return false
	}
	return len(structElement.Fields) == 0
}

func (w *TypeRWWrapper) CPPFillRecursiveChildren(visitedNodes map[*TypeRWWrapper]bool) {
	if visitedNodes[w] {
		return
	}
	visitedNodes[w] = true
	w.trw.CPPFillRecursiveChildren(visitedNodes)
}

func (w *TypeRWWrapper) CPPTypeStringInNamespace(bytesVersion bool, hppInc *DirectIncludesCPP) string {
	return w.trw.cppTypeStringInNamespace(bytesVersion, hppInc)
}

func (w *TypeRWWrapper) CPPTypeStringInNamespaceHalfResolved(bytesVersion bool, hppInc *DirectIncludesCPP, halfResolved HalfResolvedArgument) string {
	if halfResolved.Name != "" {
		return halfResolved.Name
	}
	return w.trw.cppTypeStringInNamespaceHalfResolved(bytesVersion, hppInc, halfResolved)
}

func (w *TypeRWWrapper) CPPDefaultInitializer(halfResolved HalfResolvedArgument, halfResolve bool) string {
	if halfResolve && halfResolved.Name != "" {
		return "{}"
	}
	return w.trw.cppDefaultInitializer(halfResolved, halfResolve)
}

func (w *TypeRWWrapper) cppNamespaceQualifier() string {
	var s strings.Builder
	s.WriteString("::")
	s.WriteString(w.gen.options.RootCPPNamespace)
	s.WriteString("::")
	//myWrapper.cppNamespaceQualifier = "::" + gen.options.RootCPPNamespace + "::"
	if w.tlName.Namespace != "" {
		s.WriteString(w.tlName.Namespace)
		s.WriteString("::")
	}
	return s.String()
}

func canonicalCPPName(name tlast.Name, insideNamespace string) string { // TODO
	if name.Namespace == insideNamespace {
		return name.Name
	}
	return name.Namespace + "_" + name.Name
}

func (w *TypeRWWrapper) fullyResolvedClassCppNameArgs() (string, []string) { // name in namespace, arguments decl
	cppSuffix := strings.Builder{}
	cppSuffix.WriteString(w.tlName.Name)
	var cppArgsDecl []string
	for i, a := range w.arguments {
		fieldName := w.origTL[0].TemplateArguments[i].FieldName // arguments must be the same for all union elements
		if a.isNat {
			if a.isArith {
				cppSuffix.WriteString(fieldName)
				cppArgsDecl = append(cppArgsDecl, "uint32_t "+fieldName)
			}
		} else {
			cppArgsDecl = append(cppArgsDecl, "typename "+fieldName)
		}
	}
	return cppSuffix.String(), cppArgsDecl
}

func (w *TypeRWWrapper) cppTypeStringInNamespace(bytesVersion bool, hppInc *DirectIncludesCPP, halfResolve bool, halfResolved HalfResolvedArgument) (string, string, string) {
	hppInc.ns[w.fileName] = struct{}{}
	bName := strings.Builder{}
	// bName.WriteString(w.cppNamespaceQualifier())
	bName.WriteString(w.tlName.Name)
	bCanonicalName := strings.Builder{}
	bCanonicalName.WriteString(w.tlName.Name)
	bArgs := strings.Builder{}
	for i, a := range w.arguments {
		fieldName := w.origTL[0].TemplateArguments[i].FieldName // arguments must be the same for all union elements
		if a.isNat {
			if a.isArith {
				bName.WriteString(fieldName)
				bCanonicalName.WriteString("-") // simple canonical format where each argument is preceded with '-'. Good for filenames also
				bCanonicalName.WriteString(fieldName)
				if bArgs.Len() == 0 {
					bArgs.WriteString("<")
				} else {
					bArgs.WriteString(", ")
				}
				if halfResolve {
					half := halfResolved.Args[i]
					if half.Name != "" {
						bArgs.WriteString(half.Name)
						continue
					}
				}
				bArgs.WriteString(strconv.FormatUint(uint64(a.Arith.Res), 10))
			}
		} else {
			if bArgs.Len() == 0 {
				bArgs.WriteString("<")
			} else {
				bArgs.WriteString(", ")
			}
			if halfResolve {
				half := halfResolved.Args[i]
				if half.Name != "" {
					bArgs.WriteString(half.Name)
					continue
				}
			}
			bArgs.WriteString(a.tip.CPPTypeStringInNamespace(bytesVersion, hppInc))
		}
	}
	if bArgs.Len() != 0 {
		bArgs.WriteString(">")
	}
	return bName.String(), bCanonicalName.String(), bArgs.String()
}

func (w *TypeRWWrapper) JSONHelpString() string {
	return w.CanonicalStringTop()
}

func (w *TypeRWWrapper) JSONHelpFullType(bare bool, fields []Field, natArgs []ActualNatArg) string {
	result := w.helpString2(bare, fields, &natArgs)
	if len(natArgs) != 0 {
		panic("JSONHelpFullType should consume all arguments")
	}
	return result
}

func (w *TypeRWWrapper) JSONHelpNatArg(fields []Field, natArg ActualNatArg) string {
	if natArg.isArith {
		return fmt.Sprintf("%d", natArg.Arith.Res)
	}
	if natArg.isField {
		return fields[natArg.FieldIndex].originalName
	}
	return natArg.name
}

func (w *TypeRWWrapper) helpString2(bare bool, fields []Field, natArgs *[]ActualNatArg) string {
	var s strings.Builder
	if len(w.origTL) > 1 {
		if bare {
			panic("helpString2 of bare union")
		}
		s.WriteString(w.origTL[0].TypeDecl.Name.String())
	} else {
		if bare {
			s.WriteString(w.origTL[0].Construct.Name.String())
		} else {
			s.WriteString(w.origTL[0].TypeDecl.Name.String())
		}
	}
	if len(w.arguments) == 0 {
		return s.String()
	}
	s.WriteString("<")
	for i, a := range w.arguments {
		if i != 0 {
			s.WriteString(",")
		}
		if a.isNat {
			if a.isArith {
				s.WriteString(fmt.Sprintf("%d", a.Arith.Res))
			} else {
				natArg := (*natArgs)[0]
				*natArgs = (*natArgs)[1:]
				if natArg.isField {
					s.WriteString(fields[natArg.FieldIndex].originalName)
				} else {
					s.WriteString(natArg.name)
				}
			}
		} else {
			s.WriteString(a.tip.helpString2(a.bare, fields, natArgs))
		}
	}
	s.WriteString(">")
	return s.String()
}

// same code as in func (trw *TypeRWStruct) replaceUnwrapArgs
func (w *TypeRWWrapper) transformNatArgsToChild(natArgs []ActualNatArg, childNatArgs []ActualNatArg) []ActualNatArg {
	var result []ActualNatArg
outer:
	for _, arg := range childNatArgs {
		if arg.isArith || arg.isField {
			panic("cannot transform to child arith or field nat param")
		}
		for i, p := range w.NatParams {
			if p == arg.name {
				result = append(result, natArgs[i])
				continue outer
			}
		}
		panic("nat param not found in parent nat params")
	}
	return result
}

// TODO remove skipAlias after we start generating go code like we do for C++
type TypeRW interface {
	// methods below are target language independent
	markWantsBytesVersion(visitedNodes map[*TypeRWWrapper]bool)
	fillRecursiveUnwrap(visitedNodes map[*TypeRWWrapper]bool)

	BeforeCodeGenerationStep1() // during first phase, some wr.trw are nil due to recursive types. So we delay some
	BeforeCodeGenerationStep2() // during second phase, union fields recursive bit is set

	// methods below depend on target language
	fillRecursiveChildren(visitedNodes map[*TypeRWWrapper]bool)
	IsDictKeySafe() (isSafe bool, isString bool) // natives are safe, other types TBD
	typeString2(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, isLocal bool, skipAlias bool) string
	markHasBytesVersion(visitedNodes map[*TypeRWWrapper]bool) bool
	typeResettingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, ref bool) string
	typeRandomCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, natArgs []string, ref bool) string
	typeWritingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, bare bool, natArgs []string, ref bool, last bool) string
	typeReadingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, bare bool, natArgs []string, ref bool, last bool) string
	typeJSONEmptyCondition(bytesVersion bool, val string, ref bool) string
	typeJSONWritingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, natArgs []string, ref bool) string
	typeJSONReadingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, jvalue string, val string, natArgs []string, ref bool) string
	GenerateCode(bytesVersion bool, directImports *DirectImports) string

	CPPFillRecursiveChildren(visitedNodes map[*TypeRWWrapper]bool)
	cppTypeStringInNamespace(bytesVersion bool, hppInc *DirectIncludesCPP) string
	cppTypeStringInNamespaceHalfResolved(bytesVersion bool, hppInc *DirectIncludesCPP, halfResolved HalfResolvedArgument) string
	cppDefaultInitializer(halfResolved HalfResolvedArgument, halfResolve bool) string
	CPPHasBytesVersion() bool
	CPPTypeResettingCode(bytesVersion bool, val string) string
	CPPTypeWritingCode(bytesVersion bool, val string, bare bool, natArgs []string, last bool) string
	CPPTypeReadingCode(bytesVersion bool, val string, bare bool, natArgs []string, last bool) string
	CPPGenerateCode(hpp *strings.Builder, hppInc *DirectIncludesCPP, hppIncFwd *DirectIncludesCPP, hppDet *strings.Builder, hppDetInc *DirectIncludesCPP, cppDet *strings.Builder, cppDetInc *DirectIncludesCPP, bytesVersion bool, forwardDeclaration bool)
}

type Field struct {
	originalName string
	t            *TypeRWWrapper
	bare         bool
	goName       string
	cppName      string
	recursive    bool

	fieldMask *ActualNatArg
	BitNumber uint32 // only used when fieldMask != nil

	natArgs      []ActualNatArg
	halfResolved HalfResolvedArgument

	origTL tlast.Field
}

func (f *Field) Bare() bool {
	return f.bare
}

func formatNatArg(fields []Field, arg ActualNatArg) string {
	if arg.isArith {
		return strconv.FormatUint(uint64(arg.Arith.Res), 10)
	}
	if arg.isField {
		return "item." + fields[arg.FieldIndex].goName
	}
	if strings.HasPrefix(arg.name, "nat_") {
		panic("aha!") // TODO - remove
	}
	return "nat_" + arg.name
}

func formatNatArgCPP(fields []Field, arg ActualNatArg) string { // TODO - harmonize with formatNatArg?
	if arg.isArith {
		return strconv.FormatUint(uint64(arg.Arith.Res), 10)
	}
	if arg.isField {
		return "item." + fields[arg.FieldIndex].cppName
	}
	return "nat_" + arg.name
}

func formatNatArgsCPP(fields []Field, natArgs []ActualNatArg) []string {
	var result []string
	for _, arg := range natArgs {
		result = append(result, formatNatArgCPP(fields, arg))
	}
	return result
}

func formatNatArgsCallCPP(natArgs []string) string {
	return formatNatArgsDeclCall(natArgs)
}

func formatNatArgsDeclCPP(natArgs []string) string {
	var s strings.Builder
	for _, arg := range natArgs {
		s.WriteString(fmt.Sprintf(", uint32_t nat_%s", arg))
	}
	return s.String()
}

// TODO - remove all trash functions and consolidate into 1 or 2
func formatNatArgsAddNat(natArgs []string) []string {
	var result []string
	for _, arg := range natArgs {
		result = append(result, "nat_"+arg)
	}
	return result
}

func formatNatArgs(fields []Field, natArgs []ActualNatArg) []string {
	var result []string
	for _, arg := range natArgs {
		if !arg.isArith {
			result = append(result, formatNatArg(fields, arg))
		}
	}
	return result
}

func formatNatArgsDecl(natArgs []string) string {
	var s strings.Builder
	for _, arg := range natArgs {
		s.WriteString(fmt.Sprintf(",nat_%s uint32", arg))
	}
	return s.String()
}

func formatNatArgsDeclNoComma(natArgs []string) string {
	return strings.TrimPrefix(formatNatArgsDecl(natArgs), ",")
}

// if our fun is declared as ReadBoxed(..., nat_x uint32, nat_y uint32) using formatNatArgsDecl() above,
// and we want to pass arguments to our own function, like Read(..., nat_x, nat_y)
func formatNatArgsDeclCall(natArgs []string) string {
	var s strings.Builder
	for _, arg := range natArgs {
		s.WriteString(fmt.Sprintf(", nat_%s", arg))
	}
	return s.String()
}

// simply adds commas, natArgs are already fully formatted. Difference to strings.Join is leading comma
func joinWithCommas(natArgs []string) string {
	var s strings.Builder
	for _, arg := range natArgs {
		s.WriteString(fmt.Sprintf(", %s", arg))
	}
	return s.String()
}

func addBytes(val string, bytesVersion bool) string {
	return ifString(bytesVersion, val+"Bytes", val)
}

func addBare(bare bool) string {
	return ifString(bare, "", "Boxed")
}

func addAmpersand(ref bool, val string) string {
	return ifString(ref, val, "&"+val)
}

func addAsterisk(ref bool, val string) string {
	return ifString(ref, "*"+val, val)
}

func wrapLast(last bool, code string) string {
	return ifString(last, "return "+code+"", "if err := "+code+"; err != nil { return err }")
}

func wrapLastW(last bool, code string) string {
	return ifString(last, "return "+code+"", "if w, err = "+code+"; err != nil { return w, err }")
}

func ifString(value bool, t string, f string) string {
	if value {
		return t
	}
	return f
}

var (
	camelingRegex = regexp.MustCompile(`[0-9A-Za-z]+`)
	allUpperRegex = regexp.MustCompile(`^[A-Z][A-Z0-9]+$`)
)

// TODO - investigate if this function is good
func CNameToCamelName(s string) string {
	chunks := camelingRegex.FindAllString(s, -1)
	for i, chunk := range chunks {
		if allUpperRegex.MatchString(chunk) { // TODO - why?
			chunks[i] = strings.ToUpper(chunk[:1]) + strings.ToLower(chunk[1:])
		} else {
			chunks[i] = ToUpperFirst(chunk)
		}
	}
	return strings.Join(chunks, "")
}

func ToUpperFirst(str string) string {
	for i := range str {
		if i != 0 {
			return strings.ToUpper(str[:i]) + str[i:]
		}
	}
	return strings.ToUpper(str) // zero or single rune
}

func ToLowerFirst(str string) string {
	for i := range str {
		if i != 0 {
			return strings.ToLower(str[:i]) + str[i:]
		}
	}
	return strings.ToLower(str) // zero or single rune
}
