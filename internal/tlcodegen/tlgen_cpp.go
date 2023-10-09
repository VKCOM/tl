// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlcodegen

/*
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

	cppAllInc := &DirectIncludesCPP{ns: map[string]struct{}{}}
	var hpp strings.Builder
	var hppDet strings.Builder
	var cppDet strings.Builder
	typesCounter := 0
	typesCounterBytes := 0
	internalFiles := map[InsFile][]*TypeRWWrapper{}
	for _, typeRw := range gen.generatedTypesList {
		ff := InsFile{ins: typeRw.ins, fileName: typeRw.fileName}
		internalFiles[ff] = append(internalFiles[ff], typeRw)
	}
	for ff, types := range internalFiles {
		hppInc := &DirectIncludesCPP{ns: map[string]struct{}{}}
		hppIncFwd := &DirectIncludesCPP{ns: map[string]struct{}{}}
		hppDetInc := &DirectIncludesCPP{ns: map[string]struct{}{}}
		cppDetInc := &DirectIncludesCPP{ns: map[string]struct{}{}}
		multipleDefinitions := map[string]struct{}{}
		for _, typeRw := range types {
			// log.Printf("type: %s\n", typeRw.tlName.String())
			// log.Printf("      %s\n", typeRw.resolvedType.String())
			typesCounter++
			var hppDefinition strings.Builder
			typeRw.trw.CPPGenerateCode(&hppDefinition, hppInc, hppIncFwd, &hppDet, hppDetInc, &cppDet, cppDetInc, false, false)
			def := hppDefinition.String()
			if _, ok := multipleDefinitions[def]; !ok {
				multipleDefinitions[def] = struct{}{}
				hpp.WriteString(def)
			}
			if typeRw.wantsBytesVersion && typeRw.trw.CPPHasBytesVersion() {
				hppDefinition.Reset()
				typesCounterBytes++
				typeRw.trw.CPPGenerateCode(&hppDefinition, hppInc, hppIncFwd, &hppDet, hppDetInc, &cppDet, cppDetInc, true, false)
				def = hppDefinition.String()
				if _, ok := multipleDefinitions[def]; !ok {
					multipleDefinitions[def] = struct{}{}
					hpp.WriteString(def)
				}
			}
		}
		if hpp.Len() == 0 && hppDet.Len() == 0 && cppDet.Len() == 0 {
			continue
		}
		cppAllInc.ns[ff.fileName] = struct{}{}
		hppStr := hpp.String()
		hppDetStr := hppDet.String()
		cppDetStr := cppDet.String()
		hpp.Reset()
		hppDet.Reset()
		cppDet.Reset()
		hpp.WriteString("#pragma once\n\n")
		hppDet.WriteString("#pragma once\n\n")
		hpp.WriteString(fmt.Sprintf("#include \"%s\"\n", basicTLFilepathName))
		for _, n := range hppInc.sortedNames() {
			hpp.WriteString(fmt.Sprintf("#include \"%s%s\"\n", n, hppExt))
		}
		hpp.WriteString("\n\n")
		hppDet.WriteString(fmt.Sprintf("#include \"../%s%s\"\n", ff.fileName, hppExt))
		hppDet.WriteString(fmt.Sprintf("#include \"../%s\"\n", basicTLFilepathName))
		hpp.WriteString(hppStr)
		// for _, n := range hppIncFwd.sortedNames() {
		//	hpp.WriteString(fmt.Sprintf("#include \"%s%s\"\n", n, hppExt))
		// }
		for _, n := range hppDetInc.sortedNames() {
			hppDet.WriteString(fmt.Sprintf("#include \"../%s%s\"\n", n, hppExt))
		}
		cppDet.WriteString(fmt.Sprintf("#include \"%s_details%s\"\n", ff.fileName, hppExt))
		for _, n := range cppDetInc.sortedNames() {
			cppDet.WriteString(fmt.Sprintf("#include \"%s_details%s\"\n", n, hppExt))
		}
		filepathName := ff.fileName + hppExt
		if err := gen.addCodeFile(filepathName, gen.copyrightText+hpp.String()); err != nil {
			return err
		}
		hpp.Reset()
		filepathName = filepath.Join("details", ff.fileName+"_details"+hppExt)
		if err := gen.addCodeFile(filepathName, gen.copyrightText+hppDet.String()+hppDetStr); err != nil {
			return err
		}
		hppDet.Reset()
		filepathName = filepath.Join("details", ff.fileName+"_details"+cppExt)
		if err := gen.addCodeFile(filepathName, gen.copyrightText+cppDet.String()+cppDetStr); err != nil {
			return err
		}
		cppDet.Reset()
	}
	var cppAll strings.Builder
	var cppMake strings.Builder
	var cppMakeO strings.Builder
	var cppMake1 strings.Builder
	for _, n := range cppAllInc.sortedNames() {
		cppAll.WriteString(fmt.Sprintf("#include \"details/%s%s\"\n", n+"_details", cppExt))
		cppMake1.WriteString(fmt.Sprintf("%s.o: details/%s%s details/%s%s\n", n+"_details", n+"_details", cppExt, n+"_details", hppExt))
		cppMake1.WriteString(fmt.Sprintf("\t$(CC) $(CFLAGS) -c details/%s%s\n", n+"_details", cppExt))
		cppMakeO.WriteString(fmt.Sprintf("%s.o ", n+"_details"))
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
	// if gen.options.Verbose {
	//	log.Printf("generation of serialization code finished, %d constructors processed, %d types generated", len(gen.allConstructors), typesCounter)
	//	if len(generateByteVersions) != 0 {
	//		log.Printf("    also generated byte-optimized versions of %d types by the following filter: %s", typesCounterBytes, strings.Join(generateByteVersions, ", "))
	//	}
	// }
	// if gen.options.Verbose {
	//	log.Printf("generating RPC code...")
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
*/
