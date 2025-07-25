// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlService1GetWildcardDict

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tl/tlBuiltinVectorDictionaryFieldString"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type Service1GetWildcardDict struct {
	Prefix string
}

func (Service1GetWildcardDict) TLName() string { return "service1.getWildcardDict" }
func (Service1GetWildcardDict) TLTag() uint32  { return 0x72bbc81b }

func (item *Service1GetWildcardDict) Reset() {
	item.Prefix = ""
}

func (item *Service1GetWildcardDict) Read(w []byte) (_ []byte, err error) {
	return basictl.StringRead(w, &item.Prefix)
}

func (item *Service1GetWildcardDict) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *Service1GetWildcardDict) Write(w []byte) []byte {
	w = basictl.StringWrite(w, item.Prefix)
	return w
}

func (item *Service1GetWildcardDict) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x72bbc81b); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *Service1GetWildcardDict) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *Service1GetWildcardDict) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x72bbc81b)
	return item.Write(w)
}

func (item *Service1GetWildcardDict) ReadResult(w []byte, ret *map[string]string) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x1f4c618f); err != nil {
		return w, err
	}
	return tlBuiltinVectorDictionaryFieldString.BuiltinVectorDictionaryFieldStringRead(w, ret)
}

func (item *Service1GetWildcardDict) WriteResult(w []byte, ret map[string]string) (_ []byte, err error) {
	w = basictl.NatWrite(w, 0x1f4c618f)
	w = tlBuiltinVectorDictionaryFieldString.BuiltinVectorDictionaryFieldStringWrite(w, ret)
	return w, nil
}

func (item *Service1GetWildcardDict) ReadResultJSON(legacyTypeNames bool, in *basictl.JsonLexer, ret *map[string]string) error {
	if err := tlBuiltinVectorDictionaryFieldString.BuiltinVectorDictionaryFieldStringReadJSON(legacyTypeNames, in, ret); err != nil {
		return err
	}
	return nil
}

func (item *Service1GetWildcardDict) WriteResultJSON(w []byte, ret map[string]string) (_ []byte, err error) {
	tctx := basictl.JSONWriteContext{}
	return item.writeResultJSON(&tctx, w, ret)
}

func (item *Service1GetWildcardDict) writeResultJSON(tctx *basictl.JSONWriteContext, w []byte, ret map[string]string) (_ []byte, err error) {
	w = tlBuiltinVectorDictionaryFieldString.BuiltinVectorDictionaryFieldStringWriteJSONOpt(tctx, w, ret)
	return w, nil
}

func (item *Service1GetWildcardDict) ReadResultWriteResultJSON(tctx *basictl.JSONWriteContext, r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret map[string]string
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(tctx, w, ret)
	return r, w, err
}

func (item *Service1GetWildcardDict) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	var ret map[string]string
	err := item.ReadResultJSON(true, &basictl.JsonLexer{Data: r}, &ret)
	if err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item Service1GetWildcardDict) String() string {
	return string(item.WriteJSON(nil))
}

func (item *Service1GetWildcardDict) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
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
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service1.getWildcardDict", "prefix")
				}
				if err := internal.Json2ReadString(in, &item.Prefix); err != nil {
					return err
				}
				propPrefixPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("service1.getWildcardDict", key)
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
func (item *Service1GetWildcardDict) WriteJSONGeneral(tctx *basictl.JSONWriteContext, w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(tctx, w), nil
}

func (item *Service1GetWildcardDict) WriteJSON(w []byte) []byte {
	tctx := basictl.JSONWriteContext{}
	return item.WriteJSONOpt(&tctx, w)
}
func (item *Service1GetWildcardDict) WriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte) []byte {
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

func (item *Service1GetWildcardDict) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *Service1GetWildcardDict) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("service1.getWildcardDict", err.Error())
	}
	return nil
}
