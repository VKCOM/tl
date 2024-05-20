// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlMyMcValue

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/cycle_6ca945392bbf8b14f24e5653edc8b214"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type MyMcValue struct {
	X cycle_6ca945392bbf8b14f24e5653edc8b214.Service1Value
}

func (MyMcValue) TLName() string { return "myMcValue" }
func (MyMcValue) TLTag() uint32  { return 0xe2ffd978 }

func (item *MyMcValue) Reset() {
	item.X.Reset()
}

func (item *MyMcValue) Read(w []byte) (_ []byte, err error) {
	return item.X.ReadBoxed(w)
}

// This method is general version of Write, use it instead!
func (item *MyMcValue) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *MyMcValue) Write(w []byte) []byte {
	w = item.X.WriteBoxed(w)
	return w
}

func (item *MyMcValue) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xe2ffd978); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *MyMcValue) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *MyMcValue) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0xe2ffd978)
	return item.Write(w)
}

func (item MyMcValue) String() string {
	return string(item.WriteJSON(nil))
}

func (item *MyMcValue) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propXPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "x":
				if propXPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("myMcValue", "x")
				}
				if err := item.X.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propXPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("myMcValue", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propXPresented {
		item.X.Reset()
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *MyMcValue) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *MyMcValue) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *MyMcValue) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"x":`...)
	w = item.X.WriteJSONOpt(newTypeNames, short, w)
	return append(w, '}')
}

func (item *MyMcValue) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *MyMcValue) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("myMcValue", err.Error())
	}
	return nil
}
