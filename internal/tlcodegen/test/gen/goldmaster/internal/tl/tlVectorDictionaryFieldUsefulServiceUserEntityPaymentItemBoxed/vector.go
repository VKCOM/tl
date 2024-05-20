// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlVectorDictionaryFieldUsefulServiceUserEntityPaymentItemBoxed

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlBuiltinVectorDictionaryFieldUsefulServiceUserEntityPaymentItemBoxed"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tlusefulService/tlUsefulServiceUserEntityPaymentItem"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type VectorDictionaryFieldUsefulServiceUserEntityPaymentItemBoxed map[string]tlUsefulServiceUserEntityPaymentItem.UsefulServiceUserEntityPaymentItem

func (VectorDictionaryFieldUsefulServiceUserEntityPaymentItemBoxed) TLName() string { return "vector" }
func (VectorDictionaryFieldUsefulServiceUserEntityPaymentItemBoxed) TLTag() uint32  { return 0x1cb5c415 }

func (item *VectorDictionaryFieldUsefulServiceUserEntityPaymentItemBoxed) Reset() {
	ptr := (*map[string]tlUsefulServiceUserEntityPaymentItem.UsefulServiceUserEntityPaymentItem)(item)
	tlBuiltinVectorDictionaryFieldUsefulServiceUserEntityPaymentItemBoxed.BuiltinVectorDictionaryFieldUsefulServiceUserEntityPaymentItemBoxedReset(*ptr)
}

func (item *VectorDictionaryFieldUsefulServiceUserEntityPaymentItemBoxed) FillRandom(rg *basictl.RandGenerator, nat_t uint32) {
	ptr := (*map[string]tlUsefulServiceUserEntityPaymentItem.UsefulServiceUserEntityPaymentItem)(item)
	tlBuiltinVectorDictionaryFieldUsefulServiceUserEntityPaymentItemBoxed.BuiltinVectorDictionaryFieldUsefulServiceUserEntityPaymentItemBoxedFillRandom(rg, ptr, nat_t)
}

func (item *VectorDictionaryFieldUsefulServiceUserEntityPaymentItemBoxed) Read(w []byte, nat_t uint32) (_ []byte, err error) {
	ptr := (*map[string]tlUsefulServiceUserEntityPaymentItem.UsefulServiceUserEntityPaymentItem)(item)
	return tlBuiltinVectorDictionaryFieldUsefulServiceUserEntityPaymentItemBoxed.BuiltinVectorDictionaryFieldUsefulServiceUserEntityPaymentItemBoxedRead(w, ptr, nat_t)
}

// This method is general version of Write, use it instead!
func (item *VectorDictionaryFieldUsefulServiceUserEntityPaymentItemBoxed) WriteGeneral(w []byte, nat_t uint32) (_ []byte, err error) {
	return item.Write(w, nat_t), nil
}

func (item *VectorDictionaryFieldUsefulServiceUserEntityPaymentItemBoxed) Write(w []byte, nat_t uint32) []byte {
	ptr := (*map[string]tlUsefulServiceUserEntityPaymentItem.UsefulServiceUserEntityPaymentItem)(item)
	return tlBuiltinVectorDictionaryFieldUsefulServiceUserEntityPaymentItemBoxed.BuiltinVectorDictionaryFieldUsefulServiceUserEntityPaymentItemBoxedWrite(w, *ptr, nat_t)
}

func (item *VectorDictionaryFieldUsefulServiceUserEntityPaymentItemBoxed) ReadBoxed(w []byte, nat_t uint32) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x1cb5c415); err != nil {
		return w, err
	}
	return item.Read(w, nat_t)
}

// This method is general version of WriteBoxed, use it instead!
func (item *VectorDictionaryFieldUsefulServiceUserEntityPaymentItemBoxed) WriteBoxedGeneral(w []byte, nat_t uint32) (_ []byte, err error) {
	return item.WriteBoxed(w, nat_t), nil
}

func (item *VectorDictionaryFieldUsefulServiceUserEntityPaymentItemBoxed) WriteBoxed(w []byte, nat_t uint32) []byte {
	w = basictl.NatWrite(w, 0x1cb5c415)
	return item.Write(w, nat_t)
}

func (item *VectorDictionaryFieldUsefulServiceUserEntityPaymentItemBoxed) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, nat_t uint32) error {
	ptr := (*map[string]tlUsefulServiceUserEntityPaymentItem.UsefulServiceUserEntityPaymentItem)(item)
	if err := tlBuiltinVectorDictionaryFieldUsefulServiceUserEntityPaymentItemBoxed.BuiltinVectorDictionaryFieldUsefulServiceUserEntityPaymentItemBoxedReadJSON(legacyTypeNames, in, ptr, nat_t); err != nil {
		return err
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *VectorDictionaryFieldUsefulServiceUserEntityPaymentItemBoxed) WriteJSONGeneral(w []byte, nat_t uint32) (_ []byte, err error) {
	return item.WriteJSON(w, nat_t), nil
}

func (item *VectorDictionaryFieldUsefulServiceUserEntityPaymentItemBoxed) WriteJSON(w []byte, nat_t uint32) []byte {
	return item.WriteJSONOpt(true, false, w, nat_t)
}

func (item *VectorDictionaryFieldUsefulServiceUserEntityPaymentItemBoxed) WriteJSONOpt(newTypeNames bool, short bool, w []byte, nat_t uint32) []byte {
	ptr := (*map[string]tlUsefulServiceUserEntityPaymentItem.UsefulServiceUserEntityPaymentItem)(item)
	w = tlBuiltinVectorDictionaryFieldUsefulServiceUserEntityPaymentItemBoxed.BuiltinVectorDictionaryFieldUsefulServiceUserEntityPaymentItemBoxedWriteJSONOpt(newTypeNames, short, w, *ptr, nat_t)
	return w
}
