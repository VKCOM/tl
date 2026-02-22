// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package pure

import (
	"fmt"
	"strings"

	"github.com/vkcom/tl/internal/tlast"
	"github.com/vkcom/tl/internal/utils"
)

func (k *Kernel) IsTrueType(rt tlast.TypeRef) bool {
	return rt.Type.String() == "true" || rt.Type.String() == "True"
}

func (k *Kernel) IsDict(kt *KernelType) (bool, *KernelType) {
	if kt.originTL2 || len(kt.combTL1) != 1 ||
		!strings.Contains(strings.ToLower(kt.canonicalName.Name), "dictionary") {
		return false, nil
	}
	fieldT, fieldOk := k.tips[kt.canonicalName.String()+"Field"]
	if !fieldOk || fieldT.originTL2 || len(fieldT.combTL1) != 1 {
		return false, nil
	}
	if len(fieldT.combTL1[0].TemplateArguments) == 0 ||
		len(fieldT.combTL1[0].TemplateArguments) > 2 ||
		len(fieldT.combTL1[0].TemplateArguments) != len(kt.combTL1[0].TemplateArguments) ||
		len(fieldT.combTL1[0].Fields) != 2 {
		return false, nil
	}
	// This only checks some type properties, they are enough for us for now
	for i, targ := range kt.combTL1[0].TemplateArguments {
		farg := fieldT.combTL1[0].TemplateArguments[i]
		if targ.IsNat || farg.IsNat || targ.FieldName != farg.FieldName {
			return false, nil
		}
	}
	return true, fieldT
}

func (k *Kernel) VariantNames(definition []*tlast.Combinator) ([]string, error) {
	// this mess is because there was no clear variant names in TL1
	// Removing prefix/suffix common with union name.
	// We allow relaxed case match. To use strict match, we could remove all strings.ToLower() calls below
	typePrefix := strings.ToLower(utils.ToLowerFirst(definition[0].TypeDecl.Name.Name))
	typeSuffix := strings.ToLower(definition[0].TypeDecl.Name.Name)
	for _, typ := range definition {
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

	var variantNames []string
	for _, variantDef := range definition {
		variantName := variantDef.Construct.Name.Name
		if typePrefix != "" && len(typePrefix) < len(variantName) {
			variantName = variantName[len(typePrefix):]
		} else if typeSuffix != "" && len(typeSuffix) < len(variantName) {
			variantName = variantName[:len(variantName)-len(typeSuffix)]
		}
		for len(variantName) != 0 && variantName[0] == '_' {
			variantName = variantName[1:]
		}
		if !utils.IsFirstBasicLatin(variantName) { // digit
			variantName = "v" + variantName
		}
		// check against already defined fields
		for _, usedName := range variantNames {
			if usedName == variantName {
				// We have such cases in combined.tl, for example
				// messages.oneUser#a6a042bd user_id:messages.userId = messages.ChatUsers;
				// messagesLong.oneUser#5fb6003f user_id:messagesLong.userId = messages.ChatUsers;
				variantName = variantDef.Construct.Name.Namespace + "_" + variantName // add namespace on collision
				break
			}
		}
		// check again
		for _, usedName := range variantNames {
			if usedName == variantName {
				return nil, fmt.Errorf("cannot define TL1 union - prohibited variant name collision")
			}
		}
		variantNames = append(variantNames, variantName)
	}
	return variantNames, nil
}
