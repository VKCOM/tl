// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by qtc from "qt_meta.qtpl". DO NOT EDIT.
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

func (gen *Gen2) streamgenerateMeta(qw422016 *qt422016.Writer) {
	qw422016.N().S(HeaderComment)
	qw422016.N().S(`
`)
	typeWrappers := gen.generatedTypesList

	qw422016.N().S(`package `)
	qw422016.E().S(MetaGoPackageName)
	qw422016.N().S(`

import (
    "fmt"

    `)
	qw422016.N().Q(gen.BasicPackageNameFull)
	qw422016.N().S(`
    "`)
	qw422016.N().S(gen.options.TLPackageNameFull)
	qw422016.N().S(`/internal"
)

// We can create only types which have zero type arguments and zero nat arguments
type Object interface {
	TLName() string // returns type's TL name. For union, returns constructor name depending on actual union value
	TLTag() uint32  // returns type's TL tag. For union, returns constructor tag depending on actual union value
	String() string // returns type's representation for debugging (JSON for now)

`)
	if gen.options.GenerateRandomCode {
		qw422016.N().S(`	FillRandom(rand basictl.Rand)
`)
	}
	qw422016.N().S(`	Read(w []byte) ([]byte, error) // reads type's bare TL representation by consuming bytes from the start of w and returns remaining bytes, plus error
	Write(w []byte) ([]byte, error) // appends bytes of type's bare TL representation to the end of w and returns it, plus error
	ReadBoxed(w []byte) ([]byte, error) // same as Read, but reads/checks TLTag first
	WriteBoxed(w []byte) ([]byte, error) // same as Write, but writes TLTag first

	MarshalJSON() ([]byte, error) // returns type's JSON representation, plus error
	UnmarshalJSON([]byte) error // reads type's JSON representation
	WriteJSON(w []byte) ([]byte, error) // like MarshalJSON, but appends to w and returns it
}

type Function interface {
	Object

	ReadResultWriteResultJSON(r []byte, w []byte) ([]byte, []byte, error) // combination of ReadResult(r) + WriteResultJSON(w). Returns new r, new w, plus error
	ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) // combination of ReadResultJSON(r) + WriteResult(w). Returns new r, new w, plus error

    // For transcoding short-long version during Long ID transition
	ReadResultWriteResultJSONShort(r []byte, w []byte) ([]byte, []byte, error)
}

// for quick one-liners
func GetTLName(tag uint32, notFoundName string) string {
	if item := FactoryItemByTLTag(tag); item != nil {
		return item.TLName()
	}
	return notFoundName
}

func CreateFunction(tag uint32) Function {
    if item := FactoryItemByTLTag(tag); item != nil && item.createFunction != nil {
        return item.createFunction()
    }
    return nil
}

func CreateObject(tag uint32) Object {
    if item := FactoryItemByTLTag(tag); item != nil && item.createObject != nil {
        return item.createObject()
    }
    return nil
}

// name can be in any of 3 forms "ch_proxy.insert#7cf362ba", "ch_proxy.insert" or "#7cf362ba"
func CreateFunctionFromName(name string) Function {
    if item := FactoryItemByTLName(name); item != nil && item.createFunction != nil {
        return item.createFunction()
    }
    return nil
}

// name can be in any of 3 forms "ch_proxy.insert#7cf362ba", "ch_proxy.insert" or "#7cf362ba"
func CreateObjectFromName(name string) Object {
    if item := FactoryItemByTLName(name); item != nil && item.createObject != nil {
        return item.createObject()
    }
    return nil
}

type TLItem struct {
    tag                uint32
    annotations        uint32
    tlName             string
    createFunction     func() Function
    createFunctionLong func() Function
    createObject       func() Object
}

func (item TLItem) TLTag() uint32            { return item.tag }
func (item TLItem) TLName() string           { return item.tlName }
func (item TLItem) CreateObject() Object     { return item.createObject() }
func (item TLItem) IsFunction() bool         { return item.createFunction != nil }
func (item TLItem) CreateFunction() Function { return item.createFunction() }

// For transcoding short-long version during Long ID transition
func (item TLItem) HasFunctionLong() bool        { return item.createFunctionLong != nil }
func (item TLItem) CreateFunctionLong() Function { return item.createFunctionLong() }

// Annotations
`)
	for bit, name := range gen.allAnnotations {
		goName := ToUpperFirst(name)

		qw422016.N().S(`func (item TLItem) Annotation`)
		qw422016.N().S(goName)
		qw422016.N().S(`() bool { return item.annotations & `)
		qw422016.N().S(fmt.Sprintf("%#x", 1<<bit))
		qw422016.N().S(` != 0 }
`)
	}
	qw422016.N().S(`
// TLItem serves as a single type for all enum values
func (item *TLItem) Reset()                         {}
`)
	if gen.options.GenerateRandomCode {
		qw422016.N().S(`	func (item *TLItem) FillRandom(rand basictl.Rand) {}
`)
	}
	qw422016.N().S(`func (item *TLItem) Read(w []byte) ([]byte, error)  { return w, nil }
func (item *TLItem) Write(w []byte) ([]byte, error) { return w, nil }
func (item *TLItem) ReadBoxed(w []byte) ([]byte, error) { return basictl.NatReadExactTag(w, item.tag) }
func (item *TLItem) WriteBoxed(w []byte) ([]byte, error) { return basictl.NatWrite(w, item.tag), nil }
func (item TLItem) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}
func (item *TLItem) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return internal.ErrorInvalidJSON(item.tlName, "expected json object")
	}
	for k := range _jm {
		return internal.ErrorInvalidJSONExcessElement(item.tlName, k)
	}
	return nil
}
func (item *TLItem) WriteJSON(w []byte) (_ []byte, err error) {
	w = append(w, '{')
	return append(w, '}'), nil
}
func (item *TLItem) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}
func (item *TLItem) UnmarshalJSON(b []byte) error {
	j, err := internal.JsonBytesToInterface(b)
	if err != nil {
		return internal.ErrorInvalidJSON(item.tlName, err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return internal.ErrorInvalidJSON(item.tlName, err.Error())
	}
	return nil
}


func FactoryItemByTLTag(tag uint32) *TLItem {
    return itemsByTag[tag]
}

func FactoryItemByTLName(name string) *TLItem {
    return itemsByName[name]
}

var itemsByTag = map[uint32]*TLItem {}

var itemsByName = map[string]*TLItem {}

func SetGlobalFactoryCreateForFunction(itemTag uint32, createObject func() Object, createFunction func() Function, createFunctionLong func() Function) {
    item := itemsByTag[itemTag]
    if item == nil {
        panic(fmt.Sprintf("factory cannot find function tag #%08x to set", itemTag))
    }
    item.createObject = createObject
    item.createFunction = createFunction
    item.createFunctionLong = createFunctionLong
}

func SetGlobalFactoryCreateForObject(itemTag uint32, createObject func() Object) {
    item := itemsByTag[itemTag]
    if item == nil {
        panic(fmt.Sprintf("factory cannot find item tag #%08x to set", itemTag))
    }
    item.createObject = createObject
}

func SetGlobalFactoryCreateForEnumElement(itemTag uint32) {
    item := itemsByTag[itemTag]
    if item == nil {
        panic(fmt.Sprintf("factory cannot find enum tag #%08x to set", itemTag))
    }
    item.createObject = func() Object { return item }
}

func pleaseImportFactoryObject() Object {
       panic("factory functions are not linked to reduce code bloat, please import 'gen/factory' instead of 'gen/meta'.")
}

func pleaseImportFactoryFunction() Function {
       panic("factory functions are not linked to reduce code bloat, please import 'gen/factory' instead of 'gen/meta'.")
}

func fillObject(n1 string, n2 string, item *TLItem)  {
	itemsByTag[item.tag] = item
	itemsByName[item.tlName] = item
	itemsByName[n1] = item
	itemsByName[n2] = item
	item.createObject = pleaseImportFactoryObject
	// code below is as fast, but allocates some extra strings which are already in binary const segment due to JSON code
	// itemsByName[fmt.Sprintf("%s#%08x", item.tlName, item.tag)] = item
	// itemsByName[fmt.Sprintf("#%08x", item.tag)] = item
}

func fillFunction(n1 string, n2 string, item *TLItem)  {
    fillObject(n1, n2, item)
	item.createFunction = pleaseImportFactoryFunction
}

func init() {
`)
	for _, wr := range typeWrappers {
		if wr.tlTag == 0 || !wr.isTopLevel {
			continue
		}
		if fun, ok := wr.trw.(*TypeRWStruct); ok && len(wr.NatParams) == 0 {
			if fun.ResultType != nil {
				qw422016.N().S(`fillFunction(`)
			} else {
				qw422016.N().S(`fillObject(`)
			}
			qw422016.N().S(`"`)
			wr.tlName.StreamString(qw422016)
			qw422016.N().S(`#`)
			qw422016.N().S(fmt.Sprintf("%08x", wr.tlTag))
			qw422016.N().S(`",`)
			qw422016.N().Q(fmt.Sprintf("#%08x", wr.tlTag))
			qw422016.N().S(`,&TLItem{tag:`)
			qw422016.N().S(fmt.Sprintf("%#x", wr.tlTag))
			qw422016.N().S(`, annotations:`)
			qw422016.N().S(fmt.Sprintf("%#x", wr.AnnotationsMask()))
			qw422016.N().S(`, tlName: "`)
			wr.tlName.StreamString(qw422016)
			qw422016.N().S(`"})`)
		}
		qw422016.N().S(`
`)
	}
	qw422016.N().S(`}

`)
}

func (gen *Gen2) writegenerateMeta(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	gen.streamgenerateMeta(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func (gen *Gen2) generateMeta() string {
	qb422016 := qt422016.AcquireByteBuffer()
	gen.writegenerateMeta(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}
