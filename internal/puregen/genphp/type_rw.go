// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package genphp

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/VKCOM/tl/internal/tlast"
	"github.com/VKCOM/tl/internal/utils"
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

func (d *Deconflicter) removeName(s string) {
	if d.hasConflict(s) {
		delete(d.usedNames, s)
	}
}

func (d *Deconflicter) fillGolangIdentifies() {
	d.deconflictName("Write")
	d.deconflictName("Read")
	d.deconflictName("WriteTL2")
	d.deconflictName("ReadTL2")
}

type TypeRWWrapper struct {
	gen *Gen2 // options.PHP and packages are here

	ns        *Namespace
	trw       TypeRW
	NatParams []string // external params of type Read/Write method, with nat_ prefix

	arguments []ResolvedArgument

	goGlobalName string // globally unique, so could be used also in html anchors, internal C++ function names, etc.
	goLocalName  string
	cppLocalName string

	wantsTL2      bool
	preventUnwrap bool // we can have infinite typedef loop in rare cases

	fileName string

	// php info
	phpInfo PhpClassMeta

	// tl1 info
	tlTag  uint32     // TODO - turn into function
	tlName tlast.Name // TODO - turn into function constructor name or union name for code generation
	origTL []*tlast.Combinator

	unionParent *TypeRWUnion // a bit hackish, but simple
	unionIndex  int
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
	typeName = w.origTL[0].TypeDecl.Name
	return canonicalGoName(typeName, insideNamespace), b.String()
}

type Pair[L, R any] struct {
	left  L
	right R
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

func TypeRWWrapperLessGlobal(a *TypeRWWrapper, b *TypeRWWrapper) int {
	// return stringCompare(a.CanonicalString(), b.CanonicalString()) TODO - better idea after everything is stabilized
	return stringCompare(a.goGlobalName, b.goGlobalName)
}

func (w *TypeRWWrapper) FillRecursiveUnwrap(visitedNodes map[*TypeRWWrapper]bool) {
	if visitedNodes[w] {
		return
	}
	visitedNodes[w] = true
	w.trw.fillRecursiveUnwrap(visitedNodes)
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
		!w.gen.options.PHP.InplaceSimpleStructs &&
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
	if w.gen.options.PHP.IgnoreUnusedInFunctionsTypes {
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

type TypeRW interface {
	fillRecursiveUnwrap(visitedNodes map[*TypeRWWrapper]bool)

	BeforeCodeGenerationStep1()

	IsDictKeySafe() (isSafe bool, isString bool) // integers and string are safe, other types no
	CanBeBareBoxed() (canBare bool, canBoxed bool)

	TypeRWPHPData
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
}

func (f *Field) Bare() bool {
	return f.bare
}

func (f *Field) IsAffectedByLocalFieldMask() bool {
	return f.fieldMask != nil && f.fieldMask.isField
}

func (f *Field) IsAffectedByExternalFieldMask() bool {
	return f.fieldMask != nil && !f.fieldMask.isField
}

func (f *Field) HasNatArguments() bool {
	return len(f.natArgs) != 0
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
		return b.isTL2 && b.isBit
	}
	return f.fieldMask != nil && (f.t.IsTrueType() && (f.t.tlName.String() == "true" || f.t.tlName.String() == "True"))
}

func (f *Field) TL2MaskForOP(op string) string {
	return fmt.Sprintf("tl2mask%d %s %d", *f.MaskTL2Bit/8, op, 1<<(*f.MaskTL2Bit%8))
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

func ifString(value bool, t string, f string) string {
	if value {
		return t
	}
	return f
}

func ToUpperFirst(str string) string {
	return utils.ToUpperFirst(str)
}

func ToLowerFirst(str string) string {
	return utils.ToLowerFirst(str)
}
