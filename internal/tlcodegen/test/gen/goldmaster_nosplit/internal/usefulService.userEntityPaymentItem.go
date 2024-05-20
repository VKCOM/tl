// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package internal

import (
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite

type UsefulServiceUserEntityPaymentItem struct {
	Id    string
	Promo UsefulServiceUserEntityPaymentItemPromoBoxedMaybe
}

func (UsefulServiceUserEntityPaymentItem) TLName() string {
	return "usefulService.userEntityPaymentItem"
}
func (UsefulServiceUserEntityPaymentItem) TLTag() uint32 { return 0x4f798680 }

func (item *UsefulServiceUserEntityPaymentItem) Reset() {
	item.Id = ""
	item.Promo.Reset()
}

func (item *UsefulServiceUserEntityPaymentItem) FillRandom(rg *basictl.RandGenerator, nat_fields_mask uint32) {
	item.Id = basictl.RandomString(rg)
	item.Promo.FillRandom(rg, nat_fields_mask)
}

func (item *UsefulServiceUserEntityPaymentItem) Read(w []byte, nat_fields_mask uint32) (_ []byte, err error) {
	if w, err = basictl.StringRead(w, &item.Id); err != nil {
		return w, err
	}
	return item.Promo.ReadBoxed(w, nat_fields_mask)
}

// This method is general version of Write, use it instead!
func (item *UsefulServiceUserEntityPaymentItem) WriteGeneral(w []byte, nat_fields_mask uint32) (_ []byte, err error) {
	return item.Write(w, nat_fields_mask), nil
}

func (item *UsefulServiceUserEntityPaymentItem) Write(w []byte, nat_fields_mask uint32) []byte {
	w = basictl.StringWrite(w, item.Id)
	w = item.Promo.WriteBoxed(w, nat_fields_mask)
	return w
}

func (item *UsefulServiceUserEntityPaymentItem) ReadBoxed(w []byte, nat_fields_mask uint32) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x4f798680); err != nil {
		return w, err
	}
	return item.Read(w, nat_fields_mask)
}

// This method is general version of WriteBoxed, use it instead!
func (item *UsefulServiceUserEntityPaymentItem) WriteBoxedGeneral(w []byte, nat_fields_mask uint32) (_ []byte, err error) {
	return item.WriteBoxed(w, nat_fields_mask), nil
}

func (item *UsefulServiceUserEntityPaymentItem) WriteBoxed(w []byte, nat_fields_mask uint32) []byte {
	w = basictl.NatWrite(w, 0x4f798680)
	return item.Write(w, nat_fields_mask)
}

func (item *UsefulServiceUserEntityPaymentItem) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, nat_fields_mask uint32) error {
	var propIdPresented bool
	var rawPromo []byte

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "id":
				if propIdPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("usefulService.userEntityPaymentItem", "id")
				}
				if err := Json2ReadString(in, &item.Id); err != nil {
					return err
				}
				propIdPresented = true
			case "promo":
				if rawPromo != nil {
					return ErrorInvalidJSONWithDuplicatingKeys("usefulService.userEntityPaymentItem", "promo")
				}
				rawPromo = in.Raw()
				if !in.Ok() {
					return in.Error()
				}
			default:
				return ErrorInvalidJSONExcessElement("usefulService.userEntityPaymentItem", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propIdPresented {
		item.Id = ""
	}
	var inPromoPointer *basictl.JsonLexer
	inPromo := basictl.JsonLexer{Data: rawPromo}
	if rawPromo != nil {
		inPromoPointer = &inPromo
	}
	if err := item.Promo.ReadJSON(legacyTypeNames, inPromoPointer, nat_fields_mask); err != nil {
		return err
	}

	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *UsefulServiceUserEntityPaymentItem) WriteJSONGeneral(w []byte, nat_fields_mask uint32) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w, nat_fields_mask), nil
}

func (item *UsefulServiceUserEntityPaymentItem) WriteJSON(w []byte, nat_fields_mask uint32) []byte {
	return item.WriteJSONOpt(true, false, w, nat_fields_mask)
}
func (item *UsefulServiceUserEntityPaymentItem) WriteJSONOpt(newTypeNames bool, short bool, w []byte, nat_fields_mask uint32) []byte {
	w = append(w, '{')
	backupIndexId := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"id":`...)
	w = basictl.JSONWriteString(w, item.Id)
	if (len(item.Id) != 0) == false {
		w = w[:backupIndexId]
	}
	backupIndexPromo := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"promo":`...)
	w = item.Promo.WriteJSONOpt(newTypeNames, short, w, nat_fields_mask)
	if (item.Promo.Ok) == false {
		w = w[:backupIndexPromo]
	}
	return append(w, '}')
}
