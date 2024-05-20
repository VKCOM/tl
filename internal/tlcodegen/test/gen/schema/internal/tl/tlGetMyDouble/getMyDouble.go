// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlGetMyDouble

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tl/tlMyDouble"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type GetMyDouble struct {
	X tlMyDouble.MyDouble
}

func (GetMyDouble) TLName() string { return "getMyDouble" }
func (GetMyDouble) TLTag() uint32  { return 0xb660ad10 }

func (item *GetMyDouble) Reset() {
	item.X.Reset()
}

func (item *GetMyDouble) Read(w []byte) (_ []byte, err error) {
	return item.X.Read(w)
}

// This method is general version of Write, use it instead!
func (item *GetMyDouble) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *GetMyDouble) Write(w []byte) []byte {
	w = item.X.Write(w)
	return w
}

func (item *GetMyDouble) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xb660ad10); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *GetMyDouble) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *GetMyDouble) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0xb660ad10)
	return item.Write(w)
}

func (item *GetMyDouble) ReadResult(w []byte, ret *tlMyDouble.MyDouble) (_ []byte, err error) {
	return ret.ReadBoxed(w)
}

func (item *GetMyDouble) WriteResult(w []byte, ret tlMyDouble.MyDouble) (_ []byte, err error) {
	w = ret.WriteBoxed(w)
	return w, nil
}

func (item *GetMyDouble) ReadResultJSON(legacyTypeNames bool, in *basictl.JsonLexer, ret *tlMyDouble.MyDouble) error {
	if err := ret.ReadJSON(legacyTypeNames, in); err != nil {
		return err
	}
	return nil
}

func (item *GetMyDouble) WriteResultJSON(w []byte, ret tlMyDouble.MyDouble) (_ []byte, err error) {
	return item.writeResultJSON(true, false, w, ret)
}

func (item *GetMyDouble) writeResultJSON(newTypeNames bool, short bool, w []byte, ret tlMyDouble.MyDouble) (_ []byte, err error) {
	w = ret.WriteJSONOpt(newTypeNames, short, w)
	return w, nil
}

func (item *GetMyDouble) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret tlMyDouble.MyDouble
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *GetMyDouble) ReadResultWriteResultJSONOpt(newTypeNames bool, short bool, r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret tlMyDouble.MyDouble
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(newTypeNames, short, w, ret)
	return r, w, err
}

func (item *GetMyDouble) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	var ret tlMyDouble.MyDouble
	err := item.ReadResultJSON(true, &basictl.JsonLexer{Data: r}, &ret)
	if err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item GetMyDouble) String() string {
	return string(item.WriteJSON(nil))
}

func (item *GetMyDouble) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
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
					return internal.ErrorInvalidJSONWithDuplicatingKeys("getMyDouble", "x")
				}
				if err := item.X.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propXPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("getMyDouble", key)
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
func (item *GetMyDouble) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *GetMyDouble) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *GetMyDouble) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexX := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"x":`...)
	w = item.X.WriteJSONOpt(newTypeNames, short, w)
	if (item.X != 0) == false {
		w = w[:backupIndexX]
	}
	return append(w, '}')
}

func (item *GetMyDouble) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *GetMyDouble) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("getMyDouble", err.Error())
	}
	return nil
}
