// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package pure

import (
	"fmt"

	"github.com/vkcom/tl/internal/tlast"
)

func (k *Kernel) typeCheck(tip tlast.TL2TypeDefinition, leftArgs []tlast.TL2TypeTemplate) error {
	if tip.IsAlias() {
		return k.typeCheckTypeRef(tip.TypeAlias, leftArgs)
	}
	if tip.StructType.IsUnionType {
		for _, v := range tip.StructType.UnionType.Variants {
			if err := k.typeCheckAliasFields(v.IsTypeAlias, v.TypeAlias, v.Fields, leftArgs); err != nil {
				return err
			}
		}
		return nil
	}
	return k.typeCheckAliasFields(false, tlast.TL2TypeRef{},
		tip.StructType.ConstructorFields, leftArgs)
}

func (k *Kernel) typeCheckAliasFields(isTypeAlias bool, typeAlias tlast.TL2TypeRef,
	fields []tlast.TL2Field, leftArgs []tlast.TL2TypeTemplate) error {
	if isTypeAlias {
		cat, err := k.typeCheckArgument(tlast.TL2TypeArgument{Type: typeAlias}, leftArgs)
		if err != nil {
			return err
		}
		if !cat.IsType() {
			return fmt.Errorf("type alias cannot be number")
		}
		return nil
	}
	for _, f := range fields {
		cat, err := k.typeCheckArgument(tlast.TL2TypeArgument{Type: f.Type}, leftArgs)
		if err != nil {
			return err
		}
		if !cat.IsType() {
			return fmt.Errorf("field type %s cannot be number", f.Name)
		}
	}
	return nil
}

func (k *Kernel) typeCheckTypeRef(tr tlast.TL2TypeRef, leftArgs []tlast.TL2TypeTemplate) error {
	if tr.IsBracket {
		cat, err := k.typeCheckArgument(tlast.TL2TypeArgument{Type: tr.BracketType.ArrayType}, leftArgs)
		if err != nil {
			return err
		}
		if !cat.IsType() {
			return fmt.Errorf("bracket element type cannot be number")
		}
		if tr.BracketType.HasIndex {
			_, err = k.typeCheckArgument(tr.BracketType.IndexType, leftArgs)
			if err != nil {
				return err
			}
			// can be both type (map) and number (tuple)
		}
		return nil
	}
	someType := tr.SomeType
	kt, ok := k.tips[someType.Name.String()]
	if !ok {
		return fmt.Errorf("type %s does not exist", someType.Name)
	}
	if kt.originTL2 {
		if kt.combTL2.IsFunction {
			return fmt.Errorf("cannot reference function %s", someType.Name)
		}
		if len(someType.Arguments) != len(kt.combTL2.TypeDecl.TemplateArguments) {
			return fmt.Errorf("typeref %s must have %d template arguments, has %d", someType.String(), len(kt.combTL2.TypeDecl.TemplateArguments), len(someType.Arguments))
		}
		for i, targ := range kt.combTL2.TypeDecl.TemplateArguments {
			cat, err := k.typeCheckArgument(someType.Arguments[i], leftArgs)
			if err != nil {
				return err
			}
			if targ.Category != cat {
				return fmt.Errorf("typeref %s argument %s wrong category, must be %s", someType.String(), targ.Name, targ.Category)
			}
		}
		return nil
	}
	if kt.combTL1[0].IsFunction {
		return fmt.Errorf("cannot reference function %s", someType.Name)
	}
	if len(someType.Arguments) != len(kt.combTL1[0].TemplateArguments) {
		return fmt.Errorf("typeref %s must have %d template arguments, has %d", someType.String(), len(kt.combTL1[0].TemplateArguments), len(someType.Arguments))
	}
	for i, targ := range kt.combTL1[0].TemplateArguments {
		cat, err := k.typeCheckArgument(someType.Arguments[i], leftArgs)
		if err != nil {
			return err
		}
		if targ.IsNat != (cat == tlast.TL2TypeCategoryNat) {
			return fmt.Errorf("typeref %s argument %s wrong isNat, must be %v", someType.String(), targ.FieldName, targ.IsNat)
		}
	}
	return nil
}

func (k *Kernel) typeCheckArgument(arg tlast.TL2TypeArgument, leftArgs []tlast.TL2TypeTemplate) (tlast.TL2TypeCategory, error) {
	if arg.IsNumber {
		return tlast.TL2TypeCategory{IsNatValue: true}, nil
	}
	if !arg.Type.IsBracket && arg.Type.SomeType.Name.Namespace == "" {
		for _, la := range leftArgs {
			if arg.Type.SomeType.Name.Name == la.Name {
				if len(arg.Type.SomeType.Arguments) != 0 {
					return tlast.TL2TypeCategory{}, fmt.Errorf("reference to template argument %s cannot have arguments", la.Name)
				}
				return la.Category, nil
			}
		}
		// reference to global type
	}
	if err := k.typeCheckTypeRef(arg.Type, leftArgs); err != nil {
		return tlast.TL2TypeCategory{}, err
	}
	return tlast.TL2TypeCategory{IsNatValue: false}, nil
}
