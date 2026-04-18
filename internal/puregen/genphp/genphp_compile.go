// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package genphp

import (
	"fmt"
	"slices"

	"github.com/VKCOM/tl/internal/pure"
	"github.com/VKCOM/tl/internal/utils"
)

// create parallel hierarchy of wrappers, suitable for go generator
func (gen *genphp) compile() error {
	if err := gen.addTypeWrappers(); err != nil {
		return err
	}
	for _, myWrapper := range gen.generatedTypesList {
		switch pureType := myWrapper.pureType.(type) {
		case *pure.TypeInstancePrimitive:
			if pureType.CanonicalName() == "bool" || pureType.CanonicalName() == "bit" {
				if err := gen.generateTypeBool(myWrapper, pureType, pureType.CanonicalName() == "bit"); err != nil {
					return err
				}
			} else {
				if err := gen.generateTypePrimitive(myWrapper, pureType); err != nil {
					return err
				}
			}
		case *pure.TypeInstanceStruct:
			if err := gen.generateTypeStruct(myWrapper, pureType, nil, 0); err != nil {
				return err
			}
		case *pure.TypeInstanceArray:
			if err := gen.GenerateTypeArray(myWrapper, pureType); err != nil {
				return err
			}
		case *pure.TypeInstanceDict:
			// requires fully filled element, done on the next iteration
		case *pure.TypeInstanceUnion:
			if err := gen.generateTypeUnion(myWrapper, pureType); err != nil {
				return err
			}
		default:
			return fmt.Errorf("kernel type for %s not implemented in go generator", pureType.CanonicalName())
		}
	}
	for _, myWrapper := range gen.generatedTypesList {
		switch pureType := myWrapper.pureType.(type) {
		case *pure.TypeInstanceDict:
			if err := gen.GenerateTypeDict(myWrapper, pureType); err != nil {
				return err
			}
		}
	}
	if err := gen.prepareGeneration(); err != nil {
		return err
	}
	return nil
}

func (gen *genphp) addTypeWrappers() error {
	for _, pureType := range gen.kernel.AllTypeInstances() {
		myWrapper := &TypeRWWrapper{
			gen:      gen,
			pureType: pureType,
		}
		//if kt != nil {
		//	myWrapper.goCanonicalName = kt.CanonicalName()
		//}
		gen.generatedTypes[pureType.CanonicalName()] = myWrapper
		gen.generatedTypesList = append(gen.generatedTypesList, myWrapper)

		// TODO - we'd like to change this to fileName = goCanonicalName
		fileName := myWrapper.TLName()
		for len(fileName.Name) != 0 && fileName.Name[0] == '_' {
			fileName.Name = fileName.Name[1:]
		}
		fileName.Name = utils.ToLowerFirst(fileName.Name)
		myWrapper.fileName = fileName.String()

		namespace := gen.getNamespace(pureType.Common().ArgNamespace())
		namespace.types = append(namespace.types, myWrapper)
		myWrapper.ns = namespace
	}
	return nil
}

func (gen *genphp) prepareGeneration() error {
	slices.SortStableFunc(gen.generatedTypesList, func(a, b *TypeRWWrapper) int { //  TODO - better idea?
		return stringCompare(a.pureType.CanonicalName(), b.pureType.CanonicalName())
	})

	return nil
}
