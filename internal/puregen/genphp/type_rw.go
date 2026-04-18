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

	gen *genphp // options.PHP and packages are here

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
		for _, argument := range w.pureType.KernelType().TemplateArguments() {
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

func (w *TypeRWWrapper) PHPGetNatTypeDependenciesDeclAsArray() []string {
	return w.pureType.Common().NatParams()
}

func (w *TypeRWWrapper) PhpIterateReachableTypes(reachableTypes *map[*TypeRWWrapper]bool) {
	if (*reachableTypes)[w] {
		return
	}
	(*reachableTypes)[w] = true
	w.trw.PhpIterateReachableTypes(reachableTypes)
}

type TypeRW interface {
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
	return f.pureField.IsBit()
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
