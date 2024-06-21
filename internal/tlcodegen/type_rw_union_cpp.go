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

func (trw *TypeRWUnion) CPPFillRecursiveChildren(visitedNodes map[*TypeRWWrapper]bool) {
	for _, f := range trw.Fields {
		if !f.recursive {
			f.t.CPPFillRecursiveChildren(visitedNodes)
		}
	}
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
	return fmt.Sprintf("\tif (!::%s::%sWrite%s(s, %s%s)) { return false; }", trw.wr.gen.DetailsCPPNamespace, goGlobalName, addBare(bare), val, joinWithCommas(natArgs))
}

func (trw *TypeRWUnion) CPPTypeReadingCode(bytesVersion bool, val string, bare bool, natArgs []string, last bool) string {
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
	return fmt.Sprintf("\tif (!::%s::%sRead%s(s, %s%s)) { return false; }", trw.wr.gen.DetailsCPPNamespace, goGlobalName, addBare(bare), val, joinWithCommas(natArgs))
}

func (trw *TypeRWUnion) CPPGenerateCode(hpp *strings.Builder, hppInc *DirectIncludesCPP, hppIncFwd *DirectIncludesCPP, hppDet *strings.Builder, hppDetInc *DirectIncludesCPP, cppDet *strings.Builder, cppDetInc *DirectIncludesCPP, bytesVersion bool, forwardDeclaration bool) {
	// goLocalName := addBytes(trw.goLocalName, bytesVersion)
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)

	_, myArgsDecl := trw.wr.fullyResolvedClassCppNameArgs()
	myFullType := trw.cppTypeStringInNamespace(bytesVersion, hppDetInc)
	//fmt.Printf("Ts: %s %s\n", myFullType, strings.Join(myArgsDecl, ", "))
	//fmt.Printf("    %s\n", trw.wr.cppLocalName)
	myFullTypeNoPrefix := strings.TrimPrefix(myFullType, "::") // Stupid C++ has sometimes problems with name resolution of definitions

	typeNamespace := trw.wr.gen.RootCPPNamespaceElements
	if trw.wr.tlName.Namespace != "" {
		typeNamespace = append(typeNamespace, trw.wr.tlName.Namespace)
	}

	if forwardDeclaration {
		cppStartNamespace(hpp, typeNamespace)
		if len(myArgsDecl) != 0 {
			hpp.WriteString("template<" + strings.Join(myArgsDecl, ", ") + ">\n")
		}
		hpp.WriteString("struct " + trw.wr.cppLocalName + ";")
		cppFinishNamespace(hpp, typeNamespace)
		return
	}

	cppStartNamespace(hppDet, trw.wr.gen.DetailsCPPNamespaceElements)
	hppDet.WriteString(fmt.Sprintf(`
void %[1]sReset(%[2]s& item);
bool %[1]sReadBoxed(::basictl::tl_istream & s, %[2]s& item%[3]s);
bool %[1]sWriteBoxed(::basictl::tl_ostream & s, const %[2]s& item%[3]s);
`, goGlobalName, myFullType, formatNatArgsDeclCPP(trw.wr.NatParams)))

	cppFinishNamespace(hppDet, trw.wr.gen.DetailsCPPNamespaceElements)

	for _, typeDep := range trw.AllTypeDependencies(true) {
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

	cppDet.WriteString(fmt.Sprintf(`
static const std::string_view %[1]s_tbl_tl_name[]{%[2]s};
static const uint32_t %[1]s_tbl_tl_tag[]{%[3]s};
`,
		goGlobalName,
		trw.CPPAllNames(bytesVersion),
		trw.CPPAllTags(bytesVersion)))

	if len(myArgsDecl) == 0 {
		// cppStartNamespace(cppDet, trw.wr.gen.RootCPPNamespaceElements)
		hpp.WriteString(fmt.Sprintf(`
	bool read_boxed(::basictl::tl_istream & s%[1]s);
	bool write_boxed(::basictl::tl_ostream & s%[1]s)const;
`,
			formatNatArgsDeclCPP(trw.wr.NatParams),
			trw.CPPTypeResettingCode(bytesVersion, "*this"),
			trw.CPPTypeReadingCode(bytesVersion, "*this", false, formatNatArgsAddNat(trw.wr.NatParams), true),
			trw.CPPTypeWritingCode(bytesVersion, "*this", false, formatNatArgsAddNat(trw.wr.NatParams), true)))
		cppDet.WriteString(fmt.Sprintf(`
bool %[5]s::read_boxed(::basictl::tl_istream & s%[1]s) {
%[3]s
	return true;
}
bool %[5]s::write_boxed(::basictl::tl_ostream & s%[1]s)const {
%[4]s
	return true;
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
			goGlobalName))
	}
	hpp.WriteString("};\n")
	cppFinishNamespace(hpp, typeNamespace)

	cppDet.WriteString(fmt.Sprintf(`
void %[7]s::%[1]sReset(%[2]s& item) {
	item.value.emplace<0>(); // TODO - optimize, if already 0, call Reset function
}

bool %[7]s::%[1]sReadBoxed(::basictl::tl_istream & s, %[2]s& item%[3]s) {
	uint32_t nat;
	s.nat_read(nat);
	switch (nat) {
%[5]s	default:
		return s.set_error_union_tag();
    }
	return true;
}

bool %[7]s::%[1]sWriteBoxed(::basictl::tl_ostream & s, const %[2]s& item%[3]s) {
	s.nat_write(%[1]s_tbl_tl_tag[item.value.index()]);
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
	))
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

/*
func (union *TypeRWUnion) GenerateFields(bytesVersion bool) string {
	var s strings.Builder
	for _, field := range union.Fields {
		if !field.t.IsTrueType() {
			s.WriteString(fmt.Sprintf("value%s %s%s\n", field.goName, ifString(field.recursive, "*", ""), field.t.trw.TypeString(bytesVersion)))
		}
	}
	s.WriteString("index int\n")
	return s.String()
}

func (union *TypeRWUnion) GenerateConstructorsBehavior(bytesVersion bool) string {
	var s strings.Builder
	for i, typ := range union.Fields {
		s.WriteString(fmt.Sprintf("%[3]s%[4]s%[5]s\n",
			addBytes(union.goGlobalName, bytesVersion), typ.t.trw.TypeString(bytesVersion), union.As(bytesVersion, i, typ), union.ResetTo(bytesVersion, i, typ), union.Set(bytesVersion, i, typ)))
	}
	return s.String()
}

func (union *TypeRWUnion) As(bytesVersion bool, i int, field Field) string {
	var s strings.Builder
	s.WriteString(fmt.Sprintf(`func (item %s) Is%s() bool { return item.index == %d }
`,
		addBytes(union.goGlobalName, bytesVersion), field.goName, i))
	if field.t.IsTrueType() {
		if !union.IsEnum {
			s.WriteString(fmt.Sprintf(`func (item *%[1]s) As%[5]s() (%[2]s, bool) {
	var value %[3]s
	return value, item.index == %[4]d }
`, addBytes(union.goGlobalName, bytesVersion), field.t.trw.TypeString(bytesVersion), field.t.trw.TypeString(bytesVersion), i, field.goName))
		}
	} else {
		s.WriteString(fmt.Sprintf(`func (item *%[1]s) As%[4]s() (*%[2]s, bool) {
if item.index != %[3]d {
	return nil, false
}
	return %[5]sitem.value%[4]s, true }
`, addBytes(union.goGlobalName, bytesVersion), field.t.trw.TypeString(bytesVersion), i, field.goName, ifString(field.recursive, "", "&")))
	}
	return s.String()
}

func (union *TypeRWUnion) ResetTo(bytesVersion bool, i int, field Field) string {
	if union.IsEnum {
		return ""
	}
	if field.t.IsTrueType() {
		return fmt.Sprintf(`func (item *%[1]s) ResetTo%[4]s() { item.index = %[3]d }
`, addBytes(union.goGlobalName, bytesVersion), field.t.trw.TypeString(bytesVersion), i, field.goName)
	}
	if field.recursive {
		return fmt.Sprintf(`func (item *%[1]s) ResetTo%[4]s() *%[2]s {
	item.index = %[3]d
	if item.value%[4]s == nil {
		var value %[2]s
		item.value%[4]s = &value
	} else {
		%[5]s
	}
	return item.value%[4]s }
`, addBytes(union.goGlobalName, bytesVersion), field.t.trw.TypeString(bytesVersion), i, field.goName,
			field.t.trw.TypeResettingCode(bytesVersion, "item.value"+field.goName, true))
	}
	return fmt.Sprintf(`func (item *%[1]s) ResetTo%[4]s() *%[2]s {
	item.index = %[3]d
	%[5]s
	return &item.value%[4]s }
`, addBytes(union.goGlobalName, bytesVersion), field.t.trw.TypeString(bytesVersion), i, field.goName,
		field.t.trw.TypeResettingCode(bytesVersion, "item.value"+field.goName, false))
}

func (union *TypeRWUnion) Set(bytesVersion bool, i int, field Field) string {
	if field.t.IsTrueType() {
		return fmt.Sprintf(`func (item *%s) Set%s() { item.index = %d }
`, addBytes(union.goGlobalName, bytesVersion), field.goName, i)
	}
	if field.recursive {
		return fmt.Sprintf(`func (item *%[1]s) Set%[4]s(value %[2]s) {
	item.index = %[3]d
	if item.value%[4]s == nil {
		item.value%[4]s = &value
	} else {
		*item.value%[4]s = value
	}
}
`, addBytes(union.goGlobalName, bytesVersion), field.t.trw.TypeString(bytesVersion), i, field.goName)
	}
	return fmt.Sprintf(`func (item *%[1]s) Set%[4]s(value %[2]s) {
	item.index = %[3]d
	item.value%[4]s = value
}
`, addBytes(union.goGlobalName, bytesVersion), field.t.trw.TypeString(bytesVersion), i, field.goName)
}

func (union *TypeRWUnion) GenerateReadBoxed(bytesVersion bool) string {
	var s strings.Builder
	s.WriteString(`var tag uint32
if err := basictl.NatRead(r, &tag); err != nil {
	return err
}
switch tag {
`)
	for i, field := range union.Fields {
		s.WriteString(fmt.Sprintf("case %#x:\nitem.index = %d", field.t.tlTag, i))
		if field.t.IsTrueType() {
			s.WriteString("\nreturn nil\n")
			continue
		}
		setRecursiveText := ifString(field.recursive, fmt.Sprintf(`
		if item.value%[2]s == nil {
			var value %[1]s
			item.value%[2]s = &value
		}`, field.t.trw.TypeString(bytesVersion), field.goName), "")
		s.WriteString(fmt.Sprintf("%s\n%s\n", setRecursiveText, field.t.trw.TypeReadingCode(bytesVersion, fmt.Sprintf("item.value%s", field.goName),
			true, union.Fields[0].t.NatParams, // union arg names are from first field, see same comment in generateType
			false, true)))
	}
	s.WriteString(fmt.Sprintf("default:\nreturn internal.ErrorInvalidUnionTag(\"%s\", tag)}", union.Fields[0].t.tlName))
	return s.String()
}

func (union *TypeRWUnion) GenerateWriteBoxed(bytesVersion bool) string {
	if union.IsEnum {
		return "; return nil"
	}
	var s strings.Builder
	s.WriteString("\nswitch item.index {\n")
	for i, field := range union.Fields {
		if field.t.IsTrueType() {
			s.WriteString(fmt.Sprintf("case %d: return nil\n", i))
		} else {
			s.WriteString(fmt.Sprintf("case %d:\n%s\n", i, field.t.trw.TypeWritingCode(bytesVersion,
				fmt.Sprintf("item.value%s", field.goName),
				true, union.Fields[0].t.NatParams, // union arg names are from first field, see same comment in generateType
				false, true)))
		}
	}
	s.WriteString("default: // Impossible due to panic above\nreturn nil\n}")
	return s.String()
}

func (union *TypeRWUnion) generateEnumAlias(bytesVersion bool) string {
	goName := addBytes(union.goGlobalName, false)
	var s strings.Builder
	if bytesVersion {
		return s.String()
	}
	s.WriteString(fmt.Sprintf("var _%s = [%d]internal.UnionElement{\n", goName, len(union.Fields)))
	for _, x := range union.Fields {
		s.WriteString(fmt.Sprintf(`{tlTag:%#x, tlName:"%s", tlString:"%s#%08x"},
`, x.t.tlTag, x.t.tlName, x.t.tlName, x.t.tlTag))
	}
	s.WriteString("}\n\n")
	if union.IsEnum {
		for i, x := range union.Fields {
			s.WriteString(fmt.Sprintf("func %s() %s { return %s{index:%d} }\n", x.t.trw.TypeString(false), goName, goName, i))
		}
	}
	return s.String()
}
*/

func (trw *TypeRWUnion) CPPAllTags(bytesVersion bool) string {
	var s strings.Builder
	for fieldIndex, field := range trw.Fields {
		if fieldIndex != 0 {
			s.WriteString(", ")
		}
		s.WriteString(fmt.Sprintf("0x%x", field.t.tlTag))
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
		s.WriteString(fmt.Sprintf("\tcase 0x%x:\n\t\tif (item.value.index() != %d) { item.value.emplace<%d>(); }\n", field.t.tlTag, fieldIndex, fieldIndex))
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
