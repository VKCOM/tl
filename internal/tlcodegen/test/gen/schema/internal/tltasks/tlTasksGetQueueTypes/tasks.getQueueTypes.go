// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlTasksGetQueueTypes

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tl/tlBuiltinVectorTasksQueueTypeInfo"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tltasks/tlTasksQueueTypeInfo"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type TasksGetQueueTypes struct {
	SettingsMask uint32
	StatsMask    uint32
}

func (TasksGetQueueTypes) TLName() string { return "tasks.getQueueTypes" }
func (TasksGetQueueTypes) TLTag() uint32  { return 0x5434457a }

func (item *TasksGetQueueTypes) Reset() {
	item.SettingsMask = 0
	item.StatsMask = 0
}

func (item *TasksGetQueueTypes) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.SettingsMask); err != nil {
		return w, err
	}
	return basictl.NatRead(w, &item.StatsMask)
}

// This method is general version of Write, use it instead!
func (item *TasksGetQueueTypes) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *TasksGetQueueTypes) Write(w []byte) []byte {
	w = basictl.NatWrite(w, item.SettingsMask)
	w = basictl.NatWrite(w, item.StatsMask)
	return w
}

func (item *TasksGetQueueTypes) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x5434457a); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *TasksGetQueueTypes) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *TasksGetQueueTypes) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x5434457a)
	return item.Write(w)
}

func (item *TasksGetQueueTypes) ReadResult(w []byte, ret *[]tlTasksQueueTypeInfo.TasksQueueTypeInfo) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x1cb5c415); err != nil {
		return w, err
	}
	return tlBuiltinVectorTasksQueueTypeInfo.BuiltinVectorTasksQueueTypeInfoRead(w, ret)
}

func (item *TasksGetQueueTypes) WriteResult(w []byte, ret []tlTasksQueueTypeInfo.TasksQueueTypeInfo) (_ []byte, err error) {
	w = basictl.NatWrite(w, 0x1cb5c415)
	w = tlBuiltinVectorTasksQueueTypeInfo.BuiltinVectorTasksQueueTypeInfoWrite(w, ret)
	return w, nil
}

func (item *TasksGetQueueTypes) ReadResultJSON(legacyTypeNames bool, in *basictl.JsonLexer, ret *[]tlTasksQueueTypeInfo.TasksQueueTypeInfo) error {
	if err := tlBuiltinVectorTasksQueueTypeInfo.BuiltinVectorTasksQueueTypeInfoReadJSON(legacyTypeNames, in, ret); err != nil {
		return err
	}
	return nil
}

func (item *TasksGetQueueTypes) WriteResultJSON(w []byte, ret []tlTasksQueueTypeInfo.TasksQueueTypeInfo) (_ []byte, err error) {
	return item.writeResultJSON(true, false, w, ret)
}

func (item *TasksGetQueueTypes) writeResultJSON(newTypeNames bool, short bool, w []byte, ret []tlTasksQueueTypeInfo.TasksQueueTypeInfo) (_ []byte, err error) {
	w = tlBuiltinVectorTasksQueueTypeInfo.BuiltinVectorTasksQueueTypeInfoWriteJSONOpt(newTypeNames, short, w, ret)
	return w, nil
}

func (item *TasksGetQueueTypes) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret []tlTasksQueueTypeInfo.TasksQueueTypeInfo
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *TasksGetQueueTypes) ReadResultWriteResultJSONOpt(newTypeNames bool, short bool, r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret []tlTasksQueueTypeInfo.TasksQueueTypeInfo
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(newTypeNames, short, w, ret)
	return r, w, err
}

func (item *TasksGetQueueTypes) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	var ret []tlTasksQueueTypeInfo.TasksQueueTypeInfo
	err := item.ReadResultJSON(true, &basictl.JsonLexer{Data: r}, &ret)
	if err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item TasksGetQueueTypes) String() string {
	return string(item.WriteJSON(nil))
}

func (item *TasksGetQueueTypes) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propSettingsMaskPresented bool
	var propStatsMaskPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "settings_mask":
				if propSettingsMaskPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("tasks.getQueueTypes", "settings_mask")
				}
				if err := internal.Json2ReadUint32(in, &item.SettingsMask); err != nil {
					return err
				}
				propSettingsMaskPresented = true
			case "stats_mask":
				if propStatsMaskPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("tasks.getQueueTypes", "stats_mask")
				}
				if err := internal.Json2ReadUint32(in, &item.StatsMask); err != nil {
					return err
				}
				propStatsMaskPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("tasks.getQueueTypes", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propSettingsMaskPresented {
		item.SettingsMask = 0
	}
	if !propStatsMaskPresented {
		item.StatsMask = 0
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *TasksGetQueueTypes) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *TasksGetQueueTypes) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *TasksGetQueueTypes) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexSettingsMask := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"settings_mask":`...)
	w = basictl.JSONWriteUint32(w, item.SettingsMask)
	if (item.SettingsMask != 0) == false {
		w = w[:backupIndexSettingsMask]
	}
	backupIndexStatsMask := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"stats_mask":`...)
	w = basictl.JSONWriteUint32(w, item.StatsMask)
	if (item.StatsMask != 0) == false {
		w = w[:backupIndexStatsMask]
	}
	return append(w, '}')
}

func (item *TasksGetQueueTypes) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *TasksGetQueueTypes) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("tasks.getQueueTypes", err.Error())
	}
	return nil
}
