// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by qtc from "qt_tljson_html.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

package tlcodegen

import (
	"fmt"
	"strconv"

	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

func streamtlJSON(qw422016 *qt422016.Writer, gen *Gen2, buildSHA256Checksum string) {
	qw422016.N().S(`<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <title>TL JSON help</title>
  </head>
  <body>
    <h1>Schema</h1>
    <ul>
      <li><abbr>TL</abbr> ⟷ <abbr>JSON</abbr> mapping rules: <a href="https://github.com/VKCOM/tl/blob/master/TLJSON.md">TLJSON.md</a></li>
    </ul>
    <h1>Functions</h1>
    <ul>
`)
	for _, trww := range gen.generatedTypesList {
		if fun, ok := trww.trw.(*TypeRWStruct); ok && fun.ResultType != nil {
			qw422016.N().S(`      <li>
        <a href="#`)
			qw422016.E().S(trww.JSONHelpString())
			qw422016.N().S(`">
        <code>`)
			qw422016.E().S(trww.JSONHelpString())
			qw422016.N().S(`</code></a>
        → <code>`)
			streamprintJSONHelpType(qw422016, gen, fun.ResultResolvedType, fun.ResultType, formatNatArgsJSONHelp(fun.Fields, fun.ResultNatArgs, trww.NatParams, trww.NatParams))
			qw422016.N().S(`</code>
      </li>
`)
		}
	}
	qw422016.N().S(`    </ul>
    <h1>Types</h1>
<h2 id="#">#</h2>
Builtin type <code>#</code>. Represents <code>uint32</code>. Can be used as field mask or collection size.
`)
	for _, trww := range gen.generatedTypesList {
		streamprintHTMLHelp(qw422016, gen, trww)
		qw422016.N().S(`
`)
	}
	qw422016.N().S(`  </body>
</html>
`)
}

func writetlJSON(qq422016 qtio422016.Writer, gen *Gen2, buildSHA256Checksum string) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	streamtlJSON(qw422016, gen, buildSHA256Checksum)
	qt422016.ReleaseWriter(qw422016)
}

func tlJSON(gen *Gen2, buildSHA256Checksum string) string {
	qb422016 := qt422016.AcquireByteBuffer()
	writetlJSON(qb422016, gen, buildSHA256Checksum)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

func streamprintJSONHelpType(qw422016 *qt422016.Writer, gen *Gen2, rt ResolvedType, trww *TypeRWWrapper, natArgs []string) {
	switch trw := trww.trw.(type) {
	case *TypeRWBool:
		qw422016.E().S("<bool>")
	case *TypeRWPrimitive:
		qw422016.E().S("<")
		qw422016.E().S(trw.primitiveType)
		qw422016.E().S(">")
	case *TypeRWMaybe:
		streammakeRef(qw422016, trw.goGlobalName)
	case *TypeRWStruct:
		if trw.isUnwrapTypeImpl(false) {
			streamprintJSONHelpType(qw422016, gen, trw.Fields[0].resolvedType, trw.Fields[0].t, formatNatArgsJSONHelp(trw.Fields, trw.Fields[0].natArgs, trww.NatParams, natArgs))
		} else if trw.wr.IsTrueType() {
			qw422016.E().S("{}")
		} else {
			streammakeRef(qw422016, trww.JSONHelpString())
		}
	case *TypeRWUnion:
		streammakeRef(qw422016, trww.JSONHelpString())
	case *TypeRWBrackets:
		switch {
		case trw.dictLike:
			qw422016.E().S("{")
			streamprintJSONHelpType(qw422016, gen, trw.dictKeyField.resolvedType, trw.dictKeyField.t, formatNatArgsJSONHelp(nil, trw.dictKeyField.natArgs, trww.NatParams, natArgs))
			qw422016.E().S(": ")
			streamprintJSONHelpType(qw422016, gen, trw.dictValueField.resolvedType, trw.dictValueField.t, formatNatArgsJSONHelp(nil, trw.dictValueField.natArgs, trww.NatParams, natArgs))
			qw422016.E().S("}")
		case trw.vectorLike:
			qw422016.E().S("[")
			streamprintJSONHelpType(qw422016, gen, trw.element.resolvedType, trw.element.t, formatNatArgsJSONHelp(nil, trw.element.natArgs, trww.NatParams, natArgs))
			qw422016.E().S(", ...]")
		case trw.dynamicSize:
			qw422016.E().S("[")
			qw422016.E().S(natArgs[len(natArgs)-1])
			qw422016.E().S(" × ")
			streamprintJSONHelpType(qw422016, gen, trw.element.resolvedType, trw.element.t, formatNatArgsJSONHelp(nil, trw.element.natArgs, trww.NatParams, natArgs))
			qw422016.E().S("]")
		default:
			qw422016.E().S("[")
			qw422016.E().S(strconv.Itoa(int(trw.size)))
			qw422016.E().S(" × ")
			streamprintJSONHelpType(qw422016, gen, trw.element.resolvedType, trw.element.t, formatNatArgsJSONHelp(nil, trw.element.natArgs, trww.NatParams, natArgs))
			qw422016.E().S("]")
		}
	}
}

func writeprintJSONHelpType(qq422016 qtio422016.Writer, gen *Gen2, rt ResolvedType, trww *TypeRWWrapper, natArgs []string) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	streamprintJSONHelpType(qw422016, gen, rt, trww, natArgs)
	qt422016.ReleaseWriter(qw422016)
}

func printJSONHelpType(gen *Gen2, rt ResolvedType, trww *TypeRWWrapper, natArgs []string) string {
	qb422016 := qt422016.AcquireByteBuffer()
	writeprintJSONHelpType(qb422016, gen, rt, trww, natArgs)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

func streamprintHTMLHelp(qw422016 *qt422016.Writer, gen *Gen2, trww *TypeRWWrapper) {
	if typ, ok := trww.trw.(*TypeRWStruct); ok && typ.ResultType == nil && trww.IsTrueType() {
		return
	}

	switch trw := trww.trw.(type) {
	case *TypeRWBool:
	case *TypeRWPrimitive:
		qw422016.N().S(`<h2 id="`)
		qw422016.E().S(trww.tlName.String())
		qw422016.N().S(`">`)
		qw422016.E().S(trww.tlName.String())
		qw422016.N().S(`</h2>
<p></p>
<dl>
  <dt>JSON</dt>
  <dd>`)
		qw422016.E().S(trw.primitiveType)
		qw422016.N().S(`</dd>
</dl>
`)
	case *TypeRWMaybe:
		qw422016.N().S(`<h2 id="`)
		qw422016.E().S(trw.goGlobalName)
		qw422016.N().S(`">`)
		qw422016.E().S(trw.goGlobalName)
		qw422016.N().S(`</h2>
<p></p>
<dl>
  <dt>JSON</dt>
  <dd>
    <ul>
      <li><code>{}</code></li>
      <li><code>`)
		qw422016.E().S(`{"value": `)
		streamprintJSONHelpType(qw422016, gen, trw.element.resolvedType, trw.element.t, formatNatArgsJSONHelp(nil, trw.element.natArgs, trww.NatParams, trww.NatParams))
		qw422016.E().S("}")
		qw422016.N().S(`</code></li>
    </ul>
  </dd>
  <dt>TL</dt>
  <dd>
    <ul>
    <li><code>`)
		qw422016.E().S(trww.origTL[0].String())
		qw422016.N().S(`</code></li>
    <li><code>`)
		qw422016.E().S(trww.origTL[1].String())
		qw422016.N().S(`</code></li>
    </ul>
  </dd>
</dl>
`)
	case *TypeRWStruct:
		if trw.isUnwrapTypeImpl(false) {
			return
		}
		qw422016.N().S(`<h2 id="`)
		qw422016.E().S(trww.JSONHelpString())
		qw422016.N().S(`">`)
		qw422016.E().S(trww.JSONHelpString())
		qw422016.N().S(`</h2>
<p></p>
<dl>
  <dt>JSON</dt>
  <dd><code>
`)
		if trw.ResultType != nil && trww.IsTrueType() {
			qw422016.N().S(`    {}
`)
		} else {
			qw422016.N().S(`    {
      <table>
`)
			for i, field := range trw.Fields {
				qw422016.N().S(`        <tr>
`)
				if field.t.IsTrueType() {
					qw422016.N().S(`          <td>&nbsp;&nbsp;"`)
					qw422016.E().S(field.originalName)
					qw422016.N().S(`"</td><td>: true`)
					if i != len(trw.Fields)-1 {
						qw422016.N().S(`,`)
					}
					qw422016.N().S(`</td>
`)
				} else {
					qw422016.N().S(`          <td>&nbsp;&nbsp;"`)
					qw422016.E().S(field.originalName)
					qw422016.N().S(`"</td><td>: `)
					streamprintJSONHelpType(qw422016, gen, field.resolvedType, field.t, formatNatArgsJSONHelp(trw.Fields, field.natArgs, trww.NatParams, trww.NatParams))
					if i != len(trw.Fields)-1 {
						qw422016.N().S(`,`)
					}
					qw422016.N().S(`</td>
`)
				}
				qw422016.N().S(`          <td>`)
				streamjsonCommentFieldMask(qw422016, field.fieldMask, field.BitNumber, trw.Fields)
				qw422016.N().S(`</td>
        </tr>
`)
			}
			qw422016.N().S(`      </table>
    }
`)
		}
		qw422016.N().S(`</code></dd>
  <dt>TL</dt>
  <dd>
    <code>`)
		qw422016.E().S(trww.origTL[0].String())
		qw422016.N().S(`</code>
  </dd>
</dl>
`)
	case *TypeRWUnion:
		qw422016.N().S(`<h2 id="`)
		qw422016.E().S(trww.JSONHelpString())
		qw422016.N().S(`">`)
		qw422016.E().S(trww.JSONHelpString())
		qw422016.N().S(`</h2>
<p></p>
<dl>
  <dt>JSON</dt>
  <dd>
    <ul>
`)
		if trw.IsEnum {
			for _, field := range trw.Fields {
				tag := fmt.Sprintf("%08x", field.t.tlTag)

				qw422016.N().S(`      <li><code>"`)
				qw422016.E().S(field.originalName)
				qw422016.N().S(`"</code> <small><small>(or <code>"#`)
				qw422016.E().S(tag)
				qw422016.N().S(`"</code>
      or <code>"`)
				qw422016.E().S(field.originalName)
				qw422016.N().S(`#`)
				qw422016.E().S(tag)
				qw422016.N().S(`"</code>)</small></small></li>
`)
			}
		} else {
			for _, field := range trw.Fields {
				tag := fmt.Sprintf("%08x", field.t.tlTag)

				qw422016.N().S(`      <li><code>{"type":"`)
				qw422016.E().S(field.originalName)
				qw422016.N().S(`"</code> <small><small>(or <code>"#`)
				qw422016.E().S(tag)
				qw422016.N().S(`"</code>
      or <code>"`)
				qw422016.E().S(field.originalName)
				qw422016.N().S(`#`)
				qw422016.E().S(tag)
				qw422016.N().S(`"</code>)</small></small><code>
`)
				if !field.t.IsTrueType() {
					qw422016.N().S(`,"value":`)
					streammakeRef(qw422016, field.t.JSONHelpString())
					qw422016.N().S(`
`)
				}
				qw422016.N().S(`      }</code></li>
`)
			}
		}
		qw422016.N().S(`    </ul>
  </dd>
  <dt>TL</dt>
  <dd>
    <ul>
`)
		for _, origTL := range trww.origTL {
			qw422016.N().S(`    <li><code>`)
			qw422016.E().S(origTL.String())
			qw422016.N().S(`</code></li>
`)
		}
		qw422016.N().S(`    </ul>
  </dd>
</dl>
`)
	case *TypeRWBrackets:
	}
}

func writeprintHTMLHelp(qq422016 qtio422016.Writer, gen *Gen2, trww *TypeRWWrapper) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	streamprintHTMLHelp(qw422016, gen, trww)
	qt422016.ReleaseWriter(qw422016)
}

func printHTMLHelp(gen *Gen2, trww *TypeRWWrapper) string {
	qb422016 := qt422016.AcquireByteBuffer()
	writeprintHTMLHelp(qb422016, gen, trww)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

func streamjsonCommentFieldMask(qw422016 *qt422016.Writer, fm *ActualNatArg, num uint32, fields []Field) {
	if fm == nil {
		return
	}
	if fm.isField {
		qw422016.N().S(`// `)
		qw422016.E().S(fields[fm.FieldIndex].originalName)
		qw422016.N().S(` bit #`)
		qw422016.E().S(strconv.Itoa(int(num)))
		qw422016.N().S(`
`)
	} else {
		qw422016.N().S(`// `)
		qw422016.E().S(fm.name)
		qw422016.N().S(` bit #`)
		qw422016.E().S(strconv.Itoa(int(num)))
		qw422016.N().S(`
`)
	}
}

func writejsonCommentFieldMask(qq422016 qtio422016.Writer, fm *ActualNatArg, num uint32, fields []Field) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	streamjsonCommentFieldMask(qw422016, fm, num, fields)
	qt422016.ReleaseWriter(qw422016)
}

func jsonCommentFieldMask(fm *ActualNatArg, num uint32, fields []Field) string {
	qb422016 := qt422016.AcquireByteBuffer()
	writejsonCommentFieldMask(qb422016, fm, num, fields)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

func streammakeRef(qw422016 *qt422016.Writer, s string) {
	qw422016.N().S(`<a href="#`)
	qw422016.E().S(s)
	qw422016.N().S(`">`)
	qw422016.E().S(s)
	qw422016.N().S(`</a>`)
}

func writemakeRef(qq422016 qtio422016.Writer, s string) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	streammakeRef(qw422016, s)
	qt422016.ReleaseWriter(qw422016)
}

func makeRef(s string) string {
	qb422016 := qt422016.AcquireByteBuffer()
	writemakeRef(qb422016, s)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}
