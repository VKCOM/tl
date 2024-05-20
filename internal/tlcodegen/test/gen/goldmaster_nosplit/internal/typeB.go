// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package internal

import (
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite

type TypeB struct {
	X int32
}

func (TypeB) TLName() string { return "typeB" }
func (TypeB) TLTag() uint32  { return 0x9d024802 }

func (item *TypeB) Reset() {
	item.X = 0
}

func (item *TypeB) FillRandom(rg *basictl.RandGenerator) {
	item.X = basictl.RandomInt(rg)
}

func (item *TypeB) Read(w []byte) (_ []byte, err error) {
	return basictl.IntRead(w, &item.X)
}

// This method is general version of Write, use it instead!
func (item *TypeB) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *TypeB) Write(w []byte) []byte {
	w = basictl.IntWrite(w, item.X)
	return w
}

func (item *TypeB) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x9d024802); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *TypeB) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *TypeB) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x9d024802)
	return item.Write(w)
}

func (item TypeB) String() string {
	return string(item.WriteJSON(nil))
}

func (item *TypeB) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
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
					return ErrorInvalidJSONWithDuplicatingKeys("typeB", "x")
				}
				if err := Json2ReadInt32(in, &item.X); err != nil {
					return err
				}
				propXPresented = true
			default:
				return ErrorInvalidJSONExcessElement("typeB", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propXPresented {
		item.X = 0
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *TypeB) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *TypeB) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *TypeB) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexX := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"x":`...)
	w = basictl.JSONWriteInt32(w, item.X)
	if (item.X != 0) == false {
		w = w[:backupIndexX]
	}
	return append(w, '}')
}

func (item *TypeB) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *TypeB) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("typeB", err.Error())
	}
	return nil
}
