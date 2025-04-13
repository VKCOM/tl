// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlcodegen

import (
	"fmt"
	"github.com/vkcom/tl/internal/utils"
	"golang.org/x/exp/slices"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type TypeRWCPPData interface {
	CPPFillRecursiveChildren(visitedNodes map[*TypeRWWrapper]bool)
	cppTypeStringInNamespace(bytesVersion bool, hppInc *DirectIncludesCPP) string
	cppTypeStringInNamespaceHalfResolved2(bytesVersion bool, typeReduction EvaluatedType) string
	cppTypeStringInNamespaceHalfResolved(bytesVersion bool, hppInc *DirectIncludesCPP, halfResolved HalfResolvedArgument) string
	cppDefaultInitializer(halfResolved HalfResolvedArgument, halfResolve bool) string
	CPPHasBytesVersion() bool
	CPPTypeResettingCode(bytesVersion bool, val string) string
	CPPTypeWritingJsonCode(bytesVersion bool, val string, bare bool, natArgs []string, last bool) string
	CPPTypeWritingCode(bytesVersion bool, val string, bare bool, natArgs []string, last bool) string
	CPPTypeReadingCode(bytesVersion bool, val string, bare bool, natArgs []string, last bool) string
	CPPTypeJSONEmptyCondition(bytesVersion bool, val string, ref bool, deps []string) string
	CPPGenerateCode(hpp *strings.Builder, hppInc *DirectIncludesCPP, hppIncFwd *DirectIncludesCPP, hppDet *strings.Builder, hppDetInc *DirectIncludesCPP, cppDet *strings.Builder, cppDetInc *DirectIncludesCPP, bytesVersion bool, forwardDeclaration bool)
}

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
	const basicTLPackage = "basictl"

	cppAllInc := &DirectIncludesCPP{ns: map[*TypeRWWrapper]CppIncludeInfo{}}
	typesCounter := 0

	gen.decideCppCodeDestinations(gen.generatedTypesList)

	hpps := make(map[string][]*TypeRWWrapper)
	detailsHpps := make(map[string][]*TypeRWWrapper)
	detailsCpps := make(map[string][]*TypeRWWrapper)
	groupsToDetails := make(map[string]map[string]bool)

	for _, t := range gen.generatedTypesList {
		hpps[t.fileName] = append(hpps[t.fileName], t)
		detailsHpps[t.hppDetailsFileName] = append(detailsHpps[t.hppDetailsFileName], t)
		detailsCpps[t.cppDetailsFileName] = append(detailsCpps[t.cppDetailsFileName], t)

		utils.PutPairToSetOfPairs(&groupsToDetails, t.groupName, t.cppDetailsFileName)
	}

	for group, groupDetails := range groupsToDetails {
		for det := range groupDetails {
			for _, spec := range detailsCpps[det] {
				if spec.groupName != group {
					return fmt.Errorf(`in details "%s" has different groups mentioned: "%s" and "%s"`, det, group, spec.groupName)
				}
			}
		}
	}

	createdHpps := map[string]bool{}
	createdDetailsHpps := map[string]bool{}
	createdDetailsCpps := map[string]bool{}

	headers := utils.Keys(hpps)
	sort.Strings(headers)

	for _, header := range headers {
		typeDefs := hpps[header]

		var hpp strings.Builder
		hppInc := &DirectIncludesCPP{ns: map[*TypeRWWrapper]CppIncludeInfo{}}
		hppIncFwd := &DirectIncludesCPP{ns: map[*TypeRWWrapper]CppIncludeInfo{}}
		typeDefinitions := map[string]bool{}

		for _, typeRw := range typeDefs {
			typeDefVariations := make([]TypeDefinitionVariation, 1)
			{
				if typeRw.wantsBytesVersion && typeRw.trw.CPPHasBytesVersion() {
					typeDefVariations = append(typeDefVariations, TypeDefinitionVariation{NeedBytesVersion: true})
				}
			}

			for _, typeDefVariation := range typeDefVariations {
				typesCounter++
				var hppDefinition strings.Builder
				typeRw.trw.CPPGenerateCode(&hppDefinition, hppInc, hppIncFwd, nil, nil, nil, nil, typeDefVariation.NeedBytesVersion, false)
				def := hppDefinition.String()
				if !typeDefinitions[def] {
					typeDefinitions[def] = true
					hpp.WriteString(def)
				}
			}
		}

		if hpp.Len() == 0 {
			continue
		}

		filepathName := header + hppExt

		hppStr := hpp.String()
		hpp.Reset()
		hpp.WriteString("#pragma once\n\n")
		{
			hpp.WriteString(fmt.Sprintf("#include \"%s\"\n", getCppDiff(filepathName, basicCPPTLIOStreamsPath)))
		}
		for _, headerFile := range hppInc.sortedIncludes(gen.componentsOrder, func(wrapper *TypeRWWrapper) string { return wrapper.fileName }) {
			hpp.WriteString(fmt.Sprintf("#include \"%s%s\"\n", getCppDiff(filepathName, headerFile), hppExt))
		}
		hpp.WriteString("\n\n")
		hpp.WriteString(hppStr)

		if err := gen.addCodeFile(filepathName, gen.copyrightText+hpp.String()); err != nil {
			return err
		}
		hpp.Reset()

		createdHpps[header] = true
	}

	for detailsHeader, specs := range detailsHpps {
		hppDetInc := &DirectIncludesCPP{ns: map[*TypeRWWrapper]CppIncludeInfo{}}
		var hppDet strings.Builder

		slices.SortFunc(specs, TypeComparator)
		for _, typeRw := range specs {
			typeDefVariations := make([]TypeDefinitionVariation, 1)
			{
				if typeRw.wantsBytesVersion && typeRw.trw.CPPHasBytesVersion() {
					typeDefVariations = append(typeDefVariations, TypeDefinitionVariation{NeedBytesVersion: true})
				}
			}

			for _, typeDefVariation := range typeDefVariations {
				typesCounter++
				typeRw.trw.CPPGenerateCode(nil, nil, nil, &hppDet, hppDetInc, nil, nil, typeDefVariation.NeedBytesVersion, false)
			}
		}

		if hppDet.Len() == 0 {
			continue
		}

		filepathName := detailsHeader + hppExt

		hppDetStr := hppDet.String()
		hppDet.Reset()

		hppDet.WriteString("#pragma once\n\n")

		hppDet.WriteString(fmt.Sprintf("#include \"%s\"\n", getCppDiff(filepathName, basicCPPTLIOStreamsPath)))
		hppDet.WriteString(fmt.Sprintf("#include \"%s\"\n", getCppDiff(filepathName, basicCPPTLIOThrowableStreamsPath)))

		if createdHpps[specs[0].fileName] {
			hppDet.WriteString(fmt.Sprintf("#include \"%s\"\n", getCppDiff(filepathName, specs[0].fileName+hppExt)))
		}
		includes := hppDetInc.sortedIncludes(gen.componentsOrder, func(wrapper *TypeRWWrapper) string { return wrapper.fileName })
		for _, n := range includes {
			if n == specs[0].fileName {
				continue
			}
			if !createdHpps[n] {
				continue
			}
			includePath := getCppDiff(filepathName, n+hppExt)
			hppDet.WriteString(fmt.Sprintf("#include \"%s\"\n", includePath))
		}
		hppDet.WriteString("\n")

		if err := gen.addCodeFile(filepathName, gen.copyrightText+hppDet.String()+hppDetStr); err != nil {
			return err
		}
		createdDetailsHpps[detailsHeader] = true
	}

	for detailsFile, specs := range detailsCpps {
		cppDetInc := &DirectIncludesCPP{ns: map[*TypeRWWrapper]CppIncludeInfo{}}
		var cppDet strings.Builder

		slices.SortFunc(specs, TypeComparator)
		for _, typeRw := range specs {
			typeDefVariations := make([]TypeDefinitionVariation, 1)
			{
				if typeRw.wantsBytesVersion && typeRw.trw.CPPHasBytesVersion() {
					typeDefVariations = append(typeDefVariations, TypeDefinitionVariation{NeedBytesVersion: true})
				}
			}

			for _, typeDefVariation := range typeDefVariations {
				typesCounter++
				typeRw.trw.CPPGenerateCode(nil, nil, nil, nil, nil, &cppDet, cppDetInc, typeDefVariation.NeedBytesVersion, false)
			}
		}

		if cppDet.Len() == 0 {
			continue
		}

		filepathName := detailsFile + cppExt

		// all specs in one file must be in group
		cppAllInc.ns[specs[0]] = CppIncludeInfo{-1, specs[0].groupName}
		cppDetStr := cppDet.String()
		cppDet.Reset()

		for _, spec := range specs {
			cppDetInc.ns[spec] = CppIncludeInfo{componentId: spec.typeComponent, namespace: spec.groupName}
		}
		keys := utils.Keys(cppDetInc.ns)
		fmt.Sprintln(keys)
		for _, n := range cppDetInc.sortedIncludes(gen.componentsOrder, func(wrapper *TypeRWWrapper) string { return wrapper.hppDetailsFileName }) {
			if !createdDetailsHpps[n] {
				continue
			}
			cppDet.WriteString(fmt.Sprintf("#include \"%s\"\n", getCppDiff(filepathName, n+hppExt)))
		}
		cppDet.WriteString("\n")

		if err := gen.addCodeFile(filepathName, gen.copyrightText+cppDet.String()+cppDetStr); err != nil {
			return err
		}
		createdDetailsCpps[detailsFile] = true
	}

	var cppMake strings.Builder
	var cppMakeO strings.Builder
	var cppMake1 strings.Builder

	const MakefilePath = "Makefile"

	for _, nf := range cppAllInc.splitByNamespaces() {
		// it is a group
		namespace := nf.Namespace
		namespaceDetails := namespace
		namespaceDeps := nf.Includes.sortedIncludes(gen.componentsOrder, func(wrapper *TypeRWWrapper) string { return wrapper.cppDetailsFileName })

		namespaceFilePath := "details/namespaces/" + namespaceDetails + cppExt
		if !gen.options.SplitInternal {
			namespaceFilePath = namespaceDeps[0] + cppExt
		}
		buildFilePath := filepath.Join("__build", namespaceDetails+".o")

		var cppMake1UsedFiles strings.Builder
		var cppMake1Namespace strings.Builder

		for _, n := range namespaceDeps {
			cppMake1Namespace.WriteString(fmt.Sprintf("#include \"%s\"\n", getCppDiff(namespaceFilePath, n+cppExt)))
			cppMake1UsedFiles.WriteString(getCppDiff(MakefilePath, n+cppExt))

			usedTypes := detailsCpps[n]
			usedTypes = utils.FilterSlice(usedTypes, func(w *TypeRWWrapper) bool {
				return createdDetailsHpps[w.hppDetailsFileName]
			})

			hppDets := utils.MapSlice(usedTypes, func(a *TypeRWWrapper) string {
				return a.hppDetailsFileName
			})

			hppDetsList := utils.SetToSlice(utils.SliceToSet(hppDets))

			sort.Strings(hppDetsList)
			for _, h := range hppDetsList {
				cppMake1UsedFiles.WriteString(fmt.Sprintf(" %s", getCppDiff(MakefilePath, h+hppExt)))
			}
		}

		cppMake1.WriteString(fmt.Sprintf("%s: %s %s\n", buildFilePath, namespaceFilePath, cppMake1UsedFiles.String()))
		cppMake1.WriteString(fmt.Sprintf("\t$(CC) $(CFLAGS) -o %s -c %s\n", buildFilePath, namespaceFilePath))
		cppMakeO.WriteString(fmt.Sprintf("%s ", buildFilePath))

		if gen.options.SplitInternal {
			if err := gen.addCodeFile(namespaceFilePath, cppMake1Namespace.String()); err != nil {
				return err
			}
		}
	}

	metaMake := strings.Builder{}
	factoryMake := strings.Builder{}

	if err := createMeta(gen, &metaMake); err != nil {
		return err
	}
	if err := createFactory(gen, createdHpps, &factoryMake); err != nil {
		return err
	}

	cppMake.WriteString(`
CC = g++
CFLAGS = -std=c++20 -O3 -Wno-noexcept-type -g -Wall -Wextra -Werror=return-type -Wno-unused-parameter
`)
	cppMake.WriteString("# compile all object files together\n")
	cppMake.WriteString("all: ")
	cppMake.WriteString("main.o __build/io_streams.o __build/io_throwable_streams.o ")
	cppMake.WriteString(fmt.Sprintf("%s\n", cppMakeO.String()))
	cppMake.WriteString("\t$(CC) $(CFLAGS) -o all ")
	cppMake.WriteString("main.o __build/io_streams.o __build/io_throwable_streams.o ")
	cppMake.WriteString(fmt.Sprintf("%s\n", cppMakeO.String()))
	cppMake.WriteString(`
main.o: main.cpp
	$(CC) $(CFLAGS) -c main.cpp
`)
	cppMake.WriteString("\n")
	cppMake.WriteString("# compile meta sources such as meta data collection and objects factories\n")
	cppMake.WriteString(metaMake.String())
	cppMake.WriteString(factoryMake.String())

	cppMake.WriteString("\n")

	createStreams(gen, &cppMake)

	cppMake.WriteString("\n")

	cppMake.WriteString("# compile individual namespaces\n")
	cppMake.WriteString(cppMake1.String())
	//if err := gen.addCodeFile("all.cpp", cppAll.String()); err != nil {
	//	return err
	//}
	if err := gen.addCodeFile("main.cpp", "int main() { return 0; }"); err != nil {
		return err
	}
	if err := gen.addCodeFile("Makefile", cppMake.String()); err != nil {
		return err
	}
	if err := gen.addCodeFile("__build/info.txt", ".o files here!"); err != nil {
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
		basictlCppFolder := "pkg/basictl_cpp"
		exportingFiles := []string{
			"constants.h",
			"errors.h",
			"io_connectors.h",
			"io_streams.cpp",
			"io_streams.h",
			"io_throwable_streams.cpp",
			"io_throwable_streams.h",
			"impl/string_io.h",
			"impl/string_io.cpp",
			"dependencies.h",
		}

		for _, file := range exportingFiles {
			data, err := os.ReadFile(filepath.Join(basictlCppFolder, file))
			if err != nil {
				return err
			}
			code := strings.Builder{}
			code.Write([]byte(HeaderComment))
			code.Write([]byte("\n"))
			code.Write(data)
			if err := gen.addCodeFile(filepath.Join(basicTLPackage, file), code.String()); err != nil {
				return err
			}
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

func createStreams(gen *Gen2, cppMake *strings.Builder) {
	cppMake.WriteString("# compile streams which are used to work with io\n")
	cppMake.WriteString("__build/io_streams.o: basictl/constants.h basictl/errors.h basictl/io_connectors.h basictl/io_streams.cpp basictl/io_streams.h\n")
	cppMake.WriteString("\t$(CC) $(CFLAGS) -o __build/io_streams.o -c basictl/io_streams.cpp\n")

	cppMake.WriteString("\n")

	cppMake.WriteString("__build/io_throwable_streams.o: basictl/constants.h basictl/errors.h basictl/io_connectors.h basictl/io_throwable_streams.cpp basictl/io_throwable_streams.h\n")
	cppMake.WriteString("\t$(CC) $(CFLAGS) -o __build/io_throwable_streams.o -c basictl/io_throwable_streams.cpp\n")

	cppMake.WriteString("\n")

	cppMake.WriteString("__build/string_io.o: basictl/io_connectors.h basictl/impl/string_io.cpp basictl/impl/string_io.h\n")
	cppMake.WriteString("\t$(CC) $(CFLAGS) -o __build/string_io.o -c basictl/impl/string_io.cpp\n")
}

func (gen *Gen2) decideCppCodeDestinations(allTypes []*TypeRWWrapper) {
	const NoNamespaceGroup = ""
	const CommonGroup = "__common_namespace"
	const IndependentTypes = CommonGroup

	for _, t := range allTypes {
		t.cppDetailsFileName = t.fileName
		t.hppDetailsFileName = t.cppDetailsFileName
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

	allTypesWithoutGroupMap := make(map[*TypeRWWrapper]bool)
	allTypesWithoutGroupUsages := make(map[*TypeRWWrapper]map[string]bool)
	reverseDepsEdges := make(map[*TypeRWWrapper]map[*TypeRWWrapper]bool)

	for _, t := range allTypes {
		if t.groupName != NoNamespaceGroup {
			continue
		}
		allTypesWithoutGroupMap[t] = true
	}

	for _, t := range allTypes {
		for _, dep := range t.trw.AllTypeDependencies(false, true) {
			if dep.groupName == NoNamespaceGroup {
				utils.PutPairToSetOfPairs(&allTypesWithoutGroupUsages, dep, t.groupName)
				utils.PutPairToSetOfPairs(&reverseDepsEdges, dep, t)
			}
		}
	}

	// bfs
	edges := make(map[*TypeRWWrapper][]*TypeRWWrapper)
	reverseEdges := make(map[*TypeRWWrapper][]*TypeRWWrapper)

	for _, from := range allTypes {
		for _, to := range from.trw.AllTypeDependencies(false, true) {
			if to.groupName == NoNamespaceGroup {
				edges[from] = append(edges[from], to)
				reverseEdges[to] = append(reverseEdges[to], from)
			}
		}
	}

	for _, t := range allTypes {
		if t.groupName == NoNamespaceGroup && edges[t] == nil && reverseEdges[t] == nil {
			t.groupName = IndependentTypes
		}
	}

	visited := make(map[*TypeRWWrapper]bool)
	front := make(map[*TypeRWWrapper]bool)

	for t := range edges {
		if t.groupName != NoNamespaceGroup {
			front[t] = true
		} else if t.groupName == NoNamespaceGroup && len(reverseEdges[t]) == 0 {
			front[t] = true
			t.groupName = IndependentTypes
		}
	}

	utils.AppendMap(front, &visited)

	edgesCount := make(map[*TypeRWWrapper]int)

	for len(front) != 0 {
		newFront := make(map[*TypeRWWrapper]bool)
		for v := range front {
			for _, to := range edges[v] {
				if visited[to] {
					continue
				}
				if _, ok := edgesCount[to]; !ok {
					edgesCount[to] = len(reverseEdges[to])
				}
				edgesCount[to]--
				if edgesCount[to] == 0 {
					visited[to] = true
					newFront[to] = true
					groups := make(map[string]bool)
					for _, from := range reverseEdges[to] {
						groups[from.groupName] = true
					}
					if len(groups) == 1 {
						newGroup := utils.SetToSlice(groups)[0]
						changeTypeGroup(to, newGroup, CommonGroup, IndependentTypes)
					} else if len(groups) > 1 {
						to.groupName = CommonGroup
					}
				}
			}
		}
		front = newFront
	}

	for _, t := range allTypes {
		if t.groupName == NoNamespaceGroup {
			t.groupName = CommonGroup
		}
	}

	if !gen.options.SplitInternal {
		for _, t := range allTypes {
			t.cppDetailsFileName = t.groupName
		}
	}

	for _, t := range allTypes {
		typeGroup := t.tlName.Namespace
		if typeGroup == NoNamespaceGroup {
			typeGroup = CommonGroup
		}
		if strct, isStruct := t.trw.(*TypeRWStruct); isStruct && strct.ResultType != nil {
			t.fileName = filepath.Join(typeGroup, "functions", t.fileName)
		} else {
			t.fileName = filepath.Join(typeGroup, "types", t.fileName)
		}
		t.hppDetailsFileName = filepath.Join(t.groupName, "headers", t.hppDetailsFileName)
		t.cppDetailsFileName = filepath.Join(t.groupName, "details")
	}

	//printDepsGraph(allTypes, edges)
}

func printDepsGraph(allTypes []*TypeRWWrapper, edges map[*TypeRWWrapper][]*TypeRWWrapper) {
	print("debug\n")

	vertices := make([]*TypeRWWrapper, len(allTypes))
	copy(vertices, allTypes)
	slices.SortFunc(vertices, TypeComparator)

	namespaces := make(map[string][]*TypeRWWrapper)
	for _, from := range vertices {
		namespaces[from.groupName] = append(namespaces[from.groupName], from)
	}

	namespacesNames := utils.Keys(namespaces)
	sort.Strings(namespacesNames)

	fmt.Printf("digraph G {\n")

	for _, namespace := range namespacesNames {
		fmt.Printf("\tsubgraph cluster_%[1]s {\n\t\tlabel = \"%[1]s\";\n\t\tcolor=lightgrey;\n\t\tstyle=filled;\n", namespace)
		for _, from := range namespaces[namespace] {
			color := "red"
			if from.trw.IsWrappingType() {
				color = "blue"
			}
			cppName := from.cppLocalName
			if cppName == "" {
				cppName = "__empty__"
			}
			fmt.Printf("\t\t%[1]s[color=\"%[2]s\", label=\"%[1]s,\\n%[3]s\", shape=box];\n", from.goGlobalName, color, cppName)
		}
		fmt.Printf("}\n")
	}

	for _, from := range vertices {
		for _, to := range edges[from] {
			fmt.Printf("\t%s->%s;\n", from.goGlobalName, to.goGlobalName)
		}
	}
	fmt.Printf("}\n")
}

func changeTypeGroup(to *TypeRWWrapper, newGroup string, CommonGroup string, IndependentTypes string) {
	to.groupName = newGroup
	if to.groupName != CommonGroup && to.groupName != IndependentTypes {
		to.cppDetailsFileName = to.groupName + "_" + to.cppDetailsFileName
	}
	to.hppDetailsFileName = to.cppDetailsFileName
}

func getCppDiff(base string, target string) string {
	dir1, _ := filepath.Split(base)
	dir2, file := filepath.Split(target)
	diff, _ := filepath.Rel(dir1, dir2)
	return filepath.Join(diff, file)
}

func createMeta(gen *Gen2, make *strings.Builder) error {
	meta := strings.Builder{}
	metaDetails := strings.Builder{}

	filepathName := filepath.Join("__meta", "headers"+hppExt)
	filepathDetailsName := filepath.Join("__meta", "details"+cppExt)

	meta.WriteString(fmt.Sprintf(`
#pragma once

#include <ostream>
#include <string>
#include <functional>

#include "%[1]s"
#include "%[2]s"

namespace tl2 {
    namespace meta {
        struct tl_object {
            virtual bool read(::basictl::tl_istream &s) = 0;
            virtual bool write(::basictl::tl_ostream &s) = 0;

			virtual void read_or_throw(::basictl::tl_throwable_istream &s) = 0;
            virtual void write_or_throw(::basictl::tl_throwable_ostream &s) = 0;

            virtual bool read_boxed(::basictl::tl_istream &s) = 0;
            virtual bool write_boxed(::basictl::tl_ostream &s) = 0;

			virtual void read_boxed_or_throw(::basictl::tl_throwable_istream &s) = 0;
            virtual void write_boxed_or_throw(::basictl::tl_throwable_ostream &s) = 0;
			
			virtual bool write_json(std::ostream &s) = 0;

            virtual ~tl_object() = default;
        };

        struct tl_function : public tl_object {
            virtual bool read_write_result(::basictl::tl_istream &in, ::basictl::tl_ostream &out) = 0;

            virtual ~tl_function() = default;
        };

        struct tl_item {
            uint32_t tag{};
            uint32_t annotations{};
            std::string name;

            std::function<std::unique_ptr<tl2::meta::tl_object>()> create_object;
            std::function<std::unique_ptr<tl2::meta::tl_function>()> create_function;
        };

		tl2::meta::tl_item get_item_by_name(std::string &&s);
		tl2::meta::tl_item get_item_by_tag(uint32_t &&tag);

		void set_create_object_by_name(std::string &&s, std::function<std::unique_ptr<tl2::meta::tl_object>()> &&factory);
		void set_create_function_by_name(std::string &&s, std::function<std::unique_ptr<tl2::meta::tl_function>()> &&factory);
        
    }
}`, getCppDiff(filepathName, basicCPPTLIOStreamsPath), getCppDiff(filepathName, basicCPPTLIOStreamsPath)))

	metaDetails.WriteString(fmt.Sprintf("#include \"%s\"\n", getCppDiff(filepathDetailsName, basicCPPTLIOStreamsPath)))
	metaDetails.WriteString(fmt.Sprintf(`
#include <map>

#include "%s"

namespace {
	struct tl_items {
		public:
			std::map<std::string, std::shared_ptr<tl2::meta::tl_item>> items;
			std::map<uint32_t, std::shared_ptr<tl2::meta::tl_item>> items_by_tag;
			tl_items();
	};
    
	tl_items items;
    std::function<std::unique_ptr<tl2::meta::tl_object>()> no_object_generator = []() -> std::unique_ptr<tl2::meta::tl_object> {
        throw std::runtime_error("no generation for this type of objects");
    };
    std::function<std::unique_ptr<tl2::meta::tl_function>()> no_function_generator = []() -> std::unique_ptr<tl2::meta::tl_function> {
        throw std::runtime_error("no generation for this type of functions");
    };
}

tl2::meta::tl_item tl2::meta::get_item_by_name(std::string &&s) {
    auto item = items.items.find(s);
	if (item != items.items.end()) {
        return *item->second;
    }
    throw std::runtime_error("no item with such name + \"" + s + "\"");
}

tl2::meta::tl_item tl2::meta::get_item_by_tag(std::uint32_t &&tag) {
    auto item = items.items_by_tag.find(tag);
	if (item != items.items_by_tag.end()) {
        return *item->second;
    }
    throw std::runtime_error("no item with such tag + \"" + std::to_string(tag) + "\"");
}

void tl2::meta::set_create_object_by_name(std::string &&s, std::function<std::unique_ptr<tl2::meta::tl_object>()>&& gen) {
    auto item = items.items.find(s);
	if (item != items.items.end()) {
        item->second->create_object = gen;
		return;	
    }
    throw std::runtime_error("no item with such name + \"" + s + "\"");
}

void tl2::meta::set_create_function_by_name(std::string &&s, std::function<std::unique_ptr<tl2::meta::tl_function>()>&& gen) {
    auto item = items.items.find(s);
	if (item != items.items.end()) {
        item->second->create_function = gen;
		return;	
    }
    throw std::runtime_error("no item with such name + \"" + s + "\"");
}

tl_items::tl_items() {`, getCppDiff(filepathDetailsName, filepathName)))

	for _, wr := range gen.generatedTypesList {
		if wr.tlTag == 0 || !wr.IsTopLevel() {
			continue
		}
		if _, isStruct := wr.trw.(*TypeRWStruct); isStruct {
			metaDetails.WriteString(
				fmt.Sprintf(`
	auto item%[4]d = std::shared_ptr<tl2::meta::tl_item>(new tl2::meta::tl_item{.tag=%[2]s,.annotations=%[3]s,.name="%[1]s",.create_object=no_object_generator,.create_function=no_function_generator});
	(this->items)["%[1]s"] = item%[4]d;
	(this->items_by_tag)[%[2]s] = item%[4]d;`,
					wr.tlName.String(),
					fmt.Sprintf("0x%08x", wr.tlTag),
					fmt.Sprintf("0x%x", wr.AnnotationsMask()),
					wr.tlTag,
				),
			)
		}
	}

	metaDetails.WriteString(`
}
`)
	if err := gen.addCodeFile(filepathName, gen.copyrightText+meta.String()); err != nil {
		return err
	}
	if err := gen.addCodeFile(filepathDetailsName, gen.copyrightText+metaDetails.String()); err != nil {
		return err
	}

	make.WriteString(fmt.Sprintf(`__build/__meta.o: %[1]s %[2]s
	$(CC) $(CFLAGS) -o __build/__meta.o -c %[2]s
`,
		filepathName,
		filepathDetailsName,
	))
	return nil
}

func createFactory(gen *Gen2, createdHpps map[string]bool, make *strings.Builder) error {
	factory := strings.Builder{}
	factoryDetails := strings.Builder{}

	filepathName := filepath.Join("__factory", "headers"+hppExt)
	filepathNameDetails := filepath.Join("__factory", "details"+cppExt)

	imports := DirectIncludesCPP{ns: map[*TypeRWWrapper]CppIncludeInfo{}}

	factory.WriteString(`
#pragma once
namespace tl2 {
    namespace factory {    
		void set_all_factories();
	}
}`)

	factoryDetails.WriteString(`
void tl2::factory::set_all_factories() {
`)

	for _, wr := range gen.generatedTypesList {
		if wr.tlTag == 0 || !wr.IsTopLevel() {
			continue
		}
		if strct, isStruct := wr.trw.(*TypeRWStruct); isStruct && len(wr.NatParams) == 0 {
			if strct.isTypeDef() {
				continue
			}
			hppTmpInclude := DirectIncludesCPP{ns: map[*TypeRWWrapper]CppIncludeInfo{}}
			myFullType := wr.trw.cppTypeStringInNamespace(wr.wantsBytesVersion && wr.trw.CPPHasBytesVersion(), &hppTmpInclude)
			myFullTypeNoPrefix := strings.TrimPrefix(myFullType, "::") // Stupid C++ has sometimes problems with name resolution of definitions
			myFullTypeWithUnderlines := strings.ReplaceAll(myFullTypeNoPrefix, "::", "_")

			imports.ns[wr] = CppIncludeInfo{componentId: wr.typeComponent, namespace: wr.groupName}

			implementedInterface := "tl_object"
			if strct.ResultType != nil {
				implementedInterface = "tl_function"
			}

			factoryDetails.WriteString(fmt.Sprintf(`
	struct %[3]s_%[1]s : public tl2::meta::%[1]s {
        %[2]s object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}

		void read_or_throw(::basictl::tl_throwable_istream &s) override { object.read_or_throw(s);}
		void write_or_throw(::basictl::tl_throwable_ostream &s) override { object.write_or_throw(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		void read_boxed_or_throw(::basictl::tl_throwable_istream &s) override { object.read_boxed_or_throw(s);}
		void write_boxed_or_throw(::basictl::tl_throwable_ostream &s) override { object.write_boxed_or_throw(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}
`,
				implementedInterface,
				myFullTypeNoPrefix,
				myFullTypeWithUnderlines,
			))
			if strct.ResultType != nil {
				hppTmpInclude2 := DirectIncludesCPP{ns: map[*TypeRWWrapper]CppIncludeInfo{}}
				resultType := strct.ResultType.trw.cppTypeStringInNamespace(wr.wantsBytesVersion && wr.trw.CPPHasBytesVersion(), &hppTmpInclude2)
				resultTypeNoPrefix := strings.TrimPrefix(resultType, "::") // Stupid C++ has sometimes problems with name resolution of definitions
				imports.ns[strct.ResultType] = CppIncludeInfo{componentId: strct.ResultType.typeComponent, namespace: strct.ResultType.groupName}

				factoryDetails.WriteString(fmt.Sprintf(`
		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			%[1]s result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}
`,
					resultTypeNoPrefix,
				))
			}
			factoryDetails.WriteString(`
    };`)
			factoryDetails.WriteString(fmt.Sprintf(`
	tl2::meta::set_create_object_by_name("%[1]s", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<%[2]s_%[3]s>();
	});
`,
				wr.tlName.String(),
				myFullTypeWithUnderlines,
				implementedInterface,
				myFullTypeNoPrefix,
			))
			if strct.ResultType != nil {
				factoryDetails.WriteString(fmt.Sprintf(`
	tl2::meta::set_create_function_by_name("%[1]s", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<%[2]s_%[3]s>();
	});
`,
					wr.tlName.String(),
					myFullTypeWithUnderlines,
					implementedInterface,
					myFullTypeNoPrefix,
				))
			}
		}
	}

	factoryDetails.WriteString(`
}
`)
	suffix := factoryDetails.String()
	factoryDetails.Reset()
	factoryDetails.WriteString(fmt.Sprintf("#include \"%s\"\n", getCppDiff(filepathNameDetails, filepath.Join("__meta", "headers"+hppExt))))
	factoryDetails.WriteString(fmt.Sprintf("#include \"%s\"\n\n", getCppDiff(filepathNameDetails, filepathName)))

	factoryFileDependencies := strings.Builder{}

	for _, headerFile := range imports.sortedIncludes(gen.componentsOrder, func(wrapper *TypeRWWrapper) string { return wrapper.fileName }) {
		if !createdHpps[headerFile] {
			continue
		}
		factoryDetails.WriteString(fmt.Sprintf("#include \"%s%s\"\n", getCppDiff(filepathName, headerFile), hppExt))
		factoryFileDependencies.WriteString(" " + headerFile + hppExt)
	}
	factoryDetails.WriteString(suffix)

	if err := gen.addCodeFile(filepathName, gen.copyrightText+factory.String()); err != nil {
		return err
	}
	if err := gen.addCodeFile(filepathNameDetails, gen.copyrightText+factoryDetails.String()); err != nil {
		return err
	}

	make.WriteString(fmt.Sprintf(`__build/__factory.o: %[1]s %[2]s%[3]s
	$(CC) $(CFLAGS) -o __build/__factory.o -c %[2]s
`,
		filepathName,
		filepathNameDetails,
		factoryFileDependencies.String(),
	))

	return nil
}
