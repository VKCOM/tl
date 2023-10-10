// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlast

import (
	"fmt"
	"strconv"
	"strings"
)

const ContactAuthorsString = "please check/create issue with example https://github.com/VKCOM/tl/issues" // TODO move to better place

type Name struct {
	Namespace string
	Name      string
}

type Modifier struct {
	Name string
	PR   PositionRange
}

type Constructor struct {
	Name Name
	ID   *uint32 // TODO - uint32 + bool IsExplicit

	NamePR PositionRange
	IDPR   PositionRange
}

type TemplateArgument struct {
	FieldName string
	IsNat     bool
	PR        PositionRange // TODO - split into type and name
}

type TypeDeclaration struct {
	Name      Name
	Arguments []string

	PR          PositionRange
	NamePR      PositionRange
	ArgumentsPR []PositionRange
}

type Arithmetic struct {
	Nums []uint32
	Res  uint32
}

type ArithmeticOrType struct {
	IsArith bool
	Arith   Arithmetic
	T       TypeRef // PR of T can also be used for Arith
}

type ScaleFactor struct {
	IsArith bool
	Arith   Arithmetic
	Scale   string

	PR PositionRange
}

type RepeatWithScale struct {
	ExplicitScale bool
	Scale         ScaleFactor
	Rep           []Field

	PR PositionRange
}

type FieldMask struct {
	MaskName  string
	BitNumber uint32

	PRName PositionRange
	PRBits PositionRange
}

type TypeRef struct { // due to complexity, parsing of TypeRef defined in separate tlparser_typeref.go file
	Type Name
	Args []ArithmeticOrType
	Bare bool

	PR     PositionRange
	PRArgs PositionRange
}

type Field struct {
	FieldName string
	Mask      *FieldMask
	Excl      bool

	IsRepeated  bool
	ScaleRepeat RepeatWithScale
	FieldType   TypeRef

	PR     PositionRange
	PRName PositionRange

	CommentBefore string // comment before field
	CommentRight  string // comment to the right of field
}

type Combinator struct {
	Builtin           bool
	IsFunction        bool
	Modifiers         []Modifier
	Construct         Constructor
	TemplateArguments []TemplateArgument
	Fields            []Field
	TypeDecl          TypeDeclaration
	FuncDecl          TypeRef

	OriginalDescriptor *Combinator // hack - if some replacements were made, original descriptor is saved here
	OriginalOrderIndex int         // declaration order index (needed in TLO generation)

	TemplateArgumentsPR PositionRange // especially useful when 0 arguments
	PR                  PositionRange

	CommentBefore string // comment before combinator
	CommentRight  string // comment to the right of combinator
}

type TL []*Combinator

// only trivial methods below, parsing is in tlparser_code.go and tlparser_typeref.go files

func (n Name) IsEmpty() bool {
	return n.Namespace == "" && n.Name == ""
}

func (n Name) String() string {
	if n.Namespace != "" {
		return n.Namespace + "." + n.Name
	}
	return n.Name
}

func (c Constructor) String() string {
	var s strings.Builder
	s.WriteString(c.Name.String())
	if c.ID != nil && *c.ID != 0 {
		s.WriteByte('#')
		s.WriteString(fmt.Sprintf("%08x", *c.ID))
	}
	return s.String()
}

func (ta TemplateArgument) String() string {
	if ta.IsNat {
		return fmt.Sprintf("{%v:#}", ta.FieldName)
	}
	return fmt.Sprintf("{%v:Type}", ta.FieldName)
}

func (t TypeRef) IsEmpty() bool {
	return t.Type.IsEmpty()
}

func (t TypeRef) String() string {
	var s strings.Builder
	if t.Bare {
		s.WriteByte('%')
	}
	if len(t.Args) != 0 {
		s.WriteByte('(')
		s.WriteString(t.Type.String())
		for _, x := range t.Args {
			s.WriteByte(' ')
			s.WriteString(x.String())
		}
		s.WriteByte(')')
		return s.String()
	}
	s.WriteString(t.Type.String())
	return s.String()
}

func (a Arithmetic) IsEmpty() bool {
	return len(a.Nums) == 0
}

func (a Arithmetic) String() string {
	var s strings.Builder
	for i, x := range a.Nums {
		s.WriteString(strconv.FormatUint(uint64(x), 10))
		if i != len(a.Nums)-1 {
			s.WriteString(" + ")
		}
	}
	return s.String()
}

func (aot ArithmeticOrType) String() string {
	if aot.IsArith {
		return aot.Arith.String()
	}
	return aot.T.String()
}

func (sf ScaleFactor) IsEmpty() bool {
	return sf.Arith.IsEmpty() && sf.Scale == ""
}

func (sf ScaleFactor) String() string {
	if sf.IsArith {
		return "(" + sf.Arith.String() + ")"
	}
	return sf.Scale
}

func (rws RepeatWithScale) String() string {
	var s strings.Builder
	if rws.ExplicitScale {
		s.WriteString(rws.Scale.String())
		s.WriteByte('*')
	}
	s.WriteByte('[')
	for i, x := range rws.Rep[0:] {
		if i != 0 {
			s.WriteByte(' ')
		}
		s.WriteString(x.String())
	}
	s.WriteByte(']')
	return s.String()
}

func (fm FieldMask) String() string {
	return fmt.Sprintf("%v.%v?", fm.MaskName, fm.BitNumber)
}

func (f Field) String() string {
	var s strings.Builder
	if f.FieldName != "" {
		s.WriteString(f.FieldName)
		s.WriteByte(':')
	}
	if f.Mask != nil {
		s.WriteString(f.Mask.String())
	}
	if f.Excl {
		s.WriteByte('!')
	}
	if f.IsRepeated {
		s.WriteString(f.ScaleRepeat.String())
	} else {
		s.WriteString(f.FieldType.String())
	}
	return s.String()
}

func (d TypeDeclaration) String() string {
	var s strings.Builder
	s.WriteString(d.Name.String())
	for _, x := range d.Arguments {
		s.WriteByte(' ')
		s.WriteString(x)
	}
	return s.String()
}

func (descriptor Combinator) String() string {
	var s strings.Builder
	for _, x := range descriptor.Modifiers {
		s.WriteString(x.Name)
		s.WriteByte(' ')
	}
	s.WriteString(descriptor.Construct.String())
	s.WriteByte(' ')
	for _, x := range descriptor.TemplateArguments {
		s.WriteString(x.String())
		s.WriteByte(' ')
	}
	if descriptor.Builtin {
		s.WriteString("? ")
	} else {
		for _, x := range descriptor.Fields {
			s.WriteString(x.String())
			s.WriteByte(' ')
		}
	}
	s.WriteString("= ")
	if descriptor.IsFunction {
		s.WriteString(descriptor.FuncDecl.String())
	} else {
		s.WriteString(descriptor.TypeDecl.String())
	}
	s.WriteByte(';')
	return s.String()
}

func (tl TL) String() string {
	var s strings.Builder
	functionSection := false

	for _, x := range tl {
		if x.IsFunction && !functionSection {
			s.WriteString(functionsSectionString + "\n")
			functionSection = true
		}
		if !x.IsFunction && functionSection {
			s.WriteString(typesSectionString + "\n")
			functionSection = false
		}
		s.WriteString(x.String())
		s.WriteString("\n")
	}
	return s.String()
}
