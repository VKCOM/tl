// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlListService5Output

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/cycle_16847572a0831d4cd4c0c0fb513151f3"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type ListService5Output struct {
	Flag uint32
	Head cycle_16847572a0831d4cd4c0c0fb513151f3.Service5Output // Conditional: item.Flag.0
	Tail *ListService5Output                                   // Conditional: item.Flag.0
}

func (ListService5Output) TLName() string { return "list" }
func (ListService5Output) TLTag() uint32  { return 0x02d80cdd }

func (item *ListService5Output) SetHead(v cycle_16847572a0831d4cd4c0c0fb513151f3.Service5Output) {
	item.Head = v
	item.Flag |= 1 << 0
}
func (item *ListService5Output) ClearHead() {
	item.Head.Reset()
	item.Flag &^= 1 << 0
}
func (item ListService5Output) IsSetHead() bool { return item.Flag&(1<<0) != 0 }

func (item *ListService5Output) SetTail(v ListService5Output) {
	if item.Tail == nil {
		var value ListService5Output
		item.Tail = &value
	}
	*item.Tail = v
	item.Flag |= 1 << 0
}
func (item *ListService5Output) ClearTail() {
	if item.Tail != nil {
		item.Tail.Reset()
	}
	item.Flag &^= 1 << 0
}
func (item ListService5Output) IsSetTail() bool { return item.Flag&(1<<0) != 0 }

func (item *ListService5Output) Reset() {
	item.Flag = 0
	item.Head.Reset()
	if item.Tail != nil {
		item.Tail.Reset()
	}
}

func (item *ListService5Output) FillRandom(rg *basictl.RandGenerator) {
	var maskFlag uint32
	maskFlag = basictl.RandomUint(rg)
	item.Flag = 0
	if maskFlag&(1<<0) != 0 {
		item.Flag |= (1 << 0)
	}
	if item.Flag&(1<<0) != 0 {
		item.Head.FillRandom(rg)
	} else {
		item.Head.Reset()
	}
	if item.Flag&(1<<0) != 0 {
		rg.IncreaseDepth()
		if item.Tail == nil {
			var value ListService5Output
			item.Tail = &value
		}
		item.Tail.FillRandom(rg)
		rg.DecreaseDepth()
	} else {
		if item.Tail != nil {
			item.Tail.Reset()
		}
	}
}

func (item *ListService5Output) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.Flag); err != nil {
		return w, err
	}
	if item.Flag&(1<<0) != 0 {
		if w, err = item.Head.ReadBoxed(w); err != nil {
			return w, err
		}
	} else {
		item.Head.Reset()
	}
	if item.Flag&(1<<0) != 0 {
		if item.Tail == nil {
			var value ListService5Output
			item.Tail = &value
		}
		if w, err = item.Tail.Read(w); err != nil {
			return w, err
		}
	} else {
		if item.Tail != nil {
			item.Tail.Reset()
		}
	}
	return w, nil
}

// This method is general version of Write, use it instead!
func (item *ListService5Output) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *ListService5Output) Write(w []byte) []byte {
	w = basictl.NatWrite(w, item.Flag)
	if item.Flag&(1<<0) != 0 {
		w = item.Head.WriteBoxed(w)
	}
	if item.Flag&(1<<0) != 0 {
		w = item.Tail.Write(w)
	}
	return w
}

func (item *ListService5Output) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x02d80cdd); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *ListService5Output) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *ListService5Output) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x02d80cdd)
	return item.Write(w)
}

func (item ListService5Output) String() string {
	return string(item.WriteJSON(nil))
}

func (item *ListService5Output) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propFlagPresented bool
	var propHeadPresented bool
	var propTailPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "flag":
				if propFlagPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("list", "flag")
				}
				if err := internal.Json2ReadUint32(in, &item.Flag); err != nil {
					return err
				}
				propFlagPresented = true
			case "head":
				if propHeadPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("list", "head")
				}
				if err := item.Head.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propHeadPresented = true
			case "tail":
				if propTailPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("list", "tail")
				}
				if item.Tail == nil {
					var value ListService5Output
					item.Tail = &value
				}
				if err := item.Tail.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propTailPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("list", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propFlagPresented {
		item.Flag = 0
	}
	if !propHeadPresented {
		item.Head.Reset()
	}
	if !propTailPresented {
		if item.Tail != nil {
			item.Tail.Reset()
		}
	}
	if propHeadPresented {
		item.Flag |= 1 << 0
	}
	if propTailPresented {
		item.Flag |= 1 << 0
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *ListService5Output) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *ListService5Output) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *ListService5Output) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexFlag := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"flag":`...)
	w = basictl.JSONWriteUint32(w, item.Flag)
	if (item.Flag != 0) == false {
		w = w[:backupIndexFlag]
	}
	if item.Flag&(1<<0) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"head":`...)
		w = item.Head.WriteJSONOpt(newTypeNames, short, w)
	}
	if item.Flag&(1<<0) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"tail":`...)
		w = item.Tail.WriteJSONOpt(newTypeNames, short, w)
	}
	return append(w, '}')
}

func (item *ListService5Output) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *ListService5Output) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("list", err.Error())
	}
	return nil
}
