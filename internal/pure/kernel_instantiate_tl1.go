// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package pure

import (
	"errors"
	"fmt"

	"github.com/vkcom/tl/internal/tlast"
)

func (k *Kernel) resolveMaskTL1(mask tlast.FieldMask, leftArgs []tlast.TL2TypeTemplate,
	actualArgs []LocalArgHybrid, combinatorField tlast.CombinatorField) (ActualNatArg, error) {
	for i, targ := range leftArgs {
		if targ.Name == mask.MaskName {
			actualArg := actualArgs[i]
			if actualArg.wrongTypeErr != nil {
				e1 := mask.PRName.BeautifulError(fmt.Errorf("reference %q should be to #-param or # field", targ.Name))
				return ActualNatArg{}, tlast.BeautifulError2(e1, actualArg.wrongTypeErr)
			}
			if !targ.Category.IsNatValue {
				e1 := mask.PRName.BeautifulError(fmt.Errorf("fieldMask cannot reference Type-parameter %s", targ.Name))
				e2 := targ.PR.BeautifulError(fmt.Errorf("declared here"))
				return ActualNatArg{}, tlast.BeautifulError2(e1, e2)
			}
			if len(actualArg.natArgs) > 1 {
				return ActualNatArg{}, mask.PRName.BeautifulError(fmt.Errorf("internal error: fieldMask reference len(natArg) == %d for parameter %s", len(actualArg.natArgs), targ.Name))
			}
			if actualArg.arg.IsNumber {
				return ActualNatArg{
					isNumber: true,
					number:   actualArg.arg.Number,
				}, nil
			}
			if sf := actualArg.arg.SourceField; sf != (tlast.CombinatorField{}) {
				field := &sf.Comb.Fields[sf.FieldIndex]
				if field.UsedAsSize {
					e3 := field.UsedAsSizePR.BeautifulError(fmt.Errorf("used as size here"))
					e3.PrintWarning(k.opts.ErrorWriter, nil)
					e1 := field.PRName.BeautifulError(fmt.Errorf("#-field %s is used as an field mask, while already being used as tuple size", field.FieldName))
					e2 := mask.PRName.BeautifulError(fmt.Errorf("used as mask here"))
					return ActualNatArg{}, tlast.BeautifulError2(e1, e2)
				}
				field.UsedAsMask = true
				field.UsedAsMaskPR = mask.PRName
				field.AffectedFields = append(field.AffectedFields, combinatorField)
			}
			return actualArg.natArgs[0], nil
		}
	}
	return ActualNatArg{}, mask.PRName.BeautifulError(errors.New("fieldMask reference not found"))
}

// we do not allow creating additional instances externally for now
// we identify types by TL2TypeRefs/TL2TypeNames, TL1 types are first converted into TL2 style
func (k *Kernel) GetInstance(tr tlast.TL2TypeRef) (TypeInstance, bool, error) {
	ref, bare, err := k.getInstance(tr, false)
	if err != nil {
		return nil, false, err
	}
	return ref.ins, bare, nil
}

// we identify types by TL2TypeRefs/TL2TypeNames, TL1 types are first converted into TL2 style
func (k *Kernel) getInstance(tr tlast.TL2TypeRef, create bool) (_ *TypeInstanceRef, bare bool, _ error) {
	canonicalName, bare, err := k.canonicalString(tr, true)
	if err != nil {
		return nil, false, err
	}
	if ref, ok := k.instances[canonicalName]; ok {
		return ref, bare, nil
	}
	if !create {
		return nil, false, fmt.Errorf("internal error: instance %s must exist", canonicalName)
	}
	// fmt.Printf("creating an instance of type %s\n", canonicalName)
	if br := tr.BracketType; br != nil {
		// must store pointer before children GetInstanceTL2() terminates recursion
		// this instance stays not initialized in case of error, but kernel then is not consistent anyway
		ref := k.addInstance(canonicalName, k.brackets)
		if br.HasIndex {
			if br.IndexType.IsNumber || br.IndexType.Type.String() == "*" {
				ref.ins, err = k.createTupleTL1(canonicalName, tr)
			} else {
				panic("TODO - dict2")
			}
		} else {
			ref.ins, err = k.createVectorTL1(canonicalName, tr)
		}
		if err != nil {
			return nil, false, err
		}
		return ref, bare, nil
	}
	tName := tr.SomeType.Name.String()
	kt, ok := k.tips[tName]
	if !ok {
		return nil, false, fmt.Errorf("type %s does not exist", tName)
	}
	// must store pointer before children GetInstanceTL2() terminates recursion
	// this instance stays not initialized in case of error, but kernel then is not consistent anyway
	ref := k.addInstance(canonicalName, kt)
	if kt.originTL2 {
		if !kt.combTL2.IsFunction {
			ref.ins, err = k.createOrdinaryType(canonicalName, kt, tr, kt.canonicalName, kt.combTL2.TypeDecl.Magic,
				kt.combTL2.TypeDecl.Type, kt.combTL2.TypeDecl.TemplateArguments, tr.SomeType.Arguments)
			if err != nil {
				return nil, false, err
			}
			return ref, bare, nil
		}
		funcDecl := kt.combTL2.FuncDecl
		resultAlias := false
		var resultType TypeInstance
		if resultTlName, ok := k.functionNeedsGeneratedResultType(funcDecl); ok {
			resultAlias = true
			resultIns, _, err := k.getInstance(tlast.TL2TypeRef{SomeType: tlast.TL2TypeApplication{Name: resultTlName}}, true)
			if err != nil {
				return nil, false, err
			}
			resultType = resultIns.ins // function cannot be referenced so no recursion
		} else if funcDecl.ReturnType.IsTypeAlias {
			resultAlias = true
			resultIns, _, err := k.getInstance(funcDecl.ReturnType.TypeAlias, true)
			if err != nil {
				return nil, false, err
			}
			resultType = resultIns.ins // function cannot be referenced so no recursion
		} else {
			fieldReturnType := funcDecl.ReturnType.StructType.ConstructorFields[0].Type
			// this is special case because we want no diff during migration to TL2,
			// and all TL1 functions return exactly this kind of result
			resultIns, _, err := k.getInstance(fieldReturnType, true)
			if err != nil {
				return nil, false, err
			}
			resultType = resultIns.ins // function cannotbe referenced so no recursion
		}
		ref.ins, err = k.createStruct(canonicalName, kt, tr, kt.canonicalName, funcDecl.Magic, true,
			tlast.TL2TypeRef{}, funcDecl.Arguments, nil, false, 0,
			resultType, resultAlias)
		if err != nil {
			return nil, false, err
		}
		return ref, bare, nil
	}
	switch {
	case tName == "__dict":
		fmt.Printf("creating an instance of dictionary type %s\n", canonicalName)
		ref.ins, err = k.createDictTL1(canonicalName, kt, tr)
	case tName == "__dict2":
		// fmt.Printf("creating an instance of dictionary type %s\n", canonicalName)
		ref.ins, err = k.createDict(canonicalName, kt, tr)
	case len(kt.combTL1) > 1:
		ref.ins, err = k.createUnionTL1FromTL1(canonicalName, kt, tr, kt.combTL1)
	case len(kt.combTL1) == 1:
		ref.ins, err = k.createStructTL1FromTL1(canonicalName, kt, tr, kt.combTL1[0],
			false, 0)
	default:
		return nil, false, fmt.Errorf("wrong type classification, internal error %s", canonicalName)
	}
	if err != nil {
		return nil, false, err
	}
	return ref, bare, nil
}
