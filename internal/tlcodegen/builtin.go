// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlcodegen

import (
	"fmt"
	"strings"

	"github.com/vkcom/tl/internal/tlast"
)

// This code describes part of tlgen which is dependent on VK implementation

func IsUnionBool(tlType []*tlast.Combinator) (isBool bool, falseDesc *tlast.Combinator, trueDesc *tlast.Combinator) {
	// if type is
	// 1. enum with 2 elements, 0 template arguments
	// 2. has name "Bool"
	// 3. fields have names "boolFalse" and "boolTrue"
	if len(tlType) != 2 || tlType[0].TypeDecl.Name.String() != "Bool" ||
		len(tlType[0].Fields) != 0 || len(tlType[1].Fields) != 0 || len(tlType[0].TemplateArguments) != 0 {
		return false, nil, nil
	}
	falseDesc = tlType[0]
	trueDesc = tlType[1]
	if falseDesc.Construct.Name.String() != "boolFalse" { // fix constructors order
		falseDesc, trueDesc = trueDesc, falseDesc
	}
	if falseDesc.Construct.Name.String() != "boolFalse" ||
		trueDesc.Construct.Name.String() != "boolTrue" {
		return false, nil, nil
	}
	return true, falseDesc, trueDesc
}

func IsUnionMaybe(tlType []*tlast.Combinator) (isMaybe bool, emptyDesc *tlast.Combinator, okDesc *tlast.Combinator) {
	// if type is
	// 1. union with 1 template Type arguments && 2 fields
	// 2. one field is empty, another field has itself 1 field with type from argument
	// 3. has "maybe" name
	// then it is maybe
	// reverse = false if first element is empty
	if len(tlType) != 2 || strings.ToLower(tlType[0].TypeDecl.Name.Name) != "maybe" || len(tlType) != 2 {
		return false, nil, nil
	}
	if len(tlType[0].TemplateArguments) != 1 || len(tlType[1].TemplateArguments) != 1 {
		return false, nil, nil
	}
	if tlType[0].TemplateArguments[0].IsNat || tlType[1].TemplateArguments[0].IsNat {
		return false, nil, nil
	}
	okDesc = tlType[0]
	emptyDesc = tlType[1]
	if len(tlType[0].Fields) == 0 {
		emptyDesc, okDesc = okDesc, emptyDesc
	}
	if len(emptyDesc.Fields) != 0 || len(okDesc.Fields) != 1 ||
		okDesc.Fields[0].FieldType.String() != okDesc.TemplateArguments[0].FieldName || okDesc.Fields[0].Mask != nil {
		return false, nil, nil
	}
	return true, emptyDesc, okDesc
}

// all non-trivial contents of [] is turned into new types
// we make copy deep anough to not affect original constructors
func (gen *Gen2) ReplaceSquareBracketsElem(tl tlast.TL) (tlast.TL, error) {
	tl = append([]*tlast.Combinator{}, tl...)
	constructorTags := map[uint32]*tlast.Combinator{}
	constructorNames := map[string]*tlast.Combinator{}
	typeNames := map[string]*tlast.Combinator{}
	for _, typ := range tl {
		constructorTags[typ.Crc32()] = typ                  // overwrite if same, collision check is in another place
		constructorNames[typ.Construct.Name.String()] = typ // overwrite if same, collision check is in another place
		typeNames[typ.TypeDecl.Name.String()] = typ         // overwrite if same, collision check is in another place
		for _, field := range typ.Fields {
			if field.Excl && !LegacyEnableExclamation(typ.Construct.Name.String()) {
				return tl, field.PR.BeautifulError(fmt.Errorf("new !X function wrappers are forbidden"))
			}
		}
	}
	for typeIndex := 0; typeIndex < len(tl); typeIndex++ { // We append anonymous types while iterating
		typ := tl[typeIndex]
		var newFields []tlast.Field
		replaceRepeated := func(toVector bool, insideField tlast.Field, originalCommentRight string) (tlast.TypeRef, error) {
			if len(insideField.ScaleRepeat.Rep) == 0 {
				return tlast.TypeRef{}, insideField.ScaleRepeat.PR.BeautifulError(fmt.Errorf("repetition with no fields is not allowed"))
			}
			if !insideField.ScaleRepeat.ExplicitScale {
				endRange := tlast.PositionRange{Outer: insideField.ScaleRepeat.PR.Outer, Begin: insideField.ScaleRepeat.PR.Begin, End: insideField.ScaleRepeat.PR.Begin}
				return tlast.TypeRef{}, endRange.BeautifulError(fmt.Errorf("anonymous scale repeat can be used only in top-level square brackets"))
			}
			rep := tlast.ArithmeticOrType{T: tlast.TypeRef{Type: tlast.Name{Name: insideField.ScaleRepeat.Scale.Scale}}}
			if insideField.ScaleRepeat.Scale.IsArith {
				// TODO - wrong PR here
				rep = tlast.ArithmeticOrType{IsArith: true, Arith: insideField.ScaleRepeat.Scale.Arith}
			}
			tWithArgs := insideField.ScaleRepeat.Rep[0].FieldType
			if len(insideField.ScaleRepeat.Rep) != 1 || insideField.ScaleRepeat.Rep[0].FieldName != "" || insideField.ScaleRepeat.Rep[0].Mask != nil || insideField.ScaleRepeat.Rep[0].IsRepeated {
				return tWithArgs, insideField.ScaleRepeat.PR.BeautifulError(fmt.Errorf("brackets must contain reference to single type, fields are not allowed here"))
			}
			if toVector {
				newFieldType := tlast.TypeRef{
					Type: tlast.Name{Name: BuiltinVectorName},
					Bare: true,
					Args: []tlast.ArithmeticOrType{
						{T: tWithArgs},
					},
				}
				return newFieldType, nil
			}
			newFieldType := tlast.TypeRef{
				Type: tlast.Name{Name: BuiltinTupleName},
				Bare: true,
				Args: []tlast.ArithmeticOrType{
					rep,
					{T: tWithArgs},
				},
			}
			return newFieldType, nil
		}
		for fieldIndex, field := range typ.Fields {
			if !field.IsRepeated {
				newFields = append(newFields, field)
				continue
			}
			newField := field
			// It is hard to not destroy original fields in recursive algo, so we destroy them
			toVector := false
			if !newField.ScaleRepeat.ExplicitScale {
				endRange := tlast.PositionRange{Outer: newField.ScaleRepeat.PR.Outer, Begin: newField.ScaleRepeat.PR.Begin, End: newField.ScaleRepeat.PR.Begin}
				if fieldIndex == 0 { // Allow shortcut to last template parameters
					if len(typ.TemplateArguments) == 0 {
						// hren a:[int] = Hren;
						return nil, endRange.BeautifulError(fmt.Errorf("anonymous scale repeat implicitly references non-existing template parameter"))
					}
					// hren {n:#} a:[int] = Hren n;
					a := typ.TemplateArguments[len(typ.TemplateArguments)-1]
					if !a.IsNat {
						e1 := endRange.BeautifulError(fmt.Errorf("anonymous scale repeat implicitly references last template parameter %q which should have type #", a.FieldName))
						e2 := a.PR.BeautifulError(fmt.Errorf("see here"))
						return nil, tlast.BeautifulError2(e1, e2)
					}
					newField.ScaleRepeat.Scale.Scale = a.FieldName
				} else {
					prevField := newFields[len(newFields)-1]
					if prevField.FieldType.Type.String() != "#" {
						e1 := endRange.BeautifulError(fmt.Errorf("anonymous scale repeat implicitly references previous field %q which should have type #", prevField.FieldName))
						e2 := prevField.FieldType.PR.BeautifulError(fmt.Errorf("see here"))
						return nil, tlast.BeautifulError2(e1, e2)
					}
					newField.ScaleRepeat.Scale.Scale = prevField.FieldName
					if prevField.FieldName == "" {
						// we replace 2 fields with vector
						// hren # a:[int] = Hren;
						toVector = true
						newFields = newFields[:len(newFields)-1]
					}
					// hren n:# a:[int] = Hren n;
				}
				newField.ScaleRepeat.ExplicitScale = true
			}
			var err error
			if newField.FieldType, err = replaceRepeated(toVector, newField, newField.CommentRight); err != nil {
				return nil, err
			}
			newField.IsRepeated = false
			newFields = append(newFields, newField)
		}
		newDesc := &tlast.Combinator{}
		*newDesc = *typ
		newDesc.OriginalDescriptor = typ
		newDesc.Fields = newFields
		tl[typeIndex] = newDesc
	}
	return tl, nil
}
