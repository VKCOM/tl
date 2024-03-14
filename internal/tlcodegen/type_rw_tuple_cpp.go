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

func (trw *TypeRWBrackets) CPPFillRecursiveChildren(visitedNodes map[*TypeRWWrapper]bool) {
	if trw.vectorLike || trw.dynamicSize {
		return
	}
	trw.element.t.CPPFillRecursiveChildren(visitedNodes)
}

func (trw *TypeRWBrackets) cppTypeStringInNamespace(bytesVersion bool, hppInc *DirectIncludesCPP) string {
	hppInc.ns[trw.wr.fileName] = struct{}{}
	//if trw.dictLike && !bytesVersion {
	//	TODO - which arguments must map have is very complicated
	//return fmt.Sprintf("std::map<%s, %s>",
	//	trw.dictKeyField.t.CPPTypeStringInNamespace(bytesVersion, hppInc, trw.dictKeyField.resolvedType, halfResolve),
	//	trw.dictValueField.t.CPPTypeStringInNamespace(bytesVersion, hppInc, trw.dictValueField.resolvedType, halfResolve))
	//}
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
	if trw.vectorLike || trw.dynamicSize {
		return fmt.Sprintf("std::vector<%s>", trw.element.t.CPPTypeStringInNamespaceHalfResolved(bytesVersion, hppInc, halfResolved.Args[0]))
	}
	if halfResolved.Args[1].Name != "" {
		return fmt.Sprintf("std::array<%s, %s>", trw.element.t.CPPTypeStringInNamespaceHalfResolved(bytesVersion, hppInc, halfResolved.Args[0]), halfResolved.Args[1].Name)
	}
	return fmt.Sprintf("std::array<%s, %d>", trw.element.t.CPPTypeStringInNamespaceHalfResolved(bytesVersion, hppInc, halfResolved.Args[0]), trw.size)
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
	return fmt.Sprintf("\t::%s::%sWrite%s(s, %s%s);", trw.wr.gen.DetailsCPPNamespace, goGlobalName, addBare(bare), val, joinWithCommas(natArgs))
}

func (trw *TypeRWBrackets) CPPTypeReadingCode(bytesVersion bool, val string, bare bool, natArgs []string, last bool) string {
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
	return fmt.Sprintf("\t::%s::%sRead%s(s, %s%s);", trw.wr.gen.DetailsCPPNamespace, goGlobalName, addBare(bare), val, joinWithCommas(natArgs))
}

func (trw *TypeRWBrackets) CPPGenerateCode(hpp *strings.Builder, hppInc *DirectIncludesCPP, hppIncFwd *DirectIncludesCPP, hppDet *strings.Builder, hppDetInc *DirectIncludesCPP, cppDet *strings.Builder, cppDetInc *DirectIncludesCPP, bytesVersion bool, forwardDeclaration bool) {
	cppStartNamespace(hppDet, trw.wr.gen.DetailsCPPNamespaceElements)

	hppDetCode := `
void %[1]sReset(std::array<%[2]s, %[3]d>& item);
void %[1]sRead(::basictl::tl_istream & s, std::array<%[2]s, %[3]d>& item%[4]s);
void %[1]sWrite(::basictl::tl_ostream & s, const std::array<%[2]s, %[3]d>& item%[4]s);
`
	cppCode := `
void %[8]s::%[1]sReset(std::array<%[2]s, %[3]d>& item) {
	for(auto && el : item) {
	%[7]s
	}
}

void %[8]s::%[1]sRead(::basictl::tl_istream & s, std::array<%[2]s, %[3]d>& item%[4]s) {
	for(auto && el : item) {
	%[5]s
	}
}

void %[8]s::%[1]sWrite(::basictl::tl_ostream & s, const std::array<%[2]s, %[3]d>& item%[4]s) {
	for(const auto & el : item) {
	%[6]s
	}
}
`
	// keyTypeString := ""
	// valueTypeString := ""
	// keyFieldName := ""
	// valueFieldName := ""
	switch {
	// TODO - does not work yet
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
	case trw.vectorLike:
		hppDetCode = `
void %[1]sReset(std::vector<%[2]s>& item);
void %[1]sRead(::basictl::tl_istream & s, std::vector<%[2]s>& item%[4]s);
void %[1]sWrite(::basictl::tl_ostream & s, const std::vector<%[2]s>& item%[4]s);
`
		cppCode = `
void %[8]s::%[1]sReset(std::vector<%[2]s>& item) {
	item.resize(0); // TODO - unwrap
}

void %[8]s::%[1]sRead(::basictl::tl_istream & s, std::vector<%[2]s>& item%[4]s) {
	auto len = s.nat_read();
	// TODO - check length sanity
	item.resize(len);
	for(auto && el : item) {
	%[5]s
	}
}

void %[8]s::%[1]sWrite(::basictl::tl_ostream & s, const std::vector<%[2]s>& item%[4]s) {
	s.nat_write(item.size());
	for(const auto & el : item) {
	%[6]s
	}
}
`
	case trw.dynamicSize:
		hppDetCode = `
void %[1]sReset(std::vector<%[2]s>& item);
void %[1]sRead(::basictl::tl_istream & s, std::vector<%[2]s>& item%[4]s);
void %[1]sWrite(::basictl::tl_ostream & s, const std::vector<%[2]s>& item%[4]s);
`
		cppCode = `
void %[8]s::%[1]sReset(std::vector<%[2]s>& item) {
	item.resize(0);
}

void %[8]s::%[1]sRead(::basictl::tl_istream & s, std::vector<%[2]s>& item%[4]s) {
	// TODO - check length sanity
	item.resize(nat_n);
	for(auto && el : item) {
	%[5]s
	}
}

void %[8]s::%[1]sWrite(::basictl::tl_ostream & s, const std::vector<%[2]s>& item%[4]s) {
	if (item.size() != nat_n)
		::basictl::throwSequenceLengthWrong();
	for(const auto & el : item) {
	%[6]s
	}
}
`
	}
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

	hppDet.WriteString(fmt.Sprintf(hppDetCode,
		addBytes(trw.wr.goGlobalName, bytesVersion),
		trw.element.t.CPPTypeStringInNamespace(bytesVersion, cppDetInc),
		trw.size,
		formatNatArgsDeclCPP(trw.wr.NatParams),
	))
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
	))
	cppFinishNamespace(hppDet, trw.wr.gen.DetailsCPPNamespaceElements)
}
