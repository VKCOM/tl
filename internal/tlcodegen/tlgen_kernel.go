// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlcodegen

import (
	"fmt"
	"log"
	"strings"

	"github.com/vkcom/tl/internal/purelegacy"
	"github.com/vkcom/tl/internal/tlast"
	"github.com/vkcom/tl/internal/utils"
)

// Instantiation kernel of tlgen.

func (gen *Gen2) getType(lrc LocalResolveContext, t tlast.TypeRef, unionParent *TypeRWUnion) (*TypeRWWrapper, bool, []ActualNatArg, HalfResolvedArgument, error) {
	tName := t.Type.String()
	// Each named reference is either global type, global constructor, local param or local field
	if localArg, ok := lrc.localNatArgs[tName]; ok {
		e1 := t.PR.BeautifulError(fmt.Errorf("reference to %s %q where type is required", ifString(localArg.natArg.isField, "field", "#-param"), tName))
		e2 := localArg.NamePR.BeautifulError(errSeeHere)
		return nil, false, nil, HalfResolvedArgument{}, tlast.BeautifulError2(e1, e2)
	}
	if lt, ok := lrc.localTypeArgs[tName]; ok {
		if len(t.Args) != 0 {
			e1 := t.PR.BeautifulError(fmt.Errorf("reference to template type arg %q cannot have arguments", tName))
			e2 := lt.PR.BeautifulError(fmt.Errorf("defined here"))
			return nil, false, nil, HalfResolvedArgument{}, tlast.BeautifulError2(e1, e2)
		}
		bare := lt.arg.bare
		if t.Bare { // overwrite bare
			if len(lt.arg.tip.origTL) > 1 {
				// TODO - better error. Does not reference call site
				//----- bare wrapping
				// bareWrapper {X:Type} a:%X = BareWrapper X;
				// bareWrapperTest a:(bareWrapper a.Color) = BareWrapperTest;
				e1 := t.PR.BeautifulError(fmt.Errorf("field type %q is bare, so union %q cannot be passed", tName, lt.arg.tip.CanonicalStringTop()))
				e2 := lt.PR.BeautifulError(fmt.Errorf("defined here"))
				return nil, false, nil, HalfResolvedArgument{}, tlast.BeautifulError2(e1, e2)
			}
			// myUnionA = MyUnion;
			// myUnionB b:int = MyUnion;
			// wrapper {T:Type} a:%T = Wrapper T;
			// useWarpper xx:(wrapper MyUnion) = UseWrapper;
			bare = true
			// TODO - we must perform canonical conversion of %Int to int here
		}
		return lt.arg.tip, bare, lt.natArgs, HalfResolvedArgument{Name: tName}, nil
	}
	var tlType []*tlast.Combinator

	if lt, ok := gen.typeDescriptors[tName]; ok { // order of this if-else chain is important for built-ins
		if len(lt) > 1 && t.Bare {
			// myUnionA = MyUnion;
			// myUnionB b:int = MyUnion;
			// useUnion a:%MyUnion = UseUnion;
			e1 := t.PR.BeautifulError(fmt.Errorf("reference to union %q cannot be bare", tName))
			e2 := lt[0].TypeDecl.NamePR.BeautifulError(fmt.Errorf("see more")) // TODO: maybe better message, see more about union is not very useful
			return nil, false, nil, HalfResolvedArgument{}, tlast.BeautifulError2(e1, e2)
		}
		tlType = lt
		//conName := tlType[0].Construct.Name.String()
		//if con2, ok := gen.singleConstructors[conName]; ok && t.Bare && !con2.IsFunction && con2.TypeDecl.Name.String() == "_" {
		// bare references to wrappers %int have int canonical form,
		// otherwise vectors, maybes and other templates will be generated twice
		//t.Type = tlType[0].Construct.Name
		//}
	} else if lt, ok := gen.singleConstructors[tName]; ok {
		tlType = []*tlast.Combinator{lt}
		t.Bare = true
		//if lt.TypeDecl.Name.String() != "_" {
		// We use "_" in type declaration for internal types which cannot be boxed
		// We could wish to extend this definition to user types in the future
		// If there is no boxed version, constructor name is canonical reference, otherwise
		// Type name is canonical reference. We need canonical references to avoid generating type more than once
		//t.Type = lt.TypeDecl.Name
		//}
	} else if lt, ok := gen.allConstructors[tName]; ok {
		if !lrc.allowAnyConstructor {
			e1 := t.PR.BeautifulError(fmt.Errorf("reference to %s constructor %q is not allowed", ifString(lt.IsFunction, "function", "union"), tName))
			e2 := lt.Construct.NamePR.BeautifulError(fmt.Errorf("see more"))
			return nil, false, nil, HalfResolvedArgument{}, tlast.BeautifulError2(e1, e2)
		}
		// Here type name is already in canonical form, because this code path is only internal for union members and functions
		tlType = []*tlast.Combinator{lt}
		t.Bare = true
	}
	if len(tlType) == 0 {
		return nil, false, nil, HalfResolvedArgument{}, t.PR.BeautifulError(fmt.Errorf("error resolving name %q", tName))
	}
	td := tlType[0] // for type checking, any constructor is ok for us, because they all must have the same args
	if len(td.TemplateArguments) > len(t.Args) {
		arg := td.TemplateArguments[len(t.Args)]
		e1 := t.PRArgs.CollapseToEnd().BeautifulError(fmt.Errorf("missing template argument %q here", arg.FieldName))
		e2 := arg.PR.BeautifulError(fmt.Errorf("declared here"))
		return nil, false, nil, HalfResolvedArgument{}, tlast.BeautifulError2(e1, e2)
	}
	if len(td.TemplateArguments) < len(t.Args) {
		arg := t.Args[len(td.TemplateArguments)]
		e1 := arg.T.PR.BeautifulError(fmt.Errorf("excess template argument %q here", arg.String()))
		e2 := td.TemplateArgumentsPR.BeautifulError(fmt.Errorf("arguments declared here"))
		return nil, false, nil, HalfResolvedArgument{}, tlast.BeautifulError2(e1, e2)
	}
	kernelType := &TypeRWWrapper{
		gen:         gen,
		origTL:      tlType,
		unionParent: unionParent,
	}
	var actualNatArgs []ActualNatArg
	var halfResolved HalfResolvedArgument
	for i, a := range t.Args {
		ta := td.TemplateArguments[i]
		aName := a.T.Type.String()
		if ta.IsNat {
			if a.IsArith {
				kernelType.arguments = append(kernelType.arguments, ResolvedArgument{
					isNat:   true,
					isArith: true,
					Arith:   a.Arith,
				})
				halfResolved.Args = append(halfResolved.Args, HalfResolvedArgument{}) // Empty name here
				continue
			}
			if localArg, ok := lrc.localNatArgs[aName]; ok {
				if localArg.wrongTypeErr != nil {
					e1 := a.T.PR.BeautifulError(fmt.Errorf("error resolving reference %q to #-param %q", aName, ta.FieldName))
					e2 := localArg.TypePR.BeautifulError(localArg.wrongTypeErr)
					return nil, false, nil, HalfResolvedArgument{}, tlast.BeautifulError2(e1, e2)
				}
				kernelType.arguments = append(kernelType.arguments, ResolvedArgument{
					isNat:   true, // true due to check above
					isArith: localArg.natArg.isArith,
					Arith:   localArg.natArg.Arith,
				})
				halfResolved.Args = append(halfResolved.Args, HalfResolvedArgument{Name: aName})
				if !localArg.natArg.isArith {
					actualNatArgs = append(actualNatArgs, localArg.natArg)
				}
				continue
			}
			if localArg, ok := lrc.localTypeArgs[aName]; ok {
				e1 := a.T.PR.BeautifulError(fmt.Errorf("reference to local Type-arg %q where #-arg is required", aName))
				e2 := localArg.PR.BeautifulError(fmt.Errorf("arg declared here"))
				return nil, false, nil, HalfResolvedArgument{}, tlast.BeautifulError2(e1, e2)
			}
			e1 := a.T.PR.BeautifulError(fmt.Errorf("error resolving reference %q to #-param %q", aName, ta.FieldName))
			e2 := ta.PR.BeautifulError(fmt.Errorf("see more"))
			return nil, false, nil, HalfResolvedArgument{}, tlast.BeautifulError2(e1, e2)
		}
		if a.IsArith {
			e1 := a.T.PR.BeautifulError(fmt.Errorf("passing constant %q to Type-param %q is impossible", a.Arith.String(), ta.FieldName))
			e2 := ta.PR.BeautifulError(fmt.Errorf("declared here"))
			return nil, false, nil, HalfResolvedArgument{}, tlast.BeautifulError2(e1, e2)
		}
		internalType, internalBare, internalNatArgs, internalHalfResolved, err := gen.getType(lrc, a.T, nil)
		if err != nil {
			return nil, false, nil, HalfResolvedArgument{}, err
		}
		kernelType.arguments = append(kernelType.arguments, ResolvedArgument{
			tip:  internalType,
			bare: internalBare,
		})
		halfResolved.Args = append(halfResolved.Args, internalHalfResolved)
		actualNatArgs = append(actualNatArgs, internalNatArgs...)
	}
	canonicalName := kernelType.CanonicalStringTop()
	if bt, ok := gen.builtinTypes[kernelType.CanonicalString(t.Bare)]; ok {
		return bt, true, nil, HalfResolvedArgument{}, nil
	}
	exist, ok := gen.generatedTypes[canonicalName]
	if !ok {
		// log.Printf("adding canonical type: %s\n", canonicalName)
		//log.Printf("   half resolved type: %v\n", halfResolved)
		gen.generatedTypes[canonicalName] = kernelType
		gen.generatedTypesList = append(gen.generatedTypesList, kernelType)
		// We added our type already, so others can reference it
		// Now we will iterate over our fields so all types we need are also generated
		if err := gen.generateType(kernelType); err != nil {
			return nil, false, nil, HalfResolvedArgument{}, err
		}
		if lrc.overrideFileName != "" {
			kernelType.fileName = lrc.overrideFileName
		}
		//else {
		//	if gen.options.Language == "cpp" { // Temporary solution to benchmark combined tl
		//		if resolvedType.Type.Namespace == "" && len(resolvedType.Args) == 1 && !resolvedType.Args[0].IsNat {
		//			wr.fileName = resolvedType.Args[0].TRW.fileName
		//		}
		//	}
		//}
		if kernelType.fileName == "" {
			// TODO - check this is impossible, then return LogicError
			log.Printf("Warning: empty type filename for canonical name %q, will move to 'builtin'", canonicalName)
			kernelType.fileName = "builtin"
		}
		return kernelType, t.Bare, actualNatArgs, halfResolved, nil
	}
	// exist.combinator.tips = append(exist.combinator.tips, kernelType) - TODO - collect all instantiations of combinator
	return exist, t.Bare, actualNatArgs, halfResolved, nil
}

func collectArgsNamespaces(tip *TypeRWWrapper, argNamespaces map[string]struct{}) {
	// This is policy. You can adjust it, so more or less templates instantiations
	// are moved into namespace of arguments. Code should compile anyway.
	if tip.tlName.Namespace != "" {
		argNamespaces[tip.tlName.Namespace] = struct{}{}
	}
	for _, arg := range tip.arguments {
		if !arg.isNat {
			collectArgsNamespaces(arg.tip, argNamespaces)
		}
	}
}

func (gen *Gen2) generateType(myWrapper *TypeRWWrapper) error {
	tlType := myWrapper.origTL
	lrc := LocalResolveContext{
		localTypeArgs: map[string]LocalTypeArg{},
		localNatArgs:  map[string]LocalNatArg{},
	}
	argNamespaces := map[string]struct{}{}
	for i, a := range tlType[0].TemplateArguments { // they are the same for all constructors
		if err := lrc.checkArgsCollision(a.FieldName, a.PR, errNatParamNameCollision); err != nil {
			return err
		}
		ra := myWrapper.arguments[i]
		if a.IsNat {
			lrc.localNatArgs[a.FieldName] = LocalNatArg{
				NamePR: a.PR,
				TypePR: a.PR,
				natArg: ActualNatArg{
					isArith: ra.isArith,
					Arith:   ra.Arith,
					name:    a.FieldName,
				},
			}
			if !ra.isArith {
				myWrapper.NatParams = append(myWrapper.NatParams, a.FieldName)
			}
			continue
		}
		collectArgsNamespaces(ra.tip, argNamespaces)
		natArgs := ra.tip.NatArgs(nil, a.FieldName)
		// We can select arbitrary names for arguments here, but they all must be unique per generatedType
		// The simplest idea to avoid collisions is to exploit uniqueness of field names
		if len(natArgs) == 1 { // in most common case avoid longer than necessary names.
			natArgs[0].name = a.FieldName
		}
		// We decided to use internal structure of names instead of assigning them sequential numbers for each template argument.
		// You can uncomment code below to see which naming scheme looks better
		// else {
		//	for j := range natArgs {
		//		natArgs[j].name = fmt.Sprintf("%s%d", a.FieldName, j)
		//	}
		// }
		lrc.localTypeArgs[a.FieldName] = LocalTypeArg{
			arg:     ra,
			PR:      a.PR,
			natArgs: natArgs,
		}
		for _, natArg := range natArgs {
			myWrapper.NatParams = append(myWrapper.NatParams, natArg.name)
		}
	}
	replaceNamespace := func(n string) *Namespace {
		newNamespace := n
		if n == "" && len(argNamespaces) == 1 {
			for ns := range argNamespaces {
				newNamespace = ns
			}
		}
		return gen.getNamespace(newNamespace)
	}
	//myWrapper.cppNamespaceQualifier = "::" + gen.options.RootCPPNamespace + "::"
	//if rt2.Type.Namespace != "" {
	//	myWrapper.cppNamespaceQualifier += rt2.Type.Namespace + "::"
	//}
	if len(tlType) == 1 {
		myWrapper.tlName = tlType[0].Construct.Name
		myWrapper.fileName = tlType[0].Construct.Name.String()
		namespace := replaceNamespace(myWrapper.tlName.Namespace)
		namespace.types = append(namespace.types, myWrapper)
		myWrapper.ns = namespace
		myWrapper.tlTag = tlType[0].Crc32()
		switch tlType[0].Construct.Name.String() { // TODO - better switch
		case BuiltinTupleName:
			// tl2 faking vector
			//if myWrapper.arguments[0].isArith && myWrapper.arguments[0].isTL2FakeArith {
			//	myWrapper.arguments = myWrapper.arguments[1:]
			//	_, tail := myWrapper.resolvedT2GoName("")
			//	myWrapper.goGlobalName = gen.globalDec.deconflictName("BuiltinVector" + tail)
			//	return gen.GenerateVector(myWrapper, tlType[0], lrc, 1)
			//}
			_, tail := myWrapper.resolvedT2GoName("")
			myWrapper.goGlobalName = gen.globalDec.deconflictName("BuiltinTuple" + tail)
			// built-in tuple has no local name. TODO - invent one?
			return gen.GenerateTuple(myWrapper, tlType[0], lrc)
		case BuiltinVectorName:
			_, tail := myWrapper.resolvedT2GoName("")
			myWrapper.goGlobalName = gen.globalDec.deconflictName("BuiltinVector" + tail)
			// built-in vector has no local name. TODO - invent one?
			return gen.GenerateVector(myWrapper, tlType[0], lrc, 0)
		}
		head, tail := myWrapper.resolvedT2GoName("")
		myWrapper.goGlobalName = gen.globalDec.deconflictName(head + tail)
		head, tail = myWrapper.resolvedT2GoName(myWrapper.tlName.Namespace)
		myWrapper.goLocalName = namespace.decGo.deconflictName(head + tail)
		actualName, canonicalName, _ := myWrapper.cppTypeStringInNamespace(false, &DirectIncludesCPP{ns: map[*TypeRWWrapper]CppIncludeInfo{}}, false, HalfResolvedArgument{})
		otherRW, ok := namespace.cppTemplates[canonicalName]
		if ok {
			myWrapper.cppLocalName = otherRW.cppLocalName
		} else {
			//myWrapper.cppLocalName = namespace.decCpp.deconflictName(ToUpperFirst(actualName) + tail)
			myWrapper.cppLocalName = namespace.decCpp.deconflictName(ToUpperFirst(actualName))
			namespace.cppTemplates[canonicalName] = myWrapper
		}
		return gen.generateTypeStruct(lrc, myWrapper, tlType[0])
	}
	myWrapper.tlName = tlType[0].TypeDecl.Name
	if gen.options.Language == "cpp" {
		myWrapper.fileName = tlType[0].TypeDecl.Name.String()
	} else {
		// during migration to TL2 unions (UpperCase) become normal types (lowerCase)
		// and we want to make sure diff is trivial
		fileName := tlType[0].TypeDecl.Name
		fileName.Name = ToLowerFirst(fileName.Name)
		myWrapper.fileName = fileName.String()
	}
	if isBool, falseDesc, trueDesc := IsUnionBool(tlType); isBool { // TODO - test if parts of Bool are in different namespaces
		namespace := replaceNamespace(myWrapper.tlName.Namespace)
		namespace.types = append(namespace.types, myWrapper)
		myWrapper.ns = namespace

		head, tail := myWrapper.resolvedT2GoName("")
		myWrapper.goGlobalName = gen.globalDec.deconflictName(head + tail)
		head, tail = myWrapper.resolvedT2GoName(myWrapper.tlName.Namespace)
		myWrapper.goLocalName = namespace.decGo.deconflictName(head + tail)
		myWrapper.trw = &TypeRWBool{
			isBit:       false,
			wr:          myWrapper,
			falseGoName: gen.globalDec.deconflictName(utils.CNameToCamelName(falseDesc.Construct.Name.String())),
			trueGoName:  gen.globalDec.deconflictName(utils.CNameToCamelName(trueDesc.Construct.Name.String())),
			falseTag:    falseDesc.Crc32(),
			trueTag:     trueDesc.Crc32(),
		}
		return nil
	}
	if isMaybe, emptyDesc, okDesc := IsUnionMaybe(tlType); isMaybe {
		elementT := tlast.TypeRef{Type: tlast.Name{Name: okDesc.TemplateArguments[0].FieldName}} // TODO - PR
		elementResolvedType, elementResolvedTypeBare, elementNatArgs, elementHalfResolved, err := gen.getType(lrc, elementT, nil)
		if err != nil {
			return err
		}

		namespace := replaceNamespace(elementResolvedType.tlName.Namespace)
		namespace.types = append(namespace.types, myWrapper)
		myWrapper.ns = namespace

		suffix := ifString(elementResolvedTypeBare, "Maybe", "BoxedMaybe")
		head, tail := elementResolvedType.resolvedT2GoName("")
		myWrapper.goGlobalName = gen.globalDec.deconflictName(head + tail + suffix)
		head, tail = elementResolvedType.resolvedT2GoName(elementResolvedType.tlName.Namespace)
		myWrapper.goLocalName = namespace.decGo.deconflictName(head + tail + suffix)

		res := &TypeRWMaybe{
			wr: myWrapper,
			element: Field{
				t:            elementResolvedType,
				bare:         elementResolvedTypeBare,
				natArgs:      elementNatArgs,
				halfResolved: elementHalfResolved,
			},
			emptyTag: emptyDesc.Crc32(),
			okTag:    okDesc.Crc32(),
		}
		myWrapper.fileName = elementResolvedType.fileName
		myWrapper.trw = res
		return nil
	}
	isEnum := true
	for _, typ := range tlType {
		isEnum = isEnum && len(typ.Fields) == 0
	}

	namespace := replaceNamespace(myWrapper.tlName.Namespace)
	namespace.types = append(namespace.types, myWrapper)
	myWrapper.ns = namespace

	head, tail := myWrapper.resolvedT2GoName("")
	myWrapper.goGlobalName = gen.globalDec.deconflictName(head + tail)
	head, tail = myWrapper.resolvedT2GoName(myWrapper.tlName.Namespace)
	myWrapper.goLocalName = namespace.decGo.deconflictName(head + tail)
	actualName, canonicalName, _ := myWrapper.cppTypeStringInNamespace(false, &DirectIncludesCPP{ns: map[*TypeRWWrapper]CppIncludeInfo{}}, false, HalfResolvedArgument{})
	otherRW, ok := namespace.cppTemplates[canonicalName]
	if ok {
		myWrapper.cppLocalName = otherRW.cppLocalName
	} else {
		myWrapper.cppLocalName = namespace.decCpp.deconflictName(ToUpperFirst(actualName))
		namespace.cppTemplates[canonicalName] = myWrapper
	}

	lrc.allowAnyConstructor = true
	lrc.overrideFileName = myWrapper.fileName
	if gen.options.Language == "cpp" {
		if isEnum {
			lrc.overrideFileName += "Items" // TODO - in C++ when items and union are in the same file, they must be sorted which is hard for us
		} else {
			lrc.overrideFileName = "" // each type must be in its own file to break circular dependencies
		}
	}

	res := &TypeRWUnion{
		wr:     myWrapper,
		IsEnum: isEnum,
	}
	res.fieldsDecCPP.fillCPPIdentifiers()
	res.fieldsDec.fillGolangIdentifies()
	myWrapper.trw = res

	// Removing prefix/suffix common with union name.
	// We temporarily allow relaxed case match. To use strict match, remove all strings.ToLower() calls below
	typePrefix := strings.ToLower(ToLowerFirst(tlType[0].TypeDecl.Name.Name))
	typeSuffix := strings.ToLower(tlType[0].TypeDecl.Name.Name)
	for _, typ := range tlType {
		conName := strings.ToLower(typ.Construct.Name.Name)
		// if constructor is full prefix of type, we will shorten accessors
		// ab.saveStateOne = ab.SaveState; // item.AsOne()
		// ab.saveStateTwo = ab.SaveState; // item.AsTwo()
		if !strings.HasPrefix(conName, typePrefix) { // same check as in checkUnionElementsCompatibility
			typePrefix = ""
		}
		if !strings.HasSuffix(conName, typeSuffix) {
			typeSuffix = ""
		}
	}
	for i, typ := range tlType {
		// ---- We treat
		// ab.empty = ab.Response;
		// ab.code x:int = ab.Response;
		// ab.response x:int str:string = ab.Response;
		// ---- roughly as
		// ab.empty = _;
		// ab.code x:int = _;
		// ab.response x:int str:string = _;
		// _ tag:# empty:ab.empty code:ab.code response: ab.response = ab.Response;
		fieldType := tlast.TypeRef{
			Type:   typ.Construct.Name,
			Bare:   true,
			PR:     typ.Construct.NamePR,
			PRArgs: typ.TemplateArgumentsPR,
		}
		for _, arg := range typ.TemplateArguments {
			fieldType.Args = append(fieldType.Args, tlast.ArithmeticOrType{
				IsArith: false,
				T: tlast.TypeRef{
					Type:   tlast.Name{Name: arg.FieldName},
					PR:     arg.PR,
					PRArgs: arg.PR.CollapseToEnd(),
				},
			})
		}
		fieldResolvedType, fieldResolvedTypeBare, fieldNatArgs, fieldHalfResolved, err := gen.getType(lrc, fieldType, res)
		if err != nil {
			return err
		}
		// fieldResolvedType.unionParent = res // Already set by getType
		fieldResolvedType.unionIndex = i
		if !fieldResolvedTypeBare {
			return fieldType.PR.BeautifulError(fmt.Errorf("union element resolved type %q cannot be boxed", fieldResolvedType.CanonicalStringTop()))
		}
		typeConstructName := typ.Construct.Name
		if typePrefix != "" && len(typePrefix) < len(typeConstructName.Name) {
			typeConstructName.Name = typeConstructName.Name[len(typePrefix):]
		} else if typeSuffix != "" && len(typeSuffix) < len(typeConstructName.Name) {
			typeConstructName.Name = typeConstructName.Name[:len(typeConstructName.Name)-len(typeSuffix)]
		}
		fieldGoName := canonicalGoName(typeConstructName, typ.Construct.Name.Namespace)
		if res.fieldsDec.hasConflict(fieldGoName) { // try global, if local is already used
			fieldGoName = canonicalGoName(typeConstructName, "")
		}
		fieldCPPName := canonicalCPPName(typeConstructName, typ.Construct.Name.Namespace)
		if res.fieldsDecCPP.hasConflict(fieldCPPName) { // try global, if local is already used
			fieldCPPName = canonicalCPPName(typeConstructName, "")
		}
		newField := Field{
			originalName: fieldType.Type.String(),
			t:            fieldResolvedType,
			bare:         fieldResolvedTypeBare,
			goName:       res.fieldsDec.deconflictName(fieldGoName),
			cppName:      res.fieldsDecCPP.deconflictName(fieldCPPName),
			natArgs:      fieldNatArgs,
			halfResolved: fieldHalfResolved,
			// origTL:       ?, // We do not want to set it here for now
		}
		res.Fields = append(res.Fields, newField)
	}
	return nil
}

func (gen *Gen2) generateTypeStruct(lrc LocalResolveContext, myWrapper *TypeRWWrapper, tlType *tlast.Combinator) error {
	res := &TypeRWStruct{
		wr: myWrapper,
	}
	res.fieldsDecCPP.fillCPPIdentifiers()
	res.fieldsDec.fillGolangIdentifies()
	myWrapper.trw = res
	nextTL2MaskBit := 0
	for i, field := range tlType.Fields {
		fieldType, fieldTypeBare, fieldNatArgs, fieldHalfResolved, err := gen.getType(lrc, field.FieldType, nil)
		if err != nil {
			return err
		}
		fieldName := field.FieldName
		if fieldName == "" {
			// only for typedefs, but TODO - harmonize condition with func (trw *TypeRWStruct) isTypeDef()
			fieldName = "a"
		}
		newField := Field{
			originalName: field.FieldName,
			t:            fieldType,
			bare:         fieldTypeBare,
			goName:       res.fieldsDec.deconflictName(utils.CNameToCamelName(fieldName)),
			cppName:      res.fieldsDecCPP.deconflictName(fieldName),
			natArgs:      fieldNatArgs,
			origTL:       field,

			halfResolved: fieldHalfResolved,
		}
		if field.Mask != nil {
			if field.Mask.BitNumber >= 32 {
				return field.Mask.PRBits.BeautifulError(fmt.Errorf("bitmask (%d) must be in range [0..31]", field.Mask.BitNumber))
			}
			newField.BitNumber = field.Mask.BitNumber
			localArg, ok := lrc.localNatArgs[field.Mask.MaskName]
			if !ok {
				return field.Mask.PRName.BeautifulError(fmt.Errorf("failed to resolve field mask %q reference", field.Mask.MaskName))
			}
			if localArg.wrongTypeErr != nil {
				e1 := field.Mask.PRName.BeautifulError(fmt.Errorf("field mask %q reference to field of wrong type", field.Mask.MaskName))
				e2 := localArg.TypePR.BeautifulError(localArg.wrongTypeErr)
				return tlast.BeautifulError2(e1, e2)
			}
			newField.fieldMask = &localArg.natArg
			maskBit := nextTL2MaskBit
			newField.MaskTL2Bit = &maskBit
			nextTL2MaskBit++
		}
		res.Fields = append(res.Fields, newField)
		arg := LocalNatArg{
			NamePR: field.PRName,
			TypePR: field.FieldType.PR,
			natArg: ActualNatArg{isField: true, FieldIndex: i},
		}
		if field.FieldType.Type.String() != "#" {
			arg.wrongTypeErr = fmt.Errorf("referenced field %q must have type #", field.FieldName)
		}
		if field.FieldName == "" {
			continue
		}
		if err := lrc.checkArgsCollision(field.FieldName, field.PRName, errFieldNameCollision); err != nil {
			return err
		}
		if newField.t.IsTrueType() && !newField.Bare() &&
			newField.t.origTL[0].TypeDecl.Name.String() == "True" &&
			newField.t.origTL[0].Construct.Name.String() == "true" &&
			!purelegacy.AllowTrueBoxed(myWrapper.origTL[0].Construct.Name.String(), field.FieldName) &&
			utils.DoLint(field.CommentRight) {
			// We compare type by name, because there is examples of other true types which are to be extended
			// to unions or have added fields in the future
			e1 := field.FieldType.PR.BeautifulError(fmt.Errorf("true type fields should be bare, use 'true' or '%%True' instead"))
			if gen.options.WarningsAreErrors {
				return e1
			}
			e1.PrintWarning(gen.options.ErrorWriter, nil)
		}
		if _, ok := newField.t.trw.(*TypeRWBool); ok {
			if newField.t.origTL[0].TypeDecl.Name.String() == "Bool" &&
				newField.fieldMask != nil && !newField.fieldMask.isArith && newField.fieldMask.isField &&
				!purelegacy.AllowBoolFieldsmask(myWrapper.origTL[0].Construct.Name.String(), field.FieldName) &&
				utils.DoLint(field.CommentRight) {
				// We compare type by name to make warning more narrow at first.
				e1 := field.FieldType.PR.BeautifulError(fmt.Errorf("using Bool type under fields mask produces 3rd state, you probably want to use 'true' instead of 'Bool'"))
				if gen.options.WarningsAreErrors {
					return e1
				}
				e1.PrintWarning(gen.options.ErrorWriter, nil)
			}
		}
		lrc.localNatArgs[field.FieldName] = arg
	}
	if tlType.IsFunction {
		resultResolvedType, resultResolvedTypeBare, resultNatArgs, resultHalfResolved, err := gen.getType(lrc, tlType.FuncDecl, nil)
		if err != nil {
			return err
		}
		if resultResolvedTypeBare {
			// @read a.TypeA = int;
			// @read a.TypeB = %Int;
			return tlType.FuncDecl.PR.BeautifulError(fmt.Errorf("function %q result cannot be bare", tlType.Construct.Name.String()))
		}
		res.ResultType = resultResolvedType
		res.ResultNatArgs = resultNatArgs
		res.ResultHalfResolved = resultHalfResolved
	}
	return nil
}

func (gen *Gen2) GenerateTuple(myWrapper *TypeRWWrapper, tlType *tlast.Combinator, lrc LocalResolveContext) error {
	elementT := tlast.TypeRef{Type: tlast.Name{Name: tlType.TemplateArguments[1].FieldName}} // TODO - PR
	elementResolvedType, elementResolvedTypeBare, elementNatArgs, elementHalfResolved, err := gen.getType(lrc, elementT, nil)
	if err != nil {
		return err
	}
	res := &TypeRWBrackets{
		wr:         myWrapper,
		vectorLike: false,
		element: Field{
			t:            elementResolvedType,
			bare:         elementResolvedTypeBare,
			natArgs:      elementNatArgs,
			halfResolved: elementHalfResolved,
		},
	}
	myWrapper.trw = res
	myWrapper.fileName = elementResolvedType.fileName
	res.dynamicSize = !myWrapper.arguments[0].isArith
	if myWrapper.arguments[0].isArith {
		res.size = myWrapper.arguments[0].Arith.Res
	}
	return nil
}

func (gen *Gen2) GenerateVector(myWrapper *TypeRWWrapper, tlType *tlast.Combinator, lrc LocalResolveContext, vectorElementTypeIndex int) error {
	elementT := tlast.TypeRef{Type: tlast.Name{Name: tlType.TemplateArguments[vectorElementTypeIndex].FieldName}} // TODO - PR
	elementResolvedType, elementResolvedTypeBare, elementNatArgs, elementHalfResolved, err := gen.getType(lrc, elementT, nil)
	if err != nil {
		return err
	}
	res := &TypeRWBrackets{
		wr:         myWrapper,
		vectorLike: true,
		element: Field{
			t:            elementResolvedType,
			bare:         elementResolvedTypeBare,
			natArgs:      elementNatArgs,
			halfResolved: elementHalfResolved,
		},
	}
	myWrapper.trw = res
	myWrapper.fileName = elementResolvedType.fileName
	return nil
}
