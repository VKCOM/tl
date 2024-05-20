// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlBuiltinTupleIntBoxed

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

func BuiltinTupleIntBoxedRead(w []byte, vec *[]int32, nat_n uint32) (_ []byte, err error) {
	if err = basictl.CheckLengthSanity(w, nat_n, 4); err != nil {
		return w, err
	}
	if uint32(cap(*vec)) < nat_n {
		*vec = make([]int32, nat_n)
	} else {
		*vec = (*vec)[:nat_n]
	}
	for i := range *vec {
		if w, err = basictl.NatReadExactTag(w, 0xa8509bda); err != nil {
			return w, err
		}
		if w, err = basictl.IntRead(w, &(*vec)[i]); err != nil {
			return w, err
		}
	}
	return w, nil
}

func BuiltinTupleIntBoxedWrite(w []byte, vec []int32, nat_n uint32) (_ []byte, err error) {
	if uint32(len(vec)) != nat_n {
		return w, internal.ErrorWrongSequenceLength("[]int32", len(vec), nat_n)
	}
	for _, elem := range vec {
		w = basictl.NatWrite(w, 0xa8509bda)
		w = basictl.IntWrite(w, elem)
	}
	return w, nil
}

func BuiltinTupleIntBoxedReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, vec *[]int32, nat_n uint32) error {
	if uint32(cap(*vec)) < nat_n {
		*vec = make([]int32, nat_n)
	} else {
		*vec = (*vec)[:nat_n]
	}
	index := 0
	if in != nil {
		in.Delim('[')
		if !in.Ok() {
			return internal.ErrorInvalidJSON("[]int32", "expected json array")
		}
		for ; !in.IsDelim(']'); index++ {
			if nat_n <= uint32(index) {
				return internal.ErrorInvalidJSON("[]int32", "array is longer than expected")
			}
			if err := internal.Json2ReadInt32(in, &(*vec)[index]); err != nil {
				return err
			}
			in.WantComma()
		}
		in.Delim(']')
		if !in.Ok() {
			return internal.ErrorInvalidJSON("[]int32", "expected json array's end")
		}
	}
	if uint32(index) != nat_n {
		return internal.ErrorWrongSequenceLength("[]int32", index, nat_n)
	}
	return nil
}

func BuiltinTupleIntBoxedWriteJSON(w []byte, vec []int32, nat_n uint32) (_ []byte, err error) {
	return BuiltinTupleIntBoxedWriteJSONOpt(true, false, w, vec, nat_n)
}
func BuiltinTupleIntBoxedWriteJSONOpt(newTypeNames bool, short bool, w []byte, vec []int32, nat_n uint32) (_ []byte, err error) {
	if uint32(len(vec)) != nat_n {
		return w, internal.ErrorWrongSequenceLength("[]int32", len(vec), nat_n)
	}
	w = append(w, '[')
	for _, elem := range vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = basictl.JSONWriteInt32(w, elem)
	}
	return append(w, ']'), nil
}
