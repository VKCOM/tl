// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package gengo

import (
	"fmt"

	"github.com/vkcom/tl/internal/pure"
	"github.com/vkcom/tl/internal/tlast"
	"github.com/vkcom/tl/internal/utils"
)

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
			//ct.tlTag = myWrapper.tlTag
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
	head, tail = myWrapper.resolvedT2GoName(myWrapper.ns.name)
	myWrapper.goLocalName = myWrapper.ns.decGo.deconflictName(head + tail)

	if pureType.ResultType() == nil && myWrapper.unionParent == nil && len(myWrapper.pureType.KernelType().TL1()) != 0 {
		// nowName := canonicalGoName(myWrapper.goCanonicalName, myWrapper.ns.name)
		wasName := canonicalGoName(myWrapper.pureType.KernelType().TL1()[0].TypeDecl.Name, myWrapper.ns.name)
		if head+tail != wasName+tail {
			gen.options.ReplaceStrings(".go",
				"tl"+myWrapper.ns.name+"."+wasName+tail,
				"tl"+myWrapper.ns.name+"."+myWrapper.goLocalName)
		}
	}

	//head, tail = myWrapper.resolvedT2GoName(myWrapper.tlName.Namespace)
	//if head+tail != myWrapper.goLocalName {
	//	gen.options.ReplaceStrings(".go",
	//		"tl"+myWrapper.ns.name+"."+head+tail,
	//		"tl"+myWrapper.ns.name+"."+myWrapper.goLocalName)
	//}

	res := &TypeRWStruct{
		wr:             myWrapper,
		pureTypeStruct: pureType,
	}
	res.fieldsDec.fillGolangIdentifies()
	myWrapper.trw = res
	//nextTL2MaskBit := 0
	for _, field := range pureType.Fields() {
		fieldType, err := gen.getType(field.TypeInstance())
		if err != nil {
			return err
		}
		newField := Field{
			pureField: field,
			t:         fieldType,
			bare:      field.Bare(),
			natArgs:   field.NatArgs(),
		}
		if field.Name() != "" { // empty only for typedef single field
			newField.goName = res.fieldsDec.deconflictName(utils.CNameToCamelName(field.Name()))
		}
		res.Fields = append(res.Fields, newField)
	}
	if pureType.ResultType() != nil {
		resultType, err := gen.getType(pureType.ResultType())
		if err != nil {
			return err
		}
		res.ResultType = resultType
		res.ResultNatArgs = pureType.ResultNatArgs()
	}
	return nil
}

func (gen *genGo) generateTypeBool(myWrapper *TypeRWWrapper, pureType *pure.TypeInstancePrimitive, isBit bool) error {
	head, tail := myWrapper.resolvedT2GoName("")
	myWrapper.goGlobalName = gen.globalDec.deconflictName(head + tail)
	head, tail = myWrapper.resolvedT2GoName(myWrapper.ns.name)
	myWrapper.goLocalName = myWrapper.ns.decGo.deconflictName(head + tail)
	res := &TypeRWBool{
		isBit: isBit,
		wr:    myWrapper,
	}
	if ok, falseTag, trueTag := pureType.IsTL1Bool(); ok {
		res.falseTag = falseTag
		res.trueTag = trueTag
		res.falseGoName = gen.globalDec.deconflictName("BoolFalse")
		res.trueGoName = gen.globalDec.deconflictName("BoolTrue")
	}
	myWrapper.trw = res
	return nil
}

func (gen *genGo) generateTypeUnion(myWrapper *TypeRWWrapper, pureType *pure.TypeInstanceUnion) error {
	if isMaybe, elementField := pureType.IsUnionMaybe(); isMaybe {
		fieldType, err := gen.getType(elementField.TypeInstance())
		if err != nil {
			return err
		}
		// Customizing maybe name was really stupid idea, actually.
		head, tail := myWrapper.resolvedT2GoName("")
		myWrapper.goGlobalName = gen.globalDec.deconflictName(tail + head)
		head, tail = myWrapper.resolvedT2GoName(myWrapper.ns.name)
		myWrapper.goLocalName = myWrapper.ns.decGo.deconflictName(tail + head)

		res := &TypeRWMaybe{
			wr: myWrapper,
			element: Field{
				pureField: elementField,
				t:         fieldType,
				bare:      elementField.Bare(),
				natArgs:   pureType.ElementNatArgs(),
			},
			emptyTag: pureType.VariantTypes()[0].TLTag(),
			okTag:    pureType.VariantTypes()[1].TLTag(),
		}
		myWrapper.fileNameOverride = fieldType
		myWrapper.trw = res
		return nil
	}
	head, tail := myWrapper.resolvedT2GoName("")
	myWrapper.goGlobalName = gen.globalDec.deconflictName(head + tail)
	head, tail = myWrapper.resolvedT2GoName(myWrapper.ns.name)
	myWrapper.goLocalName = myWrapper.ns.decGo.deconflictName(head + tail)

	res := &TypeRWUnion{
		pureType: pureType,
		wr:       myWrapper,
		IsEnum:   pureType.IsEnum(),
	}
	myWrapper.trw = res
	kt := pureType.KernelType()
	for i, typ := range pureType.VariantTypes() {
		variantName := pureType.VariantNames()[i]

		variantWrapper := &TypeRWWrapper{
			gen:              gen,
			pureType:         typ,
			NatParams:        myWrapper.NatParams,
			originateFromTL2: kt.OriginTL2(),
			tlTag:            typ.TLTag(),
			tlName:           typ.TLName(),
			unionParent:      res,
			unionIndex:       i,
			fileNameOverride: myWrapper,
			goCanonicalName:  typ.TLName(),
			ns:               myWrapper.ns,
		}
		if myWrapper.tlName.Namespace != variantWrapper.tlName.Namespace {
			variantWrapper.ns = gen.getNamespace(variantWrapper.tlName.Namespace)
		}
		variantWrapper.ns.types = append(variantWrapper.ns.types, variantWrapper)

		gen.generatedTypes[typ.CanonicalName()] = variantWrapper
		gen.generatedTypesList = append(gen.generatedTypesList, variantWrapper)

		if err := gen.generateTypeStruct(variantWrapper, typ); err != nil {
			return err
		}

		fieldGoName := canonicalGoName(tlast.Name{Name: variantName}, "")
		newField := Variant{
			// originalName: , we do not use filed.originalName in unions generation
			t:      variantWrapper,
			goName: res.fieldsDec.deconflictName(fieldGoName),
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
	myWrapper.fileNameOverride = fieldType
	return nil
}

func (gen *genGo) GenerateTypeDict(myWrapper *TypeRWWrapper, pureType *pure.TypeInstanceDict) error {
	fieldType, err := gen.getType(pureType.FieldType())
	if err != nil {
		return err
	}
	structElement, ok := fieldType.trw.(*TypeRWStruct)
	if !ok || len(structElement.Fields) != 2 {
		return fmt.Errorf("dict %s element is not struct with 2 fields", pureType.CanonicalName())
	}
	// TODO - better check?
	isString := structElement.Fields[0].t.pureType.CanonicalName() == "string"
	res := &TypeRWDict{
		wr: myWrapper,
		element: Field{
			t:       fieldType,
			bare:    pureType.FieldBare(),
			natArgs: pureType.FieldNatArgs(),
		},
		dictKeyString:  isString,
		dictKeyField:   structElement.Fields[0],
		dictValueField: structElement.Fields[1],
	}
	myWrapper.trw = res
	myWrapper.fileNameOverride = fieldType
	return nil
}
