// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlService2SetObjectTtl

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tl/tlTrue"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tlservice2/tlService2ObjectId"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type Service2SetObjectTtl struct {
	ObjectIdLength uint32
	ObjectId       tlService2ObjectId.Service2ObjectId
	Ttl            int32
}

func (Service2SetObjectTtl) TLName() string { return "service2.setObjectTtl" }
func (Service2SetObjectTtl) TLTag() uint32  { return 0x6f98f025 }

func (item *Service2SetObjectTtl) Reset() {
	item.ObjectIdLength = 0
	item.ObjectId.Reset()
	item.Ttl = 0
}

func (item *Service2SetObjectTtl) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.ObjectIdLength); err != nil {
		return w, err
	}
	if w, err = item.ObjectId.Read(w, item.ObjectIdLength); err != nil {
		return w, err
	}
	return basictl.IntRead(w, &item.Ttl)
}

// This method is general version of Write, use it instead!
func (item *Service2SetObjectTtl) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w)
}

func (item *Service2SetObjectTtl) Write(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, item.ObjectIdLength)
	if w, err = item.ObjectId.Write(w, item.ObjectIdLength); err != nil {
		return w, err
	}
	w = basictl.IntWrite(w, item.Ttl)
	return w, nil
}

func (item *Service2SetObjectTtl) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x6f98f025); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *Service2SetObjectTtl) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w)
}

func (item *Service2SetObjectTtl) WriteBoxed(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, 0x6f98f025)
	return item.Write(w)
}

func (item *Service2SetObjectTtl) ReadResult(w []byte, ret *tlTrue.True) (_ []byte, err error) {
	return ret.ReadBoxed(w)
}

func (item *Service2SetObjectTtl) WriteResult(w []byte, ret tlTrue.True) (_ []byte, err error) {
	w = ret.WriteBoxed(w)
	return w, nil
}

func (item *Service2SetObjectTtl) ReadResultJSON(legacyTypeNames bool, in *basictl.JsonLexer, ret *tlTrue.True) error {
	if err := ret.ReadJSON(legacyTypeNames, in); err != nil {
		return err
	}
	return nil
}

func (item *Service2SetObjectTtl) WriteResultJSON(w []byte, ret tlTrue.True) (_ []byte, err error) {
	return item.writeResultJSON(true, false, w, ret)
}

func (item *Service2SetObjectTtl) writeResultJSON(newTypeNames bool, short bool, w []byte, ret tlTrue.True) (_ []byte, err error) {
	w = ret.WriteJSONOpt(newTypeNames, short, w)
	return w, nil
}

func (item *Service2SetObjectTtl) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret tlTrue.True
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *Service2SetObjectTtl) ReadResultWriteResultJSONOpt(newTypeNames bool, short bool, r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret tlTrue.True
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(newTypeNames, short, w, ret)
	return r, w, err
}

func (item *Service2SetObjectTtl) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	var ret tlTrue.True
	err := item.ReadResultJSON(true, &basictl.JsonLexer{Data: r}, &ret)
	if err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item Service2SetObjectTtl) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func (item *Service2SetObjectTtl) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propObjectIdLengthPresented bool
	var rawObjectId []byte
	var propTtlPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "objectIdLength":
				if propObjectIdLengthPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service2.setObjectTtl", "objectIdLength")
				}
				if err := internal.Json2ReadUint32(in, &item.ObjectIdLength); err != nil {
					return err
				}
				propObjectIdLengthPresented = true
			case "objectId":
				if rawObjectId != nil {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service2.setObjectTtl", "objectId")
				}
				rawObjectId = in.Raw()
				if !in.Ok() {
					return in.Error()
				}
			case "ttl":
				if propTtlPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service2.setObjectTtl", "ttl")
				}
				if err := internal.Json2ReadInt32(in, &item.Ttl); err != nil {
					return err
				}
				propTtlPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("service2.setObjectTtl", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propObjectIdLengthPresented {
		item.ObjectIdLength = 0
	}
	if !propTtlPresented {
		item.Ttl = 0
	}
	var inObjectIdPointer *basictl.JsonLexer
	inObjectId := basictl.JsonLexer{Data: rawObjectId}
	if rawObjectId != nil {
		inObjectIdPointer = &inObjectId
	}
	if err := item.ObjectId.ReadJSON(legacyTypeNames, inObjectIdPointer, item.ObjectIdLength); err != nil {
		return err
	}

	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *Service2SetObjectTtl) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w)
}

func (item *Service2SetObjectTtl) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w)
}
func (item *Service2SetObjectTtl) WriteJSONOpt(newTypeNames bool, short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	backupIndexObjectIdLength := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"objectIdLength":`...)
	w = basictl.JSONWriteUint32(w, item.ObjectIdLength)
	if (item.ObjectIdLength != 0) == false {
		w = w[:backupIndexObjectIdLength]
	}
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"objectId":`...)
	if w, err = item.ObjectId.WriteJSONOpt(newTypeNames, short, w, item.ObjectIdLength); err != nil {
		return w, err
	}
	backupIndexTtl := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"ttl":`...)
	w = basictl.JSONWriteInt32(w, item.Ttl)
	if (item.Ttl != 0) == false {
		w = w[:backupIndexTtl]
	}
	return append(w, '}'), nil
}

func (item *Service2SetObjectTtl) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *Service2SetObjectTtl) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("service2.setObjectTtl", err.Error())
	}
	return nil
}
