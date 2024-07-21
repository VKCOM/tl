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
	"path/filepath"
	"sort"
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

	for header, typeDefs := range hpps {
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
			hpp.WriteString(fmt.Sprintf("#include \"%s\"\n", getCppDiff(filepathName, basicTLFilepathName)))
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
		hppDet.WriteString(fmt.Sprintf("#include \"%s\"\n", getCppDiff(filepathName, basicTLFilepathName)))
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
			if !createdDetailsHpps[spec.hppDetailsFileName] {
				continue
			}
			cppDetInc.ns[spec] = CppIncludeInfo{componentId: spec.typeComponent, namespace: spec.groupName}
		}
		for _, n := range cppDetInc.sortedIncludes(gen.componentsOrder, func(wrapper *TypeRWWrapper) string { return wrapper.hppDetailsFileName }) {
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
		buildFilePath := "__build/" + namespaceDetails + ".o"

		var cppMake1UsedFiles strings.Builder
		var cppMake1Namespace strings.Builder

		for _, n := range namespaceDeps {
			cppMake1Namespace.WriteString(fmt.Sprintf("#include \"%s\"\n", getCppDiff(namespaceFilePath, n+cppExt)))
			cppMake1UsedFiles.WriteString(fmt.Sprintf("%s", getCppDiff(MakefilePath, n+cppExt)))

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

	_ = createMeta(gen)
	_ = createFactory(gen)

	cppMake.WriteString(`
CC = g++
CFLAGS = -std=c++17 -O3 -Wno-noexcept-type -g -Wall -Wextra -Werror=return-type -Wno-unused-parameter
`)
	cppMake.WriteString(fmt.Sprintf("all: "))
	cppMake.WriteString(fmt.Sprintf("main.o "))
	cppMake.WriteString(fmt.Sprintf("%s\n", cppMakeO.String()))
	cppMake.WriteString(fmt.Sprintf("\t$(CC) $(CFLAGS) -o all "))
	cppMake.WriteString(fmt.Sprintf("main.o "))
	cppMake.WriteString(fmt.Sprintf("%s\n", cppMakeO.String()))
	cppMake.WriteString(`
main.o: main.cpp
	$(CC) $(CFLAGS) -c main.cpp
`)
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

func (gen *Gen2) decideCppCodeDestinations(allTypes []*TypeRWWrapper) {
	const IndependentTypes = "__independent_types"
	const NoNamespaceGroup = ""
	const CommonGroup = "__common"

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
						to.groupName = utils.SetToSlice(groups)[0]
						if to.groupName != CommonGroup && to.groupName != IndependentTypes {
							to.cppDetailsFileName = to.groupName + "_" + to.cppDetailsFileName
						}
						to.hppDetailsFileName = to.cppDetailsFileName
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
		if typeGroup == "" {
			typeGroup = CommonGroup
		}
		if strct, isStruct := t.trw.(*TypeRWStruct); isStruct && strct.ResultType != nil {
			t.fileName = filepath.Join(typeGroup, "functions", t.fileName)
		} else {
			t.fileName = filepath.Join(typeGroup, "types", t.fileName)
		}
		t.hppDetailsFileName = filepath.Join(t.groupName, "details", "headers", t.hppDetailsFileName)
		t.cppDetailsFileName = filepath.Join(t.groupName, "details", "namespace_details")
	}

}

func getCppDiff(base string, target string) string {
	dir1, _ := filepath.Split(base)
	dir2, file := filepath.Split(target)
	diff, _ := filepath.Rel(dir1, dir2)
	return filepath.Join(diff, file)
}

func createMeta(gen *Gen2) error {
	meta := strings.Builder{}
	filepathName := filepath.Join("__meta", "meta"+hppExt)

	meta.WriteString("#pragma once\n")
	meta.WriteString(fmt.Sprintf("#include \"%s\"\n", getCppDiff(filepathName, "a_tlgen_helpers_code.hpp")))
	meta.WriteString(fmt.Sprintf(`
#include <functional>
#include <map>

namespace tl2 {
namespace meta {
    struct tl_object {
        std::function<bool(::basictl::tl_istream & )> read;
        std::function<bool(::basictl::tl_ostream & )> write;

        std::function<bool(::basictl::tl_istream & )> read_boxed;
        std::function<bool(::basictl::tl_ostream & )> write_boxed;
    };

    struct tl_item {
        uint32_t tag{};
        uint32_t annotations{};
        std::string name;

        std::function<tl2::meta::tl_object()> create_object;
    };

    namespace {
        std::map<std::string, tl2::meta::tl_item> __objects;
		std::function<tl_object()> missing_generator = []() -> tl_object {
            throw std::runtime_error("no generator initialized");
        };
    }

    tl_item get_tl_item_by_name(std::string&& name) {
        if (__objects.count(name)) {
            return __objects[name];
        }
        throw std::runtime_error("no such tl (\""  + name + "\") item in system");
    }

    void set_create_object_by_name(std::string&& name, std::function<tl_object()>&& generator) {
        if (__objects.count(name)) {
            __objects[name].create_object = generator;
            return;
        }
        throw std::runtime_error("no such tl item in system");
    }

    void init_tl_objects() {`))

	for _, wr := range gen.generatedTypesList {
		if wr.tlTag == 0 || !wr.IsTopLevel() {
			continue
		}
		if strct, isStruct := wr.trw.(*TypeRWStruct); isStruct && len(wr.NatParams) == 0 {
			if strct.ResultType == nil {

				meta.WriteString(
					fmt.Sprintf(`
		__objects["%[1]s"] = tl2::meta::tl_item{.tag=%s,.annotations=%s,.name="%[1]s",.create_object=missing_generator};`,
						wr.tlName.String(),
						fmt.Sprintf("0x%08x", wr.tlTag),
						fmt.Sprintf("0x%x", wr.AnnotationsMask()),
					),
				)
			}
		}
	}

	meta.WriteString(fmt.Sprintf(`
	}
};
};
`))
	if err := gen.addCodeFile(filepathName, gen.copyrightText+meta.String()); err != nil {
		return err
	}
	return nil
}

func createFactory(gen *Gen2) error {
	factory := strings.Builder{}
	filepathName := filepath.Join("__factory", "factory"+hppExt)

	imports := DirectIncludesCPP{ns: map[*TypeRWWrapper]CppIncludeInfo{}}

	factory.WriteString(fmt.Sprintf(`

namespace tl2 {
namespace factory {
    void init_tl_create_objects() {`))

	for _, wr := range gen.generatedTypesList {
		if wr.tlTag == 0 || !wr.IsTopLevel() {
			continue
		}
		if strct, isStruct := wr.trw.(*TypeRWStruct); isStruct && len(wr.NatParams) == 0 {
			if strct.isTypeDef() {
				continue
			}
			if strct.ResultType == nil {
				hppTmpInclude := DirectIncludesCPP{ns: map[*TypeRWWrapper]CppIncludeInfo{}}
				myFullType := wr.trw.cppTypeStringInNamespace(wr.wantsBytesVersion && wr.trw.CPPHasBytesVersion(), &hppTmpInclude)
				myFullTypeNoPrefix := strings.TrimPrefix(myFullType, "::") // Stupid C++ has sometimes problems with name resolution of definitions

				factory.WriteString(
					fmt.Sprintf(`
		tl2::meta::set_create_object_by_name("%[1]s",[]() -> tl2::meta::tl_object {
        auto obj = std::make_shared<%[2]s>();
        return tl2::meta::tl_object{
                .read=[obj](auto &in) -> bool { return obj->read(in); },
                .write=[obj](auto &out) -> bool { return obj->write(out); },
                .read_boxed=[obj](auto &in) -> bool { return obj->read_boxed(in); },
                .write_boxed=[obj](auto &out) -> bool { return obj->write_boxed(out); },
        };
    });`,
						wr.tlName.String(),
						myFullTypeNoPrefix,
					),
				)
				imports.ns[wr] = CppIncludeInfo{componentId: wr.typeComponent, namespace: wr.groupName}
			}
		}
	}

	factory.WriteString(fmt.Sprintf(`
	}
};
};
`))
	suffix := factory.String()
	factory.Reset()

	factory.WriteString("#pragma once\n")
	factory.WriteString(fmt.Sprintf("#include \"%s\"\n\n", getCppDiff(filepathName, filepath.Join("__meta", "meta"+hppExt))))
	for _, headerFile := range imports.sortedIncludes(gen.componentsOrder, func(wrapper *TypeRWWrapper) string { return wrapper.fileName }) {
		factory.WriteString(fmt.Sprintf("#include \"%s%s\"\n", getCppDiff(filepathName, headerFile), hppExt))
	}
	factory.WriteString(suffix)

	if err := gen.addCodeFile(filepathName, gen.copyrightText+factory.String()); err != nil {
		return err
	}
	return nil
}
