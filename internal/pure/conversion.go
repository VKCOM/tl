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

func (k *Kernel) equalTypes(rt tlast.TypeRef, rt2 tlast.TL2TypeRef) {
	rt22 := k.convertTypeRef(rt)
	a, b := rt22.String(), rt2.String()
	if a != b {
		panic(fmt.Errorf("!equalsType %s %s", a, b))
	}
}

func (k *Kernel) equalNatArgs(rt []ActualNatArg, rt2 []ActualNatArg) {
	if len(rt) != len(rt2) {
		panic(fmt.Errorf("!equalNatArgs %d %d", len(rt), len(rt2)))
	}
	for i, na := range rt {
		nb := rt2[i]
		a := fmt.Sprintf("%v", na)
		b := fmt.Sprintf("%v", nb)
		if a != b {
			panic(fmt.Errorf("!equalNatArgs %s %s", a, b))
		}
	}
}

func (k *Kernel) convertTemplateArguments(args []tlast.TemplateArgument) []tlast.TL2TypeTemplate {
	var result []tlast.TL2TypeTemplate
	for _, arg := range args {
		result = append(result, tlast.TL2TypeTemplate{
			Name:       arg.FieldName,
			Category:   tlast.TL2TypeCategory{IsNatValue: arg.IsNat},
			PR:         arg.PR,
			PRName:     arg.PR,
			PRCategory: arg.PR,
		})
	}
	return result
}

func (k *Kernel) convertTypeArgument(tra tlast.ArithmeticOrType) tlast.TL2TypeArgument {
	if tra.IsArith {
		return tlast.TL2TypeArgument{
			Number:   tra.Arith.Res,
			IsNumber: true,
			PR:       tra.T.PR,
		}
	}
	return tlast.TL2TypeArgument{
		Type:   k.convertTypeRef(tra.T),
		Number: tra.Arith.Res,
		PR:     tra.T.PR,
	}
}

func (k *Kernel) convertTypeRef(tr tlast.TypeRef) tlast.TL2TypeRef {
	if tr.Type.String() == "__vector" && len(tr.Args) == 1 && !tr.Args[0].IsArith {
		return tlast.TL2TypeRef{
			BracketType: &tlast.TL2BracketType{
				ArrayType: k.convertTypeRef(tr.Args[0].T),
				PR:        tr.PR,
			},
		}
	}
	if tr.Type.String() == "__tuple" && len(tr.Args) == 2 && !tr.Args[1].IsArith {
		return tlast.TL2TypeRef{
			BracketType: &tlast.TL2BracketType{
				HasIndex:  true,
				IndexType: k.convertTypeArgument(tr.Args[0]),
				ArrayType: k.convertTypeRef(tr.Args[1].T),
				PR:        tr.PR,
			},
		}
	}
	//if tr.Type.String() == "__dict" && len(tr.Args) == 2 && !tr.Args[1].IsArith {
	//	return tlast.TL2TypeRef{
	//		BracketType: &tlast.TL2BracketType{
	//			IndexType: k.convertTypeArgument(tr.Args[0]),
	//			ArrayType: k.convertTypeRef(tr.Args[1].T),
	//			PR:        tr.PR,
	//		},
	//	}
	//}
	result := tlast.TL2TypeRef{
		SomeType: tlast.TL2TypeApplication{
			Name:        tlast.TL2TypeName(tr.Type),
			Arguments:   nil,
			PR:          tr.PR,
			PRName:      tr.PR,
			PRArguments: tr.PRArgs,
			Bare:        tr.Bare,
		},
		PR: tr.PR,
	}
	for _, arg := range tr.Args {
		result.SomeType.Arguments = append(result.SomeType.Arguments,
			k.convertTypeArgument(arg))
	}
	return result
}
