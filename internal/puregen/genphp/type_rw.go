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

	"github.com/VKCOM/tl/internal/pure"
	"github.com/VKCOM/tl/internal/tlast"
	"github.com/VKCOM/tl/internal/utils"
)

// During recursive generation, we store wrappers to type when they are needed, so that
// we can generate actual types later, when all references to wrappers are set
// also wrapper stores common information

type Deconflicter struct {
	usedNames map[string]bool
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

func (d *Deconflicter) DeconflictName(s string) string {
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

type TypeRWWrapper struct {
	pureType pure.TypeInstance

	gen *Gen2 // options.PHP and packages are here

	ns  *Namespace
	trw TypeRW

	fileName string

	// php info
	phpInfo PhpClassMeta

	unionParent *TypeRWUnion // a bit hackish, but simple
	unionIndex  int
}

func (w *TypeRWWrapper) TLTag() uint32 {
	return w.pureType.Common().TLTag()
}

func (w *TypeRWWrapper) TLName() tlast.TL2TypeName {
	return w.pureType.Common().TLName()
}

func (w *TypeRWWrapper) HasAnnotation(str string) bool {
	return w.pureType.KernelType().HasAnnotation(str)
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

func stringCompare(a string, b string) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
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
		for _, argument := range w.pureType.KernelType().Templates() {
			if argument.Category.IsNat() {
				return false
			}
		}
		// TODO: I HATE THIS SOOO MUCH
		if strings.ToLower(w.TLName().String()) != "true" {
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
	if w.TLName().Namespace != "" {
		group = w.TLName().Namespace
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
			valueType := phpDictionaryElement(struct_.wr)
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
		strings.HasSuffix(w.TLName().String(), "dictionary") &&
		w.TLName().Namespace == "" {
		return false
	}

	if PHPIsDict(w.pureType.KernelType()) {
		return false
	}

	if w.pureType.KernelType().CanonicalName().String() == "ReqResult" {
		_, isUnion := w.pureType.(*pure.TypeInstanceUnion)
		if isUnion {
			return false
		}
	}

	if w.pureType.KernelType().CanonicalName().String() == "RpcReqResult" {
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
	if w.TLName().String() == "rpcInvokeReq" {
		return false
	}
	if w.TLName().String() == PHPRPCFunctionMock {
		return false
	}
	if w.TLName().String() == PHPRPCFunctionResultMock {
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
	rt := w.pureType.Common().ResolvedType()
	if rt.BracketType == nil {
		for i, template := range w.pureType.KernelType().Templates() {
			tree.children = append(tree.children, nil)
			if template.Category.IsNat() {
				tree.children[i] = &TypeArgumentsTree{}
				tree.children[i].leaf = true
				tree.children[i].name = template.Name
			} else {
				tree.children[i] = &TypeArgumentsTree{}
				tree.children[i].leaf = true
				tree.children[i].name = template.Name

				tp := rt.SomeType.Arguments[i].Type
				if tp.String() != "*" {
					tree.children[i].leaf = false
					tip, _ := w.gen.getTypeWrapperMust(rt.SomeType.Arguments[i].Type)
					tip.PHPGetNatTypeDependenciesDecl(tree.children[i])
					if tree.children[i].IsEmpty() {
						tree.children[i] = nil
					}
				}
			}
		}
	} else {
		if rt.BracketType.HasIndex {
			i := len(tree.children)
			tree.children = append(tree.children, nil)

			tree.children[i] = &TypeArgumentsTree{}
			tree.children[i].leaf = true
			tree.children[i].name = "N"

			tp := rt.BracketType.IndexType
			if tp.String() != "*" {
				tree.children[i].leaf = false
				tip, _ := w.gen.getTypeWrapperMust(tp.Type)
				tip.PHPGetNatTypeDependenciesDecl(tree.children[i])
				if tree.children[i].IsEmpty() {
					tree.children[i] = nil
				}
			}
		}

		i := len(tree.children)
		tree.children = append(tree.children, nil)

		tree.children[i] = &TypeArgumentsTree{}
		tree.children[i].leaf = true
		tree.children[i].name = "X"

		tp := rt.BracketType.ArrayType
		if tp.String() != "*" {
			tree.children[i].leaf = false
			tip, _ := w.gen.getTypeWrapperMust(tp)
			tip.PHPGetNatTypeDependenciesDecl(tree.children[i])
			if tree.children[i].IsEmpty() {
				tree.children[i] = nil
			}
		}
	}
	//rt := w.pureType.Common().ResolvedType()
	//
	//handleTemplate := func(ta tlast.TL2TypeArgument) {
	//	subTree := &TypeArgumentsTree{}
	//
	//	treeIndex := len(tree.children)
	//	tree.children = append(tree.children, subTree)
	//	if ta.IsNumber {
	//		subTree.leaf = true
	//		subTree.name = ta.OriginalArgumentName
	//	} else {
	//		tip, _ := w.gen.getTypeWrapperMust(ta.Type)
	//		subTree.leaf = false
	//		subTree.name = ta.OriginalArgumentName
	//		tip.PHPGetNatTypeDependenciesDecl(subTree)
	//		if subTree.IsEmpty() {
	//			tree.children[treeIndex] = nil
	//		}
	//	}
	//}
	//
	//if rt.BracketType == nil {
	//	for _, template := range w.pureType.Common().ResolvedType().SomeType.Arguments {
	//		handleTemplate(template)
	//	}
	//} else {
	//	if rt.BracketType.HasIndex {
	//		handleTemplate(rt.BracketType.IndexType)
	//	}
	//	ta := rt.BracketType.ArrayType
	//	subTree := &TypeArgumentsTree{}
	//
	//	treeIndex := len(tree.children)
	//	tree.children = append(tree.children, subTree)
	//
	//	tip, _ := w.gen.getTypeWrapperMust(ta)
	//	subTree.leaf = false
	//	subTree.name = "X"
	//	tip.PHPGetNatTypeDependenciesDecl(subTree)
	//	if subTree.IsEmpty() {
	//		tree.children[treeIndex] = nil
	//	}
	//}
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
	pureField pure.Field

	t         *TypeRWWrapper
	goName    string
	recursive bool
}

func (f *Field) Bare() bool {
	return f.pureField.Bare()
}

// do not generate fields, but affect block position and skip during reading
// TL1: never
// TL2: _:X
func (f *Field) IsTL2Omitted() bool {
	return f.pureField.Name() == "_"
}

// generate Set/IsSet with external (TL1) or internal (TL1 & TL2) mask/
// must exactly correspond to migrator logic
// TL1: x:fm.b?true x:fm.b?True
// TL2: x:bit
func (f *Field) IsBit() bool {
	if b, ok := f.t.trw.(*TypeRWBool); ok {
		return b.isTL2 && b.isBit
	}
	return f.pureField.FieldMask() != nil && (f.t.IsTrueType() && (f.t.TLName().String() == "true" || f.t.TLName().String() == "True"))
}

func (f *Field) TL2MaskForOP(op string) string {
	return fmt.Sprintf("tl2mask%d %s %d", *f.pureField.MaskTL2Bit()/8, op, 1<<(*f.pureField.MaskTL2Bit()%8))
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
