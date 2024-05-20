// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlService3GetLimits

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tlservice3/tlService3Limits"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type Service3GetLimits struct {
}

func (Service3GetLimits) TLName() string { return "service3.getLimits" }
func (Service3GetLimits) TLTag() uint32  { return 0xeb399467 }

func (item *Service3GetLimits) Reset() {}

func (item *Service3GetLimits) Read(w []byte) (_ []byte, err error) { return w, nil }

// This method is general version of Write, use it instead!
func (item *Service3GetLimits) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *Service3GetLimits) Write(w []byte) []byte {
	return w
}

func (item *Service3GetLimits) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xeb399467); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *Service3GetLimits) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *Service3GetLimits) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0xeb399467)
	return item.Write(w)
}

func (item *Service3GetLimits) ReadResult(w []byte, ret *tlService3Limits.Service3Limits) (_ []byte, err error) {
	return ret.ReadBoxed(w)
}

func (item *Service3GetLimits) WriteResult(w []byte, ret tlService3Limits.Service3Limits) (_ []byte, err error) {
	w = ret.WriteBoxed(w)
	return w, nil
}

func (item *Service3GetLimits) ReadResultJSON(legacyTypeNames bool, in *basictl.JsonLexer, ret *tlService3Limits.Service3Limits) error {
	if err := ret.ReadJSON(legacyTypeNames, in); err != nil {
		return err
	}
	return nil
}

func (item *Service3GetLimits) WriteResultJSON(w []byte, ret tlService3Limits.Service3Limits) (_ []byte, err error) {
	return item.writeResultJSON(true, false, w, ret)
}

func (item *Service3GetLimits) writeResultJSON(newTypeNames bool, short bool, w []byte, ret tlService3Limits.Service3Limits) (_ []byte, err error) {
	w = ret.WriteJSONOpt(newTypeNames, short, w)
	return w, nil
}

func (item *Service3GetLimits) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret tlService3Limits.Service3Limits
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *Service3GetLimits) ReadResultWriteResultJSONOpt(newTypeNames bool, short bool, r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret tlService3Limits.Service3Limits
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(newTypeNames, short, w, ret)
	return r, w, err
}

func (item *Service3GetLimits) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	var ret tlService3Limits.Service3Limits
	err := item.ReadResultJSON(true, &basictl.JsonLexer{Data: r}, &ret)
	if err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item Service3GetLimits) String() string {
	return string(item.WriteJSON(nil))
}

func (item *Service3GetLimits) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			return internal.ErrorInvalidJSON("service3.getLimits", "this object can't have properties")
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *Service3GetLimits) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *Service3GetLimits) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *Service3GetLimits) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	return append(w, '}')
}

func (item *Service3GetLimits) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *Service3GetLimits) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("service3.getLimits", err.Error())
	}
	return nil
}
