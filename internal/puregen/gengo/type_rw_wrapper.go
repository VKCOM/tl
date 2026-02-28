// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package gengo

import (
	"cmp"
	"strconv"
	"strings"

	"github.com/vkcom/tl/internal/pure"
	"github.com/vkcom/tl/internal/tlast"
)

// During recursive generation, we store wrappers to type when they are needed, so that
// we can generate actual types later, when all references to wrappers are set
// also wrapper stores common information

type Deconflicter struct {
	usedNames map[string]bool
}

//func (d *Deconflicter) hasConflict(s string) bool {
//	_, ok := d.usedNames[s]
//	return ok
//}

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

func (d *Deconflicter) fillGolangIdentifies() {
	d.deconflictName("Write")
	d.deconflictName("Read")
	d.deconflictName("WriteTL2")
	d.deconflictName("ReadTL2")
}

type TypeRWWrapper struct {
	gen *genGo // options and packages are here

	pureType pure.TypeInstance

	ns        *Namespace
	ins       *InternalNamespace
	trw       TypeRW
	NatParams []string // external params of type Read/Write method

	goCanonicalName tlast.TL2TypeName // name element for names below and template full names

	goGlobalName string // globally unique
	goLocalName  string

	wantsBytesVersion bool
	hasRepairMasks    bool
	preventUnwrap     bool // we can have infinite typedef loop in rare cases

	hasBytesVersion        bool
	hasErrorInWriteMethods bool

	fileName string
	// otherwise order-dependent
	fileNameOverride *TypeRWWrapper

	originateFromTL2 bool

	tlTag  uint32
	tlName tlast.TL2TypeName

	// TODO - move into TypeRWStruct
	unionParent *TypeRWUnion // a bit hackish, but simple
	unionIndex  int

	WrLong        *TypeRWWrapper // long transitioning code
	WrWithoutLong *TypeRWWrapper // long transitioning code
}

func TypeRWWrapperLessLocal(a *TypeRWWrapper, b *TypeRWWrapper) int {
	return cmp.Compare(a.goLocalName, b.goLocalName)
}

func TypeRWWrapperLessGlobal(a *TypeRWWrapper, b *TypeRWWrapper) int {
	return cmp.Compare(a.goGlobalName, b.goGlobalName)
}

func (wr *TypeRWWrapper) HasTL2() bool {
	return wr.pureType.Common().HasTL2()
}

func (wr *TypeRWWrapper) FileName() string {
	if wr.fileNameOverride != nil {
		return wr.fileNameOverride.FileName()
	}
	return wr.fileName
}

func (wr *TypeRWWrapper) Namespace() string {
	return wr.tlName.Namespace
}

func (w *TypeRWWrapper) HasAnnotation(str string) bool {
	return w.pureType.KernelType().HasAnnotation(str)
}

func (w *TypeRWWrapper) AnnotationsMask() uint32 {
	var mask uint32
	for bit, v := range w.gen.kernel.AllAnnotations() {
		if w.HasAnnotation(v) {
			mask |= (1 << bit)
		}
	}
	return mask
}

func (w *TypeRWWrapper) DoArgumentsContainUnionTypes() bool {
	if w, ok := w.trw.(*TypeRWStruct); ok && w.ResultType != nil {
		return w.wr.containsUnion(map[*TypeRWWrapper]bool{})
	}
	return false
}

func (w *TypeRWWrapper) DoesReturnTypeContainUnionTypes() bool {
	if w, ok := w.trw.(*TypeRWStruct); ok && w.ResultType != nil {
		return w.ResultType.containsUnion(map[*TypeRWWrapper]bool{})
	}
	return false
}

func (w *TypeRWWrapper) containsUnion(visitedNodes map[*TypeRWWrapper]bool) bool {
	if _, ok := visitedNodes[w]; !ok {
		visitedNodes[w] = false
		return w.trw.ContainsUnion(visitedNodes)
	} else {
		return false
	}
}

func (w *TypeRWWrapper) resolvedT2GoNameArg(b *strings.Builder, arg tlast.TL2TypeArgument, insideNamespace string) {
	if arg.IsNumber {
		b.WriteString(strconv.FormatUint(uint64(arg.Number), 10))
		return
	}
	if arg.Type.String() == "*" {
		return
	}
	fieldType, fieldBare := w.gen.getTypeMust(arg.Type)
	head, tail := fieldType.resolvedT2GoName(insideNamespace)
	b.WriteString(head)
	if head != "Bool" && !fieldBare && !fieldType.pureType.BoxedOnly() {
		// If it cannot be bare, save on redundant suffix
		// Bool is exception, because it is bare in TL2, but boxed in TL1
		b.WriteString("Boxed")
	}
	b.WriteString(tail)
}

func (w *TypeRWWrapper) resolvedT2GoName(insideNamespace string) (head, tail string) {
	b := strings.Builder{}
	rt := w.pureType.Common().ResolvedType()
	if br := rt.BracketType; br != nil {
		if br.HasIndex {
			if br.IndexType.IsNumber || br.IndexType.Type.String() == "*" {
				head = "BuiltinTuple"
				w.resolvedT2GoNameArg(&b, br.IndexType, insideNamespace)
				//if br.IndexType.IsNumber {
				//	b.WriteString(strconv.FormatUint(uint64(br.IndexType.Number), 10))
				//}
			} else {
				head = "BuiltinDict"
				w.resolvedT2GoNameArg(&b, br.IndexType, insideNamespace)
			}
		} else {
			head = "BuiltinVector"
		}
		w.resolvedT2GoNameArg(&b, tlast.TL2TypeArgument{Type: br.ArrayType}, insideNamespace)
	} else {
		head = canonicalGoName2(w.goCanonicalName, insideNamespace)
		for _, arg := range rt.SomeType.Arguments {
			w.resolvedT2GoNameArg(&b, arg, insideNamespace)
		}
	}
	tail = b.String()
	// We keep compatibility with legacy golang naming
	// This is customization point, generated code should work with whatever naming strategy is selected here
	return
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

func (w *TypeRWWrapper) MarkHasRepairMasks(visitedNodes map[*TypeRWWrapper]bool) bool {
	if visitedNodes[w] {
		return false // We OR results of fields, so if we visited field, and it returned true, this true is already recorded
	}
	visitedNodes[w] = true
	return w.trw.markHasRepairMasks(visitedNodes)
}

func (w *TypeRWWrapper) MarkWriteHasError(visitedNodes map[*TypeRWWrapper]bool) bool {
	if visitedNodes[w] {
		return false // We OR results of fields, so if we visited field, and it returned true, this true is already recorded
	}
	visitedNodes[w] = true
	return w.trw.markWriteHasError(visitedNodes)
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
func (w *TypeRWWrapper) TypeRepairMasksCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, natArgs []string, ref bool) string {
	bytesVersion = bytesVersion && w.hasBytesVersion
	return w.trw.typeRepairMasksCode(bytesVersion, directImports, ins, val, natArgs, ref)
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
