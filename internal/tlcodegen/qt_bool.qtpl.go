// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by qtc from "qt_bool.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

package tlcodegen

import (
	"fmt"

	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

func (trw *TypeRWBool) StreamGenerateCode(qw422016 *qt422016.Writer, bytesVersion bool, directImports *DirectImports) {
	qw422016.N().S(`const ( `)
	trw.streamgenerateBoolAlias(qw422016)
	qw422016.N().S(` )

func `)
	qw422016.N().S(addBytes(trw.wr.goGlobalName, bytesVersion))
	qw422016.N().S(`ReadBoxed(w []byte, v *bool) ([]byte, error) {
    return basictl.ReadBool(w, v, `)
	qw422016.N().S(trw.falseGoName)
	qw422016.N().S(`, `)
	qw422016.N().S(trw.trueGoName)
	qw422016.N().S(`)
}

func `)
	qw422016.N().S(addBytes(trw.wr.goGlobalName, bytesVersion))
	qw422016.N().S(`WriteBoxed(w []byte, v bool) ([]byte, error) {
    if v {
        return basictl.NatWrite(w, 0x`)
	qw422016.N().S(fmt.Sprintf("%x", trw.trueTag))
	qw422016.N().S(`), nil
    }
    return basictl.NatWrite(w, 0x`)
	qw422016.N().S(fmt.Sprintf("%x", trw.falseTag))
	qw422016.N().S(`), nil
}
`)
}

func (trw *TypeRWBool) WriteGenerateCode(qq422016 qtio422016.Writer, bytesVersion bool, directImports *DirectImports) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	trw.StreamGenerateCode(qw422016, bytesVersion, directImports)
	qt422016.ReleaseWriter(qw422016)
}

func (trw *TypeRWBool) GenerateCode(bytesVersion bool, directImports *DirectImports) string {
	qb422016 := qt422016.AcquireByteBuffer()
	trw.WriteGenerateCode(qb422016, bytesVersion, directImports)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

func (trw *TypeRWBool) streamgenerateBoolAlias(qw422016 *qt422016.Writer) {
	qw422016.N().S(`    `)
	qw422016.N().S(trw.falseGoName)
	qw422016.N().S(` uint32 = `)
	qw422016.N().S(fmt.Sprintf("%#x", trw.falseTag))
	qw422016.N().S(`
    `)
	qw422016.N().S(trw.trueGoName)
	qw422016.N().S(` uint32 = `)
	qw422016.N().S(fmt.Sprintf("%#x", trw.trueTag))
	qw422016.N().S(`
`)
}

func (trw *TypeRWBool) writegenerateBoolAlias(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	trw.streamgenerateBoolAlias(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func (trw *TypeRWBool) generateBoolAlias() string {
	qb422016 := qt422016.AcquireByteBuffer()
	trw.writegenerateBoolAlias(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}
