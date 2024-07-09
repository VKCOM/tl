// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlcodegen

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"go/format"
	"log"
	"path/filepath"
	"strings"

	"github.com/TwiN/go-color"
	"golang.org/x/exp/slices"
)

type InsFile struct {
	ins      *InternalNamespace
	fileName string
}

func (gen *Gen2) generateCodeGolang(generateByteVersions []string) error {
	sortedTypes := gen.generatedTypesList
	if !gen.options.SplitInternal {
		globalIns := &InternalNamespace{DebugID: 1, Name: "internal", SubPath: "internal", Namespaces: map[string]struct{}{}, DirectImports: &DirectImports{ns: map[*InternalNamespace]struct{}{}}}
		for _, v := range sortedTypes { // start with each type in its own internal namespace
			if _, ok := v.trw.(*TypeRWPrimitive); ok {
				continue // leave namespace nil
			}
			globalIns.Types = append(globalIns.Types, v)
			globalIns.Namespaces[v.tlName.Namespace] = struct{}{}
			v.ins = globalIns
		}
	} else {
		var internalNamespaces []*InternalNamespace
		printSortedNamespaces := func() {
			// fmt.Printf("----sortedTypes----\n")
			// for _, ins := range internalNamespaces {
			//	fmt.Printf("recursive import namespace %s", ins)
			// }
		}
		nextDebugID := 0
		for _, v := range sortedTypes { // start with each type in its own internal namespace
			if _, ok := v.trw.(*TypeRWPrimitive); ok {
				continue // leave namespace nil
			}
			nextDebugID++
			e := &InternalNamespace{DebugID: nextDebugID, Types: []*TypeRWWrapper{v}, Namespaces: map[string]struct{}{}, DirectImports: &DirectImports{ns: map[*InternalNamespace]struct{}{}}}
			e.Namespaces[v.tlName.Namespace] = struct{}{}
			v.ins = e
			internalNamespaces = append(internalNamespaces, e)
		}
		// we start with each type in its own internal namespace
		// we then generate code to assemble information on which is imported by which
		for _, v := range sortedTypes {
			if v.ins == nil {
				continue
			}
			_ = v.trw.GenerateCode(false, v.ins.DirectImports)
			if v.wantsBytesVersion && v.hasBytesVersion {
				_ = v.trw.GenerateCode(true, v.ins.DirectImports)
			}
		}
		printSortedNamespaces()
		// now := time.Now()
		recursiveImports := map[*InternalNamespace][]*InternalNamespace{}
		for i := 0; i < len(internalNamespaces); {
			from := internalNamespaces[i]
			from.FindRecursiveImports(recursiveImports, nil)
			if importMe, ok := recursiveImports[from]; ok && len(importMe) != 0 {
				into := importMe[0]
				// if options.Verbose {
				// log.Printf("----merging cycle namespace %sinto  %s (total %d)", from, into, len(internalNamespaces))
				// }
				into.mergeFrom(from, internalNamespaces)
				copy(internalNamespaces[i:], internalNamespaces[i+1:])
				internalNamespaces = internalNamespaces[:len(internalNamespaces)-1]
				// if options.Verbose {
				// log.Printf("result %s", into)
				// }
				continue
			}
			i++
		}
		for _, ins := range internalNamespaces {
			if len(ins.Types) == 1 {
				t := ins.Types[0]
				ins.Name = "tl" + t.goGlobalName
				ins.SubPath = "internal/tl" + t.tlName.Namespace + "/" + ins.Name
				continue
			}
			sha := sha1.Sum([]byte(strings.Join(ins.sortedElements(), ":")))
			ins.Name = "cycle_" + hex.EncodeToString(sha[:16])
			ins.SubPath = "internal/" + ins.Name
		}
	}
	typesCounter := 0
	typesCounterBytes := 0

	internalFiles := map[InsFile][]*TypeRWWrapper{}
	for _, typeRw := range gen.generatedTypesList {
		if typeRw.ins == nil {
			continue
		}
		ff := InsFile{ins: typeRw.ins, fileName: typeRw.fileName}
		internalFiles[ff] = append(internalFiles[ff], typeRw)
	}
	var s strings.Builder
	for ff, types := range internalFiles {
		directImports := &DirectImports{ns: map[*InternalNamespace]struct{}{}}
		for _, typeRw := range types {
			_ = typeRw.trw.GenerateCode(false, directImports)
			if typeRw.wantsBytesVersion && typeRw.hasBytesVersion {
				_ = typeRw.trw.GenerateCode(true, directImports)
			}
		}
		s.WriteString(fmt.Sprintf(`%s
package %s 
import (
`,
			HeaderComment, ff.ins.Name))
		if directImports.importSort {
			s.WriteString("\"sort\"\n\n")
		}
		s.WriteString(fmt.Sprintf("\"%s\"\n", gen.BasicPackageNameFull))
		if gen.options.SplitInternal {
			s.WriteString(fmt.Sprintf("    \"%s/%s\"\n", gen.options.TLPackageNameFull, "internal"))
		}
		var sortedNames []string
		for im := range directImports.ns { // Imports of this file.
			if im != ff.ins {
				sortedNames = append(sortedNames, im.SubPath)
			}
		}
		slices.Sort(sortedNames)
		for _, n := range sortedNames {
			s.WriteString(fmt.Sprintf("    \"%s/%s\"\n", gen.options.TLPackageNameFull, n))
		}
		s.WriteString(`)

var _ = basictl.NatWrite
`)
		if gen.options.SplitInternal {
			s.WriteString("var _ = internal.ErrorInvalidEnumTag\n")
		}
		for _, typeRw := range types {
			typesCounter++
			s.WriteString("\n")
			s.WriteString(typeRw.trw.GenerateCode(false, directImports))
			if typeRw.wantsBytesVersion && typeRw.hasBytesVersion {
				typesCounterBytes++
				s.WriteString("\n")
				s.WriteString(typeRw.trw.GenerateCode(true, directImports))
			}
		}
		filepathName := filepath.Join(ff.ins.SubPath, ff.fileName+goExt)
		if err := gen.addCodeFile(filepathName, gen.copyrightText+s.String()); err != nil {
			return err
		}
		s.Reset()
	}
	if gen.options.Verbose {
		log.Printf("generation of serialization code finished, %d constructors processed, %d types generated", len(gen.allConstructors), typesCounter)
		if len(generateByteVersions) != 0 {
			log.Printf("    also generated byte-optimized versions of %d types by the following filter: %s", typesCounterBytes, strings.Join(generateByteVersions, ", "))
		}
	}
	for name, namespace := range gen.Namespaces {
		slices.SortFunc(namespace.types, func(a, b *TypeRWWrapper) int {
			return TypeRWWrapperLessLocal(a, b)
		})
		anyTypeAlias := false
		anyEnumElementAlias := false
		anyFunction := false
		for _, wr := range namespace.types {
			if wr.ShouldWriteTypeAlias() {
				anyTypeAlias = true
			}
			if wr.ShouldWriteEnumElementAlias() {
				anyEnumElementAlias = true
			}
			if fun, ok := wr.trw.(*TypeRWStruct); ok && fun.ResultType != nil {
				anyFunction = true
			}
		}
		if !anyTypeAlias && !anyEnumElementAlias && !anyFunction {
			continue
		}
		directImports := &DirectImports{ns: map[*InternalNamespace]struct{}{}}
		var sortedNames []string
		_ = gen.generateNamespacesCode(anyTypeAlias, anyFunction, name, namespace, sortedNames, directImports)
		for im := range directImports.ns { // Imports of this file.
			sortedNames = append(sortedNames, im.SubPath)
		}
		slices.Sort(sortedNames)
		filepathName := filepath.Join(gen.GlobalPackageName+name, gen.GlobalPackageName+name+goExt)
		code := gen.generateNamespacesCode(anyTypeAlias, anyFunction, name, namespace, sortedNames, directImports)
		if code == "" {
			continue
		}
		if err := gen.addCodeFile(filepathName, gen.copyrightText+code); err != nil {
			return err
		}
	}
	if gen.options.Verbose {
		log.Printf("generation of RPC code finished, %d namespaces generated", len(gen.Namespaces))
	}
	{
		filepathName := filepath.Join(ConstantsPackageName, ConstantsPackageName+goExt) // TODO if contains GlobalPackgeName as prefix, there could be name collisions
		code := gen.generateConstants(HeaderComment, ConstantsPackageName)
		if err := gen.addCodeFile(filepathName, gen.copyrightText+code); err != nil {
			return err
		}
		if gen.options.BasicPackageNameFull == "" {
			filepathName = filepath.Join(BasicTLGoPackageName, BasicTLGoPackageName+goExt) // TODO if contains GlobalPackgeName as prefix, there could be name collisions
			code = fmt.Sprintf(basicTLCodeHeader, HeaderComment, BasicTLGoPackageName) + basicTLCodeBody
			if err := gen.addCodeFile(filepathName, gen.copyrightText+code); err != nil {
				return err
			}
		} else if gen.options.Verbose {
			log.Printf("basictl code not written, expected to be available at %q", gen.options.BasicPackageNameFull)
		}
		directImports := &DirectImports{ns: map[*InternalNamespace]struct{}{}}
		var sortedNames []string
		_ = gen.generateFactory(sortedNames, directImports)
		for im := range directImports.ns { // Imports of this file.
			sortedNames = append(sortedNames, im.SubPath)
		}
		slices.Sort(sortedNames)
		if err := gen.addCodeFile(filepath.Join(FactoryGoPackageName, FactoryGoPackageName+goExt), gen.copyrightText+gen.generateFactory(sortedNames, directImports)); err != nil {
			return err
		}
		if err := gen.addCodeFile(filepath.Join(FactoryGoPackageNameBytes, FactoryGoPackageNameBytes+goExt), gen.copyrightText+gen.generateFactoryBytes(sortedNames, directImports)); err != nil {
			return err
		}
		if err := gen.addCodeFile(filepath.Join(MetaGoPackageName, MetaGoPackageName+goExt), gen.copyrightText+gen.generateMeta()); err != nil {
			return err
		}
		filepathName = filepath.Join("internal", "a_tlgen_helpers_code.go") // TODO decollision
		code = fmt.Sprintf(internalTLCodeHeader, HeaderComment, "internal") + internalTLCodeBody
		if err := gen.addCodeFile(filepathName, gen.copyrightText+code); err != nil {
			return err
		}
	}
	if gen.options.Verbose {
		log.Printf("formating generated code...")
	}
	for filepathName, code := range gen.Code {
		if !strings.HasSuffix(filepathName, goExt) {
			continue
		}
		formattedCode, err := format.Source([]byte(code))
		if err != nil {
			// We generate code still, because it will be easy to debug when the wrong file is written out
			fmt.Printf("generator %sinternal error%s: source file %q will not compile due to error: %v", color.Red, color.Reset, filepathName, err)
			continue
		}
		gen.Code[filepathName] = string(formattedCode)
	}
	return nil
}
