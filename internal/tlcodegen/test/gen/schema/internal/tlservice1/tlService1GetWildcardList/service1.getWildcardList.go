// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlService1GetWildcardList

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tl/tlBuiltinVectorString"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type Service1GetWildcardList struct {
	Prefix string
}

func (Service1GetWildcardList) TLName() string { return "service1.getWildcardList" }
func (Service1GetWildcardList) TLTag() uint32  { return 0x56b6ead4 }

func (item *Service1GetWildcardList) Reset() {
	item.Prefix = ""
}

func (item *Service1GetWildcardList) Read(w []byte) (_ []byte, err error) {
	return basictl.StringRead(w, &item.Prefix)
}

// This method is general version of Write, use it instead!
func (item *Service1GetWildcardList) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *Service1GetWildcardList) Write(w []byte) []byte {
	w = basictl.StringWrite(w, item.Prefix)
	return w
}

func (item *Service1GetWildcardList) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x56b6ead4); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *Service1GetWildcardList) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *Service1GetWildcardList) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x56b6ead4)
	return item.Write(w)
}

func (item *Service1GetWildcardList) ReadResult(w []byte, ret *[]string) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x1cb5c415); err != nil {
		return w, err
	}
	return tlBuiltinVectorString.BuiltinVectorStringRead(w, ret)
}

func (item *Service1GetWildcardList) WriteResult(w []byte, ret []string) (_ []byte, err error) {
	w = basictl.NatWrite(w, 0x1cb5c415)
	w = tlBuiltinVectorString.BuiltinVectorStringWrite(w, ret)
	return w, nil
}

func (item *Service1GetWildcardList) ReadResultJSON(legacyTypeNames bool, in *basictl.JsonLexer, ret *[]string) error {
	if err := tlBuiltinVectorString.BuiltinVectorStringReadJSON(legacyTypeNames, in, ret); err != nil {
		return err
	}
	return nil
}

func (item *Service1GetWildcardList) WriteResultJSON(w []byte, ret []string) (_ []byte, err error) {
	return item.writeResultJSON(true, false, w, ret)
}

func (item *Service1GetWildcardList) writeResultJSON(newTypeNames bool, short bool, w []byte, ret []string) (_ []byte, err error) {
	w = tlBuiltinVectorString.BuiltinVectorStringWriteJSONOpt(newTypeNames, short, w, ret)
	return w, nil
}

func (item *Service1GetWildcardList) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret []string
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *Service1GetWildcardList) ReadResultWriteResultJSONOpt(newTypeNames bool, short bool, r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret []string
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(newTypeNames, short, w, ret)
	return r, w, err
}

func (item *Service1GetWildcardList) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	var ret []string
	err := item.ReadResultJSON(true, &basictl.JsonLexer{Data: r}, &ret)
	if err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item Service1GetWildcardList) String() string {
	return string(item.WriteJSON(nil))
}

func (item *Service1GetWildcardList) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propPrefixPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "prefix":
				if propPrefixPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service1.getWildcardList", "prefix")
				}
				if err := internal.Json2ReadString(in, &item.Prefix); err != nil {
					return err
				}
				propPrefixPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("service1.getWildcardList", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propPrefixPresented {
		item.Prefix = ""
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *Service1GetWildcardList) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *Service1GetWildcardList) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *Service1GetWildcardList) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexPrefix := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"prefix":`...)
	w = basictl.JSONWriteString(w, item.Prefix)
	if (len(item.Prefix) != 0) == false {
		w = w[:backupIndexPrefix]
	}
	return append(w, '}')
}

func (item *Service1GetWildcardList) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *Service1GetWildcardList) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("service1.getWildcardList", err.Error())
	}
	return nil
}