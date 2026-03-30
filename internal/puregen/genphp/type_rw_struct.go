// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package genphp

type TypeRWStruct struct {
	wr     *TypeRWWrapper
	Fields []Field

	ResultType         *TypeRWWrapper
	ResultNatArgs      []ActualNatArg
	ResultHalfResolved HalfResolvedArgument

	fieldsDec    Deconflicter // TODO - add all generated methods here
	fieldsDecCPP Deconflicter // TODO - add all generated methods here
	setNames     []string     // method names should be the same for bytes and normal versions, so we remember them here
	clearNames   []string
	isSetNames   []string
}

func (trw *TypeRWStruct) isTypeDef() bool {
	return len(trw.Fields) == 1 && trw.Fields[0].originalName == "" && trw.Fields[0].fieldMask == nil && !trw.Fields[0].recursive
}

func (trw *TypeRWStruct) isUnwrapType() bool {
	if !trw.isTypeDef() || trw.wr.preventUnwrap {
		return false
	}
	// Motivation - we want default wrappers for primitive types, vector and tuple to generate primitive language types
	primitive, isPrimitive := trw.Fields[0].t.trw.(*TypeRWPrimitive)
	if isPrimitive && primitive.tlType == trw.wr.tlName.String() {
		return true
	}
	brackets, isBuiltinBrackets := trw.Fields[0].t.trw.(*TypeRWBrackets)
	if isBuiltinBrackets && (brackets.dictLike || trw.wr.tlName.String() == "vector" || trw.wr.tlName.String() == "tuple") {
		return true
	}
	//if trw.wr.gen.options.Language != "cpp" {
	//in combined TL Dictionary is defined via Vector.
	//dictionaryField {t:Type} key:string value:t = DictionaryField t;
	//dictionary#1f4c618f {t:Type} %(Vector %(DictionaryField t)) = Dictionary t;
	//TODO - change combined.tl to use # [] after we fully control generation of C++ & (k)PHP and remove code below
	str, isStruct := trw.Fields[0].t.trw.(*TypeRWStruct)
	if isStruct && str.wr.tlName.String() == "vector" {
		// repeat check above 1 level deeper
		brackets, isBuiltinBrackets := str.Fields[0].t.trw.(*TypeRWBrackets)
		if isBuiltinBrackets && brackets.dictLike {
			return true
		}
	}
	//}
	return false
}

func (trw *TypeRWWrapper) replaceUnwrapHalfResolvedName(topHalfResolved HalfResolvedArgument, name string) string {
	if name == "" {
		return ""
	}
	for i, arg := range trw.origTL[0].TemplateArguments {
		if arg.FieldName == name {
			return topHalfResolved.Args[i].Name
		}
	}
	return ""
}

// same code as in func (w *TypeRWWrapper) transformNatArgsToChild, replaceUnwrapArgs
func (trw *TypeRWWrapper) replaceUnwrapHalfResolved(topHalfResolved HalfResolvedArgument, halfResolved HalfResolvedArgument) HalfResolvedArgument {
	// example
	// tuple#9770768a {t:Type} {n:#} [t] = Tuple t n;
	// innerMaybe {X:#} a:(Maybe (tuple int X)) = InnerMaybe X;
	// when unwrapping we need to change tuple<int, X> into __tuple<X, int>
	// halfResolved references in field of tuple<int, X> are to "n", "t" local template args
	// we must look up in tuple<int, X> to replace "n" "t" into "X", ""
	var result HalfResolvedArgument
	result.Name = trw.replaceUnwrapHalfResolvedName(topHalfResolved, halfResolved.Name)
	for _, arg := range halfResolved.Args {
		result.Args = append(result.Args, trw.replaceUnwrapHalfResolved(topHalfResolved, arg))
	}
	return result
}

func (trw *TypeRWStruct) fillRecursiveUnwrap(visitedNodes map[*TypeRWWrapper]bool) {
	if !trw.isTypeDef() {
		return
	}
	trw.Fields[0].t.FillRecursiveUnwrap(visitedNodes)
}

func (trw *TypeRWStruct) BeforeCodeGenerationStep1() {
	trw.setNames = make([]string, len(trw.Fields))
	trw.clearNames = make([]string, len(trw.Fields))
	trw.isSetNames = make([]string, len(trw.Fields))
}

func (trw *TypeRWStruct) IsDictKeySafe() (isSafe bool, isString bool) {
	if trw.isTypeDef() {
		return trw.Fields[0].t.trw.IsDictKeySafe()
	}
	return false, false
}

func (trw *TypeRWStruct) CanBeBareBoxed() (canBare bool, canBoxed bool) {
	return true, true
}
