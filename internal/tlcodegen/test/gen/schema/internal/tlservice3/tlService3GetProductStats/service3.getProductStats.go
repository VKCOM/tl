// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlService3GetProductStats

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tl/tlBuiltinVectorInt"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tl/tlVectorService3ProductStatsOldMaybe"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type Service3GetProductStats struct {
	UserId int32
	Types  []int32
}

func (Service3GetProductStats) TLName() string { return "service3.getProductStats" }
func (Service3GetProductStats) TLTag() uint32  { return 0x261f6898 }

func (item *Service3GetProductStats) Reset() {
	item.UserId = 0
	item.Types = item.Types[:0]
}

func (item *Service3GetProductStats) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.IntRead(w, &item.UserId); err != nil {
		return w, err
	}
	return tlBuiltinVectorInt.BuiltinVectorIntRead(w, &item.Types)
}

// This method is general version of Write, use it instead!
func (item *Service3GetProductStats) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *Service3GetProductStats) Write(w []byte) []byte {
	w = basictl.IntWrite(w, item.UserId)
	w = tlBuiltinVectorInt.BuiltinVectorIntWrite(w, item.Types)
	return w
}

func (item *Service3GetProductStats) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x261f6898); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *Service3GetProductStats) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *Service3GetProductStats) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x261f6898)
	return item.Write(w)
}

func (item *Service3GetProductStats) ReadResult(w []byte, ret *tlVectorService3ProductStatsOldMaybe.VectorService3ProductStatsOldMaybe) (_ []byte, err error) {
	return ret.ReadBoxed(w)
}

func (item *Service3GetProductStats) WriteResult(w []byte, ret tlVectorService3ProductStatsOldMaybe.VectorService3ProductStatsOldMaybe) (_ []byte, err error) {
	w = ret.WriteBoxed(w)
	return w, nil
}

func (item *Service3GetProductStats) ReadResultJSON(legacyTypeNames bool, in *basictl.JsonLexer, ret *tlVectorService3ProductStatsOldMaybe.VectorService3ProductStatsOldMaybe) error {
	if err := ret.ReadJSON(legacyTypeNames, in); err != nil {
		return err
	}
	return nil
}

func (item *Service3GetProductStats) WriteResultJSON(w []byte, ret tlVectorService3ProductStatsOldMaybe.VectorService3ProductStatsOldMaybe) (_ []byte, err error) {
	return item.writeResultJSON(true, false, w, ret)
}

func (item *Service3GetProductStats) writeResultJSON(newTypeNames bool, short bool, w []byte, ret tlVectorService3ProductStatsOldMaybe.VectorService3ProductStatsOldMaybe) (_ []byte, err error) {
	w = ret.WriteJSONOpt(newTypeNames, short, w)
	return w, nil
}

func (item *Service3GetProductStats) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret tlVectorService3ProductStatsOldMaybe.VectorService3ProductStatsOldMaybe
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *Service3GetProductStats) ReadResultWriteResultJSONOpt(newTypeNames bool, short bool, r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret tlVectorService3ProductStatsOldMaybe.VectorService3ProductStatsOldMaybe
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(newTypeNames, short, w, ret)
	return r, w, err
}

func (item *Service3GetProductStats) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	var ret tlVectorService3ProductStatsOldMaybe.VectorService3ProductStatsOldMaybe
	err := item.ReadResultJSON(true, &basictl.JsonLexer{Data: r}, &ret)
	if err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item Service3GetProductStats) String() string {
	return string(item.WriteJSON(nil))
}

func (item *Service3GetProductStats) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propUserIdPresented bool
	var propTypesPresented bool

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
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service3.getProductStats", "user_id")
				}
				if err := internal.Json2ReadInt32(in, &item.UserId); err != nil {
					return err
				}
				propUserIdPresented = true
			case "types":
				if propTypesPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service3.getProductStats", "types")
				}
				if err := tlBuiltinVectorInt.BuiltinVectorIntReadJSON(legacyTypeNames, in, &item.Types); err != nil {
					return err
				}
				propTypesPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("service3.getProductStats", key)
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
	if !propTypesPresented {
		item.Types = item.Types[:0]
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *Service3GetProductStats) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *Service3GetProductStats) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *Service3GetProductStats) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexUserId := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"user_id":`...)
	w = basictl.JSONWriteInt32(w, item.UserId)
	if (item.UserId != 0) == false {
		w = w[:backupIndexUserId]
	}
	backupIndexTypes := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"types":`...)
	w = tlBuiltinVectorInt.BuiltinVectorIntWriteJSONOpt(newTypeNames, short, w, item.Types)
	if (len(item.Types) != 0) == false {
		w = w[:backupIndexTypes]
	}
	return append(w, '}')
}

func (item *Service3GetProductStats) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *Service3GetProductStats) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("service3.getProductStats", err.Error())
	}
	return nil
}
