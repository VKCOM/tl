// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package genrust

import (
	"fmt"
	"slices"
	"strings"

	"github.com/VKCOM/tl/internal/pure"
	"github.com/VKCOM/tl/internal/utils"
)

// create parallel hierarchy of wrappers, suitable for go generator
func (gen *genRust) compile() error {
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
			head, tail := myWrapper.resolvedT2GoName("")
			myWrapper.goGlobalName = gen.globalDec.DeconflictName(head + tail)
			head, tail = myWrapper.resolvedT2GoName(myWrapper.ns.name)
			myWrapper.goLocalName = myWrapper.ns.decGo.DeconflictName(head + tail)
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
			head, tail := myWrapper.resolvedT2GoName("")
			myWrapper.goGlobalName = gen.globalDec.DeconflictName(head + tail)
			head, tail = myWrapper.resolvedT2GoName(myWrapper.ns.name)
			myWrapper.goLocalName = myWrapper.ns.decGo.DeconflictName(head + tail)
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

func (gen *genRust) addTypeWrappers() error {
	for _, pureType := range gen.kernel.AllTypeInstances() {
		kt := pureType.KernelType()
		myWrapper := &TypeRWWrapper{
			gen:      gen,
			pureType: pureType,
		}
		if kt != nil {
			myWrapper.goCanonicalName = kt.CanonicalName()
		}
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

func (gen *genRust) prepareGeneration() error {
	options := gen.options

	bytesChildren := map[*TypeRWWrapper]bool{}
	typesCounterMarkBytes := 0
	for _, v := range gen.generatedTypesList {
		if gen.bytesWhiteList.HasName2(v.TLName()) {
			v.MarkWantsBytesVersion(bytesChildren)
			typesCounterMarkBytes++
		}
	}
	slices.SortStableFunc(gen.generatedTypesList, func(a, b *TypeRWWrapper) int { //  TODO - better idea?
		return TypeRWWrapperLessGlobal(a, b)
	})
	sortedTypes := gen.generatedTypesList
	// for _, st := range sortedTypes {
	//	fmt.Printf("sorted type %q\n", st.localTypeArg.rt.String())
	// }
	// in BeforeCodeGenerationStep we split recursion. Which links will be broken depends on order of nodes visited
	for _, v := range sortedTypes {
		v.trw.BeforeCodeGenerationStep1()
	}
	// in BeforeCodeGenerationStep2 we split recursion in unions.
	for _, v := range sortedTypes {
		v.trw.BeforeCodeGenerationStep2()
	}
	// TODO - long adapters
	// we link normal and long types for VK int->long conversion. This code is VK-specific and will be removed after full migration
	for _, v := range sortedTypes {
		gen.findLongAdapter(v)
		v.trw.BeforeCodeGenerationStep2()
	}
	// Order of these 2 loops is important, for example see TypeRWTuple where bytes version depends on whether it is dict_like
	for _, v := range sortedTypes {
		visitedNodes := map[*TypeRWWrapper]bool{}
		v.hasBytesVersion = v.MarkHasBytesVersion(visitedNodes)
		visitedNodes = map[*TypeRWWrapper]bool{}
		v.hasErrorInWriteMethods = v.MarkWriteHasError(visitedNodes)
		// TODO - move into pure kernel
		visitedNodes = map[*TypeRWWrapper]bool{}
		v.hasRepairMasks = v.MarkHasRepairMasks(visitedNodes)
	}

	if options.Kernel.Verbose {
		//if skippedDueToWhitelist != 0 {
		//	fmt.Printf("skipped %d object roots by the whitelist filter: %s\n", skippedDueToWhitelist, strings.Join(typesWhiteList, ", "))
		//}
		if !gen.bytesWhiteList.Empty() {
			fmt.Printf("found %d object roots for byte-optimized versions of types by the following filter: %s\n", typesCounterMarkBytes, options.BytesWhiteList)
		}
	}
	return nil
}

// this trash is to be removed with the last long adapter
func (gen *genRust) findLongAdapter(v *TypeRWWrapper) {
	//if strings.Contains(strings.ToLower(v.goCanonicalName.String()), "message") {
	//	fmt.Printf("%s %s\n", v.goCanonicalName.String(), v.pureType.CanonicalName())
	//}
	// @readwrite queueLong.getQueueKey id:long ip:int timeout:int queue:string = queueLong.TimestampKey;
	// @readwrite queue.getQueueKey id:int ip:int timeout:int queue:string = queue.TimestampKey;
	longName := v.pureType.CanonicalName()
	argsStart := strings.Index(longName, "<")
	if argsStart < 0 {
		argsStart = len(longName)
	}
	i := strings.Index(longName[:argsStart], ".")
	if i < 0 {
		return
	}
	longName = longName[:i] + "Long" + longName[i:]

	if tt, ok := gen.generatedTypes[longName]; ok {
		// fmt.Printf("long name %s discovered for %s\n", longName, v.pureType.CanonicalName())
		v.WrLong = tt
		tt.WrWithoutLong = v
		return
	}
}
