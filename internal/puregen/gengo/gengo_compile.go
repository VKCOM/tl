// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package gengo

import (
	"fmt"
	"log"
	"slices"
	"strings"

	"github.com/vkcom/tl/internal/pure"
)

// create parallel hierarchy of wrappers, suitable for go generator
func (gen *genGo) compile() error {
	if err := gen.addTypeWrappers(); err != nil {
		return err
	}
	for _, myWrapper := range gen.generatedTypesList {
		for _, arg := range myWrapper.pureType.Common().ResolvedType().Args {
			if arg.IsArith {
				myWrapper.arguments = append(myWrapper.arguments, ResolvedArgument{
					isNat:   true,
					isArith: true,
					Arith:   arg.Arith,
				})
				continue
			}
			if arg.T.String() == "*" {
				myWrapper.arguments = append(myWrapper.arguments, ResolvedArgument{
					isNat:   true,
					isArith: false,
				})
				continue
			}
			ref, err := gen.kernel.GetInstanceTL1(arg.T)
			if err != nil {
				return fmt.Errorf("internal error: cannot get type of argument %s", arg.T)
			}
			fieldType, err := gen.getType(ref)
			if err != nil {
				return err
			}
			myWrapper.arguments = append(myWrapper.arguments, ResolvedArgument{
				isNat: false,
				tip:   fieldType,
				bare:  arg.T.Bare,
			})
		}
	}
	for _, myWrapper := range gen.generatedTypesList {
		switch pureType := myWrapper.pureType.(type) {
		case *pure.TypeInstancePrimitive:
			if err := gen.generateTypePrimitive(myWrapper, pureType); err != nil {
				return err
			}
		case *pure.TypeInstanceString:
		case *pure.TypeInstanceStruct:
			head, tail := myWrapper.resolvedT2GoName("")
			myWrapper.goGlobalName = gen.globalDec.deconflictName(head + tail)
			head, tail = myWrapper.resolvedT2GoName(myWrapper.tlName.Namespace)
			myWrapper.goLocalName = myWrapper.ns.decGo.deconflictName(head + tail)
			if err := gen.generateTypeStruct(myWrapper, pureType); err != nil {
				return err
			}
		case *pure.TypeInstanceArray:
			// return fmt.Errorf("Array type for %s not implemented in go generator", in.CanonicalName())
		default:
			return fmt.Errorf("kernel type for %s not implemented in go generator", pureType.CanonicalName())
		}
	}
	if err := gen.prepareGeneration(); err != nil {
		return err
	}
	return nil
}

func (gen *genGo) addTypeWrappers() error {
	for _, pureType := range gen.kernel.AllTypeInstances() {
		myWrapper := &TypeRWWrapper{
			gen:         gen,
			pureType:    pureType,
			NatParams:   pureType.Common().NatParams,
			unionParent: nil, // TODO
			unionIndex:  0,   // TODO
		}
		gen.generatedTypes[pureType.CanonicalName()] = myWrapper
		gen.generatedTypesList = append(gen.generatedTypesList, myWrapper)

		if kt := pureType.KernelType(); kt != nil {
			myWrapper.originateFromTL2 = kt.OriginTL2()
			myWrapper.origTL = kt.TL1()
			if !myWrapper.originateFromTL2 {
				if len(myWrapper.origTL) == 1 {
					myWrapper.tlTag = myWrapper.origTL[0].Crc32()
					myWrapper.tlName = myWrapper.origTL[0].Construct.Name
					myWrapper.fileName = myWrapper.tlName.String()
				} else {
					myWrapper.tlName = myWrapper.origTL[0].TypeDecl.Name
				}
			}
			namespace := gen.getNamespace(myWrapper.tlName.Namespace)
			namespace.types = append(namespace.types, myWrapper)
			myWrapper.ns = namespace
		}
	}
	return nil
}

func (gen *genGo) prepareGeneration() error {
	options := gen.options

	bytesWhiteList := pure.NewWhiteList(options.BytesWhiteList)
	tl2WhiteList := pure.NewWhiteList(options.Kernel.TL2WhiteList)

	bytesChildren := map[*TypeRWWrapper]bool{}
	typesCounterMarkBytes := 0
	for _, v := range gen.generatedTypesList {
		if bytesWhiteList.InNameFilter(v.tlName) {
			v.MarkWantsBytesVersion(bytesChildren)
			typesCounterMarkBytes++
		}
	}
	typesCounterMarkTL2 := 0
	tl2Children := map[*TypeRWWrapper]bool{}
	for _, v := range gen.generatedTypesList {
		if tl2WhiteList.InNameFilter(v.tlName) {
			v.MarkWantsTL2(tl2Children)
			typesCounterMarkTL2++
		}
	}
	for _, v := range gen.generatedTypesList { // we do not need tl2masks in this case
		if str, ok := v.trw.(*TypeRWStruct); ok && !v.wantsTL2 {
			for i := range str.Fields {
				str.Fields[i].MaskTL2Bit = nil
			}
		}
	}
	slices.SortStableFunc(gen.generatedTypesList, func(a, b *TypeRWWrapper) int { //  TODO - better idea?
		return TypeRWWrapperLessGlobal(a, b)
	})
	sortedTypes := gen.generatedTypesList
	// for _, st := range sortedTypes {
	//	fmt.Printf("sorted type %q\n", st.localTypeArg.rt.String())
	// }
	for _, v := range sortedTypes {
		// fmt.Printf("type %s names %s %s %d\n", v.CanonicalStringTop(), v.goGlobalName, v.tlName.String(), v.tlTag)
		// if len(v.origTL) <= 1 {
		//	fmt.Printf("     %s\n", v.CanonicalString(true))
		// } else {
		//	fmt.Printf("     %s\n", v.CanonicalString(false))
		// }
		// r # [r] = S;
		visitedNodes := map[*TypeRWWrapper]bool{}
		v.trw.fillRecursiveUnwrap(visitedNodes)
		v.preventUnwrap = visitedNodes[v]
		if v.preventUnwrap {
			fmt.Printf("prevented unwrap of %v\n", v.tlName)
		}
	}
	// in BeforeCodeGenerationStep we split recursion. Which links will be broken depends on order of nodes visited
	for _, v := range sortedTypes {
		v.trw.BeforeCodeGenerationStep1()
	}
	// in BeforeCodeGenerationStep2 we split recursion in unions.
	for _, v := range sortedTypes {
		v.trw.BeforeCodeGenerationStep2()
	}
	// we link normal and long types for VK int->long conversion. This code is VK-specific and will be removed after full migration
	for _, v := range sortedTypes {
		// @readwrite queueLong.getQueueKey id:long ip:int timeout:int queue:string = queueLong.TimestampKey;
		// @readwrite queue.getQueueKey id:int ip:int timeout:int queue:string = queue.TimestampKey;
		longName := v.CanonicalStringTop()
		argsStart := strings.Index(longName, "<")
		if argsStart < 0 {
			argsStart = len(longName)
		}
		if i := strings.Index(longName[:argsStart], "."); i >= 0 {
			longName = longName[:i] + "Long" + longName[i:]

			if tt, ok := gen.generatedTypes[longName]; ok {
				// log.Printf("long name %s discovered for %s", longName, v.CanonicalStringTop())
				v.WrLong = tt
				tt.WrWithoutLong = v
			}
		}

		v.trw.BeforeCodeGenerationStep2()
	}
	// Order of these 2 loops is important, for example see TypeRWTuple where bytes version depends on whether it is dict_like
	for _, v := range sortedTypes {
		visitedNodes := map[*TypeRWWrapper]bool{}
		v.hasBytesVersion = v.MarkHasBytesVersion(visitedNodes)
		visitedNodes = map[*TypeRWWrapper]bool{}
		v.hasErrorInWriteMethods = v.MarkWriteHasError(visitedNodes)
	}

	// detect recursion loops first
	if options.Verbose {
		//if skippedDueToWhitelist != 0 {
		//	log.Printf("skipped %d object roots by the whitelist filter: %s", skippedDueToWhitelist, strings.Join(typesWhiteList, ", "))
		//}
		if !bytesWhiteList.Empty() {
			log.Printf("found %d object roots for byte-optimized versions of types by the following filter: %s", typesCounterMarkBytes, options.BytesWhiteList)
		}
		if !tl2WhiteList.Empty() {
			log.Printf("found %d object roots for TL2 versions of types by the following filter: %s", typesCounterMarkTL2, options.Kernel.TL2WhiteList)
		}
	}
	return nil
}
