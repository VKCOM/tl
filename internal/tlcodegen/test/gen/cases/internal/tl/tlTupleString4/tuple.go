// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlTupleString4

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tl/tlBuiltinTuple4String"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type TupleString4 [4]string

func (TupleString4) TLName() string { return "tuple" }
func (TupleString4) TLTag() uint32  { return 0x9770768a }

func (item *TupleString4) Reset() {
	ptr := (*[4]string)(item)
	tlBuiltinTuple4String.BuiltinTuple4StringReset(ptr)
}

func (item *TupleString4) FillRandom(rg *basictl.RandGenerator) {
	ptr := (*[4]string)(item)
	tlBuiltinTuple4String.BuiltinTuple4StringFillRandom(rg, ptr)
}

func (item *TupleString4) Read(w []byte) (_ []byte, err error) {
	ptr := (*[4]string)(item)
	return tlBuiltinTuple4String.BuiltinTuple4StringRead(w, ptr)
}

// This method is general version of Write, use it instead!
func (item *TupleString4) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *TupleString4) Write(w []byte) []byte {
	ptr := (*[4]string)(item)
	return tlBuiltinTuple4String.BuiltinTuple4StringWrite(w, ptr)
}

func (item *TupleString4) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x9770768a); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *TupleString4) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *TupleString4) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x9770768a)
	return item.Write(w)
}

func (item TupleString4) String() string {
	return string(item.WriteJSON(nil))
}

func (item *TupleString4) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	ptr := (*[4]string)(item)
	if err := tlBuiltinTuple4String.BuiltinTuple4StringReadJSON(legacyTypeNames, in, ptr); err != nil {
		return err
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *TupleString4) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSON(w), nil
}

func (item *TupleString4) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}

func (item *TupleString4) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	ptr := (*[4]string)(item)
	w = tlBuiltinTuple4String.BuiltinTuple4StringWriteJSONOpt(newTypeNames, short, w, ptr)
	return w
}
func (item *TupleString4) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *TupleString4) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("tuple", err.Error())
	}
	return nil
}
