package speed

import (
	"fmt"
	"io"
	"strconv"

	"github.com/VKCOM/tl/pkg/basictl"
	"github.com/mailru/easyjson/jlexer"
)

func JSONAppendPoint(w []byte, item *point, writeExcessField bool) []byte {
	w = append(w, '{')
	if item.x != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"one":`...)
		w = basictl.JSONWriteUint32(w, item.x)
	}
	if item.y != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"two":`...)
		w = basictl.JSONWriteUint32(w, item.y)
	}
	if item.z != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"three":`...)
		w = basictl.JSONWriteUint32(w, item.z)
	}
	if writeExcessField {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"four":`...)
		w = basictl.JSONWriteUint32(w, 1)
	}
	return append(w, '}')
}

func Json2ReadUint32(in *jlexer.Lexer, dst *uint32) error {
	if in == nil {
		*dst = 0
		return nil
	}

	switch in.CurrentToken() {
	case jlexer.TokenString:
		src := in.UnsafeString()
		value, err := strconv.ParseUint(src, 10, 32)
		if err != nil {
			return err
		}
		*dst = uint32(value)
	case jlexer.TokenNumber:
		*dst = in.Uint32()
	default:
		return fmt.Errorf("invalid json for uint32")
	}
	if !in.Ok() {
		return in.Error()
	}
	return nil
}

func JSONReadPoint(item *point, in *basictl.JsonLexer) error {
	var propXPresented bool
	var propYPresented bool
	var propZPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "one":
				if propXPresented {
					return io.EOF
				}
				if err := Json2ReadUint32(in, &item.x); err != nil {
					return err
				}
				propXPresented = true
			case "two":
				if propYPresented {
					return io.EOF
				}
				if err := Json2ReadUint32(in, &item.y); err != nil {
					return err
				}
				propYPresented = true
			case "three":
				if propZPresented {
					return io.EOF
				}
				if err := Json2ReadUint32(in, &item.z); err != nil {
					return err
				}
				propZPresented = true
			default:
				return io.EOF
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propXPresented {
		item.x = 0
	}
	if !propYPresented {
		item.y = 0
	}
	if !propZPresented {
		item.z = 0
	}
	return nil
}
