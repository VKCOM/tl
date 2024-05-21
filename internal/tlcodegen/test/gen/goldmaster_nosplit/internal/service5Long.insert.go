// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package internal

import (
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite

type Service5LongInsert struct {
	Flags uint32
	// Persistent (TrueType) // Conditional: item.Flags.0
}

func (Service5LongInsert) TLName() string { return "service5Long.insert" }
func (Service5LongInsert) TLTag() uint32  { return 0x7cf362bb }

func (item *Service5LongInsert) SetPersistent(v bool) {
	if v {
		item.Flags |= 1 << 0
	} else {
		item.Flags &^= 1 << 0
	}
}
func (item Service5LongInsert) IsSetPersistent() bool { return item.Flags&(1<<0) != 0 }

func (item *Service5LongInsert) Reset() {
	item.Flags = 0
}

func (item *Service5LongInsert) FillRandom(rg *basictl.RandGenerator) {
	var maskFlags uint32
	maskFlags = basictl.RandomUint(rg)
	item.Flags = 0
	if maskFlags&(1<<0) != 0 {
		item.Flags |= (1 << 0)
	}
}

func (item *Service5LongInsert) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.Flags); err != nil {
		return w, err
	}
	return w, nil
}

// This method is general version of Write, use it instead!
func (item *Service5LongInsert) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *Service5LongInsert) Write(w []byte) []byte {
	w = basictl.NatWrite(w, item.Flags)
	return w
}

func (item *Service5LongInsert) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x7cf362bb); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *Service5LongInsert) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *Service5LongInsert) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x7cf362bb)
	return item.Write(w)
}

func (item *Service5LongInsert) ReadResult(w []byte, ret *Service5LongOutput) (_ []byte, err error) {
	return ret.ReadBoxed(w)
}

func (item *Service5LongInsert) WriteResult(w []byte, ret Service5LongOutput) (_ []byte, err error) {
	w = ret.WriteBoxed(w)
	return w, nil
}

func (item *Service5LongInsert) ReadResultJSON(legacyTypeNames bool, in *basictl.JsonLexer, ret *Service5LongOutput) error {
	if err := ret.ReadJSON(legacyTypeNames, in); err != nil {
		return err
	}
	return nil
}

func (item *Service5LongInsert) WriteResultJSON(w []byte, ret Service5LongOutput) (_ []byte, err error) {
	return item.writeResultJSON(true, false, w, ret)
}

func (item *Service5LongInsert) writeResultJSON(newTypeNames bool, short bool, w []byte, ret Service5LongOutput) (_ []byte, err error) {
	w = ret.WriteJSONOpt(newTypeNames, short, w)
	return w, nil
}

func (item *Service5LongInsert) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret Service5LongOutput
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *Service5LongInsert) ReadResultWriteResultJSONOpt(newTypeNames bool, short bool, r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret Service5LongOutput
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(newTypeNames, short, w, ret)
	return r, w, err
}

func (item *Service5LongInsert) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	var ret Service5LongOutput
	err := item.ReadResultJSON(true, &basictl.JsonLexer{Data: r}, &ret)
	if err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item Service5LongInsert) String() string {
	return string(item.WriteJSON(nil))
}

func (item *Service5LongInsert) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
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
					return ErrorInvalidJSONWithDuplicatingKeys("service5Long.insert", "flags")
				}
				if err := Json2ReadUint32(in, &item.Flags); err != nil {
					return err
				}
				propFlagsPresented = true
			case "persistent":
				if trueTypePersistentPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("service5Long.insert", "persistent")
				}
				if err := Json2ReadBool(in, &trueTypePersistentValue); err != nil {
					return err
				}
				trueTypePersistentPresented = true
			default:
				return ErrorInvalidJSONExcessElement("service5Long.insert", key)
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
		return ErrorInvalidJSON("service5Long.insert", "fieldmask bit flags.0 is indefinite because of the contradictions in values")
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *Service5LongInsert) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *Service5LongInsert) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *Service5LongInsert) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
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

func (item *Service5LongInsert) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *Service5LongInsert) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("service5Long.insert", err.Error())
	}
	return nil
}