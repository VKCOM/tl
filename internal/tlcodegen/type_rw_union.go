// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlcodegen

import (
	"fmt"
)

type TypeRWUnion struct {
	wr     *TypeRWWrapper
	Fields []Field
	IsEnum bool

	fieldsDec    Deconflicter // TODO - add all generated methods here
	fieldsDecCPP Deconflicter // TODO - add all generated methods here
}

func (trw *TypeRWUnion) typeString2(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, isLocal bool, skipAlias bool) string {
	if isLocal {
		return addBytes(trw.wr.goLocalName, bytesVersion)
	}
	return trw.wr.ins.Prefix(directImports, ins) + addBytes(trw.wr.goGlobalName, bytesVersion)
}

func (trw *TypeRWUnion) markHasBytesVersion(visitedNodes map[*TypeRWWrapper]bool) bool {
	result := false
	for _, f := range trw.Fields {
		result = result || f.t.MarkHasBytesVersion(visitedNodes)
	}
	return result
}

func (trw *TypeRWUnion) markWriteHasError(visitedNodes map[*TypeRWWrapper]bool) bool {
	result := false
	for _, f := range trw.Fields {
		result = result || f.t.MarkWriteHasError(visitedNodes)
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

func (trw *TypeRWUnion) markWantsTL2(visitedNodes map[*TypeRWWrapper]bool) {
	for _, f := range trw.Fields {
		f.t.MarkWantsTL2(visitedNodes)
	}
}

func (trw *TypeRWUnion) FillRecursiveChildren(visitedNodes map[*TypeRWWrapper]int, generic bool) {
	if visitedNodes[trw.wr] != 0 {
		return
	}
	visitedNodes[trw.wr] = 1
	for _, f := range trw.Fields {
		if f.recursive {
			continue
		}
		f.t.trw.FillRecursiveChildren(visitedNodes, generic)
	}
	visitedNodes[trw.wr] = 2
}

func (trw *TypeRWUnion) AllPossibleRecursionProducers() []*TypeRWWrapper {
	var result []*TypeRWWrapper
	for _, typeDep := range trw.wr.arguments {
		if typeDep.tip != nil {
			result = append(result, typeDep.tip.trw.AllPossibleRecursionProducers()...)
		}
	}
	result = append(result, trw.wr)
	return result
}

func (trw *TypeRWUnion) AllTypeDependencies(generic, countFunctions bool) (res []*TypeRWWrapper) {
	for _, f := range trw.Fields {
		res = append(res, f.t)
	}
	return
}

func (trw *TypeRWUnion) IsWrappingType() bool {
	return false
}

func (trw *TypeRWUnion) ContainsUnion(visitedNodes map[*TypeRWWrapper]bool) bool {
	return true
}

func (trw *TypeRWUnion) BeforeCodeGenerationStep1() {
}

func (trw *TypeRWUnion) BeforeCodeGenerationStep2() {
	if trw.wr.gen.options.Language == "go" {
		for i, f := range trw.Fields {
			visitedNodes := map[*TypeRWWrapper]bool{}
			f.t.trw.fillRecursiveChildren(visitedNodes)
			trw.Fields[i].recursive = visitedNodes[trw.wr]
		}
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

func (trw *TypeRWUnion) CanBeBareBoxed() (canBare bool, canBoxed bool) {
	return false, true
}

func (trw *TypeRWUnion) typeResettingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, ref bool) string {
	return fmt.Sprintf("%s.Reset()", val)
}

func (trw *TypeRWUnion) typeRandomCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, natArgs []string, ref bool) string {
	return fmt.Sprintf("%s.FillRandom(rg%s)", val, joinWithCommas(natArgs))
}

func (trw *TypeRWUnion) typeWritingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, bare bool, natArgs []string, ref bool, last bool, needError bool) string {
	if bare {
		panic("trying to write bare union, please report TL which caused this")
	}
	return wrapLastW(last, fmt.Sprintf("%s.Write%s(w %s)", val, addBare(bare), joinWithCommas(natArgs)), needError)
}

func (trw *TypeRWUnion) typeReadingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, bare bool, natArgs []string, ref bool, last bool) string {
	if bare {
		panic("trying to write bare union, please report TL which caused this")
	}
	return wrapLastW(last, fmt.Sprintf("%s.Read%s(w %s)", val, addBare(bare), joinWithCommas(natArgs)), true)
}

func (trw *TypeRWUnion) typeJSONEmptyCondition(bytesVersion bool, val string, ref bool) string {
	return ""
}

func (trw *TypeRWUnion) typeJSONWritingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, natArgs []string, ref bool, needError bool) string {
	if needError {
		return fmt.Sprintf("if w, err = %s.WriteJSONOpt(tctx, w %s); err != nil { return w, err }", val, joinWithCommas(natArgs))
	} else {
		return fmt.Sprintf("w = %s.WriteJSONOpt(tctx, w %s)", val, joinWithCommas(natArgs))
	}
}

func (trw *TypeRWUnion) typeJSONReadingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, jvalue string, val string, natArgs []string, ref bool) string {
	return fmt.Sprintf("if err := %s.ReadJSONLegacy(legacyTypeNames, %s %s); err != nil { return err }", val, jvalue, joinWithCommas(natArgs))
}

func (trw *TypeRWUnion) typeJSON2ReadingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, jvalue string, val string, natArgs []string, ref bool) string {
	return fmt.Sprintf("if err := %s.ReadJSON(legacyTypeNames, %s %s); err != nil { return err }", val, jvalue, joinWithCommas(natArgs))
}

func (trw *TypeRWUnion) HasShortFieldCollision(wr *TypeRWWrapper) bool {
	//messages.peerId peer_id:int = messages.ChatId;
	//messagesLong.peerId peer_id:long = messages.ChatId;
	//
	//messages.globalChatId#07a5893d chat_id:long = messages.ChatId;
	//messagesLong.globalChatId global_id:messagesLong.GlobalId = messages.ChatId;

	for _, field := range trw.Fields {
		if field.t == wr {
			return true
		}
	}
	return false
}
