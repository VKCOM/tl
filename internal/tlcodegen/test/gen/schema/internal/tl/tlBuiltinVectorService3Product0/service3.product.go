// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlBuiltinVectorService3Product0

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tlservice3/tlService3Product0"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

func BuiltinVectorService3Product0Read(w []byte, vec *[]tlService3Product0.Service3Product0) (_ []byte, err error) {
	var l uint32
	if w, err = basictl.NatRead(w, &l); err != nil {
		return w, err
	}
	if err = basictl.CheckLengthSanity(w, l, 4); err != nil {
		return w, err
	}
	if uint32(cap(*vec)) < l {
		*vec = make([]tlService3Product0.Service3Product0, l)
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

func BuiltinVectorService3Product0Write(w []byte, vec []tlService3Product0.Service3Product0) []byte {
	w = basictl.NatWrite(w, uint32(len(vec)))
	for _, elem := range vec {
		w = elem.Write(w)
	}
	return w
}

func BuiltinVectorService3Product0ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, vec *[]tlService3Product0.Service3Product0) error {
	*vec = (*vec)[:cap(*vec)]
	index := 0
	if in != nil {
		in.Delim('[')
		if !in.Ok() {
			return internal.ErrorInvalidJSON("[]tlService3Product0.Service3Product0", "expected json array")
		}
		for ; !in.IsDelim(']'); index++ {
			if len(*vec) <= index {
				var newValue tlService3Product0.Service3Product0
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
			return internal.ErrorInvalidJSON("[]tlService3Product0.Service3Product0", "expected json array's end")
		}
	}
	*vec = (*vec)[:index]
	return nil
}

func BuiltinVectorService3Product0WriteJSON(w []byte, vec []tlService3Product0.Service3Product0) []byte {
	return BuiltinVectorService3Product0WriteJSONOpt(true, false, w, vec)
}
func BuiltinVectorService3Product0WriteJSONOpt(newTypeNames bool, short bool, w []byte, vec []tlService3Product0.Service3Product0) []byte {
	w = append(w, '[')
	for _, elem := range vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = elem.WriteJSONOpt(newTypeNames, short, w)
	}
	return append(w, ']')
}
