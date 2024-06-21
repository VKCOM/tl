// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlcodegen

import (
	"fmt"
	"path/filepath"
	"strings"
)

func cppStartNamespace(s *strings.Builder, ns []string) {
	for _, n := range ns {
		s.WriteString(fmt.Sprintf("namespace %s { ", n))
	}
	s.WriteString("\n")
}

func cppFinishNamespace(s *strings.Builder, ns []string) {
	s.WriteString(fmt.Sprintf("\n%s // namespace %s\n\n", strings.Repeat("}", len(ns)), strings.Join(ns, "::")))
}

func (gen *Gen2) generateCodeCPP(generateByteVersions []string) error {
	const basicTLFilepathName = "a_tlgen_helpers_code" + hppExt // TODO decollision

	cppAllInc := &DirectIncludesCPP{ns: map[*TypeRWWrapper]CppIncludeInfo{}}
	var hpp strings.Builder

	typesCounter := 0

	gen.decideCppCodeDestinations(gen.generatedTypesList)

	internalFiles2 := map[InsFile]map[string][]*TypeRWWrapper{}

	for _, typeRw := range gen.generatedTypesList {
		hppDef := InsFile{ins: typeRw.ins, fileName: typeRw.fileName}
		if _, ok := internalFiles2[hppDef]; !ok {
			internalFiles2[hppDef] = make(map[string][]*TypeRWWrapper)
		}
		internalFiles2[hppDef][typeRw.detailsFileName] = append(internalFiles2[hppDef][typeRw.detailsFileName], typeRw)
	}

	// for each type header ~ tl combinator
	for ff, details := range internalFiles2 {
		hppInc := &DirectIncludesCPP{ns: map[*TypeRWWrapper]CppIncludeInfo{}}
		hppIncFwd := &DirectIncludesCPP{ns: map[*TypeRWWrapper]CppIncludeInfo{}}
		typeDefinitions := map[string]bool{}

		// for each file with specifications
		for detailFile, specs := range details {
			hppDetInc := &DirectIncludesCPP{ns: map[*TypeRWWrapper]CppIncludeInfo{}}
			cppDetInc := &DirectIncludesCPP{ns: map[*TypeRWWrapper]CppIncludeInfo{}}
			var hppDet strings.Builder
			var cppDet strings.Builder

			// for each specification
			for _, typeRw := range specs {
				// init all variants for specification (ex. byte version)
				typeDefVariations := make([]TypeDefinitionVariation, 1)
				{
					if typeRw.wantsBytesVersion && typeRw.trw.CPPHasBytesVersion() {
						typeDefVariations = append(typeDefVariations, TypeDefinitionVariation{NeedBytesVersion: true})
					}
				}

				for _, typeDefVariation := range typeDefVariations {
					typesCounter++
					var hppDefinition strings.Builder
					typeRw.trw.CPPGenerateCode(&hppDefinition, hppInc, hppIncFwd, &hppDet, hppDetInc, &cppDet, cppDetInc, typeDefVariation.NeedBytesVersion, false)
					def := hppDefinition.String()
					if !typeDefinitions[def] {
						typeDefinitions[def] = true
						hpp.WriteString(def)
					}
				}
			}

			if hpp.Len() == 0 && hppDet.Len() == 0 && cppDet.Len() == 0 {
				continue
			}

			// all specs in one file must be in group
			cppAllInc.ns[specs[0]] = CppIncludeInfo{-1, specs[0].groupName}

			hppDetStr := hppDet.String()
			cppDetStr := cppDet.String()

			hppDet.Reset()
			cppDet.Reset()

			hppDet.WriteString("#pragma once\n\n")
			hppDet.WriteString(fmt.Sprintf("#include \"../%s\"\n", basicTLFilepathName))

			hppDet.WriteString(fmt.Sprintf("#include \"../%s%s\"\n", ff.fileName, hppExt))
			for _, n := range hppDetInc.sortedIncludes(gen.componentsOrder, func(wrapper *TypeRWWrapper) string { return wrapper.fileName }) {
				if n == ff.fileName {
					continue
				}
				hppDet.WriteString(fmt.Sprintf("#include \"../%s%s\"\n", n, hppExt))
			}
			hppDet.WriteString("\n")

			cppDet.WriteString(fmt.Sprintf("#include \"%s%s\"\n", detailFile, hppExt))
			for _, n := range cppDetInc.sortedIncludes(gen.componentsOrder, func(wrapper *TypeRWWrapper) string { return wrapper.detailsFileName }) {
				if n == detailFile {
					continue
				}
				cppDet.WriteString(fmt.Sprintf("#include \"%s%s\"\n", n, hppExt))
			}
			cppDet.WriteString("\n")

			filepathName := filepath.Join("details", detailFile+hppExt)
			if err := gen.addCodeFile(filepathName, gen.copyrightText+hppDet.String()+hppDetStr); err != nil {
				return err
			}
			filepathName = filepath.Join("details", detailFile+cppExt)
			if err := gen.addCodeFile(filepathName, gen.copyrightText+cppDet.String()+cppDetStr); err != nil {
				return err
			}
		}

		if hpp.Len() == 0 {
			continue
		}

		hppStr := hpp.String()
		hpp.Reset()
		hpp.WriteString("#pragma once\n\n")
		hpp.WriteString(fmt.Sprintf("#include \"%s\"\n", basicTLFilepathName))
		for _, n := range hppInc.sortedIncludes(gen.componentsOrder, func(wrapper *TypeRWWrapper) string { return wrapper.fileName }) {
			hpp.WriteString(fmt.Sprintf("#include \"%s%s\"\n", n, hppExt))
		}
		hpp.WriteString("\n\n")
		hpp.WriteString(hppStr)
		filepathName := ff.fileName + hppExt
		if err := gen.addCodeFile(filepathName, gen.copyrightText+hpp.String()); err != nil {
			return err
		}
		hpp.Reset()
	}

	var cppAll strings.Builder
	var cppMake strings.Builder
	var cppMakeO strings.Builder
	var cppMake1 strings.Builder

	for _, nf := range cppAllInc.splitByNamespaces() {
		// it is a group
		namespace := nf.Namespace
		if namespace == "" {
			namespace = "__common"
		}

		var cppMake1UsedFiles strings.Builder
		var cppMake1Namespace strings.Builder

		for _, n := range nf.Includes.sortedIncludes(gen.componentsOrder, func(wrapper *TypeRWWrapper) string { return wrapper.detailsFileName }) {
			cppAll.WriteString(fmt.Sprintf("#include \"details/%s%s\"\n", n, cppExt))
			cppMake1Namespace.WriteString(fmt.Sprintf("#include \"../%s%s\"\n", n, cppExt))
			cppMake1UsedFiles.WriteString(fmt.Sprintf("details/%s%s details/%s%s ", n, cppExt, n, hppExt))
		}

		namespaceDetails := namespace
		namespaceFilePath := "details/namespaces/" + namespaceDetails + cppExt
		buildFilePath := "build/" + namespaceDetails + ".o"

		cppMake1.WriteString(fmt.Sprintf("%s: %s %s\n", buildFilePath, namespaceFilePath, cppMake1UsedFiles.String()))
		cppMake1.WriteString(fmt.Sprintf("\t$(CC) $(CFLAGS) -o %s -c %s\n", buildFilePath, namespaceFilePath))
		cppMakeO.WriteString(fmt.Sprintf("%s ", buildFilePath))

		if err := gen.addCodeFile(namespaceFilePath, cppMake1Namespace.String()); err != nil {
			return err
		}
	}

	cppMake.WriteString(`
CC = g++
CFLAGS = -std=c++17 -O3 -Wno-noexcept-type -g -Wall -Wextra -Werror=return-type -Wno-unused-parameter
`)
	cppMake.WriteString(fmt.Sprintf("all: main.o %s\n", cppMakeO.String()))
	cppMake.WriteString(fmt.Sprintf("\t$(CC) $(CFLAGS) -o all main.o %s\n", cppMakeO.String()))
	cppMake.WriteString(`
main.o: main.cpp
	$(CC) $(CFLAGS) -c main.cpp
`)
	cppMake.WriteString(cppMake1.String())
	if err := gen.addCodeFile("all.cpp", cppAll.String()); err != nil {
		return err
	}
	if err := gen.addCodeFile("main.cpp", "int main() { return 0; }"); err != nil {
		return err
	}
	if err := gen.addCodeFile("Makefile", cppMake.String()); err != nil {
		return err
	}
	if err := gen.addCodeFile("build/info.txt", ".o files here!"); err != nil {
		return err
	}
	// if gen.options.Verbose {
	//	log.Printf("generation of serialization code finished, %d constructors processed, %d types generated", len(gen.allConstructors), typesCounter)
	//	if len(generateByteVersions) != 0 {
	//		log.Printf("    also generated byte-optimized versions of %d types by the following filter: %s", typesCounterBytes, strings.Join(generateByteVersions, ", "))
	//	}
	// }
	// if gen.options.GenerateRPCCode {
	//	for name, namespace := range gen.Namespaces {
	//		filepathName := filepath.Join(gen.GlobalPackageName+name, gen.GlobalPackageName+name+".go")
	//		// TODO - if no functions and no aliases, do not write namespace file at all
	//		code := gen.GenerateNamespacesCode(name, namespace)
	//		if _, ok := gen.Code[filepathName]; ok {
	//			return fmt.Errorf("generator %sinternal error%s: source file %q is generated twice", tlast.ColorRed, tlast.ColorReset, filepathName)
	//		}
	//		gen.Code[filepathName] = code
	//	}
	// }
	// if gen.options.Verbose {
	//	log.Printf("generation of RPC code finished, %d namespaces generated", len(gen.Namespaces))
	// }
	{
		//	filepathName := filepath.Join(BasicTLGoPackageName, BasicTLGoPackageName+".go") // TODO if contains GlobalPackgeName as prefix, there could be name collisions
		//	gen.Code[filepathName] = fmt.Sprintf(basicTLCodeHeader, HeaderComment, BasicTLGoPackageName) + basicTLCodeBody
		//	filepathName = "factory.go"
		//	gen.Code[filepathName] = gen.GenerateFactory()
		code := fmt.Sprintf(basicCPPTLCodeHeader, HeaderComment, BasicTLCPPNamespaceName) + basicCPPTLCodeBody +
			fmt.Sprintf(basicCPPTLCodeFooter, BasicTLCPPNamespaceName)
		if err := gen.addCodeFile(basicTLFilepathName, code); err != nil {
			return err
		}
	}
	// if gen.options.Verbose {
	//	log.Printf("formating generated code...")
	// }
	// for filepathName, code := range gen.Code {
	//	formattedCode, err := format.Source([]byte(code))
	//	if err != nil {
	//		// We generate code still, because it will be easy to debug when the wrong file is written out
	//		fmt.Printf("generator %sinternal error%s: source file %q will not compile due to error: %v", tlast.ColorRed, tlast.ColorReset, filepathName, err)
	//		continue
	//	}
	//	gen.Code[filepathName] = string(formattedCode)
	// }
	return nil
}

func sliceToSet[T comparable](s []T) map[T]bool {
	m := make(map[T]bool)
	for _, e := range s {
		m[e] = true
	}
	return m
}

// unstable
func setToSlice[T comparable](s map[T]bool) []T {
	m := make([]T, 0)
	for k, _ := range s {
		m = append(m, k)
	}
	return m
}

func mapSlice[A, B any](in []A, f func(A) B) (out []B) {
	for _, e := range in {
		out = append(out, f(e))
	}
	return
}

func filterSlice[A any](in []A, f func(A) bool) (out []A) {
	for _, e := range in {
		if f(e) {
			out = append(out, e)
		}
	}
	return
}

func putPairToSetOfPairs[K, V comparable](in *map[K]map[V]bool, k K, v V) {
	if _, ok := (*in)[k]; !ok {
		(*in)[k] = make(map[V]bool)
	}
	(*in)[k][v] = true
}

func reverseSetOfPairs[K, V comparable](in map[K]map[V]bool) map[V]map[K]bool {
	m := make(map[V]map[K]bool)

	for k, vs := range in {
		for v, _ := range vs {
			putPairToSetOfPairs(&m, v, k)
		}
	}

	return m
}

func findAllReachableTypeByGroup(v *TypeRWWrapper, visited *map[*TypeRWWrapper]bool, result *[]*TypeRWWrapper) {
	if v.groupName != "" {
		return
	}
	if (*visited)[v] {
		return
	}
	(*visited)[v] = true
	*result = append(*result, v)

	for _, w := range v.trw.AllTypeDependencies(false) {
		findAllReachableTypeByGroup(w, visited, result)
	}
}

func (gen *Gen2) decideCppCodeDestinations(allTypes []*TypeRWWrapper) {
	const IndependentTypes = "__independent_types"
	const NoNamespaceGroup = ""

	for _, t := range allTypes {
		t.detailsFileName = t.fileName + "_details"
		t.groupName = t.tlName.Namespace
		if t.fileName != t.tlName.String() {
			for _, t2 := range allTypes {
				if t.fileName == t2.tlName.String() {
					t.groupName = t2.tlName.Namespace
					break
				}
			}
		}
	}

	allTypesWithoutGroup := make([]*TypeRWWrapper, 0)
	allTypesWithoutGroupMap := make(map[*TypeRWWrapper]bool)

	allTypesWithoutGroupUsages := make(map[*TypeRWWrapper]map[string]bool)

	for _, t := range allTypes {
		if t.groupName != NoNamespaceGroup {
			continue
		}
		allTypesWithoutGroup = append(allTypesWithoutGroup, t)
		allTypesWithoutGroupMap[t] = true
	}

	for _, t := range allTypes {
		//if t.groupName == "" {
		//	continue
		//}
		for _, dep := range t.trw.AllTypeDependencies(false) {
			if dep.groupName == NoNamespaceGroup {
				if _, ok := allTypesWithoutGroupUsages[dep]; !ok {
					allTypesWithoutGroupUsages[dep] = make(map[string]bool)
				}
				allTypesWithoutGroupUsages[dep][t.groupName] = true
			}
		}
	}

	groupToFirstVisits := reverseSetOfPairs(allTypesWithoutGroupUsages)
	for group, firstLayer := range groupToFirstVisits {
		visited := make(map[*TypeRWWrapper]bool)
		result := make([]*TypeRWWrapper, 0)

		for v, _ := range firstLayer {
			findAllReachableTypeByGroup(v, &visited, &result)
		}

		for _, v := range result {
			putPairToSetOfPairs(&allTypesWithoutGroupUsages, v, group)
		}
	}

	for _, t := range allTypesWithoutGroup {
		usages := allTypesWithoutGroupUsages[t]

		if len(usages) == 0 {
			t.groupName = IndependentTypes
		} else if len(usages) == 1 {
			usage := setToSlice(usages)[0]
			if usage != NoNamespaceGroup {
				t.groupName = usage
				t.detailsFileName = usage + "_" + t.detailsFileName
			}
		}
	}
}
