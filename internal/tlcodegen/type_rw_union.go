// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlcodegen

import (
	"fmt"
)

type TypeRWUnion struct {
	wr           *TypeRWWrapper
	goGlobalName string
	goLocalName  string
	Fields       []Field
	IsEnum       bool

	fieldsDec    Deconflicter // TODO - add all generated methods here
	fieldsDecCPP Deconflicter // TODO - add all generated methods here
}

func (trw *TypeRWUnion) canBeBareOrBoxed(bare bool) bool {
	return !bare
}

func (trw *TypeRWUnion) typeStringGlobal(bytesVersion bool) string {
	return addBytes(trw.goGlobalName, bytesVersion)
}

func (trw *TypeRWUnion) typeString2(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, isLocal bool, skipAlias bool) string {
	if isLocal {
		return addBytes(trw.goLocalName, bytesVersion)
	}
	return trw.wr.ins.Prefix(directImports, ins) + addBytes(trw.goGlobalName, bytesVersion)
}

func (trw *TypeRWUnion) markHasBytesVersion(visitedNodes map[*TypeRWWrapper]bool) bool {
	result := false
	for _, f := range trw.Fields {
		result = result || f.t.MarkHasBytesVersion(visitedNodes)
	}
	return result
}

func (trw *TypeRWUnion) fillRecursiveUnwrap(visitedNodes map[*TypeRWWrapper]bool) {
}

func (trw *TypeRWUnion) markWantsBytesVersion(visitedNodes map[*TypeRWWrapper]bool) {
	for _, f := range trw.Fields {
		f.t.MarkWantsBytesVersion(visitedNodes)
	}
}

func (trw *TypeRWUnion) BeforeCodeGenerationStep() error {
	return nil
}

func (trw *TypeRWUnion) BeforeCodeGenerationStep2() {
	for i, f := range trw.Fields {
		visitedNodes := map[*TypeRWWrapper]bool{}
		f.t.trw.fillRecursiveChildren(visitedNodes)
		trw.Fields[i].recursive = visitedNodes[trw.wr]
	}
	//if trw.wr.gen.options.Language == "cpp" { // Temporary solution to benchmark combined tl
	//	var nf []Field
	//	for _, f := range trw.Fields {
	//		if !f.recursive {
	//			nf = append(nf, f)
	//			panic("recursive field in union " + trw.wr.tlName.String())
	//}
	//}
	//trw.Fields = nf
	//}
}

func (trw *TypeRWUnion) fillRecursiveChildren(visitedNodes map[*TypeRWWrapper]bool) {
	//if trw.wr.gen.options.Language == "cpp" { // Temporary solution to benchmark combined tl
	//	for _, f := range trw.Fields {
	//		if !f.recursive {
	//			f.t.FillRecursiveChildren(visitedNodes)
	//		}
	//	}
	//}
}

func (trw *TypeRWUnion) IsDictKeySafe() (isSafe bool, isString bool) {
	return false, false // trw.IsEnum - TODO - in the future?
}

func (trw *TypeRWUnion) typeResettingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, ref bool) string {
	return fmt.Sprintf("%s.Reset()", val)
}

func (trw *TypeRWUnion) typeRandomCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, natArgs []string, ref bool) string {
	return fmt.Sprintf("%s.FillRandom(rand%s)", val, formatNatArgsCall(natArgs))
}

func (trw *TypeRWUnion) typeWritingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, bare bool, natArgs []string, ref bool, last bool) string {
	return wrapLastW(last, fmt.Sprintf("%s.Write%s(w %s)", val, addBare(bare), formatNatArgsCall(natArgs)))
}

func (trw *TypeRWUnion) typeReadingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, bare bool, natArgs []string, ref bool, last bool) string {
	return wrapLastW(last, fmt.Sprintf("%s.Read%s(w %s)", val, addBare(bare), formatNatArgsCall(natArgs)))
}

func (trw *TypeRWUnion) typeJSONEmptyCondition(bytesVersion bool, val string, ref bool) string {
	return ""
}

func (trw *TypeRWUnion) typeJSONWritingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, natArgs []string, ref bool) string {
	return fmt.Sprintf("if w, err = %s.WriteJSON(w %s); err != nil { return w, err }", val, formatNatArgsCall(natArgs))
}

func (trw *TypeRWUnion) typeJSONReadingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, jvalue string, val string, natArgs []string, ref bool) string {
	goName := addBytes(trw.goGlobalName, bytesVersion)
	return fmt.Sprintf("if err := %s__ReadJSON(%s, %s %s); err != nil { return err }", trw.wr.ins.Prefix(directImports, ins)+goName, addAmpersand(ref, val), jvalue, formatNatArgsCall(natArgs))
}
