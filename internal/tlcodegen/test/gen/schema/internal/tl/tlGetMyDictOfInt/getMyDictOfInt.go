// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlGetMyDictOfInt

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tl/tlMyDictOfInt"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type GetMyDictOfInt struct {
	X tlMyDictOfInt.MyDictOfInt
}

func (GetMyDictOfInt) TLName() string { return "getMyDictOfInt" }
func (GetMyDictOfInt) TLTag() uint32  { return 0x166f962c }

func (item *GetMyDictOfInt) Reset() {
	item.X.Reset()
}

func (item *GetMyDictOfInt) Read(w []byte) (_ []byte, err error) {
	return item.X.ReadBoxed(w)
}

// This method is general version of Write, use it instead!
func (item *GetMyDictOfInt) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *GetMyDictOfInt) Write(w []byte) []byte {
	w = item.X.WriteBoxed(w)
	return w
}

func (item *GetMyDictOfInt) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x166f962c); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *GetMyDictOfInt) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *GetMyDictOfInt) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x166f962c)
	return item.Write(w)
}

func (item *GetMyDictOfInt) ReadResult(w []byte, ret *tlMyDictOfInt.MyDictOfInt) (_ []byte, err error) {
	return ret.ReadBoxed(w)
}

func (item *GetMyDictOfInt) WriteResult(w []byte, ret tlMyDictOfInt.MyDictOfInt) (_ []byte, err error) {
	w = ret.WriteBoxed(w)
	return w, nil
}

func (item *GetMyDictOfInt) ReadResultJSON(legacyTypeNames bool, in *basictl.JsonLexer, ret *tlMyDictOfInt.MyDictOfInt) error {
	if err := ret.ReadJSON(legacyTypeNames, in); err != nil {
		return err
	}
	return nil
}

func (item *GetMyDictOfInt) WriteResultJSON(w []byte, ret tlMyDictOfInt.MyDictOfInt) (_ []byte, err error) {
	return item.writeResultJSON(true, false, w, ret)
}

func (item *GetMyDictOfInt) writeResultJSON(newTypeNames bool, short bool, w []byte, ret tlMyDictOfInt.MyDictOfInt) (_ []byte, err error) {
	w = ret.WriteJSONOpt(newTypeNames, short, w)
	return w, nil
}

func (item *GetMyDictOfInt) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret tlMyDictOfInt.MyDictOfInt
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *GetMyDictOfInt) ReadResultWriteResultJSONOpt(newTypeNames bool, short bool, r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret tlMyDictOfInt.MyDictOfInt
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(newTypeNames, short, w, ret)
	return r, w, err
}

func (item *GetMyDictOfInt) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	var ret tlMyDictOfInt.MyDictOfInt
	err := item.ReadResultJSON(true, &basictl.JsonLexer{Data: r}, &ret)
	if err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item GetMyDictOfInt) String() string {
	return string(item.WriteJSON(nil))
}

func (item *GetMyDictOfInt) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
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
					return internal.ErrorInvalidJSONWithDuplicatingKeys("getMyDictOfInt", "x")
				}
				if err := item.X.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propXPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("getMyDictOfInt", key)
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
func (item *GetMyDictOfInt) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *GetMyDictOfInt) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *GetMyDictOfInt) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexX := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"x":`...)
	w = item.X.WriteJSONOpt(newTypeNames, short, w)
	if (len(item.X) != 0) == false {
		w = w[:backupIndexX]
	}
	return append(w, '}')
}

func (item *GetMyDictOfInt) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *GetMyDictOfInt) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("getMyDictOfInt", err.Error())
	}
	return nil
}
