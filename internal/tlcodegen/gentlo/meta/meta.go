// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package meta

import (
	"fmt"

	"github.com/vkcom/tl/internal/tlcodegen/gentlo/basictl"
	"github.com/vkcom/tl/internal/tlcodegen/gentlo/internal"
)

// We can create only types which have zero type arguments and zero nat arguments
type Object interface {
	TLName() string // returns type's TL name. For union, returns constructor name depending on actual union value
	TLTag() uint32  // returns type's TL tag. For union, returns constructor tag depending on actual union value
	String() string // returns type's representation for debugging (JSON for now)

	Read(w []byte) ([]byte, error)       // reads type's bare TL representation by consuming bytes from the start of w and returns remaining bytes, plus error
	Write(w []byte) ([]byte, error)      // appends bytes of type's bare TL representation to the end of w and returns it, plus error
	ReadBoxed(w []byte) ([]byte, error)  // same as Read, but reads/checks TLTag first
	WriteBoxed(w []byte) ([]byte, error) // same as Write, but writes TLTag first

	MarshalJSON() ([]byte, error)       // returns type's JSON representation, plus error
	UnmarshalJSON([]byte) error         // reads type's JSON representation
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
	tlName             string
	createFunction     func() Function
	createFunctionLong func() Function
	createObject       func() Object
	// TODO - annotations, etc
}

func (item TLItem) TLTag() uint32            { return item.tag }
func (item TLItem) TLName() string           { return item.tlName }
func (item TLItem) CreateObject() Object     { return item.createObject() }
func (item TLItem) IsFunction() bool         { return item.createFunction != nil }
func (item TLItem) CreateFunction() Function { return item.createFunction() }

// For transcoding short-long version during Long ID transition
func (item TLItem) HasFunctionLong() bool        { return item.createFunctionLong != nil }
func (item TLItem) CreateFunctionLong() Function { return item.createFunctionLong() }

// TLItem serves as a single type for all enum values
func (item *TLItem) Reset()                              {}
func (item *TLItem) Read(w []byte) ([]byte, error)       { return w, nil }
func (item *TLItem) Write(w []byte) ([]byte, error)      { return w, nil }
func (item *TLItem) ReadBoxed(w []byte) ([]byte, error)  { return basictl.NatReadExactTag(w, item.tag) }
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
	// code below is as fast, but allocates some extra strings which are already in binary const segment due to JSON code
	// itemsByName[fmt.Sprintf("%s#%08x", item.tlName, item.tag)] = item
	// itemsByName[fmt.Sprintf("#%08x", item.tag)] = item
}

func fillFunction(n1 string, n2 string, item *TLItem) {
	fillObject(n1, n2, item)
	item.createFunction = pleaseImportFactoryFunction
}

func init() {
	fillObject("tls.arg#29dfe61b", "#29dfe61b", &TLItem{tag: 0x29dfe61b, tlName: "tls.arg"})
	fillObject("tls.array#d9fb20de", "#d9fb20de", &TLItem{tag: 0xd9fb20de, tlName: "tls.array"})
	fillObject("tls.combinator#5c0a1ed5", "#5c0a1ed5", &TLItem{tag: 0x5c0a1ed5, tlName: "tls.combinator"})
	fillObject("tls.combinatorLeft#4c12c6d9", "#4c12c6d9", &TLItem{tag: 0x4c12c6d9, tlName: "tls.combinatorLeft"})
	fillObject("tls.combinatorLeftBuiltin#cd211f63", "#cd211f63", &TLItem{tag: 0xcd211f63, tlName: "tls.combinatorLeftBuiltin"})
	fillObject("tls.combinatorRight#2c064372", "#2c064372", &TLItem{tag: 0x2c064372, tlName: "tls.combinatorRight"})
	fillObject("tls.combinator_v4#e91692d5", "#e91692d5", &TLItem{tag: 0xe91692d5, tlName: "tls.combinator_v4"})
	fillObject("tls.exprNat#dcb49bd8", "#dcb49bd8", &TLItem{tag: 0xdcb49bd8, tlName: "tls.exprNat"})
	fillObject("tls.exprType#ecc9da78", "#ecc9da78", &TLItem{tag: 0xecc9da78, tlName: "tls.exprType"})
	fillObject("tls.natConst#8ce940b1", "#8ce940b1", &TLItem{tag: 0x8ce940b1, tlName: "tls.natConst"})
	fillObject("tls.natVar#4e8a14f0", "#4e8a14f0", &TLItem{tag: 0x4e8a14f0, tlName: "tls.natVar"})
	fillObject("tls.schema_v2#3a2f9be2", "#3a2f9be2", &TLItem{tag: 0x3a2f9be2, tlName: "tls.schema_v2"})
	fillObject("tls.schema_v3#e4a8604b", "#e4a8604b", &TLItem{tag: 0xe4a8604b, tlName: "tls.schema_v3"})
	fillObject("tls.schema_v4#90ac88d7", "#90ac88d7", &TLItem{tag: 0x90ac88d7, tlName: "tls.schema_v4"})
	fillObject("tls.type#12eb4386", "#12eb4386", &TLItem{tag: 0x12eb4386, tlName: "tls.type"})
	fillObject("tls.typeExpr#c1863d08", "#c1863d08", &TLItem{tag: 0xc1863d08, tlName: "tls.typeExpr"})
	fillObject("tls.typeVar#0142ceae", "#0142ceae", &TLItem{tag: 0x142ceae, tlName: "tls.typeVar"})
}
