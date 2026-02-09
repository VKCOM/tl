// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package gengo

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/vkcom/tl/internal/pure"
	"github.com/vkcom/tl/internal/utils"
)

// During recursive generation, we store wrappers to type when they are needed, so that
// we can generate actual types later, when all references to wrappers are set
// also wrapper stores common information

// TODO remove skipAlias after we start generating go code like we do for C++
type TypeRW interface {
	// methods below are target language independent
	markWantsBytesVersion(visitedNodes map[*TypeRWWrapper]bool)
	markWantsTL2(visitedNodes map[*TypeRWWrapper]bool)
	fillRecursiveUnwrap(visitedNodes map[*TypeRWWrapper]bool)

	FillRecursiveChildren(visitedNodes map[*TypeRWWrapper]int, generic bool)
	ContainsUnion(visitedNodes map[*TypeRWWrapper]bool) bool

	BeforeCodeGenerationStep1() // during first phase, some wr.trw are nil due to recursive types. So we delay some
	BeforeCodeGenerationStep2() // during second phase, union fields recursive bit is set

	// methods below depend on target language
	fillRecursiveChildren(visitedNodes map[*TypeRWWrapper]bool)
	IsDictKeySafe() (isSafe bool, isString bool) // integers and string are safe, other types no
	CanBeBareBoxed() (canBare bool, canBoxed bool)
	typeString2(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, isLocal bool, skipAlias bool) string
	markHasBytesVersion(visitedNodes map[*TypeRWWrapper]bool) bool
	markWriteHasError(visitedNodes map[*TypeRWWrapper]bool) bool
	typeResettingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, ref bool) string
	typeRandomCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, natArgs []string, ref bool) string
	typeWritingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, bare bool, natArgs []string, ref bool, last bool, needError bool) string
	typeReadingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, bare bool, natArgs []string, ref bool, last bool) string
	typeJSONEmptyCondition(bytesVersion bool, val string, ref bool) string
	typeJSONWritingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, natArgs []string, ref bool, needError bool) string
	typeJSONReadingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, jvalue string, val string, natArgs []string, ref bool) string
	typeJSON2ReadingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, jvalue string, val string, natArgs []string, ref bool) string
	typeJSON2ReadingRequiresContext() bool
	GenerateCode(bytesVersion bool, directImports *DirectImports) string

	TypeRWTL2
}

type Field struct {
	// TODO - store pure.Field for properties
	originalName string
	t            *TypeRWWrapper
	bare         bool
	goName       string
	recursive    bool

	fieldMask *pure.ActualNatArg
	BitNumber uint32 // only used when fieldMask != nil

	MaskTL2Bit *int

	natArgs []pure.ActualNatArg
}

func (f *Field) Bare() bool {
	return f.bare
}

func (f *Field) IsAffectingLocalFieldMasks() bool {
	return f.fieldMask != nil && f.fieldMask.IsField()
}

func (f *Field) IsAffectedByExternalFieldMask() bool {
	return f.fieldMask != nil && !f.fieldMask.IsField()
}

func (f *Field) IsTypeDependsFromLocalFields() bool {
	for _, natArg := range f.natArgs {
		if natArg.IsField() {
			return true
		}
	}
	return false
}

func (f *Field) HasNatArguments() bool {
	return len(f.natArgs) != 0
}

func (f *Field) IsLocalIndependent() bool {
	return !f.IsAffectingLocalFieldMasks() && !f.IsTypeDependsFromLocalFields()
}

// do not generate fields, but affect block position and skip during reading
// TL1: never
// TL2: _:X
func (f *Field) IsTL2Omitted() bool {
	return f.originalName == "_"
}

// generate Set/IsSet with external (TL1) or internal (TL1 & TL2) mask/
// must exactly correspond to migrator logic
// TL1: x:fm.b?true x:fm.b?True
// TL2: x:bit
func (f *Field) IsBit() bool {
	if b, ok := f.t.trw.(*TypeRWBool); ok {
		return b.isTL2 && b.isBit
	}
	return f.fieldMask != nil && (f.t.IsTrueType() && (f.t.tlName.String() == "true" || f.t.tlName.String() == "True"))
}

func (f *Field) TL2MaskForOP(op string) string {
	return fmt.Sprintf("tl2mask%d %s %d", *f.MaskTL2Bit/8, op, 1<<(*f.MaskTL2Bit%8))
}

func wrapWithError(wrap bool, wrappedType string) string {
	if !wrap {
		return wrappedType
	}
	return "(_ " + wrappedType + ", err error)"
}

func formatNatArg(fields []Field, arg pure.ActualNatArg) string {
	if arg.IsNumber() {
		return strconv.FormatUint(uint64(arg.Number()), 10)
	}
	if arg.IsField() {
		// tl2 case
		if arg.FieldIndex() < 0 {
			return fmt.Sprintf("item.mask%d", -arg.FieldIndex())
		}
		return "item." + fields[arg.FieldIndex()].goName
	}
	if strings.HasPrefix(arg.Name(), "nat_") {
		panic("aha!") // TODO - remove
	}
	return "nat_" + arg.Name()
}

func formatNatArgs(fields []Field, natArgs []pure.ActualNatArg) []string {
	var result []string
	for _, arg := range natArgs {
		if !arg.IsNumber() {
			result = append(result, formatNatArg(fields, arg))
		}
	}
	return result
}

////for tl2 to tl1 bridge
////in case of formatNatArgs(struct_.Fields, field.natArgs)
//func (f *Field) formatNatArgsOrReturnRandoms(fields []Field, rgName string) []string {
//	result := formatNatArgs(fields, f.natArgs)
//	if len(f.t.NatParams) != len(result) {
//		for i := 0; i < len(f.t.NatParams); i++ {
//			result = append(result, fmt.Sprintf(", basictl.RandomUint(%s)", rgName))
//		}
//	}
//	return result
//}

func formatNatArgsDecl(natArgs []string) string {
	var s strings.Builder
	for _, arg := range natArgs {
		s.WriteString(fmt.Sprintf(",nat_%s uint32", arg))
	}
	return s.String()
}

func formatNatArgsDeclNoComma(natArgs []string) string {
	return strings.TrimPrefix(formatNatArgsDecl(natArgs), ",")
}

// if our fun is declared as ReadBoxed(..., nat_x uint32, nat_y uint32) using formatNatArgsDecl() above,
// and we want to pass arguments to our own function, like Read(..., nat_x, nat_y)
func formatNatArgsDeclCall(natArgs []string) string {
	var s strings.Builder
	for _, arg := range natArgs {
		s.WriteString(fmt.Sprintf(", nat_%s", arg))
	}
	return s.String()
}

// simply adds commas, natArgs are already fully formatted. Difference to strings.Join is leading comma
func joinWithCommas(natArgs []string) string {
	var s strings.Builder
	for _, arg := range natArgs {
		s.WriteString(fmt.Sprintf(", %s", arg))
	}
	return s.String()
}

func addBytes(val string, bytesVersion bool) string {
	return ifString(bytesVersion, val+"Bytes", val)
}

func addBare(bare bool) string {
	return ifString(bare, "", "Boxed")
}

func addAmpersand(ref bool, val string) string {
	return ifString(ref, val, "&"+val)
}

func addAsterisk(ref bool, val string) string {
	return ifString(ref, "*"+val, val)
}

func addAsteriskAndBrackets(ref bool, val string) string {
	return ifString(ref, "(*"+val+")", val)
}

func wrapLast(last bool, code string) string {
	return ifString(last, "return "+code+"", "if err := "+code+"; err != nil { return err }")
}

func wrapLastW(last bool, code string, needError bool) string {
	if needError {
		return ifString(last, "return "+code+"", "if w, err = "+code+"; err != nil { return w, err }")
	} else {
		return ifString(last, "return "+code+"", "w = "+code)
	}
}

func ifString(value bool, t string, f string) string {
	if value {
		return t
	}
	return f
}

func ToUpperFirst(str string) string {
	return utils.ToUpperFirst(str)
}

func ToLowerFirst(str string) string {
	return utils.ToLowerFirst(str)
}

func (f *Field) EnsureRecursive(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace) string {
	if !f.recursive {
		return ""
	}
	myType := f.t.TypeString2(bytesVersion, directImports, ins, false, false)
	// new(X) does not work for some types IIRC
	return fmt.Sprintf(`	if item.%s == nil {
		var value %s
		item.%s = &value
	}
`, f.goName, myType, f.goName)
}

func (f *Field) TypeResettingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace) string {
	resetCode := f.t.TypeResettingCode(bytesVersion, directImports, ins, fmt.Sprintf("item.%s", f.goName), f.recursive)
	if f.recursive {
		return fmt.Sprintf(`	if item.%s != nil {
		%s
	}`, f.goName, resetCode)
	}
	return resetCode
}
