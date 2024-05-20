// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlBuiltinTuple3Replace14ElemElemLong

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlReplace14ElemElemLong"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

func BuiltinTuple3Replace14ElemElemLongReset(vec *[3]tlReplace14ElemElemLong.Replace14ElemElemLong) {
	for i := range *vec {
		(*vec)[i].Reset()
	}
}

func BuiltinTuple3Replace14ElemElemLongFillRandom(rg *basictl.RandGenerator, vec *[3]tlReplace14ElemElemLong.Replace14ElemElemLong, nat_tn uint32, nat_tk uint32) {
	rg.IncreaseDepth()
	for i := range *vec {
		(*vec)[i].FillRandom(rg, nat_tn, nat_tk)
	}
	rg.DecreaseDepth()
}

func BuiltinTuple3Replace14ElemElemLongRead(w []byte, vec *[3]tlReplace14ElemElemLong.Replace14ElemElemLong, nat_tn uint32, nat_tk uint32) (_ []byte, err error) {
	for i := range *vec {
		if w, err = (*vec)[i].Read(w, nat_tn, nat_tk); err != nil {
			return w, err
		}
	}
	return w, nil
}

func BuiltinTuple3Replace14ElemElemLongWrite(w []byte, vec *[3]tlReplace14ElemElemLong.Replace14ElemElemLong, nat_tn uint32, nat_tk uint32) (_ []byte, err error) {
	for _, elem := range *vec {
		if w, err = elem.Write(w, nat_tn, nat_tk); err != nil {
			return w, err
		}
	}
	return w, nil
}

func BuiltinTuple3Replace14ElemElemLongReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, vec *[3]tlReplace14ElemElemLong.Replace14ElemElemLong, nat_tn uint32, nat_tk uint32) error {
	index := 0
	if in != nil {
		in.Delim('[')
		if !in.Ok() {
			return internal.ErrorInvalidJSON("[3]tlReplace14ElemElemLong.Replace14ElemElemLong", "expected json array")
		}
		for ; !in.IsDelim(']'); index++ {
			if index == 3 {
				return internal.ErrorWrongSequenceLength("[3]tlReplace14ElemElemLong.Replace14ElemElemLong", index+1, 3)
			}
			if err := (*vec)[index].ReadJSON(legacyTypeNames, in, nat_tn, nat_tk); err != nil {
				return err
			}
			in.WantComma()
		}
		in.Delim(']')
		if !in.Ok() {
			return internal.ErrorInvalidJSON("[3]tlReplace14ElemElemLong.Replace14ElemElemLong", "expected json array's end")
		}
	}
	if index != 3 {
		return internal.ErrorWrongSequenceLength("[3]tlReplace14ElemElemLong.Replace14ElemElemLong", index+1, 3)
	}
	return nil
}

func BuiltinTuple3Replace14ElemElemLongWriteJSON(w []byte, vec *[3]tlReplace14ElemElemLong.Replace14ElemElemLong, nat_tn uint32, nat_tk uint32) (_ []byte, err error) {
	return BuiltinTuple3Replace14ElemElemLongWriteJSONOpt(true, false, w, vec, nat_tn, nat_tk)
}
func BuiltinTuple3Replace14ElemElemLongWriteJSONOpt(newTypeNames bool, short bool, w []byte, vec *[3]tlReplace14ElemElemLong.Replace14ElemElemLong, nat_tn uint32, nat_tk uint32) (_ []byte, err error) {
	w = append(w, '[')
	for _, elem := range *vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		if w, err = elem.WriteJSONOpt(newTypeNames, short, w, nat_tn, nat_tk); err != nil {
			return w, err
		}
	}
	return append(w, ']'), nil
}
