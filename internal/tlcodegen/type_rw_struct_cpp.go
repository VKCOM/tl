// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlcodegen

import (
	"fmt"
	"strings"
)

func (trw *TypeRWStruct) CPPFillRecursiveChildren(visitedNodes map[*TypeRWWrapper]bool) {
	for _, f := range trw.Fields {
		if !f.recursive {
			f.t.CPPFillRecursiveChildren(visitedNodes)
		}
	}
}

func (trw *TypeRWStruct) cppTypeStringInNamespace(bytesVersion bool, hppInc *DirectIncludesCPP) string {
	if trw.isUnwrapType() { // TODO - when replacing typedefs, we must make name resolution
		return trw.Fields[0].t.CPPTypeStringInNamespace(bytesVersion, hppInc)
	}
	_, _, args := trw.wr.cppTypeStringInNamespace(bytesVersion, hppInc, false, HalfResolvedArgument{})
	return trw.wr.cppNamespaceQualifier() + trw.wr.cppLocalName + args
}

func (trw *TypeRWStruct) cppTypeStringInNamespaceHalfResolved(bytesVersion bool, hppInc *DirectIncludesCPP, halfResolved HalfResolvedArgument) string {
	if trw.isUnwrapType() {
		halfResolvedUnwrapped := trw.wr.replaceUnwrapHalfResolved(halfResolved, trw.Fields[0].halfResolved)
		return trw.Fields[0].t.CPPTypeStringInNamespaceHalfResolved(bytesVersion, hppInc, halfResolvedUnwrapped)
	}
	_, _, args := trw.wr.cppTypeStringInNamespace(bytesVersion, hppInc, true, halfResolved)
	return trw.wr.cppNamespaceQualifier() + trw.wr.cppLocalName + args
}

func (trw *TypeRWStruct) cppTypeStringInNamespaceHalfResolved2(bytesVersion bool, typeReduction EvaluatedType) string {
	if trw.isUnwrapType() {
		eval := trw.wr.gen.typesInfo.FieldTypeReduction(typeReduction.Type, 0)
		return trw.Fields[0].t.CPPTypeStringInNamespaceHalfResolved2(bytesVersion, eval)
	}
	return trw.wr.cppNamespaceQualifier() + trw.wr.cppLocalName + trw.wr.cppTypeArguments(bytesVersion, typeReduction.Type)
}

func (trw *TypeRWStruct) cppDefaultInitializer(halfResolved HalfResolvedArgument, halfResolve bool) string {
	if trw.isUnwrapType() {
		return trw.Fields[0].t.CPPDefaultInitializer(trw.Fields[0].halfResolved, halfResolve)
	}
	return "{}"
}

func (trw *TypeRWStruct) CPPHasBytesVersion() bool {
	return false // TODO
}

func (trw *TypeRWStruct) CPPTypeResettingCode(bytesVersion bool, val string) string {
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
	if trw.isUnwrapType() {
		return trw.Fields[0].t.trw.CPPTypeResettingCode(bytesVersion, val)
	}
	return fmt.Sprintf("\t::%s::%sReset(%s);", trw.wr.gen.DetailsCPPNamespace, goGlobalName, val)
}

func (trw *TypeRWStruct) CPPTypeWritingCode(bytesVersion bool, val string, bare bool, natArgs []string, last bool) string {
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
	if trw.isUnwrapType() {
		prefix := ""
		if !bare {
			prefix = fmt.Sprintf("\ts.nat_write(0x%x);\n", trw.wr.tlTag)
		}
		return prefix + trw.Fields[0].t.trw.CPPTypeWritingCode(bytesVersion, val, trw.Fields[0].Bare(), trw.replaceUnwrapArgs(natArgs), last)
	}
	return fmt.Sprintf("\t::%s::%sWrite%s(s, %s%s);", trw.wr.gen.DetailsCPPNamespace, goGlobalName, addBare(bare), val, joinWithCommas(natArgs))
}

func (trw *TypeRWStruct) CPPTypeReadingCode(bytesVersion bool, val string, bare bool, natArgs []string, last bool) string {
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
	if trw.isUnwrapType() {
		prefix := ""
		if !bare {
			prefix = fmt.Sprintf("\ts.nat_read_exact_tag(0x%x);\n", trw.wr.tlTag)
		}
		s := prefix + trw.Fields[0].t.trw.CPPTypeReadingCode(bytesVersion, val, trw.Fields[0].Bare(), trw.replaceUnwrapArgs(natArgs), last)
		return s
	}
	s := fmt.Sprintf("\t::%s::%sRead%s(s, %s%s);", trw.wr.gen.DetailsCPPNamespace, goGlobalName, addBare(bare), val, joinWithCommas(natArgs))
	return s
}

func (trw *TypeRWStruct) CPPGenerateCode(hpp *strings.Builder, hppInc *DirectIncludesCPP, hppIncFwd *DirectIncludesCPP, hppDet *strings.Builder, hppDetInc *DirectIncludesCPP, cppDet *strings.Builder, cppDetInc *DirectIncludesCPP, bytesVersion bool, forwardDeclaration bool) {
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
	//if trw.wr.unionParent != nil && trw.wr.unionParentIsEnum {
	//	return
	//}
	_, myArgsDecl := trw.wr.fullyResolvedClassCppNameArgs()
	myFullType := trw.cppTypeStringInNamespace(bytesVersion, hppDetInc)
	// fmt.Printf("Ts: %s %s\n", myFullType, strings.Join(myArgsDecl, ", "))
	// fmt.Printf("    %s\n", trw.wr.cppLocalName)
	myFullTypeNoPrefix := strings.TrimPrefix(myFullType, "::") // Stupid C++ has sometimes problems with name resolution of definitions

	anyRecursive := false
	typeNamespace := trw.wr.gen.RootCPPNamespaceElements
	if trw.wr.tlName.Namespace != "" {
		typeNamespace = append(typeNamespace, trw.wr.tlName.Namespace)
	}

	if !forwardDeclaration {
		for _, typeDep := range trw.AllTypeDependencies() {
			if typeDep.typeComponent == trw.wr.typeComponent {
				typeDep.trw.CPPGenerateCode(hpp, nil, nil, nil, hppDetInc, nil, cppDetInc, bytesVersion, true)
			}
		}
	}
	cppStartNamespace(hpp, typeNamespace)
	// hpp.WriteString("// " + goLocalName + "\n") - uncommenting will lead to multiple definition error
	if len(myArgsDecl) != 0 {
		hpp.WriteString("template<" + strings.Join(myArgsDecl, ", ") + ">\n")
	}
	if forwardDeclaration { // TODO - does not work for typedef
		hpp.WriteString("struct " + trw.wr.cppLocalName + ";")
		cppFinishNamespace(hpp, typeNamespace)
		return
	}
	if trw.isTypeDef() {
		field := trw.Fields[0]

		if !field.t.origTL[0].Builtin && len(trw.wr.arguments) != 0 {
			// TODO TODO TODO temporary solution
			ti := trw.wr.gen.typesInfo
			tlName := trw.wr.tlName

			_, isType := ti.Types[tlName]
			typeReduction := TypeReduction{IsType: isType}
			if isType {
				typeReduction.Type = ti.Types[tlName]
			} else {
				typeReduction.Constructor = ti.Constructors[tlName]
			}
			for i, arg := range typeReduction.ReferenceType().TypeArguments {
				evalArg := EvaluatedType{}
				if arg.IsNat {
					evalArg.Index = 1
					evalArg.Variable = arg.FieldName
					if trw.wr.arguments[i].isArith {
						// set true only here
						evalArg.VariableActsAsConstant = true
					}
				} else {
					evalArg.Index = 3
					evalArg.TypeVariable = arg.FieldName
				}
				typeReduction.Arguments = append(typeReduction.Arguments, evalArg)
			}

			typeRed := ti.FieldTypeReduction(&typeReduction, 0)
			for _, typeRw := range trw.Fields[0].t.ActualTypeDependencies(typeRed) {
				hppInc.ns[typeRw.fileName] = CppIncludeInfo{componentId: typeRw.typeComponent}
			}
			hpp.WriteString(fmt.Sprintf("using %s = %s;", trw.wr.cppLocalName, field.t.CPPTypeStringInNamespaceHalfResolved2(bytesVersion, typeRed)))
		} else {
			fieldFullType := field.t.CPPTypeStringInNamespaceHalfResolved(bytesVersion, hppInc, field.halfResolved)
			hpp.WriteString(fmt.Sprintf("using %s = %s;", trw.wr.cppLocalName, fieldFullType))
		}
	} else {
		hpp.WriteString("struct " + trw.wr.cppLocalName + " {\n")
		for _, field := range trw.Fields {
			hppIncByField := DirectIncludesCPP{ns: map[string]CppIncludeInfo{}}
			fieldFullType := field.t.CPPTypeStringInNamespaceHalfResolved(bytesVersion, &hppIncByField, field.halfResolved)
			fieldsMaskComment := ""
			//if field.fieldMask != nil {
			//	fieldsMaskComment = fmt.Sprintf(" // Conditional: %s.%d", formatNatArgCPP(trw.Fields, *field.fieldMask), field.BitNumber)
			//}
			if field.recursive {
				// TODO make better
				for includeType, includeInfo := range hppIncByField.ns {
					if includeInfo.componentId == trw.wr.typeComponent {
						delete(hppIncByField.ns, includeType)
					}
				}
				anyRecursive = true // requires destructor in cpp file
				hpp.WriteString(fmt.Sprintf("\tstd::shared_ptr<%s> %s{};%s\n", fieldFullType, field.cppName, fieldsMaskComment))
			} else {
				hpp.WriteString(fmt.Sprintf("\t%s %s%s;%s\n", fieldFullType, field.cppName, field.t.CPPDefaultInitializer(field.halfResolved, true), fieldsMaskComment))
			}
			for includeType, includeInfo := range hppIncByField.ns {
				hppInc.ns[includeType] = includeInfo
			}
			//hpp.WriteString(fmt.Sprintf("\t// DebugString: %s\n", field.resolvedType.DebugString()))
		}
		if anyRecursive { // && len(trw.cppArgs) != 0
			hpp.WriteString(fmt.Sprintf("\n\t~%s() {}\n", trw.wr.cppLocalName)) // TODO - move destructor to cpp
			// cppDet.WriteString(fmt.Sprintf("%s%s::~%s() {}\n", trw.wr.cppNamespaceQualifier, goLocalName, goLocalName))
		}
		if trw.wr.tlTag != 0 { // anonymous square brackets citizens or other exotic type
			hpp.WriteString(fmt.Sprintf(`
	::basictl::string_view tl_name() const { return "%s"; }
	uint32_t tl_tag() const { return 0x%08x; }
`, trw.wr.tlName, trw.wr.tlTag))
		}
		if len(myArgsDecl) == 0 {
			// cppStartNamespace(cppDet, trw.wr.gen.RootCPPNamespaceElements)
			hpp.WriteString(fmt.Sprintf(`
	void read(::basictl::tl_istream & s%[1]s);
	void write(::basictl::tl_ostream & s%[1]s)const;
`,
				formatNatArgsDeclCPP(trw.wr.NatParams),
				trw.CPPTypeResettingCode(bytesVersion, "*this"),
				trw.CPPTypeReadingCode(bytesVersion, "*this", true, formatNatArgsAddNat(trw.wr.NatParams), true),
				trw.CPPTypeWritingCode(bytesVersion, "*this", true, formatNatArgsAddNat(trw.wr.NatParams), true)))
			cppDet.WriteString(fmt.Sprintf(`
void %[5]s::read(::basictl::tl_istream & s%[1]s) {
%[3]s
}

void %[5]s::write(::basictl::tl_ostream & s%[1]s)const {
%[4]s
}
`,
				formatNatArgsDeclCPP(trw.wr.NatParams),
				trw.CPPTypeResettingCode(bytesVersion, "*this"),
				trw.CPPTypeReadingCode(bytesVersion, "*this", true, formatNatArgsAddNat(trw.wr.NatParams), true),
				trw.CPPTypeWritingCode(bytesVersion, "*this", true, formatNatArgsAddNat(trw.wr.NatParams), true),
				myFullTypeNoPrefix))
			if trw.wr.tlTag != 0 { // anonymous square brackets citizens or other exotic type
				hpp.WriteString(fmt.Sprintf(`
	void read_boxed(::basictl::tl_istream & s%[1]s);
	void write_boxed(::basictl::tl_ostream & s%[1]s)const;
`,
					formatNatArgsDeclCPP(trw.wr.NatParams),
					trw.CPPTypeResettingCode(bytesVersion, "*this"),
					trw.CPPTypeReadingCode(bytesVersion, "*this", false, formatNatArgsAddNat(trw.wr.NatParams), true),
					trw.CPPTypeWritingCode(bytesVersion, "*this", false, formatNatArgsAddNat(trw.wr.NatParams), true)))
				cppDet.WriteString(fmt.Sprintf(`
void %[5]s::read_boxed(::basictl::tl_istream & s%[1]s) {
%[3]s
}

void %[5]s::write_boxed(::basictl::tl_ostream & s%[1]s)const {
%[4]s
}
`,
					formatNatArgsDeclCPP(trw.wr.NatParams),
					trw.CPPTypeResettingCode(bytesVersion, "*this"),
					trw.CPPTypeReadingCode(bytesVersion, "*this", false, formatNatArgsAddNat(trw.wr.NatParams), true),
					trw.CPPTypeWritingCode(bytesVersion, "*this", false, formatNatArgsAddNat(trw.wr.NatParams), true),
					myFullTypeNoPrefix))
			}
			// cppFinishNamespace(cppDet, trw.wr.gen.RootCPPNamespaceElements)
		}
		hpp.WriteString("};\n")
	}
	cppFinishNamespace(hpp, typeNamespace)

	cppDet.WriteString(fmt.Sprintf(`
void %[7]s::%[1]sReset(%[2]s& item) {
%[4]s}

void %[7]s::%[1]sRead(::basictl::tl_istream & s, %[2]s& item%[3]s) {
%[5]s}

void %[7]s::%[1]sWrite(::basictl::tl_ostream & s, const %[2]s& item%[3]s) {
%[6]s}
`,
		goGlobalName,
		myFullType,
		formatNatArgsDeclCPP(trw.wr.NatParams),
		trw.CPPResetFields(bytesVersion),
		trw.CPPReadFields(bytesVersion, hppDetInc, cppDetInc),
		trw.CPPWriteFields(bytesVersion),
		trw.wr.gen.DetailsCPPNamespace,
	))

	cppStartNamespace(hppDet, trw.wr.gen.DetailsCPPNamespaceElements)
	hppDet.WriteString(fmt.Sprintf(`
void %[1]sReset(%[2]s& item);
void %[1]sRead(::basictl::tl_istream & s, %[2]s& item%[3]s);
void %[1]sWrite(::basictl::tl_ostream & s, const %[2]s& item%[3]s);
`, goGlobalName, myFullType, formatNatArgsDeclCPP(trw.wr.NatParams)))

	if trw.wr.tlTag != 0 { // anonymous square brackets citizens or other exotic type
		hppDet.WriteString(fmt.Sprintf(`void %[1]sReadBoxed(::basictl::tl_istream & s, %[2]s& item%[3]s);
void %[1]sWriteBoxed(::basictl::tl_ostream & s, const %[2]s& item%[3]s);
`,
			goGlobalName,
			myFullType,
			formatNatArgsDeclCPP(trw.wr.NatParams)))

		s := fmt.Sprintf(`
void %[7]s::%[1]sReadBoxed(::basictl::tl_istream & s, %[2]s& item%[3]s) {
	s.nat_read_exact_tag(0x%08[9]x);
	%[7]s::%[1]sRead(s, item%[8]s);
}

void %[7]s::%[1]sWriteBoxed(::basictl::tl_ostream & s, const %[2]s& item%[3]s) {
	s.nat_write(0x%08[9]x);
	%[7]s::%[1]sWrite(s, item%[8]s);
}
`,
			goGlobalName,
			myFullType,
			formatNatArgsDeclCPP(trw.wr.NatParams),
			trw.CPPResetFields(bytesVersion),
			trw.CPPReadFields(bytesVersion, hppDetInc, cppDetInc),
			trw.CPPWriteFields(bytesVersion),
			trw.wr.gen.DetailsCPPNamespace,
			formatNatArgsCallCPP(trw.wr.NatParams),
			trw.wr.tlTag,
		)
		cppDet.WriteString(s)

	}
	cppFinishNamespace(hppDet, trw.wr.gen.DetailsCPPNamespaceElements)
}

func (trw *TypeRWStruct) CPPResetFields(bytesVersion bool) string {
	var s strings.Builder
	if trw.isTypeDef() {
		field := trw.Fields[0]
		s.WriteString(
			field.t.trw.CPPTypeResettingCode(bytesVersion, "item") + "\n")
		return s.String()
	}
	for _, field := range trw.Fields {
		if field.recursive {
			s.WriteString(fmt.Sprintf(`
	if (item.%[1]s) {
		%[2]s
	}
`, field.cppName, field.t.trw.CPPTypeResettingCode(bytesVersion, fmt.Sprintf("(*item.%s)", field.cppName))))
		} else {
			s.WriteString(field.t.trw.CPPTypeResettingCode(bytesVersion, fmt.Sprintf("item.%s", field.cppName)) + "\n")
		}
	}
	return s.String()
}

func (trw *TypeRWStruct) CPPWriteFields(bytesVersion bool) string {
	var s strings.Builder
	if trw.isTypeDef() {
		field := trw.Fields[0]
		s.WriteString(
			field.t.trw.CPPTypeWritingCode(bytesVersion, "item",
				field.Bare(), formatNatArgsCPP(trw.Fields, field.natArgs), false) + "\n")
		return s.String()
	}
	for _, field := range trw.Fields {
		indent := 0
		if field.fieldMask != nil {
			s.WriteString(fmt.Sprintf("\tif ((%s & (1<<%d)) != 0) {\n\t", formatNatArgCPP(trw.Fields, *field.fieldMask), field.BitNumber))
			indent++
		}
		s.WriteString(strings.Repeat("\t", indent))
		s.WriteString(
			field.t.trw.CPPTypeWritingCode(bytesVersion, addAsterisk(field.recursive, fmt.Sprintf("item.%s", field.cppName)),
				field.Bare(), formatNatArgsCPP(trw.Fields, field.natArgs), false) + "\n")
		if field.fieldMask != nil {
			s.WriteString("\t}\n")
		}
	}
	return s.String()
}

func (trw *TypeRWStruct) CPPReadFields(bytesVersion bool, hppDetInc *DirectIncludesCPP, cppDetInc *DirectIncludesCPP) string {
	var s strings.Builder
	if trw.isTypeDef() {
		field := trw.Fields[0]
		_ = field.t.CPPTypeStringInNamespace(bytesVersion, cppDetInc) // only fill includes
		s.WriteString(
			field.t.trw.CPPTypeReadingCode(bytesVersion, "item",
				field.Bare(), formatNatArgsCPP(trw.Fields, field.natArgs),
				false) + "\n")

		return s.String()
	}
	for _, field := range trw.Fields {
		indent := 0
		if field.fieldMask != nil {
			s.WriteString(fmt.Sprintf("\tif ((%s & (1<<%d)) != 0) {\n", formatNatArgCPP(trw.Fields, *field.fieldMask), field.BitNumber))
			indent++
		}
		if field.recursive {
			s.WriteString(strings.Repeat("\t", indent))
			s.WriteString(fmt.Sprintf("\t"+`if (!item.%[2]s) { item.%[2]s = std::make_shared<%[1]s>(); }
`, field.t.CPPTypeStringInNamespace(bytesVersion, cppDetInc), field.cppName))
		}
		_ = field.t.CPPTypeStringInNamespace(bytesVersion, cppDetInc) // only fill includes
		_ = field.t.CPPTypeStringInNamespace(bytesVersion, hppDetInc) // only fill includes
		s.WriteString(strings.Repeat("\t", indent))
		s.WriteString(
			field.t.trw.CPPTypeReadingCode(bytesVersion, addAsterisk(field.recursive, "item."+field.cppName),
				field.Bare(), formatNatArgsCPP(trw.Fields, field.natArgs),
				false) + "\n")
		if field.fieldMask != nil {
			// TODO - in case of recursive field, check for nil ptr
			s.WriteString(fmt.Sprintf("\t} else {\n\t\t%s\n\t}\n", field.t.trw.CPPTypeResettingCode(bytesVersion, addAsterisk(field.recursive, fmt.Sprintf("item.%s", field.cppName)))))
		}
	}
	return s.String()
}
