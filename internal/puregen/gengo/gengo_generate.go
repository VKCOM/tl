// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package gengo

import (
	"github.com/vkcom/tl/internal/pure"
	"github.com/vkcom/tl/internal/tlast"
	"github.com/vkcom/tl/internal/utils"
)

//	func collectArgsNamespaces(tip *TypeRWWrapper, argNamespaces map[string]struct{}) {
//		// This is policy. You can adjust it, so more or less templates instantiations
//		// are moved into namespace of arguments. Code should compile anyway.
//		if tip.tlName.Namespace != "" {
//			argNamespaces[tip.tlName.Namespace] = struct{}{}
//		}
//		for _, arg := range tip.arguments {
//			if !arg.isNat {
//				collectArgsNamespaces(arg.tip, argNamespaces)
//			}
//		}
//	}
//
//	func (gen *genGo) generateType(myWrapper *TypeRWWrapper) error {
//		tlType := myWrapper.origTL
//		argNamespaces := map[string]struct{}{}
//		for i, a := range tlType[0].TemplateArguments { // they are the same for all constructors
//			ra := myWrapper.arguments[i]
//			if a.IsNat {
//				lrc.localNatArgs[a.FieldName] = LocalNatArg{
//					NamePR: a.PR,
//					TypePR: a.PR,
//					natArg: ActualNatArg{
//						isArith: ra.isArith,
//						Arith:   ra.Arith,
//						name:    a.FieldName,
//					},
//				}
//				if !ra.isArith {
//					myWrapper.NatParams = append(myWrapper.NatParams, a.FieldName)
//				}
//				continue
//			}
//			collectArgsNamespaces(ra.tip, argNamespaces)
//			natArgs := ra.tip.NatArgs(nil, a.FieldName)
//			// We can select arbitrary names for arguments here, but they all must be unique per generatedType
//			// The simplest idea to avoid collisions is to exploit uniqueness of field names
//			if len(natArgs) == 1 { // in most common case avoid longer than necessary names.
//				natArgs[0].name = a.FieldName
//			}
//			// We decided to use internal structure of names instead of assigning them sequential numbers for each template argument.
//			// You can uncomment code below to see which naming scheme looks better
//			//	for j := range natArgs {
//			//		natArgs[j].name = fmt.Sprintf("%s%d", a.FieldName, j)
//			//	}
//			for _, natArg := range natArgs {
//				myWrapper.NatParams = append(myWrapper.NatParams, natArg.name)
//			}
//		}
//		replaceNamespace := func(n string) *Namespace {
//			newNamespace := n
//			if n == "" && len(argNamespaces) == 1 {
//				for ns := range argNamespaces {
//					newNamespace = ns
//				}
//			}
//			return gen.getNamespace(newNamespace)
//		}
//		//myWrapper.cppNamespaceQualifier = "::" + gen.options.RootCPPNamespace + "::"
//		//if rt2.Type.Namespace != "" {
//		//	myWrapper.cppNamespaceQualifier += rt2.Type.Namespace + "::"
//		//}
//		if len(tlType) == 1 {
//			myWrapper.tlName = tlType[0].Construct.Name
//			myWrapper.fileName = tlType[0].Construct.Name.String()
//			namespace := replaceNamespace(myWrapper.tlName.Namespace)
//			namespace.types = append(namespace.types, myWrapper)
//			myWrapper.ns = namespace
//			myWrapper.tlTag = tlType[0].Crc32()
//			switch tlType[0].Construct.Name.String() { // TODO - better switch
//			case BuiltinTupleName:
//				// tl2 faking vector
//				//if myWrapper.arguments[0].isArith && myWrapper.arguments[0].isTL2FakeArith {
//				//	myWrapper.arguments = myWrapper.arguments[1:]
//				//	_, tail := myWrapper.resolvedT2GoName("")
//				//	myWrapper.goGlobalName = gen.globalDec.deconflictName("BuiltinVector" + tail)
//				//	return gen.GenerateVector(myWrapper, tlType[0], lrc, 1)
//				//}
//				_, tail := myWrapper.resolvedT2GoName("")
//				myWrapper.goGlobalName = gen.globalDec.deconflictName("BuiltinTuple" + tail)
//				// built-in tuple has no local name. TODO - invent one?
//				return gen.GenerateTuple(myWrapper, tlType[0], lrc)
//			case BuiltinVectorName:
//				_, tail := myWrapper.resolvedT2GoName("")
//				myWrapper.goGlobalName = gen.globalDec.deconflictName("BuiltinVector" + tail)
//				// built-in vector has no local name. TODO - invent one?
//				return gen.GenerateVector(myWrapper, tlType[0], lrc, 0)
//			}
//			head, tail := myWrapper.resolvedT2GoName("")
//			myWrapper.goGlobalName = gen.globalDec.deconflictName(head + tail)
//			head, tail = myWrapper.resolvedT2GoName(myWrapper.tlName.Namespace)
//			myWrapper.goLocalName = namespace.decGo.deconflictName(head + tail)
//			actualName, canonicalName, _ := myWrapper.cppTypeStringInNamespace(false, &DirectIncludesCPP{ns: map[*TypeRWWrapper]CppIncludeInfo{}}, false, HalfResolvedArgument{})
//			otherRW, ok := namespace.cppTemplates[canonicalName]
//			if ok {
//				myWrapper.cppLocalName = otherRW.cppLocalName
//			} else {
//				//myWrapper.cppLocalName = namespace.decCpp.deconflictName(ToUpperFirst(actualName) + tail)
//				myWrapper.cppLocalName = namespace.decCpp.deconflictName(ToUpperFirst(actualName))
//				namespace.cppTemplates[canonicalName] = myWrapper
//			}
//			return gen.generateTypeStruct(lrc, myWrapper, tlType[0])
//		}
//		myWrapper.tlName = tlType[0].TypeDecl.Name
//		{
//			// during migration to TL2 unions (UpperCase) become normal types (lowerCase)
//			// and we want to make sure diff is trivial
//			fileName := tlType[0].TypeDecl.Name
//			fileName.Name = ToLowerFirst(fileName.Name)
//			myWrapper.fileName = fileName.String()
//		}
//		if isBool, falseDesc, trueDesc := IsUnionBool(tlType); isBool { // TODO - test if parts of Bool are in different namespaces
//			namespace := replaceNamespace(myWrapper.tlName.Namespace)
//			namespace.types = append(namespace.types, myWrapper)
//			myWrapper.ns = namespace
//
//			head, tail := myWrapper.resolvedT2GoName("")
//			myWrapper.goGlobalName = gen.globalDec.deconflictName(head + tail)
//			head, tail = myWrapper.resolvedT2GoName(myWrapper.tlName.Namespace)
//			myWrapper.goLocalName = namespace.decGo.deconflictName(head + tail)
//			myWrapper.trw = &TypeRWBool{
//				isBit:       false,
//				wr:          myWrapper,
//				falseGoName: gen.globalDec.deconflictName(utils.CNameToCamelName(falseDesc.Construct.Name.String())),
//				trueGoName:  gen.globalDec.deconflictName(utils.CNameToCamelName(trueDesc.Construct.Name.String())),
//				falseTag:    falseDesc.Crc32(),
//				trueTag:     trueDesc.Crc32(),
//			}
//			return nil
//		}
//		if isMaybe, emptyDesc, okDesc := IsUnionMaybe(tlType); isMaybe {
//			elementT := tlast.TypeRef{Type: tlast.Name{Name: okDesc.TemplateArguments[0].FieldName}} // TODO - PR
//			elementResolvedType, elementResolvedTypeBare, elementNatArgs, elementHalfResolved, err := gen.getType(lrc, elementT, nil)
//			if err != nil {
//				return err
//			}
//
//			namespace := replaceNamespace(elementResolvedType.tlName.Namespace)
//			namespace.types = append(namespace.types, myWrapper)
//			myWrapper.ns = namespace
//
//			suffix := ifString(elementResolvedTypeBare, "Maybe", "BoxedMaybe")
//			head, tail := elementResolvedType.resolvedT2GoName("")
//			myWrapper.goGlobalName = gen.globalDec.deconflictName(head + tail + suffix)
//			head, tail = elementResolvedType.resolvedT2GoName(elementResolvedType.tlName.Namespace)
//			myWrapper.goLocalName = namespace.decGo.deconflictName(head + tail + suffix)
//
//			res := &TypeRWMaybe{
//				wr: myWrapper,
//				element: Field{
//					t:            elementResolvedType,
//					bare:         elementResolvedTypeBare,
//					natArgs:      elementNatArgs,
//					halfResolved: elementHalfResolved,
//				},
//				emptyTag: emptyDesc.Crc32(),
//				okTag:    okDesc.Crc32(),
//			}
//			myWrapper.fileName = elementResolvedType.fileName
//			myWrapper.trw = res
//			return nil
//		}
//		isEnum := true
//		for _, typ := range tlType {
//			isEnum = isEnum && len(typ.Fields) == 0
//		}
//
//		namespace := replaceNamespace(myWrapper.tlName.Namespace)
//		namespace.types = append(namespace.types, myWrapper)
//		myWrapper.ns = namespace
//
//		head, tail := myWrapper.resolvedT2GoName("")
//		myWrapper.goGlobalName = gen.globalDec.deconflictName(head + tail)
//		head, tail = myWrapper.resolvedT2GoName(myWrapper.tlName.Namespace)
//		myWrapper.goLocalName = namespace.decGo.deconflictName(head + tail)
//		actualName, canonicalName, _ := myWrapper.cppTypeStringInNamespace(false, &DirectIncludesCPP{ns: map[*TypeRWWrapper]CppIncludeInfo{}}, false, HalfResolvedArgument{})
//		otherRW, ok := namespace.cppTemplates[canonicalName]
//		if ok {
//			myWrapper.cppLocalName = otherRW.cppLocalName
//		} else {
//			myWrapper.cppLocalName = namespace.decCpp.deconflictName(ToUpperFirst(actualName))
//			namespace.cppTemplates[canonicalName] = myWrapper
//		}
//
//		lrc.allowAnyConstructor = true
//		lrc.overrideFileName = myWrapper.fileName
//		if gen.options.Language == "cpp" {
//			if isEnum {
//				lrc.overrideFileName += "Items" // TODO - in C++ when items and union are in the same file, they must be sorted which is hard for us
//			} else {
//				lrc.overrideFileName = "" // each type must be in its own file to break circular dependencies
//			}
//		}
//
//		res := &TypeRWUnion{
//			wr:     myWrapper,
//			IsEnum: isEnum,
//		}
//		res.fieldsDecCPP.fillCPPIdentifiers()
//		res.fieldsDec.fillGolangIdentifies()
//		myWrapper.trw = res
//
//		// Removing prefix/suffix common with union name.
//		// We temporarily allow relaxed case match. To use strict match, remove all strings.ToLower() calls below
//		typePrefix := strings.ToLower(ToLowerFirst(tlType[0].TypeDecl.Name.Name))
//		typeSuffix := strings.ToLower(tlType[0].TypeDecl.Name.Name)
//		for _, typ := range tlType {
//			conName := strings.ToLower(typ.Construct.Name.Name)
//			// if constructor is full prefix of type, we will shorten accessors
//			// ab.saveStateOne = ab.SaveState; // item.AsOne()
//			// ab.saveStateTwo = ab.SaveState; // item.AsTwo()
//			if !strings.HasPrefix(conName, typePrefix) { // same check as in checkUnionElementsCompatibility
//				typePrefix = ""
//			}
//			if !strings.HasSuffix(conName, typeSuffix) {
//				typeSuffix = ""
//			}
//		}
//		for i, typ := range tlType {
//			// ---- We treat
//			// ab.empty = ab.Response;
//			// ab.code x:int = ab.Response;
//			// ab.response x:int str:string = ab.Response;
//			// ---- roughly as
//			// ab.empty = _;
//			// ab.code x:int = _;
//			// ab.response x:int str:string = _;
//			// _ tag:# empty:ab.empty code:ab.code response: ab.response = ab.Response;
//			fieldType := tlast.TypeRef{
//				Type:   typ.Construct.Name,
//				Bare:   true,
//				PR:     typ.Construct.NamePR,
//				PRArgs: typ.TemplateArgumentsPR,
//			}
//			for _, arg := range typ.TemplateArguments {
//				fieldType.Args = append(fieldType.Args, tlast.ArithmeticOrType{
//					IsArith: false,
//					T: tlast.TypeRef{
//						Type:   tlast.Name{Name: arg.FieldName},
//						PR:     arg.PR,
//						PRArgs: arg.PR.CollapseToEnd(),
//					},
//				})
//			}
//			fieldResolvedType, fieldResolvedTypeBare, fieldNatArgs, fieldHalfResolved, err := gen.getType(lrc, fieldType, res)
//			if err != nil {
//				return err
//			}
//			// fieldResolvedType.unionParent = res // Already set by getType
//			fieldResolvedType.unionIndex = i
//			if !fieldResolvedTypeBare {
//				return fieldType.PR.BeautifulError(fmt.Errorf("union element resolved type %q cannot be boxed", fieldResolvedType.CanonicalStringTop()))
//			}
//			typeConstructName := typ.Construct.Name
//			if typePrefix != "" && len(typePrefix) < len(typeConstructName.Name) {
//				typeConstructName.Name = typeConstructName.Name[len(typePrefix):]
//			} else if typeSuffix != "" && len(typeSuffix) < len(typeConstructName.Name) {
//				typeConstructName.Name = typeConstructName.Name[:len(typeConstructName.Name)-len(typeSuffix)]
//			}
//			fieldGoName := canonicalGoName(typeConstructName, typ.Construct.Name.Namespace)
//			if res.fieldsDec.hasConflict(fieldGoName) { // try global, if local is already used
//				fieldGoName = canonicalGoName(typeConstructName, "")
//			}
//			fieldCPPName := canonicalCPPName(typeConstructName, typ.Construct.Name.Namespace)
//			if res.fieldsDecCPP.hasConflict(fieldCPPName) { // try global, if local is already used
//				fieldCPPName = canonicalCPPName(typeConstructName, "")
//			}
//			newField := Field{
//				originalName: fieldType.Type.String(),
//				t:            fieldResolvedType,
//				bare:         fieldResolvedTypeBare,
//				goName:       res.fieldsDec.deconflictName(fieldGoName),
//				cppName:      res.fieldsDecCPP.deconflictName(fieldCPPName),
//				natArgs:      fieldNatArgs,
//				halfResolved: fieldHalfResolved,
//				// origTL:       ?, // We do not want to set it here for now
//			}
//			res.Fields = append(res.Fields, newField)
//		}
//		return nil
//	}
func (gen *genGo) generateTypePrimitive(myWrapper *TypeRWWrapper, pureType pure.TypeInstance) error {
	primitiveTypesList := []*TypeRWPrimitive{
		{
			tlType:         "#",
			goType:         "uint32",
			writeJSONValue: "basictl.JSONWriteUint32",
			readJSONValue:  gen.InternalPrefix() + "JsonReadUint32",
			readJSON2Value: gen.InternalPrefix() + "Json2ReadUint32",
			resetValue:     "%s = 0",
			randomValue:    "basictl.RandomUint",
			writeValue:     "basictl.NatWrite",
			readValue:      "basictl.NatRead",
		}, {
			tlType:         "int",
			goType:         "int32",
			writeJSONValue: "basictl.JSONWriteInt32",
			readJSONValue:  gen.InternalPrefix() + "JsonReadInt32",
			readJSON2Value: gen.InternalPrefix() + "Json2ReadInt32",
			resetValue:     "%s = 0",
			randomValue:    "basictl.RandomInt",
			writeValue:     "basictl.IntWrite",
			readValue:      "basictl.IntRead",
		}, {
			tlType:         "long",
			goType:         "int64",
			writeJSONValue: "basictl.JSONWriteInt64",
			readJSONValue:  gen.InternalPrefix() + "JsonReadInt64",
			readJSON2Value: gen.InternalPrefix() + "Json2ReadInt64",
			resetValue:     "%s = 0",
			randomValue:    "basictl.RandomLong",
			writeValue:     "basictl.LongWrite",
			readValue:      "basictl.LongRead",
		}, {
			tlType:         "float",
			goType:         "float32",
			writeJSONValue: "basictl.JSONWriteFloat32",
			readJSONValue:  gen.InternalPrefix() + "JsonReadFloat32",
			readJSON2Value: gen.InternalPrefix() + "Json2ReadFloat32",
			resetValue:     "%s = 0",
			randomValue:    "basictl.RandomFloat",
			writeValue:     "basictl.FloatWrite",
			readValue:      "basictl.FloatRead",
		}, {
			tlType:         "double",
			goType:         "float64",
			writeJSONValue: "basictl.JSONWriteFloat64",
			readJSONValue:  gen.InternalPrefix() + "JsonReadFloat64",
			readJSON2Value: gen.InternalPrefix() + "Json2ReadFloat64",
			resetValue:     "%s = 0",
			randomValue:    "basictl.RandomDouble",
			writeValue:     "basictl.DoubleWrite",
			readValue:      "basictl.DoubleRead",
		}, {
			tlType:         "string",
			goType:         "string",
			writeJSONValue: "basictl.JSONWriteString",
			readJSONValue:  gen.InternalPrefix() + "JsonReadString",
			readJSON2Value: gen.InternalPrefix() + "Json2ReadString",
			resetValue:     "%s = \"\"",
			randomValue:    "basictl.RandomString",
			writeValue:     "basictl.StringWrite",
			readValue:      "basictl.StringRead",
		},
	}
	for _, ct := range primitiveTypesList {
		if ct.goType == pureType.CanonicalName() {
			myWrapper.trw = ct
			ct.tlTag = myWrapper.tlTag
			return nil
		}
	}
	// this wrapper will crash if accessed. TODO - better idea
	myWrapper.trw = &TypeRWPrimitive{
		tlType: pureType.CanonicalName(),
		goType: pureType.CanonicalName(),
	}
	return nil // fmt.Errorf("unknown primitive type")
}

func (gen *genGo) generateTypeStruct(myWrapper *TypeRWWrapper, pureType *pure.TypeInstanceStruct) error {
	head, tail := myWrapper.resolvedT2GoName("")
	myWrapper.goGlobalName = gen.globalDec.deconflictName(head + tail)
	head, tail = myWrapper.resolvedT2GoName(myWrapper.tlName.Namespace)
	myWrapper.goLocalName = myWrapper.ns.decGo.deconflictName(head + tail)
	res := &TypeRWStruct{
		wr: myWrapper,
	}
	res.fieldsDec.fillGolangIdentifies()
	myWrapper.trw = res
	//nextTL2MaskBit := 0
	for _, field := range pureType.Fields() {
		fieldType, err := gen.getType(field.TypeInstance())
		if err != nil {
			return err
		}
		fieldName := field.Name()
		if fieldName == "" {
			// only for typedefs, but TODO - harmonize condition with func (trw *TypeRWStruct) isTypeDef()
			// TODO - should never be necessary
			fieldName = "a"
		}
		newField := Field{
			originalName: field.Name(),
			t:            fieldType,
			bare:         field.Bare(),
			goName:       res.fieldsDec.deconflictName(utils.CNameToCamelName(fieldName)),
			fieldMask:    field.FieldMask(),
			BitNumber:    field.BitNumber(),
			natArgs:      field.NatArgs(),
			// origTL:       field, - TODO
		}
		//TODO - move into pure kernel
		//if field.Mask != nil {
		//	if field.Mask.BitNumber >= 32 {
		//		return field.Mask.PRBits.BeautifulError(fmt.Errorf("bitmask (%d) must be in range [0..31]", field.Mask.BitNumber))
		//	}
		//	newField.BitNumber = field.Mask.BitNumber
		//	localArg, ok := lrc.localNatArgs[field.Mask.MaskName]
		//	if !ok {
		//		return field.Mask.PRName.BeautifulError(fmt.Errorf("failed to resolve field mask %q reference", field.Mask.MaskName))
		//	}
		//	if localArg.wrongTypeErr != nil {
		//		e1 := field.Mask.PRName.BeautifulError(fmt.Errorf("field mask %q reference to field of wrong type", field.Mask.MaskName))
		//		e2 := localArg.TypePR.BeautifulError(localArg.wrongTypeErr)
		//		return tlast.BeautifulError2(e1, e2)
		//	}
		//	newField.fieldMask = &localArg.natArg
		//	maskBit := nextTL2MaskBit
		//	newField.MaskTL2Bit = &maskBit
		//	nextTL2MaskBit++
		//}
		res.Fields = append(res.Fields, newField)

		//TODO - move into pure kernel
		//if newField.t.IsTrueType() && !newField.Bare() &&
		//	newField.t.origTL[0].TypeDecl.Name.String() == "True" &&
		//	newField.t.origTL[0].Construct.Name.String() == "true" &&
		//	!LegacyAllowTrueBoxed(myWrapper.origTL[0].Construct.Name.String(), field.FieldName) &&
		//	doLint(field.CommentRight) {
		//	// We compare type by name, because there is examples of other true types which are to be extended
		//	// to unions or have added fields in the future
		//	e1 := field.FieldType.PR.BeautifulError(fmt.Errorf("true type fields should be bare, use 'true' or '%%True' instead"))
		//	if gen.options.WarningsAreErrors {
		//		return e1
		//	}
		//	e1.PrintWarning(gen.options.ErrorWriter, nil)
		//}
		//if _, ok := newField.t.trw.(*TypeRWBool); ok {
		//	if newField.t.origTL[0].TypeDecl.Name.String() == "Bool" &&
		//		newField.fieldMask != nil && !newField.fieldMask.isArith && newField.fieldMask.isField &&
		//		!LegacyAllowBoolFieldsmask(myWrapper.origTL[0].Construct.Name.String(), field.FieldName) &&
		//		doLint(field.CommentRight) {
		//		// We compare type by name to make warning more narrow at first.
		//		e1 := field.FieldType.PR.BeautifulError(fmt.Errorf("using Bool type under fields mask produces 3rd state, you probably want to use 'true' instead of 'Bool'"))
		//		if gen.options.WarningsAreErrors {
		//			return e1
		//		}
		//		e1.PrintWarning(gen.options.ErrorWriter, nil)
		//	}
		//}
	}
	//if tlType.IsFunction {
	//	resultResolvedType, resultResolvedTypeBare, resultNatArgs, resultHalfResolved, err := gen.getType(lrc, tlType.FuncDecl, nil)
	//	if err != nil {
	//		return err
	//	}
	//	if resultResolvedTypeBare {
	//		// @read a.TypeA = int;
	//		// @read a.TypeB = %Int;
	//		return tlType.FuncDecl.PR.BeautifulError(fmt.Errorf("function %q result cannot be bare", tlType.Construct.Name.String()))
	//	}
	//	res.ResultType = resultResolvedType
	//	res.ResultNatArgs = resultNatArgs
	//}
	return nil
}

func (gen *genGo) generateTypeUnion(myWrapper *TypeRWWrapper, pureType *pure.TypeInstanceUnion) error {
	res := &TypeRWUnion{
		wr:     myWrapper,
		IsEnum: pureType.IsEnum(),
	}
	myWrapper.trw = res
	for i, typ := range pureType.VariantTypes() {
		variantName := pureType.VariantNames()[i]

		variantWrapper := &TypeRWWrapper{
			gen:         gen,
			pureType:    pureType,
			NatParams:   pureType.Common().NatParams,
			unionParent: res,
			unionIndex:  i,
		}
		gen.generatedTypes[typ.CanonicalName()] = variantWrapper
		gen.generatedTypesList = append(gen.generatedTypesList, variantWrapper)

		kt := pureType.KernelType()
		if kt.OriginTL2() {
			variantWrapper.originateFromTL2 = kt.OriginTL2()
		} else {
			variantWrapper.origTL = append(variantWrapper.origTL, kt.TL1()[i])
			variantWrapper.tlTag = variantWrapper.origTL[0].Crc32()
			variantWrapper.tlName = variantWrapper.origTL[0].Construct.Name
			variantWrapper.fileName = myWrapper.fileName
		}
		namespace := gen.getNamespace(variantWrapper.tlName.Namespace)
		namespace.types = append(namespace.types, variantWrapper)
		variantWrapper.ns = namespace

		if err := gen.generateTypeStruct(variantWrapper, typ); err != nil {
			return err
		}

		fieldGoName := canonicalGoName(tlast.Name(variantName), variantName.Namespace)
		if res.fieldsDec.hasConflict(fieldGoName) { // try global, if local is already used
			fieldGoName = canonicalGoName(tlast.Name(variantName), "")
		}
		newField := Field{
			originalName: canonicalGoName(tlast.Name(variantName), ""),
			t:            variantWrapper,
			bare:         true,
			goName:       res.fieldsDec.deconflictName(fieldGoName),
			natArgs:      pureType.ElementNatArgs(),
			// origTL:       ?, // We do not want to set it here for now
		}
		res.Fields = append(res.Fields, newField)
	}
	return nil
}

func (gen *genGo) GenerateTypeArray(myWrapper *TypeRWWrapper, pureType *pure.TypeInstanceArray) error {
	fieldType, err := gen.getType(pureType.ElemType())
	if err != nil {
		return err
	}
	res := &TypeRWBrackets{
		wr:          myWrapper,
		vectorLike:  !pureType.IsTuple(),
		dynamicSize: pureType.DynamicSize(),
		size:        pureType.Count(),
		element: Field{
			t:       fieldType,
			bare:    pureType.ElemBare(),
			natArgs: pureType.ElemNatArgs(),
		},
	}
	myWrapper.trw = res
	myWrapper.fileName = fieldType.fileName
	return nil
}
