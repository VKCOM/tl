// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlService5Query

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/cycle_16847572a0831d4cd4c0c0fb513151f3"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tlservice5/tlService5Params"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type Service5Query struct {
	Query  string
	Params tlService5Params.Service5Params
}

func (Service5Query) TLName() string { return "service5.query" }
func (Service5Query) TLTag() uint32  { return 0xb3b62513 }

func (item *Service5Query) Reset() {
	item.Query = ""
	item.Params.Reset()
}

func (item *Service5Query) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.StringRead(w, &item.Query); err != nil {
		return w, err
	}
	return item.Params.Read(w)
}

// This method is general version of Write, use it instead!
func (item *Service5Query) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *Service5Query) Write(w []byte) []byte {
	w = basictl.StringWrite(w, item.Query)
	w = item.Params.Write(w)
	return w
}

func (item *Service5Query) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xb3b62513); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *Service5Query) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *Service5Query) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0xb3b62513)
	return item.Write(w)
}

func (item *Service5Query) ReadResult(w []byte, ret *cycle_16847572a0831d4cd4c0c0fb513151f3.Service5Output) (_ []byte, err error) {
	return ret.ReadBoxed(w)
}

func (item *Service5Query) WriteResult(w []byte, ret cycle_16847572a0831d4cd4c0c0fb513151f3.Service5Output) (_ []byte, err error) {
	w = ret.WriteBoxed(w)
	return w, nil
}

func (item *Service5Query) ReadResultJSON(legacyTypeNames bool, in *basictl.JsonLexer, ret *cycle_16847572a0831d4cd4c0c0fb513151f3.Service5Output) error {
	if err := ret.ReadJSON(legacyTypeNames, in); err != nil {
		return err
	}
	return nil
}

func (item *Service5Query) WriteResultJSON(w []byte, ret cycle_16847572a0831d4cd4c0c0fb513151f3.Service5Output) (_ []byte, err error) {
	return item.writeResultJSON(true, false, w, ret)
}

func (item *Service5Query) writeResultJSON(newTypeNames bool, short bool, w []byte, ret cycle_16847572a0831d4cd4c0c0fb513151f3.Service5Output) (_ []byte, err error) {
	w = ret.WriteJSONOpt(newTypeNames, short, w)
	return w, nil
}

func (item *Service5Query) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret cycle_16847572a0831d4cd4c0c0fb513151f3.Service5Output
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *Service5Query) ReadResultWriteResultJSONOpt(newTypeNames bool, short bool, r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret cycle_16847572a0831d4cd4c0c0fb513151f3.Service5Output
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(newTypeNames, short, w, ret)
	return r, w, err
}

func (item *Service5Query) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	var ret cycle_16847572a0831d4cd4c0c0fb513151f3.Service5Output
	err := item.ReadResultJSON(true, &basictl.JsonLexer{Data: r}, &ret)
	if err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item *Service5Query) String() string {
	return string(item.WriteJSON(nil))
}

func (item *Service5Query) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propQueryPresented bool
	var propParamsPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "query":
				if propQueryPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service5.query", "query")
				}
				if err := internal.Json2ReadString(in, &item.Query); err != nil {
					return err
				}
				propQueryPresented = true
			case "params":
				if propParamsPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service5.query", "params")
				}
				if err := item.Params.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propParamsPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("service5.query", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propQueryPresented {
		item.Query = ""
	}
	if !propParamsPresented {
		item.Params.Reset()
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *Service5Query) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *Service5Query) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *Service5Query) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexQuery := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"query":`...)
	w = basictl.JSONWriteString(w, item.Query)
	if (len(item.Query) != 0) == false {
		w = w[:backupIndexQuery]
	}
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"params":`...)
	w = item.Params.WriteJSONOpt(newTypeNames, short, w)
	return append(w, '}')
}

func (item *Service5Query) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *Service5Query) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("service5.query", err.Error())
	}
	return nil
}
