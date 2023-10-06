// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlcodegen

import (
	"fmt"
	"strings"
)

// check that brackets cannot be function return type

type TypeRWBrackets struct {
	wr           *TypeRWWrapper
	vectorLike   bool // # [T], because # has no reference name
	dynamicSize  bool
	size         uint32
	goGlobalName string
	element      Field

	dictLike       bool // for now, can be true only if vectorLike is true. But should work for dynamicSize tuples, so TODO
	dictKeyString  bool
	dictKeyField   Field
	dictValueField Field
}

func (trw *TypeRWBrackets) wrapper() *TypeRWWrapper { return trw.wr }

func (trw *TypeRWBrackets) canBeBareOrBoxed(bare bool) bool {
	return bare
}

func (trw *TypeRWBrackets) typeStringGlobal(bytesVersion bool) string {
	return addBytes(trw.goGlobalName, bytesVersion)
}

func (trw *TypeRWBrackets) typeString2(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, isLocal bool, skipAlias bool) string {
	if trw.dictLike && !bytesVersion {
		return fmt.Sprintf("map[%s]%s",
			trw.dictKeyField.t.TypeString2(bytesVersion, directImports, ins, isLocal, skipAlias),
			trw.dictValueField.t.TypeString2(bytesVersion, directImports, ins, isLocal, skipAlias))
	}
	if trw.vectorLike || trw.dynamicSize {
		return fmt.Sprintf("[]%s", trw.element.t.TypeString2(bytesVersion, directImports, ins, isLocal, skipAlias))
	}
	return fmt.Sprintf("[%d]%s", trw.size, trw.element.t.TypeString2(bytesVersion, directImports, ins, isLocal, skipAlias))
}

func (trw *TypeRWBrackets) markHasBytesVersion(visitedNodes map[*TypeRWWrapper]bool) bool {
	if trw.dictLike {
		return true
	}
	return trw.element.t.MarkHasBytesVersion(visitedNodes)
}

func (trw *TypeRWBrackets) fillRecursiveUnwrap(visitedNodes map[*TypeRWWrapper]bool) {
	trw.element.t.FillRecursiveUnwrap(visitedNodes)
}

func (trw *TypeRWBrackets) markWantsBytesVersion(visitedNodes map[*TypeRWWrapper]bool) {
	trw.element.t.MarkWantsBytesVersion(visitedNodes)
}

func dictElement(wr *TypeRWWrapper) (bool, bool, Field, Field) {
	structElement, ok := wr.trw.(*TypeRWStruct)
	// TODO - better ideas?
	if !ok || len(structElement.Fields) != 2 || !strings.Contains(strings.ToLower(wr.TypeString(false)), "dictionary") {
		return false, false, Field{}, Field{}
	}
	if structElement.Fields[0].fieldMask != nil { // TODO - allowing this complicates json serialization
		return false, false, Field{}, Field{}
	}
	ok, isString := structElement.Fields[0].t.trw.IsDictKeySafe()
	return ok, isString, structElement.Fields[0], structElement.Fields[1]
}

func (trw *TypeRWBrackets) BeforeCodeGenerationStep() error {
	if trw.vectorLike {
		if ok, isString, kf, vf := dictElement(trw.element.t); ok {
			trw.dictLike = true
			trw.dictKeyString = isString
			trw.dictKeyField = kf
			trw.dictValueField = vf
		}
	}
	return trw.element.checkBareBoxed()
}

func (trw *TypeRWBrackets) BeforeCodeGenerationStep2() {
}

func (trw *TypeRWBrackets) fillRecursiveChildren(visitedNodes map[*TypeRWWrapper]bool) {
	if trw.wr.gen.options.Language == "cpp" { // Temporary solution to benchmark combined tl
		// We can make vector break the loop, but then we'd need forward declaration of each type used
		trw.element.t.FillRecursiveChildren(visitedNodes)
		return
	}
	// for golang
	if trw.vectorLike || trw.dynamicSize {
		return
	}
	trw.element.t.FillRecursiveChildren(visitedNodes)
}

func (trw *TypeRWBrackets) IsDictKeySafe() (isSafe bool, isString bool) {
	return false, false // !trw.dictLike && !trw.vectorLike && !trw.dynamicSize && trw.element.t.trw.IsDictKeySafe()
}

func (trw *TypeRWBrackets) typeResettingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, ref bool) string {
	goGlobalName := addBytes(trw.goGlobalName, bytesVersion)
	if trw.dictLike && !bytesVersion {
		return trw.wr.ins.Prefix(directImports, ins) + fmt.Sprintf("%[1]sReset(%s)", goGlobalName, addAsterisk(ref, val))
	}
	if trw.dynamicSize || trw.vectorLike {
		if ref {
			return fmt.Sprintf("*%[1]s = (*%[1]s)[:0]", val)
		}
		return fmt.Sprintf("%[1]s = %[1]s[:0]", val)
	}
	return trw.wr.ins.Prefix(directImports, ins) + fmt.Sprintf("%[1]sReset(%[2]s)", goGlobalName, addAmpersand(ref, val))
}

func (trw *TypeRWBrackets) typeRandomCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, natArgs []string, ref bool) string {
	goGlobalName := addBytes(trw.goGlobalName, bytesVersion)
	return trw.wr.ins.Prefix(directImports, ins) + fmt.Sprintf("%sFillRandom(rand, %s%s)", goGlobalName, addAmpersand(ref, val), formatNatArgsCall(natArgs))
}

func (trw *TypeRWBrackets) typeWritingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, bare bool, natArgs []string, ref bool, last bool) string {
	refVal := addAmpersand(ref, val)
	if (trw.dictLike && !bytesVersion) || trw.vectorLike || trw.dynamicSize {
		refVal = addAsterisk(ref, val) // those version pass to Write method by pointer
	}
	goGlobalName := addBytes(trw.goGlobalName, bytesVersion)
	return wrapLastW(last, trw.wr.ins.Prefix(directImports, ins)+fmt.Sprintf("%sWrite%s(w, %s%s)", goGlobalName, addBare(bare), refVal, formatNatArgsCall(natArgs)))
}

func (trw *TypeRWBrackets) typeReadingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, bare bool, natArgs []string, ref bool, last bool) string {
	goGlobalName := addBytes(trw.goGlobalName, bytesVersion)
	return wrapLastW(last, trw.wr.ins.Prefix(directImports, ins)+fmt.Sprintf("%sRead%s(w, %s%s)", goGlobalName, addBare(bare), addAmpersand(ref, val), formatNatArgsCall(natArgs)))
}

func (trw *TypeRWBrackets) typeJSONEmptyCondition(bytesVersion bool, val string, ref bool) string {
	if trw.dictLike || trw.vectorLike || trw.dynamicSize {
		return fmt.Sprintf("len(%s) != 0", addAsterisk(ref, val))
	}
	return ""
}

func (trw *TypeRWBrackets) typeJSONWritingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, natArgs []string, ref bool) string {
	refVal := addAmpersand(ref, val)
	if (trw.dictLike && !bytesVersion) || trw.vectorLike || trw.dynamicSize {
		refVal = addAsterisk(ref, val) // those version pass to Write method by pointer
	}
	goGlobalName := addBytes(trw.goGlobalName, bytesVersion)
	// Code which depends on serialization location (skipping empty array if object property) is generated in that location.
	return fmt.Sprintf("if w, err = %sWriteJSONOpt(short, w, %s%s); err != nil { return w, err }", trw.wr.ins.Prefix(directImports, ins)+goGlobalName, refVal, formatNatArgsCall(natArgs))
}

func (trw *TypeRWBrackets) typeJSONReadingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, jvalue string, val string, natArgs []string, ref bool) string {
	goGlobalName := addBytes(trw.goGlobalName, bytesVersion)
	return fmt.Sprintf("if err := %sReadJSON(%s, %s%s); err != nil { return err }", trw.wr.ins.Prefix(directImports, ins)+goGlobalName, jvalue, addAmpersand(ref, val), formatNatArgsCall(natArgs))
}
