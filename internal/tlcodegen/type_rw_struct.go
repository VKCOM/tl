// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlcodegen

import (
	"fmt"
	"log"
)

type TypeRWStruct struct {
	wr     *TypeRWWrapper
	Fields []Field

	ResultType    *TypeRWWrapper
	ResultNatArgs []ActualNatArg

	fieldsDec    Deconflicter // TODO - add all generated methods here
	fieldsDecCPP Deconflicter // TODO - add all generated methods here
	setNames     []string     // method names should be the same for bytes and normal versions, so we remember them here
	clearNames   []string
	isSetNames   []string
}

func (trw *TypeRWStruct) isTypeDef() bool {
	return len(trw.Fields) == 1 && trw.Fields[0].originalName == "" && trw.Fields[0].fieldMask == nil && !trw.Fields[0].recursive
}

func (trw *TypeRWStruct) isUnwrapTypeImpl(withAliases bool) bool {
	if !trw.isTypeDef() || trw.wr.preventUnwrap {
		return false
	}
	_, isPrimitive := trw.Fields[0].t.trw.(*TypeRWPrimitive)
	_, isBuiltinBrackets := trw.Fields[0].t.trw.(*TypeRWBrackets)
	str, _ := trw.Fields[0].t.trw.(*TypeRWStruct)
	isWrapped := str != nil && str.isUnwrapType()
	if !withAliases {
		return isPrimitive || isBuiltinBrackets || isWrapped
	}
	if str != nil && len(str.Fields) != 0 {
		br, _ := str.Fields[0].t.trw.(*TypeRWBrackets)
		isWrapped = br != nil && br.dictLike
	}
	// fmt.Printf("tlName=%s isBuiltinBrackets=%v isWrapped=%v\n", trw.wr.tlName, isBuiltinBrackets, isWrapped)
	// All custom wrappers including union elements are aliased. TODO - list of normal types from tlgen
	primitiveDefault := trw.wr.tlName.String() == "int" || trw.wr.tlName.String() == "long" || trw.wr.tlName.String() == "string" || trw.wr.tlName.String() == "float" || trw.wr.tlName.String() == "double" || trw.wr.tlName.String() == "#"
	return (isPrimitive && primitiveDefault) || isBuiltinBrackets || isWrapped
}

func (trw *TypeRWStruct) isUnwrapType() bool {
	return trw.isUnwrapTypeImpl(true)
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

func (trw *TypeRWStruct) BeforeCodeGenerationStep2() {
	if trw.wr.gen.options.Language == "cpp" { // Temporary solution to benchmark combined tl
		var nf []Field
		for _, f := range trw.Fields {
			if !f.recursive {
				nf = append(nf, f)
				//panic("recursive field in union " + trw.wr.tlName.String())
			}
		}
		trw.Fields = nf
	}
}

func (trw *TypeRWStruct) fillRecursiveChildren(visitedNodes map[*TypeRWWrapper]bool) {
	for _, f := range trw.Fields {
		if !f.recursive {
			f.t.FillRecursiveChildren(visitedNodes)
		}
	}
}

func (trw *TypeRWStruct) IsDictKeySafe() (isSafe bool, isString bool) {
	return false, false
	// In the future - support typedefs
	// if len(trw.Fields) == 0 {
	//	return false
	// }
	// for _, field := range trw.Fields {
	//	if field.recursive || !field.t.trw.IsDictKeySafe() {
	//		return false
	//	}
	// }
	// return true
}

// same code as in func (w *TypeRWWrapper) transformNatArgsToChild
func (trw *TypeRWStruct) replaceUnwrapArgs(natArgs []string) []string {
	// Caller called outer.Read(   , nat_x, nat_y)
	// outer has func Read(   ,nat_inner_x uint32, nat_inner_y uint32) {
	// which calls for example inner.Read(   , nat_inner_y, nat_inner_y)
	// in other words, outer passes some parameters to inner in some order, with potential repeats.
	// When unwrapping, we do the job of golang compiler, replacing references to outer nat parameters,
	// so that at the calling site outer.Read(   , nat_x, nat_y) is replaced to
	// inner.Read(   , nat_y, nat_y)
	var result []string
outer:
	for _, arg := range trw.Fields[0].natArgs {
		if arg.isArith || arg.isField {
			panic("cannot replace to child arith or field nat param")
		}
		for i, p := range trw.wr.NatParams {
			if p == arg.name {
				result = append(result, natArgs[i])
				continue outer
			}
		}
		log.Panicf("internal compiler error, nat parameter %s not found for unwrap type of goName %s", arg.name, trw.wr.goGlobalName)
	}
	return result
}

func (trw *TypeRWStruct) typeResettingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, ref bool) string {
	if trw.isUnwrapType() {
		return trw.Fields[0].t.TypeResettingCode(bytesVersion, directImports, ins, val, ref)
	}
	return fmt.Sprintf("%s.Reset()", val)
}

func (trw *TypeRWStruct) typeRandomCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, natArgs []string, ref bool) string {
	if trw.isUnwrapType() {
		return trw.Fields[0].t.TypeRandomCode(bytesVersion, directImports, ins, val, trw.replaceUnwrapArgs(natArgs), ref)
	}
	return fmt.Sprintf("%s.FillRandom(rand %s)", val, joinWithCommas(natArgs))
}

func (trw *TypeRWStruct) typeWritingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, bare bool, natArgs []string, ref bool, last bool) string {
	if trw.isUnwrapType() {
		prefix := ""
		if !bare {
			prefix = fmt.Sprintf("w = basictl.NatWrite(w, 0x%x)\n", trw.wr.tlTag)
		}
		return prefix + trw.Fields[0].t.TypeWritingCode(bytesVersion, directImports, ins, val, trw.Fields[0].Bare(), trw.replaceUnwrapArgs(natArgs), ref, last)
		// was
		// goName := addBytes(trw.goGlobalName, bytesVersion)
		// return wrapLastW(last, fmt.Sprintf("(*%s)(%s).Write%s(w%s)", trw.wr.ins.Prefix(ins)+goName, addAmpersand(ref, val), addBare(bare), joinWithCommas(natArgs)))
	}
	return wrapLastW(last, fmt.Sprintf("%s.Write%s(w %s)", val, addBare(bare), joinWithCommas(natArgs)))
}

func (trw *TypeRWStruct) typeReadingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, bare bool, natArgs []string, ref bool, last bool) string {
	if trw.isUnwrapType() {
		prefix := ""
		if !bare {
			prefix = fmt.Sprintf("if w, err = basictl.NatReadExactTag(w, 0x%x); err != nil {\nreturn w, err\n}\n", trw.wr.tlTag)
		}
		return prefix + trw.Fields[0].t.TypeReadingCode(bytesVersion, directImports, ins, val, trw.Fields[0].Bare(), trw.replaceUnwrapArgs(natArgs), ref, last)
		// was
		// goName := addBytes(trw.goGlobalName, bytesVersion)
		// return wrapLastW(last, fmt.Sprintf("(*%s)(%s).Read%s(w%s)", trw.wr.ins.Prefix(ins)+goName, addAmpersand(ref, val), addBare(bare), joinWithCommas(natArgs)))
	}
	return wrapLastW(last, fmt.Sprintf("%s.Read%s(w %s)", val, addBare(bare), joinWithCommas(natArgs)))
}

func (trw *TypeRWStruct) typeJSONEmptyCondition(bytesVersion bool, val string, ref bool) string {
	if trw.isUnwrapTypeImpl(false) {
		return trw.Fields[0].t.TypeJSONEmptyCondition(bytesVersion, val, ref)
	}
	return ""
}

func (trw *TypeRWStruct) typeJSONWritingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, natArgs []string, ref bool) string {
	if trw.isUnwrapType() {
		return trw.Fields[0].t.TypeJSONWritingCode(bytesVersion, directImports, ins, val, trw.replaceUnwrapArgs(natArgs), ref)
	}
	return fmt.Sprintf("if w, err = %s.WriteJSONOpt(short, w %s); err != nil { return w, err }", val, joinWithCommas(natArgs))
}

func (trw *TypeRWStruct) typeJSONReadingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, jvalue string, val string, natArgs []string, ref bool) string {
	if trw.isUnwrapType() {
		return trw.Fields[0].t.TypeJSONReadingCode(bytesVersion, directImports, ins, jvalue, val, trw.replaceUnwrapArgs(natArgs), ref)
	}
	goName := addBytes(trw.wr.goGlobalName, bytesVersion)
	return fmt.Sprintf("if err := %s__ReadJSON(%s, %s %s); err != nil { return err }", trw.wr.ins.Prefix(directImports, ins)+goName, addAmpersand(ref, val), jvalue, joinWithCommas(natArgs))
}
