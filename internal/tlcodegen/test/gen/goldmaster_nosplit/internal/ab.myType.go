// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package internal

import (
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite

type AbMyType struct {
	X int32
}

func (AbMyType) TLName() string { return "ab.myType" }
func (AbMyType) TLTag() uint32  { return 0xe0e96c86 }

func (item *AbMyType) Reset() {
	item.X = 0
}

func (item *AbMyType) FillRandom(rg *basictl.RandGenerator) {
	item.X = basictl.RandomInt(rg)
}

func (item *AbMyType) Read(w []byte) (_ []byte, err error) {
	return basictl.IntRead(w, &item.X)
}

// This method is general version of Write, use it instead!
func (item *AbMyType) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *AbMyType) Write(w []byte) []byte {
	w = basictl.IntWrite(w, item.X)
	return w
}

func (item *AbMyType) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xe0e96c86); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *AbMyType) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *AbMyType) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0xe0e96c86)
	return item.Write(w)
}

func (item AbMyType) String() string {
	return string(item.WriteJSON(nil))
}

func (item *AbMyType) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
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
					return ErrorInvalidJSONWithDuplicatingKeys("ab.myType", "x")
				}
				if err := Json2ReadInt32(in, &item.X); err != nil {
					return err
				}
				propXPresented = true
			default:
				return ErrorInvalidJSONExcessElement("ab.myType", key)
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
func (item *AbMyType) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *AbMyType) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *AbMyType) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
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

func (item *AbMyType) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *AbMyType) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("ab.myType", err.Error())
	}
	return nil
}

type AbMyTypeBoxedMaybe struct {
	Value AbMyType // not deterministic if !Ok
	Ok    bool
}

func (item *AbMyTypeBoxedMaybe) Reset() {
	item.Ok = false
}
func (item *AbMyTypeBoxedMaybe) FillRandom(rg *basictl.RandGenerator) {
	if basictl.RandomUint(rg)&1 == 1 {
		item.Ok = true
		item.Value.FillRandom(rg)
	} else {
		item.Ok = false
	}
}

func (item *AbMyTypeBoxedMaybe) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.ReadBool(w, &item.Ok, 0x27930a7b, 0x3f9c8ef8); err != nil {
		return w, err
	}
	if item.Ok {
		return item.Value.ReadBoxed(w)
	}
	return w, nil
}

// This method is general version of WriteBoxed, use it instead!
func (item *AbMyTypeBoxedMaybe) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *AbMyTypeBoxedMaybe) WriteBoxed(w []byte) []byte {
	if item.Ok {
		w = basictl.NatWrite(w, 0x3f9c8ef8)
		return item.Value.WriteBoxed(w)
	}
	return basictl.NatWrite(w, 0x27930a7b)
}

func (item *AbMyTypeBoxedMaybe) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	_ok, _jvalue, err := Json2ReadMaybe("Maybe", in)
	if err != nil {
		return err
	}
	item.Ok = _ok
	if _ok {
		var in2Pointer *basictl.JsonLexer
		if _jvalue != nil {
			in2 := basictl.JsonLexer{Data: _jvalue}
			in2Pointer = &in2
		}
		if err := item.Value.ReadJSON(legacyTypeNames, in2Pointer); err != nil {
			return err
		}
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *AbMyTypeBoxedMaybe) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *AbMyTypeBoxedMaybe) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *AbMyTypeBoxedMaybe) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	if !item.Ok {
		return append(w, "{}"...)
	}
	w = append(w, `{"ok":true`...)
	w = append(w, `,"value":`...)
	w = item.Value.WriteJSONOpt(newTypeNames, short, w)
	return append(w, '}')
}

func (item AbMyTypeBoxedMaybe) String() string {
	return string(item.WriteJSON(nil))
}

type AbMyTypeMaybe struct {
	Value AbMyType // not deterministic if !Ok
	Ok    bool
}

func (item *AbMyTypeMaybe) Reset() {
	item.Ok = false
}
func (item *AbMyTypeMaybe) FillRandom(rg *basictl.RandGenerator) {
	if basictl.RandomUint(rg)&1 == 1 {
		item.Ok = true
		item.Value.FillRandom(rg)
	} else {
		item.Ok = false
	}
}

func (item *AbMyTypeMaybe) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.ReadBool(w, &item.Ok, 0x27930a7b, 0x3f9c8ef8); err != nil {
		return w, err
	}
	if item.Ok {
		return item.Value.Read(w)
	}
	return w, nil
}

// This method is general version of WriteBoxed, use it instead!
func (item *AbMyTypeMaybe) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *AbMyTypeMaybe) WriteBoxed(w []byte) []byte {
	if item.Ok {
		w = basictl.NatWrite(w, 0x3f9c8ef8)
		return item.Value.Write(w)
	}
	return basictl.NatWrite(w, 0x27930a7b)
}

func (item *AbMyTypeMaybe) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	_ok, _jvalue, err := Json2ReadMaybe("Maybe", in)
	if err != nil {
		return err
	}
	item.Ok = _ok
	if _ok {
		var in2Pointer *basictl.JsonLexer
		if _jvalue != nil {
			in2 := basictl.JsonLexer{Data: _jvalue}
			in2Pointer = &in2
		}
		if err := item.Value.ReadJSON(legacyTypeNames, in2Pointer); err != nil {
			return err
		}
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *AbMyTypeMaybe) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *AbMyTypeMaybe) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *AbMyTypeMaybe) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	if !item.Ok {
		return append(w, "{}"...)
	}
	w = append(w, `{"ok":true`...)
	w = append(w, `,"value":`...)
	w = item.Value.WriteJSONOpt(newTypeNames, short, w)
	return append(w, '}')
}

func (item AbMyTypeMaybe) String() string {
	return string(item.WriteJSON(nil))
}
