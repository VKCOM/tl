// Code generated by qtc from "qt_maybe.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

package tlcodegen

import "fmt"

import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

func (maybe *TypeRWMaybe) StreamGenerateCode(qw422016 *qt422016.Writer, bytesVersion bool, directImports *DirectImports) {
	goName := addBytes(maybe.wr.goGlobalName, bytesVersion)
	elementTypeString := maybe.element.t.TypeString2(bytesVersion, directImports, maybe.wr.ins, false, false)
	natArgsDecl := formatNatArgsDecl(maybe.wr.NatParams)
	natArgsCall := formatNatArgsDeclCall(maybe.wr.NatParams)
	emptyTag := fmt.Sprintf("%#x", maybe.emptyTag)
	okTag := fmt.Sprintf("%#x", maybe.okTag)
	writeElementNeedsError := maybe.element.t.hasErrorInWriteMethods

	qw422016.N().S(`type `)
	qw422016.N().S(goName)
	qw422016.N().S(` struct {
    Value `)
	qw422016.N().S(elementTypeString)
	qw422016.N().S(` // not deterministic if !Ok
    Ok    bool
}

func (item *`)
	qw422016.N().S(goName)
	qw422016.N().S(`) Reset() {
    item.Ok = false
}
`)
	if maybe.wr.gen.options.GenerateRandomCode {
		qw422016.N().S(`func (item *`)
		qw422016.N().S(goName)
		qw422016.N().S(`) FillRandom(rg *basictl.RandGenerator`)
		qw422016.N().S(natArgsDecl)
		qw422016.N().S(`) {
    if basictl.RandomUint(rg) & 1 == 1 {
        item.Ok = true
        `)
		qw422016.N().S(maybe.element.t.TypeRandomCode(bytesVersion, directImports, maybe.wr.ins, "item.Value", formatNatArgs(nil, maybe.element.natArgs), false))
		qw422016.N().S(`
    } else {
        item.Ok = false
    }
}
`)
	}
	qw422016.N().S(`
func (item *`)
	qw422016.N().S(goName)
	qw422016.N().S(`) ReadBoxed(w []byte`)
	qw422016.N().S(natArgsDecl)
	qw422016.N().S(`) (_ []byte, err error) {
    if w, err = basictl.ReadBool(w, &item.Ok, `)
	qw422016.N().S(emptyTag)
	qw422016.N().S(`, `)
	qw422016.N().S(okTag)
	qw422016.N().S(`); err != nil {
        return w, err
    }
    if item.Ok {
        `)
	qw422016.N().S(maybe.element.t.TypeReadingCode(bytesVersion, directImports, maybe.wr.ins, "item.Value", maybe.element.Bare(), formatNatArgs(nil, maybe.element.natArgs), false, true))
	qw422016.N().S(`
    }
    return w, nil
}

// This method is general version of WriteBoxed, use it instead!
func (item *`)
	qw422016.N().S(goName)
	qw422016.N().S(`) WriteBoxedGeneral(w []byte`)
	qw422016.N().S(natArgsDecl)
	qw422016.N().S(`) (_ []byte, err error) {
`)
	if writeElementNeedsError {
		qw422016.N().S(`    return item.WriteBoxed(w`)
		qw422016.N().S(natArgsCall)
		qw422016.N().S(`)
`)
	} else {
		qw422016.N().S(`    return item.WriteBoxed(w`)
		qw422016.N().S(natArgsCall)
		qw422016.N().S(`), nil
`)
	}
	qw422016.N().S(`}

func (item *`)
	qw422016.N().S(goName)
	qw422016.N().S(`) WriteBoxed(w []byte`)
	qw422016.N().S(natArgsDecl)
	qw422016.N().S(`) `)
	qw422016.N().S(wrapWithError(writeElementNeedsError, "[]byte"))
	qw422016.N().S(` {
    if item.Ok {
        w = basictl.NatWrite(w, `)
	qw422016.N().S(okTag)
	qw422016.N().S(`)
        `)
	qw422016.N().S(maybe.element.t.TypeWritingCode(bytesVersion, directImports, maybe.wr.ins, "item.Value", maybe.element.Bare(), formatNatArgs(nil, maybe.element.natArgs), false, true, writeElementNeedsError))
	qw422016.N().S(`
    }
`)
	if writeElementNeedsError {
		qw422016.N().S(`    return basictl.NatWrite(w, `)
		qw422016.N().S(emptyTag)
		qw422016.N().S(`), nil
`)
	} else {
		qw422016.N().S(`    return basictl.NatWrite(w, `)
		qw422016.N().S(emptyTag)
		qw422016.N().S(`)
`)
	}
	qw422016.N().S(`}

`)
	if maybe.wr.gen.options.GenerateLegacyJsonRead {
		qw422016.N().S(`func (item *`)
		qw422016.N().S(goName)
		qw422016.N().S(`) ReadJSONLegacy(legacyTypeNames bool, j interface{}`)
		qw422016.N().S(natArgsDecl)
		qw422016.N().S(`) error {
  _ok, _jvalue, err := `)
		qw422016.N().S(maybe.wr.gen.InternalPrefix())
		qw422016.N().S(`JsonReadMaybe("`)
		maybe.wr.tlName.StreamString(qw422016)
		qw422016.N().S(`", j)
  if err != nil {
    return err
  }
  item.Ok = _ok
  if _ok {
    `)
		qw422016.N().S(maybe.element.t.TypeJSONReadingCode(bytesVersion, directImports, maybe.wr.ins, "_jvalue", "item.Value", formatNatArgs(nil, maybe.element.natArgs), false))
		qw422016.N().S(`
  }
  return nil
}

`)
	}
	qw422016.N().S(`func (item *`)
	qw422016.N().S(goName)
	qw422016.N().S(`) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer`)
	qw422016.N().S(natArgsDecl)
	qw422016.N().S(`) error {
  _ok, _jvalue, err := `)
	qw422016.N().S(maybe.wr.gen.InternalPrefix())
	qw422016.N().S(`Json2ReadMaybe("`)
	maybe.wr.tlName.StreamString(qw422016)
	qw422016.N().S(`", in)
  if err != nil {
    return err
  }
  item.Ok = _ok
  if _ok {
    var in2Pointer *basictl.JsonLexer
    if _jvalue != nil {
        in2 := basictl.JsonLexer{Data: _jvalue}
        in2Pointer = &in2
    }
    `)
	qw422016.N().S(maybe.element.t.TypeJSON2ReadingCode(bytesVersion, directImports, maybe.wr.ins, "in2Pointer", "item.Value", formatNatArgs(nil, maybe.element.natArgs), false))
	qw422016.N().S(`
  }
  return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *`)
	qw422016.N().S(goName)
	qw422016.N().S(`) WriteJSONGeneral(w []byte`)
	qw422016.N().S(natArgsDecl)
	qw422016.N().S(`) (_ []byte, err error) {
`)
	if writeElementNeedsError {
		qw422016.N().S(`    return item.WriteJSONOpt(true, false, w`)
		qw422016.N().S(natArgsCall)
		qw422016.N().S(`)
`)
	} else {
		qw422016.N().S(`    return item.WriteJSONOpt(true, false, w`)
		qw422016.N().S(natArgsCall)
		qw422016.N().S(`), nil
`)
	}
	qw422016.N().S(`}

func (item *`)
	qw422016.N().S(goName)
	qw422016.N().S(`) WriteJSON(w []byte`)
	qw422016.N().S(natArgsDecl)
	qw422016.N().S(`) `)
	qw422016.N().S(wrapWithError(writeElementNeedsError, "[]byte"))
	qw422016.N().S(` {
    return item.WriteJSONOpt(true, false, w`)
	qw422016.N().S(natArgsCall)
	qw422016.N().S(`)
}
func (item *`)
	qw422016.N().S(goName)
	qw422016.N().S(`) WriteJSONOpt(newTypeNames bool, short bool, w []byte`)
	qw422016.N().S(natArgsDecl)
	qw422016.N().S(`) `)
	qw422016.N().S(wrapWithError(writeElementNeedsError, "[]byte"))
	qw422016.N().S(` {
    if !item.Ok {
`)
	if writeElementNeedsError {
		qw422016.N().S(`        return append(w, "{}"...), nil
`)
	} else {
		qw422016.N().S(`        return append(w, "{}"...)
`)
	}
	qw422016.N().S(`    }
    w = append(w, `)
	qw422016.N().S("`")
	qw422016.N().S(`{"ok":true`)
	qw422016.N().S("`")
	qw422016.N().S(`...)
`)
	emptyCondition := maybe.element.t.TypeJSONEmptyCondition(bytesVersion, "item.Value", false)

	if emptyCondition != "" {
		qw422016.N().S(`    if `)
		qw422016.N().S(emptyCondition)
		qw422016.N().S(` {
`)
	}
	qw422016.N().S(`    w = append(w, `)
	qw422016.N().S("`")
	qw422016.N().S(`,"value":`)
	qw422016.N().S("`")
	qw422016.N().S(`...)
    `)
	qw422016.N().S(maybe.element.t.TypeJSONWritingCode(bytesVersion, directImports, maybe.wr.ins, "item.Value", formatNatArgs(nil, maybe.element.natArgs), false, writeElementNeedsError))
	qw422016.N().S(`
`)
	if emptyCondition != "" {
		qw422016.N().S(`    }
`)
	}
	if writeElementNeedsError {
		qw422016.N().S(`    return append(w, '}'), nil
`)
	} else {
		qw422016.N().S(`    return append(w, '}')
`)
	}
	qw422016.N().S(`}
`)
	if len(maybe.wr.NatParams) == 0 {
		qw422016.N().S(`
func (item `)
		qw422016.N().S(goName)
		qw422016.N().S(`) String(`)
		qw422016.N().S(formatNatArgsDeclNoComma(maybe.wr.NatParams))
		qw422016.N().S(`) string {
`)
		if writeElementNeedsError {
			qw422016.N().S(`    w, err := item.WriteJSON(nil`)
			qw422016.N().S(natArgsCall)
			qw422016.N().S(`)
    if err != nil {
        return err.Error()
    }
    return string(w)
`)
		} else {
			qw422016.N().S(`    return string(item.WriteJSON(nil`)
			qw422016.N().S(natArgsCall)
			qw422016.N().S(`))
`)
		}
		qw422016.N().S(`}

`)
	}
	qw422016.N().S(`
`)
}

func (maybe *TypeRWMaybe) WriteGenerateCode(qq422016 qtio422016.Writer, bytesVersion bool, directImports *DirectImports) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	maybe.StreamGenerateCode(qw422016, bytesVersion, directImports)
	qt422016.ReleaseWriter(qw422016)
}

func (maybe *TypeRWMaybe) GenerateCode(bytesVersion bool, directImports *DirectImports) string {
	qb422016 := qt422016.AcquireByteBuffer()
	maybe.WriteGenerateCode(qb422016, bytesVersion, directImports)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}
