// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package internal

import (
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite

func BuiltinTuple3Replace10ElemReset(vec *[3]Replace10Elem) {
	for i := range *vec {
		(*vec)[i].Reset()
	}
}

func BuiltinTuple3Replace10ElemFillRandom(rg *basictl.RandGenerator, vec *[3]Replace10Elem, nat_t uint32) {
	rg.IncreaseDepth()
	for i := range *vec {
		(*vec)[i].FillRandom(rg, nat_t)
	}
	rg.DecreaseDepth()
}

func BuiltinTuple3Replace10ElemRead(w []byte, vec *[3]Replace10Elem, nat_t uint32) (_ []byte, err error) {
	for i := range *vec {
		if w, err = (*vec)[i].Read(w, nat_t); err != nil {
			return w, err
		}
	}
	return w, nil
}

func BuiltinTuple3Replace10ElemWrite(w []byte, vec *[3]Replace10Elem, nat_t uint32) []byte {
	for _, elem := range *vec {
		w = elem.Write(w, nat_t)
	}
	return w
}

func BuiltinTuple3Replace10ElemReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, vec *[3]Replace10Elem, nat_t uint32) error {
	index := 0
	if in != nil {
		in.Delim('[')
		if !in.Ok() {
			return ErrorInvalidJSON("[3]Replace10Elem", "expected json array")
		}
		for ; !in.IsDelim(']'); index++ {
			if index == 3 {
				return ErrorWrongSequenceLength("[3]Replace10Elem", index+1, 3)
			}
			if err := (*vec)[index].ReadJSON(legacyTypeNames, in, nat_t); err != nil {
				return err
			}
			in.WantComma()
		}
		in.Delim(']')
		if !in.Ok() {
			return ErrorInvalidJSON("[3]Replace10Elem", "expected json array's end")
		}
	}
	if index != 3 {
		return ErrorWrongSequenceLength("[3]Replace10Elem", index+1, 3)
	}
	return nil
}

func BuiltinTuple3Replace10ElemWriteJSON(w []byte, vec *[3]Replace10Elem, nat_t uint32) []byte {
	return BuiltinTuple3Replace10ElemWriteJSONOpt(true, false, w, vec, nat_t)
}
func BuiltinTuple3Replace10ElemWriteJSONOpt(newTypeNames bool, short bool, w []byte, vec *[3]Replace10Elem, nat_t uint32) []byte {
	w = append(w, '[')
	for _, elem := range *vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = elem.WriteJSONOpt(newTypeNames, short, w, nat_t)
	}
	return append(w, ']')
}

func BuiltinTupleTuple3Replace10ElemFillRandom(rg *basictl.RandGenerator, vec *[][3]Replace10Elem, nat_n uint32, nat_t uint32) {
	rg.IncreaseDepth()
	*vec = make([][3]Replace10Elem, nat_n)
	for i := range *vec {
		BuiltinTuple3Replace10ElemFillRandom(rg, &(*vec)[i], nat_t)
	}
	rg.DecreaseDepth()
}

func BuiltinTupleTuple3Replace10ElemRead(w []byte, vec *[][3]Replace10Elem, nat_n uint32, nat_t uint32) (_ []byte, err error) {
	if err = basictl.CheckLengthSanity(w, nat_n, 4); err != nil {
		return w, err
	}
	if uint32(cap(*vec)) < nat_n {
		*vec = make([][3]Replace10Elem, nat_n)
	} else {
		*vec = (*vec)[:nat_n]
	}
	for i := range *vec {
		if w, err = BuiltinTuple3Replace10ElemRead(w, &(*vec)[i], nat_t); err != nil {
			return w, err
		}
	}
	return w, nil
}

func BuiltinTupleTuple3Replace10ElemWrite(w []byte, vec [][3]Replace10Elem, nat_n uint32, nat_t uint32) (_ []byte, err error) {
	if uint32(len(vec)) != nat_n {
		return w, ErrorWrongSequenceLength("[][3]Replace10Elem", len(vec), nat_n)
	}
	for _, elem := range vec {
		w = BuiltinTuple3Replace10ElemWrite(w, &elem, nat_t)
	}
	return w, nil
}

func BuiltinTupleTuple3Replace10ElemReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, vec *[][3]Replace10Elem, nat_n uint32, nat_t uint32) error {
	if uint32(cap(*vec)) < nat_n {
		*vec = make([][3]Replace10Elem, nat_n)
	} else {
		*vec = (*vec)[:nat_n]
	}
	index := 0
	if in != nil {
		in.Delim('[')
		if !in.Ok() {
			return ErrorInvalidJSON("[][3]Replace10Elem", "expected json array")
		}
		for ; !in.IsDelim(']'); index++ {
			if nat_n <= uint32(index) {
				return ErrorInvalidJSON("[][3]Replace10Elem", "array is longer than expected")
			}
			if err := BuiltinTuple3Replace10ElemReadJSON(legacyTypeNames, in, &(*vec)[index], nat_t); err != nil {
				return err
			}
			in.WantComma()
		}
		in.Delim(']')
		if !in.Ok() {
			return ErrorInvalidJSON("[][3]Replace10Elem", "expected json array's end")
		}
	}
	if uint32(index) != nat_n {
		return ErrorWrongSequenceLength("[][3]Replace10Elem", index, nat_n)
	}
	return nil
}

func BuiltinTupleTuple3Replace10ElemWriteJSON(w []byte, vec [][3]Replace10Elem, nat_n uint32, nat_t uint32) (_ []byte, err error) {
	return BuiltinTupleTuple3Replace10ElemWriteJSONOpt(true, false, w, vec, nat_n, nat_t)
}
func BuiltinTupleTuple3Replace10ElemWriteJSONOpt(newTypeNames bool, short bool, w []byte, vec [][3]Replace10Elem, nat_n uint32, nat_t uint32) (_ []byte, err error) {
	if uint32(len(vec)) != nat_n {
		return w, ErrorWrongSequenceLength("[][3]Replace10Elem", len(vec), nat_n)
	}
	w = append(w, '[')
	for _, elem := range vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = BuiltinTuple3Replace10ElemWriteJSONOpt(newTypeNames, short, w, &elem, nat_t)
	}
	return append(w, ']'), nil
}

type Replace10Elem struct {
	A int32
	B int32
}

func (item *Replace10Elem) Reset() {
	item.A = 0
	item.B = 0
}

func (item *Replace10Elem) FillRandom(rg *basictl.RandGenerator, nat_n uint32) {
	item.A = basictl.RandomInt(rg)
	item.B = basictl.RandomInt(rg)
}

func (item *Replace10Elem) Read(w []byte, nat_n uint32) (_ []byte, err error) {
	if w, err = basictl.IntRead(w, &item.A); err != nil {
		return w, err
	}
	return basictl.IntRead(w, &item.B)
}

// This method is general version of Write, use it instead!
func (item *Replace10Elem) WriteGeneral(w []byte, nat_n uint32) (_ []byte, err error) {
	return item.Write(w, nat_n), nil
}

func (item *Replace10Elem) Write(w []byte, nat_n uint32) []byte {
	w = basictl.IntWrite(w, item.A)
	w = basictl.IntWrite(w, item.B)
	return w
}

func (item *Replace10Elem) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, nat_n uint32) error {
	var propAPresented bool
	var propBPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "a":
				if propAPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("replace10Elem", "a")
				}
				if err := Json2ReadInt32(in, &item.A); err != nil {
					return err
				}
				propAPresented = true
			case "b":
				if propBPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("replace10Elem", "b")
				}
				if err := Json2ReadInt32(in, &item.B); err != nil {
					return err
				}
				propBPresented = true
			default:
				return ErrorInvalidJSONExcessElement("replace10Elem", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propAPresented {
		item.A = 0
	}
	if !propBPresented {
		item.B = 0
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *Replace10Elem) WriteJSONGeneral(w []byte, nat_n uint32) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w, nat_n), nil
}

func (item *Replace10Elem) WriteJSON(w []byte, nat_n uint32) []byte {
	return item.WriteJSONOpt(true, false, w, nat_n)
}
func (item *Replace10Elem) WriteJSONOpt(newTypeNames bool, short bool, w []byte, nat_n uint32) []byte {
	w = append(w, '{')
	backupIndexA := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"a":`...)
	w = basictl.JSONWriteInt32(w, item.A)
	if (item.A != 0) == false {
		w = w[:backupIndexA]
	}
	backupIndexB := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"b":`...)
	w = basictl.JSONWriteInt32(w, item.B)
	if (item.B != 0) == false {
		w = w[:backupIndexB]
	}
	return append(w, '}')
}
