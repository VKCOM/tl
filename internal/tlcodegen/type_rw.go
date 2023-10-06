// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlcodegen

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

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

	tlType    tlast.Name // type for canonical name generation, Type for everything except builtins and union elements
	arguments []ResolvedArgument

	wantsBytesVersion bool
	preventUnwrap     bool

	hasBytesVersion bool

	fileName string
	tlTag    uint32
	tlName   tlast.Name // constructor for code generation. Might be good idea to move it to the trw?
	origTL   []*tlast.Combinator

	isTopLevel bool

	unionParent *TypeRWWrapper // a bit hackish, but simple
	unionField  Field
	unionIndex  int
	unionIsEnum bool

	// cppNamespaceQualifier string // full qualifier, like ::tl2::tlstatshouse::, or "" for builtin types

	WrLong        *TypeRWWrapper // long transitioning code
	WrWithoutLong *TypeRWWrapper // long transitioning code
}

func (w *TypeRWWrapper) CanonicalString() string {
	var s strings.Builder
	s.WriteString(w.tlType.String())
	if len(w.arguments) == 0 {
		return s.String()
	}
	s.WriteString("<")
	for i, a := range w.arguments {
		// fieldName := t.origTL[0].TemplateArguments[i].FieldName // arguments must be the same for all union elements
		if i != 0 {
			s.WriteString(",")
		}
		if a.isNat {
			if a.isArith {
				s.WriteString(fmt.Sprintf("%d", a.Arith.Res))
			} else {
				s.WriteString("#") // TODO - write fieldName here if special argument to function is set
			}
		} else {
			if a.bare {
				s.WriteString("%")
			}
			s.WriteString(a.tip.CanonicalString())
		}
	}
	s.WriteString(">")
	return s.String()
}

// External arguments
func (w *TypeRWWrapper) NatArgs(withFixed bool, prefix string) []ActualNatArg {
	return w.natArgs(withFixed, nil, prefix)
}

func (w *TypeRWWrapper) natArgs(withFixed bool, result []ActualNatArg, prefix string) []ActualNatArg {
	for i, a := range w.arguments {
		fieldName := w.origTL[0].TemplateArguments[i].FieldName // arguments must be the same for all union elements
		if a.isNat {
			if withFixed || !a.isArith {
				result = append(result, ActualNatArg{
					isArith: a.isArith,
					Arith:   a.Arith,
					name:    prefix + fieldName,
				})
			}
		} else {
			result = a.tip.natArgs(withFixed, result, prefix+fieldName)
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
	return canonicalGoName(w.tlType, insideNamespace), b.String()
}

// for golang cycle detection
type DirectImports struct {
	ns         map[*InternalNamespace]struct{}
	importSort bool
}

// for C++ includes
/*
type DirectIncludesCPP struct {
	ns map[string]struct{}
}

func (d DirectIncludesCPP) sortedNames() []string {
	var sortedNames []string
	for im := range d.ns { // Imports of this file.
		sortedNames = append(sortedNames, im)
	}
	sort.Strings(sortedNames)
	return sortedNames
}
*/

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
	return stringCompare(a.TypeString(false), b.TypeString(false))
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

func (w *TypeRWWrapper) TypeString(bytesVersion bool) string {
	bytesVersion = bytesVersion && w.hasBytesVersion
	return w.trw.typeStringGlobal(bytesVersion)
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

/*
	func (w *TypeRWWrapper) CPPFillRecursiveChildren(visitedNodes map[*TypeRWWrapper]bool) {
		if visitedNodes[w] {
			return
		}
		visitedNodes[w] = true
		w.trw.CPPFillRecursiveChildren(visitedNodes)
	}

	func (w *TypeRWWrapper) CPPTypeStringInNamespace(bytesVersion bool, hppInc *DirectIncludesCPP, resolvedType ResolvedType, halfResolve bool) string {
		if halfResolve && resolvedType.OriginalName != "" {
			return resolvedType.OriginalName
		}
		return w.trw.cppTypeStringInNamespace(bytesVersion, hppInc, resolvedType, halfResolve)
	}

	func (w *TypeRWWrapper) CPPDefaultInitializer(resolvedType ResolvedType, halfResolve bool) string {
		if halfResolve && resolvedType.OriginalName != "" {
			return "{}"
		}
		return w.trw.cppDefaultInitializer(resolvedType, halfResolve)
	}
*/
func (w *TypeRWWrapper) JSONHelpString() string {
	return w.CanonicalString()
}

func (w *TypeRWWrapper) JSONHelpFullType(fields []Field, natArgs []ActualNatArg) string {
	result := w.helpString2(fields, &natArgs)
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

func (w *TypeRWWrapper) helpString2(fields []Field, natArgs *[]ActualNatArg) string {
	var s strings.Builder
	s.WriteString(w.tlType.String())
	if len(w.arguments) == 0 {
		return s.String()
	}
	s.WriteString("<")
	for i, a := range w.arguments {
		if i != 0 {
			s.WriteString(",")
		}
		if a.isNat {
			natArg := (*natArgs)[0]
			*natArgs = (*natArgs)[1:]
			if a.isArith != natArg.isArith || a.Arith.Res != natArg.Arith.Res {
				panic("helpString2 arith mismatch")
			}
			if natArg.isArith {
				s.WriteString(fmt.Sprintf("%d", natArg.Arith.Res))
			} else {
				if natArg.isField {
					s.WriteString(fields[natArg.FieldIndex].originalName)
				} else {
					s.WriteString(natArg.name)
				}
			}
		} else {
			if a.bare {
				s.WriteString("%")
			}
			s.WriteString(a.tip.helpString2(fields, natArgs))
		}
	}
	s.WriteString(">")
	return s.String()
}

func (w *TypeRWWrapper) transformNatArgsFromParent(parent *TypeRWWrapper, parentNatArgs []ActualNatArg, natArgs []ActualNatArg) []ActualNatArg {
	parentParams := parent.NatArgs(true, "")
	if len(parentParams) != len(parentNatArgs) {
		panic("parent nat params inconsistent length")
	}
	var result []ActualNatArg
outer:
	for _, arg := range natArgs {
		for i, p := range parentParams {
			if p.name == arg.name {
				result = append(result, parentNatArgs[i])
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
	wrapper() *TypeRWWrapper
	canBeBareOrBoxed(bare bool) bool
	markWantsBytesVersion(visitedNodes map[*TypeRWWrapper]bool)
	fillRecursiveUnwrap(visitedNodes map[*TypeRWWrapper]bool)

	BeforeCodeGenerationStep() error // during first phase, some wr.trw are nil due to recursive types. So we delay some
	BeforeCodeGenerationStep2()      // during second phase, union fields recursive bit is set

	// methods below depend on target language
	fillRecursiveChildren(visitedNodes map[*TypeRWWrapper]bool)
	IsDictKeySafe() (isSafe bool, isString bool) // natives are safe, other types TBD
	typeStringGlobal(bytesVersion bool) string
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

	//CPPFillRecursiveChildren(visitedNodes map[*TypeRWWrapper]bool)
	//cppTypeStringInNamespace(bytesVersion bool, hppInc *DirectIncludesCPP, resolvedType ResolvedType, halfResolve bool) string
	//cppDefaultInitializer(resolvedType ResolvedType, halfResolve bool) string
	//CPPHasBytesVersion() bool
	//CPPTypeResettingCode(bytesVersion bool, val string) string
	//CPPTypeWritingCode(bytesVersion bool, val string, bare bool, natArgs []string, last bool) string
	//CPPTypeReadingCode(bytesVersion bool, val string, bare bool, natArgs []string, last bool) string
	//CPPGenerateCode(hpp *strings.Builder, hppInc *DirectIncludesCPP, hppIncFwd *DirectIncludesCPP, hppDet *strings.Builder, hppDetInc *DirectIncludesCPP, cppDet *strings.Builder, cppDetInc *DirectIncludesCPP, bytesVersion bool, forwardDeclaration bool)
}

type Field struct {
	t            *TypeRWWrapper
	bare         bool
	goName       string
	cppName      string
	originalName string
	recursive    bool

	fieldMask *ActualNatArg
	BitNumber uint32 // only used when fieldMask != nil

	originalType tlast.TypeRef // TODO - remove?
	natArgs      []ActualNatArg
}

func (f *Field) Bare() bool {
	return f.bare
}

func (f *Field) checkBareBoxed() error {
	if !f.t.trw.canBeBareOrBoxed(f.Bare()) {
		e1 := f.originalType.PR.BeautifulError(fmt.Errorf("resolved type %q cannot be %s", f.t.CanonicalString(), ifString(f.Bare(), "bare", "boxed")))
		// e2 := f.resolvedType.PR.BeautifulError(fmt.Errorf("referenced here"))
		// TODO - return tlast.BeautifulError2(e1, e2)
		return e1
	}
	return nil
}

func formatNatArg(fields []Field, arg ActualNatArg) string {
	if arg.isArith {
		return strconv.FormatUint(uint64(arg.Arith.Res), 10)
	}
	if arg.isField {
		return "item." + fields[arg.FieldIndex].goName
	}
	return "nat_" + arg.name
}

/*
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
		return formatNatArgsCall(natArgs)
	}

	func formatNatArgsDeclCPP(natArgs []string) string {
		var s strings.Builder
		for _, arg := range natArgs {
			s.WriteString(fmt.Sprintf(", uint32_t %s", arg))
		}
		return s.String()
	}
*/
func formatNatArgs(fields []Field, natArgs []ActualNatArg) []string {
	var result []string
	for _, arg := range natArgs {
		if !OptimizeConstParameters || !arg.isArith {
			result = append(result, formatNatArg(fields, arg))
		}
	}
	return result
}

func formatNatArgsCall(natArgs []string) string {
	var s strings.Builder
	for _, arg := range natArgs {
		s.WriteString(fmt.Sprintf(", %s", arg))
	}
	return s.String()
}

func formatNatArgsDecl(natArgs []string) string {
	var s strings.Builder
	for _, arg := range natArgs {
		s.WriteString(fmt.Sprintf(",%s uint32", arg))
	}
	return s.String()
}

func formatNatArgsDeclNoComma(natArgs []string) string {
	s := formatNatArgsDecl(natArgs)
	if len(s) != 0 {
		return s[1:]
	}
	return s
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

func CNameToCamelName(s string) string {
	chunks := camelingRegex.FindAllString(s, -1)
	for i, chunk := range chunks {
		if allUpperRegex.MatchString(chunk) {
			chunks[i] = strings.ToUpper(chunk[:1]) + strings.ToLower(chunk[1:])
		} else {
			chunks[i] = strings.Title(chunk) //lint:ignore SA1019 "transfer to golang.org/x/text/cases in single PR"
		}
	}
	return strings.Join(chunks, "")
}

//func buildTlName(wr *TypeRWWrapper, natArgs []ActualNatArg) string {
//	switch trw := wr.trw.(type) {
//	case *TypeRWBool:
//		return "bool"
//	case *TypeRWBrackets:
//		switch {
//		case trw.dictLike:
//			return fmt.Sprintf("%s (%s)", wr.tlName, buildTlName(trw.dictKeyField.t, trw.dictValueField.natArgs))
//		case trw.vectorLike:
//			return fmt.Sprintf("%s (%s)", wr.tlName, buildTlName(trw.element.t, trw.element.natArgs))
//		case trw.dynamicSize:
//			return fmt.Sprintf("%s (%s %s)", wr.tlName, trw.element.t.tlName, trw.element.natArgs[0].name)
//		default:
//			return fmt.Sprintf("%s (%s %d)", wr.tlName, trw.element.t.tlName, trw.size)
//		}
//	case *TypeRWMaybe:
//		return fmt.Sprintf("%s (%s)", wr.tlName, buildTlName(trw.element.t, trw.element.natArgs))
//	case *TypeRWPrimitive:
//		return trw.primitiveType
//	case *TypeRWStruct:
//		log.Print("buildTlName STRUCT ", trw.goGlobalName, trw.Fields[0].originalType)
//		var templateFields strings.Builder
//		return fmt.Sprintf("%s (%s)", wr.tlName, templateFields.String())
//	case *TypeRWUnion:
//		log.Print("buildTlName UNION ", trw.goGlobalName)
//		return wr.tlName.String()
//	default:
//		return ""
//	}
//}
