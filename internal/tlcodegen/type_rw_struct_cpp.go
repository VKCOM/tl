// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlcodegen

/*
import (
	"fmt"
	"strconv"
	"strings"
)

func (trw *TypeRWStruct) CPPFillRecursiveChildren(visitedNodes map[*TypeRWWrapper]bool) {
	for _, f := range trw.Fields {
		if !f.recursive {
			f.t.CPPFillRecursiveChildren(visitedNodes)
		}
	}
}

func (trw *TypeRWStruct) cppTypeStringInNamespace(bytesVersion bool, hppInc *DirectIncludesCPP, resolvedType ResolvedType, halfResolve bool) string {
	//if trw.isUnwrapType() {
	//	TODO - when replacing typedefs, we must make name resolution
	//	return trw.Fields[0].t.CPPTypeStringInNamespace(bytesVersion, hppInc, trw.Fields[0].resolvedType, halfResolve)
	//}
	hppInc.ns[trw.wr.fileName] = struct{}{}
	bNameSuffix := strings.Builder{}
	bNameSuffix.WriteString(trw.wr.cppNamespaceQualifier)
	bNameSuffix.WriteString(trw.goLocalName)
	b := strings.Builder{}
	for _, a := range resolvedType.Args {
		if a.IsNat {
			if a.isArith {
				bNameSuffix.WriteString(a.TemplateArgName)
				if b.Len() == 0 {
					b.WriteString("<")
				} else {
					b.WriteString(", ")
				}
				if halfResolve && a.T.OriginalName != "" {
					b.WriteString(a.T.OriginalName)
				} else {
					b.WriteString(strconv.FormatUint(uint64(a.Arith.Res), 10))
				}
			}
		} else {
			if b.Len() == 0 {
				b.WriteString("<")
			} else {
				b.WriteString(", ")
			}
			b.WriteString(a.TRW.CPPTypeStringInNamespace(bytesVersion, hppInc, a.T, halfResolve))
		}
	}
	if b.Len() != 0 {
		b.WriteString(">")
	}
	return bNameSuffix.String() + b.String()
}

func (trw *TypeRWStruct) cppDefaultInitializer(resolvedType ResolvedType, halfResolve bool) string {
	if trw.isUnwrapType() {
		return trw.Fields[0].t.CPPDefaultInitializer(resolvedType, halfResolve)
	}
	return "{}"
}

func (trw *TypeRWStruct) CPPHasBytesVersion() bool {
	return false // TODO
}

func (trw *TypeRWStruct) CPPTypeResettingCode(bytesVersion bool, val string) string {
	goGlobalName := addBytes(trw.goGlobalName, bytesVersion)
	//if trw.isUnwrapType() {
	//	return trw.Fields[0].t.trw.CPPTypeResettingCode(bytesVersion, val)
	//}
	return fmt.Sprintf("\t::%s::%sReset(%s);", trw.wr.gen.DetailsCPPNamespace, goGlobalName, val)
}

func (trw *TypeRWStruct) CPPTypeWritingCode(bytesVersion bool, val string, bare bool, natArgs []string, last bool) string {
	goGlobalName := addBytes(trw.goGlobalName, bytesVersion)
	//if trw.isUnwrapType() {
	//	prefix := ""
	//	if !bare {
	//		prefix = fmt.Sprintf("\ts.nat_write(0x%x); ", trw.wr.tlTag)
	//	}
	//	return prefix + trw.Fields[0].t.trw.CPPTypeWritingCode(bytesVersion, val, trw.Fields[0].Bare(), trw.replaceUnwrapArgs(natArgs), last)
	//}
	return fmt.Sprintf("\t::%s::%sWrite%s(s, %s%s);", trw.wr.gen.DetailsCPPNamespace, goGlobalName, addBare(bare), val, formatNatArgsCallCPP(natArgs))
}

func (trw *TypeRWStruct) CPPTypeReadingCode(bytesVersion bool, val string, bare bool, natArgs []string, last bool) string {
	goGlobalName := addBytes(trw.goGlobalName, bytesVersion)
	//if trw.isUnwrapType() {
	//	prefix := ""
	//	if !bare {
	//		prefix = fmt.Sprintf("\ts.nat_read_exact_tag(0x%x); ", trw.wr.tlTag)
	//	}
	//	return prefix + trw.Fields[0].t.trw.CPPTypeReadingCode(bytesVersion, val, trw.Fields[0].Bare(), trw.replaceUnwrapArgs(natArgs), last)
	//}
	return fmt.Sprintf("\t::%s::%sRead%s(s, %s%s);", trw.wr.gen.DetailsCPPNamespace, goGlobalName, addBare(bare), val, formatNatArgsCallCPP(natArgs))
}

func (trw *TypeRWStruct) CPPGenerateCode(hpp *strings.Builder, hppInc *DirectIncludesCPP, hppIncFwd *DirectIncludesCPP, hppDet *strings.Builder, hppDetInc *DirectIncludesCPP, cppDet *strings.Builder, cppDetInc *DirectIncludesCPP, bytesVersion bool, forwardDeclaration bool) {
	goLocalName := addBytes(trw.goLocalName, bytesVersion)
	goGlobalName := addBytes(trw.goGlobalName, bytesVersion)
	//if trw.wr.unionParent != nil && trw.wr.unionIsEnum {
	//	return
	//}
	mySuffix, myArgsDecl := fullyResolvedClassCppSuffixArgs(trw.wr.resolvedType)
	myName := trw.goLocalName + mySuffix
	myFullType := trw.cppTypeStringInNamespace(bytesVersion, hppDetInc, trw.wr.resolvedType, false)
	myFullTypeNoPrefix := strings.TrimPrefix(myFullType, "::") // Stupid C++ has sometimes problems with name resolution of definitions

	anyRecursive := false
	typeNamespace := trw.wr.gen.RootCPPNamespaceElements
	if trw.wr.tlName.Namespace != "" {
		typeNamespace = append(typeNamespace, trw.wr.tlName.Namespace)
	}
	cppStartNamespace(hpp, typeNamespace)
	// hpp.WriteString("// " + goLocalName + "\n") - uncommenting will lead to multiple definition error
	if len(myArgsDecl) != 0 {
		hpp.WriteString("template<" + strings.Join(myArgsDecl, ", ") + ">\n")
	}
	if forwardDeclaration { // TODO - does not work for typedef
		hpp.WriteString("struct " + myName + ";")
		cppFinishNamespace(hpp, typeNamespace)
		return
	}
	if trw.isTypeDef() {
		field := trw.Fields[0]
		fieldFullType := field.t.CPPTypeStringInNamespace(bytesVersion, hppInc, field.resolvedType, true)
		hpp.WriteString(fmt.Sprintf("using %s = %s;", myName, fieldFullType))
	} else {
		hpp.WriteString("struct " + myName + " {\n")
		for _, field := range trw.Fields {
			fieldFullType := field.t.CPPTypeStringInNamespace(bytesVersion, hppInc, field.resolvedType, true)
			fieldsMaskComment := ""
			if field.fieldMask != nil {
				fieldsMaskComment = fmt.Sprintf(" // Conditional: %s.%d", formatNatArgCPP(trw.Fields, *field.fieldMask), field.BitNumber)
			}
			if field.recursive {
				anyRecursive = true // requires destructor in cpp file
				hpp.WriteString(fmt.Sprintf("\tstd::shared_ptr<%s> %s{};%s\n", fieldFullType, field.cppName, fieldsMaskComment))
			} else {
				hpp.WriteString(fmt.Sprintf("\t%s %s%s;%s\n", fieldFullType, field.cppName, field.t.CPPDefaultInitializer(field.resolvedType, true), fieldsMaskComment))
			}
			//hpp.WriteString(fmt.Sprintf("\t// DebugString: %s\n", field.resolvedType.DebugString()))
		}
		if anyRecursive { // && len(trw.cppArgs) != 0
			hpp.WriteString(fmt.Sprintf("\n\t~%s() {}\n", goLocalName)) // TODO - move destructor to cpp
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
				trw.CPPTypeReadingCode(bytesVersion, "*this", true, trw.wr.NatParams, true),
				trw.CPPTypeWritingCode(bytesVersion, "*this", true, trw.wr.NatParams, true)))
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
				trw.CPPTypeReadingCode(bytesVersion, "*this", true, trw.wr.NatParams, true),
				trw.CPPTypeWritingCode(bytesVersion, "*this", true, trw.wr.NatParams, true),
				myFullTypeNoPrefix))
			if trw.wr.tlTag != 0 { // anonymous square brackets citizens or other exotic type
				hpp.WriteString(fmt.Sprintf(`
	void read_boxed(::basictl::tl_istream & s%[1]s);
	void write_boxed(::basictl::tl_ostream & s%[1]s)const;
`,
					formatNatArgsDeclCPP(trw.wr.NatParams),
					trw.CPPTypeResettingCode(bytesVersion, "*this"),
					trw.CPPTypeReadingCode(bytesVersion, "*this", false, trw.wr.NatParams, true),
					trw.CPPTypeWritingCode(bytesVersion, "*this", false, trw.wr.NatParams, true)))
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
					trw.CPPTypeReadingCode(bytesVersion, "*this", false, trw.wr.NatParams, true),
					trw.CPPTypeWritingCode(bytesVersion, "*this", false, trw.wr.NatParams, true),
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
		hppDet.WriteString(fmt.Sprintf(`
inline void %[1]sReadBoxed(::basictl::tl_istream & s, %[2]s& item%[3]s) {
	s.nat_read_exact_tag(0x%08[5]x);
	%[1]sRead(s, item%[4]s);
}
inline void %[1]sWriteBoxed(::basictl::tl_ostream & s, const %[2]s& item%[3]s) {
	s.nat_write(0x%08[5]x);
	%[1]sWrite(s, item%[4]s);
}
`,
			goGlobalName,
			myFullType,
			formatNatArgsDeclCPP(trw.wr.NatParams),
			formatNatArgsCallCPP(trw.wr.NatParams),
			trw.wr.tlTag))
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
`, field.cppName, field.t.trw.CPPTypeResettingCode(bytesVersion, fmt.Sprintf("*item.%s", field.cppName))))
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
		_ = field.t.CPPTypeStringInNamespace(bytesVersion, cppDetInc, field.resolvedType, false) // only fill includes
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
`, field.t.CPPTypeStringInNamespace(bytesVersion, cppDetInc, field.resolvedType, false), field.cppName))
		}
		_ = field.t.CPPTypeStringInNamespace(bytesVersion, cppDetInc, field.resolvedType, false) // only fill includes
		_ = field.t.CPPTypeStringInNamespace(bytesVersion, hppDetInc, field.resolvedType, false) // only fill includes
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
*/
