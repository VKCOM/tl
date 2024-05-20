// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package internal

import (
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite

type Replace17 struct {
	X []int32
	A int32
	Y []int32
}

func (Replace17) TLName() string { return "replace17" }
func (Replace17) TLTag() uint32  { return 0xf46f9b9b }

func (item *Replace17) Reset() {
	item.X = item.X[:0]
	item.A = 0
	item.Y = item.Y[:0]
}

func (item *Replace17) FillRandom(rg *basictl.RandGenerator) {
	BuiltinVectorIntFillRandom(rg, &item.X)
	item.A = basictl.RandomInt(rg)
	BuiltinVectorIntFillRandom(rg, &item.Y)
}

func (item *Replace17) Read(w []byte) (_ []byte, err error) {
	if w, err = BuiltinVectorIntRead(w, &item.X); err != nil {
		return w, err
	}
	if w, err = basictl.IntRead(w, &item.A); err != nil {
		return w, err
	}
	return BuiltinVectorIntRead(w, &item.Y)
}

// This method is general version of Write, use it instead!
func (item *Replace17) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *Replace17) Write(w []byte) []byte {
	w = BuiltinVectorIntWrite(w, item.X)
	w = basictl.IntWrite(w, item.A)
	w = BuiltinVectorIntWrite(w, item.Y)
	return w
}

func (item *Replace17) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xf46f9b9b); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *Replace17) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *Replace17) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0xf46f9b9b)
	return item.Write(w)
}

func (item Replace17) String() string {
	return string(item.WriteJSON(nil))
}

func (item *Replace17) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propXPresented bool
	var propAPresented bool
	var propYPresented bool

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
					return ErrorInvalidJSONWithDuplicatingKeys("replace17", "x")
				}
				if err := BuiltinVectorIntReadJSON(legacyTypeNames, in, &item.X); err != nil {
					return err
				}
				propXPresented = true
			case "a":
				if propAPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("replace17", "a")
				}
				if err := Json2ReadInt32(in, &item.A); err != nil {
					return err
				}
				propAPresented = true
			case "y":
				if propYPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("replace17", "y")
				}
				if err := BuiltinVectorIntReadJSON(legacyTypeNames, in, &item.Y); err != nil {
					return err
				}
				propYPresented = true
			default:
				return ErrorInvalidJSONExcessElement("replace17", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propXPresented {
		item.X = item.X[:0]
	}
	if !propAPresented {
		item.A = 0
	}
	if !propYPresented {
		item.Y = item.Y[:0]
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *Replace17) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *Replace17) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *Replace17) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexX := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"x":`...)
	w = BuiltinVectorIntWriteJSONOpt(newTypeNames, short, w, item.X)
	if (len(item.X) != 0) == false {
		w = w[:backupIndexX]
	}
	backupIndexA := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"a":`...)
	w = basictl.JSONWriteInt32(w, item.A)
	if (item.A != 0) == false {
		w = w[:backupIndexA]
	}
	backupIndexY := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"y":`...)
	w = BuiltinVectorIntWriteJSONOpt(newTypeNames, short, w, item.Y)
	if (len(item.Y) != 0) == false {
		w = w[:backupIndexY]
	}
	return append(w, '}')
}

func (item *Replace17) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *Replace17) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("replace17", err.Error())
	}
	return nil
}
