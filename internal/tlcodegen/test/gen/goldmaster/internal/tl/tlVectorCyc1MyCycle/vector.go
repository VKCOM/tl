// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlVectorCyc1MyCycle

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/cycle_e10cb78db8a2766007111b86ce9e11d9"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type VectorCyc1MyCycle []cycle_e10cb78db8a2766007111b86ce9e11d9.Cyc1MyCycle

func (VectorCyc1MyCycle) TLName() string { return "vector" }
func (VectorCyc1MyCycle) TLTag() uint32  { return 0x1cb5c415 }

func (item *VectorCyc1MyCycle) Reset() {
	ptr := (*[]cycle_e10cb78db8a2766007111b86ce9e11d9.Cyc1MyCycle)(item)
	*ptr = (*ptr)[:0]
}

func (item *VectorCyc1MyCycle) FillRandom(rg *basictl.RandGenerator) {
	ptr := (*[]cycle_e10cb78db8a2766007111b86ce9e11d9.Cyc1MyCycle)(item)
	cycle_e10cb78db8a2766007111b86ce9e11d9.BuiltinVectorCyc1MyCycleFillRandom(rg, ptr)
}

func (item *VectorCyc1MyCycle) Read(w []byte) (_ []byte, err error) {
	ptr := (*[]cycle_e10cb78db8a2766007111b86ce9e11d9.Cyc1MyCycle)(item)
	return cycle_e10cb78db8a2766007111b86ce9e11d9.BuiltinVectorCyc1MyCycleRead(w, ptr)
}

// This method is general version of Write, use it instead!
func (item *VectorCyc1MyCycle) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *VectorCyc1MyCycle) Write(w []byte) []byte {
	ptr := (*[]cycle_e10cb78db8a2766007111b86ce9e11d9.Cyc1MyCycle)(item)
	return cycle_e10cb78db8a2766007111b86ce9e11d9.BuiltinVectorCyc1MyCycleWrite(w, *ptr)
}

func (item *VectorCyc1MyCycle) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x1cb5c415); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *VectorCyc1MyCycle) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *VectorCyc1MyCycle) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x1cb5c415)
	return item.Write(w)
}

func (item VectorCyc1MyCycle) String() string {
	return string(item.WriteJSON(nil))
}

func (item *VectorCyc1MyCycle) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	ptr := (*[]cycle_e10cb78db8a2766007111b86ce9e11d9.Cyc1MyCycle)(item)
	if err := cycle_e10cb78db8a2766007111b86ce9e11d9.BuiltinVectorCyc1MyCycleReadJSON(legacyTypeNames, in, ptr); err != nil {
		return err
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *VectorCyc1MyCycle) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSON(w), nil
}

func (item *VectorCyc1MyCycle) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}

func (item *VectorCyc1MyCycle) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	ptr := (*[]cycle_e10cb78db8a2766007111b86ce9e11d9.Cyc1MyCycle)(item)
	w = cycle_e10cb78db8a2766007111b86ce9e11d9.BuiltinVectorCyc1MyCycleWriteJSONOpt(newTypeNames, short, w, *ptr)
	return w
}
func (item *VectorCyc1MyCycle) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *VectorCyc1MyCycle) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("vector", err.Error())
	}
	return nil
}
