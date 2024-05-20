// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package internal

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/mailru/easyjson/jlexer"
)

type UnionElement struct {
	TLTag    uint32
	TLName   string
	TLString string
}

func ErrorClientWrite(typeName string, err error) error {
	return fmt.Errorf("failed to serialize %s request: %w", typeName, err)
}

func ErrorClientDo(typeName string, network string, actorID int64, address string, err error) error {
	return fmt.Errorf("%s request to %s://%d@%s failed: %w", typeName, network, actorID, address, err)
}

func ErrorClientReadResult(typeName string, network string, actorID int64, address string, err error) error {
	return fmt.Errorf("failed to deserialize %s response from %s://%d@%s: %w", typeName, network, actorID, address, err)
}

func ErrorServerHandle(typeName string, err error) error {
	return fmt.Errorf("failed to handle %s: %w", typeName, err)
}

func ErrorServerRead(typeName string, err error) error {
	return fmt.Errorf("failed to deserialize %s request: %w", typeName, err)
}

func ErrorServerWriteResult(typeName string, err error) error {
	return fmt.Errorf("failed to serialize %s response: %w", typeName, err)
}

func ErrorInvalidEnumTag(typeName string, tag uint32) error {
	return fmt.Errorf("invalid enum %q tag: 0x%x", typeName, tag)
}

func ErrorInvalidUnionTag(typeName string, tag uint32) error {
	return fmt.Errorf("invalid union %q tag: 0x%x", typeName, tag)
}

func ErrorWrongSequenceLength(typeName string, actual int, expected uint32) error {
	return fmt.Errorf("wrong sequence %q length: %d expected: %d", typeName, actual, expected)
}

func ErrorInvalidEnumTagJSON(typeName string, tag string) error {
	return fmt.Errorf("invalid enum %q tag: %q", typeName, tag)
}

func ErrorInvalidUnionTagJSON(typeName string, tag string) error {
	return fmt.Errorf("invalid union %q tag: %q", typeName, tag)
}

func ErrorInvalidUnionLegacyTagJSON(typeName string, tag string) error {
	return fmt.Errorf("legacy union %q tag %q, please remove suffix", typeName, tag)
}

func ErrorInvalidJSON(typeName string, msg string) error {
	return fmt.Errorf("invalid json for type %q - %s", typeName, msg)
}

func ErrorInvalidJSONWithDuplicatingKeys(typeName string, field string) error {
	return fmt.Errorf("invalid json for type %q: %q repeats several times", typeName, field)
}

func ErrorInvalidJSONExcessElement(typeName string, key string) error {
	return fmt.Errorf("invalid json object key %q", key)
}

func JsonReadUnionType(typeName string, j interface{}) (map[string]interface{}, string, error) {
	if j == nil {
		return nil, "", ErrorInvalidJSON(typeName, "expected json object")
	}
	jm, ok := j.(map[string]interface{})
	if !ok {
		return nil, "", ErrorInvalidJSON(typeName, "expected json object")
	}
	jtype, ok := jm["type"]
	if !ok {
		return nil, "", ErrorInvalidJSON(typeName, "expected 'type' key")
	}
	var ret string
	if err := JsonReadString(jtype, &ret); err != nil {
		return nil, "", err
	}
	delete(jm, "type")
	return jm, ret, nil
}

func Json2ReadUnion(typeName string, in *jlexer.Lexer) (string, []byte, error) {
	if in == nil {
		return "", nil, ErrorInvalidJSON(typeName, "expected json object")
	}
	var valueFound bool
	var valueSlice []byte

	var typeFound bool
	var typeValue string

	in.Delim('{')
	if !in.Ok() {
		return "", nil, in.Error()
	}
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(true)
		in.WantColon()
		switch key {
		case "value":
			if valueFound {
				return "", nil, ErrorInvalidJSONWithDuplicatingKeys(typeName, "value")
			}
			valueSlice = in.Raw()
			valueFound = true
		case "type":
			if typeFound {
				return "", nil, ErrorInvalidJSONWithDuplicatingKeys(typeName, "type")
			}
			typeValue = in.UnsafeString()
			typeFound = true
		default:
			return "", nil, ErrorInvalidJSON(typeName, "unexpected field \""+key+"\" in union")
		}

		in.WantComma()
	}
	in.Delim('}')
	if !in.Ok() {
		return "", nil, in.Error()
	}

	if !typeFound {
		return "", nil, ErrorInvalidJSON(typeName, "type is absent")
	}

	return typeValue, valueSlice, nil
}

func JsonReadMaybe(typeName string, j interface{}) (bool, interface{}, error) {
	if j == nil {
		return false, nil, nil
	}
	jm, ok := j.(map[string]interface{})
	if !ok {
		return false, nil, ErrorInvalidJSON(typeName, "expected json object")
	}
	jvalue := jm["value"]
	delete(jm, "value")
	jok, ok := jm["ok"]
	delete(jm, "ok")
	var dst bool
	if !ok {
		if jvalue != nil {
			dst = true
		}
	} else {
		if err := JsonReadBool(jok, &dst); err != nil {
			return false, nil, err
		}
		if !dst && jvalue != nil {
			return false, nil, ErrorInvalidJSON(typeName, "if 'ok' is set to false, 'value' should be omitted")
		}
	}
	for k := range jm {
		return false, nil, ErrorInvalidJSONExcessElement(typeName, k)
	}
	return dst, jvalue, nil
}

func Json2ReadMaybe(typeName string, in *jlexer.Lexer) (bool, []byte, error) {
	if in == nil {
		return false, nil, nil
	}
	var valueFound bool
	var valueSlice []byte

	var okFound bool
	var okValue bool

	in.Delim('{')
	if !in.Ok() {
		return false, nil, ErrorInvalidJSON(typeName, "expected json object")
	}
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(true)
		in.WantColon()
		switch key {
		case "value":
			if valueFound {
				return false, nil, ErrorInvalidJSONWithDuplicatingKeys(typeName, "value")
			}
			valueSlice = in.Raw()
			valueFound = true
		case "ok":
			if okFound {
				return false, nil, ErrorInvalidJSONWithDuplicatingKeys(typeName, "ok")
			}
			okValue = in.Bool()
			okFound = true
		default:
			return false, nil, ErrorInvalidJSON(typeName, "unexpected field \""+key+"\" in maybe")
		}

		in.WantComma()
	}
	in.Delim('}')
	if !in.Ok() {
		return false, nil, in.Error()
	}

	if okFound && !okValue && valueSlice != nil {
		return false, nil, ErrorInvalidJSON(typeName, "ok is false but value is presented in maybe")
	}
	if !okFound && valueSlice != nil {
		okValue = true
	}
	return okValue, valueSlice, nil
}

func JsonReadArray(typeName string, j interface{}) (int, []interface{}, error) {
	var arr []interface{}
	var arrok bool
	if j != nil {
		arr, arrok = j.([]interface{})
		if !arrok {
			return 0, nil, ErrorInvalidJSON(typeName, "expected json array")
		}
	}
	return len(arr), arr, nil
}

func JsonReadArrayFixedSize(typeName string, j interface{}, expectLength uint32) (int, []interface{}, error) {
	l, arr, err := JsonReadArray(typeName, j)
	if err == nil && l != int(expectLength) {
		return 0, nil, ErrorWrongSequenceLength(typeName, l, expectLength)
	}
	return l, arr, err
}

func JsonReadBool(j interface{}, dst *bool) error {
	if j == nil {
		*dst = false
		return nil
	}
	jj, ok := j.(bool)
	if !ok {
		return fmt.Errorf("invalid json for bool")
	}
	*dst = jj
	return nil
}

func Json2ReadBool(in *jlexer.Lexer, dst *bool) error {
	if in == nil {
		*dst = false
		return nil
	}
	*dst = in.Bool()
	if !in.Ok() {
		return in.Error()
	}
	return nil
}

func JsonReadString(j interface{}, dst *string) error {
	if j == nil {
		*dst = ""
		return nil
	}
	switch jj := j.(type) {
	case string:
		*dst = jj
		return nil
	case map[string]interface{}:
		iface, ok := jj["base64"]
		if !ok {
			return fmt.Errorf("invalid json for string: base64 encoded didn't match as string")
		}
		str, ok := iface.(string)
		if !ok {
			return fmt.Errorf("invalid json for string: unexpected binary string's object")
		}
		buf, err := base64.StdEncoding.DecodeString(str)
		*dst = string(buf)
		return err
	default:
		return fmt.Errorf("invalid json for string")
	}
}

func JsonReadStringBytes(j interface{}, dst *[]byte) error {
	if j == nil {
		*dst = nil
		return nil
	}
	switch jj := j.(type) {
	case string:
		*dst = append((*dst)[:0], jj...)
		return nil
	case map[string]interface{}:
		iface, ok := jj["base64"]
		if !ok {
			return fmt.Errorf("invalid json for string: base64 encoded didn't match as string")
		}
		str, ok := iface.(string)
		if !ok {
			return fmt.Errorf("invalid json for string: unexpected binary string's object")
		}
		buf, err := base64.StdEncoding.DecodeString(str)
		*dst = buf
		return err
	default:
		return fmt.Errorf("invalid json for string")
	}
}

func Json2ReadString(in *jlexer.Lexer, dst *string) error {
	if in == nil {
		*dst = ""
		return nil
	}

	switch in.CurrentToken() {
	case jlexer.TokenString:
		*dst = in.String()
	case jlexer.TokenDelim:
		var findValue = false

		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "base64":
				if findValue {
					return fmt.Errorf("base64 repeats several times")
				}
				*dst = string(in.Bytes())
				findValue = true
			default:
				return fmt.Errorf("unexpected field \"" + key + "\"")
			}

			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}

		if !findValue {
			return fmt.Errorf("base64 is absent")
		}
	default:
		return fmt.Errorf("invalid json for string")
	}
	return nil
}

func Json2ReadStringBytes(in *jlexer.Lexer, dst *[]byte) error {
	if in == nil {
		*dst = nil
		return nil
	}

	switch in.CurrentToken() {
	case jlexer.TokenString:
		*dst = append((*dst)[:0], in.String()...)
	case jlexer.TokenDelim:
		var findValue = false

		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "base64":
				if findValue {
					return fmt.Errorf("base64 repeats several times")
				}
				*dst = in.Bytes()
				findValue = true
			default:
				return fmt.Errorf("unexpected field \"" + key + "\"")
			}

			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}

		if !findValue {
			return fmt.Errorf("base64 is absent")
		}
	default:
		return fmt.Errorf("invalid json for string")
	}
	return nil
}

// We allow to specify numbers as "123", so that JS can pass through int64 and bigger numbers
func jsonNumberOrString(j interface{}) (string, bool) {
	jn, ok := j.(json.Number)
	if ok {
		return string(jn), ok
	}
	js, ok := j.(string)
	return js, ok
}

func JsonReadUint32(j interface{}, dst *uint32) error {
	if j == nil {
		*dst = 0
		return nil
	}
	jj, ok := jsonNumberOrString(j)
	if !ok {
		return fmt.Errorf("invalid json for uint32")
	}
	val, err := strconv.ParseUint(jj, 10, 32)
	if err != nil {
		return fmt.Errorf("invalid number format for uint32 %w", err)
	}
	*dst = uint32(val)
	return nil
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

func JsonReadInt32(j interface{}, dst *int32) error {
	if j == nil {
		*dst = 0
		return nil
	}
	jj, ok := jsonNumberOrString(j)
	if !ok {
		return fmt.Errorf("invalid json for int32")
	}
	val, err := strconv.ParseInt(jj, 10, 32)
	if err != nil {
		return fmt.Errorf("invalid number format for int32 %w", err)
	}
	*dst = int32(val)
	return nil
}

func Json2ReadInt32(in *jlexer.Lexer, dst *int32) error {
	if in == nil {
		*dst = 0
		return nil
	}

	switch in.CurrentToken() {
	case jlexer.TokenString:
		src := in.UnsafeString()
		value, err := strconv.ParseInt(src, 10, 32)
		if err != nil {
			return err
		}
		*dst = int32(value)
	case jlexer.TokenNumber:
		*dst = in.Int32()
	default:
		return fmt.Errorf("invalid json for int32")
	}
	if !in.Ok() {
		return in.Error()
	}
	return nil
}

func JsonReadInt64(j interface{}, dst *int64) error {
	if j == nil {
		*dst = 0
		return nil
	}
	jj, ok := jsonNumberOrString(j)
	if !ok {
		return fmt.Errorf("invalid json for int64")
	}
	val, err := strconv.ParseInt(jj, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid number format for int64 %w", err)
	}
	*dst = val
	return nil
}

func Json2ReadInt64(in *jlexer.Lexer, dst *int64) error {
	if in == nil {
		*dst = 0
		return nil
	}

	switch in.CurrentToken() {
	case jlexer.TokenString:
		src := in.UnsafeString()
		value, err := strconv.ParseInt(src, 10, 64)
		if err != nil {
			return err
		}
		*dst = value
	case jlexer.TokenNumber:
		*dst = in.Int64()
	default:
		return fmt.Errorf("invalid json for int64")
	}
	if !in.Ok() {
		return in.Error()
	}
	return nil
}

func JsonReadFloat32(j interface{}, dst *float32) error {
	if j == nil {
		*dst = 0
		return nil
	}
	jj, ok := jsonNumberOrString(j)
	if !ok {
		return fmt.Errorf("invalid json for float32")
	}
	val, err := strconv.ParseFloat(jj, 32)
	if err != nil {
		return fmt.Errorf("invalid number format for float32 %w", err)
	}
	*dst = float32(val)
	return nil
}

func Json2ReadFloat32(in *jlexer.Lexer, dst *float32) error {
	if in == nil {
		*dst = 0
		return nil
	}

	switch in.CurrentToken() {
	case jlexer.TokenString:
		src := in.UnsafeString()
		value, err := strconv.ParseFloat(src, 32)
		if err != nil {
			return err
		}
		*dst = float32(value)
	case jlexer.TokenNumber:
		*dst = in.Float32()
	default:
		return fmt.Errorf("invalid json for float32")
	}
	if !in.Ok() {
		return in.Error()
	}
	return nil
}

func JsonReadFloat64(j interface{}, dst *float64) error {
	if j == nil {
		*dst = 0
		return nil
	}
	jj, ok := jsonNumberOrString(j)
	if !ok {
		return fmt.Errorf("invalid json for float64")
	}
	val, err := strconv.ParseFloat(jj, 64)
	if err != nil {
		return fmt.Errorf("invalid number format for float64 %w", err)
	}
	*dst = val
	return nil
}

func Json2ReadFloat64(in *jlexer.Lexer, dst *float64) error {
	if in == nil {
		*dst = 0
		return nil
	}

	switch in.CurrentToken() {
	case jlexer.TokenString:
		src := in.UnsafeString()
		value, err := strconv.ParseFloat(src, 64)
		if err != nil {
			return err
		}
		*dst = value
	case jlexer.TokenNumber:
		*dst = in.Float64()
	default:
		return fmt.Errorf("invalid json for float64")
	}
	if !in.Ok() {
		return in.Error()
	}
	return nil
}

func JsonBytesToInterface(b []byte) (interface{}, error) {
	var j interface{}
	d := json.NewDecoder(bytes.NewBuffer(b))
	d.UseNumber()
	if err := d.Decode(&j); err != nil {
		return j, err
	}
	return j, nil
}
