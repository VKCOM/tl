// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlast

import "strings"

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
	Modifiers         []Modifier // TODO - rename to annotations
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

func (t TypeRef) IsEmpty() bool {
	return t.Type.IsEmpty()
}

func (a Arithmetic) IsEmpty() bool {
	return len(a.Nums) == 0
}

func (sf ScaleFactor) IsEmpty() bool {
	return sf.Arith.IsEmpty() && sf.Scale == ""
}

// we support windows-style line separators
func SplitMultilineComment(comment string) []string {
	rep := strings.ReplaceAll(comment, "\r\n", "\n")
	return strings.Split(rep, "\n")
}
