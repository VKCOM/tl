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

func (k *Kernel) typeCheckAliasFieldsTL1(fields []tlast.Field, leftArgs []tlast.TemplateArgument) error {
	//for _, f := range fields {
	//	if f.IsRepeated && f.ScaleRepeat.ExplicitScale && !f.ScaleRepeat.Scale.IsArith {
	//		arg := tlast.ArithmeticOrType{T: tlast.TypeRef{Type: tlast.Name{Name: f.ScaleRepeat.Scale.Scale}}}
	//		isNat, err := k.typeCheckArgumentTL1(arg, leftArgs)
	//		if err != nil {
	//			return err
	//		}
	//		if !isNat {
	//			return fmt.Errorf("scale repeat %s cannot be type", f.ScaleRepeat.Scale.Scale)
	//		}
	//	}
	//	// TODO - if !f.ScaleRepeat.ExplicitScale - must be only in vector and tuple definitions, which we should skip and use internal representations
	//	// TODO - the same is with various Dictionaries
	//	arg := tlast.ArithmeticOrType{T: f.FieldType}
	//	isNat, err := k.typeCheckArgumentTL1(arg, leftArgs)
	//	if err != nil {
	//		return err
	//	}
	//	if isNat {
	//		return fmt.Errorf("field type %s cannot be number", f.FieldType.String())
	//	}
	//	if f.FieldName != "" && f.FieldType.String() == "#" {
	//		// TODO - add other fields with wrong category to catch references to them
	//		leftArgs = append(leftArgs, tlast.TemplateArgument{
	//			FieldName: f.FieldName,
	//			IsNat:     true,
	//			PR:        f.PR,
	//		})
	//	}
	//}
	return nil
}

func (k *Kernel) typeCheckTypeRefTL1(tr tlast.TypeRef, leftArgs []tlast.TemplateArgument) error {
	typeName := tr.Type.String()
	// we remap TL1 type names into TL2 type names here.
	// hopefully, this does not cause too many irregularities downstream
	switch typeName {
	case "int":
		typeName = "int32"
	case "long":
		typeName = "int64"
	case "float":
		typeName = "float32"
	case "double":
		typeName = "float64"
	case "#":
		typeName = "uint32"
	}
	kt, ok := k.tips[typeName]
	if !ok {
		return fmt.Errorf("type %s does not exist", tr.Type)
	}
	if kt.originTL2 {
		return fmt.Errorf("TL1 combinator cannot reference TL2 combinator %s", tr.Type)
	}
	if kt.combTL1[0].IsFunction {
		return fmt.Errorf("cannot reference function %s", tr.Type)
	}
	if len(tr.Args) != len(kt.combTL1[0].TemplateArguments) {
		return fmt.Errorf("typeref %s must have %d template arguments, has %d", tr.Type, len(kt.combTL1[0].TemplateArguments), len(tr.Args))
	}
	for i, targ := range kt.combTL1[0].TemplateArguments {
		isNat, err := k.typeCheckArgumentTL1(tr.Args[i], leftArgs)
		if err != nil {
			return err
		}
		if targ.IsNat != isNat {
			return fmt.Errorf("typeref %s argument %s wrong IsNat, must be %v", tr.Type, targ.FieldName, targ.IsNat)
		}
	}
	return nil
}

func (k *Kernel) typeCheckArgumentTL1(arg tlast.ArithmeticOrType, leftArgs []tlast.TemplateArgument) (isNat bool, _ error) {
	if arg.IsArith {
		return true, nil
	}
	if arg.T.Type.Namespace == "" {
		for _, la := range leftArgs {
			if arg.T.Type.Name == la.FieldName {
				if len(arg.T.Args) != 0 {
					return false, fmt.Errorf("reference to template argument %s cannot have arguments", la.FieldName)
				}
				return la.IsNat, nil
			}
		}
		// reference to global type
	}
	if err := k.typeCheckTypeRefTL1(arg.T, leftArgs); err != nil {
		return false, err
	}
	return false, nil
}
