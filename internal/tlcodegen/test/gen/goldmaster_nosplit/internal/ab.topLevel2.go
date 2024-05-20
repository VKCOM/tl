// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package internal

import (
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite

type AbTopLevel2 struct {
	A NoStr
	B UseStr
}

func (AbTopLevel2) TLName() string { return "ab.topLevel2" }
func (AbTopLevel2) TLTag() uint32  { return 0xcef933fb }

func (item *AbTopLevel2) Reset() {
	item.A.Reset()
	item.B.Reset()
}

func (item *AbTopLevel2) FillRandom(rg *basictl.RandGenerator) {
	item.A.FillRandom(rg)
	item.B.FillRandom(rg)
}

func (item *AbTopLevel2) Read(w []byte) (_ []byte, err error) {
	if w, err = item.A.Read(w); err != nil {
		return w, err
	}
	return item.B.Read(w)
}

// This method is general version of Write, use it instead!
func (item *AbTopLevel2) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *AbTopLevel2) Write(w []byte) []byte {
	w = item.A.Write(w)
	w = item.B.Write(w)
	return w
}

func (item *AbTopLevel2) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xcef933fb); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *AbTopLevel2) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *AbTopLevel2) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0xcef933fb)
	return item.Write(w)
}

func (item AbTopLevel2) String() string {
	return string(item.WriteJSON(nil))
}

func (item *AbTopLevel2) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propAPresented bool
	var propBPresented bool

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
					return ErrorInvalidJSONWithDuplicatingKeys("ab.topLevel2", "a")
				}
				if err := item.A.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propAPresented = true
			case "b":
				if propBPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("ab.topLevel2", "b")
				}
				if err := item.B.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propBPresented = true
			default:
				return ErrorInvalidJSONExcessElement("ab.topLevel2", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propAPresented {
		item.A.Reset()
	}
	if !propBPresented {
		item.B.Reset()
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *AbTopLevel2) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *AbTopLevel2) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *AbTopLevel2) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"a":`...)
	w = item.A.WriteJSONOpt(newTypeNames, short, w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"b":`...)
	w = item.B.WriteJSONOpt(newTypeNames, short, w)
	return append(w, '}')
}

func (item *AbTopLevel2) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *AbTopLevel2) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("ab.topLevel2", err.Error())
	}
	return nil
}

type AbTopLevel2Bytes struct {
	A NoStr
	B UseStrBytes
}

func (AbTopLevel2Bytes) TLName() string { return "ab.topLevel2" }
func (AbTopLevel2Bytes) TLTag() uint32  { return 0xcef933fb }

func (item *AbTopLevel2Bytes) Reset() {
	item.A.Reset()
	item.B.Reset()
}

func (item *AbTopLevel2Bytes) FillRandom(rg *basictl.RandGenerator) {
	item.A.FillRandom(rg)
	item.B.FillRandom(rg)
}

func (item *AbTopLevel2Bytes) Read(w []byte) (_ []byte, err error) {
	if w, err = item.A.Read(w); err != nil {
		return w, err
	}
	return item.B.Read(w)
}

// This method is general version of Write, use it instead!
func (item *AbTopLevel2Bytes) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *AbTopLevel2Bytes) Write(w []byte) []byte {
	w = item.A.Write(w)
	w = item.B.Write(w)
	return w
}

func (item *AbTopLevel2Bytes) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xcef933fb); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *AbTopLevel2Bytes) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *AbTopLevel2Bytes) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0xcef933fb)
	return item.Write(w)
}

func (item AbTopLevel2Bytes) String() string {
	return string(item.WriteJSON(nil))
}

func (item *AbTopLevel2Bytes) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propAPresented bool
	var propBPresented bool

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
					return ErrorInvalidJSONWithDuplicatingKeys("ab.topLevel2", "a")
				}
				if err := item.A.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propAPresented = true
			case "b":
				if propBPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("ab.topLevel2", "b")
				}
				if err := item.B.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propBPresented = true
			default:
				return ErrorInvalidJSONExcessElement("ab.topLevel2", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propAPresented {
		item.A.Reset()
	}
	if !propBPresented {
		item.B.Reset()
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *AbTopLevel2Bytes) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *AbTopLevel2Bytes) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *AbTopLevel2Bytes) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"a":`...)
	w = item.A.WriteJSONOpt(newTypeNames, short, w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"b":`...)
	w = item.B.WriteJSONOpt(newTypeNames, short, w)
	return append(w, '}')
}

func (item *AbTopLevel2Bytes) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *AbTopLevel2Bytes) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("ab.topLevel2", err.Error())
	}
	return nil
}
