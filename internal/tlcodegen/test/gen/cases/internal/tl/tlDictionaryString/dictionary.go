// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlDictionaryString

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tl/tlBuiltinVectorDictionaryFieldString"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type DictionaryString map[string]string

func (DictionaryString) TLName() string { return "dictionary" }
func (DictionaryString) TLTag() uint32  { return 0x1f4c618f }

func (item *DictionaryString) Reset() {
	ptr := (*map[string]string)(item)
	tlBuiltinVectorDictionaryFieldString.BuiltinVectorDictionaryFieldStringReset(*ptr)
}

func (item *DictionaryString) FillRandom(rg *basictl.RandGenerator) {
	ptr := (*map[string]string)(item)
	tlBuiltinVectorDictionaryFieldString.BuiltinVectorDictionaryFieldStringFillRandom(rg, ptr)
}

func (item *DictionaryString) Read(w []byte) (_ []byte, err error) {
	ptr := (*map[string]string)(item)
	return tlBuiltinVectorDictionaryFieldString.BuiltinVectorDictionaryFieldStringRead(w, ptr)
}

// This method is general version of Write, use it instead!
func (item *DictionaryString) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *DictionaryString) Write(w []byte) []byte {
	ptr := (*map[string]string)(item)
	return tlBuiltinVectorDictionaryFieldString.BuiltinVectorDictionaryFieldStringWrite(w, *ptr)
}

func (item *DictionaryString) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x1f4c618f); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *DictionaryString) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *DictionaryString) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x1f4c618f)
	return item.Write(w)
}

func (item DictionaryString) String() string {
	return string(item.WriteJSON(nil))
}

func (item *DictionaryString) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	ptr := (*map[string]string)(item)
	if err := tlBuiltinVectorDictionaryFieldString.BuiltinVectorDictionaryFieldStringReadJSON(legacyTypeNames, in, ptr); err != nil {
		return err
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *DictionaryString) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSON(w), nil
}

func (item *DictionaryString) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}

func (item *DictionaryString) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	ptr := (*map[string]string)(item)
	w = tlBuiltinVectorDictionaryFieldString.BuiltinVectorDictionaryFieldStringWriteJSONOpt(newTypeNames, short, w, *ptr)
	return w
}
func (item *DictionaryString) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *DictionaryString) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("dictionary", err.Error())
	}
	return nil
}
