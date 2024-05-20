// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlService3SetLastVisitTimestamp

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tl/tlBool"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type Service3SetLastVisitTimestamp struct {
	UserId    int32
	Timestamp int32
}

func (Service3SetLastVisitTimestamp) TLName() string { return "service3.setLastVisitTimestamp" }
func (Service3SetLastVisitTimestamp) TLTag() uint32  { return 0x7909b020 }

func (item *Service3SetLastVisitTimestamp) Reset() {
	item.UserId = 0
	item.Timestamp = 0
}

func (item *Service3SetLastVisitTimestamp) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.IntRead(w, &item.UserId); err != nil {
		return w, err
	}
	return basictl.IntRead(w, &item.Timestamp)
}

// This method is general version of Write, use it instead!
func (item *Service3SetLastVisitTimestamp) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *Service3SetLastVisitTimestamp) Write(w []byte) []byte {
	w = basictl.IntWrite(w, item.UserId)
	w = basictl.IntWrite(w, item.Timestamp)
	return w
}

func (item *Service3SetLastVisitTimestamp) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x7909b020); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *Service3SetLastVisitTimestamp) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *Service3SetLastVisitTimestamp) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x7909b020)
	return item.Write(w)
}

func (item *Service3SetLastVisitTimestamp) ReadResult(w []byte, ret *bool) (_ []byte, err error) {
	return tlBool.BoolReadBoxed(w, ret)
}

func (item *Service3SetLastVisitTimestamp) WriteResult(w []byte, ret bool) (_ []byte, err error) {
	w = tlBool.BoolWriteBoxed(w, ret)
	return w, nil
}

func (item *Service3SetLastVisitTimestamp) ReadResultJSON(legacyTypeNames bool, in *basictl.JsonLexer, ret *bool) error {
	if err := internal.Json2ReadBool(in, ret); err != nil {
		return err
	}
	return nil
}

func (item *Service3SetLastVisitTimestamp) WriteResultJSON(w []byte, ret bool) (_ []byte, err error) {
	return item.writeResultJSON(true, false, w, ret)
}

func (item *Service3SetLastVisitTimestamp) writeResultJSON(newTypeNames bool, short bool, w []byte, ret bool) (_ []byte, err error) {
	w = basictl.JSONWriteBool(w, ret)
	return w, nil
}

func (item *Service3SetLastVisitTimestamp) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret bool
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *Service3SetLastVisitTimestamp) ReadResultWriteResultJSONOpt(newTypeNames bool, short bool, r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret bool
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(newTypeNames, short, w, ret)
	return r, w, err
}

func (item *Service3SetLastVisitTimestamp) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	var ret bool
	err := item.ReadResultJSON(true, &basictl.JsonLexer{Data: r}, &ret)
	if err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item Service3SetLastVisitTimestamp) String() string {
	return string(item.WriteJSON(nil))
}

func (item *Service3SetLastVisitTimestamp) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propUserIdPresented bool
	var propTimestampPresented bool

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
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service3.setLastVisitTimestamp", "user_id")
				}
				if err := internal.Json2ReadInt32(in, &item.UserId); err != nil {
					return err
				}
				propUserIdPresented = true
			case "timestamp":
				if propTimestampPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service3.setLastVisitTimestamp", "timestamp")
				}
				if err := internal.Json2ReadInt32(in, &item.Timestamp); err != nil {
					return err
				}
				propTimestampPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("service3.setLastVisitTimestamp", key)
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
	if !propTimestampPresented {
		item.Timestamp = 0
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *Service3SetLastVisitTimestamp) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *Service3SetLastVisitTimestamp) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *Service3SetLastVisitTimestamp) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexUserId := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"user_id":`...)
	w = basictl.JSONWriteInt32(w, item.UserId)
	if (item.UserId != 0) == false {
		w = w[:backupIndexUserId]
	}
	backupIndexTimestamp := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"timestamp":`...)
	w = basictl.JSONWriteInt32(w, item.Timestamp)
	if (item.Timestamp != 0) == false {
		w = w[:backupIndexTimestamp]
	}
	return append(w, '}')
}

func (item *Service3SetLastVisitTimestamp) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *Service3SetLastVisitTimestamp) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("service3.setLastVisitTimestamp", err.Error())
	}
	return nil
}
