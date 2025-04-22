// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlcodegen

import (
	"fmt"
	"strings"
)

func (trw *TypeRWBrackets) CPPTypeJSONEmptyCondition(bytesVersion bool, val string, ref bool, deps []string) string {
	if trw.dictLike || trw.vectorLike || trw.dynamicSize {
		if trw.dynamicSize && len(deps) != 0 {
			return fmt.Sprintf("(%s.size() != 0) || (%s != 0)", addAsteriskAndBrackets(ref, val), deps[0])
		}
		return fmt.Sprintf("%s.size() != 0", addAsteriskAndBrackets(ref, val))
	}
	return ""
}

func (trw *TypeRWBrackets) CPPFillRecursiveChildren(visitedNodes map[*TypeRWWrapper]bool) {
	if trw.vectorLike || trw.dynamicSize {
		return
	}
	trw.element.t.CPPFillRecursiveChildren(visitedNodes)
}

func (trw *TypeRWBrackets) cppTypeStringInNamespace(bytesVersion bool, hppInc *DirectIncludesCPP) string {
	hppInc.ns[trw.wr] = CppIncludeInfo{componentId: trw.wr.typeComponent, namespace: trw.wr.tlName.Namespace}
	//if trw.dictLike && !bytesVersion {
	//	TODO - which arguments must map have is very complicated
	//return fmt.Sprintf("std::map<%s, %s>",
	//	trw.dictKeyField.t.CPPTypeStringInNamespace(bytesVersion, hppInc, trw.dictKeyField.resolvedType, halfResolve),
	//	trw.dictValueField.t.CPPTypeStringInNamespace(bytesVersion, hppInc, trw.dictValueField.resolvedType, halfResolve))
	//}
	if trw.dictLike {
		pairType := trw.element.t.trw.(*TypeRWStruct)
		if len(pairType.wr.origTL[0].TemplateArguments) == 1 {
			keyType := pairType.Fields[0]
			valueType := pairType.Fields[1]
			return fmt.Sprintf("std::map<%[1]s, %[2]s>",
				keyType.t.CPPTypeStringInNamespace(bytesVersion, hppInc),
				valueType.t.CPPTypeStringInNamespace(bytesVersion, hppInc),
			)
		}
	}
	if trw.vectorLike || trw.dynamicSize {
		return fmt.Sprintf("std::vector<%s>", trw.element.t.CPPTypeStringInNamespace(bytesVersion, hppInc))
	}
	return fmt.Sprintf("std::array<%s, %d>", trw.element.t.CPPTypeStringInNamespace(bytesVersion, hppInc), trw.size)
}

func (trw *TypeRWBrackets) cppTypeStringInNamespaceHalfResolved(bytesVersion bool, hppInc *DirectIncludesCPP, halfResolved HalfResolvedArgument) string {
	//if trw.dictLike && !bytesVersion {
	//	TODO - which arguments must map have is very complicated
	//return fmt.Sprintf("std::map<%s, %s>",
	//	trw.dictKeyField.t.CPPTypeStringInNamespace(bytesVersion, hppInc, trw.dictKeyField.resolvedType, halfResolve),
	//	trw.dictValueField.t.CPPTypeStringInNamespace(bytesVersion, hppInc, trw.dictValueField.resolvedType, halfResolve))
	//}
	if trw.vectorLike {
		return fmt.Sprintf("std::vector<%s>", trw.element.t.CPPTypeStringInNamespaceHalfResolved(bytesVersion, hppInc, halfResolved.Args[0]))
	}
	if trw.dynamicSize {
		return fmt.Sprintf("std::vector<%s>", trw.element.t.CPPTypeStringInNamespaceHalfResolved(bytesVersion, hppInc, halfResolved.Args[1]))
	}
	if halfResolved.Args[0].Name != "" {
		return fmt.Sprintf("std::array<%s, %s>", trw.element.t.CPPTypeStringInNamespaceHalfResolved(bytesVersion, hppInc, halfResolved.Args[1]), halfResolved.Args[0].Name)
	}
	return fmt.Sprintf("std::array<%s, %d>", trw.element.t.CPPTypeStringInNamespaceHalfResolved(bytesVersion, hppInc, halfResolved.Args[1]), trw.size)
}

func (trw *TypeRWBrackets) cppTypeStringInNamespaceHalfResolved2(bytesVersion bool, typeReduction EvaluatedType) string {
	switch len(typeReduction.Type.Arguments) {
	case 1:
		if cppIsDictionaryElement(trw.element.t) && typeReduction.Type.Arguments[0].Type != nil && len(typeReduction.Type.Arguments[0].Type.Arguments) == 1 {
			pairReduction := typeReduction.Type.Arguments[0].Type
			pairType := trw.element.t.trw.(*TypeRWStruct)

			keyType := pairType.Fields[0]
			valueType := pairType.Fields[1]

			typesInfo := trw.wr.gen.typesInfo

			keyReduction := typesInfo.FieldTypeReduction(pairReduction, 0)
			valueReduction := typesInfo.FieldTypeReduction(pairReduction, 1)

			return fmt.Sprintf("std::map<%[1]s, %[2]s>",
				keyType.t.CPPTypeStringInNamespaceHalfResolved2(bytesVersion, keyReduction),
				valueType.t.CPPTypeStringInNamespaceHalfResolved2(bytesVersion, valueReduction),
			)
		}
		return fmt.Sprintf("std::vector<%s>", trw.element.t.CPPTypeStringInNamespaceHalfResolved2(bytesVersion, typeReduction.Type.Arguments[0]))
	case 2:
		if typeReduction.Type.Arguments[0].VariableActsAsConstant {
			return fmt.Sprintf("std::array<%s, %s>", trw.element.t.CPPTypeStringInNamespaceHalfResolved2(bytesVersion, typeReduction.Type.Arguments[1]), typeReduction.Type.Arguments[0].Variable)
		}
		if typeReduction.Type.Arguments[0].Index == NumberConstant {
			return fmt.Sprintf("std::array<%s, %d>", trw.element.t.CPPTypeStringInNamespaceHalfResolved2(bytesVersion, typeReduction.Type.Arguments[1]), typeReduction.Type.Arguments[0].Constant)
		}
		return fmt.Sprintf("std::vector<%s>", trw.element.t.CPPTypeStringInNamespaceHalfResolved2(bytesVersion, typeReduction.Type.Arguments[1]))
	}
	return ""
}

func (trw *TypeRWBrackets) cppDefaultInitializer(halfResolved HalfResolvedArgument, halfResolve bool) string {
	if trw.vectorLike || trw.dynamicSize {
		return ""
	}
	return "{}" // surprisingly, std::array default constructor leaves array uninitialized. Clowns!
}

func (trw *TypeRWBrackets) CPPHasBytesVersion() bool {
	if trw.dictLike {
		return true
	}
	return trw.element.t.trw.CPPHasBytesVersion()
}

func (trw *TypeRWBrackets) CPPTypeResettingCode(bytesVersion bool, val string) string {
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
	if trw.dictLike || trw.dynamicSize || trw.vectorLike {
		return fmt.Sprintf("\t%s.clear();", val)
	}
	return fmt.Sprintf("\t::%s::%sReset(%s);", trw.wr.gen.DetailsCPPNamespace, goGlobalName, val)
}

func (trw *TypeRWBrackets) CPPTypeWritingCode(bytesVersion bool, val string, bare bool, natArgs []string, last bool) string {
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
	return fmt.Sprintf("\tif (!::%s::%sWrite%s(s, %s%s)) { return false; }", trw.wr.gen.DetailsCPPNamespace, goGlobalName, addBare(bare), val, joinWithCommas(natArgs))
}

func (trw *TypeRWBrackets) CPPTypeWritingJsonCode(bytesVersion bool, val string, bare bool, natArgs []string, last bool) string {
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
	return fmt.Sprintf("\tif (!::%s::%sWriteJSON(s, %s%s)) { return false; }", trw.wr.gen.DetailsCPPNamespace, goGlobalName, val, joinWithCommas(natArgs))
}

func (trw *TypeRWBrackets) CPPTypeReadingCode(bytesVersion bool, val string, bare bool, natArgs []string, last bool) string {
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
	return fmt.Sprintf("\tif (!::%s::%sRead%s(s, %s%s)) { return false; }", trw.wr.gen.DetailsCPPNamespace, goGlobalName, addBare(bare), val, joinWithCommas(natArgs))
}

func (trw *TypeRWBrackets) CPPGenerateCode(hpp *strings.Builder, hppInc *DirectIncludesCPP, hppIncFwd *DirectIncludesCPP, hppDet *strings.Builder, hppDetInc *DirectIncludesCPP, cppDet *strings.Builder, cppDetInc *DirectIncludesCPP, bytesVersion bool, forwardDeclaration bool) {
	if forwardDeclaration {
		trw.element.t.trw.CPPGenerateCode(hpp, hppInc, hppIncFwd, hppDet, hppDetInc, cppDet, cppDetInc, bytesVersion, true)
		return
	}

	hppDetCode := `
void %[1]sReset(std::array<%[2]s, %[3]d>& item);

bool %[1]sWriteJSON(std::ostream & s, const std::array<%[2]s, %[3]d>& item%[4]s);
bool %[1]sRead(::basictl::tl_istream & s, std::array<%[2]s, %[3]d>& item%[4]s);
bool %[1]sWrite(::basictl::tl_ostream & s, const std::array<%[2]s, %[3]d>& item%[4]s);
`
	cppCode := `
void %[8]s::%[1]sReset(std::array<%[2]s, %[3]d>& item) {
	for(auto && el : item) {
	%[7]s
	}
}

bool %[8]s::%[1]sWriteJSON(std::ostream &s, const std::array<%[2]s, %[3]d>& item%[4]s) {
	s << "[";
	size_t index = 0;
	for(auto && el : item) {
	%[9]s
		if (index != item.size() - 1) {
			s << ",";
		}
		index++;
	}
	s << "]";
	return true;
}

bool %[8]s::%[1]sRead(::basictl::tl_istream & s, std::array<%[2]s, %[3]d>& item%[4]s) {
	for(auto && el : item) {
	%[5]s
	}
	return true;
}

bool %[8]s::%[1]sWrite(::basictl::tl_ostream & s, const std::array<%[2]s, %[3]d>& item%[4]s) {
	for(const auto & el : item) {
	%[6]s
	}
	return true;
}
`
	// keyTypeString := ""
	// valueTypeString := ""
	// keyFieldName := ""
	// valueFieldName := ""
	switch {
	// TODO - does not work yet
	// REAL TODO CPP
	/*
		case trw.dictLike && !bytesVersion:
				keyTypeString = trw.dictKeyField.t.TypeString(bytesVersion)
				valueTypeString = trw.dictValueField.t.TypeString(bytesVersion)
				keyFieldName = trw.dictKeyField.cppName
				valueFieldName = trw.dictValueField.cppName

				code = `func %[1]sReset(m map[%[10]s]%[11]s) {
					for k := range m {
						delete(m, k)
					}
				}

				func %[1]sRead(r *bytes.Buffer, m *map[%[10]s]%[11]s %[6]s) error {
					var l uint32
					if err := basictl.NatRead(r, &l); err != nil { // No sanity check required for map
						return err
					}
					var data map[%[10]s]%[11]s
					if *m == nil {
						data = make(map[%[10]s]%[11]s, l)
						*m = data
					} else {
						data = *m
						for k := range data {
							delete(data, k)
						}
					}
					for i := 0; i < int(l); i++ {
						var elem %[2]s
						%[14]s
						 data[elem.%[12]s] = elem.%[13]s
					}
					return nil
				}

				func %[1]sWrite(w *bytes.Buffer, m map[%[10]s]%[11]s %[6]s) error {
					basictl.NatWrite(w, uint32(len(m)))
					for key, val := range m {
						elem := %[2]s{%[12]s:key, %[13]s:val}
						%[5]s
					}
					return nil
				}

				`
	*/
	case trw.dictLike && len(trw.element.t.origTL[0].TemplateArguments) == 1:
		trw.CPPGenerateCodeMap(hpp, hppInc, hppIncFwd, hppDet, hppDetInc, cppDet, cppDetInc, bytesVersion, forwardDeclaration)
		return
	case trw.vectorLike:
		hppDetCode = `
void %[1]sReset(std::vector<%[2]s>& item);

bool %[1]sWriteJSON(std::ostream & s, const std::vector<%[2]s>& item%[4]s);
bool %[1]sRead(::basictl::tl_istream & s, std::vector<%[2]s>& item%[4]s);
bool %[1]sWrite(::basictl::tl_ostream & s, const std::vector<%[2]s>& item%[4]s);
`
		cppCode = `
void %[8]s::%[1]sReset(std::vector<%[2]s>& item) {
	item.resize(0); // TODO - unwrap
}

bool %[8]s::%[1]sWriteJSON(std::ostream & s, const std::vector<%[2]s>& item%[4]s) {
	s << "[";
	size_t index = 0;
	for(const auto & el : item) {
	%[9]s
		if (index != item.size() - 1) {
			s << ",";
		}
		index++;
	}
	s << "]";
	return true;
}

bool %[8]s::%[1]sRead(::basictl::tl_istream & s, std::vector<%[2]s>& item%[4]s) {
	uint32_t len = 0;
	if (!s.nat_read(len)) { return false; }
	// TODO - check length sanity
	item.resize(len);
	for(auto && el : item) {
	%[5]s
	}
	return true;
}

bool %[8]s::%[1]sWrite(::basictl::tl_ostream & s, const std::vector<%[2]s>& item%[4]s) {
	if (!s.nat_write(item.size())) { return false; }
	for(const auto & el : item) {
	%[6]s
	}
	return true;
}
`
	case trw.dynamicSize:
		hppDetCode = `
void %[1]sReset(std::vector<%[2]s>& item);

bool %[1]sWriteJSON(std::ostream & s, const std::vector<%[2]s>& item%[4]s);
bool %[1]sRead(::basictl::tl_istream & s, std::vector<%[2]s>& item%[4]s);
bool %[1]sWrite(::basictl::tl_ostream & s, const std::vector<%[2]s>& item%[4]s);
`
		cppCode = `
void %[8]s::%[1]sReset(std::vector<%[2]s>& item) {
	item.resize(0);
}

bool %[8]s::%[1]sWriteJSON(std::ostream & s, const std::vector<%[2]s>& item%[4]s) {
	if (item.size() != nat_n) { return false; }
	s << "[";
	size_t index = 0;
	for(const auto & el : item) {
	%[9]s
		if (index != item.size() - 1) {
			s << ",";
		}
		index++;
	}
	s << "]";
	return true;
}

bool %[8]s::%[1]sRead(::basictl::tl_istream & s, std::vector<%[2]s>& item%[4]s) {
	// TODO - check length sanity
	item.resize(nat_n);
	for(auto && el : item) {
	%[5]s
	}
	return true;
}

bool %[8]s::%[1]sWrite(::basictl::tl_ostream & s, const std::vector<%[2]s>& item%[4]s) {
	if (item.size() != nat_n) { return s.set_error_sequence_length(); }
	for(const auto & el : item) {
	%[6]s
	}
	return true;
}
`
	}
	// TODO remove such giant comments...
	/*
		_ = fmt.Sprintf(code,
			addBytes(trw.goGlobalName, bytesVersion),
			trw.element.t.CPPTypeStringInNamespace(bytesVersion, hppInc, trw.wr.resolvedType.Args[0].T, false),
			trw.size,
			trw.element.t.TypeString(bytesVersion),
			trw.wr.tlTag,
			trw.element.t.TypeReadingCode(bytesVersion, trw.wr.ins,
				"(*vec)[i]",
				trw.element.bare,
				formatNatArgsCPP(nil, trw.element.natArgs),
				false,
				false,
			),
			trw.element.t.TypeWritingCode(bytesVersion, trw.wr.ins,
				"elem",
				trw.element.bare,
				formatNatArgsCPP(nil, trw.element.natArgs),
				false,
				false,
			),
			formatNatArgsDeclCPP(trw.wr.NatParams),
			trw.typeString(bytesVersion, false),
			formatNatArgsCallCPP(trw.wr.NatParams),
			trw.element.t.TypeResettingCode(bytesVersion, trw.wr.ins, "(*vec)[i]", false),
			keyTypeString,
			valueTypeString,
			keyFieldName,
			valueFieldName,
			trw.element.t.TypeReadingCode(bytesVersion, trw.wr.ins,
				"elem",
				trw.element.bare,
				formatNatArgsCPP(nil, trw.element.natArgs),
				false,
				false,
			),
		)
	*/

	if hppDet != nil {
		cppStartNamespace(hppDet, trw.wr.gen.DetailsCPPNamespaceElements)

		hppDet.WriteString(fmt.Sprintf(hppDetCode,
			addBytes(trw.wr.goGlobalName, bytesVersion),
			trw.element.t.CPPTypeStringInNamespace(bytesVersion, &DirectIncludesCPP{ns: map[*TypeRWWrapper]CppIncludeInfo{}}),
			trw.size,
			formatNatArgsDeclCPP(trw.wr.NatParams),
		))

		cppFinishNamespace(hppDet, trw.wr.gen.DetailsCPPNamespaceElements)
	}

	if cppDet != nil {
		tt := trw.element.t.CPPTypeStringInNamespace(bytesVersion, cppDetInc)
		tr := trw.element.t.trw.CPPTypeReadingCode(bytesVersion, "el",
			trw.element.Bare(), formatNatArgsCPP(nil, trw.element.natArgs),
			true)
		tw := trw.element.t.trw.CPPTypeWritingCode(bytesVersion, "el",
			trw.element.Bare(), formatNatArgsCPP(nil, trw.element.natArgs),
			true)
		if tt == "bool" {
			// std::vector<bool> has special value-like reference type
			tr = "\tbool tmp = false;\n\t" + trw.element.t.trw.CPPTypeReadingCode(bytesVersion, "tmp",
				trw.element.Bare(), formatNatArgsCPP(nil, trw.element.natArgs),
				true) + "\n\t\tel = tmp;"
		}
		cppDet.WriteString(fmt.Sprintf(cppCode,
			addBytes(trw.wr.goGlobalName, bytesVersion),
			tt,
			trw.size,
			formatNatArgsDeclCPP(trw.wr.NatParams),
			tr,
			tw,
			trw.element.t.trw.CPPTypeResettingCode(bytesVersion, "el"),
			trw.wr.gen.DetailsCPPNamespace,
			trw.element.t.trw.CPPTypeWritingJsonCode(bytesVersion, "el", false, formatNatArgsCPP(nil, trw.element.natArgs), false),
		))
	}
}

func (trw *TypeRWBrackets) CPPGenerateCodeMap(hpp *strings.Builder, hppInc *DirectIncludesCPP, hppIncFwd *DirectIncludesCPP, hppDet *strings.Builder, hppDetInc *DirectIncludesCPP, cppDet *strings.Builder, cppDetInc *DirectIncludesCPP, bytesVersion bool, forwardDeclaration bool) {
	pairType := trw.element.t.trw.(*TypeRWStruct)

	keyValue := pairType.Fields[0]
	valueType := pairType.Fields[1]

	if hppDet != nil {
		cppStartNamespace(hppDet, trw.wr.gen.DetailsCPPNamespaceElements)

		hppDet.WriteString(fmt.Sprintf(`
void %[1]sReset(std::map<%[5]s, %[2]s>& item);

bool %[1]sWriteJSON(std::ostream & s, const std::map<%[5]s, %[2]s>& item%[4]s);
bool %[1]sRead(::basictl::tl_istream & s, std::map<%[5]s, %[2]s>& item%[4]s);
bool %[1]sWrite(::basictl::tl_ostream & s, const std::map<%[5]s, %[2]s>& item%[4]s);
`,
			addBytes(trw.wr.goGlobalName, bytesVersion),
			valueType.t.CPPTypeStringInNamespace(bytesVersion, hppDetInc),
			"",
			formatNatArgsDeclCPP(trw.wr.NatParams),
			keyValue.t.CPPTypeStringInNamespace(bytesVersion, hppDetInc),
		))

		cppFinishNamespace(hppDet, trw.wr.gen.DetailsCPPNamespaceElements)
	}

	if cppDet != nil {
		cppDet.WriteString(fmt.Sprintf(`
void %[8]s::%[1]sReset(std::map<%[13]s, %[2]s>& item) {
	item.clear(); // TODO - unwrap
}

bool %[8]s::%[1]sWriteJSON(std::ostream & s, const std::map<%[13]s, %[2]s>& item%[4]s) {
	s << "{";
	size_t index = 0;
	for(const auto & el : item) {
	%[9]s
		s << ":";
	%[10]s
		if (index != item.size() - 1) {
			s << ",";
		}
		index++;
	}
	s << "}";
	return true;
}

bool %[8]s::%[1]sRead(::basictl::tl_istream & s, std::map<%[13]s, %[2]s>& item%[4]s) {
	uint32_t len = 0;
	if (!s.nat_read(len)) { return false; }
	item.clear();
	for(uint32_t i = 0; i < len; i++) {
		%[13]s key;
	%[5]s
	%[11]s
	}
	return true;
}

bool %[8]s::%[1]sWrite(::basictl::tl_ostream & s, const std::map<%[13]s, %[2]s>& item%[4]s) {
	if (!s.nat_write(item.size())) { return false; }
	for(const auto & el : item) {
	%[6]s
	%[12]s
	}
	return true;
}
`,
			addBytes(trw.wr.goGlobalName, bytesVersion), // 1
			valueType.t.CPPTypeStringInNamespace(bytesVersion, cppDetInc),
			trw.size, // 3
			formatNatArgsDeclCPP(trw.wr.NatParams),
			keyValue.t.trw.CPPTypeReadingCode(bytesVersion, "key",
				keyValue.Bare(), formatNatArgsCPP(nil, trw.element.natArgs),
				true), // 5
			keyValue.t.trw.CPPTypeWritingCode(bytesVersion, "el.first",
				keyValue.Bare(), formatNatArgsCPP(nil, trw.element.natArgs),
				true),
			"", // 7
			trw.wr.gen.DetailsCPPNamespace,
			keyValue.t.trw.CPPTypeWritingJsonCode(bytesVersion, "el.first",
				keyValue.Bare(), formatNatArgsCPP(nil, trw.element.natArgs),
				false), // 9
			valueType.t.trw.CPPTypeWritingJsonCode(bytesVersion, "el.second",
				valueType.Bare(), formatNatArgsCPP(nil, trw.element.natArgs),
				true),
			valueType.t.trw.CPPTypeReadingCode(bytesVersion, "item[key]",
				valueType.Bare(), formatNatArgsCPP(nil, trw.element.natArgs),
				true), // 11
			valueType.t.trw.CPPTypeWritingCode(bytesVersion, "el.second",
				valueType.Bare(), formatNatArgsCPP(nil, trw.element.natArgs),
				false),
			keyValue.t.CPPTypeStringInNamespace(bytesVersion, cppDetInc), // 13
		))
	}
}
