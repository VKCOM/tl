// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package gengo

import (
	"cmp"
	"fmt"
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

	goCanonicalName tlast.Name // name element for names below and template full names

	goGlobalName string // globally unique
	goLocalName  string

	wantsBytesVersion bool
	wantsTL2          bool
	preventUnwrap     bool // we can have infinite typedef loop in rare cases

	hasBytesVersion        bool
	hasErrorInWriteMethods bool

	fileName string
	// otherwise order-dependent
	fileNameOverride *TypeRWWrapper

	originateFromTL2 bool

	// tl1 info
	tlTag  uint32
	tlName tlast.Name // constructor name, except for unions
	//origTL []*tlast.Combinator

	// TODO - move into TypeRWStruct
	unionParent *TypeRWUnion // a bit hackish, but simple
	unionIndex  int

	WrLong        *TypeRWWrapper // long transitioning code
	WrWithoutLong *TypeRWWrapper // long transitioning code

	// tl2 info (if union is not nil otherwise check there)
	//tl2Name              tlast.TL2TypeName
	//tl2Origin            *tlast.TL2Combinator
	//tl2IsResult          bool
	//tl2IsBuiltinBrackets bool
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

// Those have unique structure fully defined by the magic.
// items with condition len(w.NatParams) == 0 could be serialized independently, but if there is several type instantiations,
// they could not be distinguished by the magic. For example vector<int> and vector<long>.
//func (w *TypeRWWrapper) IsTopLevel() bool {
//	if w.originateFromTL2 {
//		if w.unionParent == nil {
//			if w.tl2IsResult {
//				return false
//			}
//			if w.tl2IsBuiltinBrackets {
//				return false
//			}
//			if w.tl2Origin != nil {
//				if w.tl2Origin.IsFunction {
//					return true
//				}
//				return len(w.tl2Origin.TypeDecl.TemplateArguments) == 0
//			}
//			return false
//		} else {
//			return false
//		}
//	}
//	if len(w.origTL) == 0 {
//		// fmt.Printf("Empty origTL for %s %v\n", w.goGlobalName, w.pureType.KernelType().)
//		return false
//	}
//	return len(w.origTL[0].TemplateArguments) == 0
//}

//func (w *TypeRWWrapper) CanonicalStringTop() string {
//	return w.CanonicalString(len(w.origTL) <= 1) // single constructors, arrays and primitives are naturally bare, unions are naturally boxed
//}
//
//func (w *TypeRWWrapper) CanonicalString(bare bool) string {
//	var s strings.Builder
//	if w.originateFromTL2 {
//		if w.unionParent == nil {
//			if w.tl2IsResult {
//				s.WriteString(w.tl2Origin.FuncDecl.Name.String() + "__Result")
//			} else if w.tl2IsBuiltinBrackets {
//				s.WriteString("__builtin_brackets")
//			} else if w.tl2Origin != nil {
//				s.WriteString(w.tl2Origin.TypeDecl.Name.String())
//			} else {
//				s.WriteString(w.tl2Name.String())
//			}
//		} else {
//			originType := w.unionParent.wr.tl2Origin
//			if w.unionParent.wr.tl2IsResult {
//				s.WriteString(originType.FuncDecl.Name.String() + "__Result" + originType.FuncDecl.ReturnType.StructType.UnionType.Variants[w.unionIndex].Name)
//			} else {
//				s.WriteString(originType.TypeDecl.Name.String() + originType.TypeDecl.Type.StructType.UnionType.Variants[w.unionIndex].Name)
//			}
//		}
//	} else {
//		if len(w.origTL) > 1 {
//			if bare {
//				panic("CanonicalString of bare union")
//			}
//			w.origTL[0].TypeDecl.Name.WriteString(&s)
//		} else if len(w.origTL) == 1 {
//			if bare {
//				w.origTL[0].Construct.Name.WriteString(&s)
//			} else {
//				w.origTL[0].TypeDecl.Name.WriteString(&s)
//			}
//		} else {
//			panic("all builtins are parsed from TL text, so must have exactly one constructor")
//		}
//	}
//	if len(w.arguments) == 0 {
//		return s.String()
//	}
//	s.WriteByte('<')
//	for i, a := range w.arguments {
//		// fieldName := t.origTL[0].TemplateArguments[i].FieldName // arguments must be the same for all union elements
//		if i != 0 {
//			s.WriteByte(',')
//		}
//		if a.isNat {
//			if a.isArith {
//				s.WriteString(strconv.FormatUint(uint64(a.Arith.Res), 10))
//			} else {
//				s.WriteString("#") // TODO - write fieldName here if special argument to function is set
//			}
//		} else {
//			s.WriteString(a.tip.CanonicalString(a.bare))
//		}
//	}
//	s.WriteByte('>')
//	return s.String()
//}

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

func (w *TypeRWWrapper) resolvedT2GoNameTail(insideNamespace string) string {
	b := strings.Builder{}
	for _, arg := range w.pureType.Common().ResolvedType().Args {
		if arg.IsArith {
			b.WriteString(strconv.FormatUint(uint64(arg.Arith.Res), 10))
			continue
		}
		if arg.T.String() == "*" {
			continue
		}
		ref, fieldBare, err := w.gen.kernel.GetInstanceTL1(arg.T)
		if err != nil {
			panic(fmt.Errorf("internal error: cannot get type of argument %s: %w", arg.T, err))
		}
		fieldType, err := w.gen.getType(ref)
		if err != nil {
			panic(fmt.Errorf("internal error: cannot get type of argument %s: %w", arg.T, err))
		}
		head, tail := fieldType.resolvedT2GoName(insideNamespace)
		b.WriteString(head)
		if head != "Bool" && !fieldBare && !fieldType.pureType.BoxedOnly() {
			// If it cannot be bare, save on redundant suffix
			// Bool is exception, because it is bare in TL2, but boxed in TL1
			b.WriteString("Boxed")
		}
		b.WriteString(tail)
	}
	return b.String()
}

func (w *TypeRWWrapper) resolvedT2GoName(insideNamespace string) (head, tail string) {
	//if w.pureType.CanonicalName() == "Vector<uint32>" {
	//	fmt.Printf("aga")
	//}
	tail = w.resolvedT2GoNameTail(insideNamespace)
	// We keep compatibility with legacy golang naming
	// This is customization point, generated code should work with whatever naming strategy is selected here
	return canonicalGoName(w.goCanonicalName, insideNamespace), tail
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
	return cmp.Compare(an, bn)
}

func TypeRWWrapperLessGlobal(a *TypeRWWrapper, b *TypeRWWrapper) int {
	// return stringCompare(a.CanonicalString(), b.CanonicalString()) TODO - better idea after everything is stabilized
	return cmp.Compare(a.goGlobalName, b.goGlobalName)
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

func (w *TypeRWWrapper) JSONHelpString() string {
	return w.pureType.CanonicalName()
}
