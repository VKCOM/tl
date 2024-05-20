// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlBuiltinVectorAColor

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tla/tlAColor"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

func BuiltinVectorAColorFillRandom(rg *basictl.RandGenerator, vec *[]tlAColor.AColor) {
	rg.IncreaseDepth()
	l := rg.LimitValue(basictl.RandomUint(rg))
	*vec = make([]tlAColor.AColor, l)
	for i := range *vec {
		(*vec)[i].FillRandom(rg)
	}
	rg.DecreaseDepth()
}
func BuiltinVectorAColorRead(w []byte, vec *[]tlAColor.AColor) (_ []byte, err error) {
	var l uint32
	if w, err = basictl.NatRead(w, &l); err != nil {
		return w, err
	}
	if err = basictl.CheckLengthSanity(w, l, 4); err != nil {
		return w, err
	}
	if uint32(cap(*vec)) < l {
		*vec = make([]tlAColor.AColor, l)
	} else {
		*vec = (*vec)[:l]
	}
	for i := range *vec {
		if w, err = (*vec)[i].ReadBoxed(w); err != nil {
			return w, err
		}
	}
	return w, nil
}

func BuiltinVectorAColorWrite(w []byte, vec []tlAColor.AColor) []byte {
	w = basictl.NatWrite(w, uint32(len(vec)))
	for _, elem := range vec {
		w = elem.WriteBoxed(w)
	}
	return w
}

func BuiltinVectorAColorReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, vec *[]tlAColor.AColor) error {
	*vec = (*vec)[:cap(*vec)]
	index := 0
	if in != nil {
		in.Delim('[')
		if !in.Ok() {
			return internal.ErrorInvalidJSON("[]tlAColor.AColor", "expected json array")
		}
		for ; !in.IsDelim(']'); index++ {
			if len(*vec) <= index {
				var newValue tlAColor.AColor
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
			return internal.ErrorInvalidJSON("[]tlAColor.AColor", "expected json array's end")
		}
	}
	*vec = (*vec)[:index]
	return nil
}

func BuiltinVectorAColorWriteJSON(w []byte, vec []tlAColor.AColor) []byte {
	return BuiltinVectorAColorWriteJSONOpt(true, false, w, vec)
}
func BuiltinVectorAColorWriteJSONOpt(newTypeNames bool, short bool, w []byte, vec []tlAColor.AColor) []byte {
	w = append(w, '[')
	for _, elem := range vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = elem.WriteJSONOpt(newTypeNames, short, w)
	}
	return append(w, ']')
}
