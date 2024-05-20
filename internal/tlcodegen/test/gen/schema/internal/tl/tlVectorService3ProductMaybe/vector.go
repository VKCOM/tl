// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlVectorService3ProductMaybe

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tl/tlBuiltinVectorService3Product"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tlservice3/tlService3Product"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type VectorService3ProductMaybe struct {
	Value []tlService3Product.Service3Product // not deterministic if !Ok
	Ok    bool
}

func (item *VectorService3ProductMaybe) Reset() {
	item.Ok = false
}

func (item *VectorService3ProductMaybe) ReadBoxed(w []byte, nat_t uint32) (_ []byte, err error) {
	if w, err = basictl.ReadBool(w, &item.Ok, 0x27930a7b, 0x3f9c8ef8); err != nil {
		return w, err
	}
	if item.Ok {
		return tlBuiltinVectorService3Product.BuiltinVectorService3ProductRead(w, &item.Value, nat_t)
	}
	return w, nil
}

// This method is general version of WriteBoxed, use it instead!
func (item *VectorService3ProductMaybe) WriteBoxedGeneral(w []byte, nat_t uint32) (_ []byte, err error) {
	return item.WriteBoxed(w, nat_t), nil
}

func (item *VectorService3ProductMaybe) WriteBoxed(w []byte, nat_t uint32) []byte {
	if item.Ok {
		w = basictl.NatWrite(w, 0x3f9c8ef8)
		return tlBuiltinVectorService3Product.BuiltinVectorService3ProductWrite(w, item.Value, nat_t)
	}
	return basictl.NatWrite(w, 0x27930a7b)
}

func (item *VectorService3ProductMaybe) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, nat_t uint32) error {
	_ok, _jvalue, err := internal.Json2ReadMaybe("Maybe", in)
	if err != nil {
		return err
	}
	item.Ok = _ok
	if _ok {
		var in2Pointer *basictl.JsonLexer
		if _jvalue != nil {
			in2 := basictl.JsonLexer{Data: _jvalue}
			in2Pointer = &in2
		}
		if err := tlBuiltinVectorService3Product.BuiltinVectorService3ProductReadJSON(legacyTypeNames, in2Pointer, &item.Value, nat_t); err != nil {
			return err
		}
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *VectorService3ProductMaybe) WriteJSONGeneral(w []byte, nat_t uint32) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w, nat_t), nil
}

func (item *VectorService3ProductMaybe) WriteJSON(w []byte, nat_t uint32) []byte {
	return item.WriteJSONOpt(true, false, w, nat_t)
}
func (item *VectorService3ProductMaybe) WriteJSONOpt(newTypeNames bool, short bool, w []byte, nat_t uint32) []byte {
	if !item.Ok {
		return append(w, "{}"...)
	}
	w = append(w, `{"ok":true`...)
	if len(item.Value) != 0 {
		w = append(w, `,"value":`...)
		w = tlBuiltinVectorService3Product.BuiltinVectorService3ProductWriteJSONOpt(newTypeNames, short, w, item.Value, nat_t)
	}
	return append(w, '}')
}
