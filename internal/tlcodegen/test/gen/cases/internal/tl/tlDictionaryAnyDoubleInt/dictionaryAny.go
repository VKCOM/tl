// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlDictionaryAnyDoubleInt

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tl/tlBuiltinVectorDictionaryFieldAnyDoubleInt"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tl/tlDictionaryFieldAnyDoubleInt"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type DictionaryAnyDoubleInt []tlDictionaryFieldAnyDoubleInt.DictionaryFieldAnyDoubleInt

func (DictionaryAnyDoubleInt) TLName() string { return "dictionaryAny" }
func (DictionaryAnyDoubleInt) TLTag() uint32  { return 0x1f4c6190 }

func (item *DictionaryAnyDoubleInt) Reset() {
	ptr := (*[]tlDictionaryFieldAnyDoubleInt.DictionaryFieldAnyDoubleInt)(item)
	*ptr = (*ptr)[:0]
}

func (item *DictionaryAnyDoubleInt) FillRandom(rg *basictl.RandGenerator) {
	ptr := (*[]tlDictionaryFieldAnyDoubleInt.DictionaryFieldAnyDoubleInt)(item)
	tlBuiltinVectorDictionaryFieldAnyDoubleInt.BuiltinVectorDictionaryFieldAnyDoubleIntFillRandom(rg, ptr)
}

func (item *DictionaryAnyDoubleInt) Read(w []byte) (_ []byte, err error) {
	ptr := (*[]tlDictionaryFieldAnyDoubleInt.DictionaryFieldAnyDoubleInt)(item)
	return tlBuiltinVectorDictionaryFieldAnyDoubleInt.BuiltinVectorDictionaryFieldAnyDoubleIntRead(w, ptr)
}

// This method is general version of Write, use it instead!
func (item *DictionaryAnyDoubleInt) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *DictionaryAnyDoubleInt) Write(w []byte) []byte {
	ptr := (*[]tlDictionaryFieldAnyDoubleInt.DictionaryFieldAnyDoubleInt)(item)
	return tlBuiltinVectorDictionaryFieldAnyDoubleInt.BuiltinVectorDictionaryFieldAnyDoubleIntWrite(w, *ptr)
}

func (item *DictionaryAnyDoubleInt) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x1f4c6190); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *DictionaryAnyDoubleInt) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *DictionaryAnyDoubleInt) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x1f4c6190)
	return item.Write(w)
}

func (item DictionaryAnyDoubleInt) String() string {
	return string(item.WriteJSON(nil))
}

func (item *DictionaryAnyDoubleInt) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	ptr := (*[]tlDictionaryFieldAnyDoubleInt.DictionaryFieldAnyDoubleInt)(item)
	if err := tlBuiltinVectorDictionaryFieldAnyDoubleInt.BuiltinVectorDictionaryFieldAnyDoubleIntReadJSON(legacyTypeNames, in, ptr); err != nil {
		return err
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *DictionaryAnyDoubleInt) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSON(w), nil
}

func (item *DictionaryAnyDoubleInt) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}

func (item *DictionaryAnyDoubleInt) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	ptr := (*[]tlDictionaryFieldAnyDoubleInt.DictionaryFieldAnyDoubleInt)(item)
	w = tlBuiltinVectorDictionaryFieldAnyDoubleInt.BuiltinVectorDictionaryFieldAnyDoubleIntWriteJSONOpt(newTypeNames, short, w, *ptr)
	return w
}
func (item *DictionaryAnyDoubleInt) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *DictionaryAnyDoubleInt) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("dictionaryAny", err.Error())
	}
	return nil
}
