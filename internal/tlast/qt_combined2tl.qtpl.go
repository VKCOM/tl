// Code generated by qtc from "qt_combined2tl.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

package tlast

import "fmt"

import "sort"

import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

func compareModifiers(a, b Modifier) bool {
	return modifierToFlag([]Modifier{a}) < modifierToFlag([]Modifier{b})
}

func (descriptor Combinator) streamcanonicalFormWithTag(qw422016 *qt422016.Writer) {
	modifiers := append([]Modifier(nil), descriptor.Modifiers...)
	sort.Slice(modifiers, func(i, j int) bool {
		return compareModifiers(modifiers[i], modifiers[j])
	})
	haveKphp := false
	for _, m := range modifiers {
		if m.Name == "@kphp" {
			haveKphp = true
		}
	}
	if haveKphp && modifiers[0].Name != "@any" {
		modifiers = append([]Modifier{{Name: "@any"}}, modifiers...)
	}

	for _, mod := range modifiers {
		qw422016.N().S(`@`)
		qw422016.N().S(mod.Name)
		qw422016.N().S(` `)
	}
	descriptor.Construct.Name.StreamString(qw422016)
	qw422016.N().S(`#`)
	qw422016.N().S(fmt.Sprintf("%08x", descriptor.Crc32()))
	qw422016.N().S(` `)
	for _, x := range descriptor.TemplateArguments {
		qw422016.N().S(x.FieldName)
		if x.IsNat {
			qw422016.N().S(":# ")
		} else {
			qw422016.N().S(":Type ")
		}
	}
	if descriptor.Builtin {
		qw422016.N().S("? ")
	}
	for _, x := range descriptor.Fields {
		if x.FieldName != "" {
			qw422016.N().S(x.FieldName)
			qw422016.N().S(`:`)
		}
		if x.Mask != nil {
			x.Mask.StreamString(qw422016)
		}
		if x.IsRepeated {
			x.ScaleRepeat.streamtoCrc32(qw422016)
		} else {
			x.FieldType.streamtoCrc32(qw422016)
		}
		qw422016.N().S(` `)
	}
	qw422016.N().S("= ")
	if descriptor.IsFunction {
		descriptor.FuncDecl.streamtoCrc32(qw422016)
	} else {
		descriptor.TypeDecl.StreamString(qw422016)
	}
}

func (descriptor Combinator) writecanonicalFormWithTag(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	descriptor.streamcanonicalFormWithTag(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func (descriptor Combinator) canonicalFormWithTag() string {
	qb422016 := qt422016.AcquireByteBuffer()
	descriptor.writecanonicalFormWithTag(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

func (tl TL) StreamGenerate2TL(qw422016 *qt422016.Writer) {
	qw422016.N().S(`int#a8509bda ? = Int`)
	qw422016.N().S(`
`)
	qw422016.N().S(`long#22076cba ? = Long`)
	qw422016.N().S(`
`)
	qw422016.N().S(`float#824dab22 ? = Float`)
	qw422016.N().S(`
`)
	qw422016.N().S(`double#2210c154 ? = Double`)
	qw422016.N().S(`
`)
	qw422016.N().S(`string#b5286e24 ? = String`)
	qw422016.N().S(`
`)
	for _, combinator := range tl {
		switch combinator.Construct.Name.String() {
		case "int", "long", "float", "double", "string":
			continue
		default:
			combinator.streamcanonicalFormWithTag(qw422016)
			qw422016.N().S(` `)
			qw422016.N().S(`//`)
			qw422016.N().S(` `)
			qw422016.N().S(` `)
			qw422016.N().S(combinator.PR.Begin.file)
			qw422016.N().S(`
`)
		}
	}
}

func (tl TL) WriteGenerate2TL(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	tl.StreamGenerate2TL(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func (tl TL) Generate2TL() string {
	qb422016 := qt422016.AcquireByteBuffer()
	tl.WriteGenerate2TL(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}
