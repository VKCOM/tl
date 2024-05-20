// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlReplace

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlReplace1"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlReplace10"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlReplace11Long"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlReplace12"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlReplace13"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlReplace13Long"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlReplace14Long"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlReplace15"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlReplace2"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlReplace3"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlReplace4"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlReplace43"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlReplace5"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlReplace6"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlReplace7"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlReplace8"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlReplace9"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type Replace struct {
	N  uint32
	A  tlReplace1.Replace1
	A1 tlReplace13.Replace13
	B  tlReplace2.Replace2
	C  tlReplace3.Replace3
	D  tlReplace4.Replace4
	D1 tlReplace43.Replace43
	E  tlReplace5.Replace5
	G  tlReplace6.Replace6
	H  tlReplace7.Replace7
	I  tlReplace8.Replace8
	J  tlReplace9.Replace9
	K  tlReplace10.Replace10
	L  tlReplace11Long.Replace11Long
	M  tlReplace12.Replace12
	O  tlReplace13Long.Replace13Long
	P  tlReplace14Long.Replace14Long
	Q  tlReplace15.Replace15
}

func (Replace) TLName() string { return "replace" }
func (Replace) TLTag() uint32  { return 0x323db63e }

func (item *Replace) Reset() {
	item.N = 0
	item.A.Reset()
	item.A1.Reset()
	item.B.Reset()
	item.C.Reset()
	item.D.Reset()
	item.D1.Reset()
	item.E.Reset()
	item.G.Reset()
	item.H.Reset()
	item.I.Reset()
	item.J.Reset()
	item.K.Reset()
	item.L.Reset()
	item.M.Reset()
	item.O.Reset()
	item.P.Reset()
	item.Q.Reset()
}

func (item *Replace) FillRandom(rg *basictl.RandGenerator) {
	var maskN uint32
	maskN = basictl.RandomUint(rg)
	maskN = rg.LimitValue(maskN)
	item.N = 0
	if maskN&(1<<0) != 0 {
		item.N |= (1 << 0)
	}
	item.A.FillRandom(rg, item.N)
	item.A1.FillRandom(rg)
	item.B.FillRandom(rg)
	item.C.FillRandom(rg)
	item.D.FillRandom(rg, item.N)
	item.D1.FillRandom(rg)
	item.E.FillRandom(rg)
	item.G.FillRandom(rg)
	item.H.FillRandom(rg)
	item.I.FillRandom(rg)
	item.J.FillRandom(rg)
	item.K.FillRandom(rg)
	item.L.FillRandom(rg, item.N)
	item.M.FillRandom(rg)
	item.O.FillRandom(rg, item.N)
	item.P.FillRandom(rg, item.N)
	item.Q.FillRandom(rg)
}

func (item *Replace) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.N); err != nil {
		return w, err
	}
	if w, err = item.A.Read(w, item.N); err != nil {
		return w, err
	}
	if w, err = item.A1.Read(w); err != nil {
		return w, err
	}
	if w, err = item.B.Read(w); err != nil {
		return w, err
	}
	if w, err = item.C.Read(w); err != nil {
		return w, err
	}
	if w, err = item.D.Read(w, item.N); err != nil {
		return w, err
	}
	if w, err = item.D1.Read(w); err != nil {
		return w, err
	}
	if w, err = item.E.Read(w); err != nil {
		return w, err
	}
	if w, err = item.G.Read(w); err != nil {
		return w, err
	}
	if w, err = item.H.Read(w); err != nil {
		return w, err
	}
	if w, err = item.I.Read(w); err != nil {
		return w, err
	}
	if w, err = item.J.Read(w); err != nil {
		return w, err
	}
	if w, err = item.K.Read(w); err != nil {
		return w, err
	}
	if w, err = item.L.Read(w, item.N); err != nil {
		return w, err
	}
	if w, err = item.M.Read(w); err != nil {
		return w, err
	}
	if w, err = item.O.Read(w, item.N); err != nil {
		return w, err
	}
	if w, err = item.P.Read(w, item.N); err != nil {
		return w, err
	}
	return item.Q.Read(w)
}

// This method is general version of Write, use it instead!
func (item *Replace) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w)
}

func (item *Replace) Write(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, item.N)
	if w, err = item.A.Write(w, item.N); err != nil {
		return w, err
	}
	w = item.A1.Write(w)
	if w, err = item.B.Write(w); err != nil {
		return w, err
	}
	w = item.C.Write(w)
	if w, err = item.D.Write(w, item.N); err != nil {
		return w, err
	}
	w = item.D1.Write(w)
	if w, err = item.E.Write(w); err != nil {
		return w, err
	}
	w = item.G.Write(w)
	if w, err = item.H.Write(w); err != nil {
		return w, err
	}
	if w, err = item.I.Write(w); err != nil {
		return w, err
	}
	if w, err = item.J.Write(w); err != nil {
		return w, err
	}
	if w, err = item.K.Write(w); err != nil {
		return w, err
	}
	if w, err = item.L.Write(w, item.N); err != nil {
		return w, err
	}
	if w, err = item.M.Write(w); err != nil {
		return w, err
	}
	if w, err = item.O.Write(w, item.N); err != nil {
		return w, err
	}
	if w, err = item.P.Write(w, item.N); err != nil {
		return w, err
	}
	if w, err = item.Q.Write(w); err != nil {
		return w, err
	}
	return w, nil
}

func (item *Replace) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x323db63e); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *Replace) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w)
}

func (item *Replace) WriteBoxed(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, 0x323db63e)
	return item.Write(w)
}

func (item Replace) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func (item *Replace) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propNPresented bool
	var rawA []byte
	var propA1Presented bool
	var propBPresented bool
	var propCPresented bool
	var rawD []byte
	var propD1Presented bool
	var propEPresented bool
	var propGPresented bool
	var propHPresented bool
	var propIPresented bool
	var propJPresented bool
	var propKPresented bool
	var rawL []byte
	var propMPresented bool
	var rawO []byte
	var rawP []byte
	var propQPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "n":
				if propNPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("replace", "n")
				}
				if err := internal.Json2ReadUint32(in, &item.N); err != nil {
					return err
				}
				propNPresented = true
			case "a":
				if rawA != nil {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("replace", "a")
				}
				rawA = in.Raw()
				if !in.Ok() {
					return in.Error()
				}
			case "a1":
				if propA1Presented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("replace", "a1")
				}
				if err := item.A1.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propA1Presented = true
			case "b":
				if propBPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("replace", "b")
				}
				if err := item.B.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propBPresented = true
			case "c":
				if propCPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("replace", "c")
				}
				if err := item.C.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propCPresented = true
			case "d":
				if rawD != nil {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("replace", "d")
				}
				rawD = in.Raw()
				if !in.Ok() {
					return in.Error()
				}
			case "d1":
				if propD1Presented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("replace", "d1")
				}
				if err := item.D1.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propD1Presented = true
			case "e":
				if propEPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("replace", "e")
				}
				if err := item.E.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propEPresented = true
			case "g":
				if propGPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("replace", "g")
				}
				if err := item.G.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propGPresented = true
			case "h":
				if propHPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("replace", "h")
				}
				if err := item.H.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propHPresented = true
			case "i":
				if propIPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("replace", "i")
				}
				if err := item.I.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propIPresented = true
			case "j":
				if propJPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("replace", "j")
				}
				if err := item.J.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propJPresented = true
			case "k":
				if propKPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("replace", "k")
				}
				if err := item.K.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propKPresented = true
			case "l":
				if rawL != nil {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("replace", "l")
				}
				rawL = in.Raw()
				if !in.Ok() {
					return in.Error()
				}
			case "m":
				if propMPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("replace", "m")
				}
				if err := item.M.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propMPresented = true
			case "o":
				if rawO != nil {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("replace", "o")
				}
				rawO = in.Raw()
				if !in.Ok() {
					return in.Error()
				}
			case "p":
				if rawP != nil {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("replace", "p")
				}
				rawP = in.Raw()
				if !in.Ok() {
					return in.Error()
				}
			case "q":
				if propQPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("replace", "q")
				}
				if err := item.Q.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propQPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("replace", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propNPresented {
		item.N = 0
	}
	if !propA1Presented {
		item.A1.Reset()
	}
	if !propBPresented {
		item.B.Reset()
	}
	if !propCPresented {
		item.C.Reset()
	}
	if !propD1Presented {
		item.D1.Reset()
	}
	if !propEPresented {
		item.E.Reset()
	}
	if !propGPresented {
		item.G.Reset()
	}
	if !propHPresented {
		item.H.Reset()
	}
	if !propIPresented {
		item.I.Reset()
	}
	if !propJPresented {
		item.J.Reset()
	}
	if !propKPresented {
		item.K.Reset()
	}
	if !propMPresented {
		item.M.Reset()
	}
	if !propQPresented {
		item.Q.Reset()
	}
	var inAPointer *basictl.JsonLexer
	inA := basictl.JsonLexer{Data: rawA}
	if rawA != nil {
		inAPointer = &inA
	}
	if err := item.A.ReadJSON(legacyTypeNames, inAPointer, item.N); err != nil {
		return err
	}

	var inDPointer *basictl.JsonLexer
	inD := basictl.JsonLexer{Data: rawD}
	if rawD != nil {
		inDPointer = &inD
	}
	if err := item.D.ReadJSON(legacyTypeNames, inDPointer, item.N); err != nil {
		return err
	}

	var inLPointer *basictl.JsonLexer
	inL := basictl.JsonLexer{Data: rawL}
	if rawL != nil {
		inLPointer = &inL
	}
	if err := item.L.ReadJSON(legacyTypeNames, inLPointer, item.N); err != nil {
		return err
	}

	var inOPointer *basictl.JsonLexer
	inO := basictl.JsonLexer{Data: rawO}
	if rawO != nil {
		inOPointer = &inO
	}
	if err := item.O.ReadJSON(legacyTypeNames, inOPointer, item.N); err != nil {
		return err
	}

	var inPPointer *basictl.JsonLexer
	inP := basictl.JsonLexer{Data: rawP}
	if rawP != nil {
		inPPointer = &inP
	}
	if err := item.P.ReadJSON(legacyTypeNames, inPPointer, item.N); err != nil {
		return err
	}

	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *Replace) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w)
}

func (item *Replace) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w)
}
func (item *Replace) WriteJSONOpt(newTypeNames bool, short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	backupIndexN := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"n":`...)
	w = basictl.JSONWriteUint32(w, item.N)
	if (item.N != 0) == false {
		w = w[:backupIndexN]
	}
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"a":`...)
	if w, err = item.A.WriteJSONOpt(newTypeNames, short, w, item.N); err != nil {
		return w, err
	}
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"a1":`...)
	w = item.A1.WriteJSONOpt(newTypeNames, short, w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"b":`...)
	if w, err = item.B.WriteJSONOpt(newTypeNames, short, w); err != nil {
		return w, err
	}
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"c":`...)
	w = item.C.WriteJSONOpt(newTypeNames, short, w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"d":`...)
	if w, err = item.D.WriteJSONOpt(newTypeNames, short, w, item.N); err != nil {
		return w, err
	}
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"d1":`...)
	w = item.D1.WriteJSONOpt(newTypeNames, short, w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"e":`...)
	if w, err = item.E.WriteJSONOpt(newTypeNames, short, w); err != nil {
		return w, err
	}
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"g":`...)
	w = item.G.WriteJSONOpt(newTypeNames, short, w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"h":`...)
	if w, err = item.H.WriteJSONOpt(newTypeNames, short, w); err != nil {
		return w, err
	}
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"i":`...)
	if w, err = item.I.WriteJSONOpt(newTypeNames, short, w); err != nil {
		return w, err
	}
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"j":`...)
	if w, err = item.J.WriteJSONOpt(newTypeNames, short, w); err != nil {
		return w, err
	}
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"k":`...)
	if w, err = item.K.WriteJSONOpt(newTypeNames, short, w); err != nil {
		return w, err
	}
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"l":`...)
	if w, err = item.L.WriteJSONOpt(newTypeNames, short, w, item.N); err != nil {
		return w, err
	}
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"m":`...)
	if w, err = item.M.WriteJSONOpt(newTypeNames, short, w); err != nil {
		return w, err
	}
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"o":`...)
	if w, err = item.O.WriteJSONOpt(newTypeNames, short, w, item.N); err != nil {
		return w, err
	}
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"p":`...)
	if w, err = item.P.WriteJSONOpt(newTypeNames, short, w, item.N); err != nil {
		return w, err
	}
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"q":`...)
	if w, err = item.Q.WriteJSONOpt(newTypeNames, short, w); err != nil {
		return w, err
	}
	return append(w, '}'), nil
}

func (item *Replace) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *Replace) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("replace", err.Error())
	}
	return nil
}
