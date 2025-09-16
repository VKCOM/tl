// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlast

import (
	"fmt"
	"hash/fnv"
	"strconv"
	"strings"
)

func hash(s string) string {
	return fmt.Sprintf("\"%x\"", hashUint32(s))
}

func hashUint32(s string) uint32 {
	h := fnv.New32a()
	_, _ = h.Write([]byte(s))
	return h.Sum32()
}

func createNode(s *strings.Builder, parent, child, nodeName string, color string) {
	_, _ = fmt.Fprintf(s, "%s [shape=record;color=%s;label=%q];\n", child, color, nodeName)
	_, _ = fmt.Fprintf(s, "%s -> %s;\n", parent, child)
}

func (tl TL) ToDot() string {
	var s, t, f strings.Builder
	s.WriteString("digraph TL { splines=\"ortho\";\nordering=in;\ngraph[ordering=in];\n")
	tlID := fmt.Sprintf("\"%pTL\"", tl)
	typesID := hash("types")
	functionsID := hash("functions")
	s.WriteString(tlID + " [shape=record;color=white;label=\"TL\"];\n")
	createNode(&t, tlID, typesID, "Types", "white")
	createNode(&f, tlID, functionsID, "Functions", "white")
	t.WriteString("subgraph types { ordering=in;\ngraph[ordering=in];\n")
	f.WriteString("subgraph functions { ordering=in;\ngraph[ordering=in];\n")
	for i := 0; i < len(tl); i++ {
		x := tl[i]
		id := fmt.Sprintf("\"%pCombinator\"", &x)
		if x.IsFunction {
			createNode(&f, functionsID, id, "Function", "white")
			x.toDot(&f)
		} else {
			createNode(&f, typesID, id, "Type", "white")
			x.toDot(&t)
		}
	}
	t.WriteString("}\n")
	f.WriteString("}\n")
	s.WriteString(t.String())
	s.WriteString(f.String())
	s.WriteString("}\n")
	return s.String()
}

func (descriptor *Combinator) toDot(s *strings.Builder) {
	parent := fmt.Sprintf("\"%pCombinator\"", descriptor)

	if descriptor.IsFunction {
		modifiersID := fmt.Sprintf("\"%pModifiers\"", &descriptor.Modifiers)
		createNode(s, parent, modifiersID, "Modifiers", "white")
		for _, m := range descriptor.Modifiers {
			modID := hash(modifiersID + m.Name)
			createNode(s, modifiersID, modID, m.Name, "yellow")
		}
	}

	constructID := fmt.Sprintf("\"%pConstructor\"", &descriptor.Construct)
	createNode(s, parent, constructID, "Constructor", "white")
	descriptor.Construct.toDot(s)

	if descriptor.TemplateArguments != nil {
		templateFieldID := fmt.Sprintf("\"%pTemplateArguments\"", &descriptor.TemplateArguments)
		createNode(s, parent, templateFieldID, "Template Field", "white")
		toDotTemplateFields(s, &descriptor.TemplateArguments)
	}

	if descriptor.Fields != nil {
		fieldID := fmt.Sprintf("\"%pFields\"", &descriptor.Fields)
		createNode(s, parent, fieldID, "Fields", "white")
		toDotFields(s, &descriptor.Fields)
	}

	if descriptor.IsFunction {
		funcDeclID := fmt.Sprintf("\"%pTypeRef\"", &descriptor.FuncDecl)
		createNode(s, parent, funcDeclID, "Func Declaration", "white")
		descriptor.FuncDecl.toDot(s)
	} else {
		typeDeclID := fmt.Sprintf("\"%pTypeDeclaration\"", &descriptor.TypeDecl)
		createNode(s, parent, typeDeclID, "Type Declaration", "white")
		descriptor.TypeDecl.toDot(s)
	}
}

func toDotTemplateFields(s *strings.Builder, tfs *[]TemplateArgument) {
	parent := fmt.Sprintf("\"%pTemplateArguments\"", tfs)
	for i := 0; i < len(*tfs); i++ {
		tfID := fmt.Sprintf("\"%pTemplateArgument\"", &(*tfs)[i])
		createNode(s, parent, tfID, strconv.Itoa(i), "white")
		(*tfs)[i].toDot(s)
	}
}

func toDotFields(s *strings.Builder, fs *[]Field) {
	parent := fmt.Sprintf("\"%pFields\"", fs)
	for i := 0; i < len(*fs); i++ {
		fID := fmt.Sprintf("\"%pField\"", &(*fs)[i])
		createNode(s, parent, fID, strconv.Itoa(i), "white")
		(*fs)[i].toDot(s)
	}
}

func (c *Constructor) toDot(s *strings.Builder) {
	parent := fmt.Sprintf("\"%pConstructor\"", c)
	nameID := fmt.Sprintf("\"%pName\"", &c.Name)
	createNode(s, parent, nameID, "Name", "white")
	c.Name.toDot(s)

	if c.ID == nil {
		return
	}
	idID := fmt.Sprintf("\"%p\"", &c.ID)
	createNode(s, parent, idID, "tag", "white")
	id := hash(idID)
	createNode(s, idID, id, fmt.Sprintf("%08x", *c.ID), "cyan")
}

func (n *Name) toDot(s *strings.Builder) {
	parent := fmt.Sprintf("\"%pName\"", n)
	id := hash(parent + n.String())
	createNode(s, parent, id, n.String(), "orange")
}

func (ta *TemplateArgument) toDot(s *strings.Builder) string {
	tfID := fmt.Sprintf("\"%pTemplateArgument\"", ta)

	nameID := hash(tfID + "Name")
	createNode(s, tfID, nameID, "Name", "white")

	id := hash(nameID)
	createNode(s, nameID, id, ta.FieldName, "orange")

	typeID := hash(tfID + "Type")
	createNode(s, tfID, typeID, "Type", "white")

	id = hash(typeID)
	if ta.IsNat {
		createNode(s, typeID, id, "#", "orange")
	} else {
		createNode(s, typeID, id, "Type", "orange")
	}

	return s.String()
}

func (f *Field) toDot(s *strings.Builder) string {
	fID := fmt.Sprintf("\"%pField\"", f)

	if f.FieldName != "" {
		nameID := hash(fID + "Name")
		createNode(s, fID, nameID, "Name", "white")

		id := hash(nameID + f.FieldName)
		createNode(s, nameID, id, f.FieldName, "orange")
	}

	if f.Excl {
		exclID := hash(fID + "!")
		createNode(s, fID, exclID, "Excl", "red")
	}

	if f.Mask != nil {
		maskID := fmt.Sprintf("\"%pFieldMask\"", f.Mask)
		createNode(s, fID, maskID, "Mask", "white")
		f.Mask.toDot(s)
	}

	var color string
	if f.FieldType.Bare {
		color = "blue"
	} else {
		color = "white"
	}
	if !f.IsRepeated {
		typeID := fmt.Sprintf("\"%pTypeRef\"", &f.FieldType)
		createNode(s, fID, typeID, "Type", color)
		f.FieldType.toDot(s)
	} else {
		rwsID := fmt.Sprintf("\"%prws\"", &f.ScaleRepeat)
		createNode(s, fID, rwsID, "Type", color)
		f.ScaleRepeat.toDot(s)
	}
	return s.String()
}

func (aot *ArithmeticOrType) toDot(s *strings.Builder) string {
	aotID := fmt.Sprintf("\"%paot\"", aot)
	if !aot.Arith.IsEmpty() {
		id := fmt.Sprintf("\"%pArithmetic\"", &aot.Arith)
		createNode(s, aotID, id, "Arithmetic", "white")
		aot.Arith.toDot(s)
	} else {
		id := fmt.Sprintf("\"%pTypeRef\"", &aot.T)
		createNode(s, aotID, id, "T", "white")
		aot.T.toDot(s)
	}
	return s.String()
}

func (a *Arithmetic) toDot(s *strings.Builder) {
	parent := fmt.Sprintf("\"%pArithmetic\"", a)
	for i := 0; i < len(a.Nums); i++ {
		id := fmt.Sprintf("\"%p\"", &(a.Nums[i]))
		createNode(s, parent, id, strconv.FormatUint(uint64(a.Nums[i]), 10), "powderblue")
	}
}

func (t *TypeRef) toDot(s *strings.Builder) {
	parent := fmt.Sprintf("\"%pTypeRef\"", t)
	if t.Type.IsEmpty() {
		return
	}
	typeID := fmt.Sprintf("\"%pName\"", &t.Type)
	if t.Bare {
		createNode(s, parent, typeID, "Type", "blue")
	} else {
		createNode(s, parent, typeID, "Type", "white")
	}
	t.Type.toDot(s)
	if t.Args == nil {
		return
	}
	argsID := fmt.Sprintf("\"%paots\"", &t.Args)
	createNode(s, parent, argsID, "Args", "white")
	for i := 0; i < len(t.Args); i++ {
		argID := fmt.Sprintf("\"%paot\"", &(t.Args[i]))
		createNode(s, argsID, argID, strconv.Itoa(i), "white")
		t.Args[i].toDot(s)
	}
}

func (rws *RepeatWithScale) toDot(s *strings.Builder) {
	parent := fmt.Sprintf("\"%prws\"", rws)
	if !rws.Scale.IsEmpty() {
		scaleID := fmt.Sprintf("\"%pScaleFactor\"", &rws.Scale)
		createNode(s, parent, scaleID, "Scale", "white")
		id := fmt.Sprintf("\"%s\"", rws.Scale.String())
		createNode(s, scaleID, id, rws.Scale.String(), "orange")
	}
	for i := 0; i < len(rws.Rep); i++ {
		id := fmt.Sprintf("\"%pField\"", &rws.Rep[i])
		createNode(s, parent, id, strconv.Itoa(i), "white")
		rws.Rep[i].toDot(s)
	}
}

func (fm *FieldMask) toDot(s *strings.Builder) string {
	parent := fmt.Sprintf("\"%pFieldMask\"", fm)

	nameID := hash(parent + "Name")
	createNode(s, parent, nameID, "Name", "white")

	id := hash(parent + fm.MaskName)
	createNode(s, nameID, id, fm.MaskName, "green")

	bitID := hash(parent + "Bit number")
	createNode(s, parent, bitID, "Bit Number", "white")

	id = hash(parent + strconv.FormatUint(uint64(fm.BitNumber), 10))
	createNode(s, bitID, id, strconv.FormatUint(uint64(fm.BitNumber), 10), "green")

	return s.String()
}

func (d *TypeDeclaration) toDot(s *strings.Builder) {
	parent := fmt.Sprintf("\"%pTypeDeclaration\"", d)

	tID := hash(parent + "Type")
	createNode(s, parent, tID, "Type", "white")

	id := fmt.Sprintf("\"%p\"", &d.Name.Name)
	createNode(s, tID, id, d.Name.String(), "orange")

	if d.Arguments != nil {
		aID := hash(parent + "Arguments")
		createNode(s, parent, aID, "Arguments", "white")
		for i := 0; i < len(d.Arguments); i++ {
			id := fmt.Sprintf("\"%p\"", &d.Arguments[i])
			createNode(s, aID, id, d.Arguments[i], "orange")
		}
	}
}
