// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlService3DeleteAllProducts

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tl/tlBool"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type Service3DeleteAllProducts struct {
	UserId    int32
	Type      int32
	StartDate int32
	EndDate   int32
}

func (Service3DeleteAllProducts) TLName() string { return "service3.deleteAllProducts" }
func (Service3DeleteAllProducts) TLTag() uint32  { return 0x4494acc2 }

func (item *Service3DeleteAllProducts) Reset() {
	item.UserId = 0
	item.Type = 0
	item.StartDate = 0
	item.EndDate = 0
}

func (item *Service3DeleteAllProducts) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.IntRead(w, &item.UserId); err != nil {
		return w, err
	}
	if w, err = basictl.IntRead(w, &item.Type); err != nil {
		return w, err
	}
	if w, err = basictl.IntRead(w, &item.StartDate); err != nil {
		return w, err
	}
	return basictl.IntRead(w, &item.EndDate)
}

// This method is general version of Write, use it instead!
func (item *Service3DeleteAllProducts) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *Service3DeleteAllProducts) Write(w []byte) []byte {
	w = basictl.IntWrite(w, item.UserId)
	w = basictl.IntWrite(w, item.Type)
	w = basictl.IntWrite(w, item.StartDate)
	w = basictl.IntWrite(w, item.EndDate)
	return w
}

func (item *Service3DeleteAllProducts) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x4494acc2); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *Service3DeleteAllProducts) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *Service3DeleteAllProducts) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x4494acc2)
	return item.Write(w)
}

func (item *Service3DeleteAllProducts) ReadResult(w []byte, ret *bool) (_ []byte, err error) {
	return tlBool.BoolReadBoxed(w, ret)
}

func (item *Service3DeleteAllProducts) WriteResult(w []byte, ret bool) (_ []byte, err error) {
	w = tlBool.BoolWriteBoxed(w, ret)
	return w, nil
}

func (item *Service3DeleteAllProducts) ReadResultJSON(legacyTypeNames bool, in *basictl.JsonLexer, ret *bool) error {
	if err := internal.Json2ReadBool(in, ret); err != nil {
		return err
	}
	return nil
}

func (item *Service3DeleteAllProducts) WriteResultJSON(w []byte, ret bool) (_ []byte, err error) {
	return item.writeResultJSON(true, false, w, ret)
}

func (item *Service3DeleteAllProducts) writeResultJSON(newTypeNames bool, short bool, w []byte, ret bool) (_ []byte, err error) {
	w = basictl.JSONWriteBool(w, ret)
	return w, nil
}

func (item *Service3DeleteAllProducts) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret bool
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *Service3DeleteAllProducts) ReadResultWriteResultJSONOpt(newTypeNames bool, short bool, r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret bool
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(newTypeNames, short, w, ret)
	return r, w, err
}

func (item *Service3DeleteAllProducts) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	var ret bool
	err := item.ReadResultJSON(true, &basictl.JsonLexer{Data: r}, &ret)
	if err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item Service3DeleteAllProducts) String() string {
	return string(item.WriteJSON(nil))
}

func (item *Service3DeleteAllProducts) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propUserIdPresented bool
	var propTypePresented bool
	var propStartDatePresented bool
	var propEndDatePresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "user_id":
				if propUserIdPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service3.deleteAllProducts", "user_id")
				}
				if err := internal.Json2ReadInt32(in, &item.UserId); err != nil {
					return err
				}
				propUserIdPresented = true
			case "type":
				if propTypePresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service3.deleteAllProducts", "type")
				}
				if err := internal.Json2ReadInt32(in, &item.Type); err != nil {
					return err
				}
				propTypePresented = true
			case "start_date":
				if propStartDatePresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service3.deleteAllProducts", "start_date")
				}
				if err := internal.Json2ReadInt32(in, &item.StartDate); err != nil {
					return err
				}
				propStartDatePresented = true
			case "end_date":
				if propEndDatePresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service3.deleteAllProducts", "end_date")
				}
				if err := internal.Json2ReadInt32(in, &item.EndDate); err != nil {
					return err
				}
				propEndDatePresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("service3.deleteAllProducts", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propUserIdPresented {
		item.UserId = 0
	}
	if !propTypePresented {
		item.Type = 0
	}
	if !propStartDatePresented {
		item.StartDate = 0
	}
	if !propEndDatePresented {
		item.EndDate = 0
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *Service3DeleteAllProducts) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *Service3DeleteAllProducts) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *Service3DeleteAllProducts) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexUserId := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"user_id":`...)
	w = basictl.JSONWriteInt32(w, item.UserId)
	if (item.UserId != 0) == false {
		w = w[:backupIndexUserId]
	}
	backupIndexType := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"type":`...)
	w = basictl.JSONWriteInt32(w, item.Type)
	if (item.Type != 0) == false {
		w = w[:backupIndexType]
	}
	backupIndexStartDate := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"start_date":`...)
	w = basictl.JSONWriteInt32(w, item.StartDate)
	if (item.StartDate != 0) == false {
		w = w[:backupIndexStartDate]
	}
	backupIndexEndDate := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"end_date":`...)
	w = basictl.JSONWriteInt32(w, item.EndDate)
	if (item.EndDate != 0) == false {
		w = w[:backupIndexEndDate]
	}
	return append(w, '}')
}

func (item *Service3DeleteAllProducts) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *Service3DeleteAllProducts) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("service3.deleteAllProducts", err.Error())
	}
	return nil
}
