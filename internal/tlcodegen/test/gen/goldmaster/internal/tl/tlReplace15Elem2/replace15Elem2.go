// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlReplace15Elem2

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type Replace15Elem2 struct {
	X int32
	Y int32
	Z int32
}

func (item *Replace15Elem2) Reset() {
	item.X = 0
	item.Y = 0
	item.Z = 0
}

func (item *Replace15Elem2) FillRandom(rg *basictl.RandGenerator, nat_n uint32) {
	item.X = basictl.RandomInt(rg)
	item.Y = basictl.RandomInt(rg)
	item.Z = basictl.RandomInt(rg)
}

func (item *Replace15Elem2) Read(w []byte, nat_n uint32) (_ []byte, err error) {
	if w, err = basictl.IntRead(w, &item.X); err != nil {
		return w, err
	}
	if w, err = basictl.IntRead(w, &item.Y); err != nil {
		return w, err
	}
	return basictl.IntRead(w, &item.Z)
}

// This method is general version of Write, use it instead!
func (item *Replace15Elem2) WriteGeneral(w []byte, nat_n uint32) (_ []byte, err error) {
	return item.Write(w, nat_n), nil
}

func (item *Replace15Elem2) Write(w []byte, nat_n uint32) []byte {
	w = basictl.IntWrite(w, item.X)
	w = basictl.IntWrite(w, item.Y)
	w = basictl.IntWrite(w, item.Z)
	return w
}

func (item *Replace15Elem2) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, nat_n uint32) error {
	var propXPresented bool
	var propYPresented bool
	var propZPresented bool

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
					return internal.ErrorInvalidJSONWithDuplicatingKeys("replace15Elem2", "x")
				}
				if err := internal.Json2ReadInt32(in, &item.X); err != nil {
					return err
				}
				propXPresented = true
			case "y":
				if propYPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("replace15Elem2", "y")
				}
				if err := internal.Json2ReadInt32(in, &item.Y); err != nil {
					return err
				}
				propYPresented = true
			case "z":
				if propZPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("replace15Elem2", "z")
				}
				if err := internal.Json2ReadInt32(in, &item.Z); err != nil {
					return err
				}
				propZPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("replace15Elem2", key)
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
	if !propYPresented {
		item.Y = 0
	}
	if !propZPresented {
		item.Z = 0
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *Replace15Elem2) WriteJSONGeneral(w []byte, nat_n uint32) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w, nat_n), nil
}

func (item *Replace15Elem2) WriteJSON(w []byte, nat_n uint32) []byte {
	return item.WriteJSONOpt(true, false, w, nat_n)
}
func (item *Replace15Elem2) WriteJSONOpt(newTypeNames bool, short bool, w []byte, nat_n uint32) []byte {
	w = append(w, '{')
	backupIndexX := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"x":`...)
	w = basictl.JSONWriteInt32(w, item.X)
	if (item.X != 0) == false {
		w = w[:backupIndexX]
	}
	backupIndexY := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"y":`...)
	w = basictl.JSONWriteInt32(w, item.Y)
	if (item.Y != 0) == false {
		w = w[:backupIndexY]
	}
	backupIndexZ := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"z":`...)
	w = basictl.JSONWriteInt32(w, item.Z)
	if (item.Z != 0) == false {
		w = w[:backupIndexZ]
	}
	return append(w, '}')
}
