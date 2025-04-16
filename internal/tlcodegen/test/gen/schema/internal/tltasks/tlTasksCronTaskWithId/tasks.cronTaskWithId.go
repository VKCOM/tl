// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlTasksCronTaskWithId

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tltasks/tlTasksCronTask"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type TasksCronTaskWithId struct {
	Id       int32
	NextTime int32
	Task     tlTasksCronTask.TasksCronTask
}

func (TasksCronTaskWithId) TLName() string { return "tasks.cronTaskWithId" }
func (TasksCronTaskWithId) TLTag() uint32  { return 0x3a958001 }

func (item *TasksCronTaskWithId) Reset() {
	item.Id = 0
	item.NextTime = 0
	item.Task.Reset()
}

func (item *TasksCronTaskWithId) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.IntRead(w, &item.Id); err != nil {
		return w, err
	}
	if w, err = basictl.IntRead(w, &item.NextTime); err != nil {
		return w, err
	}
	return item.Task.Read(w)
}

// This method is general version of Write, use it instead!
func (item *TasksCronTaskWithId) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *TasksCronTaskWithId) Write(w []byte) []byte {
	w = basictl.IntWrite(w, item.Id)
	w = basictl.IntWrite(w, item.NextTime)
	w = item.Task.Write(w)
	return w
}

func (item *TasksCronTaskWithId) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x3a958001); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *TasksCronTaskWithId) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *TasksCronTaskWithId) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x3a958001)
	return item.Write(w)
}

func (item *TasksCronTaskWithId) String() string {
	return string(item.WriteJSON(nil))
}

func (item *TasksCronTaskWithId) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propIdPresented bool
	var propNextTimePresented bool
	var propTaskPresented bool

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
					return internal.ErrorInvalidJSONWithDuplicatingKeys("tasks.cronTaskWithId", "id")
				}
				if err := internal.Json2ReadInt32(in, &item.Id); err != nil {
					return err
				}
				propIdPresented = true
			case "next_time":
				if propNextTimePresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("tasks.cronTaskWithId", "next_time")
				}
				if err := internal.Json2ReadInt32(in, &item.NextTime); err != nil {
					return err
				}
				propNextTimePresented = true
			case "task":
				if propTaskPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("tasks.cronTaskWithId", "task")
				}
				if err := item.Task.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propTaskPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("tasks.cronTaskWithId", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propIdPresented {
		item.Id = 0
	}
	if !propNextTimePresented {
		item.NextTime = 0
	}
	if !propTaskPresented {
		item.Task.Reset()
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *TasksCronTaskWithId) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *TasksCronTaskWithId) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *TasksCronTaskWithId) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexId := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"id":`...)
	w = basictl.JSONWriteInt32(w, item.Id)
	if (item.Id != 0) == false {
		w = w[:backupIndexId]
	}
	backupIndexNextTime := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"next_time":`...)
	w = basictl.JSONWriteInt32(w, item.NextTime)
	if (item.NextTime != 0) == false {
		w = w[:backupIndexNextTime]
	}
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"task":`...)
	w = item.Task.WriteJSONOpt(newTypeNames, short, w)
	return append(w, '}')
}

func (item *TasksCronTaskWithId) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *TasksCronTaskWithId) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("tasks.cronTaskWithId", err.Error())
	}
	return nil
}
