// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlast

import (
	"hash/crc32"
	"strconv"
	"strings"
	"unicode"
)

func fieldToCrc32(t TypeRef) string {
	var s strings.Builder
	if t.Bare && (len(t.Type.Name) == 0 || !unicode.IsLower(rune(t.Type.Name[0]))) { // TODO - do not set Bare true in parser in this case
		// %goft.QueryObject -> %goft.QueryObject
		// %goft.queryObject -> goft.queryObject
		s.WriteString("%")
	}
	s.WriteString(t.Type.String())
	for _, x := range t.Args {
		s.WriteByte(' ')
		s.WriteString(aotToCrc32(x))
	}
	return s.String()
}

func aotToCrc32(x ArithmeticOrType) string {
	if x.IsArith {
		return strconv.FormatUint(uint64(x.Arith.Res), 10)
	}
	return fieldToCrc32(x.T)
}

func repeatWithScaleToCrc32(x RepeatWithScale) string {
	var s strings.Builder
	if x.ExplicitScale {
		if x.Scale.IsArith {
			s.WriteString(strconv.FormatUint(uint64(x.Scale.Arith.Res), 10))
		} else {
			s.WriteString(x.Scale.Scale)
		}
		s.WriteByte('*')
	}
	s.WriteString("[")
	for _, f := range x.Rep {
		s.WriteByte(' ')
		if f.IsRepeated {
			if f.FieldName != "" {
				s.WriteString(f.FieldName)
				s.WriteByte(':')
			}
			s.WriteString(repeatWithScaleToCrc32(f.ScaleRepeat))
		} else {
			s.WriteString(f.String())
		}
	}
	s.WriteString(" ]")
	return s.String()
}

func (descriptor *Combinator) Crc32() uint32 {
	if descriptor.OriginalDescriptor != nil && descriptor.OriginalDescriptor != descriptor {
		return descriptor.OriginalDescriptor.Crc32()
	}
	if descriptor.Construct.ID != nil {
		return *descriptor.Construct.ID
	}
	var s strings.Builder
	s.WriteString(descriptor.Construct.Name.String())
	s.WriteByte(' ')
	for _, x := range descriptor.TemplateArguments {
		if x.IsNat {
			s.WriteString(x.FieldName)
			s.WriteString(":# ")
		} else {
			s.WriteString(x.FieldName)
			s.WriteString(":Type ")
		}
	}
	if descriptor.Builtin {
		s.WriteString("? ")
	}
	for _, x := range descriptor.Fields {
		if x.FieldName != "" {
			s.WriteString(x.FieldName)
			s.WriteByte(':')
		}
		if x.Mask != nil {
			s.WriteString(x.Mask.String())
		}
		if x.IsRepeated {
			s.WriteString(repeatWithScaleToCrc32(x.ScaleRepeat))
		} else {
			s.WriteString(fieldToCrc32(x.FieldType))
		}
		s.WriteByte(' ')
	}
	s.WriteString("= ")
	if descriptor.IsFunction {
		s.WriteString(fieldToCrc32(descriptor.FuncDecl))
	} else {
		s.WriteString(descriptor.TypeDecl.String())
	}
	// _, err := fmt.Fprintf(os.Stderr, "%s\n%x\n", s.String(), crc32.ChecksumIEEE([]byte(s.String())))
	// if err != nil {
	// 	os.Exit(1)
	// }

	// save for further usage
	descriptor.Construct.ID = new(uint32)
	*descriptor.Construct.ID = crc32.ChecksumIEEE([]byte(s.String()))
	return *descriptor.Construct.ID
}
