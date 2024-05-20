// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlBuiltinVectorTasksQueueTypeInfo

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tltasks/tlTasksQueueTypeInfo"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

func BuiltinVectorTasksQueueTypeInfoRead(w []byte, vec *[]tlTasksQueueTypeInfo.TasksQueueTypeInfo) (_ []byte, err error) {
	var l uint32
	if w, err = basictl.NatRead(w, &l); err != nil {
		return w, err
	}
	if err = basictl.CheckLengthSanity(w, l, 4); err != nil {
		return w, err
	}
	if uint32(cap(*vec)) < l {
		*vec = make([]tlTasksQueueTypeInfo.TasksQueueTypeInfo, l)
	} else {
		*vec = (*vec)[:l]
	}
	for i := range *vec {
		if w, err = (*vec)[i].Read(w); err != nil {
			return w, err
		}
	}
	return w, nil
}

func BuiltinVectorTasksQueueTypeInfoWrite(w []byte, vec []tlTasksQueueTypeInfo.TasksQueueTypeInfo) []byte {
	w = basictl.NatWrite(w, uint32(len(vec)))
	for _, elem := range vec {
		w = elem.Write(w)
	}
	return w
}

func BuiltinVectorTasksQueueTypeInfoReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, vec *[]tlTasksQueueTypeInfo.TasksQueueTypeInfo) error {
	*vec = (*vec)[:cap(*vec)]
	index := 0
	if in != nil {
		in.Delim('[')
		if !in.Ok() {
			return internal.ErrorInvalidJSON("[]tlTasksQueueTypeInfo.TasksQueueTypeInfo", "expected json array")
		}
		for ; !in.IsDelim(']'); index++ {
			if len(*vec) <= index {
				var newValue tlTasksQueueTypeInfo.TasksQueueTypeInfo
				*vec = append(*vec, newValue)
				*vec = (*vec)[:cap(*vec)]
			}
			if err := (*vec)[index].ReadJSON(legacyTypeNames, in); err != nil {
				return err
			}
			in.WantComma()
		}
		in.Delim(']')
		if !in.Ok() {
			return internal.ErrorInvalidJSON("[]tlTasksQueueTypeInfo.TasksQueueTypeInfo", "expected json array's end")
		}
	}
	*vec = (*vec)[:index]
	return nil
}

func BuiltinVectorTasksQueueTypeInfoWriteJSON(w []byte, vec []tlTasksQueueTypeInfo.TasksQueueTypeInfo) []byte {
	return BuiltinVectorTasksQueueTypeInfoWriteJSONOpt(true, false, w, vec)
}
func BuiltinVectorTasksQueueTypeInfoWriteJSONOpt(newTypeNames bool, short bool, w []byte, vec []tlTasksQueueTypeInfo.TasksQueueTypeInfo) []byte {
	w = append(w, '[')
	for _, elem := range vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = elem.WriteJSONOpt(newTypeNames, short, w)
	}
	return append(w, ']')
}
