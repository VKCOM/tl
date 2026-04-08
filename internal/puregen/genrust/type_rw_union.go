// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package genrust

import (
	"fmt"
	"strings"

	"github.com/VKCOM/tl/internal/pure"
	"github.com/VKCOM/tl/internal/puregen"
	"github.com/VKCOM/tl/internal/tlcodegen/codecreator"
)

type Variant struct {
	t           *TypeRWWrapper
	variantName string
	goName      string
	recursive   bool
}

func (f *Variant) EnsureRecursive(bytesVersion bool, directImports *DirectImports) string {
	if !f.recursive {
		return ""
	}
	myType := f.t.TypeString2(bytesVersion, directImports, false, false)
	return fmt.Sprintf(`	if item.value%s == nil { item.value%s = new(%s) }
`, f.goName, f.goName, myType)
}

type TypeRWUnion struct {
	wr       *TypeRWWrapper
	pureType *pure.TypeInstanceUnion
	Fields   []Variant
	IsEnum   bool

	fieldsDec puregen.Deconflicter // TODO - add all generated methods here
}

func (trw *TypeRWUnion) ElementNatArgs() []pure.ActualNatArg {
	return trw.pureType.ElementNatArgs()
}

var _ TypeRW = &TypeRWUnion{}

func (trw *TypeRWUnion) typeString2(bytesVersion bool, directImports *DirectImports, isLocal bool, skipAlias bool) string {
	if isLocal {
		return addBytes(trw.wr.goLocalName, bytesVersion)
	}
	return "crate::types::" + trw.wr.goGlobalName + "::" + addBytes(trw.wr.goGlobalName, bytesVersion)
}

func (trw *TypeRWUnion) markHasBytesVersion(visitedNodes map[*TypeRWWrapper]bool) bool {
	result := false
	for _, f := range trw.Fields {
		result = result || f.t.MarkHasBytesVersion(visitedNodes)
	}
	return result
}

func (trw *TypeRWUnion) markHasRepairMasks(visitedNodes map[*TypeRWWrapper]bool) bool {
	result := false
	for _, f := range trw.Fields {
		result = result || f.t.MarkHasRepairMasks(visitedNodes)
	}
	return result
}

func (trw *TypeRWUnion) markWriteHasError(visitedNodes map[*TypeRWWrapper]bool) bool {
	result := false
	for _, f := range trw.Fields {
		result = result || f.t.MarkWriteHasError(visitedNodes)
	}
	return result
}

func (trw *TypeRWUnion) markWantsBytesVersion(visitedNodes map[*TypeRWWrapper]bool) {
	for _, f := range trw.Fields {
		f.t.MarkWantsBytesVersion(visitedNodes)
	}
}

func (trw *TypeRWUnion) FillRecursiveChildren(visitedNodes map[*TypeRWWrapper]int, generic bool) {
	if visitedNodes[trw.wr] != 0 {
		return
	}
	visitedNodes[trw.wr] = 1
	for _, f := range trw.Fields {
		if f.recursive {
			continue
		}
		f.t.trw.FillRecursiveChildren(visitedNodes, generic)
	}
	visitedNodes[trw.wr] = 2
}

func (trw *TypeRWUnion) ContainsUnion(visitedNodes map[*TypeRWWrapper]bool) bool {
	return true
}

func (trw *TypeRWUnion) BeforeCodeGenerationStep1() {
}

func (trw *TypeRWUnion) BeforeCodeGenerationStep2() {
	for i, f := range trw.Fields {
		visitedNodes := map[*TypeRWWrapper]bool{}
		f.t.trw.fillRecursiveChildren(visitedNodes)
		trw.Fields[i].recursive = visitedNodes[trw.wr]
	}
}

func (trw *TypeRWUnion) fillRecursiveChildren(visitedNodes map[*TypeRWWrapper]bool) {
}

func (trw *TypeRWUnion) typeResettingCode(cc *codecreator.RustCodeCreator, bytesVersion bool, directImports *DirectImports, val string, ref bool) {
	cc.AddLinef("%s.Reset()", val)
}

func (trw *TypeRWUnion) typeRandomCode(bytesVersion bool, directImports *DirectImports, val string, natArgs []string, ref bool) string {
	return fmt.Sprintf("%s.FillRandom(rg%s)", val, joinWithCommas(natArgs))
}

func (trw *TypeRWUnion) typeRepairMasksCode(bytesVersion bool, directImports *DirectImports, val string, natArgs []string, ref bool) string {
	return fmt.Sprintf("%s.RepairMasks(%s)", val, strings.Join(natArgs, ","))
}

func (trw *TypeRWUnion) typeWritingCode(bytesVersion bool, directImports *DirectImports, val string, bare bool, natArgs []string, ref bool, last bool, needError bool) string {
	if bare && !trw.wr.OriginTL2() {
		panic(fmt.Errorf("trying to write bare union %s, please report TL which caused this", trw.wr.pureType.CanonicalName()))
	}
	return wrapLastW(last, fmt.Sprintf("%s.WriteTL1%s(w %s %s)", val, addBare(false), joinWithCommas(natArgs), trw.wr.fetcherCall()), needError)
}

func (trw *TypeRWUnion) typeReadingCode(cc *codecreator.RustCodeCreator, bytesVersion bool, directImports *DirectImports, val string, bare bool, natArgs []string, ref bool) {
	if bare && !trw.wr.OriginTL2() {
		panic(fmt.Errorf("trying to read bare union %s, please report TL which caused this", trw.wr.pureType.CanonicalName()))
	}
	cc.AddLinef("crate::types::%s::read_tl1_boxed(&mut %s, buf%s%s)?;", trw.wr.goGlobalName, addAsterisk(ref, val), joinWithCommas(natArgs), trw.wr.fetcherCall())
}

func (trw *TypeRWUnion) typeJSONEmptyCondition(bytesVersion bool, val string, ref bool) string {
	return ""
}

func (trw *TypeRWUnion) typeJSONWritingCode(bytesVersion bool, directImports *DirectImports, val string, natArgs []string, ref bool, needError bool) string {
	if needError {
		return fmt.Sprintf("if w, err = %s.WriteJSONOpt(jctx, w %s); err != nil { return w, err }", val, joinWithCommas(natArgs))
	} else {
		return fmt.Sprintf("w = %s.WriteJSONOpt(jctx, w %s)", val, joinWithCommas(natArgs))
	}
}

func (trw *TypeRWUnion) typeJSON2ReadingCode(bytesVersion bool, directImports *DirectImports, jvalue string, val string, natArgs []string, ref bool) string {
	return fmt.Sprintf("if err := %s.ReadJSONGeneral(jctx, %s %s); err != nil { return err }", val, jvalue, joinWithCommas(natArgs))
}

// TODO - remove with long adapters
func (trw *TypeRWUnion) HasShortFieldCollision(wr *TypeRWWrapper) bool {
	//messages.peerId peer_id:int = messages.ChatId;
	//messagesLong.peerId peer_id:long = messages.ChatId;
	//
	//messages.globalChatId#07a5893d chat_id:long = messages.ChatId;
	//messagesLong.globalChatId global_id:messagesLong.GlobalId = messages.ChatId;

	for _, field := range trw.Fields {
		if field.t == wr {
			return true
		}
	}
	return false
}

// TODO - move to separate file
func (trw *TypeRWUnion) GenerateCode(bytesVersion bool, directImports *DirectImports) string {
	cc := codecreator.NewRustCodeCreator()
	printCommentsType(trw.wr.pureType)
	cc.AddLinef("use basictl::TLRead as _;")
	cc.AddEmptyLine()
	cc.AddLinef("#[derive(Default, Debug)]")

	cc.AddLinef("pub enum %s {", trw.wr.goGlobalName)
	cc.AddBlock(func() {
		cc.AddLinef("#[default]")
		for _, field := range trw.Fields {
			if field.t.IsTrueType() {
				cc.AddLinef("%s,", field.goName)
			} else {
				fieldTypeString := field.t.TypeString2(bytesVersion, directImports, false, false)
				cc.AddLinef("%s(%s%s),", field.goName, ifString(field.recursive, "*", ""), fieldTypeString)
			}
		}
	})
	cc.AddLinef("}")
	if len(trw.wr.NatParams()) == 0 {
		cc.AddEmptyLine()
		cc.AddLinef("impl %s {", trw.wr.goGlobalName)
		cc.AddBlock(func() {
			// comment for now to avoid confusion
			//if trw.wr.HasTL2() && len(trw.wr.NatParams()) == 0 && !trw.wr.HasFetcher() {
			// for interface requirements for TL2 Type, also for tests
			//cc.AddLinef("pub fn read_tl1<B: bytes::Buf + Copy>(&mut self, buf: &mut B) -> basictl::Result<()> {")
			//cc.AddBlock(func() {
			//	cc.AddLinef("self::read_tl1_boxed(self, buf)")
			//})
			//cc.AddLinef("}")
			//}
			cc.AddLinef("pub fn read_tl1_boxed<B: bytes::Buf + Copy>(&mut self, buf: &mut B) -> basictl::Result<()> {")
			cc.AddBlock(func() {
				cc.AddLinef("self::read_tl1_boxed(self, buf)")
			})
			cc.AddLinef("}")
			for _, field := range trw.Fields {
				cc.AddLinef("pub(crate) fn reset_to_%s(&mut self) {", field.goName)
				cc.AddBlock(func() {
					if field.t.IsTrueType() {
						cc.AddLinef("*self = Self::%s;", field.goName)
					} else {
						cc.IfElse(fmt.Sprintf("let Self::%s(subValue) = self", field.goName), func() {
							field.t.TypeResettingCode(cc, bytesVersion, directImports, "subValue", true)
						}, func() {
							fieldTypeString := field.t.TypeString2(bytesVersion, directImports, false, false)
							cc.AddLinef("*self = Self::%s(%s::default());", field.goName, fieldTypeString)
						})
					}
				})
				cc.AddLinef("}")
			}
		})
		cc.AddLinef("}")
	}

	natArgsDecl := trw.wr.formatNatArgsDecl()
	cc.AddEmptyLine()
	cc.AddLinef("pub(crate) fn read_tl1_boxed<B: bytes::Buf + Copy>(value: &mut %s, buf: &mut B%s) -> basictl::Result<()> {", trw.wr.goGlobalName, natArgsDecl)
	cc.AddBlock(func() {
		if trw.wr.OriginTL2() {
			cc.AddLinef(`Err(basictl::Error::NoTL1("%s")`, trw.wr.pureType.CanonicalName())
			return
		}
		cc.AddLinef("match buf.read_u32()? {")
		cc.AddBlock(func() {
			for _, field := range trw.Fields {
				cc.AddLinef("0x%08x => {", field.t.TLTag())
				cc.AddBlock(func() {
					cc.AddLinef("value.reset_to_%s();", field.goName)
					if !field.t.IsTrueType() {
						cc.If(fmt.Sprintf("let %s::%s(subValue) = value", trw.wr.goGlobalName, field.goName), func() {
							natArgs := trw.wr.formatNatArgs(nil, trw.ElementNatArgs())
							field.t.TypeReadingCode(cc, bytesVersion, directImports, "subValue", true, natArgs, true)
						})
					}
					cc.AddLinef("Ok(())")
				})
				cc.AddLinef("}")
			}
			cc.AddLinef("other => Err(basictl::Error::UnexpectedMagic(other)),") // TODO - for type
		})
		cc.AddLinef("}")

		//switch tag {
		//	{%- for i, field := range union.Fields -%}
		//case {%s= fmt.Sprintf("0x%08x", field.t.TLTag()) %}:
		//	item.index = {%d i %}
		//	{%- if field.t.IsTrueType() -%}
		//	return w, nil
		//	{%- continue -%}
		//	{%- endif -%}
		//	{%s= field.EnsureRecursive(bytesVersion, directImports, union.wr.ins) -%}
		//	{%s=  %}
		//	{%- endfor -%}
		//default:
		//	return w, {%s= union.wr.gen.InternalPrefix()%}ErrorInvalidUnionTag({%q= tlName %}, tag)
		//}

		//for _, field := range trw.Fields {
		//	if field.IsBit() {
		//		if !field.Bare() { // special rare case for TL1, let optimizer combine 2 expressions
		//			arg := trw.wr.formatNatArg(trw.Fields, *field.FieldMask())
		//			cc.If(fmt.Sprintf("%s & (1 << %v) != 0", arg, field.BitNumber()), func() {
		//				cc.AddLinef("buf.read_exact_tag(0x%08x)?;", field.t.TLTag())
		//			})
		//		}
		//		continue
		//	}
		//	fieldAccess, fieldAsterisk := field.FieldAccess(bytesVersion, directImports)
		//	bodyFunc := func() {
		//		cc.AddLinef("%s", field.EnsureRecursive(bytesVersion, directImports))
		//		field.t.TypeReadingCode(cc, bytesVersion, directImports, fieldAccess, field.Bare(), trw.wr.formatNatArgs(trw.Fields, field.NatArgs()), fieldAsterisk)
		//	}
		//	if field.FieldMask() != nil {
		//		arg := trw.wr.formatNatArg(trw.Fields, *field.FieldMask())
		//		cc.IfElse(fmt.Sprintf("%s & (1 << %v) != 0", arg, field.BitNumber()), bodyFunc, func() {
		//			field.TypeResettingCode(cc, bytesVersion, directImports)
		//		})
		//	} else {
		//		bodyFunc()
		//	}
		//}
		//cc.AddLinef("Ok(())")
	})
	cc.AddLinef("}")
	return cc.Text()
}
