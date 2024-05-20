// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlDictionaryFieldAnyDoubleInt

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type DictionaryFieldAnyDoubleInt struct {
	Key   float64
	Value int32
}

func (DictionaryFieldAnyDoubleInt) TLName() string { return "dictionaryFieldAny" }
func (DictionaryFieldAnyDoubleInt) TLTag() uint32  { return 0x2c43a65b }

func (item *DictionaryFieldAnyDoubleInt) Reset() {
	item.Key = 0
	item.Value = 0
}

func (item *DictionaryFieldAnyDoubleInt) FillRandom(rg *basictl.RandGenerator) {
	item.Key = basictl.RandomDouble(rg)
	item.Value = basictl.RandomInt(rg)
}

func (item *DictionaryFieldAnyDoubleInt) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.DoubleRead(w, &item.Key); err != nil {
		return w, err
	}
	return basictl.IntRead(w, &item.Value)
}

// This method is general version of Write, use it instead!
func (item *DictionaryFieldAnyDoubleInt) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *DictionaryFieldAnyDoubleInt) Write(w []byte) []byte {
	w = basictl.DoubleWrite(w, item.Key)
	w = basictl.IntWrite(w, item.Value)
	return w
}

func (item *DictionaryFieldAnyDoubleInt) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x2c43a65b); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *DictionaryFieldAnyDoubleInt) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *DictionaryFieldAnyDoubleInt) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x2c43a65b)
	return item.Write(w)
}

func (item DictionaryFieldAnyDoubleInt) String() string {
	return string(item.WriteJSON(nil))
}

func (item *DictionaryFieldAnyDoubleInt) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propKeyPresented bool
	var propValuePresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "key":
				if propKeyPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("dictionaryFieldAny", "key")
				}
				if err := internal.Json2ReadFloat64(in, &item.Key); err != nil {
					return err
				}
				propKeyPresented = true
			case "value":
				if propValuePresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("dictionaryFieldAny", "value")
				}
				if err := internal.Json2ReadInt32(in, &item.Value); err != nil {
					return err
				}
				propValuePresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("dictionaryFieldAny", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propKeyPresented {
		item.Key = 0
	}
	if !propValuePresented {
		item.Value = 0
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *DictionaryFieldAnyDoubleInt) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *DictionaryFieldAnyDoubleInt) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *DictionaryFieldAnyDoubleInt) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexKey := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"key":`...)
	w = basictl.JSONWriteFloat64(w, item.Key)
	if (item.Key != 0) == false {
		w = w[:backupIndexKey]
	}
	backupIndexValue := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"value":`...)
	w = basictl.JSONWriteInt32(w, item.Value)
	if (item.Value != 0) == false {
		w = w[:backupIndexValue]
	}
	return append(w, '}')
}

func (item *DictionaryFieldAnyDoubleInt) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *DictionaryFieldAnyDoubleInt) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("dictionaryFieldAny", err.Error())
	}
	return nil
}
