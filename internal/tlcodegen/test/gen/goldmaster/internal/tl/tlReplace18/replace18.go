// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlReplace18

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlBuiltinVectorVectorVectorInt"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type Replace18 struct {
	A [][][]int32
}

func (Replace18) TLName() string { return "replace18" }
func (Replace18) TLTag() uint32  { return 0x704dd712 }

func (item *Replace18) Reset() {
	item.A = item.A[:0]
}

func (item *Replace18) FillRandom(rg *basictl.RandGenerator) {
	tlBuiltinVectorVectorVectorInt.BuiltinVectorVectorVectorIntFillRandom(rg, &item.A)
}

func (item *Replace18) Read(w []byte) (_ []byte, err error) {
	return tlBuiltinVectorVectorVectorInt.BuiltinVectorVectorVectorIntRead(w, &item.A)
}

// This method is general version of Write, use it instead!
func (item *Replace18) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *Replace18) Write(w []byte) []byte {
	w = tlBuiltinVectorVectorVectorInt.BuiltinVectorVectorVectorIntWrite(w, item.A)
	return w
}

func (item *Replace18) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x704dd712); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *Replace18) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *Replace18) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x704dd712)
	return item.Write(w)
}

func (item Replace18) String() string {
	return string(item.WriteJSON(nil))
}

func (item *Replace18) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propAPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "a":
				if propAPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("replace18", "a")
				}
				if err := tlBuiltinVectorVectorVectorInt.BuiltinVectorVectorVectorIntReadJSON(legacyTypeNames, in, &item.A); err != nil {
					return err
				}
				propAPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("replace18", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propAPresented {
		item.A = item.A[:0]
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *Replace18) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *Replace18) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *Replace18) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexA := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"a":`...)
	w = tlBuiltinVectorVectorVectorInt.BuiltinVectorVectorVectorIntWriteJSONOpt(newTypeNames, short, w, item.A)
	if (len(item.A) != 0) == false {
		w = w[:backupIndexA]
	}
	return append(w, '}')
}

func (item *Replace18) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *Replace18) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("replace18", err.Error())
	}
	return nil
}
