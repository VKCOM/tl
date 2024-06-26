// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlService5Insert

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/cycle_16847572a0831d4cd4c0c0fb513151f3"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type Service5Insert struct {
	Flags uint32
	// Persistent (TrueType) // Conditional: item.Flags.0
}

func (Service5Insert) TLName() string { return "service5.insert" }
func (Service5Insert) TLTag() uint32  { return 0x7cf362ba }

func (item *Service5Insert) SetPersistent(v bool) {
	if v {
		item.Flags |= 1 << 0
	} else {
		item.Flags &^= 1 << 0
	}
}
func (item Service5Insert) IsSetPersistent() bool { return item.Flags&(1<<0) != 0 }

func (item *Service5Insert) Reset() {
	item.Flags = 0
}

func (item *Service5Insert) FillRandom(rg *basictl.RandGenerator) {
	var maskFlags uint32
	maskFlags = basictl.RandomUint(rg)
	item.Flags = 0
	if maskFlags&(1<<0) != 0 {
		item.Flags |= (1 << 0)
	}
}

func (item *Service5Insert) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.Flags); err != nil {
		return w, err
	}
	return w, nil
}

// This method is general version of Write, use it instead!
func (item *Service5Insert) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *Service5Insert) Write(w []byte) []byte {
	w = basictl.NatWrite(w, item.Flags)
	return w
}

func (item *Service5Insert) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x7cf362ba); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *Service5Insert) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *Service5Insert) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x7cf362ba)
	return item.Write(w)
}

func (item *Service5Insert) ReadResult(w []byte, ret *cycle_16847572a0831d4cd4c0c0fb513151f3.Service5Output) (_ []byte, err error) {
	return ret.ReadBoxed(w)
}

func (item *Service5Insert) WriteResult(w []byte, ret cycle_16847572a0831d4cd4c0c0fb513151f3.Service5Output) (_ []byte, err error) {
	w = ret.WriteBoxed(w)
	return w, nil
}

func (item *Service5Insert) ReadResultJSON(legacyTypeNames bool, in *basictl.JsonLexer, ret *cycle_16847572a0831d4cd4c0c0fb513151f3.Service5Output) error {
	if err := ret.ReadJSON(legacyTypeNames, in); err != nil {
		return err
	}
	return nil
}

func (item *Service5Insert) WriteResultJSON(w []byte, ret cycle_16847572a0831d4cd4c0c0fb513151f3.Service5Output) (_ []byte, err error) {
	return item.writeResultJSON(true, false, w, ret)
}

func (item *Service5Insert) writeResultJSON(newTypeNames bool, short bool, w []byte, ret cycle_16847572a0831d4cd4c0c0fb513151f3.Service5Output) (_ []byte, err error) {
	w = ret.WriteJSONOpt(newTypeNames, short, w)
	return w, nil
}

func (item *Service5Insert) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret cycle_16847572a0831d4cd4c0c0fb513151f3.Service5Output
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *Service5Insert) ReadResultWriteResultJSONOpt(newTypeNames bool, short bool, r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret cycle_16847572a0831d4cd4c0c0fb513151f3.Service5Output
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(newTypeNames, short, w, ret)
	return r, w, err
}

func (item *Service5Insert) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	var ret cycle_16847572a0831d4cd4c0c0fb513151f3.Service5Output
	err := item.ReadResultJSON(true, &basictl.JsonLexer{Data: r}, &ret)
	if err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item Service5Insert) String() string {
	return string(item.WriteJSON(nil))
}

func (item *Service5Insert) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propFlagsPresented bool
	var trueTypePersistentPresented bool
	var trueTypePersistentValue bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "flags":
				if propFlagsPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service5.insert", "flags")
				}
				if err := internal.Json2ReadUint32(in, &item.Flags); err != nil {
					return err
				}
				propFlagsPresented = true
			case "persistent":
				if trueTypePersistentPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service5.insert", "persistent")
				}
				if err := internal.Json2ReadBool(in, &trueTypePersistentValue); err != nil {
					return err
				}
				trueTypePersistentPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("service5.insert", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propFlagsPresented {
		item.Flags = 0
	}
	if trueTypePersistentPresented {
		if trueTypePersistentValue {
			item.Flags |= 1 << 0
		}
	}
	// tries to set bit to zero if it is 1
	if trueTypePersistentPresented && !trueTypePersistentValue && (item.Flags&(1<<0) != 0) {
		return internal.ErrorInvalidJSON("service5.insert", "fieldmask bit flags.0 is indefinite because of the contradictions in values")
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *Service5Insert) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *Service5Insert) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *Service5Insert) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexFlags := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"flags":`...)
	w = basictl.JSONWriteUint32(w, item.Flags)
	if (item.Flags != 0) == false {
		w = w[:backupIndexFlags]
	}
	if item.Flags&(1<<0) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"persistent":true`...)
	}
	return append(w, '}')
}

func (item *Service5Insert) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *Service5Insert) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("service5.insert", err.Error())
	}
	return nil
}
