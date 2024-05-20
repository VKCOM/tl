// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlVectorDictionaryFieldService1Value

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/cycle_6ca945392bbf8b14f24e5653edc8b214"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tl/tlBuiltinVectorDictionaryFieldService1Value"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type VectorDictionaryFieldService1Value map[string]cycle_6ca945392bbf8b14f24e5653edc8b214.Service1Value

func (VectorDictionaryFieldService1Value) TLName() string { return "vector" }
func (VectorDictionaryFieldService1Value) TLTag() uint32  { return 0x1cb5c415 }

func (item *VectorDictionaryFieldService1Value) Reset() {
	ptr := (*map[string]cycle_6ca945392bbf8b14f24e5653edc8b214.Service1Value)(item)
	tlBuiltinVectorDictionaryFieldService1Value.BuiltinVectorDictionaryFieldService1ValueReset(*ptr)
}

func (item *VectorDictionaryFieldService1Value) Read(w []byte) (_ []byte, err error) {
	ptr := (*map[string]cycle_6ca945392bbf8b14f24e5653edc8b214.Service1Value)(item)
	return tlBuiltinVectorDictionaryFieldService1Value.BuiltinVectorDictionaryFieldService1ValueRead(w, ptr)
}

// This method is general version of Write, use it instead!
func (item *VectorDictionaryFieldService1Value) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *VectorDictionaryFieldService1Value) Write(w []byte) []byte {
	ptr := (*map[string]cycle_6ca945392bbf8b14f24e5653edc8b214.Service1Value)(item)
	return tlBuiltinVectorDictionaryFieldService1Value.BuiltinVectorDictionaryFieldService1ValueWrite(w, *ptr)
}

func (item *VectorDictionaryFieldService1Value) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x1cb5c415); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *VectorDictionaryFieldService1Value) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *VectorDictionaryFieldService1Value) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x1cb5c415)
	return item.Write(w)
}

func (item VectorDictionaryFieldService1Value) String() string {
	return string(item.WriteJSON(nil))
}

func (item *VectorDictionaryFieldService1Value) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	ptr := (*map[string]cycle_6ca945392bbf8b14f24e5653edc8b214.Service1Value)(item)
	if err := tlBuiltinVectorDictionaryFieldService1Value.BuiltinVectorDictionaryFieldService1ValueReadJSON(legacyTypeNames, in, ptr); err != nil {
		return err
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *VectorDictionaryFieldService1Value) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSON(w), nil
}

func (item *VectorDictionaryFieldService1Value) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}

func (item *VectorDictionaryFieldService1Value) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	ptr := (*map[string]cycle_6ca945392bbf8b14f24e5653edc8b214.Service1Value)(item)
	w = tlBuiltinVectorDictionaryFieldService1Value.BuiltinVectorDictionaryFieldService1ValueWriteJSONOpt(newTypeNames, short, w, *ptr)
	return w
}
func (item *VectorDictionaryFieldService1Value) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *VectorDictionaryFieldService1Value) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("vector", err.Error())
	}
	return nil
}
