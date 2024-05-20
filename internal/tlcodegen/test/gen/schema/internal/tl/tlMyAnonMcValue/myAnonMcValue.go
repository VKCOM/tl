// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlMyAnonMcValue

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/cycle_6ca945392bbf8b14f24e5653edc8b214"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type MyAnonMcValue cycle_6ca945392bbf8b14f24e5653edc8b214.Service1Value

func (MyAnonMcValue) TLName() string { return "myAnonMcValue" }
func (MyAnonMcValue) TLTag() uint32  { return 0x569310db }

func (item *MyAnonMcValue) Reset() {
	ptr := (*cycle_6ca945392bbf8b14f24e5653edc8b214.Service1Value)(item)
	ptr.Reset()
}

func (item *MyAnonMcValue) Read(w []byte) (_ []byte, err error) {
	ptr := (*cycle_6ca945392bbf8b14f24e5653edc8b214.Service1Value)(item)
	return ptr.ReadBoxed(w)
}

// This method is general version of Write, use it instead!
func (item *MyAnonMcValue) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *MyAnonMcValue) Write(w []byte) []byte {
	ptr := (*cycle_6ca945392bbf8b14f24e5653edc8b214.Service1Value)(item)
	return ptr.WriteBoxed(w)
}

func (item *MyAnonMcValue) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x569310db); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *MyAnonMcValue) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *MyAnonMcValue) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x569310db)
	return item.Write(w)
}

func (item MyAnonMcValue) String() string {
	return string(item.WriteJSON(nil))
}

func (item *MyAnonMcValue) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	ptr := (*cycle_6ca945392bbf8b14f24e5653edc8b214.Service1Value)(item)
	if err := ptr.ReadJSON(legacyTypeNames, in); err != nil {
		return err
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *MyAnonMcValue) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSON(w), nil
}

func (item *MyAnonMcValue) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}

func (item *MyAnonMcValue) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	ptr := (*cycle_6ca945392bbf8b14f24e5653edc8b214.Service1Value)(item)
	w = ptr.WriteJSONOpt(newTypeNames, short, w)
	return w
}
func (item *MyAnonMcValue) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *MyAnonMcValue) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("myAnonMcValue", err.Error())
	}
	return nil
}
