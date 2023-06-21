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
	h := fnv.New32a()
	_, _ = h.Write([]byte(s))
	return fmt.Sprintf("\"%x\"", h.Sum32())
}

func createNode(parent, child, nodeName string, color string) string {
	var s strings.Builder
	_, _ = fmt.Fprintf(&s, "%s [shape=record;color=%s;label=%q];\n", child, color, nodeName)
	_, _ = fmt.Fprintf(&s, "%s -> %s;\n", parent, child)
	return s.String()
}

func (tl *TL) ToDot() string {
	var s, t, f strings.Builder
	s.WriteString("digraph TL { splines=\"ortho\";\nordering=in;\ngraph[ordering=in];\n")
	tlID := fmt.Sprintf("\"%pTL\"", tl)
	typesID := hash("types")
	functionsID := hash("functions")
	s.WriteString(tlID + " [shape=record;color=white;label=\"TL\"];\n")
	t.WriteString(createNode(tlID, typesID, "Types", "white"))
	f.WriteString(createNode(tlID, functionsID, "Functions", "white"))
	t.WriteString("subgraph types { ordering=in;\ngraph[ordering=in];\n")
	f.WriteString("subgraph functions { ordering=in;\ngraph[ordering=in];\n")
	for i := 0; i < len(*tl); i++ {
		x := (*tl)[i]
		id := fmt.Sprintf("\"%pCombinator\"", x)
		if x.IsFunction {
			f.WriteString(createNode(functionsID, id, "Function", "white"))
			f.WriteString(x.toDot())
		} else {
			t.WriteString(createNode(typesID, id, "Type", "white"))
			t.WriteString(x.toDot())
		}
	}
	t.WriteString("}\n")
	f.WriteString("}\n")
	s.WriteString(t.String())
	s.WriteString(f.String())
	s.WriteString("}\n")
	return s.String()
}

func (descriptor *Combinator) toDot() string {
	var s strings.Builder
	parent := fmt.Sprintf("\"%pCombinator\"", descriptor)

	if descriptor.IsFunction {
		modifiersID := fmt.Sprintf("\"%pModifiers\"", &descriptor.Modifiers)
		s.WriteString(createNode(parent, modifiersID, "Modifiers", "white"))
		for _, m := range descriptor.Modifiers {
			modID := hash(modifiersID + m.Name)
			s.WriteString(createNode(modifiersID, modID, m.Name, "yellow"))
		}
	}

	constructID := fmt.Sprintf("\"%pConstructor\"", &descriptor.Construct)
	s.WriteString(createNode(parent, constructID, "Constructor", "white"))
	s.WriteString(descriptor.Construct.toDot())

	if descriptor.TemplateArguments != nil {
		templateFieldID := fmt.Sprintf("\"%pTemplateArguments\"", &descriptor.TemplateArguments)
		s.WriteString(createNode(parent, templateFieldID, "Template Field", "white"))
		s.WriteString(toDotTemplateFields(&descriptor.TemplateArguments))
	}

	if descriptor.Fields != nil {
		fieldID := fmt.Sprintf("\"%pFields\"", &descriptor.Fields)
		s.WriteString(createNode(parent, fieldID, "Fields", "white"))
		s.WriteString(toDotFields(&descriptor.Fields))
	}

	if descriptor.IsFunction {
		funcDeclID := fmt.Sprintf("\"%pTypeRef\"", &descriptor.FuncDecl)
		s.WriteString(createNode(parent, funcDeclID, "Func Declaration", "white"))
		s.WriteString(descriptor.FuncDecl.toDot())
	} else {
		typeDeclID := fmt.Sprintf("\"%pTypeDeclaration\"", &descriptor.TypeDecl)
		s.WriteString(createNode(parent, typeDeclID, "Type Declaration", "white"))
		s.WriteString(descriptor.TypeDecl.toDot())
	}
	return s.String()
}

func toDotTemplateFields(tfs *[]TemplateArgument) string {
	var s strings.Builder
	parent := fmt.Sprintf("\"%pTemplateArguments\"", tfs)
	for i := 0; i < len(*tfs); i++ {
		tfID := fmt.Sprintf("\"%pTemplateArgument\"", &(*tfs)[i])
		s.WriteString(createNode(parent, tfID, strconv.Itoa(i), "white"))
		s.WriteString((*tfs)[i].toDot())
	}
	return s.String()
}

func toDotFields(fs *[]Field) string {
	var s strings.Builder
	parent := fmt.Sprintf("\"%pFields\"", fs)
	for i := 0; i < len(*fs); i++ {
		fID := fmt.Sprintf("\"%pField\"", &(*fs)[i])
		s.WriteString(createNode(parent, fID, strconv.Itoa(i), "white"))
		s.WriteString((*fs)[i].toDot())
	}
	return s.String()
}

func (c *Constructor) toDot() string {
	var s strings.Builder
	parent := fmt.Sprintf("\"%pConstructor\"", c)
	nameID := fmt.Sprintf("\"%pName\"", &c.Name)
	s.WriteString(createNode(parent, nameID, "Name", "white"))
	s.WriteString(c.Name.toDot())

	if c.ID == nil {
		return s.String()
	}
	idID := fmt.Sprintf("\"%p\"", c.ID)
	s.WriteString(createNode(parent, idID, "tag", "white"))
	id := hash(idID)
	s.WriteString(createNode(idID, id, fmt.Sprintf("%08x", *c.ID), "cyan"))
	return s.String()
}

func (n *Name) toDot() string {
	var s strings.Builder
	parent := fmt.Sprintf("\"%pName\"", n)
	id := hash(parent + n.String())
	s.WriteString(createNode(parent, id, n.String(), "orange"))
	return s.String()
}

func (ta *TemplateArgument) toDot() string {
	var s strings.Builder
	tfID := fmt.Sprintf("\"%pTemplateArgument\"", ta)

	nameID := hash(tfID + "Name")
	s.WriteString(createNode(tfID, nameID, "Name", "white"))

	id := hash(nameID)
	s.WriteString(createNode(nameID, id, ta.FieldName, "orange"))

	typeID := hash(tfID + "Type")
	s.WriteString(createNode(tfID, typeID, "Type", "white"))

	id = hash(typeID)
	if ta.IsNat {
		s.WriteString(createNode(typeID, id, "#", "orange"))
	} else {
		s.WriteString(createNode(typeID, id, "Type", "orange"))
	}

	return s.String()
}

func (f *Field) toDot() string {
	var s strings.Builder
	fID := fmt.Sprintf("\"%pField\"", f)

	if f.FieldName != "" {
		nameID := hash(fID + "Name")
		s.WriteString(createNode(fID, nameID, "Name", "white"))

		id := hash(nameID + f.FieldName)
		s.WriteString(createNode(nameID, id, f.FieldName, "orange"))
	}

	if f.Excl {
		exclID := hash(fID + "!")
		s.WriteString(createNode(fID, exclID, "Excl", "red"))
	}

	if f.Mask != nil {
		maskID := fmt.Sprintf("\"%pFieldMask\"", f.Mask)
		s.WriteString(createNode(fID, maskID, "Mask", "white"))
		s.WriteString(f.Mask.toDot())
	}

	var color string
	if f.FieldType.Bare {
		color = "blue"
	} else {
		color = "white"
	}
	if !f.IsRepeated {
		typeID := fmt.Sprintf("\"%pTypeRef\"", &f.FieldType)
		s.WriteString(createNode(fID, typeID, "Type", color))
		s.WriteString(f.FieldType.toDot())
	} else {
		rwsID := fmt.Sprintf("\"%prws\"", &f.ScaleRepeat)
		s.WriteString(createNode(fID, rwsID, "Type", color))
		s.WriteString(f.ScaleRepeat.toDot())
	}
	return s.String()
}

func (aot *ArithmeticOrType) toDot() string {
	var s strings.Builder
	aotID := fmt.Sprintf("\"%paot\"", aot)
	if !aot.Arith.IsEmpty() {
		id := fmt.Sprintf("\"%pArithmetic\"", &aot.Arith)
		s.WriteString(createNode(aotID, id, "Arithmetic", "white"))
		s.WriteString(aot.Arith.toDot())
	} else {
		id := fmt.Sprintf("\"%pTypeRef\"", &aot.T)
		s.WriteString(createNode(aotID, id, "T", "white"))
		s.WriteString(aot.T.toDot())
	}
	return s.String()
}

func (a *Arithmetic) toDot() string {
	var s strings.Builder
	parent := fmt.Sprintf("\"%pArithmetic\"", a)
	for i := 0; i < len(a.Nums); i++ {
		id := fmt.Sprintf("\"%p\"", &(a.Nums[i]))
		s.WriteString(createNode(parent, id, strconv.FormatUint(uint64(a.Nums[i]), 10), "powderblue"))
	}
	return s.String()
}

func (t *TypeRef) toDot() string {
	var s strings.Builder
	parent := fmt.Sprintf("\"%pTypeRef\"", t)
	if t.Type.IsEmpty() {
		return s.String()
	}
	typeID := fmt.Sprintf("\"%pName\"", &t.Type)
	if t.Bare {
		s.WriteString(createNode(parent, typeID, "Type", "blue"))
	} else {
		s.WriteString(createNode(parent, typeID, "Type", "white"))
	}
	s.WriteString(t.Type.toDot())
	if t.Args == nil {
		return s.String()
	}
	argsID := fmt.Sprintf("\"%paots\"", &t.Args)
	s.WriteString(createNode(parent, argsID, "Args", "white"))
	for i := 0; i < len(t.Args); i++ {
		argID := fmt.Sprintf("\"%paot\"", &(t.Args[i]))
		s.WriteString(createNode(argsID, argID, strconv.Itoa(i), "white"))
		s.WriteString(t.Args[i].toDot())
	}
	return s.String()
}

func (rws *RepeatWithScale) toDot() string {
	var s strings.Builder
	parent := fmt.Sprintf("\"%prws\"", rws)
	if !rws.Scale.IsEmpty() {
		scaleID := fmt.Sprintf("\"%pScaleFactor\"", &rws.Scale)
		s.WriteString(createNode(parent, scaleID, "Scale", "white"))
		id := fmt.Sprintf("\"%s\"", rws.Scale.String())
		s.WriteString(createNode(scaleID, id, rws.Scale.String(), "orange"))
	}
	for i := 0; i < len(rws.Rep); i++ {
		id := fmt.Sprintf("\"%pField\"", &rws.Rep[i])
		s.WriteString(createNode(parent, id, strconv.Itoa(i), "white"))
		s.WriteString(rws.Rep[i].toDot())
	}
	return s.String()
}

func (fm *FieldMask) toDot() string {
	var s strings.Builder
	parent := fmt.Sprintf("\"%pFieldMask\"", fm)

	nameID := hash(parent + "Name")
	s.WriteString(createNode(parent, nameID, "Name", "white"))

	id := hash(parent + fm.MaskName)
	s.WriteString(createNode(nameID, id, fm.MaskName, "green"))

	bitID := hash(parent + "Bit number")
	s.WriteString(createNode(parent, bitID, "Bit Number", "white"))

	id = hash(parent + strconv.FormatUint(uint64(fm.BitNumber), 10))
	s.WriteString(createNode(bitID, id, strconv.FormatUint(uint64(fm.BitNumber), 10), "green"))

	return s.String()
}

func (d *TypeDeclaration) toDot() string {
	var s strings.Builder
	parent := fmt.Sprintf("\"%pTypeDeclaration\"", d)

	tID := hash(parent + "Type")
	s.WriteString(createNode(parent, tID, "Type", "white"))

	id := fmt.Sprintf("\"%p\"", &d.Name.Name)
	s.WriteString(createNode(tID, id, d.Name.String(), "orange"))

	if d.Arguments != nil {
		aID := hash(parent + "Arguments")
		s.WriteString(createNode(parent, aID, "Arguments", "white"))
		for i := 0; i < len(d.Arguments); i++ {
			id := fmt.Sprintf("\"%p\"", &d.Arguments[i])
			s.WriteString(createNode(aID, id, d.Arguments[i], "orange"))
		}
	}
	return s.String()
}
