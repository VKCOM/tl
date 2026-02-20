// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package gengo

//{% import "github.com/vkcom/tl/internal/utils" %}
//{% import "strings" %}
//{% import "sort" %}

import (
	"fmt"
	"strings"

	"github.com/vkcom/tl/internal/pure"
)

type TypeRWStruct struct {
	wr             *TypeRWWrapper
	pureTypeStruct *pure.TypeInstanceStruct
	Fields         []Field

	ResultType    *TypeRWWrapper
	ResultNatArgs []pure.ActualNatArg

	fieldsDec  Deconflicter // TODO - add all generated methods here
	setNames   []string     // method names should be the same for bytes and normal versions, so we remember them here
	clearNames []string
	isSetNames []string
}

var _ TypeRW = &TypeRWStruct{}

func (trw *TypeRWStruct) isTypeDef() bool {
	return len(trw.Fields) == 1 && trw.Fields[0].originalName == "" && trw.Fields[0].fieldMask == nil && !trw.Fields[0].recursive
}

func (trw *TypeRWStruct) isUnwrapType() bool {
	if !trw.isTypeDef() || trw.wr.preventUnwrap {
		return false
	}
	return trw.pureTypeStruct.IsUnwrap()
}

func (trw *TypeRWStruct) typeString2(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, isLocal bool, skipAlias bool) string {
	if !skipAlias && trw.isUnwrapType() {
		return trw.Fields[0].t.TypeString2(bytesVersion, directImports, ins, isLocal, skipAlias)
	}
	if isLocal {
		return addBytes(trw.wr.goLocalName, bytesVersion)
	}
	return trw.wr.ins.Prefix(directImports, ins) + addBytes(trw.wr.goGlobalName, bytesVersion)
}

func (trw *TypeRWStruct) markHasBytesVersion(visitedNodes map[*TypeRWWrapper]bool) bool {
	result := false
	for _, f := range trw.Fields {
		result = result || f.t.MarkHasBytesVersion(visitedNodes)
	}
	if trw.ResultType != nil {
		result = result || trw.ResultType.MarkHasBytesVersion(visitedNodes)
	}
	return result
}

func (trw *TypeRWStruct) markHasRepairMasks(visitedNodes map[*TypeRWWrapper]bool) bool {
	result := len(trw.AllNewTL2Masks()) != 0 && !trw.wr.originateFromTL2
	for _, f := range trw.Fields {
		result = result || f.t.MarkHasRepairMasks(visitedNodes)
	}
	// result type does not affect this
	return result
}

func (trw *TypeRWStruct) markWriteHasError(visitedNodes map[*TypeRWWrapper]bool) bool {
	result := false
	for _, f := range trw.Fields {
		result = result || f.t.MarkWriteHasError(visitedNodes)
	}
	if trw.ResultType != nil {
		result = result || trw.ResultType.MarkWriteHasError(visitedNodes)
	}
	return result
}

func (trw *TypeRWStruct) fillRecursiveUnwrap(visitedNodes map[*TypeRWWrapper]bool) {
	if !trw.isTypeDef() {
		return
	}
	trw.Fields[0].t.FillRecursiveUnwrap(visitedNodes)
}

func (trw *TypeRWStruct) markWantsBytesVersion(visitedNodes map[*TypeRWWrapper]bool) {
	for _, f := range trw.Fields {
		f.t.MarkWantsBytesVersion(visitedNodes)
	}
	if trw.ResultType != nil {
		trw.ResultType.MarkWantsBytesVersion(visitedNodes)
	}
}

func (trw *TypeRWStruct) markWantsTL2(visitedNodes map[*TypeRWWrapper]bool) {
	for _, f := range trw.Fields {
		f.t.MarkWantsTL2(visitedNodes)
	}
	if trw.ResultType != nil {
		trw.ResultType.MarkWantsTL2(visitedNodes)
	}
}

func (trw *TypeRWStruct) ContainsUnion(visitedNodes map[*TypeRWWrapper]bool) bool {
	for _, f := range trw.Fields {
		if f.t.containsUnion(visitedNodes) {
			return true
		}
	}
	return false
}

func (trw *TypeRWStruct) FillRecursiveChildren(visitedNodes map[*TypeRWWrapper]int, generic bool) {
	if visitedNodes[trw.wr] != 0 {
		return
	}
	visitedNodes[trw.wr] = 1
}

func (trw *TypeRWStruct) BeforeCodeGenerationStep1() {
	for i, f := range trw.Fields {
		visitedNodes := map[*TypeRWWrapper]bool{}
		f.t.trw.fillRecursiveChildren(visitedNodes)
		trw.Fields[i].recursive = visitedNodes[trw.wr]
	}

	trw.setNames = make([]string, len(trw.Fields))
	trw.clearNames = make([]string, len(trw.Fields))
	trw.isSetNames = make([]string, len(trw.Fields))
}

func (trw *TypeRWStruct) GetAllLocallyAffectedByTrueTypeFieldMasks() []Field {
	nats := make([]Field, 0)
	containingNats := make(map[int]bool)

	for _, field := range trw.Fields {
		if field.IsAffectingLocalFieldMasks() && field.t.IsTrueType() {
			index := field.fieldMask.FieldIndex()
			if _, contains := containingNats[index]; !contains {
				nats = append(nats, trw.Fields[index])
				containingNats[index] = true
			}
		}
	}

	return nats
}

func (trw *TypeRWStruct) GetAllLocallyAffectedFieldMasks() []Field {
	nats := make([]Field, 0)
	containingNats := make(map[int]bool)

	for _, field := range trw.Fields {
		if field.IsAffectingLocalFieldMasks() {
			index := field.fieldMask.FieldIndex()
			if _, contains := containingNats[index]; !contains {
				nats = append(nats, trw.Fields[index])
				containingNats[index] = true
			}
		}
	}

	return nats
}

// AllAffectedFieldMasks f must be from trw.Fields
func (trw *TypeRWStruct) AllAffectedFieldMasks(f Field) (nats []Field, bits []uint32) {
	curField := f
	for curField.IsAffectingLocalFieldMasks() {
		if curField.fieldMask.FieldIndex() < 0 {
			return
		}
		ancestor := trw.Fields[curField.fieldMask.FieldIndex()]
		nats = append(nats, ancestor)
		bits = append(bits, curField.BitNumber)
		curField = ancestor
	}

	return
}

func (trw *TypeRWStruct) BeforeCodeGenerationStep2() {
	//if trw.wr.gen.options.Language == "cpp" { // TODO - temporary solution to benchmark combined tl
	//	var nf []Field
	//	for _, f := range trw.Fields {
	//		if !f.recursive {
	//			nf = append(nf, f)
	//  // panic("recursive field in union " + trw.wr.tlName.String())
	//}
	//}
	//trw.Fields = nf
	//}
}

func (trw *TypeRWStruct) fillRecursiveChildren(visitedNodes map[*TypeRWWrapper]bool) {
	for _, f := range trw.Fields {
		if !f.recursive {
			f.t.FillRecursiveChildren(visitedNodes)
		}
	}
}

func (trw *TypeRWStruct) typeResettingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, ref bool) string {
	if trw.isUnwrapType() {
		return trw.Fields[0].t.TypeResettingCode(bytesVersion, directImports, ins, val, ref)
	}
	return fmt.Sprintf("%s.Reset()", val)
}

func (trw *TypeRWStruct) typeRandomCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, natArgs []string, ref bool) string {
	if trw.isUnwrapType() {
		return trw.Fields[0].t.TypeRandomCode(bytesVersion, directImports, ins, val, trw.pureTypeStruct.ReplaceUnwrapArgs(natArgs), ref)
	}
	return fmt.Sprintf("%s.FillRandom(rg %s)", val, joinWithCommas(natArgs))
}

func (trw *TypeRWStruct) typeRepairMasksCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, natArgs []string, ref bool) string {
	if trw.isUnwrapType() {
		return trw.Fields[0].t.TypeRepairMasksCode(bytesVersion, directImports, ins, val, trw.pureTypeStruct.ReplaceUnwrapArgs(natArgs), ref)
	}
	return fmt.Sprintf("%s.RepairMasks(%s)", val, strings.Join(natArgs, ","))
}

func (trw *TypeRWStruct) typeWritingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, bare bool, natArgs []string, ref bool, last bool, needError bool) string {
	if trw.isUnwrapType() {
		prefix := ""
		if !bare {
			prefix = fmt.Sprintf("w = basictl.NatWrite(w, 0x%x)\n", trw.wr.tlTag)
		}
		return prefix + trw.Fields[0].t.TypeWritingCode(bytesVersion, directImports, ins, val, trw.Fields[0].Bare(), trw.pureTypeStruct.ReplaceUnwrapArgs(natArgs), ref, last, needError)
		// was
		// goName := addBytes(trw.goGlobalName, bytesVersion)
		// return wrapLastW(last, fmt.Sprintf("(*%s)(%s).Write%s(w%s)", trw.wr.ins.Prefix(ins)+goName, addAmpersand(ref, val), addBare(bare), joinWithCommas(natArgs)))
	}
	return wrapLastW(last, fmt.Sprintf("%s.Write%s(w %s)", val, addBare(bare), joinWithCommas(natArgs)), needError)
}

func (trw *TypeRWStruct) typeReadingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, bare bool, natArgs []string, ref bool, last bool) string {
	if trw.isUnwrapType() {
		prefix := ""
		if !bare {
			prefix = fmt.Sprintf("if w, err = basictl.NatReadExactTag(w, 0x%x); err != nil {\nreturn w, err\n}\n", trw.wr.tlTag)
		}
		return prefix + trw.Fields[0].t.TypeReadingCode(bytesVersion, directImports, ins, val, trw.Fields[0].Bare(), trw.pureTypeStruct.ReplaceUnwrapArgs(natArgs), ref, last)
		// was
		// goName := addBytes(trw.goGlobalName, bytesVersion)
		// return wrapLastW(last, fmt.Sprintf("(*%s)(%s).Read%s(w%s)", trw.wr.ins.Prefix(ins)+goName, addAmpersand(ref, val), addBare(bare), joinWithCommas(natArgs)))
	}
	return wrapLastW(last, fmt.Sprintf("%s.Read%s(w %s)", val, addBare(bare), joinWithCommas(natArgs)), true)
}

func (trw *TypeRWStruct) typeJSONEmptyCondition(bytesVersion bool, val string, ref bool) string {
	if trw.isTypeDef() {
		return trw.Fields[0].t.TypeJSONEmptyCondition(bytesVersion, val, ref)
	}
	return ""
}

func (trw *TypeRWStruct) typeJSONWritingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, natArgs []string, ref bool, needError bool) string {
	if trw.isUnwrapType() {
		return trw.Fields[0].t.TypeJSONWritingCode(bytesVersion, directImports, ins, val, trw.pureTypeStruct.ReplaceUnwrapArgs(natArgs), ref, needError)
	}
	if needError {
		return fmt.Sprintf("if w, err = %s.WriteJSONOpt(tctx, w %s); err != nil { return w, err }", val, joinWithCommas(natArgs))
	} else {
		return fmt.Sprintf("w = %s.WriteJSONOpt(tctx, w %s)", val, joinWithCommas(natArgs))
	}
}

func (trw *TypeRWStruct) typeJSONReadingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, jvalue string, val string, natArgs []string, ref bool) string {
	if trw.isUnwrapType() {
		return trw.Fields[0].t.TypeJSONReadingCode(bytesVersion, directImports, ins, jvalue, val, trw.pureTypeStruct.ReplaceUnwrapArgs(natArgs), ref)
	}
	return fmt.Sprintf("if err := %s.ReadJSONLegacy(legacyTypeNames, %s %s); err != nil { return err }", val, jvalue, joinWithCommas(natArgs))
}

func (trw *TypeRWStruct) typeJSON2ReadingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, jvalue string, val string, natArgs []string, ref bool) string {
	if trw.isUnwrapType() {
		return trw.Fields[0].t.TypeJSON2ReadingCode(bytesVersion, directImports, ins, jvalue, val, trw.pureTypeStruct.ReplaceUnwrapArgs(natArgs), ref)
	}
	return fmt.Sprintf("if err := %s.ReadJSONGeneral(tctx, %s %s); err != nil { return err }", val, jvalue, joinWithCommas(natArgs))
}

func (trw *TypeRWStruct) typeJSON2ReadingRequiresContext() bool {
	if trw.isUnwrapType() {
		return trw.Fields[0].t.TypeJSON2ReadingRequiresContext()
	}
	return true
}
