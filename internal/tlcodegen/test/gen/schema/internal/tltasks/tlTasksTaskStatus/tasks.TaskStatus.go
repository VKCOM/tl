// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlTasksTaskStatus

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

var _TasksTaskStatus = [4]internal.UnionElement{
	{TLTag: 0xb207caaa, TLName: "tasks.taskStatusNotCurrentlyInEngine", TLString: "tasks.taskStatusNotCurrentlyInEngine#b207caaa"},
	{TLTag: 0x0aca80a9, TLName: "tasks.taskStatusScheduled", TLString: "tasks.taskStatusScheduled#0aca80a9"},
	{TLTag: 0x16739c2c, TLName: "tasks.taskStatusWaiting", TLString: "tasks.taskStatusWaiting#16739c2c"},
	{TLTag: 0x06ef70e7, TLName: "tasks.taskStatusInProgress", TLString: "tasks.taskStatusInProgress#06ef70e7"},
}

func TasksTaskStatus__MakeEnum(i int) TasksTaskStatus { return TasksTaskStatus{index: i} }

type TasksTaskStatus struct {
	index int
}

func (item TasksTaskStatus) TLName() string { return _TasksTaskStatus[item.index].TLName }
func (item TasksTaskStatus) TLTag() uint32  { return _TasksTaskStatus[item.index].TLTag }

func (item *TasksTaskStatus) Reset() { item.index = 0 }

func (item TasksTaskStatus) IsNotCurrentlyInEngine() bool { return item.index == 0 }
func (item *TasksTaskStatus) SetNotCurrentlyInEngine()    { item.index = 0 }

func (item TasksTaskStatus) IsScheduled() bool { return item.index == 1 }
func (item *TasksTaskStatus) SetScheduled()    { item.index = 1 }

func (item TasksTaskStatus) IsWaiting() bool { return item.index == 2 }
func (item *TasksTaskStatus) SetWaiting()    { item.index = 2 }

func (item TasksTaskStatus) IsInProgress() bool { return item.index == 3 }
func (item *TasksTaskStatus) SetInProgress()    { item.index = 3 }

func (item *TasksTaskStatus) ReadBoxed(w []byte) (_ []byte, err error) {
	var tag uint32
	if w, err = basictl.NatRead(w, &tag); err != nil {
		return w, err
	}
	switch tag {
	case 0xb207caaa:
		item.index = 0
		return w, nil
	case 0x0aca80a9:
		item.index = 1
		return w, nil
	case 0x16739c2c:
		item.index = 2
		return w, nil
	case 0x06ef70e7:
		item.index = 3
		return w, nil
	default:
		return w, internal.ErrorInvalidUnionTag("tasks.TaskStatus", tag)
	}
}

func (item *TasksTaskStatus) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *TasksTaskStatus) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, _TasksTaskStatus[item.index].TLTag)
	return w
}

func (item *TasksTaskStatus) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	_jtype := in.UnsafeString()
	if !in.Ok() {
		return internal.ErrorInvalidJSON("tasks.TaskStatus", "expected string")
	}
	switch _jtype {
	case "tasks.taskStatusNotCurrentlyInEngine#b207caaa", "tasks.taskStatusNotCurrentlyInEngine", "#b207caaa":
		if !legacyTypeNames && _jtype == "tasks.taskStatusNotCurrentlyInEngine#b207caaa" {
			return internal.ErrorInvalidUnionLegacyTagJSON("tasks.TaskStatus", "tasks.taskStatusNotCurrentlyInEngine#b207caaa")
		}
		item.index = 0
		return nil
	case "tasks.taskStatusScheduled#0aca80a9", "tasks.taskStatusScheduled", "#0aca80a9":
		if !legacyTypeNames && _jtype == "tasks.taskStatusScheduled#0aca80a9" {
			return internal.ErrorInvalidUnionLegacyTagJSON("tasks.TaskStatus", "tasks.taskStatusScheduled#0aca80a9")
		}
		item.index = 1
		return nil
	case "tasks.taskStatusWaiting#16739c2c", "tasks.taskStatusWaiting", "#16739c2c":
		if !legacyTypeNames && _jtype == "tasks.taskStatusWaiting#16739c2c" {
			return internal.ErrorInvalidUnionLegacyTagJSON("tasks.TaskStatus", "tasks.taskStatusWaiting#16739c2c")
		}
		item.index = 2
		return nil
	case "tasks.taskStatusInProgress#06ef70e7", "tasks.taskStatusInProgress", "#06ef70e7":
		if !legacyTypeNames && _jtype == "tasks.taskStatusInProgress#06ef70e7" {
			return internal.ErrorInvalidUnionLegacyTagJSON("tasks.TaskStatus", "tasks.taskStatusInProgress#06ef70e7")
		}
		item.index = 3
		return nil
	default:
		return internal.ErrorInvalidEnumTagJSON("tasks.TaskStatus", _jtype)
	}
}

// This method is general version of WriteJSON, use it instead!
func (item TasksTaskStatus) WriteJSONGeneral(tctx *basictl.JSONWriteContext, w []byte) ([]byte, error) {
	return item.WriteJSONOpt(tctx, w), nil
}

func (item TasksTaskStatus) WriteJSON(w []byte) []byte {
	tctx := basictl.JSONWriteContext{}
	return item.WriteJSONOpt(&tctx, w)
}
func (item TasksTaskStatus) WriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte) []byte {
	w = append(w, '"')
	if tctx.LegacyTypeNames {
		w = append(w, _TasksTaskStatus[item.index].TLString...)
	} else {
		w = append(w, _TasksTaskStatus[item.index].TLName...)
	}
	return append(w, '"')
}

func (item TasksTaskStatus) String() string {
	return string(item.WriteJSON(nil))
}

func (item *TasksTaskStatus) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *TasksTaskStatus) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("tasks.TaskStatus", err.Error())
	}
	return nil
}
