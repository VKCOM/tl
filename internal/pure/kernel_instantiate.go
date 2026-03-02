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
				ref.ins, err = k.createDict(canonicalName, tr)
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
			ref.ins, err = k.createOrdinaryTypeTL2(canonicalName, kt, tr, kt.canonicalName, kt.combTL2.TypeDecl.Magic,
				kt.combTL2.TypeDecl, kt.combTL2.TypeDecl.TemplateArguments)
			if err != nil {
				return nil, false, err
			}
			return ref, bare, nil
		}
		funcDecl := kt.combTL2.FuncDecl
		ref.ins, err = k.createStructTL2(canonicalName, kt, tr, kt.canonicalName, funcDecl.Magic, true,
			tlast.TL2TypeRef{}, funcDecl.Arguments, nil, false, 0,
			true, funcDecl)
		if err != nil {
			return nil, false, err
		}
		return ref, bare, nil
	}
	switch {
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

// alias || fields || union
func (k *Kernel) createOrdinaryTypeTL2(canonicalName string, tip *KernelType, tr tlast.TL2TypeRef,
	tlName tlast.TL2TypeName, tlTag uint32,
	def tlast.TL2TypeDeclaration,
	leftArgs []tlast.TL2TypeTemplate) (TypeInstance, error) {

	switch {
	case def.Type.IsAlias():
		return k.createAliasTL2(canonicalName, tip, tr, def)
	case def.Type.StructType.IsUnionType:
		return k.createUnionTL2(canonicalName, tip, tr, def.Type.StructType.UnionType, leftArgs)
	default:
		return k.createStructTL2(canonicalName, tip, tr,
			tlName, tlTag,
			true, def.Type.TypeAlias, def.Type.StructType.ConstructorFields,
			leftArgs, false, 0, false, tlast.TL2FuncDeclaration{})
	}
}
