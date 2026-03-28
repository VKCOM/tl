// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package genrust

//{% import "github.com/VKCOM/tl/internal/utils" %}
//{% import "strings" %}
//{% import "sort" %}

import (
	"fmt"
	"strings"

	"github.com/VKCOM/tl/internal/pure"
	"github.com/VKCOM/tl/internal/puregen"
)

type TypeRWStruct struct {
	wr             *TypeRWWrapper
	pureTypeStruct *pure.TypeInstanceStruct
	Fields         []Field

	ResultType    *TypeRWWrapper
	ResultNatArgs []pure.ActualNatArg

	unionParent *TypeRWUnion // a bit hackish, but simple
	unionIndex  int

	fieldsDec  puregen.Deconflicter // TODO - add all generated methods here
	setNames   []string             // method names should be the same for bytes and normal versions, so we remember them here
	clearNames []string
	isSetNames []string
}

var _ TypeRW = &TypeRWStruct{}

func (trw *TypeRWStruct) resultHasFetcher() bool {
	return trw.ResultType != nil && trw.ResultType.HasFetcher()
}

func (trw *TypeRWStruct) isTypedef() bool {
	return trw.pureTypeStruct.IsTypedef()
}

func (trw *TypeRWStruct) isAlias() bool {
	return trw.pureTypeStruct.IsAlias()
}

func (trw *TypeRWStruct) isUnwrapType() bool {
	return trw.pureTypeStruct.IsUnwrap()
}

func (trw *TypeRWStruct) typeString2(bytesVersion bool, directImports *DirectImports, isLocal bool, skipAlias bool) string {
	if !skipAlias && trw.isUnwrapType() {
		return trw.Fields[0].t.TypeString2(bytesVersion, directImports, isLocal, skipAlias)
	}
	if isLocal {
		return addBytes(trw.wr.goLocalName, bytesVersion)
	}
	return addBytes(trw.wr.goGlobalName, bytesVersion)
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
	result := len(trw.AllNewTL2Masks()) != 0 && !trw.wr.OriginTL2()
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

func (trw *TypeRWStruct) markWantsBytesVersion(visitedNodes map[*TypeRWWrapper]bool) {
	for _, f := range trw.Fields {
		f.t.MarkWantsBytesVersion(visitedNodes)
	}
	if trw.ResultType != nil {
		trw.ResultType.MarkWantsBytesVersion(visitedNodes)
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
		if field.IsAffectedByLocalFieldMask() && field.t.IsTrueType() {
			index := field.FieldMask().FieldIndex()
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
		if field.IsAffectedByLocalFieldMask() {
			index := field.FieldMask().FieldIndex()
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
	for curField.IsAffectedByLocalFieldMask() {
		ancestor := trw.Fields[curField.FieldMask().FieldIndex()]
		nats = append(nats, ancestor)
		bits = append(bits, curField.BitNumber())
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

func (trw *TypeRWStruct) typeResettingCode(bytesVersion bool, directImports *DirectImports, val string, ref bool) string {
	if trw.isUnwrapType() {
		return trw.Fields[0].t.TypeResettingCode(bytesVersion, directImports, val, ref)
	}
	return fmt.Sprintf("%s.Reset()", val)
}

func (trw *TypeRWStruct) typeRandomCode(bytesVersion bool, directImports *DirectImports, val string, natArgs []string, ref bool) string {
	if trw.isUnwrapType() {
		return trw.Fields[0].t.TypeRandomCode(bytesVersion, directImports, val, trw.pureTypeStruct.ReplaceUnwrapArgs(0, natArgs), ref)
	}
	return fmt.Sprintf("%s.FillRandom(rg %s)", val, joinWithCommas(natArgs))
}

func (trw *TypeRWStruct) typeRepairMasksCode(bytesVersion bool, directImports *DirectImports, val string, natArgs []string, ref bool) string {
	if trw.isUnwrapType() {
		return trw.Fields[0].t.TypeRepairMasksCode(bytesVersion, directImports, val, trw.pureTypeStruct.ReplaceUnwrapArgs(0, natArgs), ref)
	}
	return fmt.Sprintf("%s.RepairMasks(%s)", val, strings.Join(natArgs, ","))
}

func (trw *TypeRWStruct) typeWritingCode(bytesVersion bool, directImports *DirectImports, val string, bare bool, natArgs []string, ref bool, last bool, needError bool) string {
	if trw.isUnwrapType() {
		prefix := ""
		if !bare {
			prefix = fmt.Sprintf("w = basictl.NatWrite(w, 0x%x)\n", trw.wr.TLTag())
		}
		return prefix + trw.Fields[0].t.TypeWritingCode(bytesVersion, directImports, val, trw.Fields[0].Bare(), trw.pureTypeStruct.ReplaceUnwrapArgs(0, natArgs), ref, last, needError)
		// was
		// goName := addBytes(trw.goGlobalName, bytesVersion)
		// return wrapLastW(last, fmt.Sprintf("(*%s)(%s).Write%s(w%s)", trw.wr.ins.Prefix(ins)+goName, addAmpersand(ref, val), addBare(bare), joinWithCommas(natArgs)))
	}
	return wrapLastW(last, fmt.Sprintf("%s.WriteTL1%s(w %s %s)", val, addBare(bare), joinWithCommas(natArgs), trw.wr.fetcherCall()), needError)
}

func (trw *TypeRWStruct) typeReadingCode(bytesVersion bool, directImports *DirectImports, val string, bare bool, natArgs []string, ref bool, last bool) string {
	if trw.isUnwrapType() {
		prefix := ""
		if !bare {
			prefix = fmt.Sprintf("if w, err = basictl.NatReadExactTag(w, 0x%x); err != nil {\nreturn w, err\n}\n", trw.wr.TLTag())
		}
		return prefix + trw.Fields[0].t.TypeReadingCode(bytesVersion, directImports, val, trw.Fields[0].Bare(), trw.pureTypeStruct.ReplaceUnwrapArgs(0, natArgs), ref, last)
		// was
		// goName := addBytes(trw.goGlobalName, bytesVersion)
		// return wrapLastW(last, fmt.Sprintf("(*%s)(%s).Read%s(w%s)", trw.wr.ins.Prefix(ins)+goName, addAmpersand(ref, val), addBare(bare), joinWithCommas(natArgs)))
	}
	return wrapLastW(last, fmt.Sprintf("%s.ReadTL1%s(w %s %s)", val, addBare(bare), joinWithCommas(natArgs), trw.wr.fetcherCall()), true)
}

func (trw *TypeRWStruct) typeJSONEmptyCondition(bytesVersion bool, val string, ref bool) string {
	if trw.isTypedef() {
		return trw.Fields[0].t.TypeJSONEmptyCondition(bytesVersion, val, ref)
	}
	return ""
}

func (trw *TypeRWStruct) typeJSONWritingCode(bytesVersion bool, directImports *DirectImports, val string, natArgs []string, ref bool, needError bool) string {
	if trw.isUnwrapType() {
		return trw.Fields[0].t.TypeJSONWritingCode(bytesVersion, directImports, val, trw.pureTypeStruct.ReplaceUnwrapArgs(0, natArgs), ref, needError)
	}
	if needError {
		return fmt.Sprintf("if w, err = %s.WriteJSONOpt(jctx, w %s); err != nil { return w, err }", val, joinWithCommas(natArgs))
	} else {
		return fmt.Sprintf("w = %s.WriteJSONOpt(jctx, w %s)", val, joinWithCommas(natArgs))
	}
}

func (trw *TypeRWStruct) typeJSON2ReadingCode(bytesVersion bool, directImports *DirectImports, jvalue string, val string, natArgs []string, ref bool) string {
	if trw.isUnwrapType() {
		return trw.Fields[0].t.TypeJSON2ReadingCode(bytesVersion, directImports, jvalue, val, trw.pureTypeStruct.ReplaceUnwrapArgs(0, natArgs), ref)
	}
	return fmt.Sprintf("if err := %s.ReadJSONGeneral(jctx, %s %s); err != nil { return err }", val, jvalue, joinWithCommas(natArgs))
}

func (trw *TypeRWStruct) GenerateCode(bytesVersion bool, directImports *DirectImports) string {
	return ""
}
