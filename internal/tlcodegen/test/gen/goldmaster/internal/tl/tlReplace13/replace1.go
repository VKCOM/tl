// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlReplace13

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlBuiltinTuple3Int"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type Replace13 struct {
	A [3]int32
}

func (Replace13) TLName() string { return "replace1" }
func (Replace13) TLTag() uint32  { return 0x89eac43a }

func (item *Replace13) Reset() {
	tlBuiltinTuple3Int.BuiltinTuple3IntReset(&item.A)
}

func (item *Replace13) FillRandom(rg *basictl.RandGenerator) {
	tlBuiltinTuple3Int.BuiltinTuple3IntFillRandom(rg, &item.A)
}

func (item *Replace13) Read(w []byte) (_ []byte, err error) {
	return tlBuiltinTuple3Int.BuiltinTuple3IntRead(w, &item.A)
}

// This method is general version of Write, use it instead!
func (item *Replace13) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *Replace13) Write(w []byte) []byte {
	w = tlBuiltinTuple3Int.BuiltinTuple3IntWrite(w, &item.A)
	return w
}

func (item *Replace13) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x89eac43a); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *Replace13) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *Replace13) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x89eac43a)
	return item.Write(w)
}

func (item Replace13) String() string {
	return string(item.WriteJSON(nil))
}

func (item *Replace13) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
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
					return internal.ErrorInvalidJSONWithDuplicatingKeys("replace1", "a")
				}
				if err := tlBuiltinTuple3Int.BuiltinTuple3IntReadJSON(legacyTypeNames, in, &item.A); err != nil {
					return err
				}
				propAPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("replace1", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propAPresented {
		tlBuiltinTuple3Int.BuiltinTuple3IntReset(&item.A)
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *Replace13) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *Replace13) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *Replace13) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"a":`...)
	w = tlBuiltinTuple3Int.BuiltinTuple3IntWriteJSONOpt(newTypeNames, short, w, &item.A)
	return append(w, '}')
}

func (item *Replace13) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *Replace13) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("replace1", err.Error())
	}
	return nil
}
