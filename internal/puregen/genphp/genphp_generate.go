// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package genphp

import (
	"fmt"

	"github.com/VKCOM/tl/internal/pure"
	"github.com/VKCOM/tl/internal/tlast"
	"github.com/VKCOM/tl/internal/utils"
)

func (gen *genphp) generateTypePrimitive(myWrapper *TypeRWWrapper, pureType pure.TypeInstance) error {
	primitiveTypesList := []*TypeRWPrimitive{
		{
			gen:           gen,
			canonicalType: "byte",
			goType:        "byte",
		}, {
			gen:           gen,
			canonicalType: "uint32",
			goType:        "uint32",
		}, {
			gen:           gen,
			canonicalType: "int32",
			goType:        "int32",
		}, {
			gen:           gen,
			canonicalType: "int64",
			goType:        "int64",
		}, {
			gen:           gen,
			canonicalType: "uint64",
			goType:        "uint64",
		}, {
			gen:           gen,
			canonicalType: "float32",
			goType:        "float32",
		}, {
			gen:           gen,
			canonicalType: "float64",
			goType:        "float64",
		}, {
			gen:           gen,
			canonicalType: "string",
			goType:        "string",
		}, {
			gen:           gen,
			canonicalType: "__function",
			goType:        "string",
		}, {
			gen:           gen,
			canonicalType: "__function_result",
			goType:        "string",
		},
	}
	for _, ct := range primitiveTypesList {
		if ct.canonicalType == pureType.CanonicalName() {
			myWrapper.trw = ct
			//myWrapper.goCanonicalName = tlast.TL2TypeName{Name: ct.historicalName}
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

func (gen *genphp) generateTypeStruct(myWrapper *TypeRWWrapper, pureType *pure.TypeInstanceStruct,
	unionParent *TypeRWUnion, unionIndex int) error {

	//if pureType.ResultType() == nil && myWrapper.unionParent == nil && len(myWrapper.pureType.KernelType().TL1()) != 0 {
	//wasName := canonicalGoName(myWrapper.pureType.KernelType().TL1()[0].TypeDecl.Name, myWrapper.ns.name)
	//if head+tail != wasName+tail {
	//	gen.options.ReplaceStrings(".go",
	//		"tl"+myWrapper.ns.name+"."+wasName+tail,
	//		"tl"+myWrapper.ns.name+"."+myWrapper.goLocalName)
	//}
	//}

	//head, tail = myWrapper.resolvedT2GoName(myWrapper.TLName().Namespace)
	//if head+tail != myWrapper.goLocalName {
	//	gen.options.ReplaceStrings(".go",
	//		"tl"+myWrapper.ns.name+"."+head+tail,
	//		"tl"+myWrapper.ns.name+"."+myWrapper.goLocalName)
	//}

	myWrapper.unionParent = unionParent
	myWrapper.unionIndex = unionIndex

	res := &TypeRWStruct{
		wr:       myWrapper,
		pureType: pureType,
	}

	//res.fieldsDec.FillGolangIdentifies()
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

func (gen *genphp) generateTypeBool(myWrapper *TypeRWWrapper, pureType *pure.TypeInstancePrimitive, isBit bool) error {
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

func (gen *genphp) generateTypeUnion(myWrapper *TypeRWWrapper, pureType *pure.TypeInstanceUnion) error {
	if isMaybe, elementField := pureType.IsUnionMaybe(); isMaybe {
		fieldType, err := gen.getType(elementField.TypeInstance())
		if err != nil {
			return err
		}
		// Customizing maybe name was really stupid idea, actually.
		res := &TypeRWMaybe{
			wr: myWrapper,
			element: Field{
				pureField: elementField,
				t:         fieldType,
			},
			emptyTag: pureType.VariantTypes()[0].TLTag(),
			okTag:    pureType.VariantTypes()[1].TLTag(),
		}
		myWrapper.trw = res
		return nil
	}

	res := &TypeRWUnion{
		pureType: pureType,
		wr:       myWrapper,
		IsEnum:   pureType.IsEnum(),
	}
	myWrapper.trw = res
	for i, typ := range pureType.VariantTypes() {
		variantName := pureType.VariantNames()[i]

		variantWrapper := &TypeRWWrapper{
			gen:      gen,
			pureType: typ,
			ns:       myWrapper.ns,
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
		newField := Field{
			t:      variantWrapper,
			goName: res.fieldsDec.DeconflictName(fieldGoName),
		}
		res.Fields = append(res.Fields, newField)
	}
	return nil
}

func (gen *genphp) GenerateTypeArray(myWrapper *TypeRWWrapper, pureType *pure.TypeInstanceArray) error {
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
	//myWrapper.fileNameOverride = fieldType
	return nil
}

func (gen *genphp) GenerateTypeDict(myWrapper *TypeRWWrapper, pureType *pure.TypeInstanceDict) error {
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
	res := &TypeRWBrackets{
		wr: myWrapper,
		//structElement: structElement,
		vectorLike: true,
		dictLike:   true,
		element: Field{
			pureField: field,
			t:         fieldType,
		},
		dictKeyString:  isString,
		dictKeyField:   structElement.Fields[0],
		dictValueField: structElement.Fields[1],
	}
	myWrapper.trw = res
	return nil
}

func (gen *genphp) getType(t pure.TypeInstance) (*TypeRWWrapper, error) {
	result, ok := gen.generatedTypes[t.CanonicalName()]
	if !ok {
		return nil, fmt.Errorf("internal error: type %q not found", t.CanonicalName())
	}
	return result, nil
}

func (gen *genphp) getTypeMust(t pure.TypeInstance) *TypeRWWrapper {
	result, ok := gen.generatedTypes[t.CanonicalName()]
	if !ok {
		panic(fmt.Errorf("internal error: type instance %q not found", t.CanonicalName()))
	}
	return result
}

func (gen *genphp) getTypeWrapperMust(rt tlast.TL2TypeRef) (*TypeRWWrapper, bool) {
	ref, fieldBare, err := gen.kernel.GetInstance(rt)
	if err != nil {
		panic(fmt.Errorf("internal error: cannot get type of argument %s: %w", rt.String(), err))
	}
	return gen.getTypeMust(ref), fieldBare
}
