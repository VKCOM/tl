// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlcodegen

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/vkcom/tl/internal/utils"
	"golang.org/x/exp/slices"
)

type TypeRWCPPData interface {
	cppTypeStringInNamespace(bytesVersion bool, hppInc *DirectIncludesCPP) string
	cppTypeStringInNamespaceHalfResolved2(bytesVersion bool, typeReduction EvaluatedType) string
	cppTypeStringInNamespaceHalfResolved(bytesVersion bool, hppInc *DirectIncludesCPP, halfResolved HalfResolvedArgument) string
	cppDefaultInitializer(halfResolved HalfResolvedArgument, halfResolve bool) string

	CPPAllowCurrentDefinition() bool
	CPPFillRecursiveChildren(visitedNodes map[*TypeRWWrapper]bool)
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

func (gen *Gen2) generateCodeCPP(bytesWhiteList []string) error {
	cppAllInc := &DirectIncludesCPP{ns: map[*TypeRWWrapper]CppIncludeInfo{}}
	typesCounter := 0

	_deps := gen.decideCppCodeDestinations(gen.generatedTypesList)

	_, err := gen.createDependencies(_deps)
	if err != nil {
		return err
	}

	hpps := make(map[string][]*TypeRWWrapper)
	detailsHpps := make(map[string][]*TypeRWWrapper)
	detailsCpps := make(map[string][]*TypeRWWrapper)
	groupsToDetails := make(map[string]map[string]bool)

	for _, t := range gen.generatedTypesList {
		hpps[t.fileName] = append(hpps[t.fileName], t)

		if t.groupName != GhostTypes {
			detailsHpps[t.hppDetailsFileName] = append(detailsHpps[t.hppDetailsFileName], t)
			detailsCpps[t.cppDetailsFileName] = append(detailsCpps[t.cppDetailsFileName], t)

			utils.PutPairToSetOfPairs(&groupsToDetails, t.groupName, t.cppDetailsFileName)
		}
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
				if typeRw.trw.CPPAllowCurrentDefinition() {
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
		}

		if hpp.Len() == 0 {
			continue
		}

		filepathName := header + hppExt

		hppStr := hpp.String()
		hpp.Reset()
		hpp.WriteString("#pragma once\n\n")
		{
			hpp.WriteString(fmt.Sprintf("#include \"%s\"\n", filepath.Join(gen.options.RootCPP, CppBasicTLIOStreamsPath(gen))))
			hpp.WriteString(fmt.Sprintf("#include \"%s\"\n", filepath.Join(gen.options.RootCPP, CppBasicTLIOThrowableStreamsPath(gen))))
		}
		for _, headerFile := range hppInc.sortedIncludes(gen.componentsOrder, func(wrapper *TypeRWWrapper) string { return wrapper.fileName }) {
			hpp.WriteString(fmt.Sprintf("#include \"%s%s\"\n", filepath.Join(gen.options.RootCPP, headerFile), hppExt))
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

		hppDet.WriteString(fmt.Sprintf("#include \"%s\"\n", filepath.Join(gen.options.RootCPP, CppBasicTLIOStreamsPath(gen))))
		hppDet.WriteString(fmt.Sprintf("#include \"%s\"\n", filepath.Join(gen.options.RootCPP, CppBasicTLIOThrowableStreamsPath(gen))))

		if createdHpps[specs[0].fileName] {
			hppDet.WriteString(fmt.Sprintf("#include \"%s\"\n", filepath.Join(gen.options.RootCPP, specs[0].fileName+hppExt)))
		}
		includes := hppDetInc.sortedIncludes(gen.componentsOrder, func(wrapper *TypeRWWrapper) string { return wrapper.fileName })
		for _, n := range includes {
			if n == specs[0].fileName {
				continue
			}
			if !createdHpps[n] {
				continue
			}
			includePath := filepath.Join(gen.options.RootCPP, n+hppExt)
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
		//keys := utils.Keys(cppDetInc.ns)
		//fmt.Println(keys)
		for _, n := range cppDetInc.sortedIncludes(gen.componentsOrder, func(wrapper *TypeRWWrapper) string { return wrapper.hppDetailsFileName }) {
			if !createdDetailsHpps[n] {
				continue
			}
			cppDet.WriteString(fmt.Sprintf("#include \"%s\"\n", filepath.Join(gen.options.RootCPP, n+hppExt)))
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

		namespaceFilePath := filepath.Join("details", "namespaces", namespaceDetails+cppExt)
		if !gen.options.SplitInternal {
			namespaceFilePath = namespaceDeps[0] + cppExt
		}
		buildFilePath := filepath.Join("__build", namespaceDetails+".o")

		var cppMake1UsedFiles strings.Builder
		var cppMake1Namespace strings.Builder

		for _, n := range namespaceDeps {
			cppMake1Namespace.WriteString(fmt.Sprintf("#include \"%s\"\n", filepath.Join(gen.options.RootCPP, n+cppExt)))
			//cppMake1UsedFiles.WriteString(getCppDiff(MakefilePath, n+cppExt))

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
		cppMake1.WriteString(fmt.Sprintf("\t@mkdir -p __build\n\t$(CC) $(CFLAGS) -I. -o %s -c %s\n", buildFilePath, namespaceFilePath))
		cppMakeO.WriteString(fmt.Sprintf("%s ", buildFilePath))

		if gen.options.SplitInternal {
			if err := gen.addCodeFile(namespaceFilePath, gen.copyrightText+cppMake1Namespace.String()); err != nil {
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

	appendCompilerOptions(&cppMake)

	cppMake.WriteString("# compile all object files together\n")
	cppMake.WriteString("all: ")
	cppMake.WriteString("__build/main.o __build/io_streams.o __build/io_throwable_streams.o ")
	cppMake.WriteString(fmt.Sprintf("%s\n", cppMakeO.String()))
	cppMake.WriteString("\t@mkdir -p __build\n")
	cppMake.WriteString("\t$(CC) $(CFLAGS) -o all ")
	cppMake.WriteString("__build/main.o __build/io_streams.o __build/io_throwable_streams.o ")
	cppMake.WriteString(fmt.Sprintf("%s\n", cppMakeO.String()))
	cppMake.WriteString(`
__build/main.o: main.cpp
	@mkdir -p __build
	$(CC) $(CFLAGS) -c main.cpp -o __build/main.o
`)
	cppMake.WriteString("\n")
	cppMake.WriteString(metaMake.String())
	cppMake.WriteString(factoryMake.String())

	createStreams(gen, &cppMake)

	cppMake.WriteString("\n")

	cppMake.WriteString("# build object files for individual namespaces\n")
	cppMake.WriteString(cppMake1.String())

	if gen.options.GenerateCommonMakefile {
		if err := gen.addCodeFile("main.cpp", "int main() { return 0; }"); err != nil {
			return err
		}
		if err := gen.addCodeFile("Makefile", cppMake.String()); err != nil {
			return err
		}
	}
	if err = gen.addCPPBasicTLFiles(); err != nil {
		return err
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

func (gen *Gen2) createDependencies(directDeps map[string]map[string]bool) (map[string]map[string]bool, error) {
	deps := make(map[string]map[string]bool)

	for ns := range directDeps {
		visited := make(map[string]bool)
		stack := make([]string, 0)

		stack = append(stack, ns)

		for len(stack) > 0 {
			current := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if visited[current] {
				continue
			}
			visited[current] = true

			for dep := range directDeps[current] {
				if len(deps[ns]) == 0 {
					deps[ns] = make(map[string]bool)
				}
				stack = append(stack, dep)
				deps[ns][dep] = true
			}
		}
	}

	depsList := utils.Keys(deps)
	sort.Strings(depsList)

	for _, ns := range depsList {
		if cppIsSpecialNamespace(ns) && ns != CommonGroup {
			continue
		}
		nsDeps := utils.Keys(deps[ns])
		sort.Strings(nsDeps)

		code := strings.Builder{}

		code.WriteString("{\n")

		code.WriteString("\t\"io\": [")
		code.WriteString(fmt.Sprintf("\"%s\"", CppBasictlPackage(gen)))
		code.WriteString("],\n")

		code.WriteString("\t\"dependencies\":[")

		wasFirst := false
		for _, dep := range nsDeps {
			if (cppIsSpecialNamespace(dep) && dep != CommonGroup) || dep == ns {
				continue
			}
			if wasFirst {
				code.WriteString(", ")
			}
			code.WriteString(fmt.Sprintf("\"%s\"", dep))
			wasFirst = true
		}

		code.WriteString("]\n")
		code.WriteString("}")

		//if ns == CommonGroup && wasFirst {
		//	return nil, fmt.Errorf("tlgen bug: %s is not independent", CommonGroup)
		//}
		if err := gen.addCodeFile(filepath.Join(ns, "info.json"), code.String()); err != nil {
			return nil, err
		}
	}

	// create ns dependencies info file
	{
		allVertices := make(map[string]bool)
		for v, us := range directDeps {
			allVertices[v] = true
			for u := range us {
				allVertices[u] = true
			}
		}

		delete(allVertices, GhostTypes)

		order := make([]string, 0)
		visited := make(map[string]bool)
		component := make([]string, 0)
		components := make([][]string, 0)

		var df1 func(v string)
		df1 = func(v string) {
			visited[v] = true

			directDepsList := utils.Keys(directDeps[v])
			sort.Strings(directDepsList)

			for _, u := range directDepsList {
				if !visited[u] && allVertices[u] {
					df1(u)
				}
			}
			order = append(order, v)
		}

		revDirectDeps := make(map[string]map[string]bool)
		for v, us := range directDeps {
			for u := range us {
				if _, ok := revDirectDeps[u]; !ok {
					revDirectDeps[u] = make(map[string]bool)
				}
				revDirectDeps[u][v] = true
			}
		}

		var df2 func(v string)
		df2 = func(v string) {
			visited[v] = true
			component = append(component, v)
			revDepsList := utils.Keys(revDirectDeps[v])
			sort.Strings(revDepsList)

			for _, u := range revDepsList {
				if !visited[u] && allVertices[u] {
					df2(u)
				}
			}
		}

		visited = make(map[string]bool)

		listOfVertices := utils.Keys(allVertices)
		sort.Strings(listOfVertices)

		for _, v := range listOfVertices {
			if !visited[v] {
				df1(v)
			}
		}

		visited = make(map[string]bool)
		for i := range order {
			v := order[len(order)-1-i]
			if !visited[v] {
				df2(v)
				components = append(components, component)
				component = make([]string, 0)
			}
		}

		for i := range components {
			sort.Strings(components[i])
		}

		nsToComponentId := make(map[string]int)
		for id, comp := range components {
			for _, ns := range comp {
				nsToComponentId[ns] = id
			}
		}

		compDeps := make(map[int]map[int]bool)
		for ns, nsDeps := range directDeps {
			for dep := range nsDeps {
				if _, ok := compDeps[nsToComponentId[ns]]; !ok {
					compDeps[nsToComponentId[ns]] = make(map[int]bool)
				}
				compDeps[nsToComponentId[ns]][nsToComponentId[dep]] = true
			}
		}

		compVisited := make(map[int]bool)
		compOrder := make([]int, 0)

		var compDfs func(c int)
		compDfs = func(c int) {
			compVisited[c] = true
			keys := utils.Keys(compDeps[c])
			sort.Ints(keys)
			for _, cDep := range keys {
				if !compVisited[cDep] {
					compDfs(cDep)
				}
			}
			compOrder = append(compOrder, c)
		}

		for i := range components {
			if !compVisited[i] {
				compDfs(i)
			}
		}

		code := strings.Builder{}
		code.WriteString("{\n")
		code.WriteString("\t\"order\": [\n")

		for i := range compOrder {
			orderId := i
			compId := compOrder[orderId]
			code.WriteString("\t\t[")
			code.WriteString(
				strings.Join(
					utils.MapSlice(
						components[compId],
						func(s string) string { return fmt.Sprintf("\"%s\"", s) },
					),
					", ",
				),
			)
			if i == len(compOrder)-1 {
				code.WriteString("]\n")
			} else {
				code.WriteString("],\n")
			}
		}

		code.WriteString("\t]\n")
		code.WriteString("}")
		if err := gen.addCodeFile("info.json", code.String()); err != nil {
			return nil, err
		}
	}
	return deps, nil
}

func (gen *Gen2) addCPPBasicTLFiles() error {
	exportingFiles := CppCopingStreamFiles()
	exportingFilesSources := CppCopingStreamFilesText()

	for _, file := range exportingFiles {
		data := []byte(exportingFilesSources[file])

		data = bytes.ReplaceAll(data, []byte("basictl"), []byte(gen.options.BasicTLNamespace))

		code := strings.Builder{}
		code.Write([]byte(HeaderComment))
		code.Write([]byte("\n"))

		includesStart := bytes.Index(data, []byte(basictlCppIncludeStart))
		includesEnd := bytes.Index(data, []byte(basictlCppIncludeEnd))

		if includesStart != -1 && includesEnd == -1 {
			return fmt.Errorf("can't locate include block")
		}

		if includesStart != -1 {
			code.Write(data[:includesStart])

			includes := data[includesStart+len(basictlCppIncludeStart) : includesEnd]
			for _, include := range bytes.Split(includes, []byte("\n")) {
				if len(include) != 0 {
					parts := bytes.Split(include, []byte("\""))
					code.Write(parts[0])
					code.Write([]byte("\""))
					code.Write([]byte(filepath.Join(gen.options.RootCPP, CppBasictlPackage(gen), filepath.Dir(file), string(parts[1]))))
					code.Write([]byte("\""))
					code.Write(parts[2])

					code.Write([]byte("\n"))
				}
			}

			bodyCode := data[includesEnd+len(basictlCppIncludeEnd):]
			bodyCode = bytes.ReplaceAll(bodyCode, []byte(CppBasictlPackage(gen)), []byte(gen.cppBasictlNamespace()))

			code.Write(bodyCode)
		} else {
			bodyCode := data
			bodyCode = bytes.ReplaceAll(bodyCode, []byte(CppBasictlPackage(gen)), []byte(gen.cppBasictlNamespace()))

			code.Write(bodyCode)
		}

		if err := gen.addCodeFile(filepath.Join(CppBasictlPackage(gen), file), gen.copyrightText+code.String()); err != nil {
			return err
		}
	}
	{
		code := strings.Builder{}
		code.WriteString("{\n\t\"sources\":[")
		wasFirst := false
		for _, file := range exportingFiles {
			if strings.HasSuffix(file, cppExt) && !strings.Contains(file, "/") {
				if wasFirst {
					code.WriteString(", ")
				}
				wasFirst = true
				code.WriteString(fmt.Sprintf("\"%s\"", file))
			}
		}
		code.WriteString("]\n}")
		if err := gen.addCodeFile(filepath.Join(CppBasictlPackage(gen), "info.json"), code.String()); err != nil {
			return err
		}
	}
	return nil
}

func createStreams(gen *Gen2, cppMake *strings.Builder) {
	cppMake.WriteString("# compile streams which are used to work with io\n")
	cppMake.WriteString(fmt.Sprintf("__build/io_streams.o: %[1]s/constants.h %[1]s/errors.h %[1]s/io_connectors.h %[1]s/io_streams.cpp %[1]s/io_streams.h\n", CppBasictlPackage(gen)))
	cppMake.WriteString(fmt.Sprintf("\t@mkdir -p __build\n\t$(CC) $(CFLAGS) -I. -o __build/io_streams.o -c %[1]s/io_streams.cpp\n", CppBasictlPackage(gen)))

	cppMake.WriteString("\n")

	cppMake.WriteString(fmt.Sprintf("__build/io_throwable_streams.o: %[1]s/constants.h %[1]s/errors.h %[1]s/io_connectors.h %[1]s/io_throwable_streams.cpp %[1]s/io_throwable_streams.h\n", CppBasictlPackage(gen)))
	cppMake.WriteString(fmt.Sprintf("\t@mkdir -p __build\n\t$(CC) $(CFLAGS) -I. -o __build/io_throwable_streams.o -c %[1]s/io_throwable_streams.cpp\n", CppBasictlPackage(gen)))

	cppMake.WriteString("\n")

	cppMake.WriteString(fmt.Sprintf("__build/string_io.o: %[1]s/io_connectors.h %[1]s/impl/string_io.cpp %[1]s/impl/string_io.h\n", CppBasictlPackage(gen)))
	cppMake.WriteString(fmt.Sprintf("\t@mkdir -p __build\n\t$(CC) $(CFLAGS) -I. -o __build/string_io.o -c %[1]s/impl/string_io.cpp\n", CppBasictlPackage(gen)))
}

func (gen *Gen2) decideCppCodeDestinations(allTypes []*TypeRWWrapper) map[string]map[string]bool {
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

	// bfs
	edges := make(map[*TypeRWWrapper][]*TypeRWWrapper)
	reverseEdges := make(map[*TypeRWWrapper][]*TypeRWWrapper)

	for _, t := range allTypes {
		deps := t.trw.AllTypeDependencies(false, true)
		for _, dep := range deps {
			utils.PutPairToSetOfPairs(&allTypesWithoutGroupUsages, dep, t.groupName)
			utils.PutPairToSetOfPairs(&reverseDepsEdges, dep, t)

			edges[t] = append(edges[t], dep)
			reverseEdges[dep] = append(reverseEdges[dep], t)
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
		if !t.trw.IsWrappingType() && t.groupName != NoNamespaceGroup {
			front[t] = true
		} else if t.groupName == NoNamespaceGroup && len(reverseEdges[t]) == 0 {
			front[t] = true
			if t.trw.IsWrappingType() { // || strings.HasPrefix(strings.ToLower(t.cppLocalName), strings.ToLower("DictionaryField")) {
				t.groupName = GhostTypes
			} else {
				gen.decideGroupInConflict(t, edges, nil)
			}
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
					if len(groups) == 1 || (len(groups) == 2 && groups[GhostTypes]) {
						if len(groups) == 2 && groups[GhostTypes] {
							delete(groups, GhostTypes)
						}
						newGroup := utils.SetToSlice(groups)[0]
						oldGroup := to.groupName
						to.groupName = newGroup
						if to.groupName != CommonGroup && to.groupName != IndependentTypes && to.groupName != GhostTypes && oldGroup != newGroup {
							to.cppDetailsFileName = to.groupName + "_" + to.cppDetailsFileName
						}
						to.hppDetailsFileName = to.cppDetailsFileName
					} else if len(groups) > 1 {
						//if gen.options.LocalizeNamespaces {
						currentGroup := to.groupName
						if cppIsSpecialNamespace(currentGroup) {
							gen.decideGroupInConflict(to, edges, groups)
						}
						//} else {
						//	to.groupName = CommonGroup
						//}
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

	if CppPrintGraphvizRepresentation {
		//file := os.Stdout
		file, _ := os.Create("out/graphviz.dot")
		printDepsGraph(file, allTypes, edges)
	}

	deps := getNamespacesDependencies(allTypes, edges)

	if CppPrintNamespaceDependencies {
		nss := utils.Keys(deps)
		sort.Strings(nss)

		for _, ns := range nss {
			nsDeps := utils.Keys(deps[ns])
			sort.Strings(nsDeps)

			fmt.Printf("%s (%d):", ns, len(nsDeps))

			for _, dep := range nsDeps {
				fmt.Printf(" %s", dep)
			}
			fmt.Printf("\n")
		}
	}

	return deps
}

func (gen *Gen2) decideGroupInConflict(to *TypeRWWrapper, edges map[*TypeRWWrapper][]*TypeRWWrapper, groups map[string]bool) {
	deps := getSpeculativeGroupDependencies(to, &edges)
	deps = utils.FilterSlice(deps, func(group string) bool {
		return !cppIsSpecialNamespace(group)
	})
	groupsList := utils.SetToSlice(groups)
	sort.Strings(groupsList)

	//fmt.Printf("%s : <-%s, ->%s\n", to.goGlobalName, groupsList, deps)

	if len(deps) == 0 {
		to.groupName = CommonGroup
	} else {
		to.groupName = deps[0]
		to.cppDetailsFileName = to.groupName + "_" + to.cppDetailsFileName
		to.hppDetailsFileName = to.cppDetailsFileName
	}
}

func printDepsGraph(out *os.File, allTypes []*TypeRWWrapper, edges map[*TypeRWWrapper][]*TypeRWWrapper) {
	vertices := make([]*TypeRWWrapper, len(allTypes))
	copy(vertices, allTypes)
	slices.SortFunc(vertices, TypeComparator)

	namespaces := make(map[string][]*TypeRWWrapper)
	for _, from := range vertices {
		namespaces[from.groupName] = append(namespaces[from.groupName], from)
	}

	namespacesNames := utils.Keys(namespaces)
	sort.Strings(namespacesNames)

	_, _ = fmt.Fprintf(out, "digraph G {\n")

	for _, namespace := range namespacesNames {
		_, _ = fmt.Fprintf(out, "\tsubgraph cluster_%[1]s {\n\t\tlabel = \"%[1]s\";\n\t\tcolor=lightgrey;\n\t\tstyle=filled;\n", namespace)
		for _, from := range namespaces[namespace] {
			color := "red"
			if from.trw.IsWrappingType() {
				color = "blue"
			}
			cppName := from.cppLocalName
			if cppName == "" {
				cppName = "__empty__"
			}
			_, _ = fmt.Fprintf(out, "\t\t%[1]s[color=\"%[2]s\", label=\"%[1]s,\\n%[3]s\", shape=box];\n", from.goGlobalName, color, cppName)
		}
		_, _ = fmt.Fprintf(out, "}\n")
	}

	for _, from := range vertices {
		for _, to := range edges[from] {
			_, _ = fmt.Fprintf(out, "\t%s->%s;\n", from.goGlobalName, to.goGlobalName)
		}
	}
	_, _ = fmt.Fprintf(out, "}\n")
}

func getNamespacesDependencies(allTypes []*TypeRWWrapper, edges map[*TypeRWWrapper][]*TypeRWWrapper) map[string]map[string]bool {
	result := make(map[string]map[string]bool)

	for _, exactType := range allTypes {
		// to have at least zero deps
		if _, ok := result[exactType.groupName]; !ok {
			result[exactType.groupName] = make(map[string]bool)
		}
		for _, dep := range edges[exactType] {
			if len(result[exactType.groupName]) == 0 {
				result[exactType.groupName] = make(map[string]bool)
			}
			result[exactType.groupName][dep.groupName] = true
		}
	}

	return result
}

func getSpeculativeGroupDependencies(start *TypeRWWrapper, edges *map[*TypeRWWrapper][]*TypeRWWrapper) []string {
	result := make(map[string]bool)
	visited := make(map[*TypeRWWrapper]bool)

	stack := make([]*TypeRWWrapper, 0)
	stack = append(stack, start)

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if visited[current] {
			continue
		}
		visited[current] = true
		result[current.groupName] = true

		stack = append(stack, (*edges)[current]...)
	}

	keys := utils.Keys(result)
	sort.Strings(keys)

	return keys
}

func getCppDiff(base string, target string) string {
	dir1, _ := filepath.Split(base)
	dir2, file := filepath.Split(target)
	diff, _ := filepath.Rel(dir1, dir2)
	return filepath.Join(diff, file)
}

func createMeta(gen *Gen2, make *strings.Builder) error {
	if !gen.options.AddMetaData {
		return nil
	}
	meta := strings.Builder{}
	metaDetails := strings.Builder{}

	filepathName := filepath.Join("__meta", "headers"+hppExt)
	filepathDetailsName := filepath.Join("__meta", "details"+cppExt)

	meta.WriteString(fmt.Sprintf(`
#pragma once

#include <ostream>
#include <string>
#include <functional>
#include <optional>

#include "%[1]s"
#include "%[2]s"

namespace %[3]s {
	namespace meta {
		struct tl_object {
			virtual bool read(::%[4]s::tl_istream &s) = 0;
			virtual bool write(::%[4]s::tl_ostream &s) = 0;

			virtual void read(::%[4]s::tl_throwable_istream &s) = 0;
			virtual void write(::%[4]s::tl_throwable_ostream &s) = 0;

			virtual bool read_boxed(::%[4]s::tl_istream &s) = 0;
			virtual bool write_boxed(::%[4]s::tl_ostream &s) = 0;

			virtual void read_boxed(::%[4]s::tl_throwable_istream &s) = 0;
			virtual void write_boxed(::%[4]s::tl_throwable_ostream &s) = 0;

			virtual bool write_json(std::ostream &s) = 0;

			virtual ~tl_object() = default;
		};

		struct tl_function : public tl_object {
			virtual bool read_write_result(::%[4]s::tl_istream &in, ::%[4]s::tl_ostream &out) = 0;

			virtual ~tl_function() = default;
		};

		struct tl_item {
			uint32_t tag{};
			uint32_t annotations{};
			std::string name;

			bool has_create_object = false;
			bool has_create_function = false;

			std::function<std::unique_ptr<%[3]s::meta::tl_object>()> create_object;
			std::function<std::unique_ptr<%[3]s::meta::tl_function>()> create_function;
		};

		std::optional<%[3]s::meta::tl_item> get_item_by_name(std::string &&s);
		std::optional<%[3]s::meta::tl_item> get_item_by_tag(uint32_t &&tag);

		void set_create_object_by_name(std::string &&s, std::function<std::unique_ptr<%[3]s::meta::tl_object>()> &&factory);
		void set_create_function_by_name(std::string &&s, std::function<std::unique_ptr<%[3]s::meta::tl_function>()> &&factory);

	}
}`,
		filepath.Join(gen.options.RootCPP, CppBasicTLIOStreamsPath(gen)),
		filepath.Join(gen.options.RootCPP, CppBasicTLIOThrowableStreamsPath(gen)),
		gen.options.RootCPPNamespace,
		gen.cppBasictlNamespace()))

	metaDetails.WriteString(fmt.Sprintf("#include \"%s\"\n", filepath.Join(gen.options.RootCPP, CppBasicTLIOStreamsPath(gen))))
	metaDetails.WriteString(fmt.Sprintf(`
#include <map>

#include "%[1]s"

namespace {
	struct tl_items {
		public:
			std::map<std::string, std::shared_ptr<::%[2]s::meta::tl_item>> items;
			std::map<uint32_t, std::shared_ptr<::%[2]s::meta::tl_item>> items_by_tag;
			tl_items();
	};
    
	tl_items items;
    std::function<std::unique_ptr<::%[2]s::meta::tl_object>()> no_object_generator = []() -> std::unique_ptr<::%[2]s::meta::tl_object> {
        throw std::runtime_error("no generation for this type of objects");
    };
    std::function<std::unique_ptr<::%[2]s::meta::tl_function>()> no_function_generator = []() -> std::unique_ptr<::%[2]s::meta::tl_function> {
        throw std::runtime_error("no generation for this type of functions");
    };
}

std::optional<::%[2]s::meta::tl_item> %[2]s::meta::get_item_by_name(std::string &&s) {
	auto item = items.items.find(s);
	if (item != items.items.end()) {
		return *item->second;
	}
	return {};
}

std::optional<::%[2]s::meta::tl_item> %[2]s::meta::get_item_by_tag(std::uint32_t &&tag) {
	auto item = items.items_by_tag.find(tag);
	if (item != items.items_by_tag.end()) {
		return *item->second;
	}
	return {};
}

void %[2]s::meta::set_create_object_by_name(std::string &&s, std::function<std::unique_ptr<::%[2]s::meta::tl_object>()>&& gen) {
	auto item = items.items.find(s);
	if (item != items.items.end()) {
		item->second->has_create_object = true;
		item->second->create_object = gen;
		return;	
	}
	throw std::runtime_error("no item with such name + \"" + s + "\"");
}

void %[2]s::meta::set_create_function_by_name(std::string &&s, std::function<std::unique_ptr<::%[2]s::meta::tl_function>()>&& gen) {
	auto item = items.items.find(s);
	if (item != items.items.end()) {
		item->second->has_create_function = true;
		item->second->create_function = gen;
		return;	
	}
	throw std::runtime_error("no item with such name + \"" + s + "\"");
}

tl_items::tl_items() {`, filepath.Join(gen.options.RootCPP, filepathName), gen.options.RootCPPNamespace))

	for _, wr := range gen.generatedTypesList {
		if wr.tlTag == 0 || !wr.IsTopLevel() {
			continue
		}
		if _, isStruct := wr.trw.(*TypeRWStruct); isStruct {
			metaDetails.WriteString(
				fmt.Sprintf(`
	auto item%[4]d = std::shared_ptr<::%[5]s::meta::tl_item>(new ::%[5]s::meta::tl_item{.tag=%[2]s,.annotations=%[3]s,.name="%[1]s",.create_object=no_object_generator,.create_function=no_function_generator});
	(this->items)["%[1]s"] = item%[4]d;
	(this->items_by_tag)[%[2]s] = item%[4]d;`,
					wr.tlName.String(),
					fmt.Sprintf("0x%08x", wr.tlTag),
					fmt.Sprintf("0x%x", wr.AnnotationsMask()),
					wr.tlTag,
					gen.options.RootCPPNamespace,
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

	make.WriteString("# compile meta data collection\n")
	make.WriteString(fmt.Sprintf(`__build/__meta.o: %[1]s %[2]s __build
	$(CC) $(CFLAGS) -I. -o __build/__meta.o -c %[2]s
`,
		filepathName,
		filepathDetailsName,
	))
	make.WriteString("\n")
	return nil
}

func createFactory(gen *Gen2, createdHpps map[string]bool, make *strings.Builder) error {
	if !gen.options.AddFactoryData {
		return nil
	}
	factory := strings.Builder{}
	factoryDetails := strings.Builder{}

	filepathName := filepath.Join("__factory", "headers"+hppExt)
	filepathNameDetails := filepath.Join("__factory", "details"+cppExt)

	imports := DirectIncludesCPP{ns: map[*TypeRWWrapper]CppIncludeInfo{}}

	factory.WriteString(fmt.Sprintf(`
#pragma once
namespace %[1]s {
    namespace factory {    
		void set_all_factories();
	}
}`, gen.options.RootCPPNamespace))

	factoryDetails.WriteString(fmt.Sprintf(`
void %[1]s::factory::set_all_factories() {
`, gen.options.RootCPPNamespace))

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
	struct %[3]s_%[1]s : public ::%[4]s::meta::%[1]s {
		%[2]s object;

		bool read(::%[5]s::tl_istream &s) override {return object.read(s);}
		bool write(::%[5]s::tl_ostream &s) override {return object.write(s);}

		void read(::%[5]s::tl_throwable_istream &s) override { object.read(s);}
		void write(::%[5]s::tl_throwable_ostream &s) override { object.write(s);}
        
		bool read_boxed(::%[5]s::tl_istream &s) override {return object.read_boxed(s);}
		bool write_boxed(::%[5]s::tl_ostream &s) override {return object.write_boxed(s);}

		void read_boxed(::%[5]s::tl_throwable_istream &s) override { object.read_boxed(s);}
		void write_boxed(::%[5]s::tl_throwable_ostream &s) override { object.write_boxed(s);}

		bool write_json(std::ostream &s) override {return object.write_json(s);}
`,
				implementedInterface,
				myFullTypeNoPrefix,
				myFullTypeWithUnderlines,
				gen.options.RootCPPNamespace,
				gen.cppBasictlNamespace(),
			))
			if strct.ResultType != nil {
				hppTmpInclude2 := DirectIncludesCPP{ns: map[*TypeRWWrapper]CppIncludeInfo{}}
				resultType := strct.ResultType.trw.cppTypeStringInNamespace(wr.wantsBytesVersion && wr.trw.CPPHasBytesVersion(), &hppTmpInclude2)
				resultTypeNoPrefix := strings.TrimPrefix(resultType, "::") // Stupid C++ has sometimes problems with name resolution of definitions
				imports.ns[strct.ResultType] = CppIncludeInfo{componentId: strct.ResultType.typeComponent, namespace: strct.ResultType.groupName}

				factoryDetails.WriteString(fmt.Sprintf(`
		bool read_write_result(::%[2]s::tl_istream &in, ::%[2]s::tl_ostream &out) override {
			%[1]s result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}
`,
					resultTypeNoPrefix,
					gen.cppBasictlNamespace(),
				))
			}
			factoryDetails.WriteString(`
	};`)
			factoryDetails.WriteString(fmt.Sprintf(`
	::%[5]s::meta::set_create_object_by_name("%[1]s", []() -> std::unique_ptr<::%[5]s::meta::tl_object> {
		return std::make_unique<%[2]s_%[3]s>();
	});
`,
				wr.tlName.String(),
				myFullTypeWithUnderlines,
				implementedInterface,
				myFullTypeNoPrefix,
				gen.options.RootCPPNamespace,
			))
			if strct.ResultType != nil {
				factoryDetails.WriteString(fmt.Sprintf(`
	::%[5]s::meta::set_create_function_by_name("%[1]s", []() -> std::unique_ptr<::%[5]s::meta::tl_function> {
		return std::make_unique<%[2]s_%[3]s>();
	});
`,
					wr.tlName.String(),
					myFullTypeWithUnderlines,
					implementedInterface,
					myFullTypeNoPrefix,
					gen.options.RootCPPNamespace,
				))
			}
		}
	}

	factoryDetails.WriteString(`
}
`)
	suffix := factoryDetails.String()
	factoryDetails.Reset()
	factoryDetails.WriteString(fmt.Sprintf("#include \"%s\"\n", filepath.Join(gen.options.RootCPP, filepath.Join("__meta", "headers"+hppExt))))
	factoryDetails.WriteString(fmt.Sprintf("#include \"%s\"\n\n", filepath.Join(gen.options.RootCPP, filepathName)))

	factoryFileDependencies := strings.Builder{}

	for _, headerFile := range imports.sortedIncludes(gen.componentsOrder, func(wrapper *TypeRWWrapper) string { return wrapper.fileName }) {
		if !createdHpps[headerFile] {
			continue
		}
		factoryDetails.WriteString(fmt.Sprintf("#include \"%s%s\"\n", filepath.Join(gen.options.RootCPP, headerFile), hppExt))
		factoryFileDependencies.WriteString(" " + headerFile + hppExt)
	}
	factoryDetails.WriteString(suffix)

	if err := gen.addCodeFile(filepathName, gen.copyrightText+factory.String()); err != nil {
		return err
	}
	if err := gen.addCodeFile(filepathNameDetails, gen.copyrightText+factoryDetails.String()); err != nil {
		return err
	}

	make.WriteString("# compile objects factories\n")
	make.WriteString(fmt.Sprintf(`__build/__factory.o: %[1]s %[2]s%[3]s __build
	$(CC) $(CFLAGS) -I. -o __build/__factory.o -c %[2]s
`,
		filepathName,
		filepathNameDetails,
		factoryFileDependencies.String(),
	))
	make.WriteString("\n")

	return nil
}

func appendCompilerOptions(make *strings.Builder) {
	make.WriteString(`# compiler options
CC = g++
CFLAGS = -std=c++20 -O3 -Wno-noexcept-type -g -Wall -Wextra -Werror -Wunused-parameter

`)
}

func cppRunLocalLinter(code, filename string) (string, string) {
	if strings.HasSuffix(filename, hppExt) ||
		strings.HasSuffix(filename, cppExt) {
		code = strings.ReplaceAll(code, "\t", "  ")
	}
	return code, filename
}

func (gen *Gen2) cppBasictlNamespace() string {
	return fmt.Sprintf("%[1]s::%[2]s", gen.options.RootCPPNamespace, CppBasictlPackage(gen))
}
