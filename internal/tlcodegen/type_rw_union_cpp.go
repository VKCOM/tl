// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlcodegen

import (
	"fmt"
	"github.com/vkcom/tl/internal/utils"
	"strings"
)

func (trw *TypeRWUnion) CPPTypeJSONEmptyCondition(bytesVersion bool, val string, ref bool, deps []string) string {
	return ""
}

func (trw *TypeRWUnion) CPPFillRecursiveChildren(visitedNodes map[*TypeRWWrapper]bool) {
	for _, f := range trw.Fields {
		if !f.recursive {
			f.t.CPPFillRecursiveChildren(visitedNodes)
		}
	}
}

func (trw *TypeRWUnion) CPPAllowCurrentDefinition() bool {
	return true
}

func (trw *TypeRWUnion) cppTypeStringInNamespace(bytesVersion bool, hppInc *DirectIncludesCPP) string {
	_, _, args := trw.wr.cppTypeStringInNamespace(bytesVersion, hppInc, false, HalfResolvedArgument{})
	return trw.wr.cppNamespaceQualifier() + trw.wr.cppLocalName + args
}

func (trw *TypeRWUnion) cppTypeStringInNamespaceHalfResolved(bytesVersion bool, hppInc *DirectIncludesCPP, halfResolved HalfResolvedArgument) string {
	_, _, args := trw.wr.cppTypeStringInNamespace(bytesVersion, hppInc, true, halfResolved)
	return trw.wr.cppNamespaceQualifier() + trw.wr.cppLocalName + args
}

func (trw *TypeRWUnion) cppTypeStringInNamespaceHalfResolved2(bytesVersion bool, typeReduction EvaluatedType) string {
	return trw.wr.cppNamespaceQualifier() + trw.wr.cppLocalName + trw.wr.cppTypeArguments(bytesVersion, typeReduction.Type)
}

func (trw *TypeRWUnion) cppDefaultInitializer(halfResolved HalfResolvedArgument, halfResolve bool) string {
	return ""
}

func (trw *TypeRWUnion) CPPHasBytesVersion() bool {
	return false // TODO
}

func (trw *TypeRWUnion) CPPTypeResettingCode(bytesVersion bool, val string) string {
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
	return fmt.Sprintf("\t::%s::%sReset(%s);", trw.wr.gen.DetailsCPPNamespace, goGlobalName, val)
}

func (trw *TypeRWUnion) CPPTypeWritingCode(bytesVersion bool, val string, bare bool, natArgs []string, last bool) string {
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
	return fmt.Sprintf("\tif (!::%s::%sWrite%s(s, %s%s)) { return s.set_error_unknown_scenario(); }", trw.wr.gen.DetailsCPPNamespace, goGlobalName, addBare(bare), val, joinWithCommas(natArgs))
}

func (trw *TypeRWUnion) CPPTypeWritingJsonCode(bytesVersion bool, val string, bare bool, natArgs []string, last bool) string {
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
	return fmt.Sprintf("\tif (!::%s::%sWriteJSON(s, %s%s)) { return false; }", trw.wr.gen.DetailsCPPNamespace, goGlobalName, val, joinWithCommas(natArgs))
}

func (trw *TypeRWUnion) CPPTypeReadingCode(bytesVersion bool, val string, bare bool, natArgs []string, last bool) string {
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
	return fmt.Sprintf("\tif (!::%s::%sRead%s(s, %s%s)) { return s.set_error_unknown_scenario(); }", trw.wr.gen.DetailsCPPNamespace, goGlobalName, addBare(bare), val, joinWithCommas(natArgs))
}

func (trw *TypeRWUnion) CPPGenerateCode(hpp *strings.Builder, hppInc *DirectIncludesCPP, hppIncFwd *DirectIncludesCPP, hppDet *strings.Builder, hppDetInc *DirectIncludesCPP, cppDet *strings.Builder, cppDetInc *DirectIncludesCPP, bytesVersion bool, forwardDeclaration bool) {
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)

	_, myArgsDecl := trw.wr.fullyResolvedClassCppNameArgs()

	typeNamespace := trw.wr.gen.RootCPPNamespaceElements
	if trw.wr.tlName.Namespace != "" {
		typeNamespace = append(typeNamespace, trw.wr.tlName.Namespace)
	}
	if hpp != nil {
		if forwardDeclaration {
			cppStartNamespace(hpp, typeNamespace)
			if len(myArgsDecl) != 0 {
				hpp.WriteString("template<" + strings.Join(myArgsDecl, ", ") + ">\n")
			}
			hpp.WriteString("struct " + trw.wr.cppLocalName + ";")
			cppFinishNamespace(hpp, typeNamespace)
			return
		}

		for _, typeDep := range trw.AllTypeDependencies(true, false) {
			if typeDep.typeComponent == trw.wr.typeComponent {
				typeDep.trw.CPPGenerateCode(hpp, nil, nil, nil, hppDetInc, nil, cppDetInc, bytesVersion, true)
			}
		}

		cppStartNamespace(hpp, typeNamespace)
		if len(myArgsDecl) != 0 {
			hpp.WriteString("template<" + strings.Join(myArgsDecl, ", ") + ">\n")
		}
		hpp.WriteString("struct " + trw.wr.cppLocalName + " {\n")
		hpp.WriteString("\tstd::variant<")
		for filedIndex, field := range trw.Fields {
			if filedIndex != 0 {
				hpp.WriteString(", ")
			}
			if field.recursive {
				fieldFullType := field.t.CPPTypeStringInNamespaceHalfResolved(bytesVersion, hppIncFwd, field.halfResolved)
				hpp.WriteString(fmt.Sprintf("std::shared_ptr<%s>", fieldFullType))
			} else {
				fieldFullType := field.t.CPPTypeStringInNamespaceHalfResolved(bytesVersion, hppInc, field.halfResolved)
				hpp.WriteString(fieldFullType)
				//hpp.WriteString(fmt.Sprintf("\t// DebugString: %s\n", field.resolvedType.DebugString()))
			}
		}
		hpp.WriteString("> value;\n\n")
		hpp.WriteString(trw.CPPGetters(bytesVersion))
		hpp.WriteString("\n")
		hpp.WriteString(trw.CPPSetters(bytesVersion))

		hpp.WriteString(`
	std::string_view tl_name() const;
	uint32_t tl_tag() const;
`)

		if len(myArgsDecl) == 0 {
			// cppStartNamespace(cppDet, trw.wr.gen.RootCPPNamespaceElements)
			hpp.WriteString(fmt.Sprintf(`
	bool write_json(std::ostream& s%[1]s)const;

	bool read_boxed(::basictl::tl_istream & s%[1]s) noexcept;
	bool write_boxed(::basictl::tl_ostream & s%[1]s)const noexcept;
	
	void read_boxed_or_throw(::basictl::tl_throwable_istream & s%[1]s);
	void write_boxed_or_throw(::basictl::tl_throwable_ostream & s%[1]s)const;
`,
				formatNatArgsDeclCPP(trw.wr.NatParams),
				trw.CPPTypeResettingCode(bytesVersion, "*this"),
				trw.CPPTypeReadingCode(bytesVersion, "*this", false, formatNatArgsAddNat(trw.wr.NatParams), true),
				trw.CPPTypeWritingCode(bytesVersion, "*this", false, formatNatArgsAddNat(trw.wr.NatParams), true)))
		}
		hpp.WriteString("};\n")
		cppFinishNamespace(hpp, typeNamespace)
	}

	hppTmpInclude := DirectIncludesCPP{ns: map[*TypeRWWrapper]CppIncludeInfo{}}
	myFullType := trw.cppTypeStringInNamespace(bytesVersion, &hppTmpInclude)
	myFullTypeNoPrefix := strings.TrimPrefix(myFullType, "::") // Stupid C++ has sometimes problems with name resolution of definitions

	if hppDet != nil {
		utils.AppendMap(hppTmpInclude.ns, &hppDetInc.ns)

		cppStartNamespace(hppDet, trw.wr.gen.DetailsCPPNamespaceElements)
		hppDet.WriteString(fmt.Sprintf(`
void %[1]sReset(%[2]s& item) noexcept;

bool %[1]sWriteJSON(std::ostream & s, const %[2]s& item%[3]s) noexcept;
bool %[1]sReadBoxed(::basictl::tl_istream & s, %[2]s& item%[3]s) noexcept;
bool %[1]sWriteBoxed(::basictl::tl_ostream & s, const %[2]s& item%[3]s) noexcept;
`, goGlobalName, myFullType, formatNatArgsDeclCPP(trw.wr.NatParams)))
		cppFinishNamespace(hppDet, trw.wr.gen.DetailsCPPNamespaceElements)
	}

	if cppDet != nil {
		cppDet.WriteString(fmt.Sprintf(`
static const std::string_view %[1]s_tbl_tl_name[]{%[2]s};
static const uint32_t %[1]s_tbl_tl_tag[]{%[3]s};
`,
			goGlobalName,
			trw.CPPAllNames(bytesVersion),
			trw.CPPAllTags(bytesVersion)))
		if len(myArgsDecl) == 0 {
			cppDet.WriteString(fmt.Sprintf(`
bool %[5]s::write_json(std::ostream & s%[1]s)const {
%[7]s
	return true;
}
bool %[5]s::read_boxed(::basictl::tl_istream & s%[1]s) noexcept {
%[3]s
	return true;
}
bool %[5]s::write_boxed(::basictl::tl_ostream & s%[1]s)const noexcept {
%[4]s
	return true;
}

void %[5]s::read_boxed_or_throw(::basictl::tl_throwable_istream & s%[1]s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2%[8]s);
	s2.pass_data(s);
}

void %[5]s::write_boxed_or_throw(::basictl::tl_throwable_ostream & s%[1]s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2%[8]s);
	s2.pass_data(s);
}

std::string_view %[5]s::tl_name() const {
	return %[6]s_tbl_tl_name[value.index()];
}
uint32_t %[5]s::tl_tag() const {
	return %[6]s_tbl_tl_tag[value.index()];
}

`,
				formatNatArgsDeclCPP(trw.wr.NatParams),
				trw.CPPTypeResettingCode(bytesVersion, "*this"),
				trw.CPPTypeReadingCode(bytesVersion, "*this", false, formatNatArgsAddNat(trw.wr.NatParams), true),
				trw.CPPTypeWritingCode(bytesVersion, "*this", false, formatNatArgsAddNat(trw.wr.NatParams), true),
				myFullTypeNoPrefix,
				goGlobalName,
				trw.CPPTypeWritingJsonCode(bytesVersion, "*this", false, formatNatArgsAddNat(trw.wr.NatParams), true),
				formatNatArgsCallCPP(trw.wr.NatParams),
			))
		}
		cppDet.WriteString(fmt.Sprintf(`
void %[7]s::%[1]sReset(%[2]s& item) noexcept{
	item.value.emplace<0>(); // TODO - optimize, if already 0, call Reset function
}

bool %[7]s::%[1]sWriteJSON(std::ostream & s, const %[2]s& item%[3]s) noexcept {
%[8]s
	return true;
}
bool %[7]s::%[1]sReadBoxed(::basictl::tl_istream & s, %[2]s& item%[3]s) noexcept {
	uint32_t nat;
	if (!s.nat_read(nat)) { return false; }
	switch (nat) {
%[5]s	default:
		return s.set_error_union_tag();
    }
	return true;
}

bool %[7]s::%[1]sWriteBoxed(::basictl::tl_ostream & s, const %[2]s& item%[3]s) noexcept{
	if (!s.nat_write(%[1]s_tbl_tl_tag[item.value.index()])) { return false; }
	switch (item.value.index()) {
%[6]s	}
	return true;
}
`,
			goGlobalName,
			myFullType,
			formatNatArgsDeclCPP(trw.wr.NatParams),
			"",
			trw.CPPReadFields(bytesVersion, hppInc, cppDetInc),
			trw.CPPWriteFields(bytesVersion),
			trw.wr.gen.DetailsCPPNamespace,
			trw.CPPWriteJSONFields(bytesVersion),
		))
	}
	/*
			code := `type %[1]s struct {%[2]s}

		func (item %[1]s) TLName() string { return _%[9]s[item.index].tlName }
		func (item %[1]s) TLTag() uint32 { return _%[9]s[item.index].tlTag }

		func (item *%[1]s) Reset() { item.index = 0 }

		%[5]s
		func (item*%[1]s) ReadBoxed(r *bytes.Buffer%[8]s) error {
			%[6]s
		}

		func (item %[3]s%[1]s) WriteBoxed(w *bytes.Buffer%[8]s) error {
			basictl.NatWrite(w, _%[9]s[item.index].tlTag) %[7]s
		}
		`
			var s strings.Builder
			// FIXME trw.generateEnumAlias generates code for Go not for C++
			s.WriteString(trw.generateEnumAlias(bytesVersion))
			s.WriteString(fmt.Sprintf(code,
				addBytes(trw.goGlobalName, bytesVersion),
				trw.GenerateFields(bytesVersion),
				ifString(trw.IsEnum, "", "*"),
				"",
				trw.GenerateConstructorsBehavior(bytesVersion),
				trw.GenerateReadBoxed(bytesVersion),
				trw.GenerateWriteBoxed(bytesVersion),
				formatNatArgsDeclCPP(trw.wr.NatParams),
				addBytes(trw.goGlobalName, false),
			))
	*/
}

func (trw *TypeRWUnion) CPPAllTags(bytesVersion bool) string {
	var s strings.Builder
	for fieldIndex, field := range trw.Fields {
		if fieldIndex != 0 {
			s.WriteString(", ")
		}
		s.WriteString(fmt.Sprintf("0x%08x", field.t.tlTag))
	}
	return s.String()
}

func (trw *TypeRWUnion) CPPAllNames(bytesVersion bool) string {
	var s strings.Builder
	for fieldIndex, field := range trw.Fields {
		if fieldIndex != 0 {
			s.WriteString(", ")
		}
		s.WriteString(fmt.Sprintf("\"%s\"", field.t.tlName))
	}
	return s.String()
}

func (trw *TypeRWUnion) CPPWriteFields(bytesVersion bool) string {
	var s strings.Builder
	for fieldIndex, field := range trw.Fields {
		if !field.t.IsTrueType() {
			s.WriteString(fmt.Sprintf("\tcase %d:\n", fieldIndex))
			s.WriteString("\t" +
				field.t.trw.CPPTypeWritingCode(bytesVersion, addAsterisk(field.recursive, fmt.Sprintf("std::get<%d>(item.value)", fieldIndex)),
					true, formatNatArgsCPP(trw.Fields, field.natArgs), false) + "\n")
			s.WriteString("\t\tbreak;\n")
		}
	}
	return s.String()
}

func (trw *TypeRWUnion) CPPWriteJSONFields(bytesVersion bool) string {
	var s strings.Builder
	if trw.IsEnum {
		s.WriteString(fmt.Sprintf(`	s << "\"" << %s_tbl_tl_name[item.value.index()] << "\"";`, trw.wr.goGlobalName))
	} else {
		s.WriteString(fmt.Sprintf(`	s << "{";
	s << "\"type\":";
	s << "\"" << %s_tbl_tl_name[item.value.index()] << "\"";
	switch (item.value.index()) {
`, trw.wr.goGlobalName))

		for fieldIndex, field := range trw.Fields {
			indent := 1
			if !field.t.IsTrueType() {
				s.WriteString(fmt.Sprintf("%scase %d:\n", strings.Repeat("\t", indent), fieldIndex))
				indent++
				emptyCondition := field.t.trw.CPPTypeJSONEmptyCondition(bytesVersion, fmt.Sprintf("std::get<%d>(item.value)", fieldIndex), field.recursive, nil)
				if emptyCondition != "" {
					s.WriteString(fmt.Sprintf("%sif (%s) {\n", strings.Repeat("\t", indent), emptyCondition))
					indent++
				}
				s.WriteString(fmt.Sprintf(`%ss << ",\"value\":";
`, strings.Repeat("\t", indent)))
				s.WriteString(strings.Repeat("\t", indent-1) +
					field.t.trw.CPPTypeWritingJsonCode(bytesVersion, addAsterisk(field.recursive, fmt.Sprintf("std::get<%d>(item.value)", fieldIndex)),
						true, formatNatArgsCPP(trw.Fields, field.natArgs), false) + "\n")
				if emptyCondition != "" {
					indent--
					s.WriteString(fmt.Sprintf("%s}\n", strings.Repeat("\t", indent)))
				}
				s.WriteString(fmt.Sprintf("%sbreak;\n", strings.Repeat("\t", indent)))
			}
		}
		s.WriteString(`	}
	s << "}";`)
	}
	return s.String()
}

func (trw *TypeRWUnion) CPPSetters(bytesVersion bool) string {
	var s strings.Builder
	_, myArgsDecl := trw.wr.fullyResolvedClassCppNameArgs()
	for fieldIndex, field := range trw.Fields {
		if field.t.IsTrueType() {
			initValue := ""
			if len(myArgsDecl) != 0 {
				initValue = "{}"
			}
			s.WriteString(fmt.Sprintf("\tvoid set_%s() { value.emplace<%d>(%s); }\n", field.cppName, fieldIndex, initValue))
		}
	}
	return s.String()
}

func (trw *TypeRWUnion) CPPGetters(bytesVersion bool) string {
	var s strings.Builder
	for fieldIndex, field := range trw.Fields {
		s.WriteString(fmt.Sprintf("\tbool is_%s() const { return value.index() == %d; }\n", field.cppName, fieldIndex))
	}
	return s.String()
}

func (trw *TypeRWUnion) CPPReadFields(bytesVersion bool, hppInc *DirectIncludesCPP, cppDetInc *DirectIncludesCPP) string {
	var s strings.Builder
	for fieldIndex, field := range trw.Fields {
		s.WriteString(fmt.Sprintf("\tcase 0x%08x:\n\t\tif (item.value.index() != %d) { item.value.emplace<%d>(); }\n", field.t.tlTag, fieldIndex, fieldIndex))
		if !field.t.IsTrueType() {
			s.WriteString("\t")
			_ = field.t.CPPTypeStringInNamespace(bytesVersion, cppDetInc) // only fill includes
			s.WriteString(
				field.t.trw.CPPTypeReadingCode(bytesVersion, addAsterisk(field.recursive, fmt.Sprintf("std::get<%d>(item.value)", fieldIndex)),
					field.Bare(), formatNatArgsCPP(trw.Fields, field.natArgs),
					false) + "\n")
		}
		s.WriteString("\t\tbreak;\n")
	}
	return s.String()
}
