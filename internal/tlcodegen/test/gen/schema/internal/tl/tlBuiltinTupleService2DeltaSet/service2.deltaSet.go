// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlBuiltinTupleService2DeltaSet

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tlservice2/tlService2DeltaSet"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

func BuiltinTupleService2DeltaSetRead(w []byte, vec *[]tlService2DeltaSet.Service2DeltaSet, nat_n uint32, nat_tobjectIdLength uint32, nat_tintCountersNum uint32, nat_tfloatCountersNum uint32) (_ []byte, err error) {
	if err = basictl.CheckLengthSanity(w, nat_n, 4); err != nil {
		return w, err
	}
	if uint32(cap(*vec)) < nat_n {
		*vec = make([]tlService2DeltaSet.Service2DeltaSet, nat_n)
	} else {
		*vec = (*vec)[:nat_n]
	}
	for i := range *vec {
		if w, err = (*vec)[i].Read(w, nat_tobjectIdLength, nat_tintCountersNum, nat_tfloatCountersNum); err != nil {
			return w, err
		}
	}
	return w, nil
}

func BuiltinTupleService2DeltaSetWrite(w []byte, vec []tlService2DeltaSet.Service2DeltaSet, nat_n uint32, nat_tobjectIdLength uint32, nat_tintCountersNum uint32, nat_tfloatCountersNum uint32) (_ []byte, err error) {
	if uint32(len(vec)) != nat_n {
		return w, internal.ErrorWrongSequenceLength("[]tlService2DeltaSet.Service2DeltaSet", len(vec), nat_n)
	}
	for _, elem := range vec {
		if w, err = elem.Write(w, nat_tobjectIdLength, nat_tintCountersNum, nat_tfloatCountersNum); err != nil {
			return w, err
		}
	}
	return w, nil
}

func BuiltinTupleService2DeltaSetReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, vec *[]tlService2DeltaSet.Service2DeltaSet, nat_n uint32, nat_tobjectIdLength uint32, nat_tintCountersNum uint32, nat_tfloatCountersNum uint32) error {
	if uint32(cap(*vec)) < nat_n {
		*vec = make([]tlService2DeltaSet.Service2DeltaSet, nat_n)
	} else {
		*vec = (*vec)[:nat_n]
	}
	index := 0
	if in != nil {
		in.Delim('[')
		if !in.Ok() {
			return internal.ErrorInvalidJSON("[]tlService2DeltaSet.Service2DeltaSet", "expected json array")
		}
		for ; !in.IsDelim(']'); index++ {
			if nat_n <= uint32(index) {
				return internal.ErrorInvalidJSON("[]tlService2DeltaSet.Service2DeltaSet", "array is longer than expected")
			}
			if err := (*vec)[index].ReadJSON(legacyTypeNames, in, nat_tobjectIdLength, nat_tintCountersNum, nat_tfloatCountersNum); err != nil {
				return err
			}
			in.WantComma()
		}
		in.Delim(']')
		if !in.Ok() {
			return internal.ErrorInvalidJSON("[]tlService2DeltaSet.Service2DeltaSet", "expected json array's end")
		}
	}
	if uint32(index) != nat_n {
		return internal.ErrorWrongSequenceLength("[]tlService2DeltaSet.Service2DeltaSet", index, nat_n)
	}
	return nil
}

func BuiltinTupleService2DeltaSetWriteJSON(w []byte, vec []tlService2DeltaSet.Service2DeltaSet, nat_n uint32, nat_tobjectIdLength uint32, nat_tintCountersNum uint32, nat_tfloatCountersNum uint32) (_ []byte, err error) {
	return BuiltinTupleService2DeltaSetWriteJSONOpt(true, false, w, vec, nat_n, nat_tobjectIdLength, nat_tintCountersNum, nat_tfloatCountersNum)
}
func BuiltinTupleService2DeltaSetWriteJSONOpt(newTypeNames bool, short bool, w []byte, vec []tlService2DeltaSet.Service2DeltaSet, nat_n uint32, nat_tobjectIdLength uint32, nat_tintCountersNum uint32, nat_tfloatCountersNum uint32) (_ []byte, err error) {
	if uint32(len(vec)) != nat_n {
		return w, internal.ErrorWrongSequenceLength("[]tlService2DeltaSet.Service2DeltaSet", len(vec), nat_n)
	}
	w = append(w, '[')
	for _, elem := range vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		if w, err = elem.WriteJSONOpt(newTypeNames, short, w, nat_tobjectIdLength, nat_tintCountersNum, nat_tfloatCountersNum); err != nil {
			return w, err
		}
	}
	return append(w, ']'), nil
}
