// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlTupleCycleTuple

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/cycle_b51088a4226835d54f08524a36f8aa77"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type TupleCycleTuple []cycle_b51088a4226835d54f08524a36f8aa77.CycleTuple

func (TupleCycleTuple) TLName() string { return "tuple" }
func (TupleCycleTuple) TLTag() uint32  { return 0x9770768a }

func (item *TupleCycleTuple) Reset() {
	ptr := (*[]cycle_b51088a4226835d54f08524a36f8aa77.CycleTuple)(item)
	*ptr = (*ptr)[:0]
}

func (item *TupleCycleTuple) FillRandom(rg *basictl.RandGenerator, nat_n uint32) {
	ptr := (*[]cycle_b51088a4226835d54f08524a36f8aa77.CycleTuple)(item)
	cycle_b51088a4226835d54f08524a36f8aa77.BuiltinTupleCycleTupleFillRandom(rg, ptr, nat_n)
}

func (item *TupleCycleTuple) Read(w []byte, nat_n uint32) (_ []byte, err error) {
	ptr := (*[]cycle_b51088a4226835d54f08524a36f8aa77.CycleTuple)(item)
	return cycle_b51088a4226835d54f08524a36f8aa77.BuiltinTupleCycleTupleRead(w, ptr, nat_n)
}

// This method is general version of Write, use it instead!
func (item *TupleCycleTuple) WriteGeneral(w []byte, nat_n uint32) (_ []byte, err error) {
	return item.Write(w, nat_n)
}

func (item *TupleCycleTuple) Write(w []byte, nat_n uint32) (_ []byte, err error) {
	ptr := (*[]cycle_b51088a4226835d54f08524a36f8aa77.CycleTuple)(item)
	return cycle_b51088a4226835d54f08524a36f8aa77.BuiltinTupleCycleTupleWrite(w, *ptr, nat_n)
}

func (item *TupleCycleTuple) ReadBoxed(w []byte, nat_n uint32) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x9770768a); err != nil {
		return w, err
	}
	return item.Read(w, nat_n)
}

// This method is general version of WriteBoxed, use it instead!
func (item *TupleCycleTuple) WriteBoxedGeneral(w []byte, nat_n uint32) (_ []byte, err error) {
	return item.WriteBoxed(w, nat_n)
}

func (item *TupleCycleTuple) WriteBoxed(w []byte, nat_n uint32) (_ []byte, err error) {
	w = basictl.NatWrite(w, 0x9770768a)
	return item.Write(w, nat_n)
}

func (item *TupleCycleTuple) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, nat_n uint32) error {
	ptr := (*[]cycle_b51088a4226835d54f08524a36f8aa77.CycleTuple)(item)
	if err := cycle_b51088a4226835d54f08524a36f8aa77.BuiltinTupleCycleTupleReadJSON(legacyTypeNames, in, ptr, nat_n); err != nil {
		return err
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *TupleCycleTuple) WriteJSONGeneral(w []byte, nat_n uint32) (_ []byte, err error) {
	return item.WriteJSON(w, nat_n)
}

func (item *TupleCycleTuple) WriteJSON(w []byte, nat_n uint32) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w, nat_n)
}

func (item *TupleCycleTuple) WriteJSONOpt(newTypeNames bool, short bool, w []byte, nat_n uint32) (_ []byte, err error) {
	ptr := (*[]cycle_b51088a4226835d54f08524a36f8aa77.CycleTuple)(item)
	if w, err = cycle_b51088a4226835d54f08524a36f8aa77.BuiltinTupleCycleTupleWriteJSONOpt(newTypeNames, short, w, *ptr, nat_n); err != nil {
		return w, err
	}
	return w, nil
}
