// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlService5Insert

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/cycle_16847572a0831d4cd4c0c0fb513151f3"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type Service5Insert struct {
	Table string
	Data  string
}

func (Service5Insert) TLName() string { return "service5.insert" }
func (Service5Insert) TLTag() uint32  { return 0xc911ee2c }

func (item *Service5Insert) Reset() {
	item.Table = ""
	item.Data = ""
}

func (item *Service5Insert) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.StringRead(w, &item.Table); err != nil {
		return w, err
	}
	return basictl.StringRead(w, &item.Data)
}

// This method is general version of Write, use it instead!
func (item *Service5Insert) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *Service5Insert) Write(w []byte) []byte {
	w = basictl.StringWrite(w, item.Table)
	w = basictl.StringWrite(w, item.Data)
	return w
}

func (item *Service5Insert) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xc911ee2c); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *Service5Insert) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *Service5Insert) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0xc911ee2c)
	return item.Write(w)
}

func (item *Service5Insert) ReadResult(w []byte, ret *cycle_16847572a0831d4cd4c0c0fb513151f3.Service5Output) (_ []byte, err error) {
	return ret.ReadBoxed(w)
}

func (item *Service5Insert) WriteResult(w []byte, ret cycle_16847572a0831d4cd4c0c0fb513151f3.Service5Output) (_ []byte, err error) {
	w = ret.WriteBoxed(w)
	return w, nil
}

func (item *Service5Insert) ReadResultJSON(legacyTypeNames bool, in *basictl.JsonLexer, ret *cycle_16847572a0831d4cd4c0c0fb513151f3.Service5Output) error {
	if err := ret.ReadJSON(legacyTypeNames, in); err != nil {
		return err
	}
	return nil
}

func (item *Service5Insert) WriteResultJSON(w []byte, ret cycle_16847572a0831d4cd4c0c0fb513151f3.Service5Output) (_ []byte, err error) {
	return item.writeResultJSON(true, false, w, ret)
}

func (item *Service5Insert) writeResultJSON(newTypeNames bool, short bool, w []byte, ret cycle_16847572a0831d4cd4c0c0fb513151f3.Service5Output) (_ []byte, err error) {
	w = ret.WriteJSONOpt(newTypeNames, short, w)
	return w, nil
}

func (item *Service5Insert) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret cycle_16847572a0831d4cd4c0c0fb513151f3.Service5Output
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *Service5Insert) ReadResultWriteResultJSONOpt(newTypeNames bool, short bool, r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret cycle_16847572a0831d4cd4c0c0fb513151f3.Service5Output
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(newTypeNames, short, w, ret)
	return r, w, err
}

func (item *Service5Insert) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	var ret cycle_16847572a0831d4cd4c0c0fb513151f3.Service5Output
	err := item.ReadResultJSON(true, &basictl.JsonLexer{Data: r}, &ret)
	if err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item Service5Insert) String() string {
	return string(item.WriteJSON(nil))
}

func (item *Service5Insert) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propTablePresented bool
	var propDataPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "table":
				if propTablePresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service5.insert", "table")
				}
				if err := internal.Json2ReadString(in, &item.Table); err != nil {
					return err
				}
				propTablePresented = true
			case "data":
				if propDataPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service5.insert", "data")
				}
				if err := internal.Json2ReadString(in, &item.Data); err != nil {
					return err
				}
				propDataPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("service5.insert", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propTablePresented {
		item.Table = ""
	}
	if !propDataPresented {
		item.Data = ""
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *Service5Insert) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *Service5Insert) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *Service5Insert) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexTable := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"table":`...)
	w = basictl.JSONWriteString(w, item.Table)
	if (len(item.Table) != 0) == false {
		w = w[:backupIndexTable]
	}
	backupIndexData := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"data":`...)
	w = basictl.JSONWriteString(w, item.Data)
	if (len(item.Data) != 0) == false {
		w = w[:backupIndexData]
	}
	return append(w, '}')
}

func (item *Service5Insert) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *Service5Insert) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("service5.insert", err.Error())
	}
	return nil
}
