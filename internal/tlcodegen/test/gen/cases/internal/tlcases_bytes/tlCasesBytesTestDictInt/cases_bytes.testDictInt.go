// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlCasesBytesTestDictInt

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tl/tlBuiltinVectorDictionaryFieldAnyIntInt"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type CasesBytesTestDictInt struct {
	Dict map[int32]int32
}

func (CasesBytesTestDictInt) TLName() string { return "cases_bytes.testDictInt" }
func (CasesBytesTestDictInt) TLTag() uint32  { return 0x453ace07 }

func (item *CasesBytesTestDictInt) Reset() {
	tlBuiltinVectorDictionaryFieldAnyIntInt.BuiltinVectorDictionaryFieldAnyIntIntReset(item.Dict)
}

func (item *CasesBytesTestDictInt) FillRandom(rg *basictl.RandGenerator) {
	tlBuiltinVectorDictionaryFieldAnyIntInt.BuiltinVectorDictionaryFieldAnyIntIntFillRandom(rg, &item.Dict)
}

func (item *CasesBytesTestDictInt) Read(w []byte) (_ []byte, err error) {
	return tlBuiltinVectorDictionaryFieldAnyIntInt.BuiltinVectorDictionaryFieldAnyIntIntRead(w, &item.Dict)
}

// This method is general version of Write, use it instead!
func (item *CasesBytesTestDictInt) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *CasesBytesTestDictInt) Write(w []byte) []byte {
	w = tlBuiltinVectorDictionaryFieldAnyIntInt.BuiltinVectorDictionaryFieldAnyIntIntWrite(w, item.Dict)
	return w
}

func (item *CasesBytesTestDictInt) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x453ace07); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *CasesBytesTestDictInt) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *CasesBytesTestDictInt) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x453ace07)
	return item.Write(w)
}

func (item CasesBytesTestDictInt) String() string {
	return string(item.WriteJSON(nil))
}

func (item *CasesBytesTestDictInt) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propDictPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "dict":
				if propDictPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("cases_bytes.testDictInt", "dict")
				}
				if err := tlBuiltinVectorDictionaryFieldAnyIntInt.BuiltinVectorDictionaryFieldAnyIntIntReadJSON(legacyTypeNames, in, &item.Dict); err != nil {
					return err
				}
				propDictPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("cases_bytes.testDictInt", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propDictPresented {
		tlBuiltinVectorDictionaryFieldAnyIntInt.BuiltinVectorDictionaryFieldAnyIntIntReset(item.Dict)
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *CasesBytesTestDictInt) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *CasesBytesTestDictInt) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *CasesBytesTestDictInt) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexDict := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"dict":`...)
	w = tlBuiltinVectorDictionaryFieldAnyIntInt.BuiltinVectorDictionaryFieldAnyIntIntWriteJSONOpt(newTypeNames, short, w, item.Dict)
	if (len(item.Dict) != 0) == false {
		w = w[:backupIndexDict]
	}
	return append(w, '}')
}

func (item *CasesBytesTestDictInt) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *CasesBytesTestDictInt) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("cases_bytes.testDictInt", err.Error())
	}
	return nil
}
