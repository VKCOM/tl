// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlgen_easyjson

import (
	"fmt"
	"strconv"

	"github.com/mailru/easyjson/jlexer"
)

/* similar for all numbers */
func EasyJsonReadInt32(in *jlexer.Lexer, dst *int32) error {
	in.FetchToken()
	if !in.Ok() {
		return fmt.Errorf("TODO")
	}

	switch in.CurrentToken() {
	case jlexer.TokenString:
		src := in.UnsafeString()
		if !in.Ok() {
			return fmt.Errorf("TODO")
		}
		value, err := strconv.ParseInt(src, 10, 32)
		if err != nil {
			return fmt.Errorf("TODO")
		}
		*dst = int32(value)
	case jlexer.TokenNumber:
		*dst = in.Int32()
		if !in.Ok() {
			return fmt.Errorf("TODO")
		}
	default:
		return fmt.Errorf("unsupported json object to parse int32")
	}
	return nil
}

func EasyJsonReadUnion(in *jlexer.Lexer) (string, []byte, error) {
	var valueSlice []byte

	var typeFound bool
	var typeValue string

	in.Delim('{')
	if !in.Ok() {
		return "", nil, fmt.Errorf("TODO")
	}
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(true)
		in.WantColon()
		switch key {
		case "value":
			valueSlice = in.Raw()
		case "type":
			typeValue = in.String()
			typeFound = true
		default:
			return "", nil, fmt.Errorf("TODO")
		}

		in.WantComma()
	}
	in.Delim('}')

	if !typeFound {
		return "", nil, fmt.Errorf("TODO")
	}

	return typeValue, valueSlice, nil
}

func EasyJsonReadMaybe(in *jlexer.Lexer) (bool, []byte, error) {
	var valueSlice []byte

	var okFound bool
	var okValue bool

	in.Delim('{')
	if !in.Ok() {
		return false, nil, fmt.Errorf("TODO")
	}
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(true)
		in.WantColon()
		switch key {
		case "value":
			valueSlice = in.Raw()
		case "ok":
			okFound = true
			okValue = in.Bool()
			if !in.Ok() {
				return false, nil, fmt.Errorf("TODO")
			}
		default:
			return false, nil, fmt.Errorf("TODO")
		}

		in.WantComma()
	}
	in.Delim('}')

	if okFound && !okValue && valueSlice != nil {
		return false, nil, fmt.Errorf("TODO")
	}
	if !okFound && valueSlice != nil {
		okValue = true
	}
	return okValue, valueSlice, nil
}

func EasyJsonReadString(in *jlexer.Lexer, dst *string) error {
	in.FetchToken()
	if !in.Ok() {
		return fmt.Errorf("TODO")
	}

	switch in.CurrentToken() {
	case jlexer.TokenString:
		*dst = in.String()
		if !in.Ok() {
			return fmt.Errorf("TODO")
		}
	case jlexer.TokenDelim:
		var findValue = false

		in.Delim('{')
		if !in.Ok() {
			return fmt.Errorf("TODO")
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "base64":
				findValue = true
				*dst = string(in.Bytes())
				if !in.Ok() {
					return in.Error()
				}
			default:
				return fmt.Errorf("TODO")
			}

			in.WantComma()
		}
		in.Delim('}')

		if !findValue {
			return fmt.Errorf("TODO")
		}

		if !in.Ok() {
			return fmt.Errorf("TODO")
		}
	default:
		return fmt.Errorf("unsupported json object to parse string")
	}
	return nil
}

func EasyJsonReadStringBytes(in *jlexer.Lexer, dst *[]byte) error {
	in.FetchToken()
	if !in.Ok() {
		return fmt.Errorf("TODO")
	}

	switch in.CurrentToken() {
	case jlexer.TokenString:
		*dst = in.Bytes()
		if !in.Ok() {
			return fmt.Errorf("TODO")
		}
	case jlexer.TokenDelim:
		var findValue = false

		in.Delim('{')
		if !in.Ok() {
			return fmt.Errorf("TODO")
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "base64":
				findValue = true
				*dst = in.Bytes()
				if !in.Ok() {
					return in.Error()
				}
			default:
				return fmt.Errorf("TODO")
			}

			in.WantComma()
		}
		in.Delim('}')

		if !findValue {
			return fmt.Errorf("TODO")
		}

		if !in.Ok() {
			return fmt.Errorf("TODO")
		}
	default:
		return fmt.Errorf("unsupported json object to parse string")
	}
	return nil
}
