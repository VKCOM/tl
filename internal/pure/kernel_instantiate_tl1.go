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

func (k *Kernel) resolveMaskTL1(mask tlast.FieldMask, leftArgs []tlast.TemplateArgument,
	actualArgs []LocalArgHybrid, combinatorField tlast.CombinatorField) (ActualNatArg, error) {
	for i, targ := range leftArgs {
		if targ.FieldName == mask.MaskName {
			actualArg := actualArgs[i]
			if actualArg.wrongTypeErr != nil {
				e1 := mask.PRName.BeautifulError(fmt.Errorf("reference %q should be to #-param or # field", targ.FieldName))
				return ActualNatArg{}, tlast.BeautifulError2(e1, actualArg.wrongTypeErr)
			}
			if !targ.IsNat {
				e1 := mask.PRName.BeautifulError(fmt.Errorf("fieldMask cannot reference Type-parameter %s", targ.FieldName))
				e2 := targ.PR.BeautifulError(fmt.Errorf("declared here"))
				return ActualNatArg{}, tlast.BeautifulError2(e1, e2)
			}
			if len(actualArg.natArgs) > 1 {
				return ActualNatArg{}, mask.PRName.BeautifulError(fmt.Errorf("internal error: fieldMask reference len(natArg) == %d for parameter %s", len(actualArg.natArgs), targ.FieldName))
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

func (k *Kernel) GetInstanceTL1(tr tlast.TL2TypeRef) (TypeInstance, bool, error) {
	ref, bare, err := k.getInstanceTL1(tr, false)
	if err != nil {
		return nil, false, err
	}
	return ref.ins, bare, nil
}

func (k *Kernel) getInstanceTL1(tr tlast.TL2TypeRef, create bool) (_ *TypeInstanceRef, bare bool, _ error) {
	canonicalName, bare, err := k.canonicalStringTL2(tr, true)
	if err != nil {
		return nil, false, err
	}
	if ref, ok := k.instances[canonicalName]; ok {
		return ref, bare, nil
	}
	if !create {
		return nil, false, fmt.Errorf("internal error: instance %s must exist", canonicalName)
	}
	if br := tr.BracketType; br != nil {
		ref := k.addInstance(canonicalName, k.brackets)
		// must store pointer before children GetInstanceTL2() terminates recursion
		// this instance stays not initialized in case of error, but kernel then is not consistent anyway
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
	// log.Printf("creating an instance of type %s", canonicalName)
	// we must mark all usedAsNatConst, usedAsNatVariable, so
	// will do some work before looking up and returning existing instance
	tName := tr.SomeType.Name.String()
	kt, ok := k.tips[tName]
	if !ok {
		return nil, false, fmt.Errorf("type %s does not exist", tName)
	}
	// TODO - instantiate TL2 combinators here
	if kt.originTL2 {
		return nil, false, fmt.Errorf("TL1 combinator cannot reference TL2 combinator %s", tName)
	}
	td := kt.combTL1[0]
	// must store pointer before children GetInstanceTL2() terminates recursion
	// this instance stays not initialized in case of error, but kernel then is not consistent anyway
	ref := k.addInstance(canonicalName, kt)
	switch {
	case tName == "__dict":
		// log.Printf("creating an instance of dictionary type %s", canonicalName)
		ref.ins, err = k.createDictTL1(canonicalName, kt, tr, td.TemplateArguments)
	case tName == "__dict2":
		// log.Printf("creating an instance of dictionary type %s", canonicalName)
		//ref.ins, err = k.createDictTL1(canonicalName, kt, tr, td.TemplateArguments, tr.Args)
		ref.ins, err = k.createDict(canonicalName, kt, tr, td.TemplateArguments)
	case len(kt.combTL1) > 1:
		ref.ins, err = k.createUnionTL1FromTL1(canonicalName, kt, tr, kt.combTL1)
	case len(kt.combTL1) == 1:
		ref.ins, err = k.createStructTL1FromTL1(canonicalName, kt, tr,
			kt.combTL1[0],
			false, 0)
	default:
		return nil, false, fmt.Errorf("wrong type classification, internal error %s", canonicalName)
	}
	if err != nil {
		return nil, false, err
	}
	return ref, bare, nil
}
