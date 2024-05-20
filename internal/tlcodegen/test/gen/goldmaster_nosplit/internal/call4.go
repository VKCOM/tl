// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package internal

import (
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite

type Call4 struct {
	X CdTypeA
}

func (Call4) TLName() string { return "call4" }
func (Call4) TLTag() uint32  { return 0x46d7de8f }

func (item *Call4) Reset() {
	item.X.Reset()
}

func (item *Call4) FillRandom(rg *basictl.RandGenerator) {
	item.X.FillRandom(rg)
}

func (item *Call4) Read(w []byte) (_ []byte, err error) {
	return item.X.Read(w)
}

// This method is general version of Write, use it instead!
func (item *Call4) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *Call4) Write(w []byte) []byte {
	w = item.X.Write(w)
	return w
}

func (item *Call4) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x46d7de8f); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *Call4) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *Call4) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x46d7de8f)
	return item.Write(w)
}

func (item *Call4) ReadResult(w []byte, ret *AbTypeB) (_ []byte, err error) {
	return ret.ReadBoxed(w)
}

func (item *Call4) WriteResult(w []byte, ret AbTypeB) (_ []byte, err error) {
	w = ret.WriteBoxed(w)
	return w, nil
}

func (item *Call4) ReadResultJSON(legacyTypeNames bool, in *basictl.JsonLexer, ret *AbTypeB) error {
	if err := ret.ReadJSON(legacyTypeNames, in); err != nil {
		return err
	}
	return nil
}

func (item *Call4) WriteResultJSON(w []byte, ret AbTypeB) (_ []byte, err error) {
	return item.writeResultJSON(true, false, w, ret)
}

func (item *Call4) writeResultJSON(newTypeNames bool, short bool, w []byte, ret AbTypeB) (_ []byte, err error) {
	w = ret.WriteJSONOpt(newTypeNames, short, w)
	return w, nil
}

func (item *Call4) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret AbTypeB
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *Call4) ReadResultWriteResultJSONOpt(newTypeNames bool, short bool, r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret AbTypeB
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(newTypeNames, short, w, ret)
	return r, w, err
}

func (item *Call4) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	var ret AbTypeB
	err := item.ReadResultJSON(true, &basictl.JsonLexer{Data: r}, &ret)
	if err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item Call4) String() string {
	return string(item.WriteJSON(nil))
}

func (item *Call4) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
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
					return ErrorInvalidJSONWithDuplicatingKeys("call4", "x")
				}
				if err := item.X.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propXPresented = true
			default:
				return ErrorInvalidJSONExcessElement("call4", key)
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
func (item *Call4) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *Call4) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *Call4) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"x":`...)
	w = item.X.WriteJSONOpt(newTypeNames, short, w)
	return append(w, '}')
}

func (item *Call4) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *Call4) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("call4", err.Error())
	}
	return nil
}
