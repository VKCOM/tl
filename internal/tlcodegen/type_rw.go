// Copyright 2025 V Kontakte LLC
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

	"github.com/vkcom/tl/internal/utils"
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

var bannedCppFieldNames = []string{
	"and",
	"or",
	"friend",
	"xor",
	"operator",
	"errno",
	"class",
	"short",
	"default",
	"signed",
	"new",
	"private",
	"public",
	"protected",
	"struct",
	"return",
}

func (d *Deconflicter) fillCPPIdentifiers() { // TODO - full list
	d.deconflictName("int")
	d.deconflictName("double")
	d.deconflictName("float")
	d.deconflictName("long")
	d.deconflictName("else")
	d.deconflictName("inline")
	d.deconflictName("namespace")

	for _, word := range bannedCppFieldNames {
		d.deconflictName(word)
	}
}

func (d *Deconflicter) fillGolangIdentifies() {
	d.deconflictName("Write")
	d.deconflictName("Read")
	d.deconflictName("WriteTL2")
	d.deconflictName("ReadTL2")
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
	wantsTL2          bool
	preventUnwrap     bool // we can have infinite typedef loop in rare cases

	hasBytesVersion        bool
	hasErrorInWriteMethods bool

	fileName string

	originateFromTL2 bool

	// cpp info
	hppDetailsFileName string
	cppDetailsFileName string
	groupName          string
	typeComponent      int

	// php info
	phpInfo PhpClassMeta

	// tl1 info
	tlTag  uint32     // TODO - turn into function
	tlName tlast.Name // TODO - turn into function constructor name or union name for code generation
	origTL []*tlast.Combinator

	unionParent *TypeRWUnion // a bit hackish, but simple
	unionIndex  int

	WrLong        *TypeRWWrapper // long transitioning code
	WrWithoutLong *TypeRWWrapper // long transitioning code

	// tl2 info (if union is not nil otherwise check there)
	tl2Name              tlast.TL2TypeName
	tl2Origin            *tlast.TL2Combinator
	tl2IsResult          bool
	tl2IsBuiltinBrackets bool
}

func (wr *TypeRWWrapper) Namespace() string {
	if wr.originateFromTL2 {
		return wr.tl2Name.Namespace
	} else {
		return wr.tlName.Namespace
	}
}

// Those have unique structure fully defined by the magic.
// items with condition len(w.NatParams) == 0 could be serialized independently, but if there is several type instantiations,
// they could not be distinguished by the magic. For example vector<int> and vector<long>.
func (w *TypeRWWrapper) IsTopLevel() bool {
	if w.originateFromTL2 {
		if w.unionParent == nil {
			if w.tl2IsResult {
				return false
			}
			if w.tl2IsBuiltinBrackets {
				return false
			}
			if w.tl2Origin != nil {
				if w.tl2Origin.IsFunction {
					return true
				}
				return len(w.tl2Origin.TypeDecl.TemplateArguments) == 0
			}
			return false
		} else {
			return false
		}
	}
	return len(w.origTL[0].TemplateArguments) == 0
}

func (w *TypeRWWrapper) CanonicalStringTop() string {
	return w.CanonicalString(len(w.origTL) <= 1) // single constructors, arrays and primitives are naturally bare, unions are naturally boxed
}

func (w *TypeRWWrapper) CanonicalString(bare bool) string {
	var s strings.Builder
	if w.originateFromTL2 {
		if w.unionParent == nil {
			if w.tl2IsResult {
				s.WriteString(w.tl2Origin.FuncDecl.Name.String() + "__Result")
			} else if w.tl2IsBuiltinBrackets {
				s.WriteString("__builtin_brackets")
			} else if w.tl2Origin != nil {
				s.WriteString(w.tl2Origin.TypeDecl.Name.String())
			} else {
				s.WriteString(w.tl2Name.String())
			}
		} else {
			originType := w.unionParent.wr.tl2Origin
			if w.unionParent.wr.tl2IsResult {
				s.WriteString(originType.FuncDecl.Name.String() + "__Result" + originType.FuncDecl.ReturnType.UnionType.Variants[w.unionIndex].Name)
			} else {
				s.WriteString(originType.TypeDecl.Name.String() + originType.TypeDecl.Type.UnionType.Variants[w.unionIndex].Name)
			}
		}
	} else {
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
	if w.originateFromTL2 {
		for _, annotation := range w.tl2Origin.Annotations {
			if annotation.Name == str {
				return true
			}
		}
	} else {
		for _, m := range w.origTL[0].Modifiers {
			if m.Name == str {
				return true
			}
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

func (w *TypeRWWrapper) DoArgumentsContainUnionTypes() bool {
	if w, ok := w.trw.(*TypeRWStruct); ok && w.ResultType != nil {
		return w.wr.containsUnion(map[*TypeRWWrapper]bool{})
	} else {
		return false
	}
}

func (w *TypeRWWrapper) DoesReturnTypeContainUnionTypes() bool {
	if w, ok := w.trw.(*TypeRWStruct); ok && w.ResultType != nil {
		return w.ResultType.containsUnion(map[*TypeRWWrapper]bool{})
	} else {
		return false
	}
}

func (w *TypeRWWrapper) containsUnion(visitedNodes map[*TypeRWWrapper]bool) bool {
	if _, ok := visitedNodes[w]; !ok {
		visitedNodes[w] = false
		return w.trw.ContainsUnion(visitedNodes)
	} else {
		return false
	}
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

func (w *TypeRWWrapper) ActualTypeDependencies(evalType EvaluatedType) (res []*TypeRWWrapper) {
	r := make(map[*TypeRWWrapper]bool)
	w.actualTypeDependenciesRecur(evalType, &r)
	for arg := range r {
		//if br, isBr := arg.trw.(*TypeRWBrackets); isBr && br.IsBuiltinVector() {
		//	continue
		//}
		res = append(res, arg)
	}
	slices.SortFunc(res, TypeComparator)
	return
}

func (w *TypeRWWrapper) actualTypeDependenciesRecur(evalType EvaluatedType, used *map[*TypeRWWrapper]bool) {
	if evalType.Index != TypeConstant {
		return
	}
	if str, isStr := w.trw.(*TypeRWStruct); isStr && str.IsWrappingType() {
		eval := str.wr.gen.typesInfo.FieldTypeReduction(evalType.Type, 0)
		str.Fields[0].t.actualTypeDependenciesRecur(eval, used)
		return
	}
	if br, isBr := w.trw.(*TypeRWBrackets); isBr && br.dictLike && len(br.element.t.origTL[0].TemplateArguments) == 1 {
		pairType := br.element.t.trw.(*TypeRWStruct)
		pairEvalType := evalType.Type.Arguments[0]

		if pairEvalType.Type != nil {
			keyValue := pairType.Fields[0]
			valueType := pairType.Fields[1]

			evalKey := br.wr.gen.typesInfo.FieldTypeReduction(pairEvalType.Type, 0)
			evalValue := br.wr.gen.typesInfo.FieldTypeReduction(pairEvalType.Type, 1)

			keyValue.t.actualTypeDependenciesRecur(evalKey, used)
			valueType.t.actualTypeDependenciesRecur(evalValue, used)

			// add dep from builtin_vector
			if !(*used)[w] {
				(*used)[w] = true
			}
		}
		return
	}
	if !(*used)[w] {
		(*used)[w] = true
	}
	for i, arg := range w.arguments {
		if arg.tip != nil {
			arg.tip.actualTypeDependenciesRecur(evalType.Type.Arguments[i], used)
		}
	}
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
			canBare, _ := a.tip.trw.CanBeBareBoxed()
			if !a.bare && canBare { // If it cannot be bare, save on redundant suffix
				b.WriteString("Boxed")
			}
			b.WriteString(tail)
		}
	}
	// We keep compatibility with legacy golang naming
	// This is customization point, generated code should work with whatever naming strategy is selected here
	if len(w.origTL) == 1 && (w.origTL[0].TypeDecl.Name.String() == "_" || w.origTL[0].IsFunction || w.unionParent != nil) {
		return canonicalGoName(w.origTL[0].Construct.Name, insideNamespace), b.String()
	}
	var typeName tlast.Name
	if w.originateFromTL2 {
		typeName.Name = w.tl2Name.Name
		typeName.Namespace = w.tl2Name.Namespace
	} else {
		typeName = w.origTL[0].TypeDecl.Name
	}
	return canonicalGoName(typeName, insideNamespace), b.String()
}

// for golang cycle detection
type DirectImports struct {
	ns         map[*InternalNamespace]struct{}
	importSort bool
}

type CppIncludeInfo struct {
	componentId int
	namespace   string
}

// for C++ includes
type DirectIncludesCPP struct {
	ns map[*TypeRWWrapper]CppIncludeInfo
}

type TypeDefinitionVariation struct {
	NeedBytesVersion bool
}

//func (d DirectIncludesCPP) sortedNames() []string {
//	var sortedNames []string
//	for im := range d.ns { // Imports of this file.
//		sortedNames = append(sortedNames, im)
//	}
//	sort.Strings(sortedNames)
//	return sortedNames
//}

type NamespaceFiles struct {
	Namespace string
	Includes  DirectIncludesCPP
}

func (d DirectIncludesCPP) splitByNamespaces() (result []NamespaceFiles) {
	namespaces := make(map[string]int)

	for file, include := range d.ns {
		ns := include.namespace
		if namespaces[ns] == 0 {
			namespaces[ns] = len(namespaces) + 1
			result = append(result, NamespaceFiles{Namespace: ns, Includes: DirectIncludesCPP{ns: map[*TypeRWWrapper]CppIncludeInfo{}}})
		}
		result[namespaces[ns]-1].Includes.ns[file] = include
	}

	slices.SortFunc(result, func(a, b NamespaceFiles) int {
		return strings.Compare(a.Namespace, b.Namespace)
	})

	return
}

type Pair[L, R any] struct {
	left  L
	right R
}

// not stable
func mapToPairArray[L comparable, R any](m *map[L]R) (res []Pair[L, R]) {
	for k, v := range *m {
		res = append(res, Pair[L, R]{k, v})
	}
	return
}

func (d DirectIncludesCPP) sortedIncludes(componentOrder []int, typeToFile func(wrapper *TypeRWWrapper) string) (result []string) {
	result, _ = d.sortedIncludesWithMap(componentOrder, typeToFile)
	return
}

func (d DirectIncludesCPP) sortedIncludesWithMap(componentOrder []int, typeToFile func(wrapper *TypeRWWrapper) string) (result []string, mapping map[string][]*TypeRWWrapper) {
	includeNamesToTypes := make(map[string]int)
	mapping = make(map[string][]*TypeRWWrapper)

	for tp := range d.ns {
		include := typeToFile(tp)
		mapping[include] = append(mapping[include], tp)
		if _, ok := includeNamesToTypes[include]; !ok {
			includeNamesToTypes[include] = tp.typeComponent
		} else {
			includeNamesToTypes[include] = min(includeNamesToTypes[include], tp.typeComponent)
		}
	}

	includeNamesToTypesList := mapToPairArray(&includeNamesToTypes)
	slices.SortFunc(includeNamesToTypesList, func(a, b Pair[string, int]) int {
		typeDiff := a.right - b.right
		if typeDiff == 0 {
			return strings.Compare(a.left, b.left)
		} else {
			return typeDiff
		}
	})

	for _, includeInfo := range includeNamesToTypesList {
		result = append(result, includeInfo.left)
	}

	return
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
		if w.unionParent == nil || !w.unionParent.IsEnum {
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
	return w.unionParent != nil && w.unionParent.IsEnum
}

func (w *TypeRWWrapper) MarkHasBytesVersion(visitedNodes map[*TypeRWWrapper]bool) bool {
	if visitedNodes[w] {
		return false // We OR results of fields, so if we visited field, and it returned true, this true is already recorded
	}
	visitedNodes[w] = true
	return w.trw.markHasBytesVersion(visitedNodes)
}

func (w *TypeRWWrapper) MarkWriteHasError(visitedNodes map[*TypeRWWrapper]bool) bool {
	if visitedNodes[w] {
		return false // We OR results of fields, so if we visited field, and it returned true, this true is already recorded
	}
	visitedNodes[w] = true
	return w.trw.markWriteHasError(visitedNodes)
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

func (w *TypeRWWrapper) MarkWantsTL2(visitedNodes map[*TypeRWWrapper]bool) {
	if visitedNodes[w] {
		return
	}
	w.wantsTL2 = true
	visitedNodes[w] = true
	w.trw.markWantsTL2(visitedNodes)
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
func (w *TypeRWWrapper) TypeWritingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, bare bool, natArgs []string, ref bool, last bool, needError bool) string {
	bytesVersion = bytesVersion && w.hasBytesVersion
	return w.trw.typeWritingCode(bytesVersion, directImports, ins, val, bare, natArgs, ref, last, needError)
}
func (w *TypeRWWrapper) TypeReadingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, bare bool, natArgs []string, ref bool, last bool) string {
	bytesVersion = bytesVersion && w.hasBytesVersion
	return w.trw.typeReadingCode(bytesVersion, directImports, ins, val, bare, natArgs, ref, last)
}
func (w *TypeRWWrapper) TypeJSONEmptyCondition(bytesVersion bool, val string, ref bool) string {
	bytesVersion = bytesVersion && w.hasBytesVersion
	return w.trw.typeJSONEmptyCondition(bytesVersion, val, ref)
}
func (w *TypeRWWrapper) TypeJSONWritingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, natArgs []string, ref bool, needError bool) string {
	bytesVersion = bytesVersion && w.hasBytesVersion
	return w.trw.typeJSONWritingCode(bytesVersion, directImports, ins, val, natArgs, ref, needError)
}
func (w *TypeRWWrapper) TypeJSONReadingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, jvalue string, val string, natArgs []string, ref bool) string {
	bytesVersion = bytesVersion && w.hasBytesVersion
	return w.trw.typeJSONReadingCode(bytesVersion, directImports, ins, jvalue, val, natArgs, ref)
}

func (w *TypeRWWrapper) TypeJSON2ReadingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, jvalue string, val string, natArgs []string, ref bool) string {
	bytesVersion = bytesVersion && w.hasBytesVersion
	return w.trw.typeJSON2ReadingCode(bytesVersion, directImports, ins, jvalue, val, natArgs, ref)
}

func (w *TypeRWWrapper) TypeJSON2ReadingRequiresContext() bool {
	return w.trw.typeJSON2ReadingRequiresContext()
}

func (w *TypeRWWrapper) IsTrueType() bool {
	structElement, ok := w.trw.(*TypeRWStruct)
	if !ok {
		return false
	}
	return len(structElement.Fields) == 0
}

func (w *TypeRWWrapper) IsFunction() bool {
	structElement, ok := w.trw.(*TypeRWStruct)
	if !ok {
		return false
	}
	return structElement.ResultType != nil
}

func (w *TypeRWWrapper) PHPIsTrueType() bool {
	structElement, ok := w.trw.(*TypeRWStruct)
	if ok {
		unionParent := structElement.PhpConstructorNeedsUnion()
		if unionParent != nil {
			return false
		}
		for _, argument := range w.origTL[0].TemplateArguments {
			if argument.IsNat {
				return false
			}
		}
		// TODO: I HATE THIS SOOO MUCH
		if strings.ToLower(w.tlName.String()) != "true" {
			return false
		}
		return w.IsTrueType()
	}
	return false
}

func (w *TypeRWWrapper) PHPGenerateCode(code *strings.Builder, bytes bool) error {
	return w.trw.PhpGenerateCode(code, bytes)
}

func (w *TypeRWWrapper) PHPTypePathElements() []string {
	_, isStruct := w.trw.(*TypeRWStruct)
	_, isUnion := w.trw.(*TypeRWUnion)
	if !(isStruct || isUnion) {
		return nil
	}

	category := "Types"
	if strct, isStrct := w.trw.(*TypeRWStruct); isStrct && strct.ResultType != nil {
		category = "Functions"
	}
	group := "_common"
	if w.tlName.Namespace != "" {
		group = w.tlName.Namespace
	}
	return []string{"TL", group, category}
}

func (w *TypeRWWrapper) PHPDefaultValue() string {
	core := w.PHPGenCoreType()
	return core.trw.PhpDefaultValue()
}

func (w *TypeRWWrapper) PHPTypePath() string {
	path := w.PHPTypePathElements()
	if path == nil {
		return ""
	} else {
		return strings.Join(path, "\\") + "\\"
	}
}

func (w *TypeRWWrapper) PHPFilePath(bare bool) []string {
	filepathParts := []string{"VK"}
	path := fmt.Sprintf("%s.php", w.trw.PhpClassName(true, bare))
	filepathParts = append(filepathParts, strings.Split(path, "\\")...)
	return filepathParts
}

func (w *TypeRWWrapper) PHPGenCoreType() *TypeRWWrapper {
	if w.PHPUnionParent() == nil {
		struct_, isStruct := w.trw.(*TypeRWStruct)
		if isStruct && struct_.PhpCanBeSimplify() {
			return struct_.Fields[0].t.PHPGenCoreType()
		}
	}
	return w
}

func (w *TypeRWWrapper) PHPUnionParent() *TypeRWWrapper {
	if w.unionParent != nil {
		return w.unionParent.wr
	}
	if strct, isStruct := w.trw.(*TypeRWStruct); isStruct {
		unionParent := strct.PhpConstructorNeedsUnion()
		if unionParent != nil {
			return unionParent
		}
	}
	return nil
}

func (w *TypeRWWrapper) PHPIsBare() bool {
	if strct, isStruct := w.trw.(*TypeRWStruct); isStruct && strct.PhpConstructorNeedsUnion() != nil {
		return false
	}
	return true
}

func (w *TypeRWWrapper) PHPIsPrimitiveType(recursiveCheck bool) bool {
	core := w
	if recursiveCheck {
		core = w.PHPGenCoreType()
	}
	if _, isPrimitive := core.trw.(*TypeRWPrimitive); isPrimitive {
		return true
	}
	if struct_, isStruct := core.trw.(*TypeRWStruct); isStruct {
		if phpIsDictionary(struct_.wr) {
			_, _, _, valueType := isDictionaryElement(struct_.wr)
			return valueType.t.PHPIsPrimitiveType(true)
		}
	}
	if _, isBrackets := core.trw.(*TypeRWBrackets); isBrackets {
		return true
	}
	return false
}

func (w *TypeRWWrapper) PHPNeedsCode() bool {
	if w.PHPTypePath() == "" ||
		w.PHPIsPrimitiveType(true) {
		return false
	}
	// TODO
	if _, ok := w.trw.(*TypeRWStruct); ok &&
		!w.gen.options.InplaceSimpleStructs &&
		strings.HasSuffix(w.tlName.String(), "dictionary") &&
		w.tlName.Namespace == "" {
		return false
	}

	if strct, isStrct := w.trw.(*TypeRWStruct); isStrct {
		unionParent := strct.PhpConstructorNeedsUnion()
		if strct.ResultType == nil && strct.wr.PHPIsTrueType() && unionParent == nil {
			return false
		}
		if strct.ResultType != nil && strct.wr.HasAnnotation("internal") {
			return false
		}
	}
	if w.trw.PhpClassName(false, true) == "" {
		return false
	}
	if w.gen.options.IgnoreUnusedInFunctionsTypes {
		if !w.phpInfo.UsedInFunctions || w.phpInfo.UsedOnlyInInternal {
			return false
		}
	}
	if w.PHPGenCoreType() != w {
		return false
	}
	if w.trw.PhpClassNameReplaced() {
		return false
	}
	if w.tlName.String() == "rpcInvokeReq" {
		return false
	}
	if w.tlName.String() == PHPRPCFunctionMock {
		return false
	}
	if w.tlName.String() == PHPRPCFunctionResultMock {
		return false
	}

	return !w.trw.PhpClassNameReplaced()
}

type TypeArgumentsTree struct {
	name     string
	leaf     bool
	value    *string // value != nil => leaf == True
	children []*TypeArgumentsTree
}

func (t *TypeArgumentsTree) IsEmpty() bool {
	if t.leaf {
		return false
	}
	for _, child := range t.children {
		if child != nil && !child.IsEmpty() {
			return false
		}
	}
	return true
}

func (t *TypeArgumentsTree) EnumerateWithPrefixes() []string {
	const natPrefix = ""
	values := make([]string, 0)
	t.enumerateWithPrefixes(&values, "$")
	values = utils.MapSlice(values, func(s string) string { return natPrefix + s })
	return values
}

func (t *TypeArgumentsTree) EnumerateSubTreeWithPrefixes(childIndex int) []string {
	if !(0 <= childIndex && childIndex < len(t.children)) {
		panic("no such subtree")
	}
	ct := *t
	ct.children = []*TypeArgumentsTree{t.children[childIndex]}
	return ct.EnumerateWithPrefixes()
}

func (t *TypeArgumentsTree) enumerateWithPrefixes(values *[]string, curPrefix string) {
	const delimiter = "_"
	if t.leaf {
		*values = append(*values, curPrefix)
	} else {
		for _, child := range t.children {
			if child != nil {
				prefix := curPrefix + child.name
				if !child.leaf {
					prefix += delimiter
				}
				child.enumerateWithPrefixes(values, prefix)
			}
		}
	}
}

func (t *TypeArgumentsTree) ListAllValues() []string {
	if t == nil {
		return nil
	}
	values := make([]string, 0)
	listAllValuesRecursive(t, &values)
	return values
}

func listAllValuesRecursive(t *TypeArgumentsTree, values *[]string) {
	if t == nil {
		return
	}
	if t.leaf {
		if t.value != nil {
			*values = append(*values, *t.value)
		}
	} else {
		for _, child := range t.children {
			listAllValuesRecursive(child, values)
		}
	}
}

func (t *TypeArgumentsTree) FillAllLeafs() {
	values := t.EnumerateWithPrefixes()
	curIndex := 0
	fillAllLeafsRecursive(t, &curIndex, &values)
}

func (t *TypeArgumentsTree) FillAllLeafsWithValues(values []string) {
	curIndex := 0
	fillAllLeafsRecursive(t, &curIndex, &values)
}

func (t *TypeArgumentsTree) CloneValuesFrom(src *TypeArgumentsTree) {
	if t == nil {
		return
	}
	if t.leaf {
		t.leaf = src.leaf
		t.value = src.value
		t.children = src.children
		return
	}
	for i, child := range t.children {
		if child != nil {
			child.CloneValuesFrom(src.children[i])
		}
	}
}

func fillAllLeafsRecursive(t *TypeArgumentsTree, curIndex *int, values *[]string) {
	if t == nil {
		return
	}
	if t.leaf {
		t.value = &(*values)[*curIndex]
		*curIndex += 1
		return
	}
	for _, child := range t.children {
		fillAllLeafsRecursive(child, curIndex, values)
	}
}

func (w *TypeRWWrapper) PHPGetNatTypeDependenciesDeclAsArray() []string {
	t := TypeArgumentsTree{}
	w.PHPGetNatTypeDependenciesDecl(&t)
	return t.EnumerateWithPrefixes()
}

func (w *TypeRWWrapper) PHPGetNatTypeDependenciesDecl(tree *TypeArgumentsTree) {
	for i, template := range w.origTL[0].TemplateArguments {
		tree.children = append(tree.children, nil)
		actualArg := w.arguments[i]
		if template.IsNat {
			tree.children[i] = &TypeArgumentsTree{}
			tree.children[i].leaf = true
			tree.children[i].name = template.FieldName
		} else {
			tree.children[i] = &TypeArgumentsTree{}
			tree.children[i].leaf = false
			tree.children[i].name = template.FieldName
			actualArg.tip.PHPGetNatTypeDependenciesDecl(tree.children[i])
			if tree.children[i].IsEmpty() {
				tree.children[i] = nil
			}
		}
	}
}

func (w *TypeRWWrapper) PhpIterateReachableTypes(reachableTypes *map[*TypeRWWrapper]bool) {
	if (*reachableTypes)[w] {
		return
	}
	(*reachableTypes)[w] = true
	w.trw.PhpIterateReachableTypes(reachableTypes)
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

func (w *TypeRWWrapper) CPPTypeStringInNamespaceHalfResolved2(bytesVersion bool, typeReduction EvaluatedType) string {
	if typeReduction.Type == nil {
		return typeReduction.TypeVariable
	}
	return w.trw.cppTypeStringInNamespaceHalfResolved2(bytesVersion, typeReduction)
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

func (w *TypeRWWrapper) cppTypeArguments(bytesVersion bool, typeRedaction *TypeReduction) string {
	arguments := make([]string, 0)
	for i, a := range w.arguments {
		evalArg := typeRedaction.Arguments[i]
		if a.isNat {
			if evalArg.Index == NumberConstant {
				arguments = append(arguments, strconv.FormatInt(int64(evalArg.Constant), 10))
			} else if evalArg.Index == NumberVariable && evalArg.VariableActsAsConstant {
				arguments = append(arguments, evalArg.Variable)
			}
		} else {
			if evalArg.Index == TypeVariable {
				arguments = append(arguments, evalArg.TypeVariable)
			} else if evalArg.Index == TypeConstant {
				arguments = append(arguments, a.tip.CPPTypeStringInNamespaceHalfResolved2(bytesVersion, evalArg))
			}
		}
	}
	s := ""
	for i, arg := range arguments {
		if i != 0 {
			s += ", "
		}
		s += arg
	}
	if s != "" {
		s = "<" + s + ">"
	}
	return s
}

func (w *TypeRWWrapper) cppTypeStringInNamespace(bytesVersion bool, hppInc *DirectIncludesCPP, halfResolve bool, halfResolved HalfResolvedArgument) (string, string, string) {
	hppInc.ns[w] = CppIncludeInfo{w.typeComponent, w.tlName.Namespace}
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
				if halfResolve && i < len(halfResolved.Args) {
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
			if halfResolve && i < len(halfResolved.Args) {
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
	markWantsTL2(visitedNodes map[*TypeRWWrapper]bool)
	fillRecursiveUnwrap(visitedNodes map[*TypeRWWrapper]bool)

	FillRecursiveChildren(visitedNodes map[*TypeRWWrapper]int, generic bool)
	AllPossibleRecursionProducers() []*TypeRWWrapper
	AllTypeDependencies(generic, countFunctions bool) []*TypeRWWrapper
	IsWrappingType() bool
	ContainsUnion(visitedNodes map[*TypeRWWrapper]bool) bool

	BeforeCodeGenerationStep1() // during first phase, some wr.trw are nil due to recursive types. So we delay some
	BeforeCodeGenerationStep2() // during second phase, union fields recursive bit is set

	// methods below depend on target language
	fillRecursiveChildren(visitedNodes map[*TypeRWWrapper]bool)
	IsDictKeySafe() (isSafe bool, isString bool) // integers and string are safe, other types no
	CanBeBareBoxed() (canBare bool, canBoxed bool)
	typeString2(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, isLocal bool, skipAlias bool) string
	markHasBytesVersion(visitedNodes map[*TypeRWWrapper]bool) bool
	markWriteHasError(visitedNodes map[*TypeRWWrapper]bool) bool
	typeResettingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, ref bool) string
	typeRandomCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, natArgs []string, ref bool) string
	typeWritingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, bare bool, natArgs []string, ref bool, last bool, needError bool) string
	typeReadingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, bare bool, natArgs []string, ref bool, last bool) string
	typeJSONEmptyCondition(bytesVersion bool, val string, ref bool) string
	typeJSONWritingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, natArgs []string, ref bool, needError bool) string
	typeJSONReadingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, jvalue string, val string, natArgs []string, ref bool) string
	typeJSON2ReadingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, jvalue string, val string, natArgs []string, ref bool) string
	typeJSON2ReadingRequiresContext() bool
	GenerateCode(bytesVersion bool, directImports *DirectImports) string

	TypeRWCPPData
	TypeRWPHPData

	TypeRWTL2
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

	MaskTL2Bit *int

	natArgs      []ActualNatArg
	halfResolved HalfResolvedArgument

	origTL tlast.Field

	// for tl1 - tl2 smooth migration, not affecting tl1 (if it is not empty bitnumber has meaning)
	GenerateLegacySettersForTL2Name string
}

func (f *Field) Bare() bool {
	return f.bare
}

func (f *Field) IsAffectingLocalFieldMasks() bool {
	return f.fieldMask != nil && f.fieldMask.isField
}

func (f *Field) IsAffectedByExternalFieldMask() bool {
	return f.fieldMask != nil && !f.fieldMask.isField
}

func (f *Field) IsTypeDependsFromLocalFields() bool {
	for _, natArg := range f.natArgs {
		if natArg.isField {
			return true
		}
	}
	return false
}

func (f *Field) HasNatArguments() bool {
	return len(f.natArgs) != 0
}

func (f *Field) IsLocalIndependent() bool {
	return !f.IsAffectingLocalFieldMasks() && !f.IsTypeDependsFromLocalFields()
}

// do not generate fields, but affect block position and skip during reading
// TL1: never
// TL2: _:X
func (f *Field) IsTL2Omitted() bool {
	return f.originalName == "_"
}

// generate Set/IsSet with external (TL1) or internal (TL1 & TL2) mask/
// must exactly correspond to migrator logic
// TL1: x:fm.b?true x:fm.b?True
// TL2: x:bit
func (f *Field) IsBit() bool {
	if b, ok := f.t.trw.(*TypeRWBool); ok {
		return b.isTL2 && !b.isTL2Legacy
	}
	return f.fieldMask != nil && (f.t.IsTrueType() && (f.t.tlName.String() == "true" || f.t.tlName.String() == "True"))
}

func (f *Field) TL2MaskForOP(op string) string {
	return fmt.Sprintf("tl2mask%d %s %d", *f.MaskTL2Bit/8, op, 1<<(*f.MaskTL2Bit%8))
}

func wrapWithError(wrap bool, wrappedType string) string {
	if !wrap {
		return wrappedType
	}
	return "(_ " + wrappedType + ", err error)"
}

func formatNatArg(fields []Field, arg ActualNatArg) string {
	if arg.isArith {
		return strconv.FormatUint(uint64(arg.Arith.Res), 10)
	}
	if arg.isField {
		// tl2 case
		if arg.FieldIndex < 0 {
			return fmt.Sprintf("item.mask%d", -arg.FieldIndex)
		}
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
		s.WriteString(fmt.Sprintf(", [[maybe_unused]] uint32_t nat_%s", arg))
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

////for tl2 to tl1 bridge
////in case of formatNatArgs(struct_.Fields, field.natArgs)
//func (f *Field) formatNatArgsOrReturnRandoms(fields []Field, rgName string) []string {
//	result := formatNatArgs(fields, f.natArgs)
//	if len(f.t.NatParams) != len(result) {
//		for i := 0; i < len(f.t.NatParams); i++ {
//			result = append(result, fmt.Sprintf(", basictl.RandomUint(%s)", rgName))
//		}
//	}
//	return result
//}

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

func addAsteriskAndBrackets(ref bool, val string) string {
	return ifString(ref, "(*"+val+")", val)
}

func wrapLast(last bool, code string) string {
	return ifString(last, "return "+code+"", "if err := "+code+"; err != nil { return err }")
}

func wrapLastW(last bool, code string, needError bool) string {
	if needError {
		return ifString(last, "return "+code+"", "if w, err = "+code+"; err != nil { return w, err }")
	} else {
		return ifString(last, "return "+code+"", "w = "+code)
	}
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

func (f *Field) EnsureRecursive(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace) string {
	if !f.recursive {
		return ""
	}
	myType := f.t.TypeString2(bytesVersion, directImports, ins, false, false)
	return fmt.Sprintf(`	if item.%s == nil {
		var value %s
		item.%s = &value
	}
`, f.goName, myType, f.goName)
}
