// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlast

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCRC32TL(t *testing.T) {
	t.Run("Primitives", func(t *testing.T) {
		test := []struct {
			input  string
			output uint32
		}{
			{input: "int#a8509bda ? = Int;", output: 0xa8509bda},
			{input: "long#22076cba ? = Long;", output: 0x22076cba},
			{input: "float#824dab22 ? = Float;", output: 0x824dab22},
			{input: "double#2210c154 ? = Double;", output: 0x2210c154},
			{input: "string#b5286e24 ? = String;", output: 0xb5286e24},
		}
		for _, tst := range test {
			tl, err := ParseTL(tst.input)
			require.NoError(t, err)
			require.Equal(t, tst.output, tl[0].Crc32())
		}
	})
	tests := []struct {
		tlText string
		tlTag  uint32
	}{
		{`---functions--- @any  get_arrays n:# a:n*[int] b:5*[int] = Tuple int 5;`, 0x90658cdb},
		{`---functions--- @any  get_arrays#12345678 n:# a:n*[int] b:5*[int] = Tuple int 5;`, 0x12345678},
	}
	for _, test := range tests {
		tl, err := ParseTL(test.tlText)
		require.NoError(t, err)
		require.Equal(t, test.tlTag, tl[0].Crc32())
	}
}
