// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package meta

import (
	"fmt"

	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/pkg/basictl"
)

func SchemaGenerator() string { return "(devel)" }
func SchemaURL() string {
	return "https://github.com/VKCOM/tl/blob/master/internal/tlcodegen/test/tls/goldmaster.tl"
}
func SchemaCommit() string    { return "abcdefgh" }
func SchemaTimestamp() uint32 { return 301822800 }

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
func (item TLItem) AnnotationRead() bool      { return item.annotations&0x1 != 0 }
func (item TLItem) AnnotationReadwrite() bool { return item.annotations&0x2 != 0 }

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
	fillObject("a.blue#623360f3", "#623360f3", &TLItem{tag: 0x623360f3, annotations: 0x0, tlName: "a.blue", resultTypeContainsUnionTypes: false})
	fillObject("a.color#f35d7a69", "#f35d7a69", &TLItem{tag: 0xf35d7a69, annotations: 0x0, tlName: "a.color", resultTypeContainsUnionTypes: false})
	fillObject("a.green#6127e7b8", "#6127e7b8", &TLItem{tag: 0x6127e7b8, annotations: 0x0, tlName: "a.green", resultTypeContainsUnionTypes: false})
	fillObject("a.red#b83a723d", "#b83a723d", &TLItem{tag: 0xb83a723d, annotations: 0x0, tlName: "a.red", resultTypeContainsUnionTypes: false})
	fillObject("a.top2#7082d18f", "#7082d18f", &TLItem{tag: 0x7082d18f, annotations: 0x0, tlName: "a.top2", resultTypeContainsUnionTypes: false})
	fillObject("a.uNionA#a7662843", "#a7662843", &TLItem{tag: 0xa7662843, annotations: 0x0, tlName: "a.uNionA", resultTypeContainsUnionTypes: false})
	fillObject("ab.alias#944aaa97", "#944aaa97", &TLItem{tag: 0x944aaa97, annotations: 0x0, tlName: "ab.alias", resultTypeContainsUnionTypes: false})
	fillFunction("ab.call1#20c5fb2d", "#20c5fb2d", &TLItem{tag: 0x20c5fb2d, annotations: 0x1, tlName: "ab.call1", resultTypeContainsUnionTypes: false})
	fillFunction("ab.call2#77d5f057", "#77d5f057", &TLItem{tag: 0x77d5f057, annotations: 0x1, tlName: "ab.call2", resultTypeContainsUnionTypes: false})
	fillFunction("ab.call3#0a083445", "#0a083445", &TLItem{tag: 0x0a083445, annotations: 0x1, tlName: "ab.call3", resultTypeContainsUnionTypes: false})
	fillFunction("ab.call4#c1220a1e", "#c1220a1e", &TLItem{tag: 0xc1220a1e, annotations: 0x1, tlName: "ab.call4", resultTypeContainsUnionTypes: false})
	fillFunction("ab.call5#7ba4d28d", "#7ba4d28d", &TLItem{tag: 0x7ba4d28d, annotations: 0x1, tlName: "ab.call5", resultTypeContainsUnionTypes: false})
	fillFunction("ab.call6#84d815cb", "#84d815cb", &TLItem{tag: 0x84d815cb, annotations: 0x1, tlName: "ab.call6", resultTypeContainsUnionTypes: false})
	fillFunction("ab.call7#46ec10bf", "#46ec10bf", &TLItem{tag: 0x46ec10bf, annotations: 0x1, tlName: "ab.call7", resultTypeContainsUnionTypes: false})
	fillFunction("ab.call8#1b8652d9", "#1b8652d9", &TLItem{tag: 0x1b8652d9, annotations: 0x1, tlName: "ab.call8", resultTypeContainsUnionTypes: false})
	fillFunction("ab.call9#75de906c", "#75de906c", &TLItem{tag: 0x75de906c, annotations: 0x1, tlName: "ab.call9", resultTypeContainsUnionTypes: false})
	fillObject("ab.code#7651b1ac", "#7651b1ac", &TLItem{tag: 0x7651b1ac, annotations: 0x0, tlName: "ab.code", resultTypeContainsUnionTypes: false})
	fillObject("ab.counterChangeRequestPeriodsMany#14a35d80", "#14a35d80", &TLItem{tag: 0x14a35d80, annotations: 0x0, tlName: "ab.counterChangeRequestPeriodsMany", resultTypeContainsUnionTypes: false})
	fillObject("ab.counterChangeRequestPeriodsOne#d9c36de5", "#d9c36de5", &TLItem{tag: 0xd9c36de5, annotations: 0x0, tlName: "ab.counterChangeRequestPeriodsOne", resultTypeContainsUnionTypes: false})
	fillObject("ab.empty#1ec6a63e", "#1ec6a63e", &TLItem{tag: 0x1ec6a63e, annotations: 0x0, tlName: "ab.empty", resultTypeContainsUnionTypes: false})
	fillObject("ab.myType#e0e96c86", "#e0e96c86", &TLItem{tag: 0xe0e96c86, annotations: 0x0, tlName: "ab.myType", resultTypeContainsUnionTypes: false})
	fillObject("ab.testMaybe#4dac492a", "#4dac492a", &TLItem{tag: 0x4dac492a, annotations: 0x0, tlName: "ab.testMaybe", resultTypeContainsUnionTypes: false})
	fillObject("ab.topLevel1#e67bce28", "#e67bce28", &TLItem{tag: 0xe67bce28, annotations: 0x0, tlName: "ab.topLevel1", resultTypeContainsUnionTypes: false})
	fillObject("ab.topLevel2#cef933fb", "#cef933fb", &TLItem{tag: 0xcef933fb, annotations: 0x0, tlName: "ab.topLevel2", resultTypeContainsUnionTypes: false})
	fillObject("ab.typeA#a99fef6a", "#a99fef6a", &TLItem{tag: 0xa99fef6a, annotations: 0x0, tlName: "ab.typeA", resultTypeContainsUnionTypes: false})
	fillObject("ab.typeB#ff2e6d58", "#ff2e6d58", &TLItem{tag: 0xff2e6d58, annotations: 0x0, tlName: "ab.typeB", resultTypeContainsUnionTypes: false})
	fillObject("ab.typeC#69920d6e", "#69920d6e", &TLItem{tag: 0x69920d6e, annotations: 0x0, tlName: "ab.typeC", resultTypeContainsUnionTypes: false})
	fillObject("ab.typeD#76615bf1", "#76615bf1", &TLItem{tag: 0x76615bf1, annotations: 0x0, tlName: "ab.typeD", resultTypeContainsUnionTypes: false})
	fillObject("ab.useCycle#71687381", "#71687381", &TLItem{tag: 0x71687381, annotations: 0x0, tlName: "ab.useCycle", resultTypeContainsUnionTypes: false})
	fillObject("ab.useDictString#3325d884", "#3325d884", &TLItem{tag: 0x3325d884, annotations: 0x0, tlName: "ab.useDictString", resultTypeContainsUnionTypes: false})
	fillObject("au.nionA#df61f632", "#df61f632", &TLItem{tag: 0xdf61f632, annotations: 0x0, tlName: "au.nionA", resultTypeContainsUnionTypes: false})
	fillObject("b.red#a9471844", "#a9471844", &TLItem{tag: 0xa9471844, annotations: 0x0, tlName: "b.red", resultTypeContainsUnionTypes: false})
	fillFunction("call1#a7302fbc", "#a7302fbc", &TLItem{tag: 0xa7302fbc, annotations: 0x1, tlName: "call1", resultTypeContainsUnionTypes: false})
	fillFunction("call2#f02024c6", "#f02024c6", &TLItem{tag: 0xf02024c6, annotations: 0x1, tlName: "call2", resultTypeContainsUnionTypes: false})
	fillFunction("call3#6ace6718", "#6ace6718", &TLItem{tag: 0x6ace6718, annotations: 0x1, tlName: "call3", resultTypeContainsUnionTypes: false})
	fillFunction("call4#46d7de8f", "#46d7de8f", &TLItem{tag: 0x46d7de8f, annotations: 0x1, tlName: "call4", resultTypeContainsUnionTypes: false})
	fillFunction("call5#fc51061c", "#fc51061c", &TLItem{tag: 0xfc51061c, annotations: 0x1, tlName: "call5", resultTypeContainsUnionTypes: false})
	fillFunction("call6#e41e4696", "#e41e4696", &TLItem{tag: 0xe41e4696, annotations: 0x1, tlName: "call6", resultTypeContainsUnionTypes: false})
	fillFunction("call7#262a43e2", "#262a43e2", &TLItem{tag: 0x262a43e2, annotations: 0x1, tlName: "call7", resultTypeContainsUnionTypes: false})
	fillFunction("call8#7b400184", "#7b400184", &TLItem{tag: 0x7b400184, annotations: 0x1, tlName: "call8", resultTypeContainsUnionTypes: false})
	fillFunction("call9#67a0d62d", "#67a0d62d", &TLItem{tag: 0x67a0d62d, annotations: 0x1, tlName: "call9", resultTypeContainsUnionTypes: false})
	fillObject("cd.myType#eab6a6b4", "#eab6a6b4", &TLItem{tag: 0xeab6a6b4, annotations: 0x0, tlName: "cd.myType", resultTypeContainsUnionTypes: false})
	fillObject("cd.response#8c202f64", "#8c202f64", &TLItem{tag: 0x8c202f64, annotations: 0x0, tlName: "cd.response", resultTypeContainsUnionTypes: false})
	fillObject("cd.topLevel3#5cd1ca89", "#5cd1ca89", &TLItem{tag: 0x5cd1ca89, annotations: 0x0, tlName: "cd.topLevel3", resultTypeContainsUnionTypes: false})
	fillObject("cd.typeA#a831a920", "#a831a920", &TLItem{tag: 0xa831a920, annotations: 0x0, tlName: "cd.typeA", resultTypeContainsUnionTypes: false})
	fillObject("cd.typeB#377b4996", "#377b4996", &TLItem{tag: 0x377b4996, annotations: 0x0, tlName: "cd.typeB", resultTypeContainsUnionTypes: false})
	fillObject("cd.typeC#db0f93d4", "#db0f93d4", &TLItem{tag: 0xdb0f93d4, annotations: 0x0, tlName: "cd.typeC", resultTypeContainsUnionTypes: false})
	fillObject("cd.typeD#b5528285", "#b5528285", &TLItem{tag: 0xb5528285, annotations: 0x0, tlName: "cd.typeD", resultTypeContainsUnionTypes: false})
	fillObject("cd.useCycle#6ed67ca0", "#6ed67ca0", &TLItem{tag: 0x6ed67ca0, annotations: 0x0, tlName: "cd.useCycle", resultTypeContainsUnionTypes: false})
	fillObject("cyc1.myCycle#136ecc9e", "#136ecc9e", &TLItem{tag: 0x136ecc9e, annotations: 0x0, tlName: "cyc1.myCycle", resultTypeContainsUnionTypes: false})
	fillObject("cyc2.myCycle#fba5eecb", "#fba5eecb", &TLItem{tag: 0xfba5eecb, annotations: 0x0, tlName: "cyc2.myCycle", resultTypeContainsUnionTypes: false})
	fillObject("cyc3.myCycle#47866860", "#47866860", &TLItem{tag: 0x47866860, annotations: 0x0, tlName: "cyc3.myCycle", resultTypeContainsUnionTypes: false})
	fillObject("cycleTuple#c867fae3", "#c867fae3", &TLItem{tag: 0xc867fae3, annotations: 0x0, tlName: "cycleTuple", resultTypeContainsUnionTypes: false})
	fillObject("halfStr#647ddaf5", "#647ddaf5", &TLItem{tag: 0x647ddaf5, annotations: 0x0, tlName: "halfStr", resultTypeContainsUnionTypes: false})
	fillObject("hren#12ab5219", "#12ab5219", &TLItem{tag: 0x12ab5219, annotations: 0x0, tlName: "hren", resultTypeContainsUnionTypes: false})
	fillObject("int#a8509bda", "#a8509bda", &TLItem{tag: 0xa8509bda, annotations: 0x0, tlName: "int", resultTypeContainsUnionTypes: false})
	fillObject("int32#7934e71f", "#7934e71f", &TLItem{tag: 0x7934e71f, annotations: 0x0, tlName: "int32", resultTypeContainsUnionTypes: false})
	fillObject("int64#f5609de0", "#f5609de0", &TLItem{tag: 0xf5609de0, annotations: 0x0, tlName: "int64", resultTypeContainsUnionTypes: false})
	fillObject("long#22076cba", "#22076cba", &TLItem{tag: 0x22076cba, annotations: 0x0, tlName: "long", resultTypeContainsUnionTypes: false})
	fillObject("maybeTest1#c457763c", "#c457763c", &TLItem{tag: 0xc457763c, annotations: 0x0, tlName: "maybeTest1", resultTypeContainsUnionTypes: false})
	fillObject("multiPoint#0e1ae81e", "#0e1ae81e", &TLItem{tag: 0x0e1ae81e, annotations: 0x0, tlName: "multiPoint", resultTypeContainsUnionTypes: false})
	fillObject("myInt32#ba59e151", "#ba59e151", &TLItem{tag: 0xba59e151, annotations: 0x0, tlName: "myInt32", resultTypeContainsUnionTypes: false})
	fillObject("myInt64#1d95db9d", "#1d95db9d", &TLItem{tag: 0x1d95db9d, annotations: 0x0, tlName: "myInt64", resultTypeContainsUnionTypes: false})
	fillObject("myNat#c60c1b41", "#c60c1b41", &TLItem{tag: 0xc60c1b41, annotations: 0x0, tlName: "myNat", resultTypeContainsUnionTypes: false})
	fillObject("myPlus#79e0c6df", "#79e0c6df", &TLItem{tag: 0x79e0c6df, annotations: 0x0, tlName: "myPlus", resultTypeContainsUnionTypes: false})
	fillObject("myPlus3#692c291b", "#692c291b", &TLItem{tag: 0x692c291b, annotations: 0x0, tlName: "myPlus3", resultTypeContainsUnionTypes: false})
	fillObject("myZero#8d868379", "#8d868379", &TLItem{tag: 0x8d868379, annotations: 0x0, tlName: "myZero", resultTypeContainsUnionTypes: false})
	fillObject("myZero3#103a40cf", "#103a40cf", &TLItem{tag: 0x103a40cf, annotations: 0x0, tlName: "myZero3", resultTypeContainsUnionTypes: false})
	fillObject("nativeWrappers#344ddf50", "#344ddf50", &TLItem{tag: 0x344ddf50, annotations: 0x0, tlName: "nativeWrappers", resultTypeContainsUnionTypes: false})
	fillObject("noStr#3a728324", "#3a728324", &TLItem{tag: 0x3a728324, annotations: 0x0, tlName: "noStr", resultTypeContainsUnionTypes: false})
	fillObject("replace#323db63e", "#323db63e", &TLItem{tag: 0x323db63e, annotations: 0x0, tlName: "replace", resultTypeContainsUnionTypes: false})
	fillObject("replace10#fc81f008", "#fc81f008", &TLItem{tag: 0xfc81f008, annotations: 0x0, tlName: "replace10", resultTypeContainsUnionTypes: false})
	fillObject("replace12#ec121094", "#ec121094", &TLItem{tag: 0xec121094, annotations: 0x0, tlName: "replace12", resultTypeContainsUnionTypes: false})
	fillObject("replace15#2280e430", "#2280e430", &TLItem{tag: 0x2280e430, annotations: 0x0, tlName: "replace15", resultTypeContainsUnionTypes: false})
	fillObject("replace17#f46f9b9b", "#f46f9b9b", &TLItem{tag: 0xf46f9b9b, annotations: 0x0, tlName: "replace17", resultTypeContainsUnionTypes: false})
	fillObject("replace18#704dd712", "#704dd712", &TLItem{tag: 0x704dd712, annotations: 0x0, tlName: "replace18", resultTypeContainsUnionTypes: false})
	fillObject("replace2#e2d4ebee", "#e2d4ebee", &TLItem{tag: 0xe2d4ebee, annotations: 0x0, tlName: "replace2", resultTypeContainsUnionTypes: false})
	fillObject("replace3#51e324e4", "#51e324e4", &TLItem{tag: 0x51e324e4, annotations: 0x0, tlName: "replace3", resultTypeContainsUnionTypes: false})
	fillObject("replace5#8b5bc78a", "#8b5bc78a", &TLItem{tag: 0x8b5bc78a, annotations: 0x0, tlName: "replace5", resultTypeContainsUnionTypes: false})
	fillObject("replace6#abd49d06", "#abd49d06", &TLItem{tag: 0xabd49d06, annotations: 0x0, tlName: "replace6", resultTypeContainsUnionTypes: false})
	fillObject("replace7#f4c66d9f", "#f4c66d9f", &TLItem{tag: 0xf4c66d9f, annotations: 0x0, tlName: "replace7", resultTypeContainsUnionTypes: false})
	fillObject("replace8#d626c117", "#d626c117", &TLItem{tag: 0xd626c117, annotations: 0x0, tlName: "replace8", resultTypeContainsUnionTypes: false})
	fillObject("replace9#95d598c5", "#95d598c5", &TLItem{tag: 0x95d598c5, annotations: 0x0, tlName: "replace9", resultTypeContainsUnionTypes: false})
	fillObject("service5.emptyOutput#ff8f7db8", "#ff8f7db8", &TLItem{tag: 0xff8f7db8, annotations: 0x0, tlName: "service5.emptyOutput", resultTypeContainsUnionTypes: false})
	fillFunction("service5.insert#7cf362ba", "#7cf362ba", &TLItem{tag: 0x7cf362ba, annotations: 0x2, tlName: "service5.insert", resultTypeContainsUnionTypes: true})
	fillObject("service5Long.emptyOutput#ff8f7db9", "#ff8f7db9", &TLItem{tag: 0xff8f7db9, annotations: 0x0, tlName: "service5Long.emptyOutput", resultTypeContainsUnionTypes: false})
	fillFunction("service5Long.insert#7cf362bb", "#7cf362bb", &TLItem{tag: 0x7cf362bb, annotations: 0x2, tlName: "service5Long.insert", resultTypeContainsUnionTypes: true})
	fillObject("service5Long.stringOutput#dc170ff5", "#dc170ff5", &TLItem{tag: 0xdc170ff5, annotations: 0x0, tlName: "service5Long.stringOutput", resultTypeContainsUnionTypes: false})
	fillObject("service5.stringOutput#dc170ff4", "#dc170ff4", &TLItem{tag: 0xdc170ff4, annotations: 0x0, tlName: "service5.stringOutput", resultTypeContainsUnionTypes: false})
	fillObject("string#b5286e24", "#b5286e24", &TLItem{tag: 0xb5286e24, annotations: 0x0, tlName: "string", resultTypeContainsUnionTypes: false})
	fillObject("testMaybe#88920e90", "#88920e90", &TLItem{tag: 0x88920e90, annotations: 0x0, tlName: "testMaybe", resultTypeContainsUnionTypes: false})
	fillObject("testMaybe2#0aa03cf2", "#0aa03cf2", &TLItem{tag: 0x0aa03cf2, annotations: 0x0, tlName: "testMaybe2", resultTypeContainsUnionTypes: false})
	fillObject("true#3fedd339", "#3fedd339", &TLItem{tag: 0x3fedd339, annotations: 0x0, tlName: "true", resultTypeContainsUnionTypes: false})
	fillObject("typeA#157673c1", "#157673c1", &TLItem{tag: 0x157673c1, annotations: 0x0, tlName: "typeA", resultTypeContainsUnionTypes: false})
	fillObject("typeB#9d024802", "#9d024802", &TLItem{tag: 0x9d024802, annotations: 0x0, tlName: "typeB", resultTypeContainsUnionTypes: false})
	fillObject("typeC#6b8ef43f", "#6b8ef43f", &TLItem{tag: 0x6b8ef43f, annotations: 0x0, tlName: "typeC", resultTypeContainsUnionTypes: false})
	fillObject("typeD#b1f4369e", "#b1f4369e", &TLItem{tag: 0xb1f4369e, annotations: 0x0, tlName: "typeD", resultTypeContainsUnionTypes: false})
	fillObject("unionArgsUse#742161d2", "#742161d2", &TLItem{tag: 0x742161d2, annotations: 0x0, tlName: "unionArgsUse", resultTypeContainsUnionTypes: false})
	fillObject("useDictUgly#fb9ce817", "#fb9ce817", &TLItem{tag: 0xfb9ce817, annotations: 0x0, tlName: "useDictUgly", resultTypeContainsUnionTypes: false})
	fillObject("useResponse#0a63ec5f", "#0a63ec5f", &TLItem{tag: 0x0a63ec5f, annotations: 0x0, tlName: "useResponse", resultTypeContainsUnionTypes: false})
	fillObject("useStr#9aa3dee5", "#9aa3dee5", &TLItem{tag: 0x9aa3dee5, annotations: 0x0, tlName: "useStr", resultTypeContainsUnionTypes: false})
	fillObject("useTrue#dfdd4180", "#dfdd4180", &TLItem{tag: 0xdfdd4180, annotations: 0x0, tlName: "useTrue", resultTypeContainsUnionTypes: false})
	fillFunction("usefulService.getUserEntity#3c857e52", "#3c857e52", &TLItem{tag: 0x3c857e52, annotations: 0x2, tlName: "usefulService.getUserEntity", resultTypeContainsUnionTypes: false})
}
