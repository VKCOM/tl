// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package gengo

import (
	"fmt"

	"github.com/VKCOM/tl/internal/pure"
	"github.com/VKCOM/tl/internal/tlast"
	"github.com/VKCOM/tl/internal/utils"
)

func (gen *genGo) generateTypePrimitive(myWrapper *TypeRWWrapper, pureType pure.TypeInstance) error {
	primitiveTypesList := []*TypeRWPrimitive{
		{
			canonicalType:  "byte",
			historicalName: "byte",
			goType:         "byte",
			writeJSONValue: "basictl.JSONWriteByte",
			readJSON2Value: gen.InternalPrefix() + "Json2ReadByte",
			resetValue:     "%s = 0",
			randomValue:    "basictl.RandomByte",
			writeValue:     "basictl.ByteWrite",
			readValue:      "basictl.ByteRead",
		}, {
			canonicalType:  "uint32",
			historicalName: "nat",
			goType:         "uint32",
			writeJSONValue: "basictl.JSONWriteUint32",
			readJSON2Value: gen.InternalPrefix() + "Json2ReadUint32",
			resetValue:     "%s = 0",
			randomValue:    "basictl.RandomUint",
			writeValue:     "basictl.NatWrite",
			readValue:      "basictl.NatRead",
		}, {
			canonicalType:  "int32",
			historicalName: "int",
			goType:         "int32",
			writeJSONValue: "basictl.JSONWriteInt32",
			readJSON2Value: gen.InternalPrefix() + "Json2ReadInt32",
			resetValue:     "%s = 0",
			randomValue:    "basictl.RandomInt",
			writeValue:     "basictl.IntWrite",
			readValue:      "basictl.IntRead",
		}, {
			canonicalType:  "int64",
			historicalName: "long",
			goType:         "int64",
			writeJSONValue: "basictl.JSONWriteInt64",
			readJSON2Value: gen.InternalPrefix() + "Json2ReadInt64",
			resetValue:     "%s = 0",
			randomValue:    "basictl.RandomLong",
			writeValue:     "basictl.LongWrite",
			readValue:      "basictl.LongRead",
		}, {
			canonicalType:  "uint64",
			historicalName: "uint64",
			goType:         "uint64",
			writeJSONValue: "basictl.JSONWriteUint64",
			readJSON2Value: gen.InternalPrefix() + "Json2ReadUint64",
			resetValue:     "%s = 0",
			randomValue:    "basictl.RandomUint64",
			writeValue:     "basictl.Uint64Write",
			readValue:      "basictl.Uint64Read",
		}, {
			canonicalType:  "float32",
			historicalName: "float",
			goType:         "float32",
			writeJSONValue: "basictl.JSONWriteFloat32",
			readJSON2Value: gen.InternalPrefix() + "Json2ReadFloat32",
			resetValue:     "%s = 0",
			randomValue:    "basictl.RandomFloat",
			writeValue:     "basictl.FloatWrite",
			readValue:      "basictl.FloatRead",
		}, {
			canonicalType:  "float64",
			historicalName: "double",
			goType:         "float64",
			writeJSONValue: "basictl.JSONWriteFloat64",
			readJSON2Value: gen.InternalPrefix() + "Json2ReadFloat64",
			resetValue:     "%s = 0",
			randomValue:    "basictl.RandomDouble",
			writeValue:     "basictl.DoubleWrite",
			readValue:      "basictl.DoubleRead",
		}, {
			canonicalType:  "string",
			historicalName: "string",
			goType:         "string",
			writeJSONValue: "basictl.JSONWriteString",
			readJSON2Value: gen.InternalPrefix() + "Json2ReadString",
			resetValue:     "%s = \"\"",
			randomValue:    "basictl.RandomString",
			writeValue:     "basictl.StringWrite",
			readValue:      "basictl.StringRead",
		}, {
			canonicalType:  "__function",
			historicalName: "function",
			goType:         "string",
			writeJSONValue: "basictl.JSONWriteString",
			readJSON2Value: gen.InternalPrefix() + "Json2ReadString",
			resetValue:     "%s = \"\"",
			randomValue:    "basictl.RandomString",
			writeValue:     "basictl.StringWrite",
			readValue:      "basictl.StringRead",
		}, {
			canonicalType:  "__function_result",
			historicalName: "function_result",
			goType:         "string",
			writeJSONValue: "basictl.JSONWriteString",
			readJSON2Value: gen.InternalPrefix() + "Json2ReadString",
			resetValue:     "%s = \"\"",
			randomValue:    "basictl.RandomString",
			writeValue:     "basictl.StringWrite",
			readValue:      "basictl.StringRead",
		},
	}
	for _, ct := range primitiveTypesList {
		if ct.canonicalType == pureType.CanonicalName() {
			myWrapper.trw = ct
			myWrapper.goCanonicalName = tlast.TL2TypeName{Name: ct.historicalName}
			return nil
		}
	}
	// this wrapper will crash if accessed.
	// TODO - add missing types (uint64, byte, bit) aboveб error below
	// TODO - bit type might simplify generators by generating no code in TypeWriting methods
	myWrapper.trw = &TypeRWPrimitive{
		canonicalType: pureType.CanonicalName(),
		goType:        pureType.CanonicalName(),
	}
	return nil // fmt.Errorf("unknown primitive type")
}

func (gen *genGo) generateTypeStruct(myWrapper *TypeRWWrapper, pureType *pure.TypeInstanceStruct,
	unionParent *TypeRWUnion, unionIndex int) error {
	head, tail := myWrapper.resolvedT2GoName("")
	myWrapper.goGlobalName = gen.globalDec.DeconflictName(head + tail)
	head, tail = myWrapper.resolvedT2GoName(myWrapper.ns.name)
	myWrapper.goLocalName = myWrapper.ns.decGo.DeconflictName(head + tail)

	//if pureType.ResultType() == nil && myWrapper.unionParent == nil && len(myWrapper.pureType.KernelType().TL1()) != 0 {
	//wasName := canonicalGoName(myWrapper.pureType.KernelType().TL1()[0].TypeDecl.Name, myWrapper.ns.name)
	//if head+tail != wasName+tail {
	//	gen.options.ReplaceStrings(".go",
	//		"tl"+myWrapper.ns.name+"."+wasName+tail,
	//		"tl"+myWrapper.ns.name+"."+myWrapper.goLocalName)
	//}
	//}

	//head, tail = myWrapper.resolvedT2GoName(myWrapper.tlName.Namespace)
	//if head+tail != myWrapper.goLocalName {
	//	gen.options.ReplaceStrings(".go",
	//		"tl"+myWrapper.ns.name+"."+head+tail,
	//		"tl"+myWrapper.ns.name+"."+myWrapper.goLocalName)
	//}

	res := &TypeRWStruct{
		wr:             myWrapper,
		pureTypeStruct: pureType,
		unionParent:    unionParent,
		unionIndex:     unionIndex,
	}
	res.fieldsDec.FillGolangIdentifies()
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
		}
		if field.Name() != "" { // empty only for typedef single field
			newField.goName = res.fieldsDec.DeconflictName(utils.CNameToCamelName(field.Name()))
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
	myWrapper.goGlobalName = gen.globalDec.DeconflictName(head + tail)
	head, tail = myWrapper.resolvedT2GoName(myWrapper.ns.name)
	myWrapper.goLocalName = myWrapper.ns.decGo.DeconflictName(head + tail)
	res := &TypeRWBool{
		isBit: isBit,
		wr:    myWrapper,
	}
	if ok, falseTag, trueTag := pureType.IsTL1Bool(); ok {
		res.falseTag = falseTag
		res.trueTag = trueTag
		res.falseGoName = gen.globalDec.DeconflictName("BoolFalse")
		res.trueGoName = gen.globalDec.DeconflictName("BoolTrue")
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
		myWrapper.goGlobalName = gen.globalDec.DeconflictName(tail + head)
		head, tail = myWrapper.resolvedT2GoName(myWrapper.ns.name)
		myWrapper.goLocalName = myWrapper.ns.decGo.DeconflictName(tail + head)

		res := &TypeRWMaybe{
			wr: myWrapper,
			element: Field{
				pureField: elementField,
				t:         fieldType,
			},
			emptyTag: pureType.VariantTypes()[0].TLTag(),
			okTag:    pureType.VariantTypes()[1].TLTag(),
		}
		myWrapper.fileNameOverride = fieldType
		myWrapper.trw = res
		return nil
	}
	head, tail := myWrapper.resolvedT2GoName("")
	myWrapper.goGlobalName = gen.globalDec.DeconflictName(head + tail)
	head, tail = myWrapper.resolvedT2GoName(myWrapper.ns.name)
	myWrapper.goLocalName = myWrapper.ns.decGo.DeconflictName(head + tail)

	res := &TypeRWUnion{
		pureType: pureType,
		wr:       myWrapper,
		IsEnum:   pureType.IsEnum(),
	}
	myWrapper.trw = res
	for i, typ := range pureType.VariantTypes() {
		variantName := pureType.VariantNames()[i]

		variantWrapper := &TypeRWWrapper{
			gen:              gen,
			pureType:         typ,
			fileNameOverride: myWrapper,
			goCanonicalName:  typ.TLName(),
			ns:               myWrapper.ns,
		}
		if myWrapper.TLName().Namespace != variantWrapper.TLName().Namespace {
			variantWrapper.ns = gen.getNamespace(variantWrapper.TLName().Namespace)
		}
		variantWrapper.ns.types = append(variantWrapper.ns.types, variantWrapper)

		gen.generatedTypes[typ.CanonicalName()] = variantWrapper
		gen.generatedTypesList = append(gen.generatedTypesList, variantWrapper)

		if err := gen.generateTypeStruct(variantWrapper, typ, res, i); err != nil {
			return err
		}

		fieldGoName := canonicalGoName(tlast.Name{Name: variantName}, "")
		newField := Variant{
			variantName: variantName,
			t:           variantWrapper,
			goName:      res.fieldsDec.DeconflictName(fieldGoName),
		}
		res.Fields = append(res.Fields, newField)
	}
	return nil
}

func (gen *genGo) GenerateTypeArray(myWrapper *TypeRWWrapper, pureType *pure.TypeInstanceArray) error {
	field := pureType.Field()
	fieldType, err := gen.getType(field.TypeInstance())
	if err != nil {
		return err
	}
	res := &TypeRWBrackets{
		wr:          myWrapper,
		vectorLike:  !pureType.IsTuple(),
		dynamicSize: pureType.DynamicSize(),
		size:        pureType.Count(),
		element: Field{
			pureField: field,
			t:         fieldType,
		},
	}
	myWrapper.trw = res
	myWrapper.fileNameOverride = fieldType
	return nil
}

func (gen *genGo) GenerateTypeDict(myWrapper *TypeRWWrapper, pureType *pure.TypeInstanceDict) error {
	field := pureType.Field()
	fieldType, err := gen.getType(field.TypeInstance())
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
		wr:            myWrapper,
		structElement: structElement,
		element: Field{
			pureField: field,
			t:         fieldType,
		},
		dictKeyString:  isString,
		dictKeyField:   structElement.Fields[0],
		dictValueField: structElement.Fields[1],
	}
	myWrapper.trw = res
	myWrapper.fileNameOverride = fieldType
	return nil
}
