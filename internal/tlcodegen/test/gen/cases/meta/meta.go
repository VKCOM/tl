// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package meta

import (
	"fmt"

	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal"
	"github.com/vkcom/tl/pkg/basictl"
)

func SchemaGenerator() string { return "(devel)" }
func SchemaURL() string       { return "" }
func SchemaCommit() string    { return "" }
func SchemaTimestamp() uint32 { return 0 }

// We can create only types which have zero type arguments and zero nat arguments
type Object interface {
	TLName() string // returns type's TL name. For union, returns constructor name depending on actual union value
	TLTag() uint32  // returns type's TL tag. For union, returns constructor tag depending on actual union value
	String() string // returns type's representation for debugging (JSON for now)

	FillRandom(rg *basictl.RandGenerator)
	Read(w []byte) ([]byte, error)              // reads type's bare TL representation by consuming bytes from the start of w and returns remaining bytes, plus error
	ReadBoxed(w []byte) ([]byte, error)         // same as Read, but reads/checks TLTag first (this method is general version of Write, use it only when you are working with interface)
	WriteGeneral(w []byte) ([]byte, error)      // appends bytes of type's bare TL representation to the end of w and returns it, plus error
	WriteBoxedGeneral(w []byte) ([]byte, error) // same as Write, but writes TLTag first (this method is general version of WriteBoxed, use it only when you are working with interface)

	MarshalJSON() ([]byte, error) // returns type's JSON representation, plus error
	UnmarshalJSON([]byte) error   // reads type's JSON representation

	ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error
	WriteJSONGeneral(w []byte) ([]byte, error) // like MarshalJSON, but appends to w and returns it (this method is general version of WriteBoxed, use it only when you are working with interface)
}

type Function interface {
	Object

	ReadResultWriteResultJSON(r []byte, w []byte) ([]byte, []byte, error) // combination of ReadResult(r) + WriteResultJSON(w). Returns new r, new w, plus error
	ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) // combination of ReadResultJSON(r) + WriteResult(w). Returns new r, new w, plus error

	// For transcoding short-long version during Long ID and newTypeNames transition
	ReadResultWriteResultJSONOpt(newTypeNames bool, short bool, r []byte, w []byte) ([]byte, []byte, error)
}

func GetAllTLItems() []TLItem {
	var allItems []TLItem
	for _, item := range itemsByName {
		if item != nil {
			allItems = append(allItems, *item)
		}
	}
	return allItems
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

func CreateFunctionBytes(tag uint32) Function {
	if item := FactoryItemByTLTag(tag); item != nil && item.createFunctionBytes != nil {
		return item.createFunctionBytes()
	}
	return nil
}

func CreateObjectBytes(tag uint32) Object {
	if item := FactoryItemByTLTag(tag); item != nil && item.createObjectBytes != nil {
		return item.createObjectBytes()
	}
	return nil
}

// name can be in any of 3 forms "ch_proxy.insert#7cf362ba", "ch_proxy.insert" or "#7cf362ba"
func CreateFunctionFromNameBytes(name string) Function {
	if item := FactoryItemByTLName(name); item != nil && item.createFunctionBytes != nil {
		return item.createFunctionBytes()
	}
	return nil
}

// name can be in any of 3 forms "ch_proxy.insert#7cf362ba", "ch_proxy.insert" or "#7cf362ba"
func CreateObjectFromNameBytes(name string) Object {
	if item := FactoryItemByTLName(name); item != nil && item.createObjectBytes != nil {
		return item.createObjectBytes()
	}
	return nil
}

type TLItem struct {
	tag         uint32
	annotations uint32
	tlName      string

	resultTypeContainsUnionTypes bool

	createFunction          func() Function
	createFunctionLong      func() Function
	createObject            func() Object
	createFunctionBytes     func() Function
	createFunctionLongBytes func() Function
	createObjectBytes       func() Object
}

func (item TLItem) TLTag() uint32            { return item.tag }
func (item TLItem) TLName() string           { return item.tlName }
func (item TLItem) CreateObject() Object     { return item.createObject() }
func (item TLItem) IsFunction() bool         { return item.createFunction != nil }
func (item TLItem) CreateFunction() Function { return item.createFunction() }

func (item TLItem) HasResultTypeContainUnionTypes() bool { return item.resultTypeContainsUnionTypes }

// For transcoding short-long version during Long ID transition
func (item TLItem) HasFunctionLong() bool        { return item.createFunctionLong != nil }
func (item TLItem) CreateFunctionLong() Function { return item.createFunctionLong() }

// Annotations

// TLItem serves as a single type for all enum values
func (item *TLItem) Reset()                                {}
func (item *TLItem) FillRandom(rg *basictl.RandGenerator)  {}
func (item *TLItem) Read(w []byte) ([]byte, error)         { return w, nil }
func (item *TLItem) WriteGeneral(w []byte) ([]byte, error) { return w, nil }
func (item *TLItem) Write(w []byte) []byte                 { return w }
func (item *TLItem) ReadBoxed(w []byte) ([]byte, error)    { return basictl.NatReadExactTag(w, item.tag) }
func (item *TLItem) WriteBoxedGeneral(w []byte) ([]byte, error) {
	return basictl.NatWrite(w, item.tag), nil
}
func (item *TLItem) WriteBoxed(w []byte) []byte { return basictl.NatWrite(w, item.tag) }
func (item TLItem) String() string {
	return string(item.WriteJSON(nil))
}
func (item *TLItem) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	in.Delim('{')
	if !in.Ok() {
		return in.Error()
	}
	for !in.IsDelim('}') {
		return internal.ErrorInvalidJSONExcessElement(item.tlName, in.UnsafeFieldName(true))
	}
	in.Delim('}')
	if !in.Ok() {
		return in.Error()
	}
	return nil
}
func (item *TLItem) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSON(w), nil
}
func (item *TLItem) WriteJSON(w []byte) []byte {
	w = append(w, '{')
	return append(w, '}')
}
func (item *TLItem) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}
func (item *TLItem) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
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

var itemsByTag = map[uint32]*TLItem{}

var itemsByName = map[string]*TLItem{}

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

func SetGlobalFactoryCreateForFunctionBytes(itemTag uint32, createObject func() Object, createFunction func() Function, createFunctionLong func() Function) {
	item := itemsByTag[itemTag]
	if item == nil {
		panic(fmt.Sprintf("factory cannot find function tag #%08x to set", itemTag))
	}
	item.createObjectBytes = createObject
	item.createFunctionBytes = createFunction
	item.createFunctionLongBytes = createFunctionLong
}

func SetGlobalFactoryCreateForObjectBytes(itemTag uint32, createObject func() Object) {
	item := itemsByTag[itemTag]
	if item == nil {
		panic(fmt.Sprintf("factory cannot find item tag #%08x to set", itemTag))
	}
	item.createObjectBytes = createObject
}

func SetGlobalFactoryCreateForEnumElementBytes(itemTag uint32) {
	item := itemsByTag[itemTag]
	if item == nil {
		panic(fmt.Sprintf("factory cannot find enum tag #%08x to set", itemTag))
	}
	item.createObjectBytes = func() Object { return item }
}

func pleaseImportFactoryBytesObject() Object {
	panic("factory functions are not linked to reduce code bloat, please import 'gen/factory_bytes' instead of 'gen/meta'.")
}

func pleaseImportFactoryBytesFunction() Function {
	panic("factory functions are not linked to reduce code bloat, please import 'gen/factory_bytes' instead of 'gen/meta'.")
}

func pleaseImportFactoryObject() Object {
	panic("factory functions are not linked to reduce code bloat, please import 'gen/factory' instead of 'gen/meta'.")
}

func pleaseImportFactoryFunction() Function {
	panic("factory functions are not linked to reduce code bloat, please import 'gen/factory' instead of 'gen/meta'.")
}

func fillObject(n1 string, n2 string, item *TLItem) {
	itemsByTag[item.tag] = item
	itemsByName[item.tlName] = item
	itemsByName[n1] = item
	itemsByName[n2] = item
	item.createObject = pleaseImportFactoryObject
	item.createObjectBytes = pleaseImportFactoryBytesObject
	// code below is as fast, but allocates some extra strings which are already in binary const segment due to JSON code
	// itemsByName[fmt.Sprintf("%s#%08x", item.tlName, item.tag)] = item
	// itemsByName[fmt.Sprintf("#%08x", item.tag)] = item
}

func fillFunction(n1 string, n2 string, item *TLItem) {
	fillObject(n1, n2, item)
	item.createFunction = pleaseImportFactoryFunction
	item.createFunctionBytes = pleaseImportFactoryBytesFunction
}

func init() {
	fillObject("benchmarks.vruhash#d31bd0fd", "#d31bd0fd", &TLItem{tag: 0xd31bd0fd, annotations: 0x0, tlName: "benchmarks.vruhash", resultTypeContainsUnionTypes: false})
	fillObject("benchmarks.vruposition#32792c04", "#32792c04", &TLItem{tag: 0x32792c04, annotations: 0x0, tlName: "benchmarks.vruposition", resultTypeContainsUnionTypes: false})
	fillObject("benchmarks.vrutoyTopLevelContainer#fb442ca5", "#fb442ca5", &TLItem{tag: 0xfb442ca5, annotations: 0x0, tlName: "benchmarks.vrutoyTopLevelContainer", resultTypeContainsUnionTypes: false})
	fillObject("benchmarks.vrutoyTopLevelContainerWithDependency#c176008e", "#c176008e", &TLItem{tag: 0xc176008e, annotations: 0x0, tlName: "benchmarks.vrutoyTopLevelContainerWithDependency", resultTypeContainsUnionTypes: false})
	fillObject("benchmarks.vrutoytopLevelUnionBig#ef556bee", "#ef556bee", &TLItem{tag: 0xef556bee, annotations: 0x0, tlName: "benchmarks.vrutoytopLevelUnionBig", resultTypeContainsUnionTypes: false})
	fillObject("benchmarks.vrutoytopLevelUnionEmpty#ce27c770", "#ce27c770", &TLItem{tag: 0xce27c770, annotations: 0x0, tlName: "benchmarks.vrutoytopLevelUnionEmpty", resultTypeContainsUnionTypes: false})
	fillObject("cases_bytes.testArray#3762fb81", "#3762fb81", &TLItem{tag: 0x3762fb81, annotations: 0x0, tlName: "cases_bytes.testArray", resultTypeContainsUnionTypes: false})
	fillObject("cases_bytes.testDictAny#5a5fce57", "#5a5fce57", &TLItem{tag: 0x5a5fce57, annotations: 0x0, tlName: "cases_bytes.testDictAny", resultTypeContainsUnionTypes: false})
	fillObject("cases_bytes.testDictInt#453ace07", "#453ace07", &TLItem{tag: 0x453ace07, annotations: 0x0, tlName: "cases_bytes.testDictInt", resultTypeContainsUnionTypes: false})
	fillObject("cases_bytes.testDictString#6c04d6ce", "#6c04d6ce", &TLItem{tag: 0x6c04d6ce, annotations: 0x0, tlName: "cases_bytes.testDictString", resultTypeContainsUnionTypes: false})
	fillObject("cases_bytes.testDictStringString#ad69c772", "#ad69c772", &TLItem{tag: 0xad69c772, annotations: 0x0, tlName: "cases_bytes.testDictStringString", resultTypeContainsUnionTypes: false})
	fillObject("cases_bytes.testEnum1#58aad3f5", "#58aad3f5", &TLItem{tag: 0x58aad3f5, annotations: 0x0, tlName: "cases_bytes.testEnum1", resultTypeContainsUnionTypes: false})
	fillObject("cases_bytes.testEnum2#00b47add", "#00b47add", &TLItem{tag: 0x00b47add, annotations: 0x0, tlName: "cases_bytes.testEnum2", resultTypeContainsUnionTypes: false})
	fillObject("cases_bytes.testEnum3#81911ffa", "#81911ffa", &TLItem{tag: 0x81911ffa, annotations: 0x0, tlName: "cases_bytes.testEnum3", resultTypeContainsUnionTypes: false})
	fillObject("cases_bytes.testEnumContainer#32b92037", "#32b92037", &TLItem{tag: 0x32b92037, annotations: 0x0, tlName: "cases_bytes.testEnumContainer", resultTypeContainsUnionTypes: false})
	fillObject("cases_bytes.testTuple#2dd3bacf", "#2dd3bacf", &TLItem{tag: 0x2dd3bacf, annotations: 0x0, tlName: "cases_bytes.testTuple", resultTypeContainsUnionTypes: false})
	fillObject("cases_bytes.testVector#3647c8ae", "#3647c8ae", &TLItem{tag: 0x3647c8ae, annotations: 0x0, tlName: "cases_bytes.testVector", resultTypeContainsUnionTypes: false})
	fillObject("cases.myCycle1#d3ca919d", "#d3ca919d", &TLItem{tag: 0xd3ca919d, annotations: 0x0, tlName: "cases.myCycle1", resultTypeContainsUnionTypes: false})
	fillObject("cases.myCycle2#5444c9a2", "#5444c9a2", &TLItem{tag: 0x5444c9a2, annotations: 0x0, tlName: "cases.myCycle2", resultTypeContainsUnionTypes: false})
	fillObject("cases.myCycle3#7624f86b", "#7624f86b", &TLItem{tag: 0x7624f86b, annotations: 0x0, tlName: "cases.myCycle3", resultTypeContainsUnionTypes: false})
	fillObject("cases.replace7#6ccce4be", "#6ccce4be", &TLItem{tag: 0x6ccce4be, annotations: 0x0, tlName: "cases.replace7", resultTypeContainsUnionTypes: false})
	fillObject("cases.replace7plus#197858f5", "#197858f5", &TLItem{tag: 0x197858f5, annotations: 0x0, tlName: "cases.replace7plus", resultTypeContainsUnionTypes: false})
	fillObject("cases.replace7plusplus#abc39b68", "#abc39b68", &TLItem{tag: 0xabc39b68, annotations: 0x0, tlName: "cases.replace7plusplus", resultTypeContainsUnionTypes: false})
	fillObject("cases.testAllPossibleFieldConfigsContainer#e3fae936", "#e3fae936", &TLItem{tag: 0xe3fae936, annotations: 0x0, tlName: "cases.testAllPossibleFieldConfigsContainer", resultTypeContainsUnionTypes: false})
	fillObject("cases.testArray#a888030d", "#a888030d", &TLItem{tag: 0xa888030d, annotations: 0x0, tlName: "cases.testArray", resultTypeContainsUnionTypes: false})
	fillObject("cases.testBeforeReadBitValidation#9b2396db", "#9b2396db", &TLItem{tag: 0x9b2396db, annotations: 0x0, tlName: "cases.testBeforeReadBitValidation", resultTypeContainsUnionTypes: false})
	fillObject("cases.testDictAny#e29b8ae6", "#e29b8ae6", &TLItem{tag: 0xe29b8ae6, annotations: 0x0, tlName: "cases.testDictAny", resultTypeContainsUnionTypes: false})
	fillObject("cases.testDictInt#d3877643", "#d3877643", &TLItem{tag: 0xd3877643, annotations: 0x0, tlName: "cases.testDictInt", resultTypeContainsUnionTypes: false})
	fillObject("cases.testDictString#c463c79b", "#c463c79b", &TLItem{tag: 0xc463c79b, annotations: 0x0, tlName: "cases.testDictString", resultTypeContainsUnionTypes: false})
	fillObject("cases.testEnum1#6c6c55ac", "#6c6c55ac", &TLItem{tag: 0x6c6c55ac, annotations: 0x0, tlName: "cases.testEnum1", resultTypeContainsUnionTypes: false})
	fillObject("cases.testEnum2#86ea88ce", "#86ea88ce", &TLItem{tag: 0x86ea88ce, annotations: 0x0, tlName: "cases.testEnum2", resultTypeContainsUnionTypes: false})
	fillObject("cases.testEnum3#69b83e2f", "#69b83e2f", &TLItem{tag: 0x69b83e2f, annotations: 0x0, tlName: "cases.testEnum3", resultTypeContainsUnionTypes: false})
	fillObject("cases.testEnumContainer#cb684231", "#cb684231", &TLItem{tag: 0xcb684231, annotations: 0x0, tlName: "cases.testEnumContainer", resultTypeContainsUnionTypes: false})
	fillObject("cases.testLocalFieldmask#f68fd3f9", "#f68fd3f9", &TLItem{tag: 0xf68fd3f9, annotations: 0x0, tlName: "cases.testLocalFieldmask", resultTypeContainsUnionTypes: false})
	fillObject("cases.testMaybe#d6602613", "#d6602613", &TLItem{tag: 0xd6602613, annotations: 0x0, tlName: "cases.testMaybe", resultTypeContainsUnionTypes: false})
	fillObject("cases.testOutFieldMaskContainer#1850ffe4", "#1850ffe4", &TLItem{tag: 0x1850ffe4, annotations: 0x0, tlName: "cases.testOutFieldMaskContainer", resultTypeContainsUnionTypes: false})
	fillObject("cases.testRecursiveFieldMask#c58cf85e", "#c58cf85e", &TLItem{tag: 0xc58cf85e, annotations: 0x0, tlName: "cases.testRecursiveFieldMask", resultTypeContainsUnionTypes: false})
	fillObject("cases.testTuple#4b9caf8f", "#4b9caf8f", &TLItem{tag: 0x4b9caf8f, annotations: 0x0, tlName: "cases.testTuple", resultTypeContainsUnionTypes: false})
	fillObject("cases.testUnion1#4b4f09b1", "#4b4f09b1", &TLItem{tag: 0x4b4f09b1, annotations: 0x0, tlName: "cases.testUnion1", resultTypeContainsUnionTypes: false})
	fillObject("cases.testUnion2#464f96c4", "#464f96c4", &TLItem{tag: 0x464f96c4, annotations: 0x0, tlName: "cases.testUnion2", resultTypeContainsUnionTypes: false})
	fillObject("cases.testUnionContainer#4497a381", "#4497a381", &TLItem{tag: 0x4497a381, annotations: 0x0, tlName: "cases.testUnionContainer", resultTypeContainsUnionTypes: false})
	fillObject("cases.testVector#4975695c", "#4975695c", &TLItem{tag: 0x4975695c, annotations: 0x0, tlName: "cases.testVector", resultTypeContainsUnionTypes: false})
	fillObject("int#a8509bda", "#a8509bda", &TLItem{tag: 0xa8509bda, annotations: 0x0, tlName: "int", resultTypeContainsUnionTypes: false})
	fillObject("int32#7934e71f", "#7934e71f", &TLItem{tag: 0x7934e71f, annotations: 0x0, tlName: "int32", resultTypeContainsUnionTypes: false})
	fillObject("int64#f5609de0", "#f5609de0", &TLItem{tag: 0xf5609de0, annotations: 0x0, tlName: "int64", resultTypeContainsUnionTypes: false})
	fillObject("long#22076cba", "#22076cba", &TLItem{tag: 0x22076cba, annotations: 0x0, tlName: "long", resultTypeContainsUnionTypes: false})
	fillObject("string#b5286e24", "#b5286e24", &TLItem{tag: 0xb5286e24, annotations: 0x0, tlName: "string", resultTypeContainsUnionTypes: false})
	fillObject("true#3fedd339", "#3fedd339", &TLItem{tag: 0x3fedd339, annotations: 0x0, tlName: "true", resultTypeContainsUnionTypes: false})
}
