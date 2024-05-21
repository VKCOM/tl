// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlTasksGetTaskFromQueue

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tl/tlBuiltinVectorInt"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tl/tlTasksTaskInfoMaybe"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type TasksGetTaskFromQueue struct {
	TypeName string
	QueueId  []int32
}

func (TasksGetTaskFromQueue) TLName() string { return "tasks.getTaskFromQueue" }
func (TasksGetTaskFromQueue) TLTag() uint32  { return 0x6a52b698 }

func (item *TasksGetTaskFromQueue) Reset() {
	item.TypeName = ""
	item.QueueId = item.QueueId[:0]
}

func (item *TasksGetTaskFromQueue) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.StringRead(w, &item.TypeName); err != nil {
		return w, err
	}
	return tlBuiltinVectorInt.BuiltinVectorIntRead(w, &item.QueueId)
}

// This method is general version of Write, use it instead!
func (item *TasksGetTaskFromQueue) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *TasksGetTaskFromQueue) Write(w []byte) []byte {
	w = basictl.StringWrite(w, item.TypeName)
	w = tlBuiltinVectorInt.BuiltinVectorIntWrite(w, item.QueueId)
	return w
}

func (item *TasksGetTaskFromQueue) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x6a52b698); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *TasksGetTaskFromQueue) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *TasksGetTaskFromQueue) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x6a52b698)
	return item.Write(w)
}

func (item *TasksGetTaskFromQueue) ReadResult(w []byte, ret *tlTasksTaskInfoMaybe.TasksTaskInfoMaybe) (_ []byte, err error) {
	return ret.ReadBoxed(w)
}

func (item *TasksGetTaskFromQueue) WriteResult(w []byte, ret tlTasksTaskInfoMaybe.TasksTaskInfoMaybe) (_ []byte, err error) {
	w = ret.WriteBoxed(w)
	return w, nil
}

func (item *TasksGetTaskFromQueue) ReadResultJSON(legacyTypeNames bool, in *basictl.JsonLexer, ret *tlTasksTaskInfoMaybe.TasksTaskInfoMaybe) error {
	if err := ret.ReadJSON(legacyTypeNames, in); err != nil {
		return err
	}
	return nil
}

func (item *TasksGetTaskFromQueue) WriteResultJSON(w []byte, ret tlTasksTaskInfoMaybe.TasksTaskInfoMaybe) (_ []byte, err error) {
	return item.writeResultJSON(true, false, w, ret)
}

func (item *TasksGetTaskFromQueue) writeResultJSON(newTypeNames bool, short bool, w []byte, ret tlTasksTaskInfoMaybe.TasksTaskInfoMaybe) (_ []byte, err error) {
	w = ret.WriteJSONOpt(newTypeNames, short, w)
	return w, nil
}

func (item *TasksGetTaskFromQueue) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret tlTasksTaskInfoMaybe.TasksTaskInfoMaybe
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *TasksGetTaskFromQueue) ReadResultWriteResultJSONOpt(newTypeNames bool, short bool, r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret tlTasksTaskInfoMaybe.TasksTaskInfoMaybe
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(newTypeNames, short, w, ret)
	return r, w, err
}

func (item *TasksGetTaskFromQueue) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	var ret tlTasksTaskInfoMaybe.TasksTaskInfoMaybe
	err := item.ReadResultJSON(true, &basictl.JsonLexer{Data: r}, &ret)
	if err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item TasksGetTaskFromQueue) String() string {
	return string(item.WriteJSON(nil))
}

func (item *TasksGetTaskFromQueue) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propTypeNamePresented bool
	var propQueueIdPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "type_name":
				if propTypeNamePresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("tasks.getTaskFromQueue", "type_name")
				}
				if err := internal.Json2ReadString(in, &item.TypeName); err != nil {
					return err
				}
				propTypeNamePresented = true
			case "queue_id":
				if propQueueIdPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("tasks.getTaskFromQueue", "queue_id")
				}
				if err := tlBuiltinVectorInt.BuiltinVectorIntReadJSON(legacyTypeNames, in, &item.QueueId); err != nil {
					return err
				}
				propQueueIdPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("tasks.getTaskFromQueue", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propTypeNamePresented {
		item.TypeName = ""
	}
	if !propQueueIdPresented {
		item.QueueId = item.QueueId[:0]
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *TasksGetTaskFromQueue) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *TasksGetTaskFromQueue) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *TasksGetTaskFromQueue) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexTypeName := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"type_name":`...)
	w = basictl.JSONWriteString(w, item.TypeName)
	if (len(item.TypeName) != 0) == false {
		w = w[:backupIndexTypeName]
	}
	backupIndexQueueId := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"queue_id":`...)
	w = tlBuiltinVectorInt.BuiltinVectorIntWriteJSONOpt(newTypeNames, short, w, item.QueueId)
	if (len(item.QueueId) != 0) == false {
		w = w[:backupIndexQueueId]
	}
	return append(w, '}')
}

func (item *TasksGetTaskFromQueue) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *TasksGetTaskFromQueue) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("tasks.getTaskFromQueue", err.Error())
	}
	return nil
}
