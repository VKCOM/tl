// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package genphp

type TypeRWUnion struct {
	wr     *TypeRWWrapper
	Fields []Field
	IsEnum bool

	fieldsDec    Deconflicter // TODO - add all generated methods here
	fieldsDecCPP Deconflicter // TODO - add all generated methods here
}

func (trw *TypeRWUnion) fillRecursiveUnwrap(visitedNodes map[*TypeRWWrapper]bool) {
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
