// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package gengo

import (
	"strings"

	"github.com/vkcom/tl/internal/tlast"
)

// TODO - move into pure kernel as a helper/hint for language generators

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
