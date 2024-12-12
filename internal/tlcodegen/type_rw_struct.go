// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlcodegen

import (
	"fmt"
	"log"
	"sort"
	"strings"
)

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
	if trw.wr.gen.options.Language != "cpp" {
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
	}
	return false
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

func (trw *TypeRWStruct) AllPossibleRecursionProducers() []*TypeRWWrapper {
	var result []*TypeRWWrapper
	for _, typeDep := range trw.wr.arguments {
		if typeDep.tip != nil {
			result = append(result, typeDep.tip.trw.AllPossibleRecursionProducers()...)
		}
	}
	if !trw.isTypeDef() {
		result = append(result, trw.wr)
	}
	return result
}

func (trw *TypeRWStruct) AllTypeDependencies(generic, countFunctions bool) (res []*TypeRWWrapper) {
	used := make(map[*TypeRWWrapper]bool)
	ti := trw.wr.gen.typesInfo
	red := ti.TypeNameToGenericTypeReduction(trw.wr.tlName)

	for i, f := range trw.Fields {
		var deps []*TypeRWWrapper
		if generic {
			fieldRed := ti.FieldTypeReduction(&red, i)
			deps = f.t.ActualTypeDependencies(fieldRed)
		} else {
			deps = append(deps, f.t)
		}
		for _, dep := range deps {
			used[dep] = true
		}
	}

	if countFunctions && trw.ResultType != nil {
		returnRed := ti.TypeNameToGenericTypeReduction(trw.ResultType.tlName)
		for _, t := range trw.ResultType.ActualTypeDependencies(EvaluatedType{Index: TypeConstant, Type: &returnRed}) {
			used[t] = true
		}
	}

	for tp := range used {
		res = append(res, tp)
	}
	return
}

func (trw *TypeRWStruct) IsWrappingType() bool {
	return trw.isUnwrapType()
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

	ti := trw.wr.gen.typesInfo
	red := ti.TypeNameToGenericTypeReduction(trw.wr.tlName)

	for i, f := range trw.Fields {
		if f.recursive {
			continue
		}
		var typeDeps []*TypeRWWrapper
		if generic {
			typeDeps = f.t.ActualTypeDependencies(ti.FieldTypeReduction(&red, i))
		} else {
			typeDeps = f.t.trw.AllPossibleRecursionProducers()
		}
		for _, typeDep := range typeDeps {
			if visitedNodes[typeDep] == 1 {
				trw.Fields[i].recursive = true
			} else {
				typeDep.trw.FillRecursiveChildren(visitedNodes, generic)
			}
		}
	}
	visitedNodes[trw.wr] = 2
}

func (trw *TypeRWStruct) BeforeCodeGenerationStep1() {
	if trw.wr.gen.options.Language == "go" {
		for i, f := range trw.Fields {
			visitedNodes := map[*TypeRWWrapper]bool{}
			f.t.trw.fillRecursiveChildren(visitedNodes)
			trw.Fields[i].recursive = visitedNodes[trw.wr]
		}
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
			index := field.fieldMask.FieldIndex
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
			index := field.fieldMask.FieldIndex
			if _, contains := containingNats[index]; !contains {
				nats = append(nats, trw.Fields[index])
				containingNats[index] = true
			}
		}
	}

	return nats
}

type FieldNatProperties = int

const (
	FieldIsNotNat        FieldNatProperties = 0
	FieldIsNat           FieldNatProperties = 1
	FieldUsedAsFieldMask FieldNatProperties = 2
	FieldUsedAsSize      FieldNatProperties = 4
)

func (trw *TypeRWStruct) GetFieldNatProperties(fieldId int) (FieldNatProperties, []uint32) {
	if fieldId < 0 || len(trw.Fields) <= fieldId {
		return FieldIsNotNat, nil
	}
	targetField := trw.Fields[fieldId]
	pr, isPr := targetField.t.trw.(*TypeRWPrimitive)
	if !isPr || pr.tlType != "#" {
		return FieldIsNotNat, nil
	}
	result := FieldIsNat
	affectedIndexes := make(map[uint32]bool)
	natParamUsageMap := make(map[VisitedTypeNatParam]VisitResult)
	for i, f := range trw.Fields {
		if i == fieldId {
			continue
		}
		if f.fieldMask != nil &&
			f.fieldMask.isField &&
			f.fieldMask.FieldIndex == fieldId {
			affectedIndexes[f.BitNumber] = true
			result |= FieldUsedAsFieldMask
		}
		natIndexes := make([]int, 0)
		for j, natArg := range f.natArgs {
			if natArg.isField && natArg.FieldIndex == fieldId {
				natIndexes = append(natIndexes, j)
			}
		}
		for _, j := range natIndexes {
			visit(f.t, j, &natParamUsageMap, &affectedIndexes, &result)
		}
	}
	indexes := make([]uint32, 0)
	for i := range affectedIndexes {
		indexes = append(indexes, i)
	}
	// not necessary
	sort.Slice(indexes, func(i, j int) bool {
		return indexes[i] < indexes[j]
	})
	return result, indexes
}

type VisitedTypeNatParam struct {
	Type_    string
	NatIndex int
}

type VisitResult = int

const (
	VisitSuccess VisitResult = iota
	VisitFail
	VisitInProgress
)

func visit(
	t *TypeRWWrapper,
	natIndex int,
	visitResults *map[VisitedTypeNatParam]VisitResult,
	affectedIndexes *map[uint32]bool,
	natProps *FieldNatProperties,
) VisitResult {
	natParamName := t.NatParams[natIndex]
	typeName := t.goGlobalName
	key := VisitedTypeNatParam{typeName, natIndex}

	visitResult, isVisited := (*visitResults)[key]
	if isVisited {
		return visitResult
	}
	(*visitResults)[key] = VisitInProgress

	switch i := t.trw.(type) {
	case *TypeRWStruct:
		{
			for _, f := range i.Fields {
				if f.fieldMask != nil &&
					!f.fieldMask.isField &&
					!f.fieldMask.isArith &&
					natParamName == f.fieldMask.name {
					*natProps |= FieldUsedAsFieldMask
					(*affectedIndexes)[f.BitNumber] = true
					(*visitResults)[key] = VisitSuccess
				}
				natIndexes := make([]int, 0)
				for i, natParam := range f.t.NatParams {
					if natParam == natParamName {
						natIndexes = append(natIndexes, i)
					}
				}
				for _, index := range natIndexes {
					res := visit(f.t, index, visitResults, affectedIndexes, natProps)
					if res == VisitSuccess {
						(*visitResults)[key] = VisitSuccess
					}
				}
			}
		}
	case *TypeRWUnion:
		{
			for _, f := range i.Fields {
				res := visit(f.t, natIndex, visitResults, affectedIndexes, natProps)
				if res == VisitSuccess {
					(*visitResults)[key] = VisitSuccess
				}
			}
		}
	case *TypeRWMaybe:
		{
			res := visit(i.element.t, natIndex, visitResults, affectedIndexes, natProps)
			if res == VisitSuccess {
				(*visitResults)[key] = VisitSuccess
			}
		}
	case *TypeRWBrackets:
		{
			*natProps |= FieldUsedAsSize
			elementType := i.element.t
			natIndexes := make([]int, 0)
			for i, natParam := range elementType.NatParams {
				if natParam == natParamName {
					natIndexes = append(natIndexes, i)
				}
			}
			for _, index := range natIndexes {
				res := visit(elementType, index, visitResults, affectedIndexes, natProps)
				if res == VisitSuccess {
					(*visitResults)[key] = VisitSuccess
				}
			}
		}
	}

	if (*visitResults)[key] == VisitInProgress {
		(*visitResults)[key] = VisitFail
	}
	return (*visitResults)[key]
}

// AllAffectedFieldMasks f must be from trw.Fields
func (trw *TypeRWStruct) AllAffectedFieldMasks(f Field) (nats []Field, bits []uint32) {
	curField := f
	for curField.IsAffectingLocalFieldMasks() {
		ancestor := trw.Fields[curField.fieldMask.FieldIndex]
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

func (trw *TypeRWStruct) IsDictKeySafe() (isSafe bool, isString bool) {
	if trw.isTypeDef() {
		return trw.Fields[0].t.trw.IsDictKeySafe()
	}
	return false, false
}

func (trw *TypeRWStruct) CanBeBareBoxed() (canBare bool, canBoxed bool) {
	return true, true
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
	return fmt.Sprintf("%s.FillRandom(rg %s)", val, joinWithCommas(natArgs))
}

func (trw *TypeRWStruct) typeWritingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, bare bool, natArgs []string, ref bool, last bool, needError bool) string {
	if trw.isUnwrapType() {
		prefix := ""
		if !bare {
			prefix = fmt.Sprintf("w = basictl.NatWrite(w, 0x%x)\n", trw.wr.tlTag)
		}
		return prefix + trw.Fields[0].t.TypeWritingCode(bytesVersion, directImports, ins, val, trw.Fields[0].Bare(), trw.replaceUnwrapArgs(natArgs), ref, last, needError)
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
		return prefix + trw.Fields[0].t.TypeReadingCode(bytesVersion, directImports, ins, val, trw.Fields[0].Bare(), trw.replaceUnwrapArgs(natArgs), ref, last)
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
		return trw.Fields[0].t.TypeJSONWritingCode(bytesVersion, directImports, ins, val, trw.replaceUnwrapArgs(natArgs), ref, needError)
	}
	if needError {
		return fmt.Sprintf("if w, err = %s.WriteJSONOpt(newTypeNames, short, w %s); err != nil { return w, err }", val, joinWithCommas(natArgs))
	} else {
		return fmt.Sprintf("w = %s.WriteJSONOpt(newTypeNames, short, w %s)", val, joinWithCommas(natArgs))
	}
}

func (trw *TypeRWStruct) typeJSONReadingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, jvalue string, val string, natArgs []string, ref bool) string {
	if trw.isUnwrapType() {
		return trw.Fields[0].t.TypeJSONReadingCode(bytesVersion, directImports, ins, jvalue, val, trw.replaceUnwrapArgs(natArgs), ref)
	}
	return fmt.Sprintf("if err := %s.ReadJSONLegacy(legacyTypeNames, %s %s); err != nil { return err }", val, jvalue, joinWithCommas(natArgs))
}

func (trw *TypeRWStruct) typeJSON2ReadingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, jvalue string, val string, natArgs []string, ref bool) string {
	if trw.isUnwrapType() {
		return trw.Fields[0].t.TypeJSON2ReadingCode(bytesVersion, directImports, ins, jvalue, val, trw.replaceUnwrapArgs(natArgs), ref)
	}
	return fmt.Sprintf("if err := %s.ReadJSON(legacyTypeNames, %s %s); err != nil { return err }", val, jvalue, joinWithCommas(natArgs))
}

func (trw *TypeRWStruct) PhpClassName(withPath bool, bare bool) string {
	if len(trw.Fields) == 1 && trw.ResultType == nil && trw.wr.unionParent == nil {
		return trw.Fields[0].t.trw.PhpClassName(withPath, bare)
	}

	if trw.ResultType == nil && trw.wr.IsTrueType() && trw.wr.unionParent == nil {
		return "boolean"
	}

	isDict, _, _, valueType := isDictionaryElement(trw.wr)
	if isDict {
		return valueType.t.trw.PhpClassName(withPath, bare)
	}

	name := trw.wr.tlName.Name
	if !bare {
		name = trw.wr.origTL[0].TypeDecl.Name.Name
		print("debug")
	}
	if len(trw.wr.tlName.Namespace) != 0 {
		name = fmt.Sprintf("%s_%s", trw.wr.tlName.Namespace, name)
	}

	elems := make([]string, 0, len(trw.wr.arguments))
	for _, arg := range trw.wr.arguments {
		if arg.tip != nil {
			elems = append(elems, "__", arg.tip.trw.PhpClassName(false, false))
		}
	}

	name += strings.Join(elems, "")
	if withPath {
		name = trw.wr.PHPTypePath() + name
	}
	return name
}

func (trw *TypeRWStruct) PhpTypeName(withPath bool) string {
	if len(trw.Fields) == 1 && trw.ResultType == nil && trw.wr.unionParent == nil {
		return trw.Fields[0].t.trw.PhpTypeName(withPath)
	}
	return trw.PhpClassName(withPath, true)
}

func (trw *TypeRWStruct) PhpGenerateCode(code *strings.Builder, bytes bool) error {
	if isUsingTLImport(trw) ||
		trw.ResultType != nil ||
		trw.wr.unionParent != nil {
		code.WriteString("\nuse VK\\TL;\n")
	}
	code.WriteString(`
/**
 * @kphp-tl-class
 */
`)
	if "right1__test_gigi" == trw.PhpClassName(false, true) {
		print("debug")
	}
	code.WriteString(fmt.Sprintf("class %s ", trw.PhpClassName(false, true)))
	if trw.wr.unionParent != nil {
		code.WriteString(fmt.Sprintf("implements %s ", trw.wr.unionParent.PhpClassName(true, true)))
	}
	if trw.ResultType != nil {
		code.WriteString(fmt.Sprintf("implements TL\\RpcFunction "))
	}
	code.WriteString("{\n")
	// print fieldmasks
	for _, f := range trw.Fields {
		if f.fieldMask == nil {
			continue
		}
		code.WriteString(
			fmt.Sprintf(
				`
  /** Field mask for $%[1]s field */
  const BIT_%[2]s_%[3]d = (1 << %[3]d);
`,
				f.originalName,
				strings.ToUpper(f.originalName),
				f.BitNumber,
			),
		)
	}
	// print fields declarations
	for _, f := range trw.Fields {
		fieldType, defaultValue := fieldTypeAndDefaultValue(f)
		code.WriteString(
			fmt.Sprintf(
				`
  /** @var %[1]s */
  public $%[2]s = %[3]s;
`,
				fieldType,
				f.originalName,
				defaultValue,
			),
		)
	}
	// print constructor
	necessaryFieldsInConstructor := make([]Field, 0)
	usedFieldMasksIndecies := make([]int, 0)
	usedFieldMasks := make(map[int][]Field)
	for _, f := range trw.Fields {
		if f.fieldMask == nil {
			necessaryFieldsInConstructor = append(necessaryFieldsInConstructor, f)
		} else {
			index := f.fieldMask.FieldIndex
			if !f.fieldMask.isField {
				for i, argument := range trw.wr.NatParams {
					if argument == f.fieldMask.name {
						index = -(i + 1)
						break
					}
				}
			}
			if usedFieldMasks[index] == nil {
				usedFieldMasksIndecies = append(usedFieldMasksIndecies, index)
			}
			usedFieldMasks[index] = append(usedFieldMasks[index], f)
		}
	}
	// print result type for function
	if trw.ResultType != nil {
		code.WriteString(
			fmt.Sprintf(
				`
  /** Allows kphp implicitly load function result class */
  private const RESULT = %s_result::class;
`,
				trw.PhpClassName(true, true),
			),
		)
	}
	// print constructor
	code.WriteString(`
  /**
`)
	for _, f := range necessaryFieldsInConstructor {
		fieldType, _ := fieldTypeAndDefaultValue(f)
		code.WriteString(fmt.Sprintf("   * @param %[1]s $%[2]s\n", fieldType, f.originalName))
	}
	if len(necessaryFieldsInConstructor) == 0 {
		code.WriteString("   * @kphp-inline\n")
	}

	code.WriteString(`   */
`)
	code.WriteString("  public function __construct(")

	for i, f := range necessaryFieldsInConstructor {
		_, defaultValue := fieldTypeAndDefaultValue(f)
		if i != 0 {
			code.WriteString(", ")
		}
		code.WriteString(fmt.Sprintf("$%[1]s = %[2]s", f.originalName, defaultValue))
	}

	code.WriteString(") {\n")
	for _, f := range necessaryFieldsInConstructor {
		code.WriteString(fmt.Sprintf("    $this->%[1]s = $%[1]s;\n", f.originalName))
	}
	code.WriteString("  }\n")

	//sort.Ints(usedFieldMasksIndecies)
	sort.Slice(usedFieldMasksIndecies, func(i, j int) bool {
		iGEQ0 := usedFieldMasksIndecies[i] >= 0
		jGEQ0 := usedFieldMasksIndecies[j] >= 0
		if iGEQ0 && jGEQ0 {
			return usedFieldMasksIndecies[i] <= usedFieldMasksIndecies[j]
		} else if iGEQ0 && !jGEQ0 {
			return true
		} else if !iGEQ0 && jGEQ0 {
			return false
		} else {
			return usedFieldMasksIndecies[i] > usedFieldMasksIndecies[j]
		}
	})

	// print methods to calculate fieldmasks
	for _, natIndex := range usedFieldMasksIndecies {
		natName := ""
		if natIndex < 0 {
			natName = trw.wr.NatParams[-(natIndex + 1)]
		} else {
			natName = trw.Fields[natIndex].originalName
		}
		code.WriteString(`
  /**`)
		additionalArgs := make([]string, 0)
		// arguments with ambiguous existence
		for _, dependentField := range usedFieldMasks[natIndex] {
			if _, isMaybe := dependentField.t.PHPGenCoreType().trw.(*TypeRWMaybe); isMaybe {
				additionalArgs = append(additionalArgs, fmt.Sprintf("$has_%s", dependentField.originalName))
				code.WriteString(fmt.Sprintf("\n   * @param bool $has_%s", dependentField.originalName))
			}
		}
		code.WriteString(`
   * @return int
   */
`,
		)
		code.WriteString(
			fmt.Sprintf(
				"  public function calculate%[1]s(%[2]s) {\n    $mask = 0;\n",
				toPhpFieldMaskName(natName),
				strings.Join(additionalArgs, ", "),
			),
		)

		if trw.PhpClassName(false, true) == "test_namesCheck" {
			print("debug")
		}

		fields := usedFieldMasks[natIndex]
		sort.Slice(fields, func(i, j int) bool {
			if fields[i].BitNumber == fields[j].BitNumber {
				return i < j
			}
			return fields[i].BitNumber < fields[j].BitNumber
		})

		fieldsGroupedByBitNumber := make([][]Field, 0)
		for _, dependentField := range fields {
			if len(fieldsGroupedByBitNumber) == 0 ||
				fieldsGroupedByBitNumber[len(fieldsGroupedByBitNumber)-1][0].BitNumber != dependentField.BitNumber {
				fieldsGroupedByBitNumber = append(fieldsGroupedByBitNumber, make([]Field, 0))
			}
			fieldsGroupedByBitNumber[len(fieldsGroupedByBitNumber)-1] = append(fieldsGroupedByBitNumber[len(fieldsGroupedByBitNumber)-1], dependentField)
		}

		for _, dependentFields := range fieldsGroupedByBitNumber {
			conditions := make([]string, 0)
			bitConstants := make([]string, 0)
			for _, dependentField := range dependentFields {
				condition := ""
				if dependentField.t.IsTrueType() || dependentField.t.PHPNeedsCode() {
					condition = fmt.Sprintf(
						"$this->%[1]s",
						dependentField.originalName,
					)
				} else if _, isMaybe := dependentField.t.PHPGenCoreType().trw.(*TypeRWMaybe); isMaybe {
					condition = fmt.Sprintf("$has_%s", dependentField.originalName)
				} else {
					condition = fmt.Sprintf(
						"$this->%[1]s !== null",
						dependentField.originalName,
					)
				}
				conditions = append(conditions, condition)
				bitConstants = append(bitConstants, fmt.Sprintf(
					"self::BIT_%[1]s_%[2]d",
					strings.ToUpper(dependentField.originalName),
					dependentField.BitNumber))
			}

			finalCondition := conditions[0]
			finalMask := bitConstants[0]

			if len(conditions) > 1 {
				finalCondition = strings.Join(conditions, " && ")
				finalMask = "(" + strings.Join(bitConstants, " | ") + ")"
			}

			code.WriteString(
				fmt.Sprintf(
					`
    if (%[1]s) {
      $mask |= %[2]s;
    }
`,
					finalCondition,
					finalMask,
				),
			)
		}

		code.WriteString("\n    return $mask;\n")
		code.WriteString("  }\n")
	}

	// print function specific methods and types
	if trw.ResultType != nil {
		kphpSpecialCode := ""
		if trw.wr.HasAnnotation("kphp") {
			kphpSpecialCode = fmt.Sprintf(
				`

  /**
   * @param %[1]s $value
   * @return %[2]s_result
   */
  public static function createRpcServerResponse($value) {
    $response = new %[2]s_result();
    $response->value = $value;
    return $response;
  }`,
				trw.ResultType.trw.PhpTypeName(true),
				trw.PhpClassName(true, true),
			)
		}

		code.WriteString(
			fmt.Sprintf(
				`
  /**
   * @param TL\RpcFunctionReturnResult $function_return_result
   * @return %[4]s
   */
  public static function functionReturnValue($function_return_result) {
    if ($function_return_result instanceof %[1]s_result) {
      return $function_return_result->value;
    }
    warning('Unexpected result type in functionReturnValue: ' . ($function_return_result ? get_class($function_return_result) : 'null'));
    return (new %[1]s_result())->value;
  }

  /**
   * @kphp-inline
   *
   * @param TL\RpcResponse $response
   * @return %[4]s
   */
  public static function result(TL\RpcResponse $response) {
    return self::functionReturnValue($response->getResult());
  }%[5]s

  /**
   * @kphp-inline
   *
   * @return string
   */
  public function getTLFunctionName() {
    return '%[3]s';
  }
`,
				trw.PhpClassName(false, true),
				trw.PhpClassName(true, true),
				trw.wr.tlName.String(),
				trw.ResultType.trw.PhpTypeName(true),
				kphpSpecialCode,
			),
		)

	}

	code.WriteString("\n}\n")

	if trw.ResultType != nil {
		code.WriteString(
			fmt.Sprintf(
				`
/**
 * @kphp-tl-class
 */
class %[1]s_result implements TL\RpcFunctionReturnResult {

  /** @var %[2]s */
  public $value = %[3]s;

}
`,
				trw.PhpClassName(false, true),
				trw.ResultType.trw.PhpTypeName(true),
				trw.ResultType.trw.PhpDefaultValue(),
			),
		)
	}
	return nil
}

func toPhpFieldMaskName(natName string) string {
	parts := strings.Split(natName, "_")
	for i, _ := range parts {
		parts[i] = ToUpperFirst(parts[i])
	}
	return strings.Join(parts, "")
}

func isUsingTLImport(trw *TypeRWStruct) bool {
	for _, field := range trw.Fields {
		fieldType, _ := fieldTypeAndDefaultValue(field)
		if strings.Contains(fieldType, "TL\\") {
			return true
		}
	}
	return false
}

func fieldTypeAndDefaultValue(f Field) (string, string) {
	fieldType := f.t.trw.PhpTypeName(true)
	defaultValue := f.t.trw.PhpDefaultValue()
	if f.t.IsTrueType() {
		fieldType = "boolean"
		defaultValue = "true"
		if f.fieldMask != nil {
			defaultValue = "false"
		}
	} else {
		if f.fieldMask != nil {
			defaultValue = "null"
			if _, isMaybe := f.t.PHPGenCoreType().trw.(*TypeRWMaybe); !isMaybe {
				fieldType = fieldType + "|null"
			}
		}
	}
	return fieldType, defaultValue
}

func (trw *TypeRWStruct) PhpDefaultValue() string {
	core := trw.wr.PHPGenCoreType()
	if core != trw.wr {
		return core.PHPDefaultValue()
	}
	if core.IsTrueType() {
		return "true"
	}
	return "null"
}
