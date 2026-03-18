// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package gengo

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"

	"github.com/VKCOM/tl/internal/pure"
	"github.com/VKCOM/tl/internal/utils"
)

// During recursive generation, we store wrappers to type when they are needed, so that
// we can generate actual types later, when all references to wrappers are set
// also wrapper stores common information

// TODO remove skipAlias after we start generating go code like we do for C++
type TypeRW interface {
	// methods below are target language independent
	markWantsBytesVersion(visitedNodes map[*TypeRWWrapper]bool)
	markHasBytesVersion(visitedNodes map[*TypeRWWrapper]bool) bool
	markHasRepairMasks(visitedNodes map[*TypeRWWrapper]bool) bool
	markWriteHasError(visitedNodes map[*TypeRWWrapper]bool) bool

	FillRecursiveChildren(visitedNodes map[*TypeRWWrapper]int, generic bool)
	ContainsUnion(visitedNodes map[*TypeRWWrapper]bool) bool

	BeforeCodeGenerationStep1() // during first phase, some wr.trw are nil due to recursive types. So we delay some
	BeforeCodeGenerationStep2() // during second phase, union fields recursive bit is set

	fillRecursiveChildren(visitedNodes map[*TypeRWWrapper]bool)
	typeString2(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, isLocal bool, skipAlias bool) string
	typeResettingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, ref bool) string
	typeRandomCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, natArgs []string, ref bool) string
	typeRepairMasksCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, natArgs []string, ref bool) string
	typeWritingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, bare bool, natArgs []string, ref bool, last bool, needError bool) string
	typeReadingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, bare bool, natArgs []string, ref bool, last bool) string
	typeJSONEmptyCondition(bytesVersion bool, val string, ref bool) string
	typeJSONWritingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, val string, natArgs []string, ref bool, needError bool) string
	typeJSON2ReadingCode(bytesVersion bool, directImports *DirectImports, ins *InternalNamespace, jvalue string, val string, natArgs []string, ref bool) string
	typeJSON2ReadingRequiresContext() bool
	GenerateCode(bytesVersion bool, directImports *DirectImports) string

	TypeRWTL2
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
		if arg.FieldIndex() < 0 {
			panic("should be never, this was used by october kernel to pass TL2 masks")
		}
		return "item." + fields[arg.FieldIndex()].goName
	}
	return "nat_" + arg.Name()
}

func formatNatArgs(fields []Field, natArgs []pure.ActualNatArg) []string {
	var result []string
	for _, arg := range natArgs {
		result = append(result, formatNatArg(fields, arg))
	}
	return result
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

func addBytesLower(val string, bytesVersion bool) string {
	return ifString(bytesVersion, val+"_bytes", val)
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

// must be deterministic.
// we use it to mark sizes during TL2 serialization. tlTag does not work,
// because it resets to 0 for most types during TL2 migration
func someHash(str string) uint32 {
	h := sha256.Sum256([]byte(str))
	return binary.LittleEndian.Uint32(h[:])
}

func printComments(before string, right string) string {
	result := ""
	if before != "" {
		result += before + "\n"
	}
	if right != "" {
		result += right + "\n"
	}
	return result
}

func printCommentsType(pureType pure.TypeInstance) string {
	return printComments(pureType.Common().CommentBefore(), pureType.Common().CommentRight())
}

func printCommentsField(field Field) string {
	return printComments(field.pureField.CommentBefore(), field.pureField.CommentRight())
}

func ToUpperFirst(str string) string {
	return utils.ToUpperFirst(str)
}
