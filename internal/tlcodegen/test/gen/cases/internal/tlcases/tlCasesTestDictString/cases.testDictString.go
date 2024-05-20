// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlCasesTestDictString

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tl/tlBuiltinVectorDictionaryFieldInt"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type CasesTestDictString struct {
	Dict map[string]int32
}

func (CasesTestDictString) TLName() string { return "cases.testDictString" }
func (CasesTestDictString) TLTag() uint32  { return 0xc463c79b }

func (item *CasesTestDictString) Reset() {
	tlBuiltinVectorDictionaryFieldInt.BuiltinVectorDictionaryFieldIntReset(item.Dict)
}

func (item *CasesTestDictString) FillRandom(rg *basictl.RandGenerator) {
	tlBuiltinVectorDictionaryFieldInt.BuiltinVectorDictionaryFieldIntFillRandom(rg, &item.Dict)
}

func (item *CasesTestDictString) Read(w []byte) (_ []byte, err error) {
	return tlBuiltinVectorDictionaryFieldInt.BuiltinVectorDictionaryFieldIntRead(w, &item.Dict)
}

// This method is general version of Write, use it instead!
func (item *CasesTestDictString) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *CasesTestDictString) Write(w []byte) []byte {
	w = tlBuiltinVectorDictionaryFieldInt.BuiltinVectorDictionaryFieldIntWrite(w, item.Dict)
	return w
}

func (item *CasesTestDictString) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xc463c79b); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *CasesTestDictString) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *CasesTestDictString) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0xc463c79b)
	return item.Write(w)
}

func (item CasesTestDictString) String() string {
	return string(item.WriteJSON(nil))
}

func (item *CasesTestDictString) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
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
					return internal.ErrorInvalidJSONWithDuplicatingKeys("cases.testDictString", "dict")
				}
				if err := tlBuiltinVectorDictionaryFieldInt.BuiltinVectorDictionaryFieldIntReadJSON(legacyTypeNames, in, &item.Dict); err != nil {
					return err
				}
				propDictPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("cases.testDictString", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propDictPresented {
		tlBuiltinVectorDictionaryFieldInt.BuiltinVectorDictionaryFieldIntReset(item.Dict)
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *CasesTestDictString) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *CasesTestDictString) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *CasesTestDictString) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexDict := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"dict":`...)
	w = tlBuiltinVectorDictionaryFieldInt.BuiltinVectorDictionaryFieldIntWriteJSONOpt(newTypeNames, short, w, item.Dict)
	if (len(item.Dict) != 0) == false {
		w = w[:backupIndexDict]
	}
	return append(w, '}')
}

func (item *CasesTestDictString) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *CasesTestDictString) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("cases.testDictString", err.Error())
	}
	return nil
}
