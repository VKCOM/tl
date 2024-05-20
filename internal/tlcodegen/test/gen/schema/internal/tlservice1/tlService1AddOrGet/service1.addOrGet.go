// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlService1AddOrGet

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/cycle_6ca945392bbf8b14f24e5653edc8b214"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type Service1AddOrGet struct {
	Key   string
	Flags int32
	Delay int32
	Value string
}

func (Service1AddOrGet) TLName() string { return "service1.addOrGet" }
func (Service1AddOrGet) TLTag() uint32  { return 0x6a42faad }

func (item *Service1AddOrGet) Reset() {
	item.Key = ""
	item.Flags = 0
	item.Delay = 0
	item.Value = ""
}

func (item *Service1AddOrGet) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.StringRead(w, &item.Key); err != nil {
		return w, err
	}
	if w, err = basictl.IntRead(w, &item.Flags); err != nil {
		return w, err
	}
	if w, err = basictl.IntRead(w, &item.Delay); err != nil {
		return w, err
	}
	return basictl.StringRead(w, &item.Value)
}

// This method is general version of Write, use it instead!
func (item *Service1AddOrGet) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *Service1AddOrGet) Write(w []byte) []byte {
	w = basictl.StringWrite(w, item.Key)
	w = basictl.IntWrite(w, item.Flags)
	w = basictl.IntWrite(w, item.Delay)
	w = basictl.StringWrite(w, item.Value)
	return w
}

func (item *Service1AddOrGet) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x6a42faad); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *Service1AddOrGet) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *Service1AddOrGet) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x6a42faad)
	return item.Write(w)
}

func (item *Service1AddOrGet) ReadResult(w []byte, ret *cycle_6ca945392bbf8b14f24e5653edc8b214.Service1Value) (_ []byte, err error) {
	return ret.ReadBoxed(w)
}

func (item *Service1AddOrGet) WriteResult(w []byte, ret cycle_6ca945392bbf8b14f24e5653edc8b214.Service1Value) (_ []byte, err error) {
	w = ret.WriteBoxed(w)
	return w, nil
}

func (item *Service1AddOrGet) ReadResultJSON(legacyTypeNames bool, in *basictl.JsonLexer, ret *cycle_6ca945392bbf8b14f24e5653edc8b214.Service1Value) error {
	if err := ret.ReadJSON(legacyTypeNames, in); err != nil {
		return err
	}
	return nil
}

func (item *Service1AddOrGet) WriteResultJSON(w []byte, ret cycle_6ca945392bbf8b14f24e5653edc8b214.Service1Value) (_ []byte, err error) {
	return item.writeResultJSON(true, false, w, ret)
}

func (item *Service1AddOrGet) writeResultJSON(newTypeNames bool, short bool, w []byte, ret cycle_6ca945392bbf8b14f24e5653edc8b214.Service1Value) (_ []byte, err error) {
	w = ret.WriteJSONOpt(newTypeNames, short, w)
	return w, nil
}

func (item *Service1AddOrGet) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret cycle_6ca945392bbf8b14f24e5653edc8b214.Service1Value
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *Service1AddOrGet) ReadResultWriteResultJSONOpt(newTypeNames bool, short bool, r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret cycle_6ca945392bbf8b14f24e5653edc8b214.Service1Value
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(newTypeNames, short, w, ret)
	return r, w, err
}

func (item *Service1AddOrGet) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	var ret cycle_6ca945392bbf8b14f24e5653edc8b214.Service1Value
	err := item.ReadResultJSON(true, &basictl.JsonLexer{Data: r}, &ret)
	if err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item Service1AddOrGet) String() string {
	return string(item.WriteJSON(nil))
}

func (item *Service1AddOrGet) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propKeyPresented bool
	var propFlagsPresented bool
	var propDelayPresented bool
	var propValuePresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "key":
				if propKeyPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service1.addOrGet", "key")
				}
				if err := internal.Json2ReadString(in, &item.Key); err != nil {
					return err
				}
				propKeyPresented = true
			case "flags":
				if propFlagsPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service1.addOrGet", "flags")
				}
				if err := internal.Json2ReadInt32(in, &item.Flags); err != nil {
					return err
				}
				propFlagsPresented = true
			case "delay":
				if propDelayPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service1.addOrGet", "delay")
				}
				if err := internal.Json2ReadInt32(in, &item.Delay); err != nil {
					return err
				}
				propDelayPresented = true
			case "value":
				if propValuePresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service1.addOrGet", "value")
				}
				if err := internal.Json2ReadString(in, &item.Value); err != nil {
					return err
				}
				propValuePresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("service1.addOrGet", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propKeyPresented {
		item.Key = ""
	}
	if !propFlagsPresented {
		item.Flags = 0
	}
	if !propDelayPresented {
		item.Delay = 0
	}
	if !propValuePresented {
		item.Value = ""
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *Service1AddOrGet) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *Service1AddOrGet) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *Service1AddOrGet) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexKey := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"key":`...)
	w = basictl.JSONWriteString(w, item.Key)
	if (len(item.Key) != 0) == false {
		w = w[:backupIndexKey]
	}
	backupIndexFlags := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"flags":`...)
	w = basictl.JSONWriteInt32(w, item.Flags)
	if (item.Flags != 0) == false {
		w = w[:backupIndexFlags]
	}
	backupIndexDelay := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"delay":`...)
	w = basictl.JSONWriteInt32(w, item.Delay)
	if (item.Delay != 0) == false {
		w = w[:backupIndexDelay]
	}
	backupIndexValue := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"value":`...)
	w = basictl.JSONWriteString(w, item.Value)
	if (len(item.Value) != 0) == false {
		w = w[:backupIndexValue]
	}
	return append(w, '}')
}

func (item *Service1AddOrGet) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *Service1AddOrGet) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("service1.addOrGet", err.Error())
	}
	return nil
}
