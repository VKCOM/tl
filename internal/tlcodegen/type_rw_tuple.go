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
	wr          *TypeRWWrapper
	vectorLike  bool   // # [T], because # has no reference name
	dynamicSize bool   // with passed nat param
	size        uint32 // if !dynamicSize
	element     Field

	dictLike       bool // for now, can be true only if vectorLike is true. But should work for dynamicSize tuples, so TODO
	dictKeyString  bool
	dictKeyField   Field
	dictValueField Field
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

func (trw *TypeRWBrackets) markWriteHasError(visitedNodes map[*TypeRWWrapper]bool) bool {
	if trw.dynamicSize {
		return true
	}
	return trw.element.t.MarkWriteHasError(visitedNodes)
}

func (trw *TypeRWBrackets) fillRecursiveUnwrap(visitedNodes map[*TypeRWWrapper]bool) {
	trw.element.t.FillRecursiveUnwrap(visitedNodes)
}

func (trw *TypeRWBrackets) markWantsBytesVersion(visitedNodes map[*TypeRWWrapper]bool) {
	trw.element.t.MarkWantsBytesVersion(visitedNodes)
}

func isDictionaryElement(wr *TypeRWWrapper) (bool, bool, Field, Field) {
	// it is hard to mark Dictionary constructor as dictionary,
	// because it is typedef to Vector or built-in brackets.
	// TODO: FIX IT, because len(structElement.Fields) != 2 is true
	structElement, ok := wr.trw.(*TypeRWStruct)
	if !ok || len(structElement.Fields) != 2 || !strings.Contains(strings.ToLower(wr.tlName.Name), "dictionary") {
		return false, false, Field{}, Field{}
	}
	if structElement.Fields[0].fieldMask != nil { // TODO - allowing this complicates json serialization
		return false, false, Field{}, Field{}
	}
	ok, isString := structElement.Fields[0].t.trw.IsDictKeySafe()
	return ok, isString, structElement.Fields[0], structElement.Fields[1]
}

func phpIsDictionary(wr *TypeRWWrapper) bool {
	isDict, _, _, _ := isDictionaryElement(wr)
	if isDict && wr.tlName.Namespace == "" { // TODO NOT A SOLUTION, BUT...
		return true
	}
	return false
}

func cppIsDictionaryElement(wr *TypeRWWrapper) bool {
	isDict, _, _, _ := isDictionaryElement(wr)
	if isDict && wr.tlName.Namespace == "" { // TODO NOT A SOLUTION, BUT...
		return true
	}
	return false
}

func (trw *TypeRWBrackets) FillRecursiveChildren(visitedNodes map[*TypeRWWrapper]int, generic bool) {
	for _, typeDep := range trw.AllPossibleRecursionProducers() {
		typeDep.trw.FillRecursiveChildren(visitedNodes, generic)
	}
}

func (trw *TypeRWBrackets) AllPossibleRecursionProducers() []*TypeRWWrapper {
	var result []*TypeRWWrapper
	for _, typeDep := range trw.wr.arguments {
		if typeDep.tip != nil {
			result = append(result, typeDep.tip.trw.AllPossibleRecursionProducers()...)
		}
	}
	return result
}

func (trw *TypeRWBrackets) AllTypeDependencies(generic, countFunctions bool) (res []*TypeRWWrapper) {
	if !generic {
		if trw.dictLike && len(trw.element.t.origTL[0].TemplateArguments) == 1 {
			pairType := trw.element.t.trw.(*TypeRWStruct)

			keyValue := pairType.Fields[0]
			valueType := pairType.Fields[1]

			res = append(res, keyValue.t)
			res = append(res, valueType.t)
		} else {
			res = append(res, trw.element.t)
		}
	}
	return
}

func (trw *TypeRWBrackets) IsBuiltinVector() bool {
	return len(trw.wr.origTL) == 1 && trw.wr.origTL[0].Builtin
}

func (trw *TypeRWBrackets) IsWrappingType() bool {
	return trw.IsBuiltinVector()
	//if trw.IsBuiltinVector() {
	//	return trw.element.t.trw.IsWrappingType()
	//}
	//return false
}

func (trw *TypeRWBrackets) ContainsUnion(visitedNodes map[*TypeRWWrapper]bool) bool {
	return trw.element.t.containsUnion(visitedNodes)
}

func (trw *TypeRWBrackets) BeforeCodeGenerationStep1() {
	if trw.vectorLike {
		if ok, isString, kf, vf := isDictionaryElement(trw.element.t); ok {
			trw.dictLike = true
			trw.dictKeyString = isString
			trw.dictKeyField = kf
			trw.dictValueField = vf
		}
	}
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
	return false, false
}

func (trw *TypeRWBrackets) CanBeBareBoxed() (canBare bool, canBoxed bool) {
	return true, false
}

func (trw *TypeRWBrackets) typeResettingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, ref bool) string {
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
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
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
	return trw.wr.ins.Prefix(directImports, ins) + fmt.Sprintf("%sFillRandom(rg, %s%s)", goGlobalName, addAmpersand(ref, val), joinWithCommas(natArgs))
}

func (trw *TypeRWBrackets) typeWritingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, bare bool, natArgs []string, ref bool, last bool, needError bool) string {
	refVal := addAmpersand(ref, val)
	if (trw.dictLike && !bytesVersion) || trw.vectorLike || trw.dynamicSize {
		refVal = addAsterisk(ref, val) // those version pass to Write method by pointer
	}
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
	return wrapLastW(last, trw.wr.ins.Prefix(directImports, ins)+fmt.Sprintf("%sWrite%s(w, %s%s)", goGlobalName, addBare(bare), refVal, joinWithCommas(natArgs)), needError)
}

func (trw *TypeRWBrackets) typeReadingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, bare bool, natArgs []string, ref bool, last bool) string {
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
	return wrapLastW(last, trw.wr.ins.Prefix(directImports, ins)+fmt.Sprintf("%sRead%s(w, %s%s)", goGlobalName, addBare(bare), addAmpersand(ref, val), joinWithCommas(natArgs)), true)
}

func (trw *TypeRWBrackets) typeJSONEmptyCondition(bytesVersion bool, val string, ref bool) string {
	if trw.dictLike || trw.vectorLike || trw.dynamicSize {
		return fmt.Sprintf("len(%s) != 0", addAsterisk(ref, val))
	}
	return ""
}

func (trw *TypeRWBrackets) typeJSONWritingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, natArgs []string, ref bool, needError bool) string {
	refVal := addAmpersand(ref, val)
	if (trw.dictLike && !bytesVersion) || trw.vectorLike || trw.dynamicSize {
		refVal = addAsterisk(ref, val) // those version pass to Write method by pointer
	}
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
	// Code which depends on serialization location (skipping empty array if object property) is generated in that location.
	if needError {
		return fmt.Sprintf("if w, err = %sWriteJSONOpt(newTypeNames, short, w, %s%s); err != nil { return w, err }", trw.wr.ins.Prefix(directImports, ins)+goGlobalName, refVal, joinWithCommas(natArgs))
	} else {
		return fmt.Sprintf("w = %sWriteJSONOpt(newTypeNames, short, w, %s%s)", trw.wr.ins.Prefix(directImports, ins)+goGlobalName, refVal, joinWithCommas(natArgs))
	}
}

func (trw *TypeRWBrackets) typeJSONReadingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, jvalue string, val string, natArgs []string, ref bool) string {
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
	return fmt.Sprintf("if err := %sReadJSONLegacy(legacyTypeNames, %s, %s%s); err != nil { return err }", trw.wr.ins.Prefix(directImports, ins)+goGlobalName, jvalue, addAmpersand(ref, val), joinWithCommas(natArgs))
}

func (trw *TypeRWBrackets) typeJSON2ReadingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, jvalue string, val string, natArgs []string, ref bool) string {
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
	return fmt.Sprintf("if err := %sReadJSON(legacyTypeNames, %s, %s%s); err != nil { return err }", trw.wr.ins.Prefix(directImports, ins)+goGlobalName, jvalue, addAmpersand(ref, val), joinWithCommas(natArgs))
}
