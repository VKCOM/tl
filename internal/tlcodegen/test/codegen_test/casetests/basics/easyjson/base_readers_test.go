// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlgen_easyjson

import (
	"testing"

	"github.com/mailru/easyjson/jlexer"
	"github.com/stretchr/testify/assert"
)

func TestEasyJsonReadInt32(t *testing.T) {
	type subTest struct {
		input      string
		definition string
	}
	subtests := []subTest{
		{
			input:      "12345",
			definition: "normal number",
		},
		{
			input:      `"12345"`,
			definition: "number as string",
		},
	}
	for _, test := range subtests {
		t.Run(test.definition, func(st *testing.T) {
			l := jlexer.Lexer{Data: []byte(test.input)}
			var dst int32
			err := EasyJsonReadInt32(&l, &dst)
			if err != nil || dst != 12345 {
				st.Fail()
			}

		})
	}
}

func TestEasyJsonReadString(t *testing.T) {
	type subTest struct {
		input      string
		definition string
		success    bool
	}
	subtests := []subTest{
		{
			input:      `"12345"`,
			definition: "basic string",
			success:    true,
		},
		{
			input:      `{"base64": "MTIzNDU="}`,
			definition: "string as base64 object",
			success:    true,
		},
		{
			input:      `{"base64": "12345"}`,
			definition: "string as base64 object, input is incorrect",
			success:    false,
		},
		{
			input:      `{}`,
			definition: "string as base64 object, but no base64 value",
			success:    false,
		},
	}
	for _, test := range subtests {
		t.Run(test.definition, func(st *testing.T) {
			l := jlexer.Lexer{Data: []byte(test.input)}
			var dst string
			err := EasyJsonReadString(&l, &dst)
			if test.success {
				if err != nil || dst != "12345" {
					st.Fail()
				}
			} else {
				if err == nil {
					st.Fail()
				}
			}
		})
	}
}

func TestEasyJsonReadUnion(t *testing.T) {
	type subTest struct {
		input      string
		definition string
		success    bool
		typeValue  string
		valueSlice []byte
	}
	subtests := []subTest{
		{
			definition: "type and value presented, expected order",
			input:      `{"type": "#123456", "value":"something"}`,
			success:    true,
			typeValue:  "#123456",
			valueSlice: []byte(`"something"`),
		},
		{
			definition: "type and value presented, expected order",
			input:      `{"type": "#123456", "value":"something"]`,
			success:    false,
		},
		{
			definition: "type and value presented, reversed order",
			input:      `{"value":"something", "type": "#123456"}`,
			success:    true,
			typeValue:  "#123456",
			valueSlice: []byte(`"something"`),
		},
		{
			definition: "only type presented",
			input:      `{"type": "#123456"}`,
			success:    true,
			typeValue:  "#123456",
			valueSlice: nil,
		},
		{
			definition: "only value presented",
			input:      `{"value":"something"}`,
			success:    false,
		},
		{
			definition: "nothing presented",
			input:      `{}`,
			success:    false,
		},
	}
	for _, test := range subtests {
		t.Run(test.definition, func(st *testing.T) {
			l := jlexer.Lexer{Data: []byte(test.input)}
			typeValue, valueSlice, err := EasyJsonReadUnion(&l)
			if test.success {
				if err != nil {
					st.Fail()
				}
				assert.Equal(st, test.typeValue, typeValue)
				assert.Equal(st, string(test.valueSlice), string(valueSlice))
			} else {
				if err == nil {
					st.Fail()
				}
			}
		})
	}
}

func TestEasyJsonReadMaybe(t *testing.T) {
	type subTest struct {
		input      string
		definition string
		success    bool
		okValue    bool
		valueSlice []byte
	}
	subtests := []subTest{
		{
			definition: "ok and value presented, expected order",
			input:      `{"ok": true, "value":"something"}`,
			success:    true,
			okValue:    true,
			valueSlice: []byte(`"something"`),
		},
		{
			definition: "ok and value presented, expected order, but ok is false",
			input:      `{"ok": false, "value":"something"}`,
			success:    false,
		},
		{
			definition: "ok and value presented, reversed order",
			input:      `{"value":"something", "ok": true}`,
			success:    true,
			okValue:    true,
			valueSlice: []byte(`"something"`),
		},
		{
			definition: "only ok presented, ok - true",
			input:      `{"ok": true}`,
			success:    true,
			okValue:    true,
			valueSlice: nil,
		},
		{
			definition: "only ok presented, ok - false",
			input:      `{"ok": false}`,
			success:    true,
			okValue:    false,
			valueSlice: nil,
		},
		{
			definition: "only value presented",
			input:      `{"value":"something"}`,
			okValue:    true,
			success:    true,
			valueSlice: []byte(`"something"`),
		},
		{
			definition: "nothing presented",
			input:      `{}`,
			okValue:    false,
			success:    true,
		},
	}
	for _, test := range subtests {
		t.Run(test.definition, func(st *testing.T) {
			l := jlexer.Lexer{Data: []byte(test.input)}
			okValue, valueSlice, err := EasyJsonReadMaybe(&l)
			if test.success {
				if err != nil {
					st.Fail()
				}
				assert.Equal(st, test.okValue, okValue)
				assert.Equal(st, string(test.valueSlice), string(valueSlice))
			} else {
				if err == nil {
					st.Fail()
				}
			}
		})
	}
}
